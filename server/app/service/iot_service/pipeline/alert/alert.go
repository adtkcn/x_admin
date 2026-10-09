package alert

import (
	"context"
	"strconv"
	"strings"
	"time"

	"x_admin/app/model/iot_model"
	"x_admin/app/schema/iot_schema"
	"x_admin/app/service/iot_service"
	"x_admin/app/service/iot_service/pipeline/online"
	"x_admin/core"
)

// 告警状态机：把规则链告警节点脚本 alert(values, sn) 的逐条判定（0=不告警/恢复，非 0 即级别）
// 收敛成“事件级”告警（持续时长/静默/级别变更/恢复），运行态在 Redis、历史在 MySQL。
// 判据来自报文本身，无需到期扫描：每条上报顺带评估一次，只在转换点写库推 WS，
// 持续期间的重复命中不落库不广播。
//
// 单事件可改写：一台设备只维护一条“告警中”记录；新级别覆盖优先序更高（见 iot_model.AlertPriority，
// 故障危害最轻却能接管任何在告事件）、或当前故障已不再命中时，改写同条记录的 severity 并重推；
// 优先序更低则维持不动，仅脚本返回 0 才结束事件。
//
// Redis 键（{RedisPrefix}iot:alert:*）：
//   - active:<deviceId>   存在=正处于告警事件，value="<历史主键>:<级别>"，无 TTL
//   - first:<deviceId>    持续计时起点(毫秒)，仅 duration>0 时使用，TTL=时长+缓冲
//   - silence:<deviceId>  静默标记，TTL=silence，抑制其后的再次触发

// alertFirstTTLBuffer 持续计时键在时长之外的冗余存活，避免边界抖动导致从头重新计时
const alertFirstTTLBuffer = 5 * time.Minute

// init 经 online 钩子把告警收尾接出去（回调而非 import 上层，避免 online ↔ alert 循环依赖）
func init() {
	online.OnDeviceOffline(closeOnOffline) // 设备转离线：收尾其活动告警
	online.OnDeviceClear(ClearRuntime)     // 删设备：一并清告警 Redis 键
}

func alertActiveKey(deviceId string) string  { return online.RuntimeKey("alert:active:" + deviceId) }
func alertFirstKey(deviceId string) string   { return online.RuntimeKey("alert:first:" + deviceId) }
func alertSilenceKey(deviceId string) string { return online.RuntimeKey("alert:silence:" + deviceId) }

// ClearRuntime 只清设备告警运行态（三键），不做库侧收尾。活动态无 TTL 且 value 就是历史记录主键，
// 残留会让该设备后续命中永远不再新建记录，所以删库后必调。
func ClearRuntime(deviceId string) {
	ctx, cancel := online.RedisCtx()
	defer cancel()
	_ = core.Redis.Del(ctx, alertActiveKey(deviceId), alertFirstKey(deviceId), alertSilenceKey(deviceId)).Err()
}

// closeOnOffline 告警状态机由报文驱动，设备停报后永远等不到恢复转换，故离线即收尾（置恢复 + 清三键），
// 重新上线且条件命中会另起新事件。
func closeOnOffline(deviceId string) {
	ctx, cancel := online.RedisCtx()
	defer cancel()
	// 先确认有活动态才发 UPDATE：离线钩子对每台设备都跑，无告警时那条 UPDATE 永远命中 0 行
	if raw, err := core.Redis.Get(ctx, alertActiveKey(deviceId)).Result(); err == nil && raw != "" {
		if e := iot_service.AlertService.CloseActiveByDevice(deviceId); e != nil {
			core.Logger.Errorf("[iot处理] 离线收尾活动告警失败 deviceId=%s err=%v", deviceId, e)
		}
	}
	_ = core.Redis.Del(ctx, alertActiveKey(deviceId), alertFirstKey(deviceId), alertSilenceKey(deviceId)).Err()
}

// encodeActive 活动态值编码为 "<histId>:<级别>"；uuid 主键不含冒号，按首个冒号切分安全
func encodeActive(histId string, level int8) string {
	return histId + ":" + strconv.FormatInt(int64(level), 10)
}

// parseActive 解出活动态里的历史主键与级别；空值返回 ("", 0)
func parseActive(raw string) (string, int8) {
	if raw == "" {
		return "", 0
	}
	if i := strings.IndexByte(raw, ':'); i >= 0 {
		lvl, _ := strconv.ParseInt(raw[i+1:], 10, 8)
		return raw[:i], int8(lvl)
	}
	return raw, 0 // 兼容旧格式（仅存主键、无级别）：级别按 0，下条异常报文会自然纠正
}

// EvalAlert 按本次告警级别推进状态机。level 为告警节点的数字返回（0=不告警/恢复）；
// duration/silence 取自告警节点属性（秒，0=命中即告/不静默）。
func EvalAlert(msg iot_schema.IngestMsg, values map[string]any, level int8, duration, silence int) {
	ctx, cancel := online.RedisCtx()
	defer cancel()

	// 三键一次 MGet 读出：分开读是 3 个 RTT，且中间态可能被并发报文改写
	deviceId := msg.DeviceId
	rs, err := core.Redis.MGet(ctx, alertActiveKey(deviceId), alertFirstKey(deviceId), alertSilenceKey(deviceId)).Result()
	if err != nil {
		core.Logger.Errorf("[iot处理] 读告警运行态失败 deviceId=%s err=%v", deviceId, err)
		return
	}
	histId, curLevel := parseActive(strOf(rs[0]))
	firstStr := strOf(rs[1])
	silenced := len(rs) > 2 && rs[2] != nil

	// 恢复：清持续计时；正处事件则收尾（置恢复 + 清活动态 + 推恢复）
	if level == 0 {
		core.Redis.Del(ctx, alertFirstKey(deviceId))
		if histId != "" {
			core.Redis.Del(ctx, alertActiveKey(deviceId))
			if e := iot_service.AlertService.Recover(histId, msg.ReceivedAt); e != nil {
				core.Logger.Errorf("[iot处理] 告警恢复落库失败 deviceId=%s id=%s err=%v", deviceId, histId, e)
			}
			online.BroadcastAlarm(msg, values, histId, true, curLevel)
			core.Logger.Infof("[iot处理] 告警恢复 deviceId=%s sn=%s", deviceId, msg.Sn)
		}
		return
	}

	// 已在事件中：只在该改写时写库重推，否则维持不动
	if histId != "" {
		if alertRewritable(curLevel, level) {
			if e := iot_service.AlertService.UpdateSeverity(histId, level); e != nil {
				core.Logger.Errorf("[iot处理] 告警级别改写落库失败 deviceId=%s id=%s err=%v", deviceId, histId, e)
			}
			_ = core.Redis.Set(ctx, alertActiveKey(deviceId), encodeActive(histId, level), 0).Err()
			online.BroadcastAlarm(msg, values, histId, false, level)
			core.Logger.Warnf("[iot处理] 告警级别变更 deviceId=%s sn=%s id=%s %d→%d", deviceId, msg.Sn, histId, curLevel, level)
		}
		return
	}
	// 静默期内抑制再次触发
	if silenced {
		return
	}
	// 未配持续时长：命中即触发
	if duration <= 0 {
		raiseEpisode(ctx, msg, values, duration, silence, level)
		return
	}
	// 配了持续时长：首次只记起点，后续报文跨过“起点+时长”才真正触发（级别取触发这一刻的返回值）
	if firstStr == "" {
		ttl := time.Duration(duration)*time.Second + alertFirstTTLBuffer
		if e := core.Redis.Set(ctx, alertFirstKey(deviceId), msg.ReceivedAt, ttl).Err(); e != nil {
			core.Logger.Errorf("[iot处理] 写告警持续计时失败 deviceId=%s err=%v", deviceId, e)
		}
		return
	}
	if first, e := strconv.ParseInt(firstStr, 10, 64); e == nil &&
		msg.ReceivedAt-first >= int64(duration)*1000 {
		raiseEpisode(ctx, msg, values, duration, silence, level)
	}
}

// alertRewritable 判定同一条活动事件是否该把 severity 改写为 level：
//   - 当前是故障而本次不再是：设备已不再失能，把记录交还实际命中级别，否则大屏挂着过期故障
//   - 其余比覆盖优先序（只升不降）：故障能接管任何在告事件，其期间不受阈值告警干扰
func alertRewritable(curLevel, level int8) bool {
	if curLevel == iot_model.AlertSevFault {
		return level != iot_model.AlertSevFault
	}
	return iot_model.AlertPriority(level) > iot_model.AlertPriority(curLevel)
}

// raiseEpisode 触发一次告警事件：落库→写活动态(含级别)→按需静默→清计时→推广播。
// 落库失败则不建活动态，下一条异常报文会重试触发。
func raiseEpisode(ctx context.Context, msg iot_schema.IngestMsg, values map[string]any, duration, silence int, level int8) {
	id, e := iot_service.AlertService.Raise(msg.DeviceId, msg.Sn, msg.ProductId, msg.SpaceId, values, msg.Ts, level)
	if e != nil {
		core.Logger.Errorf("[iot处理] 告警触发落库失败 deviceId=%s err=%v", msg.DeviceId, e)
		return
	}
	// 活动态无 TTL：只在恢复转换点或删设备时清理
	if e := core.Redis.Set(ctx, alertActiveKey(msg.DeviceId), encodeActive(id, level), 0).Err(); e != nil {
		core.Logger.Errorf("[iot处理] 写告警活动态失败 deviceId=%s err=%v", msg.DeviceId, e)
	}
	core.Redis.Del(ctx, alertFirstKey(msg.DeviceId))
	if silence > 0 {
		_ = core.Redis.Set(ctx, alertSilenceKey(msg.DeviceId), "1",
			time.Duration(silence)*time.Second).Err()
	}
	online.BroadcastAlarm(msg, values, id, false, level)
	core.Logger.Warnf("[iot处理] 告警触发 deviceId=%s sn=%s id=%s level=%d values=%v", msg.DeviceId, msg.Sn, id, level, values)
}

// strOf 取 MGet 结果里的字符串值（键不存在时元素为 nil）
func strOf(v any) string {
	if s, ok := v.(string); ok {
		return s
	}
	return ""
}
