package commonController

import (
	"x_admin/config"
	"x_admin/core/response"
	"x_admin/middleware"
	"x_admin/schema/commonSchema"
	"x_admin/service/commonService"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

func UploadRoute(rg *gin.RouterGroup) {
	handle := uploadHandler{}

	rg = rg.Group("/common", middleware.TokenAuth())
	rg.POST("/upload/image", middleware.RecordLog("上传图片", middleware.RequestFile), handle.uploadImage)
	rg.POST("/upload/video", middleware.RecordLog("上传视频", middleware.RequestFile), handle.uploadVideo)
}

type uploadHandler struct{}

// uploadImage 上传图片
func (uh uploadHandler) uploadImage(c *gin.Context) {
	var uReq commonSchema.CommonUploadImageReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &uReq)) {
		return
	}
	file, ve := util.VerifyUtil.VerifyFile(c, "file")
	if response.IsFailWithResp(c, ve) {
		return
	}
	res, err := commonService.UploadService.UploadImage(file, uReq.Cid, config.AdminConfig.GetAdminId(c))
	response.CheckAndRespWithData(c, res, err)
}

// uploadVideo 上传视频
func (uh uploadHandler) uploadVideo(c *gin.Context) {
	var uReq commonSchema.CommonUploadImageReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &uReq)) {
		return
	}
	file, ve := util.VerifyUtil.VerifyFile(c, "file")
	if response.IsFailWithResp(c, ve) {
		return
	}
	res, err := commonService.UploadService.UploadVideo(file, uReq.Cid, config.AdminConfig.GetAdminId(c))
	response.CheckAndRespWithData(c, res, err)
}
