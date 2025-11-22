package config

type dbConfig struct {
	Type                   string `mapstructure:"Type"`                   // 数据库类型:MySQL, PostgreSQL, GaussDB, SQLite, SQLServer, TiDB
	Dsn                    string `mapstructure:"Dsn"`                    // 数据库
	MaxOpenConns           int    `mapstructure:"MaxOpenConns"`           // 数据库连接池最大值
	MaxIdleConns           int    `mapstructure:"MaxIdleConns"`           // 数据库空闲连接池最大值
	ConnMaxLifetimeSeconds int    `mapstructure:"ConnMaxLifetimeSeconds"` // 连接可复用的最大时间(秒：默认28800秒)
	TablePrefix            string `mapstructure:"TablePrefix"`            // 数据库表前缀
	SlowThreshold          int    `mapstructure:"SlowThreshold"`          // 数据库慢查询阈值(秒：默认1秒)
	LogLevel               string `mapstructure:"LogLevel"`               // 数据库日志级别
	DefaultStringSize      uint   `mapstructure:"DefaultStringSize"`      // 数据库string类型字段的默认长度:256
}

var DBConfig = dbConfig{
	Type:                   "MySQL",
	Dsn:                    "", //root:123456@tcp(127.0.0.1:3306)/x_admin?charset=utf8mb4&parseTime=True&loc=Local
	MaxOpenConns:           100,
	MaxIdleConns:           50,
	ConnMaxLifetimeSeconds: 6,
	TablePrefix:            "x_",
	SlowThreshold:          1,
	LogLevel:               "info",
	DefaultStringSize:      256,
}
