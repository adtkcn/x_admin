package commonController

import (
	"x_admin/app/middleware"
	"x_admin/app/service/commonService"
	"x_admin/core/response"

	"github.com/gin-gonic/gin"
)

func IndexRoute(rg *gin.RouterGroup) {
	handle := indexHandler{}

	rg = rg.Group("/common")
	rg.GET("/index/console", middleware.LoginAuth(), handle.console)
	rg.GET("/index/config", handle.config)
}

type indexHandler struct{}

// console 控制台
func (ih indexHandler) console(c *gin.Context) {
	res, err := commonService.IndexService.Console()
	response.CheckAndRespWithData(c, res, err)
}

// config 公共配置
func (ih indexHandler) config(c *gin.Context) {
	res, err := commonService.IndexService.Config()
	response.CheckAndRespWithData(c, res, err)
}
