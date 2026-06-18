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
// @Param			Authorization	header		string					true	"Bearer token"
// @Param			phone			body		string					true	"手机号"
// @Param			phoneCode		body		string					false	"区号(默认86)"
// @Param			code			body		string					true	"短信验证码"
// @Success		200				{object}	response.Response		"成功"
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
// @Param			Authorization	header		string					true	"Bearer token"
// @Param			email			body		string					true	"邮箱"
// @Param			code			body		string					true	"邮箱验证码"
// @Success		200				{object}	response.Response		"成功"
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
// @Param			Authorization	header		string							true	"Bearer token"
// @Success		200				{object}	response.Response{data=[]user_schema.UserAuthItem}	"成功"
// @Router			/api/user/authList [get]
func (h AuthController) GetUserAuthList(c *gin.Context) {
	userID := config.JWTConfig.GetUserID(c)
	resp, err := user_service.AuthService.GetUserAuthList(userID)
	response.CheckAndRespWithData(c, resp, err)
}
