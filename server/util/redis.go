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

// SetNX 仅当 key 不存在时设置（原子操作），常用于去重。设置成功返回 true。
func (ru redisUtil) SetNX(key string, value any, timeSec int) bool {
	ok, err := ru.redis.SetNX(context.Background(),
		config.RedisConfig.RedisPrefix+key, value, time.Duration(timeSec)*time.Second).Result()
	if err != nil {
		core.Logger.Errorf("redisUtil.SetNX err: err=[%+v]", err)
		return false
	}
	return ok
}

// Get 获取key的值
func (ru redisUtil) Get(key string) string {
	res, err := ru.redis.Get(context.Background(), config.RedisConfig.RedisPrefix+key).Result()
	if err != nil {
		core.Logger.Errorf("redisUtil.Get key=%s err: err=[%+v]", key, err)
		return ""
	}
	return res
}

// MGet 批量获取多个key的值,key 不存在时对应位置返回空串(不报错)。
func (ru redisUtil) MGet(keys ...string) []string {
	if len(keys) == 0 {
		return []string{}
	}
	fullKeys := ru.toFullKeys(keys)
	res, err := ru.redis.MGet(context.Background(), fullKeys...).Result()
	if err != nil {
		core.Logger.Errorf("redisUtil.MGet err: err=[%+v]", err)
		return make([]string, len(keys))
	}
	vals := make([]string, len(res))
	for i, v := range res {
		if s, ok := v.(string); ok {
			vals[i] = s
		}
	}
	return vals
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

// ValuesByPrefix 按前缀扫描并返回所有匹配的 key-value(使用 SCAN 游标遍历, 避免 KEYS 阻塞)。
// 注意: prefix 为业务侧传入的 key(不含 RedisPrefix), 方法内部会自动拼接前缀后匹配。
// 返回的 map 的 key 为去掉业务前缀后的剩余部分(便于调用方提取业务标识, 如 ip), value 为字符串值。
func (ru redisUtil) ValuesByPrefix(prefix string) map[string]string {
	ctx := context.Background()
	fullPrefix := config.RedisConfig.RedisPrefix + prefix + ":"
	result := map[string]string{}

	// SCAN 游标
	var cursor uint64
	for {
		keys, next, err := ru.redis.Scan(ctx, cursor, fullPrefix+"*", 100).Result()
		if err != nil {
			core.Logger.Errorf("redisUtil.ValuesByPrefix Scan err: err=[%+v]", err)
			return result
		}
		// 本批 key 一次性批量获取(MGet),减少 Redis 往返次数
		vals := ru.MGet(keys...)
		for i, fullKey := range keys {
			// MGet 中不存在的 key 返回空串(如已过期),跳过
			if vals[i] == "" {
				continue
			}
			// 去掉业务前缀
			rest := fullKey
			if len(rest) > len(prefix) {
				rest = rest[len(fullPrefix):]
			}
			result[rest] = vals[i]
		}
		cursor = next
		if cursor == 0 {
			break
		}
	}
	return result
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

// toFullKeys 为keys批量增加前缀(若key已包含前缀则不再重复添加)
func (ru redisUtil) toFullKeys(keys []string) (fullKeys []string) {
	prefix := config.RedisConfig.RedisPrefix
	for _, k := range keys {
		if prefix != "" && strings.HasPrefix(k, prefix) {
			fullKeys = append(fullKeys, k)
			continue
		}
		fullKeys = append(fullKeys, prefix+k)
	}
	return
}
