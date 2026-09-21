-- fabu 应用分发模块菜单（后端菜单驱动，插入后前端侧边栏即出现「应用分发」）
-- 说明：menu_type: M=目录(模块) C=菜单(页面) A=权限按钮
--       component 指向 src/views 下的组件路径；paths 为前端路由路径
--       超级管理员默认拥有全部菜单权限；如需普通角色可见，请在 x_system_auth_role_menu 中分配对应菜单 id。
INSERT INTO `x_system_auth_menu` (`id`, `pid`, `menu_type`, `menu_name`, `menu_icon`, `menu_sort`, `perms`, `paths`, `component`, `selected`, `params`, `is_cache`, `is_show`, `is_disable`, `create_time`, `update_time`) VALUES
    ('019fab10-c0de-7001-9001-0000000000a1', '', 'M', '应用分发', 'el-icon-TopRight', 45, '', 'fabu', '', '', '', 0, 1, 0, NOW(), NOW()),
    ('019fab10-c0de-7002-9002-0000000000a2', '019fab10-c0de-7001-9001-0000000000a1', 'C', '应用管理', 'el-icon-Mobile', 1, 'admin:fabu:app:list', 'app', 'fabu/app/index', '', '', 1, 1, 0, NOW(), NOW()),
    ('019fab10-c0de-7003-9003-0000000000a3', '019fab10-c0de-7001-9001-0000000000a1', 'C', '版本管理', 'el-icon-Download', 2, 'admin:fabu:version:list', 'version', 'fabu/version/index', '', '', 1, 1, 0, NOW(), NOW());
