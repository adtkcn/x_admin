package user_schema

// LoginReq 邮箱登录请求
type LoginReq struct {
	Email    string `json:"email"    binding:"required,email"        label:"邮箱"`
	Password string `json:"password" binding:"required,min=6,max=32" label:"密码"`
}

// PhoneLoginReq 手机号+密码登录请求
type PhoneLoginReq struct {
	Phone     string `json:"phone"     binding:"required,len=11"        label:"手机号"`
	PhoneCode string `json:"phoneCode" binding:"omitempty"              label:"区号"`
	Password  string `json:"password"  binding:"required,min=6,max=32"  label:"密码"`
}

// PhoneCodeLoginReq 手机号+短信验证码登录请求
type PhoneCodeLoginReq struct {
	Phone     string `json:"phone"     binding:"required,len=11"  label:"手机号"`
	PhoneCode string `json:"phoneCode" binding:"omitempty"        label:"区号"`
	Code      string `json:"code"      binding:"required,len=6"   label:"验证码"`
}

// LoginResp 登录响应（JWT token对）
type LoginResp struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    int    `json:"expiresIn"` // access_token 有效期(秒)
	IsNew        bool   `json:"isNew"`     // 是否新注册用户（微信登录时自动注册）
}
