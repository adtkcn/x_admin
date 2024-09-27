package monitor_error

import (
	"x_admin/core"
)

// MonitorErrorListReq 监控-错误列列表参数
type MonitorErrorListReq struct {
	ProjectKey      *string // 项目key
	EventType       *string // 事件类型
	Path            *string // URL地址
	Message         *string // 错误消息
	Stack           *string // 错误堆栈
	Md5             *string // md5
	CreateTimeStart *string // 开始创建时间
	CreateTimeEnd   *string // 结束创建时间

}

// MonitorErrorAddReq 监控-错误列新增参数
type MonitorErrorAddReq struct {
	ProjectKey string  // 项目key
	ClientId   string  // sdk生成的客户端id
	EventType  *string // 事件类型
	Path       *string // URL地址
	Message    *string // 错误消息
	Stack      *string // 错误堆栈
}

// MonitorErrorDetailReq 监控-错误列详情参数
type MonitorErrorDetailReq struct {
	Id int // 错误id
}

// MonitorErrorDelReq 监控-错误列删除参数
type MonitorErrorDelReq struct {
	Id int // 错误id
}

// MonitorErrorDelReq 监控-错误列批量删除参数
type MonitorErrorDelBatchReq struct {
	Ids string
}

// MonitorErrorResp 监控-错误列返回信息
type MonitorErrorResp struct {
	Id         int           // 错误id
	ProjectKey string        // 项目key
	EventType  string        // 事件类型
	Path       string        // URL地址
	Message    string        // 错误消息
	Stack      string        // 错误堆栈
	Md5        string        // md5
	CreateTime core.NullTime // 创建时间
}
