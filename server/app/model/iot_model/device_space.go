package iot_model

import (
	"uuid"

	"github.com/adtkcn/x_null"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

// IotDeviceSpace 设备区域树
type IotDeviceSpace struct {
	ID         string                `gorm:"primarykey;type:char(36);comment:'主键'" json:"id"`
	Pid        string                `gorm:"type:char(36);not null;default:'';index;comment:'上级区域'" json:"pid"`
	Name       string                `gorm:"type:varchar(100);not null;default:'';comment:'区域名称'" json:"name"`
	Sort       uint16                `gorm:"not null;default:0;comment:'排序'" json:"sort"`
	GeoJson    string                `gorm:"type:text;comment:'区域范围GeoJSON(Polygon)'" json:"geojson"`
	Remark     string                `gorm:"type:varchar(255);not null;default:'';comment:'备注'" json:"remark"`
	IsDelete   soft_delete.DeletedAt `gorm:"not null;default:0;softDelete:flag,DeletedAtField:DeleteTime;comment:'是否删除'" json:"-"`
	CreateTime x_null.Time           `gorm:"autoCreateTime;not null;comment:'创建时间'" json:"create_time"`
	UpdateTime x_null.Time           `gorm:"autoUpdateTime;not null;comment:'更新时间'" json:"update_time"`
	DeleteTime x_null.Time           `gorm:"default:null;comment:'删除时间'" json:"-"`
}

func (IotDeviceSpace) TableName() string { return "x_iot_device_space" }

func (m *IotDeviceSpace) BeforeCreate(tx *gorm.DB) error {
	m.ID = uuid.NewV7().String()
	return nil
}
