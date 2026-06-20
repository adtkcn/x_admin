package setting_controller

import (
	"x_admin/app/schema/setting_schema"
	"x_admin/app/service/setting_service"
	"x_admin/core/response"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

// WebsiteHandler 网站设置控制器
type WebsiteHandler struct{}

// @Summary		获取网站信息
// @Description	获取网站设置信息
// @Tags			setting_website-网站设置
// @Param			token	header		string										true	"token"
// @Success		200		{object}	response.Response{data=map[string]string}	"成功"
// @Router			/api/admin/setting/website/detail [get]
func (wh WebsiteHandler) Detail(c *gin.Context) {
	res, err := setting_service.WebsiteService.Detail()
	response.CheckAndRespWithData(c, res, err)
}

// @Summary		保存网站信息
// @Description	保存网站设置信息
// @Tags			setting_website-网站设置
// @Param			token		header		string				true	"token"
// @Param			name		body		string				false	"网站名称"
// @Param			logo		body		string				false	"网站图标"
// @Param			favicon		body		string				false	"网站LOGO"
// @Param			backdrop	body		string				false	"登录页广告图"
// @Param			shopName	body		string				false	"商城名称"
// @Param			shopLogo	body		string				false	"商城Logo"
// @Success		200			{object}	response.Response	"成功"
// @Router			/api/admin/setting/website/save [post]
func (wh WebsiteHandler) Save(c *gin.Context) {
	var wsReq setting_schema.SettingWebsiteReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &wsReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, setting_service.WebsiteService.Save(wsReq))
}
