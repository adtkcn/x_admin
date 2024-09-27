package monitor_client

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

type MonitorClientHandler struct {
	requestGroup singleflight.Group
}

//	 @Summary	监控-客户端信息列表
//	 @Tags		monitor_client-监控-客户端信息
//	 @Produce	json
//	 @Param		Token		header		string				true	"token"
//	 @Param		PageNo		query		int					true	"页码"
//	 @Param		PageSize	query		int					true	"每页数量"
//		@Param ProjectKey query string false "项目key"
//		@Param ClientId query string false "sdk生成的客户端id"
//		@Param UserId query string false "用户id"
//		@Param Os query string false "系统"
//		@Param Browser query string false "浏览器"
//		@Param City query string false "城市"
//		@Param Width query number false "屏幕"
//		@Param Height query number false "屏幕高度"
//		@Param Ua query string false "ua记录"
//		@Param CreateTimeStart  query string false "创建时间"
//		@Param CreateTimeEnd  query string false "创建时间"
//
// @Success 200	{object} response.Response{ data=response.PageResp{ lists=[]MonitorClientResp}}	"成功"
// @Router	/api/admin/monitor_client/list [get]
func (hd *MonitorClientHandler) List(c *gin.Context) {
	var page request.PageReq
	var listReq MonitorClientListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := MonitorClientService.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

//		@Summary	监控-客户端信息列表-所有
//		@Tags		monitor_client-监控-客户端信息
//	 @Produce	json
//		@Param ProjectKey query string false "项目key"
//		@Param ClientId query string false "sdk生成的客户端id"
//		@Param UserId query string false "用户id"
//		@Param Os query string false "系统"
//		@Param Browser query string false "浏览器"
//		@Param City query string false "城市"
//		@Param Width query number false "屏幕"
//		@Param Height query number false "屏幕高度"
//		@Param Ua query string false "ua记录"
//		@Param CreateTimeStart  query string false "创建时间"
//		@Param CreateTimeEnd  query string false "创建时间"
//		@Success	200			{object}	response.Response{ data=[]MonitorClientResp}	"成功"
//		@Router		/api/admin/monitor_client/listAll [get]
func (hd *MonitorClientHandler) ListAll(c *gin.Context) {
	var listReq MonitorClientListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := MonitorClientService.ListAll(listReq)
	response.CheckAndRespWithData(c, res, err)
}

// @Summary	监控-客户端信息详情
// @Tags		monitor_client-监控-客户端信息
// @Produce	json
// @Param		Token		header		string				true	"token"
// @Param		Id		query		number				false	"uuid"
// @Success	200			{object}	response.Response{ data=MonitorClientResp}	"成功"
// @Router		/api/admin/monitor_client/detail [get]
func (hd *MonitorClientHandler) Detail(c *gin.Context) {
	var detailReq MonitorClientDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err, _ := hd.requestGroup.Do("MonitorClient:Detail:"+strconv.Itoa(detailReq.Id), func() (any, error) {
		v, err := MonitorClientService.Detail(detailReq.Id)
		return v, err
	})

	response.CheckAndRespWithData(c, res, err)
}

// @Summary	监控-客户端信息新增
// @Tags		monitor_client-监控-客户端信息
// @Produce	json
// @Param		Token		header		string				true	"token"
// @Param		ProjectKey		body		string				false	"项目key"
// @Param		ClientId		body		string				false	"sdk生成的客户端id"
// @Param		UserId		body		string				false	"用户id"
// @Param		Os		body		string				false	"系统"
// @Param		Browser		body		string				false	"浏览器"
// @Param		City		body		string				false	"城市"
// @Param		Width		body		number				false	"屏幕"
// @Param		Height		body		number				false	"屏幕高度"
// @Param		Ua		body		string				false	"ua记录"
// @Success	200			{object}	response.Response	"成功"
// @Router		/api/admin/monitor_client/add [post]
func (hd *MonitorClientHandler) Add(c *gin.Context) {
	var addReq MonitorClientAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	createId, e := MonitorClientService.Add(addReq)
	response.CheckAndRespWithData(c, createId, e)
}

// @Summary	监控-客户端信息编辑
// @Tags		monitor_client-监控-客户端信息
// @Produce	json
// @Param		Token		header		string				true	"token"
// @Param		Id		body		number				false	"uuid"
// @Param		ProjectKey		body		string				false	"项目key"
// @Param		ClientId		body		string				false	"sdk生成的客户端id"
// @Param		UserId		body		string				false	"用户id"
// @Param		Os		body		string				false	"系统"
// @Param		Browser		body		string				false	"浏览器"
// @Param		City		body		string				false	"城市"
// @Param		Width		body		number				false	"屏幕"
// @Param		Height		body		number				false	"屏幕高度"
// @Param		Ua		body		string				false	"ua记录"
// @Success	200			{object}	response.Response	"成功"
// @Router		/api/admin/monitor_client/edit [post]
func (hd *MonitorClientHandler) Edit(c *gin.Context) {
	var editReq MonitorClientEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	response.CheckAndRespWithData(c, editReq.Id, MonitorClientService.Edit(editReq))
}

// @Summary	监控-客户端信息删除
// @Tags		monitor_client-监控-客户端信息
// @Produce	json
// @Param		Token		header		string				true	"token"
// @Param		Id		body		number				false	"uuid"
// @Success	200			{object}	response.Response	"成功"
// @Router		/api/admin/monitor_client/del [post]
func (hd *MonitorClientHandler) Del(c *gin.Context) {
	var delReq MonitorClientDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, MonitorClientService.Del(delReq.Id))
}

//	@Summary	监控-客户端信息删除-批量
//	@Tags		monitor_client-监控-客户端信息
//
// @Produce	json
// @Param		Token		header		string				true	"token"
// @Param		Ids		body		string				false	"逗号分割的id"
// @Success	200			{object}	response.Response	"成功"
// @Router		/api/admin/monitor_client/delBatch [post]
func (hd *MonitorClientHandler) DelBatch(c *gin.Context) {
	var delReq MonitorClientDelBatchReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	if delReq.Ids == "" {
		response.FailWithMsg(c, response.SystemError, "请选择要删除的数据")
		return
	}
	var Ids = strings.Split(delReq.Ids, ",")

	response.CheckAndResp(c, MonitorClientService.DelBatch(Ids))
}

// @Summary	监控-客户端信息导出
// @Tags		monitor_client-监控-客户端信息
// @Produce	json
// @Param		Token		header		string				true	"token"
// @Param ProjectKey query string false "项目key"
// @Param ClientId query string false "sdk生成的客户端id"
// @Param UserId query string false "用户id"
// @Param Os query string false "系统"
// @Param Browser query string false "浏览器"
// @Param City query string false "城市"
// @Param Width query number false "屏幕"
// @Param Height query number false "屏幕高度"
// @Param Ua query string false "ua记录"
// @Param CreateTimeStart  query string false "创建时间"
// @Param CreateTimeEnd  query string false "创建时间"
// @Router		/api/admin/monitor_client/ExportFile [get]
func (hd *MonitorClientHandler) ExportFile(c *gin.Context) {
	var listReq MonitorClientListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := MonitorClientService.ExportFile(listReq)
	if err != nil {
		response.FailWithMsg(c, response.SystemError, "查询信息失败")
		return
	}
	f, err := excel2.Export(res, MonitorClientService.GetExcelCol(), "Sheet1", "监控-客户端信息")
	if err != nil {
		response.FailWithMsg(c, response.SystemError, "导出失败")
		return
	}
	excel2.DownLoadExcel("监控-客户端信息"+time.Now().Format("20060102-150405"), c.Writer, f)
}

//	 @Summary	监控-客户端信息导入
//	 @Tags		monitor_client-监控-客户端信息
//	 @Produce	json
//		@Router		/api/admin/monitor_client/ImportFile [post]
func (hd *MonitorClientHandler) ImportFile(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusInternalServerError, "文件不存在")
		return
	}
	defer file.Close()
	importList := []MonitorClientResp{}
	err = excel2.GetExcelData(file, &importList, MonitorClientService.GetExcelCol())
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	err = MonitorClientService.ImportFile(importList)
	response.CheckAndResp(c, err)
}
