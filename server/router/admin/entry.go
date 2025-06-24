package admin

import (
	"x_admin/admin/common/album"
	"x_admin/admin/common/index"
	"x_admin/admin/common/upload"
	"x_admin/admin/flow"
	"x_admin/admin/generator"
	"x_admin/admin/monitor"

	"x_admin/controller/admin/settingController"
	"x_admin/controller/admin/systemController"

	"github.com/gin-gonic/gin"
)

func RegisterGroup(rg *gin.RouterGroup) {

	rg = rg.Group("/admin")
	// 所有子路由需要加上前缀 /api/admin

	upload.UploadRoute(rg)
	album.AlbumRoute(rg)
	index.IndexRoute(rg)

	monitor.MonitorRoute(rg)

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

	flow.FlowTemplateRoute(rg)
	flow.FlowApplyRoute(rg)
	flow.FlowHistoryRoute(rg)

	generator.RegisterGroup(rg)

	MonitorProjectRoute(rg)
	MonitorClientRoute(rg)
	MonitorErrorRoute(rg)

	UserProtocolRoute(rg)
}
