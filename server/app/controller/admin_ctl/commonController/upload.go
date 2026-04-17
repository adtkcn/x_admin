package commonController

import (
	"x_admin/app/middleware"
	"x_admin/app/schema/commonSchema"
	"x_admin/app/service/commonService"
	"x_admin/config"
	"x_admin/core/response"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

func UploadRoute(rg *gin.RouterGroup) {
	handle := uploadHandler{}

	rg = rg.Group("/common", middleware.LoginAuth())
	rg.POST("/upload/preUploadFile", middleware.RecordLog("文件预上传", middleware.RequestFile), handle.preUploadFile)
	rg.POST("/upload/file", middleware.RecordLog("上传文件", middleware.RequestFile), handle.uploadFile)
}

type uploadHandler struct{}

// 文件预上传
func (uh uploadHandler) preUploadFile(c *gin.Context) {
	// md5,fileName,fileSize,cid
	// 检查MD5是否已存在
	// 检查名称是否合规安全
	// 检查文件大小是否超过限制
	// 检查文件类型是否符合要求

	// 如果文件存在，复制到cid对应目录

}

// uploadFile 上传文件
func (uh uploadHandler) uploadFile(c *gin.Context) {
	var uReq commonSchema.CommonUploadImageReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &uReq)) {
		return
	}
	file, ve := util.VerifyUtil.VerifyFile(c, "file")
	if response.IsFailWithResp(c, ve) {
		return
	}
	res, err := commonService.UploadService.UploadFile(file, uReq.Cid, config.AdminConfig.GetAdminId(c))
	response.CheckAndRespWithData(c, res, err)
}
