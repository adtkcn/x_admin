package user_model

import (
	"github.com/adtkcn/x_null"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

// User 用户主表（C端用户，邮箱为主账号）
type User struct {
	ID string `gorm:"primarykey;type:char(36);comment:'UUID v7'" json:"id"`
	// 注意：本表为软删除表（IsDelete flag），不再使用 DB 唯一索引。
	// 否则被软删除的邮箱仍被唯一约束占用，导致无法用同一邮箱重新注册。
	// 邮箱唯一性改由业务层 existsActiveUser 判断（只统计未删除记录）。
	Email     string `gorm:"type:varchar(128);not null;default:'';comment:'邮箱(主账号)'" json:"email"`
	Nickname  string `gorm:"type:varchar(64);not null;default:'';comment:'昵称'" json:"nickname"`
	Avatar    string `gorm:"type:varchar(255);not null;default:'';comment:'头像'" json:"avatar"`
	Password  string `gorm:"type:varchar(255);not null;default:'';comment:'密码(bcrypt)'" json:"-"`
	Salt      string `gorm:"type:varchar(32);not null;default:'';comment:'加密盐'" json:"-"`
	Phone     string `gorm:"type:varchar(20);not null;default:'';comment:'手机号'" json:"phone"`
	PhoneCode string `gorm:"type:varchar(10);not null;default:'86';comment:'手机号区号'" json:"phoneCode"`
	Status    uint8  `gorm:"type:tinyint;not null;default:0;comment:'状态 0正常 1禁用'" json:"status"`
	TokenVersion        int64                 `gorm:"type:bigint;not null;default:0;comment:'Token版本号(踢人下线:自增使旧token失效)'" json:"-"`
	LastLoginIp         string                `gorm:"type:varchar(50);not null;default:'';comment:'最后登录IP'" json:"lastLoginIp"`
	LastLoginTime       x_null.Time           `gorm:"default:null;comment:'最后登录时间'" json:"last_login_time"`
	IsDelete            soft_delete.DeletedAt `gorm:"not null;default:0;softDelete:flag,DeletedAtField:DeleteTime;comment:'是否删除: 0=否, 1=是'" json:"-"`
	CreateTime          x_null.Time           `gorm:"autoCreateTime;not null;comment:'创建时间'" json:"create_time"`
	UpdateTime          x_null.Time           `gorm:"autoUpdateTime;not null;comment:'更新时间'" json:"update_time"`
	DeleteTime          x_null.Time           `gorm:"default:null;comment:'删除时间'" json:"-"`
}

// BeforeCreate 在创建前生成UUIDv7
func (m *User) BeforeCreate(tx *gorm.DB) (err error) {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	m.ID = id.String()
	return nil
}
