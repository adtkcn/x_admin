package middleware

import (
	"context"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
	"x_admin/config"
	"x_admin/core"

	"github.com/gin-gonic/gin"
)

const (
	// defaultCostKeep 每个接口默认保留的最近记录条数
	defaultCostKeep = 10000
	// costKeyPrefix Redis 键前缀，完整键为 req:cost:{METHOD}:{路由模板}
	costKeyPrefix = "req:cost:"
	// costTTL 记录过期时间：避免长期无流量的冷接口数据一直占用内存
	costTTL = 7 * 24 * time.Hour
	// costTTLRefresh 同一个 key 的 TTL 刷新间隔，避免每次写入都多一次 EXPIRE 命令
	costTTLRefresh = time.Hour
	// costChanSize 异步写入通道缓冲；通道满时丢弃统计，绝不阻塞业务请求
	costChanSize = 8192
	// costBatchSize 单次批量写入的最大条数
	costBatchSize = 64
	// costFlushInterval 批量写入的最大等待间隔
	costFlushInterval = 200 * time.Millisecond
	// costErrLogInterval 写入失败时的日志节流间隔（秒），避免 Redis 故障期间刷屏
	costErrLogInterval = 60
)

// costItem 待写入 Redis 的一条记录：某个接口的一次请求耗时
type costItem struct {
	key   string
	value string
	keep  int64 // 该接口需要保留的最近记录条数
}

var (
	costCh      chan costItem
	costOnce    sync.Once
	costTTLMu   sync.Mutex
	costTTLAt   = map[string]time.Time{}
	costLastErr atomic.Int64 // 上次写入失败的时间戳(秒)，用于日志节流
)

// RequestCost 请求耗时统计中间件（数据只存 Redis，不落库）
//
// 采用「按需挂载」：不做全局注册，哪个路由需要统计就单独挂上去，避免给
// 静态资源、文件流、长连接等无统计意义的接口带来额外开销。
//
// 按「路由模板」分组统计，例如 /api/admin/system/admin/:id 会归为一个 key，
// 而不是每个具体 id 一个 key，避免同接口因路径参数产生海量键。
// 每个接口只保留最近 maxKeep 条耗时记录，超出部分由 Redis 自动淘汰。
//
// 存储结构（Redis List），即 url -> [time, time, ...]：
//
//	key:   {RedisPrefix}req:cost:{METHOD}:{路由模板}
//	       例：x:req:cost:GET:/api/admin/system/admin/list
//	value: 单个请求的耗时(毫秒)，最新记录在列表右侧（RPUSH）
//	       读取示例：util.RedisUtil.LRange("req:cost:GET:/api/admin/system/admin/list", 0, -1)
//	       返回：["12","8","30","15"]  => 该接口最近若干次请求的耗时(毫秒)
//
// 写入为异步批量：业务请求只做一次数值格式化并入队，由后台协程聚合后
// 通过 pipeline 写入，因此 Redis 抖动不会拖慢或阻塞接口响应。
// 代价是进程异常退出时可能丢失内存中尚未刷出的最后一批数据。
//
// 使用示例（单个路由）：
//
//	r.GET("/list", middleware.RequestCost(10000), handler.List)
//
// 使用示例（整个路由组，组内所有接口都统计）：
//
//	adminGroup := r.Group("/admin")
//	adminGroup.Use(middleware.RequestCost(10000))
//
// 注意：maxKeep <= 0 时使用默认值 10000；不同路由可以传不同值，各自独立生效。
func RequestCost(maxKeep int64) gin.HandlerFunc {
	if maxKeep <= 0 {
		maxKeep = defaultCostKeep
	}
	keep := maxKeep

	costOnce.Do(func() {
		costCh = make(chan costItem, costChanSize)
		go costLoop()
	})

	return func(c *gin.Context) {
		// 未匹配到路由（如 404）没有路由模板，无法分组，跳过统计
		path := c.FullPath()
		if path == "" {
			c.Next()
			return
		}

		start := time.Now()
		c.Next()

		// 统计在响应写出之后执行，不占用业务处理时间
		item := costItem{
			key:   costKeyPrefix + c.Request.Method + ":" + path,
			value: strconv.FormatInt(time.Since(start).Milliseconds(), 10),
			keep:  keep,
		}
		select {
		case costCh <- item:
		default:
			// 通道已满：丢弃本次统计，保证业务请求不被监控逻辑拖慢
		}
	}
}

// costLoop 后台写入协程：攒够 costBatchSize 条或每 costFlushInterval 刷一次
func costLoop() {
	ticker := time.NewTicker(costFlushInterval)
	defer ticker.Stop()

	buf := make([]costItem, 0, costBatchSize)
	flush := func() {
		if len(buf) == 0 {
			return
		}
		costFlush(buf)
		buf = buf[:0]
	}

	for {
		select {
		case item, ok := <-costCh:
			if !ok {
				flush()
				return
			}
			buf = append(buf, item)
			if len(buf) >= costBatchSize {
				flush()
			}
		case <-ticker.C:
			flush()
		}
	}
}

// costFlush 按 key 聚合成批写入：同一接口的多条记录合并为一次 RPUSH + 一次 LTRIM
func costFlush(items []costItem) {
	ctx := context.Background()
	prefix := config.RedisConfig.RedisPrefix

	grouped := make(map[string][]any, 8)
	keeps := make(map[string]int64, 8)
	keys := make([]string, 0, 8)
	for _, it := range items {
		if _, ok := grouped[it.key]; !ok {
			keys = append(keys, it.key)
		}
		grouped[it.key] = append(grouped[it.key], it.value)
		keeps[it.key] = it.keep
	}

	now := time.Now()
	pipe := core.Redis.Pipeline()
	for _, key := range keys {
		fullKey := prefix + key
		pipe.RPush(ctx, fullKey, grouped[key]...)
		pipe.LTrim(ctx, fullKey, -keeps[key], -1) // 只保留最近 keep 条
		if costShouldExpire(key, now) {
			pipe.Expire(ctx, fullKey, costTTL)
		}
	}

	if _, err := pipe.Exec(ctx); err != nil {
		nowSec := time.Now().Unix()
		if nowSec-costLastErr.Load() > costErrLogInterval {
			costLastErr.Store(nowSec)
			core.Logger.Errorf("RequestCost flush err: err=[%+v]", err)
		}
	}
}

// costShouldExpire 判断该 key 的 TTL 是否需要刷新
func costShouldExpire(key string, now time.Time) bool {
	costTTLMu.Lock()
	defer costTTLMu.Unlock()
	if now.Sub(costTTLAt[key]) < costTTLRefresh {
		return false
	}
	costTTLAt[key] = now
	return true
}
