package monitorController

import (
	"strings"
	"x_admin/app/service/monitorService"
	"x_admin/core/response"
	"x_admin/middleware"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

func RegisterRoute(rg *gin.RouterGroup) {
	handle := monitorHandler{}

	rg = rg.Group("/monitor", middleware.TokenAuth())
	rg.GET("/cache", handle.cache)
	rg.GET("/server", handle.server)
}

type monitorHandler struct{}

// cache 缓存监控
func (mh monitorHandler) cache(c *gin.Context) {
	cmdStatsMap := util.RedisUtil.Info("commandstats")
	var stats []map[string]string
	for k, v := range cmdStatsMap {
		stats = append(stats, map[string]string{
			"name":  strings.Split(k, "_")[1],
			"value": v[strings.Index(v, "=")+1 : strings.Index(v, ",")],
		})
	}
	response.OkWithData(c, map[string]interface{}{
		"info":         util.RedisUtil.Info(),
		"commandStats": stats,
		"dbSize":       util.RedisUtil.DBSize(),
	})
}

// server 服务监控
func (mh monitorHandler) server(c *gin.Context) {
	// response.OkWithData(c, map[string]interface{}{
	// 	"cpu":  util.ServerUtil.GetCpuInfo(),
	// 	"mem":  util.ServerUtil.GetMemInfo(),
	// 	"sys":  util.ServerUtil.GetSysInfo(),
	// 	"disk": util.ServerUtil.GetDiskInfo(),
	// 	"go":   util.ServerUtil.GetGoInfo(),
	// })
	data, err := monitorService.MonitorServerService.GetAllServerLatestInfo()
	if err != nil {
		response.FailWithMsg(c, response.SystemError, "获取服务器信息失败:"+err.Error())
		return
	}
	response.OkWithData(c, data)
}
