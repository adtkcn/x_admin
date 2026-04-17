package settingController

import (
	"x_admin/app/middleware"
	"x_admin/app/schema/settingSchema"
	"x_admin/app/service/settingService"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

func DictTypeRoute(rg *gin.RouterGroup) {
	// db := core.GetDB()
	// permSrv := system.NewSystemAuthPermService(db)
	// roleSrv := system.NewSystemAuthRoleService(db, permSrv)
	// adminSrv := system.NewSystemAuthAdminService(db, permSrv, roleSrv)
	// service := system.NewSystemLoginService(db, adminSrv)

	// server := NewSettingDictTypeService(db)

	handle := dictTypeHandler{}
	notAuth := rg.Group("/setting", middleware.LoginAuth())
	notAuth.GET("/dict/type/all", handle.All)

	auth := rg.Group("/setting", middleware.PermAuth())
	auth.GET("/dict/type/list", handle.List)
	auth.GET("/dict/type/detail", handle.Detail)
	auth.POST("/dict/type/add", handle.Add)
	auth.POST("/dict/type/edit", handle.Edit)
	auth.POST("/dict/type/del", handle.Del)
}

type dictTypeHandler struct{}

// all 字典类型所有
func (dth dictTypeHandler) All(c *gin.Context) {
	res, err := settingService.DictTypeService.All()
	response.CheckAndRespWithData(c, res, err)
}

// list 字典类型列表
func (dth dictTypeHandler) List(c *gin.Context) {
	var page request.PageReq
	var listReq settingSchema.SettingDictTypeListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := settingService.DictTypeService.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

// detail 字典类型详情
func (dth dictTypeHandler) Detail(c *gin.Context) {
	var detailReq settingSchema.SettingDictTypeDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := settingService.DictTypeService.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// detail 字典类型新增
func (dth dictTypeHandler) Add(c *gin.Context) {
	var addReq settingSchema.SettingDictTypeAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, settingService.DictTypeService.Add(addReq))
}

// edit 字典类型编辑
func (dth dictTypeHandler) Edit(c *gin.Context) {
	var editReq settingSchema.SettingDictTypeEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, settingService.DictTypeService.Edit(editReq))
}

// del 字典类型删除
func (dth dictTypeHandler) Del(c *gin.Context) {
	var delReq settingSchema.SettingDictTypeDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, settingService.DictTypeService.Del(delReq))
}
