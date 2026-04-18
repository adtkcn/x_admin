package admin_route

import (
	"x_admin/app/controller/admin_ctl/generatorController"
	"x_admin/app/middleware"

	"github.com/gin-gonic/gin"
)

// GeneratorRoute 代码生成器模块路由
func GeneratorRoute(rg *gin.RouterGroup) {
	handle := generatorController.GenHandler{}

	rg = rg.Group("/gen", middleware.PermAuth())
	rg.GET("/db", handle.DbTables)
	rg.GET("/list", handle.List)
	rg.GET("/detail", handle.Detail)
	rg.POST("/importTable", handle.ImportTable)
	rg.POST("/syncTable", handle.SyncTable)
	rg.POST("/editTable", handle.EditTable)
	rg.POST("/delTable", handle.DelTable)
	rg.GET("/previewCode", handle.PreviewCode)
	rg.GET("/downloadCode", handle.DownloadCode)
}

func init() {
	routeHandlers = append(routeHandlers, GeneratorRoute)
}
