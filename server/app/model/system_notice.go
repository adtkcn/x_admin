package model

import (
	"github.com/adtkcn/x_null"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// SystemNotice 系统通知记录
type SystemNotice struct {
	ID         string      `gorm:"primarykey;type:char(36);comment:'UUIDv7'" json:"id"`
	Type       string      `gorm:"not null;default:'';type:varchar(32);comment:'通知类型: flow_pass/flow_back/flow_new/flow_finish/system'" json:"type"`
	Title      string      `gorm:"not null;default:'';type:varchar(200);comment:'通知标题'" json:"title"`
	Content    string      `gorm:"type:text;comment:'通知正文'" json:"content"`
	ReceiverID string      `gorm:"not null;type:char(36);index;comment:'接收人ID'" json:"receiverId"`
	SenderID   string      `gorm:"not null;default:'';type:char(36);comment:'发送人ID,系统通知为0'" json:"senderId"`
	URL        string      `gorm:"not null;default:'';type:varchar(500);comment:'前端跳转路径'" json:"url"`
	IsRead     uint8       `gorm:"not null;default:0;comment:'0未读 1已读'" json:"isRead"`
	IsEmailed  uint8       `gorm:"not null;default:0;comment:'0未发送 1已发送邮件'" json:"isEmailed"`
	ReadTime   x_null.Time `gorm:"default:null;comment:'阅读时间'" json:"readTime"`
	Extra      string      `gorm:"type:json;comment:'扩展数据(JSON)'" json:"extra"`
	CreateTime x_null.Time `gorm:"autoCreateTime;not null;comment:'创建时间'" json:"createTime"`
	UpdateTime x_null.Time `gorm:"autoUpdateTime;not null;comment:'更新时间'" json:"updateTime"`
}

func (m *SystemNotice) TableName() string {
	return "x_system_notice"
}

func (m *SystemNotice) BeforeCreate(tx *gorm.DB) (err error) {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	m.ID = id.String()
	return nil
}

// SystemNoticeSetting 用户通知渠道偏好
type SystemNoticeSetting struct {
	ID         string      `gorm:"primarykey;type:char(36);comment:'UUIDv7'" json:"id"`
	AdminID    string      `gorm:"not null;uniqueIndex:uk_admin_channel;type:char(36);comment:'用户ID'" json:"adminId"`
	Channel    string      `gorm:"not null;uniqueIndex:uk_admin_channel;type:varchar(32);comment:'渠道: site/email/app'" json:"channel"`
	IsEnabled  uint8       `gorm:"not null;default:0;comment:'0关闭 1开启'" json:"isEnabled"`
	CreateTime x_null.Time `gorm:"autoCreateTime;not null;comment:'创建时间'" json:"createTime"`
	UpdateTime x_null.Time `gorm:"autoUpdateTime;not null;comment:'更新时间'" json:"updateTime"`
}

func (m *SystemNoticeSetting) TableName() string {
	return "x_system_notice_setting"
}

func (m *SystemNoticeSetting) BeforeCreate(tx *gorm.DB) (err error) {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	m.ID = id.String()
	return nil
}
