package user_model

import (
	"github.com/adtkcn/x_null"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

// IdentityType 认证类型常量（手机号不在 user_auth 中，直接存 user 表）
const (
	IdentityWechatMini = "wechat_mini" // 微信小程序
	IdentityWechatMp   = "wechat_mp"   // 微信公众号
	IdentityWechatApp  = "wechat_app"  // 微信APP
	IdentityQQ         = "qq"          // QQ
)

// UserAuth 第三方认证绑定表
type UserAuth struct {
	ID           string                `gorm:"primarykey;type:char(36);comment:'UUID v7'" json:"id"`
	UserID       string                `gorm:"type:char(36);not null;index;comment:'用户ID'" json:"userId"`
	IdentityType string                `gorm:"type:varchar(20);not null;comment:'认证类型: phone/wechat_mini/wechat_mp/wechat_app/qq'" json:"identityType"`
	Identifier   string                `gorm:"type:varchar(128);not null;comment:'标识(手机号/openid/unionid/qq_openid)'" json:"identifier"`
	Credential   string                `gorm:"type:varchar(255);not null;default:'';comment:'凭证(加密存储，部分类型可为空)'" json:"-"`
	Extra        string                `gorm:"type:varchar(512);not null;default:'';comment:'扩展信息JSON(微信昵称、头像等)'" json:"extra"`
	IsDelete     soft_delete.DeletedAt `gorm:"not null;default:0;softDelete:flag,DeletedAtField:DeleteTime;comment:'是否删除: 0=否, 1=是'" json:"-"`
	CreateTime   x_null.Time           `gorm:"autoCreateTime;not null;comment:'创建时间'" json:"createTime"`
	UpdateTime   x_null.Time           `gorm:"autoUpdateTime;not null;comment:'更新时间'" json:"updateTime"`
	DeleteTime   x_null.Time           `gorm:"default:null;comment:'删除时间'" json:"-"`
}

// BeforeCreate 在创建前生成UUIDv7
func (m *UserAuth) BeforeCreate(tx *gorm.DB) (err error) {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	m.ID = id.String()
	return nil
}
