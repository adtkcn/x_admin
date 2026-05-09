package routes

import (
	"x_admin/config"
	"x_admin/core/response"

	"github.com/gin-gonic/gin"
)

var RootRouter *gin.Engine

func init() {
	// 初始化gin
	gin.SetMode(config.AppConfig.GinMode)

	RootRouter = gin.New()
	RootRouter.MaxMultipartMemory = 8 << 20 // 8 MiB
}

// initRouter 初始化router
func InitRouter() *gin.Engine {

	RootRouter.NoRoute(response.NoRoute)

	// 注册路由
	apiGroup := RootRouter.Group("/api")

	registerApiRoute(apiGroup, RootRouter)

	return RootRouter
}
