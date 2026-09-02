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
// @Tags			setting_dict_type-字典类型
// @Param			token	header		string															true	"token"
// @Success		200		{object}	response.Response{data=[]setting_schema.SettingDictTypeResp}	"成功"
// @Router			/api/admin/setting/dict_type/all [get]
func (dth DictTypeHandler) All(c *gin.Context) {
	res, err := setting_service.DictTypeService.All()
	response.JSON(c, res, err)
}

// @Summary		字典类型列表
// @Description	获取字典类型列表
// @Tags			setting_dict_type-字典类型
// @Param			token		header		string																				true	"token"
// @Param			pageNo		query		int																					true	"页码"
// @Param			pageSize	query		int																					true	"每页数量"
// @Param			dict_name	query		string																				false	"字典名称"
// @Param			dict_type	query		string																				false	"字典类型"
// @Param			dict_status	query		int8																				false	"字典状态: 0/1"
// @Success		200			{object}	response.Response{data=response.PageResp{lists=setting_schema.SettingDictTypeResp}}	"成功"
// @Router			/api/admin/setting/dict_type/list [get]
func (dth DictTypeHandler) List(c *gin.Context) {
	var page request.PageReq
	var listReq setting_schema.SettingDictTypeListReq
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := setting_service.DictTypeService.List(page, listReq)
	response.JSON(c, res, err)
}

// @Summary		字典类型详情
// @Description	获取字典类型详情
// @Tags			setting_dict_type-字典类型
// @Param			token	header		string														true	"token"
// @Param			id		query		string														true	"主键"
// @Success		200		{object}	response.Response{data=setting_schema.SettingDictTypeResp}	"成功"
// @Router			/api/admin/setting/dict_type/detail [get]
func (dth DictTypeHandler) Detail(c *gin.Context) {
	var detailReq setting_schema.SettingDictTypeDetailReq
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := setting_service.DictTypeService.Detail(detailReq.ID)
	response.JSON(c, res, err)
}

// @Summary		字典类型新增
// @Description	新增字典类型
// @Tags			setting_dict_type-字典类型
// @Param			token		header		string				true	"token"
// @Param			dict_name	body		string				true	"字典名称"
// @Param			dict_type	body		string				true	"字典类型"
// @Param			dict_remark	body		string				false	"字典备注"
// @Param			dict_status	body		int8				true	"字典状态: 0/1"
// @Success		200			{object}	response.Response	"成功"
// @Router			/api/admin/setting/dict_type/add [post]
func (dth DictTypeHandler) Add(c *gin.Context) {
	var addReq setting_schema.SettingDictTypeAddReq
	if response.IsFail(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	response.JSON(c, nil, setting_service.DictTypeService.Add(addReq))
}

// @Summary		字典类型编辑
// @Description	编辑字典类型
// @Tags			setting_dict_type-字典类型
// @Param			token		header		string				true	"token"
// @Param			id			body		string				true	"主键"
// @Param			dict_name	body		string				true	"字典名称"
// @Param			dict_type	body		string				true	"字典类型"
// @Param			dict_remark	body		string				false	"字典备注"
// @Param			dict_status	body		int8				true	"字典状态: 0/1"
// @Success		200			{object}	response.Response	"成功"
// @Router			/api/admin/setting/dict_type/edit [post]
func (dth DictTypeHandler) Edit(c *gin.Context) {
	var editReq setting_schema.SettingDictTypeEditReq
	if response.IsFail(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	response.JSON(c, nil, setting_service.DictTypeService.Edit(editReq))
}

// @Summary		字典类型删除
// @Description	删除字典类型
// @Tags			setting_dict_type-字典类型
// @Param			token	header		string				true	"token"
// @Param			ids		body		[]string			true	"主键列表"
// @Success		200		{object}	response.Response	"成功"
// @Router			/api/admin/setting/dict_type/del [post]
func (dth DictTypeHandler) Del(c *gin.Context) {
	var delReq setting_schema.SettingDictTypeDelReq
	if response.IsFail(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.JSON(c, nil, setting_service.DictTypeService.Del(delReq))
}
