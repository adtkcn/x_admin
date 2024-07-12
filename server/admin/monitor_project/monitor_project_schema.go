package monitor_project

import "x_admin/core"

//MonitorProjectListReq 错误项目列表参数
type MonitorProjectListReq struct {
	ProjectKey      string `form:"projectKey"`      // 项目uuid
	ProjectName     string `form:"projectName"`     // 项目名称
	ProjectType     string `form:"projectType"`     // 项目类型go java web node php 等
	CreateTimeStart string `form:"createTimeStart"` // 开始创建时间
	CreateTimeEnd   string `form:"createTimeEnd"`   // 结束创建时间
	UpdateTimeStart string `form:"updateTimeStart"` // 开始更新时间
	UpdateTimeEnd   string `form:"updateTimeEnd"`   // 结束更新时间
}

//MonitorProjectDetailReq 错误项目详情参数
type MonitorProjectDetailReq struct {
	Id int `form:"id"` // 项目id
}

//MonitorProjectAddReq 错误项目新增参数
type MonitorProjectAddReq struct {
	// ProjectKey  string `form:"projectKey"`  // 项目uuid
	ProjectName string `form:"projectName"` // 项目名称
	ProjectType string `form:"projectType"` // 项目类型go java web node php 等
}

//MonitorProjectEditReq 错误项目编辑参数
type MonitorProjectEditReq struct {
	Id int `form:"id"` // 项目id
	// ProjectKey  string `form:"projectKey"`  // 项目uuid
	ProjectName *string `form:"projectName"` // 项目名称
	ProjectType *string `form:"projectType"` // 项目类型go java web node php 等
}

//MonitorProjectDelReq 错误项目新增参数
type MonitorProjectDelReq struct {
	Id int `form:"id"` // 项目id
}

//MonitorProjectResp 错误项目返回信息
type MonitorProjectResp struct {
	Id          int         `json:"id" structs:"id" excel:"name:项目id;"`                                         // 项目id
	ProjectKey  string      `json:"projectKey" structs:"projectKey" excel:"name:项目uuid;"`                       // 项目uuid
	ProjectName string      `json:"projectName" structs:"projectName" excel:"name:项目名称;"`                       // 项目名称
	ProjectType string      `json:"projectType" structs:"projectType" excel:"name:项目类型go java web node php 等;"` // 项目类型go java web node php 等
	UpdateTime  core.TsTime `json:"updateTime" structs:"updateTime" excel:"name:更新时间;"`                         // 更新时间
	CreateTime  core.TsTime `json:"createTime" structs:"createTime" excel:"name:创建时间;"`                         // 创建时间
}
