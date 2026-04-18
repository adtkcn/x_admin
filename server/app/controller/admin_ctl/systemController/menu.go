package systemController

import (
	"x_admin/app/schema/systemSchema"
	"x_admin/app/service/systemService"
	"x_admin/config"
	"x_admin/core/response"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

// MenuHandler 菜单控制器
type MenuHandler struct{}

// Route 菜单路由
func (mh MenuHandler) Route(c *gin.Context) {
	adminId := config.AdminConfig.GetAdminId(c)

	res, err := systemService.MenuService.SelectMenuByAdminId(adminId)
	response.CheckAndRespWithData(c, res, err)
}

// List 菜单列表
func (mh MenuHandler) List(c *gin.Context) {
	res, err := systemService.MenuService.List()
	response.CheckAndRespWithData(c, res, err)
}

// Detail 菜单详情
func (mh MenuHandler) Detail(c *gin.Context) {
	var detailReq systemSchema.SystemAuthMenuDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := systemService.MenuService.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// Add 新增菜单
func (mh MenuHandler) Add(c *gin.Context) {
	var addReq systemSchema.SystemAuthMenuAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, systemService.MenuService.Add(addReq))
}

// Edit 编辑菜单
func (mh MenuHandler) Edit(c *gin.Context) {
	var editReq systemSchema.SystemAuthMenuEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, systemService.MenuService.Edit(editReq))
}

// Del 删除菜单
func (mh MenuHandler) Del(c *gin.Context) {
	var delReq systemSchema.SystemAuthMenuDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, systemService.MenuService.Del(delReq.ID))
}
