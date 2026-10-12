package rulechain

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
	"x_admin/app/schema/iot_schema"
	"x_admin/app/service/iot_service"
	"x_admin/config"
	"x_admin/core"
)

// 可视化规则链执行引擎：逐报文沿规则链树执行各节点。
//
// 数据流：产品 rule_chain_id → service 缓存的图原文 → 本包再缓存一层“解析好的树”
// （原文变了或超 TTL 才重解析），每条报文只走内存指针、不做 JSON 解析与库查询。
//
// 拓扑：链在语义上是树——触发器为唯一根，除并行节点外每节点至多一条出边（保存时由
// iot_service.checkGraph 约束），故解析阶段就把 LogicFlow 邻接图展开成根树，执行阶段直接下钻 children。
//
// 遍历：节点处理器只回答“本分支是否继续下钻”——过滤 false / 解码产出空置 halt 短路整条报文，
// 布尔 false 仅剪本分支；并行节点的 children 各持深拷贝的属性负载，分支间互不干扰。
//
// 依赖：本包直接 import online/alert/store 三个叶子子包，它们不反向依赖本包，
// 构成 online ← alert ← rulechain ← pipeline 的单向 DAG，无需注入式解耦。

const (
	// 两项上限都在建树阶段校验（防御前置），执行阶段只走指针、不逐层判深度
	rcMaxDepth = 64   // 规则链树最大深度，超限即判畸形链
	rcMaxSteps = 4096 // 建树允许的最多访问次数：LogicFlow 允许任意连线，遇报文放大时展开会失控
)

// rcStep 规则链树节点：LogicFlow 节点 + 已解析的后继（0 或 1 个，并行节点为 N 个）
type rcStep struct {
	node     *iot_schema.RuleNode
	children []*rcStep
}

// parsedTree 解析并缓存好的规则链树
type parsedTree struct {
	raw      string
	parsedAt time.Time
	root     *rcStep
}

var (
	treeMu  sync.RWMutex
	treeMap = map[string]*parsedTree{}
)

// loadTree 取解析好的规则链树（service 缓存原文，这里再缓存一层解析结果）；
// 返回 nil 表示链不存在/停用/无法解析（调用方按“未配置”丢弃上报）。
func loadTree(chainId string) *parsedTree {
	if chainId == "" {
		return nil
	}
	raw := iot_service.RuleChainService.CachedGraphRaw(chainId)
	if raw == "" {
		return nil
	}
	treeMu.RLock()
	entry := treeMap[chainId]
	treeMu.RUnlock()
	ttl := time.Duration(config.IotConfig.ThingCacheTTL) * time.Second
	if entry != nil && entry.raw == raw && time.Since(entry.parsedAt) < ttl {
		return entry
	}

	t := parseTree(raw)
	if t == nil {
		return nil
	}
	treeMu.Lock()
	treeMap[chainId] = t
	treeMu.Unlock()
	return t
}

// parseTree 解析 LogicFlow 原样图并展开成以触发器为根的树；缺触发器/成环/超限均返回 nil。
func parseTree(raw string) *parsedTree {
	var cg iot_schema.RuleChainGraph
	if err := json.Unmarshal([]byte(raw), &cg); err != nil {
		core.Logger.Warnf("[iot规则链] 图 JSON 解析失败 err=%v", err)
		return nil
	}
	nodes := make(map[string]*iot_schema.RuleNode, len(cg.Nodes))
	out := make(map[string][]*iot_schema.RuleEdge, len(cg.Nodes))
	var startId string
	for i := range cg.Nodes {
		node := &cg.Nodes[i]
		nodes[node.ID] = node
		if node.Type == iot_schema.RcNodeTrigger && startId == "" {
			startId = node.ID
		}
	}
	if startId == "" {
		core.Logger.Warnf("[iot规则链] 图缺少触发器节点，丢弃")
		return nil
	}
	for i := range cg.Edges {
		edge := &cg.Edges[i]
		out[edge.SourceNodeId] = append(out[edge.SourceNodeId], edge)
	}

	b := &treeBuilder{nodes: nodes, out: out, onPath: map[string]bool{}, budget: rcMaxSteps}
	root := b.expand(startId)
	if root == nil {
		core.Logger.Warnf("[iot规则链] 规则链树解析失败 nodes=%d reason=%s", len(cg.Nodes), b.bad)
		return nil
	}
	return &parsedTree{raw: raw, parsedAt: time.Now(), root: root}
}

// treeBuilder 邻接图 → 根树的展开器：onPath 标记「根→本节点」当前路径以防环，budget 防报文放大
type treeBuilder struct {
	nodes  map[string]*iot_schema.RuleNode
	out    map[string][]*iot_schema.RuleEdge
	onPath map[string]bool
	budget int
	bad    string // 首个失败原因，非空即判定畸形链
}

// fail 记录首个失败原因：递归沿路短路，不被后续错误覆盖
func (b *treeBuilder) fail(format string, args ...any) {
	if b.bad == "" {
		b.bad = fmt.Sprintf(format, args...)
	}
}

// expand 沿出边递归展开子树。onPath 只标记当前路径（返回时回溯剔除），
// 因此同一节点被多条路径引用（扇入）会展开成多份副本，与“每条路径各执行一次”的语义一致。
func (b *treeBuilder) expand(id string) *rcStep {
	if b.bad != "" {
		return nil
	}
	if b.budget <= 0 {
		b.fail("超出节点展开上限(报文放大)")
		return nil
	}
	node := b.nodes[id]
	if node == nil {
		b.fail("引用了不存在的节点 id=%s", id)
		return nil
	}
	if b.onPath[id] {
		b.fail("存在环 id=%s", id)
		return nil
	}
	b.budget--
	b.onPath[id] = true
	defer delete(b.onPath, id)
	if len(b.onPath) > rcMaxDepth {
		b.fail("链深度超上限 id=%s", id)
		return nil
	}

	step := &rcStep{node: node}
	for _, e := range b.out[id] {
		child := b.expand(e.TargetNodeId)
		if child == nil {
			return nil // bad 已由子调用记录
		}
		step.children = append(step.children, child)
	}
	return step
}

// rcShared 一条报文在所有分支间共享的可变标记（目前只有全局短路位）。以指针嵌入 ruleCtx：
// branch() 复制结构体时复制的是指针，任一分支置位对所有分支立即可见。
type rcShared struct {
	halt bool
}

// ruleCtx 单条报文在树里流转的上下文：携带产品配置与本分支的属性负载。
type ruleCtx struct {
	msg   iot_schema.IngestMsg
	model iot_service.ProductModel
	*rcShared
	values map[string]any
}

// Run 按产品的规则链处理一条上报：链缺失/停用按“未配置”丢弃。
func Run(msg iot_schema.IngestMsg, model iot_service.ProductModel) error {
	t := loadTree(model.RuleChainId)
	if t == nil {
		if config.IotConfig.LogRaw {
			core.Logger.Warnf("[iot规则链] 产品未配置有效规则链，丢弃上报 sn=%s productId=%s", msg.Sn, msg.ProductId)
		}
		return nil
	}
	walk(&ruleCtx{msg: msg, model: model, rcShared: &rcShared{}, values: msg.Values}, t.root, 0)
	return nil
}

// walk 深度优先下钻规则链树：节点处理器决定本分支是否继续，halt 后全链短路。
// 环/深度/节点数上限已在建树时校验，运行期只需 halt 检查；depth 仅用于日志呈现。
func walk(ctx *ruleCtx, step *rcStep, depth int) {
	if ctx.halt {
		return
	}
	node := step.node
	// logOn := config.IotConfig.LogRuleChain
	// if logOn {
	// 	core.Logger.Infof("进入节点%q %s sn=%s depth=%d  入参=%s",
	// 		rcNodeName(node), node.Type, ctx.msg.Sn, depth, rcValues(ctx.values))
	// }
	cont := applyRcNode(ctx, node)
	// if logOn {
	// 	core.Logger.Infof("离开节点%q %s sn=%s 出参=%s 分支=%d halt=%v",
	// 		rcNodeName(node), node.Type, ctx.msg.Sn, rcValues(ctx.values), len(step.children), ctx.halt)
	// }
	if !cont {
		return
	}

	// 并行分叉（唯一多出口）给每条分支一份独立深拷贝的上下文与属性负载，杜绝分支间相互改写；
	// 单线（≤ 1 后继）无兄弟分支，直接沿用当前上下文
	for _, child := range step.children {
		if ctx.halt {
			return
		}
		walk(ctx.branch(len(step.children) > 1), child, depth+1)
	}
}

// branch 派生流入下一节点的上下文：msg/model 不变，嵌入的 *rcShared 指针使 halt 在所有分支间共享。
func (c *ruleCtx) branch(fork bool) *ruleCtx {
	child := *c
	if fork {
		child.values = deepCloneValues(c.values)
	}
	return &child
}

// deepCloneValues 递归深拷贝属性负载，使并行分支即便被就地改写也不影响兄弟分支。
func deepCloneValues(m map[string]any) map[string]any {
	if m == nil {
		return nil
	}
	out := make(map[string]any, len(m))
	for k, v := range m {
		out[k] = deepCloneAny(v)
	}
	return out
}

// deepCloneAny 递归拷贝单个值：map/slice 新建，其余（string/float/bool/nil）不可变直接返回。
func deepCloneAny(v any) any {
	switch t := v.(type) {
	case map[string]any:
		return deepCloneValues(t)
	case []any:
		cp := make([]any, len(t))
		for i, e := range t {
			cp[i] = deepCloneAny(e)
		}
		return cp
	default:
		return v
	}
}

// rcNodeName 节点可读数名：优先用图上文本标签，缺省回退节点类型。
func rcNodeName(node *iot_schema.RuleNode) string {
	if node.Text.Value != "" {
		return node.Text.Value
	}
	return node.Type
}

// rcValues 属性负载序列化为紧凑 JSON，供日志稳定呈现（map 直接 %v 打印顺序不定）。
func rcValues(values map[string]any) string {
	if values == nil {
		return "null"
	}
	b, err := json.Marshal(values)
	if err != nil {
		return fmt.Sprintf("%v", values)
	}
	return string(b)
}
