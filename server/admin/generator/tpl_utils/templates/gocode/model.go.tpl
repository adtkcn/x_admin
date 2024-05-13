package model
import (
	"x_admin/core"
	"gorm.io/plugin/soft_delete"
)

//{{{ title (toCamelCase .EntityName) }}} {{{ .FunctionName }}}实体
type {{{ title (toCamelCase .EntityName) }}} struct {
	{{{- range .Columns }}}
    {{{- if not (contains $.SubTableFields .ColumnName) }}}
        {{{ if eq .GoField "is_delete" }}}
                IsDelete soft_delete.DeletedAt `gorm:"not null;default:0;softDelete:flag,DeletedAtField:DeleteTime;comment:'是否删除: 0=否, 1=是'"`
        {{{ else }}}
            {{{ if eq .GoType "core.TsTime" }}}
                {{{ title (toCamelCase .GoField) }}} core.TsTime `gorm:"{{{ if eq .GoField "create_time" }}}autoCreateTime;{{{ else }}}{{{if eq .GoField "update_time"}}}autoUpdateTime;{{{ end }}}{{{ end }}}comment:'{{{ .ColumnComment }}}'" excel:"name:{{{ .ColumnComment }}};"` // {{{ .ColumnComment }}}
            {{{ else }}}
                {{{ title (toCamelCase .GoField) }}} {{{ .GoType }}} `gorm:"{{{ if .IsPk }}}primarykey;{{{ end }}}comment:'{{{ .ColumnComment }}}'" excel:"name:{{{ .ColumnComment }}};"` // {{{ .ColumnComment }}}
            {{{- end }}}
        {{{- end }}}

    {{{- end }}}
    {{{- end }}}
}
