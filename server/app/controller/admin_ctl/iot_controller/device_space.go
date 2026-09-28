package iot_controller

import (
	"x_admin/app/schema/iot_schema"
	"x_admin/app/service/iot_service"
	"x_admin/core/response"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

// DeviceSpaceHandler 设备区域控制器
type DeviceSpaceHandler struct{}

// @Summary		设备区域所有
// @Description	获取所有设备区域(树形,不分页)
// @Tags			iot_deviceSpace-设备区域
// @Param			token	header		string																		true	"token"
// @Success		200		{object}	response.Response{data=[]iot_schema.DeviceSpaceResp}							"成功"
// @Router			/api/admin/iot/deviceSpace/all [get]
func (h DeviceSpaceHandler) All(c *gin.Context) {
	res, err := iot_service.DeviceSpaceService.All()
	response.JSON(c, res, err)
}

// @Summary		设备区域详情
// @Description	获取设备区域详情
// @Tags			iot_deviceSpace-设备区域
// @Param			token	header		string															true	"token"
// @Param			id		query		string															true	"主键"
// @Success		200		{object}	response.Response{data=iot_schema.DeviceSpaceResp}				"成功"
// @Router			/api/admin/iot/deviceSpace/detail [get]
func (h DeviceSpaceHandler) Detail(c *gin.Context) {
	var req iot_schema.DeviceSpaceDetailReq
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &req)) {
		return
	}
	res, err := iot_service.DeviceSpaceService.Detail(req.ID)
	response.JSON(c, res, err)
}

// @Summary		设备区域新增
// @Description	新增设备区域
// @Tags			iot_deviceSpace-设备区域
// @Param			token	header		string							true	"token"
// @Param			data	body		iot_schema.DeviceSpaceAddReq		true	"新增参数"
// @Success		200		{object}	response.Response					"成功"
// @Router			/api/admin/iot/deviceSpace/add [post]
func (h DeviceSpaceHandler) Add(c *gin.Context) {
	var req iot_schema.DeviceSpaceAddReq
	if response.IsFail(c, util.VerifyUtil.VerifyBody(c, &req)) {
		return
	}
	response.JSON(c, nil, iot_service.DeviceSpaceService.Add(req))
}

// @Summary		设备区域编辑
// @Description	编辑设备区域
// @Tags			iot_deviceSpace-设备区域
// @Param			token	header		string								true	"token"
// @Param			data	body		iot_schema.DeviceSpaceEditReq			true	"编辑参数"
// @Success		200		{object}	response.Response						"成功"
// @Router			/api/admin/iot/deviceSpace/edit [post]
func (h DeviceSpaceHandler) Edit(c *gin.Context) {
	var req iot_schema.DeviceSpaceEditReq
	if response.IsFail(c, util.VerifyUtil.VerifyBody(c, &req)) {
		return
	}
	response.JSON(c, nil, iot_service.DeviceSpaceService.Edit(req))
}

// @Summary		设备区域删除
// @Description	删除设备区域
// @Tags			iot_deviceSpace-设备区域
// @Param			token	header		string								true	"token"
// @Param			data	body		iot_schema.DeviceSpaceDelReq			true	"删除参数"
// @Success		200		{object}	response.Response						"成功"
// @Router			/api/admin/iot/deviceSpace/del [post]
func (h DeviceSpaceHandler) Del(c *gin.Context) {
	var req iot_schema.DeviceSpaceDelReq
	if response.IsFail(c, util.VerifyUtil.VerifyBody(c, &req)) {
		return
	}
	response.JSON(c, nil, iot_service.DeviceSpaceService.Del(req.ID))
}

// @Summary		设备区域排序
// @Description	同级设备区域拖拽排序
// @Tags			iot_deviceSpace-设备区域
// @Param			token	header		string								true	"token"
// @Param			data	body		iot_schema.DeviceSpaceSortReq			true	"排序参数"
// @Success		200		{object}	response.Response						"成功"
// @Router			/api/admin/iot/deviceSpace/sort [post]
func (h DeviceSpaceHandler) Sort(c *gin.Context) {
	var req iot_schema.DeviceSpaceSortReq
	if response.IsFail(c, util.VerifyUtil.VerifyBody(c, &req)) {
		return
	}
	response.JSON(c, nil, iot_service.DeviceSpaceService.Sort(req.Ids))
}
