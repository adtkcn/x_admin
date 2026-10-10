package iot_schema

import "github.com/adtkcn/x_null"

// ============ 规则链 CRUD ============

// RuleChainListReq 规则链列表参数（查询条件用 x_null.String，不传/空串=不筛选）
type RuleChainListReq struct {
	Name   x_null.String `json:"name" form:"name"`       // 规则链名称(模糊)
	IsStop x_null.String `json:"is_stop" form:"is_stop"` // 是否停用: 0正常 1停用，不传=全部
}

// RuleChainAddReq 规则链新增参数
type RuleChainAddReq struct {
	Name      string `json:"name" form:"name" binding:"required,max=100"` // 规则链名称
	Remark    string `json:"remark" form:"remark" binding:"max=255"`      // 备注
	GraphData string `json:"graph_data" form:"graph_data"`                // 规则链图JSON(LogicFlow nodes+edges)
	IsStop    uint8  `json:"is_stop" form:"is_stop" binding:"oneof=0 1"`  // 是否停用
}

// RuleChainEditReq 规则链编辑参数
type RuleChainEditReq struct {
	ID        string `json:"id" form:"id" binding:"required"`             // 主键
	Name      string `json:"name" form:"name" binding:"required,max=100"` // 规则链名称
	Remark    string `json:"remark" form:"remark" binding:"max=255"`      // 备注
	GraphData string `json:"graph_data" form:"graph_data"`                // 规则链图JSON
	IsStop    uint8  `json:"is_stop" form:"is_stop" binding:"oneof=0 1"`  // 是否停用
}

// RuleChainDelReq 规则链删除参数
type RuleChainDelReq struct {
	ID string `json:"id" form:"id" binding:"required"` // 主键
}

// RuleChainDetailReq 规则链详情参数
type RuleChainDetailReq struct {
	ID string `json:"id" form:"id" binding:"required"` // 主键
}

// RuleChainResp 规则链返回信息
type RuleChainResp struct {
	ID         string      `json:"id"`          // 主键
	Name       string      `json:"name"`        // 规则链名称
	Remark     string      `json:"remark"`      // 备注
	GraphData  string      `json:"graph_data"`  // 规则链图JSON
	IsStop     uint8       `json:"is_stop"`     // 是否停用
	CreateTime x_null.Time `json:"create_time"` // 创建时间
	UpdateTime x_null.Time `json:"update_time"` // 更新时间
}

// ============ 规则链图结构（与前端 LogicFlow getGraphData 对齐） ============
//
// 节点私有属性沿用审批流“命名空间收纳”约定：properties 只存在当前节点类型那一份，
// JSON tag（snake_case）= 命名空间 key = 前端 property.type.ts 的 NodeKey，
// 新增一种节点只在这里加一个指针字段 + 一个 struct，不污染其他节点。

// 规则链节点类型常量（前缀 rule: 与 LogicFlow 注册名、执行器分发表一致）
const (
	RcNodeTrigger  = "rule:trigger"  // 触发器：按报文类型入口过滤（当前仅属性上报 telemetry）
	RcNodeDecode   = "rule:decode"   // 解码：goja decode(values) 基于初始上报值产出属性对象
	RcNodeFilter   = "rule:filter"   // 数据过滤：goja filter(values) 返回布尔，false 丢弃整条报文
	RcNodeParallel = "rule:parallel" // 并行：唯一允许多出口的节点，所有子路径同时分流
	RcNodeBool     = "rule:bool"     // 布尔：goja filter(values) 返回 true 继续 / false 仅忽略本分支
	RcNodeAlarm    = "rule:alarm"    // 告警：goja alert(values)+持续/恢复状态机
	RcNodeState    = "rule:state"    // 状态快照：写 Redis 最新属性，供管理端/大屏即时读
	RcNodeHistory  = "rule:history"  // 历史存储：写 MySQL 时序历史表
	RcNodePush     = "rule:push"     // 大屏推送：属性广播到区域大屏（地图房间）
	RcNodeNotify   = "rule:notify"   // 通知：站内/邮件/webhook（可扩展）
)

// 节点属性命名空间 key（与前端 properties[key] 对应）
const (
	RcPropsTrigger  = "rule_trigger"
	RcPropsDecode   = "rule_decode"
	RcPropsFilter   = "rule_filter"
	RcPropsParallel = "rule_parallel"
	RcPropsBool     = "rule_bool"
	RcPropsAlarm    = "rule_alarm"
	RcPropsState    = "rule_state"
	RcPropsHistory  = "rule_history"
	RcPropsPush     = "rule_push"
	RcPropsNotify   = "rule_notify"
)

// RcTriggerProps 触发器：仅处理命中的报文类型（空=全部）。
type RcTriggerProps struct {
	MsgTypes []string `json:"msg_types"` // 可选报文类型，当前仅 telemetry
}

// RcScriptProps 单脚本节点（解码/过滤/布尔共用）：script 为空则该级透传。
// 布尔节点同样约定函数名 filter(values, sn)：true 继续、false 忽略本分支。
type RcScriptProps struct {
	Script string `json:"script"`
}

// RcAlarmProps 告警：脚本返回级别(0=不告警)，duration/recover 为状态机参数。
type RcAlarmProps struct {
	Script   string `json:"script"`
	Duration int    `json:"duration"` // 持续时长(秒)，0=命中即告
	Recover  int    `json:"recover"`  // 恢复确认时长(秒)，0=恢复正常即收；>0 需连续正常满该秒才收尾
}

// RcNotifyProps 通知节点（对齐审批流 NotifyTask，便于后续扩展 sms 等）。
type RcNotifyProps struct {
	ServiceType    string   `json:"service_type"` // site/email/webhook
	ServiceContent string   `json:"service_content"`
	ReceiverId     []string `json:"receiver_id"`
	EmailTo        []string `json:"email_to"`
	WebhookUrl     string   `json:"webhook_url"`
}

// RcNodeProps 节点私有属性：按命名空间收纳各类型那一份，未使用的不进 JSON。
// 并行/状态快照/历史存储/大屏推送无配置，不进属性表。
type RcNodeProps struct {
	Trigger *RcTriggerProps `json:"rule_trigger,omitempty"`
	Decode  *RcScriptProps  `json:"rule_decode,omitempty"`
	Filter  *RcScriptProps  `json:"rule_filter,omitempty"`
	Bool    *RcScriptProps  `json:"rule_bool,omitempty"`
	Alarm   *RcAlarmProps   `json:"rule_alarm,omitempty"`
	Notify  *RcNotifyProps  `json:"rule_notify,omitempty"`
}

// RuleNode LogicFlow 节点：仅取执行需要的字段，text 保留标签。
type RuleNode struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Text struct {
		Value string `json:"value"`
	} `json:"text"`
	Properties RcNodeProps `json:"properties"`
}

// RuleEdge LogicFlow 边：source/target；除并行外每个节点只允许一条出边（保存时校验）。
type RuleEdge struct {
	ID           string `json:"id"`
	Type         string `json:"type"`
	SourceNodeId string `json:"sourceNodeId"`
	TargetNodeId string `json:"targetNodeId"`
	Text         struct {
		Value string `json:"value"`
	} `json:"text"`
}

// RuleChainGraph 规则链有向图（LogicFlow getGraphData 原样）。
type RuleChainGraph struct {
	Nodes []RuleNode `json:"nodes"`
	Edges []RuleEdge `json:"edges"`
}
