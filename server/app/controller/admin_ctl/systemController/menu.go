package systemController

import (
	"strconv"
	. "x_admin/app/schema/systemSchema"
	"x_admin/app/service/systemService"
	"x_admin/config"
	"x_admin/core/response"
	"x_admin/middleware"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

func MenuRoute(rg *gin.RouterGroup) {
	handle := menuHandler{}
	notAuth := rg.Group("/system", middleware.LoginAuth())
	notAuth.GET("/menu/route", handle.route)

	rg = rg.Group("/system", middleware.TokenAuth())
	// rg.GET("/menu/route", handle.route)
	rg.GET("/menu/list", handle.List)
	rg.GET("/menu/detail", handle.Detail)
	rg.POST("/menu/add", handle.Add)
	rg.POST("/menu/edit", handle.Edit)
	rg.POST("/menu/del", handle.Del)
}

type menuHandler struct {
}

// route 菜单路由
func (mh menuHandler) route(c *gin.Context) {
	roleId := config.AdminConfig.GetRoleId(c)
	id, _ := strconv.ParseUint(roleId, 10, 64)
	res, err := systemService.MenuService.SelectMenuByRoleId(c, uint(id))
	response.CheckAndRespWithData(c, res, err)
}

// list 菜单列表
func (mh menuHandler) List(c *gin.Context) {
	res, err := systemService.MenuService.List()
	response.CheckAndRespWithData(c, res, err)
}

// detail 菜单详情
func (mh menuHandler) Detail(c *gin.Context) {
	var detailReq SystemAuthMenuDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := systemService.MenuService.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// add 新增菜单
func (mh menuHandler) Add(c *gin.Context) {
	var addReq SystemAuthMenuAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	response.CheckAndResp(c, systemService.MenuService.Add(addReq))
}

// edit 编辑菜单
func (mh menuHandler) Edit(c *gin.Context) {
	var editReq SystemAuthMenuEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	response.CheckAndResp(c, systemService.MenuService.Edit(editReq))
}

// del 删除菜单
func (mh menuHandler) Del(c *gin.Context) {
	var delReq SystemAuthMenuDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, systemService.MenuService.Del(delReq.ID))
}
