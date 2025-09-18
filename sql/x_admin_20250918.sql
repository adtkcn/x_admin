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
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `cid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '类目ID',
  `admin_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '管理员ID',
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
) ENGINE=InnoDB AUTO_INCREMENT=64 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='相册管理表';

-- 正在导出表  x_admin_2.x_album 的数据：~0 rows (大约)

-- 导出  表 x_admin_2.x_album_cate 结构
CREATE TABLE IF NOT EXISTS `x_album_cate` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `pid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '父级ID',
  `admin_id` int(10) unsigned NOT NULL COMMENT '管理员id',
  `name` varchar(32) NOT NULL DEFAULT '' COMMENT '分类名称',
  `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除: [0=否, 1=是]',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='相册分类表';

-- 正在导出表  x_admin_2.x_album_cate 的数据：~0 rows (大约)

-- 导出  表 x_admin_2.x_article 结构
CREATE TABLE IF NOT EXISTS `x_article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `cid` int(10) unsigned NOT NULL COMMENT '分类',
  `title` varchar(200) NOT NULL DEFAULT '' COMMENT '标题',
  `intro` varchar(200) NOT NULL DEFAULT '' COMMENT '简介',
  `summary` varchar(200) DEFAULT '' COMMENT '摘要',
  `image` varchar(200) NOT NULL DEFAULT '' COMMENT '封面',
  `content` text COMMENT '内容',
  `author` varchar(32) NOT NULL DEFAULT '' COMMENT '作者',
  `visit` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '浏览',
  `sort` int(10) unsigned NOT NULL DEFAULT '50' COMMENT '排序',
  `is_show` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否显示: 0=否, 1=是',
  `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除: 0=否, 1=是',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `cid_idx` (`cid`) USING BTREE COMMENT '分类索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='文章资讯表';

-- 正在导出表  x_admin_2.x_article 的数据：~0 rows (大约)

-- 导出  表 x_admin_2.x_article_category 结构
CREATE TABLE IF NOT EXISTS `x_article_category` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(60) NOT NULL DEFAULT '' COMMENT '名称',
  `sort` smallint(5) unsigned NOT NULL DEFAULT '50' COMMENT '排序',
  `is_show` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '是否显示: 0=否, 1=是',
  `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除: 0=否, 1=是',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='文章分类表';

-- 正在导出表  x_admin_2.x_article_category 的数据：~0 rows (大约)

-- 导出  表 x_admin_2.x_article_collect 结构
CREATE TABLE IF NOT EXISTS `x_article_collect` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `article_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '文章ID',
  `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='文章收藏表';

-- 正在导出表  x_admin_2.x_article_collect 的数据：~0 rows (大约)

-- 导出  表 x_admin_2.x_dict_data 结构
CREATE TABLE IF NOT EXISTS `x_dict_data` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `type_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '类型',
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
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='字典数据表';

-- 正在导出表  x_admin_2.x_dict_data 的数据：~20 rows (大约)
REPLACE INTO `x_dict_data` (`id`, `type_id`, `name`, `value`, `color`, `remark`, `sort`, `status`, `is_delete`, `create_time`, `update_time`, `delete_time`) VALUES
	(1, 2, '待提交', '1', '#6D85FC', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL),
	(2, 2, '审批中', '2', '#C6C150', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL),
	(3, 2, '审批成功', '3', 'green', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL),
	(4, 2, '失败', '4', 'red', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL),
	(5, 3, '待处理', '1', '#087BF6', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL),
	(6, 3, '通过', '2', 'green', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL),
	(7, 3, '拒绝', '3', 'red', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL),
	(10, 4, '假勤管理', '1', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL),
	(11, 4, '人事管理', '2', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL),
	(12, 4, '财务管理', '3', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL),
	(13, 4, '业务管理', '4', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL),
	(14, 4, '行政管理', '5', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL),
	(15, 4, '法务管理', '6', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL),
	(16, 4, '其他', '7', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL),
	(17, 5, 'web', 'web', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-06-29 00:18:02', NULL),
	(18, 5, 'go', 'go', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL),
	(19, 5, 'uniapp', 'uniapp', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL),
	(20, 5, 'node', 'node', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL),
	(22, 6, '禁用', '0', '#FF4B4B', '', 0, 1, 0, '2024-09-25 16:02:08', '2024-09-25 16:02:08', NULL),
	(23, 6, '启用', '1', '#80D251', '', 0, 1, 0, '2024-09-25 16:02:30', '2024-09-25 16:02:30', NULL);

-- 导出  表 x_admin_2.x_dict_type 结构
CREATE TABLE IF NOT EXISTS `x_dict_type` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `dict_name` varchar(100) NOT NULL DEFAULT '' COMMENT '字典名称',
  `dict_type` varchar(100) NOT NULL DEFAULT '' COMMENT '字典类型',
  `dict_remark` varchar(200) NOT NULL DEFAULT '' COMMENT '字典备注',
  `dict_status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '字典状态: 0=停用, 1=正常',
  `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除: 0=否, 1=是',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='字典类型表';

-- 正在导出表  x_admin_2.x_dict_type 的数据：~5 rows (大约)
REPLACE INTO `x_dict_type` (`id`, `dict_name`, `dict_type`, `dict_remark`, `dict_status`, `is_delete`, `create_time`, `update_time`, `delete_time`) VALUES
	(2, '审批申请状态', 'flow_apply_status', '0待提交，1审批中，2审批完成，3审批失败', 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL),
	(3, '审批历史状态', 'flow_history_status', '', 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL),
	(4, '流程分类', 'flow_group', '1假勤管理,2人事管理3财务管理4业务管理5行政管理6法务管理7其他', 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL),
	(5, '项目类型', 'project_type', '项目类型go java web node php 等', 1, 0, '2024-01-02 03:04:05', '2024-06-29 00:48:34', NULL),
	(6, '启用状态', 'status', ' 0=否, 1=是', 1, 0, '2024-09-25 15:54:56', '2024-09-25 15:54:56', NULL);

-- 导出  表 x_admin_2.x_flow_apply 结构
CREATE TABLE IF NOT EXISTS `x_flow_apply` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `template_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '模板',
  `apply_user_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '申请人id',
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
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='申请流程';

-- 正在导出表  x_admin_2.x_flow_apply 的数据：~1 rows (大约)
REPLACE INTO `x_flow_apply` (`id`, `template_id`, `apply_user_id`, `apply_user_nickname`, `flow_name`, `flow_group`, `flow_remark`, `flow_form_data`, `flow_process_data`, `flow_process_data_list`, `form_value`, `status`, `is_delete`, `create_time`, `update_time`, `delete_time`) VALUES
	(1, 1, 1, 'admin', 'a', 1, 'a', '{"widgetList":[{"key":33700,"type":"input","icon":"text-field","formItemFlag":true,"options":{"name":"input97761","label":"input","labelAlign":"","type":"text","defaultValue":"","placeholder":"","columnWidth":"200px","size":"","labelWidth":null,"labelHidden":false,"readonly":false,"disabled":false,"hidden":false,"clearable":true,"showPassword":false,"required":false,"requiredHint":"","validation":"","validationHint":"","customClass":"","labelIconClass":null,"labelIconPosition":"rear","labelTooltip":null,"minLength":null,"maxLength":null,"showWordLimit":false,"prefixIcon":"","suffixIcon":"","appendButton":false,"appendButtonDisabled":false,"buttonIcon":"custom-search","onCreated":"","onMounted":"","onInput":"","onChange":"","onFocus":"","onBlur":"","onValidate":"","onAppendButtonClick":""},"id":"input97761"}],"formConfig":{"modelName":"formData","refName":"vForm","rulesName":"rules","labelWidth":80,"labelPosition":"left","size":"","labelAlign":"label-left-align","cssCode":"","customClass":"","functions":"","layoutType":"PC","jsonVersion":3,"onFormCreated":"","onFormMounted":"","onFormDataChange":""}}', '{"nodes":[{"id":"Event_a677124","type":"bpmn:startEvent","x":300,"y":100,"properties":{"width":36,"height":36},"zIndex":1013,"text":{"x":300,"y":140,"value":"开始"}},{"id":"Activity_2c996a7","type":"bpmn:userTask","x":500,"y":120,"properties":{"width":100,"height":80},"zIndex":1015,"text":{"x":500,"y":120,"value":"审批"}},{"id":"Event_129de4e","type":"bpmn:endEvent","x":730,"y":130,"properties":{"width":36,"height":36},"zIndex":1012,"text":{"x":730,"y":170,"value":"结束"}}],"edges":[{"id":"2b28919f-30fe-4feb-9db9-0863359f87ed","type":"pro-polyline","properties":{},"sourceNodeId":"Event_a677124","targetNodeId":"Activity_2c996a7","startPoint":{"x":318,"y":100},"endPoint":{"x":450,"y":120},"zIndex":1014,"pointsList":[{"x":318,"y":100},{"x":384,"y":100},{"x":384,"y":120},{"x":450,"y":120}]},{"id":"dd95a2cd-2ecc-4e03-9d86-932673dd677b","type":"pro-polyline","properties":{},"sourceNodeId":"Activity_2c996a7","targetNodeId":"Event_129de4e","startPoint":{"x":550,"y":120},"endPoint":{"x":712,"y":130},"zIndex":1016,"pointsList":[{"x":550,"y":120},{"x":631,"y":120},{"x":631,"y":130},{"x":712,"y":130}]}]}', '[{"id":"Event_a677124","pid":0,"label":"开始","type":"bpmn:startEvent","userType":0,"userId":0,"deptId":0,"postId":0,"children":[{"id":"Activity_2c996a7","pid":"Event_a677124","label":"审批","type":"bpmn:userTask","userType":0,"userId":0,"deptId":0,"postId":0,"children":[{"id":"Event_129de4e","pid":"Activity_2c996a7","label":"结束","type":"bpmn:endEvent","userType":0,"userId":0,"deptId":0,"postId":0,"children":null}]}]},{"id":"Activity_2c996a7","pid":"Event_a677124","label":"审批","type":"bpmn:userTask","userType":0,"userId":0,"deptId":0,"postId":0,"children":[{"id":"Event_129de4e","pid":"Activity_2c996a7","label":"结束","type":"bpmn:endEvent","userType":0,"userId":0,"deptId":0,"postId":0,"children":null}]},{"id":"Event_129de4e","pid":"Activity_2c996a7","label":"结束","type":"bpmn:endEvent","userType":0,"userId":0,"deptId":0,"postId":0,"children":null},{"id":"Activity_2c996a7","pid":0,"label":"审批","type":"bpmn:userTask","userType":0,"userId":0,"deptId":0,"postId":0,"children":[{"id":"Event_129de4e","pid":"Activity_2c996a7","label":"结束","type":"bpmn:endEvent","userType":0,"userId":0,"deptId":0,"postId":0,"children":null}]},{"id":"Event_129de4e","pid":"Activity_2c996a7","label":"结束","type":"bpmn:endEvent","userType":0,"userId":0,"deptId":0,"postId":0,"children":null},{"id":"Event_129de4e","pid":0,"label":"结束","type":"bpmn:endEvent","userType":0,"userId":0,"deptId":0,"postId":0,"children":null}]', '{"input97761":"12"}', 3, 0, '2024-11-07 15:01:53', '2024-11-11 00:40:11', NULL);

-- 导出  表 x_admin_2.x_flow_history 结构
CREATE TABLE IF NOT EXISTS `x_flow_history` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '历史id',
  `apply_id` int(10) unsigned NOT NULL COMMENT '申请id',
  `template_id` int(10) unsigned DEFAULT NULL COMMENT '模板id',
  `apply_user_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '申请人id',
  `apply_user_nickname` varchar(32) NOT NULL DEFAULT '0' COMMENT '申请人昵称',
  `approver_id` int(10) unsigned NOT NULL COMMENT '审批人id',
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
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='流程历史';

-- 正在导出表  x_admin_2.x_flow_history 的数据：~3 rows (大约)
REPLACE INTO `x_flow_history` (`id`, `apply_id`, `template_id`, `apply_user_id`, `apply_user_nickname`, `approver_id`, `approver_nickname`, `node_id`, `node_type`, `node_label`, `form_value`, `pass_status`, `pass_remark`, `create_time`, `update_time`, `delete_time`) VALUES
	(1, 1, 1, 1, 'admin', 0, '', 'Event_a677124', 'bpmn:startEvent', '开始', '{"input97761":"12"}', 2, '', '2024-11-11 00:39:31', '2024-11-11 00:39:31', NULL),
	(2, 1, 1, 1, 'admin', 1, 'admin', 'Activity_2c996a7', 'bpmn:userTask', '审批', '{"input97761":"12"}', 2, '', '2024-11-11 00:39:31', '2024-11-11 00:40:11', NULL),
	(3, 1, 1, 1, 'admin', 0, '', 'Event_129de4e', 'bpmn:endEvent', '结束', '{"input97761":"12"}', 2, '', '2024-11-11 00:40:11', '2024-11-11 00:40:11', NULL);

-- 导出  表 x_admin_2.x_flow_template 结构
CREATE TABLE IF NOT EXISTS `x_flow_template` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
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
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='流程模板';

-- 正在导出表  x_admin_2.x_flow_template 的数据：~1 rows (大约)
REPLACE INTO `x_flow_template` (`id`, `flow_name`, `flow_group`, `flow_remark`, `flow_form_data`, `flow_process_data`, `flow_process_data_list`, `is_delete`, `create_time`, `update_time`, `delete_time`) VALUES
	(1, 'a', 1, 'a', '{"widgetList":[{"key":33700,"type":"input","icon":"text-field","formItemFlag":true,"options":{"name":"input97761","label":"input","labelAlign":"","type":"text","defaultValue":"","placeholder":"","columnWidth":"200px","size":"","labelWidth":null,"labelHidden":false,"readonly":false,"disabled":false,"hidden":false,"clearable":true,"showPassword":false,"required":false,"requiredHint":"","validation":"","validationHint":"","customClass":"","labelIconClass":null,"labelIconPosition":"rear","labelTooltip":null,"minLength":null,"maxLength":null,"showWordLimit":false,"prefixIcon":"","suffixIcon":"","appendButton":false,"appendButtonDisabled":false,"buttonIcon":"custom-search","onCreated":"","onMounted":"","onInput":"","onChange":"","onFocus":"","onBlur":"","onValidate":"","onAppendButtonClick":""},"id":"input97761"},{"key":60526,"type":"card","category":"container","icon":"card","widgetList":[],"options":{"name":"card21677","label":"card","hidden":false,"folded":false,"showFold":true,"cardWidth":"100%","shadow":"never","customClass":""},"id":"card21677"},{"key":27096,"type":"select","icon":"select-field","formItemFlag":true,"options":{"name":"select49056","label":"select","labelAlign":"","defaultValue":"","placeholder":"","columnWidth":"200px","size":"","labelWidth":null,"labelHidden":false,"disabled":false,"hidden":false,"clearable":true,"filterable":false,"allowCreate":false,"remote":false,"automaticDropdown":false,"multiple":false,"multipleLimit":0,"optionItems":[{"label":"select 1","value":1},{"label":"select 2","value":2},{"label":"select 3","value":3}],"required":false,"requiredHint":"","validation":"","validationHint":"","customClass":"","labelIconClass":null,"labelIconPosition":"rear","labelTooltip":null,"onCreated":"","onMounted":"","onRemoteQuery":"","onChange":"","onFocus":"","onBlur":"","onValidate":""},"id":"select49056"}],"formConfig":{"modelName":"formData","refName":"vForm","rulesName":"rules","labelWidth":80,"labelPosition":"left","size":"","labelAlign":"label-left-align","cssCode":"","customClass":"","functions":"","layoutType":"PC","jsonVersion":3,"onFormCreated":"","onFormMounted":"","onFormDataChange":""}}', '{"nodes":[{"id":"Event_a677124","type":"bpmn:startEvent","x":220,"y":198,"properties":{"width":36,"height":36},"zIndex":1001,"text":{"x":220,"y":238,"value":"开始"}},{"id":"Activity_2c996a7","type":"bpmn:userTask","x":490,"y":301,"properties":{"width":100,"height":80},"zIndex":1015,"text":{"x":490,"y":301,"value":"审批"}},{"id":"Event_129de4e","type":"bpmn:endEvent","x":730,"y":130,"properties":{"width":36,"height":36},"zIndex":1012,"text":{"x":730,"y":170,"value":"结束"}},{"id":"Activity_bbd8262","type":"bpmn:userTask","x":385,"y":146,"properties":{"width":100,"height":80},"zIndex":1005,"text":{"x":385,"y":146,"value":"审批"}}],"edges":[{"id":"2b28919f-30fe-4feb-9db9-0863359f87ed","type":"pro-polyline","properties":{},"sourceNodeId":"Event_a677124","targetNodeId":"Activity_2c996a7","sourceAnchorId":"Event_a677124_1","targetAnchorId":"Activity_2c996a7_3","startPoint":{"x":238,"y":198},"endPoint":{"x":440,"y":301},"zIndex":1014,"pointsList":[{"x":238,"y":198},{"x":268,"y":198},{"x":268,"y":301},{"x":440,"y":301}]},{"id":"dd95a2cd-2ecc-4e03-9d86-932673dd677b","type":"pro-polyline","properties":{},"sourceNodeId":"Activity_2c996a7","targetNodeId":"Event_129de4e","sourceAnchorId":"Activity_2c996a7_1","targetAnchorId":"Event_129de4e_3","startPoint":{"x":540,"y":301},"endPoint":{"x":712,"y":130},"zIndex":1016,"pointsList":[{"x":540,"y":301},{"x":682,"y":301},{"x":682,"y":130},{"x":712,"y":130}]}]}', '[{"id":"Event_a677124","pid":0,"label":"开始","type":"bpmn:startEvent","userType":0,"userId":0,"deptId":0,"postId":0,"children":[{"id":"Activity_2c996a7","pid":"Event_a677124","label":"审批","type":"bpmn:userTask","userType":0,"userId":0,"deptId":0,"postId":0,"children":[{"id":"Event_129de4e","pid":"Activity_2c996a7","label":"结束","type":"bpmn:endEvent","userType":0,"userId":0,"deptId":0,"postId":0,"children":null}]}]},{"id":"Activity_2c996a7","pid":"Event_a677124","label":"审批","type":"bpmn:userTask","userType":0,"userId":0,"deptId":0,"postId":0,"children":[{"id":"Event_129de4e","pid":"Activity_2c996a7","label":"结束","type":"bpmn:endEvent","userType":0,"userId":0,"deptId":0,"postId":0,"children":null}]},{"id":"Event_129de4e","pid":"Activity_2c996a7","label":"结束","type":"bpmn:endEvent","userType":0,"userId":0,"deptId":0,"postId":0,"children":null},{"id":"Activity_2c996a7","pid":0,"label":"审批","type":"bpmn:userTask","userType":0,"userId":0,"deptId":0,"postId":0,"children":[{"id":"Event_129de4e","pid":"Activity_2c996a7","label":"结束","type":"bpmn:endEvent","userType":0,"userId":0,"deptId":0,"postId":0,"children":null}]},{"id":"Event_129de4e","pid":"Activity_2c996a7","label":"结束","type":"bpmn:endEvent","userType":0,"userId":0,"deptId":0,"postId":0,"children":null},{"id":"Event_129de4e","pid":0,"label":"结束","type":"bpmn:endEvent","userType":0,"userId":0,"deptId":0,"postId":0,"children":null},{"id":"Activity_bbd8262","pid":0,"label":"审批","type":"bpmn:userTask","userType":0,"userId":0,"deptId":0,"postId":0,"children":null}]', 0, '2024-11-07 15:01:46', '2025-09-17 20:32:11', NULL);

-- 导出  表 x_admin_2.x_gen_table 结构
CREATE TABLE IF NOT EXISTS `x_gen_table` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
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
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='代码生成业务表';

-- 正在导出表  x_admin_2.x_gen_table 的数据：~0 rows (大约)

-- 导出  表 x_admin_2.x_gen_table_column 结构
CREATE TABLE IF NOT EXISTS `x_gen_table_column` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '列主键',
  `table_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '表外键',
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
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='代码生成字段表';

-- 正在导出表  x_admin_2.x_gen_table_column 的数据：~0 rows (大约)

-- 导出  表 x_admin_2.x_monitor_client 结构
CREATE TABLE IF NOT EXISTS `x_monitor_client` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'uuid',
  `project_key` varchar(128) NOT NULL COMMENT '项目key',
  `client_id` varchar(128) NOT NULL COMMENT 'sdk生成的客户端id',
  `user_id` varchar(128) DEFAULT NULL COMMENT '用户id',
  `os` varchar(30) DEFAULT NULL COMMENT '系统',
  `browser` varchar(30) DEFAULT NULL COMMENT '浏览器',
  `country` varchar(50) DEFAULT NULL COMMENT '国家',
  `province` varchar(50) DEFAULT NULL COMMENT '省份',
  `city` varchar(50) DEFAULT NULL COMMENT '城市',
  `operator` varchar(50) DEFAULT NULL COMMENT '电信运营商',
  `ip` varchar(50) DEFAULT NULL COMMENT 'ip',
  `ua` varchar(128) DEFAULT NULL COMMENT 'ua记录',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `project_key` (`project_key`) USING BTREE,
  KEY `user_id` (`user_id`) USING BTREE,
  KEY `client_id` (`client_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=119 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='监控-客户端信息';

-- 正在导出表  x_admin_2.x_monitor_client 的数据：~0 rows (大约)

-- 导出  表 x_admin_2.x_monitor_error 结构
CREATE TABLE IF NOT EXISTS `x_monitor_error` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '错误id',
  `project_key` varchar(128) NOT NULL COMMENT '项目key',
  `md5` varchar(32) DEFAULT NULL COMMENT 'md5',
  `event_type` varchar(20) DEFAULT NULL COMMENT '事件类型',
  `path` varchar(1000) DEFAULT NULL COMMENT 'URL地址',
  `message` text COMMENT '错误消息',
  `stack` text COMMENT '错误堆栈',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `project_key` (`project_key`,`md5`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=190 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='监控-错误列表';

-- 正在导出表  x_admin_2.x_monitor_error 的数据：~0 rows (大约)

-- 导出  表 x_admin_2.x_monitor_error_list 结构
CREATE TABLE IF NOT EXISTS `x_monitor_error_list` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `eid` varchar(32) CHARACTER SET utf8 NOT NULL COMMENT '错误表id',
  `cid` varchar(32) CHARACTER SET utf8 NOT NULL COMMENT '客户端表id',
  `width` smallint(10) unsigned DEFAULT '0' COMMENT '屏幕',
  `height` smallint(10) unsigned DEFAULT '0' COMMENT '屏幕高度',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `uid` (`cid`) USING BTREE,
  KEY `eid` (`eid`) USING BTREE
) ENGINE=MyISAM AUTO_INCREMENT=538 DEFAULT CHARSET=utf8mb4 COMMENT='错误对应的用户记录';

-- 正在导出表  x_admin_2.x_monitor_error_list 的数据：0 rows
/*!40000 ALTER TABLE `x_monitor_error_list` DISABLE KEYS */;
/*!40000 ALTER TABLE `x_monitor_error_list` ENABLE KEYS */;

-- 导出  表 x_admin_2.x_monitor_project 结构
CREATE TABLE IF NOT EXISTS `x_monitor_project` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '项目id',
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
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='监控项目';

-- 正在导出表  x_admin_2.x_monitor_project 的数据：~2 rows (大约)
REPLACE INTO `x_monitor_project` (`id`, `project_key`, `project_name`, `project_type`, `status`, `is_delete`, `create_time`, `update_time`, `delete_time`) VALUES
	(6, '6217ea4ea0044014831bd25121a3113c', 'go', 'go', 0, 0, '2024-07-12 23:17:23', '2024-11-07 15:43:08', NULL),
	(7, 'e19e3be20de94f49b68fafb4c30668bc', 'web项目', 'web', 1, 0, '2024-07-13 20:56:21', '2024-09-25 16:58:58', NULL);

-- 导出  表 x_admin_2.x_system_auth_admin 结构
CREATE TABLE IF NOT EXISTS `x_system_auth_admin` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `dept_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '部门ID',
  `post_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '岗位ID',
  `username` varchar(32) NOT NULL DEFAULT '' COMMENT '用户账号',
  `nickname` varchar(32) NOT NULL DEFAULT '' COMMENT '用户昵称',
  `password` varchar(200) NOT NULL DEFAULT '' COMMENT '用户密码',
  `avatar` varchar(200) NOT NULL DEFAULT '' COMMENT '用户头像',
  `role` varchar(200) NOT NULL DEFAULT '' COMMENT '角色主键',
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
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='系统管理成员表';

-- 正在导出表  x_admin_2.x_system_auth_admin 的数据：~1 rows (大约)
REPLACE INTO `x_system_auth_admin` (`id`, `dept_id`, `post_id`, `username`, `nickname`, `password`, `avatar`, `role`, `salt`, `sort`, `is_multipoint`, `is_disable`, `is_delete`, `last_login_ip`, `last_login_time`, `create_time`, `update_time`, `delete_time`) VALUES
	(1, 1, 3, 'admin', 'admin', '81a13dd8e25644a8823082573ca973f7', '/png/20250916/4d8d6a13a3034380b8dbf9b95591839c.png', '0', 'WFdiD', 1, 1, 0, 0, '127.0.0.1', '2025-09-18 12:09:51', '2024-01-02 03:04:05', '2025-09-18 12:09:51', NULL);

-- 导出  表 x_admin_2.x_system_auth_dept 结构
CREATE TABLE IF NOT EXISTS `x_system_auth_dept` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `pid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '上级主键',
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '部门名称',
  `duty_id` int(10) DEFAULT '0' COMMENT '负责人id',
  `duty` varchar(32) DEFAULT '' COMMENT '负责人名',
  `mobile` varchar(30) DEFAULT '' COMMENT '联系电话',
  `sort` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '排序编号',
  `is_stop` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否禁用: 0=否, 1=是',
  `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除: 0=否, 1=是',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='系统部门管理表';

-- 正在导出表  x_admin_2.x_system_auth_dept 的数据：~2 rows (大约)
REPLACE INTO `x_system_auth_dept` (`id`, `pid`, `name`, `duty_id`, `duty`, `mobile`, `sort`, `is_stop`, `is_delete`, `create_time`, `update_time`, `delete_time`) VALUES
	(1, 0, '默认部门', 1, 'admin', '18327647788', 10, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL),
	(2, 1, '指挥部', 2, '指挥部01', '17608390000', 3, 0, 0, '2024-01-02 03:04:05', '2024-07-05 15:18:25', NULL);

-- 导出  表 x_admin_2.x_system_auth_menu 结构
CREATE TABLE IF NOT EXISTS `x_system_auth_menu` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `pid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '上级菜单',
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
) ENGINE=InnoDB AUTO_INCREMENT=874 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='系统菜单管理表';

-- 正在导出表  x_admin_2.x_system_auth_menu 的数据：~133 rows (大约)
REPLACE INTO `x_system_auth_menu` (`id`, `pid`, `menu_type`, `menu_name`, `menu_icon`, `menu_sort`, `perms`, `paths`, `component`, `selected`, `params`, `is_cache`, `is_show`, `is_disable`, `create_time`, `update_time`) VALUES
	(1, 0, 'C', '工作台', 'el-icon-Monitor', 50, 'admin:common:index:console', 'workbench', 'workbench/index', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(101, 500, 'C', '管理员', 'local-icon-wode', 10, 'admin:system:admin:list', 'admin', 'system/admin/index', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2025-09-18 11:41:20'),
	(102, 101, 'A', '管理员详情', '', 0, 'admin:system:admin:detail', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(103, 101, 'A', '管理员新增', '', 0, 'admin:system:admin:add', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(104, 101, 'A', '管理员编辑', '', 0, 'admin:system:admin:edit', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(105, 101, 'A', '管理员删除', '', 0, 'admin:system:admin:del', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(106, 101, 'A', '管理员状态', '', 0, 'admin:system:admin:disable', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(110, 500, 'C', '角色管理', 'el-icon-Female', 9, 'admin:system:role:list', 'role', 'system/role/index', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2025-09-18 11:43:09'),
	(111, 110, 'A', '角色详情', '', 0, 'admin:system:role:detail', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(112, 110, 'A', '角色新增', '', 0, 'admin:system:role:add', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(113, 110, 'A', '角色编辑', '', 0, 'admin:system:role:edit', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(114, 110, 'A', '角色删除', '', 0, 'admin:system:role:del', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(120, 500, 'C', '菜单管理', 'el-icon-Operation', 0, 'admin:system:menu:list', 'menu', 'system/menu/index', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2025-09-18 11:39:11'),
	(121, 120, 'A', '菜单详情', '', 0, 'admin:system:menu:detail', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(122, 120, 'A', '菜单新增', '', 0, 'admin:system:menu:add', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(123, 120, 'A', '菜单编辑', '', 0, 'admin:system:menu:edit', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(124, 120, 'A', '菜单删除', '', 0, 'admin:system:menu:del', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(131, 500, 'C', '部门管理', 'el-icon-Coordinate', 8, 'admin:system:dept:all', 'dept', 'system/dept/index', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2025-09-18 12:17:11'),
	(132, 131, 'A', '部门详情', '', 0, 'admin:system:dept:detail', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(133, 131, 'A', '部门新增', '', 0, 'admin:system:dept:add', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(134, 131, 'A', '部门编辑', '', 0, 'admin:system:dept:edit', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(135, 131, 'A', '部门删除', '', 0, 'admin:system:dept:del', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(140, 500, 'C', '岗位管理', 'el-icon-PriceTag', 7, 'admin:system:post:list', 'post', 'system/post/index', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2025-09-18 12:12:24'),
	(141, 140, 'A', '岗位详情', '', 0, 'admin:system:post:detail', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(142, 140, 'A', '岗位新增', '', 0, 'admin:system:post:add', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(143, 140, 'A', '岗位编辑', '', 0, 'admin:system:post:edit', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(144, 140, 'A', '岗位删除', '', 0, 'admin:system:post:del', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(200, 0, 'M', '其它管理', '', 0, '', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(201, 200, 'M', '图库管理', '', 0, '', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(202, 201, 'A', '文件列表', '', 0, 'admin:common:album:albumList', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(203, 201, 'A', '文件命名', '', 0, 'admin:common:album:albumRename', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(204, 201, 'A', '文件移动', '', 0, 'admin:common:album:albumMove', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(205, 201, 'A', '文件删除', '', 0, 'admin:common:album:albumDel', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(206, 201, 'A', '分类列表', '', 0, 'admin:common:album:cateList', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(207, 201, 'A', '分类新增', '', 0, 'admin:common:album:cateAdd', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(208, 201, 'A', '分类命名', '', 0, 'admin:common:album:cateRename', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(209, 201, 'A', '分类删除', '', 0, 'admin:common:album:cateDel', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(215, 200, 'M', '上传管理', '', 0, '', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(216, 215, 'A', '上传文件', '', 0, 'admin:common:upload:file', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2025-08-25 18:28:02'),
	(500, 0, 'M', '系统设置', 'el-icon-Setting', 0, '', 'system', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2025-09-18 12:57:22'),
	(501, 500, 'M', '网站设置', 'el-icon-Basketball', 0, '', 'website', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2025-09-18 11:20:05'),
	(502, 501, 'C', '网站信息', '', 0, 'admin:setting:website:detail', 'information', 'system/website/information', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2025-09-18 12:13:52'),
	(503, 502, 'A', '保存配置', '', 0, 'admin:setting:website:save', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(505, 501, 'C', '网站备案', '', 0, 'admin:setting:copyright:detail', 'filing', 'system/website/filing', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2025-09-18 12:15:00'),
	(506, 505, 'A', '备案保存', '', 0, 'admin:setting:copyright:save', '', 'setting/website/protocol', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(515, 500, 'C', '字典管理', 'el-icon-Box', 6, 'admin:setting:dict:type:list', 'dict', 'system/dict/type/index', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2025-09-18 12:12:58'),
	(516, 515, 'A', '字典类型新增', '', 0, 'admin:setting:dict:type:add', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(517, 515, 'A', '字典类型编辑', '', 0, 'admin:setting:dict:type:edit', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(518, 515, 'A', '字典类型删除', '', 0, 'admin:setting:dict:type:del', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(520, 515, 'A', '字典数据新增', '', 0, 'admin:setting:dict:data:add', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(521, 515, 'A', '字典数据编辑', '', 0, 'admin:setting:dict:data:edit', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(522, 515, 'A', '字典数据删除', '', 0, 'admin:setting:dict:data:del', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(550, 500, 'M', '系统维护', 'el-icon-SetUp', 0, '', 'system', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(551, 550, 'C', '系统环境', '', 0, 'admin:monitor:server', 'environment', 'system/system/environment', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2025-09-18 12:22:43'),
	(552, 550, 'C', '系统缓存', '', 0, 'admin:monitor:cache', 'system/cache', 'system/system/cache', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2025-09-18 12:21:45'),
	(553, 550, 'C', '系统日志', '', 0, 'admin:system:log:operate', 'journal', 'system/system/journal', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2025-09-18 12:22:00'),
	(600, 0, 'M', '开发工具', 'el-icon-EditPen', 0, '', 'dev_tools', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(610, 600, 'C', '代码生成器', '', 0, 'admin:gen:list', 'code', 'dev_tools/code/index', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-10-31 14:59:14'),
	(611, 610, 'A', '导入数据表', '', 0, 'admin:gen:importTable', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(612, 610, 'A', '生成代码', '', 0, 'admin:gen:genCode', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(613, 610, 'A', '下载代码', '', 0, 'admin:gen:downloadCode', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(614, 610, 'A', '预览代码', '', 0, 'admin:gen:previewCode', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(616, 610, 'A', '同步表结构', '', 0, 'admin:gen:syncTable', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(617, 610, 'A', '删除数据表', '', 0, 'admin:gen:delTable', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(618, 610, 'A', '数据表详情', '', 0, 'admin:gen:detail', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(700, 0, 'M', '素材管理', 'el-icon-Picture', 43, '', 'material', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2025-06-20 01:44:00'),
	(701, 700, 'C', '素材中心', 'el-icon-PictureRounded', 0, '', 'index', 'material/index', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(775, 600, 'C', '代码生成器编辑', 'el-icon-EditPen', 0, 'admin:gen:editTable', 'code/edit', 'dev_tools/code/edit', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2025-09-18 12:29:37'),
	(776, 778, 'C', '流程模板', '', 0, 'admin:flow:flow_template:list', 'flow_template/index', 'flow/flow_template/index', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(777, 778, 'C', '我的流程', '', 0, '', 'flow_apply/index', 'flow/flow_apply/index', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(778, 0, 'M', '审批流', 'el-icon-Coordinate', 0, '', 'flow', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(780, 778, 'C', '待处理', '', 0, 'admin:flow:flow_history:list', 'flow_history/todo', 'flow/flow_history/todo', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(781, 778, 'C', '已处理', '', 0, 'admin:flow:flow_history:list', 'flow_history/done', 'flow/flow_history/done', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(782, 778, 'C', '已完成流程', '', 0, 'admin:flow:flow_history:list', 'flow_apply/finish', 'flow/flow_apply/finish', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(783, 832, 'C', '项目', '', 0, 'admin:monitor_project:list', 'project/index', 'monitor/project/index', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-10-29 15:40:44'),
	(784, 832, 'C', '用户端', '', 0, '', 'client/index', 'monitor/client/index', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-10-29 15:40:54'),
	(785, 794, 'A', '错误收集error添加', '', 0, 'admin:monitor_web:add', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(786, 794, 'A', '错误收集error编辑', '', 0, 'admin:monitor_web:edit', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(787, 794, 'A', '错误收集error删除', '', 0, 'admin:monitor_web:del', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(788, 794, 'A', '错误收集error列表', '', 0, 'admin:monitor_web:list', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(789, 794, 'A', '错误收集error全部列表', '', 0, 'admin:monitor_web:listAll', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(790, 794, 'A', '错误收集error详情', '', 0, 'admin:monitor_web:detail', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(791, 794, 'A', '错误收集error导出excel', '', 0, 'admin:monitor_web:ExportFile', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(792, 794, 'A', '错误收集error导入excel', '', 0, 'admin:monitor_web:ImportFile', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(794, 832, 'C', '错误列表', '', 0, '', 'error/index', 'monitor/error/index', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-10-29 15:39:16'),
	(795, 776, 'A', '流程模板添加', '', 0, 'admin:flow:flow_template:add', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(796, 776, 'A', '流程模板编辑', '', 0, 'admin:flow:flow_template:edit', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(797, 776, 'A', '流程模板删除', '', 0, 'admin:flow:flow_template:del', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(798, 776, 'A', '流程模板列表', '', 0, 'admin:flow:flow_template:list', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(799, 776, 'A', '流程模板全部列表', '', 0, 'admin:flow:flow_template:listAll', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(800, 776, 'A', '流程模板详情', '', 0, 'admin:flow:flow_template:detail', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(801, 777, 'A', '流程详情', '', 0, 'admin:flow:flow_apply:detail', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(802, 777, 'A', '流程添加', '', 0, 'admin:flow:flow_apply:add', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(803, 777, 'A', '流程编辑', '', 0, 'admin:flow:flow_apply:edit', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(804, 777, 'A', '流程删除', '', 0, 'admin:flow:flow_apply:del', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(805, 777, 'A', '全部流程', '', 0, 'admin:flow:flow_history:listAll', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(806, 777, 'A', '通过流程', '', 0, 'admin:flow:flow_history:pass', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(807, 777, 'A', '拒接流程', '', 0, 'admin:flow:flow_history:back', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(808, 777, 'A', '下一个流程', '', 0, 'admin:flow:flow_history:next_node', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(809, 777, 'A', '获取审批人', '', 0, 'admin:flow:flow_history:get_approver', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(810, 777, 'A', '流程列表', '', 0, 'admin:flow:flow_apply:list', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(811, 780, 'A', '审批记录列表', '', 0, 'admin:flow:flow_history:list', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(812, 780, 'A', '审批记录详情', '', 0, 'admin:flow:flow_history:detail', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(813, 783, 'A', '监控项目添加', '', 0, 'admin:monitor_project:add', '', '', '', '', 0, 1, 0, '2024-06-17 19:26:13', '2024-06-17 19:26:13'),
	(814, 783, 'A', '监控项目编辑', '', 0, 'admin:monitor_project:edit', '', '', '', '', 0, 1, 0, '2024-06-17 19:26:13', '2024-06-17 19:26:13'),
	(815, 783, 'A', '监控项目删除', '', 0, 'admin:monitor_project:del', '', '', '', '', 0, 1, 0, '2024-06-17 19:26:13', '2024-06-17 19:26:13'),
	(816, 783, 'A', '监控项目列表', '', 0, 'admin:monitor_project:list', '', '', '', '', 0, 1, 0, '2024-06-17 19:26:13', '2024-06-17 19:26:13'),
	(817, 783, 'A', '监控项目全部列表', '', 0, 'admin:monitor_project:listAll', '', '', '', '', 0, 1, 0, '2024-06-17 19:26:13', '2024-06-17 19:26:13'),
	(818, 783, 'A', '监控项目详情', '', 0, 'admin:monitor_project:detail', '', '', '', '', 0, 1, 0, '2024-06-17 19:26:13', '2024-06-17 19:26:13'),
	(819, 783, 'A', '监控项目导出excel', '', 0, 'admin:monitor_project:ExportFile', '', '', '', '', 0, 1, 0, '2024-06-17 19:26:13', '2024-06-17 19:26:13'),
	(820, 783, 'A', '监控项目导入excel', '', 0, 'admin:monitor_project:ImportFile', '', '', '', '', 0, 1, 0, '2024-06-17 19:26:13', '2024-06-17 19:26:13'),
	(821, 0, 'C', '文档', 'el-icon-Wallet', 0, '', 'https://adtkcn.github.io/x_admin/', '', '', 'a=1', 1, 1, 0, '2024-06-28 17:09:17', '2024-08-16 15:05:08'),
	(823, 0, 'C', '用户协议', 'el-icon-Coordinate', 0, '', 'user/protocol/index', 'user/protocol/index', '', '', 0, 1, 0, '2024-09-10 20:03:46', '2025-07-17 16:21:15'),
	(824, 823, 'A', '用户协议添加', '', 0, 'admin:user_protocol:add', '', '', '', '', 0, 1, 0, '2024-09-10 20:05:13', '2024-09-10 20:05:13'),
	(825, 823, 'A', '用户协议编辑', '', 0, 'admin:user_protocol:edit', '', '', '', '', 0, 1, 0, '2024-09-10 20:05:13', '2024-09-10 20:05:13'),
	(826, 823, 'A', '用户协议删除', '', 0, 'admin:user_protocol:del', '', '', '', '', 0, 1, 0, '2024-09-10 20:05:13', '2024-09-10 20:05:13'),
	(827, 823, 'A', '用户协议列表', '', 0, 'admin:user_protocol:list', '', '', '', '', 0, 1, 0, '2024-09-10 20:05:13', '2024-09-10 20:05:13'),
	(828, 823, 'A', '用户协议全部列表', '', 0, 'admin:user_protocol:listAll', '', '', '', '', 0, 1, 0, '2024-09-10 20:05:13', '2024-09-10 20:05:13'),
	(829, 823, 'A', '用户协议详情', '', 0, 'admin:user_protocol:detail', '', '', '', '', 0, 1, 0, '2024-09-10 20:05:13', '2024-09-10 20:05:13'),
	(830, 823, 'A', '用户协议导出excel', '', 0, 'admin:user_protocol:ExportFile', '', '', '', '', 0, 1, 0, '2024-09-10 20:05:13', '2024-09-10 20:05:13'),
	(831, 823, 'A', '用户协议导入excel', '', 0, 'admin:user_protocol:ImportFile', '', '', '', '', 0, 1, 0, '2024-09-10 20:05:13', '2024-09-10 20:05:13'),
	(832, 0, 'M', '错误监控', 'el-icon-Memo', 0, '', 'monitor', '', '', '', 1, 1, 0, '2024-10-29 15:32:33', '2024-10-29 15:32:33'),
	(833, 600, 'C', '错误捕获', '', 0, '', 'test', 'dev_tools/test', '', '', 1, 1, 0, '2024-10-30 18:16:56', '2025-09-18 12:28:08'),
	(864, 784, 'A', '监控-客户端信息添加', '', 0, 'admin:monitor_client:add', '', '', '', '', 0, 1, 0, '2024-11-08 14:15:28', '2024-11-08 14:15:28'),
	(865, 784, 'A', '监控-客户端信息编辑', '', 0, 'admin:monitor_client:edit', '', '', '', '', 0, 1, 0, '2024-11-08 14:15:28', '2024-11-08 14:15:28'),
	(866, 784, 'A', '监控-客户端信息删除', '', 0, 'admin:monitor_client:del', '', '', '', '', 0, 1, 0, '2024-11-08 14:15:28', '2024-11-08 14:15:28'),
	(867, 784, 'A', '监控-客户端信息删除-批量', '', 0, 'admin:monitor_client:delBatch', '', '', '', '', 0, 1, 0, '2024-11-08 14:15:28', '2024-11-08 14:15:28'),
	(868, 784, 'A', '监控-客户端信息列表', '', 0, 'admin:monitor_client:list', '', '', '', '', 0, 1, 0, '2024-11-08 14:15:28', '2024-11-08 14:15:28'),
	(869, 784, 'A', '监控-客户端信息全部列表', '', 0, 'admin:monitor_client:listAll', '', '', '', '', 0, 1, 0, '2024-11-08 14:15:28', '2024-11-08 14:15:28'),
	(870, 784, 'A', '监控-客户端信息详情', '', 0, 'admin:monitor_client:detail', '', '', '', '', 0, 1, 0, '2024-11-08 14:15:28', '2024-11-08 14:15:28'),
	(871, 784, 'A', '监控-客户端信息导出excel', '', 0, 'admin:monitor_client:ExportFile', '', '', '', '', 0, 1, 0, '2024-11-08 14:15:28', '2024-11-08 14:15:28'),
	(872, 784, 'A', '监控-客户端信息导入excel', '', 0, 'admin:monitor_client:ImportFile', '', '', '', '', 0, 1, 0, '2024-11-08 14:15:28', '2024-11-08 14:15:28'),
	(873, 600, 'C', '文件分片上传', '', 0, '', 'uploadChunk', 'dev_tools/uploadChunk', '', '', 1, 1, 0, '2025-07-19 18:37:48', '2025-09-18 12:27:57');

-- 导出  表 x_admin_2.x_system_auth_perm 结构
CREATE TABLE IF NOT EXISTS `x_system_auth_perm` (
  `id` varchar(100) NOT NULL DEFAULT '' COMMENT '主键',
  `role_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '角色ID',
  `menu_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '菜单ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='系统角色菜单表';

-- 正在导出表  x_admin_2.x_system_auth_perm 的数据：~147 rows (大约)
REPLACE INTO `x_system_auth_perm` (`id`, `role_id`, `menu_id`) VALUES
	('010eaf6fd96945678751343e7c446cb8', 1, 201),
	('0130298c38fe40feb84b9730c4727a5b', 1, 200),
	('023b1e5e7e78430ba84622824b168a5c', 1, 120),
	('04b89871ecc84a58aad426cd29c22eea', 1, 500),
	('068ea1b5a763461eb95cc6ce6b6de8b3', 1, 868),
	('06cf35e0187e4f4992484d11ebd52d0f', 1, 788),
	('09f8e446b73141cf8ee15127332b3729', 1, 502),
	('0ae242d7afcd49f6be4e00d17716b9c8', 1, 824),
	('0e3dec9ecc94449a8520d76531c65555', 1, 217),
	('157d356da1fb4e6f8b02a188772c6951', 1, 114),
	('15be80be3bde4bb2b5c14bc1d14c0d89', 1, 519),
	('162b7e475fe448da86e549f6a00502cd', 1, 830),
	('1712a8af52c6483f8440270162bd7ea3', 1, 105),
	('17f1ece467e948ab8ce5a6718e94cc6d', 1, 505),
	('1a99c882f8884188895f2516e6dd34d7', 1, 811),
	('1c81dbbf98854215b43551605e8d4d22', 1, 202),
	('1fde3bb916d943e384a16840790a5890', 1, 815),
	('216f99dd2e974c7f91857f4b0aaa2022', 1, 871),
	('22527c15af584ba78d759760cc7227d0', 1, 778),
	('22600365dcf94fcea22c5e8928e4e737', 1, 777),
	('227f1549d8cd460d9a86f99163ee8541', 1, 511),
	('23331d4a076c4caea9b15e29872a6aef', 1, 817),
	('24d7403bdde6428999a0113a51c470af', 1, 121),
	('27303a99427f49f5b53fe71541848b99', 1, 846),
	('2933c6192b004b0c983d2919535fca59', 1, 808),
	('2bb8edac06a04f8b9116b65cd270cff6', 1, 786),
	('2c34fa119755439cb0e49fc6f3a260e5', 1, 205),
	('2dc4f04a124d45d2b5896ad0568cf6a5', 1, 124),
	('3018096ff6fc413bb68691682f8a732a', 1, 100),
	('337eb76fbc7447439a693c999629d04b', 1, 518),
	('34fc47a7fd0e4409b3bf1e6d99505fe3', 1, 600),
	('3634a82a40af4369a46dbf058ae047d1', 1, 551),
	('3699548738ae44d48f37c385cf12c077', 1, 131),
	('38544db3ef65441abdf238582b12a6f5', 1, 814),
	('38be86d6990440b9aa7b95dd953d9283', 1, 834),
	('3c86a28d13d64dbcbdc50e617f19e02d', 1, 140),
	('3e3517646afb4e72b55996bbdeabed4e', 1, 807),
	('40e36b8ec4ec4f5085a01364f1f2a309', 1, 844),
	('421a4629a2d64a0db497f1d172a3f461', 1, 832),
	('4419c9318eee42b5b7954c478d4e230e', 1, 130),
	('4421cd092f2f49e4a74287dc708bbfbc', 1, 614),
	('443131145c3c47d5a8f4b7f3f0de1950', 1, 506),
	('4752c8946b1e409592d418b591e31dfc', 1, 1),
	('4895087c70d94afdb18fd5aff53c785d', 1, 613),
	('491e820b6b654db8b1d1cb55ed6a2e0e', 1, 790),
	('4979ffa5c952425fbcd388348da957b2', 1, 784),
	('49a5b64bc91646d493b78fcc1c929c18', 1, 110),
	('4c193e94525d4243a60cc7f1dace9801', 1, 821),
	('4e470fcba8f94b75ae7258fda3bf9714', 1, 823),
	('4e8fc0eeac47455bba486c353f61b249', 1, 123),
	('51a96d49f73a4945b26104af18f6dec5', 1, 701),
	('53227828c481406fad795b6bfcb481b2', 1, 612),
	('5349e8d242604c9db5112dc8794334e1', 1, 866),
	('53781886a917441b8a3f91ca51a756e2', 1, 806),
	('5888e698ffc24dd5a5db79b1424c3c82', 1, 801),
	('5b58ab3be33c4cb0bcff81d982bd74e8', 1, 845),
	('5d46f0683edf4b1cb53921f694d2bcac', 1, 550),
	('5d7ec65a31b54a5bb9187fa5cf965ef4', 1, 203),
	('5e264bbe35be4fa4b1382116824ab7d8', 1, 113),
	('5e6137c41ec24c1d90a5086bb194fa5c', 1, 829),
	('60bca7730e724a459bb541b2f178281c', 1, 101),
	('62daa21098154699abd67c23bcc0f912', 1, 106),
	('62f70e1cddd94144a2e0315efa94a783', 1, 510),
	('64140ab994194f3e829651fb0600f217', 1, 867),
	('67cfa0b6486047118721788c5ce3191d', 1, 516),
	('6d6bdddee234478789909718d5214557', 1, 144),
	('6da2b06197aa4d699781a874d3391320', 1, 794),
	('707ac42ec0a14d0882b68d0962dbd493', 1, 618),
	('775ff9632b5340449d318accf828246b', 1, 207),
	('7a5c9bb98a894133bed68f63b7d66f26', 1, 611),
	('7af1f9cb9996424394504765244e5e97', 1, 820),
	('7b85e19d96be4050a7af95348e11abb6', 1, 610),
	('7bb886763763417f8bab4d6f11278b6e', 1, 850),
	('7c28f3bbdc584c899c508c441d7e83ee', 1, 206),
	('8166e49f55d74c4d95a2d2e614cd874c', 1, 501),
	('835dcf847728459da14c1fee210b3cae', 1, 797),
	('8469fee0c6bf475fb831905c85964623', 1, 789),
	('85dc3519937c461f974c22070f546fb0', 1, 816),
	('85e0188af90d4612a00347e1125eb3f2', 1, 553),
	('8770e1a05aaf4b6b805a6f19cf8c1059', 1, 819),
	('8bff8f436f484cdc83d4440479306fee', 1, 208),
	('8c22f741ff2e44f6aaccc9dc43f26709', 1, 800),
	('8c6ff7f1f9b444d582ea45c05dd12517', 1, 864),
	('8f9c9b51f61148c7b38c984d0c01b2e8', 1, 776),
	('90bc31c28c244a0b94f04c36174ff848', 1, 848),
	('988eb38ac8cb42ba8463ed01c8618961', 1, 111),
	('9b51feb808ff4568b534cd7e12bfee62', 1, 872),
	('9c48e0641ccf414a9eb177264a60b34f', 1, 517),
	('9d3c2ed01a3449b898ee08439b4cc3b7', 1, 804),
	('9f26e153274b4464a71afeede9c131d1', 1, 780),
	('9fb9485f97864d9b9f1ae0a3aeed4f5e', 1, 852),
	('a17a6509fb3845d9a179cc4d2c94d711', 1, 798),
	('a35b2a92ad54439181ef177d60a29ae1', 1, 104),
	('a39790550a844ff9a1135043aebbc07d', 1, 796),
	('a523e3a900fa4a409e87fe66c239a25a', 1, 515),
	('a536ff1cd7ae41a9b47db3ea8391d8fc', 1, 847),
	('a6287252ca394121b0c62cbe75262d6c', 1, 775),
	('a7a7b4915a7249e18f3b05ef4b41046d', 1, 216),
	('a8883d69d6b64d8f9de808f6cecdecc8', 1, 520),
	('aa356c3528c04aef96b212184c028859', 1, 787),
	('ab881c3a8fe645dfbbed4ce3d68e808e', 1, 552),
	('af305ee7813a4f79af8cdf08d7d3df08', 1, 616),
	('af738564cafb4d9ca0ddef1eea9b130f', 1, 617),
	('afde9f4321f242fab35997cfa376565b', 1, 783),
	('b96a754d3e5d4f18b7c78938f4b2797f', 1, 112),
	('ba2cbeaa408c43e5a7a2ea268588cb14', 1, 521),
	('bb64a0acf70847b08b0e351abfc566bd', 1, 809),
	('bdd1dcfaf8d64fa998c3f70de49943f9', 1, 102),
	('be37d238b52947a08e3e34b3e7d4deb3', 1, 812),
	('beaea27ad96348e7b71f3df0ce20a9a2', 1, 133),
	('c2a02fb5b52d4246a6ccc81d3a556c6f', 1, 795),
	('c2b55653d5874aa2a1a81cce84c95619', 1, 826),
	('c40aeb2cea9f4e0d92fab9d397fdf847', 1, 851),
	('c4c97c5f1a3a45748cc5979564a38788', 1, 782),
	('c5e45b921e574e82b5828210d456bb87', 1, 143),
	('c611404c636949e09458e58ca72bc8b6', 1, 792),
	('c67b0df8543641e7b26006fb31c778ea', 1, 831),
	('c6921779d0f14961a81d48264ea82d4e', 1, 122),
	('c7ae5ebad58e484ebaddb91a51a0b47f', 1, 799),
	('c86220b9ee8b41b0838f4e081b483221', 1, 849),
	('cab18e5e6f0541e7adf8659ac2bb0d88', 1, 870),
	('cc1a5d8d865d4ff9b20ac9c55ca15d0a', 1, 818),
	('cce58db2e5774334877573d273caa169', 1, 209),
	('cd6a5b83b61b4555a92d86306ee39973', 1, 132),
	('d558a7e10c914e849be26ec1990e97ca', 1, 781),
	('d5e1c31c8222449b9beb3ad626558a7b', 1, 803),
	('d6c15bc0672e41a7aa68ce1909a4927c', 1, 802),
	('d8d73b25258040f584dcd76f4616d7b2', 1, 700),
	('d993c677e6f049ffb7fd8ee6a6cf9997', 1, 869),
	('da16de515087479baace082f0843d6cd', 1, 810),
	('dd5beee245374853bb93e1e311f54c9a', 1, 142),
	('e09760dbff1e4d91ad0320ac0e8ad6ca', 1, 204),
	('e23e0762cb874989b73bc4df4779800e', 1, 828),
	('e3a0c712ec3f4281a9c7c559adb8e739', 1, 135),
	('e588b716529f461fbc92454878e675c3', 1, 503),
	('e83eec9141f04174a376505c16c0b123', 1, 827),
	('edbf800908714038aac64adc1516757d', 1, 805),
	('eef1f4f099c0488396e75dd01103483d', 1, 522),
	('f2f4d7fae39e4706b501c1182147804e', 1, 141),
	('f350cfb1b03342ee81061fc11ee42f9f', 1, 791),
	('f4e29b50e91f4a0d8384e316cdbda8f7', 1, 865),
	('f53beaea543b46808bd503255d71eecb', 1, 813),
	('f666b9f26cda4217b8389abfc9ba0e1f', 1, 103),
	('f68605c24424413aad2493643cba8f46', 1, 215),
	('f84fc6baae0b4aa3ba1e1c179505958a', 1, 785),
	('fad966cf79824d349e3aa5bb8c7ac711', 1, 134),
	('fc68ecb541ca47ea9fde1753ce5f5cef', 1, 825);

-- 导出  表 x_admin_2.x_system_auth_post 结构
CREATE TABLE IF NOT EXISTS `x_system_auth_post` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
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
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='系统岗位管理表';

-- 正在导出表  x_admin_2.x_system_auth_post 的数据：~1 rows (大约)
REPLACE INTO `x_system_auth_post` (`id`, `code`, `name`, `remarks`, `sort`, `is_stop`, `is_delete`, `create_time`, `update_time`, `delete_time`) VALUES
	(3, 'zhihuibu01', '指挥部岗位', '', 0, 0, 0, '2024-01-02 03:04:05', '2025-07-16 17:33:24', NULL);

-- 导出  表 x_admin_2.x_system_auth_role 结构
CREATE TABLE IF NOT EXISTS `x_system_auth_role` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '角色名称',
  `remark` varchar(200) NOT NULL DEFAULT '' COMMENT '备注信息',
  `sort` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '角色排序',
  `is_disable` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否禁用: 0=否, 1=是',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='系统角色管理表';

-- 正在导出表  x_admin_2.x_system_auth_role 的数据：~1 rows (大约)
REPLACE INTO `x_system_auth_role` (`id`, `name`, `remark`, `sort`, `is_disable`, `create_time`, `update_time`) VALUES
	(1, '审核员', '1', 1, 0, '2024-01-02 03:04:05', '2024-12-04 11:53:36');

-- 导出  表 x_admin_2.x_system_config 结构
CREATE TABLE IF NOT EXISTS `x_system_config` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `type` varchar(30) DEFAULT '' COMMENT '类型',
  `name` varchar(60) NOT NULL DEFAULT '' COMMENT '键',
  `value` text COMMENT '值',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=81 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='系统全局配置表';

-- 正在导出表  x_admin_2.x_system_config 的数据：~46 rows (大约)
REPLACE INTO `x_system_config` (`id`, `type`, `name`, `value`, `create_time`, `update_time`) VALUES
	(1, 'storage', 'default', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(2, 'storage', 'local', '{"name":"本地存储"}', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(3, 'storage', 'qiniu', '{"name":"七牛云存储","bucket":"","secretKey":"","accessKey":"","domain":""}', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(4, 'storage', 'aliyun', '{"name":"阿里云存储","bucket":"","secretKey":"","accessKey":"","domain":""}', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(5, 'storage', 'qcloud', '{"name":"腾讯云存储","bucket":"","secretKey":"","accessKey":"","domain":"","region":""}', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(6, 'sms', 'default', 'aliyun', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(7, 'sms', 'aliyun', '{"name":"阿里云短信","alias":"aliyun","sign":"","appKey":"","secretKey":""}', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(8, 'sms', 'tencent', '{"name":"腾讯云短信","alias":"tencent","sign":"","appId":"","secretId":"","secretKey":""}', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(9, 'sms', 'huawei', '{"name":"华为云短信","alias":"huawei"}', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(10, 'website', 'name', 'x_admin开源系统', '2024-01-02 03:04:05', '2025-06-24 19:52:37'),
	(11, 'website', 'logo', '/api/static/backend_logo.png', '2024-01-02 03:04:05', '2025-06-24 19:52:37'),
	(12, 'website', 'favicon', '/api/static/backend_favicon.ico', '2024-01-02 03:04:05', '2025-06-24 19:52:38'),
	(13, 'website', 'backdrop', '/api/static/backend_backdrop.png', '2024-01-02 03:04:05', '2025-06-24 19:52:38'),
	(14, 'website', 'copyright', '[{"name":"蜀ICP备15007060号-1","link":"http://www.beian.gov.cn"},{"name":"x_admin","link":"http://x.adtk.cn"}]', '2024-01-02 03:04:05', '2024-06-29 00:30:54'),
	(15, 'website', 'shopName', 'x_admin开源管理系统', '2024-01-02 03:04:05', '2025-06-24 19:52:38'),
	(16, 'website', 'shopLogo', '/image/20241705/f0eb36d508834bc2ac1b8c591c563efa.png', '2024-01-02 03:04:05', '2025-06-24 19:52:38'),
	(17, 'protocol', 'service', '{"name":"服务协议","content":"\\u003cp\\u003e服务协议666\\u003c/p\\u003e"}', '2024-01-02 03:04:05', '2024-06-29 00:30:56'),
	(18, 'protocol', 'privacy', '{"name":"隐私协议","content":"\\u003cp\\u003e隐私协议\\u003c/p\\u003e"}', '2024-01-02 03:04:05', '2024-06-29 00:30:56'),
	(19, 'tabbar', 'style', '{"defaultColor":"#4A5DFF","selectedColor":"#EA5455"}', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(20, 'search', 'isHotSearch', '0', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(30, 'h5_channel', 'status', '1', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(31, 'h5_channel', 'close', '0', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(32, 'h5_channel', 'url', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(40, 'mp_channel', 'name', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(41, 'mp_channel', 'primaryId', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(42, 'mp_channel', 'appId', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(43, 'mp_channel', 'appSecret', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(44, 'mp_channel', 'qrCode', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(50, 'wx_channel', 'appId', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(51, 'wx_channel', 'appSecret', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(55, 'oa_channel', 'name', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(56, 'oa_channel', 'primaryId', ' ', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(57, 'oa_channel', 'qrCode', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(58, 'oa_channel', 'appId', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(59, 'oa_channel', 'appSecret', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(60, 'oa_channel', 'url', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(61, 'oa_channel', 'token', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(62, 'oa_channel', 'encodingAesKey', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(63, 'oa_channel', 'encryptionType', '1', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(64, 'oa_channel', 'menus', '[]', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(70, 'login', 'loginWay', '1,2', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(71, 'login', 'forceBindMobile', '0', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(72, 'login', 'openAgreement', '1', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(73, 'login', 'openOtherAuth', '1', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(74, 'login', 'autoLoginAuth', '1,2', '2024-01-02 03:04:05', '2024-01-02 03:04:05'),
	(80, 'user', 'defaultAvatar', '/api/static/default_avatar.png', '2024-01-02 03:04:05', '2024-01-02 03:04:05');

-- 导出  表 x_admin_2.x_system_log_login 结构
CREATE TABLE IF NOT EXISTS `x_system_log_login` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '注解',
  `admin_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '管理员ID',
  `username` varchar(30) NOT NULL DEFAULT '' COMMENT '登录账号',
  `ip` varchar(30) NOT NULL COMMENT '登录地址',
  `os` varchar(100) NOT NULL DEFAULT '' COMMENT '操作系统',
  `browser` varchar(100) DEFAULT '' COMMENT '浏览器',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '操作状态: 1=成功, 2=失败',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=110 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='系统登录日志表';

-- 正在导出表  x_admin_2.x_system_log_login 的数据：~0 rows (大约)

-- 导出  表 x_admin_2.x_system_log_operate 结构
CREATE TABLE IF NOT EXISTS `x_system_log_operate` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `admin_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '操作人ID',
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
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='系统操作日志表';

-- 正在导出表  x_admin_2.x_system_log_operate 的数据：~0 rows (大约)

-- 导出  表 x_admin_2.x_system_log_sms 结构
CREATE TABLE IF NOT EXISTS `x_system_log_sms` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
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
  `last_login_ip` varchar(30) NOT NULL DEFAULT '' COMMENT '最后登录IP',
  `last_login_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '最后登录时间',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='用户信息表';

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
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `tag` varchar(50) NOT NULL COMMENT '标识',
  `version` int(10) DEFAULT NULL COMMENT '排序',
  `title` varchar(50) NOT NULL COMMENT '标题',
  `content` mediumtext COMMENT '协议内容',
  `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除: 0=否, 1=是',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COMMENT='用户协议';

-- 正在导出表  x_admin_2.x_user_protocol 的数据：0 rows
/*!40000 ALTER TABLE `x_user_protocol` DISABLE KEYS */;
/*!40000 ALTER TABLE `x_user_protocol` ENABLE KEYS */;

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
