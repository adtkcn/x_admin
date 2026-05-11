package routes

import (
	"fmt"
	"x_admin/app/controller"
	"x_admin/app/controller/admin_ctl/common_controller"
	"x_admin/app/middleware"
	"x_admin/config"
	"x_admin/core/response"
	"x_admin/docs"
	"x_admin/routes/admin_route"

	"github.com/gin-gonic/gin"
)

// @Summary	获取所有接口
// @Tags		公共接口
// @Router		/api/admin/apiList [get]
func apiList(api *gin.RouterGroup, rootRouter *gin.Engine) {

	api.GET("/admin/apiList", middleware.PermAuth(), func(ctx *gin.Context) {
		var path = []string{}
		for _, route := range rootRouter.Routes() {
			path = append(path, route.Path)
		}
		response.Ok(ctx, path)
	})
}

// @Summary	swagger文档数据
// @Tags		公共接口
// @Router		/api/swagger/doc.json [get]
func swaggerDoc(api *gin.RouterGroup) {
	api.GET("/swagger/doc.json", func(c *gin.Context) {
		// 获取域名和端口号
		host := ""
		docs.SwaggerInfo.Host = fmt.Sprintf("%v", host)
		docs.SwaggerInfo.Title = config.AppConfig.AppName
		docs.SwaggerInfo.Version = config.AppConfig.Version
		c.String(200, docs.SwaggerInfo.ReadDoc())
	})
}

// initCaptchaRoute 验证码路由
func initCaptchaRoute(api *gin.RouterGroup) {
	handleCaptcha := common_controller.CaptchaHandler{}
	captchaRg := api.Group("/common/captcha")
	captchaRg.POST("/get", handleCaptcha.Get)
	captchaRg.POST("/check", handleCaptcha.Check)
}

func wsHandler(api *gin.RouterGroup) {
	api.GET("/ws", middleware.LoginAuth(), controller.WsHandler)
}

func registerApiRoute(api *gin.RouterGroup, rootRouter *gin.Engine) {
	// 静态文件路由
	api.Static("/static", "./public/static")
	rootRouter.Static(config.FileConfig.PublicPrefix, config.FileConfig.UploadDirectory)

	// 设置中间件
	rootRouter.Use(gin.Logger(), middleware.Cors(), middleware.ErrorRecover())
	apiList(api, rootRouter)

	swaggerDoc(api)
	initCaptchaRoute(api)

	wsHandler(api)
	// /api/admin
	admin_route.RegisterRoute(api)

}
