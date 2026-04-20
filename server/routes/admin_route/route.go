package admin_route

import (
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

// RegisterRoute 后台管理路由入口（按模块分组注册）
func RegisterRoute(rg *gin.RouterGroup) {

	rg = rg.Group("/admin")
	// 所有子路由需要加上前缀 /api/admin

	Autoload(rg)
}
