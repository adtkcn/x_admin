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
// @Param			email		body		string													true	"邮箱(账号)"
// @Param			password	body		string													true	"密码"
// @Param			token		body		string													true	"验证码token"
// @Param			pointJson	body		string													false	"点选坐标"
// @Param			captchaType	body		string													true	"验证码类型"
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
// @Param			token	header		string				true	"token"
// @Success		200		{object}	response.Response	"成功"
// @Router			/api/admin/system/logout [post]
func (lh LoginHandler) Logout(c *gin.Context) {
	response.CheckAndRespWithData(c, nil, system_service.LoginService.Logout(c))
}

// @Summary		忘记密码-发送验证码
// @Description	向注册邮箱发送密码重置验证码
// @Tags			system_login-登录
// @Param			email	body		string				true	"注册邮箱"
// @Success		200		{object}	response.Response	"成功"
// @Router			/api/admin/system/forgot-pwd/send-code [post]
func (lh LoginHandler) ForgotPwdSendCode(c *gin.Context) {
	var req system_schema.SystemForgotPwdSendCodeReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &req)) {
		return
	}
	err := system_service.ForgetPwdService.SendResetCode(&req)
	response.CheckAndRespWithData(c, nil, err)
}

// @Summary		忘记密码-重置密码
// @Description	使用邮箱验证码重置密码
// @Tags			system_login-登录
// @Param			email		body		string				true	"注册邮箱"
// @Param			code		body		string				true	"6位验证码"
// @Param			password	body		string				true	"新密码(MD5加密后)"
// @Success		200			{object}	response.Response	"成功"
// @Router			/api/admin/system/forgot-pwd/reset [post]
func (lh LoginHandler) ForgotPwdReset(c *gin.Context) {
	var req system_schema.SystemForgotPwdResetReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &req)) {
		return
	}
	err := system_service.ForgetPwdService.ResetPassword(&req)
	response.CheckAndRespWithData(c, nil, err)
}
