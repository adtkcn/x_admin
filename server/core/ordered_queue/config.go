package ordered_queue

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

// Backend 队列后端类型。
type Backend string

const (
	// BackendRedis Redis 后端：List 分片 + String 租约，跨实例有序。
	BackendRedis Backend = "redis"
)

// DefaultPrefix 默认键前缀。与 queue 的 "queue:" 分开，避免同名队列串到一处。
const DefaultPrefix = "oqueue:"

// DefaultShards 默认分片数 = 全局最大并行度。
// 投产后可以改，但改 N 会让同一设备的存量消息与新消息落到不同分片，顺序当场断裂；
// 必须换队列名（或等旧分片排空）后再改。故只做构造期参数，不做运行时热调。
const DefaultShards = 16

// Config OrderedQueue 配置，作为 New 工厂函数的唯一输入。
type Config struct {
	// Backend 后端类型，留空时默认 BackendRedis。
	Backend Backend
	// Prefix 队列键前缀，必须带项目前缀以隔离多实例/多环境（同 queue 约定）。
	Prefix string
	// Shards 分片数，<= 0 时取 DefaultShards。
	Shards int
}

// New 根据配置创建 OrderedQueue 实例，是唯一推荐的构造入口。
// redisClient 在 Backend == BackendRedis 时必需，传 nil 返回错误。
func New(cfg Config, redisClient *redis.Client) (OrderedQueue, error) {
	if cfg.Backend == "" {
		cfg.Backend = BackendRedis
	}
	switch cfg.Backend {
	case BackendRedis:
		if redisClient == nil {
			return nil, fmt.Errorf("ordered_queue: redis backend requires a non-nil redis client")
		}
		prefix := cfg.Prefix
		if prefix == "" {
			prefix = DefaultPrefix
		}
		shards := cfg.Shards
		if shards <= 0 {
			shards = DefaultShards
		}
		return newRedisBackend(redisClient, prefix, shards), nil
	default:
		return nil, fmt.Errorf("ordered_queue: unknown backend %q", cfg.Backend)
	}
}
