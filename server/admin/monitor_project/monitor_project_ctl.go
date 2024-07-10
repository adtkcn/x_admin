package monitor_project

import (
	"net/http"
	"strconv"
	"time"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util"
	"x_admin/util/excel"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/singleflight"
)

type MonitorProjectHandler struct {
	requestGroup singleflight.Group
}

// @Summary	错误项目列表
// @Tags		monitor_project-错误项目
// @Produce	json
// @Param		Token		header		string																	true	"token"
// @Param		PageNo		query		int																		true	"页码"
// @Param		PageSize	query		int																		true	"每页数量"
// @Param		projectKey	query		string																	false	"项目uuid."
// @Param		projectName	query		string																	false	"项目名称."
// @Param		projectType	query		string																	false	"项目类型go java web node php 等."
// @Success	200			{object}	response.Response{data=response.PageResp{lists=[]MonitorProjectResp}}	"成功"
// @Failure	400			{object}	string																	"请求错误"
// @Router		/api/admin/monitor_project/list [get]
func (hd *MonitorProjectHandler) List(c *gin.Context) {
	var page request.PageReq
	var listReq MonitorProjectListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := Service.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

// @Summary	错误项目列表-所有
// @Tags		monitor_project-错误项目
// @Produce	json
// @Success	200	{object}	response.Response{data=[]MonitorProjectResp}	"成功"
// @Router		/api/admin/monitor_project/listAll [get]
// var sg singleflight.Group

func (hd *MonitorProjectHandler) ListAll(c *gin.Context) {
	res, err, _ := hd.requestGroup.Do("monitor_project:listAll", func() (any, error) {
		v, err := Service.ListAll()
		return v, err
	})
	// fmt.Printf("v: %v, shared: %v\n", res, shared)

	// var listReq MonitorProjectListReq
	// if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
	// 	return
	// }
	response.CheckAndRespWithData(c, res, err)
}

// @Summary	错误项目详情
// @Tags		monitor_project-错误项目
// @Produce	json
// @Param		Token	header		string										true	"token"
// @Param		id		query		int											false	"项目id."
// @Success	200		{object}	response.Response{data=MonitorProjectResp}	"成功"
// @Router		/api/admin/monitor_project/detail [get]
func (hd *MonitorProjectHandler) Detail(c *gin.Context) {
	var detailReq MonitorProjectDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err, _ := hd.requestGroup.Do("monitor_project:Detail:"+strconv.Itoa(detailReq.Id), func() (any, error) {
		v, err := Service.Detail(detailReq.Id)
		return v, err
	})
	// fmt.Printf("v: %v, shared: %v\n", res, shared)
	response.CheckAndRespWithData(c, res, err)
}

// @Summary	错误项目新增
// @Tags		monitor_project-错误项目
// @Produce	json
// @Param		Token		header		string				true	"token"
// @Param		projectKey	body		string				false	"项目uuid."
// @Param		projectName	body		string				false	"项目名称."
// @Param		projectType	body		string				false	"项目类型go java web node php 等."
// @Success	200			{object}	response.Response	"成功"
// @Router		/api/admin/monitor_project/add [post]
func (hd *MonitorProjectHandler) Add(c *gin.Context) {
	var addReq MonitorProjectAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	createId, e := Service.Add(addReq)
	response.CheckAndRespWithData(c, createId, e)
}

// @Summary	错误项目编辑
// @Tags		monitor_project-错误项目
// @Produce	json
// @Param		Token		header		string				true	"token"
// @Param		id			body		int					false	"项目id."
// @Param		projectKey	body		string				false	"项目uuid."
// @Param		projectName	body		string				false	"项目名称."
// @Param		projectType	body		string				false	"项目类型go java web node php 等."
// @Success	200			{object}	response.Response	"成功"
// @Router		/api/admin/monitor_project/edit [post]
func (hd *MonitorProjectHandler) Edit(c *gin.Context) {
	var editReq MonitorProjectEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	response.CheckAndRespWithData(c, editReq.Id, Service.Edit(editReq))
}

// @Summary	错误项目删除
// @Tags		monitor_project-错误项目
// @Produce	json
// @Param		Token	header		string				true	"token"
// @Param		id		body		int					false	"项目id."
// @Success	200		{object}	response.Response	"成功"
// @Router		/api/admin/monitor_project/del [post]
func (hd *MonitorProjectHandler) Del(c *gin.Context) {
	var delReq MonitorProjectDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, Service.Del(delReq.Id))
}

// @Summary	错误项目导出
// @Tags		monitor_project-错误项目
// @Produce	json
// @Param		Token		header	string	true	"token"
// @Param		projectKey	query	string	false	"项目uuid."
// @Param		projectName	query	string	false	"项目名称."
// @Param		projectType	query	string	false	"项目类型go java web node php 等."
// @Router		/api/admin/monitor_project/ExportFile [get]
func (hd *MonitorProjectHandler) ExportFile(c *gin.Context) {
	var listReq MonitorProjectListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := Service.ExportFile(listReq)
	if err != nil {
		response.FailWithMsg(c, response.SystemError, "查询信息失败")
		return
	}
	f, err := excel.NormalDynamicExport(res, "Sheet1", "错误项目", nil)
	if err != nil {
		response.FailWithMsg(c, response.SystemError, "导出失败")
		return
	}
	excel.DownLoadExcel("错误项目"+time.Now().Format("2006-01-02 15:04:05"), c.Writer, f)
}

// @Summary	错误项目导入
// @Tags		monitor_project-错误项目
// @Produce	json
func (hd *MonitorProjectHandler) ImportFile(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusInternalServerError, "文件不存在")
		return
	}
	defer file.Close()
	importList := []MonitorProjectResp{}
	err = excel.GetExcelData(file, &importList)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	// for _, t := range importList {
	// 	fmt.Printf("%#v", t)
	// }
	err = Service.ImportFile(importList)
	response.CheckAndResp(c, err)
}
