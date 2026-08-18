package user_schema

// LoginReq 邮箱登录请求
type LoginReq struct {
	Email    string `json:"email"    binding:"required,email"        label:"邮箱"`
	Password string `json:"password" binding:"required,min=6,max=32" label:"密码"`
}

// PhoneLoginReq 手机号+密码登录请求
type PhoneLoginReq struct {
	Phone     string `json:"phone"     binding:"required,len=11"        label:"手机号"`
	PhoneCode string `json:"phone_code" binding:"omitempty"              label:"区号"`
	Password  string `json:"password"  binding:"required,min=6,max=32"  label:"密码"`
}

// PhoneCodeLoginReq 手机号+短信验证码登录请求
type PhoneCodeLoginReq struct {
	Phone     string `json:"phone"     binding:"required,len=11"  label:"手机号"`
	PhoneCode string `json:"phone_code" binding:"omitempty"        label:"区号"`
	Code      string `json:"code"      binding:"required,len=6"   label:"验证码"`
}

// LoginResp 登录响应（JWT token对）
type LoginResp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"` // access_token 有效期(秒)
	IsNew        bool   `json:"is_new"`     // 是否新注册用户（微信登录时自动注册）
	UserID       string `json:"user_id"`    // 用户ID（小程序登录后直接返回，避免二次请求）
	Nickname     string `json:"nickname"`   // 用户昵称
}
