package fabu_controller

import (
	"x_admin/app/schema/fabu_schema"
	"x_admin/app/service/fabu_service"
	"x_admin/core/response"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

// FabuWgtHandler 热更新包(wgt)管理
type FabuWgtHandler struct{}

func (h FabuWgtHandler) List(c *gin.Context) {
	var req fabu_schema.FabuWgtListReq
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &req)) {
		return
	}
	res, err := fabu_service.WgtService.List(req)
	response.JSON(c, res, err)
}

func (h FabuWgtHandler) Upload(c *gin.Context) {
	var req fabu_schema.FabuWgtUploadReq
	if response.IsFail(c, util.VerifyUtil.VerifyBody(c, &req)) {
		return
	}
	response.JSON(c, nil, fabu_service.WgtService.Upload(req))
}

func (h FabuWgtHandler) Release(c *gin.Context) {
	var req fabu_schema.FabuWgtReleaseReq
	if response.IsFail(c, util.VerifyUtil.VerifyBody(c, &req)) {
		return
	}
	response.JSON(c, nil, fabu_service.WgtService.Release(req))
}

func (h FabuWgtHandler) Del(c *gin.Context) {
	var req fabu_schema.FabuWgtDelReq
	if response.IsFail(c, util.VerifyUtil.VerifyBody(c, &req)) {
		return
	}
	response.JSON(c, nil, fabu_service.WgtService.Del(req.ID))
}
