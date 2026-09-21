package captcha_service

import (
	"strconv"
	"x_admin/config"
	"x_admin/plugin/aj-captcha-go/util"

	"github.com/redis/go-redis/v9"
)

type RedisCacheService struct {
	Cache *util.RedisUtil
}

func genKey(key string) string {
	return config.RedisConfig.RedisPrefix + key
}

// NewConfigRedisCacheService 初始化自定义redis配置
func NewConfigRedisCacheService(client redis.UniversalClient) CacheCaptchaInterface {
	redisUtils := util.NewConfigRedisUtil(client)
	return &RedisCacheService{Cache: redisUtils}
}

func (l *RedisCacheService) Get(key string) string {
	return l.Cache.Get(genKey(key))
}

func (l *RedisCacheService) Set(key string, val string, expiresInSeconds int) {
	l.Cache.Set(genKey(key), val, expiresInSeconds)
}

func (l *RedisCacheService) Delete(key string) {
	l.Cache.Delete(genKey(key))
}

func (l *RedisCacheService) Exists(key string) bool {
	return l.Cache.Exists(genKey(key))
}

func (l *RedisCacheService) GetType() string {
	return "redis"
}

func (l *RedisCacheService) Increment(key string, val int) int {
	cacheVal := l.Cache.Get(genKey(key))
	num, err := strconv.Atoi(cacheVal)
	if err != nil {
		num = 0
	}

	ret := num + val

	l.Cache.Set(genKey(key), strconv.Itoa(ret), 0)
	return ret
}
