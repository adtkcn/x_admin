package config

type RedisConfigStruct struct {
	Url             string `mapstructure:"Url"`             // Redis源配置: redis://:@127.0.0.1:6379/0
	PoolSize        int    `mapstructure:"PoolSize"`        // Redis连接池大小
	MaxIdleConns    int    `mapstructure:"MaxIdleConns"`    // Redis空闲连接池最大值
	ConnMaxLifetime int    `mapstructure:"ConnMaxLifetime"` // Redis连接可复用的最大时间(秒)
	RedisPrefix     string `mapstructure:"RedisPrefix"`     // Redis键前缀: x:如果多个基于此项目使用同一个redis，应该更改前缀
}

var RedisConfig = RedisConfigStruct{
	Url:      "redis://:@127.0.0.1:6379/0",
	PoolSize: 100,
	// 空闲连接贴近 PoolSize，避免 IoT 报文洪峰下连接被大量回收后频繁重握手
	MaxIdleConns: 50,
	// 放宽到 5 分钟：BRPOP 等长阻塞连接若按 60s 回收会打断阻塞读
	ConnMaxLifetime: 300,
	// 默认x:
	RedisPrefix: "x:",
}
