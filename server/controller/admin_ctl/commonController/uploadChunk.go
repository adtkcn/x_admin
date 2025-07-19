package commonController

import (
	"fmt"
	"net/url"
	"os"
	"regexp"
	"x_admin/core/response"
	"x_admin/service/commonService"

	"github.com/gin-gonic/gin"
)

func UploadChunkRoute(rg *gin.RouterGroup) {
	handle := uploadChunkHandler{
		uploadPath: "./uploads",
		tmpPath:    "./tmp",
	}
	os.MkdirAll(handle.uploadPath, 0755)
	os.MkdirAll(handle.tmpPath, 0755)

	rg = rg.Group("/common")
	rg.GET("/uploadChunk/CheckFileExist", handle.CheckFileExist)
	rg.GET("/uploadChunk/CheckChunkExist", handle.CheckChunkExist)
	rg.POST("/uploadChunk/UploadChunk", handle.UploadChunk)
	rg.POST("/uploadChunk/MergeChunk", handle.MergeChunk)
}

type uploadChunkHandler struct {
	uploadPath string
	tmpPath    string
}

func (uh uploadChunkHandler) CheckFileExist(c *gin.Context) {
	var fileMd5 = c.Query("fileMd5")
	var fileName = url.QueryEscape(c.Query("fileName"))
	// 正则检查MD5
	reg := regexp.MustCompile(`^[a-fA-F0-9]{32}$`)
	if !reg.MatchString(fileMd5) {
		response.FailWithMsg(c, response.SystemError, "MD5格式错误")
		return
	}
	var filePath = fmt.Sprintf("%s/%s_%s", uh.uploadPath, fileMd5, fileName)
	// 检查文件是否存在
	if commonService.UploadChunkService.CheckFileExist(filePath) {
		response.OkWithData(c, filePath)
		return
	}
	response.OkWithData(c, nil)
}

// 检查chunk是否存在
func (uh uploadChunkHandler) CheckChunkExist(c *gin.Context) {
	fileMd5 := c.Query("fileMd5") // 上传文件的md5
	index := c.Query("index")     // 分片序号
	chunkPath := fmt.Sprintf("%s/%s/%s", uh.tmpPath, fileMd5, index)
	if commonService.UploadChunkService.CheckFileExist(chunkPath) {
		response.Ok(c)
		return
	}
	response.FailWithMsg(c, response.SystemError, "分片不存在")
}

// UploadChunk 上传分片
func (uh uploadChunkHandler) UploadChunk(c *gin.Context) {
	chunk, _ := c.FormFile("chunk")  // 分片文件
	index := c.PostForm("index")     // 分片序号
	fileMd5 := c.PostForm("fileMd5") // 上传文件的md5

	chunkDir := fmt.Sprintf("%s/%s", uh.tmpPath, fileMd5)
	chunkPath := fmt.Sprintf("%s/%s", chunkDir, index)
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
	}
	bindErr := c.ShouldBindJSON(&MergeChunk)
	if bindErr != nil {
		response.FailWithMsg(c, response.SystemError, bindErr.Error())
		return
	}
	var filePath = fmt.Sprintf("%s/%s_%s", uh.uploadPath, MergeChunk.FileMd5, url.QueryEscape(MergeChunk.FileName))
	err := commonService.UploadChunkService.MergeChunk(MergeChunk.FileMd5, filePath, MergeChunk.ChunkCount)
	if err != nil {
		response.FailWithMsg(c, response.SystemError, err.Error())
		return
	}
	response.OkWithData(c, filePath)
}
