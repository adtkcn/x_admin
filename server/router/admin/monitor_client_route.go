package admin

import (
	"x_admin/admin/monitor_client"
	"x_admin/middleware"

	"github.com/gin-gonic/gin"
)

/**
集成
1. 导入
- 请先提交git避免文件覆盖!!!
- 下载并解压压缩包后，直接复制server、admin文件夹到项目根目录即可

2. 注册路由
请在 router/admin/entry.go 文件引入 MonitorClientRoute 注册路由

3. 后台手动添加菜单和按钮
admin:monitor_client:add
admin:monitor_client:edit
admin:monitor_client:del
admin:monitor_client:delBatch
admin:monitor_client:list
admin:monitor_client:listAll
admin:monitor_client:detail
admin:monitor_client:ExportFile
admin:monitor_client:ImportFile

// 列表-先添加菜单获取菜单id
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name,  paths, component, is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'C', '监控-客户端信息', 'monitor/client/index', 'monitor/client/index', 0, 1, 0, now(), now());
按钮-替换pid参数为菜单id

INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '监控-客户端信息添加','admin:monitor_client:add', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '监控-客户端信息编辑','admin:monitor_client:edit', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '监控-客户端信息删除','admin:monitor_client:del', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '监控-客户端信息删除-批量','admin:monitor_client:delBatch', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '监控-客户端信息列表','admin:monitor_client:list', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '监控-客户端信息全部列表','admin:monitor_client:listAll', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '监控-客户端信息详情','admin:monitor_client:detail', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '监控-客户端信息导出excel','admin:monitor_client:ExportFile', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '监控-客户端信息导入excel','admin:monitor_client:ImportFile', 0, 1, 0, now(), now());
*/

// MonitorClientRoute(rg)
func MonitorClientRoute(rg *gin.RouterGroup) {
	handle := monitor_client.MonitorClientHandler{}
	rg.GET("/monitor_client/add", middleware.RecordLog("监控-客户端信息新增"), handle.Add)

	r := rg.Group("/", middleware.TokenAuth())
	r.GET("/monitor_client/list", handle.List)
	r.GET("/monitor_client/listAll", handle.ListAll)
	r.GET("/monitor_client/detail", handle.Detail)
	r.GET("/monitor_client/errorUsers", handle.ErrorUsers)

	// r.POST("/monitor_client/edit",middleware.RecordLog("监控-客户端信息编辑"), handle.Edit)

	r.POST("/monitor_client/del", middleware.RecordLog("监控-客户端信息删除"), handle.Del)
	r.POST("/monitor_client/delBatch", middleware.RecordLog("监控-客户端信息删除-批量"), handle.DelBatch)

	r.GET("/monitor_client/ExportFile", middleware.RecordLog("监控-客户端信息导出"), handle.ExportFile)
	r.POST("/monitor_client/ImportFile", handle.ImportFile)
}
