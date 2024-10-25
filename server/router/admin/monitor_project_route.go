package admin

import (
	"github.com/gin-gonic/gin"
	"x_admin/middleware" 
	"x_admin/admin/monitor_project"
)

/**
集成
1. 导入
- 请先提交git避免文件覆盖!!!
- 下载并解压压缩包后，直接复制server、admin文件夹到项目根目录即可

2. 注册路由
请在 router/admin/entry.go 文件引入 MonitorProjectRoute 注册路由

3. 后台手动添加菜单和按钮
admin:monitor_project:add
admin:monitor_project:edit
admin:monitor_project:del
admin:monitor_project:delBatch
admin:monitor_project:list
admin:monitor_project:listAll
admin:monitor_project:detail
admin:monitor_project:ExportFile
admin:monitor_project:ImportFile

// 列表-先添加菜单获取菜单id
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name,  paths, component, is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'C', '监控项目', 'monitor/project/index', 'monitor/project/index', 0, 1, 0, now(), now());
按钮-替换pid参数为菜单id

INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '监控项目添加','admin:monitor_project:add', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '监控项目编辑','admin:monitor_project:edit', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '监控项目删除','admin:monitor_project:del', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '监控项目删除-批量','admin:monitor_project:delBatch', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '监控项目列表','admin:monitor_project:list', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '监控项目全部列表','admin:monitor_project:listAll', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '监控项目详情','admin:monitor_project:detail', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '监控项目导出excel','admin:monitor_project:ExportFile', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '监控项目导入excel','admin:monitor_project:ImportFile', 0, 1, 0, now(), now());
*/


// MonitorProjectRoute(rg)
func MonitorProjectRoute(rg *gin.RouterGroup) {
	handle := monitor_project.MonitorProjectHandler{}

	r := rg.Group("/", middleware.TokenAuth())
	r.GET("/monitor_project/list", handle.List)
	r.GET("/monitor_project/listAll", handle.ListAll)
	r.GET("/monitor_project/detail", handle.Detail)
	
	r.POST("/monitor_project/add",middleware.RecordLog("监控项目新增"), handle.Add)
	r.POST("/monitor_project/edit",middleware.RecordLog("监控项目编辑"), handle.Edit)
	
	r.POST("/monitor_project/del", middleware.RecordLog("监控项目删除"), handle.Del)
	r.POST("/monitor_project/delBatch", middleware.RecordLog("监控项目删除-批量"), handle.DelBatch)

	r.GET("/monitor_project/ExportFile", middleware.RecordLog("监控项目导出"), handle.ExportFile)
	r.POST("/monitor_project/ImportFile",  handle.ImportFile)
}