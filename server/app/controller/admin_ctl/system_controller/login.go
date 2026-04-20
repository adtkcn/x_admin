package system_controller

import (
	"x_admin/app/schema/common_schema"
	"x_admin/app/schema/system_schema"
	"x_admin/app/service/common_service"
	"x_admin/app/service/system_service"
	"x_admin/core/response"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

// LoginHandler 登录控制器
type LoginHandler struct{}

// Login 登录系统
func (lh LoginHandler) Login(c *gin.Context) {
	var params common_schema.ClientParams
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &params)) {
		return
	}
	err := common_service.CaptchaVerify(params)
	if err != nil {
		response.Fail(c, err.Error())
		return
	}

	var loginReq system_schema.SystemLoginReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &loginReq)) {
		return
	}
	res, err := system_service.LoginService.Login(c, &loginReq)
	response.CheckAndRespWithData(c, res, err)
}

// Logout 登录退出
func (lh LoginHandler) Logout(c *gin.Context) {
	var logoutReq system_schema.SystemLogoutReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyHeader(c, &logoutReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, system_service.LoginService.Logout(&logoutReq))
}
