package admin

import (
	"github.com/gin-gonic/gin"
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
admin:{{{ .ModuleName }}}:add
admin:{{{.ModuleName }}}:edit
admin:{{{.ModuleName }}}:del
admin:{{{.ModuleName }}}:list
admin:{{{.ModuleName }}}:listAll
admin:{{{.ModuleName }}}:detail
admin:{{{.ModuleName }}}:ExportFile
admin:{{{.ModuleName }}}:ImportFile

// 列表
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name,  paths, component, is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'C', '{{{ .FunctionName }}}', '/{{{ .ModuleName }}}/index', '{{{ .ModuleName }}}/index', 0, 1, 0, now(), now());
按钮
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '{{{ .FunctionName }}}添加','admin:{{{ .ModuleName }}}:add', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '{{{ .FunctionName }}}编辑','admin:{{{ .ModuleName }}}:edit', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '{{{ .FunctionName }}}删除','admin:{{{ .ModuleName }}}:del', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '{{{ .FunctionName }}}列表','admin:{{{ .ModuleName }}}:list', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '{{{ .FunctionName }}}全部列表','admin:{{{ .ModuleName }}}:listAll', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '{{{ .FunctionName }}}详情','admin:{{{ .ModuleName }}}:detail', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '{{{ .FunctionName }}}导出excel','admin:{{{ .ModuleName }}}:ExportFile', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '{{{ .FunctionName }}}导入excel','admin:{{{ .ModuleName }}}:ImportFile', 0, 1, 0, now(), now());
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
	rg.POST("/{{{ .ModuleName }}}/del", middleware.RecordLog("{{{ .FunctionName }}}删除"), handle.Del)
	rg.GET("/{{{ .ModuleName }}}/ExportFile", middleware.RecordLog("{{{ .FunctionName }}}导出"), handle.ExportFile)
	rg.POST("/{{{ .ModuleName }}}/ImportFile", handle.ImportFile)
}