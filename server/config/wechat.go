package config

// WechatConfig 微信配置（小程序+公众号）
var WechatConfig = wechatConfig{}

type wechatConfig struct {
	MiniAppID     string `mapstructure:"MiniAppID"`     // 小程序 AppID
	MiniSecret    string `mapstructure:"MiniSecret"`    // 小程序 AppSecret
	MpAppID       string `mapstructure:"MpAppID"`       // 公众号 AppID
	MpSecret      string `mapstructure:"MpSecret"`      // 公众号 AppSecret
	MpRedirectURI string `mapstructure:"MpRedirectURI"` // 公众号 OAuth 回调地址
}
