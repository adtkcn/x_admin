package dept

import (
	"x_admin/core/response"
	"x_admin/middleware"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

func DeptRoute(rg *gin.RouterGroup) {
	// db := core.GetDB()
	// permSrv := system.NewSystemAuthPermService(db)
	// roleSrv := system.NewSystemAuthRoleService(db, permSrv)
	// adminSrv := system.NewSystemAuthAdminService(db, permSrv, roleSrv)
	// service := system.NewSystemLoginService(db, adminSrv)
	// authSrv := system.NewSystemAuthMenuService(db, permSrv)

	handle := deptHandler{}

	rg = rg.Group("/system", middleware.TokenAuth())
	rg.GET("/dept/all", handle.All)
	rg.GET("/dept/list", handle.List)
	rg.GET("/dept/detail", handle.Detail)
	rg.POST("/dept/add", handle.Add)
	rg.POST("/dept/edit", handle.Edit)
	rg.POST("/dept/del", handle.Del)
}

type deptHandler struct {
}

// all 部门所有
func (dh deptHandler) All(c *gin.Context) {
	res, err := Service.All()
	response.CheckAndRespWithData(c, res, err)
}

// list 部门列表
func (dh deptHandler) List(c *gin.Context) {
	var listReq SystemAuthDeptListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := Service.List(listReq)
	response.CheckAndRespWithData(c, res, err)
}

// detail 部门详情
func (dh deptHandler) Detail(c *gin.Context) {
	var detailReq SystemAuthDeptDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := Service.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// add 部门新增
func (dh deptHandler) Add(c *gin.Context) {
	var addReq SystemAuthDeptAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &addReq)) {
		return
	}
	response.CheckAndResp(c, Service.Add(addReq))
}

// edit 部门编辑
func (dh deptHandler) Edit(c *gin.Context) {
	var editReq SystemAuthDeptEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &editReq)) {
		return
	}
	response.CheckAndResp(c, Service.Edit(editReq))
}

// del 部门删除
func (dh deptHandler) Del(c *gin.Context) {
	var delReq SystemAuthDeptDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, Service.Del(delReq.ID))
}
