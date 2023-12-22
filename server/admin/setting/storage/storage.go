package storage

import (
	"x_admin/core/response"
	"x_admin/middleware"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

func StorageRoute(rg *gin.RouterGroup) {
	// db := core.GetDB()
	// permSrv := system.NewSystemAuthPermService(db)
	// roleSrv := system.NewSystemAuthRoleService(db, permSrv)
	// adminSrv := system.NewSystemAuthAdminService(db, permSrv, roleSrv)
	// service := system.NewSystemLoginService(db, adminSrv)

	// server := NewSettingStorageService()

	handle := storageHandler{}

	rg = rg.Group("/setting", middleware.TokenAuth())
	rg.GET("/storage/list", handle.List)
	rg.GET("/storage/detail", handle.Detail)
	rg.POST("/storage/edit", handle.Edit)
	rg.POST("/storage/change", handle.change)
}

type storageHandler struct {
}

// list 存储列表
func (sh storageHandler) List(c *gin.Context) {
	res, err := Service.List()
	response.CheckAndRespWithData(c, res, err)
}

// detail 存储详情
func (sh storageHandler) Detail(c *gin.Context) {
	var detailReq SettingStorageDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := Service.Detail(detailReq.Alias)
	response.CheckAndRespWithData(c, res, err)
}

// edit 存储编辑
func (sh storageHandler) Edit(c *gin.Context) {
	var editReq SettingStorageEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &editReq)) {
		return
	}
	response.CheckAndResp(c, Service.Edit(editReq))
}

// change 存储切换
func (sh storageHandler) change(c *gin.Context) {
	var changeReq SettingStorageChangeReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &changeReq)) {
		return
	}
	response.CheckAndResp(c, Service.Change(changeReq.Alias, changeReq.Status))
}
