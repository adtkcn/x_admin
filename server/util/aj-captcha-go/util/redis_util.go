package util

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisUtil struct {
	Rdb redis.UniversalClient
}

func NewConfigRedisUtil(client redis.UniversalClient) *RedisUtil {
	redisUtil := &RedisUtil{
		Rdb: client,
	}
	return redisUtil
}

func (l *RedisUtil) Exists(key string) bool {
	timeVal := l.Rdb.Get(context.Background(), key+"_HoldTime").Val()
	cacheHoldTime, err := strconv.ParseInt(timeVal, 10, 64)

	if err != nil {
		return false
	}

	if cacheHoldTime == 0 {
		return true
	}

	if cacheHoldTime < time.Now().Unix() {
		l.Delete(key)
		return false
	}
	return true
}

func (l *RedisUtil) Get(key string) string {
	val := l.Rdb.Get(context.Background(), key).Val()
	return val
}

func (l *RedisUtil) Set(key string, val string, expiresInSeconds int) {
	//设置阈值，达到即clear缓存
	rdsResult := l.Rdb.Set(context.Background(), key, val, time.Duration(expiresInSeconds)*time.Second)
	fmt.Println("rdsResult: ", rdsResult.String(), "rdsErr: ", rdsResult.Err())
	if expiresInSeconds > 0 {
		// 缓存失效时间
		nowTime := time.Now().Unix() + int64(expiresInSeconds)
		l.Rdb.Set(context.Background(), key+"_HoldTime", strconv.FormatInt(nowTime, 10), time.Duration(expiresInSeconds)*time.Second)
	} else {
		l.Rdb.Set(context.Background(), key+"_HoldTime", strconv.FormatInt(0, 10), time.Duration(expiresInSeconds)*time.Second)
	}
}

func (l *RedisUtil) Delete(key string) {
	l.Rdb.Del(context.Background(), key)
	l.Rdb.Del(context.Background(), key+"_HoldTime")
}

func (l *RedisUtil) Clear() {
	//for key, _ := range l.Data {
	//	l.Delete(key)
	//}
}
