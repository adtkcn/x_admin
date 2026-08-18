package util

import (
	"net/url"
	"path"
	"strings"
	"x_admin/config"
	"x_admin/core"
)

var (
	UrlUtil = urlUtil{}

	uploadPrefix = config.FileConfig.UploadPrefix //"/api/uploads"
)

// urlUtil 文件路径处理工具
type urlUtil struct{}

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

// 去掉publicPrefix前缀
func (uu urlUtil) ToRelativeUrl(u string) string {
	// TODO: engine默认local
	if u == "" {
		return ""
	}
	up, err := url.Parse(u)
	if err != nil {
		core.Logger.Errorf("ToRelativeUrl Parse err: err=[%+v]", err)
		return u
	}
	engine := "local"
	if engine == "local" {
		lu := up.String()
		return strings.Replace(lu, uploadPrefix, "", 1)
	}
	// TODO: 其他engine
	return u
}
