package config

type ConstantConfigStruct struct {
	DateFormat string `mapstructure:"DateFormat"` // 日期格式
	TimeFormat string `mapstructure:"TimeFormat"` // 时间格式
}

var ConstantConfig = ConstantConfigStruct{
	DateFormat: "2006-01-02",
	TimeFormat: "2006-01-02 15:04:05",
}
