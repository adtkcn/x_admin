package model

import (
	"x_admin/core"

	"gorm.io/plugin/soft_delete"
)

// MonitorProject 监控项目实体
type MonitorProject struct {
	Id          int                   `gorm:"primarykey;comment:'项目id'"`            // 项目id
	ProjectKey  string                `gorm:"comment:'项目uuid'"`                     // 项目uuid
	ProjectName string                `gorm:"comment:'项目名称'"`                       // 项目名称
	ProjectType string                `gorm:"comment:'项目类型go java web node php 等'"` // 项目类型go java web node php 等
	Status      core.NullInt          `gorm:"comment:'是否启用: 0=否, 1=是'"`             // 是否启用: 0=否, 1=是
	IsDelete    soft_delete.DeletedAt `gorm:"not null;default:0;softDelete:flag,DeletedAtField:DeleteTime;comment:'是否删除: 0=否, 1=是'"`
	CreateTime  core.NullTime         `gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
	UpdateTime  core.NullTime         `gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间
	DeleteTime  core.NullTime         `gorm:"comment:'删除时间'"`                // 删除时间
}
