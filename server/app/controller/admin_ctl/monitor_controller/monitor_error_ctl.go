package monitor_controller

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"
	"x_admin/app/schema/monitor_schema"
	. "x_admin/app/schema/monitor_schema"
	"x_admin/app/service/monitor_service"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util"
	"x_admin/util/excel2"
	"x_admin/util/img_util"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/singleflight"
)

type MonitorErrorHandler struct {
	requestGroup singleflight.Group
}

// @Summary	监控-错误列列表
// @Tags		monitor_error-监控-错误列
// @Produce	json
// @Param		token			header		string																	true	"token"
// @Param		pageNo			query		int																		true	"页码"
// @Param		pageSize		query		int																		true	"每页数量"
// @Param		ProjectKey		query		string																	false	"项目key"
// @Param		EventType		query		string																	false	"事件类型"
// @Param		Path			query		string																	false	"URL地址"
// @Param		Message			query		string																	false	"错误消息"
// @Param		Stack			query		string																	false	"错误堆栈"
// @Param		Md5				query		string																	false	"md5"
// @Param		CreateTimeStart	query		string																	false	"创建时间"
// @Param		CreateTimeEnd	query		string																	false	"创建时间"
//
// @Success	200				{object}	response.Response{ data=response.PageResp{ lists=[]MonitorErrorResp}}	"成功"
// @Router		/api/admin/monitor_error/list [get]
func (hd *MonitorErrorHandler) List(c *gin.Context) {
	var page request.PageReq
	var listReq MonitorErrorListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := monitor_service.MonitorErrorService.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

// @Summary	监控-错误列列表-所有
// @Tags		monitor_error-监控-错误列
// @Produce	json
// @Param		ProjectKey		query		string										false	"项目key"
// @Param		EventType		query		string										false	"事件类型"
// @Param		Path			query		string										false	"URL地址"
// @Param		Message			query		string										false	"错误消息"
// @Param		Stack			query		string										false	"错误堆栈"
// @Param		Md5				query		string										false	"md5"
// @Param		CreateTimeStart	query		string										false	"创建时间"
// @Param		CreateTimeEnd	query		string										false	"创建时间"
// @Success	200				{object}	response.Response{ data=[]MonitorErrorResp}	"成功"
// @Router		/api/admin/monitor_error/listAll [get]
func (hd *MonitorErrorHandler) ListAll(c *gin.Context) {
	var listReq MonitorErrorListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := monitor_service.MonitorErrorService.ListAll(listReq)
	response.CheckAndRespWithData(c, res, err)
}

// @Summary	监控-错误列详情
// @Tags		monitor_error-监控-错误列
// @Produce	json
// @Param		token	header		string										true	"token"
// @Param		Id		query		number										false	"错误id"
// @Success	200		{object}	response.Response{ data=MonitorErrorResp}	"成功"
// @Router		/api/admin/monitor_error/detail [get]
func (hd *MonitorErrorHandler) Detail(c *gin.Context) {
	var detailReq MonitorErrorDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err, _ := hd.requestGroup.Do("MonitorError:Detail:"+detailReq.Id, func() (any, error) {
		v, err := monitor_service.MonitorErrorService.Detail(detailReq.Id)
		return v, err
	})

	response.CheckAndRespWithData(c, res, err)
}

// @Summary	监控-错误列新增
// @Tags		monitor_error-监控-错误列
// @Produce	json
// @Param		token		header		string				true	"token"
// @Param		ProjectKey	body		string				false	"项目key"
// @Param		EventType	body		string				false	"事件类型"
// @Param		Path		body		string				false	"URL地址"
// @Param		Message		body		string				false	"错误消息"
// @Param		Stack		body		string				false	"错误堆栈"
// @Param		Md5			body		string				false	"md5"
// @Success	200			{object}	response.Response	"成功"
// @Router		/api/admin/monitor_error/add [post]
func (hd *MonitorErrorHandler) Add(c *gin.Context) {
	data, err := url.QueryUnescape(c.Query("data"))
	if err != nil {
		c.Data(200, "image/gif", img_util.EmptyGif())
		return
	}

	var addReq []MonitorErrorAddReq
	json.Unmarshal([]byte(data), &addReq)

	ip := c.ClientIP()
	regionInfo := util.IpUtil.Parse(ip)

	for i := 0; i < len(addReq); i++ {
		var ListAddReq = monitor_schema.MonitorErrorListAddReq{
			ClientId: addReq[i].ClientId,
			UserId:   addReq[i].UserId,
			Width:    addReq[i].Width,
			Height:   addReq[i].Height,
		}
		// if ip != "" && ip != "127.0.0.1" {
		// regionInfo := util.IpUtil.Parse("118.24.157.190")

		ListAddReq.Ip = ip
		ListAddReq.City = regionInfo.City
		ListAddReq.Country = regionInfo.Country
		ListAddReq.Operator = regionInfo.Operator
		ListAddReq.Province = regionInfo.Province
		// }
		monitor_service.MonitorErrorService.Add(addReq[i], ListAddReq)
	}

	c.Data(200, "image/gif", img_util.EmptyGif())
}

// @Summary	监控-错误列删除
// @Tags		monitor_error-监控-错误列
// @Produce	json
// @Param		token	header		string				true	"token"
// @Param		Id		body		number				false	"错误id"
// @Success	200		{object}	response.Response	"成功"
// @Router		/api/admin/monitor_error/del [post]
func (hd *MonitorErrorHandler) Del(c *gin.Context) {
	var delReq MonitorErrorDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.CheckAndRespWithData(c, nil, monitor_service.MonitorErrorService.Del(delReq.Id))
}

// @Summary	监控-错误列删除-批量
// @Tags		monitor_error-监控-错误列
//
// @Produce	json
// @Param		token	header		string				true	"token"
// @Param		Ids		body		string				false	"逗号分割的id"
// @Success	200		{object}	response.Response	"成功"
// @Router		/api/admin/monitor_error/del_batch [post]
func (hd *MonitorErrorHandler) DelBatch(c *gin.Context) {
	var delReq MonitorErrorDelBatchReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	if delReq.Ids == "" {
		response.Fail(c, "请选择要删除的数据")
		return
	}
	var Ids = strings.Split(delReq.Ids, ",")

	response.CheckAndRespWithData(c, nil, monitor_service.MonitorErrorService.DelBatch(Ids))
}

// @Summary	监控-错误列导出
// @Tags		monitor_error-监控-错误列
// @Produce	json
// @Param		token			header	string	true	"token"
// @Param		ProjectKey		query	string	false	"项目key"
// @Param		EventType		query	string	false	"事件类型"
// @Param		Path			query	string	false	"URL地址"
// @Param		Message			query	string	false	"错误消息"
// @Param		Stack			query	string	false	"错误堆栈"
// @Param		Md5				query	string	false	"md5"
// @Param		CreateTimeStart	query	string	false	"创建时间"
// @Param		CreateTimeEnd	query	string	false	"创建时间"
// @Router		/api/admin/monitor_error/export_file [get]
func (hd *MonitorErrorHandler) ExportFile(c *gin.Context) {
	var listReq MonitorErrorListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := monitor_service.MonitorErrorService.ExportFile(listReq)
	if err != nil {
		response.Fail(c, "查询信息失败")
		return
	}
	f, err := excel2.Export(res, monitor_service.MonitorErrorService.GetExcelCol(), "Sheet1", "监控-错误列")
	if err != nil {
		response.Fail(c, "导出失败")
		return
	}
	excel2.DownLoadExcel("监控-错误列"+time.Now().Format("20060102-150405"), c.Writer, f)
}

// @Summary	监控-错误列导入
// @Tags		monitor_error-监控-错误列
// @Produce	json
// @Router		/api/admin/monitor_error/import_file [post]
func (hd *MonitorErrorHandler) ImportFile(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusInternalServerError, "文件不存在")
		return
	}
	defer file.Close()
	importList := []MonitorErrorResp{}
	err = excel2.GetExcelData(file, &importList, monitor_service.MonitorErrorService.GetExcelCol())
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	err = monitor_service.MonitorErrorService.ImportFile(importList)
	response.CheckAndRespWithData(c, nil, err)
}
