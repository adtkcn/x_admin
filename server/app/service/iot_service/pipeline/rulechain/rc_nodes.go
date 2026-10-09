package rulechain

import (
	"fmt"
	"strings"

	"x_admin/app/schema/iot_schema"
	"x_admin/app/schema/queue_schema"
	"x_admin/app/service/iot_service/pipeline/alert"
	"x_admin/app/service/iot_service/pipeline/online"
	"x_admin/app/service/iot_service/pipeline/store"
	"x_admin/app/service/notice_service"
	"x_admin/core"
	"x_admin/plugin/email"
)

// 规则链节点处理器注册表：每种节点一个 handler，签名统一为
//
//	handler(ctx, node) bool —— 执行节点副作用，返回本分支是否继续下钻 children（后继由 rcStep.children 决定）
//
// 新增一种节点只需：注册 handler + schema 加属性 struct + 前端加节点/属性子面板，不改遍历主逻辑；
// 未知类型按“继续下钻”兜底。
type rcHandler func(ctx *ruleCtx, node *iot_schema.RuleNode) bool

var rcHandlers = map[string]rcHandler{
	iot_schema.RcNodeTrigger: rcTrigger,
	iot_schema.RcNodeDecode:  rcDecode,
	iot_schema.RcNodeFilter:  rcFilter,
	iot_schema.RcNodeBool:    rcBool,
	iot_schema.RcNodeAlarm:   rcAlarm,
	iot_schema.RcNodeState:   rcState,
	iot_schema.RcNodeHistory: rcHistory,
	iot_schema.RcNodePush:    rcPush,
	iot_schema.RcNodeNotify:  rcNotify,
	// 并行节点的分流完全由树结构（多 children）表达，自身无副作用
	iot_schema.RcNodeParallel: rcPassthrough,
}

// applyRcNode 分派节点：未知类型按继续下钻处理（向前兼容未注册执行器的新节点）
func applyRcNode(ctx *ruleCtx, node *iot_schema.RuleNode) bool {
	if h, ok := rcHandlers[node.Type]; ok {
		return h(ctx, node)
	}
	return true
}

// nodeKey 节点脚本的编译缓键「链ID:节点ID」：nodeId 链内已唯一，链ID 做前缀让多产品共用同一条链时复用编译结果
func (c *ruleCtx) nodeKey(node *iot_schema.RuleNode) string {
	return c.model.RuleChainId + ":" + node.ID
}

// rcPassthrough 并行节点：分流由树的多 children 表达，自身无副作用，始终继续下钻。
func rcPassthrough(ctx *ruleCtx, node *iot_schema.RuleNode) bool {
	return true
}

// rcTrigger 触发器：配置了 msg_types 时只放行命中的报文类型，否则全放行。
func rcTrigger(ctx *ruleCtx, node *iot_schema.RuleNode) bool {
	if p := node.Properties.Trigger; p != nil && len(p.MsgTypes) > 0 {
		allowed := false
		for _, t := range p.MsgTypes {
			if t == ctx.msg.MsgType {
				allowed = true
				break
			}
		}
		if !allowed {
			ctx.halt = true
			return false
		}
	}
	return true
}

// rcDecode 解码节点：以当前流经属性执行 goja decode(values)，覆盖 ctx.values；产出空则丢弃本条。
func rcDecode(ctx *ruleCtx, node *iot_schema.RuleNode) bool {
	script := ""
	if node.Properties.Decode != nil {
		script = node.Properties.Decode.Script
	}
	values := execDecode(ctx.msg, ctx.nodeKey(node), script, ctx.values)
	if len(values) == 0 {
		ctx.halt = true
		ctx.values = nil
		return false
	}
	ctx.values = values
	return true
}

// rcFilter 数据过滤节点：goja filter(values) 返回 false 则丢弃整条报文（全局短路）。
func rcFilter(ctx *ruleCtx, node *iot_schema.RuleNode) bool {
	script := ""
	if node.Properties.Filter != nil {
		script = node.Properties.Filter.Script
	}
	if !execFilter(ctx.msg, ctx.nodeKey(node), script, ctx.values) {
		ctx.halt = true
		return false
	}
	return true
}

// rcBool 布尔节点：goja filter(values) 返 true 继续走子节点；false 只剪本分支（不置 halt）。
func rcBool(ctx *ruleCtx, node *iot_schema.RuleNode) bool {
	script := ""
	if node.Properties.Bool != nil {
		script = node.Properties.Bool.Script
	}
	return execFilter(ctx.msg, ctx.nodeKey(node), script, ctx.values)
}

// rcAlarm 告警节点：goja alert(values) 的级别交给告警子包状态机收敛；未配脚本则不参与告警。
func rcAlarm(ctx *ruleCtx, node *iot_schema.RuleNode) bool {
	p := node.Properties.Alarm
	if p == nil || p.Script == "" {
		return true
	}
	level := execAlertSeverity(ctx.msg, ctx.nodeKey(node), p.Script, ctx.values)
	alert.EvalAlert(ctx.msg, ctx.values, level, p.Duration, p.Silence)
	return true
}

// rcState 状态快照节点：把流经的最新属性写 Redis 快照，供管理端与大屏即时读取。
func rcState(ctx *ruleCtx, node *iot_schema.RuleNode) bool {
	online.SaveLatestValues(ctx.msg.DeviceId, ctx.values)
	return true
}

// rcHistory 历史存储节点：经存储子包追加写时序表，落库前受产品物模型白名单（ctx.model.ThingIds）约束。
func rcHistory(ctx *ruleCtx, node *iot_schema.RuleNode) bool {
	if err := store.Write(ctx.msg, ctx.values, ctx.model.ThingIds); err != nil {
		core.Logger.Errorf("[iot规则链] 历史存储写时序失败 sn=%s err=%v", ctx.msg.Sn, err)
	}
	return true
}

// rcPush 大屏推送节点：把流经的最新属性广播到设备所属区域的地图大屏房间。
func rcPush(ctx *ruleCtx, node *iot_schema.RuleNode) bool {
	online.BroadcastTelemetry(ctx.msg, ctx.values)
	return true
}

// rcNotify 通知节点：站内/邮件/webhook（邮件与 webhook 走异步队列，避免外部依赖阻塞报文处理）；单节点只配一种类型。
func rcNotify(ctx *ruleCtx, node *iot_schema.RuleNode) bool {
	p := node.Properties.Notify
	if p == nil || p.ServiceType == "" {
		return true
	}
	content := p.ServiceContent
	if content == "" {
		content = fmt.Sprintf("设备 %s 触发规则链通知", ctx.msg.Sn)
	}
	switch p.ServiceType {
	case "site":
		for _, rid := range dedupNonEmpty(p.ReceiverId) {
			if err := notice_service.NoticeService.Send(false, notice_service.NoticePayload{
				Type:       "warning",
				Title:      "IoT 规则链通知",
				Content:    content,
				ReceiverID: rid,
			}); err != nil {
				core.Logger.Errorf("[iot规则链] 站内通知失败 sn=%s receiver=%s err=%v", ctx.msg.Sn, rid, err)
			}
		}
	case "email":
		to := dedupNonEmpty(p.EmailTo)
		if len(to) == 0 {
			core.Logger.Warnf("[iot规则链] 邮件通知收件人为空，跳过 sn=%s", ctx.msg.Sn)
			return true
		}
		if err := core.Queue.Enqueue(queue_schema.QueueFlowNotifyEmail, email.EmailOptions{
			To:       to,
			Subject:  "【IoT 规则链通知】" + ctx.msg.Sn,
			HTMLBody: fmt.Sprintf("<p>%s</p>", content),
		}); err != nil {
			core.Logger.Errorf("[iot规则链] 投递邮件通知失败 sn=%s err=%v", ctx.msg.Sn, err)
		}
	case "webhook":
		if strings.TrimSpace(p.WebhookUrl) == "" {
			core.Logger.Warnf("[iot规则链] webhook 地址为空，跳过 sn=%s", ctx.msg.Sn)
			return true
		}
		if err := core.Queue.Enqueue(queue_schema.QueueFlowNotifyWebhook, queue_schema.FlowNotifyWebhookPayload{
			URL:     p.WebhookUrl,
			Content: content,
		}); err != nil {
			core.Logger.Errorf("[iot规则链] 投递 webhook 通知失败 sn=%s err=%v", ctx.msg.Sn, err)
		}
	default:
		core.Logger.Warnf("[iot规则链] 不支持的通知类型 %s sn=%s", p.ServiceType, ctx.msg.Sn)
	}
	return true
}

// dedupNonEmpty 去重且剔除空白的字符串列表。
func dedupNonEmpty(in []string) []string {
	seen := map[string]bool{}
	out := make([]string, 0, len(in))
	for _, s := range in {
		s = strings.TrimSpace(s)
		if s == "" || seen[s] {
			continue
		}
		seen[s] = true
		out = append(out, s)
	}
	return out
}
