package model

import (
	"x_admin/core"
)

// SystemLogSms 系统短信日志实体
type SystemLogSms struct {
	Id         int         `json:"id" gorm:"primarykey;" excel:"name:id;"`                             // id
	Scene      int         `json:"scene" excel:"name:场景编号;"`                                           // 场景编号
	Mobile     string      `json:"mobile" excel:"name:手机号码;"`                                          // 手机号码
	Content    string      `json:"content" excel:"name:发送内容;"`                                         // 发送内容
	Status     int         `json:"status" excel:"name:发送状态：[0=发送中, 1=发送成功, 2=发送失败];"`                  // 发送状态：[0=发送中, 1=发送成功, 2=发送失败]
	Results    string      `json:"results" excel:"name:短信结果;"`                                         // 短信结果
	SendTime   int         `mapstructure:"send_time" excel:"name:发送时间;"`                               // 发送时间
	CreateTime core.TsTime `json:"CreateTime" gorm:"autoCreateTime;comment:'创建时间'" excel:"name:创建时间;"` // 创建时间
	UpdateTime core.TsTime `json:"UpdateTime" gorm:"autoUpdateTime;comment:'更新时间'" excel:"name:更新时间;"` // 更新时间
}
