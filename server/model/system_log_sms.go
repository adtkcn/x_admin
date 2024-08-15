package model

import (
	"x_admin/core"
)

// SystemLogSms 系统短信日志实体
type SystemLogSms struct {
	Id         int           `gorm:"primarykey;comment:'id'"`                // id
	Scene      core.NullInt  `gorm:"comment:'场景编号'"`                         // 场景编号
	Mobile     string        `gorm:"comment:'手机号码'"`                         // 手机号码
	Content    string        `gorm:"comment:'发送内容'"`                         // 发送内容
	Status     core.NullInt  `gorm:"comment:'发送状态：[0=发送中, 1=发送成功, 2=发送失败]'"` // 发送状态：[0=发送中, 1=发送成功, 2=发送失败]
	Results    string        `gorm:"comment:'短信结果'"`                         // 短信结果
	SendTime   core.NullTime `gorm:"comment:'发送时间'"`                         // 发送时间
	CreateTime core.NullTime `gorm:"autoCreateTime;comment:'创建时间'"`          // 创建时间
	UpdateTime core.NullTime `gorm:"autoUpdateTime;comment:'更新时间'"`          // 更新时间
}
