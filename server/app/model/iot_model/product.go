package iot_model

import (
	"uuid"

	"github.com/adtkcn/x_null"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

// IotProduct 产品定义
type IotProduct struct {
	ID               string                `gorm:"primarykey;type:char(36);comment:'主键'" json:"id"`
	Name             string                `gorm:"type:varchar(100);not null;default:'';comment:'产品名称'" json:"name"`
	CategoryId       string                `gorm:"type:char(36);not null;default:'';index;comment:'分类ID'" json:"category_id"`
	DeviceType       string                `gorm:"type:varchar(50);not null;default:'';comment:'节点类型: DEVICE=直连设备, GATEWAY=网关设备, EDGE=边缘主机'" json:"device_type"`
	Protocol         string                `gorm:"type:varchar(20);not null;default:'';comment:'接入协议:mqtt/tcp/http/ws'" json:"protocol"`
	ThingModel       string                `gorm:"type:text;comment:'物模型JSON'" json:"thing_model"`
	DecodeScript     string                `gorm:"type:text;comment:'解码脚本'" json:"decode_script"`
	PreprocessScript string                `gorm:"type:text;comment:'预处理脚本(goja)'" json:"preprocess_script"`
	Remark           string                `gorm:"type:varchar(255);not null;default:'';comment:'备注'" json:"remark"`
	IsStop           uint8                 `gorm:"not null;default:0;comment:'是否停用: 0=否, 1=是'" json:"is_stop"`
	IsDelete         soft_delete.DeletedAt `gorm:"not null;default:0;softDelete:flag,DeletedAtField:DeleteTime;comment:'是否删除'" json:"-"`
	CreateTime       x_null.Time           `gorm:"autoCreateTime;not null;comment:'创建时间'" json:"create_time"`
	UpdateTime       x_null.Time           `gorm:"autoUpdateTime;not null;comment:'更新时间'" json:"update_time"`
	DeleteTime       x_null.Time           `gorm:"default:null;comment:'删除时间'" json:"-"`
}

func (IotProduct) TableName() string { return "x_iot_product" }

func (m *IotProduct) BeforeCreate(tx *gorm.DB) error {
	m.ID = uuid.NewV7().String()
	return nil
}
