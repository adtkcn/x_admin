package util

import (
	"bufio"
	"context"
	"strings"
	"time"
	"x_admin/config"
	"x_admin/core"

	"github.com/redis/go-redis/v9"
)

var RedisUtil = NewRedis()

func NewRedis() redisUtil {
	return redisUtil{redis: core.Redis}
}

// redisUtil Redis操作工具类
type redisUtil struct {
	redis *redis.Client
}

// stringToLines string拆分多行
func stringToLines(s string) (lines []string, err error) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	return
}

// stringToKV string拆分key和val
func stringToKV(s string) (string, string) {
	ss := strings.Split(s, ":")
	if len(ss) < 2 {
		return s, ""
	}
	return ss[0], ss[1]
}

// Info Redis服务信息
func (ru redisUtil) Info(sections ...string) (res map[string]string) {
	infoStr, err := ru.redis.Info(context.Background(), sections...).Result()
	res = map[string]string{}
	if err != nil {
		core.Logger.Errorf("redisUtil.Info err: err=[%+v]", err)
		return res
	}
	// string拆分多行
	lines, err := stringToLines(infoStr)
	if err != nil {
		core.Logger.Errorf("stringToLines err: err=[%+v]", err)
		return res
	}
	// 解析成Map
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" || strings.HasPrefix(lines[i], "# ") {
			continue
		}
		k, v := stringToKV(lines[i])
		res[k] = v
	}
	return res
}

// DBSize 当前数据库key数量
func (ru redisUtil) DBSize() int64 {
	size, err := ru.redis.DBSize(context.Background()).Result()
	if err != nil {
		core.Logger.Errorf("redisUtil.DBSize err: err=[%+v]", err)
		return 0
	}
	return size
}

// Expire 指定缓存失效时间
func (ru redisUtil) Expire(key string, timeSec int) bool {
	err := ru.redis.Expire(context.Background(), config.RedisConfig.RedisPrefix+key, time.Duration(timeSec)*time.Second).Err()
	if err != nil {
		core.Logger.Errorf("redisUtil.Expire err: err=[%+v]", err)
		return false
	}
	return true
}

// TTL 根据key获取过期时间
func (ru redisUtil) TTL(key string) int {
	td, err := ru.redis.TTL(context.Background(), config.RedisConfig.RedisPrefix+key).Result()
	if err != nil {
		core.Logger.Errorf("redisUtil.TTL err: err=[%+v]", err)
		return 0
	}
	return int(td / time.Second)
}

// Del 删除一个或多个键
func (ru redisUtil) Del(keys ...string) bool {
	fullKeys := ru.toFullKeys(keys)
	err := ru.redis.Del(context.Background(), fullKeys...).Err()
	if err != nil {
		core.Logger.Errorf("redisUtil.Del err: err=[%+v]", err)
		return false
	}
	return true
}

// Incr 将 key 中存储的数字值加 1
func (ru redisUtil) Incr(key string) int {
	res, err := ru.redis.Incr(context.Background(), config.RedisConfig.RedisPrefix+key).Result()
	if err != nil {
		core.Logger.Errorf("redisUtil.Incr err: err=[%+v]", err)
		return 0
	}
	return int(res)
}

// Exists 判断多项key是否存在
func (ru redisUtil) Exists(keys ...string) int64 {
	fullKeys := ru.toFullKeys(keys)
	// count 存在的key数量
	count, err := ru.redis.Exists(context.Background(), fullKeys...).Result()
	if err != nil {
		core.Logger.Errorf("redisUtil.Exists err: err=[%+v]", err)
		return -1
	}
	return count
}

// Set 设置键值对
func (ru redisUtil) Set(key string, value any, timeSec int) bool {
	err := ru.redis.Set(context.Background(),
		config.RedisConfig.RedisPrefix+key, value, time.Duration(timeSec)*time.Second).Err()
	if err != nil {
		core.Logger.Errorf("redisUtil.Set err: err=[%+v]", err)
		return false
	}
	return true
}

// Get 获取key的值
func (ru redisUtil) Get(key string) string {
	res, err := ru.redis.Get(context.Background(), config.RedisConfig.RedisPrefix+key).Result()
	if err != nil {
		core.Logger.Errorf("redisUtil.Get err: err=[%+v]", err)
		return ""
	}
	return res
}

// SAdd 将数据放入set缓存
func (ru redisUtil) SAdd(key string, values ...any) bool {
	err := ru.redis.SAdd(context.Background(), config.RedisConfig.RedisPrefix+key, values...).Err()
	if err != nil {
		core.Logger.Errorf("redisUtil.SAdd err: err=[%+v]", err)
		return false
	}
	return true
}

// SGet 根据key获取Set中的所有值
func (ru redisUtil) SGet(key string) []string {
	res, err := ru.redis.SMembers(context.Background(), config.RedisConfig.RedisPrefix+key).Result()
	if err != nil {
		core.Logger.Errorf("redisUtil.SGet err: err=[%+v]", err)
		return []string{}
	}
	return res
}

// HMSet 设置key, 通过字典的方式设置多个field, value对
func (ru redisUtil) HMSet(key string, value any, timeSec int) bool {
	err := ru.redis.HMSet(context.Background(), config.RedisConfig.RedisPrefix+key, value).Err()
	if err != nil {
		core.Logger.Errorf("redisUtil.HMSet err: err=[%+v]", err)
		return false
	}
	if timeSec > 0 {
		if !ru.Expire(key, timeSec) {
			return false
		}
	}
	return true
}

// HGet 获取key中field域的值
func (ru redisUtil) HGet(key string, field string) string {
	res, err := ru.redis.HGet(context.Background(), config.RedisConfig.RedisPrefix+key, field).Result()
	if err != nil {
		core.Logger.Errorf("redisUtil.HGet err: err=[%+v]", err)
		return ""
	}
	return res
}

// HGetAll
func (ru redisUtil) HGetAll(key string) map[string]string {
	res, err := ru.redis.HGetAll(context.Background(), config.RedisConfig.RedisPrefix+key).Result()
	if err != nil {
		core.Logger.Errorf("redisUtil.HGetAll err: err=[%+v]", err)
		return map[string]string{}
	}
	return res
}

// HExists 判断key中有没有field域名
func (ru redisUtil) HExists(key string, field string) bool {
	res, err := ru.redis.HExists(context.Background(), config.RedisConfig.RedisPrefix+key, field).Result()
	if err != nil {
		core.Logger.Errorf("redisUtil.HExists err: err=[%+v]", err)
		return false
	}
	return res
}

// HDel 删除hash表中的值
func (ru redisUtil) HDel(key string, fields ...string) bool {
	err := ru.redis.HDel(context.Background(), config.RedisConfig.RedisPrefix+key, fields...).Err()
	if err != nil {
		core.Logger.Errorf("redisUtil.HDel err: err=[%+v]", err)
		return false
	}
	return true
}

// Push 向列表中添加元素,并保留最新的count个元素
func (ru redisUtil) RPush(key string, value []any, count int64) bool {
	var ctx = context.Background()
	pipe := core.Redis.TxPipeline()
	pipe.RPush(ctx, config.RedisConfig.RedisPrefix+key, value...) // 推到右侧（尾部）
	if count != 0 {
		pipe.LTrim(ctx, config.RedisConfig.RedisPrefix+key, -count, -1) // 保留最新的count个元素
	}

	_, err := pipe.Exec(ctx)

	if err != nil {
		core.Logger.Errorf("redisUtil.Push err: err=[%+v]", err)
		return false
	}
	return true
}

// LIndex 获取列表中指定索引的元素
func (ru redisUtil) LRange(key string, start, stop int64) []string {
	res, err := ru.redis.LRange(context.Background(), config.RedisConfig.RedisPrefix+key, start, stop).Result()
	if err != nil {
		core.Logger.Errorf("redisUtil.LRange err: err=[%+v]", err)
		return []string{}
	}
	return res
}

// DelByPrefix 按前缀批量删除键(使用 SCAN 游标遍历, 避免 KEYS 阻塞)
// 注意: match 为业务侧传入的 key(不含 RedisPrefix), 方法内部会自动拼接前缀后匹配
func (ru redisUtil) DelByPrefix(match string) int64 {
	ctx := context.Background()
	fullMatch := config.RedisConfig.RedisPrefix + match
	var cursor uint64
	var deleted int64
	for {
		keys, next, err := ru.redis.Scan(ctx, cursor, fullMatch, 100).Result()
		if err != nil {
			core.Logger.Errorf("redisUtil.DelByPrefix Scan err: err=[%+v]", err)
			return deleted
		}
		if len(keys) > 0 {
			if n, err := ru.redis.Del(ctx, keys...).Result(); err == nil {
				deleted += n
			}
		}
		cursor = next
		if cursor == 0 {
			break
		}
	}
	return deleted
}

// toFullKeys 为keys批量增加前缀
func (ru redisUtil) toFullKeys(keys []string) (fullKeys []string) {
	for _, k := range keys {
		fullKeys = append(fullKeys, config.RedisConfig.RedisPrefix+k)
	}
	return
}
