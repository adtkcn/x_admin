package system_controller

import (
	"x_admin/app/schema/system_schema"
	"x_admin/app/service/system_service"
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

	res, err := system_service.MenuService.SelectMenuByAdminId(adminId)
	response.CheckAndRespWithData(c, res, err)
}

// List 菜单列表
func (mh MenuHandler) List(c *gin.Context) {
	res, err := system_service.MenuService.List()
	response.CheckAndRespWithData(c, res, err)
}

// Detail 菜单详情
func (mh MenuHandler) Detail(c *gin.Context) {
	var detailReq system_schema.SystemAuthMenuDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := system_service.MenuService.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// Add 新增菜单
func (mh MenuHandler) Add(c *gin.Context) {
	var addReq system_schema.SystemAuthMenuAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, system_service.MenuService.Add(addReq))
}

// Edit 编辑菜单
func (mh MenuHandler) Edit(c *gin.Context) {
	var editReq system_schema.SystemAuthMenuEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, system_service.MenuService.Edit(editReq))
}

// Del 删除菜单
func (mh MenuHandler) Del(c *gin.Context) {
	var delReq system_schema.SystemAuthMenuDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, system_service.MenuService.Del(delReq.ID))
}
