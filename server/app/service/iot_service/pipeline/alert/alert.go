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
// 收敛成“事件级”告警（持续时长/级别变更/恢复确认），运行态在 Redis、历史在 MySQL。
// 判据来自报文本身，无需到期扫描：每条上报顺带评估一次，只在转换点写库推 WS，
// 持续期间的重复命中不落库不广播。
//
// 两条独立事件流：故障(-1)与普通阈值告警(1/2/…)各自维护一条“告警中”记录，互不合并、可并存。
// 每条报文只推进其一：level==-1 只确认故障流、level>0 只确认普通流、level==0 才同时向两条流传恢复信号；
// 切到另一种类型不误判为恢复（NOOP），各自独立按恢复确认时长收尾。
//
// 记录只记最高级别：普通流的 severity 取事件内命中过的最高档（级别数值即危害序，只升不降）；
// 故障流只有单一级别，天然不变。二者都推 WS 当前级别（可升可降），故历史记录留峰值、
// 大屏实时值随当前级别升降，二者解耦；当前级别未变的重复命中不落库/广播。
//
// 恢复确认（可配 recover，秒）：报警后不会因单条恢复信号立即收尾，需恢复正常连续
// 满 recover 秒才关闭；期间再命中该流则计时作废、事件延续。recover=0 则恢复信号即收尾。
//
// 持续时长（可配 duration，秒）：命中后需连续异常满 duration 秒才触发；duration=0 则命中即告。
//
// Redis 键（{RedisPrefix}iot:alert:*，普通流无前缀、故障流 fault: 前缀，两组各三键）：
//   - [fault:]active:<deviceId>    存在=正处于该流事件，value="<历史主键>:<记录峰值级别>:<当前级别>"，无 TTL
//   - [fault:]first:<deviceId>     持续计时起点(毫秒)，仅 duration>0 时使用，TTL=时长+缓冲
//   - [fault:]recover:<deviceId>   恢复确认计时起点(毫秒)，仅 recover>0 时使用，TTL=时长+缓冲

// alertTimerTTLBuffer 持续/恢复计时键在时长之外的冗余存活，避免边界抖动导致从头重新计时
const alertTimerTTLBuffer = 5 * time.Minute

// init 经 online 钩子把告警收尾接出去（回调而非 import 上层，避免 online ↔ alert 循环依赖）
func init() {
	online.OnDeviceOffline(closeOnOffline) // 设备转离线：收尾其活动告警
	online.OnDeviceClear(ClearRuntime)     // 删设备：一并清告警 Redis 键
}

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

// allKeys 列出设备的全部告警运行态键（普通流 + 故障流各三键），供清态一次 Del
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

// ClearRuntime 只清设备告警运行态（普通流+故障流共六键），不做库侧收尾。活动态无 TTL 且 value 就是
// 历史记录主键，残留会让该设备后续命中永远不再新建记录，所以删库后必调。
func ClearRuntime(deviceId string) {
	ctx, cancel := online.RedisCtx()
	defer cancel()
	_ = core.Redis.Del(ctx, allKeys(deviceId)...).Err()
}

// closeOnOffline 告警状态机由报文驱动，设备停报后永远等不到恢复转换，故离线即收尾（两条流一并置恢复 + 清键），
// 重新上线且条件命中会另起新事件。
func closeOnOffline(deviceId string) {
	ctx, cancel := online.RedisCtx()
	defer cancel()
	// 先确认有活动态才发 UPDATE：离线钩子对每台设备都跑，无告警时那条 UPDATE 永远命中 0 行
	rs, _ := core.Redis.MGet(ctx, alertActiveKey(deviceId), alertFaultActiveKey(deviceId)).Result()
	if strOf(rs[0]) != "" || strOf(rs[1]) != "" {
		if e := iot_service.AlertService.CloseActiveByDevice(deviceId); e != nil {
			core.Logger.Errorf("[iot处理] 离线收尾活动告警失败 deviceId=%s err=%v", deviceId, e)
		}
	}
	_ = core.Redis.Del(ctx, allKeys(deviceId)...).Err()
}

// encodeActive 活动态值编码为 "<histId>:<记录峰值级别>:<当前级别>"；uuid 主键不含冒号，按冒号切分安全
func encodeActive(histId string, peak, cur int8) string {
	return histId + ":" + strconv.FormatInt(int64(peak), 10) + ":" + strconv.FormatInt(int64(cur), 10)
}

// parseActive 解出活动态里的历史主键、记录峰值级别、当前级别；空值或格式异常返回 ("", 0, 0)。
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

// EvalAlert 按本次告警级别推进两条事件流。level 为告警节点的数字返回（0=不告警/恢复，-1=故障，1/2/…=阈值告警）；
// duration/recoverDur 取自告警节点属性（秒，0=命中即告/恢复即收）；
// allow 为物模型 identifier 白名单，告警记录与广播的 values 落库前先按其裁剪（allow 空则不约束）。
func EvalAlert(msg iot_schema.IngestMsg, values map[string]any, level int8, duration, recoverDur int, allow map[string]bool) {
	values = iot_schema.ScopeValues(values, allow)
	ctx, cancel := online.RedisCtx()
	defer cancel()

	nk, fk := normalKeys(msg.DeviceId), faultKeys(msg.DeviceId)
	// 两组共六键一次 MGet 读出：分开读是多个 RTT，且中间态可能被并发报文改写
	rs, err := core.Redis.MGet(ctx,
		nk.active, nk.first, nk.recover,
		fk.active, fk.first, fk.recover).Result()
	if err != nil {
		core.Logger.Errorf("[iot处理] 读告警运行态失败 deviceId=%s err=%v", msg.DeviceId, err)
		return
	}

	// level==0 是两条流共同的恢复信号；level==-1 只确认故障流、level>0 只确认普通流，
	// 另一条流不受影响（NOOP）——故故障与普通告警可各自独立共存、各自恢复，
	// 不会因为切到另一种类型就被误判为恢复。
	vn := streamVals{strOf(rs[0]), strOf(rs[1]), strOf(rs[2])}
	vf := streamVals{strOf(rs[3]), strOf(rs[4]), strOf(rs[5])}
	switch {
	case level == 0:
		recoverStream(ctx, msg, values, nk, vn, recoverDur)
		recoverStream(ctx, msg, values, fk, vf, recoverDur)
	case level == iot_model.AlertSevFault:
		confirmStream(ctx, msg, values, fk, vf, iot_model.AlertSevFault, duration)
	default: // level>0：普通阈值告警，级别即本次返回值
		confirmStream(ctx, msg, values, nk, vn, level, duration)
	}
}

// confirmStream 本条报文确认该流处于异常：作废恢复计时；已在事件中则按“记录取峰值、WS 推当前级”
// 推进，否则按 duration 决定是否新触发。故障流级别恒为 -1，故峰值/当前不变、无升级与重推。
func confirmStream(ctx context.Context, msg iot_schema.IngestMsg, values map[string]any,
	k streamKeys, v streamVals, level int8, duration int) {

	// 恢复确认计时作废（设备重回该流告警态，之前的“疑似恢复”不算数）
	if v.recover != "" {
		core.Redis.Del(ctx, k.recover)
	}

	histId, peak, cur := parseActive(v.active)
	// 已在事件中：记录只升不降取峰值（level>peak），WS 恒推当前级别（level!=cur）——任一变化才落库/重推
	if histId != "" {
		raise := level > peak
		notify := level != cur
		if !raise && !notify {
			return
		}
		newPeak, newCur := peak, cur
		if raise {
			newPeak = level
			if e := iot_service.AlertService.UpdateSeverity(histId, level); e != nil {
				core.Logger.Errorf("[iot处理] 告警级别改写落库失败 deviceId=%s id=%s err=%v", msg.DeviceId, histId, e)
			}
		}
		if notify {
			newCur = level
			online.BroadcastAlarm(msg, values, histId, false, level)
		}
		_ = core.Redis.Set(ctx, k.active, encodeActive(histId, newPeak, newCur), 0).Err()
		return
	}

	// 新事件：未配持续时长即触发；配了则首次记起点，后续报文跨过“起点+时长”才真正触发
	if duration <= 0 {
		raiseEpisode(ctx, msg, values, k, level)
		return
	}
	if v.first == "" {
		ttl := time.Duration(duration)*time.Second + alertTimerTTLBuffer
		if e := core.Redis.Set(ctx, k.first, msg.ReceivedAt, ttl).Err(); e != nil {
			core.Logger.Errorf("[iot处理] 写告警持续计时失败 deviceId=%s err=%v", msg.DeviceId, e)
		}
		return
	}
	if start, e := strconv.ParseInt(v.first, 10, 64); e == nil &&
		msg.ReceivedAt-start >= int64(duration)*1000 {
		raiseEpisode(ctx, msg, values, k, level)
	}
}

// recoverStream 本条报文（level==0）为该流的恢复信号：清持续计时；正处事件则按 recoverDur
// 决定是否收尾（首次只记起点，待后续恢复报文跨过“起点+确认时长”才真正关闭）。
func recoverStream(ctx context.Context, msg iot_schema.IngestMsg, values map[string]any,
	k streamKeys, v streamVals, recoverDur int) {

	// 计时键存的是毫秒起点、值永不为空串：读到才删，避免稳态每条报文空转 Del
	if v.first != "" {
		core.Redis.Del(ctx, k.first)
	}
	histId, _, cur := parseActive(v.active)
	if histId == "" {
		if v.recover != "" {
			core.Redis.Del(ctx, k.recover)
		}
		return
	}
	if recoverDur <= 0 { // 未配恢复确认时长：恢复信号即收尾
		recoverEpisode(ctx, msg, values, k, histId, cur)
		return
	}
	if v.recover == "" { // 首次收到恢复：只记起点
		ttl := time.Duration(recoverDur)*time.Second + alertTimerTTLBuffer
		if e := core.Redis.Set(ctx, k.recover, msg.ReceivedAt, ttl).Err(); e != nil {
			core.Logger.Errorf("[iot处理] 写告警恢复计时失败 deviceId=%s err=%v", msg.DeviceId, e)
		}
		return
	}
	if start, e := strconv.ParseInt(v.recover, 10, 64); e == nil &&
		msg.ReceivedAt-start >= int64(recoverDur)*1000 {
		recoverEpisode(ctx, msg, values, k, histId, cur)
	}
}

// recoverEpisode 结束一次告警事件：置恢复 + 清该流活动/恢复键 + 推恢复广播（severity 推当前级别）。
func recoverEpisode(ctx context.Context, msg iot_schema.IngestMsg, values map[string]any, k streamKeys, histId string, severity int8) {
	core.Redis.Del(ctx, k.active, k.recover)
	if e := iot_service.AlertService.Recover(histId, msg.ReceivedAt); e != nil {
		core.Logger.Errorf("[iot处理] 告警恢复落库失败 deviceId=%s id=%s err=%v", msg.DeviceId, histId, e)
	}
	online.BroadcastAlarm(msg, values, histId, true, severity)
	core.Logger.Infof("[iot处理] 告警恢复 deviceId=%s sn=%s id=%s", msg.DeviceId, msg.Sn, histId)
}

// raiseEpisode 触发一次告警事件：落库→写该流活动态(含级别)→清持续计时→推广播。
// 落库失败则不建活动态，下一条命中报文会重试触发。
func raiseEpisode(ctx context.Context, msg iot_schema.IngestMsg, values map[string]any, k streamKeys, level int8) {
	id, e := iot_service.AlertService.Raise(msg.DeviceId, msg.Sn, msg.ProductId, msg.SpaceId, values, msg.Ts, level)
	if e != nil {
		core.Logger.Errorf("[iot处理] 告警触发落库失败 deviceId=%s err=%v", msg.DeviceId, e)
		return
	}
	// 活动态无 TTL：只在恢复转换点或删设备时清理
	if e := core.Redis.Set(ctx, k.active, encodeActive(id, level, level), 0).Err(); e != nil {
		core.Logger.Errorf("[iot处理] 写告警活动态失败 deviceId=%s err=%v", msg.DeviceId, e)
	}
	core.Redis.Del(ctx, k.first)
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
