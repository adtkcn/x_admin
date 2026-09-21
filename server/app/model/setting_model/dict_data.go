package setting_model

import (
	"uuid"

	"github.com/adtkcn/x_null"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

// DictData 字典数据实体
type DictData struct {
	ID         string                `gorm:"primarykey;comment:'主键'"`
	TypeId     string                `gorm:"not null;default:'';comment:'类型'"`
	Name       string                `gorm:"not null;default:'';comment:'键名''"`
	Value      string                `gorm:"not null;default:'';comment:'数值'"`
	Color      string                `gorm:"default:'';comment:'颜色'"`
	Remark     string                `gorm:"not null;default:'';comment:'备注'"`
	Sort       uint16                `gorm:"not null;default:0;comment:'排序'"`
	Status     uint8                 `gorm:"not null;default:1;comment:'字典状态: 0=停用, 1=正常'"`
	IsDelete   soft_delete.DeletedAt `gorm:"not null;default:0;softDelete:flag,DeletedAtField:DeleteTime;comment:'是否删除: 0=否, 1=是'"`
	CreateTime x_null.Time           `gorm:"autoCreateTime;not null;comment:'创建时间'"`
	UpdateTime x_null.Time           `gorm:"autoUpdateTime;not null;comment:'更新时间'"`
	DeleteTime x_null.Time           `gorm:"default:null;comment:'删除时间'"`
}

// BeforeCreate 在创建前生成UUIDv7
func (m *DictData) BeforeCreate(tx *gorm.DB) (err error) {
	id := uuid.NewV7()
	m.ID = id.String()
	return nil
}
