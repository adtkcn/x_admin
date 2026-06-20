-- 文件关联表：记录文件在业务中的引用关系
-- 配合 x_common_file_hash 使用，用于定时清理判断文件是否被业务使用
CREATE TABLE IF NOT EXISTS `x_common_file_ref` (
  `id` char(36) NOT NULL COMMENT 'UUID v7',
  `file_hash_id` char(36) NOT NULL COMMENT '文件哈希ID(x_common_file_hash.id)',
  `biz_type` varchar(50) NOT NULL DEFAULT '' COMMENT '业务类型: user_avatar/article_cover/album等',
  `biz_id` varchar(36) NOT NULL DEFAULT '' COMMENT '业务实体ID',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_file_hash_id` (`file_hash_id`),
  KEY `idx_biz` (`biz_type`, `biz_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文件关联表(记录文件被业务引用的关系)';
