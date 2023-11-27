package login

import (
	"x_admin/admin/captcha"
	"x_admin/admin/system/admin"
	"x_admin/admin/system/role"
	"x_admin/core"
	"x_admin/core/response"
	"x_admin/middleware"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

func LoginRoute(rg *gin.RouterGroup) {
	db := core.GetDB()
	permSrv := role.NewSystemAuthPermService(db)
	roleSrv := role.NewSystemAuthRoleService(db, permSrv)
	adminSrv := admin.NewSystemAuthAdminService(db, permSrv, roleSrv)
	service := NewSystemLoginService(db, adminSrv)

	handle := loginHandler{Service: service}

	rg = rg.Group("/system", middleware.TokenAuth())
	rg.POST("/login", handle.login)
	rg.POST("/logout", handle.logout)
}

type loginHandler struct {
	Service ISystemLoginService
}

// login 登录系统
func (lh loginHandler) login(c *gin.Context) {
	var params captcha.ClientParams
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &params)) {
		return
	}
	err := captcha.Verify(params)
	if err != nil {
		response.FailWithMsg(c, response.Failed, err.Error())
		return
	}

	var loginReq SystemLoginReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &loginReq)) {
		return
	}
	res, err := lh.Service.Login(c, &loginReq)
	response.CheckAndRespWithData(c, res, err)
}

// logout 登录退出
func (lh loginHandler) logout(c *gin.Context) {
	var logoutReq SystemLogoutReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyHeader(c, &logoutReq)) {
		return
	}
	response.CheckAndResp(c, lh.Service.Logout(&logoutReq))
}
