package {{{ .ModuleName }}}
import (
	"x_admin/core"

)

//{{{ toUpperCamelCase .EntityName }}}ListReq {{{ .FunctionName }}}列表参数
type {{{ toUpperCamelCase .EntityName }}}ListReq struct {
    {{{- range .Columns }}}
    {{{- if .IsQuery }}}
        {{{- if eq .HtmlType "datetime" }}}
            {{{ toUpperCamelCase .GoField }}}Start *string `` // 开始{{{ .ColumnComment }}}
            {{{ toUpperCamelCase .GoField }}}End *string `` // 结束{{{ .ColumnComment }}}
        {{{- else }}}
            {{{ toUpperCamelCase .GoField }}} *{{{.GoType }}} `` // {{{ .ColumnComment }}}
        {{{- end }}}
    {{{- end }}}
    {{{- end }}}
}



//{{{ toUpperCamelCase .EntityName }}}AddReq {{{ .FunctionName }}}新增参数
type {{{ toUpperCamelCase .EntityName }}}AddReq struct {
    {{{- range .Columns }}}
    {{{- if .IsInsert }}}
    {{{ toUpperCamelCase .GoField }}}  {{{goWithAddEditType .GoType }}}  `` // {{{ .ColumnComment }}}
    {{{- end }}}
    {{{- end }}}
}

//{{{ toUpperCamelCase .EntityName }}}EditReq {{{ .FunctionName }}}编辑参数
type {{{ toUpperCamelCase .EntityName }}}EditReq struct {
    {{{- range .Columns }}}
    {{{- if .IsEdit }}}
        {{{- if .IsPk }}}
        {{{ toUpperCamelCase .GoField }}} {{{ .GoType }}} `` // {{{ .ColumnComment }}}
        {{{- else }}}
        {{{ toUpperCamelCase .GoField }}}  {{{goWithAddEditType .GoType }}}  `` // {{{ .ColumnComment }}}
        {{{- end }}}
    {{{- end }}}
    {{{- end }}}
}

//{{{ toUpperCamelCase .EntityName }}}DetailReq {{{ .FunctionName }}}详情参数
type {{{ toUpperCamelCase .EntityName }}}DetailReq struct {
    {{{- range .Columns }}}
    {{{- if .IsPk }}}
    {{{ toUpperCamelCase .GoField }}} {{{.GoType }}} `` // {{{ .ColumnComment }}}
    {{{- end }}}
    {{{- end }}}
}

//{{{ toUpperCamelCase .EntityName }}}DelReq {{{ .FunctionName }}}删除参数
type {{{ toUpperCamelCase .EntityName }}}DelReq struct {
    {{{- range .Columns }}}
    {{{- if .IsPk }}}
    {{{ toUpperCamelCase .GoField }}} {{{ .GoType }}} `` // {{{ .ColumnComment }}}
    {{{- end }}}
    {{{- end }}}
}

//{{{ toUpperCamelCase .EntityName }}}Resp {{{ .FunctionName }}}返回信息
type {{{ toUpperCamelCase .EntityName }}}Resp struct {
	{{{- range .Columns }}}
    {{{- if or .IsList .IsPk }}}
    {{{- if .IsPk }}}
        {{{ toUpperCamelCase .GoField }}} {{{.GoType }}} `` // {{{ .ColumnComment }}}
    {{{- else }}}
        {{{ toUpperCamelCase .GoField }}} {{{goWithRespType .GoType }}} `` // {{{ .ColumnComment }}}
    {{{- end }}}
    {{{- end }}}
    {{{- end }}}
}
