package iot_model

import (
	"time"
	"uuid"

	"github.com/adtkcn/x_null"
	"gorm.io/gorm"
)

// IotDeviceData 设备时序数据（本轮以 MySQL 暂代 TDengine，见 iot.md 第五节）
//
// 建表面向“按设备 + 时间区间”查询：联合索引 (device_id, ts) 覆盖最常用的
// 单设备历史曲线；属性以 JSON 文本整包存（键受产品物模型白名单约束），避免为物模型频繁加列。
// 后续切换到 TDengine 时，本模型仅作为过渡存储，写入路径由 pipeline.writeSink 隔离，
// 换库直接改 writeSink 实现即可，主链路无需改动。
type IotDeviceData struct {
	ID         string      `gorm:"primarykey;type:char(36);comment:'主键'" json:"id"`
	DeviceId   string      `gorm:"type:char(36);not null;default:'';index:idx_iot_data_dev_ts,priority:1;comment:'设备ID'" json:"device_id"`
	Sn         string      `gorm:"type:varchar(64);not null;default:'';comment:'设备编号'" json:"sn"`
	ProductId  string      `gorm:"type:char(36);not null;default:'';index;comment:'产品ID'" json:"product_id"`
	SpaceId    string      `gorm:"type:char(36);not null;default:'';index;comment:'区域ID'" json:"space_id"`
	Ts         time.Time   `gorm:"not null;index:idx_iot_data_dev_ts,priority:2;comment:'采集时间'" json:"ts"`
	Values     string      `gorm:"type:text;comment:'属性JSON(identifier->值，规则链加工后按物模型约束)'" json:"values"`
	Protocol   string      `gorm:"type:varchar(16);not null;default:'';comment:'接入协议'" json:"protocol"`
	CreateTime x_null.Time `gorm:"autoCreateTime;not null;comment:'创建时间'" json:"create_time"`
}

func (IotDeviceData) TableName() string { return "x_iot_device_data" }

func (m *IotDeviceData) BeforeCreate(tx *gorm.DB) error {
	m.ID = uuid.NewV7().String()
	return nil
}
