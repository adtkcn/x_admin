package {{{ .ModuleName }}}
import (
	"x_admin/core"

)

//{{{ toUpperCamelCase .EntityName }}}ListReq {{{ .FunctionName }}}列表参数
type {{{ toUpperCamelCase .EntityName }}}ListReq struct {
    {{{- range .Columns }}}
    {{{- if .IsQuery }}}
        {{{- if eq .HtmlType "datetime" }}}
            {{{ toUpperCamelCase .GoField }}}Start *string `form:"{{{ .GoField }}}Start"` // 开始{{{ .ColumnComment }}}
            {{{ toUpperCamelCase .GoField }}}End *string `GoField }}}End" form:"{{{ .GoField }}}End"` // 结束{{{ .ColumnComment }}}
        {{{- else }}}
            {{{ toUpperCamelCase .GoField }}} *{{{.GoType }}} ` form:"{{{ .GoField }}}"` // {{{ .ColumnComment }}}
        {{{- end }}}
    {{{- end }}}
    {{{- end }}}
}



//{{{ toUpperCamelCase .EntityName }}}AddReq {{{ .FunctionName }}}新增参数
type {{{ toUpperCamelCase .EntityName }}}AddReq struct {
    {{{- range .Columns }}}
    {{{- if .IsInsert }}}
    {{{ toUpperCamelCase .GoField }}} interface{} `mapstructure:"{{{ .GoField }}}" json:"{{{ .GoField }}}" form:"{{{ .GoField }}}"` // {{{ .ColumnComment }}}
    {{{- end }}}
    {{{- end }}}
}

//{{{ toUpperCamelCase .EntityName }}}EditReq {{{ .FunctionName }}}编辑参数
type {{{ toUpperCamelCase .EntityName }}}EditReq struct {
    {{{- range .Columns }}}
    {{{- if .IsEdit }}}
        {{{- if .IsPk }}}
        {{{ toUpperCamelCase .GoField }}} {{{ .GoType }}} `mapstructure:"{{{ .GoField }}}" json:"{{{ .GoField }}}" form:"{{{ .GoField }}}"` // {{{ .ColumnComment }}}
        {{{- else }}}
        {{{ toUpperCamelCase .GoField }}} interface{} `mapstructure:"{{{ .GoField }}}" json:"{{{ .GoField }}}" form:"{{{ .GoField }}}"` // {{{ .ColumnComment }}}
        {{{- end }}}
    {{{- end }}}
    {{{- end }}}
}

//{{{ toUpperCamelCase .EntityName }}}DetailReq {{{ .FunctionName }}}详情参数
type {{{ toUpperCamelCase .EntityName }}}DetailReq struct {
    {{{- range .Columns }}}
    {{{- if .IsPk }}}
    {{{ toUpperCamelCase .GoField }}} {{{.GoType }}} `form:"{{{ .GoField }}}"` // {{{ .ColumnComment }}}
    {{{- end }}}
    {{{- end }}}
}

//{{{ toUpperCamelCase .EntityName }}}DelReq {{{ .FunctionName }}}删除参数
type {{{ toUpperCamelCase .EntityName }}}DelReq struct {
    {{{- range .Columns }}}
    {{{- if .IsPk }}}
    {{{ toUpperCamelCase .GoField }}} {{{ .GoType }}} `form:"{{{ .GoField }}}"` // {{{ .ColumnComment }}}
    {{{- end }}}
    {{{- end }}}
}

//{{{ toUpperCamelCase .EntityName }}}Resp {{{ .FunctionName }}}返回信息
type {{{ toUpperCamelCase .EntityName }}}Resp struct {
	{{{- range .Columns }}}
    {{{- if or .IsList .IsPk }}}
    {{{- if .IsPk }}}
        {{{ toUpperCamelCase .GoField }}} {{{ .GoType }}} `mapstructure:"{{{ .GoField }}}" json:"{{{ .GoField }}}"` // {{{ .ColumnComment }}}
    {{{- else }}}
        {{{ toUpperCamelCase .GoField }}} {{{ .GoType }}} `mapstructure:"{{{ .GoField }}}" json:"{{{ .GoField }}}"` // {{{ .ColumnComment }}}
    {{{- end }}}
    {{{- end }}}
    {{{- end }}}
}
