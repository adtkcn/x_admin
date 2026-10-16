package core

import (
	"log"

	"x_admin/config"
	"x_admin/core/ordered_queue"
)

// OrderedQueue 全局有序队列：按分片键（如 deviceId）哈希到 N 个分片，每片全局只有一个消费者，
// 同键消息严格有序、跨键并行。消费者靠 Redis 租约认领分片，多实例自动分摊、
// 实例崩溃由 TTL 到期后其他实例接管，部署侧无需任何配置。
//
// 与 Queue 的分工：Queue 是竞争消费（只看吞吐、不管顺序），OrderedQueue 是分片消费（同键有序）。
// 设备报文处理走本队列，其余异步任务仍走 Queue。
var OrderedQueue = func() ordered_queue.OrderedQueue {
	q, err := ordered_queue.New(ordered_queue.Config{
		Backend: ordered_queue.BackendRedis,
		Prefix:  config.RedisConfig.RedisPrefix + "oqueue:",
	}, Redis)
	if err != nil {
		log.Fatalf("init ordered queue error: %v", err)
	}
	return q
}()
