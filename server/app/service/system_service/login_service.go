package system_service

import (
	"errors"

	"x_admin/app/model/system_model"
	"x_admin/app/schema/system_schema"
	"x_admin/config"
	"x_admin/core"
	"x_admin/core/response"

	"x_admin/util"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var LoginService = NewSystemLoginService()

// NewSystemLoginService 初始化
func NewSystemLoginService() *systemLoginService {
	db := core.GetDB()
	return &systemLoginService{db: db}
}

// systemLoginService 系统登录服务实现类
type systemLoginService struct {
	db *gorm.DB
}

// Login 登录
func (loginSrv systemLoginService) Login(c *gin.Context, req *system_schema.SystemLoginReq) (res system_schema.SystemLoginResp, e error) {
	sysAdmin, err := AdminService.FindByEmail(req.Email)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		if e = loginSrv.RecordLoginLog(c, "", req.Email, response.LoginAccountError.Msg()); e != nil {
			return
		}
		e = response.LoginAccountError
		return
	} else if err != nil {
		core.Logger.Errorf("Login FindByEmail err: err=[%+v]", err)
		if e = loginSrv.RecordLoginLog(c, "", req.Email, response.Failed.Msg()); e != nil {
			return
		}
		e = response.Failed
		return
	}
	if sysAdmin.IsDisable == 1 {
		if e = loginSrv.RecordLoginLog(c, sysAdmin.ID, req.Email, response.LoginDisableError.Msg()); e != nil {
			return
		}
		e = response.LoginDisableError
		return
	}
	md5Pwd := util.ToolsUtil.MakeMd5(req.Password + sysAdmin.Salt)
	if sysAdmin.Password != md5Pwd {
		if e = loginSrv.RecordLoginLog(c, sysAdmin.ID, req.Email, response.LoginAccountError.Msg()); e != nil {
			return
		}
		e = response.LoginAccountError
		return
	}

	// 签发 JWT（access + refresh），token_version 内嵌于 claims，踢人靠自增 version 实现
	accessToken, refreshToken, err := util.AdminJWTUtil.GenerateTokenPair(sysAdmin.ID, sysAdmin.TokenVersion)
	if err != nil {
		if e = loginSrv.RecordLoginLog(c, sysAdmin.ID, req.Email, response.SystemError.Msg()); e != nil {
			return
		}
		e = response.SystemError
		return
	}

	AdminService.CacheAdminById(sysAdmin.ID)

	u := system_model.SystemAuthAdmin{LastLoginIp: c.ClientIP(), LastLoginTime: util.NullTimeUtil.Now()}
	// 更新登录信息
	err = loginSrv.db.Model(&sysAdmin).Updates(u).Error
	if err != nil {
		if e = loginSrv.RecordLoginLog(c, sysAdmin.ID, req.Email, response.SystemError.Msg()); e != nil {
			return
		}
		if e = response.CheckErr(err, "Login Updates err"); e != nil {
			return
		}
	}
	// 记录登录日志
	if e = loginSrv.RecordLoginLog(c, sysAdmin.ID, req.Email, ""); e != nil {
		return
	}
	// 返回登录信息（refresh_token 仅用于续签，前端按需缓存）
	return system_schema.SystemLoginResp{
		Token:        accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// Logout 退出：自增 token_version 使当前（及所有）token 失效
func (loginSrv systemLoginService) Logout(c *gin.Context) (e error) {
	adminId := config.AdminConfig.GetAdminId(c)
	if adminId == "" {
		return response.TokenEmpty
	}
	err := loginSrv.db.Model(&system_model.SystemAuthAdmin{}).
		Where("id = ?", adminId).
		Update("token_version", gorm.Expr("token_version + 1")).Error
	if e = response.CheckErr(err, "Logout Update token_version err"); e != nil {
		return
	}
	loginSrv.InvalidateTokenVersionCache(adminId)
	return
}

// IssueNewTokenForAdmin 改密码等场景下，给当前设备重新签发 access_token 实现续签（其他设备失效）
// 逻辑：原子自增 token_version（使其他设备旧 token 失效）→ 清缓存 → 用新 version 签发 access → 写入响应头
func (loginSrv systemLoginService) IssueNewTokenForAdmin(c *gin.Context, adminID string) error {
	if adminID == "" {
		return nil
	}
	// 原子自增 token_version（其他设备旧 token 因 version 不匹配而失效）
	err := loginSrv.db.Model(&system_model.SystemAuthAdmin{}).
		Where("id = ?", adminID).
		Update("token_version", gorm.Expr("token_version + 1")).Error
	if err != nil {
		return response.CheckErr(err, "IssueNewTokenForAdmin Update token_version err")
	}
	// 取最新 version
	var admin system_model.SystemAuthAdmin
	if err = loginSrv.db.Where("id = ?", adminID).First(&admin).Error; err != nil {
		return response.CheckErr(err, "IssueNewTokenForAdmin Find err")
	}
	loginSrv.InvalidateTokenVersionCache(adminID)
	// 给当前设备签发新 access_token（允许短时间旧 token 与新 token 并存，前端拿到后立即刷新）
	newAccess, _, err := util.AdminJWTUtil.GenerateTokenPair(adminID, admin.TokenVersion)
	if err != nil {
		return response.CheckErr(err, "IssueNewTokenForAdmin GenerateTokenPair err")
	}
	c.Header("X-New-Access-Token", newAccess)
	return nil
}

// InvalidateTokenVersionCache 使管理员缓存（含 token_version）失效，下次请求回源取最新版本
func (loginSrv systemLoginService) InvalidateTokenVersionCache(adminID string) {
	util.RedisUtil.Del(config.AdminConfig.BackstageAdminKey + ":" + adminID)
}

// RecordLoginLog 记录登录日志
func (loginSrv systemLoginService) RecordLoginLog(c *gin.Context, adminId string, email string, errStr string) (e error) {
	ua := util.UAUtils.Parse(c.GetHeader("user-agent"))
	var status uint8
	if errStr == "" {
		status = 1
	}
	err := loginSrv.db.Create(&system_model.SystemLogLogin{
		AdminId:    adminId,
		Email:      email,
		Ip:         c.ClientIP(),
		Os:         ua.OsName,
		Browser:    ua.BrowserName,
		Status:     status,
		CreateTime: util.NullTimeUtil.Now(),
	}).Error
	e = response.CheckErr(err, "创建记录失败")
	return
}
