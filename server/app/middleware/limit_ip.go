package middleware

import (
	"net/http"
	"strconv"

	"x_admin/util"

	"github.com/gin-gonic/gin"
)

// EmailCodeIPLimit 基于 IP 的邮箱验证码发送限流
// limit: 窗口内最大次数  ttl: 窗口秒数
//
// 使用示例：
//
//	r.POST("/sendCode", middleware.EmailCodeIPLimit(3, 60), handler)
func LimitIP(limit int, ttl int) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		key := "limit:ip:" + ip

		if exceeded(key, limit, ttl) {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"code":    429,
				"message": "操作过于频繁，请稍后重试",
			})
			return
		}
		c.Next()
	}
}

// exceeded 检查并递增 Redis 计数器，返回 true 表示已超限
func exceeded(key string, limit int, ttl int) bool {
	val := util.RedisUtil.Get(key)
	current, _ := strconv.Atoi(val)
	if current >= limit {
		return true
	}
	current++
	util.RedisUtil.Set(key, current, ttl)
	return false
}
