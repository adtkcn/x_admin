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

// @Summary		登录系统
// @Description	管理员登录系统
// @Tags			system_login-登录
// @Param			email		body		string						true	"邮箱(账号)"
// @Param			password	body		string						true	"密码"
// @Param			token		body		string						true	"验证码token"
// @Param			pointJson	body		string						false	"点选坐标"
// @Param			captchaType	body		string						true	"验证码类型"
// @Success		200			{object}	response.Response{data=system_schema.SystemLoginResp}	"成功"
// @Router			/api/admin/system/login [post]
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

// @Summary		登录退出
// @Description	管理员退出登录
// @Tags			system_login-登录
// @Param			token	header		string						true	"token"
// @Success		200		{object}	response.Response			"成功"
// @Router			/api/admin/system/logout [post]
func (lh LoginHandler) Logout(c *gin.Context) {
	var logoutReq system_schema.SystemLogoutReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyHeader(c, &logoutReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, system_service.LoginService.Logout(&logoutReq))
}
