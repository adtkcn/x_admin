package flowSchema

// FlowTemplateListReq 流程模板列表参数
type FlowTemplateListReq struct {
	FlowName            string `form:"flowName"`            // 流程名称
	FlowGroup           int    `form:"flowGroup"`           // 流程分类
	FlowRemark          string `form:"flowRemark"`          // 流程描述
	FlowFormData        string `form:"flowFormData"`        // 表单配置
	FlowProcessData     string `form:"flowProcessData"`     // 流程配置
	FlowProcessDataList string `form:"flowProcessDataList"` // 流程配置list数据
}

// FlowTemplateDetailReq 流程模板详情参数
type FlowTemplateDetailReq struct {
	Id string `form:"id"` //
}

// FlowTemplateAddReq 流程模板新增参数
type FlowTemplateAddReq struct {
	FlowName            string `form:"flowName"`            // 流程名称
	FlowGroup           int    `form:"flowGroup"`           // 流程分类
	FlowRemark          string `form:"flowRemark"`          // 流程描述
	FlowFormData        string `form:"flowFormData"`        // 表单配置
	FlowProcessData     string `form:"flowProcessData"`     // 流程配置
	FlowProcessDataList string `form:"flowProcessDataList"` // 流程配置list数据

}

// FlowTemplateEditReq 流程模板新增参数
type FlowTemplateEditReq struct {
	Id                  string `form:"id"`                  //
	FlowName            string `form:"flowName"`            // 流程名称
	FlowGroup           int    `form:"flowGroup"`           // 流程分类
	FlowRemark          string `form:"flowRemark"`          // 流程描述
	FlowFormData        string `form:"flowFormData"`        // 表单配置
	FlowProcessData     string `form:"flowProcessData"`     // 流程配置
	FlowProcessDataList string `form:"flowProcessDataList"` // 流程配置list数据

}

// FlowTemplateDelReq 流程模板新增参数
type FlowTemplateDelReq struct {
	Id string `form:"id"` //
}

// FlowTemplateResp 流程模板返回信息
type FlowTemplateResp struct {
	Id                  string `json:"id" structs:"id"`                           //
	FlowName            string `json:"flowName" structs:"flowName"`               // 流程名称
	FlowGroup           int    `json:"flowGroup" structs:"flowGroup"`             // 流程分类
	FlowRemark          string `json:"flowRemark" structs:"flowRemark"`           // 流程描述
	FlowFormData        string `json:"flowFormData" structs:"flowFormData"`       // 表单配置
	FlowProcessData     string `json:"flowProcessData" structs:"flowProcessData"` // 流程配置
	FlowProcessDataList string `form:"flowProcessDataList"`                       // 流程配置list数据
}
