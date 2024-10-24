package monitor_error

import (
	"net/http"
	"strconv"
	"strings"
	"time"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util"
	"x_admin/util/excel2"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/singleflight"
)

type MonitorErrorHandler struct {
	requestGroup singleflight.Group
}

//	@Summary	监控-错误列列表
//	@Tags		monitor_error-监控-错误列
//	@Produce	json
//	@Param		Token			header		string																	true	"token"
//	@Param		PageNo			query		int																		true	"页码"
//	@Param		PageSize		query		int																		true	"每页数量"
//	@Param		ProjectKey		query		string																	false	"项目key"
//	@Param		EventType		query		string																	false	"事件类型"
//	@Param		Path			query		string																	false	"URL地址"
//	@Param		Message			query		string																	false	"错误消息"
//	@Param		Stack			query		string																	false	"错误堆栈"
//	@Param		Md5				query		string																	false	"md5"
//	@Param		CreateTimeStart	query		string																	false	"创建时间"
//	@Param		CreateTimeEnd	query		string																	false	"创建时间"
//
//	@Success	200				{object}	response.Response{ data=response.PageResp{ lists=[]MonitorErrorResp}}	"成功"
//	@Router		/api/admin/monitor_error/list [get]
func (hd *MonitorErrorHandler) List(c *gin.Context) {
	var page request.PageReq
	var listReq MonitorErrorListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := MonitorErrorService.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

//	@Summary	监控-错误列列表-所有
//	@Tags		monitor_error-监控-错误列
//	@Produce	json
//	@Param		ProjectKey		query		string										false	"项目key"
//	@Param		EventType		query		string										false	"事件类型"
//	@Param		Path			query		string										false	"URL地址"
//	@Param		Message			query		string										false	"错误消息"
//	@Param		Stack			query		string										false	"错误堆栈"
//	@Param		Md5				query		string										false	"md5"
//	@Param		CreateTimeStart	query		string										false	"创建时间"
//	@Param		CreateTimeEnd	query		string										false	"创建时间"
//	@Success	200				{object}	response.Response{ data=[]MonitorErrorResp}	"成功"
//	@Router		/api/admin/monitor_error/listAll [get]
func (hd *MonitorErrorHandler) ListAll(c *gin.Context) {
	var listReq MonitorErrorListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := MonitorErrorService.ListAll(listReq)
	response.CheckAndRespWithData(c, res, err)
}

//	@Summary	监控-错误列详情
//	@Tags		monitor_error-监控-错误列
//	@Produce	json
//	@Param		Token	header		string										true	"token"
//	@Param		Id		query		number										false	"错误id"
//	@Success	200		{object}	response.Response{ data=MonitorErrorResp}	"成功"
//	@Router		/api/admin/monitor_error/detail [get]
func (hd *MonitorErrorHandler) Detail(c *gin.Context) {
	var detailReq MonitorErrorDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err, _ := hd.requestGroup.Do("MonitorError:Detail:"+strconv.Itoa(detailReq.Id), func() (any, error) {
		v, err := MonitorErrorService.Detail(detailReq.Id)
		return v, err
	})

	response.CheckAndRespWithData(c, res, err)
}

//	@Summary	监控-错误列新增
//	@Tags		monitor_error-监控-错误列
//	@Produce	json
//	@Param		Token		header		string				true	"token"
//	@Param		ProjectKey	body		string				false	"项目key"
//	@Param		EventType	body		string				false	"事件类型"
//	@Param		Path		body		string				false	"URL地址"
//	@Param		Message		body		string				false	"错误消息"
//	@Param		Stack		body		string				false	"错误堆栈"
//	@Param		Md5			body		string				false	"md5"
//	@Success	200			{object}	response.Response	"成功"
//	@Router		/api/admin/monitor_error/add [post]
func (hd *MonitorErrorHandler) Add(c *gin.Context) {
	var addReq MonitorErrorAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	createId, e := MonitorErrorService.Add(addReq)
	response.CheckAndRespWithData(c, createId, e)
}

//	@Summary	监控-错误列删除
//	@Tags		monitor_error-监控-错误列
//	@Produce	json
//	@Param		Token	header		string				true	"token"
//	@Param		Id		body		number				false	"错误id"
//	@Success	200		{object}	response.Response	"成功"
//	@Router		/api/admin/monitor_error/del [post]
func (hd *MonitorErrorHandler) Del(c *gin.Context) {
	var delReq MonitorErrorDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, MonitorErrorService.Del(delReq.Id))
}

//	@Summary	监控-错误列删除-批量
//	@Tags		monitor_error-监控-错误列
//
//	@Produce	json
//	@Param		Token	header		string				true	"token"
//	@Param		Ids		body		string				false	"逗号分割的id"
//	@Success	200		{object}	response.Response	"成功"
//	@Router		/api/admin/monitor_error/delBatch [post]
func (hd *MonitorErrorHandler) DelBatch(c *gin.Context) {
	var delReq MonitorErrorDelBatchReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	if delReq.Ids == "" {
		response.FailWithMsg(c, response.SystemError, "请选择要删除的数据")
		return
	}
	var Ids = strings.Split(delReq.Ids, ",")

	response.CheckAndResp(c, MonitorErrorService.DelBatch(Ids))
}

//	@Summary	监控-错误列导出
//	@Tags		monitor_error-监控-错误列
//	@Produce	json
//	@Param		Token			header	string	true	"token"
//	@Param		ProjectKey		query	string	false	"项目key"
//	@Param		EventType		query	string	false	"事件类型"
//	@Param		Path			query	string	false	"URL地址"
//	@Param		Message			query	string	false	"错误消息"
//	@Param		Stack			query	string	false	"错误堆栈"
//	@Param		Md5				query	string	false	"md5"
//	@Param		CreateTimeStart	query	string	false	"创建时间"
//	@Param		CreateTimeEnd	query	string	false	"创建时间"
//	@Router		/api/admin/monitor_error/ExportFile [get]
func (hd *MonitorErrorHandler) ExportFile(c *gin.Context) {
	var listReq MonitorErrorListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := MonitorErrorService.ExportFile(listReq)
	if err != nil {
		response.FailWithMsg(c, response.SystemError, "查询信息失败")
		return
	}
	f, err := excel2.Export(res, MonitorErrorService.GetExcelCol(), "Sheet1", "监控-错误列")
	if err != nil {
		response.FailWithMsg(c, response.SystemError, "导出失败")
		return
	}
	excel2.DownLoadExcel("监控-错误列"+time.Now().Format("20060102-150405"), c.Writer, f)
}

//	@Summary	监控-错误列导入
//	@Tags		monitor_error-监控-错误列
//	@Produce	json
//	@Router		/api/admin/monitor_error/ImportFile [post]
func (hd *MonitorErrorHandler) ImportFile(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusInternalServerError, "文件不存在")
		return
	}
	defer file.Close()
	importList := []MonitorErrorResp{}
	err = excel2.GetExcelData(file, &importList, MonitorErrorService.GetExcelCol())
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	err = MonitorErrorService.ImportFile(importList)
	response.CheckAndResp(c, err)
}
