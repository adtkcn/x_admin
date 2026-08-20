package common_model

import (
	"github.com/adtkcn/x_null"
	"uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

// Album 相册实体
// 文件信息（uri/ext/hash/size）不冗余存储，统一通过 FileHashId 关联 x_common_file_hash 查询获得。
type Album struct {
	ID         string                `gorm:"primarykey;type:char(36);comment:'主键id'"`
	Cid        string                `gorm:"not null;comment:'类目ID'"`
	AdminId    string                `gorm:"not null;default:'';comment:'管理员ID'"`
	Name       string                `gorm:"not null;default:'';comment:'文件名称''"`
	FileHashId string                `gorm:"not null;default:'';index:idx_file_hash_id;comment:'关联文件哈希ID(x_common_file_hash.id)，uri/ext/hash/size 经此关联查询'"`
	IsDelete   soft_delete.DeletedAt `gorm:"not null;default:0;softDelete:flag,DeletedAtField:DeleteTime;comment:'是否删除: 0=否, 1=是'"`
	CreateTime x_null.Time           `gorm:"autoCreateTime;not null;comment:'创建时间'"`
	UpdateTime x_null.Time           `gorm:"autoUpdateTime;not null;comment:'更新时间'"`
	DeleteTime x_null.Time           `gorm:"default:null;comment:'删除时间'"`
}

// 自动在创建时设置 UUIDv7
func (u *Album) BeforeCreate(tx *gorm.DB) error {
	id := uuid.NewV7()
	u.ID = id.String()
	return nil
}

// AlbumCate 相册分类实体
type AlbumCate struct {
	ID         string                `gorm:"primarykey;type:char(36);comment:'主键id'"`
	Pid        string                `gorm:"not null;default:'';comment:'父级ID'"`
	AdminId    string                `gorm:"not null;default:'';comment:'管理员ID'"`
	Name       string                `gorm:"not null;default:'';comment:'分类名称''"`
	IsDelete   soft_delete.DeletedAt `gorm:"not null;default:0;softDelete:flag,DeletedAtField:DeleteTime;comment:'是否删除: 0=否, 1=是'"`
	CreateTime x_null.Time           `gorm:"autoCreateTime;not null;comment:'创建时间'"`
	UpdateTime x_null.Time           `gorm:"autoUpdateTime;not null;comment:'更新时间'"`
	DeleteTime x_null.Time           `gorm:"default:null;comment:'删除时间'"`
}

// 自动在创建时设置 UUIDv7
func (u *AlbumCate) BeforeCreate(tx *gorm.DB) error {
	id := uuid.NewV7()
	u.ID = id.String()
	return nil
}
