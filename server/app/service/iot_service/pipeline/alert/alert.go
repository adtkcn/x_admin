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
	"x_admin/config"
	"x_admin/core"
)

// 告警状态机：把告警节点脚本 alert(values, sn) 的逐条判定（0=恢复，非 0 即级别）
// 收敛成“事件级”告警，运行态在 Redis、历史在 MySQL。报文驱动，只在转换点写库推 WS。
//
// 两条独立事件流（故障 -1 / 阈值告警 1/2/…）各自维护一条“告警中”记录，互不合并、严格分离：
// 每条报文只确认其一，另一条流即收恢复信号，按恢复确认时长(recover)收尾，确认期间切回则计时作废。
//
// 记录留峰值（普通流 severity 只升不降）。大屏设备点只显示一条聚合告警级，故 WS 推「设备聚合级」
// （故障优先 -1，否则普通流当前级别），仅当聚合级变化才推；延迟恢复期两条流并存时靠聚合避免误清/降级。
//
// 活动态只存 Redis，是「告警中」的唯一事实源：Redis 数据丢失（重启未持久化/驱逐）会令库里
// status=1 的记录失去收尾依据（恢复报文读不到活动态即被跳过），故进程启动、消费开始前调
// RecoverRuntime 按库回灌活动态键，让后续报文能正常恢复。
//
// 可配时长（秒，0=命中即告/恢复即收）：duration 命中后需连续异常满时长才触发；recover 恢复需连续满时长才关闭。
//
// Redis 键（{RedisPrefix}iot:alert:*，普通流无前缀、故障流 fault: 前缀，两组各三键）：
//   - [fault:]active:<deviceId>    存在=正处于该流事件，value="<历史主键>:<峰值级别>:<当前级别>"，无 TTL
//   - [fault:]first:<deviceId>     持续计时起点(毫秒)，仅 duration>0 时使用，TTL=时长+缓冲
//   - [fault:]recover:<deviceId>   恢复确认计时起点(毫秒)，仅 recover>0 时使用，TTL=时长+缓冲

// alertTimerTTLBuffer 持续/恢复计时键在时长外的冗余存活，避免边界抖动从头计时
const alertTimerTTLBuffer = 5 * time.Minute

// init 经 online 钩子接出告警收尾（回调用避免 online ↔ alert 循环依赖）
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

// allKeys 设备的全部告警运行态键（普通流+故障流共六键），供清态一次 Del
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

// RecoverRuntime 服务重启后重建告警运行态：活动态只存 Redis（无 TTL）却是状态机的唯一事实源，
// 若 Redis 数据丢失（重启未持久化/内存驱逐）而库里仍有 status=1 的「告警中」记录，
// 恢复报文读不到活动态即被 recoverStream 当 NOOP 跳过，记录会永久卡在告警中无法收尾。
// 故须在消费开始前扫一遍库：按(设备,流)回灌活动态键，使后续报文能正常恢复；同一(设备,流)
// 只留最新一条为活动态，更早的悬空记录（本缺陷历史遗留）一并收尾。
func RecoverRuntime() {
	if !config.IotConfig.Enable {
		return
	}
	var rows []iot_model.IotAlertHistory
	// 按设备+触发时间升序遍历：同一(设备,流)后读到的覆盖先读到的，天然留下最新一条
	if e := core.GetDB().Model(&iot_model.IotAlertHistory{}).
		Select("id, device_id, severity, triggered_at").
		Where("status = ?", iot_model.AlertStatusFiring).
		Order("device_id ASC, triggered_at ASC").
		Find(&rows).Error; e != nil {
		core.Logger.Errorf("[iot处理] 重建告警运行态查询失败 err=%v", e)
		return
	}
	if len(rows) == 0 {
		return
	}

	type stream struct {
		deviceId string
		fault    bool
	}
	keep := make(map[stream]iot_model.IotAlertHistory, len(rows))
	var stale []string
	for _, r := range rows {
		k := stream{r.DeviceId, r.Severity == iot_model.AlertSevFault}
		if prev, ok := keep[k]; ok {
			stale = append(stale, prev.ID) // 被更新的记录顶掉，旧记录收尾
		}
		keep[k] = r
	}

	now := time.Now().UnixMilli()
	for _, id := range stale {
		if e := iot_service.AlertService.Recover(id, now); e != nil {
			core.Logger.Errorf("[iot处理] 重建时收尾悬空告警失败 id=%s err=%v", id, e)
		}
	}

	ctx, cancel := online.RedisCtx()
	defer cancel()
	pipe := core.Redis.Pipeline()
	for k, r := range keep {
		// 库里 severity 即历史峰值（只升不降），重建后当前级取同值
		val := encodeActive(r.ID, r.Severity, r.Severity)
		if k.fault {
			pipe.Set(ctx, alertFaultActiveKey(k.deviceId), val, 0)
		} else {
			pipe.Set(ctx, alertActiveKey(k.deviceId), val, 0)
		}
	}
	if _, e := pipe.Exec(ctx); e != nil {
		core.Logger.Errorf("[iot处理] 重建告警运行态回灌失败 size=%d err=%v", len(keep), e)
		return
	}
	core.Logger.Infof("[iot处理] 重启重建告警运行态 活动=%d 收尾悬空=%d", len(keep), len(stale))
}

// ClearRuntime 清设备告警运行态（普通流+故障流共六键），不做库侧收尾；删库后必调。
func ClearRuntime(deviceId string) {
	ctx, cancel := online.RedisCtx()
	defer cancel()
	_ = core.Redis.Del(ctx, allKeys(deviceId)...).Err()
}

// closeOnOffline 设备离线即收尾两条流的活动告警并清键（报文驱动的状态机停报后等不到恢复转换）。
func closeOnOffline(deviceId string) {
	ctx, cancel := online.RedisCtx()
	defer cancel()
	// 有活动态才发 UPDATE，避免无告警时空跑
	rs, _ := core.Redis.MGet(ctx, alertActiveKey(deviceId), alertFaultActiveKey(deviceId)).Result()
	if strOf(rs[0]) != "" || strOf(rs[1]) != "" {
		if e := iot_service.AlertService.CloseActiveByDevice(deviceId); e != nil {
			core.Logger.Errorf("[iot处理] 离线收尾活动告警失败 deviceId=%s err=%v", deviceId, e)
		}
	}
	_ = core.Redis.Del(ctx, allKeys(deviceId)...).Err()
}

// encodeActive 活动态值编码为 "<histId>:<峰值级别>:<当前级别>"
func encodeActive(histId string, peak, cur int8) string {
	return histId + ":" + strconv.FormatInt(int64(peak), 10) + ":" + strconv.FormatInt(int64(cur), 10)
}

// parseActive 解出活动态里的历史主键、峰值级别、当前级别；空值或格式异常返回 ("", 0, 0)。
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

// deviceAlarm 求设备当前聚合告警态（供 WS）：故障优先(-1)，否则取普通流当前级别；两流皆无则为恢复。
// 与 MapList 取数口径一致；延迟恢复期两条流并存，必须按聚合级而非单流级推送。
func deviceAlarm(normalActive, faultActive string) (id string, recovered bool, severity int8) {
	if fid, _, _ := parseActive(faultActive); fid != "" {
		return fid, false, iot_model.AlertSevFault
	}
	if nid, _, cur := parseActive(normalActive); nid != "" {
		return nid, false, cur
	}
	return "", true, 0
}

// EvalAlert 按本次告警级别推进两条事件流。level：0=恢复，-1=故障，1/2/…=阈值告警；
// duration/recoverDur 为告警节点属性（秒）；allow 为物模型 identifier 白名单，落库前裁剪 values。
func EvalAlert(msg iot_schema.IngestMsg, values map[string]any, level int8, duration, recoverDur int, allow map[string]bool) {
	values = iot_schema.ScopeValues(values, allow)
	ctx, cancel := online.RedisCtx()
	defer cancel()

	nk, fk := normalKeys(msg.DeviceId), faultKeys(msg.DeviceId)
	// 两组共六键一次 MGet 读出，减少 RTT 并避免中间态被并发报文改写
	rs, err := core.Redis.MGet(ctx,
		nk.active, nk.first, nk.recover,
		fk.active, fk.first, fk.recover).Result()
	if err != nil {
		core.Logger.Errorf("[iot处理] 读告警运行态失败 deviceId=%s err=%v", msg.DeviceId, err)
		return
	}

	// 故障与报警严格分离：命中的流 confirm，另一条流一律发恢复信号（走 recover 确认时长）。
	vn := streamVals{strOf(rs[0]), strOf(rs[1]), strOf(rs[2])}
	vf := streamVals{strOf(rs[3]), strOf(rs[4]), strOf(rs[5])}
	prevId, prevRecovered, prevSev := deviceAlarm(vn.active, vf.active)

	var nAct, fAct string
	switch {
	case level == 0:
		nAct = recoverStream(ctx, msg, nk, vn, recoverDur)
		fAct = recoverStream(ctx, msg, fk, vf, recoverDur)
	case level == iot_model.AlertSevFault:
		fAct = confirmStream(ctx, msg, values, fk, vf, iot_model.AlertSevFault, duration)
		nAct = recoverStream(ctx, msg, nk, vn, recoverDur)
	default: // level>0：普通阈值告警，级别即本次返回值
		nAct = confirmStream(ctx, msg, values, nk, vn, level, duration)
		fAct = recoverStream(ctx, msg, fk, vf, recoverDur)
	}

	// WS 推设备聚合级（非单流级）：仅当聚合级变化才推，避免并存期误清或级别错乱。
	if id, recovered, sev := deviceAlarm(nAct, fAct); recovered != prevRecovered || sev != prevSev {
		if id == "" {
			id = prevId
		}
		online.BroadcastAlarm(msg, values, id, recovered, sev)
	}
}

// confirmStream 本条报文确认该流异常：作废恢复计时；已在事件中则按“记录取峰值”更新当前级，
// 否则按 duration 决定是否新触发。返回该流处理后的活动态值（供 EvalAlert 算聚合级）。故障流级别恒为 -1。
func confirmStream(ctx context.Context, msg iot_schema.IngestMsg, values map[string]any,
	k streamKeys, v streamVals, level int8, duration int) string {

	// 重回告警态：恢复确认计时作废
	if v.recover != "" {
		core.Redis.Del(ctx, k.recover)
	}

	histId, peak, cur := parseActive(v.active)
	// 已在事件中：level>peak 升峰值落库，level!=cur 刷新当前级；都没变则原样返回
	if histId != "" {
		if level <= peak && level == cur {
			return v.active
		}
		newPeak, newCur := peak, cur
		if level > peak {
			newPeak = level
			if e := iot_service.AlertService.UpdateSeverity(histId, level); e != nil {
				core.Logger.Errorf("[iot处理] 告警级别改写落库失败 deviceId=%s id=%s err=%v", msg.DeviceId, histId, e)
			}
		}
		if level != cur {
			newCur = level
		}
		na := encodeActive(histId, newPeak, newCur)
		_ = core.Redis.Set(ctx, k.active, na, 0).Err()
		return na
	}

	// 新事件：duration<=0 即触发；否则首次记起点，后续报文跨过“起点+时长”才触发
	if duration <= 0 {
		return raiseEpisode(ctx, msg, values, k, level)
	}
	if v.first == "" {
		ttl := time.Duration(duration)*time.Second + alertTimerTTLBuffer
		if e := core.Redis.Set(ctx, k.first, msg.ReceivedAt, ttl).Err(); e != nil {
			core.Logger.Errorf("[iot处理] 写告警持续计时失败 deviceId=%s err=%v", msg.DeviceId, e)
		}
		return v.active
	}
	if start, e := strconv.ParseInt(v.first, 10, 64); e == nil &&
		msg.ReceivedAt-start >= int64(duration)*1000 {
		return raiseEpisode(ctx, msg, values, k, level)
	}
	return v.active
}

// recoverStream 本条报文为该流的恢复信号：清持续计时；正处事件则按 recoverDur 决定是否收尾。
// 返回该流处理后的活动态值（收尾则空串，供 EvalAlert 算聚合级）。
func recoverStream(ctx context.Context, msg iot_schema.IngestMsg,
	k streamKeys, v streamVals, recoverDur int) string {

	// 读到才删，避免稳态空转 Del
	if v.first != "" {
		core.Redis.Del(ctx, k.first)
	}
	histId, _, _ := parseActive(v.active)
	if histId == "" {
		if v.recover != "" {
			core.Redis.Del(ctx, k.recover)
		}
		return v.active
	}
	if recoverDur <= 0 { // 未配恢复确认时长：恢复信号即收尾
		recoverEpisode(ctx, msg, k, histId)
		return ""
	}
	if v.recover == "" { // 首次收到恢复：只记起点，事件延续
		ttl := time.Duration(recoverDur)*time.Second + alertTimerTTLBuffer
		if e := core.Redis.Set(ctx, k.recover, msg.ReceivedAt, ttl).Err(); e != nil {
			core.Logger.Errorf("[iot处理] 写告警恢复计时失败 deviceId=%s err=%v", msg.DeviceId, e)
		}
		return v.active
	}
	if start, e := strconv.ParseInt(v.recover, 10, 64); e == nil &&
		msg.ReceivedAt-start >= int64(recoverDur)*1000 {
		recoverEpisode(ctx, msg, k, histId)
		return ""
	}
	return v.active
}

// recoverEpisode 结束一次告警事件：清该流活动/恢复键并置库恢复（WS 由 EvalAlert 按聚合级统一推）。
func recoverEpisode(ctx context.Context, msg iot_schema.IngestMsg, k streamKeys, histId string) {
	core.Redis.Del(ctx, k.active, k.recover)
	if e := iot_service.AlertService.Recover(histId, msg.ReceivedAt); e != nil {
		core.Logger.Errorf("[iot处理] 告警恢复落库失败 deviceId=%s id=%s err=%v", msg.DeviceId, histId, e)
	}
	core.Logger.Infof("[iot处理] 告警恢复 deviceId=%s sn=%s id=%s", msg.DeviceId, msg.Sn, histId)
}

// raiseEpisode 触发一次告警事件：落库→写该流活动态(含级别)→清持续计时；落库失败返回空不建活动态。
// 返回新建的活动态值（供 EvalAlert 算聚合级；WS 统一由 EvalAlert 推）。
func raiseEpisode(ctx context.Context, msg iot_schema.IngestMsg, values map[string]any, k streamKeys, level int8) string {
	id, e := iot_service.AlertService.Raise(msg.DeviceId, msg.Sn, msg.ProductId, msg.SpaceId, values, msg.Ts, level)
	if e != nil {
		core.Logger.Errorf("[iot处理] 告警触发落库失败 deviceId=%s err=%v", msg.DeviceId, e)
		return ""
	}
	// 活动态无 TTL，只在恢复转换点或删设备时清理
	na := encodeActive(id, level, level)
	if e := core.Redis.Set(ctx, k.active, na, 0).Err(); e != nil {
		core.Logger.Errorf("[iot处理] 写告警活动态失败 deviceId=%s err=%v", msg.DeviceId, e)
	}
	core.Redis.Del(ctx, k.first)
	core.Logger.Warnf("[iot处理] 告警触发 deviceId=%s sn=%s id=%s level=%d values=%v", msg.DeviceId, msg.Sn, id, level, values)
	return na
}

// strOf 取 MGet 结果里的字符串值（键不存在时元素为 nil）
func strOf(v any) string {
	if s, ok := v.(string); ok {
		return s
	}
	return ""
}
