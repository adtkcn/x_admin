package admin

import (
	"x_admin/admin/monitor_error"
	"x_admin/middleware"

	"github.com/gin-gonic/gin"
)

/**
集成
1. 导入
- 请先提交git避免文件覆盖!!!
- 下载并解压压缩包后，直接复制server、admin文件夹到项目根目录即可

2. 注册路由
请在 router/admin/entry.go 文件引入 MonitorErrorRoute 注册路由

3. 后台手动添加菜单和按钮
admin:monitor_error:add
admin:monitor_error:edit
admin:monitor_error:del
admin:monitor_error:delBatch
admin:monitor_error:list
admin:monitor_error:listAll
admin:monitor_error:detail
admin:monitor_error:ExportFile
admin:monitor_error:ImportFile

// 列表-先添加菜单获取菜单id
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name,  paths, component, is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'C', '监控-错误列', 'monitor/error/index', 'monitor/error/index', 0, 1, 0, now(), now());
按钮-替换pid参数为菜单id

INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '监控-错误列添加','admin:monitor_error:add', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '监控-错误列编辑','admin:monitor_error:edit', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '监控-错误列删除','admin:monitor_error:del', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '监控-错误列删除-批量','admin:monitor_error:delBatch', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '监控-错误列列表','admin:monitor_error:list', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '监控-错误列全部列表','admin:monitor_error:listAll', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '监控-错误列详情','admin:monitor_error:detail', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '监控-错误列导出excel','admin:monitor_error:ExportFile', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '监控-错误列导入excel','admin:monitor_error:ImportFile', 0, 1, 0, now(), now());
*/

// MonitorErrorRoute(rg)
func MonitorErrorRoute(rg *gin.RouterGroup) {
	handle := monitor_error.MonitorErrorHandler{}
	rg.GET("/monitor_error/add", middleware.RecordLog("监控-错误列新增"), handle.Add)

	r := rg.Group("/", middleware.TokenAuth())
	r.GET("/monitor_error/list", handle.List)
	r.GET("/monitor_error/listAll", handle.ListAll)
	r.GET("/monitor_error/detail", handle.Detail)

	r.POST("/monitor_error/del", middleware.RecordLog("监控-错误列删除"), handle.Del)
	r.POST("/monitor_error/delBatch", middleware.RecordLog("监控-错误列删除-批量"), handle.DelBatch)

	r.GET("/monitor_error/ExportFile", middleware.RecordLog("监控-错误列导出"), handle.ExportFile)
	r.POST("/monitor_error/ImportFile", handle.ImportFile)
}
