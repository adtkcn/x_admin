package pipeline

import (
	"context"
	"encoding/json"
	"errors"
	"time"
	"x_admin/app/schema/iot_schema"
	"x_admin/app/service/iot_service"
	"x_admin/config"
	"x_admin/core"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// Pipeline 设备报文处理层（iot.md 第一层：数据处理流水线）
//
// 链路：接入层 → core.Queue("iot:ingest") → 本包 worker →
// [刷新运行态] → [解码脚本(goja)] → [物模型过滤] → [过滤脚本(goja)] → [Redis 快照] → [Sink 存储] → [WS 广播] → [告警脚本(goja)+告警状态机]
//
// 在线方向由报文实时驱动（首条报文即上线），离线方向由到期表（ZSET）驱动：
// 每条报文/心跳 touchAlive 时 ZADD“到期时间=到达时间+设备窗口”，
// app/corn 的 iotOfflineScan 快路径每轮只弹到期成员（O(到期数)）；
// iotOfflineSweep 慢路径兜底：全扫现算并重建到期表，
// 保证判定数据可重建，Redis 丢数据不会让设备“永远在线”。
//
// 第五阶段(告警状态机) 已在 alert.go 落地：产品 goja 告警脚本逐条判定 + 持续/静默/恢复
// 收敛成事件级记录，写 x_iot_alert_history、按转换点推 iot:alarm 广播；
// 第四阶段存储出口只有一个，直接改 writeSink 实现即可（如换 TDengine）。

// StopRuntime 设备被删除后清理其运行态，避免 Redis 里残留幽灵设备
func StopRuntime(deviceId string) { clearRuntime(deviceId) }

// scanChunk 分批大小：既限住 IN 拼接出的 SQL 长度，也限住一次 Redis pipeline 里的命令数
const scanChunk = 500

// handleIngest 消费一条设备报文
func HandleIngest(_ context.Context, body []byte) error {
	var msg iot_schema.IngestMsg
	if err := json.Unmarshal(body, &msg); err != nil {
		// 坏报文无法自愈，返回错误由队列记日志即可（当前实现不重试，避免死循环）
		core.Logger.Errorf("[iot处理] 报文解析失败 err=%v body=%s", err, string(body))
		return err
	}
	if msg.DeviceId == "" {
		return errors.New("报文缺少 device_id")
	}
	if config.IotConfig.LogRaw {
		core.Logger.Infof("[iot处理] 消费报文 sn=%s type=%s values=%d", msg.Sn, msg.MsgType, len(msg.Values))
	}

	// 运行态与在线转换对所有报文类型都生效（心跳/上报/事件都会“续命”）
	touchRuntime(msg)

	switch msg.MsgType {
	case iot_schema.MsgTypeTelemetry:
		return handleTelemetry(msg)
	case iot_schema.MsgTypeHeartbeat:
		return nil
	case iot_schema.MsgTypeEvent:
		// 事件上报：本阶段只保证运行态刷新，告警联动在第五阶段接入
		core.Logger.Infof("[iot处理] 事件报文暂存待处理(告警阶段实现) sn=%s raw=%s", msg.Sn, msg.Raw)
		return nil
	default:
		return errors.New("未知报文类型: " + msg.MsgType)
	}
}

// touchRuntime 刷新在线标记，并在“离线→在线”转换时写库+广播
func touchRuntime(msg iot_schema.IngestMsg) {
	if touchAlive(msg.DeviceId, msg.OfflineTimeout) {
		changed, err := iot_service.DeviceService.MarkStatus(msg.DeviceId, 1)
		if err != nil {
			core.Logger.Errorf("[iot处理] 置在线失败 deviceId=%s err=%v", msg.DeviceId, err)
			// 写库失败撤销刚写入的到期 score：下一条报文会因“score 缺失”重新触发上线转换，
			// 重试由报文流自己完成，不需要额外的补判扫描；撤销失败则退回慢路径到期后自愈
			revokeDeadline(msg.DeviceId)
		} else if changed {
			pushStatus(msg, 1)
		}
	}
}

// handleTelemetry 属性上报：解码(goja) → 物模型过滤 → 过滤脚本(goja) → 快照 → 存储 → 广播 → 告警脚本(goja)
// 物模型与三段脚本同属产品配置，从同一条缓存一次读出，避免每报文查多次库。
func handleTelemetry(msg iot_schema.IngestMsg) error {
	model := iot_service.DeviceService.CachedProductModel(msg.ProductId)
	values := decodeMessage(msg, model.DecodeScript)
	values = filterByThingModel(msg, values, model.ThingModel)
	if len(values) == 0 {
		return nil
	}
	// 数据过滤脚本：返回 false 则丢弃本条，不进快照/存储/广播/告警
	if !passFilter(msg, values, model.FilterScript) {
		return nil
	}
	saveState(msg.DeviceId, values)
	if err := writeSink(msg, values); err != nil {
		return err
	}
	pushTelemetry(msg, values)
	// 数据告警脚本：alert(values) 返回级别数字，交给状态机（持续时长/静默/升级/恢复）收敛成事件级告警。
	// 仅在配置了告警脚本时参与，未配置则链路零额外 Redis 开销。
	if model.AlertScript != "" {
		evalAlert(msg, values, alertSeverity(msg, values, model.AlertScript), model)
	}
	return nil
}

// ============ 离线判定：到期表 ZSET（快路径 iotOfflineScan + 慢路径 iotOfflineSweep） ============

// ScanOffline 快路径（30s）：只弹 ZSET 里已到期的设备置离线。
// 到期即离线，不用 seen+当前窗口复核：改窗口后 score 是旧值的瞬时误差，
// 由下一条报文“离线→在线”转换自愈，慢路径 5 分钟内也会重建修正。
func ScanOffline() {
	if !config.IotConfig.Enable {
		return // 接入层关了就没有报文刷新 score，再扫会把所有在线设备误判为离线
	}
	markExpiredOffline()
}

// SweepOffline 慢路径（5 分钟）：到期表的安全网，回表固定一次——
// 全扫库侧在线设备，按 seen+设备窗口现算：超窗的置离线（ZSET 丢失时不至于永远在线），
// 未超窗但 score 缺失/已过期的回填到期时间（Redis 重启后重建到期表）。
func SweepOffline() {
	if !config.IotConfig.Enable {
		return
	}
	repairOnlineScores()
}

// popDueScript 单命令内原子完成“查到期+删除”（KEYS[1]=到期表，ARGV[1]=当前毫秒，ARGV[2]=单批上限），
// 避免查与删之间夹进并发报文刷新 score 后仍被误带走。
var popDueScript = redis.NewScript(`
local ids = redis.call('ZRANGEBYSCORE', KEYS[1], '-inf', ARGV[1], 'LIMIT', 0, ARGV[2])
if #ids == 0 then return ids end
redis.call('ZREM', KEYS[1], unpack(ids))
return ids
`)

// markExpiredOffline 从到期表原子弹出已到期（score ≤ 当前时间）的成员并置离线，
// 每批弹出只回一次表（取 sn/space 供广播，已离线的无需再动）。
// 到期即离线，不用 seen+当前窗口复核：弹出后才到达的新报文，
// 会因“score 不存在”重新触发离线→在线转换，下一轮到期周期自然修正。
func markExpiredOffline() {
	ctx, cancel := redisCtx()
	defer cancel()
	now := time.Now().UnixMilli()
	for {
		ids, err := popDueScript.Run(ctx, core.Redis, []string{expireZKey()}, now, scanChunk).StringSlice()
		if err != nil && !errors.Is(err, redis.Nil) {
			core.Logger.Errorf("[iot处理] 弹出到期设备失败 err=%v", err)
			return
		}
		if len(ids) == 0 {
			return
		}
		rows, err := scanCandidates(core.GetDB(), 1, ids)
		if err != nil {
			core.Logger.Errorf("[iot处理] 查到期设备信息失败 err=%v", err)
			continue
		}
		for _, row := range rows {
			changed, err := iot_service.DeviceService.MarkStatus(row.ID, 0)
			if err != nil {
				core.Logger.Errorf("[iot处理] 到期置离线失败 deviceId=%s err=%v", row.ID, err)
				continue
			}
			if changed {
				core.Logger.Infof("[iot处理] 到期判离线 deviceId=%s sn=%s timeout=%ds", row.ID, row.Sn, offlineWindow(row.DevTimeout)/time.Second)
				pushStatusOf(row.ID, row.Sn, row.SpaceId, 0, now)
			}
		}
		if len(ids) < scanChunk {
			return
		}
	}
}

// offlineCandidate 扫描候选：设备自己配置的离线判定时长。
// 时长以设备列为准（新增/编辑时已冗余产品值，产品变更会同步到设备），
// 扫描时现读库、不沿用上报侧的鉴权缓存值，改完配置下一轮就生效。
type offlineCandidate struct {
	ID         string
	Sn         string
	SpaceId    string `gorm:"column:space_id"`
	DevTimeout int    `gorm:"column:dev_timeout"`
}

// scanCandidates 按在线状态取候选设备（ids 为空表示不限定），只查设备表，
// 离线时长以设备列为准
func scanCandidates(db *gorm.DB, status uint8, ids []string) (rows []offlineCandidate, e error) {
	q := db.Table("x_iot_device AS d").
		Select("d.id, d.sn, d.space_id, d.offline_timeout AS dev_timeout").
		Where("d.status = ? AND d.is_delete = 0", status)
	if len(ids) > 0 {
		q = q.Where("d.id IN ?", ids)
	}
	e = q.Find(&rows).Error
	return
}

// expiredWithin 设备是否已超出自己的离线判定窗口（以设备配置为准，未配置时落兜底值）。
// 取不到最后报文时间即视为已超窗：要么从未上报过，要么 Redis 数据丢了，两者都该是离线。
func expiredWithin(row offlineCandidate, seen map[string]int64, now int64) (bool, time.Duration) {
	window := offlineWindow(row.DevTimeout)
	last, ok := seen[row.ID]
	if !ok {
		return true, window
	}
	return now-last > window.Milliseconds(), window
}

// repairOnlineScores 全扫库侧在线设备：超窗的置离线（ZSET 丢失时的最后一道兜底，
// 判据是 seen+设备窗口；seen 也读不到就当超窗），
// 未超窗但 score 缺失/已到期的回填 score=seen+窗口（Redis 重启后重建到期表）。
func repairOnlineScores() {
	rows, err := scanCandidates(core.GetDB(), 1, nil)
	if err != nil {
		core.Logger.Errorf("[iot处理] 慢路径扫描在线设备失败 err=%v", err)
		return
	}
	now := time.Now().UnixMilli()
	for start := 0; start < len(rows); start += scanChunk {
		chunk := rows[start:min(start+scanChunk, len(rows))]
		ids := make([]string, len(chunk))
		for i, row := range chunk {
			ids[i] = row.ID
		}
		seen, deadline, err := lastSeenAndDeadlines(ids)
		if err != nil {
			// Redis 读不动就整批跳过：把“读不到时间戳”当成“没上报过”会一波误判离线
			core.Logger.Errorf("[iot处理] 批量读报文时间/到期时间失败，本批跳过 size=%d err=%v", len(ids), err)
			continue
		}

		var repair []redis.Z
		for _, row := range chunk {
			expired, window := expiredWithin(row, seen, now)
			if expired {
				changed, err := iot_service.DeviceService.MarkStatus(row.ID, 0)
				if err != nil {
					core.Logger.Errorf("[iot处理] 慢路径置离线失败 deviceId=%s err=%v", row.ID, err)
					continue
				}
				if changed {
					core.Logger.Infof("[iot处理] 慢路径判离线 deviceId=%s sn=%s window=%s", row.ID, row.Sn, window)
					pushStatusOf(row.ID, row.Sn, row.SpaceId, 0, now)
				}
				continue
			}
			// score 缺失或已到期但窗口内实际仍有报文（如 Redis 丢数据后）：重建到期时间。
			// score 尚在未来则不动，避免用本轮的旧 seen 把并发报文刚刷新的 score 拉回去
			if sc, ok := deadline[row.ID]; !ok || sc <= float64(now) {
				repair = append(repair, redis.Z{Score: float64(seen[row.ID] + window.Milliseconds()), Member: row.ID})
			}
		}
		if len(repair) > 0 {
			ctx, cancel := redisCtx()
			if err := core.Redis.ZAdd(ctx, expireZKey(), repair...).Err(); err != nil {
				core.Logger.Errorf("[iot处理] 回填到期时间失败 size=%d err=%v", len(repair), err)
			}
			cancel()
		}
	}
}
