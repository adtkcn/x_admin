package iot_route

import "github.com/gin-gonic/gin"

type RouteHandlerFunc func(*gin.RouterGroup)

// iotRouteHandlers 设备侧路由注册函数切片，由各文件 init() 收集（同 admin_route 惯例）
var iotRouteHandlers []RouteHandlerFunc

// Autoload 自动加载所有设备侧路由
func Autoload(rg *gin.RouterGroup) {
	for _, handler := range iotRouteHandlers {
		handler(rg)
	}
}

// RegisterRoute 设备接入路由入口（/api/iot/*）
//
// 与 /api/admin、/api/web 不同：这里面向设备本身，鉴权走设备凭证中间件，
// 由各路由文件自行决定挂在哪一段，因此入口只做自动加载。
func RegisterRoute(iot *gin.RouterGroup, rootRouter *gin.Engine) {
	Autoload(iot)
}
