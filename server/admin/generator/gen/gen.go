package gen

import (
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
	var tbReq DbTablesReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &tbReq)) {
		return
	}
	res, err := Service.DbTables(page, tbReq)
	response.CheckAndRespWithData(c, res, err)
}

// list 生成列表
func (gh genHandler) List(c *gin.Context) {
	var page request.PageReq
	var listReq ListTableReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := Service.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

// detail 生成详情
func (gh genHandler) Detail(c *gin.Context) {
	var detailReq DetailTableReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := Service.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// importTable 导入表结构
func (gh genHandler) importTable(c *gin.Context) {
	var importReq ImportTableReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &importReq)) {
		return
	}
	err := Service.ImportTable(strings.Split(importReq.Tables, ","))
	response.CheckAndResp(c, err)
}

// syncTable 同步表结构
func (gh genHandler) syncTable(c *gin.Context) {
	var syncReq SyncTableReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &syncReq)) {
		return
	}
	err := Service.SyncTable(syncReq.ID)
	response.CheckAndResp(c, err)
}

// editTable 编辑表结构
func (gh genHandler) editTable(c *gin.Context) {
	var editReq EditTableReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	err := Service.EditTable(editReq)
	response.CheckAndResp(c, err)
}

// delTable 删除表结构
func (gh genHandler) delTable(c *gin.Context) {
	var delReq DelTableReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	err := Service.DelTable(delReq.Ids)
	response.CheckAndResp(c, err)
}

// previewCode 预览代码
func (gh genHandler) previewCode(c *gin.Context) {
	var previewReq PreviewCodeReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &previewReq)) {
		return
	}
	res, err := Service.PreviewCode(previewReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// // genCode 生成代码,生成到服务器自定义路径存在风险，禁止使用
// func (gh genHandler) genCode(c *gin.Context) {
// 	var genReq GenCodeReq
// 	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &genReq)) {
// 		return
// 	}
// 	for _, table := range strings.Split(genReq.Tables, ",") {
// 		err := Service.GenCode(table)
// 		if response.IsFailWithResp(c, err) {
// 			return
// 		}
// 	}
// 	response.Ok(c)
// }

// downloadCode 下载代码
func (gh genHandler) downloadCode(c *gin.Context) {
	var downloadReq DownloadReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &downloadReq)) {
		return
	}
	zipBytes, err := Service.DownloadCode(strings.Split(downloadReq.Tables, ","))
	if response.IsFailWithResp(c, err) {
		return
	}
	contentType := "application/zip"
	c.Header("Content-Type", contentType)
	c.Header("Content-Disposition", "attachment; filename=gen-"+downloadReq.Tables+".zip")
	c.Data(http.StatusOK, contentType, zipBytes)
}
