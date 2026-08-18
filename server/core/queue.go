package core

import (
	"log"

	"x_admin/config"
	"x_admin/core/queue"
)

// Queue 基础全局队列实例，基于 Redis 配置在包初始化时自动创建。
// 业务层（如 app/task）直接复用，无需自行构造 Queue。
var Queue = func() queue.Queue {
	q, err := queue.New(queue.Config{
		Backend: queue.BackendRedis,
		Prefix:  config.RedisConfig.RedisPrefix + "queue:",
	}, Redis)
	if err != nil {
		log.Fatalf("init default queue error: %v", err)
	}
	return q
}()
