package adminRoute

import (
	"x_admin/app/controller/admin_ctl"
	"x_admin/middleware"

	"github.com/gin-gonic/gin"
)

/**
集成
1. 导入
- 请先提交git避免文件覆盖!!!
- 下载并解压压缩包后，直接复制server、admin文件夹到项目根目录即可

2. 注册路由(通过init函数收集路由，Autoload自动注册)

3. 后台手动添加菜单和按钮

INSERT INTO x_system_auth_menu (pid, menu_type, menu_name,  paths, component, is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'C', '定时任务', 'system/corn/index', 'system/corn/index', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) SELECT id, 'A', '定时任务列表','admin:system_corn:list', 0, 1, 0, now(), now() FROM x_system_auth_menu WHERE component='system/corn/index';
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) SELECT id, 'A', '定时任务全部列表','admin:system_corn:listAll', 0, 1, 0, now(), now() FROM x_system_auth_menu WHERE component='system/corn/index';
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) SELECT id, 'A', '定时任务添加','admin:system_corn:add', 0, 1, 0, now(), now() FROM x_system_auth_menu WHERE component='system/corn/index';
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) SELECT id, 'A', '定时任务编辑','admin:system_corn:edit', 0, 1, 0, now(), now() FROM x_system_auth_menu WHERE component='system/corn/index';
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) SELECT id, 'A', '定时任务删除','admin:system_corn:del', 0, 1, 0, now(), now() FROM x_system_auth_menu WHERE component='system/corn/index';
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) SELECT id, 'A', '定时任务删除-批量','admin:system_corn:delBatch', 0, 1, 0, now(), now() FROM x_system_auth_menu WHERE component='system/corn/index';
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) SELECT id, 'A', '定时任务详情','admin:system_corn:detail', 0, 1, 0, now(), now() FROM x_system_auth_menu WHERE component='system/corn/index';
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) SELECT id, 'A', '定时任务导出excel','admin:system_corn:ExportFile', 0, 1, 0, now(), now() FROM x_system_auth_menu WHERE component='system/corn/index';
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) SELECT id, 'A', '定时任务导入excel','admin:system_corn:ImportFile', 0, 1, 0, now(), now() FROM x_system_auth_menu WHERE component='system/corn/index';
*/

// SystemCornRoute(rg)
func SystemCornRoute(rg *gin.RouterGroup) {
	handle := admin_ctl.SystemCornHandler{}

	r := rg.Group("/", middleware.TokenAuth())
	r.GET("/system_corn/list", handle.List)
	r.GET("/system_corn/listAll", handle.ListAll)
	r.GET("/system_corn/detail", handle.Detail)

	r.POST("/system_corn/add", middleware.RecordLog("定时任务新增"), handle.Add)
	r.POST("/system_corn/edit", middleware.RecordLog("定时任务编辑"), handle.Edit)

	r.POST("/system_corn/del", middleware.RecordLog("定时任务删除"), handle.Del)
	r.POST("/system_corn/delBatch", middleware.RecordLog("定时任务删除-批量"), handle.DelBatch)

	r.GET("/system_corn/exportFile", middleware.RecordLog("定时任务导出"), handle.ExportFile)
	r.POST("/system_corn/importFile", handle.ImportFile)

	r.GET("/system_corn/getTaskList", handle.GetTaskList)
}
func init() {
	routeHandlers = append(routeHandlers, SystemCornRoute)
}
