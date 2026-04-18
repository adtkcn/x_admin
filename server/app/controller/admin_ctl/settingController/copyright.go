package settingController

import (
	"x_admin/app/schema/settingSchema"
	"x_admin/app/service/settingService"
	"x_admin/core/response"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

// CopyrightHandler 版权设置控制器
type CopyrightHandler struct{}

// Detail 获取备案信息
func (ch CopyrightHandler) Detail(c *gin.Context) {
	res, err := settingService.CopyrightService.Detail()
	response.CheckAndRespWithData(c, res, err)
}

// Save 保存备案信息
func (ch CopyrightHandler) Save(c *gin.Context) {
	var cReqs []settingSchema.SettingCopyrightItemReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSONArray(c, &cReqs)) {
		return
	}
	response.CheckAndRespWithData(c, nil, settingService.CopyrightService.Save(cReqs))
}
