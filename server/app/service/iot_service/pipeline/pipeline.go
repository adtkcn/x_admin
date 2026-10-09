package pipeline

import (
	"context"
	"encoding/json"
	"errors"

	"x_admin/app/schema/iot_schema"
	"x_admin/app/service/iot_service"
	"x_admin/app/service/iot_service/pipeline/alert"
	"x_admin/app/service/iot_service/pipeline/online"
	"x_admin/app/service/iot_service/pipeline/rulechain"
	"x_admin/app/service/iot_service/pipeline/store"
	"x_admin/config"
	"x_admin/core"
)

// Pipeline 设备报文处理层（iot.md 第一层）：
// 接入层 → core.Queue("iot:ingest") → 本包 worker → [刷新运行态] →
// [可视化规则链逐节点执行：触发 → 解码 → 过滤 → 条件路由 → 告警/存储/推送/通知…]
//
// 本包只做“报文消费 + 对外门面”，能力已下沉为子包，依赖单向：
// online ← alert ← rulechain ← pipeline，store 为叶子旁路（时序落库、唯一数据出口）。

// StopRuntime 设备被删除后清理其运行态（含告警侧运行态，经 online 钩子收尾）
func StopRuntime(deviceId string) { online.ClearDeviceRuntime(deviceId) }

// StopAlertRuntime 只清设备告警运行态（供管理端删告警历史后编排）：
// 活动态里存的正是历史记录主键，删库不删键会让该设备后续命中永不新建告警；
// service 反向引用本包会成循环，故由控制器调门面编排。
func StopAlertRuntime(deviceId string) { alert.ClearRuntime(deviceId) }

// ScanOffline 快路径门面（供 app/corn 的 iotOfflineScan 定时调用）
func ScanOffline() { online.ScanOffline() }

// SweepOffline 慢路径门面（供 app/corn 的 iotOfflineSweep 定时调用）
func SweepOffline() { online.SweepOffline() }

// Runtime 读出 Redis 侧运行态，供管理端回读验证链路
func Runtime(deviceId string, timeout int) iot_schema.DeviceStateResp {
	return online.ReadDeviceState(deviceId, timeout)
}

// StopSink 优雅关闭时冲刷存储缓冲（供 main 调用）
func StopSink() { store.Flush() }

// QueryRange 按时间区间回读设备时序，供管理端趋势图调用
func QueryRange(deviceId string, startMs, endMs int64, limit int) []iot_schema.DeviceDataResp {
	return store.QueryRange(deviceId, startMs, endMs, limit)
}

// CleanData 清理保留期外的时序数据（供 app/corn 的 iotDataClean 定时调用）
func CleanData() { store.Clean() }

// HandleIngest 消费一条设备报文
func HandleIngest(_ context.Context, body []byte) error {
	var msg iot_schema.IngestMsg
	if err := json.Unmarshal(body, &msg); err != nil {
		// 坏报文无法自愈，返回错误由队列记日志即可（当前不重试，避免死循环）
		core.Logger.Errorf("[iot处理] 报文解析失败 err=%v body=%s", err, string(body))
		return err
	}
	if msg.DeviceId == "" {
		return errors.New("报文缺少 device_id")
	}
	if config.IotConfig.LogRaw {
		core.Logger.Infof("[iot处理] 消费报文 sn=%s type=%s values=%d", msg.Sn, msg.MsgType, len(msg.Values))
	}

	// 运行态与在线转换对所有报文类型都生效（心跳/上报都“续命”）
	online.RefreshOnline(msg)

	switch msg.MsgType {
	case iot_schema.MsgTypeTelemetry:
		return handleTelemetry(msg)
	case iot_schema.MsgTypeHeartbeat:
		return nil
	default:
		return errors.New("未知报文类型: " + msg.MsgType)
	}
}

// handleTelemetry 属性上报：交给产品的可视化规则链按有向图逐节点执行（规则链 ID 走产品缓存，避免每报文查库）
func handleTelemetry(msg iot_schema.IngestMsg) error {
	model := iot_service.DeviceService.CachedProductModel(msg.ProductId)
	return rulechain.Run(msg, model)
}
