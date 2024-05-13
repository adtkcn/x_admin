/*
 Navicat Premium Data Transfer

 Source Server         : x_admin
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : localhost:3306
 Source Schema         : x_admin

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 29/04/2024 18:50:40
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for x_album
-- ----------------------------
DROP TABLE IF EXISTS `x_album`;
CREATE TABLE `x_album`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `cid` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '类目ID',
  `aid` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '管理员ID',
  `uid` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
  `type` tinyint(2) UNSIGNED NOT NULL DEFAULT 10 COMMENT '文件类型: [10=图片, 20=视频]',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '文件名称',
  `uri` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文件路径',
  `ext` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '文件扩展',
  `size` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '文件大小',
  `is_delete` int(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除: 0=否, 1=是',

  `create_time` DATETIME NOT NULL COMMENT '创建时间',
	`update_time` DATETIME NOT NULL COMMENT '更新时间',
	`delete_time` DATETIME  NULL COMMENT '删除时间',

  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_cid`(`cid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 31 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '相册管理表' ROW_FORMAT = DYNAMIC;



-- ----------------------------
-- Table structure for x_album_cate
-- ----------------------------
DROP TABLE IF EXISTS `x_album_cate`;
CREATE TABLE `x_album_cate`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `pid` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '父级ID',
  `type` tinyint(2) UNSIGNED NOT NULL DEFAULT 10 COMMENT '类型: [10=图片, 20=视频]',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '分类名称',
  `is_delete` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除: [0=否, 1=是]',
  `create_time` DATETIME NOT NULL COMMENT '创建时间',
	`update_time` DATETIME NOT NULL COMMENT '更新时间',
	`delete_time` DATETIME  NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '相册分类表' ROW_FORMAT = DYNAMIC;
 
-- ----------------------------
-- Table structure for x_article
-- ----------------------------
DROP TABLE IF EXISTS `x_article`;
CREATE TABLE `x_article`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `cid` int(10) UNSIGNED NOT NULL COMMENT '分类',
  `title` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '标题',
  `intro` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '简介',
  `summary` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '摘要',
  `image` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '封面',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '内容',
  `author` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '作者',
  `visit` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '浏览',
  `sort` int(10) UNSIGNED NOT NULL DEFAULT 50 COMMENT '排序',
  `is_show` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否显示: 0=否, 1=是',
  `is_delete` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除: 0=否, 1=是',
  `create_time` DATETIME NOT NULL COMMENT '创建时间',
	`update_time` DATETIME NOT NULL COMMENT '更新时间',
	`delete_time` DATETIME  NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `cid_idx`(`cid`) USING BTREE COMMENT '分类索引'
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文章资讯表' ROW_FORMAT = DYNAMIC;



-- ----------------------------
-- Table structure for x_article_category
-- ----------------------------
DROP TABLE IF EXISTS `x_article_category`;
CREATE TABLE `x_article_category`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '名称',
  `sort` smallint(5) UNSIGNED NOT NULL DEFAULT 50 COMMENT '排序',
  `is_show` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '是否显示: 0=否, 1=是',
  `is_delete` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除: 0=否, 1=是',
  `create_time` DATETIME NOT NULL COMMENT '创建时间',
	`update_time` DATETIME NOT NULL COMMENT '更新时间',
	`delete_time` DATETIME  NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文章分类表' ROW_FORMAT = DYNAMIC;



-- ----------------------------
-- Table structure for x_article_collect
-- ----------------------------
DROP TABLE IF EXISTS `x_article_collect`;
CREATE TABLE `x_article_collect`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
  `article_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '文章ID',
  `is_delete` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除',
  `create_time` DATETIME NOT NULL COMMENT '创建时间',
	`update_time` DATETIME NOT NULL COMMENT '更新时间',
	`delete_time` DATETIME  NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文章收藏表' ROW_FORMAT = DYNAMIC;



-- ----------------------------
-- Table structure for x_dict_data
-- ----------------------------
DROP TABLE IF EXISTS `x_dict_data`;
CREATE TABLE `x_dict_data`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `type_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '类型',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '键名',
  `value` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '数值',
  `color` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '颜色',
  `remark` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '备注',
  `sort` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序',
  `status` tinyint(1) NOT NULL COMMENT '状态: 0=停用, 1=正常',
  `is_delete` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除: 0=否, 1=是',
  `create_time` DATETIME NOT NULL COMMENT '创建时间',
	`update_time` DATETIME NOT NULL COMMENT '更新时间',
	`delete_time` DATETIME  NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 21 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '字典数据表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_dict_data
-- ----------------------------
INSERT INTO `x_dict_data` VALUES (1, 2, '待提交', '1', '#6D85FC', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05", null);
INSERT INTO `x_dict_data` VALUES (2, 2, '审批中', '2', '#C6C150', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05", null);
INSERT INTO `x_dict_data` VALUES (3, 2, '审批成功', '3', 'green', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05", null);
INSERT INTO `x_dict_data` VALUES (4, 2, '失败', '4', 'red', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05", null);
INSERT INTO `x_dict_data` VALUES (5, 3, '待处理', '1', '#087BF6', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05", null);
INSERT INTO `x_dict_data` VALUES (6, 3, '通过', '2', 'green', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05", null);
INSERT INTO `x_dict_data` VALUES (7, 3, '拒绝', '3', 'red', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05", null);
INSERT INTO `x_dict_data` VALUES (10, 4, '假勤管理', '1', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05", null);
INSERT INTO `x_dict_data` VALUES (11, 4, '人事管理', '2', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05", null);
INSERT INTO `x_dict_data` VALUES (12, 4, '财务管理', '3', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05", null);
INSERT INTO `x_dict_data` VALUES (13, 4, '业务管理', '4', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05", null);
INSERT INTO `x_dict_data` VALUES (14, 4, '行政管理', '5', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05", null);
INSERT INTO `x_dict_data` VALUES (15, 4, '法务管理', '6', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05", null);
INSERT INTO `x_dict_data` VALUES (16, 4, '其他', '7', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05", null);
INSERT INTO `x_dict_data` VALUES (17, 5, 'web', 'web', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05", null);
INSERT INTO `x_dict_data` VALUES (18, 5, 'go', 'go', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05", null);
INSERT INTO `x_dict_data` VALUES (19, 5, 'uniapp', 'uniapp', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05", null);
INSERT INTO `x_dict_data` VALUES (20, 5, 'node', 'node', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05", null);

-- ----------------------------
-- Table structure for x_dict_type
-- ----------------------------
DROP TABLE IF EXISTS `x_dict_type`;
CREATE TABLE `x_dict_type`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `dict_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '字典名称',
  `dict_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '字典类型',
  `dict_remark` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '字典备注',
  `dict_status` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '字典状态: 0=停用, 1=正常',
  `is_delete` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除: 0=否, 1=是',
  `create_time` DATETIME NOT NULL COMMENT '创建时间',
	`update_time` DATETIME NOT NULL COMMENT '更新时间',
	`delete_time` DATETIME  NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '字典类型表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_dict_type
-- ----------------------------
INSERT INTO `x_dict_type` VALUES (2, '审批申请状态', 'flow_apply_status', '0待提交，1审批中，2审批完成，3审批失败', 1,0, "2024-01-02 03:04:05", "2024-01-02 03:04:05", null);
INSERT INTO `x_dict_type` VALUES (3, '审批历史状态', 'flow_history_status', '', 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05", null);
INSERT INTO `x_dict_type` VALUES (4, '流程分类', 'flow_group', '1假勤管理,2人事管理3财务管理4业务管理5行政管理6法务管理7其他', 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05", null);
INSERT INTO `x_dict_type` VALUES (5, '项目类型', 'project_type', '项目类型go java web node php 等', 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05", null);

-- ----------------------------
-- Table structure for x_flow_apply
-- ----------------------------
DROP TABLE IF EXISTS `x_flow_apply`;
CREATE TABLE `x_flow_apply`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `template_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '模板',
  `apply_user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '申请人id',
  `apply_user_nickname` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '0' COMMENT '申请人昵称',
  `flow_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '流程名称',
  `flow_group` tinyint(2) NOT NULL DEFAULT 0 COMMENT '流程分类',
  `flow_remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '流程描述',
  `flow_form_data` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '表单配置',
  `flow_process_data` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '流程配置',
  `flow_process_data_list` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '流程配置list',
  `form_value` mediumtext CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '表单值',
  `status` tinyint(2) UNSIGNED NULL DEFAULT 1 COMMENT '状态：1待提交，2审批中，3审批完成，4审批失败',
  `is_delete` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除: 0=否, 1=是',
  `create_time` DATETIME NOT NULL COMMENT '创建时间',
	`update_time` DATETIME NOT NULL COMMENT '更新时间',
	`delete_time` DATETIME  NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 35 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '申请流程' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for x_flow_history
-- ----------------------------
DROP TABLE IF EXISTS `x_flow_history`;
CREATE TABLE `x_flow_history`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '历史id',
  `apply_id` int(10) UNSIGNED NOT NULL COMMENT '申请id',
  `template_id` int(10) UNSIGNED NULL DEFAULT NULL COMMENT '模板id',
  `apply_user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '申请人id',
  `apply_user_nickname` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '0' COMMENT '申请人昵称',
  `approver_id` int(10) UNSIGNED NOT NULL COMMENT '审批人id',
  `approver_nickname` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '审批用户昵称',
  `node_id` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '节点',
  `node_type` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '节点类型',
  `node_label` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '节点名称',
  `form_value` mediumtext CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '表单值',
  `pass_status` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '通过状态：1待处理，2通过，3拒绝',
  `pass_remark` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '通过备注',
   `create_time` DATETIME NOT NULL COMMENT '创建时间',
	`update_time` DATETIME NOT NULL COMMENT '更新时间',
	`delete_time` DATETIME  NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `apply_id`(`apply_id`) USING BTREE,
  INDEX `approver_id`(`approver_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 98 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '流程历史' ROW_FORMAT = DYNAMIC;


-- ----------------------------
-- Table structure for x_flow_template
-- ----------------------------
DROP TABLE IF EXISTS `x_flow_template`;
CREATE TABLE `x_flow_template`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `flow_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '流程名称',
  `flow_group` tinyint(2) NOT NULL DEFAULT 0 COMMENT '流程分类',
  `flow_remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '流程描述',
  `flow_form_data` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '表单配置',
  `flow_process_data` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '流程配置',
  `flow_process_data_list` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '流程配置list',
    `create_time` DATETIME NOT NULL COMMENT '创建时间',
	`update_time` DATETIME NOT NULL COMMENT '更新时间',
	`delete_time` DATETIME  NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '流程模板' ROW_FORMAT = Dynamic;


-- ----------------------------
-- Table structure for x_gen_table
-- ----------------------------
DROP TABLE IF EXISTS `x_gen_table`;
CREATE TABLE `x_gen_table`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `table_name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '表名称',
  `table_comment` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '表描述',
  `sub_table_name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '关联表名称',
  `sub_table_fk` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '关联表外键',
  `author_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '作者的名称',
  `entity_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '实体的名称',
  `module_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '生成模块名',
  `function_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '生成功能名',
  `tree_primary` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '树主键字段',
  `tree_parent` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '树父级字段',
  `tree_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '树显示字段',
  `gen_tpl` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'crud' COMMENT '生成模板方式: [crud=单表, tree=树表]',
  `remarks` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注信息',
  `create_time` DATETIME NOT NULL COMMENT '创建时间',
	`update_time` DATETIME NOT NULL COMMENT '更新时间',

  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 25 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '代码生成业务表' ROW_FORMAT = DYNAMIC;


-- ----------------------------
-- Table structure for x_gen_table_column
-- ----------------------------
DROP TABLE IF EXISTS `x_gen_table_column`;
CREATE TABLE `x_gen_table_column`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '列主键',
  `table_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '表外键',
  `column_name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '列名称',
  `column_comment` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '列描述',
  `column_length` varchar(5) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '0' COMMENT '列长度',
  `column_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '列类型 ',
  `go_type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT 'JAVA类型',
  `go_field` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'JAVA字段',
  `is_pk` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否主键: [1=是, 0=否]',
  `is_increment` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否自增: [1=是, 0=否]',
  `is_required` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否必填: [1=是, 0=否]',
  `is_insert` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否插入字段: [1=是, 0=否]',
  `is_edit` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否编辑字段: [1=是, 0=否]',
  `is_list` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否列表字段: [1=是, 0=否]',
  `is_query` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否查询字段: [1=是, 0=否]',
  `query_type` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'EQ' COMMENT '查询方式: [等于、不等于、大于、小于、范围]',
  `html_type` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '显示类型: [文本框、文本域、下拉框、复选框、单选框、日期控件]',
  `dict_type` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '字典类型',
  `sort` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序编号',
  `create_time` DATETIME NOT NULL COMMENT '创建时间',
	`update_time` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 237 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '代码生成字段表' ROW_FORMAT = DYNAMIC;


-- ----------------------------
-- Table structure for x_monitor_client
-- ----------------------------
DROP TABLE IF EXISTS `x_monitor_client`;
CREATE TABLE `x_monitor_client`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'uuid',
  `project_key` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '项目key',
  `client_id` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'sdk生成的客户端id',
  `user_id` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户id',
  `os` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '系统',
  `browser` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '浏览器',
  `city` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '城市',
  `width` smallint(10) UNSIGNED NULL DEFAULT 0 COMMENT '屏幕',
  `height` smallint(10) UNSIGNED NULL DEFAULT 0 COMMENT '屏幕高度',
  `ua` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'ua记录',
   `create_time` DATETIME NOT NULL COMMENT '创建时间',
	`update_time` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `project_key`(`project_key`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  INDEX `client_id`(`client_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '监控-客户端信息' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_monitor_client
-- ----------------------------

-- ----------------------------
-- Table structure for x_monitor_project
-- ----------------------------
DROP TABLE IF EXISTS `x_monitor_project`;
CREATE TABLE `x_monitor_project`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '项目id',
  `project_key` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '项目uuid',
  `project_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '项目名称',
  `project_type` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '项目类型go java web node php 等',
  `is_delete` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除: 0=否, 1=是',
    `create_time` DATETIME NOT NULL COMMENT '创建时间',
	`update_time` DATETIME NOT NULL COMMENT '更新时间',
	`delete_time` DATETIME  NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `project_key`(`project_key`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 31 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '监控项目' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of x_monitor_project
-- ----------------------------

-- ----------------------------
-- Table structure for x_monitor_web
-- ----------------------------
DROP TABLE IF EXISTS `x_monitor_web`;
CREATE TABLE `x_monitor_web`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'uuid',
  `project_key` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '项目key',
  `client_id` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'sdk生成的客户端id',
  `event_type` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '事件类型',
  `page` varchar(1000) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'URL地址',
  `message` varchar(1000) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '错误消息',
  `stack` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '错误堆栈',
    `create_time` DATETIME NOT NULL COMMENT '创建时间',
	`update_time` DATETIME NOT NULL COMMENT '更新时间',
	
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `project_key`(`project_key`) USING BTREE,
  INDEX `client_id`(`client_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '监控-web端错误' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of x_monitor_web
-- ----------------------------

-- ----------------------------
-- Table structure for x_notice_setting
-- ----------------------------
DROP TABLE IF EXISTS `x_notice_setting`;
CREATE TABLE `x_notice_setting`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `scene` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '场景编号',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '场景名称',
  `remarks` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '场景描述',
  `recipient` tinyint(1) NOT NULL DEFAULT 1 COMMENT '接收人员: [1=用户, 2=平台]',
  `type` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '通知类型: [1=业务, 2=验证]',
  `system_notice` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '系统的通知设置',
  `sms_notice` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '短信的通知设置',
  `oa_notice` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '公众号通知设置',
  `mnp_notice` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '小程序通知设置',
  `is_delete` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除',
    `create_time` DATETIME NOT NULL COMMENT '创建时间',
	`update_time` DATETIME NOT NULL COMMENT '更新时间',
	`delete_time` DATETIME  NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '消息通知设置表' ROW_FORMAT = DYNAMIC;


-- ----------------------------
-- Table structure for x_system_auth_admin
-- ----------------------------
DROP TABLE IF EXISTS `x_system_auth_admin`;
CREATE TABLE `x_system_auth_admin`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `dept_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '部门ID',
  `post_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '岗位ID',
  `username` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户账号',
  `nickname` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户昵称',
  `password` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户密码',
  `avatar` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户头像',
  `role` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色主键',
  `salt` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '加密盐巴',
  `sort` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序编号',
  `is_multipoint` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '多端登录: 0=否, 1=是',
  `is_disable` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否禁用: 0=否, 1=是',
  `is_delete` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除: 0=否, 1=是',
  `last_login_ip` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '最后登录IP',
  `last_login_time` DATETIME NOT NULL COMMENT '最后登录',
  `create_time` DATETIME NOT NULL COMMENT '创建时间',
	`update_time` DATETIME NOT NULL COMMENT '更新时间',
	`delete_time` DATETIME  NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统管理成员表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_system_auth_admin
-- ----------------------------
INSERT INTO `x_system_auth_admin` VALUES (1, 1, 3, 'admin', 'admin', '7ca7e19452aa2366068785be5c7ded35', '/image/20232911/27699a1e991e4ba1ba13ee6d45185db6.jpg', '0', 'HOEp0', 1, 1, 0, 0, '127.0.0.1', "2024-01-02 03:04:05", "2024-01-02 03:04:05", "2024-01-02 03:04:05", null);
INSERT INTO `x_system_auth_admin` VALUES (2, 2, 3, 'zhihuibu01', '指挥部01', 'ea7e7f97957b7cdd2b245abc31cebaa4', '/image/20242102/1c7428a6cfdf4151a1b9cbd49ec0702c.png', '1', 'JUD5n', 1, 0, 0, 0, '127.0.0.1', "2024-01-02 03:04:05", "2024-01-02 03:04:05", "2024-01-02 03:04:05", null);

-- ----------------------------
-- Table structure for x_system_auth_dept
-- ----------------------------
DROP TABLE IF EXISTS `x_system_auth_dept`;
CREATE TABLE `x_system_auth_dept`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `pid` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '上级主键',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '部门名称',
  `duty_id` int(10) NULL DEFAULT 0 COMMENT '负责人id',
  `duty` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '负责人名',
  `mobile` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '联系电话',
  `sort` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序编号',
  `is_stop` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否禁用: 0=否, 1=是',
  `is_delete` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除: 0=否, 1=是',
  `create_time` DATETIME NOT NULL COMMENT '创建时间',
	`update_time` DATETIME NOT NULL COMMENT '更新时间',
	`delete_time` DATETIME  NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统部门管理表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_system_auth_dept
-- ----------------------------
INSERT INTO `x_system_auth_dept` VALUES (1, 0, '默认部门', 1, 'admin', '18327647788', 10, 0, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05", null);
INSERT INTO `x_system_auth_dept` VALUES (2, 1, '指挥部', 2, '指挥部01', '17608390654', 0, 0, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05", null);
INSERT INTO `x_system_auth_dept` VALUES (3, 2, '指挥部子级', 0, '', '', 0, 0, 1, "2024-01-02 03:04:05", "2024-01-02 03:04:05", null);

-- ----------------------------
-- Table structure for x_system_auth_menu
-- ----------------------------
DROP TABLE IF EXISTS `x_system_auth_menu`;
CREATE TABLE `x_system_auth_menu`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `pid` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '上级菜单',
  `menu_type` char(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '权限类型: M=目录，C=菜单，A=按钮',
  `menu_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '菜单名称',
  `menu_icon` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '菜单图标',
  `menu_sort` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '菜单排序',
  `perms` varchar(10000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '权限标识',
  `paths` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '路由地址',
  `component` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '前端组件',
  `selected` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '选中路径',
  `params` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '路由参数',
  `is_cache` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否缓存: 0=否, 1=是',
  `is_show` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '是否显示: 0=否, 1=是',
  `is_disable` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否禁用: 0=否, 1=是',
  `create_time` DATETIME NOT NULL COMMENT '创建时间',
	`update_time` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 814 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统菜单管理表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_system_auth_menu
-- ----------------------------
INSERT INTO `x_system_auth_menu` VALUES (1, 0, 'C', '工作台', 'el-icon-Monitor', 50, 'admin:common:index:console', 'workbench', 'workbench/index', '', '', 1, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (100, 500, 'M', '权限管理', 'el-icon-Lock', 44, '', 'permission', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (101, 100, 'C', '管理员', 'local-icon-wode', 0, 'admin:system:admin:list', 'admin', 'setting/permission/admin/index', '', '', 1, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (102, 101, 'A', '管理员详情', '', 0, 'admin:system:admin:detail', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (103, 101, 'A', '管理员新增', '', 0, 'admin:system:admin:add', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (104, 101, 'A', '管理员编辑', '', 0, 'admin:system:admin:edit', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (105, 101, 'A', '管理员删除', '', 0, 'admin:system:admin:del', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (106, 101, 'A', '管理员状态', '', 0, 'admin:system:admin:disable', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (110, 100, 'C', '角色管理', 'el-icon-Female', 0, 'admin:system:role:list', 'role', 'setting/permission/role/index', '', '', 1, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (111, 110, 'A', '角色详情', '', 0, 'admin:system:role:detail', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (112, 110, 'A', '角色新增', '', 0, 'admin:system:role:add', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (113, 110, 'A', '角色编辑', '', 0, 'admin:system:role:edit', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (114, 110, 'A', '角色删除', '', 0, 'admin:system:role:del', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (120, 100, 'C', '菜单管理', 'el-icon-Operation', 0, 'admin:system:menu:list', 'menu', 'setting/permission/menu/index', '', '', 1, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (121, 120, 'A', '菜单详情', '', 0, 'admin:system:menu:detail', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (122, 120, 'A', '菜单新增', '', 0, 'admin:system:menu:add', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (123, 120, 'A', '菜单编辑', '', 0, 'admin:system:menu:edit', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (124, 120, 'A', '菜单删除', '', 0, 'admin:system:menu:del', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (130, 500, 'M', '组织管理', 'el-icon-OfficeBuilding', 45, '', 'organization', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (131, 130, 'C', '部门管理', 'el-icon-Coordinate', 0, 'admin:system:dept:list', 'department', 'organization/department/index', '', '', 1, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (132, 131, 'A', '部门详情', '', 0, 'admin:system:dept:detail', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (133, 131, 'A', '部门新增', '', 0, 'admin:system:dept:add', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (134, 131, 'A', '部门编辑', '', 0, 'admin:system:dept:edit', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (135, 131, 'A', '部门删除', '', 0, 'admin:system:dept:del', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (140, 130, 'C', '岗位管理', 'el-icon-PriceTag', 0, 'admin:system:post:list', 'post', 'organization/post/index', '', '', 1, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (141, 140, 'A', '岗位详情', '', 0, 'admin:system:post:detail', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (142, 140, 'A', '岗位新增', '', 0, 'admin:system:post:add', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (143, 140, 'A', '岗位编辑', '', 0, 'admin:system:post:edit', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (144, 140, 'A', '岗位删除', '', 0, 'admin:system:post:del', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (200, 0, 'M', '其它管理', '', 0, '', '', '', '', '', 0, 0, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (201, 200, 'M', '图库管理', '', 0, '', '', '', '', '', 0, 0, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (202, 201, 'A', '文件列表', '', 0, 'admin:common:album:albumList', '', '', '', '', 0, 0, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (203, 201, 'A', '文件命名', '', 0, 'admin:common:album:albumRename', '', '', '', '', 0, 0, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (204, 201, 'A', '文件移动', '', 0, 'admin:common:album:albumMove', '', '', '', '', 0, 0, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (205, 201, 'A', '文件删除', '', 0, 'admin:common:album:albumDel', '', '', '', '', 0, 0, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (206, 201, 'A', '分类列表', '', 0, 'admin:common:album:cateList', '', '', '', '', 0, 0, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (207, 201, 'A', '分类新增', '', 0, 'admin:common:album:cateAdd', '', '', '', '', 0, 0, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (208, 201, 'A', '分类命名', '', 0, 'admin:common:album:cateRename', '', '', '', '', 0, 0, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (209, 201, 'A', '分类删除', '', 0, 'admin:common:album:cateDel', '', '', '', '', 0, 0, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (215, 200, 'M', '上传管理', '', 0, '', '', '', '', '', 0, 0, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (216, 215, 'A', '上传图片', '', 0, 'admin:common:upload:image', '', '', '', '', 0, 0, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (217, 215, 'A', '上传视频', '', 0, 'admin:common:upload:video', '', '', '', '', 0, 0, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (500, 0, 'M', '系统设置', 'el-icon-Setting', 0, '', 'setting', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (501, 500, 'M', '网站设置', 'el-icon-Basketball', 10, '', 'website', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (502, 501, 'C', '网站信息', '', 0, 'admin:setting:website:detail', 'information', 'setting/website/information', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (503, 502, 'A', '保存配置', '', 0, 'admin:setting:website:save', '', '', '', '', 0, 0, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (505, 501, 'C', '网站备案', '', 0, 'admin:setting:copyright:detail', 'filing', 'setting/website/filing', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (506, 505, 'A', '备案保存', '', 0, 'admin:setting:copyright:save', '', 'setting/website/protocol', '', '', 0, 0, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (510, 501, 'C', '政策协议', '', 0, 'admin:setting:protocol:detail', 'protocol', 'setting/website/protocol', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (511, 510, 'A', '协议保存', '', 0, 'admin:setting:protocol:save', '', '', '', '', 0, 0, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (515, 500, 'C', '字典管理', 'el-icon-Box', 0, 'admin:setting:dict:type:list', 'dict', 'setting/dict/type/index', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (516, 515, 'A', '字典类型新增', '', 0, 'admin:setting:dict:type:add', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (517, 515, 'A', '字典类型编辑', '', 0, 'admin:setting:dict:type:edit', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (518, 515, 'A', '字典类型删除', '', 0, 'admin:setting:dict:type:del', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (519, 500, 'C', '字典数据管理', '', 0, 'admin:setting:dict:data:list', 'dict/data', 'setting/dict/data/index', '/dev_tools/dict', '', 0, 0, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (520, 515, 'A', '字典数据新增', '', 0, 'admin:setting:dict:data:add', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (521, 515, 'A', '字典数据编辑', '', 0, 'admin:setting:dict:data:edit', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (522, 515, 'A', '字典数据删除', '', 0, 'admin:setting:dict:data:del', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (550, 500, 'M', '系统维护', 'el-icon-SetUp', 0, '', 'system', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (551, 550, 'C', '系统环境', '', 0, 'admin:monitor:server', 'environment', 'setting/system/environment', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (552, 550, 'C', '系统缓存', '', 0, 'admin:monitor:cache', 'cache', 'setting/system/cache', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (553, 550, 'C', '系统日志', '', 0, 'admin:system:log:operate', 'journal', 'setting/system/journal', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (600, 0, 'M', '开发工具', 'el-icon-EditPen', 0, '', 'dev_tools', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (610, 600, 'C', '代码生成器', 'el-icon-DocumentAdd', 0, 'admin:gen:list', 'code', 'dev_tools/code/index', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (611, 610, 'A', '导入数据表', '', 0, 'admin:gen:importTable', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (612, 610, 'A', '生成代码', '', 0, 'admin:gen:genCode', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (613, 610, 'A', '下载代码', '', 0, 'admin:gen:downloadCode', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (614, 610, 'A', '预览代码', '', 0, 'admin:gen:previewCode', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (616, 610, 'A', '同步表结构', '', 0, 'admin:gen:syncTable', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (617, 610, 'A', '删除数据表', '', 0, 'admin:gen:delTable', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (618, 610, 'A', '数据表详情', '', 0, 'admin:gen:detail', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (700, 0, 'M', '素材管理', 'el-icon-Picture', 43, '', 'material', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (701, 700, 'C', '素材中心', 'el-icon-PictureRounded', 0, '', 'index', 'material/index', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (775, 600, 'C', '代码生成器编辑', 'el-icon-EditPen', 0, 'admin:gen:editTable', 'dev_tools/code/edit', 'dev_tools/code/edit', '', '', 0, 0, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (776, 778, 'C', '流程模板', '', 0, 'admin:flow:flow_template:list', 'flow_template/index', 'flow/flow_template/index', '', '', 1, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (777, 778, 'C', '我的流程', '', 0, '', 'flow_apply/index', 'flow/flow_apply/index', '', '', 1, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (778, 0, 'M', '审批流', 'el-icon-Coordinate', 0, '', 'flow', '', '', '', 1, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (780, 778, 'C', '待处理', '', 0, 'admin:flow:flow_history:list', 'flow_history/todo', 'flow/flow_history/todo', '', '', 1, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (781, 778, 'C', '已处理', '', 0, 'admin:flow:flow_history:list', 'flow_history/done', 'flow/flow_history/done', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (782, 778, 'C', '已完成流程', '', 0, 'admin:flow:flow_history:list', 'flow_apply/finish', 'flow/flow_apply/finish', '', '', 1, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (783, 500, 'C', '项目监控', 'el-icon-Notebook', 0, 'admin:monitor_project:list', 'monitor_project/index', 'monitor_project/index', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (784, 500, 'C', '监控用户端', '', 0, '', 'monitor_client/index', 'monitor_client/index', '', '', 1, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (785, 794, 'A', '错误收集error添加', '', 0, 'admin:monitor_web:add', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (786, 794, 'A', '错误收集error编辑', '', 0, 'admin:monitor_web:edit', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (787, 794, 'A', '错误收集error删除', '', 0, 'admin:monitor_web:del', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (788, 794, 'A', '错误收集error列表', '', 0, 'admin:monitor_web:list', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (789, 794, 'A', '错误收集error全部列表', '', 0, 'admin:monitor_web:listAll', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (790, 794, 'A', '错误收集error详情', '', 0, 'admin:monitor_web:detail', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (791, 794, 'A', '错误收集error导出excel', '', 0, 'admin:monitor_web:ExportFile', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (792, 794, 'A', '错误收集error导入excel', '', 0, 'admin:monitor_web:ImportFile', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (794, 500, 'C', '错误收集error', '', 0, '', 'monitor_web/index', 'monitor_web/index', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (795, 776, 'A', '流程模板添加', '', 0, 'admin:flow:flow_template:add', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (796, 776, 'A', '流程模板编辑', '', 0, 'admin:flow:flow_template:edit', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (797, 776, 'A', '流程模板删除', '', 0, 'admin:flow:flow_template:del', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (798, 776, 'A', '流程模板列表', '', 0, 'admin:flow:flow_template:list', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (799, 776, 'A', '流程模板全部列表', '', 0, 'admin:flow:flow_template:listAll', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (800, 776, 'A', '流程模板详情', '', 0, 'admin:flow:flow_template:detail', '', '', '', '', 0, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (801, 777, 'A', '流程详情', '', 0, 'admin:flow:flow_apply:detail', '', '', '', '', 1, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (802, 777, 'A', '流程添加', '', 0, 'admin:flow:flow_apply:add', '', '', '', '', 1, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (803, 777, 'A', '流程编辑', '', 0, 'admin:flow:flow_apply:edit', '', '', '', '', 1, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (804, 777, 'A', '流程删除', '', 0, 'admin:flow:flow_apply:del', '', '', '', '', 1, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (805, 777, 'A', '全部流程', '', 0, 'admin:flow:flow_history:listAll', '', '', '', '', 1, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (806, 777, 'A', '通过流程', '', 0, 'admin:flow:flow_history:pass', '', '', '', '', 1, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (807, 777, 'A', '拒接流程', '', 0, 'admin:flow:flow_history:back', '', '', '', '', 1, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (808, 777, 'A', '下一个流程', '', 0, 'admin:flow:flow_history:next_node', '', '', '', '', 1, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (809, 777, 'A', '获取审批人', '', 0, 'admin:flow:flow_history:get_approver', '', '', '', '', 1, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (810, 777, 'A', '流程列表', '', 0, 'admin:flow:flow_apply:list', '', '', '', '', 1, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (811, 780, 'A', '审批记录列表', '', 0, 'admin:flow:flow_history:list', '', '', '', '', 1, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_auth_menu` VALUES (812, 780, 'A', '审批记录详情', '', 0, 'admin:flow:flow_history:detail', '', '', '', '', 1, 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");

-- ----------------------------
-- Table structure for x_system_auth_perm
-- ----------------------------
DROP TABLE IF EXISTS `x_system_auth_perm`;
CREATE TABLE `x_system_auth_perm`  (
  `id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '主键',
  `role_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '角色ID',
  `menu_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '菜单ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统角色菜单表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_system_auth_perm
-- ----------------------------
INSERT INTO `x_system_auth_perm` VALUES ('02849ad51fd24d0f959ad0629e9e1fe4', 1, 522);
INSERT INTO `x_system_auth_perm` VALUES ('04487e8700f2406abfad957f3732ee30', 1, 784);
INSERT INTO `x_system_auth_perm` VALUES ('04a34fc8ae12416c95595ed4c93828b0', 1, 797);
INSERT INTO `x_system_auth_perm` VALUES ('0c2921d838874c7091f8e312d8985193', 1, 804);
INSERT INTO `x_system_auth_perm` VALUES ('115dd9eeb1404d89b2fb6827de49a737', 1, 812);
INSERT INTO `x_system_auth_perm` VALUES ('11ae6d9c46364ad093149803836758c2', 1, 780);
INSERT INTO `x_system_auth_perm` VALUES ('12023b0aaae14a48bcfb9011f069ed1e', 1, 217);
INSERT INTO `x_system_auth_perm` VALUES ('13d8c8ead1ca49e9aab50977671ebbce', 1, 798);
INSERT INTO `x_system_auth_perm` VALUES ('142baed090b141b2a9ede7d96bf4df78', 1, 552);
INSERT INTO `x_system_auth_perm` VALUES ('154aaec657a74034b20eea3991fe5dbb', 1, 205);
INSERT INTO `x_system_auth_perm` VALUES ('18bb6d0fd189410fad80cb0b64e5f74f', 1, 140);
INSERT INTO `x_system_auth_perm` VALUES ('1ac303aef5794defaab58613ae1efc7a', 1, 502);
INSERT INTO `x_system_auth_perm` VALUES ('1e048396a2924be8bef72822f460bd9f', 1, 131);
INSERT INTO `x_system_auth_perm` VALUES ('1f9e87e5ffa04992bded01c6dc2350aa', 1, 550);
INSERT INTO `x_system_auth_perm` VALUES ('21c7ca19973549aca58612e7cb1fbeb2', 1, 505);
INSERT INTO `x_system_auth_perm` VALUES ('2437bd9c41ee4222bcaf6df3407fb20a', 1, 100);
INSERT INTO `x_system_auth_perm` VALUES ('25866e25dbb84299acac12ca92745755', 1, 123);
INSERT INTO `x_system_auth_perm` VALUES ('259e59808a7e45eaa5f3a4bad9c1e0aa', 1, 506);
INSERT INTO `x_system_auth_perm` VALUES ('28b7b9c680514dc2bbd7f548c35a2d73', 1, 792);
INSERT INTO `x_system_auth_perm` VALUES ('2aff7d76f4fa4f058d93be667fadbef8', 1, 785);
INSERT INTO `x_system_auth_perm` VALUES ('2f2af99726b444cab4d08fd4a2d3af0f', 1, 794);
INSERT INTO `x_system_auth_perm` VALUES ('333e4bab351246a8908fbe6b4e73aef5', 1, 519);
INSERT INTO `x_system_auth_perm` VALUES ('365393451dcb485eb18adb3422d89dc6', 1, 781);
INSERT INTO `x_system_auth_perm` VALUES ('36ce29f4f2de480495ece60f51ff3e94', 1, 791);
INSERT INTO `x_system_auth_perm` VALUES ('370c8eb2e71442ff9f652c1f5021b423', 1, 133);
INSERT INTO `x_system_auth_perm` VALUES ('3b659baff55d42ffa688c1f6fc984d70', 1, 617);
INSERT INTO `x_system_auth_perm` VALUES ('424a2b7dd0d54db28045edb064e7747e', 1, 500);
INSERT INTO `x_system_auth_perm` VALUES ('47e6e92c3f8d47a0a7ecf6055c75bcd6', 1, 208);
INSERT INTO `x_system_auth_perm` VALUES ('4bf2e6a46d164a7d8018c8df067126ab', 1, 207);
INSERT INTO `x_system_auth_perm` VALUES ('4c10437c35d04c3c88e2d8a63e98c925', 1, 134);
INSERT INTO `x_system_auth_perm` VALUES ('4c2b595615ea49628e989a9266e3e253', 1, 204);
INSERT INTO `x_system_auth_perm` VALUES ('4cd78098aea84e5e9e1d79452a607658', 1, 808);
INSERT INTO `x_system_auth_perm` VALUES ('4d993bc9a28e462dbad7be6101b50535', 1, 811);
INSERT INTO `x_system_auth_perm` VALUES ('4de195ebe48f47fea47ef972b227720e', 1, 113);
INSERT INTO `x_system_auth_perm` VALUES ('4deb9fc89d634115b3b7d30b571aa4dd', 1, 102);
INSERT INTO `x_system_auth_perm` VALUES ('4f9f5cbe148e4c9e82bd03ed24bb934a', 1, 805);
INSERT INTO `x_system_auth_perm` VALUES ('5272e3b59007467c92a5b953ffc6900c', 1, 520);
INSERT INTO `x_system_auth_perm` VALUES ('52a5530f482f4ec3a07aa7941df40dff', 1, 700);
INSERT INTO `x_system_auth_perm` VALUES ('56664480a15a4ee898ac9aee1c1014f3', 1, 802);
INSERT INTO `x_system_auth_perm` VALUES ('59618b0354574f8d8949de4efa170471', 1, 776);
INSERT INTO `x_system_auth_perm` VALUES ('5b5a0b548bfe4e1fb51b0592122eb048', 1, 616);
INSERT INTO `x_system_auth_perm` VALUES ('5cafd9be52b6436090ff76520e454c23', 1, 611);
INSERT INTO `x_system_auth_perm` VALUES ('6083696b56f047c3923a1284278925a5', 1, 104);
INSERT INTO `x_system_auth_perm` VALUES ('632edb42cb694d8d9534e942d6338293', 1, 206);
INSERT INTO `x_system_auth_perm` VALUES ('67cfb9970a954c55b06ba50a62b47611', 1, 143);
INSERT INTO `x_system_auth_perm` VALUES ('6bfb96696d204ae29bbe6f9f3742488e', 1, 701);
INSERT INTO `x_system_auth_perm` VALUES ('6c5f2568cd074b2086bfdf3db9e999ad', 1, 787);
INSERT INTO `x_system_auth_perm` VALUES ('6cc0625e26424d0b88c202fd6f38af88', 1, 142);
INSERT INTO `x_system_auth_perm` VALUES ('6e0896d04ba34f2f901a7ed764957c6e', 1, 803);
INSERT INTO `x_system_auth_perm` VALUES ('7197d7b4eafd47f3819f70fd67b95169', 1, 103);
INSERT INTO `x_system_auth_perm` VALUES ('776061d01f0743af97431d76a7d29a3b', 1, 510);
INSERT INTO `x_system_auth_perm` VALUES ('7b6a55357b524c61950a558ebb050383', 1, 807);
INSERT INTO `x_system_auth_perm` VALUES ('7cd0fdb89fb64622bcc20041749ad98a', 1, 130);
INSERT INTO `x_system_auth_perm` VALUES ('836584161e314db4a7c9d3e17b10dffb', 1, 788);
INSERT INTO `x_system_auth_perm` VALUES ('8424d4ee6ec844b1add5e9386b3aa936', 1, 610);
INSERT INTO `x_system_auth_perm` VALUES ('88ee4db9432c4610a1158e698cc02c7c', 1, 518);
INSERT INTO `x_system_auth_perm` VALUES ('8995c474475f43c492c343edf21e4f3f', 1, 614);
INSERT INTO `x_system_auth_perm` VALUES ('8b1776a8436146e4963565cff4231b78', 1, 795);
INSERT INTO `x_system_auth_perm` VALUES ('927e29c0b369442bb0e88cb77a4e4d5b', 1, 202);
INSERT INTO `x_system_auth_perm` VALUES ('9430404611924c809be637162f445b88', 1, 600);
INSERT INTO `x_system_auth_perm` VALUES ('946d97b5bf284d12a058d3eb92b410f9', 1, 120);
INSERT INTO `x_system_auth_perm` VALUES ('97f0c2ad897c4e03a5be79d4e6a8bc4a', 1, 778);
INSERT INTO `x_system_auth_perm` VALUES ('9b06d4afc00c4198b19b7159cf82f796', 1, 132);
INSERT INTO `x_system_auth_perm` VALUES ('9cac9258695947988efbea8188633dc4', 1, 209);
INSERT INTO `x_system_auth_perm` VALUES ('9e4e502124a541fd8af7dc740e1bcf5e', 1, 141);
INSERT INTO `x_system_auth_perm` VALUES ('a3852632d33f41f9a34748e68cc3f73a', 1, 105);
INSERT INTO `x_system_auth_perm` VALUES ('a9a1234793464bbcb2872964598035bf', 1, 806);
INSERT INTO `x_system_auth_perm` VALUES ('ab2a34c2b175448487971f5d10d4fd84', 1, 201);
INSERT INTO `x_system_auth_perm` VALUES ('ab39e031d9554d0f8c9bd23a22d2992b', 1, 799);
INSERT INTO `x_system_auth_perm` VALUES ('ab6044f2957e40ec9dab8ef4604a980a', 1, 517);
INSERT INTO `x_system_auth_perm` VALUES ('b22848d6194449719affc0378c5defa8', 1, 501);
INSERT INTO `x_system_auth_perm` VALUES ('b26e5a7ab57b4191bca9b6ef05d7e766', 1, 515);
INSERT INTO `x_system_auth_perm` VALUES ('b7a30ea8a3534fb48778f10a6cbff2eb', 1, 521);
INSERT INTO `x_system_auth_perm` VALUES ('b8ce62b38c0148e0ad7d2f83b077fb52', 1, 786);
INSERT INTO `x_system_auth_perm` VALUES ('b90ad473fa0b431b8b7f21f282e27f63', 1, 553);
INSERT INTO `x_system_auth_perm` VALUES ('bb27c9499e7a4c93a15e5a448291e9aa', 1, 122);
INSERT INTO `x_system_auth_perm` VALUES ('bc079980f626418e8d7236480d11b507', 1, 783);
INSERT INTO `x_system_auth_perm` VALUES ('bd06cabd21f44d3786f2a6e9b1131d5d', 1, 110);
INSERT INTO `x_system_auth_perm` VALUES ('bec81045070c462986c8c6c4cc7555ca', 1, 613);
INSERT INTO `x_system_auth_perm` VALUES ('c20a057a443c4e2f990819800c23d9f5', 1, 800);
INSERT INTO `x_system_auth_perm` VALUES ('c5a3a048e8fe40678c411e76bd3a7c69', 1, 124);
INSERT INTO `x_system_auth_perm` VALUES ('c80498f6a0424195afe7e40bcd287262', 1, 516);
INSERT INTO `x_system_auth_perm` VALUES ('cc3020f5228e468d87f91e83ad61e818', 1, 789);
INSERT INTO `x_system_auth_perm` VALUES ('ce6e84d612c942dba783b5d2df6ed595', 1, 612);
INSERT INTO `x_system_auth_perm` VALUES ('cfa5702a424347099eea48f28d57714f', 1, 782);
INSERT INTO `x_system_auth_perm` VALUES ('d0d98c528c9840a89f4aa04ed0415c10', 1, 203);
INSERT INTO `x_system_auth_perm` VALUES ('d1232298befc44d5beca54260e93761b', 1, 216);
INSERT INTO `x_system_auth_perm` VALUES ('d14f51d7515e4bf0a092d3bf0a4225ac', 1, 790);
INSERT INTO `x_system_auth_perm` VALUES ('d1784f85d159497ca13dfc7a1749e612', 1, 135);
INSERT INTO `x_system_auth_perm` VALUES ('d2b7aa04883d47f0bbb9414d5e736205', 1, 503);
INSERT INTO `x_system_auth_perm` VALUES ('d5d6a09bc7494c919e03472dca3faad3', 1, 801);
INSERT INTO `x_system_auth_perm` VALUES ('d6388c117a364080a7d169ff60afe5f2', 1, 111);
INSERT INTO `x_system_auth_perm` VALUES ('e0af9fc2032a4c4a8284a420c1dfcf20', 1, 144);
INSERT INTO `x_system_auth_perm` VALUES ('e401edd59c5b4bb8a4a1cd2391c28476', 1, 1);
INSERT INTO `x_system_auth_perm` VALUES ('e475f99a770f4fa791c3b2b139c9f98c', 1, 810);
INSERT INTO `x_system_auth_perm` VALUES ('e6b216b706f44aefaf43a154fda4d1cd', 1, 618);
INSERT INTO `x_system_auth_perm` VALUES ('e813a85f488741ed90c77e48b20b3f3e', 1, 809);
INSERT INTO `x_system_auth_perm` VALUES ('ed406503650c44849f6b82bfb78ff432', 1, 775);
INSERT INTO `x_system_auth_perm` VALUES ('edcbeefe1c6d457aab65a5640c5e8802', 1, 112);
INSERT INTO `x_system_auth_perm` VALUES ('ee441f534ab049a5adb590bcbedda29e', 1, 200);
INSERT INTO `x_system_auth_perm` VALUES ('f21e0aeb1a504690a933b83b2ae3c8d5', 1, 796);
INSERT INTO `x_system_auth_perm` VALUES ('f21e3c5152064ba79c3d2ee0a3036022', 1, 215);
INSERT INTO `x_system_auth_perm` VALUES ('f2696be26aa44060bb3330deb301f2c0', 1, 114);
INSERT INTO `x_system_auth_perm` VALUES ('f42ae39a8c5e464fa1301f6dd1d3bae6', 1, 121);
INSERT INTO `x_system_auth_perm` VALUES ('f91a6aacb7114e569f4953e224145335', 1, 511);
INSERT INTO `x_system_auth_perm` VALUES ('f932b361ab53400c86f73a183d1cd248', 1, 101);
INSERT INTO `x_system_auth_perm` VALUES ('fc973331a38249889b3743dde4abede9', 1, 777);
INSERT INTO `x_system_auth_perm` VALUES ('fcbd54d2a0e04e82b5847a1426ed29ba', 1, 551);
INSERT INTO `x_system_auth_perm` VALUES ('feeddde0528c4660bf5ddbcd7e6c46ac', 1, 106);

-- ----------------------------
-- Table structure for x_system_auth_post
-- ----------------------------
DROP TABLE IF EXISTS `x_system_auth_post`;
CREATE TABLE `x_system_auth_post`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `code` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '岗位编码',
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '岗位名称',
  `remarks` varchar(250) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '岗位备注',
  `sort` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '岗位排序',
  `is_stop` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否停用: 0=否, 1=是',
  `is_delete` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除: 0=否, 1=是',
    `create_time` DATETIME NOT NULL COMMENT '创建时间',
	`update_time` DATETIME NOT NULL COMMENT '更新时间',
	`delete_time` DATETIME  NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统岗位管理表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_system_auth_post
-- ----------------------------
INSERT INTO `x_system_auth_post` VALUES (2, 'gw0001', '默认岗位', '', 3, 0, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05", null);
INSERT INTO `x_system_auth_post` VALUES (3, 'zhihuibu01', '指挥部岗位', '', 0, 0, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05", null);
INSERT INTO `x_system_auth_post` VALUES (4, '1', '11', '', 0, 0, 1, "2024-01-02 03:04:05", "2024-01-02 03:04:05", null);
INSERT INTO `x_system_auth_post` VALUES (5, '2', '2', '2', 1, 0, 1, "2024-01-02 03:04:05", "2024-01-02 03:04:05", null);

-- ----------------------------
-- Table structure for x_system_auth_role
-- ----------------------------
DROP TABLE IF EXISTS `x_system_auth_role`;
CREATE TABLE `x_system_auth_role`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色名称',
  `remark` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注信息',
  `sort` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '角色排序',
  `is_disable` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否禁用: 0=否, 1=是',
    `create_time` DATETIME NOT NULL COMMENT '创建时间',
	`update_time` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统角色管理表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_system_auth_role
-- ----------------------------
INSERT INTO `x_system_auth_role` VALUES (1, '审核员', '审核数据', 1, 0, "2024-01-02 03:04:05", "2024-01-02 03:04:05");

-- ----------------------------
-- Table structure for x_system_config
-- ----------------------------
DROP TABLE IF EXISTS `x_system_config`;
CREATE TABLE `x_system_config`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `type` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '类型',
  `name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '键',
  `value` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '值',
`create_time` DATETIME NOT NULL COMMENT '创建时间',
	`update_time` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 81 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统全局配置表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_system_config
-- ----------------------------
INSERT INTO `x_system_config` VALUES (1, 'storage', 'default', '', "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_config` VALUES (2, 'storage', 'local', '{\"name\":\"本地存储\"}', "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_config` VALUES (3, 'storage', 'qiniu', '{\"name\":\"七牛云存储\",\"bucket\":\"\",\"secretKey\":\"\",\"accessKey\":\"\",\"domain\":\"\"}', "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_config` VALUES (4, 'storage', 'aliyun', '{\"name\":\"阿里云存储\",\"bucket\":\"\",\"secretKey\":\"\",\"accessKey\":\"\",\"domain\":\"\"}', "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_config` VALUES (5, 'storage', 'qcloud', '{\"name\":\"腾讯云存储\",\"bucket\":\"\",\"secretKey\":\"\",\"accessKey\":\"\",\"domain\":\"\",\"region\":\"\"}', "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_config` VALUES (6, 'sms', 'default', 'aliyun', "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_config` VALUES (7, 'sms', 'aliyun', '{\"name\":\"阿里云短信\",\"alias\":\"aliyun\",\"sign\":\"\",\"appKey\":\"\",\"secretKey\":\"\"}', "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_config` VALUES (8, 'sms', 'tencent', '{\"name\":\"腾讯云短信\",\"alias\":\"tencent\",\"sign\":\"\",\"appId\":\"\",\"secretId\":\"\",\"secretKey\":\"\"}', "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_config` VALUES (9, 'sms', 'huawei', '{\"name\":\"华为云短信\",\"alias\":\"huawei\"}', "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_config` VALUES (10, 'website', 'name', 'x_admin开源系统', "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_config` VALUES (11, 'website', 'logo', '/api/static/backend_logo.png', "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_config` VALUES (12, 'website', 'favicon', '/api/static/backend_favicon.ico', "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_config` VALUES (13, 'website', 'backdrop', '/api/static/backend_backdrop.png', "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_config` VALUES (14, 'website', 'copyright', '[{\"name\":\"x_admin开源管理系统\",\"link\":\"http://www.beian.gov.cn\"}]', "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_config` VALUES (15, 'website', 'shopName', 'x_admin开源管理系统', "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_config` VALUES (16, 'website', 'shopLogo', '/api/static/shop_logo.png', "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_config` VALUES (17, 'protocol', 'service', '{\"name\":\"服务协议\",\"content\":\"\\u003cp\\u003e服务协议\\u003c/p\\u003e\"}', "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_config` VALUES (18, 'protocol', 'privacy', '{\"name\":\"隐私协议\",\"content\":\"\\u003cp\\u003e隐私协议\\u003c/p\\u003e\"}', "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_config` VALUES (19, 'tabbar', 'style', '{\"defaultColor\":\"#4A5DFF\",\"selectedColor\":\"#EA5455\"}', "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_config` VALUES (20, 'search', 'isHotSearch', '0', "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_config` VALUES (30, 'h5_channel', 'status', '1', "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_config` VALUES (31, 'h5_channel', 'close', '0', "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_config` VALUES (32, 'h5_channel', 'url', '', "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_config` VALUES (40, 'mp_channel', 'name', '', "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_config` VALUES (41, 'mp_channel', 'primaryId', '', "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_config` VALUES (42, 'mp_channel', 'appId', '', "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_config` VALUES (43, 'mp_channel', 'appSecret', '', "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_config` VALUES (44, 'mp_channel', 'qrCode', '', "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_config` VALUES (50, 'wx_channel', 'appId', '', "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_config` VALUES (51, 'wx_channel', 'appSecret', '', "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_config` VALUES (55, 'oa_channel', 'name', '', "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_config` VALUES (56, 'oa_channel', 'primaryId', ' ', "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_config` VALUES (57, 'oa_channel', 'qrCode', '', "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_config` VALUES (58, 'oa_channel', 'appId', '', "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_config` VALUES (59, 'oa_channel', 'appSecret', '', "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_config` VALUES (60, 'oa_channel', 'url', '', "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_config` VALUES (61, 'oa_channel', 'token', '', "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_config` VALUES (62, 'oa_channel', 'encodingAesKey', '', "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_config` VALUES (63, 'oa_channel', 'encryptionType', '1', "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_config` VALUES (64, 'oa_channel', 'menus', '[]', "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_config` VALUES (70, 'login', 'loginWay', '1,2', "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_config` VALUES (71, 'login', 'forceBindMobile', '0', "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_config` VALUES (72, 'login', 'openAgreement', '1', "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_config` VALUES (73, 'login', 'openOtherAuth', '1', "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_config` VALUES (74, 'login', 'autoLoginAuth', '1,2', "2024-01-02 03:04:05", "2024-01-02 03:04:05");
INSERT INTO `x_system_config` VALUES (80, 'user', 'defaultAvatar', '/api/static/default_avatar.png', "2024-01-02 03:04:05", "2024-01-02 03:04:05");

-- ----------------------------
-- Table structure for x_system_log_login
-- ----------------------------
DROP TABLE IF EXISTS `x_system_log_login`;
CREATE TABLE `x_system_log_login`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '注解',
  `admin_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '管理员ID',
  `username` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '登录账号',
  `ip` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '登录地址',
  `os` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '操作系统',
  `browser` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '浏览器',
  `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '操作状态: 1=成功, 2=失败',
  `create_time` DATETIME NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 139 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统登录日志表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_system_log_login
-- ----------------------------
INSERT INTO `x_system_log_login` VALUES (135, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1,"2024-01-02 03:04:05");
INSERT INTO `x_system_log_login` VALUES (136, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1,"2024-01-02 03:04:05");
INSERT INTO `x_system_log_login` VALUES (137, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1,"2024-01-02 03:04:05");
INSERT INTO `x_system_log_login` VALUES (138, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1,"2024-01-02 03:04:05");

-- ----------------------------
-- Table structure for x_system_log_operate
-- ----------------------------
DROP TABLE IF EXISTS `x_system_log_operate`;
CREATE TABLE `x_system_log_operate`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `admin_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '操作人ID',
  `type` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '请求类型: GET/POST/PUT',
  `title` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '操作标题',
  `ip` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '请求IP',
  `url` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '请求接口',
  `method` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '请求方法',
  `args` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '请求参数',
  `error` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '错误信息',
  `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '执行状态: 1=成功, 2=失败',
  `start_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '开始时间',
  `end_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '结束时间',
  `task_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '执行耗时',
  `create_time` DATETIME NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 656 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统操作日志表' ROW_FORMAT = DYNAMIC;


-- ----------------------------
-- Table structure for x_system_log_sms
-- ----------------------------
DROP TABLE IF EXISTS `x_system_log_sms`;
CREATE TABLE `x_system_log_sms`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
  `scene` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '场景编号',
  `mobile` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '手机号码',
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '发送内容',
  `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '发送状态：[0=发送中, 1=发送成功, 2=发送失败]',
  `results` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '短信结果',
  `send_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '发送时间',
    `create_time` DATETIME NOT NULL COMMENT '创建时间',
	`update_time` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统短信日志表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_system_log_sms
-- ----------------------------

-- ----------------------------
-- Table structure for x_user
-- ----------------------------
DROP TABLE IF EXISTS `x_user`;
CREATE TABLE `x_user`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `sn` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '编号',
  `avatar` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '头像',
  `real_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '真实姓名',
  `nickname` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户昵称',
  `username` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户账号',
  `password` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户密码',
  `mobile` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户电话',
  `salt` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '加密盐巴',
  `sex` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户性别: [1=男, 2=女]',
  `channel` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '注册渠道: [1=微信小程序, 2=微信公众号, 3=手机H5, 4=电脑PC, 5=苹果APP, 6=安卓APP]',
  `is_disable` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否禁用: [0=否, 1=是]',
  `is_delete` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除: [0=否, 1=是]',
  `last_login_ip` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '最后登录IP',
  `last_login_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '最后登录时间',
    `create_time` DATETIME NOT NULL COMMENT '创建时间',
	`update_time` DATETIME NOT NULL COMMENT '更新时间',
	`delete_time` DATETIME  NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户信息表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_user
-- ----------------------------

-- ----------------------------
-- Table structure for x_user_auth
-- ----------------------------
DROP TABLE IF EXISTS `x_user_auth`;
CREATE TABLE `x_user_auth`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
  `openid` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Openid',
  `unionid` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'Unionid',
  `client` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '客户端类型: [1=微信小程序, 2=微信公众号, 3=手机H5, 4=电脑PC, 5=苹果APP, 6=安卓APP]',
    `create_time` DATETIME NOT NULL COMMENT '创建时间',
	`update_time` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `openid`(`openid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户授权表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_user_auth
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
