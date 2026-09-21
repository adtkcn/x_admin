package admin_ctl

import (
	"x_admin/app/schema"
	"x_admin/app/service"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

// @Summary	用户列表
// @Tags		user-用户
// @Router		/api/admin/user/list [get]
func (hd *UserHandler) List(c *gin.Context) {
	var page request.PageReq
	var listReq schema.UserListReq
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := service.UserService.List(page, listReq)
	response.JSON(c, res, err)
}

// @Summary	用户详情
// @Router		/api/admin/user/detail [get]
func (hd *UserHandler) Detail(c *gin.Context) {
	var detailReq schema.UserPrimarykey
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := service.UserService.Detail(detailReq.Id)
	response.JSON(c, res, err)
}

// @Summary	用户编辑
// @Router		/api/admin/user/edit [post]
func (hd *UserHandler) Edit(c *gin.Context) {
	var editReq schema.UserEditReq
	if response.IsFail(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	err := service.UserService.Edit(editReq)
	response.JSON(c, editReq.Id, err)
}

// @Summary	用户禁用/启用
// @Router		/api/admin/user/disable [post]
func (hd *UserHandler) Disable(c *gin.Context) {
	var req schema.UserDisableReq
	if response.IsFail(c, util.VerifyUtil.VerifyJSON(c, &req)) {
		return
	}
	err := service.UserService.Disable(req)
	response.JSON(c, nil, err)
}

// @Summary	用户踢下线
// @Router		/api/admin/user/kick [post]
func (hd *UserHandler) Kick(c *gin.Context) {
	var req schema.UserPrimarykey
	if response.IsFail(c, util.VerifyUtil.VerifyJSON(c, &req)) {
		return
	}
	err := service.UserService.Kick(req.Id)
	response.JSON(c, nil, err)
}
