package routes

import (
	"x_admin/app/middleware"
	"x_admin/config"
	"x_admin/core/response"

	"github.com/gin-gonic/gin"
)

// InitRouter 初始化router
func InitRouter() *gin.Engine {
	// 初始化gin
	gin.SetMode(config.AppConfig.GinMode)

	var RootRouter = gin.New()
	RootRouter.MaxMultipartMemory = 50 << 20 // 50 MiB

	// 全局中间件必须在创建任何路由分组之前注册：
	// gin 的 Group() 会快照当前 handlers，之后再 RootRouter.Use() 不会回流到已建分组，
	// 否则 Cors 等中间件对 /api 路由不生效（跨域头不会写出）
	RootRouter.Use(gin.Logger(), middleware.Cors(), middleware.ErrorRecover())

	// 404 路由
	RootRouter.NoRoute(response.NoRoute)

	// 注册/api路由
	apiGroup := RootRouter.Group("/api")
	{
		registerApiRoute(apiGroup, RootRouter)
	}

	return RootRouter
}
