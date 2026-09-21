package admin_route

import (
	"x_admin/app/controller/admin_ctl"
	"x_admin/app/middleware"

	"github.com/gin-gonic/gin"
)

func SystemCornRoute(rg *gin.RouterGroup) {
	handle := admin_ctl.SystemCornHandler{}

	r := rg.Group("/", middleware.PermAuth())
	r.GET("/system_corn/list", handle.List)
	r.GET("/system_corn/list_all", handle.ListAll)
	r.GET("/system_corn/detail", handle.Detail)

	r.POST("/system_corn/add", middleware.RecordLog("定时任务新增"), handle.Add)
	r.POST("/system_corn/edit", middleware.RecordLog("定时任务编辑"), handle.Edit)

	r.POST("/system_corn/del", middleware.RecordLog("定时任务删除"), handle.Del)
	r.POST("/system_corn/del_batch", middleware.RecordLog("定时任务删除-批量"), handle.DelBatch)

	r.GET("/system_corn/export_file", middleware.RecordLog("定时任务导出"), handle.ExportFile)
	r.POST("/system_corn/import_file", handle.ImportFile)

	r.GET("/system_corn/getTaskList", handle.GetTaskList)
}
func init() {
	routeHandlers = append(routeHandlers, SystemCornRoute)
}
