package flow_schema

import "github.com/adtkcn/x_null"

// FlowApplyListReq 申请流程列表参数
type FlowApplyListReq struct {
	TemplateId          string `json:"template_id" form:"template_id"`                       // 模板
	ApplyUserId         string `json:"apply_user_id" form:"apply_user_id"`                   // 申请人id
	ApplyUserNickname   string `json:"apply_user_nickname" form:"apply_user_nickname"`       // 申请人昵称
	FlowName            string `json:"flow_name" form:"flow_name"`                           // 流程名称
	FlowGroup           int    `json:"flow_group" form:"flow_group"`                         // 流程分类
	FlowRemark          string `json:"flow_remark" form:"flow_remark"`                       // 流程描述
	FlowFormData        string `json:"flow_form_data" form:"flow_form_data"`                 // 表单配置
	FlowProcessData     string `json:"flow_process_data" form:"flow_process_data"`           // 流程配置
	FlowProcessDataList string `json:"flow_process_data_list" form:"flow_process_data_list"` // 流程配置list数据
	FormValue           string `json:"form_value" form:"form_value"`                         // 表单值
	Status              int    `json:"status" form:"status"`                                 // 状态：1待提交，2审批中，3审批完成，4审批失败
}

// FlowApplyDetailReq 申请流程详情参数
type FlowApplyDetailReq struct {
	Id string `json:"id" form:"id"` //
}

// FlowApplyAddReq 申请流程新增参数
type FlowApplyAddReq struct {
	TemplateId        string `json:"template_id" form:"template_id"`                 // 模板
	ApplyUserId       string `json:"apply_user_id" form:"apply_user_id"`             // 申请人id
	ApplyUserNickname string `json:"apply_user_nickname" form:"apply_user_nickname"` // 申请人昵称
	FlowName          string `json:"flow_name" form:"flow_name"`                     // 流程名称
	FormValue         string `json:"form_value" form:"form_value"`                   // 表单值
	Status            int    `json:"status" form:"status"`                           // 状态：1待提交，2审批中，3审批完成，4审批失败
}

// FlowApplyEditReq 申请流程新增参数
type FlowApplyEditReq struct {
	Id        string `json:"id" form:"id"`                 //
	FlowName  string `json:"flow_name" form:"flow_name"`   // 流程名称
	FormValue string `json:"form_value" form:"form_value"` // 表单值
	Status    int    `json:"status" form:"status"`         // 状态：1待提交，2审批中，3审批完成，4审批失败
}

// FlowApplyDelReq 申请流程新增参数
type FlowApplyDelReq struct {
	Id string `json:"id" form:"id"` //
}

// FlowApplyResp 申请流程返回信息
type FlowApplyResp struct {
	Id                  string      `json:"id"`                     //
	TemplateId          string      `json:"template_id"`            // 模板
	ApplyUserId         string      `json:"apply_user_id"`          // 申请人id
	ApplyUserNickname   string      `json:"apply_user_nickname"`    // 申请人昵称
	FlowName            string      `json:"flow_name"`              // 流程名称
	FlowGroup           int         `json:"flow_group"`             // 流程分类
	FlowRemark          string      `json:"flow_remark"`            // 流程描述
	FlowFormData        string      `json:"flow_form_data"`         // 表单配置
	FlowProcessData     string      `json:"flow_process_data"`      // 流程配置
	FlowProcessDataList string      `json:"flow_process_data_list"` // 流程配置list数据
	FormValue           string      `json:"form_value"`             // 表单值
	Status              int         `json:"status"`                 // 状态：1待提交，2审批中，3审批完成，4审批失败
	UpdateTime          x_null.Time `json:"update_time"`            // 更新时间
	CreateTime          x_null.Time `json:"create_time"`            // 创建时间
}
