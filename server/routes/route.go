package routes

import (
	"fmt"
	"x_admin/app/controller"
	"x_admin/app/controller/admin_ctl/commonController"
	"x_admin/config"
	"x_admin/core/response"
	"x_admin/docs"
	"x_admin/middleware"
	"x_admin/routes/adminRoute"

	"github.com/gin-gonic/gin"
)

// @Summary	获取所有接口
// @Tags		公共接口
// @Router		/api/admin/apiList [get]
func apiList(api *gin.RouterGroup, rootRouter *gin.Engine) {

	api.GET("/admin/apiList", middleware.TokenAuth(), func(ctx *gin.Context) {
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
func swaggerJson(api *gin.RouterGroup) {
	api.GET("/swagger/doc.json", func(c *gin.Context) {
		// 获取域名和端口号
		host := c.Request.Host
		// port := c.Request.Port
		// docs.SwaggerInfo.Host = fmt.Sprintf("%v:%v", host, config.AppConfig.Port)
		docs.SwaggerInfo.Host = fmt.Sprintf("%v", host)
		docs.SwaggerInfo.Title = config.AppConfig.AppName
		docs.SwaggerInfo.Version = config.AppConfig.Version
		c.String(200, docs.SwaggerInfo.ReadDoc())
	})
}

// @Summary	ws通用接口
// @schemes	ws
// @Tags		公共接口
// @Success	101	{string}	string	"ws连接成功"
// @Router		/api/ws [get]
func wsHandler(api *gin.RouterGroup) {
	api.GET("/ws", middleware.LoginAuth(), controller.WsHandler)
}

func RegisterRoute(api *gin.RouterGroup, rootRouter *gin.Engine) {

	apiList(api, rootRouter)

	swaggerJson(api)

	wsHandler(api)
	// /api/admin
	adminRoute.RegisterRoute(api)

	// /api/common/captcha 验证码
	commonController.CaptchaRoute(api)

}
