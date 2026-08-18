package user_schema

import "github.com/adtkcn/x_null"

// RefreshTokenReq 刷新token请求
type RefreshTokenReq struct {
	RefreshToken string `json:"refresh_token" binding:"required" label:"refreshToken"`
}

// BindPhoneReq 绑定手机号请求（需要短信验证码）
type BindPhoneReq struct {
	Phone     string `json:"phone"     binding:"required,len=11"   label:"手机号"`
	PhoneCode string `json:"phone_code" binding:"omitempty"         label:"区号"`
	Code      string `json:"code"      binding:"required,len=6"    label:"短信验证码"`
}

// UnbindPhoneReq 解绑手机号请求（需要邮箱验证码，因为手机号可能已注销）
type UnbindPhoneReq struct {
	Email string `json:"email" binding:"required,email" label:"邮箱"`
	Code  string `json:"code"  binding:"required,len=6" label:"邮箱验证码"`
}

// UserAuthItem 绑定项信息
type UserAuthItem struct {
	IdentityType string      `json:"identity_type"`
	Identifier   string      `json:"identifier"`
	CreateTime   x_null.Time `json:"create_time"`
}
