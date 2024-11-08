package monitor_slow

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util"
	"x_admin/util/excel2"
	"x_admin/util/img_util"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/singleflight"
)

type MonitorSlowHandler struct {
	requestGroup singleflight.Group
}

//	@Summary	监控-错误列列表
//	@Tags		monitor_slow-监控-错误列
//	@Produce	json
//	@Param		Token			header		string																	true	"token"
//	@Param		PageNo			query		int																		true	"页码"
//	@Param		PageSize		query		int																		true	"每页数量"
//	@Param		ProjectKey		query		string																	false	"项目key"
//	@Param		ClientId		query		string																	false	"sdk生成的客户端id"
//	@Param		UserId			query		string																	false	"用户id"
//	@Param		Path			query		string																	false	"URL地址"
//	@Param		Time			query		number																	false	"时间"
//	@Param		CreateTimeStart	query		string																	false	"创建时间"
//	@Param		CreateTimeEnd	query		string																	false	"创建时间"
//
//	@Success	200				{object}	response.Response{ data=response.PageResp{ lists=[]MonitorSlowResp}}	"成功"
//	@Router		/api/admin/monitor_slow/list [get]
func (hd *MonitorSlowHandler) List(c *gin.Context) {
	var page request.PageReq
	var listReq MonitorSlowListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := MonitorSlowService.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

//	@Summary	监控-错误列列表-所有
//	@Tags		monitor_slow-监控-错误列
//	@Produce	json
//	@Param		ProjectKey		query		string										false	"项目key"
//	@Param		ClientId		query		string										false	"sdk生成的客户端id"
//	@Param		UserId			query		string										false	"用户id"
//	@Param		Path			query		string										false	"URL地址"
//	@Param		Time			query		number										false	"时间"
//	@Param		CreateTimeStart	query		string										false	"创建时间"
//	@Param		CreateTimeEnd	query		string										false	"创建时间"
//	@Success	200				{object}	response.Response{ data=[]MonitorSlowResp}	"成功"
//	@Router		/api/admin/monitor_slow/listAll [get]
func (hd *MonitorSlowHandler) ListAll(c *gin.Context) {
	var listReq MonitorSlowListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := MonitorSlowService.ListAll(listReq)
	response.CheckAndRespWithData(c, res, err)
}

//	@Summary	监控-错误列详情
//	@Tags		monitor_slow-监控-错误列
//	@Produce	json
//	@Param		Token	header		string										true	"token"
//	@Param		Id		query		number										false	"错误id"
//	@Success	200		{object}	response.Response{ data=MonitorSlowResp}	"成功"
//	@Router		/api/admin/monitor_slow/detail [get]
func (hd *MonitorSlowHandler) Detail(c *gin.Context) {
	var detailReq MonitorSlowDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err, _ := hd.requestGroup.Do("MonitorSlow:Detail:"+strconv.Itoa(detailReq.Id), func() (any, error) {
		v, err := MonitorSlowService.Detail(detailReq.Id)
		return v, err
	})

	response.CheckAndRespWithData(c, res, err)
}

//	@Summary	监控-错误列新增
//	@Tags		monitor_slow-监控-错误列
//	@Produce	json
//	@Success	200	{object}	response.Response	"成功"
//	@Router		/api/admin/monitor_slow/add [get]
func (hd *MonitorSlowHandler) Add(c *gin.Context) {
	data, err := url.QueryUnescape(c.Query("data"))
	if err != nil {
		c.Data(200, "image/gif", img_util.EmptyGif())
		return
	}

	var addReq MonitorSlowAddReq
	json.Unmarshal([]byte(data), &addReq)

	MonitorSlowService.Add(addReq)
	c.Data(200, "image/gif", img_util.EmptyGif())
}

//	@Summary	监控-错误列编辑
//	@Tags		monitor_slow-监控-错误列
//	@Produce	json
//	@Param		Token		header		string				true	"token"
//	@Param		Id			body		number				false	"错误id"
//	@Param		ProjectKey	body		string				false	"项目key"
//	@Param		ClientId	body		string				false	"sdk生成的客户端id"
//	@Param		UserId		body		string				false	"用户id"
//	@Param		Path		body		string				false	"URL地址"
//	@Param		Time		body		number				false	"时间"
//	@Success	200			{object}	response.Response	"成功"
//	@Router		/api/admin/monitor_slow/edit [post]
func (hd *MonitorSlowHandler) Edit(c *gin.Context) {
	var editReq MonitorSlowEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	response.CheckAndRespWithData(c, editReq.Id, MonitorSlowService.Edit(editReq))
}

//	@Summary	监控-错误列删除
//	@Tags		monitor_slow-监控-错误列
//	@Produce	json
//	@Param		Token	header		string				true	"token"
//	@Param		Id		body		number				false	"错误id"
//	@Success	200		{object}	response.Response	"成功"
//	@Router		/api/admin/monitor_slow/del [post]
func (hd *MonitorSlowHandler) Del(c *gin.Context) {
	var delReq MonitorSlowDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, MonitorSlowService.Del(delReq.Id))
}

//	@Summary	监控-错误列删除-批量
//	@Tags		monitor_slow-监控-错误列
//
//	@Produce	json
//	@Param		Token	header		string				true	"token"
//	@Param		Ids		body		string				false	"逗号分割的id"
//	@Success	200		{object}	response.Response	"成功"
//	@Router		/api/admin/monitor_slow/delBatch [post]
func (hd *MonitorSlowHandler) DelBatch(c *gin.Context) {
	var delReq MonitorSlowDelBatchReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	if delReq.Ids == "" {
		response.FailWithMsg(c, response.SystemError, "请选择要删除的数据")
		return
	}
	var Ids = strings.Split(delReq.Ids, ",")

	response.CheckAndResp(c, MonitorSlowService.DelBatch(Ids))
}

//	@Summary	监控-错误列导出
//	@Tags		monitor_slow-监控-错误列
//	@Produce	json
//	@Param		Token			header	string	true	"token"
//	@Param		ProjectKey		query	string	false	"项目key"
//	@Param		ClientId		query	string	false	"sdk生成的客户端id"
//	@Param		UserId			query	string	false	"用户id"
//	@Param		Path			query	string	false	"URL地址"
//	@Param		Time			query	number	false	"时间"
//	@Param		CreateTimeStart	query	string	false	"创建时间"
//	@Param		CreateTimeEnd	query	string	false	"创建时间"
//	@Router		/api/admin/monitor_slow/ExportFile [get]
func (hd *MonitorSlowHandler) ExportFile(c *gin.Context) {
	var listReq MonitorSlowListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := MonitorSlowService.ExportFile(listReq)
	if err != nil {
		response.FailWithMsg(c, response.SystemError, "查询信息失败")
		return
	}
	f, err := excel2.Export(res, MonitorSlowService.GetExcelCol(), "Sheet1", "监控-错误列")
	if err != nil {
		response.FailWithMsg(c, response.SystemError, "导出失败")
		return
	}
	excel2.DownLoadExcel("监控-错误列"+time.Now().Format("20060102-150405"), c.Writer, f)
}

//	@Summary	监控-错误列导入
//	@Tags		monitor_slow-监控-错误列
//	@Produce	json
//	@Router		/api/admin/monitor_slow/ImportFile [post]
func (hd *MonitorSlowHandler) ImportFile(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusInternalServerError, "文件不存在")
		return
	}
	defer file.Close()
	importList := []MonitorSlowResp{}
	err = excel2.GetExcelData(file, &importList, MonitorSlowService.GetExcelCol())
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	err = MonitorSlowService.ImportFile(importList)
	response.CheckAndResp(c, err)
}
