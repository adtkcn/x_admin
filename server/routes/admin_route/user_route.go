package admin_route

import (
	"x_admin/app/controller/admin_ctl"
	"x_admin/app/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoute(rg *gin.RouterGroup) {
	handle := admin_ctl.UserHandler{}
	r := rg.Group("/", middleware.PermAuth())
	r.GET("/user/list", handle.List)
	r.GET("/user/detail", handle.Detail)
	r.POST("/user/edit", middleware.RecordLog("用户编辑"), handle.Edit)
	r.POST("/user/disable", middleware.RecordLog("用户禁用/启用"), handle.Disable)
	r.POST("/user/kick", middleware.RecordLog("用户踢下线"), handle.Kick)
}

func init() {
	routeHandlers = append(routeHandlers, UserRoute)
}
