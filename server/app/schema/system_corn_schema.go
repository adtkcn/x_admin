package schema
import (
	"x_admin/core"
)

type SystemCornPrimarykey struct {
        Id string // taskid
}

//SystemCornListReq 定时任务列表参数
type SystemCornListReq struct {
            TaskName core.NullString // 任务名称
            TaskCode core.NullString // 任务编码
            CornExpr core.NullString // corn表达式
            Disabled core.NullInt // 禁用
            CreatedBy core.NullString // 创建人
            CreateTimeStart core.NullString // 开始创建时间
            CreateTimeEnd core.NullString // 结束创建时间
            UpdateTimeStart core.NullString // 开始更新时间
            UpdateTimeEnd core.NullString // 结束更新时间
}



//SystemCornAddReq 定时任务新增参数
type SystemCornAddReq struct {
    TaskName  core.NullString  // 任务名称
    TaskCode  core.NullString  // 任务编码
    CornExpr  core.NullString  // corn表达式
    Disabled  core.NullInt  // 禁用
}

//SystemCornEditReq 定时任务编辑参数
type SystemCornEditReq struct {
    Id string // taskid
    TaskName  core.NullString  // 任务名称
    TaskCode  core.NullString  // 任务编码
    CornExpr  core.NullString  // corn表达式
    Disabled  core.NullInt  // 禁用
}

 

//SystemCornDelBatchReq 定时任务批量删除参数
type SystemCornDelBatchReq struct {
	Ids string
}

//SystemCornResp 定时任务返回信息
type SystemCornResp struct {
        Id string `swaggertype:"string"`// taskid
        TaskName core.NullString `swaggertype:"string"`// 任务名称
        TaskCode core.NullString `swaggertype:"string"`// 任务编码
        CornExpr core.NullString `swaggertype:"string"`// corn表达式
        Disabled core.NullInt `swaggertype:"number"`// 禁用
        CreatedBy core.NullString `swaggertype:"string"`// 创建人
        CreateTime core.NullTime `swaggertype:"string"`// 创建时间
        UpdateTime core.NullTime `swaggertype:"string"`// 更新时间
}
