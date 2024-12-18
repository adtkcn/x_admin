package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

var Config = loadConfig(".")

// #region envConfig
// envConfig 环境配置
type envConfig struct {
	RootPath        string // 项目根目录
	GinMode         string `mapstructure:"GIN_MODE"`        // gin运行模式
	PublicUrl       string `mapstructure:"PUBLIC_URL"`      // 对外发布的Url
	OssDomain       string `mapstructure:"OSS_DOMAIN"`      // OSS域名
	ServerPort      int    `mapstructure:"SERVER_PORT"`     // 服务运行端口
	DisallowModify  bool   `mapstructure:"DISALLOW_MODIFY"` // 禁止修改操作 (演示功能,限制POST请求)
	PublicPrefix    string // 资源访问前缀
	UploadDirectory string `mapstructure:"UPLOAD_DIRECTORY"` // 上传文件路径
	RedisUrl        string `mapstructure:"REDIS_URL"`        // Redis源配置
	// RedisPoolSize            int           // Redis连接池大小
	RedisMaxIdleConns        int           // Redis空闲连接池最大值
	RedisConnMaxLifetime     time.Duration // Redis连接可复用的最大时间(秒：默认60秒)
	DatabaseUrl              string        `mapstructure:"DATABASE_URL"` // 数据源配置
	DbTablePrefix            string        // Mysql表前缀
	DbDefaultStringSize      uint          // 数据库string类型字段的默认长度
	DbMaxIdleConns           int           // 数据库空闲连接池最大值
	DbMaxOpenConns           int           // 数据库连接池最大值
	DbConnMaxLifetimeSeconds int16         // 连接可复用的最大时间(秒：默认28800秒)，请根据这个sql查处的时间设置： show variables like 'wait_timeout'
	Version                  string        // 版本
	Secret                   string        // 系统加密字符

	RedisPrefix     string   // Redis键前缀
	UploadImageSize int64    // 上传图片限制
	UploadVideoSize int64    // 上传视频限制
	UploadImageExt  []string // 上传图片扩展
	UploadVideoExt  []string // 上传视频扩展
}

// #endregion envConfig

// loadConfig 加载配置
func loadConfig(envPath string) envConfig {
	var cfgPath string
	flag.StringVar(&cfgPath, "c", "", "config file envPath.")
	flag.Parse()
	if cfgPath == "" {
		viper.AddConfigPath(envPath)
		viper.SetConfigFile(".env")
	} else {
		viper.SetConfigFile(cfgPath)
	}
	viper.AutomaticEnv()

	rootPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("rootPath:", rootPath)
	config := envConfig{
		RootPath: rootPath,
		GinMode:  "debug",
		// 服务运行端口
		ServerPort: 8000,

		OssDomain: "",

		// 禁止修改操作 (演示功能,限制POST请求)
		DisallowModify: false,
		// 资源访问前缀
		PublicPrefix: "/api/uploads",
		// 上传文件路径
		UploadDirectory: "/tmp/uploads/x_admin-go/",
		// Redis源配置
		RedisUrl: "redis://localhost:6379",
		// RedisPoolSize:        100,
		RedisMaxIdleConns:    80,
		RedisConnMaxLifetime: 60 * time.Second,
		// 数据源配置
		DatabaseUrl:         "x_admin:x_admin@tcp(localhost:3306)/x_admin?charset=utf8mb4&parseTime=True&loc=Local",
		DbTablePrefix:       "x_",
		DbDefaultStringSize: 256,
		DbMaxIdleConns:      10,
		DbMaxOpenConns:      100,
		// 连接可复用的最大时间(秒：默认28800秒)
		DbConnMaxLifetimeSeconds: 28800,
		// 全局配置
		// 版本
		Version: "v1.1.0",
		// 系统加密字符
		Secret: "UVTIyzCy",

		// Redis键前缀
		RedisPrefix: "x:",
		// 上传图片限制
		UploadImageSize: 1024 * 1024 * 10,
		// 上传视频限制
		UploadVideoSize: 1024 * 1024 * 30,
		// 上传图片扩展
		UploadImageExt: []string{"png", "jpg", "jpeg", "gif", "ico", "bmp", "webp"},
		// 上传视频扩展
		UploadVideoExt: []string{"mp4", "mp3", "avi", "flv", "rmvb", "mov"},
	}
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal("loadConfig ReadInConfig err:", err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("loadConfig Unmarshal err:", err)
	}
	// PublicUrl未设置设置默认值
	// if config.PublicUrl == "" {
	// 	// config.PublicUrl = "http://127.0.0.1:" + strconv.Itoa(config.ServerPort)
	// }
	return config
}
