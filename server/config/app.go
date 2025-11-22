package config

type appConfig struct {
	AppName string `mapstructure:"AppName"` // 应用名称
	Version string `mapstructure:"Version"` // 应用版本
	Port    int    `mapstructure:"Port"`    // 应用端口

	OssDomain      string `mapstructure:"OssDomain"`      // OSS域名
	GinMode        string `mapstructure:"GinMode"`        // 应用模式: debug,release
	DisallowModify bool   `mapstructure:"DisallowModify"` // 禁止修改操作 (演示功能,限制POST请求)
}

var AppConfig = appConfig{
	AppName:        "x_admin",
	Version:        "1.0.0",
	Port:           8080,
	GinMode:        "",
	OssDomain:      "",
	DisallowModify: false,
}
