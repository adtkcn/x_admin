package config

import (
	"flag"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type config struct {
	APP   *AppConfigStruct
	DB    *DBConfigStruct
	REDIS *RedisConfigStruct
	FILE  *FileConfigStruct
	GeTui *GeTuiConfigStruct
	Email *EmailConfigStruct
	Log   *LogConfigStruct
}

var Config = loadConfig(config{
	APP:   &AppConfig,
	DB:    &DBConfig,
	REDIS: &RedisConfig,
	FILE:  &FileConfig,
	GeTui: &GeTuiConfig,
	Email: &EmailConfig,
	Log:   &LogConfig,
})

func loadConfig(config config) config {
	var envFilePath string
	// 读取命令行参数 -env 配置文件路径，默认运行目录下的.env文件,使用：-env=.env
	flag.StringVar(&envFilePath, "env", "", "-env 配置文件路径，默认运行目录下的.env文件")
	flag.Parse()
	if envFilePath == "" {
		envFilePath = ".env"
	}
	// viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigFile(envFilePath)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("loadConfig ReadInConfig err:", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("loadConfig Unmarshal err:", err)
	}
	return config
}

func init() {
	fmt.Printf("AppConfig: %+v\n", AppConfig)
	fmt.Printf("DBConfig: %+v\n", DBConfig)
	fmt.Printf("RedisConfig: %+v\n", RedisConfig)
	fmt.Printf("FileConfig: %+v\n", FileConfig)
	fmt.Printf("GeTuiConfig: %+v\n", GeTuiConfig)
	fmt.Printf("EmailConfig: %+v\n", EmailConfig)
	fmt.Printf("LogConfig: %+v\n", LogConfig)
}
