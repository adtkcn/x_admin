package core

import (
	"context"
	"log"
	"time"
	"x_admin/config"

	"github.com/redis/go-redis/v9"
)

var Redis = initRedis()

// initRedis 初始化redis客户端
func initRedis() *redis.Client {
	opt, err := redis.ParseURL(config.Config.RedisUrl)
	if err != nil {
		log.Fatal("initRedis redis.ParseURL err: ", err)
	}
	// opt.PoolSize = config.Config.RedisPoolSize
	opt.MaxIdleConns = config.Config.RedisMaxIdleConns
	opt.ConnMaxLifetime = config.Config.RedisConnMaxLifetime

	client := redis.NewClient(opt)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err = client.Ping(ctx).Result()
	if err != nil {
		log.Fatal("initRedis client.Ping err: ", err)
	}
	return client
}
