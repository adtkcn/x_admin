package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PageResp 分页响应结构
type PageResp struct {
	Count    int64 `json:"count"`    // 总数
	PageNo   int   `json:"pageNo"`   // 当前页码
	PageSize int   `json:"pageSize"` // 每页数量
	Lists    any   `json:"lists"`    // 数据列表
}

// Send 底层发送函数（HTTP 状态恒为 200，成败由 body.code 表达）
func Send(c *gin.Context, code int, msg string, data any) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: msg,
		Data:    data,
	})
}

// Ok 成功响应快捷函数
func Ok(c *gin.Context, data ...any) {
	var respData any
	if len(data) > 0 {
		respData = data[0]
	}
	Send(c, Success.Code(), Success.Msg(), respData)
}
