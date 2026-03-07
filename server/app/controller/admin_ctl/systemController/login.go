package systemController

import (
	"x_admin/app/schema/commonSchema"
	"x_admin/app/schema/systemSchema"
	"x_admin/app/service/commonService"
	"x_admin/app/service/systemService"
	"x_admin/core/response"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

func LoginRoute(rg *gin.RouterGroup) {

	handle := loginHandler{}

	rg = rg.Group("/system")
	rg.POST("/login", handle.login)
	rg.POST("/logout", handle.logout)
}

type loginHandler struct{}

// login 登录系统
func (lh loginHandler) login(c *gin.Context) {
	var params commonSchema.ClientParams
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &params)) {
		return
	}
	err := commonService.CaptchaVerify(params)
	if err != nil {
		response.Fail(c, err.Error())
		return
	}

	var loginReq systemSchema.SystemLoginReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &loginReq)) {
		return
	}
	res, err := systemService.LoginService.Login(c, &loginReq)
	response.CheckAndRespWithData(c, res, err)
}

// logout 登录退出
func (lh loginHandler) logout(c *gin.Context) {
	var logoutReq systemSchema.SystemLogoutReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyHeader(c, &logoutReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, systemService.LoginService.Logout(&logoutReq))
}
