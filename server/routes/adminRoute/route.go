package adminRoute

import (
	"x_admin/app/controller/admin_ctl/commonController"
	"x_admin/app/controller/admin_ctl/generatorController"
	"x_admin/app/controller/admin_ctl/monitorController"
	"x_admin/app/controller/admin_ctl/settingController"
	"x_admin/app/controller/admin_ctl/systemController"

	"github.com/gin-gonic/gin"
)

type RouteHandlerFunc func(*gin.RouterGroup)

// routeHandlers 全局的路由注册函数切片，用于自动加载路由，通过每个文件的init()收集
var routeHandlers []RouteHandlerFunc

// Autoload 自动加载所有路由
func Autoload(rg *gin.RouterGroup) {
	for _, handler := range routeHandlers {
		handler(rg)
	}
}

func RegisterRoute(rg *gin.RouterGroup) {

	rg = rg.Group("/admin")
	// 所有子路由需要加上前缀 /api/admin

	commonController.UploadRoute(rg)
	commonController.UploadChunkRoute(rg)
	commonController.AlbumRoute(rg)
	commonController.IndexRoute(rg)
	commonController.GeTuiRoute(rg) //个推

	monitorController.RegisterRoute(rg)

	settingController.CopyrightRoute(rg)
	settingController.DictDataRoute(rg)
	settingController.DictTypeRoute(rg)
	settingController.WebsiteRoute(rg)

	systemController.LoginRoute(rg)
	systemController.AdminRoute(rg)
	systemController.MenuRoute(rg)
	systemController.PostRoute(rg)

	systemController.DeptRoute(rg)
	systemController.RoleRoute(rg)
	systemController.LogRoute(rg)
	generatorController.GenRoute(rg)
	// FlowTemplateRoute(rg)
	// FlowApplyRoute(rg)
	// FlowHistoryRoute(rg)

	// MonitorProjectRoute(rg)
	// MonitorClientRoute(rg)
	// MonitorErrorRoute(rg)

	// UserProtocolRoute(rg)

	Autoload(rg)
}
