package monitor_project

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

type MonitorProjectHandler struct {
	requestGroup singleflight.Group
}

//	 @Summary	监控项目列表
//	 @Tags		monitor_project-监控项目
//	 @Produce	json
//	 @Param		Token		header		string				true	"token"
//	 @Param		PageNo		query		int					true	"页码"
//	 @Param		PageSize	query		int					true	"每页数量"
//		@Param ProjectKey query string false "项目uuid"
//		@Param ProjectName query string false "项目名称"
//		@Param ProjectType query string false "项目类型go java web node php 等"
//		@Param Status query number false "是否启用: 0=否, 1=是"
//		@Param CreateTimeStart  query string false "创建时间"
//		@Param CreateTimeEnd  query string false "创建时间"
//		@Param UpdateTimeStart  query string false "更新时间"
//		@Param UpdateTimeEnd  query string false "更新时间"
//
// @Success 200	{object} response.Response{ data=response.PageResp{ lists=[]MonitorProjectResp}}	"成功"
// @Router	/api/admin/monitor_project/list [get]
func (hd *MonitorProjectHandler) List(c *gin.Context) {
	var page request.PageReq
	var listReq MonitorProjectListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := MonitorProjectService.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

//		@Summary	监控项目列表-所有
//		@Tags		monitor_project-监控项目
//	 @Produce	json
//		@Param ProjectKey query string false "项目uuid"
//		@Param ProjectName query string false "项目名称"
//		@Param ProjectType query string false "项目类型go java web node php 等"
//		@Param Status query number false "是否启用: 0=否, 1=是"
//		@Param CreateTimeStart  query string false "创建时间"
//		@Param CreateTimeEnd  query string false "创建时间"
//		@Param UpdateTimeStart  query string false "更新时间"
//		@Param UpdateTimeEnd  query string false "更新时间"
//		@Success	200			{object}	response.Response{ data=[]MonitorProjectResp}	"成功"
//		@Router		/api/admin/monitor_project/listAll [get]
func (hd *MonitorProjectHandler) ListAll(c *gin.Context) {
	var listReq MonitorProjectListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := MonitorProjectService.ListAll(listReq)
	response.CheckAndRespWithData(c, res, err)
}

// @Summary	监控项目详情
// @Tags		monitor_project-监控项目
// @Produce	json
// @Param		Token		header		string				true	"token"
// @Param		Id		query		number				false	"项目id"
// @Success	200			{object}	response.Response{ data=MonitorProjectResp}	"成功"
// @Router		/api/admin/monitor_project/detail [get]
func (hd *MonitorProjectHandler) Detail(c *gin.Context) {
	var detailReq MonitorProjectDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err, _ := hd.requestGroup.Do("MonitorProject:Detail:"+strconv.Itoa(detailReq.Id), func() (any, error) {
		v, err := MonitorProjectService.Detail(detailReq.Id)
		return v, err
	})

	response.CheckAndRespWithData(c, res, err)
}

// @Summary	监控项目新增
// @Tags		monitor_project-监控项目
// @Produce	json
// @Param		Token		header		string				true	"token"
// @Param		ProjectKey		body		string				false	"项目uuid"
// @Param		ProjectName		body		string				false	"项目名称"
// @Param		ProjectType		body		string				false	"项目类型go java web node php 等"
// @Param		Status		body		number				false	"是否启用: 0=否, 1=是"
// @Success	200			{object}	response.Response	"成功"
// @Router		/api/admin/monitor_project/add [post]
func (hd *MonitorProjectHandler) Add(c *gin.Context) {
	var addReq MonitorProjectAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	createId, e := MonitorProjectService.Add(addReq)
	response.CheckAndRespWithData(c, createId, e)
}

// @Summary	监控项目编辑
// @Tags		monitor_project-监控项目
// @Produce	json
// @Param		Token		header		string				true	"token"
// @Param		Id		body		number				false	"项目id"
// @Param		ProjectKey		body		string				false	"项目uuid"
// @Param		ProjectName		body		string				false	"项目名称"
// @Param		ProjectType		body		string				false	"项目类型go java web node php 等"
// @Param		Status		body		number				false	"是否启用: 0=否, 1=是"
// @Success	200			{object}	response.Response	"成功"
// @Router		/api/admin/monitor_project/edit [post]
func (hd *MonitorProjectHandler) Edit(c *gin.Context) {
	var editReq MonitorProjectEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	response.CheckAndRespWithData(c, editReq.Id, MonitorProjectService.Edit(editReq))
}

// @Summary	监控项目删除
// @Tags		monitor_project-监控项目
// @Produce	json
// @Param		Token		header		string				true	"token"
// @Param		Id		body		number				false	"项目id"
// @Success	200			{object}	response.Response	"成功"
// @Router		/api/admin/monitor_project/del [post]
func (hd *MonitorProjectHandler) Del(c *gin.Context) {
	var delReq MonitorProjectDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, MonitorProjectService.Del(delReq.Id))
}

//	@Summary	监控项目删除-批量
//	@Tags		monitor_project-监控项目
//
// @Produce	json
// @Param		Token		header		string				true	"token"
// @Param		Ids		body		string				false	"逗号分割的id"
// @Success	200			{object}	response.Response	"成功"
// @Router		/api/admin/monitor_project/delBatch [post]
func (hd *MonitorProjectHandler) DelBatch(c *gin.Context) {
	var delReq MonitorProjectDelBatchReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	if delReq.Ids == "" {
		response.FailWithMsg(c, response.SystemError, "请选择要删除的数据")
		return
	}
	var Ids = strings.Split(delReq.Ids, ",")

	response.CheckAndResp(c, MonitorProjectService.DelBatch(Ids))
}

// @Summary	监控项目导出
// @Tags		monitor_project-监控项目
// @Produce	json
// @Param		Token		header		string				true	"token"
// @Param ProjectKey query string false "项目uuid"
// @Param ProjectName query string false "项目名称"
// @Param ProjectType query string false "项目类型go java web node php 等"
// @Param Status query number false "是否启用: 0=否, 1=是"
// @Param CreateTimeStart  query string false "创建时间"
// @Param CreateTimeEnd  query string false "创建时间"
// @Param UpdateTimeStart  query string false "更新时间"
// @Param UpdateTimeEnd  query string false "更新时间"
// @Router		/api/admin/monitor_project/ExportFile [get]
func (hd *MonitorProjectHandler) ExportFile(c *gin.Context) {
	var listReq MonitorProjectListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := MonitorProjectService.ExportFile(listReq)
	if err != nil {
		response.FailWithMsg(c, response.SystemError, "查询信息失败")
		return
	}
	f, err := excel2.Export(res, MonitorProjectService.GetExcelCol(), "Sheet1", "监控项目")
	if err != nil {
		response.FailWithMsg(c, response.SystemError, "导出失败")
		return
	}
	excel2.DownLoadExcel("监控项目"+time.Now().Format("20060102-150405"), c.Writer, f)
}

//	 @Summary	监控项目导入
//	 @Tags		monitor_project-监控项目
//	 @Produce	json
//		@Router		/api/admin/monitor_project/ImportFile [post]
func (hd *MonitorProjectHandler) ImportFile(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusInternalServerError, "文件不存在")
		return
	}
	defer file.Close()
	importList := []MonitorProjectResp{}
	err = excel2.GetExcelData(file, &importList, MonitorProjectService.GetExcelCol())
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	err = MonitorProjectService.ImportFile(importList)
	response.CheckAndResp(c, err)
}
