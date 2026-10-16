package iot_ctl

import (
	"strconv"
	"time"
	"x_admin/app/middleware"
	"x_admin/app/schema/iot_schema"
	"x_admin/app/schema/queue_schema"
	"x_admin/config"
	"x_admin/core"
	"x_admin/core/response"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

// HttpGateway 设备侧 HTTP 接入适配器（iot.md 第一层：接入层）：
// 取 middleware.DeviceAuth 已校验的凭证上下文 → 报文合法性检查（条数上限、时间戳兑底）
// → 组装与协议无关的标准信封 iot_schema.IngestMsg → 投递 core.Queue("iot:ingest") 后立即返回。
// 绝不在请求线程里做解码/物模型校验/入库，保证设备洪峰下接入层耗时恒定；
// 后续 MQTT/TCP/WS 适配器复用同一信封与同一队列。
var HttpGateway = httpGateway{}

type httpGateway struct{}

// Telemetry 设备属性/遥测上报
//
//	@Summary		设备属性上报(HTTP接入)
//	@Description	设备上报属性键值，接入层校验凭证后投递处理队列，不同步处理
//	@Tags			iot_gateway-设备接入
//	@Accept			json
//	@Produce		json
//	@Param			X-Device-Sn		header		string							true	"设备编号"
//	@Param			X-Access-Token	header		string							true	"设备接入凭证"
//	@Param			data			body		iot_schema.DeviceTelemetryReq	true	"上报报文"
//	@Success		200				{object}	response.Response{}				"成功"
//	@Router			/api/iot/telemetry [post]
func (g httpGateway) Telemetry(c *gin.Context) {
	var req iot_schema.DeviceTelemetryReq
	if response.IsFail(c, util.VerifyUtil.VerifyBody(c, &req)) {
		return
	}
	info, ok := middleware.DeviceFromCtx(c)
	if !ok {
		response.Fail(c, response.TokenEmpty.SetMessage("设备凭证上下文缺失!"))
		return
	}
	// 防报文放大：单条报文属性条数上限由 IOT.MaxValues 控制
	if len(req.Values) > config.IotConfig.MaxValues {
		response.Fail(c, response.AssertArgumentError.SetMessage(
			"单次上报属性条数超过上限 "+strconv.Itoa(config.IotConfig.MaxValues)+"!"))
		return
	}
	msg := g.envelope(iot_schema.MsgTypeTelemetry, info, req.Ts, req.Values)
	g.push(c, msg)
}

// Heartbeat 设备心跳
//
//	@Summary		设备心跳(HTTP接入)
//	@Description	仅刷新在线态；超时未收到任何报文由处理层判离线
//	@Tags			iot_gateway-设备接入
//	@Accept			json
//	@Produce		json
//	@Param			X-Device-Sn		header		string							true	"设备编号"
//	@Param			X-Access-Token	header		string							true	"设备接入凭证"
//	@Param			data			body		iot_schema.DeviceHeartbeatReq	false	"心跳报文(可空)"
//	@Success		200				{object}	response.Response{}				"成功"
//	@Router			/api/iot/heartbeat [post]
func (g httpGateway) Heartbeat(c *gin.Context) {
	var req iot_schema.DeviceHeartbeatReq
	// 心跳允许空 body，解析失败不阻断
	_ = c.ShouldBindJSON(&req)
	info, ok := middleware.DeviceFromCtx(c)
	if !ok {
		response.Fail(c, response.TokenEmpty.SetMessage("设备凭证上下文缺失!"))
		return
	}
	msg := g.envelope(iot_schema.MsgTypeHeartbeat, info, req.Ts, nil)
	g.push(c, msg)
}

// envelope 组装标准信封：设备侧时间戳缺失或明显异常（未来/过早）时用服务端接收时间兜底
func (g httpGateway) envelope(msgType string, info *iot_schema.DeviceAuthInfo, ts int64, values map[string]any) iot_schema.IngestMsg {
	now := time.Now().UnixMilli()
	if ts <= 0 || ts > now+60_000 {
		ts = now
	}
	return iot_schema.IngestMsg{
		MsgType:        msgType,
		DeviceId:       info.DeviceId,
		Sn:             info.Sn,
		ProductId:      info.ProductId,
		SpaceId:        info.SpaceId,
		Protocol:       "http",
		OfflineTimeout: info.OfflineTimeout,
		Values:         values,
		Ts:             ts,
		ReceivedAt:     now,
	}
}

// push 投递队列并统一应答
func (g httpGateway) push(c *gin.Context, msg iot_schema.IngestMsg) {
	if config.IotConfig.LogRaw {
		core.Logger.Infof("[iot网关] 收到报文 sn=%s type=%s values=%d",
			msg.Sn, msg.MsgType, len(msg.Values))
	}
	if e := core.OrderedQueue.Enqueue(queue_schema.QueueIotIngest, msg.DeviceId, msg); e != nil {
		core.Logger.Errorf("[iot网关] 投递 %s 失败 sn=%s err=%v", queue_schema.QueueIotIngest, msg.Sn, e)
		response.Fail(c, response.SystemError.SetMessage("报文入队失败，请稍后重试!"))
		return
	}
	// 返回接收时间供设备侧对账；不做同步处理承诺
	response.JSON(c, gin.H{
		"msg_type":    msg.MsgType,
		"received_at": msg.ReceivedAt,
	}, nil)
}
