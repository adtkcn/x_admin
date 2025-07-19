package adminRoute

import (
	"x_admin/controller/admin_ctl/commonController"
	"x_admin/controller/admin_ctl/generatorController"
	"x_admin/controller/admin_ctl/monitorController"
	"x_admin/controller/admin_ctl/settingController"
	"x_admin/controller/admin_ctl/systemController"

	"github.com/gin-gonic/gin"
)

func RegisterRoute(rg *gin.RouterGroup) {

	rg = rg.Group("/admin")
	// 所有子路由需要加上前缀 /api/admin

	commonController.UploadRoute(rg)
	commonController.UploadChunkRoute(rg)
	commonController.AlbumRoute(rg)
	commonController.IndexRoute(rg)

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
	FlowTemplateRoute(rg)
	FlowApplyRoute(rg)
	FlowHistoryRoute(rg)

	MonitorProjectRoute(rg)
	MonitorClientRoute(rg)
	MonitorErrorRoute(rg)

	UserProtocolRoute(rg)
}
