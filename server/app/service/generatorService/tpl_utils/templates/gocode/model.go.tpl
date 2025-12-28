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
{{{- if ne .ID "" }}}
    {{{- if .IsUid }}}
        {{{.GoField}}}   {{{.GoNullType }}}                    `gorm:"column:{{{.ColumnName}}};type:{{{toSqlType .ColumnType .ColumnLength}}};comment:'{{{ .ColumnComment }}}'"`
        {{{.GoField}}}User system_model.SystemAuthAdminSimple `gorm:"foreignKey:{{{.GoField}}}"`
    {{{- else if eq .ColumnName "is_delete" }}}
        IsDelete soft_delete.DeletedAt `gorm:"column:{{{.ColumnName}}};type:{{{toSqlType .ColumnType .ColumnLength}}};not null;default:0;softDelete:flag,DeletedAtField:DeleteTime;comment:'是否删除: 0=否, 1=是'"`
    {{{- else if eq .GoType "time.Time" }}}
        {{{ .GoField }}} {{{.GoNullType }}} `gorm:"column:{{{.ColumnName}}};type:{{{toSqlType .ColumnType .ColumnLength}}};{{{ if eq .GoField "CreateTime" }}}autoCreateTime;{{{ else }}}{{{if eq .GoField "UpdateTime"}}}autoUpdateTime;{{{ end }}}{{{ end }}}comment:'{{{ .ColumnComment }}}'"`
    {{{- else if .IsPk }}}
        {{{ .GoField }}} {{{.GoType }}} `gorm:"column:{{{.ColumnName}}};type:{{{toSqlType .ColumnType .ColumnLength}}};primarykey;{{{ if .IsIncrement}}}autoIncrement;{{{end}}}comment:'{{{ .ColumnComment }}}'"`
    {{{- else }}}
        {{{ .GoField }}} {{{.GoNullType }}} `gorm:"column:{{{.ColumnName}}};type:{{{toSqlType .ColumnType .ColumnLength}}};comment:'{{{ .ColumnComment }}}'"`
    {{{- end }}}
{{{- end }}}
{{{- end }}}
}
{{{- if eq .PrimaryKeyGoType "string" }}}
// 自动在创建时设置 UUIDv7
func (u *{{{ toUpperCamelCase .EntityName }}}) BeforeCreate(tx *gorm.DB) error {
    if u.{{{toUpperCamelCase .PrimaryKey }}} == "" {
        id, err := uuid.NewV7()
        if err != nil {
            return err
        }
        u.{{{toUpperCamelCase .PrimaryKey }}} = id.String()
    }
	return nil
}
{{{- end }}}