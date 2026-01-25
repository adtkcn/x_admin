package systemController

import (
	. "x_admin/app/schema/systemSchema"
	"x_admin/app/service/systemService"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/middleware"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

func LogRoute(rg *gin.RouterGroup) {
	handle := logHandler{}

	rg = rg.Group("/system", middleware.TokenAuth())
	rg.GET("/log/operate", handle.operate)
	rg.GET("/log/login", handle.login)
}

type logHandler struct {
	// Service ISystemLogsServer
}

// operate 操作日志
func (lh logHandler) operate(c *gin.Context) {
	var page request.PageReq
	var logReq SystemLogOperateReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &logReq)) {
		return
	}
	res, err := systemService.LogsService.Operate(page, logReq)
	response.CheckAndRespWithData(c, res, err)
}

// login 登录日志
func (lh logHandler) login(c *gin.Context) {
	var page request.PageReq
	var logReq SystemLogLoginReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &logReq)) {
		return
	}
	res, err := systemService.LogsService.Login(page, logReq)
	response.CheckAndRespWithData(c, res, err)
}
