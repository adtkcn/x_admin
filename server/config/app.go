package config

type appConfig struct {
	AppName string `mapstructure:"APP_NAME"` // 应用名称
	Version string `mapstructure:"VERSION"`  // 应用版本
	Secret  string `mapstructure:"SECRET"`   // 应用系统加密字符
	Port    int    `mapstructure:"PORT"`     // 应用端口

	OssDomain      string `mapstructure:"OSS_DOMAIN"`      // OSS域名
	GinMode        string `mapstructure:"GIN_MODE"`        // 应用模式: debug,release
	DisallowModify bool   `mapstructure:"DISALLOW_MODIFY"` // 禁止修改操作 (演示功能,限制POST请求)
}

var AppConfig = appConfig{
	AppName:        "x_admin",
	Version:        "1.0.0",
	Secret:         "UVTIyzCy",
	Port:           8080,
	GinMode:        "",
	OssDomain:      "",
	DisallowModify: false,
}
