package admin

import (
	"x_admin/admin/monitor_slow"
	"x_admin/middleware"

	"github.com/gin-gonic/gin"
)

/**
集成
1. 导入
- 请先提交git避免文件覆盖!!!
- 下载并解压压缩包后，直接复制server、admin文件夹到项目根目录即可

2. 注册路由
请在 router/admin/entry.go 文件引入 MonitorSlowRoute 注册路由

3. 后台手动添加菜单和按钮
admin:monitor_slow:add
admin:monitor_slow:edit
admin:monitor_slow:del
admin:monitor_slow:delBatch
admin:monitor_slow:list
admin:monitor_slow:listAll
admin:monitor_slow:detail
admin:monitor_slow:ExportFile
admin:monitor_slow:ImportFile

// 列表-先添加菜单获取菜单id
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name,  paths, component, is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'C', '慢接口', 'monitor/slow/index', 'monitor/slow/index', 0, 1, 0, now(), now());
按钮-替换pid参数为菜单id

INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (834, 'A', '慢接口添加','admin:monitor_slow:add', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (834, 'A', '慢接口编辑','admin:monitor_slow:edit', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (834, 'A', '慢接口删除','admin:monitor_slow:del', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (834, 'A', '慢接口删除-批量','admin:monitor_slow:delBatch', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (834, 'A', '慢接口列表','admin:monitor_slow:list', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (834, 'A', '慢接口全部列表','admin:monitor_slow:listAll', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (834, 'A', '慢接口详情','admin:monitor_slow:detail', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (834, 'A', '慢接口导出excel','admin:monitor_slow:ExportFile', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (834, 'A', '慢接口导入excel','admin:monitor_slow:ImportFile', 0, 1, 0, now(), now());
*/

// MonitorSlowRoute(rg)
func MonitorSlowRoute(rg *gin.RouterGroup) {
	handle := monitor_slow.MonitorSlowHandler{}
	rg.GET("/monitor_slow/add", middleware.RecordLog("慢接口新增"), handle.Add)
	r := rg.Group("/", middleware.TokenAuth())
	r.GET("/monitor_slow/list", handle.List)
	r.GET("/monitor_slow/listAll", handle.ListAll)
	r.GET("/monitor_slow/detail", handle.Detail)

	r.POST("/monitor_slow/del", middleware.RecordLog("慢接口删除"), handle.Del)
	r.POST("/monitor_slow/delBatch", middleware.RecordLog("慢接口删除-批量"), handle.DelBatch)

	r.GET("/monitor_slow/ExportFile", middleware.RecordLog("慢接口导出"), handle.ExportFile)
	// r.POST("/monitor_slow/ImportFile", handle.ImportFile)
}
