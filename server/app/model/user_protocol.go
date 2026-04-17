package model

import (
	"x_admin/app/model/system_model"

	"github.com/adtkcn/x_null"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

// UserProtocol 用户协议实体
type UserProtocol struct {
	Id            string                             `gorm:"primarykey;type:char(36);comment:''"` //
	Tag           x_null.String                      `gorm:"comment:'标识'"`                        // 标识
	Version       x_null.Int64                       `gorm:"column:version;comment:'排序'"`         // 排序
	Title         x_null.String                      `gorm:"comment:'标题'"`                        // 标题
	Content       x_null.String                      `gorm:"comment:'协议内容'"`                      // 协议内容
	CreatedBy     x_null.String                      `gorm:"column:created_by;type:char(36);comment:'创建人'"`
	CreatedByUser system_model.SystemAuthAdminSimple `gorm:"foreignKey:CreatedBy"`
	IsDelete      soft_delete.DeletedAt              `gorm:"not null;default:0;softDelete:flag,DeletedAtField:DeleteTime;comment:'是否删除: 0=否, 1=是'"`
	CreateTime    x_null.Time                        `gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
	UpdateTime    x_null.Time                        `gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间
	DeleteTime    x_null.Time                        `gorm:"comment:'删除时间'"`                // 删除时间
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
