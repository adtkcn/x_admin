package dict_type

import (
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/middleware"
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

	rg = rg.Group("/setting", middleware.TokenAuth())
	rg.GET("/dict/type/all", handle.All)
	rg.GET("/dict/type/list", handle.List)
	rg.GET("/dict/type/detail", handle.Detail)
	rg.POST("/dict/type/add", handle.Add)
	rg.POST("/dict/type/edit", handle.Edit)
	rg.POST("/dict/type/del", handle.Del)
}

type dictTypeHandler struct{}

// all 字典类型所有
func (dth dictTypeHandler) All(c *gin.Context) {
	res, err := Service.All()
	response.CheckAndRespWithData(c, res, err)
}

// list 字典类型列表
func (dth dictTypeHandler) List(c *gin.Context) {
	var page request.PageReq
	var listReq SettingDictTypeListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := Service.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

// detail 字典类型详情
func (dth dictTypeHandler) Detail(c *gin.Context) {
	var detailReq SettingDictTypeDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := Service.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// detail 字典类型新增
func (dth dictTypeHandler) Add(c *gin.Context) {
	var addReq SettingDictTypeAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	response.CheckAndResp(c, Service.Add(addReq))
}

// edit 字典类型编辑
func (dth dictTypeHandler) Edit(c *gin.Context) {
	var editReq SettingDictTypeEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	response.CheckAndResp(c, Service.Edit(editReq))
}

// del 字典类型删除
func (dth dictTypeHandler) Del(c *gin.Context) {
	var delReq SettingDictTypeDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, Service.Del(delReq))
}
