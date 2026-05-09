package common_controller

import (
	"x_admin/app/service/common_service"
	"x_admin/core/response"

	"github.com/gin-gonic/gin"
)

// IndexHandler 首页控制器
type IndexHandler struct{}

// @Summary		控制台数据
// @Description	获取控制台统计数据
// @Tags			common_index-首页
// @Param			token	header		string						true	"token"
// @Success		200		{object}	response.Response			"成功"
// @Router			/api/admin/common/index/console [get]
func (ih IndexHandler) Console(c *gin.Context) {
	res, err := common_service.IndexService.Console()
	response.CheckAndRespWithData(c, res, err)
}

// @Summary		公共配置
// @Description	获取系统公共配置
// @Tags			common_index-首页
// @Param			token	header		string						true	"token"
// @Success		200		{object}	response.Response			"成功"
// @Router			/api/admin/common/index/config [get]
func (ih IndexHandler) Config(c *gin.Context) {
	res, err := common_service.IndexService.Config()
	response.CheckAndRespWithData(c, res, err)
}
