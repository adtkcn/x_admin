package router

import (
	"x_admin/controller/admin/commonController"
	"x_admin/core/response"
	"x_admin/middleware"
	"x_admin/router/admin"

	"github.com/gin-gonic/gin"
)

func RegisterRoute(api *gin.RouterGroup, rootRouter *gin.Engine) {

	// /api/admin/apiList 获取所有接口
	api.GET("/admin/apiList", middleware.TokenAuth(), func(ctx *gin.Context) {
		var path = []string{}
		for _, route := range rootRouter.Routes() {
			// fmt.Printf("%s 127.0.0.1:%v%s\n", route.Method, config.Config.ServerPort, route.Path)
			path = append(path, route.Path)
		}
		response.Result(ctx, response.Success, path)
	})

	// /api/admin
	admin.RegisterRoute(api)
	// /api/common/captcha 验证码
	commonController.CaptchaRoute(api)
}
