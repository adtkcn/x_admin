package common_model

import (
	"github.com/adtkcn/x_null"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CommonFileHash 文件哈希记录表（用于秒传）
type CommonFileHash struct {
	ID         string      `gorm:"primaryKey;type:char(36);comment:'主键id'"`
	FileMd5    string      `gorm:"not null;uniqueIndex:idx_md5;comment:'文件MD5值'"`
	FileSize   int64       `gorm:"not null;default:0;comment:'文件大小(字节)'"`
	FilePath   string      `gorm:"not null;comment:'存储路径/对象key'"`
	Ext        string      `gorm:"not null;default:'';comment:'文件扩展名'"`
	CreateTime x_null.Time `gorm:"autoCreateTime;not null;comment:'创建时间'"`
}

// TableName 指定表名
func (CommonFileHash) TableName() string {
	return "x_common_file_hash"
}

// BeforeCreate 自动在创建时设置 UUIDv7
func (u *CommonFileHash) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	u.ID = id.String()
	return nil
}
