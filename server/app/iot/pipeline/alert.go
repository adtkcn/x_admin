package pipeline

import (
	"context"
	"strconv"
	"strings"
	"time"
	"x_admin/app/schema/iot_schema"
	"x_admin/app/service/iot_service"
	"x_admin/core"
)

// 告警状态机（iot.md 第五阶段）：产品 goja 告警脚本 alert(values) 给出“本条告警级别”的数字判定
// （0=不告警/恢复，1/2/3=轻微/严重/故障），本文件把逐条判定收敛成“事件级”告警——
// 持续时长、静默、级别升级、恢复，落在 Redis 运行态 + MySQL 历史。
//
// 判据来自报文本身（异常=有数据），因此无需像离线判定那样另起到期表扫描：
// 每条上报顺带评估一次即可，运行态全部在 Redis，异常→恢复/升级的转换点才写一次库、推一次 WS，
// 持续期间的同级别重复命中不再落库/广播（避免高频报文把告警打成刷屏）。
//
// 单事件可升级：一台设备同时只维护一条“告警中”记录；级别升高就更新同条记录的 severity 并重推，
// 级别回落不降级（仍是同一条），仅当脚本返回 0 才结束本次事件。
//
// Redis 键（{RedisPrefix}iot:alert:*）：
//   - active:<deviceId>   存在=正处于一次告警事件，value="<历史主键>:<级别>"；恢复或删设备时删除，无 TTL
//   - first:<deviceId>    持续计时的起点(毫秒)，仅在配置了 alert_duration 时使用，TTL=时长+缓冲
//   - silence:<deviceId>  静默标记，TTL=alert_silence，触发时写入，抑制其后的再次触发

// alertFirstTTLBuffer 持续计时键在时长之外的冗余存活，避免边界抖动导致从头重新计时
const alertFirstTTLBuffer = 5 * time.Minute

func alertActiveKey(deviceId string) string  { return rtKey("alert:active:" + deviceId) }
func alertFirstKey(deviceId string) string   { return rtKey("alert:first:" + deviceId) }
func alertSilenceKey(deviceId string) string { return rtKey("alert:silence:" + deviceId) }

// encodeActive 活动态值编码为 "<histId>:<级别>"；uuid 主键不含冒号，解析按首个冒号切分安全
func encodeActive(histId string, level uint8) string {
	return histId + ":" + strconv.FormatUint(uint64(level), 10)
}

// parseActive 解出活动态里的历史主键与级别；空值返回 ("", 0)
func parseActive(raw string) (string, uint8) {
	if raw == "" {
		return "", 0
	}
	if i := strings.IndexByte(raw, ':'); i >= 0 {
		lvl, _ := strconv.ParseUint(raw[i+1:], 10, 8)
		return raw[:i], uint8(lvl)
	}
	return raw, 0 // 兼容旧格式（仅存主键、无级别）：级别按 0，下条异常报文会自然纠正
}

// evalAlert 依据本次告警级别推进状态机。level 为 alert(values) 的数字返回（0=不告警/恢复）；
// 仅在产品配置了告警脚本时调用（未配置则不参与告警，链路上零额外 Redis 开销）。
func evalAlert(msg iot_schema.IngestMsg, values map[string]any, level uint8, model iot_service.ProductModel) {
	ctx, cancel := redisCtx()
	defer cancel()

	// 三个运行态键一次 MGet 读出：分开读会有 3 个 RTT，且中间态可能被并发报文改写
	deviceId := msg.DeviceId
	rs, err := core.Redis.MGet(ctx, alertActiveKey(deviceId), alertFirstKey(deviceId), alertSilenceKey(deviceId)).Result()
	if err != nil {
		core.Logger.Errorf("[iot处理] 读告警运行态失败 deviceId=%s err=%v", deviceId, err)
		return
	}
	histId, curLevel := parseActive(strOf(rs[0]))
	firstStr := strOf(rs[1])
	silenced := len(rs) > 2 && rs[2] != nil

	// 恢复：清持续计时；若正处告警事件则收尾（置恢复 + 清活动态 + 推恢复）
	if level == 0 {
		core.Redis.Del(ctx, alertFirstKey(deviceId))
		if histId != "" {
			core.Redis.Del(ctx, alertActiveKey(deviceId))
			if e := iot_service.AlertService.Recover(histId, msg.ReceivedAt); e != nil {
				core.Logger.Errorf("[iot处理] 告警恢复落库失败 deviceId=%s id=%s err=%v", deviceId, histId, e)
			}
			pushAlarm(msg, values, histId, true, curLevel)
			core.Logger.Infof("[iot处理] 告警恢复 deviceId=%s sn=%s", deviceId, msg.Sn)
		}
		return
	}

	// 异常且已在事件中：级别升高才升级（更新同条记录 + 重推），否则维持不动（不逐条写库/广播）
	if histId != "" {
		if level > curLevel {
			if e := iot_service.AlertService.UpdateSeverity(histId, level); e != nil {
				core.Logger.Errorf("[iot处理] 告警升级落库失败 deviceId=%s id=%s err=%v", deviceId, histId, e)
			}
			_ = core.Redis.Set(ctx, alertActiveKey(deviceId), encodeActive(histId, level), 0).Err()
			pushAlarm(msg, values, histId, false, level)
			core.Logger.Warnf("[iot处理] 告警升级 deviceId=%s sn=%s id=%s %d→%d", deviceId, msg.Sn, histId, curLevel, level)
		}
		return
	}
	// 静默期内抑制再次触发
	if silenced {
		return
	}
	// 未配持续时长：命中即触发
	if model.AlertDuration <= 0 {
		raiseEpisode(ctx, msg, values, model, level)
		return
	}
	// 配了持续时长：首次只记起点，后续报文跨过“起点+时长”才真正触发（级别取触发这一刻的返回值）
	if firstStr == "" {
		ttl := time.Duration(model.AlertDuration)*time.Second + alertFirstTTLBuffer
		if e := core.Redis.Set(ctx, alertFirstKey(deviceId), msg.ReceivedAt, ttl).Err(); e != nil {
			core.Logger.Errorf("[iot处理] 写告警持续计时失败 deviceId=%s err=%v", deviceId, e)
		}
		return
	}
	if first, e := strconv.ParseInt(firstStr, 10, 64); e == nil &&
		msg.ReceivedAt-first >= int64(model.AlertDuration)*1000 {
		raiseEpisode(ctx, msg, values, model, level)
	}
}

// raiseEpisode 触发一次告警事件：落库(status=1,severity=level)→写 Redis 活动态(含级别)→按需静默→清计时→推广播。
// 落库失败则不建立活动态，本条不告警，下一条异常报文会重新尝试触发。
func raiseEpisode(ctx context.Context, msg iot_schema.IngestMsg, values map[string]any, model iot_service.ProductModel, level uint8) {
	id, e := iot_service.AlertService.Raise(msg.DeviceId, msg.Sn, msg.ProductId, msg.SpaceId, values, msg.Ts, level)
	if e != nil {
		core.Logger.Errorf("[iot处理] 告警触发落库失败 deviceId=%s err=%v", msg.DeviceId, e)
		return
	}
	// 活动态无 TTL：异常→恢复的转换点或删设备(StopRuntime)时才清理，语义上“未确认的活跃告警”应一直驻留
	if e := core.Redis.Set(ctx, alertActiveKey(msg.DeviceId), encodeActive(id, level), 0).Err(); e != nil {
		core.Logger.Errorf("[iot处理] 写告警活动态失败 deviceId=%s err=%v", msg.DeviceId, e)
	}
	core.Redis.Del(ctx, alertFirstKey(msg.DeviceId))
	if model.AlertSilence > 0 {
		_ = core.Redis.Set(ctx, alertSilenceKey(msg.DeviceId), "1",
			time.Duration(model.AlertSilence)*time.Second).Err()
	}
	pushAlarm(msg, values, id, false, level)
	core.Logger.Warnf("[iot处理] 告警触发 deviceId=%s sn=%s id=%s level=%d values=%v", msg.DeviceId, msg.Sn, id, level, values)
}

// strOf 取 MGet 结果里的字符串值（键不存在时元素为 nil）
func strOf(v any) string {
	if s, ok := v.(string); ok {
		return s
	}
	return ""
}
