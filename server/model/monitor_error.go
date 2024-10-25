package model

import (
	"x_admin/core"
)

// MonitorError 监控-错误列实体
type MonitorError struct {
	Id         int           `gorm:"primarykey;comment:'错误id'"`     // 错误id
	ProjectKey string        `gorm:"comment:'项目key'"`               // 项目key
	EventType  string        `gorm:"comment:'事件类型'"`                // 事件类型
	Path       string        `gorm:"comment:'URL地址'"`               // URL地址
	Message    string        `gorm:"comment:'错误消息'"`                // 错误消息
	Stack      string        `gorm:"comment:'错误堆栈'"`                // 错误堆栈
	Md5        string        `gorm:"comment:'md5'"`                 // md5
	CreateTime core.NullTime `gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
}
