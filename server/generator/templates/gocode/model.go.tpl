package model

//{{{ title (toCamelCase .EntityName) }}} {{{ .FunctionName }}}实体
type {{{ title (toCamelCase .EntityName) }}} struct {
	{{{- range .Columns }}}
    {{{- if not (contains $.SubTableFields .ColumnName) }}}
    {{{ title (toCamelCase .GoField) }}} {{{ if eq .GoType "core.TsTime" }}} int64 {{{ else }}} {{{ .GoType }}} {{{ end }}} `gorm:"{{{ if .IsPk }}}primarykey;{{{ end }}}comment:'{{{ .ColumnComment }}}'"` // {{{ .ColumnComment }}}
    {{{- end }}}
    {{{- end }}}
}
