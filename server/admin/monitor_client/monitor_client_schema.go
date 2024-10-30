package monitor_client

import (
	"x_admin/core"
)

// MonitorClientListReq 监控-客户端信息列表参数
type MonitorClientListReq struct {
	ProjectKey *string // 项目key
	ClientId   *string // sdk生成的客户端id
	UserId     *string // 用户id
	Os         *string // 系统
	Browser    *string // 浏览器

	Country  *string // 国家
	Province *string // 省份
	City     *string // 城市
	Operator *string // 电信运营商
	Ip       *string // ip

	Width           *int    // 屏幕
	Height          *int    // 屏幕高度
	Ua              *string // ua记录
	CreateTimeStart *string // 开始创建时间
	CreateTimeEnd   *string // 结束创建时间

}

// MonitorClientAddReq 监控-客户端信息新增参数
type MonitorClientAddReq struct {
	ProjectKey *string      // 项目key
	ClientId   *string      // sdk生成的客户端id
	UserId     *string      // 用户id
	Os         *string      // 系统
	Browser    *string      // 浏览器
	Country    *string      // 国家
	Province   *string      // 省份
	City       *string      // 城市
	Operator   *string      // 电信运营商
	Ip         *string      // ip
	Width      core.NullInt // 屏幕
	Height     core.NullInt // 屏幕高度
	Ua         *string      // ua记录

}

// MonitorClientEditReq 监控-客户端信息编辑参数
type MonitorClientEditReq struct {
	Id         int          // uuid
	ProjectKey *string      // 项目key
	ClientId   *string      // sdk生成的客户端id
	UserId     *string      // 用户id
	Os         *string      // 系统
	Browser    *string      // 浏览器
	Country    *string      // 国家
	Province   *string      // 省份
	City       *string      // 城市
	Operator   *string      // 电信运营商
	Ip         *string      // ip
	Width      core.NullInt // 屏幕
	Height     core.NullInt // 屏幕高度
	Ua         *string      // ua记录
}

// MonitorClientDetailReq 监控-客户端信息详情参数
type MonitorClientDetailReq struct {
	Id int // uuid
}

// MonitorClientDelReq 监控-客户端信息删除参数
type MonitorClientDelReq struct {
	Id int // uuid
}

// MonitorClientDelReq 监控-客户端信息批量删除参数
type MonitorClientDelBatchReq struct {
	Ids string
}

// MonitorClientResp 监控-客户端信息返回信息
type MonitorClientResp struct {
	Id         int           // uuid
	ProjectKey string        // 项目key
	ClientId   string        // sdk生成的客户端id
	UserId     string        // 用户id
	Os         string        // 系统
	Browser    string        // 浏览器
	Country    string        // 国家
	Province   string        // 省份
	City       string        // 城市
	Operator   string        // 电信运营商
	Ip         string        // ip
	Width      core.NullInt  // 屏幕
	Height     core.NullInt  // 屏幕高度
	Ua         string        // ua记录
	CreateTime core.NullTime // 创建时间
}
