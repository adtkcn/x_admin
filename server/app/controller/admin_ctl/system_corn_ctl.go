package admin_ctl

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"x_admin/app/schema"
	"x_admin/app/service/cornService"

	"x_admin/config"
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util"
	"x_admin/util/excel2"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/singleflight"
)

type SystemCornHandler struct {
	requestGroup singleflight.Group
}

//	 @Summary	定时任务列表
//	 @Tags		system_corn-定时任务
//	 @Produce	json
//	 @Param		token		header		string				true	"token"
//	 @Param		pageNo		query		int					true	"页码"
//	 @Param		pageSize	query		int					true	"每页数量"
//		@Param TaskName query string false "任务名称"
//		@Param TaskCode query string false "任务编码"
//		@Param CornExpr query string false "corn表达式"
//		@Param Status query number false "状态"
//		@Param CreatedBy query string false "创建人"
//		@Param CreateTimeStart  query string false "创建时间"
//		@Param CreateTimeEnd  query string false "创建时间"
//		@Param UpdateTimeStart  query string false "更新时间"
//		@Param UpdateTimeEnd  query string false "更新时间"
//
// @Success 200	{object} response.Response{data=response.PageResp{lists=[]schema.SystemCornResp}}	"成功"
// @Router	/api/admin/system_corn/list [get]
func (hd *SystemCornHandler) List(c *gin.Context) {
	var page request.PageReq
	var listReq schema.SystemCornListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := cornService.SystemCornService.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

//		@Summary	定时任务列表-所有
//		@Tags		system_corn-定时任务
//	 @Produce	json
//	 @Param		token		header		string				true	"token"
//		@Param TaskName query string false "任务名称"
//		@Param TaskCode query string false "任务编码"
//		@Param CornExpr query string false "corn表达式"
//		@Param Status query number false "状态"
//		@Param CreatedBy query string false "创建人"
//		@Param CreateTimeStart  query string false "创建时间"
//		@Param CreateTimeEnd  query string false "创建时间"
//		@Param UpdateTimeStart  query string false "更新时间"
//		@Param UpdateTimeEnd  query string false "更新时间"
//		@Success	200			{object}	response.Response{data=[]SystemCornResp}	"成功"
//		@Router		/api/admin/system_corn/listAll [get]
func (hd *SystemCornHandler) ListAll(c *gin.Context) {
	var listReq schema.SystemCornListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := cornService.SystemCornService.ListAll(listReq)
	response.CheckAndRespWithData(c, res, err)
}

// @Summary	定时任务详情
// @Tags		system_corn-定时任务
// @Produce	json
// @Param		token		header		string				true	"token"
// @Param		Id		query		string				false	"taskid"
// @Success	200			{object}	response.Response{data=SystemCornResp}	"成功"
// @Router		/api/admin/system_corn/detail [get]
func (hd *SystemCornHandler) Detail(c *gin.Context) {
	var detailReq schema.SystemCornPrimarykey
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err, _ := hd.requestGroup.Do(fmt.Sprintf("SystemCorn:Detail:%v", detailReq.Id), func() (any, error) {
		v, err := cornService.SystemCornService.Detail(detailReq.Id)
		return v, err
	})

	response.CheckAndRespWithData(c, res, err)
}

// @Summary	定时任务新增
// @Tags		system_corn-定时任务
// @Produce	json
// @Param		token		header		string				true	"token"
// @Param		TaskName		body		string				false	"任务名称"
// @Param		TaskCode		body		string				false	"任务编码"
// @Param		CornExpr		body		string				false	"corn表达式"
// @Param		Status		body		number				false	"状态"
// @Success	200			{object}	response.Response	"成功"
// @Router		/api/admin/system_corn/add [post]
func (hd *SystemCornHandler) Add(c *gin.Context) {
	var addReq schema.SystemCornAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	var adminId = config.AdminConfig.GetAdminId(c)
	createId, e := cornService.SystemCornService.Add(addReq, adminId)
	response.CheckAndRespWithData(c, createId, e)
}

// @Summary	定时任务编辑
// @Tags		system_corn-定时任务
// @Produce	json
// @Param		token		header		string				true	"token"
// @Param		Id		body		string				false	"taskid"
// @Param		TaskName		body		string				false	"任务名称"
// @Param		TaskCode		body		string				false	"任务编码"
// @Param		CornExpr		body		string				false	"corn表达式"
// @Param		Status		body		number				false	"状态"
// @Success	200			{object}	response.Response	"成功"
// @Router		/api/admin/system_corn/edit [post]
func (hd *SystemCornHandler) Edit(c *gin.Context) {
	var editReq schema.SystemCornEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	response.CheckAndRespWithData(c, editReq.Id, cornService.SystemCornService.Edit(editReq))
}

// @Summary	定时任务删除
// @Tags		system_corn-定时任务
// @Produce	json
// @Param		token		header		string				true	"token"
// @Param		Id		body		string				false	"taskid"
// @Success	200			{object}	response.Response	"成功"
// @Router		/api/admin/system_corn/del [post]
func (hd *SystemCornHandler) Del(c *gin.Context) {
	var delReq schema.SystemCornPrimarykey
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, cornService.SystemCornService.Del(delReq.Id))
}

//	@Summary	定时任务删除-批量
//	@Tags		system_corn-定时任务
//
// @Produce	json
// @Param		token		header		string				true	"token"
// @Param		Ids		body		string				false	"逗号分割的id"
// @Success	200			{object}	response.Response	"成功"
// @Router		/api/admin/system_corn/delBatch [post]
func (hd *SystemCornHandler) DelBatch(c *gin.Context) {
	var delReq schema.SystemCornDelBatchReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	if delReq.Ids == "" {
		response.FailWithMsg(c, response.SystemError, "请选择要删除的数据")
		return
	}
	var Ids = strings.Split(delReq.Ids, ",")

	response.CheckAndResp(c, cornService.SystemCornService.DelBatch(Ids))
}

//		@Summary	定时任务导出
//		@Tags		system_corn-定时任务
//		@Produce	octet-stream,json
//		@Param		token		header		string				true	"token"
//		@Param TaskName query string false "任务名称"
//		@Param TaskCode query string false "任务编码"
//		@Param CornExpr query string false "corn表达式"
//		@Param Status query number false "状态"
//		@Param CreatedBy query string false "创建人"
//		@Param CreateTimeStart  query string false "创建时间"
//		@Param CreateTimeEnd  query string false "创建时间"
//		@Param UpdateTimeStart  query string false "更新时间"
//		@Param UpdateTimeEnd  query string false "更新时间"
//	 @Success	200		{file} string	"成功"
//	 @Failure	500 	{object}	response.Response	"失败"
//		@Router		/api/admin/system_corn/ExportFile [get]
func (hd *SystemCornHandler) ExportFile(c *gin.Context) {
	var listReq schema.SystemCornListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := cornService.SystemCornService.ExportFile(listReq)
	if err != nil {
		response.FailWithMsg(c, response.SystemError, "查询信息失败")
		return
	}
	f, err := excel2.Export(res, cornService.SystemCornService.GetExcelCol(), "Sheet1", "定时任务")
	if err != nil {
		response.FailWithMsg(c, response.SystemError, "导出失败")
		return
	}
	excel2.DownLoadExcel("定时任务"+time.Now().Format("20060102-150405"), c.Writer, f)
}

//	 @Summary	定时任务导入
//	 @Tags		system_corn-定时任务
//	 @Produce	json
//	 @Param		token		header		string				true	"token"
//	 @Param		file	formData	file	true	"导入文件"
//	 @Success	200		{object}	response.Response	"成功"
//		@Router		/api/admin/system_corn/ImportFile [post]
func (hd *SystemCornHandler) ImportFile(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusInternalServerError, "文件不存在")
		return
	}
	defer file.Close()
	importList := []schema.SystemCornResp{}
	err = excel2.GetExcelData(file, &importList, cornService.SystemCornService.GetExcelCol())
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	err = cornService.SystemCornService.ImportFile(importList)
	response.CheckAndResp(c, err)
}

// @Summary	获取任务列表
// @Tags		system_corn-定时任务
// @Produce	json
// @Param		token		header		string				true	"token"
// @Success	200			{object}	response.Response	"成功"
// @Router		/api/admin/system_corn/getTaskList [get]
func (hd *SystemCornHandler) GetTaskList(c *gin.Context) {
	var taskList = cornService.SystemCornService.GetTaskList()
	response.CheckAndRespWithData(c, taskList, nil)
}
