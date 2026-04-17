package systemController

import (
	"x_admin/app/middleware"
	"x_admin/app/schema/systemSchema"
	"x_admin/app/service/systemService"
	"x_admin/core/response"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

func DeptRoute(rg *gin.RouterGroup) {
	handle := deptHandler{}
	notAuth := rg.Group("/system", middleware.LoginAuth())
	notAuth.GET("/dept/list", handle.List)

	rg = rg.Group("/system", middleware.PermAuth())
	rg.GET("/dept/all", handle.All)
	rg.GET("/dept/detail", handle.Detail)
	rg.POST("/dept/add", handle.Add)
	rg.POST("/dept/edit", handle.Edit)
	rg.POST("/dept/del", handle.Del)
}

type deptHandler struct {
}

// all 部门所有
func (dh deptHandler) All(c *gin.Context) {
	res, err := systemService.DeptService.All()
	response.CheckAndRespWithData(c, res, err)
}

// list 部门列表
func (dh deptHandler) List(c *gin.Context) {
	var listReq systemSchema.SystemAuthDeptListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := systemService.DeptService.List(listReq)
	response.CheckAndRespWithData(c, res, err)
}

// detail 部门详情
func (dh deptHandler) Detail(c *gin.Context) {
	var detailReq systemSchema.SystemAuthDeptDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := systemService.DeptService.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// add 部门新增
func (dh deptHandler) Add(c *gin.Context) {
	var addReq systemSchema.SystemAuthDeptAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &addReq)) {
		return
	}
	err := systemService.DeptService.Add(addReq)
	response.CheckAndRespWithData(c, nil, err)
}

// edit 部门编辑
func (dh deptHandler) Edit(c *gin.Context) {
	var editReq systemSchema.SystemAuthDeptEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &editReq)) {
		return
	}
	err := systemService.DeptService.Edit(editReq)
	response.CheckAndRespWithData(c, nil, err)
}

// del 部门删除
func (dh deptHandler) Del(c *gin.Context) {
	var delReq systemSchema.SystemAuthDeptDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &delReq)) {
		return
	}
	err := systemService.DeptService.Del(delReq.ID)
	response.CheckAndRespWithData(c, nil, err)
}
