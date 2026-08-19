package routes

import (
	"x_admin/app/controller"
	"x_admin/app/controller/admin_ctl/common_controller"
	"x_admin/app/middleware"
	"x_admin/config"
	"x_admin/routes/admin_route"
	"x_admin/routes/web_route"

	"github.com/gin-gonic/gin"
)

// captchaRoute 验证码路由
func captchaRoute(api *gin.RouterGroup) {
	handleCaptcha := common_controller.CaptchaHandler{}

	api.POST("/common/captcha/get", handleCaptcha.Get)
	api.POST("/common/captcha/check", handleCaptcha.Check)
}

func wsRoute(api *gin.RouterGroup) {
	api.GET("/ws", middleware.LoginAuth(), controller.WsHandler)
}

func registerApiRoute(api *gin.RouterGroup, rootRouter *gin.Engine) {
	// 静态文件路由
	api.Static("/static", "./public/static")
	rootRouter.Static(config.FileConfig.UploadPrefix, config.FileConfig.UploadDirectory)

	// 设置中间件
	rootRouter.Use(gin.Logger(), middleware.Cors(), middleware.ErrorRecover())

	captchaRoute(api)
	wsRoute(api)
	adminGroup := api.Group("/admin")
	{
		// /api/admin
		admin_route.RegisterRoute(adminGroup, rootRouter)
	}

	webGroup := api.Group("/web")
	{
		// /api/web
		web_route.RegisterRoute(webGroup, rootRouter)
	}
}
