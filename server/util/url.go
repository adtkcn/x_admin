package util

import (
	"path"
	"strings"
	"time"
	"x_admin/config"
)

var (
	UrlUtil = urlUtil{}

	uploadPrefix = config.FileConfig.UploadPrefix //"/api/uploads"
)

// urlUtil 文件路径处理工具
type urlUtil struct{}

// HashUrl 由文件哈希ID拼访问URL：/api/uploads/<id>
// 实际文件流由路由 GET /api/uploads/:id 按 id 查 x_common_file_hash.FilePath 返回，无需扩展名。
func (uu urlUtil) HashUrl(id string) string {
	if id == "" {
		return ""
	}
	return path.Join(uploadPrefix, id)
}

// ToAbsoluteUrl 转绝对路径
func (uu urlUtil) ToAbsoluteUrl(u string) string {
	// TODO: engine默认local
	if u == "" {
		return ""
	}

	// 处理/api/static/开头的路径
	if strings.HasPrefix(u, "/api/static/") {
		return u
	}
	engine := "local"
	if engine == "local" {
		return path.Join(uploadPrefix, u)
	}
	// TODO: 其他engine
	return u
}

// GetFileExt 获取文件扩展名
func (uu urlUtil) GetFileExt(fileName string) string {
	fileExt := strings.ToLower(strings.Replace(path.Ext(fileName), ".", "", 1))
	return fileExt
}

// 生成文件存储路径
func (uu urlUtil) BuildFileSavePath(fileName string) string {
	ext := strings.ToLower(path.Ext(fileName))
	now := time.Now()
	// 年月日/时/分
	datePath := path.Join(now.Format("20060102"), now.Format("15"), now.Format("04"))
	return path.Join(datePath, ToolsUtil.MakeUuidV7()+ext)
}
