package model

import (
	"uuid"

	"github.com/adtkcn/x_null"
	"gorm.io/gorm"
)

// SystemNotice 系统通知记录
type SystemNotice struct {
	ID         string      `gorm:"primarykey;type:char(36);comment:'UUIDv7'"`
	Type       string      `gorm:"not null;default:'';type:varchar(32);comment:'通知类型:  success|danger|primary|info|warning '"`
	Title      string      `gorm:"not null;default:'';type:varchar(200);comment:'通知标题'"`
	Content    string      `gorm:"type:text;comment:'通知正文'"`
	ReceiverID string      `gorm:"not null;type:char(36);index;comment:'接收人ID'"`
	SenderID   string      `gorm:"not null;default:'';type:char(36);comment:'发送人ID,系统通知为空'"`
	URL        string      `gorm:"not null;default:'';type:varchar(500);comment:'前端跳转路径'"`
	IsRead     uint8       `gorm:"not null;default:0;comment:'0未读 1已读'"`
	IsEmailed  int8        `gorm:"not null;default:0;type:tinyint;comment:'邮件推送状态: -1不发送 0待发送 1发送中(已进入队列) 2发送成功 3发送失败'"`
	ReadTime   x_null.Time `gorm:"default:null;comment:'阅读时间'"`
	Extra      string      `gorm:"type:json;comment:'扩展数据(JSON)'"`
	CreateTime x_null.Time `gorm:"autoCreateTime;not null;comment:'创建时间'"`
	UpdateTime x_null.Time `gorm:"autoUpdateTime;not null;comment:'更新时间'"`
}

func (m *SystemNotice) TableName() string {
	return "x_system_notice"
}

func (m *SystemNotice) BeforeCreate(tx *gorm.DB) (err error) {
	id := uuid.NewV7()
	m.ID = id.String()
	return nil
}

// SystemNoticeSetting 用户通知渠道偏好
type SystemNoticeSetting struct {
	ID         string      `gorm:"primarykey;type:char(36);comment:'UUIDv7'"`
	AdminID    string      `gorm:"not null;uniqueIndex:uk_admin_channel;type:char(36);comment:'用户ID'"`
	Channel    string      `gorm:"not null;uniqueIndex:uk_admin_channel;type:varchar(32);comment:'渠道: email/app'"`
	IsEnabled  uint8       `gorm:"not null;default:0;comment:'0关闭 1开启'"`
	CreateTime x_null.Time `gorm:"autoCreateTime;not null;comment:'创建时间'"`
	UpdateTime x_null.Time `gorm:"autoUpdateTime;not null;comment:'更新时间'"`
}

func (m *SystemNoticeSetting) TableName() string {
	return "x_system_notice_setting"
}

func (m *SystemNoticeSetting) BeforeCreate(tx *gorm.DB) (err error) {
	id := uuid.NewV7()
	m.ID = id.String()
	return nil
}
