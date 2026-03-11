package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Cors CORS（跨域资源共享）中间件 - 默认配置
func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
		AllowMethods: []string{"OPTIONS", "GET", "POST", "DELETE", "PUT"},
		MaxAge:       1 * time.Hour,
	})
}

// CorsStrict 严格的CORS（跨域资源共享）中间件
func CorsStrict() gin.HandlerFunc {
	return cors.New(cors.Config{
		// 只允许特定的来源，从配置文件中读取
		// AllowOrigins: []string{""},
		// 只允许必要的头部
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
			"token", // 项目使用的token头部
		},
		// 只允许必要的HTTP方法
		AllowMethods: []string{
			"OPTIONS",
			"GET",
			"POST",
			"DELETE",
			"PUT",
		},
		// 允许携带凭证（如cookie）
		AllowCredentials: true,
		// 预检请求的有效期
		MaxAge: 1 * time.Hour,
	})
}
