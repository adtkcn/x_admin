package user_schema

import "github.com/adtkcn/x_null"

// UserInfoResp 用户信息响应
type UserInfoResp struct {
	ID            string      `json:"id"`
	Email         string      `json:"email"`
	Nickname      string      `json:"nickname"`
	Avatar        string      `json:"avatar"`
	Phone         string      `json:"phone"`
	PhoneCode     string      `json:"phone_code"`
	Status        uint8       `json:"status"`
	LastLoginIp   string      `json:"last_login_ip"`
	LastLoginTime x_null.Time `json:"last_login_time"`
	CreateTime    x_null.Time `json:"create_time"`
}

// UpdateUserReq 更新用户信息请求
type UpdateUserReq struct {
	Nickname string `json:"nickname" binding:"omitempty,min=2,max=20" label:"昵称"`
	Avatar   string `json:"avatar"   binding:"omitempty"              label:"头像"`
}

// ChangePasswordReq 修改密码请求
type ChangePasswordReq struct {
	OldPassword string `json:"old_password" binding:"required"            label:"原密码"`
	NewPassword string `json:"new_password" binding:"required,min=6,max=32" label:"新密码"`
}
