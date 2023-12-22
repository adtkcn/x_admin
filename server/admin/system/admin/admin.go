package admin

import (
	"x_admin/config"
	"x_admin/core/request"
	"x_admin/core/response"

	"x_admin/util"

	"github.com/gin-gonic/gin"
)

// func AdminRoute(rg *gin.RouterGroup) {
// 	db := core.GetDB()

// 	permSrv := role.NewSystemAuthPermService(db)
// 	roleSrv := role.NewSystemAuthRoleService(db, permSrv)
// 	adminSrv := NewSystemAuthAdminService(db, permSrv, roleSrv)
// 	// service := NewSystemLoginService(db, adminSrv)

// 	handle := AdminHandler{Service: adminSrv}

// 	rg = rg.Group("/system", middleware.TokenAuth())

// 	rg.GET("/admin/self", handle.self)
// 	rg.GET("/admin/list", handle.List)
// 	rg.GET("/admin/detail", handle.Detail)
// 	rg.POST("/admin/add", middleware.RecordLog("管理员新增"), handle.Add)
// 	rg.POST("/admin/edit", middleware.RecordLog("管理员编辑"), handle.Edit)
// 	rg.POST("/admin/upInfo", middleware.RecordLog("管理员更新"), handle.upInfo)
// 	rg.POST("/admin/del", middleware.RecordLog("管理员删除"), handle.Del)
// 	rg.POST("/admin/disable", middleware.RecordLog("管理员状态切换"), handle.disable)
// }

type AdminHandler struct {
	// Service ISystemAuthAdminService
}

// self 管理员信息
func (ah AdminHandler) Self(c *gin.Context) {
	adminId := config.AdminConfig.GetAdminId(c)
	res, err := Service.Self(adminId)
	response.CheckAndRespWithData(c, res, err)
}

// list 管理员列表
func (ah AdminHandler) List(c *gin.Context) {
	var page request.PageReq
	var listReq SystemAuthAdminListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := Service.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

// detail 管理员详细
func (ah AdminHandler) Detail(c *gin.Context) {
	var detailReq SystemAuthAdminDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := Service.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// add 管理员新增
func (ah AdminHandler) Add(c *gin.Context) {
	var addReq SystemAuthAdminAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	response.CheckAndResp(c, Service.Add(addReq))
}

// edit 管理员编辑
func (ah AdminHandler) Edit(c *gin.Context) {
	var editReq SystemAuthAdminEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	response.CheckAndResp(c, Service.Edit(c, editReq))
}

// upInfo 管理员更新
func (ah AdminHandler) UpInfo(c *gin.Context) {
	var updateReq SystemAuthAdminUpdateReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &updateReq)) {
		return
	}
	response.CheckAndResp(c, Service.Update(
		c, updateReq, config.AdminConfig.GetAdminId(c)))
}

// del 管理员删除
func (ah AdminHandler) Del(c *gin.Context) {
	var delReq SystemAuthAdminDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, Service.Del(c, delReq.ID))
}

// disable 管理员状态切换
func (ah AdminHandler) Disable(c *gin.Context) {
	var disableReq SystemAuthAdminDisableReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &disableReq)) {
		return
	}
	response.CheckAndResp(c, Service.Disable(c, disableReq.ID))
}
