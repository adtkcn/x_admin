package monitor_error_list

import (
	"x_admin/core"
)

// MonitorErrorListAddReq 错误对应的用户记录新增参数
type MonitorErrorListAddReq struct {
	ErrorId    string // 错误id
	ClientId   string // 客户端id
	ProjectKey string // 项目id
}

// MonitorErrorListResp 错误对应的用户记录返回信息
type MonitorErrorListResp struct {
	Id         int           // 项目id
	ErrorId    string        // 错误id
	ClientId   string        // 客户端id
	ProjectKey string        // 项目id
	CreateTime core.NullTime // 创建时间
}