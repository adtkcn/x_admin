package model

import (
	"x_admin/core"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// MonitorClient 监控-客户端信息实体
type MonitorClient struct {
	Id         string `gorm:"primarykey;type:char(36);comment:'uuid'"` // uuid
	ProjectKey string `gorm:"comment:'项目key'"`                         // 项目key
	ClientId   string `gorm:"comment:'sdk生成的客户端id'"`                   // sdk生成的客户端id
	// UserId     string        `gorm:"comment:'用户id'"`                          // 用户id
	Os      string `gorm:"comment:'系统'"`  // 系统
	Browser string `gorm:"comment:'浏览器'"` // 浏览器
	// Country    string        `gorm:"comment:'国家'"`                            // 国家
	// Province   string        `gorm:"comment:'省份'"`                            // 省份
	// City       string        `gorm:"comment:'城市'"`                            // 城市
	// Operator   string        `gorm:"comment:'电信运营商'"`                         // 电信运营商
	// Ip         string        `gorm:"comment:'ip'"`                            // ip
	Ua         string        `gorm:"comment:'ua记录'"`                // ua记录
	CreateTime core.NullTime `gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
}

// BeforeCreate 在创建前生成UUIDv7
func (m *MonitorClient) BeforeCreate(tx *gorm.DB) (err error) {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	m.Id = id.String()
	return nil
}
