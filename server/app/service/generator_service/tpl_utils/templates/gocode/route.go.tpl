package admin_route

import (
	"github.com/gin-gonic/gin"
	"x_admin/app/middleware" 
	"x_admin/app/controller/admin_ctl/{{{.Domain}}}_controller"
)

/**
集成
1. 导入
- 请先提交git避免文件覆盖!!!
- 下载并解压压缩包后，直接复制server、admin文件夹到项目根目录即可

2. 后台手动添加菜单和按钮

INSERT INTO x_system_auth_menu (id,pid, menu_type, menu_name,  paths, component, is_cache, is_show, is_disable, create_time, update_time) VALUES ('{{{makeID}}}',"", 'C', '{{{ .FunctionName }}}', '{{{.Domain}}}/{{{.ModuleName}}}/index', '{{{.Domain}}}/{{{.ModuleName}}}/index', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (id,pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) SELECT '{{{makeID}}}',id, 'A', '{{{ .FunctionName }}}列表','admin:{{{ .ModuleName }}}:list', 0, 1, 0, now(), now() FROM x_system_auth_menu WHERE component='{{{.Domain}}}/{{{.ModuleName}}}/index';
INSERT INTO x_system_auth_menu (id,pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) SELECT '{{{makeID}}}',id, 'A', '{{{ .FunctionName }}}全部列表','admin:{{{ .ModuleName }}}:list_all', 0, 1, 0, now(), now() FROM x_system_auth_menu WHERE component='{{{.Domain}}}/{{{.ModuleName}}}/index';
INSERT INTO x_system_auth_menu (id,pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) SELECT '{{{makeID}}}',id, 'A', '{{{ .FunctionName }}}添加','admin:{{{ .ModuleName }}}:add', 0, 1, 0, now(), now() FROM x_system_auth_menu WHERE component='{{{.Domain}}}/{{{.ModuleName}}}/index';
INSERT INTO x_system_auth_menu (id,pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) SELECT '{{{makeID}}}',id, 'A', '{{{ .FunctionName }}}编辑','admin:{{{ .ModuleName }}}:edit', 0, 1, 0, now(), now() FROM x_system_auth_menu WHERE component='{{{.Domain}}}/{{{.ModuleName}}}/index';
INSERT INTO x_system_auth_menu (id,pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) SELECT '{{{makeID}}}',id, 'A', '{{{ .FunctionName }}}删除','admin:{{{ .ModuleName }}}:del', 0, 1, 0, now(), now() FROM x_system_auth_menu WHERE component='{{{.Domain}}}/{{{.ModuleName}}}/index';
INSERT INTO x_system_auth_menu (id,pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) SELECT '{{{makeID}}}',id, 'A', '{{{ .FunctionName }}}删除-批量','admin:{{{ .ModuleName }}}:del_batch', 0, 1, 0, now(), now() FROM x_system_auth_menu WHERE component='{{{.Domain}}}/{{{.ModuleName}}}/index';
INSERT INTO x_system_auth_menu (id,pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) SELECT '{{{makeID}}}',id, 'A', '{{{ .FunctionName }}}详情','admin:{{{ .ModuleName }}}:detail', 0, 1, 0, now(), now() FROM x_system_auth_menu WHERE component='{{{.Domain}}}/{{{.ModuleName}}}/index';
INSERT INTO x_system_auth_menu (id,pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) SELECT '{{{makeID}}}',id, 'A', '{{{ .FunctionName }}}导出excel','admin:{{{ .ModuleName }}}:export_file', 0, 1, 0, now(), now() FROM x_system_auth_menu WHERE component='{{{.Domain}}}/{{{.ModuleName}}}/index';
INSERT INTO x_system_auth_menu (id,pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) SELECT '{{{makeID}}}',id, 'A', '{{{ .FunctionName }}}导入excel','admin:{{{ .ModuleName }}}:import_file', 0, 1, 0, now(), now() FROM x_system_auth_menu WHERE component='{{{.Domain}}}/{{{.ModuleName}}}/index';
*/



func {{{ toUpperCamelCase .ModuleName }}}Route(rg *gin.RouterGroup) {
	handle := {{{.Domain}}}_controller.{{{ toUpperCamelCase .EntityName }}}Handler{}

	r := rg.Group("/", middleware.PermAuth())
	r.GET("/{{{ .ModuleName }}}/list", handle.List)
	r.GET("/{{{ .ModuleName }}}/list_all", handle.ListAll)
	r.GET("/{{{ .ModuleName }}}/detail", handle.Detail)
	
	r.POST("/{{{ .ModuleName }}}/add",middleware.RecordLog("{{{ .FunctionName }}}新增"), handle.Add)
	r.POST("/{{{ .ModuleName }}}/edit",middleware.RecordLog("{{{ .FunctionName }}}编辑"), handle.Edit)
	
	r.POST("/{{{ .ModuleName }}}/del", middleware.RecordLog("{{{ .FunctionName }}}删除"), handle.Del)
	r.POST("/{{{ .ModuleName }}}/del_batch", middleware.RecordLog("{{{ .FunctionName }}}删除-批量"), handle.DelBatch)

	r.GET("/{{{ .ModuleName }}}/export_file", middleware.RecordLog("{{{ .FunctionName }}}导出"), handle.ExportFile)
	r.POST("/{{{ .ModuleName }}}/import_file", middleware.RecordLog("{{{ .FunctionName }}}导入"),  handle.ImportFile)
}
func init() {
	routeHandlers = append(routeHandlers, {{{ toUpperCamelCase .ModuleName }}}Route)
}
