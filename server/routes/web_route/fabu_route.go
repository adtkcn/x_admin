package web_route

import (
	"x_admin/app/controller/web_ctl"

	"github.com/gin-gonic/gin"
)

// FabuWebRoute 公开下载/安装路由（无需登录）
func FabuWebRoute(rg *gin.RouterGroup) {
	handle := web_ctl.FabuController{}
	fabuRg := rg.Group("/fabu")
	{
		fabuRg.GET("/plist/:appId/:versionId", handle.Plist)
		fabuRg.GET("/app/:shortUrl", handle.AppInfo)
		fabuRg.GET("/download/:shortUrl", handle.Download)
		fabuRg.GET("/count/:appId/:versionId", handle.Count)
		fabuRg.GET("/version/checkupdate", handle.CheckUpdate)
	}
}

func init() {
	webRouteHandlers = append(webRouteHandlers, FabuWebRoute)
}
