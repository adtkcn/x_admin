package user_schema

// RegisterReq 注册请求
type RegisterReq struct {
	Email    string `json:"email"    binding:"required,email"                    label:"邮箱"`
	Password string `json:"password" binding:"required,min=6,max=32"             label:"密码"`
	Code     string `json:"code"     binding:"required,len=6"                    label:"验证码"`
	Nickname string `json:"nickname" binding:"omitempty,min=2,max=20"            label:"昵称"`
}
