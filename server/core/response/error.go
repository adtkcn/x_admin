package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 本文件的函数会真实修改 HTTP 状态码，仅用于「非业务 JSON 响应」场景：
// 文件流/静态资源/网关探测等（例如图片流 404、S3 分片协议）。
//
// 普通业务接口请统一使用 response.go 中的 JSON / IsFail / Fail / FailMsg，
// 它们返回的 HTTP 状态恒为 200，成败一律以 body.code 表达。

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

// NoRoute 无路由响应
func NoRoute(c *gin.Context) {
	HTTPError(c, 404, "请求接口不存在")
}

// NotFound 404 错误
func NotFound(c *gin.Context, msg string) {
	HTTPError(c, http.StatusNotFound, msg)
}

// InternalError 500 错误
func InternalError(c *gin.Context, msg string) {
	HTTPError(c, http.StatusInternalServerError, msg)
}
