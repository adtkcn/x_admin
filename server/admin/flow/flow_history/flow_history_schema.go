package flow_history

import "x_admin/core"

//FlowHistoryListReq 流程历史列表参数
type FlowHistoryListReq struct {
	ApplyId           int    `form:"applyId"`           // 申请id
	TemplateId        int    `form:"templateId"`        // 模板id
	ApplyUserId       int    `form:"applyUserId"`       // 申请人id
	ApplyUserNickname string `form:"applyUserNickname"` // 申请人昵称
	ApproverId        int    `form:"approverId"`        // 审批人id
	ApproverNickname  string `form:"approverNickname"`  // 审批用户昵称
	NodeId            string `form:"nodeId"`            // 节点
	FormValue         string `form:"formValue"`         // 表单值
	PassStatus        int    `form:"passStatus"`        // 通过状态：1待处理，2通过，3拒绝
	PassRemark        string `form:"passRemark"`        // 通过备注
}

//FlowHistoryDetailReq 流程历史详情参数
type FlowHistoryDetailReq struct {
	Id int `form:"id"` // 历史id
}

//FlowHistoryAddReq 流程历史新增参数
type FlowHistoryAddReq struct {
	ApplyId           int    `form:"applyId"`           // 申请id
	TemplateId        int    `form:"templateId"`        // 模板id
	ApplyUserId       int    `form:"applyUserId"`       // 申请人id
	ApplyUserNickname string `form:"applyUserNickname"` // 申请人昵称
	ApproverId        int    `form:"approverId"`        // 审批人id
	ApproverNickname  string `form:"approverNickname"`  // 审批用户昵称
	NodeId            string `form:"nodeId"`            // 节点
	FormValue         string `form:"formValue"`         // 表单值
	PassStatus        int    `form:"passStatus"`        // 通过状态：1待处理，2通过，3拒绝
	PassRemark        string `form:"passRemark"`        // 通过备注
}

//FlowHistoryEditReq 流程历史新增参数
type FlowHistoryEditReq struct {
	Id                int    `form:"id"`                // 历史id
	ApplyId           int    `form:"applyId"`           // 申请id
	TemplateId        int    `form:"templateId"`        // 模板id
	ApplyUserId       int    `form:"applyUserId"`       // 申请人id
	ApplyUserNickname string `form:"applyUserNickname"` // 申请人昵称
	ApproverId        int    `form:"approverId"`        // 审批人id
	ApproverNickname  string `form:"approverNickname"`  // 审批用户昵称
	NodeId            string `form:"nodeId"`            // 节点
	FormValue         string `form:"formValue"`         // 表单值
	PassStatus        int    `form:"passStatus"`        // 通过状态：1待处理，2通过，3拒绝
	PassRemark        string `form:"passRemark"`        // 通过备注
}

//FlowHistoryDelReq 流程历史新增参数
type FlowHistoryDelReq struct {
	Id int `form:"id"` // 历史id
}

//FlowHistoryResp 流程历史返回信息
type FlowHistoryResp struct {
	Id                int         `json:"id" structs:"id"`                               // 历史id
	ApplyId           int         `json:"applyId" structs:"applyId"`                     // 申请id
	TemplateId        int         `json:"templateId" structs:"templateId"`               // 模板id
	ApplyUserId       int         `json:"applyUserId" structs:"applyUserId"`             // 申请人id
	ApplyUserNickname string      `json:"applyUserNickname" structs:"applyUserNickname"` // 申请人昵称
	ApproverId        int         `json:"approverId" structs:"approverId"`               // 审批人id
	ApproverNickname  string      `json:"approverNickname" structs:"approverNickname"`   // 审批用户昵称
	NodeId            string      `json:"nodeId" structs:"nodeId"`                       // 节点
	FormValue         string      `json:"formValue" structs:"formValue"`                 // 表单值
	PassStatus        int         `json:"passStatus" structs:"passStatus"`               // 通过状态：1待处理，2通过，3拒绝
	PassRemark        string      `json:"passRemark" structs:"passRemark"`               // 通过备注
	UpdateTime        core.TsTime `json:"updateTime" structs:"updateTime"`               // 更新时间
	CreateTime        core.TsTime `json:"createTime" structs:"createTime"`               // 创建时间
}

type FlowTree struct {
	Id    string `json:"id"`
	Pid   string `json:"pid"`
	Label string `json:"label"`
	Type  string `json:"type"`

	UserId int `json:"userId"`
	DeptId int `json:"deptId"`
	PostId int `json:"postId"`

	FieldAuth map[string]int `json:"fieldAuth"`

	Children *[]FlowTree `json:"children"`
}
type NextNodeReq struct {
	ApplyId int `form:"applyId"` // 申请id
	// CurrentNodeId   string `form:"currentNodeId"` // 流程里的节点id
	// FormValue       string `form:"formValue"`
	// NextNodeAdminId int `form:"nextNodeAdminId"` // 下一个节点的审批用户id
}
type PassReq struct {
	ApplyId int `form:"applyId"` // 申请id

	// CurrentNodeId   string `form:"currentNodeId"` // 流程里的节点id
	// FormValue       string `form:"formValue"`
	NextNodeAdminId int    `form:"nextNodeAdminId"` // 下一个节点的审批用户id
	PassRemark      string `form:"passRemark"`      // 通过备注
}
