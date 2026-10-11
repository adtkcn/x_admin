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

// 增（写入）：规则链「历史存储」节点加工后的属性只经本文件落库，是全流水线唯一的数据落地出口。
// 单条 INSERT 在高频报文下会把 worker 拖在数据库往返上，故先投内存缓冲，后台批写协程按
// “攒满 batchSize 行或到 flushEvery”批量落库；缓冲满（数据库堵）时退回同步单条写，
// 宁可压慢 worker 也不丢数据。

// Write 存储出口：属性先按物模型白名单裁剪，再投递批写缓冲（满则同步单写）
func Write(msg iot_schema.IngestMsg, values map[string]any, allow map[string]bool) error {
	scoped := iot_schema.ScopeValues(values, allow)
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

// 批写缓冲：chan 容量给足瞬时余量，批写失败只记日志（时序数据允许极小概率丢批）
const (
	bufSize    = 2048
	flushEvery = 2 * time.Second
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
