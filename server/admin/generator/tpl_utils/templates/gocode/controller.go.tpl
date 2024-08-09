package {{{ .ModuleName }}}

import (
	"net/http"
	"strconv"
	"time"
	"github.com/gin-gonic/gin" 
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util"
	"x_admin/util/excel"
	"golang.org/x/sync/singleflight"
)

 
type {{{ toUpperCamelCase .ModuleName }}}Handler struct {
	requestGroup singleflight.Group
}

//  @Summary	{{{ .FunctionName }}}列表
//  @Tags		{{{ .ModuleName }}}-{{{ .FunctionName }}}
//  @Produce	json
//  @Param		Token		header		string				true	"token"
//  @Param		PageNo		query		int					true	"页码"
//  @Param		PageSize	query		int					true	"每页数量"
{{{- range .Columns }}}
{{{- if .IsQuery }}}
{{{- if eq .HtmlType "datetime" }}}
//	@Param {{{ toCamelCase .GoField }}}Start  query {{{ .GoType }}} false "{{{ .ColumnComment }}}"
//	@Param {{{ toCamelCase .GoField }}}End  query {{{ .GoType }}} false "{{{ .ColumnComment }}}"	
{{{- else }}}
//	@Param {{{ toCamelCase .GoField }}} query {{{ .GoType }}} false "{{{ .ColumnComment }}}"
{{{- end }}}
{{{- end }}}
{{{- end }}}
//@Success 200	{object} {{{getPageResp (toUpperCamelCase .EntityName)  }}}	"成功"
//@Router	/api/admin/{{{ .ModuleName }}}/list [get]
func (hd *{{{  toUpperCamelCase .ModuleName }}}Handler) List(c *gin.Context) {
	var page request.PageReq
	var listReq {{{ toUpperCamelCase .EntityName }}}ListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := {{{ toUpperCamelCase .EntityName }}}Service.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

//	@Summary	{{{ .FunctionName }}}列表-所有
//	@Tags		{{{ .ModuleName }}}-{{{ .FunctionName }}}
//  @Produce	json
{{{- range .Columns }}}
{{{- if .IsQuery }}}
{{{- if eq .HtmlType "datetime" }}}
//	@Param {{{ toCamelCase .GoField }}}Start  query {{{ .GoType }}} false "{{{ .ColumnComment }}}"
//	@Param {{{ toCamelCase .GoField }}}End  query {{{ .GoType }}} false "{{{ .ColumnComment }}}"	
{{{- else }}}
//	@Param {{{ toCamelCase .GoField }}} query {{{ .GoType }}} false "{{{ .ColumnComment }}}"
{{{- end }}}
{{{- end }}}
{{{- end }}}
//	@Success	200			{object}	response.Response{ data=[]{{{ toUpperCamelCase .EntityName }}}Resp}	"成功"
//	@Router		/api/admin/{{{ .ModuleName }}}/listAll [get]
func (hd *{{{  toUpperCamelCase .ModuleName }}}Handler) ListAll(c *gin.Context) {
	var listReq {{{ toUpperCamelCase .EntityName }}}ListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := {{{ toUpperCamelCase .EntityName }}}Service.ListAll(listReq)
	response.CheckAndRespWithData(c, res, err)
}

//	@Summary	{{{ .FunctionName }}}详情
//	@Tags		{{{ .ModuleName }}}-{{{ .FunctionName }}}
//	@Produce	json
//	@Param		Token		header		string				true	"token"
{{{- range .Columns }}}
{{{- if .IsPk }}}
//	@Param		{{{ toCamelCase .GoField }}}		query		{{{ .GoType }}}				false	"{{{ .ColumnComment }}}"
{{{- end }}}
{{{- end }}}
//	@Success	200			{object}	response.Response{ data={{{ toUpperCamelCase .EntityName }}}Resp}	"成功"
//	@Router		/api/admin/{{{ .ModuleName }}}/detail [get]
func (hd *{{{  toUpperCamelCase .ModuleName }}}Handler) Detail(c *gin.Context) {
	var detailReq {{{ toUpperCamelCase .EntityName }}}DetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err, _ := hd.requestGroup.Do("{{{ toUpperCamelCase .EntityName }}}:Detail:"+strconv.Itoa(detailReq.{{{ toUpperCamelCase .PrimaryKey }}}), func() (any, error) {
		v, err := {{{ toUpperCamelCase .EntityName }}}Service.Detail(detailReq.{{{ toUpperCamelCase .PrimaryKey }}})
		return v, err
	})

	response.CheckAndRespWithData(c, res, err)
}


//	@Summary	{{{ .FunctionName }}}新增
//	@Tags		{{{ .ModuleName }}}-{{{ .FunctionName }}}
//	@Produce	json
//	@Param		Token		header		string				true	"token"
{{{- range .Columns }}}
{{{- if .IsInsert }}}
//	@Param		{{{ toCamelCase .GoField }}}		body		{{{ .GoType }}}				false	"{{{ .ColumnComment }}}"
{{{- end }}}
{{{- end }}}
//	@Success	200			{object}	response.Response	"成功"
//	@Router		/api/admin/{{{ .ModuleName }}}/add [post]
func (hd *{{{  toUpperCamelCase .ModuleName }}}Handler) Add(c *gin.Context) {
	var addReq {{{ toUpperCamelCase .EntityName }}}AddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	createId, e := {{{ toUpperCamelCase .EntityName }}}Service.Add(addReq)
	response.CheckAndRespWithData(c,createId, e)
}
//	@Summary	{{{ .FunctionName }}}编辑
//	@Tags		{{{ .ModuleName }}}-{{{ .FunctionName }}}
//	@Produce	json
//	@Param		Token		header		string				true	"token"
{{{- range .Columns }}}
{{{- if .IsEdit }}}
//	@Param		{{{ toCamelCase .GoField }}}		body		{{{ .GoType }}}				false	"{{{ .ColumnComment }}}"
{{{- end }}}
{{{- end }}}
//	@Success	200			{object}	response.Response	"成功"
//	@Router		/api/admin/{{{ .ModuleName }}}/edit [post]
func (hd *{{{  toUpperCamelCase .ModuleName }}}Handler) Edit(c *gin.Context) {
	var editReq {{{ toUpperCamelCase .EntityName }}}EditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	response.CheckAndRespWithData(c,editReq.{{{ toUpperCamelCase .PrimaryKey }}}, {{{ toUpperCamelCase .EntityName }}}Service.Edit(editReq))
}
//	@Summary	{{{ .FunctionName }}}删除
//	@Tags		{{{ .ModuleName }}}-{{{ .FunctionName }}}
//	@Produce	json
//	@Param		Token		header		string				true	"token"
{{{- range .Columns }}}
{{{- if .IsPk }}}
//	@Param		{{{ toCamelCase .GoField }}}		body		{{{ .GoType }}}				false	"{{{ .ColumnComment }}}"
{{{- end }}}
{{{- end }}}
//	@Success	200			{object}	response.Response	"成功"
//	@Router		/api/admin/{{{ .ModuleName }}}/del [post]
func (hd *{{{  toUpperCamelCase .ModuleName }}}Handler) Del(c *gin.Context) {
	var delReq {{{ toUpperCamelCase .EntityName }}}DelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, {{{ toUpperCamelCase .EntityName }}}Service.Del(delReq.{{{ toUpperCamelCase .PrimaryKey }}}))
}



//	@Summary	{{{ .FunctionName }}}导出
//	@Tags		{{{ .ModuleName }}}-{{{ .FunctionName }}}
//	@Produce	json
//	@Param		Token		header		string				true	"token"
{{{- range .Columns }}}
{{{- if .IsQuery }}}
{{{- if eq .HtmlType "datetime" }}}
//	@Param {{{ toCamelCase .GoField }}}Start  query {{{ .GoType }}} false "{{{ .ColumnComment }}}"
//	@Param {{{ toCamelCase .GoField }}}End  query {{{ .GoType }}} false "{{{ .ColumnComment }}}"	
{{{- else }}}
//	@Param {{{ toCamelCase .GoField }}} query {{{ .GoType }}} false "{{{ .ColumnComment }}}"
{{{- end }}}
{{{- end }}}
{{{- end }}}
//	@Router		/api/admin/{{{ .ModuleName }}}/ExportFile [get]
func (hd *{{{  toUpperCamelCase .ModuleName }}}Handler) ExportFile(c *gin.Context) {
	var listReq {{{ toUpperCamelCase .EntityName }}}ListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := {{{ toUpperCamelCase .EntityName }}}Service.ExportFile(listReq)
	if err != nil {
		response.FailWithMsg(c, response.SystemError, "查询信息失败")
		return
	}
	f, err := excel.NormalDynamicExport(res, "Sheet1", "{{{ .FunctionName }}}", nil)
	if err != nil {
		response.FailWithMsg(c, response.SystemError, "导出失败")
		return
	}
	excel.DownLoadExcel("{{{ .FunctionName }}}" + time.Now().Format("20060102-150405"), c.Writer, f)
}

//  @Summary	{{{ .FunctionName }}}导入
//  @Tags		{{{ .ModuleName }}}-{{{ .FunctionName }}}
//  @Produce	json
//	@Router		/api/admin/{{{ .ModuleName }}}/ImportFile [post]
func (hd *{{{  toUpperCamelCase .ModuleName }}}Handler) ImportFile(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusInternalServerError, "文件不存在")
		return
	}
	defer file.Close()
	importList := []{{{ toUpperCamelCase .EntityName }}}Resp{}
	err = excel.GetExcelData(file, &importList)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	err = {{{ toUpperCamelCase .EntityName }}}Service.ImportFile(importList)
	response.CheckAndResp(c, err)
}