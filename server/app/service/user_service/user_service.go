package user_service

import (
	"errors"
	"strconv"
	"x_admin/app/model/user_model"
	"x_admin/app/schema/user_schema"
	"x_admin/config"
	"x_admin/core"
	"x_admin/core/response"
	"x_admin/util"
	"x_admin/util/convert_util"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var UserService = NewUserService()

// NewUserService 初始化
func NewUserService() *userService {
	db := core.GetDB()
	return &userService{db: db}
}

type userService struct {
	db *gorm.DB
}

// hashPassword bcrypt 加密密码
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// checkPassword bcrypt 校验密码
func checkPassword(hashedPassword, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}

// existsActiveUser 判断是否存在「未删除」的用户（按自定义条件）。
// 软删除（is_delete=1）的记录不计入，避免"已删除账号仍占用邮箱/手机号"，
// 导致用户被软删除后无法用同一邮箱/手机号重新注册。
func (s *userService) existsActiveUser(query string, args ...any) (bool, error) {
	var count int64
	if err := s.db.Model(&user_model.User{}).
		Where(query, args...).
		Where("is_delete = 0").
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

// tokenVersionCacheKey 生成 token_version 缓存 key
func tokenVersionCacheKey(userID string) string {
	return "user:tv:" + userID
}

// GetCachedTokenVersion 获取缓存的 token_version（Redis 优先，miss 回源 MySQL，TTL 5min）
func (s *userService) GetCachedTokenVersion(userID string) (int64, error) {
	cacheKey := tokenVersionCacheKey(userID)
	cached := util.RedisUtil.Get(cacheKey)
	if cached != "" {
		v, err := strconv.ParseInt(cached, 10, 64)
		if err == nil {
			return v, nil
		}
	}
	// 回源 MySQL
	version, err := s.GetUserTokenVersion(userID)
	if err != nil {
		return 0, err
	}
	// 写入缓存，5分钟
	util.RedisUtil.Set(cacheKey, strconv.FormatInt(version, 10), 300)
	return version, nil
}

// InvalidateTokenVersionCache 清除 token_version 缓存（踢人/重置密码后立即生效）
func (s *userService) InvalidateTokenVersionCache(userID string) {
	util.RedisUtil.Del(tokenVersionCacheKey(userID))
}

// Register 邮箱注册
func (s *userService) Register(req *user_schema.RegisterReq) error {
	// 校验验证码
	if err := util.EmailCodeUtil.VerifyCode(req.Email, util.CodeSceneRegister, req.Code, ""); err != nil {
		return response.Failed.SetMessage(err.Error())
	}

	// 检查邮箱是否已注册（只统计未删除的账号，软删除后可重新注册）
	exists, err := s.existsActiveUser("email = ?", req.Email)
	if err != nil {
		return response.CheckErr(err, "检查邮箱失败")
	}
	if exists {
		return response.Failed.SetMessage("该邮箱已被注册")
	}

	// 加密密码
	hashedPwd, err := hashPassword(req.Password)
	if err != nil {
		return response.CheckErr(err, "密码加密失败")
	}

	// 生成昵称
	nickname := req.Nickname
	if nickname == "" {
		nickname = "用户" + util.ToolsUtil.RandomString(6)
	}

	// 创建用户
	user := user_model.User{
		Email:    req.Email,
		Password: hashedPwd,
		Salt:     util.ToolsUtil.RandomString(16),
		Nickname: nickname,
	}
	if err := s.db.Create(&user).Error; err != nil {
		return response.CheckErr(err, "注册失败")
	}
	return nil
}

// Login 邮箱登录，返回 JWT token 对
func (s *userService) Login(c *gin.Context, req *user_schema.LoginReq) (user_schema.LoginResp, error) {
	var user user_model.User
	err := s.db.Where("email = ?", req.Email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user_schema.LoginResp{}, response.LoginAccountError
		}
		return user_schema.LoginResp{}, response.CheckErr(err, "查询用户失败")
	}

	return s.doLogin(c, &user, req.Password)
}

// PhoneLogin 手机号+密码登录
func (s *userService) PhoneLogin(c *gin.Context, req *user_schema.PhoneLoginReq) (user_schema.LoginResp, error) {
	var user user_model.User
	err := s.db.Where("phone = ? AND phone != ''", req.Phone).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user_schema.LoginResp{}, response.LoginAccountError
		}
		return user_schema.LoginResp{}, response.CheckErr(err, "查询用户失败")
	}

	return s.doLogin(c, &user, req.Password)
}

// PhoneCodeLogin 手机号+短信验证码登录（未注册手机号不自动注册，需先绑定手机）
func (s *userService) PhoneCodeLogin(c *gin.Context, req *user_schema.PhoneCodeLoginReq) (user_schema.LoginResp, error) {
	// 校验短信验证码
	if err := util.SmsCodeUtil.VerifyCode(req.Phone, util.SmsSceneLogin, req.Code); err != nil {
		return user_schema.LoginResp{}, response.Failed.SetMessage(err.Error())
	}

	var user user_model.User
	err := s.db.Where("phone = ? AND phone != ''", req.Phone).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user_schema.LoginResp{}, response.Failed.SetMessage("该手机号未绑定账号，请先注册并绑定手机号")
		}
		return user_schema.LoginResp{}, response.CheckErr(err, "查询用户失败")
	}

	if user.Status == 1 {
		return user_schema.LoginResp{}, response.LoginDisableError
	}

	// 生成 JWT token 对
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
	}, nil
}

// doLogin 统一登录逻辑（校验密码+状态+生成token+更新登录信息）
func (s *userService) doLogin(c *gin.Context, user *user_model.User, password string) (user_schema.LoginResp, error) {
	// 校验密码
	if !checkPassword(user.Password, password) {
		return user_schema.LoginResp{}, response.LoginAccountError
	}

	// 校验状态
	if user.Status == 1 {
		return user_schema.LoginResp{}, response.LoginDisableError
	}

	// 生成 JWT token 对（仅传 userID + tokenVersion）
	accessToken, refreshToken, err := util.JWTUtil.GenerateTokenPair(user.ID, user.TokenVersion)
	if err != nil {
		return user_schema.LoginResp{}, response.CheckErr(err, "生成token失败")
	}

	// 更新最后登录信息
	clientIP := c.ClientIP()
	s.db.Model(user).Updates(map[string]any{
		"last_login_ip":   clientIP,
		"last_login_time": util.NullTimeUtil.Now(),
	})

	return user_schema.LoginResp{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    config.JWTConfig.AccessExpireSec,
	}, nil
}

// RefreshToken 刷新 token
func (s *userService) RefreshToken(req *user_schema.RefreshTokenReq) (user_schema.LoginResp, error) {
	// 解析 refresh_token
	claims, err := util.JWTUtil.ParseRefreshToken(req.RefreshToken)
	if err != nil {
		return user_schema.LoginResp{}, response.TokenInvalid.SetMessage("refreshToken已失效")
	}

	// 查询用户并校验 token_version
	var user user_model.User
	if err := s.db.Where("id = ?", claims.UserID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user_schema.LoginResp{}, response.TokenInvalid.SetMessage("用户不存在")
		}
		return user_schema.LoginResp{}, response.CheckErr(err, "查询用户失败")
	}

	// token_version 校验（踢人下线检测）
	if claims.TokenVersion != user.TokenVersion {
		return user_schema.LoginResp{}, response.TokenInvalid.SetMessage("token已失效，请重新登录")
	}

	if user.Status == 1 {
		return user_schema.LoginResp{}, response.LoginDisableError
	}

	// 生成新 token 对
	accessToken, refreshToken, err := util.JWTUtil.GenerateTokenPair(user.ID, user.TokenVersion)
	if err != nil {
		return user_schema.LoginResp{}, response.CheckErr(err, "生成token失败")
	}

	return user_schema.LoginResp{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    config.JWTConfig.AccessExpireSec,
	}, nil
}

// GetUserInfo 获取用户信息
func (s *userService) GetUserInfo(userID string) (user_schema.UserInfoResp, error) {
	var user user_model.User
	if err := s.db.Where("id = ?", userID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user_schema.UserInfoResp{}, response.Failed.SetMessage("用户不存在")
		}
		return user_schema.UserInfoResp{}, response.CheckErr(err, "获取用户信息失败")
	}
	var resp user_schema.UserInfoResp
	convert_util.Copy(&resp, user)
	return resp, nil
}

// UpdateUserInfo 更新用户信息
func (s *userService) UpdateUserInfo(userID string, req *user_schema.UpdateUserReq) error {
	updates := map[string]any{}
	if req.Nickname != "" {
		updates["nickname"] = req.Nickname
	}
	if req.Avatar != "" {
		updates["avatar"] = req.Avatar
	}
	if len(updates) == 0 {
		return nil
	}
	if err := s.db.Model(&user_model.User{}).Where("id = ?", userID).Updates(updates).Error; err != nil {
		return response.CheckErr(err, "更新失败")
	}
	return nil
}

// SendEmailCode 发送邮箱验证码（注册/重置密码/解绑手机场景）
func (s *userService) SendEmailCode(req *user_schema.SendCodeReq) error {
	// 注册场景：检查邮箱未注册（只统计未删除的账号）
	if req.Scene == util.CodeSceneRegister {
		exists, err := s.existsActiveUser("email = ?", req.Email)
		if err != nil {
			return response.CheckErr(err, "检查邮箱失败")
		}
		if exists {
			return response.Failed.SetMessage("该邮箱已被注册")
		}
	}
	// 重置密码场景：检查邮箱已注册
	if req.Scene == util.CodeSceneReset {
		var existUser user_model.User
		err := s.db.Where("email = ?", req.Email).First(&existUser).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return response.Failed.SetMessage("该邮箱未注册")
			}
			return response.CheckErr(err, "检查邮箱失败")
		}
	}
	// 解绑手机场景：检查邮箱已注册
	if req.Scene == util.CodeSceneUnbind {
		var existUser user_model.User
		err := s.db.Where("email = ?", req.Email).First(&existUser).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return response.Failed.SetMessage("该邮箱未注册")
			}
			return response.CheckErr(err, "检查邮箱失败")
		}
	}
	return util.EmailCodeUtil.SendCode(req.Email, req.Scene, "")
}

// SendSmsCode 发送短信验证码（绑定手机/短信登录/手机重置密码场景）
func (s *userService) SendSmsCode(req *user_schema.SendSmsCodeReq) error {
	// 绑定/登录场景：检查手机号是否已被绑定
	if req.Scene == util.SmsSceneBind || req.Scene == util.SmsSceneLogin {
		var existUser user_model.User
		err := s.db.Where("phone = ? AND phone != ''", req.Phone).First(&existUser).Error
		if err == nil && req.Scene == util.SmsSceneBind {
			return response.Failed.SetMessage("该手机号已被其他账号绑定")
		}
	}
	// 重置密码场景：检查手机号已绑定
	if req.Scene == util.SmsSceneReset {
		var existUser user_model.User
		err := s.db.Where("phone = ? AND phone != ''", req.Phone).First(&existUser).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return response.Failed.SetMessage("该手机号未绑定账号")
			}
			return response.CheckErr(err, "检查手机号失败")
		}
	}
	return util.SmsCodeUtil.SendCode(req.Phone, req.Scene)
}

// ResetPassword 邮箱重置密码
func (s *userService) ResetPassword(req *user_schema.ResetPasswordReq) error {
	// 校验验证码
	if err := util.EmailCodeUtil.VerifyCode(req.Email, util.CodeSceneReset, req.Code, ""); err != nil {
		return response.Failed.SetMessage(err.Error())
	}

	// 查找用户
	var user user_model.User
	if err := s.db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.Failed.SetMessage("该邮箱未注册")
		}
		return response.CheckErr(err, "查询用户失败")
	}

	return s.doResetPassword(&user, req.Password)
}

// ResetPhonePassword 手机号重置密码
func (s *userService) ResetPhonePassword(req *user_schema.ResetPhonePasswordReq) error {
	// 校验短信验证码
	if err := util.SmsCodeUtil.VerifyCode(req.Phone, util.SmsSceneReset, req.Code); err != nil {
		return response.Failed.SetMessage(err.Error())
	}

	// 查找用户
	var user user_model.User
	if err := s.db.Where("phone = ? AND phone != ''", req.Phone).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.Failed.SetMessage("该手机号未绑定账号")
		}
		return response.CheckErr(err, "查询用户失败")
	}

	return s.doResetPassword(&user, req.Password)
}

// doResetPassword 统一重置密码逻辑（加密+更新+踢人下线）
func (s *userService) doResetPassword(user *user_model.User, password string) error {
	hashedPwd, err := hashPassword(password)
	if err != nil {
		return response.CheckErr(err, "密码加密失败")
	}

	if err := s.db.Model(user).Updates(map[string]any{
		"password":      hashedPwd,
		"token_version": gorm.Expr("token_version + 1"),
	}).Error; err != nil {
		return response.CheckErr(err, "重置密码失败")
	}

	// 立即清除 Redis 缓存，使旧 token 在下次鉴权时即被拒绝
	s.InvalidateTokenVersionCache(user.ID)
	return nil
}

// KickOffline 踢人下线（自增 token_version，使该用户所有旧 token 失效）
func (s *userService) KickOffline(userID string) error {
	if err := s.db.Model(&user_model.User{}).Where("id = ?", userID).
		Update("token_version", gorm.Expr("token_version + 1")).Error; err != nil {
		return response.CheckErr(err, "操作失败")
	}
	// 立即清除 Redis 缓存
	s.InvalidateTokenVersionCache(userID)
	return nil
}

// GetUserTokenVersion 获取用户当前 token_version（直接查 MySQL，供 GetCachedTokenVersion 回源使用）
func (s *userService) GetUserTokenVersion(userID string) (int64, error) {
	var user user_model.User
	if err := s.db.Select("token_version").Where("id = ?", userID).First(&user).Error; err != nil {
		return 0, err
	}
	return user.TokenVersion, nil
}

// ChangePassword 修改密码（校验原密码 + bcrypt 更新 + 踢人下线）
func (s *userService) ChangePassword(userID, oldPwd, newPwd string) error {
	var user user_model.User
	if err := s.db.Where("id = ?", userID).First(&user).Error; err != nil {
		return response.CheckErr(err, "用户不存在")
	}
	if !checkPassword(user.Password, oldPwd) {
		return response.Failed.SetMessage("原密码错误")
	}
	hashed, err := hashPassword(newPwd)
	if err != nil {
		return response.CheckErr(err, "密码加密失败")
	}
	if err := s.db.Model(&user).Updates(map[string]any{
		"password":      hashed,
		"token_version": gorm.Expr("token_version + 1"),
	}).Error; err != nil {
		return response.CheckErr(err, "修改密码失败")
	}
	// 立即清除 Redis 缓存，使旧 token 失效
	s.InvalidateTokenVersionCache(userID)
	return nil
}
