package model

import (
	"x_admin/core"
)

// MonitorErrorList 错误对应的用户记录实体
type MonitorErrorList struct {
	Id         int           `gorm:"primarykey;comment:'id'"`
	Eid        string        `gorm:"comment:'错误id'"`                // 错误表id
	Cid        string        `gorm:"comment:'客户端id'"`               // 客户端表id
	Width      core.NullInt  `gorm:"comment:'屏幕'"`                  // 屏幕
	Height     core.NullInt  `gorm:"comment:'屏幕高度'"`                // 屏幕高度
	CreateTime core.NullTime `gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
}
