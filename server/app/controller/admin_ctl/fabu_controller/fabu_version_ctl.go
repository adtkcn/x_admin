package fabu_controller

import (
	"x_admin/app/schema/fabu_schema"
	"x_admin/app/service/fabu_service"
	"x_admin/core/response"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

// FabuVersionHandler 版本管理
type FabuVersionHandler struct{}

func (h FabuVersionHandler) List(c *gin.Context) {
	var req fabu_schema.FabuVersionListReq
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &req)) {
		return
	}
	res, err := fabu_service.VersionService.List(req)
	response.JSON(c, res, err)
}

func (h FabuVersionHandler) Upload(c *gin.Context) {
	var req fabu_schema.FabuVersionUploadReq
	if response.IsFail(c, util.VerifyUtil.VerifyBody(c, &req)) {
		return
	}
	res, e := fabu_service.VersionService.Upload(req)
	response.JSON(c, res, e)
}

func (h FabuVersionHandler) Release(c *gin.Context) {
	var req fabu_schema.FabuVersionReleaseReq
	if response.IsFail(c, util.VerifyUtil.VerifyBody(c, &req)) {
		return
	}
	response.JSON(c, nil, fabu_service.VersionService.Release(req.AppId, req.ID))
}

func (h FabuVersionHandler) Cancel(c *gin.Context) {
	var req fabu_schema.FabuVersionReleaseReq
	if response.IsFail(c, util.VerifyUtil.VerifyBody(c, &req)) {
		return
	}
	response.JSON(c, nil, fabu_service.VersionService.Cancel(req.AppId, req.ID))
}

func (h FabuVersionHandler) Gray(c *gin.Context) {
	var req fabu_schema.FabuVersionGrayReq
	if response.IsFail(c, util.VerifyUtil.VerifyBody(c, &req)) {
		return
	}
	response.JSON(c, nil, fabu_service.VersionService.Gray(req.AppId, req.ID, req.Gray))
}

func (h FabuVersionHandler) UpdateMode(c *gin.Context) {
	var req fabu_schema.FabuVersionUpdateModeReq
	if response.IsFail(c, util.VerifyUtil.VerifyBody(c, &req)) {
		return
	}
	response.JSON(c, nil, fabu_service.VersionService.UpdateMode(req.AppId, req.ID, req.UpdateMode))
}

func (h FabuVersionHandler) Del(c *gin.Context) {
	var req fabu_schema.FabuVersionDelReq
	if response.IsFail(c, util.VerifyUtil.VerifyBody(c, &req)) {
		return
	}
	response.JSON(c, nil, fabu_service.VersionService.Del(req.AppId, req.ID))
}
