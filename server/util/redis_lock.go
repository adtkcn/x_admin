package util

import (
	"context"
	"fmt"
	"time"
	"x_admin/core"

	"github.com/redis/go-redis/v9"
)

type RedisLock struct {
	client   *redis.Client
	key      string
	value    string // 用于标识锁的持有者（防止误删）
	expire   time.Duration
	retry    int
	interval time.Duration
	ctx      context.Context
}

func NewRedisLock(key string, expire time.Duration) *RedisLock {

	value := ToolsUtil.RandomString(10)
	// value := "1"
	return &RedisLock{
		client:   core.Redis,
		key:      key,
		value:    value,
		expire:   expire,
		retry:    3, // 默认重试 3 次
		interval: 50 * time.Millisecond,
		ctx:      context.Background(),
	}
}

// TryLock 尝试获取锁（非阻塞）
func (rl *RedisLock) TryLock() bool {
	result, err := rl.client.SetNX(rl.ctx, rl.key, rl.value, rl.expire).Result()
	if err != nil {
		return false
	}
	return result
}

// Lock 阻塞式加锁（带重试）
func (rl *RedisLock) Lock() bool {
	for i := 0; i < rl.retry; i++ {
		if rl.TryLock() {
			return true
		}
		time.Sleep(rl.interval)
	}
	return false
}

// Unlock 释放锁（使用 Lua 脚本保证原子性）
func (rl *RedisLock) Unlock() error {
	luaScript := `
		if redis.call("GET", KEYS[1]) == ARGV[1] then
			return redis.call("DEL", KEYS[1])
		else
			return 0
		end
	`
	result, err := rl.client.Eval(rl.ctx, luaScript, []string{rl.key}, rl.value).Result()
	if err != nil {
		return err
	}
	if result.(int64) != 1 {
		return fmt.Errorf("unlock failed: lock not held or expired")
	}
	return nil
}
