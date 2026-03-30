package model

import (
	"github.com/adtkcn/x_null"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// MonitorError 监控-错误列实体
type MonitorError struct {
	Id         string      `gorm:"primarykey;type:char(36);comment:'错误id'"` // 错误id
	ProjectKey string      `gorm:"comment:'项目key'"`                         // 项目key
	EventType  string      `gorm:"comment:'事件类型'"`                          // 事件类型
	Path       string      `gorm:"comment:'URL地址'"`                         // URL地址
	Message    string      `gorm:"comment:'错误消息'"`                          // 错误消息
	Stack      string      `gorm:"comment:'错误堆栈'"`                          // 错误堆栈
	Md5        string      `gorm:"comment:'md5'"`                           // md5
	CreateTime x_null.Time `gorm:"autoCreateTime;comment:'创建时间'"`           // 创建时间
}

// BeforeCreate 在创建前生成UUIDv7
func (m *MonitorError) BeforeCreate(tx *gorm.DB) (err error) {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	m.Id = id.String()
	return nil
}
