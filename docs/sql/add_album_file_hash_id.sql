-- 移除 x_common_file_ref 表后，相册表直接关联文件哈希，用于清理时判断文件是否被相册引用
-- 执行前请确认实际表名为 x_album（与代码中 gorm model 的 TableName 一致）
ALTER TABLE `x_album`
  ADD COLUMN `file_hash_id` varchar(36) NOT NULL DEFAULT '' COMMENT '关联文件哈希ID(x_common_file_hash.id)',
  ADD INDEX `idx_file_hash_id` (`file_hash_id`);
