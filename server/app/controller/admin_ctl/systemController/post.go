package systemController

import (
	"x_admin/app/schema/systemSchema"
	"x_admin/app/service/systemService"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

// PostHandler 岗位控制器
type PostHandler struct{}

// All 岗位所有
func (ph PostHandler) All(c *gin.Context) {
	res, err := systemService.PostService.All()
	response.CheckAndRespWithData(c, res, err)
}

// List 岗位列表
func (ph PostHandler) List(c *gin.Context) {
	var page request.PageReq
	var listReq systemSchema.SystemAuthPostListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := systemService.PostService.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

// Detail 岗位详情
func (ph PostHandler) Detail(c *gin.Context) {
	var detailReq systemSchema.SystemAuthPostDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := systemService.PostService.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// Add 岗位新增
func (ph PostHandler) Add(c *gin.Context) {
	var addReq systemSchema.SystemAuthPostAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &addReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, systemService.PostService.Add(addReq))
}

// Edit 岗位编辑
func (ph PostHandler) Edit(c *gin.Context) {
	var editReq systemSchema.SystemAuthPostEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &editReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, systemService.PostService.Edit(editReq))
}

// Del 岗位删除
func (ph PostHandler) Del(c *gin.Context) {
	var delReq systemSchema.SystemAuthPostDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &delReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, systemService.PostService.Del(delReq.ID))
}
