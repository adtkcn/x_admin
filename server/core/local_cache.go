package core

import (
	"strings"
	"time"

	"github.com/jellydator/ttlcache/v3"
)

// 项目统一本地缓存（进程内 TTL，非 Redis，无网络、无持久化）。
//
// 单实例 + 命名空间 key：所有业务共用一个 LocalCache，key 形如
// "<域>:<子域>:<标识>"（例 `iot:product:123`、`iot:rulechain:abc`），
// 避免跨模块键冲突。域内可 `CacheDeletePrefix` 做前缀清理。
//
// 类型：值统一存 any；读写走包级泛型函数 `CacheGet[T](key)` 拿回原类型，
// miss 与"值类型不匹配"都返回零值 + false。之所以用包级函数而不是方法：
// Go 方法不能自带类型参数。
//
// 生命周期：本文件 init 中启动后台清理协程；main 优雅关闭前调 `StopLocalCache` 收尾。
//
// 与 util.CacheUtil 的分工：后者是 Redis 上的 1h KV 简易封装；本文件是进程内热路径缓存。
//
// 一致性：多实例部署要即时失效的，业务自行经 pubsub 广播 → 订阅端 Delete/DeletePrefix
// （参考 iot_service/local_cache.go）。全局默认 TTL 仅作为「广播未达 + 业务忘设 ttl」时的兜底窗口。

// LocalCache 全局本地缓存实例；一般通过 CacheGet/CacheSet/CacheDelete/CacheDeletePrefix
// 访问，特殊场景（如遍历 Keys/Len）可直接用底层 ttlcache API。
var LocalCache = ttlcache.New(
	ttlcache.WithTTL[string, any](1*time.Minute),
	ttlcache.WithDisableTouchOnHit[string, any](),
)

func init() {
	go LocalCache.Start()
}

// CacheGet 泛型读取：miss / 过期 / 类型不匹配均返回零值 + false。
func CacheGet[T any](key string) (T, bool) {
	var zero T
	item := LocalCache.Get(key)
	if item == nil {
		return zero, false
	}
	v, ok := item.Value().(T)
	if !ok {
		return zero, false
	}
	return v, true
}

// CacheSet 写入：ttl<=0 使用 LocalCache 默认（5 分钟），否则用传入 ttl。
// 业务侧推荐显式传 `IotConfig.LocalCacheTTL` 等配置化时长，避免依赖默认。
func CacheSet(key string, val any, ttl time.Duration) {
	if ttl <= 0 {
		ttl = ttlcache.DefaultTTL
	}
	LocalCache.Set(key, val, ttl)
}

// CacheDelete 删除单键；不存在时 no-op。
func CacheDelete(key string) { LocalCache.Delete(key) }

// CacheDeletePrefix 删除全部以 prefix 起始的键（供「某域全清」场景，如 pubsub 收到 `*`）。
// 先 Keys() 快照（含内部互斥）再逐个 Delete，避免在 Delete 持锁期间做范围遍历。
func CacheDeletePrefix(prefix string) {
	for _, k := range LocalCache.Keys() {
		if strings.HasPrefix(k, prefix) {
			LocalCache.Delete(k)
		}
	}
}

// StopLocalCache 停后台清理协程；幂等由 ttlcache 保证。请在进程优雅关闭时（Redis 关前）调用一次。
func StopLocalCache() { LocalCache.Stop() }
