package systemController

import (
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/middleware"
	. "x_admin/schema/systemSchema"
	"x_admin/service/systemService"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

func RoleRoute(rg *gin.RouterGroup) {

	handle := RoleHandler{}

	rg = rg.Group("/system", middleware.TokenAuth())
	rg.GET("/role/all", handle.All)
	rg.GET("/role/list", middleware.RecordLog("角色列表"), handle.List)
	rg.GET("/role/detail", middleware.RecordLog("角色详情"), handle.Detail)
	rg.POST("/role/add", middleware.RecordLog("角色新增"), handle.Add)
	rg.POST("/role/edit", middleware.RecordLog("角色编辑"), handle.Edit)
	rg.POST("/role/del", middleware.RecordLog("角色删除"), handle.Del)
}

type RoleHandler struct {
	// Service ISystemAuthRoleService
}

// all 角色所有
func (rh RoleHandler) All(c *gin.Context) {
	res, err := systemService.RoleService.All()
	response.CheckAndRespWithData(c, res, err)
}

// list 角色列表
func (rh RoleHandler) List(c *gin.Context) {
	var page request.PageReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	res, err := systemService.RoleService.List(page)
	response.CheckAndRespWithData(c, res, err)
}

// detail 角色详情
func (rh RoleHandler) Detail(c *gin.Context) {
	var detailReq SystemAuthRoleDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := systemService.RoleService.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// add 新增角色
func (rh RoleHandler) Add(c *gin.Context) {
	var addReq SystemAuthRoleAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	response.CheckAndResp(c, systemService.RoleService.Add(addReq))
}

// edit 编辑角色
func (rh RoleHandler) Edit(c *gin.Context) {
	var editReq SystemAuthRoleEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	response.CheckAndResp(c, systemService.RoleService.Edit(editReq))
}

// del 删除角色
func (rh RoleHandler) Del(c *gin.Context) {
	var delReq SystemAuthRoleDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, systemService.RoleService.Del(delReq.ID))
}
