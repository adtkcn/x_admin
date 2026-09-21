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

// @Summary		字典数据所有
// @Description	获取字典数据列表(不分页)
// @Tags			setting_dictData-字典数据
// @Param			token		header		string															true	"token"
// @Param			dict_type	query		string															false	"字典类型"
// @Param			name		query		string															false	"键"
// @Param			value		query		string															false	"值"
// @Param			status		query		int8															false	"状态: 0=停用,1=启用"
// @Success		200			{object}	response.Response{data=[]setting_schema.SettingDictDataResp}	"成功"
// @Router			/api/admin/setting/dictData/all [get]
func (ddh DictDataHandler) All(c *gin.Context) {
	var allReq setting_schema.SettingDictDataListReq
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &allReq)) {
		return
	}
	res, err := setting_service.DictDataService.All(allReq)
	response.JSON(c, res, err)
}

// @Summary		字典数据详情
// @Description	获取字典数据详情
// @Tags			setting_dictData-字典数据
// @Param			token	header		string														true	"token"
// @Param			id		query		string														true	"主键"
// @Success		200		{object}	response.Response{data=setting_schema.SettingDictDataResp}	"成功"
// @Router			/api/admin/setting/dictData/detail [get]
func (ddh DictDataHandler) Detail(c *gin.Context) {
	var detailReq setting_schema.SettingDictDataDetailReq
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := setting_service.DictDataService.Detail(detailReq.ID)
	response.JSON(c, res, err)
}

// @Summary		字典数据新增
// @Description	新增字典数据
// @Tags			setting_dictData-字典数据
// @Param			token	header		string				true	"token"
// @Param			type_id	body		string				true	"类型ID"
// @Param			name	body		string				true	"键"
// @Param			value	body		string				true	"值"
// @Param			color	body		string				false	"颜色"
// @Param			remark	body		string				false	"备注"
// @Param			sort	body		int					false	"排序"
// @Param			status	body		int8				false	"状态: 0=停用,1=启用"
// @Success		200		{object}	response.Response	"成功"
// @Router			/api/admin/setting/dictData/add [post]
func (ddh DictDataHandler) Add(c *gin.Context) {
	var addReq setting_schema.SettingDictDataAddReq
	if response.IsFail(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	response.JSON(c, nil, setting_service.DictDataService.Add(addReq))
}

// @Summary		字典数据编辑
// @Description	编辑字典数据
// @Tags			setting_dictData-字典数据
// @Param			token	header		string				true	"token"
// @Param			id		body		string				true	"主键"
// @Param			type_id	body		string				true	"类型ID"
// @Param			name	body		string				true	"键"
// @Param			value	body		string				true	"值"
// @Param			color	body		string				false	"颜色"
// @Param			remark	body		string				false	"备注"
// @Param			sort	body		int					false	"排序"
// @Param			status	body		int8				false	"状态: 0=停用,1=启用"
// @Success		200		{object}	response.Response	"成功"
// @Router			/api/admin/setting/dictData/edit [post]
func (ddh DictDataHandler) Edit(c *gin.Context) {
	var editReq setting_schema.SettingDictDataEditReq
	if response.IsFail(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	response.JSON(c, nil, setting_service.DictDataService.Edit(editReq))
}

// @Summary		字典数据删除
// @Description	删除字典数据
// @Tags			setting_dictData-字典数据
// @Param			token	header		string				true	"token"
// @Param			ids		body		[]string			true	"主键列表"
// @Success		200		{object}	response.Response	"成功"
// @Router			/api/admin/setting/dictData/del [post]
func (ddh DictDataHandler) Del(c *gin.Context) {
	var delReq setting_schema.SettingDictDataDelReq
	if response.IsFail(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.JSON(c, nil, setting_service.DictDataService.Del(delReq))
}
