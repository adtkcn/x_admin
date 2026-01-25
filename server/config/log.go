package config

type LogConfigStruct struct {
	// 日志级别 debug, info, warn, error
	Level string `mapstructure:"Level"`

	// 是否开启控制台输出
	EnableConsole bool `mapstructure:"EnableConsole"`

	// 是否开启文件输出
	EnableFile bool `mapstructure:"EnableFile"`

	// 日志文件路径
	Filename string `mapstructure:"Filename"`

	// 单个文件最大 MB
	MaxSize int `mapstructure:"MaxSize"`

	// 最多保留几个备份文件
	MaxBackups int `mapstructure:"MaxBackups"`

	// 文件最多保存多少天
	MaxAge int `mapstructure:"MaxAge"`

	// 是否压缩旧文件
	Compress bool `mapstructure:"Compress"`
}

var LogConfig = LogConfigStruct{
	Level:         "debug",
	EnableConsole: true,
	EnableFile:    true,
	Filename:      "./logs/likeadmin.log",
	MaxSize:       100,
	MaxBackups:    7,
	MaxAge:        30,
	Compress:      false,
}
