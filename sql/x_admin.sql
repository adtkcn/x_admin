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
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_cid` (`cid`) USING BTREE,
  KEY `admin_id` (`admin_id`),
  KEY `hash` (`hash`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='相册管理表';

-- 正在导出表  x_admin_2.x_album 的数据：~3 rows (大约)
INSERT INTO `x_album` (`id`, `cid`, `admin_id`, `uid`, `name`, `uri`, `ext`, `hash`, `size`, `is_delete`, `create_time`, `update_time`, `delete_time`) VALUES
	('019b34af-4726-793d-92dd-d7aa34d1c41b', '0', '1', 0, 'logo.png', 'png/20251219/019b34af4723708d8c6964602edeb1f2.png', 'png', 'c84e7b0fc945d625e85bd0790223287b', 99626, 0, '2025-12-19 11:37:40', '2025-12-19 11:37:40', NULL),
	('019ce5fd-2e11-7be1-9b6f-198dee819200', '0', '1', 0, '019b308a1ca57a17b3e617bd481c00c1.png', 'png/20260313/019ce5fd-2e01-7db0-ab0e-c1ace63e55cb.png', 'png', 'c84e7b0fc945d625e85bd0790223287b', 99626, 1, '2026-03-13 14:58:20', '2026-03-28 16:03:19', '2026-03-28 16:03:19'),
	('019ce828-1687-7654-9b3c-a9572da4933a', '0', '1', 0, '019b308a1ca57a17b3e617bd481c00c1.png', 'png/20260314/019ce828-167c-7266-a5f4-6f1b50800d20.png', 'png', 'c84e7b0fc945d625e85bd0790223287b', 99626, 1, '2026-03-14 01:04:27', '2026-03-28 16:03:17', '2026-03-28 16:03:17');

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

-- 正在导出表  x_admin_2.x_album_cate 的数据：~0 rows (大约)

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

-- 正在导出表  x_admin_2.x_dict_data 的数据：~20 rows (大约)
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
  PRIMARY KEY (`id`) USING BTREE,
  KEY `apply_user_id` (`apply_user_id`),
  KEY `is_delete` (`is_delete`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='申请流程';

-- 正在导出表  x_admin_2.x_flow_apply 的数据：~1 rows (大约)
INSERT INTO `x_flow_apply` (`id`, `template_id`, `apply_user_id`, `apply_user_nickname`, `flow_name`, `flow_group`, `flow_remark`, `flow_form_data`, `flow_process_data`, `flow_process_data_list`, `form_value`, `status`, `is_delete`, `create_time`, `update_time`, `delete_time`) VALUES
	('019b37f0-c72c-7537-b64f-3dc82946db4b', '019b37f0-9d38-743b-ae7a-8680dfa5a838', '1', 'admin', '1_2025-12-20-0248', 1, '1', '[{"type":"input","field":"F78bmjd81k41aec","title":"输入框","info":"","$required":false,"_fc_id":"id_Fes7mjd81k41afc","name":"ref_Frxtmjd81k41agc","display":true,"hidden":false,"_fc_drag_tag":"input"}]', '{"nodes":[{"id":"Event_d91de17","type":"bpmn:startEvent","x":345.6000061035156,"y":159.60000610351562,"properties":{"width":36,"height":36},"zIndex":1020,"text":{"x":345.6000061035156,"y":199.60000610351562,"value":"开始"}},{"id":"Event_bf14d1c","type":"bpmn:endEvent","x":810.4000244140625,"y":160.40000915527344,"properties":{"width":36,"height":36},"zIndex":1014,"text":{"x":810.4000244140625,"y":200.40000915527344,"value":"结束"}},{"id":"Activity_3ef229e","type":"bpmn:userTask","x":599.2000122070312,"y":160.39999389648438,"properties":{"userType":3,"userId":"1","deptId":null,"postId":null,"fieldAuth":{"F78bmjd81k41aec":1},"gateway":[],"width":100,"height":80},"zIndex":1026,"text":{"x":599.2000122070312,"y":160.39999389648438,"value":"审批"}}],"edges":[{"id":"fc6fc2bc-2670-4355-bbe7-fe8634ede6f6","type":"pro-polyline","properties":{},"sourceNodeId":"Event_d91de17","targetNodeId":"Activity_3ef229e","sourceAnchorId":"Event_d91de17_1","targetAnchorId":"Activity_3ef229e_3","startPoint":{"x":363.6000061035156,"y":159.60000610351562},"endPoint":{"x":549.2000122070312,"y":160.39999389648438},"zIndex":1021,"pointsList":[{"x":363.6000061035156,"y":159.60000610351562},{"x":456.40000915527344,"y":159.60000610351562},{"x":456.40000915527344,"y":160.39999389648438},{"x":549.2000122070312,"y":160.39999389648438}]},{"id":"3bd10a02-68f6-4600-a149-50a0d075af69","type":"pro-polyline","properties":{},"sourceNodeId":"Activity_3ef229e","targetNodeId":"Event_bf14d1c","sourceAnchorId":"Activity_3ef229e_1","targetAnchorId":"Event_bf14d1c_3","startPoint":{"x":649.2000122070312,"y":160.39999389648438},"endPoint":{"x":792.4000244140625,"y":160.40000915527344},"zIndex":1023,"pointsList":[{"x":649.2000122070312,"y":160.39999389648438},{"x":720.8000183105469,"y":160.39999389648438},{"x":720.8000183105469,"y":160.40000915527344},{"x":792.4000244140625,"y":160.40000915527344}]}]}', '[{"id":"Event_d91de17","pid":0,"label":"开始","type":"bpmn:startEvent","userType":0,"userId":0,"deptId":0,"postId":0,"children":[{"id":"Activity_3ef229e","pid":"Event_d91de17","label":"审批","type":"bpmn:userTask","fieldAuth":{"F78bmjd81k41aec":1},"gateway":[],"userType":3,"userId":"1","deptId":0,"postId":0,"children":[{"id":"Event_bf14d1c","pid":"Activity_3ef229e","label":"结束","type":"bpmn:endEvent","userType":0,"userId":0,"deptId":0,"postId":0,"children":null}]}]},{"id":"Activity_3ef229e","pid":"Event_d91de17","label":"审批","type":"bpmn:userTask","fieldAuth":{"F78bmjd81k41aec":1},"gateway":[],"userType":3,"userId":"1","deptId":0,"postId":0,"children":[{"id":"Event_bf14d1c","pid":"Activity_3ef229e","label":"结束","type":"bpmn:endEvent","userType":0,"userId":0,"deptId":0,"postId":0,"children":null}]},{"id":"Event_bf14d1c","pid":"Activity_3ef229e","label":"结束","type":"bpmn:endEvent","userType":0,"userId":0,"deptId":0,"postId":0,"children":null},{"id":"Event_bf14d1c","pid":0,"label":"结束","type":"bpmn:endEvent","userType":0,"userId":0,"deptId":0,"postId":0,"children":null},{"id":"Activity_3ef229e","pid":0,"label":"审批","type":"bpmn:userTask","fieldAuth":{"F78bmjd81k41aec":1},"gateway":[],"userType":3,"userId":"1","deptId":0,"postId":0,"children":[{"id":"Event_bf14d1c","pid":"Activity_3ef229e","label":"结束","type":"bpmn:endEvent","userType":0,"userId":0,"deptId":0,"postId":0,"children":null}]},{"id":"Event_bf14d1c","pid":"Activity_3ef229e","label":"结束","type":"bpmn:endEvent","userType":0,"userId":0,"deptId":0,"postId":0,"children":null}]', '{"F78bmjd81k41aec":"1"}', 3, 0, '2025-12-20 02:48:05', '2025-12-20 02:50:30', NULL);

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

-- 正在导出表  x_admin_2.x_flow_history 的数据：~3 rows (大约)
INSERT INTO `x_flow_history` (`id`, `apply_id`, `template_id`, `apply_user_id`, `apply_user_nickname`, `approver_id`, `approver_nickname`, `node_id`, `node_type`, `node_label`, `form_value`, `pass_status`, `pass_remark`, `is_delete`, `create_time`, `update_time`, `delete_time`) VALUES
	('019b37f1-1252-7310-9c9f-ef81789cd979', '019b37f0-c72c-7537-b64f-3dc82946db4b', '019b37f0-9d38-743b-ae7a-8680dfa5a838', '1', 'admin', '', '', 'Event_d91de17', 'bpmn:startEvent', '开始', '{"F78bmjd81k41aec":"1"}', 2, '', 0, '2025-12-20 02:48:24', '2025-12-20 02:48:24', NULL),
	('019b37f1-1252-7311-b77f-b43664066dd7', '019b37f0-c72c-7537-b64f-3dc82946db4b', '019b37f0-9d38-743b-ae7a-8680dfa5a838', '1', 'admin', '1', 'admin', 'Activity_3ef229e', 'bpmn:userTask', '审批', '{"F78bmjd81k41aec":"1"}', 2, '1', 0, '2025-12-20 02:48:24', '2025-12-20 02:50:30', NULL),
	('019b37f2-fe7c-7bc9-b10d-619b46887133', '019b37f0-c72c-7537-b64f-3dc82946db4b', '019b37f0-9d38-743b-ae7a-8680dfa5a838', '1', 'admin', '', '', 'Event_bf14d1c', 'bpmn:endEvent', '结束', '{"F78bmjd81k41aec":"1"}', 2, '', 0, '2025-12-20 02:50:30', '2025-12-20 02:50:30', NULL);

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

-- 正在导出表  x_admin_2.x_flow_template 的数据：~1 rows (大约)
INSERT INTO `x_flow_template` (`id`, `flow_name`, `flow_group`, `flow_remark`, `flow_form_data`, `flow_process_data`, `flow_process_data_list`, `is_delete`, `create_time`, `update_time`, `delete_time`) VALUES
	('019b37f0-9d38-743b-ae7a-8680dfa5a838', '1', 1, '1', '[{"type":"input","field":"F78bmjd81k41aec","title":"输入框","info":"","$required":false,"_fc_id":"id_Fes7mjd81k41afc","name":"ref_Frxtmjd81k41agc","display":true,"hidden":false,"_fc_drag_tag":"input"}]', '{"nodes":[{"id":"Event_d91de17","type":"bpmn:startEvent","x":345.6000061035156,"y":159.60000610351562,"properties":{"width":36,"height":36},"zIndex":1020,"text":{"x":345.6000061035156,"y":199.60000610351562,"value":"开始"}},{"id":"Event_bf14d1c","type":"bpmn:endEvent","x":810.4000244140625,"y":160.40000915527344,"properties":{"width":36,"height":36},"zIndex":1014,"text":{"x":810.4000244140625,"y":200.40000915527344,"value":"结束"}},{"id":"Activity_3ef229e","type":"bpmn:userTask","x":599.2000122070312,"y":160.39999389648438,"properties":{"userType":3,"userId":"1","deptId":null,"postId":null,"fieldAuth":{"F78bmjd81k41aec":1},"gateway":[],"width":100,"height":80},"zIndex":1026,"text":{"x":599.2000122070312,"y":160.39999389648438,"value":"审批"}}],"edges":[{"id":"fc6fc2bc-2670-4355-bbe7-fe8634ede6f6","type":"pro-polyline","properties":{},"sourceNodeId":"Event_d91de17","targetNodeId":"Activity_3ef229e","sourceAnchorId":"Event_d91de17_1","targetAnchorId":"Activity_3ef229e_3","startPoint":{"x":363.6000061035156,"y":159.60000610351562},"endPoint":{"x":549.2000122070312,"y":160.39999389648438},"zIndex":1021,"pointsList":[{"x":363.6000061035156,"y":159.60000610351562},{"x":456.40000915527344,"y":159.60000610351562},{"x":456.40000915527344,"y":160.39999389648438},{"x":549.2000122070312,"y":160.39999389648438}]},{"id":"3bd10a02-68f6-4600-a149-50a0d075af69","type":"pro-polyline","properties":{},"sourceNodeId":"Activity_3ef229e","targetNodeId":"Event_bf14d1c","sourceAnchorId":"Activity_3ef229e_1","targetAnchorId":"Event_bf14d1c_3","startPoint":{"x":649.2000122070312,"y":160.39999389648438},"endPoint":{"x":792.4000244140625,"y":160.40000915527344},"zIndex":1023,"pointsList":[{"x":649.2000122070312,"y":160.39999389648438},{"x":720.8000183105469,"y":160.39999389648438},{"x":720.8000183105469,"y":160.40000915527344},{"x":792.4000244140625,"y":160.40000915527344}]}]}', '[{"id":"Event_d91de17","pid":0,"label":"开始","type":"bpmn:startEvent","userType":0,"userId":0,"deptId":0,"postId":0,"children":[{"id":"Activity_3ef229e","pid":"Event_d91de17","label":"审批","type":"bpmn:userTask","fieldAuth":{"F78bmjd81k41aec":1},"gateway":[],"userType":3,"userId":"1","deptId":0,"postId":0,"children":[{"id":"Event_bf14d1c","pid":"Activity_3ef229e","label":"结束","type":"bpmn:endEvent","userType":0,"userId":0,"deptId":0,"postId":0,"children":null}]}]},{"id":"Activity_3ef229e","pid":"Event_d91de17","label":"审批","type":"bpmn:userTask","fieldAuth":{"F78bmjd81k41aec":1},"gateway":[],"userType":3,"userId":"1","deptId":0,"postId":0,"children":[{"id":"Event_bf14d1c","pid":"Activity_3ef229e","label":"结束","type":"bpmn:endEvent","userType":0,"userId":0,"deptId":0,"postId":0,"children":null}]},{"id":"Event_bf14d1c","pid":"Activity_3ef229e","label":"结束","type":"bpmn:endEvent","userType":0,"userId":0,"deptId":0,"postId":0,"children":null},{"id":"Event_bf14d1c","pid":0,"label":"结束","type":"bpmn:endEvent","userType":0,"userId":0,"deptId":0,"postId":0,"children":null},{"id":"Activity_3ef229e","pid":0,"label":"审批","type":"bpmn:userTask","fieldAuth":{"F78bmjd81k41aec":1},"gateway":[],"userType":3,"userId":"1","deptId":0,"postId":0,"children":[{"id":"Event_bf14d1c","pid":"Activity_3ef229e","label":"结束","type":"bpmn:endEvent","userType":0,"userId":0,"deptId":0,"postId":0,"children":null}]},{"id":"Event_bf14d1c","pid":"Activity_3ef229e","label":"结束","type":"bpmn:endEvent","userType":0,"userId":0,"deptId":0,"postId":0,"children":null}]', 0, '2025-12-20 02:47:54', '2025-12-20 02:47:54', NULL);

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

-- 正在导出表  x_admin_2.x_gen_table_column 的数据：~0 rows (大约)

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

-- 正在导出表  x_admin_2.x_monitor_client 的数据：~4 rows (大约)
INSERT INTO `x_monitor_client` (`id`, `project_key`, `client_id`, `os`, `browser`, `ua`, `create_time`) VALUES
	('019b3fa3-0dae-7c9b-baeb-bc639a461f7e', 'e19e3be20de94f49b68fafb4c30668bc', 'b2524cb0-d9c8-11f0-a305-47d852bfb336', 'Windows', 'Edge', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/146.0.0.0 Safari/537.36 Edg/146.0.0.0', '2025-12-21 14:40:09'),
	('019cde6b-7d19-7bab-8f41-6147f229fa3f', 'e19e3be20de94f49b68fafb4c30668bc', '67c11650-dcc9-11f0-b463-1509576f75a9', 'Windows', 'Firefox', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:148.0) Gecko/20100101 Firefox/148.0', '2026-03-12 03:41:52'),
	('019cfefb-0a04-7c0d-be9b-36142babd78a', 'e19e3be20de94f49b68fafb4c30668bc', '41b213a0-227a-11f1-8ca5-0d6c441551d7', 'Windows', 'Edge', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/146.0.0.0 Safari/537.36 Edg/146.0.0.0', '2026-03-18 11:26:30'),
	('019cff02-8745-750f-a267-f49339acec0d', 'e19e3be20de94f49b68fafb4c30668bc', '664166c0-227b-11f1-9124-afe167399c62', 'Windows', 'Edge', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/146.0.0.0 Safari/537.36 Edg/146.0.0.0', '2026-03-18 11:34:41');

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

-- 正在导出表  x_admin_2.x_monitor_error_list 的数据：0 rows

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

-- 正在导出表  x_admin_2.x_monitor_project 的数据：~4 rows (大约)
INSERT INTO `x_monitor_project` (`id`, `project_key`, `project_name`, `project_type`, `status`, `is_delete`, `create_time`, `update_time`, `delete_time`) VALUES
	('019b30ca-32af-7b43-98c0-a6631648e4ef', '019b30ca32af7b42a08d04e8ebac4888', '1', 'web', 1, 1, '2025-12-18 17:28:36', '2025-12-19 16:09:52', '2025-12-19 16:09:52'),
	('6', '6217ea4ea0044014831bd25121a3113c', 'go', 'go', 0, 1, '2024-07-12 23:17:23', '2026-01-26 00:32:42', '2026-01-26 00:32:42'),
	('7', 'e19e3be20de94f49b68fafb4c30668bc', 'web项目', 'web', 1, 0, '2024-07-13 20:56:21', '2024-09-25 16:58:58', NULL),
	('8', '019ab6afadb5798397361fd25d973ed8', 'a', 'go', 1, 1, '2025-11-25 00:25:57', '2025-12-19 16:09:56', '2025-12-19 16:09:56');

-- 导出  表 x_admin_2.x_system_auth_admin 结构
CREATE TABLE IF NOT EXISTS `x_system_auth_admin` (
  `id` char(36) NOT NULL COMMENT '主键',
  `username` varchar(32) NOT NULL DEFAULT '' COMMENT '用户账号',
  `nickname` varchar(32) NOT NULL DEFAULT '' COMMENT '用户昵称',
  `password` varchar(200) NOT NULL DEFAULT '' COMMENT '用户密码',
  `dept_id` char(36) NOT NULL COMMENT '部门ID',
  `post_id` char(36) NOT NULL COMMENT '岗位ID',
  `avatar` varchar(200) NOT NULL DEFAULT '' COMMENT '用户头像',
  `salt` varchar(20) NOT NULL DEFAULT '' COMMENT '加密盐巴',
  `sort` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '排序编号',
  `is_multipoint` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '多端登录: 0=否, 1=是',
  `is_disable` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否禁用: 0=否, 1=是',
  `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除: 0=否, 1=是',
  `last_login_ip` varchar(39) NOT NULL DEFAULT '' COMMENT '最后登录IP',
  `last_login_time` datetime DEFAULT NULL COMMENT '最后登录',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `username` (`username`),
  KEY `is_delete` (`is_delete`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='系统管理成员表';

-- 正在导出表  x_admin_2.x_system_auth_admin 的数据：~1 rows (大约)
INSERT INTO `x_system_auth_admin` (`id`, `username`, `nickname`, `password`, `dept_id`, `post_id`, `avatar`, `salt`, `sort`, `is_multipoint`, `is_disable`, `is_delete`, `last_login_ip`, `last_login_time`, `create_time`, `update_time`, `delete_time`) VALUES
	('1', 'admin', 'admin', '81a13dd8e25644a8823082573ca973f7', '', '', '/png/20250916/4d8d6a13a3034380b8dbf9b95591839c.png', 'WFdiD', 1, 1, 0, 0, '127.0.0.1', '2026-03-30 11:27:30', '2024-01-02 03:04:05', '2026-03-30 11:27:30', NULL);

-- 导出  表 x_admin_2.x_system_auth_admin_role 结构
CREATE TABLE IF NOT EXISTS `x_system_auth_admin_role` (
  `id` char(36) NOT NULL COMMENT 'uuid',
  `admin_id` char(36) NOT NULL COMMENT '管理员ID',
  `role_id` char(36) NOT NULL COMMENT '角色ID',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_admin_role` (`admin_id`,`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户角色关联表';

-- 正在导出表  x_admin_2.x_system_auth_admin_role 的数据：~0 rows (大约)

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

-- 正在导出表  x_admin_2.x_system_auth_dept 的数据：~3 rows (大约)
INSERT INTO `x_system_auth_dept` (`id`, `pid`, `name`, `duty_id`, `duty`, `mobile`, `sort`, `is_stop`, `is_delete`, `create_time`, `update_time`, `delete_time`) VALUES
	('019b37e4-edbb-75fb-8aeb-b00e2591c7e1', '', '1', '', '', '13111111111', 0, 0, 1, '2025-12-20 02:35:08', '2025-12-20 02:35:08', '2026-03-12 10:58:21'),
	('019b37e6-6c40-7d4c-a796-4749d5cab39c', '019b37e4-edbb-75fb-8aeb-b00e2591c7e1', '2', '019b360a-4bd9-718c-ba71-46429ffd8b53', '333', '13222222222', 0, 0, 1, '2025-12-20 02:36:46', '2025-12-20 02:42:56', '2026-03-11 21:38:48'),
	('019ce028-6982-7dc9-a683-aaae0a861547', '', '顶级部门', '', '', '', 0, 0, 0, '2026-03-12 11:47:50', '2026-03-12 11:47:50', '2026-03-12 00:00:00');

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

-- 正在导出表  x_admin_2.x_system_auth_menu 的数据：~141 rows (大约)
INSERT INTO `x_system_auth_menu` (`id`, `pid`, `menu_type`, `menu_name`, `menu_icon`, `menu_sort`, `perms`, `paths`, `component`, `selected`, `params`, `is_cache`, `is_show`, `is_disable`, `create_time`, `update_time`) VALUES
	('019ce5eb-2ff4-7007-9ded-7d60d5b8c1ee', '', 'C', '工作台', 'el-icon-Monitor', 50, 'admin:common:index:console', 'workbench', 'workbench/index', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7008-a001-f311accef586', '019ce5eb-2ff4-702c-b7a4-5ced23241090', 'C', '管理员', 'local-icon-wode', 10, 'admin:system:admin:list', 'admin', 'system/admin/index', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7009-92eb-df4f0f8b0de4', '019ce5eb-2ff4-7008-a001-f311accef586', 'A', '管理员详情', '', 0, 'admin:system:admin:detail', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-700a-a80b-9f093885dc21', '019ce5eb-2ff4-7008-a001-f311accef586', 'A', '管理员新增', '', 0, 'admin:system:admin:add', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-700b-ae74-d742cc570476', '019ce5eb-2ff4-7008-a001-f311accef586', 'A', '管理员编辑', '', 0, 'admin:system:admin:edit', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-700c-9f62-45eca0c41903', '019ce5eb-2ff4-7008-a001-f311accef586', 'A', '管理员删除', '', 0, 'admin:system:admin:del', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-700d-96f1-d3607a8e93a0', '019ce5eb-2ff4-7008-a001-f311accef586', 'A', '管理员状态', '', 0, 'admin:system:admin:disable', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-700e-805c-527092dc3c0c', '019ce5eb-2ff4-702c-b7a4-5ced23241090', 'C', '角色管理', 'el-icon-Female', 9, 'admin:system:role:list', 'role', 'system/role/index', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-700f-8a12-913d79bdcd8b', '019ce5eb-2ff4-700e-805c-527092dc3c0c', 'A', '角色详情', '', 0, 'admin:system:role:detail', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7010-ac3b-4cff20f1ffc9', '019ce5eb-2ff4-700e-805c-527092dc3c0c', 'A', '角色新增', '', 0, 'admin:system:role:add', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7011-9a36-ff221ddc089a', '019ce5eb-2ff4-700e-805c-527092dc3c0c', 'A', '角色编辑', '', 0, 'admin:system:role:edit', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7012-9089-67fa5fffa6d0', '019ce5eb-2ff4-700e-805c-527092dc3c0c', 'A', '角色删除', '', 0, 'admin:system:role:del', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7013-9081-f329cca853d4', '019ce5eb-2ff4-702c-b7a4-5ced23241090', 'C', '菜单管理', 'el-icon-Operation', 0, 'admin:system:menu:list', 'menu', 'system/menu/index', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7014-a4c6-795d061beafd', '019ce5eb-2ff4-7013-9081-f329cca853d4', 'A', '菜单详情', '', 0, 'admin:system:menu:detail', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7015-bc54-638c532c8439', '019ce5eb-2ff4-7013-9081-f329cca853d4', 'A', '菜单新增', '', 0, 'admin:system:menu:add', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7016-9a6d-410566a343d0', '019ce5eb-2ff4-7013-9081-f329cca853d4', 'A', '菜单编辑', '', 0, 'admin:system:menu:edit', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7017-acea-b202413a5b6f', '019ce5eb-2ff4-7013-9081-f329cca853d4', 'A', '菜单删除', '', 0, 'admin:system:menu:del', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7018-a11a-261f7666a756', '019ce5eb-2ff4-702c-b7a4-5ced23241090', 'C', '部门管理', 'el-icon-Coordinate', 8, 'admin:system:dept:all', 'dept', 'system/dept/index', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7019-94da-4620361c2f80', '019ce5eb-2ff4-7018-a11a-261f7666a756', 'A', '部门详情', '', 0, 'admin:system:dept:detail', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-701a-8a5f-aed6841aedee', '019ce5eb-2ff4-7018-a11a-261f7666a756', 'A', '部门新增', '', 0, 'admin:system:dept:add', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-701b-9ec2-08b1723c8f94', '019ce5eb-2ff4-7018-a11a-261f7666a756', 'A', '部门编辑', '', 0, 'admin:system:dept:edit', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-701c-aa4b-454ff0dbbbbe', '019ce5eb-2ff4-7018-a11a-261f7666a756', 'A', '部门删除', '', 0, 'admin:system:dept:del', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-701d-94d5-942851553293', '019ce5eb-2ff4-702c-b7a4-5ced23241090', 'C', '岗位管理', 'el-icon-PriceTag', 7, 'admin:system:post:list', 'post', 'system/post/index', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-701e-96a8-77399dca62d8', '019ce5eb-2ff4-701d-94d5-942851553293', 'A', '岗位详情', '', 0, 'admin:system:post:detail', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-701f-ac43-b9e3caf48a00', '019ce5eb-2ff4-701d-94d5-942851553293', 'A', '岗位新增', '', 0, 'admin:system:post:add', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7020-a977-fe50308cb050', '019ce5eb-2ff4-701d-94d5-942851553293', 'A', '岗位编辑', '', 0, 'admin:system:post:edit', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7021-a5c4-aa0223b40e1d', '019ce5eb-2ff4-701d-94d5-942851553293', 'A', '岗位删除', '', 0, 'admin:system:post:del', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7022-8cf6-824b19fbad9b', '', 'M', '其它管理', '', 0, '', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7023-a0e3-e68d7027713a', '019ce5eb-2ff4-7022-8cf6-824b19fbad9b', 'M', '图库管理', '', 0, '', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7024-bd00-810e268c0a09', '019ce5eb-2ff4-7023-a0e3-e68d7027713a', 'A', '文件列表', '', 0, 'admin:common:album:albumList', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7025-951c-efdc178f4f83', '019ce5eb-2ff4-7023-a0e3-e68d7027713a', 'A', '文件命名', '', 0, 'admin:common:album:albumRename', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7026-9485-8ed0fa735490', '019ce5eb-2ff4-7023-a0e3-e68d7027713a', 'A', '文件移动', '', 0, 'admin:common:album:albumMove', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7027-b942-ec388c3c037f', '019ce5eb-2ff4-7023-a0e3-e68d7027713a', 'A', '文件删除', '', 0, 'admin:common:album:albumDel', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7028-a665-e8feeb8677d4', '019ce5eb-2ff4-7023-a0e3-e68d7027713a', 'A', '分类列表', '', 0, 'admin:common:album:cateList', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7029-91a2-2f2f861b9f69', '019ce5eb-2ff4-7023-a0e3-e68d7027713a', 'A', '分类新增', '', 0, 'admin:common:album:cateAdd', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-702a-bfcc-9ae35237f87d', '019ce5eb-2ff4-7023-a0e3-e68d7027713a', 'A', '分类命名', '', 0, 'admin:common:album:cateRename', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-702b-a401-72c5679ea1ba', '019ce5eb-2ff4-7023-a0e3-e68d7027713a', 'A', '分类删除', '', 0, 'admin:common:album:cateDel', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-702c-b7a4-5ced23241090', '', 'M', '系统设置', 'el-icon-Setting', 0, '', 'system', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-702d-af6d-3aaa056c89f0', '019ce5eb-2ff4-702c-b7a4-5ced23241090', 'M', '网站设置', 'el-icon-Basketball', 0, '', 'website', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-702e-b017-86fa59d4df22', '019ce5eb-2ff4-702d-af6d-3aaa056c89f0', 'C', '网站信息', '', 0, 'admin:setting:website:detail', 'information', 'system/website/information', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-702f-be79-838649533d1c', '019ce5eb-2ff4-702e-b017-86fa59d4df22', 'A', '保存配置', '', 0, 'admin:setting:website:save', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7030-9943-039710e189e7', '019ce5eb-2ff4-702d-af6d-3aaa056c89f0', 'C', '网站备案', '', 0, 'admin:setting:copyright:detail', 'filing', 'system/website/filing', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7031-a6c2-5ab1516b3b59', '019ce5eb-2ff4-7030-9943-039710e189e7', 'A', '备案保存', '', 0, 'admin:setting:copyright:save', '', 'setting/website/protocol', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7032-a90f-7951e72cbf6b', '019ce5eb-2ff4-702c-b7a4-5ced23241090', 'C', '字典管理', 'el-icon-Box', 6, 'admin:setting:dict:type:list', 'dict', 'system/dict/type/index', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7033-a584-3aeee8230a15', '019ce5eb-2ff4-7032-a90f-7951e72cbf6b', 'A', '字典类型新增', '', 0, 'admin:setting:dict:type:add', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7034-99f8-b986cfeb3310', '019ce5eb-2ff4-7032-a90f-7951e72cbf6b', 'A', '字典类型编辑', '', 0, 'admin:setting:dict:type:edit', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7035-816d-b35131d9238f', '019ce5eb-2ff4-7032-a90f-7951e72cbf6b', 'A', '字典类型删除', '', 0, 'admin:setting:dict:type:del', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7036-9677-6180966ff9e3', '019ce5eb-2ff4-7032-a90f-7951e72cbf6b', 'A', '字典数据新增', '', 0, 'admin:setting:dict:data:add', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7037-8241-1e0ce987af13', '019ce5eb-2ff4-7032-a90f-7951e72cbf6b', 'A', '字典数据编辑', '', 0, 'admin:setting:dict:data:edit', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7038-8c61-4baef3ae40be', '019ce5eb-2ff4-7032-a90f-7951e72cbf6b', 'A', '字典数据删除', '', 0, 'admin:setting:dict:data:del', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7039-aa74-0654a9debe14', '019ce5eb-2ff4-702c-b7a4-5ced23241090', 'M', '系统维护', 'el-icon-SetUp', 0, '', 'system', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-703a-9bb6-8fd7e7b59265', '019ce5eb-2ff4-7039-aa74-0654a9debe14', 'C', '系统环境', '', 0, 'admin:monitor:server', 'environment', 'system/system/environment', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-703b-a0cb-c0f7f73125af', '019ce5eb-2ff4-7039-aa74-0654a9debe14', 'C', '系统缓存', '', 0, 'admin:monitor:cache', 'system/cache', 'system/system/cache', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-703c-acfd-51c3c6931aad', '019ce5eb-2ff4-7039-aa74-0654a9debe14', 'C', '系统日志', '', 0, 'admin:system:log:operate', 'journal', 'system/system/journal', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-703d-91fc-49816776ebc2', '', 'M', '开发工具', 'el-icon-EditPen', 0, '', 'dev_tools', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-703e-90c6-c72919dad93a', '019ce5eb-2ff4-703d-91fc-49816776ebc2', 'C', '代码生成器', '', 0, 'admin:gen:list', 'code', 'dev_tools/code/index', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-703f-9c4d-bc1bc9f3bacb', '019ce5eb-2ff4-703e-90c6-c72919dad93a', 'A', '导入数据表', '', 0, 'admin:gen:importTable', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7040-8826-0badcd7e7140', '019ce5eb-2ff4-703e-90c6-c72919dad93a', 'A', '生成代码', '', 0, 'admin:gen:genCode', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7041-a819-e73a8b611a0f', '019ce5eb-2ff4-703e-90c6-c72919dad93a', 'A', '下载代码', '', 0, 'admin:gen:downloadCode', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7042-8574-3b7309b272d6', '019ce5eb-2ff4-703e-90c6-c72919dad93a', 'A', '预览代码', '', 0, 'admin:gen:previewCode', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7043-ba57-0672c3925a3a', '019ce5eb-2ff4-703e-90c6-c72919dad93a', 'A', '同步表结构', '', 0, 'admin:gen:syncTable', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7044-ae7f-c7f5c46df45e', '019ce5eb-2ff4-703e-90c6-c72919dad93a', 'A', '删除数据表', '', 0, 'admin:gen:delTable', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7045-940f-b6e7408fe1a7', '019ce5eb-2ff4-703e-90c6-c72919dad93a', 'A', '数据表详情', '', 0, 'admin:gen:detail', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7046-a60d-9fb5c79c8419', '', 'C', '素材中心', 'el-icon-PictureRounded', 44, '', 'material/index', 'material/index', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7047-97d9-5544db8fbdb0', '019ce5eb-2ff4-703d-91fc-49816776ebc2', 'C', '代码生成器编辑', 'el-icon-EditPen', 0, 'admin:gen:editTable', 'code/edit', 'dev_tools/code/edit', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7048-95e9-0ebc5c2c9e1f', '019ce5eb-2ff4-704a-9b92-e6758e30020c', 'C', '流程模板', '', 0, 'admin:flow:flow_template:list', 'flow_template/index', 'flow/flow_template/index', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7049-9a86-4165a60b86ce', '019ce5eb-2ff4-704a-9b92-e6758e30020c', 'C', '我的流程', '', 0, '', 'flow_apply/index', 'flow/flow_apply/index', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-704a-9b92-e6758e30020c', '', 'M', '审批流', 'el-icon-Coordinate', 0, '', 'flow', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-704b-b6d8-1c5e85f7991f', '019ce5eb-2ff4-704a-9b92-e6758e30020c', 'C', '待处理', '', 0, 'admin:flow:flow_history:list', 'flow_history/todo', 'flow/flow_history/todo', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-704c-a884-f3c71f811e8d', '019ce5eb-2ff4-704a-9b92-e6758e30020c', 'C', '已处理', '', 0, 'admin:flow:flow_history:list', 'flow_history/done', 'flow/flow_history/done', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-704d-b054-c1c93cecc208', '019ce5eb-2ff4-704a-9b92-e6758e30020c', 'C', '已完成流程', '', 0, 'admin:flow:flow_history:list', 'flow_apply/finish', 'flow/flow_apply/finish', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-704e-be16-81a6a37b37ec', '019ce5eb-2ff4-787a-a1f3-260f06dccb3f', 'C', '项目', '', 0, 'admin:monitor_project:list', 'project/index', 'monitor/project/index', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-704f-9425-ccb9c15ea9ca', '019ce5eb-2ff4-787a-a1f3-260f06dccb3f', 'C', '用户端', '', 0, '', 'client/index', 'monitor/client/index', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7050-a30f-ff418ce9fc5a', '019ce5eb-2ff4-7058-beb1-9f6e1c2d442d', 'A', '错误收集error添加', '', 0, 'admin:monitor_web:add', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7051-87bb-f80a2e4a0ec9', '019ce5eb-2ff4-7058-beb1-9f6e1c2d442d', 'A', '错误收集error编辑', '', 0, 'admin:monitor_web:edit', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7052-9964-22a83c0c53eb', '019ce5eb-2ff4-7058-beb1-9f6e1c2d442d', 'A', '错误收集error删除', '', 0, 'admin:monitor_web:del', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7053-8fe7-082967e1ee31', '019ce5eb-2ff4-7058-beb1-9f6e1c2d442d', 'A', '错误收集error列表', '', 0, 'admin:monitor_web:list', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7054-be98-6ed7beccaf29', '019ce5eb-2ff4-7058-beb1-9f6e1c2d442d', 'A', '错误收集error全部列表', '', 0, 'admin:monitor_web:listAll', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7055-9408-8ef856ecaa19', '019ce5eb-2ff4-7058-beb1-9f6e1c2d442d', 'A', '错误收集error详情', '', 0, 'admin:monitor_web:detail', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7056-a809-fd2f1c890c70', '019ce5eb-2ff4-7058-beb1-9f6e1c2d442d', 'A', '错误收集error导出excel', '', 0, 'admin:monitor_web:ExportFile', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7057-8e27-37c0577dd656', '019ce5eb-2ff4-7058-beb1-9f6e1c2d442d', 'A', '错误收集error导入excel', '', 0, 'admin:monitor_web:ImportFile', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7058-beb1-9f6e1c2d442d', '019ce5eb-2ff4-787a-a1f3-260f06dccb3f', 'C', '错误列表', '', 0, '', 'error/index', 'monitor/error/index', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7059-a9ae-820badcd1d2b', '019ce5eb-2ff4-7048-95e9-0ebc5c2c9e1f', 'A', '流程模板添加', '', 0, 'admin:flow:flow_template:add', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-705a-acfb-67e94904a749', '019ce5eb-2ff4-7048-95e9-0ebc5c2c9e1f', 'A', '流程模板编辑', '', 0, 'admin:flow:flow_template:edit', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-705b-9c11-d7e6238aaa44', '019ce5eb-2ff4-7048-95e9-0ebc5c2c9e1f', 'A', '流程模板删除', '', 0, 'admin:flow:flow_template:del', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-705c-bc05-6e8fc7daccf9', '019ce5eb-2ff4-7048-95e9-0ebc5c2c9e1f', 'A', '流程模板列表', '', 0, 'admin:flow:flow_template:list', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-705d-b762-f2c42925f3a7', '019ce5eb-2ff4-7048-95e9-0ebc5c2c9e1f', 'A', '流程模板全部列表', '', 0, 'admin:flow:flow_template:listAll', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-705e-9ff2-d8e5e6f683e7', '019ce5eb-2ff4-7048-95e9-0ebc5c2c9e1f', 'A', '流程模板详情', '', 0, 'admin:flow:flow_template:detail', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-705f-8571-a86f8d1c1b5e', '019ce5eb-2ff4-7049-9a86-4165a60b86ce', 'A', '流程详情', '', 0, 'admin:flow:flow_apply:detail', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7060-930d-3c9282a4be74', '019ce5eb-2ff4-7049-9a86-4165a60b86ce', 'A', '流程添加', '', 0, 'admin:flow:flow_apply:add', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7061-9fe8-0b6e8666ddce', '019ce5eb-2ff4-7049-9a86-4165a60b86ce', 'A', '流程编辑', '', 0, 'admin:flow:flow_apply:edit', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7062-9209-87da1b0ce74a', '019ce5eb-2ff4-7049-9a86-4165a60b86ce', 'A', '流程删除', '', 0, 'admin:flow:flow_apply:del', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7063-8277-9b9a738ba5f4', '019ce5eb-2ff4-7049-9a86-4165a60b86ce', 'A', '全部流程', '', 0, 'admin:flow:flow_history:listAll', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7064-afa7-131504da5a75', '019ce5eb-2ff4-7049-9a86-4165a60b86ce', 'A', '通过流程', '', 0, 'admin:flow:flow_history:pass', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7065-9051-3ddcf43feeed', '019ce5eb-2ff4-7049-9a86-4165a60b86ce', 'A', '拒接流程', '', 0, 'admin:flow:flow_history:back', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7066-a125-2a2a2e0dfef6', '019ce5eb-2ff4-7049-9a86-4165a60b86ce', 'A', '下一个流程', '', 0, 'admin:flow:flow_history:next_node', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7067-a81a-65eadb4f0d03', '019ce5eb-2ff4-7049-9a86-4165a60b86ce', 'A', '获取审批人', '', 0, 'admin:flow:flow_history:get_approver', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7068-af86-2ae2e80b36b8', '019ce5eb-2ff4-7049-9a86-4165a60b86ce', 'A', '流程列表', '', 0, 'admin:flow:flow_apply:list', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7069-b472-558305293aae', '019ce5eb-2ff4-704b-b6d8-1c5e85f7991f', 'A', '审批记录列表', '', 0, 'admin:flow:flow_history:list', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-706a-8a81-8efba30d1d02', '019ce5eb-2ff4-704b-b6d8-1c5e85f7991f', 'A', '审批记录详情', '', 0, 'admin:flow:flow_history:detail', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-706b-ad9d-22f1ecd14f3f', '019ce5eb-2ff4-704e-be16-81a6a37b37ec', 'A', '监控项目添加', '', 0, 'admin:monitor_project:add', '', '', '', '', 0, 1, 0, '2024-06-17 19:26:13', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-706c-b656-8811bcb8c98b', '019ce5eb-2ff4-704e-be16-81a6a37b37ec', 'A', '监控项目编辑', '', 0, 'admin:monitor_project:edit', '', '', '', '', 0, 1, 0, '2024-06-17 19:26:13', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-706d-bdb4-231e14406cc6', '019ce5eb-2ff4-704e-be16-81a6a37b37ec', 'A', '监控项目删除', '', 0, 'admin:monitor_project:del', '', '', '', '', 0, 1, 0, '2024-06-17 19:26:13', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-706e-b332-8531ef43ec21', '019ce5eb-2ff4-704e-be16-81a6a37b37ec', 'A', '监控项目列表', '', 0, 'admin:monitor_project:list', '', '', '', '', 0, 1, 0, '2024-06-17 19:26:13', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-706f-8231-e354f3fedc43', '019ce5eb-2ff4-704e-be16-81a6a37b37ec', 'A', '监控项目全部列表', '', 0, 'admin:monitor_project:listAll', '', '', '', '', 0, 1, 0, '2024-06-17 19:26:13', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7070-a51b-4b06026046da', '019ce5eb-2ff4-704e-be16-81a6a37b37ec', 'A', '监控项目详情', '', 0, 'admin:monitor_project:detail', '', '', '', '', 0, 1, 0, '2024-06-17 19:26:13', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7071-b33a-1fd653485aaa', '019ce5eb-2ff4-704e-be16-81a6a37b37ec', 'A', '监控项目导出excel', '', 0, 'admin:monitor_project:ExportFile', '', '', '', '', 0, 1, 0, '2024-06-17 19:26:13', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7072-a15d-5bd0c408fd4e', '019ce5eb-2ff4-704e-be16-81a6a37b37ec', 'A', '监控项目导入excel', '', 0, 'admin:monitor_project:ImportFile', '', '', '', '', 0, 1, 0, '2024-06-17 19:26:13', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7073-ac0f-5319fa644d6f', '019ce5eb-2ff4-703d-91fc-49816776ebc2', 'C', '开发文档', '', 0, '', 'https://adtkcn.github.io/x_admin/', '', '', 'a=1', 1, 1, 0, '2024-06-28 17:09:17', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7074-9803-09dc3e1bfa14', '', 'C', '用户协议', 'el-icon-Coordinate', 0, '', 'user/protocol/index', 'user/protocol/index', '', '', 0, 1, 0, '2024-09-10 20:03:46', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7075-a471-a6c3a568b7a2', '019ce5eb-2ff4-7074-9803-09dc3e1bfa14', 'A', '用户协议添加', '', 0, 'admin:user_protocol:add', '', '', '', '', 0, 1, 0, '2024-09-10 20:05:13', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7076-83a1-54f5dde120d0', '019ce5eb-2ff4-7074-9803-09dc3e1bfa14', 'A', '用户协议编辑', '', 0, 'admin:user_protocol:edit', '', '', '', '', 0, 1, 0, '2024-09-10 20:05:13', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7077-82f8-d08be2d7b56c', '019ce5eb-2ff4-7074-9803-09dc3e1bfa14', 'A', '用户协议删除', '', 0, 'admin:user_protocol:del', '', '', '', '', 0, 1, 0, '2024-09-10 20:05:13', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7875-bcf2-a081f814ce3a', '019ce5eb-2ff4-7074-9803-09dc3e1bfa14', 'A', '用户协议列表', '', 0, 'admin:user_protocol:list', '', '', '', '', 0, 1, 0, '2024-09-10 20:05:13', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7876-b02f-a0ac8db35b8d', '019ce5eb-2ff4-7074-9803-09dc3e1bfa14', 'A', '用户协议全部列表', '', 0, 'admin:user_protocol:listAll', '', '', '', '', 0, 1, 0, '2024-09-10 20:05:13', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7877-bdfc-255d59bb0d4d', '019ce5eb-2ff4-7074-9803-09dc3e1bfa14', 'A', '用户协议详情', '', 0, 'admin:user_protocol:detail', '', '', '', '', 0, 1, 0, '2024-09-10 20:05:13', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7878-8461-51e42ff9ab89', '019ce5eb-2ff4-7074-9803-09dc3e1bfa14', 'A', '用户协议导出excel', '', 0, 'admin:user_protocol:ExportFile', '', '', '', '', 0, 1, 0, '2024-09-10 20:05:13', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7879-9f4f-7cc56e5180e4', '019ce5eb-2ff4-7074-9803-09dc3e1bfa14', 'A', '用户协议导入excel', '', 0, 'admin:user_protocol:ImportFile', '', '', '', '', 0, 1, 0, '2024-09-10 20:05:13', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-787a-a1f3-260f06dccb3f', '', 'M', '错误监控', 'el-icon-Memo', 0, '', 'monitor', '', '', '', 1, 1, 0, '2024-10-29 15:32:33', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-787b-b112-f60e4bfc247d', '019ce5eb-2ff4-703d-91fc-49816776ebc2', 'C', '错误捕获', '', 0, '', 'test', 'dev_tools/test', '', '', 1, 1, 0, '2024-10-30 18:16:56', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-787c-b1be-b2f73bb41208', '019ce5eb-2ff4-704f-9425-ccb9c15ea9ca', 'A', '监控-客户端信息添加', '', 0, 'admin:monitor_client:add', '', '', '', '', 0, 1, 0, '2024-11-08 14:15:28', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-787d-b343-3e71b3b093cd', '019ce5eb-2ff4-704f-9425-ccb9c15ea9ca', 'A', '监控-客户端信息编辑', '', 0, 'admin:monitor_client:edit', '', '', '', '', 0, 1, 0, '2024-11-08 14:15:28', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-787e-ab57-24e83c88cf3c', '019ce5eb-2ff4-704f-9425-ccb9c15ea9ca', 'A', '监控-客户端信息删除', '', 0, 'admin:monitor_client:del', '', '', '', '', 0, 1, 0, '2024-11-08 14:15:28', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-787f-8ac6-06e5744c2a3c', '019ce5eb-2ff4-704f-9425-ccb9c15ea9ca', 'A', '监控-客户端信息删除-批量', '', 0, 'admin:monitor_client:delBatch', '', '', '', '', 0, 1, 0, '2024-11-08 14:15:28', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7880-b4ec-1dca66823b36', '019ce5eb-2ff4-704f-9425-ccb9c15ea9ca', 'A', '监控-客户端信息列表', '', 0, 'admin:monitor_client:list', '', '', '', '', 0, 1, 0, '2024-11-08 14:15:28', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7881-87fa-ec74b915bf78', '019ce5eb-2ff4-704f-9425-ccb9c15ea9ca', 'A', '监控-客户端信息全部列表', '', 0, 'admin:monitor_client:listAll', '', '', '', '', 0, 1, 0, '2024-11-08 14:15:28', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7882-9f80-051f5094cf44', '019ce5eb-2ff4-704f-9425-ccb9c15ea9ca', 'A', '监控-客户端信息详情', '', 0, 'admin:monitor_client:detail', '', '', '', '', 0, 1, 0, '2024-11-08 14:15:28', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7883-9bbf-e1ff3021210e', '019ce5eb-2ff4-704f-9425-ccb9c15ea9ca', 'A', '监控-客户端信息导出excel', '', 0, 'admin:monitor_client:ExportFile', '', '', '', '', 0, 1, 0, '2024-11-08 14:15:28', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7884-b460-932a4c6a28eb', '019ce5eb-2ff4-704f-9425-ccb9c15ea9ca', 'A', '监控-客户端信息导入excel', '', 0, 'admin:monitor_client:ImportFile', '', '', '', '', 0, 1, 0, '2024-11-08 14:15:28', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7885-88fa-dc8df33c6016', '019ce5eb-2ff4-703d-91fc-49816776ebc2', 'C', '文件分片上传', '', 0, '', 'uploadChunk', 'dev_tools/uploadChunk', '', '', 1, 1, 0, '2025-07-19 18:37:48', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7886-a509-fdcb98731a88', '019ce5eb-2ff4-703d-91fc-49816776ebc2', 'C', '接口文档', '', 0, '', 'apiDocs', 'IframeComponent', '', '{"url": "/api/static/api/index.html"}', 1, 1, 0, '2025-12-01 00:13:54', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7887-8030-9434abdf59d4', '019ce5eb-2ff4-702c-b7a4-5ced23241090', 'C', '定时任务', 'el-icon-AlarmClock', 0, '', 'system/corn/index', 'system/corn/index', '', '', 0, 1, 0, '2025-12-21 18:07:03', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7888-af98-ea8f5ae8ae73', '019ce5eb-2ff4-7887-8030-9434abdf59d4', 'A', '定时任务列表', '', 0, 'admin:system_corn:list', '', '', '', '', 0, 1, 0, '2025-12-21 18:07:03', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7889-86e7-6c66544f3afb', '019ce5eb-2ff4-7887-8030-9434abdf59d4', 'A', '定时任务全部列表', '', 0, 'admin:system_corn:listAll', '', '', '', '', 0, 1, 0, '2025-12-21 18:07:03', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-788a-a8d6-218425893903', '019ce5eb-2ff4-7887-8030-9434abdf59d4', 'A', '定时任务添加', '', 0, 'admin:system_corn:add', '', '', '', '', 0, 1, 0, '2025-12-21 18:07:03', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-788b-9db3-36db2b18b47f', '019ce5eb-2ff4-7887-8030-9434abdf59d4', 'A', '定时任务编辑', '', 0, 'admin:system_corn:edit', '', '', '', '', 0, 1, 0, '2025-12-21 18:07:03', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-788c-a00c-f51525910432', '019ce5eb-2ff4-7887-8030-9434abdf59d4', 'A', '定时任务删除', '', 0, 'admin:system_corn:del', '', '', '', '', 0, 1, 0, '2025-12-21 18:07:03', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-788d-a814-812da1f5a963', '019ce5eb-2ff4-7887-8030-9434abdf59d4', 'A', '定时任务删除-批量', '', 0, 'admin:system_corn:delBatch', '', '', '', '', 0, 1, 0, '2025-12-21 18:07:03', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-788e-b1fa-649242c6ae6f', '019ce5eb-2ff4-7887-8030-9434abdf59d4', 'A', '定时任务详情', '', 0, 'admin:system_corn:detail', '', '', '', '', 0, 1, 0, '2025-12-21 18:07:03', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-788f-b821-c8566fbbfade', '019ce5eb-2ff4-7887-8030-9434abdf59d4', 'A', '定时任务导出excel', '', 0, 'admin:system_corn:ExportFile', '', '', '', '', 0, 1, 0, '2025-12-21 18:07:03', '2026-03-13 14:38:42'),
	('019ce5eb-2ff4-7890-838a-398223a07c8a', '019ce5eb-2ff4-7887-8030-9434abdf59d4', 'A', '定时任务导入excel', '', 0, 'admin:system_corn:ImportFile', '', '', '', '', 0, 1, 0, '2025-12-21 18:07:03', '2026-03-13 14:38:42');

-- 导出  表 x_admin_2.x_system_auth_perm 结构
CREATE TABLE IF NOT EXISTS `x_system_auth_perm` (
  `id` char(36) NOT NULL COMMENT '主键',
  `role_id` char(36) NOT NULL COMMENT '角色ID',
  `menu_id` char(36) NOT NULL COMMENT '菜单ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='系统角色菜单表';

-- 正在导出表  x_admin_2.x_system_auth_perm 的数据：~0 rows (大约)

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

-- 正在导出表  x_admin_2.x_system_auth_post 的数据：~0 rows (大约)

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
INSERT INTO `x_system_auth_role` (`id`, `name`, `remark`, `sort`, `is_disable`, `create_time`, `update_time`) VALUES
	('019cdd61-db45-7333-873d-cfd9ea9713c3', '系统管理', '系统管理', 0, 0, '2026-03-11 22:51:43', '2026-03-12 21:55:13');

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

-- 正在导出表  x_admin_2.x_system_config 的数据：~7 rows (大约)
INSERT INTO `x_system_config` (`id`, `type`, `name`, `value`, `create_time`, `update_time`) VALUES
	('10', 'website', 'name', 'x_admin开源系统', '2024-01-02 03:04:05', '2025-06-24 19:52:37'),
	('11', 'website', 'logo', '/api/static/backend_logo.png', '2024-01-02 03:04:05', '2025-06-24 19:52:37'),
	('12', 'website', 'favicon', '/api/static/backend_favicon.ico', '2024-01-02 03:04:05', '2025-06-24 19:52:38'),
	('13', 'website', 'backdrop', '/api/static/backend_backdrop.png', '2024-01-02 03:04:05', '2025-06-24 19:52:38'),
	('14', 'website', 'copyright', '[{"name":"蜀ICP备15007060号-1","link":"http://www.beian.gov.cn"},{"name":"x_admin","link":"http://x.adtk.cn"}]', '2024-01-02 03:04:05', '2025-12-19 00:27:43'),
	('15', 'website', 'shopName', 'x_admin开源管理系统', '2024-01-02 03:04:05', '2025-06-24 19:52:38'),
	('16', 'website', 'shopLogo', '/image/20241705/f0eb36d508834bc2ac1b8c591c563efa.png', '2024-01-02 03:04:05', '2025-06-24 19:52:38');

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

-- 正在导出表  x_admin_2.x_system_corn 的数据：~1 rows (大约)
INSERT INTO `x_system_corn` (`id`, `task_name`, `task_code`, `corn_expr`, `status`, `created_by`, `is_delete`, `create_time`, `update_time`, `delete_time`) VALUES
	('019b4069-06de-78d3-8737-4d5bcf424988', 'a2', 'exampleTask', '*/5 * * * * *', 1, '1', 0, '2025-12-21 18:16:23', '2026-01-25 23:22:16', NULL);

-- 导出  表 x_admin_2.x_system_log_login 结构
CREATE TABLE IF NOT EXISTS `x_system_log_login` (
  `id` char(36) NOT NULL COMMENT '注解',
  `admin_id` char(36) NOT NULL COMMENT '管理员ID',
  `username` varchar(30) NOT NULL DEFAULT '' COMMENT '登录账号',
  `ip` varchar(39) NOT NULL COMMENT '登录地址',
  `os` varchar(100) NOT NULL DEFAULT '' COMMENT '操作系统',
  `browser` varchar(100) DEFAULT '' COMMENT '浏览器',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '操作状态: 1=成功, 2=失败',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='系统登录日志表';

-- 正在导出表  x_admin_2.x_system_log_login 的数据：~0 rows (大约)

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

-- 正在导出表  x_admin_2.x_system_log_operate 的数据：~0 rows (大约)

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
  `last_login_ip` varchar(39) NOT NULL DEFAULT '' COMMENT '最后登录IP',
  `last_login_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '最后登录时间',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='用户信息表';

-- 正在导出表  x_admin_2.x_user 的数据：~0 rows (大约)

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

-- 正在导出表  x_admin_2.x_user_protocol 的数据：9 rows
INSERT INTO `x_user_protocol` (`id`, `tag`, `version`, `title`, `content`, `created_by`, `is_delete`, `create_time`, `update_time`, `delete_time`) VALUES
	('019b35e3-6ca9-7b57-97f6-13aa822dab20', '1', 1, '1', '<p>1</p>', '1', 0, '2025-12-19 17:14:15', '2025-12-21 22:42:40', NULL),
	('019b35f2-268b-7854-a6d2-71e8750358cd', '2', 2, '2', '<p>2</p>', '019b360a-4bd9-718c-ba71-46429ffd8b53', 0, '2025-12-19 17:30:20', '2025-12-19 17:30:20', NULL),
	('019b37cc-3523-7e5b-a4a3-0161eec677a0', '1', NULL, '', '<p>1</p>', '1', 0, '2025-12-19 17:14:15', '2025-12-20 02:09:37', NULL),
	('019b37cc-3523-7e5c-9450-f7d2cb97b08b', '2', 2, '2', '<p>2</p>', '1', 1, '2025-12-19 17:30:20', '2025-12-20 02:08:13', '2026-03-11 19:17:09'),
	('019b4169-f2b4-74f8-8d00-95290f098670', '1', NULL, '', '<p>1</p>', '1', 0, '2025-12-19 17:14:15', '2026-03-11 17:45:07', NULL),
	('019b4169-f2b4-74f9-9b42-224cc3b4d00d', '', NULL, '', '<p>123</p>', '019b360a-4bd9-718c-ba71-46429ffd8b53', 1, '2025-12-19 17:30:20', '2025-12-19 17:30:20', '2026-03-30 13:25:36'),
	('019b4175-8fec-7766-b115-2bf93d2ddd23', '33', 33, '33', '<p>33</p>', '1', 1, '2025-12-19 17:14:15', '2025-12-19 17:14:15', '2026-03-11 18:06:02'),
	('019b4175-8fec-7767-9f86-490fd125a350', '2', NULL, '', '<p>2</p>', '1', 1, '2025-12-19 17:30:20', '2026-03-11 16:19:42', '2026-03-11 18:39:01'),
	('019d3d3a-b82d-7a4b-824b-82b89fcdb58a', '222', 222, '222', '<p>222</p>', '1', 0, '2026-03-30 13:32:31', '2026-03-30 13:32:31', NULL);

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
