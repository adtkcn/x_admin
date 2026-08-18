package schema

import (
	"x_admin/app/schema/system_schema"

	"github.com/adtkcn/x_null"
)

type SystemCornPrimarykey struct {
	Id string `json:"id" form:"id"`
}

// SystemCornListReq 定时任务列表参数
type SystemCornListReq struct {
	TaskName        x_null.String `json:"task_name" form:"task_name"`             // 任务名称
	TaskCode        x_null.String `json:"task_code" form:"task_code"`             // 任务编码
	CornExpr        x_null.String `json:"corn_expr" form:"corn_expr"`             // corn表达式
	Status          x_null.Int64  `json:"status" form:"status"`                   // 状态
	CreatedBy       x_null.String `json:"created_by" form:"created_by"`           // 创建人
	Nickname        x_null.String `json:"nickname" form:"nickname"`               // 创建人名称
	CreateTimeStart x_null.String `json:"create_time_start" form:"create_time_start"` // 开始创建时间
	CreateTimeEnd   x_null.String `json:"create_time_end" form:"create_time_end"`     // 结束创建时间
	UpdateTimeStart x_null.String `json:"update_time_start" form:"update_time_start"` // 开始更新时间
	UpdateTimeEnd   x_null.String `json:"update_time_end" form:"update_time_end"`     // 结束更新时间
}

// SystemCornAddReq 定时任务新增参数
type SystemCornAddReq struct {
	TaskName x_null.String `json:"task_name" form:"task_name"` // 任务名称
	TaskCode x_null.String `json:"task_code" form:"task_code"` // 任务编码
	CornExpr x_null.String `json:"corn_expr" form:"corn_expr"` // corn表达式
	Status   x_null.Int64  `json:"status" form:"status"`       // 状态
}

// SystemCornEditReq 定时任务编辑参数
type SystemCornEditReq struct {
	Id       string        `json:"id" form:"id"`
	TaskName x_null.String `json:"task_name" form:"task_name"` // 任务名称
	TaskCode x_null.String `json:"task_code" form:"task_code"` // 任务编码
	CornExpr x_null.String `json:"corn_expr" form:"corn_expr"` // corn表达式
	Status   x_null.Int64  `json:"status" form:"status"`       // 状态
}

// SystemCornDelBatchReq 定时任务批量删除参数
type SystemCornDelBatchReq struct {
	Ids string `json:"ids" form:"ids"`
}

// SystemCornResp 定时任务返回信息
type SystemCornResp struct {
	Id            string                                  `json:"id" swaggertype:"string"` // ID
	TaskName      x_null.String                           `json:"task_name" swaggertype:"string"` // 任务名称
	TaskCode      x_null.String                           `json:"task_code" swaggertype:"string"` // 任务编码
	CornExpr      x_null.String                           `json:"corn_expr" swaggertype:"string"` // corn表达式
	Status        x_null.Int64                            `json:"status" swaggertype:"number"`    // 状态
	CreatedBy     x_null.String                           `json:"created_by" swaggertype:"string"` // 创建人
	CreatedByUser system_schema.SystemAuthAdminSimpleInfo `json:"created_by_user" swaggertype:"string"` // 创建人
	CreateTime    x_null.Time                             `json:"create_time" swaggertype:"string"` // 创建时间
	UpdateTime    x_null.Time                             `json:"update_time" swaggertype:"string"` // 更新时间
}
