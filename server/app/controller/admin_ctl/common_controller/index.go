package common_controller

import (
	"x_admin/app/service/common_service"
	"x_admin/core/response"

	"github.com/gin-gonic/gin"
)

// IndexHandler 首页控制器
type IndexHandler struct{}

// Console 控制台
func (ih IndexHandler) Console(c *gin.Context) {
	res, err := common_service.IndexService.Console()
	response.CheckAndRespWithData(c, res, err)
}

// Config 公共配置
func (ih IndexHandler) Config(c *gin.Context) {
	res, err := common_service.IndexService.Config()
	response.CheckAndRespWithData(c, res, err)
}
