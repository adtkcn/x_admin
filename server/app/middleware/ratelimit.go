package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/ratelimit"
)

// 限制每秒请求数量
// r := rg.Group("/", middleware.RateLimiterMiddleware(200))
func RateLimiterMiddleware(rate int) gin.HandlerFunc {
	rl := ratelimit.New(rate)

	return func(c *gin.Context) {
		// Try to take a token from the rate limiter.
		now := rl.Take()

		// Check if we've waited longer than expected. If so, this means the rate
		// limit has been exceeded.
		if now.After(time.Now()) {
			c.AbortWithStatus(http.StatusTooManyRequests)
			return
		}

		c.Next()
	}
}
