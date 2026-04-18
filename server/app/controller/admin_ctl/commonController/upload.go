package commonController

import (
	"x_admin/app/schema/commonSchema"
	"x_admin/app/service/commonService"
	"x_admin/config"
	"x_admin/core/response"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

// UploadHandler 上传控制器
type UploadHandler struct{}

// PreUploadFile 文件预上传
func (uh UploadHandler) PreUploadFile(c *gin.Context) {
	// md5,fileName,fileSize,cid
	// 检查MD5是否已存在
	// 检查名称是否合规安全
	// 检查文件大小是否超过限制
	// 检查文件类型是否符合要求

	// 如果文件存在，复制到cid对应目录

}

// UploadFile 上传文件
func (uh UploadHandler) UploadFile(c *gin.Context) {
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
