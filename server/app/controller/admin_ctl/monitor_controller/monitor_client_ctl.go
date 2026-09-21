package monitor_controller

import (
	"encoding/json/v2"
	"net/url"
	"strings"
	"time"
	"x_admin/app/schema/monitor_schema"

	"x_admin/app/service/monitor_service"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util"
	"x_admin/util/excel2"
	"x_admin/util/img_util"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/singleflight"
)

type MonitorClientHandler struct {
	requestGroup singleflight.Group
}

// @Summary	监控-客户端信息列表
// @Tags		monitor_client-监控-客户端信息
// @Produce	json
// @Param		token				header		string																					true	"token"
// @Param		page_no				query		int																						true	"页码"
// @Param		page_size			query		int																						true	"每页数量"
// @Param		project_key			query		string																					false	"项目key"
// @Param		client_id			query		string																					false	"sdk生成的客户端id"
// @Param		os					query		string																					false	"系统"
// @Param		browser				query		string																					false	"浏览器"
// @Param		ua					query		string																					false	"ua记录"
// @Param		create_time_start	query		string																					false	"开始创建时间"
// @Param		create_time_end		query		string																					false	"结束创建时间"
// @Success	200					{object}	response.Response{ data=response.PageResp{ lists=[]monitor_schema.MonitorClientResp}}	"成功"
// @Router		/api/admin/monitor_client/list [get]
func (hd *MonitorClientHandler) List(c *gin.Context) {
	var page request.PageReq
	var listReq monitor_schema.MonitorClientListReq
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := monitor_service.MonitorClientService.List(page, listReq)
	response.JSON(c, res, err)
}

// @Summary	监控-客户端信息列表-所有
// @Tags		monitor_client-监控-客户端信息
// @Produce	json
// @Param		project_key			query		string														false	"项目key"
// @Param		client_id			query		string														false	"sdk生成的客户端id"
// @Param		os					query		string														false	"系统"
// @Param		browser				query		string														false	"浏览器"
// @Param		ua					query		string														false	"ua记录"
// @Param		create_time_start	query		string														false	"开始创建时间"
// @Param		create_time_end		query		string														false	"结束创建时间"
// @Success	200					{object}	response.Response{ data=[]monitor_schema.MonitorClientResp}	"成功"
// @Router		/api/admin/monitor_client/list_all [get]
func (hd *MonitorClientHandler) ListAll(c *gin.Context) {
	var listReq monitor_schema.MonitorClientListReq
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := monitor_service.MonitorClientService.ListAll(listReq)
	response.JSON(c, res, err)
}

// @Summary	监控-客户端信息错误用户列表
// @Tags		monitor_client-监控-客户端信息
// @Produce	json
// @Param		token	header		string				true	"token"
// @Param		id		query		string				false	"uuid"
// @Success	200		{object}	response.Response	"成功"
// @Router		/api/admin/monitor_client/error_users [get]
func (hd *MonitorClientHandler) ErrorUsers(c *gin.Context) {
	var Req monitor_schema.MonitorClientDetailReq
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &Req)) {
		return
	}
	res, err := monitor_service.MonitorClientService.ErrorUsers(Req.Id)
	response.JSON(c, res, err)
}

// @Summary	监控-客户端信息详情
// @Tags		monitor_client-监控-客户端信息
// @Produce	json
// @Param		token	header		string														true	"token"
// @Param		id		query		string														false	"uuid"
// @Success	200		{object}	response.Response{ data=monitor_schema.MonitorClientResp}	"成功"
// @Router		/api/admin/monitor_client/detail [get]
func (hd *MonitorClientHandler) Detail(c *gin.Context) {
	var detailReq monitor_schema.MonitorClientDetailReq
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err, _ := hd.requestGroup.Do("MonitorClient:Detail:"+detailReq.Id, func() (any, error) {
		v, err := monitor_service.MonitorClientService.Detail(detailReq.Id)
		return v, err
	})

	response.JSON(c, res, err)
}

// @Summary	监控-客户端信息新增
// @Tags		monitor_client-监控-客户端信息
// @Produce	json
// @Param		token		header		string				true	"token"
// @Param		project_key	body		string				false	"项目key"
// @Param		client_id	body		string				false	"sdk生成的客户端id"
// @Param		os			body		string				false	"系统"
// @Param		browser		body		string				false	"浏览器"
// @Param		ua			body		string				false	"ua记录"
// @Success	200			{object}	response.Response	"成功"
// @Router		/api/admin/monitor_client/add [post]
func (hd *MonitorClientHandler) Add(c *gin.Context) {
	data, err := url.QueryUnescape(c.Query("data"))
	if err != nil {
		c.Data(200, "image/gif", img_util.EmptyGif())
		return
	}
	var addReq monitor_schema.MonitorClientAddReq
	json.Unmarshal([]byte(data), &addReq)

	uaStr := c.GetHeader("user-agent")
	if uaStr != "" {
		ua := util.UAUtils.Parse(uaStr)
		addReq.Ua = uaStr
		addReq.Os = ua.OsName
		addReq.Browser = ua.BrowserName
	}

	monitor_service.MonitorClientService.Add(addReq)

	c.Data(200, "image/gif", img_util.EmptyGif())
}

// @Summary	监控-客户端信息删除
// @Tags		monitor_client-监控-客户端信息
// @Produce	json
// @Param		token	header		string				true	"token"
// @Param		id		body		string				false	"uuid"
// @Success	200		{object}	response.Response	"成功"
// @Router		/api/admin/monitor_client/del [post]
func (hd *MonitorClientHandler) Del(c *gin.Context) {
	var delReq monitor_schema.MonitorClientDelReq
	if response.IsFail(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.JSON(c, nil, monitor_service.MonitorClientService.Del(delReq.Id))
}

// @Summary	监控-客户端信息删除-批量
// @Tags		monitor_client-监控-客户端信息
//
// @Produce	json
// @Param		token	header		string				true	"token"
// @Param		ids		body		string				false	"逗号分割的id"
// @Success	200		{object}	response.Response	"成功"
// @Router		/api/admin/monitor_client/del_batch [post]
func (hd *MonitorClientHandler) DelBatch(c *gin.Context) {
	var delReq monitor_schema.MonitorClientDelBatchReq
	if response.IsFail(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	if delReq.Ids == "" {
		response.FailMsg(c, "请选择要删除的数据")
		return
	}
	var Ids = strings.Split(delReq.Ids, ",")

	response.JSON(c, nil, monitor_service.MonitorClientService.DelBatch(Ids))
}

// @Summary	监控-客户端信息导出
// @Tags		monitor_client-监控-客户端信息
// @Produce	json
// @Param		token				header	string	true	"token"
// @Param		project_key			query	string	false	"项目key"
// @Param		client_id			query	string	false	"sdk生成的客户端id"
// @Param		os					query	string	false	"系统"
// @Param		browser				query	string	false	"浏览器"
// @Param		ua					query	string	false	"ua记录"
// @Param		create_time_start	query	string	false	"开始创建时间"
// @Param		create_time_end		query	string	false	"结束创建时间"
// @Router		/api/admin/monitor_client/export_file [get]
func (hd *MonitorClientHandler) ExportFile(c *gin.Context) {
	var listReq monitor_schema.MonitorClientListReq
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := monitor_service.MonitorClientService.ExportFile(listReq)
	if err != nil {
		response.Fail(c, response.CheckErr(err, "查询信息失败"))
		return
	}
	f, err := excel2.Export(res, monitor_service.MonitorClientService.GetExcelCol(), "Sheet1", "监控-客户端信息")
	if err != nil {
		response.Fail(c, response.CheckErr(err, "导出失败"))
		return
	}
	excel2.DownLoadExcel("监控-客户端信息"+time.Now().Format("20060102-150405"), c.Writer, f)
}

// @Summary	监控-客户端信息导入
// @Tags		monitor_client-监控-客户端信息
// @Produce	json
// @Router		/api/admin/monitor_client/import_file [post]
func (hd *MonitorClientHandler) ImportFile(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		response.Fail(c, response.CheckErr(err, "文件不存在"))
		return
	}
	defer file.Close()
	importList := []monitor_schema.MonitorClientResp{}
	err = excel2.GetExcelData(file, &importList, monitor_service.MonitorClientService.GetExcelCol())
	if err != nil {
		response.Fail(c, response.CheckErr(err, "文件解析失败"))
		return
	}

	err = monitor_service.MonitorClientService.ImportFile(importList)
	response.JSON(c, nil, err)
}
