-- 可视化规则链菜单（挂在「物联网」目录 ...001 下，替代产品级解码/过滤/告警脚本的可视化配置入口）
-- menu_type: M=目录 C=菜单(页面) A=权限按钮；perms 与后端 PermAuth 依据请求路径生成一致（/api/admin/iot/ruleChain/xxx => admin:iot:ruleChain:xxx）
-- 超级管理员默认拥有全部菜单权限；普通角色请在 x_system_auth_perm 中分配对应菜单 id。
INSERT INTO `x_system_auth_menu` (`id`, `pid`, `menu_type`, `menu_name`, `menu_icon`, `menu_sort`, `perms`, `paths`, `component`, `selected`, `params`, `is_cache`, `is_show`, `is_disable`, `create_time`, `update_time`) VALUES
    -- 规则链
    ('019fcccc-0000-7000-9000-000000000061', '019fcccc-0000-7000-9000-000000000001', 'C', '规则链', 'el-icon-Share', 6, 'admin:iot:ruleChain:list', 'ruleChain', 'iot/rule_chain/index', '', '', 1, 1, 0, NOW(), NOW()),
    ('019fcccc-0000-7000-9000-000000000062', '019fcccc-0000-7000-9000-000000000061', 'A', '详情', '', 0, 'admin:iot:ruleChain:detail', '', '', '', '', 0, 0, 0, NOW(), NOW()),
    ('019fcccc-0000-7000-9000-000000000063', '019fcccc-0000-7000-9000-000000000061', 'A', '新增', '', 0, 'admin:iot:ruleChain:add', '', '', '', '', 0, 0, 0, NOW(), NOW()),
    ('019fcccc-0000-7000-9000-000000000064', '019fcccc-0000-7000-9000-000000000061', 'A', '编辑', '', 0, 'admin:iot:ruleChain:edit', '', '', '', '', 0, 0, 0, NOW(), NOW()),
    ('019fcccc-0000-7000-9000-000000000065', '019fcccc-0000-7000-9000-000000000061', 'A', '删除', '', 0, 'admin:iot:ruleChain:del', '', '', '', '', 0, 0, 0, NOW(), NOW());
