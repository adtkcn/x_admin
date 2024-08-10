package model
import (
	"x_admin/core"
	"gorm.io/plugin/soft_delete"
)

//{{{ toUpperCamelCase .EntityName }}} {{{ .FunctionName }}}实体
type {{{ toUpperCamelCase .EntityName }}} struct {
	{{{- range .Columns }}}
    {{{- if not (contains $.SubTableFields .ColumnName) }}}
        {{{- if eq .GoField "is_delete" }}}
                IsDelete soft_delete.DeletedAt `mapstructure:"{{{ .GoField }}}" gorm:"not null;default:0;softDelete:flag,DeletedAtField:DeleteTime;comment:'是否删除: 0=否, 1=是'"`
        {{{- else }}}
            {{{- if eq .GoType "core.TsTime" }}}
                {{{ toUpperCamelCase .GoField }}} core.TsTime `mapstructure:"{{{ .GoField }}}" gorm:"{{{ if eq .GoField "create_time" }}}autoCreateTime;{{{ else }}}{{{if eq .GoField "update_time"}}}autoUpdateTime;{{{ end }}}{{{ end }}}comment:'{{{ .ColumnComment }}}'" excel:"name:{{{ .ColumnComment }}};"` // {{{ .ColumnComment }}}
            {{{- else if .IsPk }}}
                {{{ toUpperCamelCase .GoField }}} {{{ .GoType }}} `mapstructure:"{{{ .GoField }}}" gorm:"primarykey;" excel:"name:{{{ .ColumnComment }}};"` // {{{ .ColumnComment }}}
            {{{- else }}}
                {{{ toUpperCamelCase .GoField }}} {{{ .GoType }}} `mapstructure:"{{{ .GoField }}}" excel:"name:{{{ .ColumnComment }}};"` // {{{ .ColumnComment }}}
            {{{- end }}}
        {{{- end }}}

    {{{- end }}}
    {{{- end }}}
}
