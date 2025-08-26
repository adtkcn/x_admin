package config

type geTuiConfig struct {
	HOST         string `mapstructure:"HOST"`         // 主机地址
	APPID        string `mapstructure:"APPID"`        // 应用ID
	APPKEY       string `mapstructure:"APPKEY"`       // 应用KEY
	APPSECRET    string `mapstructure:"APPSECRET"`    // 应用Secret
	MASTERSECRET string `mapstructure:"MASTERSECRET"` // 主Secret
	PackName     string `mapstructure:"PackName"`     // 包名
}

var GeTuiConfig = geTuiConfig{
	HOST:         "",
	APPID:        "",
	APPKEY:       "",
	APPSECRET:    "",
	MASTERSECRET: "",
	PackName:     "",
}
