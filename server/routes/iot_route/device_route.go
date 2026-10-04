package iot_route

import (
	"x_admin/app/iot/gateway"
	"x_admin/app/middleware"

	"github.com/gin-gonic/gin"
)

// initDeviceRoute 设备侧 HTTP 接入路由（iot.md 第一层：接入层）
//
// 整组都要过 DeviceAuth：设备凭证与后台登录态无关，走 X-Device-Sn + X-Access-Token。
func initDeviceRoute(rg *gin.RouterGroup) {
	ingress := rg.Group("", middleware.DeviceAuth())
	{
		ingress.POST("/telemetry", gateway.HttpGateway.Telemetry)
		ingress.POST("/heartbeat", gateway.HttpGateway.Heartbeat)
	}
}

func init() {
	iotRouteHandlers = append(iotRouteHandlers, initDeviceRoute)
}
