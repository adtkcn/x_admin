package common_model

import (
	"x_admin/core"

	"gorm.io/plugin/soft_delete"
)

// Album 相册实体
type Album struct {
	ID         uint                  `gorm:"primarykey;comment:'主键ID'"`
	Cid        uint                  `gorm:"not null;default:0;comment:'类目ID'"`
	Aid        uint                  `gorm:"not null;default:0;comment:'管理ID'"`
	Uid        uint                  `gorm:"not null;default:0;comment:'用户ID'"`
	Type       int                   `gorm:"not null;default:10;comment:'文件类型: [10=图片, 20=视频]''"`
	Name       string                `gorm:"not null;default:'';comment:'文件名称''"`
	Uri        string                `gorm:"not null;comment:'文件路径'"`
	Ext        string                `gorm:"not null;default:'';comment:'文件扩展'"`
	Size       int64                 `gorm:"not null;default:0;comment:文件大小"`
	IsDelete   soft_delete.DeletedAt `gorm:"not null;default:0;softDelete:flag,DeletedAtField:DeleteTime;comment:'是否删除: 0=否, 1=是'"`
	CreateTime core.NullTime         `gorm:"autoCreateTime;not null;comment:'创建时间'"`
	UpdateTime core.NullTime         `gorm:"autoUpdateTime;not null;comment:'更新时间'"`
	DeleteTime core.NullTime         `gorm:"default:null;comment:'删除时间'"`
}

// AlbumCate 相册分类实体
type AlbumCate struct {
	ID         uint                  `gorm:"primarykey;comment:'主键ID'"`
	Pid        uint                  `gorm:"not null;default:0;comment:'父级ID'"`
	Type       int                   `gorm:"not null;default:10;comment:'文件类型: [10=图片, 20=视频]''"`
	Name       string                `gorm:"not null;default:'';comment:'分类名称''"`
	IsDelete   soft_delete.DeletedAt `gorm:"not null;default:0;softDelete:flag,DeletedAtField:DeleteTime;comment:'是否删除: 0=否, 1=是'"`
	CreateTime core.NullTime         `gorm:"autoCreateTime;not null;comment:'创建时间'"`
	UpdateTime core.NullTime         `gorm:"autoUpdateTime;not null;comment:'更新时间'"`
	DeleteTime core.NullTime         `gorm:"default:null;comment:'删除时间'"`
}
