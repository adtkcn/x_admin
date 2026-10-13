package core

import (
	"log"

	"x_admin/config"
	"x_admin/core/pubsub"
)

// IotBus IoT 侧事件总线（跨实例广播）：目前用于「产品模型 / 规则链执行树」本地缓存失效。
//
// 与 Ws 使用独立频道前缀（{RedisPrefix}iot:bus:），职责隔离、事件类型互不重叠。
// 消费方在 iot_service 包 init 中订阅；生产方在管理端 CRUD 后 Emit。
// 单机部署时也走 Redis 广播（多实例语义一致，代码分支不区分）。
var IotBus pubsub.Emitter

func init() {
	em, err := pubsub.New(pubsub.Config{
		Backend: pubsub.BackendRedis,
		Prefix:  config.RedisConfig.RedisPrefix + "iot:bus:",
	}, Redis)
	if err != nil {
		log.Fatalf("[iot-bus] init pubsub error: %v", err)
	}
	IotBus = em
}
