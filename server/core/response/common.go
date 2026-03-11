package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PageResp 分页响应结构
type PageResp struct {
	Count    int64       `json:"count"`    // 总数
	PageNo   int         `json:"pageNo"`   // 当前页码
	PageSize int         `json:"pageSize"` // 每页数量
	Lists    interface{} `json:"lists"`    // 数据列表
}

// ========== 快捷响应函数 ==========
// Send 发送响应的内部函数
func Send(c *gin.Context, code int, msg string, data interface{}) {
	status := http.StatusOK
	// if code >= 500 {
	// 	status = http.StatusInternalServerError
	// }

	c.JSON(status, Response{
		Code:    code,
		Message: msg,
		Data:    data,
	})
}

// Ok 成功响应快捷函数
func Ok(c *gin.Context, data ...interface{}) {
	var respData interface{}
	if len(data) > 0 {
		respData = data[0]
	}
	Send(c, 200, "成功", respData)
}

// Fail 失败响应快捷函数（支持字符串消息）
func Fail(c *gin.Context, msg string) {
	Send(c, 300, msg, nil)
}

// FailWithResp 失败响应快捷函数（支持 RespType - 兼容旧代码）
func FailWithResp(c *gin.Context, resp RespType) {
	Send(c, resp.Code(), resp.Msg(), resp.Data())
}
