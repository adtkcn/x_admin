package admin_route

import (
	"x_admin/app/controller/admin_ctl/fabu_controller"
	"x_admin/app/middleware"

	"github.com/gin-gonic/gin"
)

func initFabuRoute(rg *gin.RouterGroup) {
	handleApp := fabu_controller.FabuAppHandler{}
	handleVersion := fabu_controller.FabuVersionHandler{}
	handleWgt := fabu_controller.FabuWgtHandler{}

	auth := rg.Group("/fabu", middleware.LoginAuth())
	{
		auth.GET("/app/list", handleApp.List)
		auth.GET("/app/detail", handleApp.Detail)
		auth.POST("/app/add", handleApp.Add)
		auth.POST("/app/edit", handleApp.Edit)
		auth.POST("/app/del", handleApp.Del)

		auth.GET("/version/list", handleVersion.List)
		auth.POST("/version/upload", handleVersion.Upload)
		auth.POST("/version/release", handleVersion.Release)
		auth.POST("/version/cancel", handleVersion.Cancel)
		auth.POST("/version/gray", handleVersion.Gray)
		auth.POST("/version/updateMode", handleVersion.UpdateMode)
		auth.POST("/version/del", handleVersion.Del)

		auth.GET("/wgt/list", handleWgt.List)
		auth.POST("/wgt/upload", handleWgt.Upload)
		auth.POST("/wgt/del", handleWgt.Del)
	}
}

func init() {
	routeHandlers = append(routeHandlers, initFabuRoute)
}
