package core

import (
	"log"
	"time"
	"x_admin/config"

	"github.com/redis/go-redis/v9"
)

var Redis = initRedis()

// initRedis 初始化redis客户端
func initRedis() *redis.Client {
	opt, err := redis.ParseURL(config.RedisConfig.Url)
	if err != nil {
		log.Fatal("解析redis url失败: ", err)
	}
	// 启用 RESP3 协议：pubsub 走独立 push 通道，不再挤占普通命令回复
	opt.Protocol = 3
	opt.MaxIdleConns = config.RedisConfig.MaxIdleConns
	opt.ConnMaxLifetime = time.Duration(config.RedisConfig.ConnMaxLifetime) * time.Second

	client := redis.NewClient(opt)

	return client
}
