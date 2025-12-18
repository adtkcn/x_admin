package common_model

import (
	"x_admin/core"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

// Album 相册实体
type Album struct {
	ID         string                `gorm:"primarykey;type:char(36);comment:'主键id'"`
	Cid        string                `gorm:"not null;comment:'类目ID'"`
	AdminId    uint                  `gorm:"not null;default:0;comment:'管理员ID'"`
	Uid        uint                  `gorm:"not null;default:0;comment:'用户ID'"`
	Name       string                `gorm:"not null;default:'';comment:'文件名称''"`
	Uri        string                `gorm:"not null;comment:'文件路径'"`
	Ext        string                `gorm:"not null;default:'';comment:'文件扩展'"`
	Hash       string                `gorm:"not null;default:'';comment:'文件hash'"`
	Size       int64                 `gorm:"not null;default:0;comment:文件大小"`
	IsDelete   soft_delete.DeletedAt `gorm:"not null;default:0;softDelete:flag,DeletedAtField:DeleteTime;comment:'是否删除: 0=否, 1=是'"`
	CreateTime core.NullTime         `gorm:"autoCreateTime;not null;comment:'创建时间'"`
	UpdateTime core.NullTime         `gorm:"autoUpdateTime;not null;comment:'更新时间'"`
	DeleteTime core.NullTime         `gorm:"default:null;comment:'删除时间'"`
}

// 自动在创建时设置 UUIDv7
func (u *Album) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	u.ID = id.String()
	return nil
}

// AlbumCate 相册分类实体
type AlbumCate struct {
	ID         string                `gorm:"primarykey;type:char(36);comment:'主键id'"`
	Pid        string                `gorm:"not null;default:'';comment:'父级ID'"`
	AdminId    uint                  `gorm:"not null;default:0;comment:'管理员ID'"`
	Name       string                `gorm:"not null;default:'';comment:'分类名称''"`
	IsDelete   soft_delete.DeletedAt `gorm:"not null;default:0;softDelete:flag,DeletedAtField:DeleteTime;comment:'是否删除: 0=否, 1=是'"`
	CreateTime core.NullTime         `gorm:"autoCreateTime;not null;comment:'创建时间'"`
	UpdateTime core.NullTime         `gorm:"autoUpdateTime;not null;comment:'更新时间'"`
	DeleteTime core.NullTime         `gorm:"default:null;comment:'删除时间'"`
}

// 自动在创建时设置 UUIDv7
func (u *AlbumCate) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	u.ID = id.String()
	return nil
}
