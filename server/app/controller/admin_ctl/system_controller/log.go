package system_controller

import (
	"x_admin/app/schema/system_schema"
	"x_admin/app/service/system_service"
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
	var logReq system_schema.SystemLogOperateReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &logReq)) {
		return
	}
	res, err := system_service.LogsService.OperateLog(page, logReq)
	response.CheckAndRespWithData(c, res, err)
}

// Login 登录日志
func (lh LogHandler) Login(c *gin.Context) {
	var page request.PageReq
	var logReq system_schema.SystemLogLoginReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &logReq)) {
		return
	}
	res, err := system_service.LogsService.LoginLog(page, logReq)
	response.CheckAndRespWithData(c, res, err)
}
