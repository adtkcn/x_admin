package common_controller

import (
	"x_admin/app/service/common_service"
	"x_admin/app/service/iot_service"
	"x_admin/core/response"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/singleflight"
)

// IndexHandler 首页控制器
type IndexHandler struct {
	// 防止缓存击穿场景，将大量重复数据库查询合并为一次。
	requestGroup singleflight.Group
}

// @Summary		控制台数据
// @Description	获取控制台统计数据
// @Tags			common_index-首页
// @Param			token	header		string				true	"token"
// @Success		200		{object}	response.Response	"成功"
// @Router			/api/admin/common/index/console [get]
func (ih *IndexHandler) Console(c *gin.Context) {
	res, err, _ := ih.requestGroup.Do("Index:Console", func() (any, error) {
		data, e := common_service.IndexService.Console()
		if e != nil {
			return nil, e
		}
		// 跨域数据在 controller 组合：设备总量/在线量 + 今日告警分级统计合并进 today
		deviceTotal, deviceOnline, e := iot_service.DeviceService.GetConsoleCount()
		if e != nil {
			return nil, e
		}
		alertMinor, alertSerious, alertFault, e := iot_service.AlertService.GetConsoleCount()
		if e != nil {
			return nil, e
		}
		today := data["today"].(map[string]any)
		today["deviceTotal"] = deviceTotal
		today["deviceOnline"] = deviceOnline
		today["alertMinor"] = alertMinor
		today["alertSerious"] = alertSerious
		today["alertFault"] = alertFault
		return data, nil
	})
	response.JSON(c, res, err)
}

// @Summary		公共配置
// @Description	获取系统公共配置
// @Tags			common_index-首页
// @Param			token	header		string				true	"token"
// @Success		200		{object}	response.Response	"成功"
// @Router			/api/admin/common/index/config [get]
func (ih *IndexHandler) Config(c *gin.Context) {
	res, err, _ := ih.requestGroup.Do("Index:Config", func() (any, error) {
		return common_service.IndexService.Config()
	})
	response.JSON(c, res, err)
}
