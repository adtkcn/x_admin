package schema
import (
	"x_admin/core"
)

type {{{ toUpperCamelCase .EntityName }}}Primarykey struct {
    {{{- range .Columns }}}
    {{{- if .IsPk }}}
        {{{ .GoField }}} {{{.GoType }}} // {{{ .ColumnComment }}}
    {{{- end }}}
    {{{- end }}}
}

//{{{ toUpperCamelCase .EntityName }}}ListReq {{{ .FunctionName }}}列表参数
type {{{ toUpperCamelCase .EntityName }}}ListReq struct {
    {{{- range .Columns }}}
    {{{- if .IsQuery }}}
        {{{- if eq .HtmlType "datetime" }}}
            {{{ .GoField }}}Start x_null.String // 开始{{{ .ColumnComment }}}
            {{{ .GoField }}}End x_null.String // 结束{{{ .ColumnComment }}}
        {{{- else }}}
            {{{ .GoField }}} {{{ .GoNullType }}} // {{{ .ColumnComment }}}
        {{{- end }}}
    {{{- end }}}
    {{{- end }}}
}



//{{{ toUpperCamelCase .EntityName }}}AddReq {{{ .FunctionName }}}新增参数
type {{{ toUpperCamelCase .EntityName }}}AddReq struct {
    {{{- range .Columns }}}
    {{{- if .IsInsert }}}
    {{{ .GoField }}}  {{{ .GoNullType }}} `binding:"{{{ if eq .IsRequired 1 }}}required;{{{ end }}}"`  // {{{ .ColumnComment }}}
    {{{- end }}}
    {{{- end }}}
}

//{{{ toUpperCamelCase .EntityName }}}EditReq {{{ .FunctionName }}}编辑参数
type {{{ toUpperCamelCase .EntityName }}}EditReq struct {
    {{{- range .Columns }}}
    {{{- if .IsEdit }}}
        {{{- if .IsPk }}}
    {{{ .GoField }}} {{{ .GoType }}} // {{{ .ColumnComment }}}
        {{{- else }}}
    {{{ .GoField }}}  {{{ .GoNullType }}}  // {{{ .ColumnComment }}}
        {{{- end }}}
    {{{- end }}}
    {{{- end }}}
}

 

//{{{ toUpperCamelCase .EntityName }}}DelBatchReq {{{ .FunctionName }}}批量删除参数
type {{{ toUpperCamelCase .EntityName }}}DelBatchReq struct {
	Ids string
}

//{{{ toUpperCamelCase .EntityName }}}Resp {{{ .FunctionName }}}返回信息
type {{{ toUpperCamelCase .EntityName }}}Resp struct {
	{{{- range .Columns }}}
    {{{- if or .IsList .IsPk }}}
    {{{- if .IsPk }}}
        {{{ .GoField }}} {{{.GoType }}} `swaggertype:"{{{ .SwagType }}}"`// {{{ .ColumnComment }}}
    {{{- else }}}
        {{{ .GoField }}} {{{ .GoNullType }}} `swaggertype:"{{{ .SwagType }}}"`// {{{ .ColumnComment }}}
    {{{- end }}}
    {{{- end }}}
    {{{- end }}}
}
