package model

import "x_admin/core"

//MonitorWeb 错误收集error实体
type MonitorWeb struct {
	Id int `gorm:"primarykey;comment:'uuid'"` // uuid

	ProjectKey string `gorm:"comment:'项目key'"` // 项目key

	ClientId string `gorm:"comment:'sdk生成的客户端id'"` // sdk生成的客户端id

	EventType string `gorm:"comment:'事件类型'"` // 事件类型

	Page string `gorm:"comment:'URL地址'"` // URL地址

	Message string `gorm:"comment:'错误消息'"` // 错误消息

	Stack string `gorm:"comment:'错误堆栈'"` // 错误堆栈

	ClientTime core.NullTime `gorm:"comment:'客户端时间'"` // 客户端时间

	CreateTime core.NullTime `gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间

}
