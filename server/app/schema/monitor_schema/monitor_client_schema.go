package monitor_schema

import (
	"mime/multipart"
)

// MonitorClientListReq 监控-客户端信息列表
type MonitorClientListReq struct {
	PageNo          *int    `json:"page_no" form:"page_no"`         // 页码
	PageSize        *int    `json:"page_size" form:"page_size"`     // 每页数量
	ProjectKey      *string `json:"project_key" form:"project_key"`
	ClientId        *string `json:"client_id" form:"client_id"`
	UserId          *string `json:"user_id" form:"user_id"`
	Os              *string `json:"os" form:"os"`
	Browser         *string `json:"browser" form:"browser"`
	Country         *string `json:"country" form:"country"`
	Province        *string `json:"province" form:"province"`
	City            *string `json:"city" form:"city"`
	Operator        *string `json:"operator" form:"operator"`
	Ip              *string `json:"ip" form:"ip"`
	Ua              *string `json:"ua" form:"ua"`
	CreateTimeStart *string `json:"create_time_start" form:"create_time_start"`
	CreateTimeEnd   *string `json:"create_time_end" form:"create_time_end"`
}

// MonitorClientListResp 监控-客户端信息列表返回
type MonitorClientListResp struct {
	PageNo   int                 `json:"page_no"`   // 页码
	PageSize int                 `json:"page_size"` // 每页数量
	Count    int64               `json:"count"`     // 总数
	Lists    []MonitorClientResp `json:"lists"`     // 列表
}

// MonitorClientListAllResp 监控-客户端信息列表全部返回
type MonitorClientListAllResp struct {
	Lists []MonitorClientResp `json:"lists"`
}

// MonitorClientResp 监控-客户端信息
type MonitorClientResp struct {
	Id         string `json:"id"`          // 主键
	ProjectKey string `json:"project_key"` // 项目key
	ClientId   string `json:"client_id"`   // 客户端id
	UserId     string `json:"user_id"`     // 用户id
	Os         string `json:"os"`          // 操作系统
	Browser    string `json:"browser"`     // 浏览器
	Country    string `json:"country"`     // 国家
	Province   string `json:"province"`    // 省份
	City       string `json:"city"`        // 城市
	Operator   string `json:"operator"`    // 运营商
	Ip         string `json:"ip"`          // IP地址
	Ua         string `json:"ua"`          // 用户代理
	CreateTime string `json:"create_time"` // 创建时间
	UpdateTime string `json:"update_time"` // 更新时间
	Width      int    `json:"width"`       // 屏幕宽度
	Height     int    `json:"height"`      // 屏幕高度
	IsDelete   int    `json:"is_delete"`   // 是否删除
}

// MonitorClientDetailReq 监控-客户端信息详情
type MonitorClientDetailReq struct {
	Id string `json:"id" form:"id" v:"required#客户端id不能为空"`
}

// MonitorClientAddReq 监控-客户端信息新增
type MonitorClientAddReq struct {
	ProjectKey string `json:"project_key"`
	ClientId   string `json:"client_id"`
	UserId     string `json:"user_id"`
	Os         string `json:"os"`
	Browser    string `json:"browser"`
	Ip         string `json:"ip"`
	Ua         string `json:"ua"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
}

// MonitorClientEditReq 监控-客户端信息编辑
type MonitorClientEditReq struct {
	Id         string `json:"id" v:"required#客户端id不能为空"`
	ProjectKey string `json:"project_key"`
	ClientId   string `json:"client_id"`
	UserId     string `json:"user_id"`
	Os         string `json:"os"`
	Browser    string `json:"browser"`
	Ip         string `json:"ip"`
	Ua         string `json:"ua"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
}

// MonitorClientDelReq 监控-客户端信息删除
type MonitorClientDelReq struct {
	Id string `json:"id" form:"id" v:"required#客户端id不能为空"`
}

// MonitorClientDelBatchReq 监控-客户端信息删除
type MonitorClientDelBatchReq struct {
	Ids string `json:"ids" form:"ids" v:"required#客户端ids不能为空"`
}

// MonitorClientErrorUsersReq 监控-客户端信息错误用户列表
type MonitorClientErrorUsersReq struct {
	Id string `json:"id" form:"id" v:"required#客户端id不能为空"`
}

// MonitorClientExportReq 监控-客户端信息导出
type MonitorClientExportReq struct {
	ProjectKey      *string `json:"project_key" form:"project_key"`
	ClientId        *string `json:"client_id" form:"client_id"`
	Os              *string `json:"os" form:"os"`
	Browser         *string `json:"browser" form:"browser"`
	Ua              *string `json:"ua" form:"ua"`
	CreateTimeStart *string `json:"create_time_start" form:"create_time_start"`
	CreateTimeEnd   *string `json:"create_time_end" form:"create_time_end"`
}

// MonitorClientImportReq 导入
type MonitorClientImportReq struct {
	File *multipart.FileHeader `json:"file" form:"file"`
}

