package system_log_sms

import (
	"net/http"
	"strconv"
	"time"
	"github.com/gin-gonic/gin" 
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util"
	"x_admin/util/excel"
	"golang.org/x/sync/singleflight"
)

 
type SystemLogSmsHandler struct {
	requestGroup singleflight.Group
}

//  @Summary	系统短信日志列表
//  @Tags		system_log_sms-系统短信日志
//  @Produce	json
//  @Param		Token		header		string				true	"token"
//  @Param		PageNo		query		int					true	"页码"
//  @Param		PageSize	query		int					true	"每页数量"
//	@Param scene query int false "场景编号"
//	@Param mobile query string false "手机号码"
//	@Param content query string false "发送内容"
//	@Param status query int false "发送状态：[0=发送中, 1=发送成功, 2=发送失败]"
//	@Param results query string false "短信结果"
//	@Param send_time query int false "发送时间"
//	@Param create_timeStart  query core.TsTime false "创建时间"
//	@Param create_timeEnd  query core.TsTime false "创建时间"
//	@Param update_timeStart  query core.TsTime false "更新时间"
//	@Param update_timeEnd  query core.TsTime false "更新时间"
//@Success 200	{object} response.Response{ data=response.PageResp{ lists= []SystemLogSmsResp}}	"成功"
//@Router	/api/admin/system_log_sms/list [get]
func (hd *SystemLogSmsHandler) List(c *gin.Context) {
	var page request.PageReq
	var listReq SystemLogSmsListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := SystemLogSmsService.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

//	@Summary	系统短信日志列表-所有
//	@Tags		system_log_sms-系统短信日志
//  @Produce	json
//	@Param scene query int false "场景编号"
//	@Param mobile query string false "手机号码"
//	@Param content query string false "发送内容"
//	@Param status query int false "发送状态：[0=发送中, 1=发送成功, 2=发送失败]"
//	@Param results query string false "短信结果"
//	@Param send_time query int false "发送时间"
//	@Param create_timeStart  query core.TsTime false "创建时间"
//	@Param create_timeEnd  query core.TsTime false "创建时间"
//	@Param update_timeStart  query core.TsTime false "更新时间"
//	@Param update_timeEnd  query core.TsTime false "更新时间"
//	@Success	200			{object}	response.Response{ data=[]SystemLogSmsResp}	"成功"
//	@Router		/api/admin/system_log_sms/listAll [get]
func (hd *SystemLogSmsHandler) ListAll(c *gin.Context) {
	var listReq SystemLogSmsListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := SystemLogSmsService.ListAll(listReq)
	response.CheckAndRespWithData(c, res, err)
}

//	@Summary	系统短信日志详情
//	@Tags		system_log_sms-系统短信日志
//	@Produce	json
//	@Param		Token		header		string				true	"token"
//	@Param		id		query		int				false	"id"
//	@Success	200			{object}	response.Response{ data=SystemLogSmsResp}	"成功"
//	@Router		/api/admin/system_log_sms/detail [get]
func (hd *SystemLogSmsHandler) Detail(c *gin.Context) {
	var detailReq SystemLogSmsDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err, _ := hd.requestGroup.Do("SystemLogSms:Detail:"+strconv.Itoa(detailReq.Id), func() (any, error) {
		v, err := SystemLogSmsService.Detail(detailReq.Id)
		return v, err
	})

	response.CheckAndRespWithData(c, res, err)
}


//	@Summary	系统短信日志新增
//	@Tags		system_log_sms-系统短信日志
//	@Produce	json
//	@Param		Token		header		string				true	"token"
//	@Param		scene		body		int				false	"场景编号"
//	@Param		mobile		body		string				false	"手机号码"
//	@Param		content		body		string				false	"发送内容"
//	@Param		status		body		int				false	"发送状态：[0=发送中, 1=发送成功, 2=发送失败]"
//	@Param		results		body		string				false	"短信结果"
//	@Param		send_time		body		int				false	"发送时间"
//	@Success	200			{object}	response.Response	"成功"
//	@Router		/api/admin/system_log_sms/add [post]
func (hd *SystemLogSmsHandler) Add(c *gin.Context) {
	var addReq SystemLogSmsAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	createId, e := SystemLogSmsService.Add(addReq)
	response.CheckAndRespWithData(c,createId, e)
}
//	@Summary	系统短信日志编辑
//	@Tags		system_log_sms-系统短信日志
//	@Produce	json
//	@Param		Token		header		string				true	"token"
//	@Param		id		body		int				false	"id"
//	@Param		scene		body		int				false	"场景编号"
//	@Param		mobile		body		string				false	"手机号码"
//	@Param		content		body		string				false	"发送内容"
//	@Param		status		body		int				false	"发送状态：[0=发送中, 1=发送成功, 2=发送失败]"
//	@Param		results		body		string				false	"短信结果"
//	@Param		send_time		body		int				false	"发送时间"
//	@Success	200			{object}	response.Response	"成功"
//	@Router		/api/admin/system_log_sms/edit [post]
func (hd *SystemLogSmsHandler) Edit(c *gin.Context) {
	var editReq SystemLogSmsEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	response.CheckAndRespWithData(c,editReq.Id, SystemLogSmsService.Edit(editReq))
}
//	@Summary	系统短信日志删除
//	@Tags		system_log_sms-系统短信日志
//	@Produce	json
//	@Param		Token		header		string				true	"token"
//	@Param		id		body		int				false	"id"
//	@Success	200			{object}	response.Response	"成功"
//	@Router		/api/admin/system_log_sms/del [post]
func (hd *SystemLogSmsHandler) Del(c *gin.Context) {
	var delReq SystemLogSmsDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, SystemLogSmsService.Del(delReq.Id))
}



//	@Summary	系统短信日志导出
//	@Tags		system_log_sms-系统短信日志
//	@Produce	json
//	@Param		Token		header		string				true	"token"
//	@Param scene query int false "场景编号"
//	@Param mobile query string false "手机号码"
//	@Param content query string false "发送内容"
//	@Param status query int false "发送状态：[0=发送中, 1=发送成功, 2=发送失败]"
//	@Param results query string false "短信结果"
//	@Param send_time query int false "发送时间"
//	@Param create_timeStart  query core.TsTime false "创建时间"
//	@Param create_timeEnd  query core.TsTime false "创建时间"
//	@Param update_timeStart  query core.TsTime false "更新时间"
//	@Param update_timeEnd  query core.TsTime false "更新时间"
//	@Router		/api/admin/system_log_sms/ExportFile [get]
func (hd *SystemLogSmsHandler) ExportFile(c *gin.Context) {
	var listReq SystemLogSmsListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := SystemLogSmsService.ExportFile(listReq)
	if err != nil {
		response.FailWithMsg(c, response.SystemError, "查询信息失败")
		return
	}
	f, err := excel.NormalDynamicExport(res, "Sheet1", "系统短信日志", nil)
	if err != nil {
		response.FailWithMsg(c, response.SystemError, "导出失败")
		return
	}
	excel.DownLoadExcel("系统短信日志" + time.Now().Format("20060102-150405"), c.Writer, f)
}

//  @Summary	系统短信日志导入
//  @Tags		system_log_sms-系统短信日志
//  @Produce	json
//	@Router		/api/admin/system_log_sms/ImportFile [post]
func (hd *SystemLogSmsHandler) ImportFile(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusInternalServerError, "文件不存在")
		return
	}
	defer file.Close()
	importList := []SystemLogSmsResp{}
	err = excel.GetExcelData(file, &importList)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	err = SystemLogSmsService.ImportFile(importList)
	response.CheckAndResp(c, err)
}