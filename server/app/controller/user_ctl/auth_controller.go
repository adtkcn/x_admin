package user_ctl

import (
	"x_admin/app/schema/user_schema"
	"x_admin/app/service/user_service"
	"x_admin/config"
	"x_admin/core/response"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

// AuthController 认证绑定控制器
type AuthController struct{}

// @Summary		绑定手机号
// @Description	绑定手机号到当前用户（需短信验证码）
// @Tags			user_auth-用户绑定
// @Param			Authorization	header		string				true	"Bearer token"
// @Param			phone			body		string				true	"手机号"
// @Param			phoneCode		body		string				false	"区号(默认86)"
// @Param			code			body		string				true	"短信验证码"
// @Success		200				{object}	response.Response	"成功"
// @Router			/api/user/bindPhone [post]
func (h AuthController) BindPhone(c *gin.Context) {
	var req user_schema.BindPhoneReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &req)) {
		return
	}
	userID := config.JWTConfig.GetUserID(c)
	err := user_service.AuthService.BindPhone(userID, &req)
	response.CheckAndRespWithData(c, nil, err)
}

// @Summary		解绑手机号
// @Description	解绑当前用户手机号（需邮箱验证码确认身份）
// @Tags			user_auth-用户绑定
// @Param			Authorization	header		string				true	"Bearer token"
// @Param			email			body		string				true	"邮箱"
// @Param			code			body		string				true	"邮箱验证码"
// @Success		200				{object}	response.Response	"成功"
// @Router			/api/user/unbindPhone [post]
func (h AuthController) UnbindPhone(c *gin.Context) {
	var req user_schema.UnbindPhoneReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &req)) {
		return
	}
	userID := config.JWTConfig.GetUserID(c)
	err := user_service.AuthService.UnbindPhone(userID, &req)
	response.CheckAndRespWithData(c, nil, err)
}

// @Summary		获取绑定列表
// @Description	获取当前用户所有第三方绑定信息（不含手机号，手机号在用户信息中）
// @Tags			user_auth-用户绑定
// @Param			Authorization	header		string												true	"Bearer token"
// @Success		200				{object}	response.Response{data=[]user_schema.UserAuthItem}	"成功"
// @Router			/api/user/authList [get]
func (h AuthController) GetUserAuthList(c *gin.Context) {
	userID := config.JWTConfig.GetUserID(c)
	resp, err := user_service.AuthService.GetUserAuthList(userID)
	response.CheckAndRespWithData(c, resp, err)
}

// ---- 微信小程序 ----

// @Summary		小程序登录
// @Description	微信小程序登录（wx.login code → openid → 自动注册/登录），返回JWT token对
// @Tags			user_wechat-微信登录
// @Param			code	body		string											true	"小程序 wx.login 返回的 code"
// @Success		200		{object}	response.Response{data=user_schema.LoginResp}	"成功(isNew=true表示新注册用户)"
// @Router			/api/user/wechatMiniLogin [post]
func (h AuthController) WechatMiniLogin(c *gin.Context) {
	var req user_schema.WechatMiniLoginReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &req)) {
		return
	}
	resp, err := user_service.WechatService.MiniLogin(c, &req)
	response.CheckAndRespWithData(c, resp, err)
}

// @Summary		绑定小程序
// @Description	绑定微信小程序到当前已登录用户
// @Tags			user_wechat-微信登录
// @Param			Authorization	header		string				true	"Bearer token"
// @Param			code			body		string				true	"小程序 wx.login code"
// @Success		200				{object}	response.Response	"成功"
// @Router			/api/user/bindWechatMini [post]
func (h AuthController) BindWechatMini(c *gin.Context) {
	var req user_schema.WechatBindReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &req)) {
		return
	}
	userID := config.JWTConfig.GetUserID(c)
	err := user_service.WechatService.BindMini(userID, &req)
	response.CheckAndRespWithData(c, nil, err)
}

// ---- 公众号 ----

// @Summary		公众号登录
// @Description	微信公众号 OAuth 登录（前端跳转微信授权页获取 code → openid → 自动注册/登录）
// @Tags			user_wechat-微信登录
// @Param			code	body		string											true	"微信 OAuth 回调返回的 code"
// @Success		200		{object}	response.Response{data=user_schema.LoginResp}	"成功(isNew=true表示新注册用户)"
// @Router			/api/user/wechatMpLogin [post]
func (h AuthController) WechatMpLogin(c *gin.Context) {
	var req user_schema.WechatMpLoginReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &req)) {
		return
	}
	resp, err := user_service.WechatService.MpLogin(c, &req)
	response.CheckAndRespWithData(c, resp, err)
}

// @Summary		绑定公众号
// @Description	绑定微信公众号到当前已登录用户
// @Tags			user_wechat-微信登录
// @Param			Authorization	header		string				true	"Bearer token"
// @Param			code			body		string				true	"微信 OAuth code"
// @Success		200				{object}	response.Response	"成功"
// @Router			/api/user/bindWechatMp [post]
func (h AuthController) BindWechatMp(c *gin.Context) {
	var req user_schema.WechatBindReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &req)) {
		return
	}
	userID := config.JWTConfig.GetUserID(c)
	err := user_service.WechatService.BindMp(userID, &req)
	response.CheckAndRespWithData(c, nil, err)
}

// ---- 通用解绑 ----

// @Summary		解绑微信
// @Description	解绑微信小程序或公众号（identityType: wechat_mini / wechat_mp）
// @Tags			user_wechat-微信登录
// @Param			Authorization	header		string				true	"Bearer token"
// @Param			identityType	body		string				true	"wechat_mini/wechat_mp"
// @Success		200				{object}	response.Response	"成功"
// @Router			/api/user/unbindWechat [post]
func (h AuthController) UnbindWechat(c *gin.Context) {
	var req user_schema.WechatUnbindReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &req)) {
		return
	}
	userID := config.JWTConfig.GetUserID(c)
	err := user_service.WechatService.UnbindWechat(userID, &req)
	response.CheckAndRespWithData(c, nil, err)
}
