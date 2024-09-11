package router

import (
	"x_admin/admin/common/captcha"
	"x_admin/router/admin"

	"github.com/gin-gonic/gin"
)

func RegisterGroup(rg *gin.RouterGroup) {
	// 一下路由前缀为/api
	admin.RegisterGroup(rg)

	captcha.CaptchaRoute(rg)
}
