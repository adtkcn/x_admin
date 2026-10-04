package pipeline

import (
	"encoding/json"
	"time"

	"x_admin/app/model/iot_model"
	"x_admin/app/schema/iot_schema"
	"x_admin/core"
)

// writeSink 存储出口（流水线第四级）：把加工后的属性落到存储介质。
//
// 全链路只有一个存储出口，故不再做多 Sink 注册表。本期按需求用 MySQL 顶替
// iot.md 第四阶段的 TDengine；后续换库时直接改本函数的实现即可，
// handleTelemetry 与接入层都不受影响。
func writeSink(msg iot_schema.IngestMsg, values map[string]any) error {
	bs, err := json.Marshal(values)
	if err != nil {
		return err
	}
	return core.GetDB().Create(&iot_model.IotDeviceData{
		DeviceId:  msg.DeviceId,
		Sn:        msg.Sn,
		ProductId: msg.ProductId,
		SpaceId:   msg.SpaceId,
		Ts:        time.UnixMilli(msg.Ts),
		Values:    string(bs),
		Raw:       msg.Raw,
		Protocol:  msg.Protocol,
	}).Error
}

// QueryRecent 读取某设备最近 N 条时序数据（供管理端回读验证链路）。
// 放在 pipeline 而非 iot_service，是因为它与存储形态强绑定：换 TDengine 时
// 只需改这里，流水线其余部分不动。返回值把属性 JSON 解回 map，前端直接可用。
func QueryRecent(deviceId string, limit int) []iot_schema.DeviceDataResp {
	if limit <= 0 || limit > 500 {
		limit = 50
	}
	var rows []iot_model.IotDeviceData
	if err := core.GetDB().Where("device_id = ?", deviceId).
		Order("ts desc").Limit(limit).Find(&rows).Error; err != nil {
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
			Raw:      row.Raw,
			Protocol: row.Protocol,
		})
	}
	return list
}
