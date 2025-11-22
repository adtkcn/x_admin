package config

type redisConfig struct {
	Url             string `mapstructure:"Url"`             // Redis源配置: redis://:@127.0.0.1:6379/0
	PoolSize        int    `mapstructure:"PoolSize"`        // Redis连接池大小
	MaxIdleConns    int    `mapstructure:"MaxIdleConns"`    // Redis空闲连接池最大值
	ConnMaxLifetime int    `mapstructure:"ConnMaxLifetime"` // Redis连接可复用的最大时间(秒：默认60秒)
	RedisPrefix     string `mapstructure:"RedisPrefix"`     // Redis键前缀: x:
}

var RedisConfig = redisConfig{
	Url:             "redis://:@127.0.0.1:6379/0",
	PoolSize:        100,
	MaxIdleConns:    10,
	ConnMaxLifetime: 60,
	// 默认x:
	RedisPrefix: "x:",
}
