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

// @Summary		获取备案信息
// @Description	获取网站备案信息
// @Tags			setting_copyright-版权设置
// @Param			token	header		string										true	"token"
// @Success		200		{object}	response.Response{data=[]map[string]any}	"成功"
// @Router			/api/admin/setting/copyright/detail [get]
func (ch CopyrightHandler) Detail(c *gin.Context) {
	res, err := setting_service.CopyrightService.Detail()
	response.JSON(c, res, err)
}

// @Summary		保存备案信息
// @Description	保存网站备案信息
// @Tags			setting_copyright-版权设置
// @Param			token	header		string										true	"token"
// @Param			list	body		[]setting_schema.SettingCopyrightItemReq	true	"备案信息列表"
// @Success		200		{object}	response.Response							"成功"
// @Router			/api/admin/setting/copyright/save [post]
func (ch CopyrightHandler) Save(c *gin.Context) {
	var cReqs []setting_schema.SettingCopyrightItemReq
	if response.IsFail(c, util.VerifyUtil.VerifyJSONArray(c, &cReqs)) {
		return
	}
	response.JSON(c, nil, setting_service.CopyrightService.Save(cReqs))
}
