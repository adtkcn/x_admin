package system_controller

import (
	"x_admin/app/schema/system_schema"
	"x_admin/app/service/system_service"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

// RoleHandler 角色控制器
type RoleHandler struct{}

// @Summary		角色所有
// @Description	获取所有角色列表(不分页)
// @Tags			system_role-角色
// @Param			token	header		string																true	"token"
// @Success		200		{object}	response.Response{data=[]system_schema.SystemAuthRoleSimpleResp}	"成功"
// @Router			/api/admin/system/role/all [get]
func (rh RoleHandler) All(c *gin.Context) {
	res, err := system_service.RoleService.All()
	response.CheckAndRespWithData(c, res, err)
}

// @Summary		角色列表
// @Description	获取角色列表
// @Tags			system_role-角色
// @Param			token		header		string																				true	"token"
// @Param			pageNo		query		int																					true	"页码"
// @Param			pageSize	query		int																					true	"每页数量"
// @Success		200			{object}	response.Response{data=response.PageResp{lists=system_schema.SystemAuthRoleResp}}	"成功"
// @Router			/api/admin/system/role/list [get]
func (rh RoleHandler) List(c *gin.Context) {
	var page request.PageReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	res, err := system_service.RoleService.List(page)
	response.CheckAndRespWithData(c, res, err)
}

// @Summary		角色详情
// @Description	获取角色详情
// @Tags			system_role-角色
// @Param			token	header		string														true	"token"
// @Param			id		query		string														true	"主键"
// @Success		200		{object}	response.Response{data=system_schema.SystemAuthRoleResp}	"成功"
// @Router			/api/admin/system/role/detail [get]
func (rh RoleHandler) Detail(c *gin.Context) {
	var detailReq system_schema.SystemAuthRoleDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := system_service.RoleService.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// @Summary		新增角色
// @Description	新增角色
// @Tags			system_role-角色
// @Param			token		header		string				true	"token"
// @Param			name		body		string				true	"角色名称"
// @Param			sort		body		int					false	"角色排序"
// @Param			isDisable	body		uint8				false	"是否禁用: [0=否, 1=是]"
// @Param			remark		body		string				false	"角色备注"
// @Param			menuIds		body		string				false	"关联菜单"
// @Success		200			{object}	response.Response	"成功"
// @Router			/api/admin/system/role/add [post]
func (rh RoleHandler) Add(c *gin.Context) {
	var addReq system_schema.SystemAuthRoleAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, system_service.RoleService.Add(addReq))
}

// @Summary		编辑角色
// @Description	编辑角色
// @Tags			system_role-角色
// @Param			token		header		string				true	"token"
// @Param			id			body		string				true	"主键"
// @Param			name		body		string				true	"角色名称"
// @Param			sort		body		int					false	"角色排序"
// @Param			isDisable	body		uint8				false	"是否禁用: [0=否, 1=是]"
// @Param			remark		body		string				false	"角色备注"
// @Param			menuIds		body		string				false	"关联菜单"
// @Success		200			{object}	response.Response	"成功"
// @Router			/api/admin/system/role/edit [post]
func (rh RoleHandler) Edit(c *gin.Context) {
	var editReq system_schema.SystemAuthRoleEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, system_service.RoleService.Edit(editReq))
}

// @Summary		删除角色
// @Description	删除角色
// @Tags			system_role-角色
// @Param			token	header		string				true	"token"
// @Param			id		body		string				true	"主键"
// @Success		200		{object}	response.Response	"成功"
// @Router			/api/admin/system/role/del [post]
func (rh RoleHandler) Del(c *gin.Context) {
	var delReq system_schema.SystemAuthRoleDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, system_service.RoleService.Del(delReq.ID))
}
