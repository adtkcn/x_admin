package middleware

import (
	"runtime/debug"
	"x_admin/core"
	"x_admin/core/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ErrorRecover 异常恢复中间件
func ErrorRecover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				switch v := r.(type) {
				// 业务异常：沿用其业务码与文案
				case response.RespType:
					core.Logger.WithOptions(zap.AddCallerSkip(2)).Warnf(
						"Request Fail by recover: url=[%s], resp=[%+v]", c.Request.URL.Path, v)

					response.Fail(c, v)
				// 其他类型：细节留在日志，对外统一系统错误
				default:
					core.Logger.Errorf("stacktrace from panic: %+v\n%s", r, string(debug.Stack()))
					response.Fail(c, response.SystemError)
				}
				c.Abort()
			}
		}()
		c.Next()
	}
}
