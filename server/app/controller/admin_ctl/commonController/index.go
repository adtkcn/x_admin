package commonController

import (
	"x_admin/app/service/commonService"
	"x_admin/core/response"

	"github.com/gin-gonic/gin"
)

// IndexHandler 首页控制器
type IndexHandler struct{}

// Console 控制台
func (ih IndexHandler) Console(c *gin.Context) {
	res, err := commonService.IndexService.Console()
	response.CheckAndRespWithData(c, res, err)
}

// Config 公共配置
func (ih IndexHandler) Config(c *gin.Context) {
	res, err := commonService.IndexService.Config()
	response.CheckAndRespWithData(c, res, err)
}
