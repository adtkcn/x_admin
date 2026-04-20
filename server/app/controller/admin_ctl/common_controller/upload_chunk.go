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
	// 获取文件后缀
	ext := filepath.Ext(fileName)

	return fmt.Sprintf("%s/%s%s", uh.UploadPath, fileMd5, ext)
}
func (uh UploadChunkHandler) GetChunkDir(fileMd5 string, chunkSize string) string {
	return fmt.Sprintf("%s/%s_%s", uh.TmpPath, fileMd5, chunkSize)
}
func (uh UploadChunkHandler) GetChunkPath(fileMd5 string, chunkSize string, index string) string {
	return fmt.Sprintf("%s/%s_%s/%s", uh.TmpPath, fileMd5, chunkSize, index)
}

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
	//文件类型白名单
	var whiteList = []string{
		".jpg",
		".jpeg",
		".png",
		".gif",
		".bmp",
		".svg",
		".webp",
		".avif",
		".txt",
		".mp4",
		".avi",
		".mov",
		".wmv",
		".flv",
		".mkv",
		".mp3",
		".wav",
		".aac",
		".flac",
		".ogg",
		".m4a",
		".pdf",
		".doc",
		".docx",
		".xls",
		".xlsx",
		".ppt",
		".pptx",
		".zip",
		".rar",
		".7z",
		".tar",
		".gz",
		".bz2",
		".xz",
		".msi",
		".exe",
		".dmg",
		".iso",
		".app",
		".deb",
		".rpm",
		".pkg",
		".apk",
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

	// var fileName = url.QueryEscape(c.Query("fileName"))
	// 正则检查MD5
	reg := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	if !reg.MatchString(fileMd5) {
		response.Fail(c, "文件hash错误")
		return
	}
	var filePath = uh.GetFilePath(fileMd5, fileName)
	// 检查文件是否存在
	if common_service.UploadChunkService.CheckFileExist(filePath) {
		response.Ok(c, filePath)
		return
	}
	response.Ok(c, nil)
}
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

// UploadChunk 上传分片
func (uh UploadChunkHandler) UploadChunk(c *gin.Context) {
	chunk, _ := c.FormFile("chunk")      // 分片文件
	chunkSize := c.PostForm("chunkSize") // 分片分割的大小
	index := c.PostForm("index")         // 分片序号
	fileMd5 := c.PostForm("fileMd5")     // 上传文件的md5
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
func (uh UploadChunkHandler) MergeChunk(c *gin.Context) {
	var mergeReq struct {
		FileMd5    string `json:"fileMd5"`    // 上传文件的md5
		FileName   string `json:"fileName"`   // 文件名
		ChunkCount int    `json:"chunkCount"` // 分片数量
		ChunkSize  int    `json:"chunkSize"`  // 分片分割的大小,作用：确保不同分片大小不放在同一目录
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
