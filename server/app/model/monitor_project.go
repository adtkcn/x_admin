package model

import (
	"github.com/adtkcn/x_null"
	"uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

// MonitorProject 监控项目实体
type MonitorProject struct {
	Id          string                `gorm:"primarykey;type:char(36);comment:'项目id'"` // 项目id
	ProjectKey  string                `gorm:"comment:'项目uuid'"`                        // 项目uuid
	ProjectName string                `gorm:"comment:'项目名称'"`                          // 项目名称
	ProjectType string                `gorm:"comment:'项目类型go java web node php 等'"`    // 项目类型go java web node php 等
	Status      x_null.Int64          `gorm:"comment:'是否启用: 0=否, 1=是'"`                // 是否启用: 0=否, 1=是
	IsDelete    soft_delete.DeletedAt `gorm:"not null;default:0;softDelete:flag,DeletedAtField:DeleteTime;comment:'是否删除: 0=否, 1=是'"`
	CreateTime  x_null.Time           `gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
	UpdateTime  x_null.Time           `gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间
	DeleteTime  x_null.Time           `gorm:"comment:'删除时间'"`                // 删除时间
}

// BeforeCreate 在创建前生成UUIDv7
func (m *MonitorProject) BeforeCreate(tx *gorm.DB) (err error) {
	id := uuid.NewV7()
	m.Id = id.String()
	return nil
}
