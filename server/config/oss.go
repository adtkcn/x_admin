package config

// OssConfig 对象存储配置（兼容阿里云 OSS / MinIO 的 S3 协议）
var OssConfig = ossConfig{}

type ossConfig struct {
	Endpoint        string `mapstructure:"Endpoint"`
	PublicEndpoint  string `mapstructure:"PublicEndpoint"`
	Bucket          string `mapstructure:"Bucket"`
	AccessKeyId     string `mapstructure:"AccessKeyId"`
	AccessKeySecret string `mapstructure:"AccessKeySecret"`
	Region          string `mapstructure:"Region"`
}
