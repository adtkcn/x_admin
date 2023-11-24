package log

import (
	"time"
	"x_admin/core"
)

// //SystemLogOperateReq 操作日志列表参数
type SystemLogOperateReq struct {
	Title     string    `form:"title"`                                       // 操作标题
	Username  string    `form:"username"`                                    // 用户账号
	Ip        string    `form:"ip"`                                          // 请求IP
	Type      string    `form:"type" binding:"omitempty,oneof=GET POST PUT"` // 请求类型: GET/POST/PUT
	Status    int       `form:"status" binding:"omitempty,oneof=1 2"`        // 执行状态: [1=成功, 2=失败]
	Url       string    `form:"url"`                                         // 请求地址
	StartTime time.Time `form:"startTime" time_format:"2006-01-02"`          // 开始时间
	EndTime   time.Time `form:"endTime" time_format:"2006-01-02"`            // 结束时间
}
type SystemLogLoginReq struct {
	Username  string    `form:"username"`                             // 登录账号
	Status    int       `form:"status" binding:"omitempty,oneof=1 2"` // 执行状态: [1=成功, 2=失败]
	StartTime time.Time `form:"startTime" time_format:"2006-01-02"`   // 开始时间
	EndTime   time.Time `form:"endTime" time_format:"2006-01-02"`     // 结束时间
}

// SystemLogOperateResp 操作日志返回信息
type SystemLogOperateResp struct {
	ID         uint        `json:"id" structs:"id"`                 // 主键
	Username   string      `json:"username" structs:"username"`     // 用户账号
	Nickname   string      `json:"nickname" structs:"nickname"`     // 用户昵称
	Type       string      `json:"type" structs:"type"`             // 请求类型: GET/POST/PUT
	Title      string      `json:"title" structs:"title"`           // 操作标题
	Method     string      `json:"method" structs:"method"`         // 请求方式
	Ip         string      `json:"ip" structs:"ip"`                 // 请求IP
	Url        string      `json:"url" structs:"url"`               // 请求地址
	Args       string      `json:"args" structs:"args"`             // 请求参数
	Error      string      `json:"error" structs:"error"`           // 错误信息
	Status     int         `json:"status" structs:"status"`         // 执行状态: [1=成功, 2=失败]
	TaskTime   string      `json:"taskTime" structs:"taskTime"`     // 执行耗时
	StartTime  core.TsTime `json:"startTime" structs:"startTime"`   // 开始时间
	EndTime    core.TsTime `json:"endTime" structs:"endTime"`       // 结束时间
	CreateTime core.TsTime `json:"createTime" structs:"createTime"` // 创建时间
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
