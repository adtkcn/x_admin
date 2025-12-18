-- --------------------------------------------------------
-- 主机:                           127.0.0.1
-- 服务器版本:                        5.7.26-log - MySQL Community Server (GPL)
-- 服务器操作系统:                      Win64
-- HeidiSQL 版本:                  12.3.0.6589
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
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_cid` (`cid`) USING BTREE,
  KEY `admin_id` (`admin_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='相册管理表';

-- 正在导出表  x_admin_2.x_album 的数据：~0 rows (大约)
DELETE FROM `x_album`;

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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='相册分类表';

-- 正在导出表  x_admin_2.x_album_cate 的数据：~0 rows (大约)
DELETE FROM `x_album_cate`;

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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='字典数据表';

-- 正在导出表  x_admin_2.x_dict_data 的数据：~20 rows (大约)
DELETE FROM `x_dict_data`;
INSERT INTO `x_dict_data` (`id`, `type_id`, `name`, `value`, `color`, `remark`, `sort`, `status`, `is_delete`, `create_time`, `update_time`, `delete_time`) VALUES
	('1', '2', '待提交', '1', '#6D85FC', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL),
	('10', '4', '假勤管理', '1', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL),
	('11', '4', '人事管理', '2', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL),
	('12', '4', '财务管理', '3', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL),
	('13', '4', '业务管理', '4', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL),
	('14', '4', '行政管理', '5', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL),
	('15', '4', '法务管理', '6', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL),
	('16', '4', '其他', '7', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL),
	('17', '5', 'web', 'web', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-06-29 00:18:02', NULL),
	('18', '5', 'go', 'go', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL),
	('19', '5', 'uniapp', 'uniapp', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL),
	('2', '2', '审批中', '2', '#C6C150', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL),
	('20', '5', 'node', 'node', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL),
	('22', '6', '禁用', '0', '#FF4B4B', '', 0, 1, 0, '2024-09-25 16:02:08', '2024-09-25 16:02:08', NULL),
	('23', '6', '启用', '1', '#80D251', '', 0, 1, 0, '2024-09-25 16:02:30', '2024-09-25 16:02:30', NULL),
	('3', '2', '审批成功', '3', 'green', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL),
	('4', '2', '失败', '4', 'red', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL),
	('5', '3', '待处理', '1', '#087BF6', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL),
	('6', '3', '通过', '2', 'green', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL),
	('7', '3', '拒绝', '3', 'red', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL);

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

-- 正在导出表  x_admin_2.x_dict_type 的数据：~5 rows (大约)
DELETE FROM `x_dict_type`;
INSERT INTO `x_dict_type` (`id`, `dict_name`, `dict_type`, `dict_remark`, `dict_status`, `is_delete`, `create_time`, `update_time`, `delete_time`) VALUES
	('2', '审批申请状态', 'flow_apply_status', '0待提交，1审批中，2审批完成，3审批失败', 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL),
	('3', '审批历史状态', 'flow_history_status', '', 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL),
	('4', '流程分类', 'flow_group', '1假勤管理,2人事管理3财务管理4业务管理5行政管理6法务管理7其他', 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL),
	('5', '项目类型', 'project_type', '项目类型go java web node php 等', 1, 0, '2024-01-02 03:04:05', '2024-06-29 00:48:34', NULL),
	('6', '启用状态', 'status', ' 0=否, 1=是', 1, 0, '2024-09-25 15:54:56', '2024-09-25 15:54:56', NULL);

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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='申请流程';

-- 正在导出表  x_admin_2.x_flow_apply 的数据：~0 rows (大约)
DELETE FROM `x_flow_apply`;

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
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `apply_id` (`apply_id`) USING BTREE,
  KEY `approver_id` (`approver_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='流程历史';

-- 正在导出表  x_admin_2.x_flow_history 的数据：~0 rows (大约)
DELETE FROM `x_flow_history`;

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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='流程模板';

-- 正在导出表  x_admin_2.x_flow_template 的数据：~0 rows (大约)
DELETE FROM `x_flow_template`;

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

-- 正在导出表  x_admin_2.x_gen_table 的数据：~0 rows (大约)
DELETE FROM `x_gen_table`;

-- 导出  表 x_admin_2.x_gen_table_column 结构
CREATE TABLE IF NOT EXISTS `x_gen_table_column` (
  `id` char(36) NOT NULL COMMENT '列主键',
  `table_id` char(36) NOT NULL COMMENT '表外键',
  `column_name` varchar(200) NOT NULL DEFAULT '' COMMENT '列名称',
  `column_comment` varchar(200) NOT NULL DEFAULT '' COMMENT '列描述',
  `column_length` varchar(5) DEFAULT '0' COMMENT '列长度',
  `column_type` varchar(100) NOT NULL DEFAULT '' COMMENT '列类型 ',
  `go_type` varchar(50) NOT NULL DEFAULT '0' COMMENT 'JAVA类型',
  `go_field` varchar(100) NOT NULL DEFAULT '' COMMENT 'JAVA字段',
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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='代码生成字段表';

-- 正在导出表  x_admin_2.x_gen_table_column 的数据：~0 rows (大约)
DELETE FROM `x_gen_table_column`;

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
  KEY `project_key` (`project_key`) USING BTREE,
  KEY `client_id` (`client_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='监控-客户端信息';

-- 正在导出表  x_admin_2.x_monitor_client 的数据：~0 rows (大约)
DELETE FROM `x_monitor_client`;

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

-- 正在导出表  x_admin_2.x_monitor_error 的数据：~0 rows (大约)
DELETE FROM `x_monitor_error`;

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
  `ip` varchar(50) CHARACTER SET utf8 DEFAULT NULL COMMENT 'ip',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `client_id` (`client_id`) USING BTREE,
  KEY `eid` (`error_id`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COMMENT='错误对应的用户记录';

-- 正在导出表  x_admin_2.x_monitor_error_list 的数据：0 rows
DELETE FROM `x_monitor_error_list`;
/*!40000 ALTER TABLE `x_monitor_error_list` DISABLE KEYS */;
/*!40000 ALTER TABLE `x_monitor_error_list` ENABLE KEYS */;

-- 导出  表 x_admin_2.x_monitor_project 结构
CREATE TABLE IF NOT EXISTS `x_monitor_project` (
  `id` char(36) NOT NULL DEFAULT '' COMMENT '项目id',
  `project_key` varchar(32) NOT NULL COMMENT 'project_key',
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

-- 正在导出表  x_admin_2.x_monitor_project 的数据：~4 rows (大约)
DELETE FROM `x_monitor_project`;
INSERT INTO `x_monitor_project` (`id`, `project_key`, `project_name`, `project_type`, `status`, `is_delete`, `create_time`, `update_time`, `delete_time`) VALUES
	('019b30ca-32af-7b43-98c0-a6631648e4ef', '019b30ca32af7b42a08d04e8ebac4888', '1', 'web', 1, 0, '2025-12-18 17:28:36', '2025-12-18 17:28:36', NULL),
	('6', '6217ea4ea0044014831bd25121a3113c', 'go', 'go', 0, 0, '2024-07-12 23:17:23', '2024-11-07 15:43:08', NULL),
	('7', 'e19e3be20de94f49b68fafb4c30668bc', 'web项目', 'web', 1, 0, '2024-07-13 20:56:21', '2024-09-25 16:58:58', NULL),
	('8', '019ab6afadb5798397361fd25d973ed8', 'a', 'go', 1, 0, '2025-11-25 00:25:57', '2025-11-25 00:25:57', NULL);

-- 导出  表 x_admin_2.x_system_auth_admin 结构
CREATE TABLE IF NOT EXISTS `x_system_auth_admin` (
  `id` char(36) NOT NULL COMMENT '主键',
  `dept_id` char(36) NOT NULL COMMENT '部门ID',
  `post_id` char(36) NOT NULL COMMENT '岗位ID',
  `username` varchar(32) NOT NULL DEFAULT '' COMMENT '用户账号',
  `nickname` varchar(32) NOT NULL DEFAULT '' COMMENT '用户昵称',
  `password` varchar(200) NOT NULL DEFAULT '' COMMENT '用户密码',
  `avatar` varchar(200) NOT NULL DEFAULT '' COMMENT '用户头像',
  `role` char(36) NOT NULL DEFAULT '' COMMENT '角色主键',
  `salt` varchar(20) NOT NULL DEFAULT '' COMMENT '加密盐巴',
  `sort` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '排序编号',
  `is_multipoint` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '多端登录: 0=否, 1=是',
  `is_disable` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否禁用: 0=否, 1=是',
  `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除: 0=否, 1=是',
  `last_login_ip` varchar(20) NOT NULL DEFAULT '' COMMENT '最后登录IP',
  `last_login_time` datetime DEFAULT NULL COMMENT '最后登录',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='系统管理成员表';

-- 正在导出表  x_admin_2.x_system_auth_admin 的数据：~1 rows (大约)
DELETE FROM `x_system_auth_admin`;
INSERT INTO `x_system_auth_admin` (`id`, `dept_id`, `post_id`, `username`, `nickname`, `password`, `avatar`, `role`, `salt`, `sort`, `is_multipoint`, `is_disable`, `is_delete`, `last_login_ip`, `last_login_time`, `create_time`, `update_time`, `delete_time`) VALUES
	('1', '1', '3', 'admin', 'admin', '81a13dd8e25644a8823082573ca973f7', '/png/20250916/4d8d6a13a3034380b8dbf9b95591839c.png', '0', 'WFdiD', 1, 1, 0, 0, '127.0.0.1', '2025-12-18 23:49:21', '2024-01-02 03:04:05', '2025-12-18 23:49:22', NULL);

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

-- 正在导出表  x_admin_2.x_system_auth_dept 的数据：~2 rows (大约)
DELETE FROM `x_system_auth_dept`;
INSERT INTO `x_system_auth_dept` (`id`, `pid`, `name`, `duty_id`, `duty`, `mobile`, `sort`, `is_stop`, `is_delete`, `create_time`, `update_time`, `delete_time`) VALUES
	('019b3244-cf89-7417-88b2-d6df02f5c132', '1', '11', '', '', '', 0, 0, 0, '2025-12-19 00:22:09', '2025-12-19 00:22:38', '2025-12-19 00:00:00'),
	('1', '0', '默认部门', '1', 'admin', '18327647788', 10, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL);

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

-- 正在导出表  x_admin_2.x_system_auth_menu 的数据：~131 rows (大约)
DELETE FROM `x_system_auth_menu`;
INSERT INTO `x_system_auth_menu` (`id`, `pid`, `menu_type`, `menu_name`, `menu_icon`, `menu_sort`, `perms`, `paths`, `component`, `selected`, `params`, `is_cache`, `is_show`, `is_disable`, `create_time`, `update_time`) VALUES
	('1', '0', 'C', '工作台', 'el-icon-Monitor', 50, 'admin:common:index:console', 'workbench', 'workbench/index', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('101', '500', 'C', '管理员', 'local-icon-wode', 10, 'admin:system:admin:list', 'admin', 'system/admin/index', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2025-09-18 11:41:20'),
	('102', '101', 'A', '管理员详情', '', 0, 'admin:system:admin:detail', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('103', '101', 'A', '管理员新增', '', 0, 'admin:system:admin:add', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('104', '101', 'A', '管理员编辑', '', 0, 'admin:system:admin:edit', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('105', '101', 'A', '管理员删除', '', 0, 'admin:system:admin:del', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('106', '101', 'A', '管理员状态', '', 0, 'admin:system:admin:disable', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('110', '500', 'C', '角色管理', 'el-icon-Female', 9, 'admin:system:role:list', 'role', 'system/role/index', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2025-09-18 11:43:09'),
	('111', '110', 'A', '角色详情', '', 0, 'admin:system:role:detail', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('112', '110', 'A', '角色新增', '', 0, 'admin:system:role:add', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('113', '110', 'A', '角色编辑', '', 0, 'admin:system:role:edit', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('114', '110', 'A', '角色删除', '', 0, 'admin:system:role:del', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('120', '500', 'C', '菜单管理', 'el-icon-Operation', 0, 'admin:system:menu:list', 'menu', 'system/menu/index', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2025-09-18 11:39:11'),
	('121', '120', 'A', '菜单详情', '', 0, 'admin:system:menu:detail', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('122', '120', 'A', '菜单新增', '', 0, 'admin:system:menu:add', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('123', '120', 'A', '菜单编辑', '', 0, 'admin:system:menu:edit', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('124', '120', 'A', '菜单删除', '', 0, 'admin:system:menu:del', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('131', '500', 'C', '部门管理', 'el-icon-Coordinate', 8, 'admin:system:dept:all', 'dept', 'system/dept/index', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2025-09-18 12:17:11'),
	('132', '131', 'A', '部门详情', '', 0, 'admin:system:dept:detail', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('133', '131', 'A', '部门新增', '', 0, 'admin:system:dept:add', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('134', '131', 'A', '部门编辑', '', 0, 'admin:system:dept:edit', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('135', '131', 'A', '部门删除', '', 0, 'admin:system:dept:del', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('140', '500', 'C', '岗位管理', 'el-icon-PriceTag', 7, 'admin:system:post:list', 'post', 'system/post/index', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2025-09-18 12:12:24'),
	('141', '140', 'A', '岗位详情', '', 0, 'admin:system:post:detail', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('142', '140', 'A', '岗位新增', '', 0, 'admin:system:post:add', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('143', '140', 'A', '岗位编辑', '', 0, 'admin:system:post:edit', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('144', '140', 'A', '岗位删除', '', 0, 'admin:system:post:del', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('200', '0', 'M', '其它管理', '', 0, '', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('201', '200', 'M', '图库管理', '', 0, '', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('202', '201', 'A', '文件列表', '', 0, 'admin:common:album:albumList', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('203', '201', 'A', '文件命名', '', 0, 'admin:common:album:albumRename', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('204', '201', 'A', '文件移动', '', 0, 'admin:common:album:albumMove', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('205', '201', 'A', '文件删除', '', 0, 'admin:common:album:albumDel', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('206', '201', 'A', '分类列表', '', 0, 'admin:common:album:cateList', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('207', '201', 'A', '分类新增', '', 0, 'admin:common:album:cateAdd', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('208', '201', 'A', '分类命名', '', 0, 'admin:common:album:cateRename', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('209', '201', 'A', '分类删除', '', 0, 'admin:common:album:cateDel', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('500', '0', 'M', '系统设置', 'el-icon-Setting', 0, '', 'system', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2025-09-18 12:57:22'),
	('501', '500', 'M', '网站设置', 'el-icon-Basketball', 0, '', 'website', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2025-09-18 11:20:05'),
	('502', '501', 'C', '网站信息', '', 0, 'admin:setting:website:detail', 'information', 'system/website/information', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2025-09-18 12:13:52'),
	('503', '502', 'A', '保存配置', '', 0, 'admin:setting:website:save', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('505', '501', 'C', '网站备案', '', 0, 'admin:setting:copyright:detail', 'filing', 'system/website/filing', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2025-09-18 12:15:00'),
	('506', '505', 'A', '备案保存', '', 0, 'admin:setting:copyright:save', '', 'setting/website/protocol', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('515', '500', 'C', '字典管理', 'el-icon-Box', 6, 'admin:setting:dict:type:list', 'dict', 'system/dict/type/index', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2025-09-18 12:12:58'),
	('516', '515', 'A', '字典类型新增', '', 0, 'admin:setting:dict:type:add', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('517', '515', 'A', '字典类型编辑', '', 0, 'admin:setting:dict:type:edit', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('518', '515', 'A', '字典类型删除', '', 0, 'admin:setting:dict:type:del', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('520', '515', 'A', '字典数据新增', '', 0, 'admin:setting:dict:data:add', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('521', '515', 'A', '字典数据编辑', '', 0, 'admin:setting:dict:data:edit', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('522', '515', 'A', '字典数据删除', '', 0, 'admin:setting:dict:data:del', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('550', '500', 'M', '系统维护', 'el-icon-SetUp', 0, '', 'system', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('551', '550', 'C', '系统环境', '', 0, 'admin:monitor:server', 'environment', 'system/system/environment', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2025-09-18 12:22:43'),
	('552', '550', 'C', '系统缓存', '', 0, 'admin:monitor:cache', 'system/cache', 'system/system/cache', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2025-09-18 12:21:45'),
	('553', '550', 'C', '系统日志', '', 0, 'admin:system:log:operate', 'journal', 'system/system/journal', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2025-09-18 12:22:00'),
	('600', '0', 'M', '开发工具', 'el-icon-EditPen', 0, '', 'dev_tools', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('610', '600', 'C', '代码生成器', '', 0, 'admin:gen:list', 'code', 'dev_tools/code/index', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-10-31 14:59:14'),
	('611', '610', 'A', '导入数据表', '', 0, 'admin:gen:importTable', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('612', '610', 'A', '生成代码', '', 0, 'admin:gen:genCode', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('613', '610', 'A', '下载代码', '', 0, 'admin:gen:downloadCode', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('614', '610', 'A', '预览代码', '', 0, 'admin:gen:previewCode', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('616', '610', 'A', '同步表结构', '', 0, 'admin:gen:syncTable', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('617', '610', 'A', '删除数据表', '', 0, 'admin:gen:delTable', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('618', '610', 'A', '数据表详情', '', 0, 'admin:gen:detail', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('701', '0', 'C', '素材中心', 'el-icon-PictureRounded', 44, '', 'material/index', 'material/index', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2025-09-19 00:39:29'),
	('775', '600', 'C', '代码生成器编辑', 'el-icon-EditPen', 0, 'admin:gen:editTable', 'code/edit', 'dev_tools/code/edit', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2025-09-18 12:29:37'),
	('776', '778', 'C', '流程模板', '', 0, 'admin:flow:flow_template:list', 'flow_template/index', 'flow/flow_template/index', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('777', '778', 'C', '我的流程', '', 0, '', 'flow_apply/index', 'flow/flow_apply/index', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('778', '0', 'M', '审批流', 'el-icon-Coordinate', 0, '', 'flow', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('780', '778', 'C', '待处理', '', 0, 'admin:flow:flow_history:list', 'flow_history/todo', 'flow/flow_history/todo', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('781', '778', 'C', '已处理', '', 0, 'admin:flow:flow_history:list', 'flow_history/done', 'flow/flow_history/done', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('782', '778', 'C', '已完成流程', '', 0, 'admin:flow:flow_history:list', 'flow_apply/finish', 'flow/flow_apply/finish', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('783', '832', 'C', '项目', '', 0, 'admin:monitor_project:list', 'project/index', 'monitor/project/index', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-10-29 15:40:44'),
	('784', '832', 'C', '用户端', '', 0, '', 'client/index', 'monitor/client/index', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-10-29 15:40:54'),
	('785', '794', 'A', '错误收集error添加', '', 0, 'admin:monitor_web:add', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('786', '794', 'A', '错误收集error编辑', '', 0, 'admin:monitor_web:edit', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('787', '794', 'A', '错误收集error删除', '', 0, 'admin:monitor_web:del', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('788', '794', 'A', '错误收集error列表', '', 0, 'admin:monitor_web:list', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('789', '794', 'A', '错误收集error全部列表', '', 0, 'admin:monitor_web:listAll', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('790', '794', 'A', '错误收集error详情', '', 0, 'admin:monitor_web:detail', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('791', '794', 'A', '错误收集error导出excel', '', 0, 'admin:monitor_web:ExportFile', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('792', '794', 'A', '错误收集error导入excel', '', 0, 'admin:monitor_web:ImportFile', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('794', '832', 'C', '错误列表', '', 0, '', 'error/index', 'monitor/error/index', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-10-29 15:39:16'),
	('795', '776', 'A', '流程模板添加', '', 0, 'admin:flow:flow_template:add', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('796', '776', 'A', '流程模板编辑', '', 0, 'admin:flow:flow_template:edit', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('797', '776', 'A', '流程模板删除', '', 0, 'admin:flow:flow_template:del', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('798', '776', 'A', '流程模板列表', '', 0, 'admin:flow:flow_template:list', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('799', '776', 'A', '流程模板全部列表', '', 0, 'admin:flow:flow_template:listAll', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('800', '776', 'A', '流程模板详情', '', 0, 'admin:flow:flow_template:detail', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('801', '777', 'A', '流程详情', '', 0, 'admin:flow:flow_apply:detail', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('802', '777', 'A', '流程添加', '', 0, 'admin:flow:flow_apply:add', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('803', '777', 'A', '流程编辑', '', 0, 'admin:flow:flow_apply:edit', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('804', '777', 'A', '流程删除', '', 0, 'admin:flow:flow_apply:del', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('805', '777', 'A', '全部流程', '', 0, 'admin:flow:flow_history:listAll', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('806', '777', 'A', '通过流程', '', 0, 'admin:flow:flow_history:pass', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('807', '777', 'A', '拒接流程', '', 0, 'admin:flow:flow_history:back', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('808', '777', 'A', '下一个流程', '', 0, 'admin:flow:flow_history:next_node', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('809', '777', 'A', '获取审批人', '', 0, 'admin:flow:flow_history:get_approver', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('810', '777', 'A', '流程列表', '', 0, 'admin:flow:flow_apply:list', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('811', '780', 'A', '审批记录列表', '', 0, 'admin:flow:flow_history:list', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('812', '780', 'A', '审批记录详情', '', 0, 'admin:flow:flow_history:detail', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('813', '783', 'A', '监控项目添加', '', 0, 'admin:monitor_project:add', '', '', '', '', 0, 1, 0, '2024-06-17 19:26:13', '2024-06-17 19:26:13'),
	('814', '783', 'A', '监控项目编辑', '', 0, 'admin:monitor_project:edit', '', '', '', '', 0, 1, 0, '2024-06-17 19:26:13', '2024-06-17 19:26:13'),
	('815', '783', 'A', '监控项目删除', '', 0, 'admin:monitor_project:del', '', '', '', '', 0, 1, 0, '2024-06-17 19:26:13', '2024-06-17 19:26:13'),
	('816', '783', 'A', '监控项目列表', '', 0, 'admin:monitor_project:list', '', '', '', '', 0, 1, 0, '2024-06-17 19:26:13', '2024-06-17 19:26:13'),
	('817', '783', 'A', '监控项目全部列表', '', 0, 'admin:monitor_project:listAll', '', '', '', '', 0, 1, 0, '2024-06-17 19:26:13', '2024-06-17 19:26:13'),
	('818', '783', 'A', '监控项目详情', '', 0, 'admin:monitor_project:detail', '', '', '', '', 0, 1, 0, '2024-06-17 19:26:13', '2024-06-17 19:26:13'),
	('819', '783', 'A', '监控项目导出excel', '', 0, 'admin:monitor_project:ExportFile', '', '', '', '', 0, 1, 0, '2024-06-17 19:26:13', '2024-06-17 19:26:13'),
	('820', '783', 'A', '监控项目导入excel', '', 0, 'admin:monitor_project:ImportFile', '', '', '', '', 0, 1, 0, '2024-06-17 19:26:13', '2024-06-17 19:26:13'),
	('821', '600', 'C', '开发文档', '', 0, '', 'https://adtkcn.github.io/x_admin/', '', '', 'a=1', 1, 1, 0, '2024-06-28 17:09:17', '2025-12-01 00:44:37'),
	('823', '0', 'C', '用户协议', 'el-icon-Coordinate', 0, '', 'user/protocol/index', 'user/protocol/index', '', '', 0, 1, 0, '2024-09-10 20:03:46', '2025-07-17 16:21:15'),
	('824', '823', 'A', '用户协议添加', '', 0, 'admin:user_protocol:add', '', '', '', '', 0, 1, 0, '2024-09-10 20:05:13', '2024-09-10 20:05:13'),
	('825', '823', 'A', '用户协议编辑', '', 0, 'admin:user_protocol:edit', '', '', '', '', 0, 1, 0, '2024-09-10 20:05:13', '2024-09-10 20:05:13'),
	('826', '823', 'A', '用户协议删除', '', 0, 'admin:user_protocol:del', '', '', '', '', 0, 1, 0, '2024-09-10 20:05:13', '2024-09-10 20:05:13'),
	('827', '823', 'A', '用户协议列表', '', 0, 'admin:user_protocol:list', '', '', '', '', 0, 1, 0, '2024-09-10 20:05:13', '2024-09-10 20:05:13'),
	('828', '823', 'A', '用户协议全部列表', '', 0, 'admin:user_protocol:listAll', '', '', '', '', 0, 1, 0, '2024-09-10 20:05:13', '2024-09-10 20:05:13'),
	('829', '823', 'A', '用户协议详情', '', 0, 'admin:user_protocol:detail', '', '', '', '', 0, 1, 0, '2024-09-10 20:05:13', '2024-09-10 20:05:13'),
	('830', '823', 'A', '用户协议导出excel', '', 0, 'admin:user_protocol:ExportFile', '', '', '', '', 0, 1, 0, '2024-09-10 20:05:13', '2024-09-10 20:05:13'),
	('831', '823', 'A', '用户协议导入excel', '', 0, 'admin:user_protocol:ImportFile', '', '', '', '', 0, 1, 0, '2024-09-10 20:05:13', '2024-09-10 20:05:13'),
	('832', '0', 'M', '错误监控', 'el-icon-Memo', 0, '', 'monitor', '', '', '', 1, 1, 0, '2024-10-29 15:32:33', '2024-10-29 15:32:33'),
	('833', '600', 'C', '错误捕获', '', 0, '', 'test', 'dev_tools/test', '', '', 1, 1, 0, '2024-10-30 18:16:56', '2025-09-18 12:28:08'),
	('864', '784', 'A', '监控-客户端信息添加', '', 0, 'admin:monitor_client:add', '', '', '', '', 0, 1, 0, '2024-11-08 14:15:28', '2024-11-08 14:15:28'),
	('865', '784', 'A', '监控-客户端信息编辑', '', 0, 'admin:monitor_client:edit', '', '', '', '', 0, 1, 0, '2024-11-08 14:15:28', '2024-11-08 14:15:28'),
	('866', '784', 'A', '监控-客户端信息删除', '', 0, 'admin:monitor_client:del', '', '', '', '', 0, 1, 0, '2024-11-08 14:15:28', '2024-11-08 14:15:28'),
	('867', '784', 'A', '监控-客户端信息删除-批量', '', 0, 'admin:monitor_client:delBatch', '', '', '', '', 0, 1, 0, '2024-11-08 14:15:28', '2024-11-08 14:15:28'),
	('868', '784', 'A', '监控-客户端信息列表', '', 0, 'admin:monitor_client:list', '', '', '', '', 0, 1, 0, '2024-11-08 14:15:28', '2024-11-08 14:15:28'),
	('869', '784', 'A', '监控-客户端信息全部列表', '', 0, 'admin:monitor_client:listAll', '', '', '', '', 0, 1, 0, '2024-11-08 14:15:28', '2024-11-08 14:15:28'),
	('870', '784', 'A', '监控-客户端信息详情', '', 0, 'admin:monitor_client:detail', '', '', '', '', 0, 1, 0, '2024-11-08 14:15:28', '2024-11-08 14:15:28'),
	('871', '784', 'A', '监控-客户端信息导出excel', '', 0, 'admin:monitor_client:ExportFile', '', '', '', '', 0, 1, 0, '2024-11-08 14:15:28', '2024-11-08 14:15:28'),
	('872', '784', 'A', '监控-客户端信息导入excel', '', 0, 'admin:monitor_client:ImportFile', '', '', '', '', 0, 1, 0, '2024-11-08 14:15:28', '2024-11-08 14:15:28'),
	('873', '600', 'C', '文件分片上传', '', 0, '', 'uploadChunk', 'dev_tools/uploadChunk', '', '', 1, 1, 0, '2025-07-19 18:37:48', '2025-09-18 12:27:57'),
	('874', '600', 'C', '接口文档', '', 0, '', 'apiDocs', 'IframeComponent', '', '{"url": "/api/static/api/index.html"}', 1, 1, 0, '2025-12-01 00:13:54', '2025-12-01 00:43:28');

-- 导出  表 x_admin_2.x_system_auth_perm 结构
CREATE TABLE IF NOT EXISTS `x_system_auth_perm` (
  `id` char(36) NOT NULL COMMENT '主键',
  `role_id` char(36) NOT NULL COMMENT '角色ID',
  `menu_id` char(36) NOT NULL COMMENT '菜单ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='系统角色菜单表';

-- 正在导出表  x_admin_2.x_system_auth_perm 的数据：~2 rows (大约)
DELETE FROM `x_system_auth_perm`;
INSERT INTO `x_system_auth_perm` (`id`, `role_id`, `menu_id`) VALUES
	('019b322c-6cbb-750d-a999-a0098992e16b', '1', '1'),
	('019b322c-6cbb-750e-ac54-455eca342d23', '1', '701');

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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='系统岗位管理表';

-- 正在导出表  x_admin_2.x_system_auth_post 的数据：~1 rows (大约)
DELETE FROM `x_system_auth_post`;
INSERT INTO `x_system_auth_post` (`id`, `code`, `name`, `remarks`, `sort`, `is_stop`, `is_delete`, `create_time`, `update_time`, `delete_time`) VALUES
	('3', 'zhihuibu01', '指挥部岗位', '', 0, 0, 0, '2024-01-02 03:04:05', '2025-07-16 17:33:24', NULL);

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

-- 正在导出表  x_admin_2.x_system_auth_role 的数据：~1 rows (大约)
DELETE FROM `x_system_auth_role`;
INSERT INTO `x_system_auth_role` (`id`, `name`, `remark`, `sort`, `is_disable`, `create_time`, `update_time`) VALUES
	('1', '审核员', '1', 1, 0, '2024-01-02 03:04:05', '2025-12-18 23:55:31');

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

-- 正在导出表  x_admin_2.x_system_config 的数据：~46 rows (大约)
DELETE FROM `x_system_config`;
INSERT INTO `x_system_config` (`id`, `type`, `name`, `value`, `create_time`, `update_time`) VALUES
	('1', 'storage', 'default', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('10', 'website', 'name', 'x_admin开源系统', '2024-01-02 03:04:05', '2025-06-24 19:52:37'),
	('11', 'website', 'logo', '/api/static/backend_logo.png', '2024-01-02 03:04:05', '2025-06-24 19:52:37'),
	('12', 'website', 'favicon', '/api/static/backend_favicon.ico', '2024-01-02 03:04:05', '2025-06-24 19:52:38'),
	('13', 'website', 'backdrop', '/api/static/backend_backdrop.png', '2024-01-02 03:04:05', '2025-06-24 19:52:38'),
	('14', 'website', 'copyright', '[{"name":"蜀ICP备15007060号-1","link":"http://www.beian.gov.cn"},{"name":"x_admin","link":"http://x.adtk.cn"}]', '2024-01-02 03:04:05', '2025-12-19 00:27:43'),
	('15', 'website', 'shopName', 'x_admin开源管理系统', '2024-01-02 03:04:05', '2025-06-24 19:52:38'),
	('16', 'website', 'shopLogo', '/image/20241705/f0eb36d508834bc2ac1b8c591c563efa.png', '2024-01-02 03:04:05', '2025-06-24 19:52:38'),
	('17', 'protocol', 'service', '{"name":"服务协议","content":"\\u003cp\\u003e服务协议666\\u003c/p\\u003e"}', '2024-01-02 03:04:05', '2024-06-29 00:30:56'),
	('18', 'protocol', 'privacy', '{"name":"隐私协议","content":"\\u003cp\\u003e隐私协议\\u003c/p\\u003e"}', '2024-01-02 03:04:05', '2024-06-29 00:30:56'),
	('19', 'tabbar', 'style', '{"defaultColor":"#4A5DFF","selectedColor":"#EA5455"}', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('2', 'storage', 'local', '{"name":"本地存储"}', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('20', 'search', 'isHotSearch', '0', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('3', 'storage', 'qiniu', '{"name":"七牛云存储","bucket":"","secretKey":"","accessKey":"","domain":""}', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('30', 'h5_channel', 'status', '1', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('31', 'h5_channel', 'close', '0', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('32', 'h5_channel', 'url', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('4', 'storage', 'aliyun', '{"name":"阿里云存储","bucket":"","secretKey":"","accessKey":"","domain":""}', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('40', 'mp_channel', 'name', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('41', 'mp_channel', 'primaryId', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('42', 'mp_channel', 'appId', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('43', 'mp_channel', 'appSecret', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('44', 'mp_channel', 'qrCode', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('5', 'storage', 'qcloud', '{"name":"腾讯云存储","bucket":"","secretKey":"","accessKey":"","domain":"","region":""}', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('50', 'wx_channel', 'appId', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('51', 'wx_channel', 'appSecret', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('55', 'oa_channel', 'name', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('56', 'oa_channel', 'primaryId', ' ', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('57', 'oa_channel', 'qrCode', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('58', 'oa_channel', 'appId', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('59', 'oa_channel', 'appSecret', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('6', 'sms', 'default', 'aliyun', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('60', 'oa_channel', 'url', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('61', 'oa_channel', 'token', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('62', 'oa_channel', 'encodingAesKey', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('63', 'oa_channel', 'encryptionType', '1', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('64', 'oa_channel', 'menus', '[]', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('7', 'sms', 'aliyun', '{"name":"阿里云短信","alias":"aliyun","sign":"","appKey":"","secretKey":""}', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('70', 'login', 'loginWay', '1,2', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('71', 'login', 'forceBindMobile', '0', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('72', 'login', 'openAgreement', '1', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('73', 'login', 'openOtherAuth', '1', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('74', 'login', 'autoLoginAuth', '1,2', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('8', 'sms', 'tencent', '{"name":"腾讯云短信","alias":"tencent","sign":"","appId":"","secretId":"","secretKey":""}', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('80', 'user', 'defaultAvatar', '/api/static/default_avatar.png', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	('9', 'sms', 'huawei', '{"name":"华为云短信","alias":"huawei"}', '2024-01-02 03:04:05', '2024-01-02 03:04:05');

-- 导出  表 x_admin_2.x_system_log_login 结构
CREATE TABLE IF NOT EXISTS `x_system_log_login` (
  `id` char(50) NOT NULL COMMENT '注解',
  `admin_id` char(50) NOT NULL COMMENT '管理员ID',
  `username` varchar(30) NOT NULL DEFAULT '' COMMENT '登录账号',
  `ip` varchar(30) NOT NULL COMMENT '登录地址',
  `os` varchar(100) NOT NULL DEFAULT '' COMMENT '操作系统',
  `browser` varchar(100) DEFAULT '' COMMENT '浏览器',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '操作状态: 1=成功, 2=失败',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='系统登录日志表';

-- 正在导出表  x_admin_2.x_system_log_login 的数据：~0 rows (大约)
DELETE FROM `x_system_log_login`;

-- 导出  表 x_admin_2.x_system_log_operate 结构
CREATE TABLE IF NOT EXISTS `x_system_log_operate` (
  `id` char(36) NOT NULL COMMENT '主键',
  `admin_id` char(36) NOT NULL COMMENT '操作人ID',
  `type` varchar(30) NOT NULL DEFAULT '' COMMENT '请求类型: GET/POST/PUT',
  `title` varchar(30) DEFAULT '' COMMENT '操作标题',
  `ip` varchar(30) NOT NULL DEFAULT '' COMMENT '请求IP',
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

-- 正在导出表  x_admin_2.x_system_log_operate 的数据：~2 rows (大约)
DELETE FROM `x_system_log_operate`;
INSERT INTO `x_system_log_operate` (`id`, `admin_id`, `type`, `title`, `ip`, `url`, `method`, `args`, `error`, `status`, `start_time`, `end_time`, `task_time`, `create_time`) VALUES
	('1', '1', 'POST', '用户协议编辑', '127.0.0.1', '/api/admin/user_protocol/edit', 'x_admin/app/controller/admin_ctl.(*UserProtocolHandler).Edit-fm', '{"Content":"\\u003cp\\u003e123\\u003c/p\\u003e","Id":"019ab685-34a5-7bb0-b0f1-2755c0b9c4fb","Tag":"123","Title":"123","Version":166}', 'Bad field type core.NullInt', 2, '2025-12-05 22:50:42', '2025-12-05 22:50:42', 1, '2025-12-05 22:50:42'),
	('2', '1', 'POST', '用户协议编辑', '127.0.0.1', '/api/admin/user_protocol/edit', 'x_admin/app/controller/admin_ctl.(*UserProtocolHandler).Edit-fm', '{"Content":"\\u003cp\\u003e123\\u003c/p\\u003e","Id":"019ab685-34a5-7bb0-b0f1-2755c0b9c4fb","Tag":"123","Title":"123","Version":166}', 'Bad field type core.NullInt', 2, '2025-12-05 22:51:31', '2025-12-05 22:52:01', 30460, '2025-12-05 22:52:01');

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

-- 正在导出表  x_admin_2.x_system_log_sms 的数据：~0 rows (大约)
DELETE FROM `x_system_log_sms`;

-- 导出  表 x_admin_2.x_user 结构
CREATE TABLE IF NOT EXISTS `x_user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `sn` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '编号',
  `avatar` varchar(200) NOT NULL DEFAULT '' COMMENT '头像',
  `real_name` varchar(32) NOT NULL DEFAULT '' COMMENT '真实姓名',
  `nickname` varchar(32) NOT NULL DEFAULT '' COMMENT '用户昵称',
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `mobile` varchar(32) NOT NULL DEFAULT '' COMMENT '用户电话',
  `salt` varchar(32) NOT NULL DEFAULT '' COMMENT '加密盐巴',
  `sex` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '用户性别: [1=男, 2=女]',
  `channel` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '注册渠道: [1=微信小程序, 2=微信公众号, 3=手机H5, 4=电脑PC, 5=苹果APP, 6=安卓APP]',
  `is_disable` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否禁用: [0=否, 1=是]',
  `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除: [0=否, 1=是]',
  `last_login_ip` varchar(30) NOT NULL DEFAULT '' COMMENT '最后登录IP',
  `last_login_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '最后登录时间',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='用户信息表';

-- 正在导出表  x_admin_2.x_user 的数据：~0 rows (大约)
DELETE FROM `x_user`;

-- 导出  表 x_admin_2.x_user_auth 结构
CREATE TABLE IF NOT EXISTS `x_user_auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `openid` varchar(200) NOT NULL DEFAULT '' COMMENT 'Openid',
  `unionid` varchar(200) NOT NULL DEFAULT '' COMMENT 'Unionid',
  `client` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '客户端类型: [1=微信小程序, 2=微信公众号, 3=手机H5, 4=电脑PC, 5=苹果APP, 6=安卓APP]',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `openid` (`openid`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='用户授权表';

-- 正在导出表  x_admin_2.x_user_auth 的数据：~0 rows (大约)
DELETE FROM `x_user_auth`;

-- 导出  表 x_admin_2.x_user_protocol 结构
CREATE TABLE IF NOT EXISTS `x_user_protocol` (
  `id` char(36) NOT NULL COMMENT 'uuid',
  `tag` varchar(50) DEFAULT NULL COMMENT '标识',
  `version` int(10) DEFAULT NULL COMMENT '版本',
  `title` varchar(50) NOT NULL COMMENT '标题',
  `content` mediumtext COMMENT '协议内容',
  `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除: 0=否, 1=是',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `created_by` int(11) NOT NULL DEFAULT '0' COMMENT '创建人',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COMMENT='用户协议';

-- 正在导出表  x_admin_2.x_user_protocol 的数据：0 rows
DELETE FROM `x_user_protocol`;
/*!40000 ALTER TABLE `x_user_protocol` DISABLE KEYS */;
/*!40000 ALTER TABLE `x_user_protocol` ENABLE KEYS */;

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
