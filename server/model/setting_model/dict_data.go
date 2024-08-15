package setting_model

import (
	"x_admin/core"

	"gorm.io/plugin/soft_delete"
)

// DictData 字典数据实体
type DictData struct {
	ID         uint                  `gorm:"primarykey;comment:'主键'"`
	TypeId     uint                  `gorm:"not null;default:0;comment:'类型'"`
	Name       string                `gorm:"not null;default:'';comment:'键名''"`
	Value      string                `gorm:"not null;default:'';comment:'数值'"`
	Color      string                `gorm:"default:'';comment:'颜色'"`
	Remark     string                `gorm:"not null;default:'';comment:'备注'"`
	Sort       uint16                `gorm:"not null;default:0;comment:'排序'"`
	Status     uint8                 `gorm:"not null;default:1;comment:'字典状态: 0=停用, 1=正常'"`
	IsDelete   soft_delete.DeletedAt `gorm:"not null;default:0;softDelete:flag,DeletedAtField:DeleteTime;comment:'是否删除: 0=否, 1=是'"`
	CreateTime core.NullTime         `gorm:"autoCreateTime;not null;comment:'创建时间'"`
	UpdateTime core.NullTime         `gorm:"autoUpdateTime;not null;comment:'更新时间'"`
	DeleteTime core.NullTime         `gorm:"default:null;comment:'删除时间'"`
}
