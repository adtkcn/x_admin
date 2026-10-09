package store

import (
	"encoding/json"
	"sync"
	"time"

	"x_admin/app/model/iot_model"
	"x_admin/app/schema/iot_schema"
	"x_admin/config"
	"x_admin/core"
)

// 时序存储出口：规则链「历史存储」节点加工后的属性只经本包落库，是全流水线唯一的数据落地
// 出口——换存储介质（如 MySQL→TDengine）只改本包。落库属性受物模型白名单约束（由调用方传 allow），
// 本包不查库、不引入 service 依赖。
//
// 本期用 MySQL 顶替 TDengine：单条 INSERT 在高频报文下会把 worker 拖在数据库往返上，
// 故先投内存缓冲，后台批写协程按“攒满 batchSize 行或到 flushEvery”批量落库；
// 缓冲满（数据库堵）时退回同步单条写，宁可压慢 worker 也不丢数据。

// Write 存储出口：属性先按物模型白名单裁剪，再投递批写缓冲（满则同步单写）
func Write(msg iot_schema.IngestMsg, values map[string]any, allow map[string]bool) error {
	scoped := scopeValues(values, allow)
	if len(scoped) == 0 {
		core.Logger.Warnf("[iot存储] 属性无一命中物模型白名单，跳过落库 sn=%s values=%d", msg.Sn, len(values))
		return nil
	}
	bs, err := json.Marshal(scoped)
	if err != nil {
		return err
	}
	row := iot_model.IotDeviceData{
		DeviceId:  msg.DeviceId,
		Sn:        msg.Sn,
		ProductId: msg.ProductId,
		SpaceId:   msg.SpaceId,
		Ts:        time.UnixMilli(msg.Ts),
		Values:    string(bs),
		Protocol:  msg.Protocol,
	}
	startBatcher()
	select {
	case bufCh <- row:
		return nil
	default:
		return batchCreateRows([]iot_model.IotDeviceData{row})
	}
}

// scopeValues 按物模型白名单裁剪属性。allow 为空表示未配物模型（或旧形态无法解析），
// 此时不约束、直接返回原 map，避免配置缺失反而误丢数据。
func scopeValues(values map[string]any, allow map[string]bool) map[string]any {
	if len(allow) == 0 {
		return values
	}
	out := make(map[string]any, len(values))
	for k, v := range values {
		if allow[k] {
			out[k] = v
		}
	}
	return out
}

// 批写缓冲：chan 容量给足瞬时余量，批写失败只记日志（时序数据允许极小概率丢批）
const (
	bufSize      = 2048
	flushEvery   = 2 * time.Second
	cleanBatch   = 2000 // 保留期清理的单批删除行数，避免大事务锁表
	queryMaxRows = 5000 // 区间回读的单次行数上限（趋势图取数容量）
)

var (
	bufCh     chan iot_model.IotDeviceData
	startOnce sync.Once
)

// startBatcher 首条数据到达时才惰性拉起批写协程，流水线未启用时不留空转任务
func startBatcher() {
	startOnce.Do(func() {
		bufCh = make(chan iot_model.IotDeviceData, bufSize)
		go flusher()
	})
}

func flusher() {
	ticker := time.NewTicker(flushEvery)
	defer ticker.Stop()
	buf := make([]iot_model.IotDeviceData, 0, config.IotConfig.DataBatchSize)
	for {
		select {
		case row := <-bufCh:
			buf = append(buf, row)
			if len(buf) >= config.IotConfig.DataBatchSize {
				batchCreateRows(buf)
				buf = buf[:0]
			}
		case <-ticker.C:
			if len(buf) > 0 {
				batchCreateRows(buf)
				buf = buf[:0]
			}
		}
	}
}

// Flush 进程退出（队列已停、数据库未关）时把缓冲剩余数据同步落库，由 main 的优雅关闭序列调用
func Flush() {
	if bufCh == nil {
		return
	}
	buf := make([]iot_model.IotDeviceData, 0, config.IotConfig.DataBatchSize)
	for len(bufCh) > 0 {
		buf = append(buf, <-bufCh)
		if len(buf) >= config.IotConfig.DataBatchSize {
			batchCreateRows(buf)
			buf = buf[:0]
		}
	}
	if len(buf) > 0 {
		batchCreateRows(buf)
	}
}

func batchCreateRows(rows []iot_model.IotDeviceData) error {
	if len(rows) == 0 {
		return nil
	}
	if err := core.GetDB().Create(&rows).Error; err != nil {
		core.Logger.Errorf("[iot处理] 时序批量写入失败 size=%d err=%v", len(rows), err)
		return err
	}
	return nil
}

// QueryRange 按时间区间读设备时序（毫秒时间戳，0=不限），ts 倒序，供管理端趋势图/回读
func QueryRange(deviceId string, startMs, endMs int64, limit int) []iot_schema.DeviceDataResp {
	if limit <= 0 || limit > queryMaxRows {
		limit = 1000
	}
	q := core.GetDB().Where("device_id = ?", deviceId)
	if startMs > 0 {
		q = q.Where("ts >= ?", time.UnixMilli(startMs))
	}
	if endMs > 0 {
		q = q.Where("ts <= ?", time.UnixMilli(endMs))
	}
	var rows []iot_model.IotDeviceData
	if err := q.Order("ts desc").Limit(limit).Find(&rows).Error; err != nil {
		core.Logger.Errorf("[iot处理] 回读时序失败 deviceId=%s err=%v", deviceId, err)
		return []iot_schema.DeviceDataResp{}
	}
	list := make([]iot_schema.DeviceDataResp, 0, len(rows))
	for _, row := range rows {
		vals := map[string]any{}
		if row.Values != "" {
			if err := json.Unmarshal([]byte(row.Values), &vals); err != nil {
				core.Logger.Warnf("[iot处理] 时序属性反编译失败 id=%s err=%v", row.ID, err)
			}
		}
		list = append(list, iot_schema.DeviceDataResp{
			Ts:       row.Ts.UnixMilli(),
			Values:   vals,
			Protocol: row.Protocol,
		})
	}
	return list
}

// Clean 删除保留期之外的时序数据（定时任务调）：DataKeepDays=0 为永久保留，按批小事务避免锁表。
func Clean() {
	days := config.IotConfig.DataKeepDays
	if days <= 0 {
		return
	}
	cutoff := time.Now().AddDate(0, 0, -days)
	for {
		res := core.GetDB().Where("ts < ?", cutoff).Limit(cleanBatch).Delete(&iot_model.IotDeviceData{})
		if res.Error != nil {
			core.Logger.Errorf("[iot处理] 时序保留期清理失败 err=%v", res.Error)
			return
		}
		if res.RowsAffected < int64(cleanBatch) {
			return
		}
	}
}
