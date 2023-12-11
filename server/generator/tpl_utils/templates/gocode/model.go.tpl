package model

//{{{ title (toCamelCase .EntityName) }}} {{{ .FunctionName }}}实体
type {{{ title (toCamelCase .EntityName) }}} struct {
	{{{- range .Columns }}}
    {{{- if not (contains $.SubTableFields .ColumnName) }}}
    {{{ if eq .GoType "core.TsTime" }}}
        {{{ title (toCamelCase .GoField) }}} int64 `gorm:"{{{ if eq .GoField "create_time" }}}autoCreateTime;{{{ else }}}{{{if eq .GoField "update_time"}}}autoUpdateTime;{{{ end }}}{{{ end }}}comment:'{{{ .ColumnComment }}}'"` // {{{ .ColumnComment }}}
    {{{ else }}}
        {{{ title (toCamelCase .GoField) }}} {{{ .GoType }}} `gorm:"{{{ if .IsPk }}}primarykey;{{{ end }}}comment:'{{{ .ColumnComment }}}'"` // {{{ .ColumnComment }}}
    {{{- end }}}
    {{{- end }}}
    {{{- end }}}
}
