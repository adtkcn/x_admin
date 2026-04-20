package setting_controller

import (
	"x_admin/app/schema/setting_schema"
	"x_admin/app/service/setting_service"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

// DictTypeHandler 字典类型控制器
type DictTypeHandler struct{}

// All 字典类型所有
func (dth DictTypeHandler) All(c *gin.Context) {
	res, err := setting_service.DictTypeService.All()
	response.CheckAndRespWithData(c, res, err)
}

// List 字典类型列表
func (dth DictTypeHandler) List(c *gin.Context) {
	var page request.PageReq
	var listReq setting_schema.SettingDictTypeListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := setting_service.DictTypeService.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

// Detail 字典类型详情
func (dth DictTypeHandler) Detail(c *gin.Context) {
	var detailReq setting_schema.SettingDictTypeDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := setting_service.DictTypeService.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// Add 字典类型新增
func (dth DictTypeHandler) Add(c *gin.Context) {
	var addReq setting_schema.SettingDictTypeAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, setting_service.DictTypeService.Add(addReq))
}

// Edit 字典类型编辑
func (dth DictTypeHandler) Edit(c *gin.Context) {
	var editReq setting_schema.SettingDictTypeEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, setting_service.DictTypeService.Edit(editReq))
}

// Del 字典类型删除
func (dth DictTypeHandler) Del(c *gin.Context) {
	var delReq setting_schema.SettingDictTypeDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, setting_service.DictTypeService.Del(delReq))
}
