package config

import "time"

type dbConfig struct {
	Type                   string        `mapstructure:"TYPE"`                      // 数据库类型:MySQL, PostgreSQL, GaussDB, SQLite, SQLServer, TiDB
	Dsn                    string        `mapstructure:"DSN"`                       // 数据库
	MaxOpenConns           int           `mapstructure:"MAX_OPEN_CONNS"`            // 数据库连接池最大值
	MaxIdleConns           int           `mapstructure:"MAX_IDLE_CONNS"`            // 数据库空闲连接池最大值
	ConnMaxLifetimeSeconds int           `mapstructure:"CONN_MAX_LIFETIME_SECONDS"` // 连接可复用的最大时间(秒：默认28800秒)
	TablePrefix            string        `mapstructure:"TABLE_PREFIX"`              // 数据库表前缀
	SlowThreshold          time.Duration `mapstructure:"SLOW_THRESHOLD"`            // 数据库慢查询阈值(秒：默认1秒)
	LogLevel               string        `mapstructure:"LOG_LEVEL"`                 // 数据库日志级别
	DefaultStringSize      uint          `mapstructure:"DEFAULT_STRING_SIZE"`       // 数据库string类型字段的默认长度:256
}

var DBConfig = dbConfig{
	Type:                   "mysql",
	Dsn:                    "", //root:123456@tcp(127.0.0.1:3306)/x_admin?charset=utf8mb4&parseTime=True&loc=Local
	MaxOpenConns:           100,
	MaxIdleConns:           10,
	ConnMaxLifetimeSeconds: 60 * 60 * 24,
	TablePrefix:            "x_",
	SlowThreshold:          1 * time.Second,
	LogLevel:               "info",
}
