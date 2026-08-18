package config

// WechatConfig 微信配置（小程序+公众号+支付）
var WechatConfig = wechatConfig{
	MiniAppID:  "wxfd733ac8440284e5",
	MiniSecret: "f479c9c66968374b49a55fd239204a01",
}

// wechatPaymentConfig 微信支付配置（v3 API）
type wechatPaymentConfig struct {
	MchID       string `mapstructure:"MchID"`       // 商户号
	MchApiV3Key string `mapstructure:"MchApiV3Key"` // APIv3 密钥（用于回调报文解密）
	KeyPath     string `mapstructure:"KeyPath"`     // 商户 API 私钥 pem 路径（apiclient_key.pem）
	SerialNo    string `mapstructure:"SerialNo"`    // 商户 API 证书序列号
	NotifyURL   string `mapstructure:"NotifyURL"`   // 支付结果通知地址
}

type wechatConfig struct {
	MiniAppID     string              `mapstructure:"MiniAppID"`     // 小程序 AppID
	MiniSecret    string              `mapstructure:"MiniSecret"`    // 小程序 AppSecret
	MpAppID       string              `mapstructure:"MpAppID"`       // 公众号 AppID
	MpSecret      string              `mapstructure:"MpSecret"`      // 公众号 AppSecret
	MpRedirectURI string              `mapstructure:"MpRedirectURI"` // 公众号 OAuth 回调地址
	Payment       wechatPaymentConfig `mapstructure:"Payment"`
}
