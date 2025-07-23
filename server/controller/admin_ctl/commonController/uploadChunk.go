package commonController

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"x_admin/core/response"
	"x_admin/service/commonService"

	"github.com/gin-gonic/gin"
)

func UploadChunkRoute(rg *gin.RouterGroup) {
	handle := uploadChunkHandler{
		uploadPath: "./uploads",
		tmpPath:    "./uploads/.tmp",
	}
	os.MkdirAll(handle.uploadPath, 0755)
	os.MkdirAll(handle.tmpPath, 0755)

	rg = rg.Group("/common")
	rg.GET("/uploadChunk/CheckFileExist", handle.CheckFileExist)
	// rg.GET("/uploadChunk/CheckChunkExist", handle.CheckChunkExist)
	rg.GET("/uploadChunk/HasChunk", handle.HasChunk)

	rg.POST("/uploadChunk/UploadChunk", handle.UploadChunk)
	rg.POST("/uploadChunk/MergeChunk", handle.MergeChunk)
}

type uploadChunkHandler struct {
	uploadPath string
	tmpPath    string
}

func (uh uploadChunkHandler) getFilePath(fileMd5 string) string {
	return fmt.Sprintf("%s/%s", uh.uploadPath, fileMd5)
}
func (uh uploadChunkHandler) getChunkDir(fileMd5 string, chunkSize string) string {
	return fmt.Sprintf("%s/%s_%s", uh.tmpPath, fileMd5, chunkSize)
}
func (uh uploadChunkHandler) getChunkPath(fileMd5 string, chunkSize string, index string) string {
	return fmt.Sprintf("%s/%s_%s/%s", uh.tmpPath, fileMd5, chunkSize, index)
}

func (uh uploadChunkHandler) CheckFileExist(c *gin.Context) {
	var fileMd5 = c.Query("fileMd5")
	var fileName = c.Query("fileName")
	if fileMd5 == "" {
		response.FailWithMsg(c, response.SystemError, "文件hash错误")
		return
	}
	if fileName == "" {
		response.FailWithMsg(c, response.SystemError, "文件名错误")
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
	}
	var fileExt = ""
	for _, ext := range whiteList {
		if strings.HasSuffix(fileName, ext) {
			fileExt = ext
			break
		}
	}
	if fileExt == "" {
		response.FailWithMsg(c, response.SystemError, "文件类型错误")
		return
	}

	// var fileName = url.QueryEscape(c.Query("fileName"))
	// 正则检查MD5
	reg := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	if !reg.MatchString(fileMd5) {
		response.FailWithMsg(c, response.SystemError, "文件hash错误")
		return
	}
	var filePath = uh.getFilePath(fileMd5)
	// 检查文件是否存在
	if commonService.UploadChunkService.CheckFileExist(filePath) {
		response.OkWithData(c, filePath)
		return
	}
	response.OkWithData(c, nil)
}
func (uh uploadChunkHandler) HasChunk(c *gin.Context) {
	var fileMd5 = c.Query("fileMd5")
	var chunkSize = c.Query("chunkSize")
	if fileMd5 == "" {
		response.FailWithMsg(c, response.SystemError, "文件hash错误")
		return
	}
	if chunkSize == "" {
		response.FailWithMsg(c, response.SystemError, "分片大小错误")
		return
	}
	var chunkDir = uh.getChunkDir(fileMd5, chunkSize)
	var HasChunk = commonService.UploadChunkService.HasChunk(chunkDir)
	response.OkWithData(c, HasChunk)
}

// 检查chunk是否存在
// func (uh uploadChunkHandler) CheckChunkExist(c *gin.Context) {
// 	fileMd5 := c.Query("fileMd5") // 上传文件的md5
// 	chunkSize := c.Query("chunkSize")
// 	index := c.Query("index") // 分片序号
// 	// chunkPath := fmt.Sprintf("%s/%s/%s", uh.tmpPath, fileMd5, index)
// 	chunkPath := uh.getChunkPath(fileMd5, chunkSize, index)
// 	if commonService.UploadChunkService.CheckFileExist(chunkPath) {
// 		response.OkWithData(c, 1)
// 		return
// 	}
// 	response.OkWithData(c, 0)
// }

// UploadChunk 上传分片
func (uh uploadChunkHandler) UploadChunk(c *gin.Context) {
	chunk, _ := c.FormFile("chunk")      // 分片文件
	chunkSize := c.PostForm("chunkSize") // 分片分割的大小
	index := c.PostForm("index")         // 分片序号
	fileMd5 := c.PostForm("fileMd5")     // 上传文件的md5
	if fileMd5 == "" {
		response.FailWithMsg(c, response.SystemError, "文件hash错误")
		return
	}
	if chunkSize == "" {
		response.FailWithMsg(c, response.SystemError, "分片大小错误")
		return
	}
	if index == "" {
		response.FailWithMsg(c, response.SystemError, "分片序号错误")
		return
	}
	if chunk == nil {
		response.FailWithMsg(c, response.SystemError, "分片文件错误")
		return
	}

	chunkDir := uh.getChunkDir(fileMd5, chunkSize)
	chunkPath := uh.getChunkPath(fileMd5, chunkSize, index)
	err := commonService.UploadChunkService.UploadChunk(chunkDir, chunkPath, chunk)
	if err != nil {
		response.FailWithMsg(c, response.SystemError, err.Error())
		return
	}
	response.Ok(c)
}
func (uh uploadChunkHandler) MergeChunk(c *gin.Context) {
	var MergeChunk struct {
		FileMd5    string `json:"fileMd5"`    // 上传文件的md5
		FileName   string `json:"fileName"`   // 文件名
		ChunkCount int    `json:"chunkCount"` // 分片数量
		ChunkSize  int    `json:"chunkSize"`  // 分片分割的大小,作用：确保不同分片大小不放在同一目录
	}
	bindErr := c.ShouldBindJSON(&MergeChunk)
	if bindErr != nil {
		response.FailWithMsg(c, response.SystemError, bindErr.Error())
		return
	}
	if MergeChunk.FileMd5 == "" {
		response.FailWithMsg(c, response.SystemError, "文件hash错误")
		return
	}
	if MergeChunk.FileName == "" {
		response.FailWithMsg(c, response.SystemError, "文件名错误")
		return
	}
	if MergeChunk.ChunkCount <= 0 {
		response.FailWithMsg(c, response.SystemError, "分片数量错误")
		return
	}
	if MergeChunk.ChunkSize <= 0 {
		response.FailWithMsg(c, response.SystemError, "分片大小错误")
		return
	}
	var filePath = uh.getFilePath(MergeChunk.FileMd5)
	var chunkDir = uh.getChunkDir(MergeChunk.FileMd5, fmt.Sprintf("%d", MergeChunk.ChunkSize))
	err := commonService.UploadChunkService.MergeChunk(chunkDir, filePath, MergeChunk.ChunkCount)
	if err != nil {
		response.FailWithMsg(c, response.SystemError, err.Error())
		return
	}
	response.OkWithData(c, filePath)
}
