package iot_controller

import (
	"x_admin/app/schema/iot_schema"
	"x_admin/app/service/iot_service"
	iot_pipeline "x_admin/app/service/iot_service/pipeline"
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
// @Param			token		header		string																	true	"token"
// @Param			pageNo		query		int																		false	"页码"
// @Param			pageSize	query		int																		false	"每页数量"
// @Param			name		query		string																	false	"设备名称"
// @Param			sn			query		string																	false	"设备编号"
// @Param			product_id	query		string																	false	"产品ID"
// @Param			space_id	query		string																	false	"区域ID(含子区域)"
// @Param			status		query		string																	false	"在线状态: 0离线 1在线，不传=全部"
// @Param			start_time	query		string																	false	"创建时间起(2006-01-02 15:04:05)"
// @Param			end_time	query		string																	false	"创建时间止(2006-01-02 15:04:05)"
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
// @Param			token	header		string											true	"token"
// @Param			id		query		string											true	"主键"
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
// @Success		200		{object}	response.Response		"成功"
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
// @Param			token	header		string						true	"token"
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
// @Success		200		{object}	response.Response		"成功"
// @Router			/api/admin/iot/device/del [post]
func (h DeviceHandler) Del(c *gin.Context) {
	var req iot_schema.DeviceDelReq
	if response.IsFail(c, util.VerifyUtil.VerifyBody(c, &req)) {
		return
	}
	e := iot_service.DeviceService.Del(req.ID)
	if e == nil {
		// 设备删除后同步清理 Redis 运行态（alive/state/seen/watch/alert），
		// 否则幽灵设备在标记 TTL 内仍会被回读为在线；service 反向引用 pipeline
		// 会成包循环，故由同时依赖两边的控制器编排这一步。
		// 库里仍“告警中”的记录一并收尾，避免残留永不恢复的悬空告警。
		iot_pipeline.StopRuntime(req.ID)
		_ = iot_service.AlertService.CloseActiveByDevice(req.ID)
	}
	response.JSON(c, nil, e)
}

// @Summary		设备运行态快照
// @Description	回读 Redis 中的在线判定/最后报文时间与最新属性快照，用于验证接入→处理→运行态链路
// @Tags			iot_device-设备
// @Param			token	header		string												true	"token"
// @Param			id		query		string												true	"设备主键"
// @Success		200		{object}	response.Response{data=iot_schema.DeviceStateResp}	"成功"
// @Router			/api/admin/iot/device/state [get]
func (h DeviceHandler) State(c *gin.Context) {
	var req iot_schema.DeviceStateReq
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &req)) {
		return
	}
	detail, err := iot_service.DeviceService.Detail(req.ID)
	if response.IsFail(c, err) {
		return
	}
	// 运行态在 pipeline（Redis），库侧状态在 iot_service；pipeline 依赖 iot_service，
	// 故由控制器合并两侧数据，避免 service 反向引用接入层造成包循环
	// 生效的离线判定时长以设备记录为准（未配置时落兜底值），与定时扫描同源，改完配置即时反映
	timeout, err := iot_service.DeviceService.EffectiveOfflineTimeout(req.ID)
	if response.IsFail(c, err) {
		return
	}
	res := iot_pipeline.Runtime(req.ID, timeout)
	res.Sn = detail.Sn
	res.Name = detail.Name
	res.DBStatus = detail.Status
	res.LastOnline = detail.LastOnlineTime
	res.LastOffline = detail.LastOfflineTime
	response.JSON(c, res, nil)
}

// @Summary		设备时序数据回读
// @Description	按时间区间读取属性上报历史（本轮存 MySQL，后续换 TDengine 接口不变）
// @Tags			iot_device-设备
// @Param			token		header		string												true	"token"
// @Param			id			query		string												true	"设备主键"
// @Param			limit		query		int													false	"条数(缺省1000,上限5000)"
// @Param			start_ms	query		int64												false	"开始时间(毫秒),0=不限"
// @Param			end_ms		query		int64												false	"结束时间(毫秒),0=不限"
// @Success		200			{object}	response.Response{data=[]iot_schema.DeviceDataResp}	"成功"
// @Router			/api/admin/iot/device/data [get]
func (h DeviceHandler) Data(c *gin.Context) {
	var req iot_schema.DeviceDataReq
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &req)) {
		return
	}
	if _, err := iot_service.DeviceService.Detail(req.ID); response.IsFail(c, err) {
		return
	}
	response.JSON(c, iot_pipeline.QueryRange(req.ID, req.StartMs, req.EndMs, req.Limit), nil)
}

// @Summary		大屏地图设备点
// @Description	不分页返回带坐标的设备点（含在线状态与活动告警级别 0=无 -1=故障 1=轻微 2=严重），实时增量由 WS 房间 iot:map:{spaceId} 广播
// @Tags			iot_device-设备
// @Param			token		header		string												true	"token"
// @Param			space_ids	query		string												false	"区域ID列表(逗号分隔),缺省全部"
// @Success		200			{object}	response.Response{data=[]iot_schema.DeviceMapItem}	"成功"
// @Router			/api/admin/iot/device/map [get]
func (h DeviceHandler) Map(c *gin.Context) {
	var req iot_schema.DeviceMapReq
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &req)) {
		return
	}
	res, err := iot_service.DeviceService.MapList(req.SpaceIds)
	response.JSON(c, res, err)
}
