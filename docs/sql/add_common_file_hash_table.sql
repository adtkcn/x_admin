-- 创建文件哈希记录表（用于秒传功能）
CREATE TABLE IF NOT EXISTS `common_file_hash` (
  `id` varchar(36) NOT NULL COMMENT 'UUID',
  `file_md5` varchar(64) NOT NULL COMMENT '文件MD5',
  `file_size` bigint NOT NULL COMMENT '文件大小（字节）',
  `file_path` varchar(500) NOT NULL COMMENT '文件路径',
  `ext` varchar(20) DEFAULT '' COMMENT '文件扩展名',
  `created_at` bigint NOT NULL COMMENT '创建时间（Unix毫秒时间戳）',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_md5` (`file_md5`),
  KEY `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文件哈希记录表（秒传）';
