package common_controller

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"x_admin/app/service/common_service"
	"x_admin/core/response"

	"github.com/gin-gonic/gin"
)

// UploadChunkHandler 分片上传控制器
type UploadChunkHandler struct {
	UploadPath string
	TmpPath    string
}

// NewUploadChunkHandler 创建分片上传处理器（初始化目录）
func NewUploadChunkHandler(uploadPath, tmpPath string) UploadChunkHandler {
	os.MkdirAll(uploadPath, 0755)
	os.MkdirAll(tmpPath, 0755)
	return UploadChunkHandler{UploadPath: uploadPath, TmpPath: tmpPath}
}

func (uh UploadChunkHandler) GetFilePath(fileMd5 string, fileName string) string {
	ext := filepath.Ext(fileName)
	return fmt.Sprintf("%s/%s%s", uh.UploadPath, fileMd5, ext)
}
func (uh UploadChunkHandler) GetChunkDir(fileMd5 string, chunkSize string) string {
	return fmt.Sprintf("%s/%s_%s", uh.TmpPath, fileMd5, chunkSize)
}
func (uh UploadChunkHandler) GetChunkPath(fileMd5 string, chunkSize string, index string) string {
	return fmt.Sprintf("%s/%s_%s/%s", uh.TmpPath, fileMd5, chunkSize, index)
}

// @Summary		检查文件是否存在
// @Description	检查文件是否已存在
// @Tags			common_uploadChunk-分片上传
// @Param			token		header		string						true	"token"
// @Param			fileMd5		query		string						true	"文件MD5"
// @Param			fileName	query		string						true	"文件名"
// @Success		200			{object}	response.Response			"成功"
// @Router			/api/admin/common/uploadChunk/checkFileExist [get]
func (uh UploadChunkHandler) CheckFileExist(c *gin.Context) {
	var fileMd5 = c.Query("fileMd5")
	var fileName = c.Query("fileName")
	if fileMd5 == "" {
		response.Fail(c, "文件hash错误")
		return
	}
	if fileName == "" {
		response.Fail(c, "文件名错误")
		return
	}
	var whiteList = []string{
		".jpg", ".jpeg", ".png", ".gif", ".bmp", ".svg", ".webp", ".avif",
		".txt", ".mp4", ".avi", ".mov", ".wmv", ".flv", ".mkv",
		".mp3", ".wav", ".aac", ".flac", ".ogg", ".m4a",
		".pdf", ".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx",
		".zip", ".rar", ".7z", ".tar", ".gz", ".bz2", ".xz",
		".msi", ".exe", ".dmg", ".iso", ".app", ".deb", ".rpm", ".pkg", ".apk",
	}
	var fileExt = ""
	for _, ext := range whiteList {
		if strings.HasSuffix(fileName, ext) {
			fileExt = ext
			break
		}
	}
	if fileExt == "" {
		response.Fail(c, "文件类型错误")
		return
	}

	reg := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	if !reg.MatchString(fileMd5) {
		response.Fail(c, "文件hash错误")
		return
	}
	var filePath = uh.GetFilePath(fileMd5, fileName)
	if common_service.UploadChunkService.CheckFileExist(filePath) {
		response.Ok(c, filePath)
		return
	}
	response.Ok(c, nil)
}

// @Summary		检查已上传分片
// @Description	检查已上传的分片列表
// @Tags			common_uploadChunk-分片上传
// @Param			token		header		string						true	"token"
// @Param			fileMd5		query		string						true	"文件MD5"
// @Param			chunkSize	query		string						true	"分片大小"
// @Success		200			{object}	response.Response			"成功"
// @Router			/api/admin/common/uploadChunk/hasChunk [get]
func (uh UploadChunkHandler) HasChunk(c *gin.Context) {
	var fileMd5 = c.Query("fileMd5")
	var chunkSize = c.Query("chunkSize")
	if fileMd5 == "" {
		response.Fail(c, "文件hash错误")
		return
	}
	if chunkSize == "" {
		response.Fail(c, "分片大小错误")
		return
	}
	var chunkDir = uh.GetChunkDir(fileMd5, chunkSize)
	var HasChunk = common_service.UploadChunkService.HasChunk(chunkDir)
	response.Ok(c, HasChunk)
}

// @Summary		上传分片
// @Description	上传文件分片
// @Tags			common_uploadChunk-分片上传
// @Param			token		header		string						true	"token"
// @Param			chunk		formData	file						true	"分片文件"
// @Param			chunkSize	formData	string						true	"分片大小"
// @Param			index		formData	string						true	"分片序号"
// @Param			fileMd5		formData	string						true	"文件MD5"
// @Success		200			{object}	response.Response			"成功"
// @Router			/api/admin/common/uploadChunk/upload [post]
func (uh UploadChunkHandler) UploadChunk(c *gin.Context) {
	chunk, _ := c.FormFile("chunk")
	chunkSize := c.PostForm("chunkSize")
	index := c.PostForm("index")
	fileMd5 := c.PostForm("fileMd5")
	if fileMd5 == "" {
		response.Fail(c, "文件hash错误")
		return
	}
	if chunkSize == "" {
		response.Fail(c, "分片大小错误")
		return
	}
	if index == "" {
		response.Fail(c, "分片序号错误")
		return
	}
	if chunk == nil {
		response.Fail(c, "分片文件错误")
		return
	}

	chunkDir := uh.GetChunkDir(fileMd5, chunkSize)
	chunkPath := uh.GetChunkPath(fileMd5, chunkSize, index)
	err := common_service.UploadChunkService.UploadChunk(chunkDir, chunkPath, chunk)
	if err != nil {
		response.Fail(c, err.Error())
		return
	}
	response.Ok(c)
}

// @Summary		合并分片
// @Description	合并所有分片为完整文件
// @Tags			common_uploadChunk-分片上传
// @Param			token		header		string						true	"token"
// @Param			fileMd5		body		string						true	"文件MD5"
// @Param			fileName	body		string						true	"文件名"
// @Param			chunkCount	body		int							true	"分片数量"
// @Param			chunkSize	body		int							true	"分片大小"
// @Success		200			{object}	response.Response			"成功"
// @Router			/api/admin/common/uploadChunk/merge [post]
func (uh UploadChunkHandler) MergeChunk(c *gin.Context) {
	var mergeReq struct {
		FileMd5    string `json:"fileMd5"`
		FileName   string `json:"fileName"`
		ChunkCount int    `json:"chunkCount"`
		ChunkSize  int    `json:"chunkSize"`
	}
	bindErr := c.ShouldBindJSON(&mergeReq)
	if bindErr != nil {
		response.Fail(c, bindErr.Error())
		return
	}
	if mergeReq.FileMd5 == "" {
		response.Fail(c, "文件hash错误")
		return
	}
	if mergeReq.FileName == "" {
		response.Fail(c, "文件名错误")
		return
	}
	if mergeReq.ChunkCount <= 0 {
		response.Fail(c, "分片数量错误")
		return
	}
	if mergeReq.ChunkSize <= 0 {
		response.Fail(c, "分片大小错误")
		return
	}
	var filePath = uh.GetFilePath(mergeReq.FileMd5, mergeReq.FileName)
	var chunkDir = uh.GetChunkDir(mergeReq.FileMd5, fmt.Sprintf("%d", mergeReq.ChunkSize))
	err := common_service.UploadChunkService.MergeChunk(chunkDir, filePath, mergeReq.ChunkCount)
	if err != nil {
		response.Fail(c, err.Error())
		return
	}
	response.Ok(c, filePath)
}
