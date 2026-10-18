package alert

import (
	"strconv"
	"strings"
	"time"

	"x_admin/app/service/iot_service/pipeline/online"
)

// alertTimerTTLBuffer 持续/恢复计时键在时长外的冗余存活，避免边界抖动从头计时
const alertTimerTTLBuffer = 5 * time.Minute

func alertActiveKey(deviceId string) string  { return online.RuntimeKey("alert:active:" + deviceId) }
func alertFirstKey(deviceId string) string   { return online.RuntimeKey("alert:first:" + deviceId) }
func alertRecoverKey(deviceId string) string { return online.RuntimeKey("alert:recover:" + deviceId) }

func alertFaultActiveKey(deviceId string) string {
	return online.RuntimeKey("alert:fault:active:" + deviceId)
}
func alertFaultFirstKey(deviceId string) string {
	return online.RuntimeKey("alert:fault:first:" + deviceId)
}
func alertFaultRecoverKey(deviceId string) string {
	return online.RuntimeKey("alert:fault:recover:" + deviceId)
}

// allKeys 设备全部告警运行态键（普通流+故障流共六键），供清态一次 Del
func allKeys(deviceId string) []string {
	return []string{
		alertActiveKey(deviceId), alertFirstKey(deviceId), alertRecoverKey(deviceId),
		alertFaultActiveKey(deviceId), alertFaultFirstKey(deviceId), alertFaultRecoverKey(deviceId),
	}
}

// streamKeys 一条事件流的三把运行态键；streamVals 为其 MGet 读回的原始值
type streamKeys struct{ active, first, recover string }
type streamVals struct{ active, first, recover string }

func normalKeys(deviceId string) streamKeys {
	return streamKeys{alertActiveKey(deviceId), alertFirstKey(deviceId), alertRecoverKey(deviceId)}
}
func faultKeys(deviceId string) streamKeys {
	return streamKeys{alertFaultActiveKey(deviceId), alertFaultFirstKey(deviceId), alertFaultRecoverKey(deviceId)}
}

// encodeActive 活动态值编码为 "<histId>:<峰值级别>:<当前级别>"
func encodeActive(histId string, peak, cur int8) string {
	return histId + ":" + strconv.FormatInt(int64(peak), 10) + ":" + strconv.FormatInt(int64(cur), 10)
}

// parseActive 解出活动态里的历史主键、峰值级别、当前级别；空值或格式异常返回 ("", 0, 0)
func parseActive(raw string) (histId string, peak, cur int8) {
	parts := strings.SplitN(raw, ":", 3)
	if len(parts) < 3 {
		return "", 0, 0
	}
	histId = parts[0]
	p, _ := strconv.ParseInt(parts[1], 10, 8)
	peak = int8(p)
	c, _ := strconv.ParseInt(parts[2], 10, 8)
	cur = int8(c)
	return histId, peak, cur
}

// strOf 取 MGet 结果里的字符串值（键不存在时元素为 nil）
func strOf(v any) string {
	if s, ok := v.(string); ok {
		return s
	}
	return ""
}
