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

// @Summary		菜单路由
// @Description	获取当前管理员的菜单路由
// @Tags			system_menu-菜单
// @Param			token	header		string														true	"token"
// @Success		200		{object}	response.Response{data=[]system_schema.SystemAuthMenuResp}	"成功"
// @Router			/api/admin/system/menu/route [get]
func (mh MenuHandler) Route(c *gin.Context) {
	adminId := config.AdminConfig.GetAdminId(c)

	res, err := system_service.MenuService.SelectMenuByAdminId(adminId)
	response.CheckAndRespWithData(c, res, err)
}

// @Summary		菜单列表
// @Description	获取菜单列表
// @Tags			system_menu-菜单
// @Param			token	header		string														true	"token"
// @Success		200		{object}	response.Response{data=[]system_schema.SystemAuthMenuResp}	"成功"
// @Router			/api/admin/system/menu/list [get]
func (mh MenuHandler) List(c *gin.Context) {
	res, err := system_service.MenuService.List()
	response.CheckAndRespWithData(c, res, err)
}

// @Summary		菜单详情
// @Description	获取菜单详情
// @Tags			system_menu-菜单
// @Param			token	header		string														true	"token"
// @Param			id		query		string														true	"主键"
// @Success		200		{object}	response.Response{data=system_schema.SystemAuthMenuResp}	"成功"
// @Router			/api/admin/system/menu/detail [get]
func (mh MenuHandler) Detail(c *gin.Context) {
	var detailReq system_schema.SystemAuthMenuDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := system_service.MenuService.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// @Summary		新增菜单
// @Description	新增菜单
// @Tags			system_menu-菜单
// @Param			token		header		string				true	"token"
// @Param			pid			body		string				false	"上级菜单"
// @Param			menuType	body		string				true	"权限类型: [M=目录, C=菜单, A=按钮]"
// @Param			menuName	body		string				true	"菜单名称"
// @Param			menuIcon	body		string				false	"菜单图标"
// @Param			menuSort	body		int					false	"菜单排序"
// @Param			perms		body		string				false	"权限标识"
// @Param			paths		body		string				false	"路由地址"
// @Param			component	body		string				false	"前端组件"
// @Param			selected	body		string				false	"选中路径"
// @Param			params		body		string				false	"路由参数"
// @Param			isCache		body		uint8				false	"是否缓存: [0=否, 1=是]"
// @Param			isShow		body		uint8				false	"是否显示: [0=否, 1=是]"
// @Param			isDisable	body		uint8				false	"是否禁用: [0=否, 1=是]"
// @Success		200			{object}	response.Response	"成功"
// @Router			/api/admin/system/menu/add [post]
func (mh MenuHandler) Add(c *gin.Context) {
	var addReq system_schema.SystemAuthMenuAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, system_service.MenuService.Add(addReq))
}

// @Summary		编辑菜单
// @Description	编辑菜单
// @Tags			system_menu-菜单
// @Param			token		header		string				true	"token"
// @Param			id			body		string				true	"主键"
// @Param			pid			body		string				false	"上级菜单"
// @Param			menuType	body		string				true	"权限类型: [M=目录, C=菜单, A=按钮]"
// @Param			menuName	body		string				true	"菜单名称"
// @Param			menuIcon	body		string				false	"菜单图标"
// @Param			menuSort	body		int					false	"菜单排序"
// @Param			perms		body		string				false	"权限标识"
// @Param			paths		body		string				false	"路由地址"
// @Param			component	body		string				false	"前端组件"
// @Param			selected	body		string				false	"选中路径"
// @Param			params		body		string				false	"路由参数"
// @Param			isCache		body		uint8				false	"是否缓存: [0=否, 1=是]"
// @Param			isShow		body		uint8				false	"是否显示: [0=否, 1=是]"
// @Param			isDisable	body		uint8				false	"是否禁用: [0=否, 1=是]"
// @Success		200			{object}	response.Response	"成功"
// @Router			/api/admin/system/menu/edit [post]
func (mh MenuHandler) Edit(c *gin.Context) {
	var editReq system_schema.SystemAuthMenuEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, system_service.MenuService.Edit(editReq))
}

// @Summary		删除菜单
// @Description	删除菜单
// @Tags			system_menu-菜单
// @Param			token	header		string				true	"token"
// @Param			id		body		string				true	"主键"
// @Success		200		{object}	response.Response	"成功"
// @Router			/api/admin/system/menu/del [post]
func (mh MenuHandler) Del(c *gin.Context) {
	var delReq system_schema.SystemAuthMenuDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, system_service.MenuService.Del(delReq.ID))
}
