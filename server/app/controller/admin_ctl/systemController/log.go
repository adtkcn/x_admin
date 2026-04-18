package systemController

import (
	"x_admin/app/schema/systemSchema"
	"x_admin/app/service/systemService"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

// LogHandler 日志控制器
type LogHandler struct{}

// Operate 操作日志
func (lh LogHandler) Operate(c *gin.Context) {
	var page request.PageReq
	var logReq systemSchema.SystemLogOperateReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &logReq)) {
		return
	}
	res, err := systemService.LogsService.OperateLog(page, logReq)
	response.CheckAndRespWithData(c, res, err)
}

// Login 登录日志
func (lh LogHandler) Login(c *gin.Context) {
	var page request.PageReq
	var logReq systemSchema.SystemLogLoginReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &logReq)) {
		return
	}
	res, err := systemService.LogsService.LoginLog(page, logReq)
	response.CheckAndRespWithData(c, res, err)
}
