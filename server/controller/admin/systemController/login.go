package systemController

import (
	"x_admin/admin/common/captcha"
	"x_admin/core/response"
	"x_admin/middleware"
	. "x_admin/schema/systemSchema"
	"x_admin/service/systemService"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

func LoginRoute(rg *gin.RouterGroup) {

	handle := loginHandler{}

	rg = rg.Group("/system", middleware.TokenAuth())
	rg.POST("/login", handle.login)
	rg.POST("/logout", handle.logout)
}

type loginHandler struct{}

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
	res, err := systemService.LoginService.Login(c, &loginReq)
	response.CheckAndRespWithData(c, res, err)
}

// logout 登录退出
func (lh loginHandler) logout(c *gin.Context) {
	var logoutReq SystemLogoutReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyHeader(c, &logoutReq)) {
		return
	}
	response.CheckAndResp(c, systemService.LoginService.Logout(&logoutReq))
}
