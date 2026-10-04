package pipeline

import (
	"sync"
	"x_admin/app/model/iot_model"
	"x_admin/app/schema/iot_schema"
	"x_admin/core"

	"github.com/dop251/goja"
)

// goja 脚本执行（流水线前三级的脚本部分）：产品可配置三段 JS——
//   - decode(raw)   → 物模型属性对象（解码级）
//   - filter(values) → 布尔，true 放行、false 丢弃本条上报（过滤级）
//   - alert(values)  → 告警级别数字（0=不告警，1/2/3=轻微/严重/故障），>0 即异常（告警级）
//
// 解码入参是原始报文串，过滤/告警入参是解码后的属性对象。约定各段脚本
// 分别定义同名全局函数；未配置对应脚本时该级为“透传/默认值”，不阻断链路。
//
// 编译一次、每条报文新建 Runtime：goja.Runtime 不是并发安全的，跨 goroutine 复用需加锁、
// 反而把并发流水线串成串行；Program 复用省掉重复编译，新建 Runtime 的开销可接受。
// 编译缓存以「产品ID:阶段」为键，脚本原文变更时（产品编辑后缓存失效→回源拿到新脚本）按内容比对重编译。

// compiledScript 某产品某阶段当前编译好的程序，raw 变了就重编译
type compiledScript struct {
	raw  string
	prog *goja.Program
}

var (
	scriptMu  sync.RWMutex
	scriptMap = map[string]*compiledScript{}
)

// decodeMessage 执行产品解码脚本；脚本为空时原样返回上报 values。
// 返回 nil 表示本条报文丢弃（编译/执行失败），交由上层按空值短路。
func decodeMessage(msg iot_schema.IngestMsg, script string) map[string]any {
	if script == "" {
		return msg.Values
	}
	prog := compile(msg.ProductId, "decode", script)
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
		core.Logger.Warnf("[iot处理] 解码脚本返回值不是对象 sn=%s", msg.Sn)
		return nil
	}
}

// passFilter 执行数据过滤脚本：返回 true 放行进入快照/存储/广播，false 丢弃本条。
// 脚本为空默认放行；编译/执行/返回非布尔都按 false（丢弃）处理并记日志——
// 过滤是显式闸门，脚本没跑起来时宁可不落地，避免半成品数据入库。
func passFilter(msg iot_schema.IngestMsg, values map[string]any, script string) bool {
	if script == "" {
		return true
	}
	return boolScript(msg, "filter", values, script)
}

// alertSeverity 执行数据告警脚本：返回告警级别（0=不告警/恢复，>0=对应级别）。
// 脚本为空默认 0；兼容旧布尔写法：true → 默认级别、false → 0；非布尔/非数字也按 0。
func alertSeverity(msg iot_schema.IngestMsg, values map[string]any, script string) uint8 {
	if script == "" {
		return 0
	}
	prog := compile(msg.ProductId, "alert", script)
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
		core.Logger.Warnf("[iot处理] alert 脚本返回值不是级别数字/布尔 sn=%s", msg.Sn)
		return 0
	}
}

// boolScript 跑一次返回布尔的脚本（filter 用），失败一律降级为 false。
func boolScript(msg iot_schema.IngestMsg, fnName string, values map[string]any, script string) bool {
	prog := compile(msg.ProductId, fnName, script)
	if prog == nil {
		return false
	}
	val, ok := callScript(prog, fnName, values, msg.Sn)
	if !ok {
		return false
	}
	b, isBool := val.Export().(bool)
	if !isBool {
		core.Logger.Warnf("[iot处理] %s 脚本返回值不是布尔 sn=%s", fnName, msg.Sn)
		return false
	}
	return b
}

// compile 取「产品:阶段」编译好的程序，脚本原文与缓存不一致时重新编译
func compile(productId, stage, script string) *goja.Program {
	key := productId + ":" + stage
	scriptMu.RLock()
	entry := scriptMap[key]
	scriptMu.RUnlock()
	if entry != nil && entry.raw == script {
		return entry.prog
	}

	prog, err := goja.Compile(stage+".js", script, false)
	if err != nil {
		core.Logger.Errorf("[iot处理] %s 脚本编译失败 productId=%s err=%v", stage, productId, err)
		return nil
	}
	scriptMu.Lock()
	scriptMap[key] = &compiledScript{raw: script, prog: prog}
	scriptMu.Unlock()
	return prog
}

// callScript 在独立 Runtime 里跑一次程序并调用全局函数 fn(arg)。
// 脚本可能抛异常或死循环触发栈溢出，统一 recover 兜住，单条报文脚本失败不拖垮 worker。
func callScript(prog *goja.Program, fnName string, arg any, sn string) (val goja.Value, ok bool) {
	defer func() {
		if r := recover(); r != nil {
			core.Logger.Errorf("[iot处理] %s 脚本 panic sn=%s err=%v", fnName, sn, r)
			val, ok = nil, false
		}
	}()

	vm := goja.New()
	if _, err := vm.RunProgram(prog); err != nil {
		core.Logger.Errorf("[iot处理] %s 脚本执行失败 sn=%s err=%v", fnName, sn, err)
		return nil, false
	}
	fn, isFn := goja.AssertFunction(vm.Get(fnName))
	if !isFn {
		core.Logger.Errorf("[iot处理] %s 脚本未定义全局函数 %s sn=%s", fnName, fnName, sn)
		return nil, false
	}
	v, err := fn(goja.Undefined(), vm.ToValue(arg))
	if err != nil {
		core.Logger.Errorf("[iot处理] %s 脚本调用失败 sn=%s err=%v", fnName, sn, err)
		return nil, false
	}
	return v, true
}
