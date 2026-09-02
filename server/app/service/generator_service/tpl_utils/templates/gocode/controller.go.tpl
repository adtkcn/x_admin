package {{{.Domain}}}_controller

import (
	"fmt"
	"strings"
	"time"
	"github.com/gin-gonic/gin" 
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/config"
	"x_admin/util"
	"x_admin/util/excel2"
	"golang.org/x/sync/singleflight"
	"x_admin/app/schema/{{{.Domain}}}_schema"
	"x_admin/app/service/{{{.Domain}}}_service"
)

 
type {{{ toUpperCamelCase .ModuleName }}}Handler struct {
	requestGroup singleflight.Group
}

//  @Summary	{{{ .FunctionName }}}列表
//  @Tags		{{{ .ModuleName }}}-{{{ .FunctionName }}}
//  @Produce	json
//  @Param		token		header		string				true	"token"
//  @Param		pageNo		query		int					true	"页码"
//  @Param		pageSize	query		int					true	"每页数量"
{{{- range .Columns }}}
{{{- if .IsQuery }}}
{{{- if eq .HtmlType "datetime" }}}
//	@Param {{{ .TsField }}}_start  query {{{.TsType }}} false "开始{{{ .ColumnComment }}}"
//	@Param {{{ .TsField }}}_end  query {{{.TsType }}} false "结束{{{ .ColumnComment }}}"	
{{{- else }}}
//	@Param {{{ .TsField }}} query {{{.TsType }}} false "{{{ .ColumnComment }}}"
{{{- end }}}
{{{- end }}}
{{{- end }}}
//@Success 200	{object} {{{getPageResp (toUpperCamelCase .EntityName)  }}}	"成功"
//@Router	/api/admin/{{{ .ModuleName }}}/list [get]
func (hd *{{{  toUpperCamelCase .ModuleName }}}Handler) List(c *gin.Context) {
	var page request.PageReq
	var listReq {{{.Domain}}}_schema.{{{ toUpperCamelCase .EntityName }}}ListReq
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := {{{.Domain}}}_service.{{{ toUpperCamelCase .EntityName }}}Service.List(page, listReq)
	response.JSON(c, res, err)
}

//	@Summary	{{{ .FunctionName }}}列表-所有
//	@Tags		{{{ .ModuleName }}}-{{{ .FunctionName }}}
//  @Produce	json
//  @Param		token		header		string				true	"token"
{{{- range .Columns }}}
{{{- if .IsQuery }}}	
{{{- if eq .HtmlType "datetime" }}}
//	@Param {{{ .TsField }}}_start  query {{{.TsType }}} false "开始{{{ .ColumnComment }}}"
//	@Param {{{ .TsField }}}_end  query {{{.TsType }}} false "结束{{{ .ColumnComment }}}"	
{{{- else }}}
//	@Param {{{ .TsField }}} query {{{.TsType }}} false "{{{ .ColumnComment }}}"
{{{- end }}}
{{{- end }}}
{{{- end }}}
//	@Success	200			{object}	response.Response{data=[]{{{.Domain}}}_schema.{{{ toUpperCamelCase .EntityName }}}Resp}	"成功"
//	@Router		/api/admin/{{{ .ModuleName }}}/list_all [get]
func (hd *{{{  toUpperCamelCase .ModuleName }}}Handler) ListAll(c *gin.Context) {
	var listReq {{{.Domain}}}_schema.{{{ toUpperCamelCase .EntityName }}}ListReq
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := {{{.Domain}}}_service.{{{ toUpperCamelCase .EntityName }}}Service.ListAll(listReq)
	response.JSON(c, res, err)
}

//	@Summary	{{{ .FunctionName }}}详情
//	@Tags		{{{ .ModuleName }}}-{{{ .FunctionName }}}
//	@Produce	json
//	@Param		token		header		string				true	"token"
{{{- range .Columns }}}
{{{- if .IsPk }}}
//	@Param		{{{ .TsField }}}		query		{{{.TsType }}}				false	"{{{ .ColumnComment }}}"
{{{- end }}}
{{{- end }}}
//	@Success	200			{object}	response.Response{data={{{.Domain}}}_schema.{{{ toUpperCamelCase .EntityName }}}Resp}	"成功"
//	@Router		/api/admin/{{{ .ModuleName }}}/detail [get]
func (hd *{{{  toUpperCamelCase .ModuleName }}}Handler) Detail(c *gin.Context) {
	var detailReq {{{.Domain}}}_schema.{{{ toUpperCamelCase .EntityName }}}Primarykey
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err, _ := hd.requestGroup.Do(fmt.Sprintf("{{{ toUpperCamelCase .EntityName }}}:Detail:%v", detailReq.{{{ .PrimaryGoField }}}), func() (any, error) {
		v, err := {{{.Domain}}}_service.{{{ toUpperCamelCase .EntityName }}}Service.Detail(detailReq.{{{ .PrimaryGoField }}})
		return v, err
	})

	response.JSON(c, res, err)
}


//	@Summary	{{{ .FunctionName }}}新增
//	@Tags		{{{ .ModuleName }}}-{{{ .FunctionName }}}
//	@Produce	json
//	@Param		token		header		string				true	"token"
{{{- range .Columns }}}
{{{- if .IsInsert }}}
//	@Param		{{{ .TsField }}}		body		{{{.TsType }}}				false	"{{{ .ColumnComment }}}"
{{{- end }}}
{{{- end }}}
//	@Success	200			{object}	response.Response	"成功"
//	@Router		/api/admin/{{{ .ModuleName }}}/add [post]
func (hd *{{{  toUpperCamelCase .ModuleName }}}Handler) Add(c *gin.Context) {
	var addReq {{{.Domain}}}_schema.{{{ toUpperCamelCase .EntityName }}}AddReq
	if response.IsFail(c, util.VerifyUtil.VerifyJSON(c, &addReq)) {
		return
	}
	
	var adminId = config.AdminConfig.GetAdminId(c)// 创建人
	
	createId, e := {{{.Domain}}}_service.{{{ toUpperCamelCase .EntityName }}}Service.Add(addReq, adminId)
	response.JSON(c,createId, e)
}
//	@Summary	{{{ .FunctionName }}}编辑
//	@Tags		{{{ .ModuleName }}}-{{{ .FunctionName }}}
//	@Produce	json
//	@Param		token		header		string				true	"token"
{{{- range .Columns }}}
{{{- if .IsEdit }}}
//	@Param		{{{ .TsField }}}		body		{{{.TsType }}}				false	"{{{ .ColumnComment }}}"
{{{- end }}}
{{{- end }}}
//	@Success	200			{object}	response.Response	"成功"
//	@Router		/api/admin/{{{ .ModuleName }}}/edit [post]
func (hd *{{{  toUpperCamelCase .ModuleName }}}Handler) Edit(c *gin.Context) {
	var editReq {{{.Domain}}}_schema.{{{ toUpperCamelCase .EntityName }}}EditReq
	if response.IsFail(c, util.VerifyUtil.VerifyJSON(c, &editReq)) {
		return
	}
	response.JSON(c,editReq.{{{ .PrimaryGoField }}}, {{{.Domain}}}_service.{{{ toUpperCamelCase .EntityName }}}Service.Edit(editReq))
}
//	@Summary	{{{ .FunctionName }}}删除
//	@Tags		{{{ .ModuleName }}}-{{{ .FunctionName }}}
//	@Produce	json
//	@Param		token		header		string				true	"token"
{{{- range .Columns }}}
{{{- if .IsPk }}}
//	@Param		{{{ .TsField }}}		body		{{{.TsType }}}				false	"{{{ .ColumnComment }}}"
{{{- end }}}
{{{- end }}}
//	@Success	200			{object}	response.Response	"成功"
//	@Router		/api/admin/{{{ .ModuleName }}}/del [post]
func (hd *{{{  toUpperCamelCase .ModuleName }}}Handler) Del(c *gin.Context) {
	var delReq {{{.Domain}}}_schema.{{{ toUpperCamelCase .EntityName }}}Primarykey
	if response.IsFail(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	response.JSON(c, nil, {{{.Domain}}}_service.{{{ toUpperCamelCase .EntityName }}}Service.Del(delReq.{{{ .PrimaryGoField }}}))
}

//	@Summary	{{{ .FunctionName }}}删除-批量
//	@Tags		{{{ .ModuleName }}}-{{{ .FunctionName }}}
// @Produce	json
// @Param		token		header		string				true	"token"
// @Param		ids		body		string				false	"逗号分割的id"
// @Success	200			{object}	response.Response	"成功"
// @Router		/api/admin/{{{ .ModuleName }}}/del_batch [post]
func (hd *{{{  toUpperCamelCase .ModuleName }}}Handler) DelBatch(c *gin.Context) {
	var delReq {{{.Domain}}}_schema.{{{ toUpperCamelCase .EntityName }}}DelBatchReq
	if response.IsFail(c, util.VerifyUtil.VerifyJSON(c, &delReq)) {
		return
	}
	if delReq.Ids == "" {
		response.FailMsg(c, "请选择要删除的数据")
		return
	}
	var ids = strings.Split(delReq.Ids, ",")

	response.JSON(c, nil, {{{.Domain}}}_service.{{{ toUpperCamelCase .EntityName }}}Service.DelBatch(ids))
}



//	@Summary	{{{ .FunctionName }}}导出
//	@Tags		{{{ .ModuleName }}}-{{{ .FunctionName }}}
//	@Produce	octet-stream,json
//	@Param		token		header		string				true	"token"
{{{- range .Columns }}}
{{{- if .IsQuery }}}
{{{- if eq .HtmlType "datetime" }}}
//	@Param {{{ .TsField }}}_start  query {{{.TsType }}} false "{{{ .ColumnComment }}}"
//	@Param {{{ .TsField }}}_end  query {{{.TsType }}} false "{{{ .ColumnComment }}}"	
{{{- else }}}
//	@Param {{{ .TsField }}} query {{{.TsType }}} false "{{{ .ColumnComment }}}"
{{{- end }}}
{{{- end }}}
{{{- end }}}
//  @Success	200		{file} string	"成功"
//  @Failure	500 	{object}	response.Response	"失败"
//	@Router		/api/admin/{{{ .ModuleName }}}/export_file [get]
func (hd *{{{  toUpperCamelCase .ModuleName }}}Handler) ExportFile(c *gin.Context) {
	var listReq {{{.Domain}}}_schema.{{{ toUpperCamelCase .EntityName }}}ListReq
	if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := {{{.Domain}}}_service.{{{ toUpperCamelCase .EntityName }}}Service.ExportFile(listReq)
	if err != nil {
		response.Fail(c, response.CheckErr(err, "查询信息失败"))
		return
	}
	f, err := excel2.Export(res,{{{.Domain}}}_service.{{{ toUpperCamelCase .EntityName }}}Service.GetExcelCol(), "Sheet1", "{{{ .FunctionName }}}")
	if err != nil {
		response.Fail(c, response.CheckErr(err, "导出失败"))
		return
	}
	excel2.DownLoadExcel("{{{ .FunctionName }}}" + time.Now().Format("20060102-150405"), c.Writer, f)
}

//  @Summary	{{{ .FunctionName }}}导入
//  @Tags		{{{ .ModuleName }}}-{{{ .FunctionName }}}
//  @Produce	json
//  @Param		token		header		string				true	"token"
//  @Param		file	formData	file	true	"导入文件"
//  @Success	200		{object}	response.Response	"成功"
//	@Router		/api/admin/{{{ .ModuleName }}}/import_file [post]
func (hd *{{{  toUpperCamelCase .ModuleName }}}Handler) ImportFile(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		response.Fail(c, response.CheckErr(err, "文件不存在"))
		return
	}
	defer file.Close()
	importList := []{{{.Domain}}}_schema.{{{ toUpperCamelCase .EntityName }}}Resp{}
	err = excel2.GetExcelData(file, &importList,{{{.Domain}}}_service.{{{ toUpperCamelCase .EntityName }}}Service.GetExcelCol())
	if err != nil {
		response.Fail(c, response.CheckErr(err, "文件解析失败"))
		return
	}

	err = {{{.Domain}}}_service.{{{ toUpperCamelCase .EntityName }}}Service.ImportFile(importList)
	response.JSON(c, nil, err)
}