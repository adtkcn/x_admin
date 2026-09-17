package fabu_model

import (
	"uuid"

	"github.com/adtkcn/x_null"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

// FabuAppVersion 应用版本表
type FabuAppVersion struct {
	ID            string                `gorm:"primarykey;type:char(36);comment:'主键'" json:"id"`
	AppId         string                `gorm:"type:char(36);not null;default:'';index:idx_app;comment:'应用id'" json:"app_id"`
	Version       string                `gorm:"type:varchar(64);not null;default:'';comment:'版本号'" json:"version"`
	VersionCode   int                   `gorm:"not null;default:0;comment:'版本code'" json:"version_code"`
	Size          int64                 `gorm:"not null;default:0;comment:'包大小(字节)'" json:"size"`
	Md5           string                `gorm:"type:varchar(64);not null;default:'';comment:'文件MD5'" json:"md5"`
	DownloadUrl   string                `gorm:"type:varchar(512);not null;default:'';comment:'下载地址'" json:"download_url"`
	InstallUrl    string                `gorm:"type:varchar(512);not null;default:'';comment:'安装地址(iOS plist)'" json:"install_url"`
	Released      bool                  `gorm:"not null;default:0;comment:'是否已发布'" json:"released"`
	UpdateMode    int                   `gorm:"not null;default:0;comment:'更新模式 0普通 1强制 2静默'" json:"update_mode"`
	Gray          bool                  `gorm:"not null;default:0;comment:'是否灰度'" json:"gray"`
	DownloadTimes int                   `gorm:"not null;default:0;comment:'下载次数'" json:"download_times"`
	IsDelete      soft_delete.DeletedAt `gorm:"not null;default:0;softDelete:flag,DeletedAtField:DeleteTime;comment:'是否删除'" json:"-"`
	CreateTime    x_null.Time           `gorm:"autoCreateTime;not null;comment:'创建时间'" json:"create_time"`
	UpdateTime    x_null.Time           `gorm:"autoUpdateTime;not null;comment:'更新时间'" json:"update_time"`
	DeleteTime    x_null.Time           `gorm:"default:null;comment:'删除时间'" json:"-"`
}

func (FabuAppVersion) TableName() string { return "x_fabu_app_version" }

func (m *FabuAppVersion) BeforeCreate(tx *gorm.DB) error {
	id := uuid.NewV7()
	m.ID = id.String()
	return nil
}
