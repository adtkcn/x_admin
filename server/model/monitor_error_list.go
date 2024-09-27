package model

import (
	"x_admin/core"
)

// MonitorErrorList 错误对应的用户记录实体
type MonitorErrorList struct {
	Id         int           `gorm:"primarykey;comment:'项目id'"`     // 项目id
	ErrorId    string        `gorm:"comment:'错误id'"`                // 错误id
	ClientId   string        `gorm:"comment:'客户端id'"`               // 客户端id
	ProjectKey string        `gorm:"comment:'项目key'"`               // 项目id
	CreateTime core.NullTime `gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
}
