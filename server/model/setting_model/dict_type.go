package setting_model

import "x_admin/core"

//DictType 字典类型实体
type DictType struct {
	ID         uint        `gorm:"primarykey;comment:'主键'"`
	DictName   string      `gorm:"not null;default:'';comment:'字典名称''"`
	DictType   string      `gorm:"not null;default:'';comment:'字典类型'"`
	DictRemark string      `gorm:"not null;default:'';comment:'字典备注'"`
	DictStatus uint8       `gorm:"not null;default:1;comment:'字典状态: 0=停用, 1=正常'"`
	IsDelete   uint8       `gorm:"not null;default:0;comment:'是否删除: 0=否, 1=是'"`
	CreateTime core.TsTime `gorm:"autoCreateTime;not null;comment:'创建时间'"`
	UpdateTime core.TsTime `gorm:"autoUpdateTime;not null;comment:'更新时间'"`
	DeleteTime core.TsTime `gorm:"default:null;comment:'删除时间'"`
}
