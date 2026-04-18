package systemController

import (
	"x_admin/app/schema/systemSchema"
	"x_admin/app/service/systemService"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

// RoleHandler 角色控制器
type RoleHandler struct{}

// All 角色所有
func (rh RoleHandler) All(c *gin.Context) {
	res, err := systemService.RoleService.All()
	response.CheckAndRespWithData(c, res, err)
}

// List 角色列表
func (rh RoleHandler) List(c *gin.Context) {
	var page request.PageReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	res, err := systemService.RoleService.List(page)
	response.CheckAndRespWithData(c, res, err)
}

// Detail 角色详情
func (rh RoleHandler) Detail(c *gin.Context) {
	var detailReq systemSchema.SystemAuthRoleDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := systemService.RoleService.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// Add 新增角色
func (rh RoleHandler) Add(c *gin.Context) {
	var addReq systemSchema.SystemAuthRoleAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, systemService.RoleService.Add(addReq))
}

// Edit 编辑角色
func (rh RoleHandler) Edit(c *gin.Context) {
	var editReq systemSchema.SystemAuthRoleEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, systemService.RoleService.Edit(editReq))
}

// Del 删除角色
func (rh RoleHandler) Del(c *gin.Context) {
	var delReq systemSchema.SystemAuthRoleDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, systemService.RoleService.Del(delReq.ID))
}
