package core

import (
	"x_admin/config"
	"x_admin/core/ws"
)

// Ws WebSocket 管理器,作为基础设施方便全局使用
var Ws = ws.NewManager()

func init() {
	// 根据集群模式配置选择 PubSub 实现
	// 房间管理两种模式共用本地内存实现，跨实例通信完全通过 PubSub 完成
	var ps ws.PubSub
	var rdb = Redis
	var prefix = config.RedisConfig.RedisPrefix
	if config.AppConfig.ClusterMode == true {
		// 集群模式：Redis Pub/Sub
		channel := prefix + "ws:broadcast"
		ps = ws.NewRedisPubSub(Redis, channel)
	} else {
		// 单机模式：本地 channel
		ps = ws.NewLocalPubSub()
		rdb = nil // 单机模式不使用 Redis 在线计数
	}

	Ws.Init(ps, rdb, prefix)
	go Ws.Start()
}
