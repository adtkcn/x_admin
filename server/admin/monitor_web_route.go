package admin

import (
	"github.com/gin-gonic/gin"
	"x_admin/middleware" 
	"x_admin/admin/monitor_web"
)

/**
集成
1. 导入
- 请先提交git避免文件覆盖!!!
- 下载并解压压缩包后，直接复制server、admin文件夹到项目根目录即可

2. 注册路由
请在 admin/entry.go 文件引入MonitorWebRoute注册路由

3. 后台手动添加菜单和按钮
admin:monitor_web:add
admin:monitor_web:edit
admin:monitor_web:del
admin:monitor_web:list
admin:monitor_web:listAll
admin:monitor_web:detail
admin:monitor_web:ExportFile
admin:monitor_web:ImportFile

// 列表
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name,  paths, component, is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'C', '错误收集error', '/monitor_web/index', 'monitor_web/index', 0, 1, 0, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());
按钮
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '错误收集error添加','admin:monitor_web:add', 0, 1, 0, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '错误收集error编辑','admin:monitor_web:edit', 0, 1, 0, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '错误收集error删除','admin:monitor_web:del', 0, 1, 0, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '错误收集error列表','admin:monitor_web:list', 0, 1, 0, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '错误收集error全部列表','admin:monitor_web:listAll', 0, 1, 0, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '错误收集error详情','admin:monitor_web:detail', 0, 1, 0, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '错误收集error导出excel','admin:monitor_web:ExportFile', 0, 1, 0, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '错误收集error导入excel','admin:monitor_web:ImportFile', 0, 1, 0, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());
*/


// MonitorWebRoute(rg)
func MonitorWebRoute(rg *gin.RouterGroup) {
	handle := monitor_web.MonitorWebHandler{}

	rg = rg.Group("/", middleware.TokenAuth())
	rg.GET("/monitor_web/list", handle.List)
	rg.GET("/monitor_web/listAll", handle.ListAll)
	rg.GET("/monitor_web/detail", handle.Detail)
	rg.POST("/monitor_web/add",middleware.RecordLog("错误收集error新增"), handle.Add)
	rg.POST("/monitor_web/edit",middleware.RecordLog("错误收集error编辑"), handle.Edit)
	rg.POST("/monitor_web/del", middleware.RecordLog("错误收集error删除"), handle.Del)
	rg.GET("/monitor_web/ExportFile", middleware.RecordLog("错误收集error导出"), handle.ExportFile)
	rg.POST("/monitor_web/ImportFile", handle.ImportFile)
}