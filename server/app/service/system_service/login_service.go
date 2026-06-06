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

	token := util.ToolsUtil.MakeUuidV7()

	// 缓存登录信息
	util.RedisUtil.Set(config.AdminConfig.BackstageTokenKey+token, sysAdmin.ID, config.AdminConfig.TokenExpire)
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
	// 返回登录信息
	return system_schema.SystemLoginResp{Token: token}, nil
}

// Logout 退出
func (loginSrv systemLoginService) Logout(req *system_schema.SystemLogoutReq) (e error) {
	util.RedisUtil.Del(config.AdminConfig.BackstageTokenKey + req.Token)
	return
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
