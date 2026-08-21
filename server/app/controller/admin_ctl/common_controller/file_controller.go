package common_controller

import (
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"x_admin/app/service/common_service"
	"x_admin/config"
	"x_admin/core"
	"x_admin/core/response"

	"github.com/gin-gonic/gin"
)

// FileHandler 文件流控制器
type FileHandler struct{}

// Serve 按文件哈希ID返回文件流
// @Summary		按文件哈希ID获取文件流
// @Description	通过 id 查询 x_common_file_hash.FilePath，读取物理文件并以流形式返回
// @Tags			common_file-文件
// @Param			id	path		string	true	"文件哈希ID"
// @Success		200	{file}		binary	"文件流"
// @Router			/api/uploads/{id} [get]
func (fh FileHandler) Serve(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	// 通过 id 查询磁盘存储 key（含扩展名）
	filePath := common_service.FileHashService.GetFilePath(id)
	if filePath == "" {
		response.NotFound(c, "文件不存在")
		return
	}

	absPath := filepath.Join(config.FileConfig.UploadDirectory, filePath)
	f, err := os.Open(absPath)
	if err != nil {
		core.Logger.Errorf("FileHandler.Serve open err: id=%s path=%s err=%+v", id, absPath, err)
		response.NotFound(c, "文件不存在")
		return
	}
	defer f.Close()

	ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(filePath), "."))
	ctype := mime.TypeByExtension("." + ext)
	if ctype == "" {
		ctype = "application/octet-stream"
	}
	c.Header("Content-Type", ctype)
	c.Header("Cache-Control", "public, max-age=31536000, immutable")
	http.ServeContent(c.Writer, c.Request, filepath.Base(absPath), time.Time{}, f)
}
