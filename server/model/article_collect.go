package model

import (
	"x_admin/core"

	"gorm.io/plugin/soft_delete"
)

// ArticleCollect 文章收藏实体
type ArticleCollect struct {
	Id         int                   `gorm:"primarykey;comment:'主键'"` // 主键
	UserId     int                   `gorm:"comment:'用户ID'"`          // 用户ID
	ArticleId  int                   `gorm:"comment:'文章ID'"`          // 文章ID
	IsDelete   soft_delete.DeletedAt `gorm:"not null;default:0;softDelete:flag,DeletedAtField:DeleteTime;comment:'是否删除: 0=否, 1=是'"`
	UpdateTime core.NullTime         `gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间
	CreateTime core.NullTime         `gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
	DeleteTime core.NullTime         `gorm:"default:null;comment:'删除时间'"`   // 删除时间
}
