package alert

import (
	"context"
	"strconv"
	"time"

	"x_admin/app/model/iot_model"
	"x_admin/app/schema/iot_schema"
	"x_admin/app/service/iot_service"
	"x_admin/app/service/iot_service/pipeline/online"
	"x_admin/core"
)

// 告警状态机：把告警脚本 run(values, sn) 的逐条判定（0=恢复，非0=级别）收敛为「事件级」告警，
// 活动态在 Redis、历史在 MySQL，报文驱动，仅在转换点写库推 WS。
//
// 故障流(-1)与阈值流(1/2/…)各自独立维护一条「告警中」记录、严格分离：每条报文只确认其一，
// 另一条即收恢复信号。记录留峰值（当前级只升不降）。大屏只展示聚合级（故障优先，否则普通流），
// 仅聚合级变化才推 WS，故两流并存的延迟恢复期靠聚合避免误清/降级。
//
// 活动态仅存 Redis 且无 TTL，是状态机唯一事实源；Redis 丢失会令库中告警中记录失去收尾依据，
// 故启动时 RecoverRuntime 按库回灌（见 runtime.go）。duration/recover 为可配时长（秒，0=立即）。
// 键定义与编解码见 keys.go。

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

	// 故障与报警严格分离：命中的流 confirm，另一条流一律发恢复信号（走 recover 确认时长）
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

	// WS 推设备聚合级（非单流级）：仅当聚合级变化才推，避免并存期误清或级别错乱
	if id, recovered, sev := deviceAlarm(nAct, fAct); recovered != prevRecovered || sev != prevSev {
		if id == "" {
			id = prevId
		}
		online.BroadcastAlarm(msg, values, id, recovered, sev)
	}
}

// deviceAlarm 求设备当前聚合告警态（供 WS）：故障优先(-1)，否则取普通流当前级别；两流皆无则为恢复。
// 与 MapList 取数口径一致；延迟恢复期两流并存，必须按聚合级而非单流级推送。
func deviceAlarm(normalActive, faultActive string) (id string, recovered bool, severity int8) {
	if fid, _, _ := parseActive(faultActive); fid != "" {
		return fid, false, iot_model.AlertSevFault
	}
	if nid, _, cur := parseActive(normalActive); nid != "" {
		return nid, false, cur
	}
	return "", true, 0
}

// confirmStream 本条报文确认该流异常：作废恢复计时；已在事件中则按“记录取峰值”更新当前级，
// 否则按 duration 决定是否新触发。返回处理后的活动态值（供 EvalAlert 算聚合级）。故障流级别恒为 -1。
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
// 返回处理后的活动态值（收尾则空串，供 EvalAlert 算聚合级）。
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

// recoverEpisode 结束一次告警事件：清该流活动/恢复键并置库恢复（WS 由 EvalAlert 按聚合级统一推）
func recoverEpisode(ctx context.Context, msg iot_schema.IngestMsg, k streamKeys, histId string) {
	core.Redis.Del(ctx, k.active, k.recover)
	if e := iot_service.AlertService.Recover(histId, msg.ReceivedAt); e != nil {
		core.Logger.Errorf("[iot处理] 告警恢复落库失败 deviceId=%s id=%s err=%v", msg.DeviceId, histId, e)
	}
	core.Logger.Infof("[iot处理] 告警恢复 deviceId=%s sn=%s id=%s", msg.DeviceId, msg.Sn, histId)
}

// raiseEpisode 触发一次告警事件：落库→写该流活动态(含级别)→清持续计时；落库失败返回空不建活动态。
// 返回新建的活动态值（供 EvalAlert 算聚合级；WS 统一由 EvalAlert 推）
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
