package iot_model

import (
	"uuid"

	"github.com/adtkcn/x_null"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

// IotDevice 设备
type IotDevice struct {
	ID              string                `gorm:"primarykey;type:char(36);comment:'主键'" json:"id"`
	Name            string                `gorm:"type:varchar(100);not null;default:'';comment:'设备名称'" json:"name"`
	Sn              string                `gorm:"type:varchar(64);not null;default:'';uniqueIndex;comment:'设备编号'" json:"sn"`
	ProductId       string                `gorm:"type:char(36);not null;default:'';index;comment:'产品ID'" json:"product_id"`
	SpaceId         string                `gorm:"type:char(36);not null;default:'';index;comment:'区域ID'" json:"space_id"`
	GatewayId       string                `gorm:"type:char(36);not null;default:'';index;comment:'网关设备ID'" json:"gateway_id"`
	Status          uint8                 `gorm:"not null;default:0;comment:'在线状态: 0=离线, 1=在线（由设备接入层根据心跳/连接事件自动维护，不开放人工编辑）'" json:"status"`
	OfflineTimeout  int                   `gorm:"not null;default:0;comment:'离线判定时长(秒): 以设备为准（冗余+自定义），新增/编辑未指定时写入产品值；仍为 0 则用兜底值 300 秒'" json:"offline_timeout"`
	LastOnlineTime  x_null.Time           `gorm:"default:null;comment:'最后上线时间'" json:"last_online_time"`
	LastOfflineTime x_null.Time           `gorm:"default:null;comment:'最后离线时间'" json:"last_offline_time"`
	Longitude       float64               `gorm:"type:decimal(10,7);not null;default:0;comment:'经度'" json:"longitude"`
	Latitude        float64               `gorm:"type:decimal(10,7);not null;default:0;comment:'纬度'" json:"latitude"`
	AccessToken     string                `gorm:"type:varchar(64);not null;default:'';comment:'接入凭证'" json:"access_token"`
	Remark          string                `gorm:"type:varchar(255);not null;default:'';comment:'备注'" json:"remark"`
	IsDelete        soft_delete.DeletedAt `gorm:"not null;default:0;softDelete:flag,DeletedAtField:DeleteTime;comment:'是否删除'" json:"-"`
	CreateTime      x_null.Time           `gorm:"autoCreateTime;not null;comment:'创建时间'" json:"create_time"`
	UpdateTime      x_null.Time           `gorm:"autoUpdateTime;not null;comment:'更新时间'" json:"update_time"`
	DeleteTime      x_null.Time           `gorm:"default:null;comment:'删除时间'" json:"-"`
}

func (IotDevice) TableName() string { return "x_iot_device" }

func (m *IotDevice) BeforeCreate(tx *gorm.DB) error {
	m.ID = uuid.NewV7().String()
	return nil
}
