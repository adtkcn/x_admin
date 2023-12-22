package post

import (
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/middleware"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

func PostRoute(rg *gin.RouterGroup) {
	// db := core.GetDB()
	// permSrv := system.NewSystemAuthPermService(db)
	// roleSrv := system.NewSystemAuthRoleService(db, permSrv)
	// adminSrv := system.NewSystemAuthAdminService(db, permSrv, roleSrv)
	// service := system.NewSystemLoginService(db, adminSrv)

	// server := NewSystemAuthPostService(db)

	handle := postHandler{}

	rg = rg.Group("/system", middleware.TokenAuth())
	rg.GET("/post/all", handle.All)
	rg.GET("/post/list", handle.List)
	rg.GET("/post/detail", handle.Detail)
	rg.POST("/post/add", handle.Add)
	rg.POST("/post/edit", handle.Edit)
	rg.POST("/post/del", handle.Del)
}

type postHandler struct {
}

// all 岗位所有
func (ph postHandler) All(c *gin.Context) {
	res, err := Service.All()
	response.CheckAndRespWithData(c, res, err)
}

// list 岗位列表
func (ph postHandler) List(c *gin.Context) {
	var page request.PageReq
	var listReq SystemAuthPostListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := Service.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

// detail 岗位详情
func (ph postHandler) Detail(c *gin.Context) {
	var detailReq SystemAuthPostDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := Service.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// add 岗位新增
func (ph postHandler) Add(c *gin.Context) {
	var addReq SystemAuthPostAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &addReq)) {
		return
	}
	response.CheckAndResp(c, Service.Add(addReq))
}

// edit 岗位编辑
func (ph postHandler) Edit(c *gin.Context) {
	var editReq SystemAuthPostEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &editReq)) {
		return
	}
	response.CheckAndResp(c, Service.Edit(editReq))
}

// del 岗位删除
func (ph postHandler) Del(c *gin.Context) {
	var delReq SystemAuthPostDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, Service.Del(delReq.ID))
}
