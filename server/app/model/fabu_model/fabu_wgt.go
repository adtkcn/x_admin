package fabu_model

import (
	"uuid"

	"github.com/adtkcn/x_null"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

// FabuWgt 热更新包子表（uni-app wgt），隶属于某个应用版本
type FabuWgt struct {
	ID           string                `gorm:"primarykey;type:char(36);comment:'主键'" json:"id"`
	AppId        string                `gorm:"type:char(36);not null;default:'';index:idx_app;comment:'应用id(冗余)'" json:"app_id"`
	VersionId    string                `gorm:"type:char(36);not null;default:'';index:idx_version;comment:'所属版本id'" json:"version_id"`
	Version      string                `gorm:"type:varchar(64);not null;default:'';comment:'版本号'" json:"version"`
	VersionCode  int                   `gorm:"not null;default:0;comment:'版本code'" json:"version_code"`
	DownloadUrl  string                `gorm:"type:varchar(512);not null;default:'';comment:'下载地址'" json:"download_url"`
	Md5          string                `gorm:"type:varchar(64);not null;default:'';comment:'文件MD5'" json:"md5"`
	Size         int64                 `gorm:"not null;default:0;comment:'包大小(字节)'" json:"size"`
	IsDelete     soft_delete.DeletedAt `gorm:"not null;default:0;softDelete:flag,DeletedAtField:DeleteTime;comment:'是否删除'" json:"-"`
	CreateTime   x_null.Time           `gorm:"autoCreateTime;not null;comment:'创建时间'" json:"create_time"`
	UpdateTime   x_null.Time           `gorm:"autoUpdateTime;not null;comment:'更新时间'" json:"update_time"`
	DeleteTime   x_null.Time           `gorm:"default:null;comment:'删除时间'" json:"-"`
}

func (FabuWgt) TableName() string { return "x_fabu_wgt" }

func (m *FabuWgt) BeforeCreate(tx *gorm.DB) error {
	id := uuid.NewV7()
	m.ID = id.String()
	return nil
}
