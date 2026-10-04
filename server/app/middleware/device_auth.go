package middleware

import (
	"strings"
	"x_admin/app/schema/iot_schema"
	"x_admin/app/service/iot_service"
	"x_admin/config"
	"x_admin/core/response"

	"github.com/gin-gonic/gin"
)

// DeviceAuth 设备接入鉴权中间件
// 与后台/用户登录态完全无关：仅认设备凭证（X-Device-Sn + X-Access-Token），
// 鉴权结果 *iot_schema.DeviceAuthInfo 写入 Context 供接入层组装报文。
func DeviceAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		sn := strings.TrimSpace(c.GetHeader(iot_schema.HeaderDeviceSn))
		token := strings.TrimSpace(c.GetHeader(iot_schema.HeaderAccessToken))
		if sn == "" || token == "" {
			response.Fail(c, response.TokenEmpty.SetMessage(
				"缺少接入凭证头 "+iot_schema.HeaderDeviceSn+"/"+iot_schema.HeaderAccessToken))
			c.Abort()
			return
		}

		info, err := iot_service.DeviceService.Authenticate(sn, token)
		if err != nil {
			response.Fail(c, err)
			c.Abort()
			return
		}
		if info.ProductStop {
			response.Fail(c, response.AssertArgumentError.SetMessage("设备所属产品已停用!"))
			c.Abort()
			return
		}
		// 协议准入：产品明确配置了接入协议且不是 http 时拒绝走 HTTP 通道；
		// 空值放行以兼容历史数据（尚未填写 protocol 的产品）
		if info.Protocol != "" && !strings.EqualFold(info.Protocol, "http") {
			response.Fail(c, response.AssertArgumentError.SetMessage(
				"该产品接入协议为 "+info.Protocol+"，不接受 HTTP 上报!"))
			c.Abort()
			return
		}

		c.Set(config.DeviceCtxKey, info)
		c.Next()
	}
}

// DeviceFromCtx 读取 DeviceAuth 写入的设备信息，须保证路由已挂 DeviceAuth()
func DeviceFromCtx(c *gin.Context) (*iot_schema.DeviceAuthInfo, bool) {
	v, ok := c.Get(config.DeviceCtxKey)
	if !ok {
		return nil, false
	}
	info, ok := v.(*iot_schema.DeviceAuthInfo)
	return info, ok && info != nil
}
