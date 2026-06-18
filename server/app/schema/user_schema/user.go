package user_schema

import "github.com/adtkcn/x_null"

// UserInfoResp 用户信息响应
type UserInfoResp struct {
	ID            string      `json:"id"`
	Email         string      `json:"email"`
	Nickname      string      `json:"nickname"`
	Avatar        string      `json:"avatar"`
	Phone         string      `json:"phone"`
	PhoneCode     string      `json:"phoneCode"`
	Status        uint8       `json:"status"`
	LastLoginIp   string      `json:"lastLoginIp"`
	LastLoginTime x_null.Time `json:"lastLoginTime"`
	CreateTime    x_null.Time `json:"createTime"`
}

// UpdateUserReq 更新用户信息请求
type UpdateUserReq struct {
	Nickname string `json:"nickname" binding:"omitempty,min=2,max=20" label:"昵称"`
	Avatar   string `json:"avatar"   binding:"omitempty"              label:"头像"`
}
