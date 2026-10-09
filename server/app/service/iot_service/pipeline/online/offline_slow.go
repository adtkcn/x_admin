package online

import (
	"time"
	"x_admin/config"
	"x_admin/core"

	"github.com/redis/go-redis/v9"
)

// SweepOffline 慢路径（5 分钟）：到期表的安全网，全扫库侧在线设备按 seen+窗口现算——
// 超窗的置离线（ZSET 丢失时不至于永远在线），未超窗但 score 缺失/过期的回填到期时间（Redis 重启后重建）。
func SweepOffline() {
	if !config.IotConfig.Enable {
		return
	}
	repairOnlineScores()
}

// isExpired 设备是否已超出自己的离线判定窗口。
// 取不到最后报文时间即视为超窗：要么从未上报，要么 Redis 数据丢了，两者都该是离线。
func isExpired(row offlineCandidate, seen map[string]int64, now int64) (bool, time.Duration) {
	window := offlineWindow(row.DevTimeout)
	last, ok := seen[row.ID]
	if !ok {
		return true, window
	}
	return now-last > window.Milliseconds(), window
}

// repairOnlineScores 逐批复核在线设备：超窗的置离线，未超窗但 score 缺失/已到期的回填 score=seen+窗口。
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
			expired, window := isExpired(row, seen, now)
			if expired {
				changed, err := markOfflineDB(row.ID)
				if err != nil {
					core.Logger.Errorf("[iot处理] 慢路径置离线失败 deviceId=%s err=%v", row.ID, err)
					continue
				}
				if changed {
					core.Logger.Infof("[iot处理] 慢路径判离线 deviceId=%s sn=%s window=%s", row.ID, row.Sn, window)
					BroadcastStatus(row.ID, row.Sn, row.SpaceId, 0, now)
					fireOfflineHooks(row.ID)
				}
				continue
			}
			// score 缺失或已到期但窗口内确有报文（如 Redis 丢数据）：重建到期时间。
			// score 尚在未来则不动，避免用本批旧 seen 把并发报文刚刷新的 score 拉回去
			if sc, ok := deadline[row.ID]; !ok || sc <= float64(now) {
				repair = append(repair, redis.Z{Score: float64(seen[row.ID] + window.Milliseconds()), Member: row.ID})
			}
		}
		if len(repair) > 0 {
			ctx, cancel := RedisCtx()
			if err := core.Redis.ZAdd(ctx, expireZKey(), repair...).Err(); err != nil {
				core.Logger.Errorf("[iot处理] 回填到期时间失败 size=%d err=%v", len(repair), err)
			}
			cancel()
		}
	}
}
