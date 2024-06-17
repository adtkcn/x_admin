package monitor_client

import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin" 
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util"
	"x_admin/util/excel"
)

 
type MonitorClientHandler struct {}

//  @Summary	监控-客户端信息列表
//  @Tags		monitor_client-监控-客户端信息
//  @Produce	json
//  @Param		Token		header		string				true	"token"
//  @Param		PageNo		query		int					true	"页码"
//  @Param		PageSize	query		int					true	"每页数量"
//	@Param projectKey query string false "项目key."
//	@Param clientId query string false "sdk生成的客户端id."
//	@Param userId query string false "用户id."
//	@Param os query string false "系统."
//	@Param browser query string false "浏览器."
//	@Param city query string false "城市."
//	@Param width query int false "屏幕."
//	@Param height query int false "屏幕高度."
//	@Param ua query string false "ua记录."
//	@Param createTimeStart  query core.TsTime false "创建时间."
//	@Param createTimeEnd  query core.TsTime false "创建时间."
//	@Param clientTimeStart  query core.TsTime false "更新时间."
//	@Param clientTimeEnd  query core.TsTime false "更新时间."
//	@Success	200			{object}	[]MonitorClientResp	"成功"
//	@Failure	400			{object}	string				"请求错误"
//	@Router		/api/admin/monitor_client/list [get]
func (hd MonitorClientHandler) List(c *gin.Context) {
	var page request.PageReq
	var listReq MonitorClientListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := Service.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

//	@Summary	监控-客户端信息列表-所有
//	@Tags		monitor_client-监控-客户端信息
//  @Produce	json
//	@Param projectKey query string false "项目key."
//	@Param clientId query string false "sdk生成的客户端id."
//	@Param userId query string false "用户id."
//	@Param os query string false "系统."
//	@Param browser query string false "浏览器."
//	@Param city query string false "城市."
//	@Param width query int false "屏幕."
//	@Param height query int false "屏幕高度."
//	@Param ua query string false "ua记录."
//	@Param createTimeStart  query core.TsTime false "创建时间."
//	@Param createTimeEnd  query core.TsTime false "创建时间."
//	@Param clientTimeStart  query core.TsTime false "更新时间."
//	@Param clientTimeEnd  query core.TsTime false "更新时间."
//	@Success	200			{object}	[]MonitorClientResp	"成功"
//	@Router		/api/admin/monitor_client/listAll [get]
func (hd MonitorClientHandler) ListAll(c *gin.Context) {
	var listReq MonitorClientListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := Service.ListAll(listReq)
	response.CheckAndRespWithData(c, res, err)
}

//	@Summary	监控-客户端信息详情
//	@Tags		monitor_client-监控-客户端信息
//	@Produce	json
//	@Param		Token		header		string				true	"token"
//	@Param		id		query		int				false	"uuid."
//	@Success	200			{object}	MonitorClientResp	"成功"
//	@Router		/api/admin/monitor_client/detail [get]
func (hd MonitorClientHandler) Detail(c *gin.Context) {
	var detailReq MonitorClientDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := Service.Detail(detailReq.Id)
	response.CheckAndRespWithData(c, res, err)
}


//	@Summary	监控-客户端信息新增
//	@Tags		monitor_client-监控-客户端信息
//	@Produce	json
//	@Param		Token		header		string				true	"token"
//	@Param		clientId		body		string				false	"sdk生成的客户端id."
//	@Param		userId		body		string				false	"用户id."
//	@Param		os		body		string				false	"系统."
//	@Param		browser		body		string				false	"浏览器."
//	@Param		city		body		string				false	"城市."
//	@Param		width		body		int				false	"屏幕."
//	@Param		height		body		int				false	"屏幕高度."
//	@Param		ua		body		string				false	"ua记录."
//	@Param		clientTime		body		core.TsTime				false	"更新时间."
//	@Success	200			{object}	response.RespType	"成功"
//	@Router		/api/admin/monitor_client/add [post]
func (hd MonitorClientHandler) Add(c *gin.Context) {
	var addReq MonitorClientAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	response.CheckAndResp(c, Service.Add(addReq))
}
//	@Summary	监控-客户端信息编辑
//	@Tags		monitor_client-监控-客户端信息
//	@Produce	json
//	@Param		Token		header		string				true	"token"
//	@Param		clientId		body		string				false	"sdk生成的客户端id."
//	@Param		userId		body		string				false	"用户id."
//	@Param		os		body		string				false	"系统."
//	@Param		browser		body		string				false	"浏览器."
//	@Param		city		body		string				false	"城市."
//	@Param		width		body		int				false	"屏幕."
//	@Param		height		body		int				false	"屏幕高度."
//	@Param		ua		body		string				false	"ua记录."
//	@Param		clientTime		body		core.TsTime				false	"更新时间."
//	@Success	200			{object}	response.RespType	"成功"
//	@Router		/api/admin/monitor_client/edit [post]
func (hd MonitorClientHandler) Edit(c *gin.Context) {
	var editReq MonitorClientEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	response.CheckAndResp(c, Service.Edit(editReq))
}
//	@Summary	监控-客户端信息删除
//	@Tags		monitor_client-监控-客户端信息
//	@Produce	json
//	@Param		Token		header		string				true	"token"
//	@Param		id		body		int				false	"uuid."
//	@Success	200			{object}	response.RespType	"成功"
//	@Router		/api/admin/monitor_client/del [post]
func (hd MonitorClientHandler) Del(c *gin.Context) {
	var delReq MonitorClientDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, Service.Del(delReq.Id))
}



//	@Summary	监控-客户端信息导出
//	@Tags		monitor_client-监控-客户端信息
//	@Produce	json
//	@Param		Token		header		string				true	"token"
//	@Param		projectKey		query		string				false	"项目key."
//	@Param		clientId		query		string				false	"sdk生成的客户端id."
//	@Param		userId		query		string				false	"用户id."
//	@Param		os		query		string				false	"系统."
//	@Param		browser		query		string				false	"浏览器."
//	@Param		city		query		string				false	"城市."
//	@Param		width		query		int				false	"屏幕."
//	@Param		height		query		int				false	"屏幕高度."
//	@Param		ua		query		string				false	"ua记录."
//	@Param		createTime		query		core.TsTime				false	"创建时间."
//	@Param		clientTime		query		core.TsTime				false	"更新时间."
//	@Router		/api/admin/monitor_client/ExportFile [get]
func (hd MonitorClientHandler) ExportFile(c *gin.Context) {
	var listReq MonitorClientListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := Service.ExportFile(listReq)
	if err != nil {
		response.FailWithMsg(c, response.SystemError, "查询信息失败")
		return
	}
	f, err := excel.NormalDynamicExport(res, "Sheet1", "监控-客户端信息", nil)
	if err != nil {
		response.FailWithMsg(c, response.SystemError, "导出失败")
		return
	}
	excel.DownLoadExcel("监控-客户端信息" + time.Now().Format("20060102-150405"), c.Writer, f)
}

//  @Summary	监控-客户端信息导入
//  @Tags		monitor_client-监控-客户端信息
//  @Produce	json
func (hd MonitorClientHandler) ImportFile(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusInternalServerError, "文件不存在")
		return
	}
	defer file.Close()
	importList := []MonitorClientResp{}
	err = excel.GetExcelData(file, &importList)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
//	for _, t := range importList {
//		fmt.Printf("%#v", t)
//	}
	err = Service.ImportFile(importList)
	response.CheckAndResp(c, err)
}