package common_model

import (
	"github.com/adtkcn/x_null"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CommonFileRef 文件关联表 — 记录文件在业务中的引用关系
// 配合 x_common_file_hash 使用，用于判断文件是否被业务使用
type CommonFileRef struct {
	ID         string      `gorm:"primaryKey;type:char(36);comment:'主键id'"`
	FileHashID string      `gorm:"not null;index:idx_file_hash_id;comment:'文件哈希ID(x_common_file_hash.id)'"`
	BizType    string      `gorm:"not null;default:'';index:idx_biz;comment:'业务类型: user_avatar/article_cover/album等'"`
	BizID      string      `gorm:"not null;default:'';index:idx_biz;comment:'业务实体ID'"`
	CreateTime x_null.Time `gorm:"autoCreateTime;not null;comment:'创建时间'"`
}

// TableName 指定表名
func (CommonFileRef) TableName() string {
	return "x_common_file_ref"
}

// BeforeCreate 自动在创建时设置 UUIDv7
func (u *CommonFileRef) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	u.ID = id.String()
	return nil
}
