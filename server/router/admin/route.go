package admin

import (
	"x_admin/admin/generator"

	"x_admin/controller/admin/commonController"
	"x_admin/controller/admin/monitorController"
	"x_admin/controller/admin/settingController"
	"x_admin/controller/admin/systemController"

	"github.com/gin-gonic/gin"
)

func RegisterRoute(rg *gin.RouterGroup) {

	rg = rg.Group("/admin")
	// 所有子路由需要加上前缀 /api/admin

	commonController.UploadRoute(rg)
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

	FlowTemplateRoute(rg)
	FlowApplyRoute(rg)
	FlowHistoryRoute(rg)

	generator.RegisterGroup(rg)

	MonitorProjectRoute(rg)
	MonitorClientRoute(rg)
	MonitorErrorRoute(rg)

	UserProtocolRoute(rg)
}
