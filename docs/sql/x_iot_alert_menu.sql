-- IoT 告警历史菜单（挂在既有「物联网」目录 019fcccc-0000-7000-9000-000000000001 下）
-- menu_type: M=目录 C=菜单(页面) A=权限按钮；perms 与后端 PermAuth 依据请求路径生成的标识一致
-- 数据由处理层告警状态机自动写入，管理端提供列表查询与删除两个入口。
-- 超级管理员默认拥有全部菜单权限；普通角色请在 x_system_auth_perm 中分配对应菜单 id。
INSERT INTO `x_system_auth_menu` (`id`, `pid`, `menu_type`, `menu_name`, `menu_icon`, `menu_sort`, `perms`, `paths`, `component`, `selected`, `params`, `is_cache`, `is_show`, `is_disable`, `create_time`, `update_time`) VALUES
    -- 告警历史
    ('019fcccc-0000-7000-9000-000000000051', '019fcccc-0000-7000-9000-000000000001', 'C', '告警历史', 'el-icon-Bell', 5, 'admin:iot:alert:list', 'alert', 'iot/alert/index', '', '', 1, 1, 0, NOW(), NOW()),
    ('019fcccc-0000-7000-9000-000000000052', '019fcccc-0000-7000-9000-000000000051', 'A', '删除', '', 0, 'admin:iot:alert:del', '', '', '', '', 0, 0, 0, NOW(), NOW());
