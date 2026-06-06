package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/ratelimit"
)

// RateLimiterMiddleware 全局限流中间件（基于 uber-go/ratelimit 漏桶算法）
// rate: 每秒允许通过的请求数，如 200 表示每秒最多 200 个请求
//
// 使用示例：
//
//	r := rg.Group("/api", middleware.RateLimiterMiddleware(200))
//
// 注意：
//   - 这是进程级别的全局限流，不区分用户或 IP
//   - 超限时直接返回 429 Too Many Requests，不执行后续 handler
func RateLimiterMiddleware(rate int) gin.HandlerFunc {
	// 创建漏桶限流器，每秒产生 rate 个令牌
	rl := ratelimit.New(rate)

	return func(c *gin.Context) {
		// 尝试获取一个令牌，返回预计可执行时间
		now := rl.Take()

		// 如果预计时间在未来（当前时间之后），说明令牌不足，触发限流
		if now.After(time.Now()) {
			c.AbortWithStatus(http.StatusTooManyRequests)
			return
		}

		c.Next()
	}
}
