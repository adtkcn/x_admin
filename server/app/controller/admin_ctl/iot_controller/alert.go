package iot_controller

import (
	"x_admin/app/schema/iot_schema"
	"x_admin/app/service/iot_service"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

// AlertHandler 告警历史控制器（数据由处理层告警状态机自动写入，管理端只读/清理）
type AlertHandler struct{}

// @Summary		告警历史列表
// @Description	分页查询设备告警事件（触发→恢复为一条记录），可按设备/编号/产品/区域/状态/时间区间过滤
// @Tags			iot_alert-告警
// @Param			token		header	string										true	"token"
// @Param			pageNo		query	int											false	"页码"
// @Param			pageSize	query	int											false	"每页数量"
// @Param			device_id	query	string										false	"设备ID"
// @Param			sn			query	string										false	"设备编号"
// @Param			product_id	query	string										false	"产品ID"
// @Param			space_id	query	string										false	"区域ID"
// @Param			status		query	int8										false	"状态: -1全部 0已恢复 1告警中"
// @Param			severity	query		int8									false	"级别: -1全部 1轻微 2严重 3故障"
// @Param			start_time	query	string										false	"触发时间起"
// @Param			end_time	query	string										false	"触发时间止"
// @Success		200			{object}	response.Response{data=response.PageResp{lists=iot_schema.AlertResp}}	"成功"
// @Router			/api/admin/iot/alert/list [get]
func (h AlertHandler) List(c *gin.Context) {
	var page request.PageReq
	var req iot_schema.AlertListReq
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &req)) {
		return
	}
	res, err := iot_service.AlertService.List(page, req)
	response.JSON(c, res, err)
}

// @Summary		告警历史删除
// @Description	批量删除告警历史记录
// @Tags			iot_alert-告警
// @Param			token	header	string					true	"token"
// @Param			data	body	iot_schema.AlertDelReq	true	"删除参数"
// @Success		200		{object}	response.Response		"成功"
// @Router			/api/admin/iot/alert/del [post]
func (h AlertHandler) Del(c *gin.Context) {
	var req iot_schema.AlertDelReq
	if response.IsFail(c, util.VerifyUtil.VerifyBody(c, &req)) {
		return
	}
	response.JSON(c, nil, iot_service.AlertService.Del(req.IDs))
}
