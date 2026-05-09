package system_controller

import (
	"x_admin/app/schema/system_schema"
	"x_admin/app/service/system_service"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

// PostHandler 岗位控制器
type PostHandler struct{}

// @Summary		岗位所有
// @Description	获取所有岗位列表(不分页)
// @Tags			system_post-岗位
// @Param			token	header		string					true	"token"
// @Success		200		{object}	response.Response{data=[]system_schema.SystemAuthPostResp}	"成功"
// @Router			/api/admin/system/post/all [get]
func (ph PostHandler) All(c *gin.Context) {
	res, err := system_service.PostService.All()
	response.CheckAndRespWithData(c, res, err)
}

// @Summary		岗位列表
// @Description	获取岗位列表
// @Tags			system_post-岗位
// @Param			token		header		string					true	"token"
// @Param			pageNo		query		int						true	"页码"
// @Param			pageSize	query		int						true	"每页数量"
// @Param			name		query		string					false	"岗位名称"
// @Param			code		query		string					false	"岗位编码"
// @Param			isStop		query		int8					false	"是否停用: [0=否, 1=是]"
// @Success		200			{object}	response.Response{data=response.PageResp{lists=system_schema.SystemAuthPostResp}}	"成功"
// @Router			/api/admin/system/post/list [get]
func (ph PostHandler) List(c *gin.Context) {
	var page request.PageReq
	var listReq system_schema.SystemAuthPostListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := system_service.PostService.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

// @Summary		岗位详情
// @Description	获取岗位详情
// @Tags			system_post-岗位
// @Param			token	header		string					true	"token"
// @Param			id		query		string					true	"主键"
// @Success		200		{object}	response.Response{data=system_schema.SystemAuthPostResp}	"成功"
// @Router			/api/admin/system/post/detail [get]
func (ph PostHandler) Detail(c *gin.Context) {
	var detailReq system_schema.SystemAuthPostDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := system_service.PostService.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// @Summary		岗位新增
// @Description	新增岗位
// @Tags			system_post-岗位
// @Param			token		header		string					true	"token"
// @Param			name		body		string					true	"岗位名称"
// @Param			code		body		string					false	"岗位编码"
// @Param			remarks		body		string					false	"岗位备注"
// @Param			isStop		body		uint8					false	"是否停用: [0=否, 1=是]"
// @Param			sort		body		int						false	"排序"
// @Success		200			{object}	response.Response		"成功"
// @Router			/api/admin/system/post/add [post]
func (ph PostHandler) Add(c *gin.Context) {
	var addReq system_schema.SystemAuthPostAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &addReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, system_service.PostService.Add(addReq))
}

// @Summary		岗位编辑
// @Description	编辑岗位
// @Tags			system_post-岗位
// @Param			token		header		string					true	"token"
// @Param			id			body		string					true	"主键"
// @Param			name		body		string					true	"岗位名称"
// @Param			code		body		string					false	"岗位编码"
// @Param			remarks		body		string					false	"岗位备注"
// @Param			isStop		body		uint8					false	"是否停用: [0=否, 1=是]"
// @Param			sort		body		int						false	"排序"
// @Success		200			{object}	response.Response		"成功"
// @Router			/api/admin/system/post/edit [post]
func (ph PostHandler) Edit(c *gin.Context) {
	var editReq system_schema.SystemAuthPostEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &editReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, system_service.PostService.Edit(editReq))
}

// @Summary		岗位删除
// @Description	删除岗位
// @Tags			system_post-岗位
// @Param			token	header		string					true	"token"
// @Param			id		body		string					true	"主键"
// @Success		200		{object}	response.Response		"成功"
// @Router			/api/admin/system/post/del [post]
func (ph PostHandler) Del(c *gin.Context) {
	var delReq system_schema.SystemAuthPostDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &delReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, system_service.PostService.Del(delReq.ID))
}
