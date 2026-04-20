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

// DbTables 数据表列表
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

// list 生成列表
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

// detail 生成详情
func (gh GenHandler) Detail(c *gin.Context) {
	var detailReq generator_schema.DetailTableReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := generator_service.GenerateService.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// ImportTable 导入表结构
func (gh GenHandler) ImportTable(c *gin.Context) {
	var importReq generator_schema.ImportTableReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &importReq)) {
		return
	}
	err := generator_service.GenerateService.ImportTable(strings.Split(importReq.Tables, ","))
	response.CheckAndRespWithData(c, nil, err)
}

// SyncTable 同步表结构
func (gh GenHandler) SyncTable(c *gin.Context) {
	var syncReq generator_schema.SyncTableReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &syncReq)) {
		return
	}
	err := generator_service.GenerateService.SyncTable(syncReq.ID)
	response.CheckAndRespWithData(c, nil, err)
}

// EditTable 编辑表结构
func (gh GenHandler) EditTable(c *gin.Context) {
	var editReq generator_schema.EditTableReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	err := generator_service.GenerateService.EditTable(editReq)
	response.CheckAndRespWithData(c, nil, err)
}

// DelTable 删除表结构
func (gh GenHandler) DelTable(c *gin.Context) {
	var delReq generator_schema.DelTableReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	err := generator_service.GenerateService.DelTable(delReq.Ids)
	response.CheckAndRespWithData(c, nil, err)
}

// PreviewCode 预览代码
func (gh GenHandler) PreviewCode(c *gin.Context) {
	var previewReq generator_schema.PreviewCodeReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &previewReq)) {
		return
	}
	res, err := generator_service.GenerateService.PreviewCode(previewReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// DownloadCode 下载代码
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
