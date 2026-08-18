package common_controller

import (
	"x_admin/app/schema/common_schema"
	"x_admin/app/service/common_service"
	"x_admin/core"
	"x_admin/core/response"
	"x_admin/plugin/storage"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

// UploadHandler 上传控制器
type UploadHandler struct{}

// @Summary		上传文件
// @Description	上传文件
// @Tags			common_upload-上传
// @Param			token	header		string														true	"token"
// @Param			cid		body		string														false	"分类ID"
// @Param			file	formData	file														true	"文件"
// @Success		200		{object}	response.Response{data=common_schema.CommonUploadFileResp}	"成功"
// @Router			/api/admin/common/upload/file [post]
func (uh UploadHandler) UploadFile(c *gin.Context) {
	var uReq common_schema.CommonUploadImageReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &uReq)) {
		return
	}
	file, ve := util.VerifyUtil.VerifyFile(c, "file")
	if response.IsFailWithResp(c, ve) {
		return
	}
	res, err := common_service.UploadService.UploadFile(file)
	// 组装与相册列表一致的 url/path 字段，便于前端直接展示已上传文件
	// 本地引擎下 GetObjectURL(filePath) = /api/uploads/<filePath>，与相册 row.uri 完全对齐
	engine := storage.GetStorageEngine()
	url, _ := engine.GetObjectURL(res.FilePath)
	resp := common_schema.CommonUploadFileResp{
		ID:         res.ID,
		FileHashId: res.ID,
		Name:       file.Filename,
		Uri:        url,          // 访问地址（完整可访问 URL）
		Path:       res.FilePath, // 相对路径
		Ext:        res.Ext,
		Size:       res.FileSize,
		Instant:    false,
	}
	response.CheckAndRespWithData(c, resp, err)
}

// @Summary		文件秒传检查
// @Description	根据文件 MD5 查询是否已上传，命中则返回已有的文件哈希记录ID
// @Tags			common_upload-上传
// @Param			token		header		string				true	"token"
// @Param			file_md5	body		string				true	"文件MD5"
// @Param			file_name	body		string				true	"文件名"
// @Success		200			{object}	response.Response	"成功"
// @Router			/api/admin/common/upload/checkInstant [post]
func (uh UploadHandler) CheckInstant(c *gin.Context) {
	var req struct {
		FileMd5  string `json:"file_md5" binding:"required"`
		FileName string `json:"file_name" binding:"required"`
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &req)) {
		return
	}
	record, err := common_service.FileHashService.FindByMd5(req.FileMd5)
	if err != nil {
		core.Logger.Errorf("CheckInstant err: %v", err)
	}
	if record != nil {
		engine := storage.GetStorageEngine()
		url, _ := engine.GetObjectURL(record.FilePath)
		resp := common_schema.CommonUploadFileResp{
			ID:         record.ID,
			FileHashId: record.ID,
			Name:       req.FileName,
			Uri:        url,             // 访问地址（完整可访问 URL）
			Path:       record.FilePath, // 相对路径
			Ext:        record.Ext,
			Size:       record.FileSize,
			Instant:    true,
		}
		response.CheckAndRespWithData(c, resp, nil)
		return
	}

	// 秒传未命中：返回统一结构（instant=false），上传流程继续
	response.CheckAndRespWithData(c, common_schema.CommonUploadFileResp{Instant: false}, nil)
}
