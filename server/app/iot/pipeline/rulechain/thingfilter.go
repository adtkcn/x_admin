package rulechain

import (
	"encoding/json"
	"sync"
	"time"

	"x_admin/app/schema/iot_schema"
	"x_admin/config"
	"x_admin/core"
)

// 物模型白名单过滤（规则链「物模型过滤」节点）：只放行产品 thing_model 里声明过的属性，
// 防止设备乱报字段把 Redis 快照与时序表撑爆。
//
// 物模型为必填配置：解析不出任何 identifier（未配置或 JSON 里没有 identifier 字段）
// 即视为产品未建模，整条上报丢弃——数据落地前必须有明确的物模型约束。

// identSet 某产品的属性白名单
type identSet struct {
	raw    string // 解析来源的物模型原文，变了就重新解析
	names  map[string]struct{}
	parsed time.Time
}

var (
	identMu  sync.RWMutex
	identMap = map[string]*identSet{}
)

// loadIdentifiers 取产品属性白名单：物模型原文由调用方从缓存读出后传入，
// 这里再缓存一层“解析结果”，避免每条报文都做一次 JSON 解析（原文变了或超过 TTL 才重解析）。
func loadIdentifiers(productId, raw string) map[string]struct{} {
	if raw == "" {
		return nil
	}
	identMu.RLock()
	entry := identMap[productId]
	identMu.RUnlock()
	ttl := time.Duration(config.IotConfig.ThingCacheTTL) * time.Second
	if entry != nil && entry.raw == raw && time.Since(entry.parsed) < ttl {
		return entry.names
	}

	names := parseIdentifiers(raw)
	identMu.Lock()
	identMap[productId] = &identSet{raw: raw, names: names, parsed: time.Now()}
	identMu.Unlock()
	return names
}

// parseIdentifiers 递归收集 identifier/key 字段值，兼容常见物模型写法：
//
//	[{"identifier":"temperature"}, ...]（本项目扁平数组约定）
//	{"properties":[{"identifier":"temperature"}, ...]}
//
// 解析不出任何 identifier 时返回 nil，由调用方按“未建模”丢弃上报。
func parseIdentifiers(raw string) map[string]struct{} {
	var node any
	if err := json.Unmarshal([]byte(raw), &node); err != nil {
		core.Logger.Warnf("[iot处理] 物模型 JSON 解析失败，按未建模处理 err=%v", err)
		return nil
	}
	names := map[string]struct{}{}
	collectIdentifiers(node, names)
	if len(names) == 0 {
		return nil
	}
	return names
}

func collectIdentifiers(node any, out map[string]struct{}) {
	switch t := node.(type) {
	case map[string]any:
		for k, v := range t {
			if k == "identifier" || k == "key" {
				if name, ok := v.(string); ok && name != "" {
					out[name] = struct{}{}
					continue // 已将本对象计为属性，不再深入其元数据（name/dataType 等）
				}
			}
			collectIdentifiers(v, out)
		}
	case []any:
		for _, item := range t {
			collectIdentifiers(item, out)
		}
	}
}

// filterByThingModel 按白名单裁剪属性。物模型为必填：白名单为空（未配置或解析不出属性）
// 时整条丢弃，防止未建模数据污染存储。
func filterByThingModel(msg iot_schema.IngestMsg, values map[string]any, thingRaw string) map[string]any {
	if len(values) == 0 {
		return values
	}
	allow := loadIdentifiers(msg.ProductId, thingRaw)
	if len(allow) == 0 {
		core.Logger.Warnf("[iot处理] 产品未配置有效物模型，丢弃上报 sn=%s productId=%s", msg.Sn, msg.ProductId)
		return nil
	}
	filtered := make(map[string]any, len(values))
	dropped := 0
	for k, v := range values {
		if _, ok := allow[k]; ok {
			filtered[k] = v
		} else {
			dropped++
		}
	}
	if dropped > 0 && config.IotConfig.LogRaw {
		core.Logger.Infof("[iot处理] 丢弃未建模属性 sn=%s count=%d", msg.Sn, dropped)
	}
	return filtered
}
