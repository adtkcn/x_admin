package alert

import (
	"time"

	"x_admin/app/model/iot_model"
	"x_admin/app/service/iot_service"
	"x_admin/app/service/iot_service/pipeline/online"
	"x_admin/config"
	"x_admin/core"
)

// init 经 online 钩子接出告警收尾（回调用避免 online ↔ alert 循环依赖）
func init() {
	online.OnDeviceOffline(closeOnOffline) // 设备转离线：收尾其活动告警
	online.OnDeviceClear(ClearRuntime)     // 删设备：一并清告警 Redis 键
}

// RecoverRuntime 服务重启后按库回灌告警活动态：活动态仅存 Redis 却是状态机唯一事实源，
// Redis 丢失而库里仍有 status=1 记录时，恢复报文读不到活动态会被当 NOOP 跳过、永久卡告警中。
// 故消费开始前扫库：按(设备,流)回灌最新活动态，更早的悬空记录一并收尾。
func RecoverRuntime() {
	if !config.IotConfig.Enable {
		return
	}
	var rows []iot_model.IotAlertHistory
	// 按设备+触发时间升序：同一(设备,流)后读到的覆盖先读到的，天然留下最新一条
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
			stale = append(stale, prev.ID)
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

// ClearRuntime 清设备告警运行态（普通流+故障流共六键），不做库侧收尾；删库后必调
func ClearRuntime(deviceId string) {
	ctx, cancel := online.RedisCtx()
	defer cancel()
	_ = core.Redis.Del(ctx, allKeys(deviceId)...).Err()
}

// closeOnOffline 设备离线收尾两条流的活动告警并清键（停报后状态机等不到恢复转换）
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
