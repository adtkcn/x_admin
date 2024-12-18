/*
 Navicat Premium Data Transfer

 Source Server         : 5_7_root
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : localhost:3306
 Source Schema         : x_admin_2

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 25/10/2024 18:35:03
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
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_cid`(`cid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 35 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '相册管理表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_album
-- ----------------------------
INSERT INTO `x_album` VALUES (31, 6, 1, 0, 10, '6dd3ee16f5822373ffb6318c426cf13 (2).png', 'image/20241305/089ab8212e1e4b4ebe6c5cbe516ea3a5.png', 'png', 11914, 0, '2024-05-13 19:12:39', '2024-05-13 19:12:39', NULL);
INSERT INTO `x_album` VALUES (32, 0, 1, 0, 10, 'appicon.png', 'image/20241705/f0eb36d508834bc2ac1b8c591c563efa.png', 'png', 9873, 0, '2024-05-17 12:15:05', '2024-05-17 12:15:05', NULL);
INSERT INTO `x_album` VALUES (33, 0, 1, 0, 10, 'COMPUMAX_1p2x61.jpg', 'image/20240908/98587b5a927a4a18b5f7d165a3e60287.jpg', 'jpg', 613802, 0, '2024-08-09 00:21:03', '2024-08-09 00:21:03', NULL);
INSERT INTO `x_album` VALUES (34, 6, 1, 0, 10, 'Alx_gp73pq - 副本.png', 'image/20242410/7324e2b149f6412f8af21993e13c2fe9.png', 'png', 7138309, 0, '2024-10-24 18:16:44', '2024-10-24 18:16:44', NULL);

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
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '相册分类表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_album_cate
-- ----------------------------
INSERT INTO `x_album_cate` VALUES (6, 0, 10, 'a', 0, '2024-05-13 19:12:28', '2024-05-13 19:12:28', NULL);
INSERT INTO `x_album_cate` VALUES (7, 0, 10, '1', 0, '2024-08-11 03:43:51', '2024-08-11 03:43:51', NULL);

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
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `cid_idx`(`cid`) USING BTREE COMMENT '分类索引'
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文章资讯表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_article
-- ----------------------------

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
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文章分类表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_article_category
-- ----------------------------

-- ----------------------------
-- Table structure for x_article_collect
-- ----------------------------
DROP TABLE IF EXISTS `x_article_collect`;
CREATE TABLE `x_article_collect`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
  `article_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '文章ID',
  `is_delete` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文章收藏表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_article_collect
-- ----------------------------

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
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 24 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '字典数据表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_dict_data
-- ----------------------------
INSERT INTO `x_dict_data` VALUES (1, 2, '待提交', '1', '#6D85FC', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL);
INSERT INTO `x_dict_data` VALUES (2, 2, '审批中', '2', '#C6C150', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL);
INSERT INTO `x_dict_data` VALUES (3, 2, '审批成功', '3', 'green', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL);
INSERT INTO `x_dict_data` VALUES (4, 2, '失败', '4', 'red', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL);
INSERT INTO `x_dict_data` VALUES (5, 3, '待处理', '1', '#087BF6', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL);
INSERT INTO `x_dict_data` VALUES (6, 3, '通过', '2', 'green', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL);
INSERT INTO `x_dict_data` VALUES (7, 3, '拒绝', '3', 'red', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL);
INSERT INTO `x_dict_data` VALUES (10, 4, '假勤管理', '1', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL);
INSERT INTO `x_dict_data` VALUES (11, 4, '人事管理', '2', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL);
INSERT INTO `x_dict_data` VALUES (12, 4, '财务管理', '3', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL);
INSERT INTO `x_dict_data` VALUES (13, 4, '业务管理', '4', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL);
INSERT INTO `x_dict_data` VALUES (14, 4, '行政管理', '5', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL);
INSERT INTO `x_dict_data` VALUES (15, 4, '法务管理', '6', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL);
INSERT INTO `x_dict_data` VALUES (16, 4, '其他', '7', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL);
INSERT INTO `x_dict_data` VALUES (17, 5, 'web', 'web', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-06-29 00:18:02', NULL);
INSERT INTO `x_dict_data` VALUES (18, 5, 'go', 'go', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL);
INSERT INTO `x_dict_data` VALUES (19, 5, 'uniapp', 'uniapp', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL);
INSERT INTO `x_dict_data` VALUES (20, 5, 'node', 'node', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL);
INSERT INTO `x_dict_data` VALUES (21, 5, '1', '1', '', '', 0, 1, 1, '2024-06-29 00:31:08', '2024-06-29 00:31:11', '2024-06-29 00:31:11');
INSERT INTO `x_dict_data` VALUES (22, 6, '禁用', '0', '#FF4B4B', '', 0, 1, 0, '2024-09-25 16:02:08', '2024-09-25 16:02:08', NULL);
INSERT INTO `x_dict_data` VALUES (23, 6, '启用', '1', '#80D251', '', 0, 1, 0, '2024-09-25 16:02:30', '2024-09-25 16:02:30', NULL);

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
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '字典类型表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_dict_type
-- ----------------------------
INSERT INTO `x_dict_type` VALUES (2, '审批申请状态', 'flow_apply_status', '0待提交，1审批中，2审批完成，3审批失败', 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL);
INSERT INTO `x_dict_type` VALUES (3, '审批历史状态', 'flow_history_status', '', 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL);
INSERT INTO `x_dict_type` VALUES (4, '流程分类', 'flow_group', '1假勤管理,2人事管理3财务管理4业务管理5行政管理6法务管理7其他', 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL);
INSERT INTO `x_dict_type` VALUES (5, '项目类型1', 'project_type', '项目类型go java web node php 等', 1, 0, '2024-01-02 03:04:05', '2024-06-29 00:48:34', NULL);
INSERT INTO `x_dict_type` VALUES (6, '启用状态', 'status', ' 0=否, 1=是', 1, 0, '2024-09-25 15:54:56', '2024-09-25 15:54:56', NULL);

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
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '申请流程' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_flow_apply
-- ----------------------------

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
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `apply_id`(`apply_id`) USING BTREE,
  INDEX `approver_id`(`approver_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 19 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '流程历史' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_flow_history
-- ----------------------------

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
  `is_delete` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除: 0=否, 1=是',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `delete_time` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '流程模板' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_flow_template
-- ----------------------------

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
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 27 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '代码生成业务表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_gen_table
-- ----------------------------

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
  `list_all_api` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '下拉框列表数据来源api',
  `sort` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序编号',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 228 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '代码生成字段表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_gen_table_column
-- ----------------------------

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
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `project_key`(`project_key`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  INDEX `client_id`(`client_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '监控-客户端信息' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_monitor_client
-- ----------------------------
INSERT INTO `x_monitor_client` VALUES (1, 'e19e3be20de94f49b68fafb4c30668bc', '1', '11', '1', '1', '1', 1, 1, '1', '2024-09-25 16:16:03');
INSERT INTO `x_monitor_client` VALUES (2, '6217ea4ea0044014831bd25121a3113c', '2', '2', '2', '2', '2', 2, 2, '2', '2024-10-11 12:17:26');

-- ----------------------------
-- Table structure for x_monitor_error
-- ----------------------------
DROP TABLE IF EXISTS `x_monitor_error`;
CREATE TABLE `x_monitor_error`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '错误id',
  `project_key` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '项目key',
  `md5` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'md5',
  `event_type` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '事件类型',
  `path` varchar(1000) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'URL地址',
  `message` varchar(1000) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '错误消息',
  `stack` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '错误堆栈',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `project_key`(`project_key`, `md5`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 41 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '监控-错误列表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_monitor_error
-- ----------------------------
INSERT INTO `x_monitor_error` VALUES (39, '2a01a60efed14d6eab279ba5c664ceac', 'd5d9182bdad80a3397f248a2d820958c', '1', '1', '', '', '2024-09-25 18:33:17');
INSERT INTO `x_monitor_error` VALUES (40, '2a01a60efed14d6eab279ba5c664ceac', '4f42e2b532ee0f7924d43f63caaa26d2', '1', '1', '1', '1', '2024-09-25 18:35:33');

-- ----------------------------
-- Table structure for x_monitor_error_list
-- ----------------------------
DROP TABLE IF EXISTS `x_monitor_error_list`;
CREATE TABLE `x_monitor_error_list`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `error_id` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '错误id',
  `client_id` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '客户端id',
  `project_key` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '项目key',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `uid`(`client_id`) USING BTREE,
  INDEX `eid`(`error_id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '错误对应的用户记录' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of x_monitor_error_list
-- ----------------------------
INSERT INTO `x_monitor_error_list` VALUES (6, '40', '1', '2a01a60efed14d6eab279ba5c664ceac', '2024-09-25 18:35:33');
INSERT INTO `x_monitor_error_list` VALUES (5, '39', '1', '2a01a60efed14d6eab279ba5c664ceac', '2024-09-25 18:34:24');
INSERT INTO `x_monitor_error_list` VALUES (4, '39', '1', '2a01a60efed14d6eab279ba5c664ceac', '2024-09-25 18:33:17');

-- ----------------------------
-- Table structure for x_monitor_project
-- ----------------------------
DROP TABLE IF EXISTS `x_monitor_project`;
CREATE TABLE `x_monitor_project`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '项目id',
  `project_key` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'project_key',
  `project_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '项目名称',
  `project_type` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '项目类型go java web node php 等',
  `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否启用: 0=否, 1=是',
  `is_delete` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除: 0=否, 1=是',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `project_key`(`project_key`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '监控项目' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_monitor_project
-- ----------------------------
INSERT INTO `x_monitor_project` VALUES (1, '2a01a60efed14d6eab279ba5c664ceac', 'x_admin', 'node', 0, 0, '2024-05-18 15:59:53', '2024-07-09 12:06:42', NULL);
INSERT INTO `x_monitor_project` VALUES (2, '9a1d5d657d884bbfb7df0ddf032a8504', '2', 'go', 0, 1, '2024-06-17 18:58:57', '2024-06-29 00:31:27', NULL);
INSERT INTO `x_monitor_project` VALUES (3, '603055b4b53e45a4b6fb725147d7fd42', 'ad', 'web', 0, 1, '2024-07-08 15:26:56', '2024-07-10 12:21:11', NULL);
INSERT INTO `x_monitor_project` VALUES (4, '95cf6a6a37e549019c90c2ed34704f1f', 'a', 'uniapp', 0, 1, '2024-07-09 00:45:18', '2024-07-09 00:49:27', NULL);
INSERT INTO `x_monitor_project` VALUES (5, '57377a40408f4bfd974047fbf0a18e24', 'go项目', 'go', 0, 0, '2024-07-10 12:20:37', '2024-09-25 16:50:45', NULL);
INSERT INTO `x_monitor_project` VALUES (6, '6217ea4ea0044014831bd25121a3113c', 'go', 'go', 0, 0, '2024-07-12 23:17:23', '2024-07-12 23:17:23', NULL);
INSERT INTO `x_monitor_project` VALUES (7, 'e19e3be20de94f49b68fafb4c30668bc', 'web项目', 'web', 1, 0, '2024-07-13 20:56:21', '2024-09-25 16:58:58', NULL);
INSERT INTO `x_monitor_project` VALUES (8, 'e0cb2d091ce1408babcf17b48eccebf6', '1', 'web', 1, 1, '2024-10-12 10:35:32', '2024-10-12 10:35:38', '2024-10-12 10:35:38');

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
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '消息通知设置表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_notice_setting
-- ----------------------------

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
  `last_login_time` datetime NULL DEFAULT NULL COMMENT '最后登录',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统管理成员表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_system_auth_admin
-- ----------------------------
INSERT INTO `x_system_auth_admin` VALUES (1, 1, 3, 'admin', 'admin', '81a13dd8e25644a8823082573ca973f7', '/image/20232911/27699a1e991e4ba1ba13ee6d45185db6.jpg', '0', 'WFdiD', 1, 1, 0, 0, '127.0.0.1', '2024-10-25 18:17:07', '2024-01-02 03:04:05', '2024-10-25 18:17:07', NULL);

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
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统部门管理表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_system_auth_dept
-- ----------------------------
INSERT INTO `x_system_auth_dept` VALUES (1, 0, '默认部门', 1, 'admin', '18327647788', 10, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL);
INSERT INTO `x_system_auth_dept` VALUES (2, 1, '指挥部', 2, '指挥部01', '17608390000', 3, 0, 0, '2024-01-02 03:04:05', '2024-07-05 15:18:25', NULL);
INSERT INTO `x_system_auth_dept` VALUES (3, 2, '指挥部子级', 0, '', '', 0, 0, 1, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL);

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
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 832 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统菜单管理表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_system_auth_menu
-- ----------------------------
INSERT INTO `x_system_auth_menu` VALUES (1, 0, 'C', '工作台', 'el-icon-Monitor', 50, 'admin:common:index:console', 'workbench', 'workbench/index', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (100, 500, 'M', '权限管理', 'el-icon-Lock', 44, '', 'permission', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (101, 100, 'C', '管理员', 'local-icon-wode', 0, 'admin:system:admin:list', 'admin', 'setting/permission/admin/index', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (102, 101, 'A', '管理员详情', '', 0, 'admin:system:admin:detail', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (103, 101, 'A', '管理员新增', '', 0, 'admin:system:admin:add', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (104, 101, 'A', '管理员编辑', '', 0, 'admin:system:admin:edit', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (105, 101, 'A', '管理员删除', '', 0, 'admin:system:admin:del', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (106, 101, 'A', '管理员状态', '', 0, 'admin:system:admin:disable', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (110, 100, 'C', '角色管理', 'el-icon-Female', 0, 'admin:system:role:list', 'role', 'setting/permission/role/index', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (111, 110, 'A', '角色详情', '', 0, 'admin:system:role:detail', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (112, 110, 'A', '角色新增', '', 0, 'admin:system:role:add', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (113, 110, 'A', '角色编辑', '', 0, 'admin:system:role:edit', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (114, 110, 'A', '角色删除', '', 0, 'admin:system:role:del', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (120, 100, 'C', '菜单管理', 'el-icon-Operation', 0, 'admin:system:menu:list', 'menu', 'setting/permission/menu/index', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (121, 120, 'A', '菜单详情', '', 0, 'admin:system:menu:detail', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (122, 120, 'A', '菜单新增', '', 0, 'admin:system:menu:add', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (123, 120, 'A', '菜单编辑', '', 0, 'admin:system:menu:edit', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (124, 120, 'A', '菜单删除', '', 0, 'admin:system:menu:del', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (130, 500, 'M', '组织管理', 'el-icon-OfficeBuilding', 45, '', 'organization', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (131, 130, 'C', '部门管理', 'el-icon-Coordinate', 0, 'admin:system:dept:all', 'department', 'organization/department/index', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (132, 131, 'A', '部门详情', '', 0, 'admin:system:dept:detail', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (133, 131, 'A', '部门新增', '', 0, 'admin:system:dept:add', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (134, 131, 'A', '部门编辑', '', 0, 'admin:system:dept:edit', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (135, 131, 'A', '部门删除', '', 0, 'admin:system:dept:del', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (140, 130, 'C', '岗位管理', 'el-icon-PriceTag', 0, 'admin:system:post:list', 'post', 'organization/post/index', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (141, 140, 'A', '岗位详情', '', 0, 'admin:system:post:detail', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (142, 140, 'A', '岗位新增', '', 0, 'admin:system:post:add', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (143, 140, 'A', '岗位编辑', '', 0, 'admin:system:post:edit', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (144, 140, 'A', '岗位删除', '', 0, 'admin:system:post:del', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (200, 0, 'M', '其它管理', '', 0, '', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (201, 200, 'M', '图库管理', '', 0, '', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (202, 201, 'A', '文件列表', '', 0, 'admin:common:album:albumList', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (203, 201, 'A', '文件命名', '', 0, 'admin:common:album:albumRename', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (204, 201, 'A', '文件移动', '', 0, 'admin:common:album:albumMove', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (205, 201, 'A', '文件删除', '', 0, 'admin:common:album:albumDel', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (206, 201, 'A', '分类列表', '', 0, 'admin:common:album:cateList', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (207, 201, 'A', '分类新增', '', 0, 'admin:common:album:cateAdd', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (208, 201, 'A', '分类命名', '', 0, 'admin:common:album:cateRename', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (209, 201, 'A', '分类删除', '', 0, 'admin:common:album:cateDel', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (215, 200, 'M', '上传管理', '', 0, '', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (216, 215, 'A', '上传图片', '', 0, 'admin:common:upload:image', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (217, 215, 'A', '上传视频', '', 0, 'admin:common:upload:video', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (500, 0, 'M', '系统设置', 'el-icon-Setting', 0, '', 'setting', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (501, 500, 'M', '网站设置', 'el-icon-Basketball', 10, '', 'website', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (502, 501, 'C', '网站信息', '', 0, 'admin:setting:website:detail', 'information', 'setting/website/information', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (503, 502, 'A', '保存配置', '', 0, 'admin:setting:website:save', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (505, 501, 'C', '网站备案', '', 0, 'admin:setting:copyright:detail', 'filing', 'setting/website/filing', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (506, 505, 'A', '备案保存', '', 0, 'admin:setting:copyright:save', '', 'setting/website/protocol', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (510, 501, 'C', '政策协议', '', 0, 'admin:setting:protocol:detail', 'protocol', 'setting/website/protocol', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (511, 510, 'A', '协议保存', '', 0, 'admin:setting:protocol:save', '', '', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (515, 500, 'C', '字典管理', 'el-icon-Box', 0, 'admin:setting:dict:type:list', 'dict', 'setting/dict/type/index', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (516, 515, 'A', '字典类型新增', '', 0, 'admin:setting:dict:type:add', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (517, 515, 'A', '字典类型编辑', '', 0, 'admin:setting:dict:type:edit', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (518, 515, 'A', '字典类型删除', '', 0, 'admin:setting:dict:type:del', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (519, 500, 'C', '字典数据管理', '', 0, 'admin:setting:dict:data:list', 'dict/data', 'setting/dict/data/index', '/dev_tools/dict', '', 0, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (520, 515, 'A', '字典数据新增', '', 0, 'admin:setting:dict:data:add', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (521, 515, 'A', '字典数据编辑', '', 0, 'admin:setting:dict:data:edit', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (522, 515, 'A', '字典数据删除', '', 0, 'admin:setting:dict:data:del', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (550, 500, 'M', '系统维护', 'el-icon-SetUp', 0, '', 'system', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (551, 550, 'C', '系统环境', '', 0, 'admin:monitor:server', 'environment', 'setting/system/environment', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (552, 550, 'C', '系统缓存', '', 0, 'admin:monitor:cache', 'cache', 'setting/system/cache', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (553, 550, 'C', '系统日志', '', 0, 'admin:system:log:operate', 'journal', 'setting/system/journal', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (600, 0, 'M', '开发工具', 'el-icon-EditPen', 0, '', 'dev_tools', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (610, 600, 'C', '代码生成器', 'el-icon-DocumentAdd', 0, 'admin:gen:list', 'code', 'dev_tools/code/index', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (611, 610, 'A', '导入数据表', '', 0, 'admin:gen:importTable', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (612, 610, 'A', '生成代码', '', 0, 'admin:gen:genCode', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (613, 610, 'A', '下载代码', '', 0, 'admin:gen:downloadCode', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (614, 610, 'A', '预览代码', '', 0, 'admin:gen:previewCode', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (616, 610, 'A', '同步表结构', '', 0, 'admin:gen:syncTable', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (617, 610, 'A', '删除数据表', '', 0, 'admin:gen:delTable', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (618, 610, 'A', '数据表详情', '', 0, 'admin:gen:detail', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (700, 0, 'M', '素材管理', 'el-icon-Picture', 43, '', 'material', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (701, 700, 'C', '素材中心', 'el-icon-PictureRounded', 0, '', 'index', 'material/index', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (775, 600, 'C', '代码生成器编辑', 'el-icon-EditPen', 0, 'admin:gen:editTable', 'dev_tools/code/edit', 'dev_tools/code/edit', '', '', 0, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (776, 778, 'C', '流程模板', '', 0, 'admin:flow:flow_template:list', 'flow_template/index', 'flow/flow_template/index', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (777, 778, 'C', '我的流程', '', 0, '', 'flow_apply/index', 'flow/flow_apply/index', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (778, 0, 'M', '审批流', 'el-icon-Coordinate', 0, '', 'flow', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (780, 778, 'C', '待处理', '', 0, 'admin:flow:flow_history:list', 'flow_history/todo', 'flow/flow_history/todo', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (781, 778, 'C', '已处理', '', 0, 'admin:flow:flow_history:list', 'flow_history/done', 'flow/flow_history/done', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (782, 778, 'C', '已完成流程', '', 0, 'admin:flow:flow_history:list', 'flow_apply/finish', 'flow/flow_apply/finish', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (783, 500, 'C', '项目监控', 'el-icon-Notebook', 0, 'admin:monitor_project:list', 'monitor/project/index', 'monitor/project/index', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-09-25 16:14:03');
INSERT INTO `x_system_auth_menu` VALUES (784, 500, 'C', '监控用户端', '', 0, '', 'monitor/client/index', 'monitor/client/index', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-09-25 16:13:05');
INSERT INTO `x_system_auth_menu` VALUES (785, 794, 'A', '错误收集error添加', '', 0, 'admin:monitor_web:add', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (786, 794, 'A', '错误收集error编辑', '', 0, 'admin:monitor_web:edit', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (787, 794, 'A', '错误收集error删除', '', 0, 'admin:monitor_web:del', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (788, 794, 'A', '错误收集error列表', '', 0, 'admin:monitor_web:list', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (789, 794, 'A', '错误收集error全部列表', '', 0, 'admin:monitor_web:listAll', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (790, 794, 'A', '错误收集error详情', '', 0, 'admin:monitor_web:detail', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (791, 794, 'A', '错误收集error导出excel', '', 0, 'admin:monitor_web:ExportFile', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (792, 794, 'A', '错误收集error导入excel', '', 0, 'admin:monitor_web:ImportFile', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (794, 500, 'C', '错误列表', '', 0, '', 'monitor/error/index', 'monitor/error/index', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-09-25 16:38:39');
INSERT INTO `x_system_auth_menu` VALUES (795, 776, 'A', '流程模板添加', '', 0, 'admin:flow:flow_template:add', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (796, 776, 'A', '流程模板编辑', '', 0, 'admin:flow:flow_template:edit', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (797, 776, 'A', '流程模板删除', '', 0, 'admin:flow:flow_template:del', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (798, 776, 'A', '流程模板列表', '', 0, 'admin:flow:flow_template:list', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (799, 776, 'A', '流程模板全部列表', '', 0, 'admin:flow:flow_template:listAll', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (800, 776, 'A', '流程模板详情', '', 0, 'admin:flow:flow_template:detail', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (801, 777, 'A', '流程详情', '', 0, 'admin:flow:flow_apply:detail', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (802, 777, 'A', '流程添加', '', 0, 'admin:flow:flow_apply:add', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (803, 777, 'A', '流程编辑', '', 0, 'admin:flow:flow_apply:edit', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (804, 777, 'A', '流程删除', '', 0, 'admin:flow:flow_apply:del', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (805, 777, 'A', '全部流程', '', 0, 'admin:flow:flow_history:listAll', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (806, 777, 'A', '通过流程', '', 0, 'admin:flow:flow_history:pass', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (807, 777, 'A', '拒接流程', '', 0, 'admin:flow:flow_history:back', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (808, 777, 'A', '下一个流程', '', 0, 'admin:flow:flow_history:next_node', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (809, 777, 'A', '获取审批人', '', 0, 'admin:flow:flow_history:get_approver', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (810, 777, 'A', '流程列表', '', 0, 'admin:flow:flow_apply:list', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (811, 780, 'A', '审批记录列表', '', 0, 'admin:flow:flow_history:list', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (812, 780, 'A', '审批记录详情', '', 0, 'admin:flow:flow_history:detail', '', '', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (813, 783, 'A', '监控项目添加', '', 0, 'admin:monitor_project:add', '', '', '', '', 0, 1, 0, '2024-06-17 19:26:13', '2024-06-17 19:26:13');
INSERT INTO `x_system_auth_menu` VALUES (814, 783, 'A', '监控项目编辑', '', 0, 'admin:monitor_project:edit', '', '', '', '', 0, 1, 0, '2024-06-17 19:26:13', '2024-06-17 19:26:13');
INSERT INTO `x_system_auth_menu` VALUES (815, 783, 'A', '监控项目删除', '', 0, 'admin:monitor_project:del', '', '', '', '', 0, 1, 0, '2024-06-17 19:26:13', '2024-06-17 19:26:13');
INSERT INTO `x_system_auth_menu` VALUES (816, 783, 'A', '监控项目列表', '', 0, 'admin:monitor_project:list', '', '', '', '', 0, 1, 0, '2024-06-17 19:26:13', '2024-06-17 19:26:13');
INSERT INTO `x_system_auth_menu` VALUES (817, 783, 'A', '监控项目全部列表', '', 0, 'admin:monitor_project:listAll', '', '', '', '', 0, 1, 0, '2024-06-17 19:26:13', '2024-06-17 19:26:13');
INSERT INTO `x_system_auth_menu` VALUES (818, 783, 'A', '监控项目详情', '', 0, 'admin:monitor_project:detail', '', '', '', '', 0, 1, 0, '2024-06-17 19:26:13', '2024-06-17 19:26:13');
INSERT INTO `x_system_auth_menu` VALUES (819, 783, 'A', '监控项目导出excel', '', 0, 'admin:monitor_project:ExportFile', '', '', '', '', 0, 1, 0, '2024-06-17 19:26:13', '2024-06-17 19:26:13');
INSERT INTO `x_system_auth_menu` VALUES (820, 783, 'A', '监控项目导入excel', '', 0, 'admin:monitor_project:ImportFile', '', '', '', '', 0, 1, 0, '2024-06-17 19:26:13', '2024-06-17 19:26:13');
INSERT INTO `x_system_auth_menu` VALUES (821, 0, 'C', '文档', 'el-icon-Wallet', 0, '', 'https://adtkcn.github.io/x_admin/', '', '', 'a=1', 1, 1, 0, '2024-06-28 17:09:17', '2024-08-16 15:05:08');
INSERT INTO `x_system_auth_menu` VALUES (823, 0, 'C', '用户协议', '', 0, '', 'user/protocol/index', 'user/protocol/index', '', '', 0, 1, 0, '2024-09-10 20:03:46', '2024-09-10 22:47:36');
INSERT INTO `x_system_auth_menu` VALUES (824, 823, 'A', '用户协议添加', '', 0, 'admin:user_protocol:add', '', '', '', '', 0, 1, 0, '2024-09-10 20:05:13', '2024-09-10 20:05:13');
INSERT INTO `x_system_auth_menu` VALUES (825, 823, 'A', '用户协议编辑', '', 0, 'admin:user_protocol:edit', '', '', '', '', 0, 1, 0, '2024-09-10 20:05:13', '2024-09-10 20:05:13');
INSERT INTO `x_system_auth_menu` VALUES (826, 823, 'A', '用户协议删除', '', 0, 'admin:user_protocol:del', '', '', '', '', 0, 1, 0, '2024-09-10 20:05:13', '2024-09-10 20:05:13');
INSERT INTO `x_system_auth_menu` VALUES (827, 823, 'A', '用户协议列表', '', 0, 'admin:user_protocol:list', '', '', '', '', 0, 1, 0, '2024-09-10 20:05:13', '2024-09-10 20:05:13');
INSERT INTO `x_system_auth_menu` VALUES (828, 823, 'A', '用户协议全部列表', '', 0, 'admin:user_protocol:listAll', '', '', '', '', 0, 1, 0, '2024-09-10 20:05:13', '2024-09-10 20:05:13');
INSERT INTO `x_system_auth_menu` VALUES (829, 823, 'A', '用户协议详情', '', 0, 'admin:user_protocol:detail', '', '', '', '', 0, 1, 0, '2024-09-10 20:05:13', '2024-09-10 20:05:13');
INSERT INTO `x_system_auth_menu` VALUES (830, 823, 'A', '用户协议导出excel', '', 0, 'admin:user_protocol:ExportFile', '', '', '', '', 0, 1, 0, '2024-09-10 20:05:13', '2024-09-10 20:05:13');
INSERT INTO `x_system_auth_menu` VALUES (831, 823, 'A', '用户协议导入excel', '', 0, 'admin:user_protocol:ImportFile', '', '', '', '', 0, 1, 0, '2024-09-10 20:05:13', '2024-09-10 20:05:13');

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
INSERT INTO `x_system_auth_perm` VALUES ('00b23681b95c4432a56fd7f5438d29ff', 1, 111);
INSERT INTO `x_system_auth_perm` VALUES ('015a14213276483ab2d5ff7904ad00ac', 1, 519);
INSERT INTO `x_system_auth_perm` VALUES ('06289de4159748d783ea5c24ae36cc26', 1, 102);
INSERT INTO `x_system_auth_perm` VALUES ('066c5d6eb0f043d3b6cd0c08b4178d49', 1, 206);
INSERT INTO `x_system_auth_perm` VALUES ('0e20554114934323b68e2c1935466384', 1, 612);
INSERT INTO `x_system_auth_perm` VALUES ('0f6ba3d081ad4d9b86f0ae47c7b9107f', 1, 610);
INSERT INTO `x_system_auth_perm` VALUES ('10124b836e4244c79535351dabd80528', 1, 701);
INSERT INTO `x_system_auth_perm` VALUES ('117be250cd6f4637843aec001b26d644', 1, 700);
INSERT INTO `x_system_auth_perm` VALUES ('126b45b2c73a436ba90fe54c2e780981', 1, 784);
INSERT INTO `x_system_auth_perm` VALUES ('1852c3f2614b459eaffb11b69ff7d3ac', 1, 809);
INSERT INTO `x_system_auth_perm` VALUES ('1937a72912fc4564af391ea8a1045b37', 1, 775);
INSERT INTO `x_system_auth_perm` VALUES ('1aabdba5cacc4c5fba84c1c58433e5cf', 1, 200);
INSERT INTO `x_system_auth_perm` VALUES ('2436b8caae694cfe8b1d56cea92c3410', 1, 829);
INSERT INTO `x_system_auth_perm` VALUES ('2479a60e4c47414282d11cc627e26df6', 1, 828);
INSERT INTO `x_system_auth_perm` VALUES ('257937880a6843dbb60e6a171a8436ea', 1, 144);
INSERT INTO `x_system_auth_perm` VALUES ('28f9a3ea7cfd4c47ae7f19c2d17001e8', 1, 207);
INSERT INTO `x_system_auth_perm` VALUES ('29760cf1193e4a989cd5e5442dab3b4b', 1, 106);
INSERT INTO `x_system_auth_perm` VALUES ('2ee8928a0d594dd2a6d6f210666ab3a6', 1, 802);
INSERT INTO `x_system_auth_perm` VALUES ('2f03c73106f548258742681ab95c8ed2', 1, 796);
INSERT INTO `x_system_auth_perm` VALUES ('2f4b7c6cf88d4b0ca27d015470d9a005', 1, 822);
INSERT INTO `x_system_auth_perm` VALUES ('3072e49dc9a645feb3eb05b8d464eeba', 1, 825);
INSERT INTO `x_system_auth_perm` VALUES ('326f68c509a7457b80296ab1a6ea141e', 1, 780);
INSERT INTO `x_system_auth_perm` VALUES ('3525ee3d7ce6490885ca3f3889fbd110', 1, 505);
INSERT INTO `x_system_auth_perm` VALUES ('38f456adb1644b7fa83f8045d38a8d3e', 1, 618);
INSERT INTO `x_system_auth_perm` VALUES ('3b39276ec49342ab8b0b49ce759cb20b', 1, 613);
INSERT INTO `x_system_auth_perm` VALUES ('4047a2c7efbb49fa95abb419c49ab4cb', 1, 807);
INSERT INTO `x_system_auth_perm` VALUES ('45b8230de8454d9299c5557f857f14d7', 1, 132);
INSERT INTO `x_system_auth_perm` VALUES ('4829c386f37e4ad0bfba9691745e2ef0', 1, 522);
INSERT INTO `x_system_auth_perm` VALUES ('485fe03be60c4b208a3a5932c1b82451', 1, 203);
INSERT INTO `x_system_auth_perm` VALUES ('4a602dbb4e1549c49ae2dea8d93ffc06', 1, 816);
INSERT INTO `x_system_auth_perm` VALUES ('4ae52b2643284ca3afa0160765d2aa85', 1, 798);
INSERT INTO `x_system_auth_perm` VALUES ('4d701fd8c34041d49cdca9c6a19ad383', 1, 810);
INSERT INTO `x_system_auth_perm` VALUES ('4e2fcc51fe774a5ab6514009b3dc263c', 1, 818);
INSERT INTO `x_system_auth_perm` VALUES ('4f8ca12187054b9ca61a9c9b0437e615', 1, 794);
INSERT INTO `x_system_auth_perm` VALUES ('50a90d5574594cd18a9c97857027f417', 1, 815);
INSERT INTO `x_system_auth_perm` VALUES ('513ab57533d544858d4d1abda4a72960', 1, 103);
INSERT INTO `x_system_auth_perm` VALUES ('5216160327b74c9b950667036cb99e8b', 1, 777);
INSERT INTO `x_system_auth_perm` VALUES ('531b5f44f62a4f1c85ec56f6dc56bdd7', 1, 820);
INSERT INTO `x_system_auth_perm` VALUES ('53789a2311394a01a0b9e8957136a9e8', 1, 515);
INSERT INTO `x_system_auth_perm` VALUES ('5449939f72aa4c50b9df93992e19c51a', 1, 122);
INSERT INTO `x_system_auth_perm` VALUES ('56112a65356a451db622ce3b431ca73d', 1, 783);
INSERT INTO `x_system_auth_perm` VALUES ('5841ea91228347718c4d7d7c43fb8fd2', 1, 792);
INSERT INTO `x_system_auth_perm` VALUES ('58504dd855314786a11663c65c4ce15f', 1, 134);
INSERT INTO `x_system_auth_perm` VALUES ('613ad8755a644db79541bf25b1f010f8', 1, 201);
INSERT INTO `x_system_auth_perm` VALUES ('61c0692345c14dea9d8ac26656418791', 1, 1);
INSERT INTO `x_system_auth_perm` VALUES ('62411ba321544c289cd87234da2fd003', 1, 600);
INSERT INTO `x_system_auth_perm` VALUES ('6372e3c7064a4d75b43f57621ffcefcf', 1, 553);
INSERT INTO `x_system_auth_perm` VALUES ('6499ca918ee04a1fb76220da8c4c2794', 1, 801);
INSERT INTO `x_system_auth_perm` VALUES ('64baaf548bc342c3af9322a30768bf96', 1, 516);
INSERT INTO `x_system_auth_perm` VALUES ('65fe5c85bf2349aca368aed82c7f1538', 1, 113);
INSERT INTO `x_system_auth_perm` VALUES ('664180087ea24ce8b6c0dc1cb1436c44', 1, 811);
INSERT INTO `x_system_auth_perm` VALUES ('67f07d19bc9a49919692dd33eae90692', 1, 808);
INSERT INTO `x_system_auth_perm` VALUES ('67f1013bc0d644c5b12a08cd48c85bbc', 1, 823);
INSERT INTO `x_system_auth_perm` VALUES ('6cb18648246049909a53c8bd15b529e8', 1, 124);
INSERT INTO `x_system_auth_perm` VALUES ('6e4fcd3359ed48cd9c8cbfab8bee7c10', 1, 789);
INSERT INTO `x_system_auth_perm` VALUES ('6e5d3b14c120473da6af38046727fc29', 1, 791);
INSERT INTO `x_system_auth_perm` VALUES ('6ebd40320bab491785d1a940f9cd31b5', 1, 510);
INSERT INTO `x_system_auth_perm` VALUES ('715bc10138aa44e2b0fd48f72efe26e5', 1, 114);
INSERT INTO `x_system_auth_perm` VALUES ('71a975a2fa694c089f876ace329e70e3', 1, 831);
INSERT INTO `x_system_auth_perm` VALUES ('7356ec71b47945bc8758235859ad6d30', 1, 209);
INSERT INTO `x_system_auth_perm` VALUES ('742e2c95cc5d42bca97845b71b969f12', 1, 551);
INSERT INTO `x_system_auth_perm` VALUES ('77d0aa7e90744bb191112a64b305aacb', 1, 502);
INSERT INTO `x_system_auth_perm` VALUES ('7a6c10af221f4bf4a6b3a75f49e0c25f', 1, 216);
INSERT INTO `x_system_auth_perm` VALUES ('7f2b461215ac46d88d4faf184885df1b', 1, 806);
INSERT INTO `x_system_auth_perm` VALUES ('802afa2ffa1d4680b177dd542669505e', 1, 819);
INSERT INTO `x_system_auth_perm` VALUES ('8039993b87fe4eeb82b596ae56596b01', 1, 787);
INSERT INTO `x_system_auth_perm` VALUES ('80b771ff355a48178510d7f57b4507a6', 1, 110);
INSERT INTO `x_system_auth_perm` VALUES ('8410858035a14989afb56b6aeafed26d', 1, 552);
INSERT INTO `x_system_auth_perm` VALUES ('857e8369ce4f41ac981a27b6e22b2ad4', 1, 550);
INSERT INTO `x_system_auth_perm` VALUES ('8f3508393a0b49368cc0642908e99029', 1, 786);
INSERT INTO `x_system_auth_perm` VALUES ('90c582e4b81b4012925c94c0d228f6a2', 1, 776);
INSERT INTO `x_system_auth_perm` VALUES ('97241dd15a18425da170e1e3aa275f33', 1, 812);
INSERT INTO `x_system_auth_perm` VALUES ('9ab2879262944c22b180812a6838c236', 1, 521);
INSERT INTO `x_system_auth_perm` VALUES ('9ac4019e53754a678e179251ea9bbd23', 1, 140);
INSERT INTO `x_system_auth_perm` VALUES ('9cf5f295c1c2459f81013a71d3808b91', 1, 506);
INSERT INTO `x_system_auth_perm` VALUES ('a33dbcc175b04b8786516a466b7d49c9', 1, 143);
INSERT INTO `x_system_auth_perm` VALUES ('a7c63c20539649abb9291f9ce8a59320', 1, 100);
INSERT INTO `x_system_auth_perm` VALUES ('a9d84e3316884b77819dffb7067f26f6', 1, 813);
INSERT INTO `x_system_auth_perm` VALUES ('adbd5b8fe1ba46fbaef7aa759883c97b', 1, 803);
INSERT INTO `x_system_auth_perm` VALUES ('ae71e16ab1e4453e8b49a3b7f81e72b2', 1, 517);
INSERT INTO `x_system_auth_perm` VALUES ('ae8f4a6240b24c128fd1963d207d5f22', 1, 799);
INSERT INTO `x_system_auth_perm` VALUES ('b1926fb4bbfd49fd864c53d87f748ccb', 1, 205);
INSERT INTO `x_system_auth_perm` VALUES ('b244feef3f6d48e8b07a25b4cd8f88b1', 1, 611);
INSERT INTO `x_system_auth_perm` VALUES ('b35a70ed6f564d5896c4721edc22a4e6', 1, 141);
INSERT INTO `x_system_auth_perm` VALUES ('b6379a839ec248929ad702910ba115c8', 1, 518);
INSERT INTO `x_system_auth_perm` VALUES ('b6462df064a6453d93ac66b20fa454d6', 1, 511);
INSERT INTO `x_system_auth_perm` VALUES ('b7375b434e5d48909709a7d9d71e067f', 1, 614);
INSERT INTO `x_system_auth_perm` VALUES ('b8149f7108284f0c948734960f350248', 1, 131);
INSERT INTO `x_system_auth_perm` VALUES ('b94b7f9b2bd843859169390c385cf12d', 1, 617);
INSERT INTO `x_system_auth_perm` VALUES ('b952a010c84440f5ae30cc9b7dc3600d', 1, 520);
INSERT INTO `x_system_auth_perm` VALUES ('bad3f621a6bf41be90b52cbd900a8090', 1, 503);
INSERT INTO `x_system_auth_perm` VALUES ('bb3765c66fc440519d20442afd1e5ad1', 1, 824);
INSERT INTO `x_system_auth_perm` VALUES ('bbdff371c01f4c668a93414222918f26', 1, 788);
INSERT INTO `x_system_auth_perm` VALUES ('bccfdfafe7ab41e69d8f1507c503ab00', 1, 133);
INSERT INTO `x_system_auth_perm` VALUES ('bf0e3f14fcde49afb7d5d84f4484fe8b', 1, 500);
INSERT INTO `x_system_auth_perm` VALUES ('bf76b5e1d11b48648712d9327b28c295', 1, 215);
INSERT INTO `x_system_auth_perm` VALUES ('c40c5e94d93f4cb4a65790ca6e775112', 1, 785);
INSERT INTO `x_system_auth_perm` VALUES ('c58aac49117b451a813dbb55771b1460', 1, 104);
INSERT INTO `x_system_auth_perm` VALUES ('c73861e354de430a8888272dde38ee08', 1, 797);
INSERT INTO `x_system_auth_perm` VALUES ('c93abc921a6a490eb5939d8ad2484183', 1, 208);
INSERT INTO `x_system_auth_perm` VALUES ('cb00e88d06234cc99b2d7f78a7fdc7b5', 1, 817);
INSERT INTO `x_system_auth_perm` VALUES ('cc18bc1d8b09481ebff2cbb080bd77db', 1, 814);
INSERT INTO `x_system_auth_perm` VALUES ('cd49713040df4f6fb02a5e696ab41042', 1, 142);
INSERT INTO `x_system_auth_perm` VALUES ('d066f83036974df081f3e99fd87949cc', 1, 778);
INSERT INTO `x_system_auth_perm` VALUES ('d145f604248c439994999fe262b35364', 1, 501);
INSERT INTO `x_system_auth_perm` VALUES ('d3c11aedec9d4663904a5ee1fd72cd1d', 1, 826);
INSERT INTO `x_system_auth_perm` VALUES ('d60fe11f05eb4c74aa2ada0d289a24fb', 1, 830);
INSERT INTO `x_system_auth_perm` VALUES ('dad0f8c0fffe436f9914d26959680eda', 1, 105);
INSERT INTO `x_system_auth_perm` VALUES ('db6cb5870a5747698abb281b6a41362b', 1, 804);
INSERT INTO `x_system_auth_perm` VALUES ('dbbc46d9f44346d8bae7cd75f6aaf47e', 1, 616);
INSERT INTO `x_system_auth_perm` VALUES ('dde39b9cef884c6c8a6d4c07304ce420', 1, 790);
INSERT INTO `x_system_auth_perm` VALUES ('e2d2f485941f4d58b74bc62f8aabb608', 1, 121);
INSERT INTO `x_system_auth_perm` VALUES ('e3d9361e9ef14545b97c943e8ecbf244', 1, 130);
INSERT INTO `x_system_auth_perm` VALUES ('e453e0a958ab4a38890e1e9aef9837ae', 1, 821);
INSERT INTO `x_system_auth_perm` VALUES ('e5c01fa8e4004a54b6d08bc47b061c66', 1, 120);
INSERT INTO `x_system_auth_perm` VALUES ('e814631ae1474c00adda3949bf68c7d9', 1, 217);
INSERT INTO `x_system_auth_perm` VALUES ('e82b9544d5984bf4b4d9a7bcb433287f', 1, 800);
INSERT INTO `x_system_auth_perm` VALUES ('e9ee9bc1a9464edd9b8c833156a73788', 1, 782);
INSERT INTO `x_system_auth_perm` VALUES ('f0dcd90a6d564381a129ee3597c13c0f', 1, 123);
INSERT INTO `x_system_auth_perm` VALUES ('f4c5738b28f54592b5165293d39735ee', 1, 781);
INSERT INTO `x_system_auth_perm` VALUES ('f6262ef5adac4f0094c31807f45e29a1', 1, 827);
INSERT INTO `x_system_auth_perm` VALUES ('f6b002373b6646d9896f514b9481e070', 1, 202);
INSERT INTO `x_system_auth_perm` VALUES ('f78473a721824e0ba754c9b8468adc04', 1, 204);
INSERT INTO `x_system_auth_perm` VALUES ('f79aaccc502749ff90b1633ca71f1f6e', 1, 795);
INSERT INTO `x_system_auth_perm` VALUES ('f86c95e1db2b40d7beb0afa4e86225dc', 1, 135);
INSERT INTO `x_system_auth_perm` VALUES ('fa90f76b6f114166be51b23ff8bb06f5', 1, 101);
INSERT INTO `x_system_auth_perm` VALUES ('fb43a7eeccc54eaf9c7e34ef054ed60c', 1, 805);
INSERT INTO `x_system_auth_perm` VALUES ('fe7a53bc5d14423f83f540114022db94', 1, 112);

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
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统岗位管理表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_system_auth_post
-- ----------------------------
INSERT INTO `x_system_auth_post` VALUES (3, 'zhihuibu01', '指挥部岗位', '', 0, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL);

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
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统角色管理表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_system_auth_role
-- ----------------------------
INSERT INTO `x_system_auth_role` VALUES (1, '审核员', '1', 1, 0, '2024-01-02 03:04:05', '2024-09-10 20:05:50');

-- ----------------------------
-- Table structure for x_system_config
-- ----------------------------
DROP TABLE IF EXISTS `x_system_config`;
CREATE TABLE `x_system_config`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `type` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '类型',
  `name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '键',
  `value` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '值',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 81 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统全局配置表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_system_config
-- ----------------------------
INSERT INTO `x_system_config` VALUES (1, 'storage', 'default', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_config` VALUES (2, 'storage', 'local', '{\"name\":\"本地存储\"}', '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_config` VALUES (3, 'storage', 'qiniu', '{\"name\":\"七牛云存储\",\"bucket\":\"\",\"secretKey\":\"\",\"accessKey\":\"\",\"domain\":\"\"}', '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_config` VALUES (4, 'storage', 'aliyun', '{\"name\":\"阿里云存储\",\"bucket\":\"\",\"secretKey\":\"\",\"accessKey\":\"\",\"domain\":\"\"}', '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_config` VALUES (5, 'storage', 'qcloud', '{\"name\":\"腾讯云存储\",\"bucket\":\"\",\"secretKey\":\"\",\"accessKey\":\"\",\"domain\":\"\",\"region\":\"\"}', '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_config` VALUES (6, 'sms', 'default', 'aliyun', '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_config` VALUES (7, 'sms', 'aliyun', '{\"name\":\"阿里云短信\",\"alias\":\"aliyun\",\"sign\":\"\",\"appKey\":\"\",\"secretKey\":\"\"}', '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_config` VALUES (8, 'sms', 'tencent', '{\"name\":\"腾讯云短信\",\"alias\":\"tencent\",\"sign\":\"\",\"appId\":\"\",\"secretId\":\"\",\"secretKey\":\"\"}', '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_config` VALUES (9, 'sms', 'huawei', '{\"name\":\"华为云短信\",\"alias\":\"huawei\"}', '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_config` VALUES (10, 'website', 'name', 'x_admin开源系统', '2024-01-02 03:04:05', '2024-06-29 00:30:51');
INSERT INTO `x_system_config` VALUES (11, 'website', 'logo', '/api/static/backend_logo.png', '2024-01-02 03:04:05', '2024-06-29 00:30:51');
INSERT INTO `x_system_config` VALUES (12, 'website', 'favicon', '/api/static/backend_favicon.ico', '2024-01-02 03:04:05', '2024-06-29 00:30:51');
INSERT INTO `x_system_config` VALUES (13, 'website', 'backdrop', '/api/static/backend_backdrop.png', '2024-01-02 03:04:05', '2024-06-29 00:30:51');
INSERT INTO `x_system_config` VALUES (14, 'website', 'copyright', '[{\"name\":\"蜀ICP备15007060号-1\",\"link\":\"http://www.beian.gov.cn\"},{\"name\":\"x_admin\",\"link\":\"http://x.adtk.cn\"}]', '2024-01-02 03:04:05', '2024-06-29 00:30:54');
INSERT INTO `x_system_config` VALUES (15, 'website', 'shopName', 'x_admin开源管理系统', '2024-01-02 03:04:05', '2024-06-29 00:30:51');
INSERT INTO `x_system_config` VALUES (16, 'website', 'shopLogo', '/image/20241705/f0eb36d508834bc2ac1b8c591c563efa.png', '2024-01-02 03:04:05', '2024-06-29 00:30:51');
INSERT INTO `x_system_config` VALUES (17, 'protocol', 'service', '{\"name\":\"服务协议\",\"content\":\"\\u003cp\\u003e服务协议666\\u003c/p\\u003e\"}', '2024-01-02 03:04:05', '2024-06-29 00:30:56');
INSERT INTO `x_system_config` VALUES (18, 'protocol', 'privacy', '{\"name\":\"隐私协议\",\"content\":\"\\u003cp\\u003e隐私协议\\u003c/p\\u003e\"}', '2024-01-02 03:04:05', '2024-06-29 00:30:56');
INSERT INTO `x_system_config` VALUES (19, 'tabbar', 'style', '{\"defaultColor\":\"#4A5DFF\",\"selectedColor\":\"#EA5455\"}', '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_config` VALUES (20, 'search', 'isHotSearch', '0', '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_config` VALUES (30, 'h5_channel', 'status', '1', '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_config` VALUES (31, 'h5_channel', 'close', '0', '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_config` VALUES (32, 'h5_channel', 'url', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_config` VALUES (40, 'mp_channel', 'name', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_config` VALUES (41, 'mp_channel', 'primaryId', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_config` VALUES (42, 'mp_channel', 'appId', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_config` VALUES (43, 'mp_channel', 'appSecret', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_config` VALUES (44, 'mp_channel', 'qrCode', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_config` VALUES (50, 'wx_channel', 'appId', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_config` VALUES (51, 'wx_channel', 'appSecret', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_config` VALUES (55, 'oa_channel', 'name', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_config` VALUES (56, 'oa_channel', 'primaryId', ' ', '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_config` VALUES (57, 'oa_channel', 'qrCode', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_config` VALUES (58, 'oa_channel', 'appId', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_config` VALUES (59, 'oa_channel', 'appSecret', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_config` VALUES (60, 'oa_channel', 'url', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_config` VALUES (61, 'oa_channel', 'token', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_config` VALUES (62, 'oa_channel', 'encodingAesKey', '', '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_config` VALUES (63, 'oa_channel', 'encryptionType', '1', '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_config` VALUES (64, 'oa_channel', 'menus', '[]', '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_config` VALUES (70, 'login', 'loginWay', '1,2', '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_config` VALUES (71, 'login', 'forceBindMobile', '0', '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_config` VALUES (72, 'login', 'openAgreement', '1', '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_config` VALUES (73, 'login', 'openOtherAuth', '1', '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_config` VALUES (74, 'login', 'autoLoginAuth', '1,2', '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_config` VALUES (80, 'user', 'defaultAvatar', '/api/static/default_avatar.png', '2024-01-02 03:04:05', '2024-01-02 03:04:05');

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
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 164 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统登录日志表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_system_log_login
-- ----------------------------

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
  `start_time` datetime NULL DEFAULT NULL COMMENT '开始时间',
  `end_time` datetime NULL DEFAULT NULL COMMENT '结束时间',
  `task_time` int(11) NULL DEFAULT NULL COMMENT '执行耗时',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 623 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统操作日志表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_system_log_operate
-- ----------------------------

-- ----------------------------
-- Table structure for x_system_log_sms
-- ----------------------------
DROP TABLE IF EXISTS `x_system_log_sms`;
CREATE TABLE `x_system_log_sms`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
  `scene` int(11) UNSIGNED NULL DEFAULT 0 COMMENT '场景编号',
  `mobile` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '手机号码',
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '发送内容',
  `status` tinyint(1) UNSIGNED NULL DEFAULT 0 COMMENT '发送状态：[0=发送中, 1=发送成功, 2=发送失败]',
  `results` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '短信结果',
  `send_time` datetime NULL DEFAULT NULL COMMENT '发送时间',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 117 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统短信日志表' ROW_FORMAT = DYNAMIC;

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
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime NULL DEFAULT NULL COMMENT '删除时间',
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
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `openid`(`openid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户授权表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_user_auth
-- ----------------------------

-- ----------------------------
-- Table structure for x_user_protocol
-- ----------------------------
DROP TABLE IF EXISTS `x_user_protocol`;
CREATE TABLE `x_user_protocol`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标题',
  `content` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '协议内容',
  `sort` double(3, 2) NULL DEFAULT NULL COMMENT '排序',
  `is_delete` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除: 0=否, 1=是',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户协议' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of x_user_protocol
-- ----------------------------
INSERT INTO `x_user_protocol` VALUES (1, '123', '<p>123123</p><p>123123</p><p>123123<img src=\"/api/uploads/image/20241705/f0eb36d508834bc2ac1b8c591c563efa.png\" alt=\"\" data-href=\"\" style=\"\"/></p>', 0.10, 1, '2024-09-11 00:02:48', '2024-09-11 13:14:01', '2024-09-12 11:41:00');
INSERT INTO `x_user_protocol` VALUES (2, '123', '<p>123123</p><p>123123</p><p>123123<img src=\"/api/uploads/image/20241705/f0eb36d508834bc2ac1b8c591c563efa.png\" alt=\"\" data-href=\"\" style=\"\"/></p>', NULL, 1, '2024-09-11 00:02:48', '2024-09-11 12:58:17', '2024-09-11 12:58:17');
INSERT INTO `x_user_protocol` VALUES (3, '123', '<p>123123</p><p>123123</p><p>123123<img src=\"/api/uploads/image/20241705/f0eb36d508834bc2ac1b8c591c563efa.png\" alt=\"\" data-href=\"\" style=\"\"/></p>', NULL, 1, '2024-09-11 00:02:48', '2024-09-11 13:14:01', '2024-09-12 12:20:51');
INSERT INTO `x_user_protocol` VALUES (4, '1', '<p>1</p>', 1.00, 1, '2024-09-12 16:40:10', '2024-09-12 16:40:10', '2024-09-12 16:40:21');
INSERT INTO `x_user_protocol` VALUES (5, '2', '<p>2</p>', 2.00, 1, '2024-09-12 16:40:16', '2024-09-12 16:40:16', '2024-09-13 09:02:24');
INSERT INTO `x_user_protocol` VALUES (6, '123', '<p>123123</p><p>123123</p><p>123123<img src=\"/api/uploads/image/20241705/f0eb36d508834bc2ac1b8c591c563efa.png\" alt=\"\" data-href=\"\" style=\"\"/></p>', NULL, 1, '2024-09-11 00:02:48', '2024-09-11 13:14:01', '2024-09-13 09:02:24');
INSERT INTO `x_user_protocol` VALUES (7, '3', '<p>3</p>', 3.00, 0, '2024-09-13 09:01:53', '2024-09-13 09:01:53', NULL);
INSERT INTO `x_user_protocol` VALUES (8, '4', '<p>4</p>', 4.00, 1, '2024-09-13 09:04:43', '2024-09-13 09:04:43', '2024-09-18 23:09:41');

SET FOREIGN_KEY_CHECKS = 1;
