package upload

import (
	"x_admin/admin/common/album"
	"x_admin/config"
	"x_admin/core/response"
	"x_admin/middleware"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

func UploadRoute(rg *gin.RouterGroup) {
	// db := core.GetDB()
	// permSrv := system.NewSystemAuthPermService(db)
	// roleSrv := system.NewSystemAuthRoleService(db, permSrv)
	// adminSrv := system.NewSystemAuthAdminService(db, permSrv, roleSrv)
	// service := system.NewSystemLoginService(db, adminSrv)
	// AlbumServer := album.NewAlbumService()
	// server := NewUploadService()

	handle := uploadHandler{}

	rg = rg.Group("/common", middleware.TokenAuth())
	rg.POST("/upload/image", middleware.RecordLog("上传图片", middleware.RequestFile), handle.uploadImage)
	rg.POST("/upload/video", middleware.RecordLog("上传视频", middleware.RequestFile), handle.uploadVideo)
}

type uploadHandler struct{}

// uploadImage 上传图片
func (uh uploadHandler) uploadImage(c *gin.Context) {
	var uReq album.CommonUploadImageReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &uReq)) {
		return
	}
	file, ve := util.VerifyUtil.VerifyFile(c, "file")
	if response.IsFailWithResp(c, ve) {
		return
	}
	res, err := Service.UploadImage(file, uReq.Cid, config.AdminConfig.GetAdminId(c))
	response.CheckAndRespWithData(c, res, err)
}

// uploadVideo 上传视频
func (uh uploadHandler) uploadVideo(c *gin.Context) {
	var uReq album.CommonUploadImageReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &uReq)) {
		return
	}
	file, ve := util.VerifyUtil.VerifyFile(c, "file")
	if response.IsFailWithResp(c, ve) {
		return
	}
	res, err := Service.UploadVideo(file, uReq.Cid, config.AdminConfig.GetAdminId(c))
	response.CheckAndRespWithData(c, res, err)
}
