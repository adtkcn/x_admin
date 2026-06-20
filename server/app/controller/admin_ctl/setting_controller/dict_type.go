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

// @Summary		字典类型所有
// @Description	获取所有字典类型列表(不分页)
// @Tags			setting_dictType-字典类型
// @Param			token	header		string															true	"token"
// @Success		200		{object}	response.Response{data=[]setting_schema.SettingDictTypeResp}	"成功"
// @Router			/api/admin/setting/dictType/all [get]
func (dth DictTypeHandler) All(c *gin.Context) {
	res, err := setting_service.DictTypeService.All()
	response.CheckAndRespWithData(c, res, err)
}

// @Summary		字典类型列表
// @Description	获取字典类型列表
// @Tags			setting_dictType-字典类型
// @Param			token		header		string																				true	"token"
// @Param			pageNo		query		int																					true	"页码"
// @Param			pageSize	query		int																					true	"每页数量"
// @Param			dictName	query		string																				false	"字典名称"
// @Param			dictType	query		string																				false	"字典类型"
// @Param			dictStatus	query		int8																				false	"字典状态: 0/1"
// @Success		200			{object}	response.Response{data=response.PageResp{lists=setting_schema.SettingDictTypeResp}}	"成功"
// @Router			/api/admin/setting/dictType/list [get]
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

// @Summary		字典类型详情
// @Description	获取字典类型详情
// @Tags			setting_dictType-字典类型
// @Param			token	header		string														true	"token"
// @Param			id		query		string														true	"主键"
// @Success		200		{object}	response.Response{data=setting_schema.SettingDictTypeResp}	"成功"
// @Router			/api/admin/setting/dictType/detail [get]
func (dth DictTypeHandler) Detail(c *gin.Context) {
	var detailReq setting_schema.SettingDictTypeDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := setting_service.DictTypeService.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// @Summary		字典类型新增
// @Description	新增字典类型
// @Tags			setting_dictType-字典类型
// @Param			token		header		string				true	"token"
// @Param			dictName	body		string				true	"字典名称"
// @Param			dictType	body		string				true	"字典类型"
// @Param			dictRemark	body		string				false	"字典备注"
// @Param			dictStatus	body		int8				true	"字典状态: 0/1"
// @Success		200			{object}	response.Response	"成功"
// @Router			/api/admin/setting/dictType/add [post]
func (dth DictTypeHandler) Add(c *gin.Context) {
	var addReq setting_schema.SettingDictTypeAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, setting_service.DictTypeService.Add(addReq))
}

// @Summary		字典类型编辑
// @Description	编辑字典类型
// @Tags			setting_dictType-字典类型
// @Param			token		header		string				true	"token"
// @Param			id			body		string				true	"主键"
// @Param			dictName	body		string				true	"字典名称"
// @Param			dictType	body		string				true	"字典类型"
// @Param			dictRemark	body		string				false	"字典备注"
// @Param			dictStatus	body		int8				true	"字典状态: 0/1"
// @Success		200			{object}	response.Response	"成功"
// @Router			/api/admin/setting/dictType/edit [post]
func (dth DictTypeHandler) Edit(c *gin.Context) {
	var editReq setting_schema.SettingDictTypeEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, setting_service.DictTypeService.Edit(editReq))
}

// @Summary		字典类型删除
// @Description	删除字典类型
// @Tags			setting_dictType-字典类型
// @Param			token	header		string				true	"token"
// @Param			ids		body		[]string			true	"主键列表"
// @Success		200		{object}	response.Response	"成功"
// @Router			/api/admin/setting/dictType/del [post]
func (dth DictTypeHandler) Del(c *gin.Context) {
	var delReq setting_schema.SettingDictTypeDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, setting_service.DictTypeService.Del(delReq))
}
