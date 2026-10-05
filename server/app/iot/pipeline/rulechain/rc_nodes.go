package rulechain

import (
	"fmt"
	"strings"
	"x_admin/app/iot/pipeline/alert"
	"x_admin/app/iot/pipeline/online"
	"x_admin/app/iot/pipeline/store"
	"x_admin/app/schema/iot_schema"
	"x_admin/app/schema/queue_schema"
	"x_admin/app/service/notice_service"
	"x_admin/core"
	"x_admin/plugin/email"
)

// 规则链节点处理器注册表：每种节点一个 handler，签名统一
//
//	handler(ctx, node, g) []string  —— 执行节点副作用，返回要跟进的出边目标节点ID
//
// 新增一种节点（如发送通知之外的 sms/推送第三方）只需：注册一个 handler + schema 加属性 struct
// + 前端加节点/属性子面板，不改遍历主逻辑；未知类型按“透传全部出边”兜底。
type rcHandler func(ctx *ruleCtx, node *iot_schema.RuleNode, g *parsedGraph) []string

var rcHandlers = map[string]rcHandler{
	iot_schema.RcNodeTrigger:     rcTrigger,
	iot_schema.RcNodeDecode:      rcDecode,
	iot_schema.RcNodeThingFilter: rcThingFilter,
	iot_schema.RcNodeFilter:      rcFilter,
	iot_schema.RcNodeParallel:    rcParallel,
	iot_schema.RcNodeBool:        rcBool,
	iot_schema.RcNodeAlarm:       rcAlarm,
	iot_schema.RcNodeState:       rcState,
	iot_schema.RcNodeHistory:     rcHistory,
	iot_schema.RcNodePush:        rcPush,
	iot_schema.RcNodeNotify:      rcNotify,
}

// applyRcNode 分派节点：未知类型按透传全部出边处理（向前兼容新增节点未注册执行器的场景）。
func applyRcNode(ctx *ruleCtx, node *iot_schema.RuleNode, g *parsedGraph) []string {
	if h, ok := rcHandlers[node.Type]; ok {
		return h(ctx, node, g)
	}
	return allTargets(g, node.ID)
}

// allTargets 节点的全部出边目标（无条件流转的节点用）。
func allTargets(g *parsedGraph, nodeId string) []string {
	edges := g.out[nodeId]
	if len(edges) == 0 {
		return nil
	}
	ids := make([]string, 0, len(edges))
	for _, e := range edges {
		ids = append(ids, e.TargetNodeId)
	}
	return ids
}

// rcTrigger 触发器：配置了 msg_types 时只放行命中的报文类型，否则全放行。
func rcTrigger(ctx *ruleCtx, node *iot_schema.RuleNode, g *parsedGraph) []string {
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
			return nil
		}
	}
	return allTargets(g, node.ID)
}

// rcDecode 解码节点：执行 goja decode(raw)，覆盖 ctx.values；产出空则丢弃本条。
func rcDecode(ctx *ruleCtx, node *iot_schema.RuleNode, g *parsedGraph) []string {
	script := ""
	if node.Properties.Decode != nil {
		script = node.Properties.Decode.Script
	}
	values := execDecode(ctx.msg, scriptCacheKey(ctx.model.RuleChainId, node.ID), script)
	if len(values) == 0 {
		ctx.halt = true
		ctx.values = nil
		return nil
	}
	ctx.values = values
	return allTargets(g, node.ID)
}

// rcThingFilter 物模型白名单过滤：固定开启，无配置；裁剪后为空则丢弃整条报文。
func rcThingFilter(ctx *ruleCtx, node *iot_schema.RuleNode, g *parsedGraph) []string {
	ctx.values = filterByThingModel(ctx.msg, ctx.values, ctx.model.ThingModel)
	if len(ctx.values) == 0 {
		ctx.halt = true
		return nil
	}
	return allTargets(g, node.ID)
}

// rcFilter 数据过滤节点：goja filter(values) 返回 false 则丢弃整条报文（全局短路）。
func rcFilter(ctx *ruleCtx, node *iot_schema.RuleNode, g *parsedGraph) []string {
	script := ""
	if node.Properties.Filter != nil {
		script = node.Properties.Filter.Script
	}
	if !execFilter(ctx.msg, scriptCacheKey(ctx.model.RuleChainId, node.ID), script, ctx.values) {
		ctx.halt = true
		return nil
	}
	return allTargets(g, node.ID)
}

// rcParallel 并行节点：所有出边子路径同时分流，唯一允许多出口的节点类型。
func rcParallel(ctx *ruleCtx, node *iot_schema.RuleNode, g *parsedGraph) []string {
	return allTargets(g, node.ID)
}

// rcBool 布尔节点：goja filter(values) 返回 true 继续走子节点；
// false 只剪掉本分支（不置 halt），并行的其他分支照常执行。
func rcBool(ctx *ruleCtx, node *iot_schema.RuleNode, g *parsedGraph) []string {
	script := ""
	if node.Properties.Bool != nil {
		script = node.Properties.Bool.Script
	}
	if !execFilter(ctx.msg, scriptCacheKey(ctx.model.RuleChainId, node.ID), script, ctx.values) {
		return nil
	}
	return allTargets(g, node.ID)
}

// rcAlarm 告警节点：goja alert(values) 返回级别，交给告警子包的状态机（持续/静默/升级/恢复）收敛。
// 未配脚本则不参与告警，链路上零额外 Redis 开销。
func rcAlarm(ctx *ruleCtx, node *iot_schema.RuleNode, g *parsedGraph) []string {
	p := node.Properties.Alarm
	if p == nil || p.Script == "" {
		return allTargets(g, node.ID)
	}
	level := execAlertSeverity(ctx.msg, scriptCacheKey(ctx.model.RuleChainId, node.ID), p.Script, ctx.values)
	alert.EvalAlert(ctx.msg, ctx.values, level, p.Duration, p.Silence)
	return allTargets(g, node.ID)
}

// rcState 状态快照节点：把流经的最新属性写 Redis 快照，供管理端与大屏即时读取。
func rcState(ctx *ruleCtx, node *iot_schema.RuleNode, g *parsedGraph) []string {
	online.SaveLatestValues(ctx.msg.DeviceId, ctx.values)
	return allTargets(g, node.ID)
}

// rcHistory 历史存储节点：把流经的属性经存储子包追加写时序历史表。
func rcHistory(ctx *ruleCtx, node *iot_schema.RuleNode, g *parsedGraph) []string {
	if err := store.Write(ctx.msg, ctx.values); err != nil {
		core.Logger.Errorf("[iot规则链] 历史存储写时序失败 sn=%s err=%v", ctx.msg.Sn, err)
	}
	return allTargets(g, node.ID)
}

// rcPush 大屏推送节点：把流经的最新属性广播到设备所属区域的地图大屏房间。
func rcPush(ctx *ruleCtx, node *iot_schema.RuleNode, g *parsedGraph) []string {
	online.BroadcastTelemetry(ctx.msg, ctx.values)
	return allTargets(g, node.ID)
}

// rcNotify 通知节点：站内/邮件/webhook（对齐审批流 NotifyTask，后续可扩 sms 等）。
// 邮件与 webhook 走异步队列，避免外部依赖阻塞报文处理；单节点只配一种类型。
func rcNotify(ctx *ruleCtx, node *iot_schema.RuleNode, g *parsedGraph) []string {
	p := node.Properties.Notify
	if p == nil || p.ServiceType == "" {
		return allTargets(g, node.ID)
	}
	content := p.ServiceContent
	if content == "" {
		content = fmt.Sprintf("设备 %s 触发规则链通知", ctx.msg.Sn)
	}
	switch p.ServiceType {
	case "site":
		for _, rid := range p.ReceiverId {
			if rid == "" {
				continue
			}
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
			return allTargets(g, node.ID)
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
			return allTargets(g, node.ID)
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
	return allTargets(g, node.ID)
}

// dedupNonEmpty 去重且剔除空白的字符串列表。
func dedupNonEmpty(in []string) []string {
	seen := map[string]struct{}{}
	out := make([]string, 0, len(in))
	for _, s := range in {
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}
		if _, ok := seen[s]; ok {
			continue
		}
		seen[s] = struct{}{}
		out = append(out, s)
	}
	return out
}
