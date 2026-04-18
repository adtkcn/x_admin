package settingController

import (
	"x_admin/app/schema/settingSchema"
	"x_admin/app/service/settingService"
	"x_admin/core/response"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

// DictDataHandler 字典数据控制器
type DictDataHandler struct{}

// All 字典数据所有
func (ddh DictDataHandler) All(c *gin.Context) {
	var allReq settingSchema.SettingDictDataListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &allReq)) {
		return
	}
	res, err := settingService.DictDataService.All(allReq)
	response.CheckAndRespWithData(c, res, err)
}

// Detail 字典数据详情
func (ddh DictDataHandler) Detail(c *gin.Context) {
	var detailReq settingSchema.SettingDictDataDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := settingService.DictDataService.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// Add 字典数据新增
func (ddh DictDataHandler) Add(c *gin.Context) {
	var addReq settingSchema.SettingDictDataAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, settingService.DictDataService.Add(addReq))
}

// Edit 字典数据编辑
func (ddh DictDataHandler) Edit(c *gin.Context) {
	var editReq settingSchema.SettingDictDataEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, settingService.DictDataService.Edit(editReq))
}

// Del 字典数据删除
func (ddh DictDataHandler) Del(c *gin.Context) {
	var delReq settingSchema.SettingDictDataDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, settingService.DictDataService.Del(delReq))
}
