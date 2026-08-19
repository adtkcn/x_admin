package queue_delay

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

// Backend 延迟队列后端类型。
type Backend string

const (
	// BackendRedis Redis 后端，基于 ZSET 的到期轮询延迟队列。
	BackendRedis Backend = "redis"
)
const DefaultPrefix = "queue_delay:"

// Config 延迟队列配置，作为 New 工厂函数的唯一输入。
type Config struct {
	// Backend 后端类型，留空时默认 BackendRedis。
	Backend Backend
	// Prefix 队列键前缀，建议带业务/项目前缀以隔离。
	// 队列名 name 最终映射到 Redis 键 {Prefix}{name}（ZSET）。
	Prefix string
}

// New 根据配置创建 DelayedQueue 实例，是唯一推荐的构造入口。
// redisClient 在 Backend == BackendRedis 时需要，传 nil 会返回错误。
//
// 新增后端只需实现 DelayedQueue 接口并在此 switch 中注册，调用方无需改动。
func New(cfg Config, redisClient *redis.Client) (DelayedQueue, error) {
	if cfg.Backend == "" {
		cfg.Backend = BackendRedis
	}
	switch cfg.Backend {
	case BackendRedis:
		if redisClient == nil {
			return nil, fmt.Errorf("queue_delay: redis backend requires a non-nil redis client")
		}
		prefix := cfg.Prefix
		if prefix == "" {
			prefix = DefaultPrefix
		}
		return newRedisDelayedBackend(redisClient, prefix), nil
	default:
		return nil, fmt.Errorf("queue_delay: unknown backend %q", cfg.Backend)
	}
}
