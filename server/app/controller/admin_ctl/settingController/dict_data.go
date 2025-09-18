package settingController

import (
	"x_admin/app/schema/settingSchema"

	"x_admin/app/service/settingService"
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
	notAuth := rg.Group("/setting", middleware.LoginAuth())
	notAuth.GET("/dict/data/all", handle.All)

	rg = rg.Group("/setting", middleware.TokenAuth())
	rg.GET("/dict/data/detail", handle.Detail)
	rg.POST("/dict/data/add", handle.Add)
	rg.POST("/dict/data/edit", handle.Edit)
	rg.POST("/dict/data/del", handle.Del)
}

type dictDataHandler struct{}

// all 字典数据所有
func (ddh dictDataHandler) All(c *gin.Context) {
	var allReq settingSchema.SettingDictDataListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &allReq)) {
		return
	}
	res, err := settingService.DictDataService.All(allReq)
	response.CheckAndRespWithData(c, res, err)
}

// detail 字典数据详情
func (ddh dictDataHandler) Detail(c *gin.Context) {
	var detailReq settingSchema.SettingDictDataDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := settingService.DictDataService.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// detail 字典数据新增
func (ddh dictDataHandler) Add(c *gin.Context) {
	var addReq settingSchema.SettingDictDataAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	response.CheckAndResp(c, settingService.DictDataService.Add(addReq))
}

// edit 字典数据编辑
func (ddh dictDataHandler) Edit(c *gin.Context) {
	var editReq settingSchema.SettingDictDataEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	response.CheckAndResp(c, settingService.DictDataService.Edit(editReq))
}

// del 字典数据删除
func (ddh dictDataHandler) Del(c *gin.Context) {
	var delReq settingSchema.SettingDictDataDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, settingService.DictDataService.Del(delReq))
}
