package system_schema

import "github.com/adtkcn/x_null"

// //SystemLogOperateReq 操作日志列表参数
type SystemLogOperateReq struct {
	Title     string `json:"title" form:"title"`                                       // 操作标题
	Email     string `json:"email" form:"email"`                                       // 邮箱(账号)
	Ip        string `json:"ip" form:"ip"`                                          // 请求IP
	Type      string `json:"type" form:"type" binding:"omitempty,oneof=GET POST PUT"` // 请求类型: GET/POST/PUT
	Status    int    `json:"status" form:"status" binding:"omitempty,oneof=1 2"`        // 执行状态: [1=成功, 2=失败]
	Url       string `json:"url" form:"url"`                                         // 请求地址
	StartTime string `json:"start_time" form:"start_time" time_format:"2006-01-02"`         // 开始时间
	EndTime   string `json:"end_time" form:"end_time" time_format:"2006-01-02"`           // 结束时间
}

// type SystemLogLoginReq struct {
// 	Username  string `form:"username"`                             // 登录账号
// 	Status    int    `form:"status" binding:"omitempty,oneof=1 2"` // 执行状态: [1=成功, 2=失败]
// 	StartTime string `form:"startTime" time_format:"2006-01-02"`   // 开始时间
// 	EndTime   string `form:"endTime" time_format:"2006-01-02"`     // 结束时间
// }

// SystemLogOperateResp 操作日志返回信息
type SystemLogOperateResp struct {
	ID         string      `json:"id"`          // 主键
	Email      string      `json:"email"`       // 邮箱(账号)
	Nickname   string      `json:"nickname"`    // 用户昵称
	Type       string      `json:"type"`        // 请求类型: GET/POST/PUT
	Title      string      `json:"title"`       // 操作标题
	Method     string      `json:"method"`      // 请求方式
	Ip         string      `json:"ip"`          // 请求IP
	Url        string      `json:"url"`         // 请求地址
	Args       string      `json:"args"`        // 请求参数
	Error      string      `json:"error"`       // 错误信息
	Status     int         `json:"status"`      // 执行状态: [1=成功, 2=失败]
	TaskTime   string      `json:"task_time"`   // 执行耗时
	StartTime  x_null.Time `json:"start_time"`  // 开始时间
	EndTime    x_null.Time `json:"end_time"`    // 结束时间
	CreateTime x_null.Time `json:"create_time"` // 创建时间
}
