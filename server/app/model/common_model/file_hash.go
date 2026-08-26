package common_model

import (
	"uuid"

	"github.com/adtkcn/x_null"
	"gorm.io/gorm"
)

// CommonFileHash 文件哈希记录表（用于秒传）
type CommonFileHash struct {
	ID             string      `gorm:"primaryKey;type:char(36);comment:'主键id'"`
	Pid            string      `gorm:"not null;default:'';index:idx_pid;comment:'父文件id（主记录为空，webp/缩略图等派生记录指向原图id）'"`
	Quality        int         `gorm:"not null;default:0;comment:'转码质量（webp 有损压缩质量，0 表示原图/未转码）'"`
	ScaleWidth     int         `gorm:"not null;default:0;comment:'缩放宽度（0 表示不缩放）'"`
	ScaleHeight    int         `gorm:"not null;default:0;comment:'缩放高度（0 表示不缩放）'"`
	FileMd5        string      `gorm:"not null;uniqueIndex:idx_md5;comment:'文件MD5值（派生记录填派生文件自身的MD5）'"`
	FileSize       int64       `gorm:"not null;default:0;comment:'文件大小(字节)'"`
	FilePath       string      `gorm:"not null;comment:'存储路径/对象key'"`
	Ext            string      `gorm:"not null;default:'';comment:'文件扩展名'"`
	LastAccessTime x_null.Time `gorm:"default:null;comment:'最后访问时间(标记冷热文件,后续用于清理)'"`
	CreateTime     x_null.Time `gorm:"autoCreateTime;not null;comment:'创建时间'"`
}

// TableName 指定表名
func (CommonFileHash) TableName() string {
	return "x_common_file_hash"
}

// BeforeCreate 自动在创建时设置 UUIDv7
func (u *CommonFileHash) BeforeCreate(tx *gorm.DB) error {
	id := uuid.NewV7()
	u.ID = id.String()
	return nil
}
