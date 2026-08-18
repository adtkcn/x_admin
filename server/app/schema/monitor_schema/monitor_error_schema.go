package monitor_schema

import (
	"mime/multipart"
)

// MonitorErrorListReq 监控错误列表
type MonitorErrorListReq struct {
	PageNo          *int    `json:"page_no" form:"page_no"`
	PageSize        *int    `json:"page_size" form:"page_size"`
	ProjectKey      *string `json:"project_key" form:"project_key"`
	EventType       *string `json:"event_type" form:"event_type"`
	Path            *string `json:"path" form:"path"`
	Message         *string `json:"message" form:"message"`
	Stack           *string `json:"stack" form:"stack"`
	Md5             *string `json:"md5" form:"md5"`
	CreateTimeStart *string `json:"create_time_start" form:"create_time_start"`
	CreateTimeEnd   *string `json:"create_time_end" form:"create_time_end"`
}

// MonitorErrorListResp 监控错误列表返回
type MonitorErrorListResp struct {
	PageNo   int                `json:"page_no"`
	PageSize int                `json:"page_size"`
	Count    int64              `json:"count"`
	Lists    []MonitorErrorResp `json:"lists"`
}

// MonitorErrorListAllResp 监控错误列表返回
type MonitorErrorListAllResp struct {
	Lists []MonitorErrorResp `json:"lists"`
}

// MonitorErrorResp 监控错误
type MonitorErrorResp struct {
	Id         string `json:"id"`          // 主键
	ProjectKey string `json:"project_key"` // 项目key
	EventType  string `json:"event_type"`  // 事件类型
	Path       string `json:"path"`        // URL地址
	Message    string `json:"message"`     // 错误消息
	Stack      string `json:"stack"`       // 错误堆栈
	Md5        string `json:"md5"`         // md5值
	CreateTime string `json:"create_time"` // 创建时间
	UpdateTime string `json:"update_time"` // 更新时间
	UserId     string `json:"user_id"`     // 用户id
	ClientId   string `json:"client_id"`   // 客户端id
	IsDelete   int    `json:"is_delete"`   // 是否删除
}

// MonitorErrorDetailReq 监控错误详情
type MonitorErrorDetailReq struct {
	Id string `json:"id" form:"id" v:"required#错误id不能为空"`
}

// MonitorErrorAddReq 监控错误新增
type MonitorErrorAddReq struct {
	ProjectKey string `json:"project_key"`
	EventType  string `json:"event_type"`
	Path       string `json:"path"`
	Message    string `json:"message"`
	Stack      string `json:"stack"`
	Md5        string `json:"md5"`
	ClientId   string `json:"client_id"`
	UserId     string `json:"user_id"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
}

// MonitorErrorEditReq 监控错误编辑
type MonitorErrorEditReq struct {
	Id         string `json:"id" v:"required#错误id不能为空"`
	ProjectKey string `json:"project_key"`
	EventType  string `json:"event_type"`
	Path       string `json:"path"`
	Message    string `json:"message"`
	Stack      string `json:"stack"`
	Md5        string `json:"md5"`
	ClientId   string `json:"client_id"`
	UserId     string `json:"user_id"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
}

// MonitorErrorDelReq 监控错误删除
type MonitorErrorDelReq struct {
	Id string `json:"id" form:"id" v:"required#错误id不能为空"`
}

// MonitorErrorDelBatchReq 监控错误删除
type MonitorErrorDelBatchReq struct {
	Ids string `json:"ids" form:"ids" v:"required#错误ids不能为空"`
}

// MonitorErrorExportReq 监控错误导出
type MonitorErrorExportReq struct {
	ProjectKey      *string `json:"project_key" form:"project_key"`
	EventType       *string `json:"event_type" form:"event_type"`
	Path            *string `json:"path" form:"path"`
	Message         *string `json:"message" form:"message"`
	Stack           *string `json:"stack" form:"stack"`
	Md5             *string `json:"md5" form:"md5"`
	CreateTimeStart *string `json:"create_time_start" form:"create_time_start"`
	CreateTimeEnd   *string `json:"create_time_end" form:"create_time_end"`
}

// MonitorErrorImportReq 导入
type MonitorErrorImportReq struct {
	File *multipart.FileHeader `json:"file" form:"file"`
}
