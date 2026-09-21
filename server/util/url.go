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
func (uu urlUtil) HashUrl(id string, file_name string) string {
	if id == "" {
		return ""
	}
	return path.Join(uploadPrefix, id, file_name)
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

// 生成文件存储路径 年月日/时/uuid.ext
func (uu urlUtil) BuildFileSavePath(fileName string) string {
	ext := strings.ToLower(path.Ext(fileName))
	now := time.Now()
	id := ToolsUtil.MakeUuidV7()
	// 取id后两位作为目录，分散目录锁争抢
	idHash := id[len(id)-2:]
	datePath := path.Join(now.Format("20060102"), now.Format("15"), idHash)
	return path.Join(datePath, id+ext)
}

// ReplaceExt 将 path 的扩展名替换为 newExt（保留目录与文件名主体）。
func (uu urlUtil) ReplaceExt(filePath, newExt string) string {
	ext := path.Ext(filePath)
	base := filePath
	if ext != "" {
		base = strings.TrimSuffix(filePath, ext)
	}
	return base + "." + newExt
}
