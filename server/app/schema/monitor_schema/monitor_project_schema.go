package monitor_schema

import (
	"mime/multipart"
)

// MonitorProjectListReq 监控-项目列表
type MonitorProjectListReq struct {
	PageNo          *int    `json:"page_no" form:"page_no"`
	PageSize        *int    `json:"page_size" form:"page_size"`
	ProjectKey      *string `json:"project_key" form:"project_key"`
	ProjectName     *string `json:"project_name" form:"project_name"`
	ProjectType     *string `json:"project_type" form:"project_type"`
	Status          *int    `json:"status" form:"status"`
	CreateTimeStart *string `json:"create_time_start" form:"create_time_start"`
	CreateTimeEnd   *string `json:"create_time_end" form:"create_time_end"`
	UpdateTimeStart *string `json:"update_time_start" form:"update_time_start"`
	UpdateTimeEnd   *string `json:"update_time_end" form:"update_time_end"`
}

// MonitorProjectListResp 监控-项目列表返回
type MonitorProjectListResp struct {
	PageNo   int                  `json:"page_no"`
	PageSize int                  `json:"page_size"`
	Count    int64                `json:"count"`
	Lists    []MonitorProjectResp `json:"lists"`
}

// MonitorProjectListAllResp 监控-项目列表全部返回
type MonitorProjectListAllResp struct {
	Lists []MonitorProjectResp `json:"lists"`
}

// MonitorProjectResp 监控-项目
type MonitorProjectResp struct {
	Id          string `json:"id"`           // 主键
	ProjectKey  string `json:"project_key"`  // 项目key
	ProjectName string `json:"project_name"` // 项目名称
	ProjectType string `json:"project_type"` // 项目类型
	Status      int    `json:"status"`       // 状态
	CreateTime  string `json:"create_time"`  // 创建时间
	UpdateTime  string `json:"update_time"`  // 更新时间
	IsDelete    int    `json:"is_delete"`    // 是否删除
}

// MonitorProjectDetailReq 监控-项目详情
type MonitorProjectDetailReq struct {
	Id string `json:"id" form:"id" binding:"required"`
}

// MonitorProjectAddReq 监控-项目新增
type MonitorProjectAddReq struct {
	ProjectKey  string `json:"project_key"`
	ProjectName string `json:"project_name"`
	ProjectType string `json:"project_type"`
	Status      int    `json:"status"`
}

// MonitorProjectEditReq 监控-项目编辑
type MonitorProjectEditReq struct {
	Id          string `json:"id" binding:"required"`
	ProjectKey  string `json:"project_key"`
	ProjectName string `json:"project_name"`
	ProjectType string `json:"project_type"`
	Status      int    `json:"status"`
}

// MonitorProjectDelReq 监控-项目删除
type MonitorProjectDelReq struct {
	Id string `json:"id" form:"id" binding:"required"`
}

// MonitorProjectDelBatchReq 监控-项目删除
type MonitorProjectDelBatchReq struct {
	Ids string `json:"ids" form:"ids" binding:"required"`
}

// MonitorProjectExportReq 监控-项目导出
type MonitorProjectExportReq struct {
	ProjectKey      *string `json:"project_key" form:"project_key"`
	ProjectName     *string `json:"project_name" form:"project_name"`
	ProjectType     *string `json:"project_type" form:"project_type"`
	Status          *int    `json:"status" form:"status"`
	CreateTimeStart *string `json:"create_time_start" form:"create_time_start"`
	CreateTimeEnd   *string `json:"create_time_end" form:"create_time_end"`
	UpdateTimeStart *string `json:"update_time_start" form:"update_time_start"`
	UpdateTimeEnd   *string `json:"update_time_end" form:"update_time_end"`
}

// MonitorProjectImportReq 导入
type MonitorProjectImportReq struct {
	File *multipart.FileHeader `json:"file" form:"file"`
}
