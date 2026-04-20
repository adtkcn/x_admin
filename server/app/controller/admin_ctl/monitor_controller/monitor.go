package monitor_controller

import (
	"strings"
	"x_admin/app/service/monitor_service"
	"x_admin/core/response"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

// MonitorHandler 监控控制器（服务端+缓存）
type MonitorHandler struct{}

// Cache 缓存监控
func (mh MonitorHandler) Cache(c *gin.Context) {
	cmdStatsMap := util.RedisUtil.Info("commandstats")
	var stats []map[string]string
	for k, v := range cmdStatsMap {
		stats = append(stats, map[string]string{
			"name":  strings.Split(k, "_")[1],
			"value": v[strings.Index(v, "=")+1 : strings.Index(v, ",")],
		})
	}
	response.Ok(c, map[string]any{
		"info":         util.RedisUtil.Info(),
		"commandStats": stats,
		"dbSize":       util.RedisUtil.DBSize(),
	})
}

// Server 服务监控
func (mh MonitorHandler) Server(c *gin.Context) {
	data, err := monitor_service.MonitorServerService.GetAllServerLatestInfo()
	if err != nil {
		response.Fail(c, "获取服务器信息失败:"+err.Error())
		return
	}
	response.Ok(c, data)
}
