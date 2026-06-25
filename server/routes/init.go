package routes

import (
	"x_admin/config"
	"x_admin/core/response"

	"github.com/gin-gonic/gin"
)

// initRouter 初始化router
func InitRouter() *gin.Engine {
	// 初始化gin
	gin.SetMode(config.AppConfig.GinMode)

	var RootRouter = gin.New()
	RootRouter.MaxMultipartMemory = 8 << 20 // 8 MiB

	// 404 路由
	RootRouter.NoRoute(response.NoRoute)

	// 注册/api路由
	apiGroup := RootRouter.Group("/api")
	{
		registerApiRoute(apiGroup, RootRouter)
	}

	return RootRouter
}
