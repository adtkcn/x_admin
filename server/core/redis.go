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
		log.Fatal("initRedis redis.ParseURL err: ", err)
	}
	// opt.PoolSize = config.Config.RedisPoolSize
	opt.MaxIdleConns = config.RedisConfig.MaxIdleConns
	opt.ConnMaxLifetime = time.Duration(config.RedisConfig.ConnMaxLifetime) * time.Second

	client := redis.NewClient(opt)

	return client
}
