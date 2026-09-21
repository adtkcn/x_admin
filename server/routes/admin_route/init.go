package admin_route

import (
	"x_admin/app/middleware"
	"x_admin/core/response"

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

// @Summary	获取所有接口
// @Tags		公共接口
// @Router		/api/admin/apiList [get]
func apiList(api *gin.RouterGroup, rootRouter *gin.Engine) {
	api.GET("/apiList", middleware.PermAuth(), func(ctx *gin.Context) {
		var path = []string{}
		for _, route := range rootRouter.Routes() {
			path = append(path, route.Path)
		}
		response.Ok(ctx, path)
	})
}

// RegisterRoute 后台管理路由入口（按模块分组注册）
func RegisterRoute(admin *gin.RouterGroup, rootRouter *gin.Engine) {
	apiList(admin, rootRouter)
	Autoload(admin)
}
