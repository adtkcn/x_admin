package core

import (
	"log"

	"x_admin/config"
	"x_admin/core/queue"
	"x_admin/core/queue_delay"
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

// QueueDelay 全局延迟投递队列实例，与 Queue 独立，基于 Redis ZSET 到期轮询。
// 业务层用于定时/延迟消息（如订单超时关闭、延时通知）。
var QueueDelay = func() queue_delay.DelayedQueue {
	q, err := queue_delay.New(queue_delay.Config{
		Backend: queue_delay.BackendRedis,
		Prefix:  config.RedisConfig.RedisPrefix + "queue_delay:",
	}, Redis)
	if err != nil {
		log.Fatalf("init default queue_delay error: %v", err)
	}
	return q
}()
