package schema
import (
	"x_admin/core"
)

type {{{ toUpperCamelCase .EntityName }}}Primarykey struct {
    {{{- range .Columns }}}
    {{{- if .IsPk }}}
        {{{ toUpperCamelCase .GoField }}} {{{.GoType }}} // {{{ .ColumnComment }}}
    {{{- end }}}
    {{{- end }}}
}

//{{{ toUpperCamelCase .EntityName }}}ListReq {{{ .FunctionName }}}列表参数
type {{{ toUpperCamelCase .EntityName }}}ListReq struct {
    {{{- range .Columns }}}
    {{{- if .IsQuery }}}
        {{{- if eq .ColumnName "created_by" }}}
            {{{ toUpperCamelCase .GoField }}} {{{goWithAddEditType .GoType }}} // {{{ .ColumnComment }}}
            Nickname core.NullString // 创建人昵称
        {{{- else if eq .HtmlType "datetime" }}}
            {{{ toUpperCamelCase .GoField }}}Start core.NullString // 开始{{{ .ColumnComment }}}
            {{{ toUpperCamelCase .GoField }}}End core.NullString // 结束{{{ .ColumnComment }}}
        {{{- else }}}
            {{{ toUpperCamelCase .GoField }}} {{{goWithAddEditType .GoType }}} // {{{ .ColumnComment }}}
        {{{- end }}}
    {{{- end }}}
    {{{- end }}}
}



//{{{ toUpperCamelCase .EntityName }}}AddReq {{{ .FunctionName }}}新增参数
type {{{ toUpperCamelCase .EntityName }}}AddReq struct {
    {{{- range .Columns }}}
    {{{- if .IsInsert }}}
    {{{ toUpperCamelCase .GoField }}}  {{{goWithAddEditType .GoType }}} `binding:"{{{ if eq .IsRequired 1 }}}required;{{{ end }}}"`  // {{{ .ColumnComment }}}
    {{{- end }}}
    {{{- end }}}
}

//{{{ toUpperCamelCase .EntityName }}}EditReq {{{ .FunctionName }}}编辑参数
type {{{ toUpperCamelCase .EntityName }}}EditReq struct {
    {{{- range .Columns }}}
    {{{- if .IsEdit }}}
        {{{- if .IsPk }}}
    {{{ toUpperCamelCase .GoField }}} {{{ .GoType }}} // {{{ .ColumnComment }}}
        {{{- else }}}
    {{{ toUpperCamelCase .GoField }}}  {{{goWithAddEditType .GoType }}}  // {{{ .ColumnComment }}}
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
        {{{ toUpperCamelCase .GoField }}} {{{.GoType }}} `swaggertype:"{{{goToTsType .GoType }}}"`// {{{ .ColumnComment }}}
    {{{- else }}}
        {{{ toUpperCamelCase .GoField }}} {{{goWithRespType .GoType }}} `swaggertype:"{{{goToTsType .GoType }}}"`// {{{ .ColumnComment }}}
    {{{- end }}}
    {{{- end }}}
    {{{- end }}}
}
