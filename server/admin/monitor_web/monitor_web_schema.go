package monitor_web

import "x_admin/core"

//MonitorWebListReq 错误收集error列表参数
type MonitorWebListReq struct {
	ProjectKey      string `form:"projectKey"`      // 项目key
	ClientId        string `form:"clientId"`        // sdk生成的客户端id
	EventType       string `form:"eventType"`       // 事件类型
	Page            string `form:"page"`            // URL地址
	Message         string `form:"message"`         // 错误消息
	Stack           string `form:"stack"`           // 错误堆栈
	ClientTimeStart string `form:"clientTimeStart"` // 开始客户端时间
	ClientTimeEnd   string `form:"clientTimeEnd"`   // 结束客户端时间
	CreateTimeStart string `form:"createTimeStart"` // 开始创建时间
	CreateTimeEnd   string `form:"createTimeEnd"`   // 结束创建时间
}

//MonitorWebDetailReq 错误收集error详情参数
type MonitorWebDetailReq struct {
	Id int `form:"id"` // uuid
}

//MonitorWebAddReq 错误收集error新增参数
type MonitorWebAddReq struct {
	ProjectKey string      `form:"projectKey"` // 项目key
	ClientId   string      `form:"clientId"`   // sdk生成的客户端id
	EventType  string      `form:"eventType"`  // 事件类型
	Page       string      `form:"page"`       // URL地址
	Message    string      `form:"message"`    // 错误消息
	Stack      string      `form:"stack"`      // 错误堆栈
	ClientTime core.TsTime `form:"clientTime"` // 客户端时间
}

//MonitorWebEditReq 错误收集error编辑参数
type MonitorWebEditReq struct {
	Id         int         `form:"id"`         // uuid
	ProjectKey string      `form:"projectKey"` // 项目key
	ClientId   string      `form:"clientId"`   // sdk生成的客户端id
	EventType  string      `form:"eventType"`  // 事件类型
	Page       string      `form:"page"`       // URL地址
	Message    string      `form:"message"`    // 错误消息
	Stack      string      `form:"stack"`      // 错误堆栈
	ClientTime core.TsTime `form:"clientTime"` // 客户端时间
}

//MonitorWebDelReq 错误收集error新增参数
type MonitorWebDelReq struct {
	Id int `form:"id"` // uuid
}

//MonitorWebResp 错误收集error返回信息
type MonitorWebResp struct {
	Id         int         `json:"id" structs:"id"`                                       // uuid
	ProjectKey string      `json:"projectKey" structs:"projectKey" excel:"name:项目key;"`   // 项目key
	ClientId   string      `json:"clientId" structs:"clientId" excel:"name:sdk生成的客户端id;"` // sdk生成的客户端id
	EventType  string      `json:"eventType" structs:"eventType" excel:"name:事件类型;"`      // 事件类型
	Page       string      `json:"page" structs:"page" excel:"name:URL地址;"`               // URL地址
	Message    string      `json:"message" structs:"message" excel:"name:错误消息;"`          // 错误消息
	Stack      string      `json:"stack" structs:"stack" excel:"name:错误堆栈;"`              // 错误堆栈
	ClientTime core.TsTime `json:"clientTime" structs:"clientTime" excel:"name:客户端时间;"`   // 客户端时间
	CreateTime core.TsTime `json:"createTime" structs:"createTime" excel:"name:创建时间;"`    // 创建时间
}
