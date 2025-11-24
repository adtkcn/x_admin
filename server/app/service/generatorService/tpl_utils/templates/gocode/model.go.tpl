package model
import (
    "github.com/google/uuid"
    "gorm.io/gorm"
	"x_admin/core"
	"gorm.io/plugin/soft_delete"
)

//{{{ toUpperCamelCase .EntityName }}} {{{ .FunctionName }}}实体
type {{{ toUpperCamelCase .EntityName }}} struct {
{{{- range .Columns }}}
{{{- if not (contains $.SubTableFields .ColumnName) }}}
    {{{- if eq .GoField "is_delete" }}}
        IsDelete soft_delete.DeletedAt `gorm:"column:{{{.ColumnName}}};not null;default:0;softDelete:flag,DeletedAtField:DeleteTime;comment:'是否删除: 0=否, 1=是'"`
    {{{- else }}}
    {{{- if eq .GoType "core.NullTime" }}}
        {{{ toUpperCamelCase .GoField }}} core.NullTime `gorm:"column:{{{.ColumnName}}};{{{ if eq .GoField "create_time" }}}autoCreateTime;{{{ else }}}{{{if eq .GoField "update_time"}}}autoUpdateTime;{{{ end }}}{{{ end }}}comment:'{{{ .ColumnComment }}}'"`
    {{{- else if .IsPk }}}
        {{{ toUpperCamelCase .GoField }}} {{{.GoType }}} `gorm:"column:{{{.ColumnName}}};primarykey;{{{ if .IsIncrement}}}autoIncrement;{{{end}}}comment:'{{{ .ColumnComment }}}'"`
    {{{- else }}}
        {{{ toUpperCamelCase .GoField }}} {{{goWithRespType .GoType }}} `gorm:"column:{{{.ColumnName}}};comment:'{{{ .ColumnComment }}}'"`
    {{{- end }}}
    {{{- end }}}

{{{- end }}}
{{{- end }}}
}
{{{- if eq .PrimaryKeyGoType "string" }}}
// 自动在创建时设置 UUIDv7
func (u *{{{ toUpperCamelCase .EntityName }}}) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	u.Id = id.String()
	return nil
}
{{{- end }}}