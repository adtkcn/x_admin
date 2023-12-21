package flow_apply

import "x_admin/core"

//FlowApplyListReq 申请流程列表参数
type FlowApplyListReq struct {
	TemplateId          int    `form:"templateId"`          // 模板
	ApplyUserId         int    `form:"applyUserId"`         // 申请人id
	ApplyUserNickname   string `form:"applyUserNickname"`   // 申请人昵称
	FlowName            string `form:"flowName"`            // 流程名称
	FlowGroup           int    `form:"flowGroup"`           // 流程分类
	FlowRemark          string `form:"flowRemark"`          // 流程描述
	FlowFormData        string `form:"flowFormData"`        // 表单配置
	FlowProcessData     string `form:"flowProcessData"`     // 流程配置
	FlowProcessDataList string `form:"flowProcessDataList"` // 流程配置list数据
	FormValue           string `form:"formValue"`           // 表单值
	Status              int    `form:"status"`              // 状态：1待提交，2审批中，3审批完成，4审批失败
}

//FlowApplyDetailReq 申请流程详情参数
type FlowApplyDetailReq struct {
	Id int `form:"id"` //
}

//FlowApplyAddReq 申请流程新增参数
type FlowApplyAddReq struct {
	TemplateId        int    `form:"templateId"`        // 模板
	ApplyUserId       int    `form:"applyUserId"`       // 申请人id
	ApplyUserNickname string `form:"applyUserNickname"` // 申请人昵称
	FlowName          string `form:"flowName"`          // 流程名称
	// FlowGroup           int    `form:"flowGroup"`           // 流程分类
	// FlowRemark          string `form:"flowRemark"`          // 流程描述
	// FlowFormData        string `form:"flowFormData"`        // 表单配置
	// FlowProcessData     string `form:"flowProcessData"`     // 流程配置
	// FlowProcessDataList string `form:"flowProcessDataList"` // 流程配置list数据
	FormValue string `form:"formValue"` // 表单值
	Status    int    `form:"status"`    // 状态：1待提交，2审批中，3审批完成，4审批失败
}

//FlowApplyEditReq 申请流程新增参数
type FlowApplyEditReq struct {
	Id int `form:"id"` //
	// TemplateId          int    `form:"templateId"`          // 模板
	// ApplyUserId         int    `form:"applyUserId"`         // 申请人id
	// ApplyUserNickname   string `form:"applyUserNickname"`   // 申请人昵称
	FlowName string `form:"flowName"` // 流程名称
	// FlowGroup           int    `form:"flowGroup"`           // 流程分类
	// FlowRemark          string `form:"flowRemark"`          // 流程描述
	// FlowFormData        string `form:"flowFormData"`        // 表单配置
	// FlowProcessData     string `form:"flowProcessData"`     // 流程配置
	// FlowProcessDataList string `form:"flowProcessDataList"` // 流程配置list数据
	FormValue string `form:"formValue"` // 表单值
	Status    int    `form:"status"`    // 状态：1待提交，2审批中，3审批完成，4审批失败
}

//FlowApplyDelReq 申请流程新增参数
type FlowApplyDelReq struct {
	Id int `form:"id"` //
}

//FlowApplyResp 申请流程返回信息
type FlowApplyResp struct {
	Id                  int         `json:"id" structs:"id"`                               //
	TemplateId          int         `json:"templateId" structs:"templateId"`               // 模板
	ApplyUserId         int         `json:"applyUserId" structs:"applyUserId"`             // 申请人id
	ApplyUserNickname   string      `json:"applyUserNickname" structs:"applyUserNickname"` // 申请人昵称
	FlowName            string      `json:"flowName" structs:"flowName"`                   // 流程名称
	FlowGroup           int         `json:"flowGroup" structs:"flowGroup"`                 // 流程分类
	FlowRemark          string      `json:"flowRemark" structs:"flowRemark"`               // 流程描述
	FlowFormData        string      `json:"flowFormData" structs:"flowFormData"`           // 表单配置
	FlowProcessData     string      `json:"flowProcessData" structs:"flowProcessData"`     // 流程配置
	FlowProcessDataList string      `json:"flowProcessDataList"`                           // 流程配置list数据
	FormValue           string      `json:"formValue"`                                     // 表单值
	Status              int         `json:"status" structs:"status"`                       // 状态：1待提交，2审批中，3审批完成，4审批失败
	UpdateTime          core.TsTime `json:"updateTime" structs:"updateTime"`               // 更新时间
	CreateTime          core.TsTime `json:"createTime" structs:"createTime"`               // 创建时间
}
