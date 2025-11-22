package config

type emailConfig struct {
	Host     string `mapstructure:"Host"`     // SMTP服务器地址，如 "smtp.qq.com"
	Port     int    `mapstructure:"Port"`     // SMTP服务器端口，如 465
	SSL      bool   `mapstructure:"SSL"`      // 是否启用SSL/TLS，如 true
	Username string `mapstructure:"Username"` // SMTP服务器用户名，如 "adtkcn@qq.com"
	Password string `mapstructure:"Password"` // SMTP服务器密码，如 "授权码"
	Timeout  int    `mapstructure:"Timeout"`  // 超时时间，单位秒，默认 10s
}

var EmailConfig = emailConfig{
	Host:     "smtp.qq.com",
	Port:     465,
	SSL:      true,
	Username: "",
	Password: "",
	Timeout:  10,
}
