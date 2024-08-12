package model

import (
	"x_admin/core"
)

// SystemLogSms 系统短信日志实体
type SystemLogSms struct {
	Id         int         `gorm:"primarykey;"`                   // id
	Scene      int         ``                                     // 场景编号
	Mobile     string      ``                                     // 手机号码
	Content    string      ``                                     // 发送内容
	Status     int         ``                                     // 发送状态：[0=发送中, 1=发送成功, 2=发送失败]
	Results    string      ``                                     // 短信结果
	SendTime   int         ``                                     // 发送时间
	CreateTime core.TsTime `gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
	UpdateTime core.TsTime `gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间
}

// func (t *SystemLogSms) BeforeCreate(tx *gorm.DB) (err error) {
// 	// u.Id = uuid.New()
// 	// t.CreateTime = core.NowTime().String()

// 	t.CreateTime = core.NowTime()
// 	t.UpdateTime = core.NowTime()
// 	tx.Statement.SetColumn("CreateTime", t.CreateTime.String())
// 	tx.Statement.SetColumn("UpdateTime", t.UpdateTime.String())
// 	return
// }
// func (t *SystemLogSms) BeforeSave(tx *gorm.DB) (err error) {

// 	t.UpdateTime = core.NowTime()
// 	tx.Statement.SetColumn("UpdateTime", t.UpdateTime.String())
// 	return
// }
// func (t *SystemLogSms) BeforeDelete(tx *gorm.DB) (err error) {
// 	t.UpdateTime = core.NowTime()
// 	tx.Statement.SetColumn("UpdateTime", t.UpdateTime.String())
// 	return
// }
