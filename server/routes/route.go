package routes

import (
	"x_admin/app/controller/admin_ctl/commonController"
	"x_admin/core/response"
	"x_admin/middleware"
	"x_admin/routes/adminRoute"

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
	adminRoute.RegisterRoute(api)

	// /api/common/captcha 验证码
	commonController.CaptchaRoute(api)
}
