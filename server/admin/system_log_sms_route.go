package admin

import (
	"x_admin/admin/system_log_sms"
	"x_admin/middleware"

	"github.com/gin-gonic/gin"
)

/**
集成
1. 导入
- 请先提交git避免文件覆盖!!!
- 下载并解压压缩包后，直接复制server、admin文件夹到项目根目录即可

2. 注册路由
请在 admin/entry.go 文件引入SystemLogSmsRoute注册路由

3. 后台手动添加菜单和按钮
admin:system_log_sms:add
admin:system_log_sms:edit
admin:system_log_sms:del
admin:system_log_sms:list
admin:system_log_sms:listAll
admin:system_log_sms:detail
admin:system_log_sms:ExportFile
admin:system_log_sms:ImportFile

// 列表-先添加菜单获取菜单id
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name,  paths, component, is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'C', '系统短信日志', 'system/log/sms/index', 'system_log_sms/index', 0, 1, 0, now(), now());
按钮-替换pid参数为菜单id

INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '系统短信日志添加','admin:system_log_sms:add', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '系统短信日志编辑','admin:system_log_sms:edit', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '系统短信日志删除','admin:system_log_sms:del', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '系统短信日志列表','admin:system_log_sms:list', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '系统短信日志全部列表','admin:system_log_sms:listAll', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '系统短信日志详情','admin:system_log_sms:detail', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '系统短信日志导出excel','admin:system_log_sms:ExportFile', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '系统短信日志导入excel','admin:system_log_sms:ImportFile', 0, 1, 0, now(), now());
*/

// SystemLogSmsRoute(rg)
func SystemLogSmsRoute(rg *gin.RouterGroup) {
	handle := system_log_sms.SystemLogSmsHandler{}

	r := rg.Group("/", middleware.TokenAuth())
	r.GET("/system_log_sms/list", handle.List)
	r.GET("/system_log_sms/listAll", handle.ListAll)
	r.GET("/system_log_sms/detail", handle.Detail)
	r.POST("/system_log_sms/add", middleware.RecordLog("系统短信日志新增"), handle.Add)
	r.POST("/system_log_sms/edit", middleware.RecordLog("系统短信日志编辑"), handle.Edit)
	r.POST("/system_log_sms/del", middleware.RecordLog("系统短信日志删除"), handle.Del)
	r.GET("/system_log_sms/ExportFile", middleware.RecordLog("系统短信日志导出"), handle.ExportFile)
	r.POST("/system_log_sms/ImportFile", handle.ImportFile)
}
