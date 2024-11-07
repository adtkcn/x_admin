package model

import (
	"x_admin/core"
)

// MonitorSlow 监控-错误列实体
type MonitorSlow struct {
	Id         int            `gorm:"primarykey;comment:'错误id'"`     // 错误id
	ProjectKey string         `gorm:"comment:'项目key'"`               // 项目key
	ClientId   string         `gorm:"comment:'sdk生成的客户端id'"`         // sdk生成的客户端id
	UserId     string         `gorm:"comment:'用户id'"`                // 用户id
	Path       string         `gorm:"comment:'URL地址'"`               // URL地址
	Time       core.NullFloat `gorm:"comment:'时间'"`                  // 时间
	CreateTime core.NullTime  `gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
}
