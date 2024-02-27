package admin

import (
	"github.com/gin-gonic/gin"
	"x_admin/core" 
	"x_admin/middleware" 
	"x_admin/admin/{{{ .ModuleName }}}"
)

/**
集成
1. 导入
- 请先提交git避免文件覆盖!!!
- 下载并解压压缩包后，直接复制server、admin文件夹到项目根目录即可

2. 注册路由
请在 admin/entry.go 文件引入{{{ title (toCamelCase .ModuleName) }}}Route注册路由

3. 后台手动添加菜单和按钮
{{{ .ModuleName }}}:add
{{{.ModuleName }}}:edit
{{{.ModuleName }}}:del
{{{.ModuleName }}}:list
{{{.ModuleName }}}:listAll
{{{.ModuleName }}}:detail
*/


// {{{ title (toCamelCase .ModuleName) }}}Route(rg)
func {{{ title (toCamelCase .ModuleName) }}}Route(rg *gin.RouterGroup) {
	handle := {{{ .ModuleName}}}.{{{ title (toCamelCase .EntityName) }}}Handler{}

	rg = rg.Group("/", middleware.TokenAuth())
	rg.GET("/{{{ .ModuleName }}}/list", handle.List)
	rg.GET("/{{{ .ModuleName }}}/listAll", handle.ListAll)
	rg.GET("/{{{ .ModuleName }}}/detail", handle.Detail)
	rg.POST("/{{{ .ModuleName }}}/add",middleware.RecordLog("{{{ .FunctionName }}}新增"), handle.Add)
	rg.POST("/{{{ .ModuleName }}}/edit",middleware.RecordLog("{{{ .FunctionName }}}编辑"), handle.Edit)
	rg.POST("/{{{ .ModuleName }}}/del", middleware.RecordLog("{{{ .FunctionName }}}删除"),  ,handle.Del)
	rg.GET("/{{{ .ModuleName }}}/ExportFile", middleware.RecordLog("{{{ .FunctionName }}}导出"), handle.ExportFile)
	rg.POST("/{{{ .ModuleName }}}/ImportFile", handle.ImportFile)
}