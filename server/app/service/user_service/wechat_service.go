package user_service

import (
	"context"
	"encoding/json"
	"errors"
	"x_admin/app/model/user_model"
	"x_admin/app/schema/user_schema"
	"x_admin/config"
	"x_admin/core"
	"x_admin/core/response"
	"x_admin/plugin"
	"x_admin/util"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var WechatService = NewWechatService()

// NewWechatService 初始化
func NewWechatService() *wechatService {
	db := core.GetDB()
	return &wechatService{db: db}
}

type wechatService struct {
	db *gorm.DB
}

// ---- 微信小程序 ----

// MiniLogin 小程序登录（code → openid → 查找/自动注册用户 → JWT）
func (s *wechatService) MiniLogin(c *gin.Context, req *user_schema.WechatMiniLoginReq) (user_schema.LoginResp, error) {
	app := plugin.GetMiniProgramApp()
	if app == nil {
		return user_schema.LoginResp{}, response.Failed.SetMessage("小程序未配置")
	}

	// 调用微信 jscode2session 获取 openid
	result, err := app.Auth.Session(context.Background(), req.Code)
	if err != nil {
		core.Logger.Errorf("MiniLogin Session err: %v", err)
		return user_schema.LoginResp{}, response.Failed.SetMessage("微信登录失败")
	}
	if result.OpenID == "" {
		return user_schema.LoginResp{}, response.Failed.SetMessage("微信登录失败：未获取到openid")
	}

	return s.wechatLogin(c, user_model.IdentityWechatMini, result.OpenID, result.SessionKey, "")
}

// BindMini 绑定小程序到当前用户
func (s *wechatService) BindMini(userID string, req *user_schema.WechatBindReq) error {
	app := plugin.GetMiniProgramApp()
	if app == nil {
		return response.Failed.SetMessage("小程序未配置")
	}

	result, err := app.Auth.Session(context.Background(), req.Code)
	if err != nil {
		core.Logger.Errorf("BindMini Session err: %v", err)
		return response.Failed.SetMessage("微信验证失败")
	}
	if result.OpenID == "" {
		return response.Failed.SetMessage("微信验证失败：未获取到openid")
	}

	return s.bindWechat(userID, user_model.IdentityWechatMini, result.OpenID, result.SessionKey, "")
}

// ---- 公众号 ----

// MpLogin 公众号登录（OAuth code → openid → 查找/自动注册用户 → JWT）
func (s *wechatService) MpLogin(c *gin.Context, req *user_schema.WechatMpLoginReq) (user_schema.LoginResp, error) {
	app := plugin.GetOfficialAccountApp()
	if app == nil {
		return user_schema.LoginResp{}, response.Failed.SetMessage("公众号未配置")
	}

	// 通过 OAuth code 换取 token + openid
	tokenResp, err := app.OAuth.TokenFromCode(req.Code)
	if err != nil {
		core.Logger.Errorf("MpLogin TokenFromCode err: %v", err)
		return user_schema.LoginResp{}, response.Failed.SetMessage("微信登录失败")
	}

	openID, ok := (*tokenResp)["openid"].(string)
	if !ok || openID == "" {
		return user_schema.LoginResp{}, response.Failed.SetMessage("微信登录失败：未获取到openid")
	}

	accessToken, _ := (*tokenResp)["access_token"].(string)
	unionID, _ := (*tokenResp)["unionid"].(string)

	// 尝试获取用户昵称头像（snsapi_userinfo 授权才能拿到）
	extra := ""
	if accessToken != "" {
		userInfo, err := app.OAuth.UserFromToken(accessToken, openID)
		if err == nil && userInfo != nil {
			extraBytes, _ := json.Marshal(userInfo)
			extra = string(extraBytes)
		}
	}

	_ = unionID // 预留 unionID 后续使用
	return s.wechatLogin(c, user_model.IdentityWechatMp, openID, accessToken, extra)
}

// BindMp 绑定公众号到当前用户
func (s *wechatService) BindMp(userID string, req *user_schema.WechatBindReq) error {
	app := plugin.GetOfficialAccountApp()
	if app == nil {
		return response.Failed.SetMessage("公众号未配置")
	}

	tokenResp, err := app.OAuth.TokenFromCode(req.Code)
	if err != nil {
		core.Logger.Errorf("BindMp TokenFromCode err: %v", err)
		return response.Failed.SetMessage("微信验证失败")
	}

	openID, ok := (*tokenResp)["openid"].(string)
	if !ok || openID == "" {
		return response.Failed.SetMessage("微信验证失败：未获取到openid")
	}

	accessToken, _ := (*tokenResp)["access_token"].(string)
	extra := ""
	if accessToken != "" {
		userInfo, err := app.OAuth.UserFromToken(accessToken, openID)
		if err == nil && userInfo != nil {
			extraBytes, _ := json.Marshal(userInfo)
			extra = string(extraBytes)
		}
	}

	return s.bindWechat(userID, user_model.IdentityWechatMp, openID, accessToken, extra)
}

// ---- 通用解绑 ----

// UnbindWechat 解绑微信（小程序/公众号）
func (s *wechatService) UnbindWechat(userID string, req *user_schema.WechatUnbindReq) error {
	var auth user_model.UserAuth
	err := s.db.Where("user_id = ? AND identity_type = ?", userID, req.IdentityType).First(&auth).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.Failed.SetMessage("未绑定该微信账号")
		}
		return response.CheckErr(err, "查询绑定记录失败")
	}

	// 软删除绑定记录
	if err := s.db.Delete(&auth).Error; err != nil {
		return response.CheckErr(err, "解绑失败")
	}
	return nil
}

// ---- 内部方法 ----

// wechatLogin 微信通用登录逻辑：通过 openid 查找用户，不存在则自动注册
func (s *wechatService) wechatLogin(c *gin.Context, identityType, openID, credential, extra string) (user_schema.LoginResp, error) {
	// 查找绑定记录
	var auth user_model.UserAuth
	err := s.db.Where("identity_type = ? AND identifier = ?", identityType, openID).First(&auth).Error

	isNew := false
	var user user_model.User

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 未绑定 → 自动注册用户
		isNew = true
		nickname := "微信用户" + util.ToolsUtil.RandomString(6)

		// 从 extra 中尝试提取昵称
		if extra != "" {
			var extraMap map[string]any
			if json.Unmarshal([]byte(extra), &extraMap) == nil {
				if n, ok := extraMap["nickname"].(string); ok && n != "" {
					nickname = n
				}
			}
		}

		// 用 openid 生成唯一邮箱占位（避免 email unique 冲突）
		email := openID + "@" + identityType

		err = s.db.Transaction(func(tx *gorm.DB) error {
			// 创建用户
			newUser := user_model.User{
				Email:    email,
				Nickname: nickname,
				Salt:     util.ToolsUtil.RandomString(16),
			}
			if err := tx.Create(&newUser).Error; err != nil {
				return err
			}
			user = newUser

			// 创建绑定记录
			bindAuth := user_model.UserAuth{
				UserID:       newUser.ID,
				IdentityType: identityType,
				Identifier:   openID,
				Credential:   credential,
				Extra:        extra,
			}
			return tx.Create(&bindAuth).Error
		})
		if err != nil {
			return user_schema.LoginResp{}, response.CheckErr(err, "自动注册失败")
		}
	} else if err != nil {
		return user_schema.LoginResp{}, response.CheckErr(err, "查询绑定记录失败")
	} else {
		// 已绑定 → 查找用户
		if err := s.db.Where("id = ?", auth.UserID).First(&user).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return user_schema.LoginResp{}, response.Failed.SetMessage("关联用户不存在")
			}
			return user_schema.LoginResp{}, response.CheckErr(err, "查询用户失败")
		}

		// 更新凭证（session_key 可能会变）
		s.db.Model(&auth).Updates(map[string]any{
			"credential": credential,
			"extra":      extra,
		})
	}

	// 校验状态
	if user.Status == 1 {
		return user_schema.LoginResp{}, response.LoginDisableError
	}

	// 生成 JWT
	accessToken, refreshToken, err := util.JWTUtil.GenerateTokenPair(user.ID, user.TokenVersion)
	if err != nil {
		return user_schema.LoginResp{}, response.CheckErr(err, "生成token失败")
	}

	// 更新最后登录信息
	clientIP := c.ClientIP()
	s.db.Model(&user).Updates(map[string]any{
		"last_login_ip":   clientIP,
		"last_login_time": util.NullTimeUtil.Now(),
	})

	return user_schema.LoginResp{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    config.JWTConfig.AccessExpireSec,
		IsNew:        isNew,
		UserID:       user.ID,
		Nickname:     user.Nickname,
	}, nil
}

// bindWechat 微信通用绑定逻辑
func (s *wechatService) bindWechat(userID, identityType, openID, credential, extra string) error {
	// 检查该 openid 是否已被绑定
	var existAuth user_model.UserAuth
	err := s.db.Where("identity_type = ? AND identifier = ?", identityType, openID).First(&existAuth).Error
	if err == nil {
		if existAuth.UserID == userID {
			return response.Failed.SetMessage("已绑定该微信账号")
		}
		return response.Failed.SetMessage("该微信账号已被其他用户绑定")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return response.CheckErr(err, "检查绑定失败")
	}

	// 检查当前用户是否已绑定该类型
	var existUserAuth user_model.UserAuth
	err = s.db.Where("user_id = ? AND identity_type = ?", userID, identityType).First(&existUserAuth).Error
	if err == nil {
		return response.Failed.SetMessage("已绑定该类型微信，请先解绑")
	}

	// 创建绑定记录
	auth := user_model.UserAuth{
		UserID:       userID,
		IdentityType: identityType,
		Identifier:   openID,
		Credential:   credential,
		Extra:        extra,
	}
	if err := s.db.Create(&auth).Error; err != nil {
		return response.CheckErr(err, "绑定失败")
	}
	return nil
}
