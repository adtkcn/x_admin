package online

import (
	"errors"
	"time"
	"x_admin/config"
	"x_admin/core"

	"github.com/redis/go-redis/v9"
)

// ScanOffline 快路径（每 5s）：只弹到期表 ZSET 里已到期的设备置离线。
// 到期即离线，不用 seen+当前窗口复核：改窗口的瞬时误差由下一条报文的上线转换自愈，
// 慢路径 5 分钟内也会重建修正。
func ScanOffline() {
	if !config.IotConfig.Enable {
		return // 接入层关了就没有报文刷新 score，再扫会把所有在线设备误判为离线
	}
	markExpiredOffline()
}

// popDueScript 单命令内原子完成“查到期+删除”（KEYS[1]=到期表，ARGV[1]=当前毫秒，ARGV[2]=单批上限），
// 避免查与删之间夹进并发报文刷新 score 后仍被误带走。
var popDueScript = redis.NewScript(`
local ids = redis.call('ZRANGEBYSCORE', KEYS[1], '-inf', ARGV[1], 'LIMIT', 0, ARGV[2])
if #ids == 0 then return ids end
redis.call('ZREM', KEYS[1], unpack(ids))
return ids
`)

// markExpiredOffline 从到期表原子弹出已到期成员并置离线，每批只回一次表（取 sn/space 供广播）。
// 弹出后才到达的新报文会因“score 不存在”重新触发上线转换，下一轮到期自然修正。
func markExpiredOffline() {
	ctx, cancel := RedisCtx()
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
			// 整轮停下而不是接着弹：库不可用时继续弹出只会把越来越多成员的到期记录
			// 从表里带走却无法置离线。仍在表里的成员下个 tick 重试；已弹出的靠该设备
			// 下一条报文刷新 score 回表，真正不再上报的由慢路径按 seen+窗口收尾
			return
		}
		for _, row := range rows {
			changed, err := markOfflineDB(row.ID)
			if err != nil {
				core.Logger.Errorf("[iot处理] 到期置离线失败 deviceId=%s err=%v", row.ID, err)
				continue
			}
			if changed {
				core.Logger.Infof("[iot处理] 到期判离线 deviceId=%s sn=%s timeout=%ds", row.ID, row.Sn, offlineWindow(row.DevTimeout)/time.Second)
				BroadcastStatus(row.ID, row.Sn, row.SpaceId, 0, now)
				fireOfflineHooks(row.ID)
			}
		}
		if len(ids) < scanChunk {
			return
		}
	}
}
