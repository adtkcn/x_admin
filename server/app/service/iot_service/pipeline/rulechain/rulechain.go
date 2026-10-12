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

// 可视化规则链执行引擎：逐报文沿规则链树执行各节点。
//
// 数据流：产品 rule_chain_id → service 缓存的执行树原文（tree_data）→ 本包再缓存一层“解析好的树”
// （原文变了或超 TTL 才重解析），每条报文只走内存指针、不做 JSON 解析与库查询。
//
// 拓扑：链在语义上是一棵树——触发器为唯一根，除并行节点外每节点至多一条后继。这棵树由前端
// 从 LogicFlow 图展开后直传（graph_data 仅用于绘图回填），后端不再做图→树展开，直接反序列化下钻。
//
// 遍历：节点处理器只回答“本分支是否继续下钻”——过滤 false / 解码产出空置 halt 短路整条报文，
// 布尔 false 仅剪本分支；并行节点的 children 各持深拷贝的属性负载，分支间互不干扰。
//
// 依赖：本包直接 import online/alert/store 三个叶子子包，它们不反向依赖本包，
// 构成 online ← alert ← rulechain ← pipeline 的单向 DAG，无需注入式解耦。

// parsedTree 解析并缓存好的规则链树（直接复用前端提交的 RuleTreeNode）
type parsedTree struct {
	raw      string
	parsedAt time.Time
	root     *iot_schema.RuleTreeNode
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
	raw := iot_service.RuleChainService.CachedTreeRaw(chainId)
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

// parseTree 反序列化前端提交的执行树 JSON 为 RuleTreeNode（JSON 本身即树形、无共享引用不会成环）；
// 解析失败或缺触发器根节点返回 nil（保存时已校过根类型，此处只兜住明显畸形）。
func parseTree(raw string) *parsedTree {
	var tn iot_schema.RuleTreeNode
	if err := json.Unmarshal([]byte(raw), &tn); err != nil {
		core.Logger.Warnf("[iot规则链] 执行树 JSON 解析失败 err=%v", err)
		return nil
	}
	if tn.Type != iot_schema.RcNodeTrigger {
		core.Logger.Warnf("[iot规则链] 执行树缺少触发器根节点，丢弃")
		return nil
	}
	return &parsedTree{raw: raw, parsedAt: time.Now(), root: &tn}
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
	walk(&ruleCtx{msg: msg, model: model, rcShared: &rcShared{}, values: msg.Values}, t.root)
	return nil
}

// walk 深度优先下钻规则链树：节点处理器决定本分支是否继续，halt 后全链短路。
// 后继由 node.Children 直接表达（非并行为 0/1 个，并行为 N 个）。
func walk(ctx *ruleCtx, node *iot_schema.RuleTreeNode) {
	if ctx.halt {
		return
	}
	if !applyRcNode(ctx, node) {
		return
	}

	// 并行分叉（唯一多出口）给每条分支一份独立深拷贝的上下文与属性负载，杜绝分支间相互改写；
	// 单线（≤ 1 后继）无兄弟分支，直接沿用当前上下文
	for _, child := range node.Children {
		if ctx.halt {
			return
		}
		walk(ctx.branch(len(node.Children) > 1), child)
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
