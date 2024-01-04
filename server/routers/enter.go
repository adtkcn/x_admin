package routers

import (
	"x_admin/admin"
	"x_admin/admin/common/captcha"

	"github.com/gin-gonic/gin"
)

func RegisterGroup(rg *gin.RouterGroup) {
	// 一下路由前缀为/api
	admin.RegisterGroup(rg)

	captcha.CaptchaRoute(rg)
}
