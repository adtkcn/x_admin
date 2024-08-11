package model

import (
	"x_admin/core"
)

// SystemLogSms 系统短信日志实体
type SystemLogSms struct {
	Id         int         `json:"id" gorm:"primarykey;"`                           // id
	Scene      int         `json:"scene"`                                           // 场景编号
	Mobile     string      `json:"mobile"`                                          // 手机号码
	Content    string      `json:"content"`                                         // 发送内容
	Status     int         `json:"status"`                                          // 发送状态：[0=发送中, 1=发送成功, 2=发送失败]
	Results    string      `json:"results"`                                         // 短信结果
	SendTime   int         `json:"send_time" mapstructure:"send_time"`              // 发送时间
	CreateTime core.TsTime `json:"CreateTime" gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
	UpdateTime core.TsTime `json:"UpdateTime" gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间
}
