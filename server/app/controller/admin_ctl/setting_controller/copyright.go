package setting_controller

import (
	"x_admin/app/schema/setting_schema"
	"x_admin/app/service/setting_service"
	"x_admin/core/response"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

// CopyrightHandler 版权设置控制器
type CopyrightHandler struct{}

// Detail 获取备案信息
func (ch CopyrightHandler) Detail(c *gin.Context) {
	res, err := setting_service.CopyrightService.Detail()
	response.CheckAndRespWithData(c, res, err)
}

// Save 保存备案信息
func (ch CopyrightHandler) Save(c *gin.Context) {
	var cReqs []setting_schema.SettingCopyrightItemReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSONArray(c, &cReqs)) {
		return
	}
	response.CheckAndRespWithData(c, nil, setting_service.CopyrightService.Save(cReqs))
}
