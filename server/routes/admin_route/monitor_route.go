package admin_route

import (
	"x_admin/app/controller/admin_ctl/monitor_controller"
	"x_admin/app/middleware"

	"github.com/gin-gonic/gin"
)

// MonitorRoute 系统监控模块路由（服务端监控、缓存监控 + 客户端/错误/项目子路由）
func MonitorRoute(rg *gin.RouterGroup) {

	// 服务端监控 + 缓存监控
	handleMonitor := monitor_controller.MonitorHandler{}
	monitorRg := rg.Group("/monitor", middleware.PermAuth())
	monitorRg.GET("/cache", handleMonitor.Cache)
	monitorRg.GET("/server", handleMonitor.Server)

}
func MonitorClientRoute(rg *gin.RouterGroup) {
	handle := monitor_controller.MonitorClientHandler{}
	rg.GET("/monitor_client/add", middleware.RecordLog("监控-客户端信息新增"), handle.Add)

	r := rg.Group("/", middleware.PermAuth())
	r.GET("/monitor_client/list", handle.List)
	r.GET("/monitor_client/listAll", handle.ListAll)
	r.GET("/monitor_client/detail", handle.Detail)
	r.GET("/monitor_client/errorUsers", handle.ErrorUsers)

	// r.POST("/monitor_client/edit",middleware.RecordLog("监控-客户端信息编辑"), handle.Edit)

	r.POST("/monitor_client/del", middleware.RecordLog("监控-客户端信息删除"), handle.Del)
	r.POST("/monitor_client/delBatch", middleware.RecordLog("监控-客户端信息删除-批量"), handle.DelBatch)

	r.GET("/monitor_client/exportFile", middleware.RecordLog("监控-客户端信息导出"), handle.ExportFile)
	r.POST("/monitor_client/importFile", handle.ImportFile)
}

func MonitorErrorRoute(rg *gin.RouterGroup) {
	handle := monitor_controller.MonitorErrorHandler{}
	rg.GET("/monitor_error/add", handle.Add)

	r := rg.Group("/", middleware.PermAuth())
	r.GET("/monitor_error/list", handle.List)
	r.GET("/monitor_error/listAll", handle.ListAll)
	r.GET("/monitor_error/detail", handle.Detail)

	r.POST("/monitor_error/del", middleware.RecordLog("监控-错误列删除"), handle.Del)
	r.POST("/monitor_error/delBatch", middleware.RecordLog("监控-错误列删除-批量"), handle.DelBatch)

	r.GET("/monitor_error/exportFile", middleware.RecordLog("监控-错误列导出"), handle.ExportFile)
	r.POST("/monitor_error/importFile", handle.ImportFile)
}

func MonitorProjectRoute(rg *gin.RouterGroup) {
	handle := monitor_controller.MonitorProjectHandler{}

	r := rg.Group("/", middleware.PermAuth())
	r.GET("/monitor_project/list", handle.List)
	r.GET("/monitor_project/listAll", handle.ListAll)
	r.GET("/monitor_project/detail", handle.Detail)

	r.POST("/monitor_project/add", middleware.RecordLog("监控项目新增"), handle.Add)
	r.POST("/monitor_project/edit", middleware.RecordLog("监控项目编辑"), handle.Edit)

	r.POST("/monitor_project/del", middleware.RecordLog("监控项目删除"), handle.Del)
	r.POST("/monitor_project/delBatch", middleware.RecordLog("监控项目删除-批量"), handle.DelBatch)

	r.GET("/monitor_project/exportFile", middleware.RecordLog("监控项目导出"), handle.ExportFile)
	r.POST("/monitor_project/importFile", handle.ImportFile)
}

func init() {
	routeHandlers = append(routeHandlers, MonitorRoute, MonitorClientRoute, MonitorErrorRoute, MonitorProjectRoute)
}
