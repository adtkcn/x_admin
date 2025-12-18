package model

import (
	"x_admin/core"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

// UserProtocol 用户协议实体
type UserProtocol struct {
	Id      string          `gorm:"primarykey;type:char(36);comment:''"` //
	Tag     core.NullString `gorm:"comment:'标识'"`                        // 标识
	Version core.NullInt    `gorm:"column:version;comment:'排序'"`         // 排序
	Title   core.NullString `gorm:"comment:'标题'"`                        // 标题
	Content core.NullString `gorm:"comment:'协议内容'"`                      // 协议内容
	// Sort       core.NullFloat        `gorm:"comment:'排序'"`          // 排序
	CreateBy   string                `gorm:"comment:'创建人'"` // 创建人
	IsDelete   soft_delete.DeletedAt `gorm:"not null;default:0;softDelete:flag,DeletedAtField:DeleteTime;comment:'是否删除: 0=否, 1=是'"`
	CreateTime core.NullTime         `gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
	UpdateTime core.NullTime         `gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间
	DeleteTime core.NullTime         `gorm:"comment:'删除时间'"`                // 删除时间
}

// 自动在创建时设置 UUIDv7
func (u *UserProtocol) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	u.Id = id.String()
	return nil
}
