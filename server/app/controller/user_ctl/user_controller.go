package user_ctl

import (
	"x_admin/app/schema/user_schema"
	"x_admin/app/service/user_service"
	"x_admin/config"
	"x_admin/core/response"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

// UserController 用户控制器
type UserController struct{}

// @Summary		用户注册
// @Description	邮箱+密码+验证码注册
// @Tags			user-用户系统
// @Param			email		body		string				true	"邮箱"
// @Param			password	body		string				true	"密码(6-32位)"
// @Param			code		body		string				true	"6位验证码"
// @Param			nickname	body		string				false	"昵称"
// @Success		200			{object}	response.Response	"成功"
// @Router			/api/user/register [post]
func (h UserController) Register(c *gin.Context) {
	var req user_schema.RegisterReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &req)) {
		return
	}
	err := user_service.UserService.Register(&req)
	response.CheckAndRespWithData(c, nil, err)
}

// @Summary		邮箱登录
// @Description	邮箱+密码登录，返回JWT token对
// @Tags			user-用户系统
// @Param			email		body		string											true	"邮箱"
// @Param			password	body		string											true	"密码"
// @Success		200			{object}	response.Response{data=user_schema.LoginResp}	"成功"
// @Router			/api/user/login [post]
func (h UserController) Login(c *gin.Context) {
	var req user_schema.LoginReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &req)) {
		return
	}
	resp, err := user_service.UserService.Login(c, &req)
	response.CheckAndRespWithData(c, resp, err)
}

// @Summary		手机号+密码登录
// @Description	手机号+密码登录，返回JWT token对
// @Tags			user-用户系统
// @Param			phone		body		string											true	"手机号"
// @Param			phoneCode	body		string											false	"区号(默认86)"
// @Param			password	body		string											true	"密码"
// @Success		200			{object}	response.Response{data=user_schema.LoginResp}	"成功"
// @Router			/api/user/phoneLogin [post]
func (h UserController) PhoneLogin(c *gin.Context) {
	var req user_schema.PhoneLoginReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &req)) {
		return
	}
	resp, err := user_service.UserService.PhoneLogin(c, &req)
	response.CheckAndRespWithData(c, resp, err)
}

// @Summary		手机号+短信验证码登录
// @Description	手机号+短信验证码登录
// @Tags			user-用户系统
// @Param			phone		body		string											true	"手机号"
// @Param			phoneCode	body		string											false	"区号(默认86)"
// @Param			code		body		string											true	"6位验证码"
// @Success		200			{object}	response.Response{data=user_schema.LoginResp}	"成功"
// @Router			/api/user/phoneCodeLogin [post]
func (h UserController) PhoneCodeLogin(c *gin.Context) {
	var req user_schema.PhoneCodeLoginReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &req)) {
		return
	}
	resp, err := user_service.UserService.PhoneCodeLogin(c, &req)
	response.CheckAndRespWithData(c, resp, err)
}

// @Summary		刷新token
// @Description	使用refresh_token换取新的token对
// @Tags			user-用户系统
// @Param			refreshToken	body		string											true	"refreshToken"
// @Success		200				{object}	response.Response{data=user_schema.LoginResp}	"成功"
// @Router			/api/user/refresh [post]
func (h UserController) RefreshToken(c *gin.Context) {
	var req user_schema.RefreshTokenReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &req)) {
		return
	}
	resp, err := user_service.UserService.RefreshToken(&req)
	response.CheckAndRespWithData(c, resp, err)
}

// @Summary		获取用户信息
// @Description	获取当前登录用户信息
// @Tags			user-用户系统
// @Param			Authorization	header		string												true	"Bearer token"
// @Success		200				{object}	response.Response{data=user_schema.UserInfoResp}	"成功"
// @Router			/api/user/info [get]
func (h UserController) GetUserInfo(c *gin.Context) {
	userID := config.JWTConfig.GetUserID(c)
	resp, err := user_service.UserService.GetUserInfo(userID)
	response.CheckAndRespWithData(c, resp, err)
}

// @Summary		更新用户信息
// @Description	更新昵称、头像等
// @Tags			user-用户系统
// @Param			Authorization	header		string				true	"Bearer token"
// @Param			nickname		body		string				false	"昵称"
// @Param			avatar			body		string				false	"头像"
// @Success		200				{object}	response.Response	"成功"
// @Router			/api/user/info [post]
func (h UserController) UpdateUserInfo(c *gin.Context) {
	var req user_schema.UpdateUserReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &req)) {
		return
	}
	userID := config.JWTConfig.GetUserID(c)
	err := user_service.UserService.UpdateUserInfo(userID, &req)
	response.CheckAndRespWithData(c, nil, err)
}

// @Summary		发送邮箱验证码
// @Description	发送注册/重置密码/解绑手机邮箱验证码
// @Tags			user-用户系统
// @Param			email	body		string				true	"邮箱"
// @Param			scene	body		string				true	"场景:register/reset/unbind"
// @Success		200		{object}	response.Response	"成功"
// @Router			/api/user/sendCode [post]
func (h UserController) SendEmailCode(c *gin.Context) {
	var req user_schema.SendCodeReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &req)) {
		return
	}
	err := user_service.UserService.SendEmailCode(&req)
	response.CheckAndRespWithData(c, nil, err)
}

// @Summary		发送短信验证码
// @Description	发送绑定手机/短信登录/手机重置密码短信验证码
// @Tags			user-用户系统
// @Param			phone	body		string				true	"手机号"
// @Param			scene	body		string				true	"场景:sms_bind/sms_login/sms_reset"
// @Success		200		{object}	response.Response	"成功"
// @Router			/api/user/sendSmsCode [post]
func (h UserController) SendSmsCode(c *gin.Context) {
	var req user_schema.SendSmsCodeReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &req)) {
		return
	}
	err := user_service.UserService.SendSmsCode(&req)
	response.CheckAndRespWithData(c, nil, err)
}

// @Summary		邮箱重置密码
// @Description	使用邮箱验证码重置密码
// @Tags			user-用户系统
// @Param			email		body		string				true	"邮箱"
// @Param			code		body		string				true	"6位验证码"
// @Param			password	body		string				true	"新密码"
// @Success		200			{object}	response.Response	"成功"
// @Router			/api/user/resetPassword [post]
func (h UserController) ResetPassword(c *gin.Context) {
	var req user_schema.ResetPasswordReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &req)) {
		return
	}
	err := user_service.UserService.ResetPassword(&req)
	response.CheckAndRespWithData(c, nil, err)
}

// @Summary		手机号重置密码
// @Description	使用短信验证码重置密码
// @Tags			user-用户系统
// @Param			phone		body		string				true	"手机号"
// @Param			code		body		string				true	"6位验证码"
// @Param			password	body		string				true	"新密码"
// @Success		200			{object}	response.Response	"成功"
// @Router			/api/user/resetPhonePassword [post]
func (h UserController) ResetPhonePassword(c *gin.Context) {
	var req user_schema.ResetPhonePasswordReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &req)) {
		return
	}
	err := user_service.UserService.ResetPhonePassword(&req)
	response.CheckAndRespWithData(c, nil, err)
}

// @Summary		踢人下线
// @Description	使当前用户所有旧token失效（修改密码后建议调用）
// @Tags			user-用户系统
// @Param			Authorization	header		string				true	"Bearer token"
// @Success		200				{object}	response.Response	"成功"
// @Router			/api/user/kickOffline [post]
func (h UserController) KickOffline(c *gin.Context) {
	userID := config.JWTConfig.GetUserID(c)
	err := user_service.UserService.KickOffline(userID)
	response.CheckAndRespWithData(c, nil, err)
}
