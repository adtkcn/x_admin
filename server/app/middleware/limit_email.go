package middleware

import (
	"bytes"
	"encoding/json/v2"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// EmailCodeEmailLimit 基于邮箱的验证码发送限流（会读取 body 提取 email）
// limit: 窗口内最大次数  ttl: 窗口秒数
//
// 使用示例：
//
//	r.POST("/sendCode", middleware.EmailLimit(2, 60), handler)
func LimitEmail(limit int, ttl int) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := extractEmail(c)
		if email == "" {
			c.Next()
			return
		}

		key := "limit:email:" + email
		if exceeded(key, limit, ttl) {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"code":    429,
				"message": "该邮箱验证码发送过于频繁，请稍后重试",
			})
			return
		}
		c.Next()
	}
}

// extractEmail 从请求 body 中提取 email 字段，并恢复 body 供后续 handler 使用
func extractEmail(c *gin.Context) string {
	bodyBytes, err := c.GetRawData()
	if err != nil {
		return ""
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	var req struct {
		Email string `json:"email"`
	}
	if json.Unmarshal(bodyBytes, &req) != nil {
		return ""
	}
	return req.Email
}
