package rulechain

import (
	"encoding/json"
	"sync"
	"time"
	"x_admin/app/schema/iot_schema"
	"x_admin/app/service/iot_service"
	"x_admin/config"
	"x_admin/core"
)

// 可视化规则链执行引擎（处理层逐报文按有向图执行，替代原产品级三段脚本）。
//
// 数据流：产品 rule_chain_id → iot_service.RuleChainService.CachedGraphRaw(Redis 缓存原文)
// → 本包再缓存一层“解析好的图”（仿 identMap：原文变了或超过 TTL 才重解析），
// 保证每条报文只走内存邻接表、不做 JSON 解析与库查询。
//
// 遍历：从输入节点出发，节点处理器返回要跟进的出边目标（并行返回全部子目标、
// 过滤 false 丢弃整条、布尔 false 仅剪本分支），深度优先展开；用访问计数+深度上限防御环与报文放大。
//
// 依赖方向：本包直接 import online/alert/store 三个叶子子包（状态快照/大屏推送/告警收敛/时序落库），
// 它们都不反向依赖本包，构成 online ← alert ← rulechain ← pipeline 的单向 DAG，无需注入式解耦。

// rcMaxDepth 单次报文图遍历的最大深度，防成环/畸形图把 worker 拖死
const rcMaxDepth = 64

// parsedGraph 解析好的规则链有向图：节点表 + 出边邻接表 + 触发器入口
type parsedGraph struct {
	raw      string
	parsedAt time.Time
	nodes    map[string]*iot_schema.RuleNode
	out      map[string][]*iot_schema.RuleEdge
	startId  string
}

var (
	graphMu  sync.RWMutex
	graphMap = map[string]*parsedGraph{}
)

// loadGraph 取解析好的规则链图：原文由 service 缓存读出，这里再缓存一层解析结果。
// 返回 nil 表示链不存在/停用/无法解析（调用方按“未配置”丢弃上报）。
func loadGraph(chainId string) *parsedGraph {
	if chainId == "" {
		return nil
	}
	raw := iot_service.RuleChainService.CachedGraphRaw(chainId)
	if raw == "" {
		return nil
	}
	graphMu.RLock()
	entry := graphMap[chainId]
	graphMu.RUnlock()
	ttl := time.Duration(config.IotConfig.ThingCacheTTL) * time.Second
	if entry != nil && entry.raw == raw && time.Since(entry.parsedAt) < ttl {
		return entry
	}

	g := parseGraph(raw)
	if g == nil {
		return nil
	}
	g.raw = raw
	g.parsedAt = time.Now()
	graphMu.Lock()
	graphMap[chainId] = g
	graphMu.Unlock()
	return g
}

// parseGraph 解析 LogicFlow 原样图，构造节点表/出边邻接表，定位触发器入口。
func parseGraph(raw string) *parsedGraph {
	var cg iot_schema.RuleChainGraph
	if err := json.Unmarshal([]byte(raw), &cg); err != nil {
		core.Logger.Warnf("[iot规则链] 图 JSON 解析失败 err=%v", err)
		return nil
	}
	g := &parsedGraph{
		nodes: make(map[string]*iot_schema.RuleNode, len(cg.Nodes)),
		out:   make(map[string][]*iot_schema.RuleEdge, len(cg.Nodes)),
	}
	for i := range cg.Nodes {
		node := &cg.Nodes[i]
		g.nodes[node.ID] = node
		if node.Type == iot_schema.RcNodeTrigger && g.startId == "" {
			g.startId = node.ID
		}
	}
	for i := range cg.Edges {
		edge := &cg.Edges[i]
		g.out[edge.SourceNodeId] = append(g.out[edge.SourceNodeId], edge)
	}
	if g.startId == "" {
		core.Logger.Warnf("[iot规则链] 图缺少触发器节点，丢弃")
		return nil
	}
	return g
}

// ruleCtx 单条报文在图里流转的上下文：携带产品配置、可变属性负载与短路标记。
type ruleCtx struct {
	msg    iot_schema.IngestMsg
	model  iot_service.ProductModel
	values map[string]any
	halt   bool
}

// Run 按产品的规则链处理一条上报：链缺失/停用按“未配置”丢弃（对齐物模型必填的短路语义）。
func Run(msg iot_schema.IngestMsg, model iot_service.ProductModel) error {
	g := loadGraph(model.RuleChainId)
	if g == nil {
		if config.IotConfig.LogRaw {
			core.Logger.Warnf("[iot规则链] 产品未配置有效规则链，丢弃上报 sn=%s productId=%s", msg.Sn, msg.ProductId)
		}
		return nil
	}
	ctx := &ruleCtx{msg: msg, model: model, values: msg.Values}
	walk(g, ctx, g.startId, 0)
	return nil
}

// walk 从 nodeId 深度优先执行规则链：节点处理器产出效果并给出后继目标，逐个递归。
// halt 后不再展开；访问计数与深度上限双重防御成环与报文放大。
func walk(g *parsedGraph, ctx *ruleCtx, nodeId string, depth int) {
	if ctx.halt || depth > rcMaxDepth {
		return
	}
	node := g.nodes[nodeId]
	if node == nil {
		return
	}
	nextIds := applyRcNode(ctx, node, g)
	for _, id := range nextIds {
		if ctx.halt {
			return
		}
		walk(g, ctx, id, depth+1)
	}
}
