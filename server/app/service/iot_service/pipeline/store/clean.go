package store

import (
	"time"

	"x_admin/app/model/iot_model"
	"x_admin/config"
	"x_admin/core"
)

// 删（清理）：保留期外的时序数据由定时任务按批删除。
const cleanBatch = 2000 // 保留期清理的单批删除行数，避免大事务锁表

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
