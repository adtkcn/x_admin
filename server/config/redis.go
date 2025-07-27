package config

import "time"

type redisConfig struct {
	Url             string        `mapstructure:"URL"`               // Redis源配置: redis://:@127.0.0.1:6379/0
	PoolSize        int           `mapstructure:"POOL_SIZE"`         // Redis连接池大小
	MaxIdleConns    int           `mapstructure:"MAX_IDLE_CONNS"`    // Redis空闲连接池最大值
	ConnMaxLifetime time.Duration `mapstructure:"CONN_MAX_LIFETIME"` // Redis连接可复用的最大时间(秒：默认60秒)
	RedisPrefix     string        `mapstructure:"REDIS_PREFIX"`      // Redis键前缀: x:
}

var RedisConfig = redisConfig{
	Url:             "redis://:@127.0.0.1:6379/0",
	PoolSize:        100,
	MaxIdleConns:    10,
	ConnMaxLifetime: 60 * time.Second,
	// 默认x:
	RedisPrefix: "x:",
}
