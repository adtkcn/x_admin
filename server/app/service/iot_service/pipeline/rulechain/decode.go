package rulechain

import (
	"math"
	"sync"
	"x_admin/app/schema/iot_schema"
	"x_admin/config"
	"x_admin/core"

	"github.com/dop251/goja"
)

// goja 脚本执行原语（解码/过滤/告警节点共用），均约定定义同名全局函数、以 sn 作第 2 个参数；
// 脚本为空即“透传/默认值”，不阻断链路：
//   - decode(values, sn) → 物模型属性对象（values 入链时为设备初始上报值）
//   - filter(values, sn) → 布尔，true 放行、false 丢弃本条
//   - alert(values, sn)  → 告警级别数字（0=不告警/恢复；危害升序：-1=故障、1=轻微、2=严重，3+ 预留）
//
// 编译一次、每条报文新建 Runtime：goja.Runtime 非并发安全，跨 goroutine 复用要加锁，
// 反而把并发流水线串成串行；Program 复用省掉编译，新建 Runtime 开销可接受。
// 缓存键由调用方传入（规则链下是「链ID:节点ID」——脚本内嵌在图节点里，归属是链而非产品），
// 原文变更时按内容比对重编译。

// compiledScript 某缓存键当前编译好的程序，raw 变了就重编译
type compiledScript struct {
	raw  string
	prog *goja.Program
}

var (
	scriptMu  sync.RWMutex
	scriptMap = map[string]*compiledScript{}
)

// execDecode 执行解码脚本：以设备初始上报值 values 为入参，script 为空时原值透传；返回 nil 表示本条丢弃。
func execDecode(msg iot_schema.IngestMsg, cacheKey, script string, values map[string]any) map[string]any {
	if script == "" {
		return values
	}
	val, ok := runScript("decode", msg, cacheKey, script, values)
	if !ok {
		return nil
	}
	out := val.Export()
	m, isMap := out.(map[string]any)
	if !isMap {
		core.Logger.Warnf("[iot规则链] decode 脚本返回值不是对象 sn=%s type=%T value=%v", msg.Sn, out, out)
		return nil
	}
	return m
}

// execFilter 执行过滤脚本：script 为空默认放行。编译/执行/返回非布尔都按丢弃处理——
// 过滤是显式闸门，脚本没跑起来宁可不落地。
func execFilter(msg iot_schema.IngestMsg, cacheKey, script string, values map[string]any) bool {
	if script == "" {
		return true
	}
	val, ok := runScript("filter", msg, cacheKey, script, values)
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

// execAlertSeverity 执行告警脚本：script 为空或返回 0 都为不告警；只认数字级别，不兼容其它写法。
func execAlertSeverity(msg iot_schema.IngestMsg, cacheKey, script string, values map[string]any) int8 {
	if script == "" {
		return 0
	}
	val, ok := runScript("alert", msg, cacheKey, script, values)
	if !ok {
		return 0
	}

	out := val.Export()
	// 告警是否命中完全取决于这个返回值，联调时需看到它的原始类型与值
	if config.IotConfig.LogRuleChain {
		core.Logger.Infof("[iot规则链] alert 脚本返回 sn=%s type=%T value=%v", msg.Sn, out, out)
	}
	switch v := out.(type) {
	case int64:
		// 0 是唯一的“不告警/恢复”哨兵，负数（故障）是合法级别，不能按 <=0 一律丢掉
		if v == 0 {
			return 0
		}
		// 截断进 int8：库列是 tinyint，脚本乱返回大数会写库失败或回绕成假级别
		if v > math.MaxInt8 {
			v = math.MaxInt8
		}
		if v < math.MinInt8 {
			v = math.MinInt8
		}
		return int8(v)
	default:
		core.Logger.Warnf("[iot规则链] alert 脚本返回值不是级别数字 sn=%s type=%T value=%v", msg.Sn, out, out)
		return 0
	}
}

// runScript 编译并按名调用全局函数 fn(values, sn)；ok=false 表示脚本未跑起来（编译失败/panic/未定义函数），
// 降级语义由调用方决定。
func runScript(stage string, msg iot_schema.IngestMsg, cacheKey, script string, values map[string]any) (goja.Value, bool) {
	prog := compile(cacheKey, stage, script)
	if prog == nil {
		return nil, false
	}
	return callScript(prog, stage, values, msg.Sn)
}

// compile 取缓存键对应的编译程序，脚本原文与缓存不一致时重编译
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
// 脚本可能抛异常或死循环触发栈溢出，统一 recover 兜住，单条报文失败不拖垮 worker。
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
