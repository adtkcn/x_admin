package login

import (
	"time"
	"x_admin/core"
)

type SystemLogLoginReq struct {
	Username  string    `form:"username"`                             // 登录账号
	Status    int       `form:"status" binding:"omitempty,oneof=1 2"` // 执行状态: [1=成功, 2=失败]
	StartTime time.Time `form:"startTime" time_format:"2006-01-02"`   // 开始时间
	EndTime   time.Time `form:"endTime" time_format:"2006-01-02"`     // 结束时间
}

type SystemLoginResp struct {
	Token string `json:"token"`
}

// SystemLoginReq 系统登录参数
type SystemLoginReq struct {
	Username string `json:"username" binding:"required,min=2,max=20"` // 账号
	Password string `json:"password" binding:"required,min=6,max=32"` // 密码
}

// SystemLogoutReq 登录退出参数
type SystemLogoutReq struct {
	Token string `header:"token" binding:"required"` // 令牌
}

// SystemLogLoginResp 登录日志返回信息
type SystemLogLoginResp struct {
	ID         uint        `json:"id" structs:"id"`                 // 主键
	Username   string      `json:"username" structs:"username"`     // 登录账号
	Ip         string      `json:"ip" structs:"ip"`                 // 来源IP
	Os         string      `json:"os" structs:"os"`                 // 操作系统
	Browser    string      `json:"browser" structs:"browser"`       // 浏览器
	Status     int         `json:"status" structs:"status"`         // 操作状态: [1=成功, 2=失败]
	CreateTime core.TsTime `json:"createTime" structs:"createTime"` // 创建时间
}
