package config

import (
	"flag"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type config struct {
	APP   *appConfig
	DB    *dbConfig
	REDIS *redisConfig
	FILE  *fileConfig
}

var Config = loadConfig(config{
	APP:   &AppConfig,
	DB:    &DBConfig,
	REDIS: &RedisConfig,
	FILE:  &FileConfig,
})

func loadConfig(config config) config {
	var envFilePath string
	flag.StringVar(&envFilePath, "env", "", "-env 配置文件路径，默认运行目录下的.env文件")
	flag.Parse()
	if envFilePath == "" {
		envFilePath = ".env"
	}
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
	// fmt.Println(Config)
	fmt.Println("AppConfig:", AppConfig)
	fmt.Println("DBConfig:", DBConfig)
	fmt.Println("RedisConfig:", RedisConfig)
	fmt.Println("FileConfig:", FileConfig)
}

// loadConfig 加载配置
// func loadConfig(envPath string) envConfig {
// 	var cfgPath string
// 	flag.StringVar(&cfgPath, "c", "", "config file envPath.")
// 	flag.Parse()
// 	if cfgPath == "" {
// 		viper.AddConfigPath(envPath)
// 		viper.SetConfigFile(".env")
// 	} else {
// 		viper.SetConfigFile(cfgPath)
// 	}
// 	viper.AutomaticEnv()
// 	err := viper.ReadInConfig()
// 	if err != nil {
// 		log.Fatal("loadConfig ReadInConfig err:", err)
// 	}
// 	err = viper.Unmarshal(&config)
// 	if err != nil {
// 		log.Fatal("loadConfig Unmarshal err:", err)
// 	}
// 	return config
// }
