package iot_controller

import (
	"x_admin/app/schema/iot_schema"
	"x_admin/app/service/iot_service"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

// DeviceHandler 设备控制器
type DeviceHandler struct{}

// @Summary		设备列表
// @Description	分页获取设备列表
// @Tags			iot_device-设备
// @Param			token		header		string										true	"token"
// @Param			pageNo		query		int											false	"页码"
// @Param			pageSize	query		int											false	"每页数量"
// @Param			name		query		string										false	"设备名称"
// @Param			sn			query		string										false	"设备编号"
// @Param			product_id	query		string										false	"产品ID"
// @Param			space_id	query		string										false	"区域ID"
// @Param			status		query		int8										false	"在线状态: -1全部 0离线 1在线"
// @Success		200			{object}	response.Response{data=response.PageResp{lists=iot_schema.DeviceResp}}	"成功"
// @Router			/api/admin/iot/device/list [get]
func (h DeviceHandler) List(c *gin.Context) {
	var page request.PageReq
	var req iot_schema.DeviceListReq
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &req)) {
		return
	}
	res, err := iot_service.DeviceService.List(page, req)
	response.JSON(c, res, err)
}

// @Summary		设备详情
// @Description	获取设备详情
// @Tags			iot_device-设备
// @Param			token	header		string							true	"token"
// @Param			id		query		string							true	"主键"
// @Success		200		{object}	response.Response{data=iot_schema.DeviceResp}	"成功"
// @Router			/api/admin/iot/device/detail [get]
func (h DeviceHandler) Detail(c *gin.Context) {
	var req iot_schema.DeviceDetailReq
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &req)) {
		return
	}
	res, err := iot_service.DeviceService.Detail(req.ID)
	response.JSON(c, res, err)
}

// @Summary		设备新增
// @Description	新增设备(自动生成接入凭证)
// @Tags			iot_device-设备
// @Param			token	header		string					true	"token"
// @Param			data	body		iot_schema.DeviceAddReq	true	"新增参数"
// @Success		200		{object}	response.Response			"成功"
// @Router			/api/admin/iot/device/add [post]
func (h DeviceHandler) Add(c *gin.Context) {
	var req iot_schema.DeviceAddReq
	if response.IsFail(c, util.VerifyUtil.VerifyBody(c, &req)) {
		return
	}
	response.JSON(c, nil, iot_service.DeviceService.Add(req))
}

// @Summary		设备编辑
// @Description	编辑设备
// @Tags			iot_device-设备
// @Param			token	header		string					true	"token"
// @Param			data	body		iot_schema.DeviceEditReq	true	"编辑参数"
// @Success		200		{object}	response.Response			"成功"
// @Router			/api/admin/iot/device/edit [post]
func (h DeviceHandler) Edit(c *gin.Context) {
	var req iot_schema.DeviceEditReq
	if response.IsFail(c, util.VerifyUtil.VerifyBody(c, &req)) {
		return
	}
	response.JSON(c, nil, iot_service.DeviceService.Edit(req))
}

// @Summary		设备删除
// @Description	删除设备
// @Tags			iot_device-设备
// @Param			token	header		string					true	"token"
// @Param			data	body		iot_schema.DeviceDelReq	true	"删除参数"
// @Success		200		{object}	response.Response			"成功"
// @Router			/api/admin/iot/device/del [post]
func (h DeviceHandler) Del(c *gin.Context) {
	var req iot_schema.DeviceDelReq
	if response.IsFail(c, util.VerifyUtil.VerifyBody(c, &req)) {
		return
	}
	response.JSON(c, nil, iot_service.DeviceService.Del(req.ID))
}
