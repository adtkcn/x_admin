package iot_service

import (
	"time"

	"x_admin/config"
	"x_admin/core"
)

// IoT 域本地缓存 key 前缀与跨实例失效订阅。
// 缓存实例本身在 core.LocalCache；此处只做键命名空间与 pubsub 接线。

const (
	BusEvtProduct = "product" // payload=产品 ID（空 / `*` 全清）
	BusEvtChain   = "chain"   // payload=规则链 ID（空 / `*` 全清）

	localProductPrefix = "iot:product:"
	localChainPrefix   = "iot:rulechain:"
)

// localTTL 本域兜底窗口，远小于 Redis 层 ThingCacheTTL；跨实例广播未达时靠它过期回源。
var localTTL = time.Duration(config.IotConfig.LocalCacheTTL) * time.Second

func localProductKey(id string) string { return localProductPrefix + id }
func localChainKey(id string) string   { return localChainPrefix + id }

func init() {
	if core.IotBus == nil {
		return
	}
	core.IotBus.On(BusEvtProduct, func(p []byte) {
		invalidateLocal(string(p), localProductPrefix, localProductKey)
	})
	core.IotBus.On(BusEvtChain, func(p []byte) {
		invalidateLocal(string(p), localChainPrefix, localChainKey)
	})
}

// invalidateLocal 单键或全清（id 空 / `*` 视作按前缀清）。
func invalidateLocal(id, prefix string, keyFn func(string) string) {
	if id == "" || id == "*" {
		core.CacheDeletePrefix(prefix)
		return
	}
	core.CacheDelete(keyFn(id))
}

// emitInvalidate 广播失效；IotBus 为 nil（测试环境）时静默。
func emitInvalidate(evt, id string) {
	if core.IotBus == nil {
		return
	}
	if err := core.IotBus.Emit(evt, []byte(id)); err != nil {
		core.Logger.Errorf("[iot-bus] 广播失效失败 evt=%s id=%s err=%v", evt, id, err)
	}
}
