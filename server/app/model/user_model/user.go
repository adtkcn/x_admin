package user_model

import (
	"github.com/adtkcn/x_null"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

// User 用户主表（C端用户，邮箱为主账号）
type User struct {
	ID            string                `gorm:"primarykey;type:char(36);comment:'UUID v7'" json:"id"`
	Email         string                `gorm:"type:varchar(128);uniqueIndex;not null;default:'';comment:'邮箱(主账号)'" json:"email"`
	Nickname      string                `gorm:"type:varchar(64);not null;default:'';comment:'昵称'" json:"nickname"`
	Avatar        string                `gorm:"type:varchar(255);not null;default:'';comment:'头像'" json:"avatar"`
	Password      string                `gorm:"type:varchar(255);not null;default:'';comment:'密码(bcrypt)'" json:"-"`
	Salt          string                `gorm:"type:varchar(32);not null;default:'';comment:'加密盐'" json:"-"`
	Phone         string                `gorm:"type:varchar(20);not null;default:'';comment:'手机号'" json:"phone"`
	PhoneCode     string                `gorm:"type:varchar(10);not null;default:'86';comment:'手机号区号'" json:"phoneCode"`
	Status        uint8                 `gorm:"type:tinyint;not null;default:0;comment:'状态 0正常 1禁用'" json:"status"`
	TokenVersion  int64                 `gorm:"type:bigint;not null;default:0;comment:'Token版本号(踢人下线:自增使旧token失效)'" json:"-"`
	LastLoginIp   string                `gorm:"type:varchar(50);not null;default:'';comment:'最后登录IP'" json:"lastLoginIp"`
	LastLoginTime x_null.Time           `gorm:"default:null;comment:'最后登录时间'" json:"lastLoginTime"`
	IsDelete      soft_delete.DeletedAt `gorm:"not null;default:0;softDelete:flag,DeletedAtField:DeleteTime;comment:'是否删除: 0=否, 1=是'" json:"-"`
	CreateTime    x_null.Time           `gorm:"autoCreateTime;not null;comment:'创建时间'" json:"createTime"`
	UpdateTime    x_null.Time           `gorm:"autoUpdateTime;not null;comment:'更新时间'" json:"updateTime"`
	DeleteTime    x_null.Time           `gorm:"default:null;comment:'删除时间'" json:"-"`
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
