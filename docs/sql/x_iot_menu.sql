-- iot 物联网元数据管理模块菜单（后端菜单驱动，插入后前端侧边栏即出现「物联网」）
-- menu_type: M=目录(模块) C=菜单(页面) A=权限按钮
-- component 指向 src/views 下的组件路径；paths 为前端路由路径
-- perms 与后端 PermAuth 依据请求路径生成的标识一致（/api/admin/iot/xxx/yyy => admin:iot:xxx:yyy）
-- 超级管理员默认拥有全部菜单权限；如需普通角色可见/可操作，请在 x_system_auth_perm 中分配对应菜单 id。
INSERT INTO `x_system_auth_menu` (`id`, `pid`, `menu_type`, `menu_name`, `menu_icon`, `menu_sort`, `perms`, `paths`, `component`, `selected`, `params`, `is_cache`, `is_show`, `is_disable`, `create_time`, `update_time`) VALUES
    -- 目录
    ('019fcccc-0000-7000-9000-000000000001', '', 'M', '物联网', 'el-icon-Cpu', 50, '', 'iot', '', '', '', 0, 1, 0, NOW(), NOW()),

    -- 产品分类
    ('019fcccc-0000-7000-9000-000000000011', '019fcccc-0000-7000-9000-000000000001', 'C', '产品分类', 'el-icon-Folder', 1, 'admin:iot:productCategory:detail', 'productCategory', 'iot/product_category/index', '', '', 1, 1, 0, NOW(), NOW()),
    ('019fcccc-0000-7000-9000-000000000012', '019fcccc-0000-7000-9000-000000000011', 'A', '新增', '', 0, 'admin:iot:productCategory:add', '', '', '', '', 0, 0, 0, NOW(), NOW()),
    ('019fcccc-0000-7000-9000-000000000013', '019fcccc-0000-7000-9000-000000000011', 'A', '编辑', '', 0, 'admin:iot:productCategory:edit', '', '', '', '', 0, 0, 0, NOW(), NOW()),
    ('019fcccc-0000-7000-9000-000000000014', '019fcccc-0000-7000-9000-000000000011', 'A', '删除', '', 0, 'admin:iot:productCategory:del', '', '', '', '', 0, 0, 0, NOW(), NOW()),
    ('019fcccc-0000-7000-9000-000000000015', '019fcccc-0000-7000-9000-000000000011', 'A', '排序', '', 0, 'admin:iot:productCategory:sort', '', '', '', '', 0, 0, 0, NOW(), NOW()),

    -- 产品
    ('019fcccc-0000-7000-9000-000000000021', '019fcccc-0000-7000-9000-000000000001', 'C', '产品管理', 'el-icon-Goods', 2, 'admin:iot:product:list', 'product', 'iot/product/index', '', '', 1, 1, 0, NOW(), NOW()),
    ('019fcccc-0000-7000-9000-000000000022', '019fcccc-0000-7000-9000-000000000021', 'A', '详情', '', 0, 'admin:iot:product:detail', '', '', '', '', 0, 0, 0, NOW(), NOW()),
    ('019fcccc-0000-7000-9000-000000000023', '019fcccc-0000-7000-9000-000000000021', 'A', '新增', '', 0, 'admin:iot:product:add', '', '', '', '', 0, 0, 0, NOW(), NOW()),
    ('019fcccc-0000-7000-9000-000000000024', '019fcccc-0000-7000-9000-000000000021', 'A', '编辑', '', 0, 'admin:iot:product:edit', '', '', '', '', 0, 0, 0, NOW(), NOW()),
    ('019fcccc-0000-7000-9000-000000000025', '019fcccc-0000-7000-9000-000000000021', 'A', '删除', '', 0, 'admin:iot:product:del', '', '', '', '', 0, 0, 0, NOW(), NOW()),

    -- 设备区域
    ('019fcccc-0000-7000-9000-000000000031', '019fcccc-0000-7000-9000-000000000001', 'C', '设备区域', 'el-icon-OfficeBuilding', 3, 'admin:iot:deviceSpace:detail', 'deviceSpace', 'iot/device_space/index', '', '', 1, 1, 0, NOW(), NOW()),
    ('019fcccc-0000-7000-9000-000000000032', '019fcccc-0000-7000-9000-000000000031', 'A', '新增', '', 0, 'admin:iot:deviceSpace:add', '', '', '', '', 0, 0, 0, NOW(), NOW()),
    ('019fcccc-0000-7000-9000-000000000033', '019fcccc-0000-7000-9000-000000000031', 'A', '编辑', '', 0, 'admin:iot:deviceSpace:edit', '', '', '', '', 0, 0, 0, NOW(), NOW()),
    ('019fcccc-0000-7000-9000-000000000034', '019fcccc-0000-7000-9000-000000000031', 'A', '删除', '', 0, 'admin:iot:deviceSpace:del', '', '', '', '', 0, 0, 0, NOW(), NOW()),
    ('019fcccc-0000-7000-9000-000000000035', '019fcccc-0000-7000-9000-000000000031', 'A', '排序', '', 0, 'admin:iot:deviceSpace:sort', '', '', '', '', 0, 0, 0, NOW(), NOW()),

    -- 设备
    ('019fcccc-0000-7000-9000-000000000041', '019fcccc-0000-7000-9000-000000000001', 'C', '设备管理', 'el-icon-Monitor', 4, 'admin:iot:device:list', 'device', 'iot/device/index', '', '', 1, 1, 0, NOW(), NOW()),
    ('019fcccc-0000-7000-9000-000000000042', '019fcccc-0000-7000-9000-000000000041', 'A', '详情', '', 0, 'admin:iot:device:detail', '', '', '', '', 0, 0, 0, NOW(), NOW()),
    ('019fcccc-0000-7000-9000-000000000043', '019fcccc-0000-7000-9000-000000000041', 'A', '新增', '', 0, 'admin:iot:device:add', '', '', '', '', 0, 0, 0, NOW(), NOW()),
    ('019fcccc-0000-7000-9000-000000000044', '019fcccc-0000-7000-9000-000000000041', 'A', '编辑', '', 0, 'admin:iot:device:edit', '', '', '', '', 0, 0, 0, NOW(), NOW()),
    ('019fcccc-0000-7000-9000-000000000045', '019fcccc-0000-7000-9000-000000000041', 'A', '删除', '', 0, 'admin:iot:device:del', '', '', '', '', 0, 0, 0, NOW(), NOW());
