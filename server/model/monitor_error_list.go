package model

import (
	"x_admin/core"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// MonitorErrorList 错误对应的用户记录实体
type MonitorErrorList struct {
	Id       string `gorm:"primarykey;type:char(36);comment:'id'"` // id
	ErrorId  string `gorm:"comment:'错误id'"`                        // 错误表id
	ClientId string `gorm:"comment:'sdk生成的客户端id'"`                 // sdk生成的客户端id
	UserId   string `gorm:"comment:'业务中用户id'"`                     // 用户id

	Width  core.NullInt `gorm:"comment:'屏幕'"`   // 屏幕
	Height core.NullInt `gorm:"comment:'屏幕高度'"` // 屏幕高度

	Country  string `gorm:"comment:'国家'"`    // 国家
	Province string `gorm:"comment:'省份'"`    // 省份
	City     string `gorm:"comment:'城市'"`    // 城市
	Operator string `gorm:"comment:'电信运营商'"` // 电信运营商
	Ip       string `gorm:"comment:'ip'"`    // ip
	// Ua       string `gorm:"comment:'ua记录'"`  // ua记录

	CreateTime core.NullTime `gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
}

// BeforeCreate 在创建前生成UUIDv7
func (m *MonitorErrorList) BeforeCreate(tx *gorm.DB) (err error) {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	m.Id = id.String()
	return nil
}
