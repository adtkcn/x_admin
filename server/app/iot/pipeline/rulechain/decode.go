package rulechain

import (
	"sync"
	"x_admin/app/model/iot_model"
	"x_admin/app/schema/iot_schema"
	"x_admin/core"

	"github.com/dop251/goja"
)

// goja 脚本执行原语（规则链的解码/过滤/告警节点共用），函数均以 sn 作为第 2 个参数：
//   - decode(raw, sn)    → 物模型属性对象（解码节点）
//   - filter(values, sn) → 布尔，true 放行、false 丢弃本条（过滤节点）
//   - alert(values, sn)  → 告警级别数字（0=不告警，1/2/3=轻微/严重/故障），>0 即异常（告警节点）
//
// 各段脚本约定定义同名全局函数；脚本为空时该级为“透传/默认值”，不阻断链路。
//
// 编译一次、每条报文新建 Runtime：goja.Runtime 不是并发安全的，跨 goroutine 复用需加锁、
// 反而把并发流水线串成串行；Program 复用省掉重复编译，新建 Runtime 的开销可接受。
// 编译缓存以调用方传入的 cacheKey 为键（规则链下用「链:节点ID」——脚本内嵌在图节点里，
// 归属是链而非产品，多产品共用同一条链时天然复用同一份编译结果），脚本原文变更时按内容比对重编译。

// compiledScript 某缓存键当前编译好的程序，raw 变了就重编译
type compiledScript struct {
	raw  string
	prog *goja.Program
}

var (
	scriptMu  sync.RWMutex
	scriptMap = map[string]*compiledScript{}
)

// scriptCacheKey 节点脚本的编译缓存键：规则链ID+节点ID。nodeId 在链内已唯一（节点类型/脚本角色
// 随之确定，无需再拼节点类型）；链ID 做前缀让多产品共用同一条链时复用同一份编译结果。
func scriptCacheKey(chainId, nodeId string) string {
	return chainId + ":" + nodeId
}

// execDecode 执行解码脚本：script 为空时透传上报 values；返回 nil 表示本条丢弃。
func execDecode(msg iot_schema.IngestMsg, cacheKey, script string) map[string]any {
	if script == "" {
		return msg.Values
	}
	prog := compile(cacheKey, "decode", script)
	if prog == nil {
		return nil
	}
	val, ok := callScript(prog, "decode", msg.Raw, msg.Sn)
	if !ok {
		return nil
	}
	switch m := val.Export().(type) {
	case map[string]any:
		return m
	default:
		core.Logger.Warnf("[iot规则链] 解码脚本返回值不是对象 sn=%s", msg.Sn)
		return nil
	}
}

// execFilter 执行过滤脚本：script 为空默认放行；返回 true 放行，false 丢弃。
// 编译/执行/返回非布尔都按 false（丢弃）处理——过滤是显式闸门，脚本没跑起来宁可不落地。
func execFilter(msg iot_schema.IngestMsg, cacheKey, script string, values map[string]any) bool {
	if script == "" {
		return true
	}
	prog := compile(cacheKey, "filter", script)
	if prog == nil {
		return false
	}
	val, ok := callScript(prog, "filter", values, msg.Sn)
	if !ok {
		return false
	}
	b, isBool := val.Export().(bool)
	if !isBool {
		core.Logger.Warnf("[iot规则链] filter 脚本返回值不是布尔 sn=%s", msg.Sn)
		return false
	}
	return b
}

// execAlertSeverity 执行告警脚本：script 为空返回 0；兼容旧布尔写法（true→默认级别）。
func execAlertSeverity(msg iot_schema.IngestMsg, cacheKey, script string, values map[string]any) uint8 {
	if script == "" {
		return 0
	}
	prog := compile(cacheKey, "alert", script)
	if prog == nil {
		return 0
	}
	val, ok := callScript(prog, "alert", values, msg.Sn)
	if !ok {
		return 0
	}
	switch v := val.Export().(type) {
	case bool:
		if v {
			return iot_model.DefaultAlertSeverity
		}
		return 0
	case float64:
		if v <= 0 {
			return 0
		}
		return uint8(v)
	default:
		core.Logger.Warnf("[iot规则链] alert 脚本返回值不是级别数字/布尔 sn=%s", msg.Sn)
		return 0
	}
}

// compile 取缓存键对应的编译程序，脚本原文与缓存不一致时重新编译
func compile(cacheKey, stage, script string) *goja.Program {
	key := cacheKey + ":" + stage
	scriptMu.RLock()
	entry := scriptMap[key]
	scriptMu.RUnlock()
	if entry != nil && entry.raw == script {
		return entry.prog
	}

	prog, err := goja.Compile(stage+".js", script, false)
	if err != nil {
		core.Logger.Errorf("[iot规则链] %s 脚本编译失败 key=%s err=%v", stage, key, err)
		return nil
	}
	scriptMu.Lock()
	scriptMap[key] = &compiledScript{raw: script, prog: prog}
	scriptMu.Unlock()
	return prog
}

// callScript 在独立 Runtime 里跑一次程序并调用全局函数 fn(arg, sn)。
// 脚本可能抛异常或死循环触发栈溢出，统一 recover 兜住，单条报文脚本失败不拖垮 worker。
func callScript(prog *goja.Program, fnName string, arg any, sn string) (val goja.Value, ok bool) {
	defer func() {
		if r := recover(); r != nil {
			core.Logger.Errorf("[iot规则链] %s 脚本 panic sn=%s err=%v", fnName, sn, r)
			val, ok = nil, false
		}
	}()

	vm := goja.New()
	if _, err := vm.RunProgram(prog); err != nil {
		core.Logger.Errorf("[iot规则链] %s 脚本执行失败 sn=%s err=%v", fnName, sn, err)
		return nil, false
	}
	fn, isFn := goja.AssertFunction(vm.Get(fnName))
	if !isFn {
		core.Logger.Errorf("[iot规则链] %s 脚本未定义全局函数 %s sn=%s", fnName, fnName, sn)
		return nil, false
	}
	v, err := fn(goja.Undefined(), vm.ToValue(arg), vm.ToValue(sn))
	if err != nil {
		core.Logger.Errorf("[iot规则链] %s 脚本调用失败 sn=%s err=%v", fnName, sn, err)
		return nil, false
	}
	return v, true
}
