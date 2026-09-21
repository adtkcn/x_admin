package routes

import (
	"x_admin/app/controller"
	"x_admin/app/controller/admin_ctl/common_controller"
	"x_admin/app/middleware"
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

	// 文件流路由：按 file_hash_id 返回物理文件（取代原静态目录映射）
	handleFile := common_controller.UploadHandler{}
	api.GET("/uploads/:id/:file_name", handleFile.Serve)

	// 全局中间件（gin.Logger/Cors/ErrorRecover）已在 InitRouter 创建 /api 分组前统一注册，
	// 此处不可再 Use，否则对已建分组无效

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
