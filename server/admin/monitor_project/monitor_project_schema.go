package monitor_project
import (
	"x_admin/core"

)

//MonitorProjectListReq 监控项目列表参数
type MonitorProjectListReq struct {
            ProjectKey *string // 项目uuid
            ProjectName *string // 项目名称
            ProjectType *string // 项目类型go java web node php 等
            Status *int // 是否启用: 0=否, 1=是
            CreateTimeStart *string // 开始创建时间
            CreateTimeEnd *string // 结束创建时间
            UpdateTimeStart *string // 开始更新时间
            UpdateTimeEnd *string // 结束更新时间
}



//MonitorProjectAddReq 监控项目新增参数
type MonitorProjectAddReq struct {
    ProjectKey  *string  // 项目uuid
    ProjectName  *string  // 项目名称
    ProjectType  *string  // 项目类型go java web node php 等
    Status  core.NullInt  // 是否启用: 0=否, 1=是
}

//MonitorProjectEditReq 监控项目编辑参数
type MonitorProjectEditReq struct {
        Id int // 项目id
        ProjectKey  *string  // 项目uuid
        ProjectName  *string  // 项目名称
        ProjectType  *string  // 项目类型go java web node php 等
        Status  core.NullInt  // 是否启用: 0=否, 1=是
}

//MonitorProjectDetailReq 监控项目详情参数
type MonitorProjectDetailReq struct {
    Id int // 项目id
}

//MonitorProjectDelReq 监控项目删除参数
type MonitorProjectDelReq struct {
    Id int // 项目id
}

//MonitorProjectDelReq 监控项目批量删除参数
type MonitorProjectDelBatchReq struct {
	Ids string
}

//MonitorProjectResp 监控项目返回信息
type MonitorProjectResp struct {
        Id int // 项目id
        ProjectKey string // 项目uuid
        ProjectName string // 项目名称
        ProjectType string // 项目类型go java web node php 等
        Status core.NullInt // 是否启用: 0=否, 1=是
        CreateTime core.NullTime // 创建时间
        UpdateTime core.NullTime // 更新时间
}
