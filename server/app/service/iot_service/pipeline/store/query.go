package store

import (
	"encoding/json"
	"time"

	"x_admin/app/model/iot_model"
	"x_admin/app/schema/iot_schema"
	"x_admin/core"
)

// 查（回读）：按时间区间读取已落库的时序数据，供管理端趋势图/历史回读。

// QueryRange 按时间区间读设备时序（毫秒时间戳，0=不限），ts 倒序，供管理端趋势图/回读
// limit<=0 表示不限制条数（区间内全量回读）；limit>0 才加 LIMIT 子句
func QueryRange(deviceId string, startMs, endMs int64, limit int) []iot_schema.DeviceDataResp {
	q := core.GetDB().Where("device_id = ?", deviceId)
	if startMs > 0 {
		q = q.Where("ts >= ?", time.UnixMilli(startMs))
	}
	if endMs > 0 {
		q = q.Where("ts <= ?", time.UnixMilli(endMs))
	}
	db := q.Order("ts desc")
	if limit > 0 {
		db = db.Limit(limit)
	}
	var rows []iot_model.IotDeviceData
	if err := db.Find(&rows).Error; err != nil {
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
