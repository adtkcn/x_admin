package dict_data

import (
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/middleware"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

func DictDataRoute(rg *gin.RouterGroup) {
	// db := core.GetDB()
	// permSrv := system.NewSystemAuthPermService(db)
	// roleSrv := system.NewSystemAuthRoleService(db, permSrv)
	// adminSrv := system.NewSystemAuthAdminService(db, permSrv, roleSrv)
	// service := system.NewSystemLoginService(db, adminSrv)

	// authSrv := NewSettingDictDataService(db)

	handle := dictDataHandler{}

	rg = rg.Group("/setting", middleware.TokenAuth())
	rg.GET("/dict/data/all", handle.All)
	rg.GET("/dict/data/list", handle.List)
	rg.GET("/dict/data/detail", handle.Detail)
	rg.POST("/dict/data/add", handle.Add)
	rg.POST("/dict/data/edit", handle.Edit)
	rg.POST("/dict/data/del", handle.Del)
}

type dictDataHandler struct{}

// all 字典数据所有
func (ddh dictDataHandler) All(c *gin.Context) {
	var allReq SettingDictDataListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &allReq)) {
		return
	}
	res, err := Service.All(allReq)
	response.CheckAndRespWithData(c, res, err)
}

// list 字典数据列表
func (ddh dictDataHandler) List(c *gin.Context) {
	var page request.PageReq
	var listReq SettingDictDataListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := Service.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

// detail 字典数据详情
func (ddh dictDataHandler) Detail(c *gin.Context) {
	var detailReq SettingDictDataDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := Service.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// detail 字典数据新增
func (ddh dictDataHandler) Add(c *gin.Context) {
	var addReq SettingDictDataAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	response.CheckAndResp(c, Service.Add(addReq))
}

// edit 字典数据编辑
func (ddh dictDataHandler) Edit(c *gin.Context) {
	var editReq SettingDictDataEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	response.CheckAndResp(c, Service.Edit(editReq))
}

// del 字典数据删除
func (ddh dictDataHandler) Del(c *gin.Context) {
	var delReq SettingDictDataDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, Service.Del(delReq))
}
