-- 用户角色关联表
CREATE TABLE IF NOT EXISTS `x_system_auth_admin_role` (
    `id` char(36) NOT NULL COMMENT 'uuid',
    `admin_id` char(36) NOT NULL COMMENT '管理员ID',
    `role_id` char(36) NOT NULL COMMENT '角色ID',
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`),
    KEY `idx_admin_role` (`admin_id`, `role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户角色关联表';

-- 迁移现有数据：将原有用户的role_id迁移到关联表
INSERT INTO `x_system_auth_admin_role` (`id`, `admin_id`, `role_id`, `created_at`)
SELECT UUID(), `id`, `role_id`, NOW()
FROM `x_system_auth_admin`
WHERE `role_id` IS NOT NULL AND `role_id` != '';

-- 注意：迁移完成后，可以考虑删除 x_system_auth_admin 表中的 role_id 字段
-- 但为了向后兼容，暂时保留该字段
