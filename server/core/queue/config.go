package queue

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

// Backend 队列后端类型。
type Backend string

const (
	// BackendRedis Redis 后端，基于 Redis List 的跨实例工作队列。
	BackendRedis Backend = "redis"
)
const DefaultPrefix = "queue:"

// Config Queue 配置，作为 New 工厂函数的唯一输入。
type Config struct {
	// Backend 后端类型，留空时默认 BackendRedis。
	Backend Backend
	// Prefix 队列键前缀，建议带业务/项目前缀以隔离。
	// 队列名 name 最终映射到 Redis 键 {Prefix}{name}。
	Prefix string
}

// New 根据配置创建 Queue 实例，是唯一推荐的构造入口。
// redisClient 在 Backend == BackendRedis 时需要，传 nil 会返回错误。
//
// 新增后端（NATS / Kafka / 数据库轮询等）只需实现 Queue 接口并在此 switch 中注册，
// 调用方无需改动（保留可扩展性）。
func New(cfg Config, redisClient *redis.Client) (Queue, error) {
	if cfg.Backend == "" {
		cfg.Backend = BackendRedis
	}
	switch cfg.Backend {
	case BackendRedis:
		if redisClient == nil {
			return nil, fmt.Errorf("queue: redis backend requires a non-nil redis client")
		}
		prefix := cfg.Prefix
		if prefix == "" {
			prefix = DefaultPrefix
		}

		return newRedisBackend(redisClient, prefix), nil // 已应用 default 后的 prefix，避免传空
	default:
		return nil, fmt.Errorf("queue: unknown backend %q", cfg.Backend)
	}
}
