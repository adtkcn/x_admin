package config

type geTuiConfig struct {
	Host         string `mapstructure:"Host"`         // 主机地址
	AppID        string `mapstructure:"AppID"`        // 应用ID
	AppKEY       string `mapstructure:"AppKEY"`       // 应用KEY
	AppSECRET    string `mapstructure:"AppSECRET"`    // 应用Secret
	MasterSecret string `mapstructure:"MasterSecret"` // 主Secret
	PackName     string `mapstructure:"PackName"`     // 包名
}

var GeTuiConfig = geTuiConfig{
	Host:         "",
	AppID:        "",
	AppKEY:       "",
	AppSECRET:    "",
	MasterSecret: "",
	PackName:     "",
}
