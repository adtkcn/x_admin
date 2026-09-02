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

// @Summary		操作日志
// @Description	获取操作日志列表
// @Tags			system_log-日志
// @Param			token		header		string																				true	"token"
// @Param			pageNo		query		int																					true	"页码"
// @Param			pageSize	query		int																					true	"每页数量"
// @Param			email		query		string																				false	"邮箱(账号)"
// @Param			title		query		string																				false	"操作标题"
// @Param			type		query		string																				false	"请求类型: GET/POST/PUT"
// @Param			ip			query		string																				false	"请求IP"
// @Param			status		query		int																					false	"执行状态: [1=成功, 2=失败]"
// @Param			url			query		string																				false	"请求地址"
// @Param			start_time	query		string																				false	"开始时间"
// @Param			end_time	query		string																				false	"结束时间"
// @Success		200			{object}	response.Response{data=response.PageResp{lists=system_schema.SystemLogOperateResp}}	"成功"
// @Router			/api/admin/system/log/operate [get]
func (lh LogHandler) Operate(c *gin.Context) {
	var page request.PageReq
	var logReq system_schema.SystemLogOperateReq
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &logReq)) {
		return
	}
	res, err := system_service.LogsService.OperateLog(page, logReq)
	response.JSON(c, res, err)
}

// @Summary		登录日志
// @Description	获取登录日志列表
// @Tags			system_log-日志
// @Param			token		header		string																				true	"token"
// @Param			pageNo		query		int																					true	"页码"
// @Param			pageSize	query		int																					true	"每页数量"
// @Param			email		query		string																				false	"邮箱(账号)"
// @Param			status		query		int																					false	"执行状态: [1=成功, 2=失败]"
// @Param			start_time	query		string																				false	"开始时间"
// @Param			end_time	query		string																				false	"结束时间"
// @Success		200			{object}	response.Response{data=response.PageResp{lists=system_schema.SystemLogLoginResp}}	"成功"
// @Router			/api/admin/system/log/login [get]
func (lh LogHandler) Login(c *gin.Context) {
	var page request.PageReq
	var logReq system_schema.SystemLogLoginReq
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &logReq)) {
		return
	}
	res, err := system_service.LogsService.LoginLog(page, logReq)
	response.JSON(c, res, err)
}
