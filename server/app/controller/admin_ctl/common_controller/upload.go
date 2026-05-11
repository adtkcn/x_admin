package common_controller

import (
	"x_admin/app/schema/common_schema"
	"x_admin/app/service/common_service"
	"x_admin/config"
	"x_admin/core/response"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

// UploadHandler 上传控制器
type UploadHandler struct{}

// @Summary		文件预上传
// @Description	文件预上传检查
// @Tags			common_upload-上传
// @Param			token		header		string						true	"token"
// @Param			md5			body		string						true	"文件MD5"
// @Param			fileName	body		string						true	"文件名"
// @Param			fileSize	body		int64						true	"文件大小"
// @Param			cid			body		string						false	"分类ID"
// @Success		200			{object}	response.Response			"成功"
// @Router			/api/admin/common/upload/preUpload [post]
func (uh UploadHandler) PreUploadFile(c *gin.Context) {
	// md5,fileName,fileSize,cid
	// 检查MD5是否已存在
	// 检查名称是否合规安全
	// 检查文件大小是否超过限制
	// 检查文件类型是否符合要求

	// 如果文件存在，复制到cid对应目录

}

// @Summary		上传文件
// @Description	上传文件
// @Tags			common_upload-上传
// @Param			token	header		string						true	"token"
// @Param			cid		body		string						false	"分类ID"
// @Param			file	formData	file						true	"文件"
// @Success		200		{object}	response.Response{data=common_schema.CommonUploadFileResp}	"成功"
// @Router			/api/admin/common/upload/upload [post]
func (uh UploadHandler) UploadFile(c *gin.Context) {
	var uReq common_schema.CommonUploadImageReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &uReq)) {
		return
	}
	file, ve := util.VerifyUtil.VerifyFile(c, "file")
	if response.IsFailWithResp(c, ve) {
		return
	}
	res, err := common_service.UploadService.UploadFile(file, uReq.Cid, config.AdminConfig.GetAdminId(c))
	response.CheckAndRespWithData(c, res, err)
}
