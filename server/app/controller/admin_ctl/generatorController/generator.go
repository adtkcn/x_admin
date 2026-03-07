package generatorController

import (
	"x_admin/app/schema/generatorSchema"
	"x_admin/app/service/generatorService"
	"x_admin/core/request"
	"x_admin/core/response"

	"net/http"
	"strings"
	"x_admin/middleware"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

func GenRoute(rg *gin.RouterGroup) {
	handle := genHandler{}

	rg = rg.Group("/gen", middleware.TokenAuth())
	rg.GET("/db", handle.dbTables)
	rg.GET("/list", handle.List)
	rg.GET("/detail", handle.Detail)
	rg.POST("/importTable", handle.importTable)
	rg.POST("/syncTable", handle.syncTable)
	rg.POST("/editTable", handle.editTable)
	rg.POST("/delTable", handle.delTable)
	rg.GET("/previewCode", handle.previewCode)
	rg.GET("/downloadCode", handle.downloadCode)
}

type genHandler struct {
	// Service IGenerateService
}

// dbTables 数据表列表
func (gh genHandler) dbTables(c *gin.Context) {
	var page request.PageReq
	var tbReq generatorSchema.DbTablesReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &tbReq)) {
		return
	}
	res, err := generatorService.GenerateService.DbTables(page, tbReq)
	response.CheckAndRespWithData(c, res, err)
}

// list 生成列表
func (gh genHandler) List(c *gin.Context) {
	var page request.PageReq
	var listReq generatorSchema.ListTableReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := generatorService.GenerateService.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

// detail 生成详情
func (gh genHandler) Detail(c *gin.Context) {
	var detailReq generatorSchema.DetailTableReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := generatorService.GenerateService.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// importTable 导入表结构
func (gh genHandler) importTable(c *gin.Context) {
	var importReq generatorSchema.ImportTableReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &importReq)) {
		return
	}
	err := generatorService.GenerateService.ImportTable(strings.Split(importReq.Tables, ","))
	response.CheckAndRespWithData(c, nil, err)
}

// syncTable 同步表结构
func (gh genHandler) syncTable(c *gin.Context) {
	var syncReq generatorSchema.SyncTableReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &syncReq)) {
		return
	}
	err := generatorService.GenerateService.SyncTable(syncReq.ID)
	response.CheckAndRespWithData(c, nil, err)
}

// editTable 编辑表结构
func (gh genHandler) editTable(c *gin.Context) {
	var editReq generatorSchema.EditTableReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	err := generatorService.GenerateService.EditTable(editReq)
	response.CheckAndRespWithData(c, nil, err)
}

// delTable 删除表结构
func (gh genHandler) delTable(c *gin.Context) {
	var delReq generatorSchema.DelTableReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	err := generatorService.GenerateService.DelTable(delReq.Ids)
	response.CheckAndRespWithData(c, nil, err)
}

// previewCode 预览代码
func (gh genHandler) previewCode(c *gin.Context) {
	var previewReq generatorSchema.PreviewCodeReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &previewReq)) {
		return
	}
	res, err := generatorService.GenerateService.PreviewCode(previewReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// downloadCode 下载代码
func (gh genHandler) downloadCode(c *gin.Context) {
	var downloadReq generatorSchema.DownloadReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &downloadReq)) {
		return
	}
	zipBytes, err := generatorService.GenerateService.DownloadCode(strings.Split(downloadReq.Tables, ","))
	if response.IsFailWithResp(c, err) {
		return
	}
	contentType := "application/zip"
	c.Header("Content-Type", contentType)
	c.Header("Content-Disposition", "attachment; filename=gen-"+downloadReq.Tables+".zip")
	c.Data(http.StatusOK, contentType, zipBytes)
}
