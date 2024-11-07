package monitor_slow

import (
	"x_admin/core"
)

// MonitorSlowListReq 监控-错误列列表参数
type MonitorSlowListReq struct {
	ProjectKey      *string  // 项目key
	ClientId        *string  // sdk生成的客户端id
	UserId          *string  // 用户id
	Path            *string  // URL地址
	Time            *float64 // 时间
	CreateTimeStart *string  // 开始创建时间
	CreateTimeEnd   *string  // 结束创建时间
}

// MonitorSlowAddReq 监控-错误列新增参数
type MonitorSlowAddReq struct {
	ProjectKey *string        // 项目key
	ClientId   *string        // sdk生成的客户端id
	UserId     *string        // 用户id
	Path       *string        // URL地址
	Time       core.NullFloat // 时间
}

// MonitorSlowEditReq 监控-错误列编辑参数
type MonitorSlowEditReq struct {
	Id         int            // 错误id
	ProjectKey *string        // 项目key
	ClientId   *string        // sdk生成的客户端id
	UserId     *string        // 用户id
	Path       *string        // URL地址
	Time       core.NullFloat // 时间
}

// MonitorSlowDetailReq 监控-错误列详情参数
type MonitorSlowDetailReq struct {
	Id int // 错误id
}

// MonitorSlowDelReq 监控-错误列删除参数
type MonitorSlowDelReq struct {
	Id int // 错误id
}

// MonitorSlowDelReq 监控-错误列批量删除参数
type MonitorSlowDelBatchReq struct {
	Ids string
}

// MonitorSlowResp 监控-错误列返回信息
type MonitorSlowResp struct {
	Id         int            // 错误id
	ProjectKey string         // 项目key
	ClientId   string         // sdk生成的客户端id
	UserId     string         // 用户id
	Path       string         // URL地址
	Time       core.NullFloat // 时间
	CreateTime core.NullTime  // 创建时间
}
