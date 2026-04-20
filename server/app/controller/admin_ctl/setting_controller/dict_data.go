package setting_controller

import (
	"x_admin/app/schema/setting_schema"
	"x_admin/app/service/setting_service"
	"x_admin/core/response"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

// DictDataHandler 字典数据控制器
type DictDataHandler struct{}

// All 字典数据所有
func (ddh DictDataHandler) All(c *gin.Context) {
	var allReq setting_schema.SettingDictDataListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &allReq)) {
		return
	}
	res, err := setting_service.DictDataService.All(allReq)
	response.CheckAndRespWithData(c, res, err)
}

// Detail 字典数据详情
func (ddh DictDataHandler) Detail(c *gin.Context) {
	var detailReq setting_schema.SettingDictDataDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := setting_service.DictDataService.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// Add 字典数据新增
func (ddh DictDataHandler) Add(c *gin.Context) {
	var addReq setting_schema.SettingDictDataAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, setting_service.DictDataService.Add(addReq))
}

// Edit 字典数据编辑
func (ddh DictDataHandler) Edit(c *gin.Context) {
	var editReq setting_schema.SettingDictDataEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, setting_service.DictDataService.Edit(editReq))
}

// Del 字典数据删除
func (ddh DictDataHandler) Del(c *gin.Context) {
	var delReq setting_schema.SettingDictDataDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, setting_service.DictDataService.Del(delReq))
}
