package monitor_model

import (
	"uuid"

	"github.com/adtkcn/x_null"
	"gorm.io/gorm"
)

// MonitorClient 监控-客户端信息实体
type MonitorClient struct {
	Id         string      `gorm:"primarykey;type:char(36);comment:'uuid'"` // uuid
	ProjectKey string      `gorm:"comment:'项目key'"`                         // 项目key
	ClientId   string      `gorm:"uniqueIndex;comment:'sdk生成的客户端id'"`       // sdk生成的客户端id
	Os         string      `gorm:"comment:'系统'"`                            // 系统
	Browser    string      `gorm:"comment:'浏览器'"`                           // 浏览器
	Ua         string      `gorm:"comment:'ua记录'"`                          // ua记录
	CreateTime x_null.Time `gorm:"autoCreateTime;comment:'创建时间'"`           // 创建时间
}

// BeforeCreate 在创建前生成UUIDv7
func (m *MonitorClient) BeforeCreate(tx *gorm.DB) (err error) {
	id := uuid.NewV7()
	m.Id = id.String()
	return nil
}
