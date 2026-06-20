package generator_controller

import (
	"x_admin/app/schema/generator_schema"
	"x_admin/app/service/generator_service"
	"x_admin/core/request"
	"x_admin/core/response"

	"net/http"
	"strings"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

// GenHandler 代码生成器控制器
type GenHandler struct{}

// @Summary		数据表列表
// @Description	获取数据库表列表
// @Tags			generator-代码生成器
// @Param			token			header		string																			true	"token"
// @Param			pageNo			query		int																				true	"页码"
// @Param			pageSize		query		int																				true	"每页数量"
// @Param			tableName		query		string																			false	"表名"
// @Param			tableComment	query		string																			false	"表描述"
// @Success		200				{object}	response.Response{data=response.PageResp{lists=generator_schema.DbTableResp}}	"成功"
// @Router			/api/admin/generator/dbTables [get]
func (gh GenHandler) DbTables(c *gin.Context) {
	var page request.PageReq
	var tbReq generator_schema.DbTablesReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &tbReq)) {
		return
	}
	res, err := generator_service.GenerateService.DbTables(page, tbReq)
	response.CheckAndRespWithData(c, res, err)
}

// @Summary		生成列表
// @Description	获取生成列表
// @Tags			generator-代码生成器
// @Param			token			header		string																			true	"token"
// @Param			pageNo			query		int																				true	"页码"
// @Param			pageSize		query		int																				true	"每页数量"
// @Param			tableName		query		string																			false	"表名"
// @Param			tableComment	query		string																			false	"表描述"
// @Param			startTime		query		time.Time																		false	"开始时间"
// @Param			endTime			query		time.Time																		false	"结束时间"
// @Success		200				{object}	response.Response{data=response.PageResp{lists=generator_schema.GenTableResp}}	"成功"
// @Router			/api/admin/generator/list [get]
func (gh GenHandler) List(c *gin.Context) {
	var page request.PageReq
	var listReq generator_schema.ListTableReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := generator_service.GenerateService.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

// @Summary		生成详情
// @Description	获取生成详情
// @Tags			generator-代码生成器
// @Param			token	header		string													true	"token"
// @Param			id		query		string													true	"主键"
// @Success		200		{object}	response.Response{data=generator_schema.GenTableResp}	"成功"
// @Router			/api/admin/generator/detail [get]
func (gh GenHandler) Detail(c *gin.Context) {
	var detailReq generator_schema.DetailTableReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := generator_service.GenerateService.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// @Summary		导入表结构
// @Description	导入数据库表结构
// @Tags			generator-代码生成器
// @Param			token	header		string				true	"token"
// @Param			tables	query		string				true	"表名列表(逗号分隔)"
// @Success		200		{object}	response.Response	"成功"
// @Router			/api/admin/generator/importTable [get]
func (gh GenHandler) ImportTable(c *gin.Context) {
	var importReq generator_schema.ImportTableReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &importReq)) {
		return
	}
	err := generator_service.GenerateService.ImportTable(strings.Split(importReq.Tables, ","))
	response.CheckAndRespWithData(c, nil, err)
}

// @Summary		同步表结构
// @Description	同步数据库表结构
// @Tags			generator-代码生成器
// @Param			token	header		string				true	"token"
// @Param			id		query		string				true	"主键"
// @Success		200		{object}	response.Response	"成功"
// @Router			/api/admin/generator/syncTable [get]
func (gh GenHandler) SyncTable(c *gin.Context) {
	var syncReq generator_schema.SyncTableReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &syncReq)) {
		return
	}
	err := generator_service.GenerateService.SyncTable(syncReq.ID)
	response.CheckAndRespWithData(c, nil, err)
}

// @Summary		编辑表结构
// @Description	编辑表结构信息
// @Tags			generator-代码生成器
// @Param			token			header		string							true	"token"
// @Param			id				body		string							true	"主键"
// @Param			tableName		body		string							true	"表名"
// @Param			entityName		body		string							true	"实体名称"
// @Param			tableComment	body		string							true	"表描述"
// @Param			authorName		body		string							false	"作者名称"
// @Param			remarks			body		string							false	"备注信息"
// @Param			genTpl			body		string							false	"生成模板方式: [crud=单表, tree=树表]"
// @Param			moduleName		body		string							true	"模块名"
// @Param			functionName	body		string							true	"功能名"
// @Param			treePrimary		body		string							false	"树表主键"
// @Param			treeParent		body		string							false	"树表父键"
// @Param			treeName		body		string							false	"树表名称"
// @Param			subTableName	body		string							false	"子表名称"
// @Param			subTableFk		body		string							false	"子表外键"
// @Param			columns			body		[]generator_schema.EditColumn	true	"字段列表"
// @Success		200				{object}	response.Response				"成功"
// @Router			/api/admin/generator/editTable [post]
func (gh GenHandler) EditTable(c *gin.Context) {
	var editReq generator_schema.EditTableReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	err := generator_service.GenerateService.EditTable(editReq)
	response.CheckAndRespWithData(c, nil, err)
}

// @Summary		删除表结构
// @Description	删除表结构
// @Tags			generator-代码生成器
// @Param			token	header		string				true	"token"
// @Param			ids		body		[]string			true	"主键列表"
// @Success		200		{object}	response.Response	"成功"
// @Router			/api/admin/generator/delTable [post]
func (gh GenHandler) DelTable(c *gin.Context) {
	var delReq generator_schema.DelTableReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	err := generator_service.GenerateService.DelTable(delReq.Ids)
	response.CheckAndRespWithData(c, nil, err)
}

// @Summary		预览代码
// @Description	预览生成的代码
// @Tags			generator-代码生成器
// @Param			token	header		string										true	"token"
// @Param			id		query		string										true	"主键"
// @Success		200		{object}	response.Response{data=map[string]string}	"成功"
// @Router			/api/admin/generator/previewCode [get]
func (gh GenHandler) PreviewCode(c *gin.Context) {
	var previewReq generator_schema.PreviewCodeReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &previewReq)) {
		return
	}
	res, err := generator_service.GenerateService.PreviewCode(previewReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// @Summary		下载代码
// @Description	下载生成的代码
// @Tags			generator-代码生成器
// @Param			token	header	string	true	"token"
// @Param			tables	query	string	true	"表名列表(逗号分隔)"
// @Success		200		"文件流(zip格式)"
// @Router			/api/admin/generator/downloadCode [get]
func (gh GenHandler) DownloadCode(c *gin.Context) {
	var downloadReq generator_schema.DownloadReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &downloadReq)) {
		return
	}
	zipBytes, err := generator_service.GenerateService.DownloadCode(strings.Split(downloadReq.Tables, ","))
	if response.IsFailWithResp(c, err) {
		return
	}
	contentType := "application/zip"
	c.Header("Content-Type", contentType)
	c.Header("Content-Disposition", "attachment; filename=gen-"+downloadReq.Tables+".zip")
	c.Data(http.StatusOK, contentType, zipBytes)
}
