package flow_schema

import "github.com/adtkcn/x_null"

// FlowHistoryListReq 流程历史列表参数
type FlowHistoryListReq struct {
	ApplyId           string `json:"apply_id" form:"apply_id"`                       // 申请id
	TemplateId        string `json:"template_id" form:"template_id"`                 // 模板id
	ApplyUserId       string `json:"apply_user_id" form:"apply_user_id"`             // 申请人id
	ApplyUserNickname string `json:"apply_user_nickname" form:"apply_user_nickname"` // 申请人昵称
	ApproverId        string `json:"approver_id" form:"approver_id"`                 // 审批人id
	ApproverNickname  string `json:"approver_nickname" form:"approver_nickname"`     // 审批用户昵称
	NodeId            string `json:"node_id" form:"node_id"`                         // 节点
	NodeLabel         string `json:"node_label" form:"node_label"`                   // 节点名称
	NodeType          string `json:"node_type" form:"node_type"`                     // 节点类型
	FormValue         string `json:"form_value" form:"form_value"`                   // 表单值
	PassStatus        int    `json:"pass_status" form:"pass_status"`                 // 通过状态：1待处理，2通过，3拒绝
	PassRemark        string `json:"pass_remark" form:"pass_remark"`                 // 通过备注
	IsShow            int    `json:"is_show" form:"is_show"`                         // 是否显示：0隐藏，1显示
}

// FlowHistoryDetailReq 流程历史详情参数
type FlowHistoryDetailReq struct {
	Id string `json:"id" form:"id"` // 历史id
}

// FlowHistoryAddReq 流程历史新增参数
type FlowHistoryAddReq struct {
	ApplyId           string `json:"apply_id" form:"apply_id"`                       // 申请id
	TemplateId        string `json:"template_id" form:"template_id"`                 // 模板id
	ApplyUserId       string `json:"apply_user_id" form:"apply_user_id"`             // 申请人id
	ApplyUserNickname string `json:"apply_user_nickname" form:"apply_user_nickname"` // 申请人昵称
	ApproverId        string `json:"approver_id" form:"approver_id"`                 // 审批人id
	ApproverNickname  string `json:"approver_nickname" form:"approver_nickname"`     // 审批用户昵称
	NodeId            string `json:"node_id" form:"node_id"`                         // 节点
	NodeLabel         string `json:"node_label" form:"node_label"`                   // 节点名称
	NodeType          string `json:"node_type" form:"node_type"`                     // 节点类型
	FormValue         string `json:"form_value" form:"form_value"`                   // 表单值
	PassStatus        int    `json:"pass_status" form:"pass_status"`                 // 通过状态：1待处理，2通过，3拒绝
	PassRemark        string `json:"pass_remark" form:"pass_remark"`                 // 通过备注
}

// FlowHistoryEditReq 流程历史新增参数
type FlowHistoryEditReq struct {
	Id                string `json:"id" form:"id"`                                   // 历史id
	ApplyId           string `json:"apply_id" form:"apply_id"`                       // 申请id
	TemplateId        string `json:"template_id" form:"template_id"`                 // 模板id
	ApplyUserId       string `json:"apply_user_id" form:"apply_user_id"`             // 申请人id
	ApplyUserNickname string `json:"apply_user_nickname" form:"apply_user_nickname"` // 申请人昵称
	ApproverId        string `json:"approver_id" form:"approver_id"`                 // 审批人id
	ApproverNickname  string `json:"approver_nickname" form:"approver_nickname"`     // 审批用户昵称
	NodeId            string `json:"node_id" form:"node_id"`                         // 节点
	NodeLabel         string `json:"node_label" form:"node_label"`                   // 节点名称
	NodeType          string `json:"node_type" form:"node_type"`                     // 节点类型
	FormValue         string `json:"form_value" form:"form_value"`                   // 表单值
	PassStatus        int    `json:"pass_status" form:"pass_status"`                 // 通过状态：1待处理，2通过，3拒绝
	PassRemark        string `json:"pass_remark" form:"pass_remark"`                 // 通过备注
}

// FlowHistoryDelReq 流程历史新增参数
type FlowHistoryDelReq struct {
	Id string `json:"id" form:"id"` // 历史id
}

// FlowHistoryResp 流程历史返回信息
type FlowHistoryResp struct {
	Id                string      `json:"id"`                  // 历史id
	ApplyId           string      `json:"apply_id"`            // 申请id
	TemplateId        string      `json:"template_id"`         // 模板id
	ApplyUserId       string      `json:"apply_user_id"`       // 申请人id
	ApplyUserNickname string      `json:"apply_user_nickname"` // 申请人昵称
	ApproverId        string      `json:"approver_id"`         // 审批人id
	ApproverNickname  string      `json:"approver_nickname"`   // 审批用户昵称
	NodeId            string      `json:"node_id"`             // 节点
	NodeType          string      `json:"node_type"`           // 节点类型
	NodeLabel         string      `json:"node_label"`          // 节点名称
	FormValue         string      `json:"form_value"`          // 表单值
	PassStatus        int         `json:"pass_status"`         // 通过状态：1待处理，2通过，3拒绝
	PassRemark        string      `json:"pass_remark"`         // 通过备注
	IsShow            int         `json:"is_show"`             // 是否显示：0隐藏，1显示
	UpdateTime        x_null.Time `json:"update_time"`         // 更新时间
	CreateTime        x_null.Time `json:"create_time"`         // 创建时间
}

// 流程节点类型常量，统一使用 BPMN 标准类型标识，避免散落的字符串字面量拼写错误
const (
	NodeStartEvent       = "bpmn:startEvent"
	NodeUserTask         = "bpmn:userTask"
	NodeNotifyTask       = "bpmn:notifyTask"
	NodeExclusiveGateway = "bpmn:exclusiveGateway"
	NodeEndEvent         = "bpmn:endEvent"
)

// 流程节点属性命名空间 key（snake_case 短名），与前端 properties[key] 对应。
// 前端 LogicFlow 节点 properties 以本 key 收纳各节点类型的私有数据，例如
// properties["user_task"] = { "user_type": 3, "user_id": "1", ... }
const (
	NodePropsStartEvent       = "start_event"
	NodePropsUserTask         = "user_task"
	NodePropsNotifyTask       = "notify_task"
	NodePropsExclusiveGateway = "exclusive_gateway"
	NodePropsEndEvent         = "end_event"
)

// GatewayCondition 网关流转条件
type GatewayCondition struct {
	Id        string `json:"id"`        // 表单项 id
	Condition string `json:"condition"` // 判断符：== != >= <= include
	Value     string `json:"value"`     // 比较值
}

// 各节点类型私有属性（独立 struct，互不干扰，新增节点类型只加这里）。
// JSON tag 全部小写+下划线，与前端 properties[key] 完全一致。

// StartEventProps 开始节点私有属性
type StartEventProps struct {
	FieldAuth map[string]int `json:"field_auth"` // 表单项权限：1读写 2只读 3隐藏
}

// UserTaskProps 审批节点私有属性
type UserTaskProps struct {
	FieldAuth map[string]int `json:"field_auth"` // 表单项权限：1读写 2只读 3隐藏
	UserType  int            `json:"user_type"`  // 1指定部门、岗位 2用户部门负责人 3指定审批人
	UserId    string         `json:"user_id"`
	DeptId    string         `json:"dept_id"`
	PostId    string         `json:"post_id"`
}

// NotifyTaskProps 通知节点私有属性
type NotifyTaskProps struct {
	ServiceType    string   `json:"service_type"`    // site=站内消息 email=邮件 webhook=回调
	ServiceContent string   `json:"service_content"` // 消息内容
	ReceiverId     []string `json:"receiver_id"`     // 站内消息接收人 admin_id 列表，为空则默认通知申请人
	EmailTo        []string `json:"email_to"`        // 邮件收件邮箱数组（手填 + 用户邮箱）
	WebhookUrl     string   `json:"webhook_url"`     // webhook 场景：回调地址
}

// GatewayProps 排他网关节点私有属性
type GatewayProps struct {
	Gateway []GatewayCondition `json:"gateway"`
}

// EndEventProps 结束节点私有属性（无配置项）
type EndEventProps struct{}

// NodeProps 节点私有属性。
// 按节点类型命名空间 key 收纳各自那一份独立 struct，与前端 properties 结构一致；
// 指针类型 + omitempty：未使用的命名空间不出现在 JSON 里，新增节点类型只加一个字段，不污染其他节点。
type NodeProps struct {
	StartEvent       *StartEventProps `json:"start_event,omitempty"`
	UserTask         *UserTaskProps   `json:"user_task,omitempty"`
	NotifyTask       *NotifyTaskProps `json:"notify_task,omitempty"`
	ExclusiveGateway *GatewayProps    `json:"exclusive_gateway,omitempty"`
	EndEvent         *EndEventProps   `json:"end_event,omitempty"`
}

// FlowTree 流程节点树结构
type FlowTree struct {
	Id       string      `json:"id"`
	Pid      string      `json:"pid"`
	Label    string      `json:"label"`
	Type     string      `json:"type"` // bpmn:startEvent / bpmn:userTask / ...
	Props    NodeProps   `json:"props"`
	Children *[]FlowTree `json:"children"`
}
type NextNodeReq struct {
	ApplyId string `json:"apply_id" form:"apply_id"` // 申请id
}
type PassReq struct {
	ApplyId         string `json:"apply_id" form:"apply_id"`                     // 申请id
	NextNodeAdminId string `json:"next_node_admin_id" form:"next_node_admin_id"` // 下一个节点的审批用户id
	PassRemark      string `json:"pass_remark" form:"pass_remark"`               // 通过备注
}
type BackReq struct {
	ApplyId   string `json:"apply_id" form:"apply_id"`     // 申请 id
	HistoryId string `json:"history_id" form:"history_id"` //审批节点
	Remark    string `json:"remark" form:"remark"`         // 备注
}

// FlowNotifyEmailPayload 流程邮件通知异步任务载荷
type FlowNotifyEmailPayload struct {
	To       []string `json:"to"`        // 收件人邮箱列表
	Subject  string   `json:"subject"`   // 邮件主题
	HTMLBody string   `json:"html_body"` // 邮件 HTML 正文
}

// FlowNotifyWebhookPayload 流程 Webhook 回调异步任务载荷
type FlowNotifyWebhookPayload struct {
	URL     string `json:"url"`     // 回调地址
	Content string `json:"content"` // 回调内容
}
