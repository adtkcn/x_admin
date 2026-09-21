package user_schema

// SendCodeReq 发送邮箱验证码请求
type SendCodeReq struct {
	Email string `json:"email" binding:"required,email" label:"邮箱"`
	Scene string `json:"scene" binding:"required,oneof=register reset bind unbind" label:"场景"`
}

// SendSmsCodeReq 发送短信验证码请求
type SendSmsCodeReq struct {
	Phone string `json:"phone" binding:"required,len=11" label:"手机号"`
	Scene string `json:"scene" binding:"required,oneof=sms_bind sms_login sms_reset" label:"场景"`
}

// ResetPasswordReq 邮箱重置密码请求
type ResetPasswordReq struct {
	Email    string `json:"email"    binding:"required,email"          label:"邮箱"`
	Code     string `json:"code"     binding:"required,len=6"          label:"验证码"`
	Password string `json:"password" binding:"required,min=6,max=32"   label:"新密码"`
}

// ResetPhonePasswordReq 手机号重置密码请求
type ResetPhonePasswordReq struct {
	Phone    string `json:"phone"    binding:"required,len=11"         label:"手机号"`
	Code     string `json:"code"     binding:"required,len=6"          label:"验证码"`
	Password string `json:"password" binding:"required,min=6,max=32"   label:"新密码"`
}
