package core

import (
	"log"

	"x_admin/config"
	"x_admin/core/pubsub"
	"x_admin/core/ws"
)

// Ws WebSocket 管理器,作为基础设施方便全局使用
var Ws = ws.NewManager()

func init() {
	// 事件总线使用 Redis 后端（跨实例广播）；房间管理由 ws 内部本地内存维护
	var rdb = Redis
	var prefix = config.RedisConfig.RedisPrefix

	em, err := pubsub.New(pubsub.Config{
		Backend: pubsub.BackendRedis,
		Prefix:  prefix + "ws:", // 事件频道前缀
	}, Redis)
	if err != nil {
		log.Fatalf("[ws] init pubsub error: %v", err)
	}

	if !config.AppConfig.ClusterMode {
		rdb = nil // 单机模式不使用 Redis 在线计数
	}

	Ws.Init(em, rdb, prefix)
	go Ws.Start()
}
