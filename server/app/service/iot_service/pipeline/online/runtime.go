package online

import (
	"context"
	"encoding/json"
	"strconv"
	"time"
	"x_admin/app/schema/iot_schema"
	"x_admin/app/service/iot_service"
	"x_admin/config"
	"x_admin/core"

	"github.com/redis/go-redis/v9"
)

// online 子包：设备运行态（在线/离线判定、最新属性快照、状态与遥测广播）。
// 运行态一律放 Redis（元数据在 MySQL、时序在存储层），键统一为 {RedisPrefix}iot:<用途>:<标识>。
// 在线方向由报文实时驱动，离线方向由到期表（ZSET）驱动。
// 本包不反向依赖告警：离线/删设备的告警收尾经 hooks.go 由上层注入。

// RuntimeKey 拼接运行态键（导出供 alert 等同级子包共用同一命名空间）
func RuntimeKey(suffix string) string {
	return config.RedisConfig.RedisPrefix + "iot:" + suffix
}

// expireZKey 到期表（ZSET）：member=设备id，score=离线到期时间(毫秒)。
// 在线判定的唯一事实源，一份数据同时表达“谁在线”和“何时到期”；
// score 只在收到报文时刷新，改窗口后的瞬时误差由下一条报文/慢路径自愈。
func expireZKey() string { return RuntimeKey("expire:zset") }

// stateKey 最新属性快照：HASH，field=identifier，value=字符串化的值
func stateKey(deviceId string) string { return RuntimeKey("state:" + deviceId) }

// seenKey 最后一条报文的服务端接收时间(毫秒)，供管理端回读与慢路径重建到期表；
// TTL 必须长于任何可能的判定窗口（见 keepAlive）
func seenKey(deviceId string) string { return RuntimeKey("seen:" + deviceId) }

// RedisCtx 统一给 Redis 操作加超时，避免设备洪峰下 worker 卡死在慢命令上
func RedisCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 2*time.Second)
}

// offlineWindow 设备生效的离线判定窗口：直接等于设备配置的 offline_timeout（秒，已冗余产品值），
// 无宽限期；未配置时落 iot_service 的兜底值
func offlineWindow(timeout int) time.Duration {
	if timeout <= 0 {
		timeout = iot_service.DefaultOfflineTimeout
	}
	return time.Duration(timeout) * time.Second
}

// keepAlive 刷新到期表 score（=now+窗口），返回刷新前是否无有效 score（true=离线→在线）。
// 上下线两个方向都由到期表驱动，不在此处复核旧 score。
func keepAlive(deviceId string, timeout int) (wasAbsent bool) {
	ctx, cancel := RedisCtx()
	defer cancel()
	zkey := expireZKey()
	window := offlineWindow(timeout)
	// 时间戳必须比判定窗口活得更久：否则长窗设备会在窗口内先丢掉时间戳而被慢路径误判离线
	seenTTL := time.Duration(config.IotConfig.StateTTL) * time.Second
	if seenTTL < window {
		seenTTL = window + time.Minute
	}
	now := time.Now().UnixMilli()
	pipe := core.Redis.Pipeline()
	score := pipe.ZScore(ctx, zkey, deviceId)
	pipe.ZAdd(ctx, zkey, redis.Z{Score: float64(now + window.Milliseconds()), Member: deviceId})
	pipe.Set(ctx, seenKey(deviceId), now, seenTTL)
	if _, err := pipe.Exec(ctx); err != nil && err != redis.Nil {
		core.Logger.Errorf("[iot处理] 刷新到期表失败 deviceId=%s err=%v", deviceId, err)
		// Redis 异常时保守报“未发生转换”：否则不可用的每秒每条报文都会 SELECT+UPDATE 设备表，
		// 在依赖故障时反而拖垮库；标记恢复后首条报文仍会正常判为上线
		return false
	}
	if sc, err := score.Result(); err == nil {
		return sc <= float64(now) // score 已到期：这条报文算离线→在线
	}
	if err := score.Err(); err == nil || err == redis.Nil {
		return true // 无成员（从未上报/到期被弹/数据丢）：同样算离线→在线
	}
	return false // ZScore 真故障：按上面的保守策略处理
}

// cancelDeadline 从到期表撤销设备记录（上线转换写库失败后调）：
// 下一条报文会因“score 缺失”重新触发上线转换，重试搭在报文流上，不需额外补判扫描
func cancelDeadline(deviceId string) {
	ctx, cancel := RedisCtx()
	defer cancel()
	if err := core.Redis.ZRem(ctx, expireZKey(), deviceId).Err(); err != nil {
		core.Logger.Errorf("[iot处理] 撤销到期记录失败 deviceId=%s err=%v", deviceId, err)
	}
}

// RefreshOnline 刷新在线标记，并在“离线→在线”转换时写库+广播（心跳/上报都走这里）
func RefreshOnline(msg iot_schema.IngestMsg) {
	if keepAlive(msg.DeviceId, msg.OfflineTimeout) {
		changed, err := iot_service.DeviceService.MarkStatus(msg.DeviceId, 1)
		if err != nil {
			core.Logger.Errorf("[iot处理] 置在线失败 deviceId=%s err=%v", msg.DeviceId, err)
			// 写库失败撤销刚写的到期 score，重试由下一条报文完成；撤销失败则退到慢路径到期后自愈
			cancelDeadline(msg.DeviceId)
		} else if changed {
			broadcastStatus(msg, 1)
		}
	}
}

// lastSeenAndDeadlines 一次 Redis 往返批量读多台设备的“最后报文时间+到期表 score”。
// 取不到的 id 不出现在对应返回 map 里，调用方以“缺键”区分未设置；Redis 真故障才返回错误。
func lastSeenAndDeadlines(deviceIds []string) (seen map[string]int64, deadline map[string]float64, e error) {
	seen = make(map[string]int64, len(deviceIds))
	deadline = make(map[string]float64, len(deviceIds))
	if len(deviceIds) == 0 {
		return
	}
	ctx, cancel := RedisCtx()
	defer cancel()
	pipe := core.Redis.Pipeline()
	gets := make([]*redis.StringCmd, len(deviceIds))
	scores := make([]*redis.FloatCmd, len(deviceIds))
	for i, id := range deviceIds {
		gets[i] = pipe.Get(ctx, seenKey(id))
		scores[i] = pipe.ZScore(ctx, expireZKey(), id)
	}
	if _, err := pipe.Exec(ctx); err != nil && err != redis.Nil {
		return seen, deadline, err
	}
	for i, id := range deviceIds {
		if v, err := gets[i].Int64(); err == nil && v > 0 {
			seen[id] = v
		}
		if sc, err := scores[i].Result(); err == nil {
			deadline[id] = sc
		}
	}
	return seen, deadline, nil
}

// SaveLatestValues 写入最新属性快照（HSET 覆盖同名 identifier）并续期 TTL；时序历史不在这里。
func SaveLatestValues(deviceId string, values map[string]any) {
	if len(values) == 0 {
		return
	}
	ctx, cancel := RedisCtx()
	defer cancel()
	// 逐字段下 HSET：Redis < 4.0 的 HSET 只接受一对 field/value，一次传多对会报
	// “wrong number of arguments”（本项目环境实测）；拆成多条仍可由 pipeline 一个往返完成
	key := stateKey(deviceId)
	pipe := core.Redis.Pipeline()
	for k, v := range values {
		pipe.HSet(ctx, key, k, stringify(v))
	}
	pipe.Expire(ctx, key, time.Duration(config.IotConfig.StateTTL)*time.Second)
	if _, err := pipe.Exec(ctx); err != nil {
		core.Logger.Errorf("[iot处理] 写属性快照失败 deviceId=%s err=%v", deviceId, err)
	}
}

// ClearDeviceRuntime 清除设备运行态（删设备时调），并从到期表移除，避免残留成员被快路径弹出不存在的设备
func ClearDeviceRuntime(deviceId string) {
	ctx, cancel := RedisCtx()
	defer cancel()
	_ = core.Redis.Del(ctx, stateKey(deviceId), seenKey(deviceId)).Err()
	_ = core.Redis.ZRem(ctx, expireZKey(), deviceId).Err()
	for _, fn := range clearDeviceHooks {
		fn(deviceId)
	}
}

// stringify 属性值统一转字符串存 HASH；对象/数组转 JSON 文本，避免丢失结构
func stringify(v any) string {
	switch t := v.(type) {
	case nil:
		return ""
	case string:
		return t
	case bool:
		return strconv.FormatBool(t)
	case float64:
		return strconv.FormatFloat(t, 'f', -1, 64)
	default:
		bs, err := json.Marshal(t)
		if err != nil {
			return ""
		}
		return string(bs)
	}
}

// ReadDeviceState 读出 Redis 侧运行态，供管理端回读验证链路（库侧状态由调用方合并）。
// 在线与剩余时长直读到期表 score（与快路径判定同源），score 缺失时退回“seen+生效窗口”推算；
// 三个键并进一个 pipeline，避免 3 个 RTT 与中间态变迁。
func ReadDeviceState(deviceId string, timeout int) iot_schema.DeviceStateResp {
	ctx, cancel := RedisCtx()
	defer cancel()
	pipe := core.Redis.Pipeline()
	score := pipe.ZScore(ctx, expireZKey(), deviceId)
	seen := pipe.Get(ctx, seenKey(deviceId))
	state := pipe.HGetAll(ctx, stateKey(deviceId))
	if _, err := pipe.Exec(ctx); err != nil && err != redis.Nil {
		core.Logger.Errorf("[iot处理] 读运行态失败 deviceId=%s err=%v", deviceId, err)
	}

	window := offlineWindow(timeout)
	nowMs := time.Now().UnixMilli()
	resp := iot_schema.DeviceStateResp{
		DeviceId:       deviceId,
		OfflineTimeout: timeout,
		Values:         map[string]string{},
	}
	if v, err := state.Result(); err == nil && v != nil {
		resp.Values = v
	}
	if ms, err := strconv.ParseInt(seen.Val(), 10, 64); err == nil && ms > 0 {
		resp.LastSeenMs = ms
		resp.UpdatedAt = time.UnixMilli(ms).Format("2006-01-02 15:04:05")
	}
	if sc, err := score.Result(); err == nil {
		if sc > float64(nowMs) { // 到期表未到期即在线，剩余时长直接由 score 算
			resp.Online = true
			resp.OfflineIn = int64(time.Duration(sc - float64(nowMs)).Seconds())
		}
	} else if err != redis.Nil {
		core.Logger.Errorf("[iot处理] 读到期表失败 deviceId=%s err=%v", deviceId, err)
	}
	if !resp.Online && resp.LastSeenMs > 0 {
		// 到期表无成员但窗口内确有报文（Redis 被清理/慢路径待重建）：按 seen+窗口推算
		if remain := window - time.Since(time.UnixMilli(resp.LastSeenMs)); remain > 0 {
			resp.Online = true
			resp.OfflineIn = int64(remain.Seconds())
		}
	}
	return resp
}

// DeviceRoom 单设备实时房间名（WS 房间标识，不带 Redis 前缀）
func DeviceRoom(deviceId string) string { return "iot:device:" + deviceId }

// MapRoom 区域大屏房间名（对应地图房间广播）
func MapRoom(spaceId string) string { return "iot:map:" + spaceId }

// broadcastStatus 在线状态变更广播（上报路径，带完整信封）
func broadcastStatus(msg iot_schema.IngestMsg, status uint8) {
	BroadcastStatus(msg.DeviceId, msg.Sn, msg.SpaceId, status, msg.ReceivedAt)
}

// BroadcastStatus 广播状态转换：仅在发生上线/离线转换时调用，避免高频报文刷爆 WS；
// 定时扫描路径没有信封，所以字段拆开传。
func BroadcastStatus(deviceId, sn, spaceId string, status uint8, ts int64) {
	if !config.IotConfig.PushStateChange {
		return
	}
	payload := map[string]any{
		"device_id": deviceId,
		"sn":        sn,
		"space_id":  spaceId,
		"status":    status,
		"ts":        ts,
	}
	core.Ws.SendToRoom(DeviceRoom(deviceId), "iot:status", payload)
	if spaceId != "" {
		core.Ws.SendToRoom(MapRoom(spaceId), "iot:status", payload)
	}
}

// BroadcastTelemetry 大屏推送（规则链「大屏推送」节点专用）：只广播到设备所属区域的地图房间；
// 设备详情页实时值由「状态快照」节点 + 轮询/回读承接。
func BroadcastTelemetry(msg iot_schema.IngestMsg, values map[string]any) {
	if !config.IotConfig.PushStateChange || len(values) == 0 || msg.SpaceId == "" {
		return
	}
	payload := map[string]any{
		"device_id": msg.DeviceId,
		"sn":        msg.Sn,
		"space_id":  msg.SpaceId,
		"ts":        msg.Ts,
		"values":    values,
	}
	core.Ws.SendToRoom(MapRoom(msg.SpaceId), "iot:telemetry", payload)
}

// BroadcastAlarm 告警事件广播：recovered=false 为告警中、true 为恢复；severity 为恢复前的级别。
// 前端可据 alertId 把同一条告警的“触发→级别变更→恢复”合并展示。
func BroadcastAlarm(msg iot_schema.IngestMsg, values map[string]any, alertId string, recovered bool, severity int8) {
	if !config.IotConfig.PushStateChange {
		return
	}
	status := uint8(1) // 1=告警中
	if recovered {
		status = 0 // 0=已恢复
	}
	payload := map[string]any{
		"device_id": msg.DeviceId,
		"sn":        msg.Sn,
		"space_id":  msg.SpaceId,
		"alert_id":  alertId,
		"status":    status,
		"severity":  severity,
		"ts":        msg.Ts,
		"values":    values,
	}
	core.Ws.SendToRoom(DeviceRoom(msg.DeviceId), "iot:alarm", payload)
	if msg.SpaceId != "" {
		core.Ws.SendToRoom(MapRoom(msg.SpaceId), "iot:alarm", payload)
	}
}
