package system_schema

import "github.com/adtkcn/x_null"

type SystemLogLoginReq struct {
	Email     string `json:"email" form:"email"`                                // 登录邮箱
	Status    int    `json:"status" form:"status" binding:"omitempty,oneof=1 2"` // 执行状态: [1=成功, 2=失败]
	StartTime string `json:"start_time" form:"start_time" time_format:"2006-01-02"`  // 开始时间
	EndTime   string `json:"end_time" form:"end_time" time_format:"2006-01-02"`    // 结束时间
}

type SystemLoginResp struct {
	Token        string `json:"token"`         // access_token (JWT)
	RefreshToken string `json:"refresh_token"` // refresh_token (JWT)
}

// SystemLoginReq 系统登录参数
type SystemLoginReq struct {
	Email    string `json:"email" binding:"required,min=5,max=200"`   // 邮箱(账号)
	Password string `json:"password" binding:"required,min=6,max=32"` // 密码
}

// SystemLogLoginResp 登录日志返回信息
type SystemLogLoginResp struct {
	ID         string      `json:"id"`          // 主键
	Email      string      `json:"email"`       // 登录邮箱
	Ip         string      `json:"ip"`          // 来源IP
	Os         string      `json:"os"`          // 操作系统
	Browser    string      `json:"browser"`     // 浏览器
	Status     int         `json:"status"`      // 操作状态: [1=成功, 2=失败]
	CreateTime x_null.Time `json:"create_time"` // 创建时间
}

// SystemForgotPwdSendCodeReq 忘记密码-发送验证码请求
type SystemForgotPwdSendCodeReq struct {
	Email string `json:"email" binding:"required,email,min=5,max=200"` // 注册邮箱
}

// SystemForgotPwdResetReq 忘记密码-重置密码请求
type SystemForgotPwdResetReq struct {
	Email    string `json:"email" binding:"required,email,min=5,max=200"` // 注册邮箱
	Code     string `json:"code" binding:"required,len=6"`                // 6位验证码
	Password string `json:"password" binding:"required,min=6,max=32"`     // 新密码(MD5加密后)
}
