package {{{.Domain}}}_schema
import (
    "x_admin/app/schema/system_schema"
	"github.com/adtkcn/x_null"
)

type {{{ toUpperCamelCase .EntityName }}}Primarykey struct {
    {{{- range .Columns }}}
    {{{- if .IsPk }}}
        {{{ .GoField }}} {{{.GoType }}} `form:"{{{ .TsField }}}" json:"{{{ .TsField }}}"` // {{{ .ColumnComment }}}
    {{{- end }}}
    {{{- end }}}
}

//{{{ toUpperCamelCase .EntityName }}}ListReq {{{ .FunctionName }}}列表参数
type {{{ toUpperCamelCase .EntityName }}}ListReq struct {
    {{{- range .Columns }}}
    {{{- if .IsQuery }}}
        {{{- if eq .HtmlType "datetime" }}}
            {{{ .GoField }}}Start x_null.String `form:"{{{ .TsField }}}_start" json:"{{{ .TsField }}}_start"` // 开始{{{ .ColumnComment }}}
            {{{ .GoField }}}End x_null.String `form:"{{{ .TsField }}}_end" json:"{{{ .TsField }}}_end"` // 结束{{{ .ColumnComment }}}
        {{{- else }}}
            {{{ .GoField }}} {{{ .GoNullType }}} `form:"{{{ .TsField }}}" json:"{{{ .TsField }}}"` // {{{ .ColumnComment }}}
        {{{- end }}}
    {{{- end }}}
    {{{- end }}}
}



//{{{ toUpperCamelCase .EntityName }}}AddReq {{{ .FunctionName }}}新增参数
type {{{ toUpperCamelCase .EntityName }}}AddReq struct {
    {{{- range .Columns }}}
    {{{- if .IsInsert }}}
    {{{ .GoField }}}  {{{ .GoNullType }}} `binding:"{{{ if eq .IsRequired 1 }}}required{{{ end }}}" form:"{{{ .TsField }}}" json:"{{{ .TsField }}}"`  // {{{ .ColumnComment }}}
    {{{- end }}}
    {{{- end }}}
}

//{{{ toUpperCamelCase .EntityName }}}EditReq {{{ .FunctionName }}}编辑参数
type {{{ toUpperCamelCase .EntityName }}}EditReq struct {
    {{{- range .Columns }}}
    {{{- if .IsEdit }}}
        {{{- if .IsPk }}}
    {{{ .GoField }}} {{{ .GoType }}} `form:"{{{ .TsField }}}" json:"{{{ .TsField }}}"` // {{{ .ColumnComment }}}
        {{{- else }}}
    {{{ .GoField }}}  {{{ .GoNullType }}} `form:"{{{ .TsField }}}" json:"{{{ .TsField }}}"` // {{{ .ColumnComment }}}
        {{{- end }}}
    {{{- end }}}
    {{{- end }}}
}

 

//{{{ toUpperCamelCase .EntityName }}}DelBatchReq {{{ .FunctionName }}}批量删除参数
type {{{ toUpperCamelCase .EntityName }}}DelBatchReq struct {
	Ids string `form:"ids" json:"ids"`
}

//{{{ toUpperCamelCase .EntityName }}}Resp {{{ .FunctionName }}}返回信息
type {{{ toUpperCamelCase .EntityName }}}Resp struct {
	{{{- range .Columns }}}
    {{{- if or .IsList .IsPk }}}
    {{{- if .IsPk }}}
        {{{ .GoField }}} {{{.GoType }}} `swaggertype:"{{{ .SwagType }}}" json:"{{{ .TsField }}}"`// {{{ .ColumnComment }}}
    {{{- else }}}
        {{{ .GoField }}} {{{ .GoNullType }}} `swaggertype:"{{{ .SwagType }}}" json:"{{{ .TsField }}}"`// {{{ .ColumnComment }}}
    {{{- end }}}
    {{{- end }}}
    {{{- end }}}
}
