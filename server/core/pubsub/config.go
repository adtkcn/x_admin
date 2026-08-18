package pubsub

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

// Backend 事件总线后端类型。
type Backend string

const (
	// BackendRedis Redis 后端，基于 Redis Pub/Sub 跨实例广播（单机 / 集群均可用）。
	BackendRedis Backend = "redis"
)

// DefaultPrefix Redis 频道前缀为空时的中性默认值（不含业务前缀）。
const DefaultPrefix = "pubsub:"

// Config Emitter 配置，作为 New 工厂函数的唯一输入。
type Config struct {
	// Backend 后端类型，留空时默认 BackendRedis。
	Backend Backend
	// Prefix 事件频道前缀，建议带业务/项目前缀以隔离。
	// 事件类型 eventType 最终映射到 Redis 频道 {Prefix}{eventType}。
	// 留空时使用 DefaultPrefix。
	Prefix string
}

// New 根据配置创建 Emitter 实例，是唯一推荐的构造入口。
// redisClient 在 Backend == BackendRedis 时需要，传 nil 会返回错误。
//
// 新增后端（NATS / Kafka 等）只需实现 Emitter 接口并在此 switch 中注册分支，
// 调用方无需改动（保留可扩展性）。
func New(cfg Config, redisClient *redis.Client) (Emitter, error) {
	if cfg.Backend == "" {
		cfg.Backend = BackendRedis
	}
	switch cfg.Backend {
	case BackendRedis:
		if redisClient == nil {
			return nil, fmt.Errorf("pubsub: redis backend requires a non-nil redis client")
		}
		prefix := cfg.Prefix
		if prefix == "" {
			prefix = DefaultPrefix
		}
		return newRedisEmitter(redisClient, prefix), nil
	default:
		return nil, fmt.Errorf("pubsub: unknown backend %q", cfg.Backend)
	}
}
