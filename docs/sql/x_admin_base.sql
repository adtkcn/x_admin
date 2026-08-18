-- --------------------------------------------------------
-- 主机:                           127.0.0.1
-- 服务器版本:                        5.7.26-log - MySQL Community Server (GPL)
-- 服务器操作系统:                      Win64
-- HeidiSQL 版本:                  12.16.0.7229
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

-- 导出  表 x_admin_2.x_album 结构
CREATE TABLE IF NOT EXISTS `x_album` (
  `id` char(36) NOT NULL COMMENT '主键ID',
  `cid` char(36) NOT NULL DEFAULT '' COMMENT '类目ID',
  `admin_id` char(36) NOT NULL COMMENT '管理员ID',
  `uid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '文件名称',
  `uri` varchar(200) NOT NULL COMMENT '文件路径',
  `ext` varchar(10) NOT NULL DEFAULT '' COMMENT '文件扩展',
  `hash` varchar(32) NOT NULL COMMENT '文件md5_size',
  `size` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '文件大小',
  `is_delete` int(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除: 0=否, 1=是',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
  `file_hash_id` varchar(36) NOT NULL DEFAULT '' COMMENT '关联文件哈希ID(x_common_file_hash.id)',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_cid` (`cid`) USING BTREE,
  KEY `admin_id` (`admin_id`),
  KEY `hash` (`hash`),
  KEY `idx_file_hash_id` (`file_hash_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='相册管理表';

-- 数据导出被取消选择。

-- 导出  表 x_admin_2.x_album_cate 结构
CREATE TABLE IF NOT EXISTS `x_album_cate` (
  `id` char(36) NOT NULL COMMENT '主键ID',
  `pid` char(36) NOT NULL DEFAULT '' COMMENT '父级ID',
  `admin_id` char(36) NOT NULL COMMENT '管理员id',
  `name` varchar(32) NOT NULL DEFAULT '' COMMENT '分类名称',
  `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除: [0=否, 1=是]',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `pid` (`pid`),
  KEY `admin_id` (`admin_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='相册分类表';

-- 数据导出被取消选择。

-- 导出  表 x_admin_2.x_common_file_hash 结构
CREATE TABLE IF NOT EXISTS `x_common_file_hash` (
  `id` varchar(36) NOT NULL COMMENT 'UUID',
  `file_md5` varchar(64) NOT NULL COMMENT '文件MD5',
  `file_size` bigint(20) NOT NULL COMMENT '文件大小（字节）',
  `file_path` varchar(500) NOT NULL COMMENT '文件路径',
  `ext` varchar(20) DEFAULT '' COMMENT '文件扩展名',
  `create_time` datetime NOT NULL COMMENT '创建时间（Unix毫秒时间戳）',
  `last_access_time` datetime DEFAULT NULL COMMENT '最后访问时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_md5` (`file_md5`),
  KEY `idx_created_at` (`create_time`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文件哈希记录表（秒传）';

-- 数据导出被取消选择。

-- 数据导出被取消选择。

-- 导出  表 x_admin_2.x_dict_data 结构
CREATE TABLE IF NOT EXISTS `x_dict_data` (
  `id` char(36) NOT NULL COMMENT '主键',
  `type_id` char(36) NOT NULL COMMENT '类型',
  `name` varchar(100) NOT NULL COMMENT '键名',
  `value` varchar(200) NOT NULL COMMENT '数值',
  `color` varchar(20) DEFAULT NULL COMMENT '颜色',
  `remark` varchar(200) NOT NULL COMMENT '备注',
  `sort` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `status` tinyint(1) NOT NULL COMMENT '状态: 0=停用, 1=正常',
  `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除: 0=否, 1=是',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `type_id` (`type_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='字典数据表';

-- 数据导出被取消选择。

-- 导出  表 x_admin_2.x_dict_type 结构
CREATE TABLE IF NOT EXISTS `x_dict_type` (
  `id` char(36) NOT NULL COMMENT '主键',
  `dict_name` varchar(100) NOT NULL DEFAULT '' COMMENT '字典名称',
  `dict_type` varchar(100) NOT NULL DEFAULT '' COMMENT '字典类型',
  `dict_remark` varchar(200) NOT NULL DEFAULT '' COMMENT '字典备注',
  `dict_status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '字典状态: 0=停用, 1=正常',
  `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除: 0=否, 1=是',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='字典类型表';

-- 数据导出被取消选择。

-- 导出  表 x_admin_2.x_flow_apply 结构
CREATE TABLE IF NOT EXISTS `x_flow_apply` (
  `id` char(36) NOT NULL,
  `template_id` char(36) NOT NULL COMMENT '模板',
  `apply_user_id` char(36) NOT NULL COMMENT '申请人id',
  `apply_user_nickname` varchar(32) NOT NULL DEFAULT '0' COMMENT '申请人昵称',
  `flow_name` varchar(255) NOT NULL COMMENT '流程名称',
  `flow_group` tinyint(2) NOT NULL DEFAULT '0' COMMENT '流程分类',
  `flow_remark` varchar(255) DEFAULT NULL COMMENT '流程描述',
  `flow_form_data` longtext COMMENT '表单配置',
  `flow_process_data` longtext COMMENT '流程配置',
  `flow_process_data_list` longtext COMMENT '流程配置list',
  `form_value` mediumtext NOT NULL COMMENT '表单值',
  `status` tinyint(2) unsigned DEFAULT '1' COMMENT '状态：1待提交，2审批中，3审批完成，4审批失败',
  `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除: 0=否, 1=是',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `apply_user_id` (`apply_user_id`),
  KEY `is_delete` (`is_delete`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='申请流程';

-- 数据导出被取消选择。

-- 导出  表 x_admin_2.x_flow_history 结构
CREATE TABLE IF NOT EXISTS `x_flow_history` (
  `id` char(36) NOT NULL COMMENT '历史id',
  `apply_id` char(36) NOT NULL DEFAULT '' COMMENT '申请id',
  `template_id` char(36) DEFAULT NULL COMMENT '模板id',
  `apply_user_id` char(36) NOT NULL COMMENT '申请人id',
  `apply_user_nickname` varchar(32) NOT NULL DEFAULT '0' COMMENT '申请人昵称',
  `approver_id` char(36) NOT NULL COMMENT '审批人id',
  `approver_nickname` varchar(32) NOT NULL DEFAULT '' COMMENT '审批用户昵称',
  `node_id` varchar(50) NOT NULL COMMENT '节点',
  `node_type` varchar(30) NOT NULL COMMENT '节点类型',
  `node_label` varchar(50) DEFAULT NULL COMMENT '节点名称',
  `form_value` mediumtext NOT NULL COMMENT '表单值',
  `pass_status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '通过状态：1待处理，2通过，3拒绝',
  `pass_remark` varchar(200) NOT NULL COMMENT '通过备注',
  `is_show` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否在已完成显示: 0=否, 1=是',
  `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除: 0=否, 1=是',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `apply_id` (`apply_id`) USING BTREE,
  KEY `approver_id` (`approver_id`) USING BTREE,
  KEY `apply_user_id` (`apply_user_id`),
  KEY `is_delete` (`is_delete`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='流程历史';

-- 数据导出被取消选择。

-- 导出  表 x_admin_2.x_flow_template 结构
CREATE TABLE IF NOT EXISTS `x_flow_template` (
  `id` char(36) NOT NULL,
  `flow_name` varchar(255) NOT NULL COMMENT '流程名称',
  `flow_group` tinyint(2) NOT NULL DEFAULT '0' COMMENT '流程分类',
  `flow_remark` varchar(255) DEFAULT NULL COMMENT '流程描述',
  `flow_form_data` longtext COMMENT '表单配置',
  `flow_process_data` longtext COMMENT '流程配置',
  `flow_process_data_list` longtext COMMENT '流程配置list',
  `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除: 0=否, 1=是',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `is_delete` (`is_delete`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='流程模板';

-- 数据导出被取消选择。

-- 导出  表 x_admin_2.x_gen_table 结构
CREATE TABLE IF NOT EXISTS `x_gen_table` (
  `id` char(36) NOT NULL COMMENT '主键',
  `table_name` varchar(200) NOT NULL DEFAULT '' COMMENT '表名称',
  `table_comment` varchar(200) NOT NULL DEFAULT '' COMMENT '表描述',
  `sub_table_name` varchar(200) NOT NULL DEFAULT '' COMMENT '关联表名称',
  `sub_table_fk` varchar(200) NOT NULL DEFAULT '' COMMENT '关联表外键',
  `author_name` varchar(100) NOT NULL DEFAULT '' COMMENT '作者的名称',
  `entity_name` varchar(100) NOT NULL DEFAULT '' COMMENT '实体的名称',
  `module_name` varchar(60) NOT NULL DEFAULT '' COMMENT '生成模块名',
  `function_name` varchar(60) NOT NULL DEFAULT '' COMMENT '生成功能名',
  `tree_primary` varchar(60) NOT NULL DEFAULT '' COMMENT '树主键字段',
  `tree_parent` varchar(60) NOT NULL DEFAULT '' COMMENT '树父级字段',
  `tree_name` varchar(60) NOT NULL DEFAULT '' COMMENT '树显示字段',
  `gen_tpl` varchar(20) NOT NULL DEFAULT 'crud' COMMENT '生成模板方式: [crud=单表, tree=树表]',
  `remarks` varchar(200) NOT NULL DEFAULT '' COMMENT '备注信息',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='代码生成业务表';

-- 数据导出被取消选择。

-- 导出  表 x_admin_2.x_gen_table_column 结构
CREATE TABLE IF NOT EXISTS `x_gen_table_column` (
  `id` char(36) NOT NULL COMMENT '列主键',
  `table_id` char(36) NOT NULL COMMENT '表外键',
  `column_name` varchar(200) NOT NULL DEFAULT '' COMMENT 'sql列名称',
  `column_comment` varchar(200) NOT NULL DEFAULT '' COMMENT 'sql列描述',
  `column_length` varchar(5) DEFAULT '0' COMMENT 'sql列长度',
  `column_type` varchar(100) NOT NULL DEFAULT '' COMMENT 'sql列类型 ',
  `go_type` varchar(50) NOT NULL DEFAULT '0' COMMENT 'go类型',
  `go_field` varchar(100) NOT NULL DEFAULT '' COMMENT 'go字段',
  `is_pk` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否主键: [1=是, 0=否]',
  `is_increment` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否自增: [1=是, 0=否]',
  `is_required` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否必填: [1=是, 0=否]',
  `is_insert` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否插入字段: [1=是, 0=否]',
  `is_edit` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否编辑字段: [1=是, 0=否]',
  `is_list` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否列表字段: [1=是, 0=否]',
  `is_query` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否查询字段: [1=是, 0=否]',
  `query_type` varchar(30) NOT NULL DEFAULT 'EQ' COMMENT '查询方式: [等于、不等于、大于、小于、范围]',
  `html_type` varchar(30) NOT NULL DEFAULT '' COMMENT '显示类型: [文本框、文本域、下拉框、复选框、单选框、日期控件]',
  `dict_type` varchar(200) NOT NULL DEFAULT '' COMMENT '字典类型',
  `list_all_api` varchar(200) NOT NULL DEFAULT '' COMMENT '下拉框列表数据来源api',
  `sort` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '排序编号',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `table_id` (`table_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='代码生成字段表';

-- 数据导出被取消选择。

-- 导出  表 x_admin_2.x_monitor_client 结构
CREATE TABLE IF NOT EXISTS `x_monitor_client` (
  `id` char(36) NOT NULL COMMENT 'uuid',
  `project_key` varchar(128) NOT NULL COMMENT '项目key',
  `client_id` varchar(128) NOT NULL COMMENT 'sdk生成的客户端id',
  `os` varchar(30) DEFAULT NULL COMMENT '系统',
  `browser` varchar(30) DEFAULT NULL COMMENT '浏览器',
  `ua` varchar(128) DEFAULT NULL COMMENT 'ua记录',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `client_id` (`client_id`),
  KEY `project_key` (`project_key`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='监控-客户端信息';

-- 数据导出被取消选择。

-- 导出  表 x_admin_2.x_monitor_error 结构
CREATE TABLE IF NOT EXISTS `x_monitor_error` (
  `id` char(36) NOT NULL COMMENT '错误id',
  `project_key` varchar(128) NOT NULL COMMENT '项目key',
  `md5` varchar(32) DEFAULT NULL COMMENT 'md5',
  `event_type` varchar(20) DEFAULT NULL COMMENT '事件类型',
  `path` varchar(1000) DEFAULT NULL COMMENT 'URL地址',
  `message` text COMMENT '错误消息',
  `stack` text COMMENT '错误堆栈',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `project_key` (`project_key`,`md5`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='监控-错误列表';

-- 数据导出被取消选择。

-- 导出  表 x_admin_2.x_monitor_error_list 结构
CREATE TABLE IF NOT EXISTS `x_monitor_error_list` (
  `id` char(36) NOT NULL COMMENT 'id',
  `error_id` char(36) CHARACTER SET utf8 NOT NULL COMMENT '错误表id',
  `client_id` char(36) CHARACTER SET utf8 NOT NULL COMMENT '客户端表id',
  `user_id` char(36) CHARACTER SET utf8 DEFAULT NULL COMMENT '业务中用户id',
  `width` smallint(10) unsigned DEFAULT '0' COMMENT '屏幕',
  `height` smallint(10) unsigned DEFAULT '0' COMMENT '屏幕高度',
  `country` varchar(50) CHARACTER SET utf8 DEFAULT NULL COMMENT '国家',
  `province` varchar(50) CHARACTER SET utf8 DEFAULT NULL COMMENT '省份',
  `city` varchar(50) CHARACTER SET utf8 DEFAULT NULL COMMENT '城市',
  `operator` varchar(50) CHARACTER SET utf8 DEFAULT NULL COMMENT '电信运营商',
  `ip` varchar(39) CHARACTER SET utf8 DEFAULT NULL COMMENT 'ip',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `client_id` (`client_id`) USING BTREE,
  KEY `eid` (`error_id`) USING BTREE,
  KEY `create_time` (`create_time`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COMMENT='错误对应的用户记录';

-- 数据导出被取消选择。

-- 导出  表 x_admin_2.x_monitor_project 结构
CREATE TABLE IF NOT EXISTS `x_monitor_project` (
  `id` char(36) NOT NULL DEFAULT '' COMMENT '项目id',
  `project_key` char(36) NOT NULL DEFAULT '' COMMENT 'project_key',
  `project_name` varchar(50) NOT NULL COMMENT '项目名称',
  `project_type` varchar(20) DEFAULT '' COMMENT '项目类型go java web node php 等',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否启用: 0=否, 1=是',
  `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除: 0=否, 1=是',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `project_key` (`project_key`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='监控项目';

-- 数据导出被取消选择。

-- 导出  表 x_admin_2.x_system_auth_admin 结构
CREATE TABLE IF NOT EXISTS `x_system_auth_admin` (
  `id` char(36) NOT NULL COMMENT '主键',
  `email` varchar(200) NOT NULL DEFAULT '' COMMENT '邮箱(账号)',
  `nickname` varchar(32) NOT NULL DEFAULT '' COMMENT '用户昵称',
  `password` varchar(200) NOT NULL DEFAULT '' COMMENT '用户密码',
  `dept_id` char(36) NOT NULL COMMENT '部门ID',
  `post_id` char(36) NOT NULL COMMENT '岗位ID',
  `avatar` varchar(200) NOT NULL DEFAULT '' COMMENT '用户头像',
  `salt` varchar(20) NOT NULL DEFAULT '' COMMENT '加密盐巴',
  `sort` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '排序编号',
  `is_multipoint` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '多端登录: 0=否, 1=是',
  `token_version` bigint(20) NOT NULL DEFAULT '0' COMMENT 'Token版本号(踢人下线:自增使旧token失效)',
  `is_disable` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否禁用: 0=否, 1=是',
  `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除: 0=否, 1=是',
  `last_login_ip` varchar(39) NOT NULL DEFAULT '' COMMENT '最后登录IP',
  `last_login_time` datetime DEFAULT NULL COMMENT '最后登录',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `is_delete` (`is_delete`),
  KEY `idx_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='系统管理成员表';

-- 数据导出被取消选择。

-- 导出  表 x_admin_2.x_system_auth_admin_role 结构
CREATE TABLE IF NOT EXISTS `x_system_auth_admin_role` (
  `id` char(36) NOT NULL COMMENT 'uuid',
  `admin_id` char(36) NOT NULL COMMENT '管理员ID',
  `role_id` char(36) NOT NULL COMMENT '角色ID',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_admin_role` (`admin_id`,`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户角色关联表';

-- 数据导出被取消选择。

-- 导出  表 x_admin_2.x_system_auth_dept 结构
CREATE TABLE IF NOT EXISTS `x_system_auth_dept` (
  `id` char(36) NOT NULL COMMENT '主键',
  `pid` char(36) NOT NULL COMMENT '上级主键',
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '部门名称',
  `duty_id` char(36) DEFAULT NULL COMMENT '负责人id',
  `duty` varchar(32) NOT NULL DEFAULT '' COMMENT '负责人名',
  `mobile` varchar(30) DEFAULT '' COMMENT '联系电话',
  `sort` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '排序编号',
  `is_stop` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否禁用: 0=否, 1=是',
  `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除: 0=否, 1=是',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='系统部门管理表';

-- 数据导出被取消选择。

-- 导出  表 x_admin_2.x_system_auth_menu 结构
CREATE TABLE IF NOT EXISTS `x_system_auth_menu` (
  `id` char(36) NOT NULL COMMENT '主键',
  `pid` char(36) NOT NULL COMMENT '上级菜单',
  `menu_type` char(2) NOT NULL DEFAULT '' COMMENT '权限类型: M=目录，C=菜单，A=按钮',
  `menu_name` varchar(100) NOT NULL DEFAULT '' COMMENT '菜单名称',
  `menu_icon` varchar(100) NOT NULL DEFAULT '' COMMENT '菜单图标',
  `menu_sort` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '菜单排序',
  `perms` varchar(10000) NOT NULL DEFAULT '' COMMENT '权限标识',
  `paths` varchar(100) NOT NULL DEFAULT '' COMMENT '路由地址',
  `component` varchar(200) NOT NULL DEFAULT '' COMMENT '前端组件',
  `selected` varchar(200) NOT NULL DEFAULT '' COMMENT '选中路径',
  `params` varchar(200) NOT NULL DEFAULT '' COMMENT '路由参数',
  `is_cache` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否缓存: 0=否, 1=是',
  `is_show` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '是否显示: 0=否, 1=是',
  `is_disable` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否禁用: 0=否, 1=是',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='系统菜单管理表';

-- 数据导出被取消选择。

-- 导出  表 x_admin_2.x_system_auth_perm 结构
CREATE TABLE IF NOT EXISTS `x_system_auth_perm` (
  `id` char(36) NOT NULL COMMENT '主键',
  `role_id` char(36) NOT NULL COMMENT '角色ID',
  `menu_id` char(36) NOT NULL COMMENT '菜单ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='系统角色菜单表';

-- 数据导出被取消选择。

-- 导出  表 x_admin_2.x_system_auth_post 结构
CREATE TABLE IF NOT EXISTS `x_system_auth_post` (
  `id` char(36) NOT NULL COMMENT '主键',
  `code` varchar(30) NOT NULL DEFAULT '' COMMENT '岗位编码',
  `name` varchar(30) NOT NULL DEFAULT '' COMMENT '岗位名称',
  `remarks` varchar(250) NOT NULL DEFAULT '' COMMENT '岗位备注',
  `sort` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '岗位排序',
  `is_stop` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否停用: 0=否, 1=是',
  `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除: 0=否, 1=是',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `is_delete` (`is_delete`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='系统岗位管理表';

-- 数据导出被取消选择。

-- 导出  表 x_admin_2.x_system_auth_role 结构
CREATE TABLE IF NOT EXISTS `x_system_auth_role` (
  `id` char(36) NOT NULL COMMENT '主键',
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '角色名称',
  `remark` varchar(200) NOT NULL DEFAULT '' COMMENT '备注信息',
  `sort` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '角色排序',
  `is_disable` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否禁用: 0=否, 1=是',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='系统角色管理表';

-- 数据导出被取消选择。

-- 导出  表 x_admin_2.x_system_config 结构
CREATE TABLE IF NOT EXISTS `x_system_config` (
  `id` char(36) NOT NULL COMMENT '主键',
  `type` varchar(30) DEFAULT '' COMMENT '类型',
  `name` varchar(60) NOT NULL DEFAULT '' COMMENT '键',
  `value` text COMMENT '值',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='系统全局配置表';

-- 数据导出被取消选择。

-- 导出  表 x_admin_2.x_system_corn 结构
CREATE TABLE IF NOT EXISTS `x_system_corn` (
  `id` char(36) NOT NULL COMMENT 'taskid',
  `task_name` char(100) NOT NULL COMMENT '任务名称',
  `task_code` char(100) NOT NULL COMMENT '任务编码',
  `corn_expr` char(100) NOT NULL COMMENT 'corn表达式',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态',
  `created_by` char(36) NOT NULL COMMENT '创建人',
  `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除: 0=否, 1=是',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='定时任务';

-- 数据导出被取消选择。

-- 导出  表 x_admin_2.x_system_log_login 结构
CREATE TABLE IF NOT EXISTS `x_system_log_login` (
  `id` char(36) NOT NULL COMMENT '注解',
  `admin_id` char(36) NOT NULL COMMENT '管理员ID',
  `email` varchar(200) NOT NULL DEFAULT '' COMMENT '登录邮箱',
  `ip` varchar(39) NOT NULL COMMENT '登录地址',
  `os` varchar(100) NOT NULL DEFAULT '' COMMENT '操作系统',
  `browser` varchar(100) DEFAULT '' COMMENT '浏览器',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '操作状态: 1=成功, 2=失败',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='系统登录日志表';

-- 数据导出被取消选择。

-- 导出  表 x_admin_2.x_system_log_operate 结构
CREATE TABLE IF NOT EXISTS `x_system_log_operate` (
  `id` char(36) NOT NULL COMMENT '主键',
  `admin_id` char(36) NOT NULL COMMENT '操作人ID',
  `type` varchar(30) NOT NULL DEFAULT '' COMMENT '请求类型: GET/POST/PUT',
  `title` varchar(30) DEFAULT '' COMMENT '操作标题',
  `ip` varchar(39) NOT NULL DEFAULT '' COMMENT '请求IP',
  `url` varchar(200) NOT NULL DEFAULT '' COMMENT '请求接口',
  `method` varchar(200) NOT NULL DEFAULT '' COMMENT '请求方法',
  `args` text COMMENT '请求参数',
  `error` text COMMENT '错误信息',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '执行状态: 1=成功, 2=失败',
  `start_time` datetime DEFAULT NULL COMMENT '开始时间',
  `end_time` datetime DEFAULT NULL COMMENT '结束时间',
  `task_time` int(11) DEFAULT NULL COMMENT '执行耗时',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='系统操作日志表';

-- 数据导出被取消选择。

-- 导出  表 x_admin_2.x_system_log_sms 结构
CREATE TABLE IF NOT EXISTS `x_system_log_sms` (
  `id` char(36) NOT NULL COMMENT 'id',
  `scene` int(11) unsigned DEFAULT '0' COMMENT '场景编号',
  `mobile` varchar(11) DEFAULT '' COMMENT '手机号码',
  `content` varchar(255) DEFAULT '' COMMENT '发送内容',
  `status` tinyint(1) unsigned DEFAULT '0' COMMENT '发送状态：[0=发送中, 1=发送成功, 2=发送失败]',
  `results` text COMMENT '短信结果',
  `send_time` datetime DEFAULT NULL COMMENT '发送时间',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='系统短信日志表';

-- 数据导出被取消选择。

-- 导出  表 x_admin_2.x_system_notice 结构
CREATE TABLE IF NOT EXISTS `x_system_notice` (
  `id` char(36) NOT NULL COMMENT 'UUIDv7',
  `type` varchar(32) NOT NULL DEFAULT '' COMMENT '通知类型success|danger|primary|info|warning',
  `title` varchar(200) NOT NULL DEFAULT '' COMMENT '通知标题',
  `content` text COMMENT '通知正文',
  `receiver_id` char(36) NOT NULL COMMENT '接收人ID',
  `sender_id` char(36) NOT NULL DEFAULT '' COMMENT '发送人ID',
  `url` varchar(500) NOT NULL DEFAULT '' COMMENT '跳转路径',
  `is_read` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0未读 1已读',
  `is_emailed` tinyint(4) NOT NULL DEFAULT '0' COMMENT '-1不发送 0待发送 1发送中(已进入队列) 2发送成功 3发送失败',
  `read_time` datetime DEFAULT NULL COMMENT '阅读时间',
  `extra` text COMMENT '扩展数据',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_receiver_id` (`receiver_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='系统通知记录';

-- 数据导出被取消选择。

-- 导出  表 x_admin_2.x_system_notice_setting 结构
CREATE TABLE IF NOT EXISTS `x_system_notice_setting` (
  `id` char(36) NOT NULL COMMENT 'UUIDv7',
  `admin_id` char(36) NOT NULL COMMENT '用户ID',
  `channel` varchar(32) NOT NULL COMMENT '渠道: site/email/app',
  `is_enabled` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0关闭 1开启',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_admin_channel` (`admin_id`,`channel`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户通知渠道偏好';

-- 数据导出被取消选择。

-- 导出  表 x_admin_2.x_user 结构
CREATE TABLE IF NOT EXISTS `x_user` (
  `id` char(36) NOT NULL COMMENT 'UUID v7',
  `email` varchar(128) NOT NULL DEFAULT '' COMMENT '邮箱(主账号)',
  `nickname` varchar(64) NOT NULL DEFAULT '' COMMENT '昵称',
  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '密码(bcrypt)',
  `salt` varchar(32) NOT NULL DEFAULT '' COMMENT '加密盐',
  `phone` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号',
  `phone_code` varchar(10) NOT NULL DEFAULT '86' COMMENT '手机号区号',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态 0正常 1禁用',
  `token_version` bigint(20) NOT NULL DEFAULT '0' COMMENT 'Token版本号(踢人下线:自增使旧token失效)',
  `last_login_ip` varchar(50) NOT NULL DEFAULT '' COMMENT '最后登录IP',
  `last_login_time` datetime(3) DEFAULT NULL COMMENT '最后登录时间',
  `membership_level` varchar(20) NOT NULL DEFAULT 'free' COMMENT '会员等级 free/monthly/quarterly/annual/permanent',
  `membership_expires_at` datetime(3) DEFAULT NULL COMMENT '会员过期时间',
  `daily_quota` bigint(20) NOT NULL DEFAULT '0' COMMENT '每日录音次数配额',
  `quota_used_today` bigint(20) NOT NULL DEFAULT '0' COMMENT '今日已用录音次数',
  `total_quota_seconds` bigint(20) NOT NULL DEFAULT '0' COMMENT '语音时长总配额(秒)',
  `used_quota_seconds` bigint(20) NOT NULL DEFAULT '0' COMMENT '已用语音时长(秒)',
  `quota_reset_at` datetime(3) DEFAULT NULL COMMENT '配额重置时间',
  `is_delete` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除: 0=否, 1=是',
  `create_time` datetime(3) NOT NULL COMMENT '创建时间',
  `update_time` datetime(3) NOT NULL COMMENT '更新时间',
  `delete_time` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_phone` (`phone`),
  KEY `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户主表';

-- 数据导出被取消选择。

-- 导出  表 x_admin_2.x_user_auth 结构
CREATE TABLE IF NOT EXISTS `x_user_auth` (
  `id` char(36) NOT NULL COMMENT 'UUID v7',
  `user_id` char(36) NOT NULL COMMENT '用户ID',
  `identity_type` varchar(20) NOT NULL COMMENT '认证类型: phone/wechat_mini/wechat_mp/wechat_app/qq',
  `identifier` varchar(128) NOT NULL COMMENT '标识(手机号/openid/unionid/qq_openid)',
  `credential` varchar(255) NOT NULL DEFAULT '' COMMENT '凭证(加密存储，部分类型可为空)',
  `extra` varchar(512) NOT NULL DEFAULT '' COMMENT '扩展信息JSON(微信昵称、头像等)',
  `is_delete` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除: 0=否, 1=是',
  `create_time` datetime(3) NOT NULL COMMENT '创建时间',
  `update_time` datetime(3) NOT NULL COMMENT '更新时间',
  `delete_time` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_type_identifier` (`identity_type`,`identifier`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_x_user_auth_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='第三方认证绑定表';

-- 数据导出被取消选择。

-- 导出  表 x_admin_2.x_user_protocol 结构
CREATE TABLE IF NOT EXISTS `x_user_protocol` (
  `id` char(36) NOT NULL COMMENT 'uuid',
  `tag` varchar(50) DEFAULT NULL COMMENT '标识',
  `version` int(10) DEFAULT NULL COMMENT '版本',
  `title` varchar(50) NOT NULL COMMENT '标题',
  `content` mediumtext COMMENT '协议内容',
  `created_by` char(36) NOT NULL COMMENT '创建人',
  `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除: 0=否, 1=是',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COMMENT='用户协议';

-- 数据导出被取消选择。

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
