package login

import (
	"errors"
	"runtime/debug"
	"strconv"
	"x_admin/admin/system/admin"
	"x_admin/config"
	"x_admin/core"
	"x_admin/core/response"
	"x_admin/model/system_model"
	"x_admin/util"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ISystemLoginService interface {
	Login(c *gin.Context, req *SystemLoginReq) (res SystemLoginResp, e error)
	Logout(req *SystemLogoutReq) (e error)
	RecordLoginLog(c *gin.Context, adminId uint, username string, errStr string) (e error)
}

var Service = NewSystemLoginService()

// NewSystemLoginService 初始化
func NewSystemLoginService() ISystemLoginService {
	db := core.GetDB()
	return &systemLoginService{db: db}
}

// systemLoginService 系统登录服务实现类
type systemLoginService struct {
	db *gorm.DB
}

// Login 登录
func (loginSrv systemLoginService) Login(c *gin.Context, req *SystemLoginReq) (res SystemLoginResp, e error) {
	sysAdmin, err := admin.Service.FindByUsername(req.Username)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		if e = loginSrv.RecordLoginLog(c, 0, req.Username, response.LoginAccountError.Msg()); e != nil {
			return
		}
		e = response.LoginAccountError
		return
	} else if err != nil {
		core.Logger.Errorf("Login FindByUsername err: err=[%+v]", err)
		if e = loginSrv.RecordLoginLog(c, 0, req.Username, response.Failed.Msg()); e != nil {
			return
		}
		e = response.Failed
		return
	}
	if sysAdmin.IsDelete == 1 {
		if e = loginSrv.RecordLoginLog(c, 0, req.Username, response.LoginAccountError.Msg()); e != nil {
			return
		}
		e = response.LoginAccountError
		return
	}
	if sysAdmin.IsDisable == 1 {
		if e = loginSrv.RecordLoginLog(c, sysAdmin.ID, req.Username, response.LoginDisableError.Msg()); e != nil {
			return
		}
		e = response.LoginDisableError
		return
	}
	md5Pwd := util.ToolsUtil.MakeMd5(req.Password + sysAdmin.Salt)
	if sysAdmin.Password != md5Pwd {
		if e = loginSrv.RecordLoginLog(c, sysAdmin.ID, req.Username, response.LoginAccountError.Msg()); e != nil {
			return
		}
		e = response.LoginAccountError
		return
	}
	defer func() {
		if r := recover(); r != nil {
			switch r.(type) {
			// 自定义类型
			case response.RespType:
				panic(r)
			// 其他类型
			default:
				core.Logger.Errorf("stacktrace from panic: %+v\n%s", r, string(debug.Stack()))
				loginSrv.RecordLoginLog(c, sysAdmin.ID, req.Username, response.Failed.Msg())
				panic(response.Failed)
			}
		}
	}()
	token := util.ToolsUtil.MakeToken()
	adminIdStr := strconv.FormatUint(uint64(sysAdmin.ID), 10)

	// 缓存登录信息
	util.RedisUtil.Set(config.AdminConfig.BackstageTokenKey+token, adminIdStr, 7200)
	admin.Service.CacheAdminUserByUid(sysAdmin.ID)

	u := system_model.SystemAuthAdmin{LastLoginIp: c.ClientIP(), LastLoginTime: util.NullTimeUtil.Now()}
	// 更新登录信息
	err = loginSrv.db.Model(&sysAdmin).Updates(u).Error
	if err != nil {
		if e = loginSrv.RecordLoginLog(c, sysAdmin.ID, req.Username, response.SystemError.Msg()); e != nil {
			return
		}
		if e = response.CheckErr(err, "Login Updates err"); e != nil {
			return
		}
	}
	// 记录登录日志
	if e = loginSrv.RecordLoginLog(c, sysAdmin.ID, req.Username, ""); e != nil {
		return
	}
	// 返回登录信息
	return SystemLoginResp{Token: token}, nil
}

// Logout 退出
func (loginSrv systemLoginService) Logout(req *SystemLogoutReq) (e error) {
	util.RedisUtil.Del(config.AdminConfig.BackstageTokenKey + req.Token)
	return
}

// RecordLoginLog 记录登录日志
func (loginSrv systemLoginService) RecordLoginLog(c *gin.Context, adminId uint, username string, errStr string) (e error) {
	ua := core.UAParser.Parse(c.GetHeader("user-agent"))
	var status uint8
	if errStr == "" {
		status = 1
	}
	err := loginSrv.db.Create(&system_model.SystemLogLogin{
		AdminId:    adminId,
		Username:   username,
		Ip:         c.ClientIP(),
		Os:         ua.Os.Family,
		Browser:    ua.UserAgent.Family,
		Status:     status,
		CreateTime: util.NullTimeUtil.Now(),
	}).Error
	e = response.CheckErr(err, "创建记录失败")
	return
}
