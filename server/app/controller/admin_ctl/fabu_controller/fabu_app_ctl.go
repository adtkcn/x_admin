package fabu_controller

import (
	"x_admin/app/schema/fabu_schema"
	"x_admin/app/service/fabu_service"
	"x_admin/core/response"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

// FabuAppHandler 应用管理
type FabuAppHandler struct{}

func (h FabuAppHandler) List(c *gin.Context) {
	var req fabu_schema.FabuAppListReq
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &req)) {
		return
	}
	res, err := fabu_service.AppService.List(req)
	response.JSON(c, res, err)
}

func (h FabuAppHandler) Detail(c *gin.Context) {
	var req fabu_schema.FabuAppDetailReq
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &req)) {
		return
	}
	res, err := fabu_service.AppService.Detail(req.ID)
	response.JSON(c, res, err)
}

func (h FabuAppHandler) Add(c *gin.Context) {
	var req fabu_schema.FabuAppAddReq
	if response.IsFail(c, util.VerifyUtil.VerifyBody(c, &req)) {
		return
	}
	response.JSON(c, nil, fabu_service.AppService.Add(req))
}

func (h FabuAppHandler) Edit(c *gin.Context) {
	var req fabu_schema.FabuAppEditReq
	if response.IsFail(c, util.VerifyUtil.VerifyBody(c, &req)) {
		return
	}
	response.JSON(c, nil, fabu_service.AppService.Edit(req))
}

func (h FabuAppHandler) Del(c *gin.Context) {
	var req fabu_schema.FabuAppDelReq
	if response.IsFail(c, util.VerifyUtil.VerifyBody(c, &req)) {
		return
	}
	response.JSON(c, nil, fabu_service.AppService.Del(req.ID))
}
