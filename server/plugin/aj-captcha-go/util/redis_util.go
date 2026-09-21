package util

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisUtil struct {
	Rdb redis.UniversalClient
}

func NewConfigRedisUtil(client redis.UniversalClient) *RedisUtil {
	return &RedisUtil{
		Rdb: client,
	}
}

// Exists 检查 key 是否存在（利用 Redis 原生 TTL 管理过期）
func (l *RedisUtil) Exists(key string) bool {
	result, err := l.Rdb.Exists(context.Background(), key).Result()
	if err != nil {
		return false
	}
	return result > 0
}

// Get 获取 key 的值（Redis 自动管理过期）
func (l *RedisUtil) Get(key string) string {
	val, err := l.Rdb.Get(context.Background(), key).Result()
	if err != nil {
		return ""
	}
	return val
}

// Set 设置 key 的值，expiresInSeconds 为过期时间（秒）
func (l *RedisUtil) Set(key string, val string, expiresInSeconds int) {
	var expiration time.Duration
	if expiresInSeconds > 0 {
		expiration = time.Duration(expiresInSeconds) * time.Second
	}
	l.Rdb.Set(context.Background(), key, val, expiration)
}

// Delete 删除 key
func (l *RedisUtil) Delete(key string) {
	l.Rdb.Del(context.Background(), key)
}

// Clear 清空（不实现，Redis 由 TTL 自动管理）
func (l *RedisUtil) Clear() {
}
