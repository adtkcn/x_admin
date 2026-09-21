package flow_schema

// FlowTemplateListReq 流程模板列表参数
type FlowTemplateListReq struct {
	FlowName            string `json:"flow_name" form:"flow_name"`                           // 流程名称
	FlowGroup           int    `json:"flow_group" form:"flow_group"`                         // 流程分类
	FlowRemark          string `json:"flow_remark" form:"flow_remark"`                       // 流程描述
	FlowFormData        string `json:"flow_form_data" form:"flow_form_data"`                 // 表单配置
	FlowProcessData     string `json:"flow_process_data" form:"flow_process_data"`           // 流程配置
	FlowProcessDataList string `json:"flow_process_data_list" form:"flow_process_data_list"` // 流程配置list数据
}

// FlowTemplateDetailReq 流程模板详情参数
type FlowTemplateDetailReq struct {
	Id string `json:"id" form:"id"` //
}

// FlowTemplateAddReq 流程模板新增参数
type FlowTemplateAddReq struct {
	FlowName            string `json:"flow_name" form:"flow_name"`                           // 流程名称
	FlowGroup           int    `json:"flow_group" form:"flow_group"`                         // 流程分类
	FlowRemark          string `json:"flow_remark" form:"flow_remark"`                       // 流程描述
	FlowFormData        string `json:"flow_form_data" form:"flow_form_data"`                 // 表单配置
	FlowProcessData     string `json:"flow_process_data" form:"flow_process_data"`           // 流程配置
	FlowProcessDataList string `json:"flow_process_data_list" form:"flow_process_data_list"` // 流程配置list数据

}

// FlowTemplateEditReq 流程模板新增参数
type FlowTemplateEditReq struct {
	Id                  string `json:"id" form:"id"`                                         //
	FlowName            string `json:"flow_name" form:"flow_name"`                           // 流程名称
	FlowGroup           int    `json:"flow_group" form:"flow_group"`                         // 流程分类
	FlowRemark          string `json:"flow_remark" form:"flow_remark"`                       // 流程描述
	FlowFormData        string `json:"flow_form_data" form:"flow_form_data"`                 // 表单配置
	FlowProcessData     string `json:"flow_process_data" form:"flow_process_data"`           // 流程配置
	FlowProcessDataList string `json:"flow_process_data_list" form:"flow_process_data_list"` // 流程配置list数据

}

// FlowTemplateDelReq 流程模板新增参数
type FlowTemplateDelReq struct {
	Id string `json:"id" form:"id"` //
}

// FlowTemplateResp 流程模板返回信息
type FlowTemplateResp struct {
	Id                  string `json:"id"`                     //
	FlowName            string `json:"flow_name"`              // 流程名称
	FlowGroup           int    `json:"flow_group"`             // 流程分类
	FlowRemark          string `json:"flow_remark"`            // 流程描述
	FlowFormData        string `json:"flow_form_data"`         // 表单配置
	FlowProcessData     string `json:"flow_process_data"`      // 流程配置
	FlowProcessDataList string `json:"flow_process_data_list"` // 流程配置list数据
}
