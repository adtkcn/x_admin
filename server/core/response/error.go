package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NoRoute 无路由响应
func NoRoute(c *gin.Context) {
	Send(c, 404, "请求接口不存在", nil)
}

// NoMethod 无方法响应
func NoMethod(c *gin.Context) {
	Send(c, 405, "请求方法不允许", nil)
}

// HTTPError HTTP 错误响应
func HTTPError(c *gin.Context, status int, msg string) {
	c.JSON(status, Response{
		Code:    status,
		Message: msg,
		Data:    nil,
	})
}

// BadRequest 400 错误
func BadRequest(c *gin.Context, msg string) {
	HTTPError(c, http.StatusBadRequest, msg)
}

// Unauthorized 401 错误
func Unauthorized(c *gin.Context, msg string) {
	HTTPError(c, http.StatusUnauthorized, msg)
}

// Forbidden 403 错误
func Forbidden(c *gin.Context, msg string) {
	HTTPError(c, http.StatusForbidden, msg)
}

// NotFound 404 错误
func NotFound(c *gin.Context, msg string) {
	HTTPError(c, http.StatusNotFound, msg)
}

// InternalError 500 错误
func InternalError(c *gin.Context, msg string) {
	HTTPError(c, http.StatusInternalServerError, msg)
}
