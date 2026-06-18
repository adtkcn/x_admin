-- 用户主表（手机号在此表，非 user_auth 表）
CREATE TABLE IF NOT EXISTS `x_user` (
  `id` char(36) NOT NULL COMMENT 'UUID v7',
  `email` varchar(128) NOT NULL DEFAULT '' COMMENT '邮箱(主账号)',
  `nickname` varchar(64) NOT NULL DEFAULT '' COMMENT '昵称',
  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '密码(bcrypt)',
  `salt` varchar(32) NOT NULL DEFAULT '' COMMENT '加密盐',
  `phone` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号(可选，绑定后存入)',
  `phone_code` varchar(10) NOT NULL DEFAULT '86' COMMENT '手机号区号',
  `status` tinyint NOT NULL DEFAULT 0 COMMENT '状态 0正常 1禁用',
  `token_version` bigint NOT NULL DEFAULT 0 COMMENT 'Token版本号(踢人下线:自增使旧token失效)',
  `last_login_ip` varchar(50) NOT NULL DEFAULT '' COMMENT '最后登录IP',
  `last_login_time` datetime DEFAULT NULL COMMENT '最后登录时间',
  `is_delete` bigint NOT NULL DEFAULT 0 COMMENT '是否删除: 0=否, 1=是',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_email` (`email`),
  KEY `idx_phone` (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户主表';

-- 第三方认证绑定表（不含手机号，手机号直接存 x_user 表）
CREATE TABLE IF NOT EXISTS `x_user_auth` (
  `id` char(36) NOT NULL COMMENT 'UUID v7',
  `user_id` char(36) NOT NULL COMMENT '用户ID',
  `identity_type` varchar(20) NOT NULL COMMENT '认证类型: wechat_mini/wechat_mp/wechat_app/qq',
  `identifier` varchar(128) NOT NULL COMMENT '标识(openid/unionid/qq_openid)',
  `credential` varchar(255) NOT NULL DEFAULT '' COMMENT '凭证(加密存储，部分类型可为空)',
  `extra` varchar(512) NOT NULL DEFAULT '' COMMENT '扩展信息JSON(微信昵称、头像等)',
  `is_delete` bigint NOT NULL DEFAULT 0 COMMENT '是否删除: 0=否, 1=是',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_type_identifier` (`identity_type`, `identifier`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='第三方认证绑定表';

