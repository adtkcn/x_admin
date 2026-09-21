package fabu_model

import (
	"uuid"

	"github.com/adtkcn/x_null"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

// FabuApp 应用表
type FabuApp struct {
	ID            string                `gorm:"primarykey;type:char(36);comment:'主键'" json:"id"`
	Name          string                `gorm:"type:varchar(255);not null;default:'';comment:'应用名称'" json:"name"`
	Platform      string                `gorm:"type:varchar(16);not null;default:'';comment:'平台 ios/android'" json:"platform"`
	BundleId      string                `gorm:"type:varchar(255);not null;default:'';index:idx_bundle;comment:'包名/BundleId'" json:"bundle_id"`
	BundleName    string                `gorm:"type:varchar(255);not null;default:'';comment:'BundleName'" json:"bundle_name"`
	ShortUrl      string                `gorm:"type:varchar(64);not null;default:'';uniqueIndex:idx_short;comment:'短链'" json:"short_url"`
	Icon          string                `gorm:"type:varchar(512);not null;default:'';comment:'图标URL'" json:"icon"`
	DownloadTimes int                   `gorm:"not null;default:0;comment:'累计下载次数'" json:"download_times"`
	IsDelete      soft_delete.DeletedAt `gorm:"not null;default:0;softDelete:flag,DeletedAtField:DeleteTime;comment:'是否删除'" json:"-"`
	CreateTime    x_null.Time           `gorm:"autoCreateTime;not null;comment:'创建时间'" json:"create_time"`
	UpdateTime    x_null.Time           `gorm:"autoUpdateTime;not null;comment:'更新时间'" json:"update_time"`
	DeleteTime    x_null.Time           `gorm:"default:null;comment:'删除时间'" json:"-"`
}

func (FabuApp) TableName() string { return "x_fabu_app" }

func (m *FabuApp) BeforeCreate(tx *gorm.DB) error {
	id := uuid.NewV7()
	m.ID = id.String()
	return nil
}
