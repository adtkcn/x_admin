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
	file, err := c.FormFile("file")
	if err != nil {
		response.JSON(c, nil, response.ParamsValidError.SetMessage("请上传文件"))
		return
	}
	versionId := c.PostForm("version_id")
	if versionId == "" {
		response.JSON(c, nil, response.ParamsValidError.SetMessage("缺少 version_id"))
		return
	}
	e := fabu_service.WgtService.Upload(file, versionId)
	response.JSON(c, nil, e)
}

func (h FabuWgtHandler) Del(c *gin.Context) {
	var req fabu_schema.FabuWgtDelReq
	if response.IsFail(c, util.VerifyUtil.VerifyBody(c, &req)) {
		return
	}
	response.JSON(c, nil, fabu_service.WgtService.Del(req.ID))
}
