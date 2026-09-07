package admin_route

import (
	"x_admin/app/controller/admin_ctl"
	"x_admin/app/middleware"

	"github.com/gin-gonic/gin"
)

func UserProtocolRoute(rg *gin.RouterGroup) {
	handle := admin_ctl.UserProtocolHandler{}

	r := rg.Group("/", middleware.PermAuth())
	r.GET("/user_protocol/list", middleware.RequestCost(10000), handle.List)
	r.GET("/user_protocol/list_all", handle.ListAll)
	r.GET("/user_protocol/detail", handle.Detail)

	r.POST("/user_protocol/add", middleware.RecordLog("用户协议新增"), handle.Add)
	r.POST("/user_protocol/edit", middleware.RecordLog("用户协议编辑"), handle.Edit)

	r.POST("/user_protocol/del", middleware.RecordLog("用户协议删除"), handle.Del)
	r.POST("/user_protocol/del_batch", middleware.RecordLog("用户协议删除-批量"), handle.DelBatch)

	r.GET("/user_protocol/export_file", middleware.RecordLog("用户协议导出"), handle.ExportFile)
	r.POST("/user_protocol/import_file", handle.ImportFile)
}
func init() {
	routeHandlers = append(routeHandlers, UserProtocolRoute)
}
