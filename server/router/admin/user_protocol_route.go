package admin

import (
	"x_admin/admin/user_protocol"
	"x_admin/middleware"

	"github.com/gin-gonic/gin"
)

/**
集成
1. 导入
- 请先提交git避免文件覆盖!!!
- 下载并解压压缩包后，直接复制server、admin文件夹到项目根目录即可

2. 注册路由
请在 admin/entry.go 文件引入UserProtocolRoute注册路由

3. 后台手动添加菜单和按钮
admin:user_protocol:add
admin:user_protocol:edit
admin:user_protocol:del
admin:user_protocol:list
admin:user_protocol:listAll
admin:user_protocol:detail
admin:user_protocol:ExportFile
admin:user_protocol:ImportFile

// 列表-先添加菜单获取菜单id
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name,  paths, component, is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'C', '用户协议', 'user/protocol/index', 'user/protocol/index', 0, 1, 0, now(), now());
按钮-替换pid参数为菜单id

INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '用户协议添加','admin:user_protocol:add', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '用户协议编辑','admin:user_protocol:edit', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '用户协议删除','admin:user_protocol:del', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '用户协议批量删除','admin:user_protocol:delBatch', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '用户协议列表','admin:user_protocol:list', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '用户协议全部列表','admin:user_protocol:listAll', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '用户协议详情','admin:user_protocol:detail', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '用户协议导出excel','admin:user_protocol:ExportFile', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms,is_cache, is_show, is_disable, create_time, update_time) VALUES (0, 'A', '用户协议导入excel','admin:user_protocol:ImportFile', 0, 1, 0, now(), now());
*/

// UserProtocolRoute(rg)
func UserProtocolRoute(rg *gin.RouterGroup) {
	handle := user_protocol.UserProtocolHandler{}

	r := rg.Group("/", middleware.TokenAuth())
	r.GET("/user_protocol/list", handle.List)
	r.GET("/user_protocol/listAll", handle.ListAll)
	r.GET("/user_protocol/detail", handle.Detail)
	r.POST("/user_protocol/add", middleware.RecordLog("用户协议新增"), handle.Add)
	r.POST("/user_protocol/edit", middleware.RecordLog("用户协议编辑"), handle.Edit)
	r.POST("/user_protocol/del", middleware.RecordLog("用户协议删除"), handle.Del)

	r.POST("/user_protocol/delBatch", middleware.RecordLog("用户协议批量删除"), handle.DelBatch)

	r.GET("/user_protocol/ExportFile", middleware.RecordLog("用户协议导出"), handle.ExportFile)
	r.POST("/user_protocol/ImportFile", handle.ImportFile)
}
