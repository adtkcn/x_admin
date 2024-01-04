package {{{ .ModuleName }}}

import (
	"github.com/gin-gonic/gin" 
	"x_admin/core/request"
	"x_admin/core/response"
	"x_admin/util"
)

 
type {{{ title (toCamelCase .ModuleName) }}}Handler struct {}

//	@Summary	{{{ .FunctionName }}}列表
//	@Tags		{{{ .ModuleName }}}-{{{ .FunctionName }}}
//	@Produce	json
//	@Param		Token		header		string				true	"token"
//	@Param		PageNo		query		int					true	"页码"
//	@Param		PageSize	query		int					true	"每页数量"
{{{- range .Columns }}}
{{{- if .IsQuery }}}
//	@Param		{{{ toCamelCase .GoField }}}		query		{{{ .GoType }}}				false	"{{{ .ColumnComment }}}."
{{{- end }}}
{{{- end }}}
//	@Success	200			{object}	[]{{{ title (toCamelCase .EntityName) }}}Resp	"成功"
//	@Failure	400			{object}	string				"请求错误"
//	@Router		/api/admin/{{{ .ModuleName }}}/list [get]
func (hd {{{  title (toCamelCase .ModuleName) }}}Handler) List(c *gin.Context) {
	var page request.PageReq
	var listReq {{{ title (toCamelCase .EntityName) }}}ListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := Service.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

//	@Summary	{{{ .FunctionName }}}列表-所有
//	@Tags		{{{ .ModuleName }}}-{{{ .FunctionName }}}
//  @Produce	json
//	@Success	200			{object}	[]{{{ title (toCamelCase .EntityName) }}}Resp	"成功"
//	@Router		/api/admin/{{{ .ModuleName }}}/listAll [get]
func (hd {{{  title (toCamelCase .ModuleName) }}}Handler) ListAll(c *gin.Context) {
	res, err := Service.ListAll()
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	response.CheckAndRespWithData(c, res, err)
}

//	@Summary	{{{ .FunctionName }}}详情
//	@Tags		{{{ .ModuleName }}}-{{{ .FunctionName }}}
//	@Produce	json
//	@Param		Token		header		string				true	"token"
{{{- range .Columns }}}
{{{- if .IsPk }}}
//	@Param		{{{ toCamelCase .GoField }}}		query		{{{ .GoType }}}				false	"{{{ .ColumnComment }}}."
{{{- end }}}
{{{- end }}}
//	@Success	200			{object}	{{{ title (toCamelCase .EntityName) }}}Resp	"成功"
//	@Router		/api/admin/{{{ .ModuleName }}}/detail [get]
func (hd {{{  title (toCamelCase .ModuleName) }}}Handler) Detail(c *gin.Context) {
	var detailReq {{{ title (toCamelCase .EntityName) }}}DetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := Service.Detail(detailReq.{{{ title (toCamelCase .PrimaryKey) }}})
	response.CheckAndRespWithData(c, res, err)
}


//	@Summary	{{{ .FunctionName }}}新增
//	@Tags		{{{ .ModuleName }}}-{{{ .FunctionName }}}
//	@Produce	json
//	@Param		Token		header		string				true	"token"
{{{- range .Columns }}}
{{{- if .IsInsert }}}
//	@Param		{{{ toCamelCase .GoField }}}		body		{{{ .GoType }}}				false	"{{{ .ColumnComment }}}."
{{{- end }}}
{{{- end }}}
//	@Success	200			{object}	response.RespType	"成功"
//	@Router		/api/admin/{{{ .ModuleName }}}/add [post]
func (hd {{{  title (toCamelCase .ModuleName) }}}Handler) Add(c *gin.Context) {
	var addReq {{{ title (toCamelCase .EntityName) }}}AddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &addReq)) {
		return
	}
	response.CheckAndResp(c, Service.Add(addReq))
}
//	@Summary	{{{ .FunctionName }}}编辑
//	@Tags		{{{ .ModuleName }}}-{{{ .FunctionName }}}
//	@Produce	json
//	@Param		Token		header		string				true	"token"
{{{- range .Columns }}}
{{{- if .IsEdit }}}
//	@Param		{{{ toCamelCase .GoField }}}		body		{{{ .GoType }}}				false	"{{{ .ColumnComment }}}."
{{{- end }}}
{{{- end }}}
//	@Success	200			{object}	response.RespType	"成功"
//	@Router		/api/admin/{{{ .ModuleName }}}/edit [post]
func (hd {{{  title (toCamelCase .ModuleName) }}}Handler) Edit(c *gin.Context) {
	var editReq {{{ title (toCamelCase .EntityName) }}}EditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &editReq)) {
		return
	}
	response.CheckAndResp(c, Service.Edit(editReq))
}
//	@Summary	{{{ .FunctionName }}}删除
//	@Tags		{{{ .ModuleName }}}-{{{ .FunctionName }}}
//	@Produce	json
//	@Param		Token		header		string				true	"token"
{{{- range .Columns }}}
{{{- if .IsPk }}}
//	@Param		{{{ toCamelCase .GoField }}}		body		{{{ .GoType }}}				false	"{{{ .ColumnComment }}}."
{{{- end }}}
{{{- end }}}
//	@Success	200			{object}	response.RespType	"成功"
//	@Router		/api/admin/{{{ .ModuleName }}}/del [post]
func (hd {{{  title (toCamelCase .ModuleName) }}}Handler) Del(c *gin.Context) {
	var delReq {{{ title (toCamelCase .EntityName) }}}DelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, Service.Del(delReq.{{{ title (toCamelCase .PrimaryKey) }}}))
}
