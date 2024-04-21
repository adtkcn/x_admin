package monitor_web

import (
	"fmt"
	"net/http"
	"time"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util"
	"x_admin/util/excel"

	"github.com/gin-gonic/gin"
)

type MonitorWebHandler struct{}

//	@Summary	错误收集error列表
//	@Tags		monitor_web-错误收集error
//	@Produce	json
//	@Param		Token		header		string				true	"token"
//	@Param		PageNo		query		int					true	"页码"
//	@Param		PageSize	query		int					true	"每页数量"
//	@Param		projectKey	query		string				false	"项目key."
//	@Param		clientId	query		string				false	"sdk生成的客户端id."
//	@Param		eventType	query		string				false	"事件类型."
//	@Param		page		query		string				false	"URL地址."
//	@Param		message		query		string				false	"错误消息."
//	@Param		stack		query		string				false	"错误堆栈."
//	@Param		clientTime	query		int					false	"客户端时间."
//	@Success	200			{object}	[]MonitorWebResp	"成功"
//	@Failure	400			{object}	string				"请求错误"
//	@Router		/api/admin/monitor_web/list [get]
func (hd MonitorWebHandler) List(c *gin.Context) {
	var page request.PageReq
	var listReq MonitorWebListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := Service.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

//	@Summary	错误收集error列表-所有
//	@Tags		monitor_web-错误收集error
//	@Produce	json
//	@Success	200	{object}	[]MonitorWebResp	"成功"
//	@Router		/api/admin/monitor_web/listAll [get]
func (hd MonitorWebHandler) ListAll(c *gin.Context) {
	//var listReq MonitorWebListReq
	//if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
	//	return
	//}
	res, err := Service.ListAll()

	response.CheckAndRespWithData(c, res, err)
}

//	@Summary	错误收集error详情
//	@Tags		monitor_web-错误收集error
//	@Produce	json
//	@Param		Token	header		string			true	"token"
//	@Param		id		query		int				false	"uuid."
//	@Success	200		{object}	MonitorWebResp	"成功"
//	@Router		/api/admin/monitor_web/detail [get]
func (hd MonitorWebHandler) Detail(c *gin.Context) {
	var detailReq MonitorWebDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := Service.Detail(detailReq.Id)
	response.CheckAndRespWithData(c, res, err)
}

//	@Summary	错误收集error新增
//	@Tags		monitor_web-错误收集error
//	@Produce	json
//	@Param		Token		header		string				true	"token"
//	@Param		projectKey	body		string				false	"项目key."
//	@Param		clientId	body		string				false	"sdk生成的客户端id."
//	@Param		eventType	body		string				false	"事件类型."
//	@Param		page		body		string				false	"URL地址."
//	@Param		message		body		string				false	"错误消息."
//	@Param		stack		body		string				false	"错误堆栈."
//	@Param		clientTime	body		int					false	"客户端时间."
//	@Success	200			{object}	response.RespType	"成功"
//	@Router		/api/admin/monitor_web/add [post]
func (hd MonitorWebHandler) Add(c *gin.Context) {
	var addReq MonitorWebAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	response.CheckAndResp(c, Service.Add(addReq))
}

//	@Summary	错误收集error编辑
//	@Tags		monitor_web-错误收集error
//	@Produce	json
//	@Param		Token		header		string				true	"token"
//	@Param		id			body		int					false	"uuid."
//	@Param		projectKey	body		string				false	"项目key."
//	@Param		clientId	body		string				false	"sdk生成的客户端id."
//	@Param		eventType	body		string				false	"事件类型."
//	@Param		page		body		string				false	"URL地址."
//	@Param		message		body		string				false	"错误消息."
//	@Param		stack		body		string				false	"错误堆栈."
//	@Param		clientTime	body		int					false	"客户端时间."
//	@Success	200			{object}	response.RespType	"成功"
//	@Router		/api/admin/monitor_web/edit [post]
func (hd MonitorWebHandler) Edit(c *gin.Context) {
	var editReq MonitorWebEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	response.CheckAndResp(c, Service.Edit(editReq))
}

//	@Summary	错误收集error删除
//	@Tags		monitor_web-错误收集error
//	@Produce	json
//	@Param		Token	header		string				true	"token"
//	@Param		id		body		int					false	"uuid."
//	@Success	200		{object}	response.RespType	"成功"
//	@Router		/api/admin/monitor_web/del [post]
func (hd MonitorWebHandler) Del(c *gin.Context) {
	var delReq MonitorWebDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, Service.Del(delReq.Id))
}

//	@Summary	错误收集error导出
//	@Tags		monitor_web-错误收集error
//	@Produce	json
//	@Param		Token		header	string	true	"token"
//	@Param		projectKey	query	string	false	"项目key."
//	@Param		clientId	query	string	false	"sdk生成的客户端id."
//	@Param		eventType	query	string	false	"事件类型."
//	@Param		page		query	string	false	"URL地址."
//	@Param		message		query	string	false	"错误消息."
//	@Param		stack		query	string	false	"错误堆栈."
//	@Param		clientTime	query	int		false	"客户端时间."
//	@Router		/api/admin/monitor_web/ExportFile [get]
func (hd MonitorWebHandler) ExportFile(c *gin.Context) {
	var listReq MonitorWebListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := Service.ExportFile(listReq)
	if err != nil {
		response.FailWithMsg(c, response.SystemError, "查询信息失败")
		return
	}
	f, err := excel.NormalDynamicExport(res, "Sheet1", "错误收集error", "", true, false, nil)
	if err != nil {
		response.FailWithMsg(c, response.SystemError, "导出失败")
		return
	}
	excel.DownLoadExcel("错误收集error"+time.Now().Format("20060102-150405"), c.Writer, f)
}

//	@Summary	错误收集error导入
//	@Tags		monitor_web-错误收集error
//	@Produce	json
func (hd MonitorWebHandler) ImportFile(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusInternalServerError, "文件不存在")
		return
	}
	defer file.Close()
	importList := []MonitorWebResp{}
	err = excel.GetExcelData(file, &importList)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	for _, t := range importList {
		fmt.Printf("%#v", t)
	}
	err = Service.ImportFile(importList)
	response.CheckAndResp(c, err)
}
