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

// Detail 获取网站信息
func (wh WebsiteHandler) Detail(c *gin.Context) {
	res, err := setting_service.WebsiteService.Detail()
	response.CheckAndRespWithData(c, res, err)
}

// Save 保存网站信息
func (wh WebsiteHandler) Save(c *gin.Context) {
	var wsReq setting_schema.SettingWebsiteReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &wsReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, setting_service.WebsiteService.Save(wsReq))
}
