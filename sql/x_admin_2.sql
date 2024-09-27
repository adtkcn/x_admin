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

 Date: 27/09/2024 10:35:26
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
) ENGINE = InnoDB AUTO_INCREMENT = 34 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '相册管理表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_album
-- ----------------------------
INSERT INTO `x_album` VALUES (31, 6, 1, 0, 10, '6dd3ee16f5822373ffb6318c426cf13 (2).png', 'image/20241305/089ab8212e1e4b4ebe6c5cbe516ea3a5.png', 'png', 11914, 0, '2024-05-13 19:12:39', '2024-05-13 19:12:39', NULL);
INSERT INTO `x_album` VALUES (32, 0, 1, 0, 10, 'appicon.png', 'image/20241705/f0eb36d508834bc2ac1b8c591c563efa.png', 'png', 9873, 0, '2024-05-17 12:15:05', '2024-05-17 12:15:05', NULL);
INSERT INTO `x_album` VALUES (33, 0, 1, 0, 10, 'COMPUMAX_1p2x61.jpg', 'image/20240908/98587b5a927a4a18b5f7d165a3e60287.jpg', 'jpg', 613802, 0, '2024-08-09 00:21:03', '2024-08-09 00:21:03', NULL);

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
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '申请流程' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_flow_apply
-- ----------------------------
INSERT INTO `x_flow_apply` VALUES (5, 10, 1, 'admin', 'a1', 1, '1', '{\"widgetList\":[{\"key\":69408,\"type\":\"input\",\"icon\":\"text-field\",\"formItemFlag\":true,\"options\":{\"name\":\"input30124\",\"label\":\"input\",\"labelAlign\":\"\",\"type\":\"text\",\"defaultValue\":\"\",\"placeholder\":\"\",\"columnWidth\":\"200px\",\"size\":\"\",\"labelWidth\":null,\"labelHidden\":false,\"readonly\":false,\"disabled\":false,\"hidden\":false,\"clearable\":true,\"showPassword\":false,\"required\":false,\"requiredHint\":\"\",\"validation\":\"\",\"validationHint\":\"\",\"customClass\":[],\"labelIconClass\":null,\"labelIconPosition\":\"rear\",\"labelTooltip\":null,\"minLength\":null,\"maxLength\":null,\"showWordLimit\":false,\"prefixIcon\":\"\",\"suffixIcon\":\"\",\"appendButton\":false,\"appendButtonDisabled\":false,\"buttonIcon\":\"custom-search\",\"onCreated\":\"\",\"onMounted\":\"\",\"onInput\":\"\",\"onChange\":\"\",\"onFocus\":\"\",\"onBlur\":\"\",\"onValidate\":\"\",\"onAppendButtonClick\":\"\"},\"id\":\"input30124\"}],\"formConfig\":{\"modelName\":\"formData\",\"refName\":\"vForm\",\"rulesName\":\"rules\",\"labelWidth\":80,\"labelPosition\":\"left\",\"size\":\"\",\"labelAlign\":\"label-left-align\",\"cssCode\":\"\",\"customClass\":[],\"functions\":\"\",\"layoutType\":\"PC\",\"jsonVersion\":3,\"onFormCreated\":\"\",\"onFormMounted\":\"\",\"onFormDataChange\":\"\"}}', '{\"nodes\":[{\"id\":\"Event_6bbfd06\",\"type\":\"bpmn:startEvent\",\"x\":280,\"y\":180,\"properties\":{},\"zIndex\":1004,\"text\":{\"x\":280,\"y\":220,\"value\":\"开始\"}},{\"id\":\"Event_da0ad7c\",\"type\":\"bpmn:endEvent\",\"x\":600,\"y\":220,\"properties\":{},\"zIndex\":1022,\"text\":{\"x\":600,\"y\":260,\"value\":\"结束\"}},{\"id\":\"Activity_03030b3\",\"type\":\"bpmn:userTask\",\"x\":410,\"y\":170,\"properties\":{\"userType\":2,\"userId\":\"\",\"deptId\":\"\",\"postId\":\"\",\"fieldAuth\":{},\"gateway\":[]},\"zIndex\":1008,\"text\":{\"x\":410,\"y\":170,\"value\":\"审批\"}}],\"edges\":[{\"id\":\"45b30f0a-2b12-4cbc-a202-07362358182f\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Event_6bbfd06\",\"targetNodeId\":\"Activity_03030b3\",\"startPoint\":{\"x\":298,\"y\":180},\"endPoint\":{\"x\":360,\"y\":170},\"properties\":{},\"zIndex\":1005,\"pointsList\":[{\"x\":298,\"y\":180},{\"x\":329,\"y\":180},{\"x\":329,\"y\":170},{\"x\":360,\"y\":170}]},{\"id\":\"c07add9e-9e85-4318-900a-afab4f9b86eb\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Activity_03030b3\",\"targetNodeId\":\"Event_da0ad7c\",\"startPoint\":{\"x\":460,\"y\":170},\"endPoint\":{\"x\":600,\"y\":202},\"properties\":{},\"zIndex\":1007,\"pointsList\":[{\"x\":460,\"y\":170},{\"x\":600,\"y\":170},{\"x\":600,\"y\":202}]}]}', '[{\"id\":\"Event_6bbfd06\",\"pid\":0,\"label\":\"开始\",\"type\":\"bpmn:startEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_03030b3\",\"pid\":\"Event_6bbfd06\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{},\"gateway\":[],\"userType\":2,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_da0ad7c\",\"pid\":\"Activity_03030b3\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0}]}]},{\"id\":\"Activity_03030b3\",\"pid\":\"Event_6bbfd06\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{},\"gateway\":[],\"userType\":2,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_da0ad7c\",\"pid\":\"Activity_03030b3\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0}]},{\"id\":\"Event_da0ad7c\",\"pid\":\"Activity_03030b3\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0}]', '{\"input30124\":\"2\"}', 3, 1, '2024-07-03 14:55:47', '2024-07-30 18:10:15', '2024-07-30 18:10:14');

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
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '流程历史' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_flow_history
-- ----------------------------
INSERT INTO `x_flow_history` VALUES (9, 5, 10, 1, 'admin', 0, '', 'Event_6bbfd06', 'bpmn:startEvent', '开始', '{\"input30124\":\"1\"}', 2, '', '2024-07-03 17:04:35', '2024-07-03 17:04:35', NULL);
INSERT INTO `x_flow_history` VALUES (10, 5, 10, 1, 'admin', 1, 'admin', 'Activity_03030b3', 'bpmn:userTask', '审批', '{\"input30124\":\"1\"}', 3, '1', '2024-07-03 17:04:35', '2024-07-03 17:11:50', NULL);
INSERT INTO `x_flow_history` VALUES (11, 5, 10, 1, 'admin', 0, '', 'Event_6bbfd06', 'bpmn:startEvent', '开始', '{\"input30124\":\"1\"}', 2, '', '2024-07-03 17:05:49', '2024-07-03 17:24:53', NULL);
INSERT INTO `x_flow_history` VALUES (12, 5, 10, 1, 'admin', 1, 'admin', 'Activity_03030b3', 'bpmn:userTask', '审批', '{\"input30124\":\"1\"}', 2, '', '2024-07-03 17:24:52', '2024-07-13 20:00:46', NULL);
INSERT INTO `x_flow_history` VALUES (13, 5, 10, 1, 'admin', 0, '', 'Event_da0ad7c', 'bpmn:endEvent', '结束', '{\"input30124\":\"1\"}', 2, '', '2024-07-13 20:00:46', '2024-07-13 20:00:46', NULL);

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
INSERT INTO `x_flow_template` VALUES (10, 'a1', 1, '1', '{\"widgetList\":[{\"key\":69408,\"type\":\"input\",\"icon\":\"text-field\",\"formItemFlag\":true,\"options\":{\"name\":\"input30124\",\"label\":\"input\",\"labelAlign\":\"\",\"type\":\"text\",\"defaultValue\":\"\",\"placeholder\":\"\",\"columnWidth\":\"200px\",\"size\":\"\",\"labelWidth\":null,\"labelHidden\":false,\"readonly\":false,\"disabled\":false,\"hidden\":false,\"clearable\":true,\"showPassword\":false,\"required\":false,\"requiredHint\":\"\",\"validation\":\"\",\"validationHint\":\"\",\"customClass\":[],\"labelIconClass\":null,\"labelIconPosition\":\"rear\",\"labelTooltip\":null,\"minLength\":null,\"maxLength\":null,\"showWordLimit\":false,\"prefixIcon\":\"\",\"suffixIcon\":\"\",\"appendButton\":false,\"appendButtonDisabled\":false,\"buttonIcon\":\"custom-search\",\"onCreated\":\"\",\"onMounted\":\"\",\"onInput\":\"\",\"onChange\":\"\",\"onFocus\":\"\",\"onBlur\":\"\",\"onValidate\":\"\",\"onAppendButtonClick\":\"\"},\"id\":\"input30124\"}],\"formConfig\":{\"modelName\":\"formData\",\"refName\":\"vForm\",\"rulesName\":\"rules\",\"labelWidth\":80,\"labelPosition\":\"left\",\"size\":\"\",\"labelAlign\":\"label-left-align\",\"cssCode\":\"\",\"customClass\":[],\"functions\":\"\",\"layoutType\":\"PC\",\"jsonVersion\":3,\"onFormCreated\":\"\",\"onFormMounted\":\"\",\"onFormDataChange\":\"\"}}', '{\"nodes\":[{\"id\":\"Event_6bbfd06\",\"type\":\"bpmn:startEvent\",\"x\":280,\"y\":180,\"properties\":{},\"zIndex\":1004,\"text\":{\"x\":280,\"y\":220,\"value\":\"开始\"}},{\"id\":\"Event_da0ad7c\",\"type\":\"bpmn:endEvent\",\"x\":600,\"y\":220,\"properties\":{},\"zIndex\":1022,\"text\":{\"x\":600,\"y\":260,\"value\":\"结束\"}},{\"id\":\"Activity_03030b3\",\"type\":\"bpmn:userTask\",\"x\":410,\"y\":170,\"properties\":{\"userType\":2,\"userId\":\"\",\"deptId\":\"\",\"postId\":\"\",\"fieldAuth\":{},\"gateway\":[]},\"zIndex\":1008,\"text\":{\"x\":410,\"y\":170,\"value\":\"审批\"}}],\"edges\":[{\"id\":\"45b30f0a-2b12-4cbc-a202-07362358182f\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Event_6bbfd06\",\"targetNodeId\":\"Activity_03030b3\",\"startPoint\":{\"x\":298,\"y\":180},\"endPoint\":{\"x\":360,\"y\":170},\"properties\":{},\"zIndex\":1005,\"pointsList\":[{\"x\":298,\"y\":180},{\"x\":329,\"y\":180},{\"x\":329,\"y\":170},{\"x\":360,\"y\":170}]},{\"id\":\"c07add9e-9e85-4318-900a-afab4f9b86eb\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Activity_03030b3\",\"targetNodeId\":\"Event_da0ad7c\",\"startPoint\":{\"x\":460,\"y\":170},\"endPoint\":{\"x\":600,\"y\":202},\"properties\":{},\"zIndex\":1007,\"pointsList\":[{\"x\":460,\"y\":170},{\"x\":600,\"y\":170},{\"x\":600,\"y\":202}]}]}', '[{\"id\":\"Event_6bbfd06\",\"pid\":0,\"label\":\"开始\",\"type\":\"bpmn:startEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_03030b3\",\"pid\":\"Event_6bbfd06\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{},\"gateway\":[],\"userType\":2,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_da0ad7c\",\"pid\":\"Activity_03030b3\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0}]}]},{\"id\":\"Activity_03030b3\",\"pid\":\"Event_6bbfd06\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{},\"gateway\":[],\"userType\":2,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_da0ad7c\",\"pid\":\"Activity_03030b3\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0}]},{\"id\":\"Event_da0ad7c\",\"pid\":\"Activity_03030b3\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0}]', 0, '2024-05-14 00:50:31', '2024-07-03 14:06:54', NULL);
INSERT INTO `x_flow_template` VALUES (11, '123', 3, '2', '{\"widgetList\":[{\"key\":69408,\"type\":\"input\",\"icon\":\"text-field\",\"formItemFlag\":true,\"options\":{\"name\":\"input30124\",\"label\":\"input\",\"labelAlign\":\"\",\"type\":\"text\",\"defaultValue\":\"\",\"placeholder\":\"\",\"columnWidth\":\"200px\",\"size\":\"\",\"labelWidth\":null,\"labelHidden\":false,\"readonly\":false,\"disabled\":false,\"hidden\":false,\"clearable\":true,\"showPassword\":false,\"required\":false,\"requiredHint\":\"\",\"validation\":\"\",\"validationHint\":\"\",\"customClass\":[],\"labelIconClass\":null,\"labelIconPosition\":\"rear\",\"labelTooltip\":null,\"minLength\":null,\"maxLength\":null,\"showWordLimit\":false,\"prefixIcon\":\"\",\"suffixIcon\":\"\",\"appendButton\":false,\"appendButtonDisabled\":false,\"buttonIcon\":\"custom-search\",\"onCreated\":\"\",\"onMounted\":\"\",\"onInput\":\"\",\"onChange\":\"\",\"onFocus\":\"\",\"onBlur\":\"\",\"onValidate\":\"\",\"onAppendButtonClick\":\"\"},\"id\":\"input30124\"},{\"key\":51339,\"type\":\"textarea\",\"icon\":\"textarea-field\",\"formItemFlag\":true,\"options\":{\"name\":\"textarea14310\",\"label\":\"textarea\",\"labelAlign\":\"\",\"rows\":3,\"defaultValue\":\"\",\"placeholder\":\"\",\"columnWidth\":\"200px\",\"size\":\"\",\"labelWidth\":null,\"labelHidden\":false,\"readonly\":false,\"disabled\":false,\"hidden\":false,\"required\":false,\"requiredHint\":\"\",\"validation\":\"\",\"validationHint\":\"\",\"customClass\":[],\"labelIconClass\":null,\"labelIconPosition\":\"rear\",\"labelTooltip\":null,\"minLength\":null,\"maxLength\":null,\"showWordLimit\":false,\"onCreated\":\"\",\"onMounted\":\"\",\"onInput\":\"\",\"onChange\":\"\",\"onFocus\":\"\",\"onBlur\":\"\",\"onValidate\":\"\"},\"id\":\"textarea14310\"},{\"key\":18101,\"type\":\"select\",\"icon\":\"select-field\",\"formItemFlag\":true,\"options\":{\"name\":\"select80984\",\"label\":\"select\",\"labelAlign\":\"\",\"defaultValue\":\"\",\"placeholder\":\"\",\"columnWidth\":\"200px\",\"size\":\"\",\"labelWidth\":null,\"labelHidden\":false,\"disabled\":false,\"hidden\":false,\"clearable\":true,\"filterable\":false,\"allowCreate\":false,\"remote\":false,\"automaticDropdown\":false,\"multiple\":false,\"multipleLimit\":0,\"optionItems\":[{\"label\":\"select 1\",\"value\":1},{\"label\":\"select 2\",\"value\":2},{\"label\":\"select 3\",\"value\":3}],\"required\":false,\"requiredHint\":\"\",\"validation\":\"\",\"validationHint\":\"\",\"customClass\":[],\"labelIconClass\":null,\"labelIconPosition\":\"rear\",\"labelTooltip\":null,\"onCreated\":\"\",\"onMounted\":\"\",\"onRemoteQuery\":\"\",\"onChange\":\"\",\"onFocus\":\"\",\"onBlur\":\"\",\"onValidate\":\"\"},\"id\":\"select80984\"}],\"formConfig\":{\"modelName\":\"formData\",\"refName\":\"vForm\",\"rulesName\":\"rules\",\"labelWidth\":80,\"labelPosition\":\"left\",\"size\":\"\",\"labelAlign\":\"label-left-align\",\"cssCode\":\"\",\"customClass\":[],\"functions\":\"\",\"layoutType\":\"PC\",\"jsonVersion\":3,\"onFormCreated\":\"\",\"onFormMounted\":\"\",\"onFormDataChange\":\"\"}}', '{\"nodes\":[{\"id\":\"Event_010c426\",\"type\":\"bpmn:startEvent\",\"x\":430,\"y\":190,\"properties\":{},\"zIndex\":1005,\"text\":{\"x\":430,\"y\":230,\"value\":\"开始\"}},{\"id\":\"Activity_77aa833\",\"type\":\"bpmn:userTask\",\"x\":550,\"y\":190,\"properties\":{\"userType\":1,\"userId\":null,\"deptId\":2,\"postId\":2,\"fieldAuth\":{\"input30124\":3,\"textarea14310\":3,\"select80984\":3},\"gateway\":[]},\"zIndex\":1039,\"text\":{\"x\":550,\"y\":190,\"value\":\"审批\"}},{\"id\":\"Event_ff8d83d\",\"type\":\"bpmn:endEvent\",\"x\":1090,\"y\":190,\"properties\":{},\"zIndex\":1038,\"text\":{\"x\":1090,\"y\":230,\"value\":\"结束\"}},{\"id\":\"Gateway_68e52a5\",\"type\":\"bpmn:exclusiveGateway\",\"x\":700,\"y\":130,\"properties\":{\"userType\":null,\"userId\":null,\"deptId\":null,\"postId\":null,\"fieldAuth\":{\"input30124\":1,\"textarea14310\":1,\"select80984\":1},\"gateway\":[{\"id\":\"input30124\",\"value\":\"1\",\"condition\":\">=\"}]},\"zIndex\":1046,\"text\":{\"x\":700,\"y\":170,\"value\":\"条件\"}},{\"id\":\"Gateway_436010b\",\"type\":\"bpmn:exclusiveGateway\",\"x\":700,\"y\":260,\"properties\":{\"userType\":null,\"userId\":null,\"deptId\":null,\"postId\":null,\"fieldAuth\":{\"input30124\":1,\"textarea14310\":1,\"select80984\":1},\"gateway\":[{\"id\":\"input30124\",\"value\":\"0\",\"condition\":\"<=\"}]},\"zIndex\":1047,\"text\":{\"x\":700,\"y\":300,\"value\":\"条件\"}},{\"id\":\"Activity_2d912e5\",\"type\":\"bpmn:userTask\",\"x\":890,\"y\":130,\"properties\":{\"userType\":null,\"userId\":null,\"deptId\":null,\"postId\":null,\"fieldAuth\":{\"input30124\":1,\"textarea14310\":1,\"select80984\":1},\"gateway\":[]},\"zIndex\":1049,\"text\":{\"x\":890,\"y\":130,\"value\":\"审批\"}},{\"id\":\"Activity_9fced2f\",\"type\":\"bpmn:userTask\",\"x\":890,\"y\":260,\"properties\":{\"userType\":null,\"userId\":null,\"deptId\":null,\"postId\":null,\"fieldAuth\":{\"input30124\":1,\"textarea14310\":1,\"select80984\":1},\"gateway\":[]},\"zIndex\":1048,\"text\":{\"x\":890,\"y\":260,\"value\":\"审批\"}}],\"edges\":[{\"id\":\"b042a41a-d2d4-4483-8e90-1acfcc990738\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Event_010c426\",\"targetNodeId\":\"Activity_77aa833\",\"startPoint\":{\"x\":448,\"y\":190},\"endPoint\":{\"x\":500,\"y\":190},\"properties\":{},\"zIndex\":1006,\"pointsList\":[{\"x\":448,\"y\":190},{\"x\":478,\"y\":190},{\"x\":478,\"y\":190},{\"x\":470,\"y\":190},{\"x\":470,\"y\":190},{\"x\":500,\"y\":190}]},{\"id\":\"34eba9ed-2afa-4867-aea1-4f0d73fe019e\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Activity_77aa833\",\"targetNodeId\":\"Gateway_68e52a5\",\"startPoint\":{\"x\":600,\"y\":190},\"endPoint\":{\"x\":675,\"y\":130},\"properties\":{},\"zIndex\":1011,\"pointsList\":[{\"x\":600,\"y\":190},{\"x\":645,\"y\":190},{\"x\":645,\"y\":130},{\"x\":675,\"y\":130}]},{\"id\":\"d50a1fda-8753-4645-8909-c8493dad946c\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Activity_77aa833\",\"targetNodeId\":\"Gateway_436010b\",\"startPoint\":{\"x\":600,\"y\":190},\"endPoint\":{\"x\":675,\"y\":260},\"properties\":{},\"zIndex\":1013,\"pointsList\":[{\"x\":600,\"y\":190},{\"x\":645,\"y\":190},{\"x\":645,\"y\":260},{\"x\":675,\"y\":260}]},{\"id\":\"2b968694-52b7-4bda-a8f1-785c5caed473\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Gateway_68e52a5\",\"targetNodeId\":\"Activity_2d912e5\",\"startPoint\":{\"x\":725,\"y\":130},\"endPoint\":{\"x\":840,\"y\":130},\"properties\":{},\"zIndex\":1026,\"pointsList\":[{\"x\":725,\"y\":130},{\"x\":840,\"y\":130}]},{\"id\":\"f91e85de-c5d0-4d6f-bdc2-25a2352b95ad\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Gateway_436010b\",\"targetNodeId\":\"Activity_9fced2f\",\"startPoint\":{\"x\":725,\"y\":260},\"endPoint\":{\"x\":840,\"y\":260},\"properties\":{},\"zIndex\":1029,\"pointsList\":[{\"x\":725,\"y\":260},{\"x\":840,\"y\":260}]},{\"id\":\"55e2d6ca-2437-4707-a411-f9243bd0c780\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Activity_2d912e5\",\"targetNodeId\":\"Event_ff8d83d\",\"startPoint\":{\"x\":940,\"y\":130},\"endPoint\":{\"x\":1072,\"y\":190},\"properties\":{},\"zIndex\":1034,\"pointsList\":[{\"x\":940,\"y\":130},{\"x\":1042,\"y\":130},{\"x\":1042,\"y\":190},{\"x\":1072,\"y\":190}]},{\"id\":\"03e08ae1-ac6f-410c-b81b-1c8a4b8e3cd3\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Activity_9fced2f\",\"targetNodeId\":\"Event_ff8d83d\",\"startPoint\":{\"x\":940,\"y\":260},\"endPoint\":{\"x\":1072,\"y\":190},\"properties\":{},\"zIndex\":1037,\"pointsList\":[{\"x\":940,\"y\":260},{\"x\":1042,\"y\":260},{\"x\":1042,\"y\":190},{\"x\":1072,\"y\":190}]}]}', '[{\"id\":\"Event_010c426\",\"pid\":0,\"label\":\"开始\",\"type\":\"bpmn:startEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_77aa833\",\"pid\":\"Event_010c426\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"input30124\":3,\"textarea14310\":3,\"select80984\":3},\"gateway\":[],\"userType\":1,\"userId\":0,\"deptId\":2,\"postId\":2,\"children\":[{\"id\":\"Gateway_68e52a5\",\"pid\":\"Activity_77aa833\",\"label\":\"条件\",\"type\":\"bpmn:exclusiveGateway\",\"fieldAuth\":{\"input30124\":1,\"textarea14310\":1,\"select80984\":1},\"gateway\":[{\"id\":\"input30124\",\"value\":\"1\",\"condition\":\">=\"}],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_2d912e5\",\"pid\":\"Gateway_68e52a5\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"input30124\":1,\"textarea14310\":1,\"select80984\":1},\"gateway\":[],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_ff8d83d\",\"pid\":\"Activity_2d912e5\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":null}]}]},{\"id\":\"Gateway_436010b\",\"pid\":\"Activity_77aa833\",\"label\":\"条件\",\"type\":\"bpmn:exclusiveGateway\",\"fieldAuth\":{\"input30124\":1,\"textarea14310\":1,\"select80984\":1},\"gateway\":[{\"id\":\"input30124\",\"value\":\"0\",\"condition\":\"<=\"}],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_9fced2f\",\"pid\":\"Gateway_436010b\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"input30124\":1,\"textarea14310\":1,\"select80984\":1},\"gateway\":[],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_ff8d83d\",\"pid\":\"Activity_9fced2f\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":null}]}]}]}]},{\"id\":\"Activity_77aa833\",\"pid\":\"Event_010c426\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"input30124\":3,\"textarea14310\":3,\"select80984\":3},\"gateway\":[],\"userType\":1,\"userId\":0,\"deptId\":2,\"postId\":2,\"children\":[{\"id\":\"Gateway_68e52a5\",\"pid\":\"Activity_77aa833\",\"label\":\"条件\",\"type\":\"bpmn:exclusiveGateway\",\"fieldAuth\":{\"input30124\":1,\"textarea14310\":1,\"select80984\":1},\"gateway\":[{\"id\":\"input30124\",\"value\":\"1\",\"condition\":\">=\"}],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_2d912e5\",\"pid\":\"Gateway_68e52a5\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"input30124\":1,\"textarea14310\":1,\"select80984\":1},\"gateway\":[],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_ff8d83d\",\"pid\":\"Activity_2d912e5\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":null}]}]},{\"id\":\"Gateway_436010b\",\"pid\":\"Activity_77aa833\",\"label\":\"条件\",\"type\":\"bpmn:exclusiveGateway\",\"fieldAuth\":{\"input30124\":1,\"textarea14310\":1,\"select80984\":1},\"gateway\":[{\"id\":\"input30124\",\"value\":\"0\",\"condition\":\"<=\"}],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_9fced2f\",\"pid\":\"Gateway_436010b\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"input30124\":1,\"textarea14310\":1,\"select80984\":1},\"gateway\":[],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_ff8d83d\",\"pid\":\"Activity_9fced2f\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":null}]}]}]},{\"id\":\"Gateway_68e52a5\",\"pid\":\"Activity_77aa833\",\"label\":\"条件\",\"type\":\"bpmn:exclusiveGateway\",\"fieldAuth\":{\"input30124\":1,\"textarea14310\":1,\"select80984\":1},\"gateway\":[{\"id\":\"input30124\",\"value\":\"1\",\"condition\":\">=\"}],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_2d912e5\",\"pid\":\"Gateway_68e52a5\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"input30124\":1,\"textarea14310\":1,\"select80984\":1},\"gateway\":[],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_ff8d83d\",\"pid\":\"Activity_2d912e5\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":null}]}]},{\"id\":\"Activity_2d912e5\",\"pid\":\"Gateway_68e52a5\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"input30124\":1,\"textarea14310\":1,\"select80984\":1},\"gateway\":[],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_ff8d83d\",\"pid\":\"Activity_2d912e5\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":null}]},{\"id\":\"Event_ff8d83d\",\"pid\":\"Activity_2d912e5\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":null},{\"id\":\"Gateway_436010b\",\"pid\":\"Activity_77aa833\",\"label\":\"条件\",\"type\":\"bpmn:exclusiveGateway\",\"fieldAuth\":{\"input30124\":1,\"textarea14310\":1,\"select80984\":1},\"gateway\":[{\"id\":\"input30124\",\"value\":\"0\",\"condition\":\"<=\"}],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_9fced2f\",\"pid\":\"Gateway_436010b\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"input30124\":1,\"textarea14310\":1,\"select80984\":1},\"gateway\":[],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_ff8d83d\",\"pid\":\"Activity_9fced2f\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":null}]}]},{\"id\":\"Activity_9fced2f\",\"pid\":\"Gateway_436010b\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"input30124\":1,\"textarea14310\":1,\"select80984\":1},\"gateway\":[],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_ff8d83d\",\"pid\":\"Activity_9fced2f\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":null}]},{\"id\":\"Event_ff8d83d\",\"pid\":\"Activity_9fced2f\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":null},{\"id\":\"Activity_77aa833\",\"pid\":0,\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"input30124\":3,\"textarea14310\":3,\"select80984\":3},\"gateway\":[],\"userType\":1,\"userId\":0,\"deptId\":2,\"postId\":2,\"children\":[{\"id\":\"Gateway_68e52a5\",\"pid\":\"Activity_77aa833\",\"label\":\"条件\",\"type\":\"bpmn:exclusiveGateway\",\"fieldAuth\":{\"input30124\":1,\"textarea14310\":1,\"select80984\":1},\"gateway\":[{\"id\":\"input30124\",\"value\":\"1\",\"condition\":\">=\"}],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_2d912e5\",\"pid\":\"Gateway_68e52a5\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"input30124\":1,\"textarea14310\":1,\"select80984\":1},\"gateway\":[],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_ff8d83d\",\"pid\":\"Activity_2d912e5\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":null}]}]},{\"id\":\"Gateway_436010b\",\"pid\":\"Activity_77aa833\",\"label\":\"条件\",\"type\":\"bpmn:exclusiveGateway\",\"fieldAuth\":{\"input30124\":1,\"textarea14310\":1,\"select80984\":1},\"gateway\":[{\"id\":\"input30124\",\"value\":\"0\",\"condition\":\"<=\"}],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_9fced2f\",\"pid\":\"Gateway_436010b\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"input30124\":1,\"textarea14310\":1,\"select80984\":1},\"gateway\":[],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_ff8d83d\",\"pid\":\"Activity_9fced2f\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":null}]}]}]},{\"id\":\"Gateway_68e52a5\",\"pid\":\"Activity_77aa833\",\"label\":\"条件\",\"type\":\"bpmn:exclusiveGateway\",\"fieldAuth\":{\"input30124\":1,\"textarea14310\":1,\"select80984\":1},\"gateway\":[{\"id\":\"input30124\",\"value\":\"1\",\"condition\":\">=\"}],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_2d912e5\",\"pid\":\"Gateway_68e52a5\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"input30124\":1,\"textarea14310\":1,\"select80984\":1},\"gateway\":[],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_ff8d83d\",\"pid\":\"Activity_2d912e5\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":null}]}]},{\"id\":\"Activity_2d912e5\",\"pid\":\"Gateway_68e52a5\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"input30124\":1,\"textarea14310\":1,\"select80984\":1},\"gateway\":[],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_ff8d83d\",\"pid\":\"Activity_2d912e5\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":null}]},{\"id\":\"Event_ff8d83d\",\"pid\":\"Activity_2d912e5\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":null},{\"id\":\"Gateway_436010b\",\"pid\":\"Activity_77aa833\",\"label\":\"条件\",\"type\":\"bpmn:exclusiveGateway\",\"fieldAuth\":{\"input30124\":1,\"textarea14310\":1,\"select80984\":1},\"gateway\":[{\"id\":\"input30124\",\"value\":\"0\",\"condition\":\"<=\"}],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_9fced2f\",\"pid\":\"Gateway_436010b\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"input30124\":1,\"textarea14310\":1,\"select80984\":1},\"gateway\":[],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_ff8d83d\",\"pid\":\"Activity_9fced2f\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":null}]}]},{\"id\":\"Activity_9fced2f\",\"pid\":\"Gateway_436010b\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"input30124\":1,\"textarea14310\":1,\"select80984\":1},\"gateway\":[],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_ff8d83d\",\"pid\":\"Activity_9fced2f\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":null}]},{\"id\":\"Event_ff8d83d\",\"pid\":\"Activity_9fced2f\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":null},{\"id\":\"Event_ff8d83d\",\"pid\":0,\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":null},{\"id\":\"Gateway_68e52a5\",\"pid\":0,\"label\":\"条件\",\"type\":\"bpmn:exclusiveGateway\",\"fieldAuth\":{\"input30124\":1,\"textarea14310\":1,\"select80984\":1},\"gateway\":[{\"id\":\"input30124\",\"value\":\"1\",\"condition\":\">=\"}],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_2d912e5\",\"pid\":\"Gateway_68e52a5\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"input30124\":1,\"textarea14310\":1,\"select80984\":1},\"gateway\":[],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_ff8d83d\",\"pid\":\"Activity_2d912e5\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":null}]}]},{\"id\":\"Activity_2d912e5\",\"pid\":\"Gateway_68e52a5\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"input30124\":1,\"textarea14310\":1,\"select80984\":1},\"gateway\":[],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_ff8d83d\",\"pid\":\"Activity_2d912e5\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":null}]},{\"id\":\"Event_ff8d83d\",\"pid\":\"Activity_2d912e5\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":null},{\"id\":\"Gateway_436010b\",\"pid\":0,\"label\":\"条件\",\"type\":\"bpmn:exclusiveGateway\",\"fieldAuth\":{\"input30124\":1,\"textarea14310\":1,\"select80984\":1},\"gateway\":[{\"id\":\"input30124\",\"value\":\"0\",\"condition\":\"<=\"}],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_9fced2f\",\"pid\":\"Gateway_436010b\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"input30124\":1,\"textarea14310\":1,\"select80984\":1},\"gateway\":[],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_ff8d83d\",\"pid\":\"Activity_9fced2f\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":null}]}]},{\"id\":\"Activity_9fced2f\",\"pid\":\"Gateway_436010b\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"input30124\":1,\"textarea14310\":1,\"select80984\":1},\"gateway\":[],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_ff8d83d\",\"pid\":\"Activity_9fced2f\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":null}]},{\"id\":\"Event_ff8d83d\",\"pid\":\"Activity_9fced2f\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":null},{\"id\":\"Activity_2d912e5\",\"pid\":0,\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"input30124\":1,\"textarea14310\":1,\"select80984\":1},\"gateway\":[],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_ff8d83d\",\"pid\":\"Activity_2d912e5\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":null}]},{\"id\":\"Event_ff8d83d\",\"pid\":\"Activity_2d912e5\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":null},{\"id\":\"Activity_9fced2f\",\"pid\":0,\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"input30124\":1,\"textarea14310\":1,\"select80984\":1},\"gateway\":[],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_ff8d83d\",\"pid\":\"Activity_9fced2f\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":null}]},{\"id\":\"Event_ff8d83d\",\"pid\":\"Activity_9fced2f\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":null}]', 0, '2024-05-14 01:16:44', '2024-08-06 01:04:37', NULL);
INSERT INTO `x_flow_template` VALUES (12, '111', 2, '111', '{\"widgetList\":[{\"key\":71691,\"type\":\"input\",\"icon\":\"text-field\",\"formItemFlag\":true,\"options\":{\"name\":\"input51067\",\"label\":\"input\",\"labelAlign\":\"\",\"type\":\"text\",\"defaultValue\":\"\",\"placeholder\":\"\",\"columnWidth\":\"200px\",\"size\":\"\",\"labelWidth\":null,\"labelHidden\":false,\"readonly\":false,\"disabled\":false,\"hidden\":false,\"clearable\":true,\"showPassword\":false,\"required\":false,\"requiredHint\":\"\",\"validation\":\"\",\"validationHint\":\"\",\"customClass\":\"\",\"labelIconClass\":null,\"labelIconPosition\":\"rear\",\"labelTooltip\":null,\"minLength\":null,\"maxLength\":null,\"showWordLimit\":false,\"prefixIcon\":\"\",\"suffixIcon\":\"\",\"appendButton\":false,\"appendButtonDisabled\":false,\"buttonIcon\":\"custom-search\",\"onCreated\":\"\",\"onMounted\":\"\",\"onInput\":\"\",\"onChange\":\"\",\"onFocus\":\"\",\"onBlur\":\"\",\"onValidate\":\"\",\"onAppendButtonClick\":\"\"},\"id\":\"input51067\"},{\"key\":56231,\"type\":\"textarea\",\"icon\":\"textarea-field\",\"formItemFlag\":true,\"options\":{\"name\":\"textarea49728\",\"label\":\"textarea\",\"labelAlign\":\"\",\"rows\":3,\"defaultValue\":\"\",\"placeholder\":\"\",\"columnWidth\":\"200px\",\"size\":\"\",\"labelWidth\":null,\"labelHidden\":false,\"readonly\":false,\"disabled\":false,\"hidden\":false,\"required\":false,\"requiredHint\":\"\",\"validation\":\"\",\"validationHint\":\"\",\"customClass\":\"\",\"labelIconClass\":null,\"labelIconPosition\":\"rear\",\"labelTooltip\":null,\"minLength\":null,\"maxLength\":null,\"showWordLimit\":false,\"onCreated\":\"\",\"onMounted\":\"\",\"onInput\":\"\",\"onChange\":\"\",\"onFocus\":\"\",\"onBlur\":\"\",\"onValidate\":\"\"},\"id\":\"textarea49728\"}],\"formConfig\":{\"modelName\":\"formData\",\"refName\":\"vForm\",\"rulesName\":\"rules\",\"labelWidth\":80,\"labelPosition\":\"left\",\"size\":\"\",\"labelAlign\":\"label-left-align\",\"cssCode\":\"\",\"customClass\":\"\",\"functions\":\"\",\"layoutType\":\"PC\",\"jsonVersion\":3,\"onFormCreated\":\"\",\"onFormMounted\":\"\",\"onFormDataChange\":\"\"}}', '{\"nodes\":[{\"id\":\"Event_3a7a9f5\",\"type\":\"bpmn:startEvent\",\"x\":200,\"y\":110,\"properties\":{},\"zIndex\":1007,\"text\":{\"x\":200,\"y\":150,\"value\":\"开始\"}},{\"id\":\"Activity_db9bf6e\",\"type\":\"bpmn:userTask\",\"x\":370,\"y\":140,\"properties\":{},\"zIndex\":1009,\"text\":{\"x\":370,\"y\":140,\"value\":\"审批\"}},{\"id\":\"Event_6637742\",\"type\":\"bpmn:endEvent\",\"x\":520,\"y\":160,\"properties\":{},\"zIndex\":1006,\"text\":{\"x\":520,\"y\":200,\"value\":\"结束\"}}],\"edges\":[{\"id\":\"2c9c10f5-8d24-4ae2-b3be-3e93a15389b0\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Event_3a7a9f5\",\"targetNodeId\":\"Activity_db9bf6e\",\"startPoint\":{\"x\":218,\"y\":110},\"endPoint\":{\"x\":320,\"y\":140},\"properties\":{},\"zIndex\":1008,\"pointsList\":[{\"x\":218,\"y\":110},{\"x\":269,\"y\":110},{\"x\":269,\"y\":140},{\"x\":320,\"y\":140}]},{\"id\":\"3eaee3e9-34ae-48ec-8b71-1ef57434b0f2\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Activity_db9bf6e\",\"targetNodeId\":\"Event_6637742\",\"startPoint\":{\"x\":420,\"y\":140},\"endPoint\":{\"x\":538,\"y\":160},\"properties\":{},\"zIndex\":1010,\"pointsList\":[{\"x\":420,\"y\":140},{\"x\":450,\"y\":140},{\"x\":450,\"y\":112},{\"x\":568,\"y\":112},{\"x\":568,\"y\":160},{\"x\":538,\"y\":160}]}]}', '[{\"id\":\"Event_3a7a9f5\",\"pid\":0,\"label\":\"开始\",\"type\":\"bpmn:startEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_db9bf6e\",\"pid\":\"Event_3a7a9f5\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_6637742\",\"pid\":\"Activity_db9bf6e\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":null}]}]},{\"id\":\"Activity_db9bf6e\",\"pid\":\"Event_3a7a9f5\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_6637742\",\"pid\":\"Activity_db9bf6e\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":null}]},{\"id\":\"Event_6637742\",\"pid\":\"Activity_db9bf6e\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":null},{\"id\":\"Activity_db9bf6e\",\"pid\":0,\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_6637742\",\"pid\":\"Activity_db9bf6e\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":null}]},{\"id\":\"Event_6637742\",\"pid\":\"Activity_db9bf6e\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":null},{\"id\":\"Event_6637742\",\"pid\":0,\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":null}]', 1, '2024-08-16 00:53:50', '2024-08-16 00:53:50', '2024-08-16 00:54:32');

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
INSERT INTO `x_gen_table` VALUES (22, 'x_user_protocol', '用户协议', '', '', '', 'userProtocol', 'user_protocol', '用户协议', '', '', '', 'crud', '', '2024-09-11 12:38:29', '2024-09-11 12:38:29');
INSERT INTO `x_gen_table` VALUES (23, 'x_monitor_project', '监控项目', '', '', '', 'monitorProject', 'monitor_project', '监控项目', '', '', '', 'crud', '', '2024-09-25 15:38:03', '2024-09-25 15:57:32');
INSERT INTO `x_gen_table` VALUES (24, 'x_monitor_client', '监控-客户端信息', '', '', '', 'monitorClient', 'monitor_client', '监控-客户端信息', '', '', '', 'crud', '', '2024-09-25 16:08:23', '2024-09-25 16:09:38');
INSERT INTO `x_gen_table` VALUES (25, 'x_monitor_error', '监控-错误列表', '', '', '', 'monitorError', 'monitor_error', '监控-错误列', '', '', '', 'crud', '', '2024-09-25 16:08:23', '2024-09-25 16:24:38');
INSERT INTO `x_gen_table` VALUES (26, 'x_monitor_error_list', '错误对应的用户记录', '', '', '', 'monitorErrorList', 'monitor_error_list', '错误对应的用户记录', '', '', '', 'crud', '', '2024-09-25 16:08:23', '2024-09-25 16:08:23');

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
INSERT INTO `x_gen_table_column` VALUES (185, 22, 'id', '', '11', 'int', 'int', 'id', 1, 1, 1, 0, 1, 0, 0, '=', 'number', '', '', 1, '2024-09-11 12:38:29', '2024-09-11 12:38:29');
INSERT INTO `x_gen_table_column` VALUES (186, 22, 'title', '标题', '50', 'varchar', 'string', 'title', 0, 0, 1, 1, 1, 1, 1, 'LIKE', 'input', '', '', 2, '2024-09-11 12:38:31', '2024-09-11 12:38:31');
INSERT INTO `x_gen_table_column` VALUES (187, 22, 'content', '协议内容', '0', 'mediumtext', 'string', 'content', 0, 0, 1, 1, 1, 1, 1, '=', 'editor', '', '', 3, '2024-09-11 12:38:31', '2024-09-11 12:38:31');
INSERT INTO `x_gen_table_column` VALUES (188, 22, 'sort', '排序', '0', 'double', 'float64', 'sort', 0, 0, 1, 1, 1, 1, 1, '=', 'number', '', '', 4, '2024-09-11 12:38:31', '2024-09-11 12:38:31');
INSERT INTO `x_gen_table_column` VALUES (189, 22, 'is_delete', '是否删除: 0=否, 1=是', '1', 'tinyint', 'int', 'is_delete', 0, 0, 0, 0, 0, 0, 0, '=', 'number', '', '', 5, '2024-09-11 12:38:31', '2024-09-11 12:38:31');
INSERT INTO `x_gen_table_column` VALUES (190, 22, 'create_time', '创建时间', '0', 'datetime', 'core.NullTime', 'create_time', 0, 0, 0, 0, 0, 1, 1, '=', 'datetime', '', '', 6, '2024-09-11 12:38:34', '2024-09-11 12:38:34');
INSERT INTO `x_gen_table_column` VALUES (191, 22, 'update_time', '更新时间', '0', 'datetime', 'core.NullTime', 'update_time', 0, 0, 0, 0, 0, 1, 1, '=', 'datetime', '', '', 7, '2024-09-11 12:38:34', '2024-09-11 12:38:34');
INSERT INTO `x_gen_table_column` VALUES (192, 22, 'delete_time', '删除时间', '0', 'datetime', 'core.NullTime', 'delete_time', 0, 0, 0, 0, 0, 0, 0, '=', 'datetime', '', '', 8, '2024-09-11 12:38:34', '2024-09-11 12:38:34');
INSERT INTO `x_gen_table_column` VALUES (193, 23, 'id', '项目id', '11', 'int', 'int', 'id', 1, 1, 1, 0, 1, 0, 0, '=', 'number', '', '', 0, '2024-09-25 15:38:03', '2024-09-25 15:57:32');
INSERT INTO `x_gen_table_column` VALUES (194, 23, 'project_key', '项目uuid', '32', 'varchar', 'string', 'project_key', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', '', 0, '2024-09-25 15:38:09', '2024-09-25 15:57:32');
INSERT INTO `x_gen_table_column` VALUES (195, 23, 'project_name', '项目名称', '50', 'varchar', 'string', 'project_name', 0, 0, 1, 1, 1, 1, 1, 'LIKE', 'input', '', '', 0, '2024-09-25 15:38:09', '2024-09-25 15:57:32');
INSERT INTO `x_gen_table_column` VALUES (196, 23, 'project_type', '项目类型go java web node php 等', '20', 'varchar', 'string', 'project_type', 0, 0, 1, 1, 1, 1, 1, '=', 'select', 'project_type', '', 0, '2024-09-25 15:38:09', '2024-09-25 15:57:32');
INSERT INTO `x_gen_table_column` VALUES (197, 23, 'status', '是否启用: 0=否, 1=是', '1', 'tinyint', 'int', 'status', 0, 0, 1, 1, 1, 1, 1, '=', 'select', 'status', '', 0, '2024-09-25 15:38:09', '2024-09-25 15:57:32');
INSERT INTO `x_gen_table_column` VALUES (198, 23, 'is_delete', '是否删除: 0=否, 1=是', '1', 'tinyint', 'int', 'is_delete', 0, 0, 0, 0, 0, 0, 0, '=', 'number', '', '', 0, '2024-09-25 15:38:09', '2024-09-25 15:57:32');
INSERT INTO `x_gen_table_column` VALUES (199, 23, 'create_time', '创建时间', '0', 'datetime', 'core.NullTime', 'create_time', 0, 0, 0, 0, 0, 1, 1, '=', 'datetime', '', '', 0, '2024-09-25 15:38:09', '2024-09-25 15:57:32');
INSERT INTO `x_gen_table_column` VALUES (200, 23, 'update_time', '更新时间', '0', 'datetime', 'core.NullTime', 'update_time', 0, 0, 0, 0, 0, 1, 1, '=', 'datetime', '', '', 0, '2024-09-25 15:38:09', '2024-09-25 15:57:32');
INSERT INTO `x_gen_table_column` VALUES (201, 23, 'delete_time', '删除时间', '0', 'datetime', 'core.NullTime', 'delete_time', 0, 0, 0, 0, 0, 0, 0, '=', 'datetime', '', '', 0, '2024-09-25 15:38:09', '2024-09-25 15:57:32');
INSERT INTO `x_gen_table_column` VALUES (202, 24, 'id', 'uuid', '11', 'int', 'int', 'id', 1, 1, 1, 0, 1, 0, 0, '=', 'number', '', '', 0, '2024-09-25 16:08:23', '2024-09-25 16:09:38');
INSERT INTO `x_gen_table_column` VALUES (203, 24, 'project_key', '项目key', '128', 'varchar', 'string', 'project_key', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', '', 0, '2024-09-25 16:08:23', '2024-09-25 16:09:38');
INSERT INTO `x_gen_table_column` VALUES (204, 24, 'client_id', 'sdk生成的客户端id', '128', 'varchar', 'string', 'client_id', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', '', 0, '2024-09-25 16:08:23', '2024-09-25 16:09:38');
INSERT INTO `x_gen_table_column` VALUES (205, 24, 'user_id', '用户id', '128', 'varchar', 'string', 'user_id', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', '', 0, '2024-09-25 16:08:23', '2024-09-25 16:09:38');
INSERT INTO `x_gen_table_column` VALUES (206, 24, 'os', '系统', '30', 'varchar', 'string', 'os', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', '', 0, '2024-09-25 16:08:23', '2024-09-25 16:09:38');
INSERT INTO `x_gen_table_column` VALUES (207, 24, 'browser', '浏览器', '30', 'varchar', 'string', 'browser', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', '', 0, '2024-09-25 16:08:23', '2024-09-25 16:09:38');
INSERT INTO `x_gen_table_column` VALUES (208, 24, 'city', '城市', '50', 'varchar', 'string', 'city', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', '', 0, '2024-09-25 16:08:23', '2024-09-25 16:09:38');
INSERT INTO `x_gen_table_column` VALUES (209, 24, 'width', '屏幕', '10', 'smallint', 'int', 'width', 0, 0, 1, 1, 1, 1, 1, '=', 'number', '', '', 0, '2024-09-25 16:08:23', '2024-09-25 16:09:38');
INSERT INTO `x_gen_table_column` VALUES (210, 24, 'height', '屏幕高度', '10', 'smallint', 'int', 'height', 0, 0, 1, 1, 1, 1, 1, '=', 'number', '', '', 0, '2024-09-25 16:08:23', '2024-09-25 16:09:38');
INSERT INTO `x_gen_table_column` VALUES (211, 24, 'ua', 'ua记录', '128', 'varchar', 'string', 'ua', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', '', 0, '2024-09-25 16:08:23', '2024-09-25 16:09:38');
INSERT INTO `x_gen_table_column` VALUES (212, 24, 'create_time', '创建时间', '0', 'datetime', 'core.NullTime', 'create_time', 0, 0, 0, 0, 0, 1, 1, '=', 'datetime', '', '', 0, '2024-09-25 16:08:23', '2024-09-25 16:09:38');
INSERT INTO `x_gen_table_column` VALUES (213, 24, 'client_time', '更新时间', '0', 'datetime', 'core.NullTime', 'client_time', 0, 0, 1, 1, 1, 1, 1, '=', 'datetime', '', '', 0, '2024-09-25 16:08:23', '2024-09-25 16:09:38');
INSERT INTO `x_gen_table_column` VALUES (214, 25, 'id', '错误id', '11', 'int', 'int', 'id', 1, 1, 1, 0, 1, 0, 0, '=', 'number', '', '', 0, '2024-09-25 16:08:23', '2024-09-25 16:24:38');
INSERT INTO `x_gen_table_column` VALUES (215, 25, 'project_key', '项目key', '128', 'varchar', 'string', 'project_key', 0, 0, 1, 1, 1, 1, 1, '=', 'select', '', '/api/admin/monitor_project/listAll', 0, '2024-09-25 16:08:23', '2024-09-25 16:24:38');
INSERT INTO `x_gen_table_column` VALUES (216, 25, 'event_type', '事件类型', '20', 'varchar', 'string', 'event_type', 0, 0, 1, 1, 1, 1, 1, '=', 'textarea', '', '', 0, '2024-09-25 16:08:23', '2024-09-25 16:24:38');
INSERT INTO `x_gen_table_column` VALUES (217, 25, 'path', 'URL地址', '1000', 'varchar', 'string', 'path', 0, 0, 1, 1, 1, 1, 1, '=', 'textarea', '', '', 0, '2024-09-25 16:08:23', '2024-09-25 16:24:38');
INSERT INTO `x_gen_table_column` VALUES (218, 25, 'message', '错误消息', '1000', 'varchar', 'string', 'message', 0, 0, 1, 1, 1, 1, 1, '=', 'textarea', '', '', 0, '2024-09-25 16:08:23', '2024-09-25 16:24:38');
INSERT INTO `x_gen_table_column` VALUES (219, 25, 'stack', '错误堆栈', '0', 'text', 'string', 'stack', 0, 0, 1, 1, 1, 1, 1, '=', 'textarea', '', '', 0, '2024-09-25 16:08:23', '2024-09-25 16:24:38');
INSERT INTO `x_gen_table_column` VALUES (220, 25, 'md5', 'md5', '32', 'varchar', 'string', 'md5', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', '', 0, '2024-09-25 16:08:23', '2024-09-25 16:24:38');
INSERT INTO `x_gen_table_column` VALUES (221, 25, 'create_time', '创建时间', '0', 'datetime', 'core.NullTime', 'create_time', 0, 0, 0, 0, 0, 1, 1, '=', 'datetime', '', '', 0, '2024-09-25 16:08:23', '2024-09-25 16:24:38');
INSERT INTO `x_gen_table_column` VALUES (222, 25, 'client_time', '更新时间', '0', 'datetime', 'core.NullTime', 'client_time', 0, 0, 1, 1, 1, 1, 1, '=', 'datetime', '', '', 0, '2024-09-25 16:08:23', '2024-09-25 16:24:38');
INSERT INTO `x_gen_table_column` VALUES (223, 26, 'id', '项目id', '11', 'int', 'int', 'id', 1, 1, 1, 0, 1, 0, 0, '=', 'number', '', '', 1, '2024-09-25 16:08:23', '2024-09-25 16:08:23');
INSERT INTO `x_gen_table_column` VALUES (224, 26, 'error_id', '错误id', '32', 'varchar', 'string', 'error_id', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', '', 2, '2024-09-25 16:08:23', '2024-09-25 16:08:23');
INSERT INTO `x_gen_table_column` VALUES (225, 26, 'client_id', '客户端id', '32', 'varchar', 'string', 'client_id', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', '', 3, '2024-09-25 16:08:23', '2024-09-25 16:08:23');
INSERT INTO `x_gen_table_column` VALUES (226, 26, 'project_key', '项目id', '32', 'varchar', 'string', 'project_key', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', '', 4, '2024-09-25 16:08:23', '2024-09-25 16:08:23');
INSERT INTO `x_gen_table_column` VALUES (227, 26, 'create_time', '创建时间', '0', 'datetime', 'core.NullTime', 'create_time', 0, 0, 0, 0, 0, 1, 1, '=', 'datetime', '', '', 5, '2024-09-25 16:08:23', '2024-09-25 16:08:23');

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
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `client_time` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `project_key`(`project_key`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  INDEX `client_id`(`client_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '监控-客户端信息' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_monitor_client
-- ----------------------------
INSERT INTO `x_monitor_client` VALUES (1, 'e19e3be20de94f49b68fafb4c30668bc', '1', '11', '1', '1', '1', 1, 1, '1', '2024-09-25 16:16:03', '2024-09-25 04:16:02');

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
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '监控项目' ROW_FORMAT = DYNAMIC;

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
INSERT INTO `x_system_auth_admin` VALUES (1, 1, 3, 'admin', 'admin', '81a13dd8e25644a8823082573ca973f7', '/image/20232911/27699a1e991e4ba1ba13ee6d45185db6.jpg', '0', 'WFdiD', 1, 1, 0, 0, '127.0.0.1', '2024-09-25 22:17:45', '2024-01-02 03:04:05', '2024-09-25 22:17:45', NULL);
INSERT INTO `x_system_auth_admin` VALUES (2, 2, 3, 'a2', 'a2', '246994e6b031dd3db14ae444901848c3', '/image/20241705/f0eb36d508834bc2ac1b8c591c563efa.png', '1', 'jXvQo', 1, 0, 0, 0, '192.168.43.34', '2024-06-17 20:54:25', '2024-06-17 20:29:54', '2024-08-16 14:39:10', NULL);
INSERT INTO `x_system_auth_admin` VALUES (3, 1, 2, 'test', 'test', 'a26519d1e28191e0cc643cf997925c82', '/image/20241705/f0eb36d508834bc2ac1b8c591c563efa.png', '2', 'ilEvk', 1, 0, 0, 1, '127.0.0.1', '2024-07-24 13:21:18', '2024-07-24 12:09:44', '2024-07-24 13:49:52', '2024-07-24 13:49:52');

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
INSERT INTO `x_system_auth_dept` VALUES (4, 2, '1', 0, '', '13801310000', 0, 0, 1, '2024-07-01 16:43:20', '2024-07-01 16:43:20', '2024-07-01 16:43:24');

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
INSERT INTO `x_system_auth_post` VALUES (2, 'gw0001', '默认岗位', '', 3, 0, 1, '2024-01-02 03:04:05', '2024-06-29 00:30:21', '2024-08-16 14:39:17');
INSERT INTO `x_system_auth_post` VALUES (3, 'zhihuibu01', '指挥部岗位', '', 0, 0, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05', NULL);
INSERT INTO `x_system_auth_post` VALUES (7, 'zhihuibu02', 'zhihuibu02', '2', 1, 0, 1, '2024-05-13 19:13:44', '2024-05-13 19:13:49', NULL);
INSERT INTO `x_system_auth_post` VALUES (8, 'zhihuibu03', 'zhihuibu03', '3', 1, 0, 1, '2024-05-13 19:57:38', '2024-05-13 19:57:38', '2024-05-13 19:57:55');
INSERT INTO `x_system_auth_post` VALUES (9, '1', '1', '1', 1, 0, 1, '2024-05-19 00:07:46', '2024-05-19 00:07:46', '2024-05-19 00:07:49');
INSERT INTO `x_system_auth_post` VALUES (10, '7', '1', '', 1, 0, 0, '2024-08-16 14:39:22', '2024-08-16 14:39:32', NULL);

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
) ENGINE = InnoDB AUTO_INCREMENT = 154 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统登录日志表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_system_log_login
-- ----------------------------
INSERT INTO `x_system_log_login` VALUES (1, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-06-14 19:39:17');
INSERT INTO `x_system_log_login` VALUES (2, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-06-14 23:13:30');
INSERT INTO `x_system_log_login` VALUES (3, 1, 'admin', '192.168.43.34', 'Windows', 'Edge', 1, '2024-06-17 14:56:19');
INSERT INTO `x_system_log_login` VALUES (4, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-06-17 15:32:21');
INSERT INTO `x_system_log_login` VALUES (5, 1, 'admin', '192.168.43.34', 'Windows', 'Edge', 0, '2024-06-17 15:58:41');
INSERT INTO `x_system_log_login` VALUES (6, 1, 'admin', '192.168.43.34', 'Android', 'Chrome Mobile', 0, '2024-06-17 16:45:41');
INSERT INTO `x_system_log_login` VALUES (7, 1, 'admin', '192.168.43.34', 'Android', 'Chrome Mobile', 0, '2024-06-17 16:49:21');
INSERT INTO `x_system_log_login` VALUES (8, 1, 'admin', '192.168.43.34', 'Android', 'Chrome Mobile', 1, '2024-06-17 17:07:50');
INSERT INTO `x_system_log_login` VALUES (9, 1, 'admin', '192.168.43.34', 'Android', 'Chrome Mobile', 0, '2024-06-17 17:08:18');
INSERT INTO `x_system_log_login` VALUES (10, 0, '123', '192.168.43.34', 'Android', 'Chrome Mobile', 0, '2024-06-17 17:15:39');
INSERT INTO `x_system_log_login` VALUES (11, 1, 'admin', '192.168.43.34', 'Android', 'Chrome Mobile', 0, '2024-06-17 17:17:13');
INSERT INTO `x_system_log_login` VALUES (12, 1, 'admin', '192.168.43.34', 'Android', 'Chrome Mobile', 0, '2024-06-17 17:18:24');
INSERT INTO `x_system_log_login` VALUES (13, 1, 'admin', '192.168.43.34', 'iOS', 'Mobile Safari', 0, '2024-06-17 17:20:16');
INSERT INTO `x_system_log_login` VALUES (14, 1, 'admin', '192.168.43.34', 'iOS', 'Mobile Safari', 0, '2024-06-17 17:20:24');
INSERT INTO `x_system_log_login` VALUES (15, 1, 'admin', '192.168.43.34', 'iOS', 'Mobile Safari', 0, '2024-06-17 17:24:36');
INSERT INTO `x_system_log_login` VALUES (16, 1, 'admin', '192.168.43.34', 'iOS', 'Mobile Safari', 0, '2024-06-17 17:25:53');
INSERT INTO `x_system_log_login` VALUES (17, 1, 'admin', '192.168.43.34', 'iOS', 'Mobile Safari', 0, '2024-06-17 17:26:05');
INSERT INTO `x_system_log_login` VALUES (18, 0, 'sdff', '192.168.43.1', 'Android', 'Chrome Mobile WebView', 0, '2024-06-17 17:27:41');
INSERT INTO `x_system_log_login` VALUES (19, 1, 'admin', '192.168.43.1', 'Android', 'Chrome Mobile WebView', 1, '2024-06-17 17:27:57');
INSERT INTO `x_system_log_login` VALUES (20, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-06-17 17:32:36');
INSERT INTO `x_system_log_login` VALUES (21, 1, 'admin', '192.168.43.34', 'Android', 'Chrome Mobile', 1, '2024-06-17 17:37:52');
INSERT INTO `x_system_log_login` VALUES (22, 1, 'admin', '192.168.43.34', 'Android', 'Chrome Mobile', 1, '2024-06-17 17:44:19');
INSERT INTO `x_system_log_login` VALUES (23, 1, 'admin', '192.168.43.34', 'Android', 'Chrome Mobile', 1, '2024-06-17 18:03:07');
INSERT INTO `x_system_log_login` VALUES (24, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-06-17 18:05:06');
INSERT INTO `x_system_log_login` VALUES (25, 1, 'admin', '192.168.43.34', 'iOS', 'Mobile Safari', 1, '2024-06-17 18:34:48');
INSERT INTO `x_system_log_login` VALUES (26, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-06-17 20:18:01');
INSERT INTO `x_system_log_login` VALUES (27, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-06-17 20:26:01');
INSERT INTO `x_system_log_login` VALUES (28, 2, 'a2', '192.168.43.34', 'Android', 'Chrome Mobile', 1, '2024-06-17 20:30:24');
INSERT INTO `x_system_log_login` VALUES (29, 2, 'a2', '192.168.43.34', 'Android', 'Chrome Mobile', 1, '2024-06-17 20:54:25');
INSERT INTO `x_system_log_login` VALUES (30, 1, 'admin', '192.168.43.34', 'iOS', 'Mobile Safari', 1, '2024-06-17 21:54:04');
INSERT INTO `x_system_log_login` VALUES (31, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-06-28 17:04:12');
INSERT INTO `x_system_log_login` VALUES (32, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-06-28 23:13:55');
INSERT INTO `x_system_log_login` VALUES (33, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-07-01 16:41:28');
INSERT INTO `x_system_log_login` VALUES (34, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-07-01 16:42:26');
INSERT INTO `x_system_log_login` VALUES (35, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-07-01 21:23:49');
INSERT INTO `x_system_log_login` VALUES (36, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-07-03 11:41:48');
INSERT INTO `x_system_log_login` VALUES (37, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-07-03 13:42:46');
INSERT INTO `x_system_log_login` VALUES (38, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-07-03 16:54:26');
INSERT INTO `x_system_log_login` VALUES (39, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-07-03 22:45:27');
INSERT INTO `x_system_log_login` VALUES (40, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-07-04 21:15:45');
INSERT INTO `x_system_log_login` VALUES (41, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-07-05 10:00:45');
INSERT INTO `x_system_log_login` VALUES (42, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-07-05 11:38:52');
INSERT INTO `x_system_log_login` VALUES (43, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-07-05 23:04:18');
INSERT INTO `x_system_log_login` VALUES (44, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-07-08 15:05:00');
INSERT INTO `x_system_log_login` VALUES (45, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-07-08 15:18:02');
INSERT INTO `x_system_log_login` VALUES (46, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-07-08 15:18:58');
INSERT INTO `x_system_log_login` VALUES (47, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-07-08 15:19:37');
INSERT INTO `x_system_log_login` VALUES (48, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-07-08 23:56:26');
INSERT INTO `x_system_log_login` VALUES (49, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-07-09 08:29:46');
INSERT INTO `x_system_log_login` VALUES (50, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-07-09 08:49:02');
INSERT INTO `x_system_log_login` VALUES (51, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-07-09 15:04:20');
INSERT INTO `x_system_log_login` VALUES (52, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-07-09 18:36:26');
INSERT INTO `x_system_log_login` VALUES (53, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-07-10 00:19:09');
INSERT INTO `x_system_log_login` VALUES (54, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-07-10 11:02:31');
INSERT INTO `x_system_log_login` VALUES (55, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-07-12 19:29:58');
INSERT INTO `x_system_log_login` VALUES (56, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-07-12 23:04:55');
INSERT INTO `x_system_log_login` VALUES (57, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-07-13 14:57:11');
INSERT INTO `x_system_log_login` VALUES (58, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-07-13 18:57:59');
INSERT INTO `x_system_log_login` VALUES (59, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-07-24 12:08:27');
INSERT INTO `x_system_log_login` VALUES (60, 3, 'test', '127.0.0.1', 'Windows', 'Firefox', 1, '2024-07-24 12:10:07');
INSERT INTO `x_system_log_login` VALUES (61, 3, 'test', '127.0.0.1', 'Windows', 'Firefox', 1, '2024-07-24 12:29:19');
INSERT INTO `x_system_log_login` VALUES (62, 3, 'test', '127.0.0.1', 'Windows', 'Firefox', 1, '2024-07-24 12:30:22');
INSERT INTO `x_system_log_login` VALUES (63, 3, 'test', '127.0.0.1', 'Windows', 'Firefox', 1, '2024-07-24 12:30:56');
INSERT INTO `x_system_log_login` VALUES (64, 3, 'test', '127.0.0.1', 'Windows', 'Firefox', 1, '2024-07-24 12:31:13');
INSERT INTO `x_system_log_login` VALUES (65, 3, 'test', '127.0.0.1', 'Windows', 'Firefox', 1, '2024-07-24 12:32:50');
INSERT INTO `x_system_log_login` VALUES (66, 3, 'test', '127.0.0.1', 'Windows', 'Firefox', 1, '2024-07-24 12:35:11');
INSERT INTO `x_system_log_login` VALUES (67, 3, 'test', '127.0.0.1', 'Windows', 'Firefox', 1, '2024-07-24 12:35:38');
INSERT INTO `x_system_log_login` VALUES (68, 3, 'test', '127.0.0.1', 'Windows', 'Firefox', 1, '2024-07-24 12:44:25');
INSERT INTO `x_system_log_login` VALUES (69, 3, 'test', '127.0.0.1', 'Windows', 'Firefox', 1, '2024-07-24 12:45:29');
INSERT INTO `x_system_log_login` VALUES (70, 3, 'test', '127.0.0.1', 'Windows', 'Firefox', 1, '2024-07-24 13:21:03');
INSERT INTO `x_system_log_login` VALUES (71, 3, 'test', '127.0.0.1', 'Windows', 'Firefox', 1, '2024-07-24 13:21:18');
INSERT INTO `x_system_log_login` VALUES (72, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-07-29 16:41:06');
INSERT INTO `x_system_log_login` VALUES (73, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-07-30 17:09:19');
INSERT INTO `x_system_log_login` VALUES (74, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-07-31 10:32:46');
INSERT INTO `x_system_log_login` VALUES (75, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-05 11:19:25');
INSERT INTO `x_system_log_login` VALUES (76, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-05 16:45:06');
INSERT INTO `x_system_log_login` VALUES (77, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-05 16:48:46');
INSERT INTO `x_system_log_login` VALUES (78, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-06 00:42:44');
INSERT INTO `x_system_log_login` VALUES (79, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-06 03:28:24');
INSERT INTO `x_system_log_login` VALUES (80, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-07 21:52:49');
INSERT INTO `x_system_log_login` VALUES (81, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-07 21:53:31');
INSERT INTO `x_system_log_login` VALUES (82, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-08 12:32:38');
INSERT INTO `x_system_log_login` VALUES (83, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-08 16:04:43');
INSERT INTO `x_system_log_login` VALUES (84, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-08 22:24:52');
INSERT INTO `x_system_log_login` VALUES (85, 1, 'admin', '192.168.43.34', 'Windows', 'Edge', 1, '2024-08-09 00:48:20');
INSERT INTO `x_system_log_login` VALUES (86, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-09 11:38:46');
INSERT INTO `x_system_log_login` VALUES (87, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-09 17:00:44');
INSERT INTO `x_system_log_login` VALUES (88, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-09 21:04:48');
INSERT INTO `x_system_log_login` VALUES (89, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-10 15:18:30');
INSERT INTO `x_system_log_login` VALUES (90, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-11 01:38:37');
INSERT INTO `x_system_log_login` VALUES (91, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-11 02:05:42');
INSERT INTO `x_system_log_login` VALUES (92, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-11 02:08:23');
INSERT INTO `x_system_log_login` VALUES (93, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-11 02:10:44');
INSERT INTO `x_system_log_login` VALUES (94, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-11 02:13:20');
INSERT INTO `x_system_log_login` VALUES (95, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-11 02:18:49');
INSERT INTO `x_system_log_login` VALUES (96, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-11 02:19:11');
INSERT INTO `x_system_log_login` VALUES (97, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-11 02:19:56');
INSERT INTO `x_system_log_login` VALUES (98, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-11 02:21:11');
INSERT INTO `x_system_log_login` VALUES (99, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-11 02:21:31');
INSERT INTO `x_system_log_login` VALUES (100, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-11 02:24:57');
INSERT INTO `x_system_log_login` VALUES (101, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-11 02:28:52');
INSERT INTO `x_system_log_login` VALUES (102, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-11 02:39:34');
INSERT INTO `x_system_log_login` VALUES (103, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-11 02:48:34');
INSERT INTO `x_system_log_login` VALUES (104, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-11 02:51:41');
INSERT INTO `x_system_log_login` VALUES (105, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-11 03:41:02');
INSERT INTO `x_system_log_login` VALUES (106, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-11 03:51:24');
INSERT INTO `x_system_log_login` VALUES (107, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-11 18:51:28');
INSERT INTO `x_system_log_login` VALUES (108, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-12 00:49:03');
INSERT INTO `x_system_log_login` VALUES (109, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-12 13:28:22');
INSERT INTO `x_system_log_login` VALUES (110, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-12 13:29:40');
INSERT INTO `x_system_log_login` VALUES (111, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-12 15:02:00');
INSERT INTO `x_system_log_login` VALUES (112, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-12 16:36:17');
INSERT INTO `x_system_log_login` VALUES (113, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-12 16:42:20');
INSERT INTO `x_system_log_login` VALUES (114, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-12 18:07:31');
INSERT INTO `x_system_log_login` VALUES (115, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-12 18:20:36');
INSERT INTO `x_system_log_login` VALUES (116, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-12 22:22:18');
INSERT INTO `x_system_log_login` VALUES (117, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-13 01:18:37');
INSERT INTO `x_system_log_login` VALUES (118, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-13 12:15:54');
INSERT INTO `x_system_log_login` VALUES (119, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-13 16:06:42');
INSERT INTO `x_system_log_login` VALUES (120, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-13 22:09:51');
INSERT INTO `x_system_log_login` VALUES (121, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-14 10:47:13');
INSERT INTO `x_system_log_login` VALUES (122, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-14 16:53:17');
INSERT INTO `x_system_log_login` VALUES (123, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-14 19:01:56');
INSERT INTO `x_system_log_login` VALUES (124, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-15 11:10:19');
INSERT INTO `x_system_log_login` VALUES (125, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-15 14:12:56');
INSERT INTO `x_system_log_login` VALUES (126, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-15 16:40:59');
INSERT INTO `x_system_log_login` VALUES (127, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-15 19:50:21');
INSERT INTO `x_system_log_login` VALUES (128, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-16 11:20:51');
INSERT INTO `x_system_log_login` VALUES (129, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-16 17:34:57');
INSERT INTO `x_system_log_login` VALUES (130, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-16 17:36:05');
INSERT INTO `x_system_log_login` VALUES (131, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-16 17:37:18');
INSERT INTO `x_system_log_login` VALUES (132, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-16 17:38:40');
INSERT INTO `x_system_log_login` VALUES (133, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-16 17:54:14');
INSERT INTO `x_system_log_login` VALUES (134, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-08-19 13:29:46');
INSERT INTO `x_system_log_login` VALUES (135, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-09-10 19:45:24');
INSERT INTO `x_system_log_login` VALUES (136, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-09-10 22:08:52');
INSERT INTO `x_system_log_login` VALUES (137, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-09-11 02:07:41');
INSERT INTO `x_system_log_login` VALUES (138, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-09-11 12:13:09');
INSERT INTO `x_system_log_login` VALUES (139, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-09-11 18:02:30');
INSERT INTO `x_system_log_login` VALUES (140, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-09-12 10:59:00');
INSERT INTO `x_system_log_login` VALUES (141, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-09-12 16:34:07');
INSERT INTO `x_system_log_login` VALUES (142, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-09-12 18:56:13');
INSERT INTO `x_system_log_login` VALUES (143, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-09-12 19:31:26');
INSERT INTO `x_system_log_login` VALUES (144, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-09-12 22:38:26');
INSERT INTO `x_system_log_login` VALUES (145, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-09-13 08:57:30');
INSERT INTO `x_system_log_login` VALUES (146, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-09-13 16:30:21');
INSERT INTO `x_system_log_login` VALUES (147, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-09-14 00:25:55');
INSERT INTO `x_system_log_login` VALUES (148, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-09-14 10:06:26');
INSERT INTO `x_system_log_login` VALUES (149, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-09-18 15:27:00');
INSERT INTO `x_system_log_login` VALUES (150, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-09-18 23:01:58');
INSERT INTO `x_system_log_login` VALUES (151, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-09-25 14:50:26');
INSERT INTO `x_system_log_login` VALUES (152, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-09-25 15:41:57');
INSERT INTO `x_system_log_login` VALUES (153, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-09-25 22:17:45');

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
INSERT INTO `x_system_log_operate` VALUES (1, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-06-17 17:33:11', '2024-06-17 17:33:11', 4, '2024-06-17 17:33:11');
INSERT INTO `x_system_log_operate` VALUES (2, 1, 'POST', '错误收集error删除', '127.0.0.1', '/api/admin/monitor_web/del', 'x_admin/admin/monitor_web.MonitorWebHandler.Del-fm', '{\"id\":3}', '', 1, '2024-06-17 17:33:35', '2024-06-17 17:33:35', 99, '2024-06-17 17:33:35');
INSERT INTO `x_system_log_operate` VALUES (3, 1, 'POST', '错误收集error删除', '127.0.0.1', '/api/admin/monitor_web/del', 'x_admin/admin/monitor_web.MonitorWebHandler.Del-fm', '{\"id\":2}', '', 1, '2024-06-17 17:33:38', '2024-06-17 17:33:38', 44, '2024-06-17 17:33:38');
INSERT INTO `x_system_log_operate` VALUES (4, 1, 'POST', '错误收集error删除', '127.0.0.1', '/api/admin/monitor_web/del', 'x_admin/admin/monitor_web.MonitorWebHandler.Del-fm', '{\"id\":1}', '', 1, '2024-06-17 17:33:39', '2024-06-17 17:33:39', 8, '2024-06-17 17:33:39');
INSERT INTO `x_system_log_operate` VALUES (5, 1, 'POST', '错误收集error删除', '127.0.0.1', '/api/admin/monitor_web/del', 'x_admin/admin/monitor_web.MonitorWebHandler.Del-fm', '{\"id\":4}', '', 1, '2024-06-17 17:33:44', '2024-06-17 17:33:44', 3, '2024-06-17 17:33:44');
INSERT INTO `x_system_log_operate` VALUES (6, 1, 'POST', '错误收集error删除', '127.0.0.1', '/api/admin/monitor_web/del', 'x_admin/admin/monitor_web.MonitorWebHandler.Del-fm', '{\"id\":5}', '', 1, '2024-06-17 17:33:45', '2024-06-17 17:33:45', 3, '2024-06-17 17:33:45');
INSERT INTO `x_system_log_operate` VALUES (7, 1, 'POST', '错误收集error删除', '127.0.0.1', '/api/admin/monitor_web/del', 'x_admin/admin/monitor_web.MonitorWebHandler.Del-fm', '{\"id\":6}', '', 1, '2024-06-17 17:33:46', '2024-06-17 17:33:46', 3, '2024-06-17 17:33:46');
INSERT INTO `x_system_log_operate` VALUES (8, 1, 'POST', '错误收集error删除', '127.0.0.1', '/api/admin/monitor_web/del', 'x_admin/admin/monitor_web.MonitorWebHandler.Del-fm', '{\"id\":11}', '', 1, '2024-06-17 17:34:23', '2024-06-17 17:34:23', 4, '2024-06-17 17:34:23');
INSERT INTO `x_system_log_operate` VALUES (9, 1, 'POST', '错误收集error删除', '127.0.0.1', '/api/admin/monitor_web/del', 'x_admin/admin/monitor_web.MonitorWebHandler.Del-fm', '{\"id\":12}', '', 1, '2024-06-17 17:34:25', '2024-06-17 17:34:25', 3, '2024-06-17 17:34:25');
INSERT INTO `x_system_log_operate` VALUES (10, 1, 'POST', '错误收集error删除', '127.0.0.1', '/api/admin/monitor_web/del', 'x_admin/admin/monitor_web.MonitorWebHandler.Del-fm', '{\"id\":14}', '', 1, '2024-06-17 17:34:26', '2024-06-17 17:34:26', 49, '2024-06-17 17:34:26');
INSERT INTO `x_system_log_operate` VALUES (11, 1, 'POST', '错误收集error删除', '127.0.0.1', '/api/admin/monitor_web/del', 'x_admin/admin/monitor_web.MonitorWebHandler.Del-fm', '{\"id\":15}', '', 1, '2024-06-17 17:34:28', '2024-06-17 17:34:28', 4, '2024-06-17 17:34:28');
INSERT INTO `x_system_log_operate` VALUES (12, 1, 'POST', '错误收集error删除', '127.0.0.1', '/api/admin/monitor_web/del', 'x_admin/admin/monitor_web.MonitorWebHandler.Del-fm', '{\"id\":16}', '', 1, '2024-06-17 17:34:29', '2024-06-17 17:34:29', 4, '2024-06-17 17:34:29');
INSERT INTO `x_system_log_operate` VALUES (13, 1, 'POST', '错误收集error删除', '127.0.0.1', '/api/admin/monitor_web/del', 'x_admin/admin/monitor_web.MonitorWebHandler.Del-fm', '{\"id\":7}', '', 1, '2024-06-17 17:34:31', '2024-06-17 17:34:31', 2, '2024-06-17 17:34:31');
INSERT INTO `x_system_log_operate` VALUES (14, 1, 'POST', '错误收集error删除', '127.0.0.1', '/api/admin/monitor_web/del', 'x_admin/admin/monitor_web.MonitorWebHandler.Del-fm', '{\"id\":10}', '', 1, '2024-06-17 17:34:33', '2024-06-17 17:34:33', 2, '2024-06-17 17:34:33');
INSERT INTO `x_system_log_operate` VALUES (15, 1, 'POST', '错误收集error删除', '127.0.0.1', '/api/admin/monitor_web/del', 'x_admin/admin/monitor_web.MonitorWebHandler.Del-fm', '{\"id\":9}', '', 1, '2024-06-17 17:34:35', '2024-06-17 17:34:35', 5, '2024-06-17 17:34:35');
INSERT INTO `x_system_log_operate` VALUES (16, 1, 'POST', '错误收集error删除', '127.0.0.1', '/api/admin/monitor_web/del', 'x_admin/admin/monitor_web.MonitorWebHandler.Del-fm', '{\"id\":8}', '', 1, '2024-06-17 17:34:36', '2024-06-17 17:34:36', 8, '2024-06-17 17:34:36');
INSERT INTO `x_system_log_operate` VALUES (17, 1, 'POST', '错误项目编辑', '192.168.43.34', '/api/admin/monitor_project/edit', 'x_admin/admin/monitor_project.MonitorProjectHandler.Edit-fm', '{\"id\":null,\"projectKey\":\"1\",\"projectName\":\"1\",\"projectType\":\"web\"}', 'Error #01: 500:数据不存在!\n', 2, '2024-06-17 18:55:37', '2024-06-17 18:55:37', 148, '2024-06-17 18:55:37');
INSERT INTO `x_system_log_operate` VALUES (18, 1, 'POST', '错误项目新增', '192.168.43.34', '/api/admin/monitor_project/add', 'x_admin/admin/monitor_project.MonitorProjectHandler.Add-fm', '{\"id\":null,\"projectKey\":\"1\",\"projectName\":\"2\",\"projectType\":\"web\"}', '', 1, '2024-06-17 18:58:57', '2024-06-17 18:58:57', 6, '2024-06-17 18:58:57');
INSERT INTO `x_system_log_operate` VALUES (19, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-06-17 19:27:21', '2024-06-17 19:27:21', 10, '2024-06-17 19:27:21');
INSERT INTO `x_system_log_operate` VALUES (20, 1, 'GET', '角色详情', '127.0.0.1', '/api/admin/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=1', '', 1, '2024-06-17 19:27:24', '2024-06-17 19:27:24', 137, '2024-06-17 19:27:24');
INSERT INTO `x_system_log_operate` VALUES (21, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-06-17 19:27:35', '2024-06-17 19:27:35', 13, '2024-06-17 19:27:35');
INSERT INTO `x_system_log_operate` VALUES (22, 1, 'GET', '角色详情', '127.0.0.1', '/api/admin/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=1', '', 1, '2024-06-17 19:27:39', '2024-06-17 19:27:39', 13, '2024-06-17 19:27:39');
INSERT INTO `x_system_log_operate` VALUES (23, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-06-17 19:27:58', '2024-06-17 19:27:58', 3, '2024-06-17 19:27:58');
INSERT INTO `x_system_log_operate` VALUES (24, 1, 'GET', '角色详情', '127.0.0.1', '/api/admin/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=1', '', 1, '2024-06-17 19:28:00', '2024-06-17 19:28:00', 9, '2024-06-17 19:28:00');
INSERT INTO `x_system_log_operate` VALUES (25, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-06-17 19:28:21', '2024-06-17 19:28:21', 4, '2024-06-17 19:28:21');
INSERT INTO `x_system_log_operate` VALUES (26, 1, 'GET', '角色详情', '127.0.0.1', '/api/admin/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=1', '', 1, '2024-06-17 19:28:24', '2024-06-17 19:28:24', 13, '2024-06-17 19:28:24');
INSERT INTO `x_system_log_operate` VALUES (27, 1, 'GET', '角色详情', '127.0.0.1', '/api/admin/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=1', '', 1, '2024-06-17 19:28:29', '2024-06-17 19:28:29', 5, '2024-06-17 19:28:29');
INSERT INTO `x_system_log_operate` VALUES (28, 1, 'POST', '角色编辑', '127.0.0.1', '/api/admin/system/role/edit', 'x_admin/admin/system/role.RoleHandler.Edit-fm', '{\"id\":1,\"isDisable\":0,\"menuIds\":\"1,700,701,200,201,202,203,204,205,206,207,208,209,215,216,217,500,130,131,132,133,134,135,140,141,142,143,144,100,101,102,103,104,105,106,110,111,112,113,114,120,121,122,123,124,501,502,503,505,506,510,511,515,516,517,518,520,521,522,519,550,551,552,553,783,813,814,815,816,817,818,819,820,784,794,785,786,787,788,789,790,791,792,600,610,611,612,613,614,616,617,618,775,778,776,795,796,797,798,799,800,777,801,802,803,804,805,806,807,808,809,810,780,811,812,781,782\",\"menus\":[1,700,701,200,201,202,203,204,205,206,207,208,209,215,216,217,500,130,131,132,133,134,135,140,141,142,143,144,100,101,102,103,104,105,106,110,111,112,113,114,120,121,122,123,124,501,502,503,505,506,510,511,515,516,517,518,520,521,522,519,550,551,552,553,783,813,814,815,816,817,818,819,820,784,794,785,786,787,788,789,790,791,792,600,610,611,612,613,614,616,617,618,775,778,776,795,796,797,798,799,800,777,801,802,803,804,805,806,807,808,809,810,780,811,812,781,782],\"name\":\"审核员\",\"remark\":\"审核数据\",\"sort\":1}', '', 1, '2024-06-17 19:29:03', '2024-06-17 19:29:03', 211, '2024-06-17 19:29:03');
INSERT INTO `x_system_log_operate` VALUES (29, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-06-17 19:29:04', '2024-06-17 19:29:04', 8, '2024-06-17 19:29:04');
INSERT INTO `x_system_log_operate` VALUES (30, 1, 'POST', '错误项目编辑', '192.168.43.34', '/api/admin/monitor_project/edit', 'x_admin/admin/monitor_project.MonitorProjectHandler.Edit-fm', '{\"createTime\":\"2024-06-17 18:58:57\",\"id\":2,\"projectKey\":\"9a1d5d657d884bbfb7df0ddf032a8504\",\"projectName\":\"2\",\"projectType\":\"go\",\"updateTime\":\"2024-06-17 18:58:57\"}', '', 1, '2024-06-17 19:59:29', '2024-06-17 19:59:29', 34, '2024-06-17 19:59:29');
INSERT INTO `x_system_log_operate` VALUES (31, 1, 'POST', '错误项目编辑', '192.168.43.34', '/api/admin/monitor_project/edit', 'x_admin/admin/monitor_project.MonitorProjectHandler.Edit-fm', '{\"createTime\":\"2024-06-17 18:58:57\",\"id\":2,\"projectKey\":\"9a1d5d657d884bbfb7df0ddf032a8504\",\"projectName\":\"2\",\"projectType\":\"go\",\"updateTime\":\"2024-06-17 19:59:30\"}', '', 1, '2024-06-17 20:25:15', '2024-06-17 20:25:15', 4, '2024-06-17 20:25:15');
INSERT INTO `x_system_log_operate` VALUES (32, 1, 'POST', '错误项目编辑', '192.168.43.34', '/api/admin/monitor_project/edit', 'x_admin/admin/monitor_project.MonitorProjectHandler.Edit-fm', '{\"createTime\":\"2024-06-17 18:58:57\",\"id\":2,\"projectKey\":\"9a1d5d657d884bbfb7df0ddf032a8504\",\"projectName\":\"2\",\"projectType\":\"go\",\"updateTime\":\"2024-06-17 20:25:16\"}', '', 1, '2024-06-17 20:25:17', '2024-06-17 20:25:17', 4, '2024-06-17 20:25:17');
INSERT INTO `x_system_log_operate` VALUES (33, 1, 'POST', '管理员新增', '127.0.0.1', '/api/admin/system/admin/add', 'x_admin/admin/system/admin.AdminHandler.Add-fm', '{\"avatar\":\"/api/uploads/image/20241705/f0eb36d508834bc2ac1b8c591c563efa.png\",\"deptId\":1,\"id\":\"\",\"isDisable\":0,\"nickname\":\"a2\",\"password\":\"da00d628804752a870ca30e83c97f2ba\",\"passwordConfirm\":\"da00d628804752a870ca30e83c97f2ba\",\"postId\":2,\"role\":1,\"sort\":1,\"username\":\"a2\"}', 'Error #01: 300:密码必须在6~20位\n', 2, '2024-06-17 20:26:46', '2024-06-17 20:26:46', 11, '2024-06-17 20:26:46');
INSERT INTO `x_system_log_operate` VALUES (34, 1, 'POST', '管理员新增', '127.0.0.1', '/api/admin/system/admin/add', 'x_admin/admin/system/admin.AdminHandler.Add-fm', '{\"avatar\":\"/api/uploads/image/20241705/f0eb36d508834bc2ac1b8c591c563efa.png\",\"deptId\":1,\"id\":\"\",\"isDisable\":0,\"nickname\":\"a2\",\"password\":\"da00d628804752a870ca30e83c97f2ba\",\"passwordConfirm\":\"da00d628804752a870ca30e83c97f2ba\",\"postId\":2,\"role\":1,\"sort\":1,\"username\":\"a2\"}', 'Error #01: 300:密码必须在6~20位\n', 2, '2024-06-17 20:26:57', '2024-06-17 20:26:57', 8, '2024-06-17 20:26:57');
INSERT INTO `x_system_log_operate` VALUES (35, 1, 'POST', '管理员新增', '127.0.0.1', '/api/admin/system/admin/add', 'x_admin/admin/system/admin.AdminHandler.Add-fm', '{\"avatar\":\"/api/uploads/image/20241705/f0eb36d508834bc2ac1b8c591c563efa.png\",\"deptId\":1,\"id\":\"\",\"isDisable\":0,\"nickname\":\"a2\",\"password\":\"da00d628804752a870ca30e83c97f2ba\",\"passwordConfirm\":\"da00d628804752a870ca30e83c97f2ba\",\"postId\":2,\"role\":1,\"sort\":1,\"username\":\"a2\"}', 'Error #01: 300:密码必须在6~20位\n', 2, '2024-06-17 20:27:02', '2024-06-17 20:27:02', 4, '2024-06-17 20:27:02');
INSERT INTO `x_system_log_operate` VALUES (36, 1, 'POST', '管理员新增', '127.0.0.1', '/api/admin/system/admin/add', 'x_admin/admin/system/admin.AdminHandler.Add-fm', '{\"avatar\":\"/api/uploads/image/20241705/f0eb36d508834bc2ac1b8c591c563efa.png\",\"deptId\":1,\"id\":\"\",\"isDisable\":0,\"nickname\":\"a2\",\"password\":\"da00d628804752a870ca30e83c97f2ba\",\"passwordConfirm\":\"da00d628804752a870ca30e83c97f2ba\",\"postId\":2,\"role\":1,\"sort\":1,\"username\":\"a2\"}', '', 1, '2024-06-17 20:29:54', '2024-06-17 20:29:54', 19, '2024-06-17 20:29:54');
INSERT INTO `x_system_log_operate` VALUES (37, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-06-17 20:52:23', '2024-06-17 20:52:23', 3, '2024-06-17 20:52:23');
INSERT INTO `x_system_log_operate` VALUES (38, 1, 'GET', '角色详情', '127.0.0.1', '/api/admin/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=1', '', 1, '2024-06-17 20:52:24', '2024-06-17 20:52:24', 9, '2024-06-17 20:52:24');
INSERT INTO `x_system_log_operate` VALUES (39, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-06-17 20:52:48', '2024-06-17 20:52:48', 10, '2024-06-17 20:52:48');
INSERT INTO `x_system_log_operate` VALUES (40, 1, 'GET', '角色详情', '127.0.0.1', '/api/admin/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=1', '', 1, '2024-06-17 20:52:50', '2024-06-17 20:52:50', 10, '2024-06-17 20:52:50');
INSERT INTO `x_system_log_operate` VALUES (41, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-06-17 21:05:23', '2024-06-17 21:05:23', 3, '2024-06-17 21:05:23');
INSERT INTO `x_system_log_operate` VALUES (42, 1, 'GET', '角色详情', '127.0.0.1', '/api/admin/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=1', '', 1, '2024-06-17 21:05:25', '2024-06-17 21:05:25', 4, '2024-06-17 21:05:25');
INSERT INTO `x_system_log_operate` VALUES (43, 1, 'POST', '角色编辑', '127.0.0.1', '/api/admin/system/role/edit', 'x_admin/admin/system/role.RoleHandler.Edit-fm', '{\"id\":1,\"isDisable\":0,\"menuIds\":\"1,700,701,200,201,202,203,204,205,206,207,208,209,215,216,217,500,130,131,132,133,134,135,140,141,142,143,144,100,101,102,103,104,105,106,110,111,112,113,114,120,121,122,123,124,501,502,503,505,506,510,511,515,516,517,518,520,521,522,519,550,551,552,553,783,813,814,815,816,817,818,819,820,784,794,785,786,787,788,789,790,791,792,600,610,611,612,613,614,616,617,618,775,778,776,795,796,797,798,799,800,777,801,802,803,804,805,806,807,808,809,810,780,811,812,781,782\",\"menus\":[1,700,701,200,201,202,203,204,205,206,207,208,209,215,216,217,500,130,131,132,133,134,135,140,141,142,143,144,100,101,102,103,104,105,106,110,111,112,113,114,120,121,122,123,124,501,502,503,505,506,510,511,515,516,517,518,520,521,522,519,550,551,552,553,783,813,814,815,816,817,818,819,820,784,794,785,786,787,788,789,790,791,792,600,610,611,612,613,614,616,617,618,775,778,776,795,796,797,798,799,800,777,801,802,803,804,805,806,807,808,809,810,780,811,812,781,782],\"name\":\"审核员\",\"remark\":\"审核数据\",\"sort\":1}', '', 1, '2024-06-17 21:05:30', '2024-06-17 21:05:31', 216, '2024-06-17 21:05:31');
INSERT INTO `x_system_log_operate` VALUES (44, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-06-17 21:05:31', '2024-06-17 21:05:31', 5, '2024-06-17 21:05:31');
INSERT INTO `x_system_log_operate` VALUES (45, 1, 'GET', '角色详情', '127.0.0.1', '/api/admin/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=1', '', 1, '2024-06-17 21:05:48', '2024-06-17 21:05:48', 10, '2024-06-17 21:05:48');
INSERT INTO `x_system_log_operate` VALUES (46, 1, 'POST', '角色编辑', '127.0.0.1', '/api/admin/system/role/edit', 'x_admin/admin/system/role.RoleHandler.Edit-fm', '{\"id\":1,\"isDisable\":0,\"menuIds\":\"1,700,701,200,201,202,203,204,205,206,207,208,209,215,216,217,500,130,131,132,133,134,135,140,141,142,143,144,100,101,102,103,104,105,106,110,111,112,113,114,120,121,122,123,124,501,502,503,505,506,510,511,515,516,517,518,520,521,522,519,550,551,552,553,783,813,814,815,816,817,818,819,820,784,794,785,786,787,788,789,790,791,792,600,610,611,612,613,614,616,617,618,775,778,776,795,796,797,798,799,800,777,801,802,803,804,805,806,807,808,809,810,780,811,812,781,782\",\"menus\":[1,700,701,200,201,202,203,204,205,206,207,208,209,215,216,217,500,130,131,132,133,134,135,140,141,142,143,144,100,101,102,103,104,105,106,110,111,112,113,114,120,121,122,123,124,501,502,503,505,506,510,511,515,516,517,518,520,521,522,519,550,551,552,553,783,813,814,815,816,817,818,819,820,784,794,785,786,787,788,789,790,791,792,600,610,611,612,613,614,616,617,618,775,778,776,795,796,797,798,799,800,777,801,802,803,804,805,806,807,808,809,810,780,811,812,781,782],\"name\":\"审核员\",\"remark\":\"审核数据\",\"sort\":1}', '', 1, '2024-06-17 21:05:49', '2024-06-17 21:05:49', 23, '2024-06-17 21:05:49');
INSERT INTO `x_system_log_operate` VALUES (47, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-06-17 21:05:50', '2024-06-17 21:05:50', 2, '2024-06-17 21:05:50');
INSERT INTO `x_system_log_operate` VALUES (48, 2, 'POST', '错误项目编辑', '192.168.43.34', '/api/admin/monitor_project/edit', 'x_admin/admin/monitor_project.MonitorProjectHandler.Edit-fm', '{\"createTime\":\"2024-06-17 18:58:57\",\"id\":2,\"projectKey\":\"9a1d5d657d884bbfb7df0ddf032a8504\",\"projectName\":\"2\",\"projectType\":\"go\",\"updateTime\":\"2024-06-17 20:25:18\"}', '', 1, '2024-06-17 21:10:48', '2024-06-17 21:10:48', 3, '2024-06-17 21:10:48');
INSERT INTO `x_system_log_operate` VALUES (49, 1, 'GET', '角色详情', '127.0.0.1', '/api/admin/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=1', '', 1, '2024-06-17 21:12:43', '2024-06-17 21:12:43', 8, '2024-06-17 21:12:43');
INSERT INTO `x_system_log_operate` VALUES (50, 1, 'POST', '客户端信息新增', '127.0.0.1', '/api/admin/monitor_client/add', 'x_admin/admin/monitor_client.MonitorClientHandler.Add-fm', '{\"browser\":\"1\",\"city\":\"1\",\"clientId\":\"1\",\"clientTime\":\"2024-06-19 12:00:00\",\"height\":\"1\",\"id\":\"\",\"os\":\"1\",\"ua\":\"1\",\"userId\":\"1\",\"width\":\"1\"}', 'Error #01: 310:参数校验错误\n', 2, '2024-06-17 22:27:53', '2024-06-17 22:27:54', 91, '2024-06-17 22:27:54');
INSERT INTO `x_system_log_operate` VALUES (51, 1, 'POST', '客户端信息新增', '127.0.0.1', '/api/admin/monitor_client/add', 'x_admin/admin/monitor_client.MonitorClientHandler.Add-fm', '{\"browser\":\"1\",\"city\":\"1\",\"clientId\":\"1\",\"clientTime\":\"2024-06-19 12:00:00\",\"height\":\"1\",\"id\":\"\",\"os\":\"1\",\"ua\":\"1\",\"userId\":\"1\",\"width\":\"1\"}', 'Error #01: 310:参数校验错误\n', 2, '2024-06-17 22:28:11', '2024-06-17 22:28:11', 1, '2024-06-17 22:28:11');
INSERT INTO `x_system_log_operate` VALUES (52, 1, 'GET', '服务监控', '127.0.0.1', '/api/admin/monitor/server', 'x_admin/admin/monitor.monitorHandler.server-fm', '', '', 1, '2024-06-28 17:07:36', '2024-06-28 17:07:36', 59, '2024-06-28 17:07:36');
INSERT INTO `x_system_log_operate` VALUES (53, 1, 'GET', '缓存监控', '127.0.0.1', '/api/admin/monitor/cache', 'x_admin/admin/monitor.monitorHandler.cache-fm', '', '', 1, '2024-06-28 17:07:37', '2024-06-28 17:07:37', 88, '2024-06-28 17:07:37');
INSERT INTO `x_system_log_operate` VALUES (54, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-06-28 17:08:11', '2024-06-28 17:08:11', 17, '2024-06-28 17:08:11');
INSERT INTO `x_system_log_operate` VALUES (55, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-06-28 17:08:15', '2024-06-28 17:08:15', 2, '2024-06-28 17:08:15');
INSERT INTO `x_system_log_operate` VALUES (56, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-06-28 17:09:27', '2024-06-28 17:09:27', 2, '2024-06-28 17:09:27');
INSERT INTO `x_system_log_operate` VALUES (57, 1, 'GET', '角色详情', '127.0.0.1', '/api/admin/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=1', '', 1, '2024-06-28 17:09:29', '2024-06-28 17:09:29', 13, '2024-06-28 17:09:29');
INSERT INTO `x_system_log_operate` VALUES (58, 1, 'GET', '角色详情', '127.0.0.1', '/api/admin/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=1', '', 1, '2024-06-28 17:09:32', '2024-06-28 17:09:32', 5, '2024-06-28 17:09:32');
INSERT INTO `x_system_log_operate` VALUES (59, 1, 'POST', '角色编辑', '127.0.0.1', '/api/admin/system/role/edit', 'x_admin/admin/system/role.RoleHandler.Edit-fm', '{\"id\":1,\"isDisable\":0,\"menuIds\":\"1,700,701,200,201,202,203,204,205,206,207,208,209,215,216,217,500,130,131,132,133,134,135,140,141,142,143,144,100,101,102,103,104,105,106,110,111,112,113,114,120,121,122,123,124,501,502,503,505,506,510,511,515,516,517,518,520,521,522,519,550,551,552,553,783,813,814,815,816,817,818,819,820,784,794,785,786,787,788,789,790,791,792,600,610,611,612,613,614,616,617,618,775,778,776,795,796,797,798,799,800,777,801,802,803,804,805,806,807,808,809,810,780,811,812,781,782,821\",\"menus\":[1,700,701,200,201,202,203,204,205,206,207,208,209,215,216,217,500,130,131,132,133,134,135,140,141,142,143,144,100,101,102,103,104,105,106,110,111,112,113,114,120,121,122,123,124,501,502,503,505,506,510,511,515,516,517,518,520,521,522,519,550,551,552,553,783,813,814,815,816,817,818,819,820,784,794,785,786,787,788,789,790,791,792,600,610,611,612,613,614,616,617,618,775,778,776,795,796,797,798,799,800,777,801,802,803,804,805,806,807,808,809,810,780,811,812,781,782,821],\"name\":\"审核员\",\"remark\":\"审核数据\",\"sort\":1}', '', 1, '2024-06-28 17:09:35', '2024-06-28 17:09:35', 65, '2024-06-28 17:09:35');
INSERT INTO `x_system_log_operate` VALUES (60, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-06-28 17:09:36', '2024-06-28 17:09:36', 8, '2024-06-28 17:09:36');
INSERT INTO `x_system_log_operate` VALUES (61, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-06-28 17:09:39', '2024-06-28 17:09:39', 3, '2024-06-28 17:09:39');
INSERT INTO `x_system_log_operate` VALUES (62, 1, 'POST', '错误项目编辑', '127.0.0.1', '/api/admin/monitor_project/edit', 'x_admin/admin/monitor_project.MonitorProjectHandler.Edit-fm', '{\"id\":2,\"projectName\":\"2\",\"projectType\":\"go\"}', '', 1, '2024-06-29 00:18:52', '2024-06-29 00:18:52', 4, '2024-06-29 00:18:52');
INSERT INTO `x_system_log_operate` VALUES (63, 1, 'POST', '管理员编辑', '127.0.0.1', '/api/admin/system/admin/edit', 'x_admin/admin/system/admin.AdminHandler.Edit-fm', '{\"avatar\":\"/api/uploads/image/20241705/f0eb36d508834bc2ac1b8c591c563efa.png\",\"deptId\":1,\"id\":2,\"isDisable\":0,\"nickname\":\"a2\",\"password\":\"\",\"passwordConfirm\":\"\",\"postId\":2,\"role\":1,\"sort\":1,\"username\":\"a2\"}', '', 1, '2024-06-29 00:30:28', '2024-06-29 00:30:28', 17, '2024-06-29 00:30:28');
INSERT INTO `x_system_log_operate` VALUES (64, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-06-29 00:30:30', '2024-06-29 00:30:30', 2, '2024-06-29 00:30:30');
INSERT INTO `x_system_log_operate` VALUES (65, 1, 'GET', '角色详情', '127.0.0.1', '/api/admin/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=1', '', 1, '2024-06-29 00:30:32', '2024-06-29 00:30:32', 4, '2024-06-29 00:30:32');
INSERT INTO `x_system_log_operate` VALUES (66, 1, 'POST', '角色编辑', '127.0.0.1', '/api/admin/system/role/edit', 'x_admin/admin/system/role.RoleHandler.Edit-fm', '{\"id\":1,\"isDisable\":0,\"menuIds\":\"506,785,105,206,133,819,809,795,515,205,522,799,776,798,101,501,111,130,216,600,503,131,551,204,775,202,124,818,519,612,511,611,617,505,201,217,802,805,553,200,803,784,135,800,520,120,814,804,777,817,821,521,810,516,790,778,510,613,502,792,550,203,518,121,616,207,132,113,143,811,103,134,781,552,820,140,610,122,112,500,517,789,813,142,806,614,208,815,782,110,701,100,808,104,812,788,102,141,215,794,797,123,807,106,796,783,816,618,144,791,801,114,780,786,209,1,700,787\",\"menus\":[506,785,105,206,133,819,809,795,515,205,522,799,776,798,101,501,111,130,216,600,503,131,551,204,775,202,124,818,519,612,511,611,617,505,201,217,802,805,553,200,803,784,135,800,520,120,814,804,777,817,821,521,810,516,790,778,510,613,502,792,550,203,518,121,616,207,132,113,143,811,103,134,781,552,820,140,610,122,112,500,517,789,813,142,806,614,208,815,782,110,701,100,808,104,812,788,102,141,215,794,797,123,807,106,796,783,816,618,144,791,801,114,780,786,209,1,700,787],\"name\":\"审核员\",\"remark\":\"审核数据\",\"sort\":1}', '', 1, '2024-06-29 00:30:33', '2024-06-29 00:30:33', 34, '2024-06-29 00:30:33');
INSERT INTO `x_system_log_operate` VALUES (67, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-06-29 00:30:33', '2024-06-29 00:30:33', 5, '2024-06-29 00:30:33');
INSERT INTO `x_system_log_operate` VALUES (68, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-06-29 00:30:36', '2024-06-29 00:30:36', 1, '2024-06-29 00:30:36');
INSERT INTO `x_system_log_operate` VALUES (69, 1, 'GET', '服务监控', '127.0.0.1', '/api/admin/monitor/server', 'x_admin/admin/monitor.monitorHandler.server-fm', '', '', 1, '2024-06-29 00:31:15', '2024-06-29 00:31:15', 10, '2024-06-29 00:31:15');
INSERT INTO `x_system_log_operate` VALUES (70, 1, 'GET', '缓存监控', '127.0.0.1', '/api/admin/monitor/cache', 'x_admin/admin/monitor.monitorHandler.cache-fm', '', '', 1, '2024-06-29 00:31:17', '2024-06-29 00:31:17', 1, '2024-06-29 00:31:17');
INSERT INTO `x_system_log_operate` VALUES (71, 1, 'POST', '错误项目删除', '127.0.0.1', '/api/admin/monitor_project/del', 'x_admin/admin/monitor_project.MonitorProjectHandler.Del-fm', '{\"id\":2}', '', 1, '2024-06-29 00:31:27', '2024-06-29 00:31:27', 3, '2024-06-29 00:31:27');
INSERT INTO `x_system_log_operate` VALUES (72, 1, 'POST', '错误收集error删除', '127.0.0.1', '/api/admin/monitor_web/del', 'x_admin/admin/monitor_web.MonitorWebHandler.Del-fm', '{\"id\":17}', '', 1, '2024-06-29 00:42:57', '2024-06-29 00:42:57', 4, '2024-06-29 00:42:57');
INSERT INTO `x_system_log_operate` VALUES (73, 1, 'POST', '错误收集error删除', '127.0.0.1', '/api/admin/monitor_web/del', 'x_admin/admin/monitor_web.MonitorWebHandler.Del-fm', '{\"id\":18}', '', 1, '2024-06-29 00:42:59', '2024-06-29 00:42:59', 3, '2024-06-29 00:42:59');
INSERT INTO `x_system_log_operate` VALUES (74, 1, 'POST', '错误收集error新增', '127.0.0.1', '/api/admin/monitor_web/add', 'x_admin/admin/monitor_web.MonitorWebHandler.Add-fm', '{\"clientId\":\"1\",\"clientTime\":\"1\",\"eventType\":\"1\",\"id\":\"\",\"message\":\"11\",\"page\":\"1\",\"projectKey\":\"1\",\"stack\":\"1\"}', '', 1, '2024-06-29 00:43:08', '2024-06-29 00:43:08', 3, '2024-06-29 00:43:08');
INSERT INTO `x_system_log_operate` VALUES (75, 1, 'GET', '错误收集error导出', '127.0.0.1', '/api/admin/monitor_web/ExportFile', 'x_admin/admin/monitor_web.MonitorWebHandler.ExportFile-fm', 'token=373bbfc6547feffd04882380b95d503etdx1R3&clientTimeEnd=&clientTimeStart=&createTimeEnd=&createTimeStart=&eventType=&message=&page=&stack=', '', 1, '2024-06-29 00:44:34', '2024-06-29 00:44:35', 224, '2024-06-29 00:44:35');
INSERT INTO `x_system_log_operate` VALUES (76, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-01 16:43:30', '2024-07-01 16:43:30', 2, '2024-07-01 16:43:30');
INSERT INTO `x_system_log_operate` VALUES (77, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-05 23:04:44', '2024-07-05 23:04:44', 4, '2024-07-05 23:04:44');
INSERT INTO `x_system_log_operate` VALUES (78, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-05 23:05:52', '2024-07-05 23:05:52', 2, '2024-07-05 23:05:52');
INSERT INTO `x_system_log_operate` VALUES (79, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-08 15:21:01', '2024-07-08 15:21:01', 3, '2024-07-08 15:21:01');
INSERT INTO `x_system_log_operate` VALUES (80, 1, 'POST', '错误项目编辑', '127.0.0.1', '/api/admin/monitor_project/edit', 'x_admin/admin/monitor_project.MonitorProjectHandler.Edit-fm', '{\"id\":1,\"projectName\":\"x_admin\",\"projectType\":\"web\"}', '', 1, '2024-07-08 15:21:41', '2024-07-08 15:21:41', 13, '2024-07-08 15:21:41');
INSERT INTO `x_system_log_operate` VALUES (81, 1, 'POST', '错误项目新增', '127.0.0.1', '/api/admin/monitor_project/add', 'x_admin/admin/monitor_project.MonitorProjectHandler.Add-fm', '{\"id\":\"\",\"projectName\":\"ad\",\"projectType\":\"web\"}', '', 1, '2024-07-08 15:26:56', '2024-07-08 15:26:56', 15, '2024-07-08 15:26:56');
INSERT INTO `x_system_log_operate` VALUES (82, 1, 'POST', '错误项目编辑', '127.0.0.1', '/api/admin/monitor_project/edit', 'x_admin/admin/monitor_project.MonitorProjectHandler.Edit-fm', '{\"id\":3,\"projectName\":\"ad\",\"projectType\":\"go\"}', '', 1, '2024-07-09 00:44:50', '2024-07-09 00:44:50', 32, '2024-07-09 00:44:50');
INSERT INTO `x_system_log_operate` VALUES (83, 1, 'POST', '错误项目新增', '127.0.0.1', '/api/admin/monitor_project/add', 'x_admin/admin/monitor_project.MonitorProjectHandler.Add-fm', '{\"id\":\"\",\"projectName\":\"a\",\"projectType\":\"uniapp\"}', '', 1, '2024-07-09 00:45:18', '2024-07-09 00:45:18', 2, '2024-07-09 00:45:18');
INSERT INTO `x_system_log_operate` VALUES (84, 1, 'POST', '错误项目删除', '127.0.0.1', '/api/admin/monitor_project/del', 'x_admin/admin/monitor_project.MonitorProjectHandler.Del-fm', '{\"id\":4}', '', 1, '2024-07-09 00:49:26', '2024-07-09 00:49:26', 11, '2024-07-09 00:49:26');
INSERT INTO `x_system_log_operate` VALUES (85, 1, 'POST', '错误项目编辑', '127.0.0.1', '/api/admin/monitor_project/edit', 'x_admin/admin/monitor_project.MonitorProjectHandler.Edit-fm', '{\"id\":1,\"projectName\":\"x_admin\",\"projectType\":\"node\"}', '', 1, '2024-07-09 12:06:41', '2024-07-09 12:06:41', 12, '2024-07-09 12:06:41');
INSERT INTO `x_system_log_operate` VALUES (86, 1, 'POST', '错误项目编辑', '127.0.0.1', '/api/admin/monitor_project/edit', 'x_admin/admin/monitor_project.(*MonitorProjectHandler).Edit-fm', '{\"id\":3,\"projectName\":\"ad\",\"projectType\":\"web\"}', '', 1, '2024-07-09 15:04:47', '2024-07-09 15:04:47', 13, '2024-07-09 15:04:47');
INSERT INTO `x_system_log_operate` VALUES (87, 1, 'GET', '服务监控', '127.0.0.1', '/api/admin/monitor/server', 'x_admin/admin/monitor.monitorHandler.server-fm', '', '', 1, '2024-07-09 15:26:38', '2024-07-09 15:26:38', 9, '2024-07-09 15:26:38');
INSERT INTO `x_system_log_operate` VALUES (88, 1, 'GET', '缓存监控', '127.0.0.1', '/api/admin/monitor/cache', 'x_admin/admin/monitor.monitorHandler.cache-fm', '', '', 1, '2024-07-09 15:26:42', '2024-07-09 15:26:42', 1, '2024-07-09 15:26:42');
INSERT INTO `x_system_log_operate` VALUES (89, 1, 'GET', '服务监控', '127.0.0.1', '/api/admin/monitor/server', 'x_admin/admin/monitor.monitorHandler.server-fm', '', '', 1, '2024-07-09 15:26:43', '2024-07-09 15:26:43', 12, '2024-07-09 15:26:43');
INSERT INTO `x_system_log_operate` VALUES (90, 1, 'GET', '缓存监控', '127.0.0.1', '/api/admin/monitor/cache', 'x_admin/admin/monitor.monitorHandler.cache-fm', '', '', 1, '2024-07-09 15:26:47', '2024-07-09 15:26:47', 1, '2024-07-09 15:26:47');
INSERT INTO `x_system_log_operate` VALUES (91, 1, 'POST', '错误项目编辑', '127.0.0.1', '/api/admin/monitor_project/edit', 'x_admin/admin/monitor_project.(*MonitorProjectHandler).Edit-fm', '{\"id\":3,\"projectName\":\"ad\",\"projectType\":\"uniapp\"}', '', 1, '2024-07-09 15:34:46', '2024-07-09 15:34:46', 17, '2024-07-09 15:34:46');
INSERT INTO `x_system_log_operate` VALUES (92, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-10 00:47:56', '2024-07-10 00:47:56', 5, '2024-07-10 00:47:56');
INSERT INTO `x_system_log_operate` VALUES (93, 1, 'POST', '错误项目编辑', '127.0.0.1', '/api/admin/monitor_project/edit', 'x_admin/admin/monitor_project.(*MonitorProjectHandler).Edit-fm', '{\"id\":3,\"projectName\":\"ad\",\"projectType\":\"web\"}', '', 1, '2024-07-10 12:20:19', '2024-07-10 12:20:19', 4, '2024-07-10 12:20:19');
INSERT INTO `x_system_log_operate` VALUES (94, 1, 'POST', '错误项目新增', '127.0.0.1', '/api/admin/monitor_project/add', 'x_admin/admin/monitor_project.(*MonitorProjectHandler).Add-fm', '{\"id\":\"\",\"projectName\":\"无\",\"projectType\":\"web\"}', '', 1, '2024-07-10 12:20:37', '2024-07-10 12:20:37', 5, '2024-07-10 12:20:37');
INSERT INTO `x_system_log_operate` VALUES (95, 1, 'POST', '错误项目删除', '127.0.0.1', '/api/admin/monitor_project/del', 'x_admin/admin/monitor_project.(*MonitorProjectHandler).Del-fm', '{\"id\":3}', '', 1, '2024-07-10 12:21:11', '2024-07-10 12:21:11', 3, '2024-07-10 12:21:11');
INSERT INTO `x_system_log_operate` VALUES (96, 1, 'POST', '错误项目编辑', '127.0.0.1', '/api/admin/monitor_project/edit', 'x_admin/admin/monitor_project.(*MonitorProjectHandler).Edit-fm', '{\"id\":5,\"projectName\":\"无\",\"projectType\":\"web\"}', '', 1, '2024-07-10 12:22:44', '2024-07-10 12:22:44', 6, '2024-07-10 12:22:44');
INSERT INTO `x_system_log_operate` VALUES (97, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-12 19:30:31', '2024-07-12 19:30:31', 1, '2024-07-12 19:30:31');
INSERT INTO `x_system_log_operate` VALUES (98, 1, 'POST', '错误项目编辑', '127.0.0.1', '/api/admin/monitor_project/edit', 'x_admin/admin/monitor_project.(*MonitorProjectHandler).Edit-fm', '{\"id\":5,\"projectName\":\"\",\"projectType\":\"web\"}', '', 1, '2024-07-12 19:35:49', '2024-07-12 19:35:49', 38, '2024-07-12 19:35:49');
INSERT INTO `x_system_log_operate` VALUES (99, 1, 'POST', '错误项目编辑', '127.0.0.1', '/api/admin/monitor_project/edit', 'x_admin/admin/monitor_project.(*MonitorProjectHandler).Edit-fm', '{\"id\":5,\"projectName\":\"\",\"projectType\":\"web\"}', '', 1, '2024-07-12 19:37:47', '2024-07-12 19:55:10', 1043116, '2024-07-12 19:55:10');
INSERT INTO `x_system_log_operate` VALUES (100, 1, 'POST', '错误项目编辑', '127.0.0.1', '/api/admin/monitor_project/edit', 'x_admin/admin/monitor_project.(*MonitorProjectHandler).Edit-fm', '{\"id\":5,\"projectName\":\"\",\"projectType\":\"web\"}', '', 1, '2024-07-12 20:09:18', '2024-07-12 20:09:20', 2684, '2024-07-12 20:09:20');
INSERT INTO `x_system_log_operate` VALUES (101, 1, 'POST', '错误项目编辑', '127.0.0.1', '/api/admin/monitor_project/edit', 'x_admin/admin/monitor_project.(*MonitorProjectHandler).Edit-fm', '{\"id\":5,\"projectName\":\"123\",\"projectType\":\"web\"}', '', 1, '2024-07-12 20:09:50', '2024-07-12 20:09:58', 7841, '2024-07-12 20:09:58');
INSERT INTO `x_system_log_operate` VALUES (102, 1, 'POST', '错误项目编辑', '127.0.0.1', '/api/admin/monitor_project/edit', 'x_admin/admin/monitor_project.(*MonitorProjectHandler).Edit-fm', '{\"id\":5,\"projectName\":\"\",\"projectType\":\"web\"}', '', 1, '2024-07-12 20:10:11', '2024-07-12 20:10:34', 23515, '2024-07-12 20:10:34');
INSERT INTO `x_system_log_operate` VALUES (103, 1, 'POST', '错误项目编辑', '127.0.0.1', '/api/admin/monitor_project/edit', 'x_admin/admin/monitor_project.(*MonitorProjectHandler).Edit-fm', '{\"id\":5,\"projectName\":\"2\",\"projectType\":\"web\"}', '', 1, '2024-07-12 20:12:38', '2024-07-12 20:12:41', 2666, '2024-07-12 20:12:41');
INSERT INTO `x_system_log_operate` VALUES (104, 1, 'POST', '错误项目编辑', '127.0.0.1', '/api/admin/monitor_project/edit', 'x_admin/admin/monitor_project.(*MonitorProjectHandler).Edit-fm', '{\"id\":5,\"projectName\":\"2\",\"projectType\":\"web\"}', '', 1, '2024-07-12 20:12:49', '2024-07-12 20:13:07', 17698, '2024-07-12 20:13:07');
INSERT INTO `x_system_log_operate` VALUES (105, 1, 'POST', '错误项目编辑', '127.0.0.1', '/api/admin/monitor_project/edit', 'x_admin/admin/monitor_project.(*MonitorProjectHandler).Edit-fm', '{\"id\":5,\"projectName\":\"\",\"projectType\":\"web\"}', '', 1, '2024-07-12 20:13:13', '2024-07-12 20:18:46', 332888, '2024-07-12 20:18:46');
INSERT INTO `x_system_log_operate` VALUES (106, 1, 'POST', '错误项目编辑', '127.0.0.1', '/api/admin/monitor_project/edit', 'x_admin/admin/monitor_project.(*MonitorProjectHandler).Edit-fm', '{\"id\":5,\"projectType\":\"web\"}', '', 1, '2024-07-12 20:18:50', '2024-07-12 20:20:37', 107466, '2024-07-12 20:20:37');
INSERT INTO `x_system_log_operate` VALUES (107, 1, 'POST', '错误项目编辑', '127.0.0.1', '/api/admin/monitor_project/edit', 'x_admin/admin/monitor_project.(*MonitorProjectHandler).Edit-fm', '{\"id\":5,\"projectName\":\"\",\"projectType\":\"web\"}', '', 1, '2024-07-12 20:21:06', '2024-07-12 20:21:13', 7055, '2024-07-12 20:21:13');
INSERT INTO `x_system_log_operate` VALUES (108, 1, 'POST', '错误项目编辑', '127.0.0.1', '/api/admin/monitor_project/edit', 'x_admin/admin/monitor_project.(*MonitorProjectHandler).Edit-fm', '{\"id\":5,\"projectName\":\"\",\"projectType\":\"web\"}', '', 1, '2024-07-12 20:29:14', '2024-07-12 20:29:36', 21934, '2024-07-12 20:29:36');
INSERT INTO `x_system_log_operate` VALUES (109, 1, 'POST', '错误项目编辑', '127.0.0.1', '/api/admin/monitor_project/edit', 'x_admin/admin/monitor_project.(*MonitorProjectHandler).Edit-fm', '{\"id\":5,\"projectType\":\"web\"}', '', 1, '2024-07-12 20:29:58', '2024-07-12 20:33:18', 200419, '2024-07-12 20:33:18');
INSERT INTO `x_system_log_operate` VALUES (110, 1, 'POST', '错误项目编辑', '127.0.0.1', '/api/admin/monitor_project/edit', 'x_admin/admin/monitor_project.(*MonitorProjectHandler).Edit-fm', '{\"id\":5,\"projectType\":\"uniapp\"}', '', 1, '2024-07-12 20:33:23', '2024-07-12 20:33:45', 21818, '2024-07-12 20:33:45');
INSERT INTO `x_system_log_operate` VALUES (111, 1, 'POST', '错误项目编辑', '127.0.0.1', '/api/admin/monitor_project/edit', 'x_admin/admin/monitor_project.(*MonitorProjectHandler).Edit-fm', '{\"id\":5,\"projectType\":\"uniapp\"}', '', 1, '2024-07-12 20:36:55', '2024-07-12 20:37:14', 19042, '2024-07-12 20:37:14');
INSERT INTO `x_system_log_operate` VALUES (112, 1, 'POST', '错误项目编辑', '127.0.0.1', '/api/admin/monitor_project/edit', 'x_admin/admin/monitor_project.(*MonitorProjectHandler).Edit-fm', '{\"id\":5,\"projectType\":\"go\"}', '', 1, '2024-07-12 20:37:21', '2024-07-12 20:37:58', 37170, '2024-07-12 20:37:58');
INSERT INTO `x_system_log_operate` VALUES (113, 1, 'POST', '错误项目编辑', '127.0.0.1', '/api/admin/monitor_project/edit', 'x_admin/admin/monitor_project.(*MonitorProjectHandler).Edit-fm', '{\"id\":5,\"projectName\":\"2\",\"projectType\":\"uniapp\"}', '', 1, '2024-07-12 20:38:15', '2024-07-12 20:38:42', 27731, '2024-07-12 20:38:42');
INSERT INTO `x_system_log_operate` VALUES (114, 1, 'POST', '错误项目编辑', '127.0.0.1', '/api/admin/monitor_project/edit', 'x_admin/admin/monitor_project.(*MonitorProjectHandler).Edit-fm', '{\"id\":5,\"projectName\":\"\",\"projectType\":\"uniapp\"}', '', 1, '2024-07-12 20:38:52', '2024-07-12 20:39:18', 26213, '2024-07-12 20:39:18');
INSERT INTO `x_system_log_operate` VALUES (115, 1, 'POST', '错误项目编辑', '127.0.0.1', '/api/admin/monitor_project/edit', 'x_admin/admin/monitor_project.(*MonitorProjectHandler).Edit-fm', '{\"id\":5,\"projectName\":\"1\",\"projectType\":\"go\"}', '', 1, '2024-07-12 20:43:28', '2024-07-12 20:43:30', 2528, '2024-07-12 20:43:30');
INSERT INTO `x_system_log_operate` VALUES (116, 1, 'POST', '错误项目编辑', '127.0.0.1', '/api/admin/monitor_project/edit', 'x_admin/admin/monitor_project.(*MonitorProjectHandler).Edit-fm', '{\"id\":5,\"projectName\":\"\",\"projectType\":\"web\"}', '', 1, '2024-07-12 20:43:37', '2024-07-12 20:43:37', 5, '2024-07-12 20:43:37');
INSERT INTO `x_system_log_operate` VALUES (117, 1, 'POST', '错误项目编辑', '127.0.0.1', '/api/admin/monitor_project/edit', 'x_admin/admin/monitor_project.(*MonitorProjectHandler).Edit-fm', '{\"id\":5,\"projectName\":\"2\",\"projectType\":\"go\"}', '', 1, '2024-07-12 20:44:02', '2024-07-12 20:44:02', 5, '2024-07-12 20:44:02');
INSERT INTO `x_system_log_operate` VALUES (118, 1, 'POST', '错误项目编辑', '127.0.0.1', '/api/admin/monitor_project/edit', 'x_admin/admin/monitor_project.(*MonitorProjectHandler).Edit-fm', '{\"id\":5,\"projectName\":\"2\",\"projectType\":\"go\"}', '', 1, '2024-07-12 20:50:28', '2024-07-12 20:50:28', 7, '2024-07-12 20:50:28');
INSERT INTO `x_system_log_operate` VALUES (119, 1, 'POST', '错误项目编辑', '127.0.0.1', '/api/admin/monitor_project/edit', 'x_admin/admin/monitor_project.(*MonitorProjectHandler).Edit-fm', '{\"id\":5,\"projectName\":\"\",\"projectType\":\"go\"}', '', 1, '2024-07-12 20:50:54', '2024-07-12 20:50:54', 5, '2024-07-12 20:50:54');
INSERT INTO `x_system_log_operate` VALUES (120, 1, 'POST', '错误项目编辑', '127.0.0.1', '/api/admin/monitor_project/edit', 'x_admin/admin/monitor_project.(*MonitorProjectHandler).Edit-fm', '{\"id\":5,\"projectName\":\"111\",\"projectType\":\"go\"}', '', 1, '2024-07-12 23:05:28', '2024-07-12 23:05:28', 3, '2024-07-12 23:05:28');
INSERT INTO `x_system_log_operate` VALUES (121, 1, 'POST', '错误项目编辑', '127.0.0.1', '/api/admin/monitor_project/edit', 'x_admin/admin/monitor_project.(*MonitorProjectHandler).Edit-fm', '{\"id\":5,\"projectName\":\"\",\"projectType\":\"go\"}', '', 1, '2024-07-12 23:05:44', '2024-07-12 23:05:47', 2968, '2024-07-12 23:05:47');
INSERT INTO `x_system_log_operate` VALUES (122, 1, 'POST', '错误项目编辑', '127.0.0.1', '/api/admin/monitor_project/edit', 'x_admin/admin/monitor_project.(*MonitorProjectHandler).Edit-fm', '{\"id\":5,\"projectName\":\"\",\"projectType\":\"go\"}', '', 1, '2024-07-12 23:06:08', '2024-07-12 23:06:36', 28503, '2024-07-12 23:06:46');
INSERT INTO `x_system_log_operate` VALUES (123, 1, 'POST', '错误项目编辑', '127.0.0.1', '/api/admin/monitor_project/edit', 'x_admin/admin/monitor_project.(*MonitorProjectHandler).Edit-fm', '{\"id\":5,\"projectName\":\"1\",\"projectType\":\"go\"}', '', 1, '2024-07-12 23:07:00', '2024-07-12 23:07:02', 2183, '2024-07-12 23:07:02');
INSERT INTO `x_system_log_operate` VALUES (124, 1, 'POST', '错误项目编辑', '127.0.0.1', '/api/admin/monitor_project/edit', 'x_admin/admin/monitor_project.(*MonitorProjectHandler).Edit-fm', '{\"id\":5,\"projectName\":\"\",\"projectType\":\"go\"}', '', 1, '2024-07-12 23:07:10', '2024-07-12 23:07:31', 21690, '2024-07-12 23:07:31');
INSERT INTO `x_system_log_operate` VALUES (125, 1, 'POST', '错误项目编辑', '127.0.0.1', '/api/admin/monitor_project/edit', 'x_admin/admin/monitor_project.(*MonitorProjectHandler).Edit-fm', '{\"id\":5,\"projectType\":\"go\"}', 'Error #01: 500:编辑失败\n', 2, '2024-07-12 23:08:07', '2024-07-12 23:08:29', 21723, '2024-07-12 23:08:29');
INSERT INTO `x_system_log_operate` VALUES (126, 1, 'POST', '错误项目新增', '127.0.0.1', '/api/admin/monitor_project/add', 'x_admin/admin/monitor_project.(*MonitorProjectHandler).Add-fm', '{\"id\":\"\",\"projectName\":\"go\",\"projectType\":\"go\"}', '', 1, '2024-07-12 23:17:23', '2024-07-12 23:17:23', 8, '2024-07-12 23:17:23');
INSERT INTO `x_system_log_operate` VALUES (127, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-13 20:17:37', '2024-07-13 20:17:37', 2, '2024-07-13 20:17:37');
INSERT INTO `x_system_log_operate` VALUES (128, 1, 'GET', '服务监控', '127.0.0.1', '/api/admin/monitor/server', 'x_admin/admin/monitor.monitorHandler.server-fm', '', '', 1, '2024-07-13 20:54:39', '2024-07-13 20:54:39', 6, '2024-07-13 20:54:39');
INSERT INTO `x_system_log_operate` VALUES (129, 1, 'GET', '缓存监控', '127.0.0.1', '/api/admin/monitor/cache', 'x_admin/admin/monitor.monitorHandler.cache-fm', '', '', 1, '2024-07-13 20:54:41', '2024-07-13 20:54:41', 1, '2024-07-13 20:54:41');
INSERT INTO `x_system_log_operate` VALUES (130, 1, 'GET', '服务监控', '127.0.0.1', '/api/admin/monitor/server', 'x_admin/admin/monitor.monitorHandler.server-fm', '', '', 1, '2024-07-13 20:54:43', '2024-07-13 20:54:43', 7, '2024-07-13 20:54:43');
INSERT INTO `x_system_log_operate` VALUES (131, 1, 'GET', '服务监控', '127.0.0.1', '/api/admin/monitor/server', 'x_admin/admin/monitor.monitorHandler.server-fm', '', '', 1, '2024-07-13 20:54:53', '2024-07-13 20:54:53', 7, '2024-07-13 20:54:53');
INSERT INTO `x_system_log_operate` VALUES (132, 1, 'GET', '服务监控', '127.0.0.1', '/api/admin/monitor/server', 'x_admin/admin/monitor.monitorHandler.server-fm', '', '', 1, '2024-07-13 20:55:03', '2024-07-13 20:55:03', 6, '2024-07-13 20:55:03');
INSERT INTO `x_system_log_operate` VALUES (133, 1, 'GET', '服务监控', '127.0.0.1', '/api/admin/monitor/server', 'x_admin/admin/monitor.monitorHandler.server-fm', '', '', 1, '2024-07-13 20:55:13', '2024-07-13 20:55:13', 7, '2024-07-13 20:55:13');
INSERT INTO `x_system_log_operate` VALUES (134, 1, 'GET', '服务监控', '127.0.0.1', '/api/admin/monitor/server', 'x_admin/admin/monitor.monitorHandler.server-fm', '', '', 1, '2024-07-13 20:55:23', '2024-07-13 20:55:23', 7, '2024-07-13 20:55:23');
INSERT INTO `x_system_log_operate` VALUES (135, 1, 'POST', '错误项目新增', '127.0.0.1', '/api/admin/monitor_project/add', 'x_admin/admin/monitor_project.(*MonitorProjectHandler).Add-fm', '{\"id\":\"\",\"projectName\":\"1\",\"projectType\":\"go\"}', '', 1, '2024-07-13 20:56:21', '2024-07-13 20:56:21', 40, '2024-07-13 20:56:21');
INSERT INTO `x_system_log_operate` VALUES (136, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-24 12:08:39', '2024-07-24 12:08:39', 9, '2024-07-24 12:08:39');
INSERT INTO `x_system_log_operate` VALUES (137, 1, 'POST', '角色新增', '127.0.0.1', '/api/admin/system/role/add', 'x_admin/admin/system/role.RoleHandler.Add-fm', '{\"id\":\"\",\"isDisable\":0,\"menuIds\":\"\",\"menus\":[],\"name\":\"测试\",\"remark\":\"\",\"sort\":0}', '', 1, '2024-07-24 12:08:46', '2024-07-24 12:08:46', 2, '2024-07-24 12:08:46');
INSERT INTO `x_system_log_operate` VALUES (138, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-24 12:08:46', '2024-07-24 12:08:46', 2, '2024-07-24 12:08:46');
INSERT INTO `x_system_log_operate` VALUES (139, 1, 'GET', '角色详情', '127.0.0.1', '/api/admin/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=2', '', 1, '2024-07-24 12:08:48', '2024-07-24 12:08:48', 19, '2024-07-24 12:08:48');
INSERT INTO `x_system_log_operate` VALUES (140, 1, 'POST', '角色编辑', '127.0.0.1', '/api/admin/system/role/edit', 'x_admin/admin/system/role.RoleHandler.Edit-fm', '{\"id\":2,\"isDisable\":0,\"menuIds\":\"500,1,700,701,783,813,814,815,816,817,818,819,820,784,794,785,786,787,788,789,790,791,792\",\"menus\":[500,1,700,701,783,813,814,815,816,817,818,819,820,784,794,785,786,787,788,789,790,791,792],\"name\":\"测试\",\"remark\":\"\",\"sort\":0}', '', 1, '2024-07-24 12:09:16', '2024-07-24 12:09:16', 39, '2024-07-24 12:09:16');
INSERT INTO `x_system_log_operate` VALUES (141, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-24 12:09:16', '2024-07-24 12:09:16', 2, '2024-07-24 12:09:16');
INSERT INTO `x_system_log_operate` VALUES (142, 1, 'POST', '管理员新增', '127.0.0.1', '/api/admin/system/admin/add', 'x_admin/admin/system/admin.AdminHandler.Add-fm', '{\"avatar\":\"/api/uploads/image/20241705/f0eb36d508834bc2ac1b8c591c563efa.png\",\"deptId\":1,\"id\":\"\",\"isDisable\":0,\"nickname\":\"test\",\"password\":\"da00d628804752a870ca30e83c97f2ba\",\"passwordConfirm\":\"da00d628804752a870ca30e83c97f2ba\",\"postId\":2,\"role\":1,\"sort\":1,\"username\":\"test\"}', '', 1, '2024-07-24 12:09:44', '2024-07-24 12:09:44', 17, '2024-07-24 12:09:44');
INSERT INTO `x_system_log_operate` VALUES (143, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-24 12:10:18', '2024-07-24 12:10:18', 3, '2024-07-24 12:10:18');
INSERT INTO `x_system_log_operate` VALUES (144, 1, 'GET', '角色详情', '127.0.0.1', '/api/admin/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=2', '', 1, '2024-07-24 12:10:20', '2024-07-24 12:10:20', 4, '2024-07-24 12:10:20');
INSERT INTO `x_system_log_operate` VALUES (145, 3, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-24 12:11:47', '2024-07-24 12:11:47', 11, '2024-07-24 12:11:47');
INSERT INTO `x_system_log_operate` VALUES (146, 1, 'POST', '角色编辑', '127.0.0.1', '/api/admin/system/role/edit', 'x_admin/admin/system/role.RoleHandler.Edit-fm', '{\"id\":2,\"isDisable\":0,\"menuIds\":\"1,700,701,200,201,202,203,204,205,206,207,208,209,215,216,217,500,130,131,132,133,134,135,140,141,142,143,144,100,101,102,103,104,105,106,110,111,112,113,114,120,121,122,123,124,501,502,503,505,506,510,511,515,516,517,518,520,521,522,519,550,551,552,553,783,813,814,815,816,817,818,819,820,784,794,785,786,787,788,789,790,791,792,600,610,611,612,613,614,616,617,618,775,778,776,795,796,797,798,799,800,777,801,802,803,804,805,806,807,808,809,810,780,811,812,781,782,821\",\"menus\":[1,700,701,200,201,202,203,204,205,206,207,208,209,215,216,217,500,130,131,132,133,134,135,140,141,142,143,144,100,101,102,103,104,105,106,110,111,112,113,114,120,121,122,123,124,501,502,503,505,506,510,511,515,516,517,518,520,521,522,519,550,551,552,553,783,813,814,815,816,817,818,819,820,784,794,785,786,787,788,789,790,791,792,600,610,611,612,613,614,616,617,618,775,778,776,795,796,797,798,799,800,777,801,802,803,804,805,806,807,808,809,810,780,811,812,781,782,821],\"name\":\"测试\",\"remark\":\"\",\"sort\":0}', '', 1, '2024-07-24 12:13:13', '2024-07-24 12:13:13', 34, '2024-07-24 12:13:13');
INSERT INTO `x_system_log_operate` VALUES (147, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-24 12:13:13', '2024-07-24 12:13:13', 7, '2024-07-24 12:13:13');
INSERT INTO `x_system_log_operate` VALUES (148, 3, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-24 12:26:25', '2024-07-24 12:26:25', 3, '2024-07-24 12:26:25');
INSERT INTO `x_system_log_operate` VALUES (149, 3, 'GET', '服务监控', '127.0.0.1', '/api/admin/monitor/server', 'x_admin/admin/monitor.monitorHandler.server-fm', '', '', 1, '2024-07-24 12:26:35', '2024-07-24 12:26:35', 7, '2024-07-24 12:26:35');
INSERT INTO `x_system_log_operate` VALUES (150, 3, 'GET', '缓存监控', '127.0.0.1', '/api/admin/monitor/cache', 'x_admin/admin/monitor.monitorHandler.cache-fm', '', '', 1, '2024-07-24 12:26:36', '2024-07-24 12:26:36', 1, '2024-07-24 12:26:36');
INSERT INTO `x_system_log_operate` VALUES (151, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-24 12:26:59', '2024-07-24 12:26:59', 12, '2024-07-24 12:26:59');
INSERT INTO `x_system_log_operate` VALUES (152, 1, 'GET', '角色详情', '127.0.0.1', '/api/admin/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=2', '', 1, '2024-07-24 12:27:02', '2024-07-24 12:27:02', 6, '2024-07-24 12:27:02');
INSERT INTO `x_system_log_operate` VALUES (153, 1, 'POST', '角色编辑', '127.0.0.1', '/api/admin/system/role/edit', 'x_admin/admin/system/role.RoleHandler.Edit-fm', '{\"id\":2,\"isDisable\":0,\"menuIds\":\"1,700,701,200,201,202,203,204,205,206,207,208,209,215,216,217,500,130,131,132,133,134,135,140,141,142,143,144,100,101,102,103,104,105,106,110,111,112,113,114,120,121,122,123,124,501,502,503,505,506,510,511,515,516,517,518,520,521,522,519,550,551,552,553,783,813,814,815,816,817,818,819,820,784,794,785,786,787,788,789,790,791,792,778,776,795,796,797,798,799,800,777,801,802,803,804,805,806,807,808,809,810,780,811,812,781,782,821\",\"menus\":[1,700,701,200,201,202,203,204,205,206,207,208,209,215,216,217,500,130,131,132,133,134,135,140,141,142,143,144,100,101,102,103,104,105,106,110,111,112,113,114,120,121,122,123,124,501,502,503,505,506,510,511,515,516,517,518,520,521,522,519,550,551,552,553,783,813,814,815,816,817,818,819,820,784,794,785,786,787,788,789,790,791,792,778,776,795,796,797,798,799,800,777,801,802,803,804,805,806,807,808,809,810,780,811,812,781,782,821],\"name\":\"测试\",\"remark\":\"\",\"sort\":0}', '', 1, '2024-07-24 12:27:07', '2024-07-24 12:27:07', 21, '2024-07-24 12:27:07');
INSERT INTO `x_system_log_operate` VALUES (154, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-24 12:27:07', '2024-07-24 12:27:07', 3, '2024-07-24 12:27:07');
INSERT INTO `x_system_log_operate` VALUES (155, 1, 'GET', '角色详情', '127.0.0.1', '/api/admin/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=2', '', 1, '2024-07-24 12:27:35', '2024-07-24 12:27:35', 9, '2024-07-24 12:27:35');
INSERT INTO `x_system_log_operate` VALUES (156, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-24 12:28:43', '2024-07-24 12:28:43', 27, '2024-07-24 12:28:43');
INSERT INTO `x_system_log_operate` VALUES (157, 1, 'GET', '角色详情', '127.0.0.1', '/api/admin/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=2', '', 1, '2024-07-24 12:28:45', '2024-07-24 12:28:45', 4, '2024-07-24 12:28:45');
INSERT INTO `x_system_log_operate` VALUES (158, 1, 'POST', '角色编辑', '127.0.0.1', '/api/admin/system/role/edit', 'x_admin/admin/system/role.RoleHandler.Edit-fm', '{\"id\":2,\"isDisable\":0,\"menuIds\":\"\",\"menus\":[],\"name\":\"测试\",\"remark\":\"\",\"sort\":0}', '', 1, '2024-07-24 12:28:52', '2024-07-24 12:28:52', 35, '2024-07-24 12:28:52');
INSERT INTO `x_system_log_operate` VALUES (159, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-24 12:28:52', '2024-07-24 12:28:52', 2, '2024-07-24 12:28:52');
INSERT INTO `x_system_log_operate` VALUES (160, 1, 'GET', '角色详情', '127.0.0.1', '/api/admin/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=2', '', 1, '2024-07-24 12:29:32', '2024-07-24 12:29:32', 3, '2024-07-24 12:29:32');
INSERT INTO `x_system_log_operate` VALUES (161, 1, 'POST', '管理员编辑', '127.0.0.1', '/api/admin/system/admin/edit', 'x_admin/admin/system/admin.AdminHandler.Edit-fm', '{\"avatar\":\"/api/uploads/image/20241705/f0eb36d508834bc2ac1b8c591c563efa.png\",\"deptId\":1,\"id\":3,\"isDisable\":0,\"nickname\":\"test\",\"password\":\"\",\"passwordConfirm\":\"\",\"postId\":2,\"role\":2,\"sort\":1,\"username\":\"test\"}', '', 1, '2024-07-24 12:30:09', '2024-07-24 12:30:09', 10, '2024-07-24 12:30:09');
INSERT INTO `x_system_log_operate` VALUES (162, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-24 12:30:47', '2024-07-24 12:30:47', 10, '2024-07-24 12:30:47');
INSERT INTO `x_system_log_operate` VALUES (163, 1, 'GET', '角色详情', '127.0.0.1', '/api/admin/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=2', '', 1, '2024-07-24 12:30:49', '2024-07-24 12:30:49', 8, '2024-07-24 12:30:49');
INSERT INTO `x_system_log_operate` VALUES (164, 1, 'POST', '角色编辑', '127.0.0.1', '/api/admin/system/role/edit', 'x_admin/admin/system/role.RoleHandler.Edit-fm', '{\"id\":2,\"isDisable\":0,\"menuIds\":\"1\",\"menus\":[1],\"name\":\"测试\",\"remark\":\"\",\"sort\":0}', '', 1, '2024-07-24 12:31:01', '2024-07-24 12:31:01', 44, '2024-07-24 12:31:01');
INSERT INTO `x_system_log_operate` VALUES (165, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-24 12:31:01', '2024-07-24 12:31:01', 2, '2024-07-24 12:31:01');
INSERT INTO `x_system_log_operate` VALUES (166, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-24 12:34:53', '2024-07-24 12:34:53', 3, '2024-07-24 12:34:53');
INSERT INTO `x_system_log_operate` VALUES (167, 1, 'GET', '角色详情', '127.0.0.1', '/api/admin/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=2', '', 1, '2024-07-24 12:34:55', '2024-07-24 12:34:55', 8, '2024-07-24 12:34:55');
INSERT INTO `x_system_log_operate` VALUES (168, 1, 'POST', '角色编辑', '127.0.0.1', '/api/admin/system/role/edit', 'x_admin/admin/system/role.RoleHandler.Edit-fm', '{\"id\":2,\"isDisable\":0,\"menuIds\":\"\",\"menus\":[],\"name\":\"测试\",\"remark\":\"\",\"sort\":0}', '', 1, '2024-07-24 12:34:58', '2024-07-24 12:34:58', 8, '2024-07-24 12:34:58');
INSERT INTO `x_system_log_operate` VALUES (169, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-24 12:34:58', '2024-07-24 12:34:58', 6, '2024-07-24 12:34:58');
INSERT INTO `x_system_log_operate` VALUES (170, 1, 'GET', '角色详情', '127.0.0.1', '/api/admin/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=2', '', 1, '2024-07-24 12:35:22', '2024-07-24 12:35:22', 3, '2024-07-24 12:35:22');
INSERT INTO `x_system_log_operate` VALUES (171, 1, 'POST', '角色编辑', '127.0.0.1', '/api/admin/system/role/edit', 'x_admin/admin/system/role.RoleHandler.Edit-fm', '{\"id\":2,\"isDisable\":0,\"menuIds\":\"1\",\"menus\":[1],\"name\":\"测试\",\"remark\":\"\",\"sort\":0}', '', 1, '2024-07-24 12:35:25', '2024-07-24 12:35:25', 47, '2024-07-24 12:35:25');
INSERT INTO `x_system_log_operate` VALUES (172, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-24 12:35:26', '2024-07-24 12:35:26', 3, '2024-07-24 12:35:26');
INSERT INTO `x_system_log_operate` VALUES (173, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-24 12:37:06', '2024-07-24 12:37:06', 12, '2024-07-24 12:37:06');
INSERT INTO `x_system_log_operate` VALUES (174, 1, 'GET', '角色详情', '127.0.0.1', '/api/admin/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=2', '', 1, '2024-07-24 12:37:08', '2024-07-24 12:37:08', 2, '2024-07-24 12:37:08');
INSERT INTO `x_system_log_operate` VALUES (175, 1, 'POST', '角色编辑', '127.0.0.1', '/api/admin/system/role/edit', 'x_admin/admin/system/role.RoleHandler.Edit-fm', '{\"id\":2,\"isDisable\":0,\"menuIds\":\"1\",\"menus\":[1],\"name\":\"测试\",\"remark\":\"\",\"sort\":0}', '', 1, '2024-07-24 12:43:13', '2024-07-24 12:43:13', 37, '2024-07-24 12:43:13');
INSERT INTO `x_system_log_operate` VALUES (176, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-24 12:43:14', '2024-07-24 12:43:14', 7, '2024-07-24 12:43:14');
INSERT INTO `x_system_log_operate` VALUES (177, 1, 'GET', '角色详情', '127.0.0.1', '/api/admin/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=2', '', 1, '2024-07-24 12:43:16', '2024-07-24 12:43:16', 5, '2024-07-24 12:43:16');
INSERT INTO `x_system_log_operate` VALUES (178, 1, 'POST', '角色编辑', '127.0.0.1', '/api/admin/system/role/edit', 'x_admin/admin/system/role.RoleHandler.Edit-fm', '{\"id\":2,\"isDisable\":0,\"menuIds\":\"\",\"menus\":[],\"name\":\"测试\",\"remark\":\"\",\"sort\":0}', '', 1, '2024-07-24 12:43:38', '2024-07-24 12:43:38', 7, '2024-07-24 12:43:38');
INSERT INTO `x_system_log_operate` VALUES (179, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-24 12:43:38', '2024-07-24 12:43:38', 2, '2024-07-24 12:43:38');
INSERT INTO `x_system_log_operate` VALUES (180, 1, 'GET', '角色详情', '127.0.0.1', '/api/admin/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=2', '', 1, '2024-07-24 12:43:52', '2024-07-24 12:43:52', 3, '2024-07-24 12:43:52');
INSERT INTO `x_system_log_operate` VALUES (181, 1, 'GET', '角色详情', '127.0.0.1', '/api/admin/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=1', '', 1, '2024-07-24 12:43:55', '2024-07-24 12:43:55', 8, '2024-07-24 12:43:55');
INSERT INTO `x_system_log_operate` VALUES (182, 1, 'GET', '角色详情', '127.0.0.1', '/api/admin/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=2', '', 1, '2024-07-24 12:45:44', '2024-07-24 12:45:44', 3, '2024-07-24 12:45:44');
INSERT INTO `x_system_log_operate` VALUES (183, 1, 'POST', '角色编辑', '127.0.0.1', '/api/admin/system/role/edit', 'x_admin/admin/system/role.RoleHandler.Edit-fm', '{\"id\":2,\"isDisable\":0,\"menuIds\":\"1\",\"menus\":[1],\"name\":\"测试\",\"remark\":\"\",\"sort\":0}', '', 1, '2024-07-24 12:45:47', '2024-07-24 12:45:47', 43, '2024-07-24 12:45:47');
INSERT INTO `x_system_log_operate` VALUES (184, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-24 12:45:47', '2024-07-24 12:45:47', 3, '2024-07-24 12:45:47');
INSERT INTO `x_system_log_operate` VALUES (185, 1, 'GET', '角色详情', '127.0.0.1', '/api/admin/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=2', '', 1, '2024-07-24 12:45:57', '2024-07-24 12:45:57', 2, '2024-07-24 12:45:57');
INSERT INTO `x_system_log_operate` VALUES (186, 1, 'POST', '角色编辑', '127.0.0.1', '/api/admin/system/role/edit', 'x_admin/admin/system/role.RoleHandler.Edit-fm', '{\"id\":2,\"isDisable\":0,\"menuIds\":\"\",\"menus\":[],\"name\":\"测试\",\"remark\":\"\",\"sort\":0}', '', 1, '2024-07-24 12:46:00', '2024-07-24 12:46:00', 9, '2024-07-24 12:46:00');
INSERT INTO `x_system_log_operate` VALUES (187, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-24 12:46:01', '2024-07-24 12:46:01', 2, '2024-07-24 12:46:01');
INSERT INTO `x_system_log_operate` VALUES (188, 1, 'GET', '角色详情', '127.0.0.1', '/api/admin/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=2', '', 1, '2024-07-24 12:46:08', '2024-07-24 12:46:08', 3, '2024-07-24 12:46:08');
INSERT INTO `x_system_log_operate` VALUES (189, 1, 'POST', '角色编辑', '127.0.0.1', '/api/admin/system/role/edit', 'x_admin/admin/system/role.RoleHandler.Edit-fm', '{\"id\":2,\"isDisable\":0,\"menuIds\":\"1\",\"menus\":[1],\"name\":\"测试\",\"remark\":\"\",\"sort\":0}', '', 1, '2024-07-24 12:46:10', '2024-07-24 12:46:10', 13, '2024-07-24 12:46:10');
INSERT INTO `x_system_log_operate` VALUES (190, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-24 12:46:11', '2024-07-24 12:46:11', 3, '2024-07-24 12:46:11');
INSERT INTO `x_system_log_operate` VALUES (191, 1, 'GET', '角色详情', '127.0.0.1', '/api/admin/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=2', '', 1, '2024-07-24 12:46:15', '2024-07-24 12:46:15', 6, '2024-07-24 12:46:15');
INSERT INTO `x_system_log_operate` VALUES (192, 1, 'POST', '角色编辑', '127.0.0.1', '/api/admin/system/role/edit', 'x_admin/admin/system/role.RoleHandler.Edit-fm', '{\"id\":2,\"isDisable\":0,\"menuIds\":\"\",\"menus\":[],\"name\":\"测试\",\"remark\":\"\",\"sort\":0}', '', 1, '2024-07-24 12:46:19', '2024-07-24 12:46:19', 13, '2024-07-24 12:46:19');
INSERT INTO `x_system_log_operate` VALUES (193, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-24 12:46:19', '2024-07-24 12:46:19', 2, '2024-07-24 12:46:19');
INSERT INTO `x_system_log_operate` VALUES (194, 1, 'GET', '角色详情', '127.0.0.1', '/api/admin/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=2', '', 1, '2024-07-24 12:46:26', '2024-07-24 12:46:26', 2, '2024-07-24 12:46:26');
INSERT INTO `x_system_log_operate` VALUES (195, 1, 'POST', '角色编辑', '127.0.0.1', '/api/admin/system/role/edit', 'x_admin/admin/system/role.RoleHandler.Edit-fm', '{\"id\":2,\"isDisable\":0,\"menuIds\":\"1\",\"menus\":[1],\"name\":\"测试\",\"remark\":\"\",\"sort\":0}', '', 1, '2024-07-24 12:46:33', '2024-07-24 12:46:33', 13, '2024-07-24 12:46:33');
INSERT INTO `x_system_log_operate` VALUES (196, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-24 12:46:34', '2024-07-24 12:46:34', 3, '2024-07-24 12:46:34');
INSERT INTO `x_system_log_operate` VALUES (197, 1, 'GET', '角色详情', '127.0.0.1', '/api/admin/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=2', '', 1, '2024-07-24 12:46:58', '2024-07-24 12:46:58', 3, '2024-07-24 12:46:58');
INSERT INTO `x_system_log_operate` VALUES (198, 1, 'POST', '角色编辑', '127.0.0.1', '/api/admin/system/role/edit', 'x_admin/admin/system/role.RoleHandler.Edit-fm', '{\"id\":2,\"isDisable\":0,\"menuIds\":\"\",\"menus\":[],\"name\":\"测试\",\"remark\":\"\",\"sort\":0}', 'Error #01: 500:Edit Transaction err\n', 2, '2024-07-24 12:47:03', '2024-07-24 12:50:43', 220047, '2024-07-24 12:50:43');
INSERT INTO `x_system_log_operate` VALUES (199, 1, 'POST', '角色编辑', '127.0.0.1', '/api/admin/system/role/edit', 'x_admin/admin/system/role.RoleHandler.Edit-fm', '{\"id\":2,\"isDisable\":0,\"menuIds\":\"\",\"menus\":[],\"name\":\"测试\",\"remark\":\"\",\"sort\":0}', '', 1, '2024-07-24 12:51:10', '2024-07-24 12:51:12', 2400, '2024-07-24 12:51:12');
INSERT INTO `x_system_log_operate` VALUES (200, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-24 12:51:12', '2024-07-24 12:51:12', 9, '2024-07-24 12:51:12');
INSERT INTO `x_system_log_operate` VALUES (201, 1, 'GET', '角色详情', '127.0.0.1', '/api/admin/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=2', '', 1, '2024-07-24 12:51:17', '2024-07-24 12:51:17', 3, '2024-07-24 12:51:17');
INSERT INTO `x_system_log_operate` VALUES (202, 1, 'POST', '角色编辑', '127.0.0.1', '/api/admin/system/role/edit', 'x_admin/admin/system/role.RoleHandler.Edit-fm', '{\"id\":2,\"isDisable\":0,\"menuIds\":\"1\",\"menus\":[1],\"name\":\"测试\",\"remark\":\"\",\"sort\":0}', '', 1, '2024-07-24 12:51:20', '2024-07-24 12:51:22', 2130, '2024-07-24 12:51:22');
INSERT INTO `x_system_log_operate` VALUES (203, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-24 12:51:22', '2024-07-24 12:51:22', 5, '2024-07-24 12:51:22');
INSERT INTO `x_system_log_operate` VALUES (204, 1, 'GET', '角色详情', '127.0.0.1', '/api/admin/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=2', '', 1, '2024-07-24 12:51:27', '2024-07-24 12:51:27', 3, '2024-07-24 12:51:27');
INSERT INTO `x_system_log_operate` VALUES (205, 1, 'POST', '角色编辑', '127.0.0.1', '/api/admin/system/role/edit', 'x_admin/admin/system/role.RoleHandler.Edit-fm', '{\"id\":2,\"isDisable\":0,\"menuIds\":\"\",\"menus\":[],\"name\":\"测试\",\"remark\":\"\",\"sort\":0}', 'Error #01: 500:Edit Transaction err\n', 2, '2024-07-24 12:51:32', '2024-07-24 12:54:35', 183251, '2024-07-24 12:54:35');
INSERT INTO `x_system_log_operate` VALUES (206, 1, 'POST', '角色编辑', '127.0.0.1', '/api/admin/system/role/edit', 'x_admin/admin/system/role.RoleHandler.Edit-fm', '{\"id\":2,\"isDisable\":0,\"menuIds\":\"\",\"menus\":[],\"name\":\"测试\",\"remark\":\"\",\"sort\":0}', '', 1, '2024-07-24 12:55:23', '2024-07-24 12:58:37', 194091, '2024-07-24 12:58:37');
INSERT INTO `x_system_log_operate` VALUES (207, 1, 'POST', '角色编辑', '127.0.0.1', '/api/admin/system/role/edit', 'x_admin/admin/system/role.RoleHandler.Edit-fm', '{\"id\":2,\"isDisable\":0,\"menuIds\":\"\",\"menus\":[],\"name\":\"测试\",\"remark\":\"\",\"sort\":0}', '', 1, '2024-07-24 12:58:44', '2024-07-24 13:00:45', 120970, '2024-07-24 13:00:45');
INSERT INTO `x_system_log_operate` VALUES (208, 1, 'POST', '角色编辑', '127.0.0.1', '/api/admin/system/role/edit', 'x_admin/admin/system/role.RoleHandler.Edit-fm', '{\"id\":2,\"isDisable\":0,\"menuIds\":\"1\",\"menus\":[1],\"name\":\"测试\",\"remark\":\"\",\"sort\":0}', '', 1, '2024-07-24 13:01:01', '2024-07-24 13:01:52', 50542, '2024-07-24 13:01:52');
INSERT INTO `x_system_log_operate` VALUES (209, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-24 13:01:52', '2024-07-24 13:01:52', 3, '2024-07-24 13:01:52');
INSERT INTO `x_system_log_operate` VALUES (210, 1, 'GET', '角色详情', '127.0.0.1', '/api/admin/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=2', '', 1, '2024-07-24 13:02:06', '2024-07-24 13:02:06', 7, '2024-07-24 13:02:06');
INSERT INTO `x_system_log_operate` VALUES (211, 1, 'GET', '角色详情', '127.0.0.1', '/api/admin/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=2', '', 1, '2024-07-24 13:02:10', '2024-07-24 13:02:10', 2, '2024-07-24 13:02:10');
INSERT INTO `x_system_log_operate` VALUES (212, 1, 'POST', '角色编辑', '127.0.0.1', '/api/admin/system/role/edit', 'x_admin/admin/system/role.RoleHandler.Edit-fm', '{\"id\":2,\"isDisable\":0,\"menuIds\":\"\",\"menus\":[],\"name\":\"测试\",\"remark\":\"\",\"sort\":0}', 'Error #01: 500:Edit Transaction err\n', 2, '2024-07-24 13:02:19', '2024-07-24 13:06:04', 225558, '2024-07-24 13:06:04');
INSERT INTO `x_system_log_operate` VALUES (213, 1, 'POST', '角色编辑', '127.0.0.1', '/api/admin/system/role/edit', 'x_admin/admin/system/role.RoleHandler.Edit-fm', '{\"id\":2,\"isDisable\":0,\"menuIds\":\"1\",\"menus\":[1],\"name\":\"测试\",\"remark\":\"\",\"sort\":0}', '', 1, '2024-07-24 13:09:15', '2024-07-24 13:09:19', 3327, '2024-07-24 13:09:19');
INSERT INTO `x_system_log_operate` VALUES (214, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-24 13:09:19', '2024-07-24 13:09:19', 3, '2024-07-24 13:09:19');
INSERT INTO `x_system_log_operate` VALUES (215, 1, 'GET', '角色详情', '127.0.0.1', '/api/admin/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=2', '', 1, '2024-07-24 13:09:32', '2024-07-24 13:09:32', 3, '2024-07-24 13:09:32');
INSERT INTO `x_system_log_operate` VALUES (216, 1, 'POST', '角色编辑', '127.0.0.1', '/api/admin/system/role/edit', 'x_admin/admin/system/role.RoleHandler.Edit-fm', '{\"id\":2,\"isDisable\":0,\"menuIds\":\"\",\"menus\":[],\"name\":\"测试\",\"remark\":\"\",\"sort\":0}', '', 1, '2024-07-24 13:09:34', '2024-07-24 13:09:41', 6556, '2024-07-24 13:09:41');
INSERT INTO `x_system_log_operate` VALUES (217, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-24 13:09:41', '2024-07-24 13:09:41', 3, '2024-07-24 13:09:41');
INSERT INTO `x_system_log_operate` VALUES (218, 1, 'GET', '角色详情', '127.0.0.1', '/api/admin/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=2', '', 1, '2024-07-24 13:09:44', '2024-07-24 13:09:44', 3, '2024-07-24 13:09:44');
INSERT INTO `x_system_log_operate` VALUES (219, 1, 'POST', '角色编辑', '127.0.0.1', '/api/admin/system/role/edit', 'x_admin/admin/system/role.RoleHandler.Edit-fm', '{\"id\":2,\"isDisable\":0,\"menuIds\":\"1\",\"menus\":[1],\"name\":\"测试\",\"remark\":\"\",\"sort\":0}', '', 1, '2024-07-24 13:09:48', '2024-07-24 13:09:48', 8, '2024-07-24 13:09:48');
INSERT INTO `x_system_log_operate` VALUES (220, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-24 13:09:49', '2024-07-24 13:09:49', 3, '2024-07-24 13:09:49');
INSERT INTO `x_system_log_operate` VALUES (221, 1, 'GET', '角色详情', '127.0.0.1', '/api/admin/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=2', '', 1, '2024-07-24 13:09:50', '2024-07-24 13:09:50', 4, '2024-07-24 13:09:50');
INSERT INTO `x_system_log_operate` VALUES (222, 1, 'POST', '角色编辑', '127.0.0.1', '/api/admin/system/role/edit', 'x_admin/admin/system/role.RoleHandler.Edit-fm', '{\"id\":2,\"isDisable\":0,\"menuIds\":\"\",\"menus\":[],\"name\":\"测试\",\"remark\":\"\",\"sort\":0}', '', 1, '2024-07-24 13:10:02', '2024-07-24 13:10:02', 13, '2024-07-24 13:10:02');
INSERT INTO `x_system_log_operate` VALUES (223, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-24 13:10:02', '2024-07-24 13:10:02', 3, '2024-07-24 13:10:02');
INSERT INTO `x_system_log_operate` VALUES (224, 1, 'GET', '角色详情', '127.0.0.1', '/api/admin/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=2', '', 1, '2024-07-24 13:10:10', '2024-07-24 13:10:10', 2, '2024-07-24 13:10:10');
INSERT INTO `x_system_log_operate` VALUES (225, 1, 'POST', '角色编辑', '127.0.0.1', '/api/admin/system/role/edit', 'x_admin/admin/system/role.RoleHandler.Edit-fm', '{\"id\":2,\"isDisable\":0,\"menuIds\":\"1\",\"menus\":[1],\"name\":\"测试\",\"remark\":\"\",\"sort\":0}', '', 1, '2024-07-24 13:10:12', '2024-07-24 13:10:12', 14, '2024-07-24 13:10:12');
INSERT INTO `x_system_log_operate` VALUES (226, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-24 13:10:13', '2024-07-24 13:10:13', 45, '2024-07-24 13:10:13');
INSERT INTO `x_system_log_operate` VALUES (227, 1, 'GET', '角色详情', '127.0.0.1', '/api/admin/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=2', '', 1, '2024-07-24 13:10:15', '2024-07-24 13:10:15', 3, '2024-07-24 13:10:15');
INSERT INTO `x_system_log_operate` VALUES (228, 1, 'POST', '角色编辑', '127.0.0.1', '/api/admin/system/role/edit', 'x_admin/admin/system/role.RoleHandler.Edit-fm', '{\"id\":2,\"isDisable\":0,\"menuIds\":\"\",\"menus\":[],\"name\":\"测试\",\"remark\":\"\",\"sort\":0}', '', 1, '2024-07-24 13:20:50', '2024-07-24 13:20:50', 12, '2024-07-24 13:20:50');
INSERT INTO `x_system_log_operate` VALUES (229, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-24 13:20:50', '2024-07-24 13:20:50', 8, '2024-07-24 13:20:50');
INSERT INTO `x_system_log_operate` VALUES (230, 1, 'GET', '角色详情', '127.0.0.1', '/api/admin/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=2', '', 1, '2024-07-24 13:21:07', '2024-07-24 13:21:07', 4, '2024-07-24 13:21:07');
INSERT INTO `x_system_log_operate` VALUES (231, 1, 'POST', '角色编辑', '127.0.0.1', '/api/admin/system/role/edit', 'x_admin/admin/system/role.RoleHandler.Edit-fm', '{\"id\":2,\"isDisable\":0,\"menuIds\":\"700,701\",\"menus\":[700,701],\"name\":\"测试\",\"remark\":\"\",\"sort\":0}', '', 1, '2024-07-24 13:21:11', '2024-07-24 13:21:11', 16, '2024-07-24 13:21:11');
INSERT INTO `x_system_log_operate` VALUES (232, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-24 13:21:11', '2024-07-24 13:21:11', 3, '2024-07-24 13:21:11');
INSERT INTO `x_system_log_operate` VALUES (233, 1, 'GET', '角色详情', '127.0.0.1', '/api/admin/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=2', '', 1, '2024-07-24 13:21:22', '2024-07-24 13:21:22', 2, '2024-07-24 13:21:22');
INSERT INTO `x_system_log_operate` VALUES (234, 1, 'POST', '角色编辑', '127.0.0.1', '/api/admin/system/role/edit', 'x_admin/admin/system/role.RoleHandler.Edit-fm', '{\"id\":2,\"isDisable\":0,\"menuIds\":\"1,700,701\",\"menus\":[1,700,701],\"name\":\"测试\",\"remark\":\"\",\"sort\":0}', '', 1, '2024-07-24 13:21:24', '2024-07-24 13:21:24', 15, '2024-07-24 13:21:24');
INSERT INTO `x_system_log_operate` VALUES (235, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-24 13:21:24', '2024-07-24 13:21:24', 2, '2024-07-24 13:21:24');
INSERT INTO `x_system_log_operate` VALUES (236, 1, 'POST', '角色删除', '127.0.0.1', '/api/admin/system/role/del', 'x_admin/admin/system/role.RoleHandler.Del-fm', '{\"id\":2}', 'Error #01: 313:角色已被管理员使用,请先移除!\n', 2, '2024-07-24 13:49:22', '2024-07-24 13:49:22', 186, '2024-07-24 13:49:22');
INSERT INTO `x_system_log_operate` VALUES (237, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-24 13:49:30', '2024-07-24 13:49:30', 2, '2024-07-24 13:49:30');
INSERT INTO `x_system_log_operate` VALUES (238, 1, 'POST', '角色删除', '127.0.0.1', '/api/admin/system/role/del', 'x_admin/admin/system/role.RoleHandler.Del-fm', '{\"id\":2}', 'Error #01: 313:角色已被管理员使用,请先移除!\n', 2, '2024-07-24 13:49:35', '2024-07-24 13:49:35', 1, '2024-07-24 13:49:35');
INSERT INTO `x_system_log_operate` VALUES (239, 1, 'POST', '管理员删除', '127.0.0.1', '/api/admin/system/admin/del', 'x_admin/admin/system/admin.AdminHandler.Del-fm', '{\"id\":3}', '', 1, '2024-07-24 13:49:52', '2024-07-24 13:49:52', 3, '2024-07-24 13:49:52');
INSERT INTO `x_system_log_operate` VALUES (240, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-24 13:49:55', '2024-07-24 13:49:55', 2, '2024-07-24 13:49:55');
INSERT INTO `x_system_log_operate` VALUES (241, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-24 13:49:58', '2024-07-24 13:49:58', 3, '2024-07-24 13:49:58');
INSERT INTO `x_system_log_operate` VALUES (242, 1, 'POST', '角色删除', '127.0.0.1', '/api/admin/system/role/del', 'x_admin/admin/system/role.RoleHandler.Del-fm', '{\"id\":2}', '', 1, '2024-07-24 13:50:05', '2024-07-24 13:50:05', 11, '2024-07-24 13:50:05');
INSERT INTO `x_system_log_operate` VALUES (243, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-24 13:50:05', '2024-07-24 13:50:05', 4, '2024-07-24 13:50:05');
INSERT INTO `x_system_log_operate` VALUES (244, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-24 14:26:16', '2024-07-24 14:26:16', 2, '2024-07-24 14:26:16');
INSERT INTO `x_system_log_operate` VALUES (245, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-29 16:43:17', '2024-07-29 16:43:17', 4, '2024-07-29 16:43:17');
INSERT INTO `x_system_log_operate` VALUES (246, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-29 16:48:11', '2024-07-29 16:48:11', 2, '2024-07-29 16:48:11');
INSERT INTO `x_system_log_operate` VALUES (247, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-07-30 17:10:07', '2024-07-30 17:10:07', 2, '2024-07-30 17:10:07');
INSERT INTO `x_system_log_operate` VALUES (248, 1, 'GET', '服务监控', '127.0.0.1', '/api/admin/monitor/server', 'x_admin/admin/monitor.monitorHandler.server-fm', '', '', 1, '2024-07-30 17:58:02', '2024-07-30 17:58:02', 12, '2024-07-30 17:58:02');
INSERT INTO `x_system_log_operate` VALUES (249, 1, 'GET', '缓存监控', '127.0.0.1', '/api/admin/monitor/cache', 'x_admin/admin/monitor.monitorHandler.cache-fm', '', '', 1, '2024-07-30 17:58:03', '2024-07-30 17:58:03', 10, '2024-07-30 17:58:03');
INSERT INTO `x_system_log_operate` VALUES (250, 1, 'GET', '服务监控', '127.0.0.1', '/api/admin/monitor/server', 'x_admin/admin/monitor.monitorHandler.server-fm', '', '', 1, '2024-07-30 18:00:33', '2024-07-30 18:00:33', 11, '2024-07-30 18:00:33');
INSERT INTO `x_system_log_operate` VALUES (251, 1, 'GET', '缓存监控', '127.0.0.1', '/api/admin/monitor/cache', 'x_admin/admin/monitor.monitorHandler.cache-fm', '', '', 1, '2024-07-30 18:00:36', '2024-07-30 18:00:36', 1, '2024-07-30 18:00:36');
INSERT INTO `x_system_log_operate` VALUES (252, 1, 'GET', '服务监控', '127.0.0.1', '/api/admin/monitor/server', 'x_admin/admin/monitor.monitorHandler.server-fm', '', '', 1, '2024-07-30 18:39:07', '2024-07-30 18:39:07', 9, '2024-07-30 18:39:07');
INSERT INTO `x_system_log_operate` VALUES (253, 1, 'GET', '缓存监控', '127.0.0.1', '/api/admin/monitor/cache', 'x_admin/admin/monitor.monitorHandler.cache-fm', '', '', 1, '2024-07-30 18:39:08', '2024-07-30 18:39:08', 0, '2024-07-30 18:39:08');
INSERT INTO `x_system_log_operate` VALUES (254, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-08-05 18:44:26', '2024-08-05 18:44:26', 22, '2024-08-05 18:44:26');
INSERT INTO `x_system_log_operate` VALUES (255, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-08-05 18:44:56', '2024-08-05 18:44:56', 2, '2024-08-05 18:44:56');
INSERT INTO `x_system_log_operate` VALUES (256, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-08-08 01:14:24', '2024-08-08 01:14:24', 11, '2024-08-08 01:14:24');
INSERT INTO `x_system_log_operate` VALUES (257, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-08-08 12:32:52', '2024-08-08 12:32:52', 3, '2024-08-08 12:32:52');
INSERT INTO `x_system_log_operate` VALUES (258, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-08-08 12:33:02', '2024-08-08 12:33:02', 1, '2024-08-08 12:33:02');
INSERT INTO `x_system_log_operate` VALUES (259, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-08-08 12:33:04', '2024-08-08 12:33:04', 3, '2024-08-08 12:33:04');
INSERT INTO `x_system_log_operate` VALUES (260, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e\\u003cbr\\u003e\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"123\",\"results\":2,\"scene\":\"123\",\"sendTime\":\"123\",\"status\":1}', 'Error #01: 310:参数校验错误\n', 2, '2024-08-08 13:11:47', '2024-08-08 13:11:47', 1, '2024-08-08 13:11:47');
INSERT INTO `x_system_log_operate` VALUES (261, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e\\u003cbr\\u003e\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"123\",\"results\":2,\"scene\":\"123\",\"sendTime\":\"123\",\"status\":1}', 'Error #01: 310:参数校验错误\n', 2, '2024-08-08 13:14:45', '2024-08-08 13:14:45', 1, '2024-08-08 13:14:45');
INSERT INTO `x_system_log_operate` VALUES (262, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e111\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"111\",\"results\":2,\"scene\":\"111\",\"sendTime\":\"1\",\"status\":1}', 'Error #01: 310:参数校验错误\n', 2, '2024-08-08 16:18:24', '2024-08-08 16:18:24', 1, '2024-08-08 16:18:24');
INSERT INTO `x_system_log_operate` VALUES (263, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e111\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"111\",\"results\":2,\"scene\":\"111\",\"sendTime\":\"1\",\"status\":1}', '', 1, '2024-08-08 16:19:41', '2024-08-08 16:21:28', 106456, '2024-08-08 16:21:28');
INSERT INTO `x_system_log_operate` VALUES (264, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e111\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"111\",\"results\":2,\"scene\":\"111\",\"sendTime\":\"111\",\"status\":1}', 'Error #01: 310:参数校验错误\n', 2, '2024-08-08 16:50:33', '2024-08-08 16:50:33', 1, '2024-08-08 16:50:33');
INSERT INTO `x_system_log_operate` VALUES (265, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e111\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"111\",\"results\":2,\"scene\":\"111\",\"sendTime\":\"111\",\"status\":1}', '', 1, '2024-08-08 16:51:14', '2024-08-08 16:51:51', 37457, '2024-08-08 16:51:51');
INSERT INTO `x_system_log_operate` VALUES (266, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e2222\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"2222\",\"results\":2,\"scene\":\"2222\",\"sendTime\":\"222\",\"status\":1}', '', 1, '2024-08-08 16:52:17', '2024-08-08 16:52:20', 2993, '2024-08-08 16:52:20');
INSERT INTO `x_system_log_operate` VALUES (267, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e111\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"11\",\"results\":2,\"scene\":\"11\",\"sendTime\":\"1112\",\"status\":1}', 'Error #01: 310:参数校验错误\n', 2, '2024-08-08 22:32:39', '2024-08-08 22:32:39', 0, '2024-08-08 22:32:39');
INSERT INTO `x_system_log_operate` VALUES (268, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e111\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"11\",\"results\":2,\"scene\":\"11\",\"sendTime\":\"1112\",\"status\":1}', 'Error #01: 310:参数校验错误\n', 2, '2024-08-08 22:33:43', '2024-08-08 22:33:43', 0, '2024-08-08 22:33:43');
INSERT INTO `x_system_log_operate` VALUES (269, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e111\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"11\",\"results\":2,\"scene\":\"11\",\"sendTime\":\"1112\",\"status\":1}', 'Error #01: 310:参数校验错误\n', 2, '2024-08-08 22:36:09', '2024-08-08 22:41:15', 305380, '2024-08-08 22:41:15');
INSERT INTO `x_system_log_operate` VALUES (270, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e111\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"11\",\"results\":1,\"scene\":\"11\",\"sendTime\":\"1112\",\"status\":3}', 'Error #01: 310:参数校验错误\n', 2, '2024-08-08 22:41:15', '2024-08-08 22:41:41', 26261, '2024-08-08 22:41:41');
INSERT INTO `x_system_log_operate` VALUES (271, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e111\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"11\",\"results\":1,\"scene\":\"11\",\"sendTime\":\"1112\",\"status\":3}', 'Error #01: 310:参数校验错误\n', 2, '2024-08-08 22:42:08', '2024-08-08 22:44:07', 118673, '2024-08-08 22:44:07');
INSERT INTO `x_system_log_operate` VALUES (272, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e111\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"11\",\"results\":1,\"scene\":\"11\",\"sendTime\":\"1112\",\"status\":3}', 'Error #01: 310:参数校验错误\n', 2, '2024-08-08 22:46:58', '2024-08-08 22:47:09', 10971, '2024-08-08 22:47:09');
INSERT INTO `x_system_log_operate` VALUES (273, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e111\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"11\",\"results\":1,\"scene\":\"11\",\"sendTime\":\"1112\",\"status\":3}', 'Error #01: 310:参数校验错误\n', 2, '2024-08-08 22:48:30', '2024-08-08 22:50:36', 125168, '2024-08-08 22:50:36');
INSERT INTO `x_system_log_operate` VALUES (274, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e111\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"11\",\"results\":1,\"scene\":\"11\",\"sendTime\":\"1112\",\"status\":3}', '', 1, '2024-08-08 22:52:03', '2024-08-08 22:52:18', 14845, '2024-08-08 22:52:18');
INSERT INTO `x_system_log_operate` VALUES (275, 1, 'POST', '系统短信日志删除', '127.0.0.1', '/api/admin/system_log_sms/del', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Del-fm', '{\"id\":1}', '', 1, '2024-08-08 23:40:05', '2024-08-08 23:40:05', 9, '2024-08-08 23:40:05');
INSERT INTO `x_system_log_operate` VALUES (276, 1, 'POST', '系统短信日志删除', '127.0.0.1', '/api/admin/system_log_sms/del', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Del-fm', '{\"id\":2}', '', 1, '2024-08-08 23:40:08', '2024-08-08 23:40:08', 3, '2024-08-08 23:40:08');
INSERT INTO `x_system_log_operate` VALUES (277, 1, 'POST', '系统短信日志删除', '127.0.0.1', '/api/admin/system_log_sms/del', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Del-fm', '{\"id\":3}', '', 1, '2024-08-08 23:40:18', '2024-08-08 23:40:18', 2, '2024-08-08 23:40:18');
INSERT INTO `x_system_log_operate` VALUES (278, 1, 'POST', '系统短信日志删除', '127.0.0.1', '/api/admin/system_log_sms/del', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Del-fm', '{\"id\":4}', '', 1, '2024-08-08 23:40:20', '2024-08-08 23:40:20', 4, '2024-08-08 23:40:20');
INSERT INTO `x_system_log_operate` VALUES (279, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e111\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"11\",\"results\":\"111\",\"scene\":\"11\",\"sendTime\":\"123\",\"status\":2}', '', 1, '2024-08-08 23:40:56', '2024-08-08 23:40:56', 4, '2024-08-08 23:40:56');
INSERT INTO `x_system_log_operate` VALUES (280, 1, 'POST', '系统短信日志删除', '127.0.0.1', '/api/admin/system_log_sms/del', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Del-fm', '{\"id\":5}', '', 1, '2024-08-08 23:41:29', '2024-08-08 23:41:29', 3, '2024-08-08 23:41:29');
INSERT INTO `x_system_log_operate` VALUES (281, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"1\",\"results\":\"1\",\"scene\":\"1\",\"sendTime\":\"1\",\"status\":1}', '', 1, '2024-08-08 23:41:45', '2024-08-08 23:41:45', 6, '2024-08-08 23:41:45');
INSERT INTO `x_system_log_operate` VALUES (282, 1, 'POST', '上传图片', '127.0.0.1', '/api/admin/common/upload/image', 'x_admin/admin/common/upload.uploadHandler.uploadImage-fm', 'COMPUMAX_1p2x61.jpg', '', 1, '2024-08-09 00:21:03', '2024-08-09 00:21:03', 204, '2024-08-09 00:21:03');
INSERT INTO `x_system_log_operate` VALUES (283, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e\\u003cimg src=\\\"/api/uploads/image/20240908/98587b5a927a4a18b5f7d165a3e60287.jpg\\\" alt=\\\"\\\" data-href=\\\"\\\" style=\\\"width: 520.00px;height: 292.50px;\\\"/\\u003e\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"1\",\"results\":\"1\",\"scene\":\"1\",\"sendTime\":\"1\",\"status\":1}', '', 1, '2024-08-09 00:21:49', '2024-08-09 00:21:49', 3, '2024-08-09 00:21:49');
INSERT INTO `x_system_log_operate` VALUES (284, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e123 \\u0026nbsp; \\u0026nbsp;\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"123\",\"results\":\"123\",\"scene\":\"\",\"sendTime\":\"123\",\"status\":1}', 'Error #01: 310:参数校验错误\n', 2, '2024-08-09 02:29:27', '2024-08-09 02:29:27', 0, '2024-08-09 02:29:27');
INSERT INTO `x_system_log_operate` VALUES (285, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e123\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"123\",\"results\":\"123\",\"scene\":\"\",\"sendTime\":\"123\",\"status\":1}', 'Error #01: 310:参数校验错误\n', 2, '2024-08-09 02:29:34', '2024-08-09 02:29:34', 0, '2024-08-09 02:29:34');
INSERT INTO `x_system_log_operate` VALUES (286, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e123\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"123\",\"results\":\"123\",\"scene\":\"\",\"sendTime\":\"123\",\"status\":1}', 'Error #01: 310:参数校验错误\n', 2, '2024-08-09 02:30:01', '2024-08-09 02:30:01', 1, '2024-08-09 02:30:01');
INSERT INTO `x_system_log_operate` VALUES (287, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e123\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"123\",\"results\":\"123\",\"scene\":\"\",\"sendTime\":\"123\",\"status\":1}', 'Error #01: 310:参数校验错误\n', 2, '2024-08-09 02:32:14', '2024-08-09 02:32:14', 1, '2024-08-09 02:32:14');
INSERT INTO `x_system_log_operate` VALUES (288, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e123\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"123\",\"results\":\"123\",\"scene\":\"\",\"sendTime\":\"123\",\"status\":1}', 'Error #01: 310:参数校验错误\n', 2, '2024-08-09 02:33:38', '2024-08-09 02:33:38', 0, '2024-08-09 02:33:38');
INSERT INTO `x_system_log_operate` VALUES (289, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-08-09 11:39:01', '2024-08-09 11:39:01', 3, '2024-08-09 11:39:01');
INSERT INTO `x_system_log_operate` VALUES (290, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"content\":\"\\u003cp\\u003e\\u003cimg src=\\\"/api/uploads/image/20240908/98587b5a927a4a18b5f7d165a3e60287.jpg\\\" alt=\\\"\\\" data-href=\\\"\\\" style=\\\"width: 520.00px;height: 292.50px;\\\"/\\u003e\\u003c/p\\u003e\",\"id\":7,\"mobile\":\"2\",\"results\":\"2\",\"scene\":\"2\",\"sendTime\":\"2\",\"status\":2}', '', 1, '2024-08-09 11:45:45', '2024-08-09 11:46:27', 42762, '2024-08-09 11:46:27');
INSERT INTO `x_system_log_operate` VALUES (291, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=68ee106c69b3815ca7b11a362eb210e0BcwiJu&content&createTimeEnd&createTimeStart&mobile&results&scene=&sendTime&status&updateTimeEnd&updateTimeStart', 'Error #01: 310:参数校验错误\n', 2, '2024-08-09 12:23:50', '2024-08-09 12:23:50', 0, '2024-08-09 12:23:50');
INSERT INTO `x_system_log_operate` VALUES (292, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=68ee106c69b3815ca7b11a362eb210e0BcwiJu&content&createTimeEnd&createTimeStart&mobile&results&scene&sendTime&status&updateTimeEnd&updateTimeStart', 'Error #01: 310:参数校验错误\n', 2, '2024-08-09 12:24:05', '2024-08-09 12:24:05', 0, '2024-08-09 12:24:05');
INSERT INTO `x_system_log_operate` VALUES (293, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=68ee106c69b3815ca7b11a362eb210e0BcwiJu', '', 1, '2024-08-09 12:24:38', '2024-08-09 12:24:38', 15, '2024-08-09 12:24:38');
INSERT INTO `x_system_log_operate` VALUES (294, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=68ee106c69b3815ca7b11a362eb210e0BcwiJu&content&createTimeEnd&createTimeStart&mobile&results&scene&sendTime&status&updateTimeEnd&updateTimeStart', 'Error #01: 310:参数校验错误\n', 2, '2024-08-09 12:27:33', '2024-08-09 12:27:33', 1, '2024-08-09 12:27:33');
INSERT INTO `x_system_log_operate` VALUES (295, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=68ee106c69b3815ca7b11a362eb210e0BcwiJu&content&createTimeEnd&createTimeStart&mobile&results&scene&sendTime&status&updateTimeEnd&updateTimeStart', 'Error #01: 310:参数校验错误\n', 2, '2024-08-09 13:15:46', '2024-08-09 13:15:46', 1, '2024-08-09 13:15:46');
INSERT INTO `x_system_log_operate` VALUES (296, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=68ee106c69b3815ca7b11a362eb210e0BcwiJu', '', 1, '2024-08-09 12:27:49', '2024-08-09 13:15:46', 2876109, '2024-08-09 13:15:46');
INSERT INTO `x_system_log_operate` VALUES (297, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"id\":6,\"mobile\":\"1\",\"results\":\"1\",\"scene\":1,\"sendTime\":1,\"status\":1}', '', 1, '2024-08-09 17:03:13', '2024-08-09 17:03:13', 74, '2024-08-09 17:03:13');
INSERT INTO `x_system_log_operate` VALUES (298, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"id\":6,\"mobile\":\"1\",\"results\":\"1\",\"scene\":\"\",\"sendTime\":1,\"status\":1}', 'Error #01: 310:参数校验错误\n', 2, '2024-08-09 17:03:19', '2024-08-09 17:03:19', 1, '2024-08-09 17:03:19');
INSERT INTO `x_system_log_operate` VALUES (299, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"id\":6,\"mobile\":\"1\",\"results\":\"1\",\"scene\":\"\",\"sendTime\":1,\"status\":1}', 'Error #01: 310:参数校验错误\n', 2, '2024-08-09 17:03:24', '2024-08-09 17:03:24', 0, '2024-08-09 17:03:24');
INSERT INTO `x_system_log_operate` VALUES (300, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"id\":7,\"mobile\":\"1\",\"results\":\"1\",\"scene\":2,\"sendTime\":2,\"status\":2}', '', 1, '2024-08-09 17:13:44', '2024-08-09 17:13:44', 2, '2024-08-09 17:13:44');
INSERT INTO `x_system_log_operate` VALUES (301, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"id\":7,\"mobile\":\"1\",\"results\":\"1\",\"scene\":\"\",\"sendTime\":2,\"status\":2}', 'Error #01: 310:参数校验错误\n', 2, '2024-08-09 17:13:52', '2024-08-09 17:13:52', 0, '2024-08-09 17:13:52');
INSERT INTO `x_system_log_operate` VALUES (302, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"id\":7,\"mobile\":\"1\",\"results\":\"1\",\"scene\":\"\",\"sendTime\":2,\"status\":2}', 'Error #01: 310:参数校验错误\n', 2, '2024-08-09 17:14:42', '2024-08-09 17:14:42', 0, '2024-08-09 17:14:42');
INSERT INTO `x_system_log_operate` VALUES (303, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"id\":7,\"mobile\":\"1\",\"results\":\"1\",\"scene\":\"\",\"sendTime\":2,\"status\":2}', '', 1, '2024-08-09 17:17:31', '2024-08-09 17:17:31', 10, '2024-08-09 17:17:31');
INSERT INTO `x_system_log_operate` VALUES (304, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"id\":7,\"mobile\":\"1\",\"results\":\"1\",\"scene\":\"\",\"sendTime\":2,\"status\":2}', '', 1, '2024-08-09 17:18:15', '2024-08-09 17:22:32', 257080, '2024-08-09 17:22:32');
INSERT INTO `x_system_log_operate` VALUES (305, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"id\":7,\"mobile\":\"1\",\"results\":\"1\",\"scene\":\"\",\"sendTime\":2,\"status\":2}', '', 1, '2024-08-09 17:47:28', '2024-08-09 17:48:28', 60716, '2024-08-09 17:48:28');
INSERT INTO `x_system_log_operate` VALUES (306, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"id\":7,\"mobile\":\"1\",\"results\":\"1\",\"scene\":\"\",\"sendTime\":2,\"status\":2}', '', 1, '2024-08-09 17:48:23', '2024-08-09 17:48:28', 5370, '2024-08-09 17:48:28');
INSERT INTO `x_system_log_operate` VALUES (307, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"id\":7,\"mobile\":\"222\",\"results\":\"1\",\"scene\":\"\",\"sendTime\":2,\"status\":2}', '', 1, '2024-08-09 18:54:26', '2024-08-09 18:54:46', 20314, '2024-08-09 18:54:46');
INSERT INTO `x_system_log_operate` VALUES (308, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"content\":\"\\u003cp\\u003e1444\\u003c/p\\u003e\",\"id\":7,\"mobile\":\"333445566\",\"results\":\"1\",\"scene\":\"\",\"sendTime\":2,\"status\":2}', 'reflect: Field index out of range', 2, '2024-08-09 19:29:31', '2024-08-09 19:29:31', 6, '2024-08-09 19:29:31');
INSERT INTO `x_system_log_operate` VALUES (309, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"content\":\"\\u003cp\\u003e1444\\u003c/p\\u003e\",\"id\":7,\"mobile\":\"333445566\",\"results\":\"1\",\"scene\":\"\",\"sendTime\":2,\"status\":2}', 'reflect: Field index out of range', 2, '2024-08-09 19:30:54', '2024-08-09 19:30:54', 2, '2024-08-09 19:30:54');
INSERT INTO `x_system_log_operate` VALUES (310, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"content\":\"\\u003cp\\u003e1444\\u003c/p\\u003e\",\"id\":7,\"mobile\":\"333445566\",\"results\":\"1\",\"scene\":\"\",\"sendTime\":2,\"status\":2}', 'reflect: Field index out of range', 2, '2024-08-09 19:31:29', '2024-08-09 19:31:39', 10132, '2024-08-09 19:31:39');
INSERT INTO `x_system_log_operate` VALUES (311, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"content\":\"\\u003cp\\u003e1444\\u003c/p\\u003e\",\"id\":7,\"mobile\":\"333445566\",\"results\":\"1\",\"scene\":\"\",\"sendTime\":2,\"status\":2}', 'Error #01: 310:参数校验错误\n', 2, '2024-08-09 19:32:59', '2024-08-09 19:32:59', 1, '2024-08-09 19:32:59');
INSERT INTO `x_system_log_operate` VALUES (312, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"content\":\"\\u003cp\\u003e1444\\u003c/p\\u003e\",\"id\":7,\"mobile\":\"333445566\",\"results\":\"1\",\"scene\":\"\",\"sendTime\":2,\"status\":2}', 'reflect: Field index out of range', 2, '2024-08-09 19:34:17', '2024-08-09 19:34:22', 4955, '2024-08-09 19:34:22');
INSERT INTO `x_system_log_operate` VALUES (313, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"content\":\"\\u003cp\\u003e1444\\u003c/p\\u003e\",\"id\":7,\"mobile\":\"333445566\",\"results\":\"1\",\"scene\":\"\",\"sendTime\":2,\"status\":2}', 'Error #01: 500:编辑失败\n', 2, '2024-08-09 19:40:59', '2024-08-09 19:46:38', 338427, '2024-08-09 19:46:38');
INSERT INTO `x_system_log_operate` VALUES (314, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"content\":\"\\u003cp\\u003e1444\\u003c/p\\u003e\",\"id\":7,\"mobile\":\"333445566\",\"results\":\"1\",\"scene\":\"\",\"sendTime\":2,\"status\":2}', 'Error #01: 500:编辑失败\n', 2, '2024-08-09 19:46:38', '2024-08-09 19:46:40', 2154, '2024-08-09 19:46:40');
INSERT INTO `x_system_log_operate` VALUES (315, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"content\":\"\\u003cp\\u003e1444\\u003c/p\\u003e\",\"id\":7,\"mobile\":\"333445566\",\"results\":\"1\",\"scene\":\"\",\"sendTime\":2,\"status\":2}', 'Error #01: 500:编辑失败\n', 2, '2024-08-09 20:15:39', '2024-08-09 20:15:52', 13208, '2024-08-09 20:15:52');
INSERT INTO `x_system_log_operate` VALUES (316, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"content\":\"\\u003cp\\u003e122\\u003c/p\\u003e\",\"id\":7,\"mobile\":\"222222\",\"results\":\"1\",\"scene\":\"2222\",\"sendTime\":2,\"status\":2}', '', 1, '2024-08-09 21:05:03', '2024-08-09 21:05:14', 11601, '2024-08-09 21:05:14');
INSERT INTO `x_system_log_operate` VALUES (317, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"content\":\"\\u003cp\\u003e122\\u003c/p\\u003e\",\"id\":7,\"mobile\":\"222222\",\"results\":\"1\",\"scene\":\"0\",\"sendTime\":2,\"status\":2}', '', 1, '2024-08-09 21:05:30', '2024-08-09 21:05:30', 3, '2024-08-09 21:05:30');
INSERT INTO `x_system_log_operate` VALUES (318, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"content\":\"\\u003cp\\u003e122\\u003c/p\\u003e\",\"id\":7,\"mobile\":\"222222\",\"results\":\"1\",\"scene\":2222,\"sendTime\":2,\"status\":2}', '', 1, '2024-08-09 21:08:33', '2024-08-09 21:08:33', 12, '2024-08-09 21:08:33');
INSERT INTO `x_system_log_operate` VALUES (319, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"content\":\"\\u003cp\\u003e122\\u003c/p\\u003e\",\"id\":7,\"mobile\":\"222222\",\"results\":\"1\",\"scene\":\"2222333\",\"sendTime\":2,\"status\":2}', '', 1, '2024-08-09 21:08:40', '2024-08-09 21:08:40', 4, '2024-08-09 21:08:40');
INSERT INTO `x_system_log_operate` VALUES (320, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"content\":\"\\u003cp\\u003e122\\u003c/p\\u003e\",\"id\":7,\"mobile\":\"222222\",\"results\":\"1\",\"scene\":\"\",\"sendTime\":2,\"status\":2}', 'Error #01: 500:编辑失败\n', 2, '2024-08-09 21:08:53', '2024-08-09 21:08:53', 3, '2024-08-09 21:08:53');
INSERT INTO `x_system_log_operate` VALUES (321, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=91df9a27c406e1ec61cd95432389be61zQdlTO&', '', 1, '2024-08-09 22:42:09', '2024-08-09 22:42:12', 3233, '2024-08-09 22:42:12');
INSERT INTO `x_system_log_operate` VALUES (322, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=91df9a27c406e1ec61cd95432389be61zQdlTO&', '', 1, '2024-08-09 22:43:15', '2024-08-09 22:43:19', 4356, '2024-08-09 22:43:19');
INSERT INTO `x_system_log_operate` VALUES (323, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"content\":\"\\u003cp\\u003e122\\u003c/p\\u003e\",\"id\":7,\"mobile\":\"222222\",\"results\":\"1\",\"scene\":\"\",\"sendTime\":2,\"status\":2}', 'Error #01: 500:编辑失败\n', 2, '2024-08-09 22:47:02', '2024-08-09 22:47:02', 2, '2024-08-09 22:47:02');
INSERT INTO `x_system_log_operate` VALUES (324, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"content\":\"\\u003cp\\u003e122\\u003c/p\\u003e\",\"id\":7,\"mobile\":\"222222\",\"results\":\"1\",\"scene\":\"0\",\"sendTime\":2,\"status\":2}', '', 1, '2024-08-09 22:47:14', '2024-08-09 22:47:14', 5, '2024-08-09 22:47:14');
INSERT INTO `x_system_log_operate` VALUES (325, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"content\":\"\\u003cp\\u003e122\\u003c/p\\u003e\",\"id\":7,\"mobile\":\"2\",\"results\":\"1\",\"scene\":0,\"sendTime\":2,\"status\":2}', '', 1, '2024-08-09 22:47:27', '2024-08-09 22:47:27', 7, '2024-08-09 22:47:27');
INSERT INTO `x_system_log_operate` VALUES (326, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"content\":\"\\u003cp\\u003e\\u003cbr\\u003e\\u003c/p\\u003e\",\"id\":7,\"mobile\":\"2\",\"results\":\"1\",\"scene\":0,\"sendTime\":2,\"status\":2}', '', 1, '2024-08-09 22:47:37', '2024-08-09 22:47:37', 6, '2024-08-09 22:47:37');
INSERT INTO `x_system_log_operate` VALUES (327, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"content\":\"\\u003cp\\u003e111\\u003c/p\\u003e\",\"id\":7,\"mobile\":\"2\",\"results\":\"1\",\"scene\":0,\"sendTime\":2,\"status\":2}', '', 1, '2024-08-09 22:50:57', '2024-08-09 22:50:57', 5, '2024-08-09 22:50:57');
INSERT INTO `x_system_log_operate` VALUES (328, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"content\":\"\\u003cp\\u003e\\u003cbr\\u003e\\u003c/p\\u003e\",\"id\":7,\"mobile\":\"2\",\"results\":\"1\",\"scene\":0,\"sendTime\":2,\"status\":2}', '', 1, '2024-08-09 22:51:15', '2024-08-09 22:51:15', 4, '2024-08-09 22:51:15');
INSERT INTO `x_system_log_operate` VALUES (329, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-08-09 22:53:22', '2024-08-09 22:53:22', 2, '2024-08-09 22:53:22');
INSERT INTO `x_system_log_operate` VALUES (330, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"content\":\"\\u003cp\\u003e\\u003cbr\\u003e\\u003c/p\\u003e\",\"id\":7,\"results\":\"1\",\"scene\":0,\"sendTime\":2,\"status\":2}', '', 1, '2024-08-09 22:54:00', '2024-08-09 22:54:00', 4, '2024-08-09 22:54:00');
INSERT INTO `x_system_log_operate` VALUES (331, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"content\":\"\\u003cp\\u003e\\u003cbr\\u003e\\u003c/p\\u003e\",\"id\":7,\"results\":\"1\",\"scene\":0,\"sendTime\":2,\"status\":2}', '', 1, '2024-08-09 22:54:17', '2024-08-09 22:54:17', 4, '2024-08-09 22:54:17');
INSERT INTO `x_system_log_operate` VALUES (332, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"content\":\"\\u003cp\\u003e\\u003cbr\\u003e\\u003c/p\\u003e\",\"id\":7,\"mobile\":\"2333\",\"results\":\"1\",\"scene\":0,\"sendTime\":2,\"status\":2}', '', 1, '2024-08-09 22:54:29', '2024-08-09 22:54:29', 5, '2024-08-09 22:54:29');
INSERT INTO `x_system_log_operate` VALUES (333, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"content\":\"\\u003cp\\u003e\\u003cbr\\u003e\\u003c/p\\u003e\",\"id\":7,\"results\":\"1\",\"scene\":0,\"sendTime\":2,\"status\":2}', '', 1, '2024-08-09 22:59:22', '2024-08-09 22:59:58', 35528, '2024-08-09 22:59:58');
INSERT INTO `x_system_log_operate` VALUES (334, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"content\":\"\\u003cp\\u003e\\u003cbr\\u003e\\u003c/p\\u003e\",\"id\":7,\"results\":\"1\",\"scene\":0,\"sendTime\":2,\"status\":2}', '', 1, '2024-08-09 23:00:06', '2024-08-09 23:00:36', 30478, '2024-08-09 23:00:36');
INSERT INTO `x_system_log_operate` VALUES (335, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"content\":\"\\u003cp\\u003e\\u003cbr\\u003e\\u003c/p\\u003e\",\"id\":7,\"mobile\":\"\",\"results\":\"1\",\"scene\":0,\"sendTime\":2,\"status\":2}', '', 1, '2024-08-09 23:01:03', '2024-08-09 23:03:07', 124714, '2024-08-09 23:03:07');
INSERT INTO `x_system_log_operate` VALUES (336, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e11\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"11\",\"results\":\"1\",\"scene\":\"\",\"sendTime\":\"1\",\"status\":1}', 'Error #01: 310:参数校验错误\n', 2, '2024-08-09 23:13:38', '2024-08-09 23:13:38', 0, '2024-08-09 23:13:38');
INSERT INTO `x_system_log_operate` VALUES (337, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e11\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"11\",\"results\":\"1\",\"scene\":\"0\",\"sendTime\":\"1\",\"status\":1}', 'Error #01: 310:参数校验错误\n', 2, '2024-08-09 23:13:58', '2024-08-09 23:13:58', 0, '2024-08-09 23:13:58');
INSERT INTO `x_system_log_operate` VALUES (338, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e11\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"11\",\"results\":\"1\",\"scene\":\"0\",\"sendTime\":\"1\",\"status\":1}', 'Error #01: 310:参数校验错误\n', 2, '2024-08-09 23:14:25', '2024-08-09 23:14:25', 0, '2024-08-09 23:14:25');
INSERT INTO `x_system_log_operate` VALUES (339, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e11\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"11\",\"results\":\"1\",\"scene\":\"1\",\"sendTime\":\"1\",\"status\":1}', 'Error #01: 310:参数校验错误\n', 2, '2024-08-09 23:14:49', '2024-08-09 23:14:49', 0, '2024-08-09 23:14:49');
INSERT INTO `x_system_log_operate` VALUES (340, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e12 \\u0026nbsp; \\u0026nbsp;\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"\",\"results\":\"12\",\"scene\":\"\",\"sendTime\":\"12\",\"status\":1}', 'Error #01: 310:参数校验错误\n', 2, '2024-08-09 23:28:53', '2024-08-09 23:28:53', 0, '2024-08-09 23:28:53');
INSERT INTO `x_system_log_operate` VALUES (341, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e12 \\u0026nbsp; \\u0026nbsp;\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"\",\"results\":\"12\",\"scene\":\"\",\"sendTime\":\"12\",\"status\":1}', '', 1, '2024-08-09 23:31:08', '2024-08-09 23:31:08', 4, '2024-08-09 23:31:08');
INSERT INTO `x_system_log_operate` VALUES (342, 1, 'POST', '系统短信日志删除', '127.0.0.1', '/api/admin/system_log_sms/del', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Del-fm', '{\"id\":8}', '', 1, '2024-08-09 23:31:20', '2024-08-09 23:31:20', 3, '2024-08-09 23:31:20');
INSERT INTO `x_system_log_operate` VALUES (343, 1, 'POST', '系统短信日志删除', '127.0.0.1', '/api/admin/system_log_sms/del', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Del-fm', '{\"id\":7}', '', 1, '2024-08-09 23:31:23', '2024-08-09 23:31:23', 10, '2024-08-09 23:31:23');
INSERT INTO `x_system_log_operate` VALUES (344, 1, 'POST', '系统短信日志删除', '127.0.0.1', '/api/admin/system_log_sms/del', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Del-fm', '{\"id\":6}', '', 1, '2024-08-09 23:31:25', '2024-08-09 23:31:25', 3, '2024-08-09 23:31:25');
INSERT INTO `x_system_log_operate` VALUES (345, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"\",\"results\":\"1\",\"scene\":\"\",\"sendTime\":\"1\",\"status\":1}', '', 1, '2024-08-09 23:31:42', '2024-08-09 23:31:42', 2, '2024-08-09 23:31:42');
INSERT INTO `x_system_log_operate` VALUES (346, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e2\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"2\",\"results\":\"2\",\"scene\":\"2\",\"sendTime\":\"2\",\"status\":1}', '', 1, '2024-08-09 23:31:56', '2024-08-09 23:31:56', 3, '2024-08-09 23:31:56');
INSERT INTO `x_system_log_operate` VALUES (347, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e3\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"3\",\"results\":\"3\",\"scene\":\"3\",\"sendTime\":\"3\",\"status\":3}', '', 1, '2024-08-09 23:36:17', '2024-08-09 23:36:17', 31, '2024-08-09 23:36:17');
INSERT INTO `x_system_log_operate` VALUES (348, 1, 'POST', '系统短信日志删除', '127.0.0.1', '/api/admin/system_log_sms/del', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Del-fm', '{\"id\":9}', '', 1, '2024-08-09 23:37:34', '2024-08-09 23:37:34', 4, '2024-08-09 23:37:34');
INSERT INTO `x_system_log_operate` VALUES (349, 1, 'POST', '系统短信日志删除', '127.0.0.1', '/api/admin/system_log_sms/del', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Del-fm', '{\"id\":11}', '', 1, '2024-08-09 23:37:35', '2024-08-09 23:37:35', 4, '2024-08-09 23:37:35');
INSERT INTO `x_system_log_operate` VALUES (350, 1, 'POST', '系统短信日志删除', '127.0.0.1', '/api/admin/system_log_sms/del', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Del-fm', '{\"id\":10}', '', 1, '2024-08-09 23:37:38', '2024-08-09 23:37:38', 3, '2024-08-09 23:37:38');
INSERT INTO `x_system_log_operate` VALUES (351, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"1\",\"results\":\"1\",\"scene\":\"1\",\"sendTime\":\"1\",\"status\":2}', '', 1, '2024-08-09 23:38:01', '2024-08-09 23:38:46', 45466, '2024-08-09 23:38:46');
INSERT INTO `x_system_log_operate` VALUES (352, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e2\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"2\",\"results\":\"2\",\"scene\":2,\"sendTime\":2,\"status\":\"0\"}', '', 1, '2024-08-09 23:41:50', '2024-08-09 23:42:11', 21788, '2024-08-09 23:42:11');
INSERT INTO `x_system_log_operate` VALUES (353, 1, 'POST', '系统短信日志删除', '127.0.0.1', '/api/admin/system_log_sms/del', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Del-fm', '{\"id\":12}', '', 1, '2024-08-09 23:43:00', '2024-08-09 23:43:00', 6, '2024-08-09 23:43:00');
INSERT INTO `x_system_log_operate` VALUES (354, 1, 'POST', '系统短信日志删除', '127.0.0.1', '/api/admin/system_log_sms/del', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Del-fm', '{\"id\":13}', '', 1, '2024-08-09 23:43:02', '2024-08-09 23:43:02', 3, '2024-08-09 23:43:02');
INSERT INTO `x_system_log_operate` VALUES (355, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"1\",\"results\":\"1\",\"scene\":1,\"sendTime\":1,\"status\":1}', '', 1, '2024-08-09 23:43:12', '2024-08-09 23:43:29', 17423, '2024-08-09 23:43:29');
INSERT INTO `x_system_log_operate` VALUES (356, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"content\":\"\\u003cp\\u003e2\\u003c/p\\u003e\",\"id\":14,\"mobile\":\"2\",\"results\":\"2\",\"scene\":2,\"sendTime\":2,\"status\":3}', '', 1, '2024-08-09 23:43:46', '2024-08-09 23:43:51', 4603, '2024-08-09 23:43:51');
INSERT INTO `x_system_log_operate` VALUES (357, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"content\":\"\\u003cp\\u003e2\\u003c/p\\u003e\",\"id\":14,\"mobile\":\"2\",\"results\":\"2\",\"scene\":2,\"sendTime\":2,\"status\":3}', '', 1, '2024-08-09 23:44:03', '2024-08-09 23:44:55', 52054, '2024-08-09 23:44:55');
INSERT INTO `x_system_log_operate` VALUES (358, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"content\":\"\\u003cp\\u003e2\\u003c/p\\u003e\",\"id\":14,\"mobile\":\"2\",\"results\":\"2\",\"scene\":2,\"sendTime\":3,\"status\":3}', '', 1, '2024-08-09 23:45:17', '2024-08-09 23:46:05', 48775, '2024-08-09 23:46:05');
INSERT INTO `x_system_log_operate` VALUES (359, 1, 'POST', '系统短信日志删除', '127.0.0.1', '/api/admin/system_log_sms/del', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Del-fm', '{\"id\":14}', '', 1, '2024-08-09 23:53:47', '2024-08-09 23:53:47', 2, '2024-08-09 23:53:47');
INSERT INTO `x_system_log_operate` VALUES (360, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e11\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"11\",\"results\":\"1\",\"scene\":null,\"sendTime\":\"1\",\"status\":1}', '500:系统错误', 2, '2024-08-09 23:54:54', '2024-08-09 23:55:09', 14362, '2024-08-09 23:55:09');
INSERT INTO `x_system_log_operate` VALUES (361, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e11\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"11\",\"results\":\"1\",\"scene\":null,\"sendTime\":\"1\",\"status\":1}', '500:系统错误', 2, '2024-08-09 23:56:22', '2024-08-09 23:56:31', 9172, '2024-08-09 23:56:31');
INSERT INTO `x_system_log_operate` VALUES (362, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"1\",\"results\":\"1\",\"scene\":\"1\",\"sendTime\":\"1\",\"status\":1}', '', 1, '2024-08-10 00:49:27', '2024-08-10 00:49:30', 3294, '2024-08-10 00:49:30');
INSERT INTO `x_system_log_operate` VALUES (363, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e2\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"2\",\"results\":\"2\",\"scene\":\"2\",\"sendTime\":\"2\",\"status\":2}', '', 1, '2024-08-10 00:50:15', '2024-08-10 00:51:26', 71343, '2024-08-10 00:51:26');
INSERT INTO `x_system_log_operate` VALUES (364, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e2\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"2\",\"results\":\"2\",\"scene\":\"2\",\"sendTime\":\"2\",\"status\":2}', '', 1, '2024-08-10 01:09:01', '2024-08-10 01:10:02', 61525, '2024-08-10 01:10:02');
INSERT INTO `x_system_log_operate` VALUES (365, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"content\":\"\\u003cp\\u003e2\\u003c/p\\u003e\",\"id\":17,\"mobile\":\"2\",\"results\":\"2\",\"scene\":2,\"sendTime\":0,\"status\":2}', '', 1, '2024-08-10 01:13:39', '2024-08-10 01:14:05', 26027, '2024-08-10 01:14:05');
INSERT INTO `x_system_log_operate` VALUES (366, 1, 'POST', '系统短信日志删除', '127.0.0.1', '/api/admin/system_log_sms/del', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Del-fm', '{\"id\":17}', '', 1, '2024-08-10 01:14:13', '2024-08-10 01:14:13', 2, '2024-08-10 01:14:13');
INSERT INTO `x_system_log_operate` VALUES (367, 1, 'POST', '系统短信日志删除', '127.0.0.1', '/api/admin/system_log_sms/del', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Del-fm', '{\"id\":16}', '', 1, '2024-08-10 01:14:15', '2024-08-10 01:14:15', 3, '2024-08-10 01:14:15');
INSERT INTO `x_system_log_operate` VALUES (368, 1, 'POST', '系统短信日志删除', '127.0.0.1', '/api/admin/system_log_sms/del', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Del-fm', '{\"id\":15}', '', 1, '2024-08-10 01:14:17', '2024-08-10 01:14:17', 3, '2024-08-10 01:14:17');
INSERT INTO `x_system_log_operate` VALUES (369, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"1\",\"results\":\"1\",\"scene\":\"1\",\"sendTime\":\"1\",\"status\":2}', '', 1, '2024-08-10 01:15:09', '2024-08-10 01:15:54', 44578, '2024-08-10 01:15:54');
INSERT INTO `x_system_log_operate` VALUES (370, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"id\":18,\"mobile\":\"2\",\"results\":\"1\",\"scene\":\"2\",\"sendTime\":0,\"status\":2}', '', 1, '2024-08-10 01:24:57', '2024-08-10 01:25:01', 3801, '2024-08-10 01:25:01');
INSERT INTO `x_system_log_operate` VALUES (371, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e3\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"3\",\"results\":\"3\",\"scene\":\"3\",\"sendTime\":\"2\",\"status\":3}', '', 1, '2024-08-10 01:25:27', '2024-08-10 01:25:30', 2816, '2024-08-10 01:25:30');
INSERT INTO `x_system_log_operate` VALUES (372, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e4\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"4\",\"results\":\"4\",\"scene\":\"4\",\"sendTime\":\"4\",\"status\":3}', '', 1, '2024-08-10 01:25:58', '2024-08-10 01:27:33', 94998, '2024-08-10 01:27:33');
INSERT INTO `x_system_log_operate` VALUES (373, 1, 'POST', '系统短信日志删除', '127.0.0.1', '/api/admin/system_log_sms/del', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Del-fm', '{\"id\":20}', '', 1, '2024-08-10 01:33:31', '2024-08-10 01:33:31', 2, '2024-08-10 01:33:31');
INSERT INTO `x_system_log_operate` VALUES (374, 1, 'POST', '系统短信日志删除', '127.0.0.1', '/api/admin/system_log_sms/del', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Del-fm', '{\"id\":19}', '', 1, '2024-08-10 01:33:34', '2024-08-10 01:33:34', 3, '2024-08-10 01:33:34');
INSERT INTO `x_system_log_operate` VALUES (375, 1, 'POST', '系统短信日志删除', '127.0.0.1', '/api/admin/system_log_sms/del', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Del-fm', '{\"id\":18}', '', 1, '2024-08-10 01:33:36', '2024-08-10 01:33:36', 3, '2024-08-10 01:33:36');
INSERT INTO `x_system_log_operate` VALUES (376, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"1\",\"results\":\"1\",\"scene\":\"1\",\"send_time\":\"1\",\"status\":1}', '', 1, '2024-08-10 01:33:46', '2024-08-10 01:34:01', 15102, '2024-08-10 01:34:01');
INSERT INTO `x_system_log_operate` VALUES (377, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e2\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"2\",\"results\":\"2\",\"scene\":\"2\",\"send_time\":\"2\",\"status\":2}', '', 1, '2024-08-10 01:39:10', '2024-08-10 01:39:13', 2521, '2024-08-10 01:39:13');
INSERT INTO `x_system_log_operate` VALUES (378, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e3\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"3\",\"results\":\"3\",\"scene\":\"3\",\"send_time\":\"3\",\"status\":2}', '', 1, '2024-08-10 01:39:32', '2024-08-10 01:40:09', 37101, '2024-08-10 01:40:09');
INSERT INTO `x_system_log_operate` VALUES (379, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e33\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"33\",\"results\":\"3\",\"scene\":3,\"send_time\":3,\"status\":2}', '', 1, '2024-08-10 15:21:06', '2024-08-10 15:22:38', 92219, '2024-08-10 15:22:38');
INSERT INTO `x_system_log_operate` VALUES (380, 1, 'POST', '系统短信日志删除', '127.0.0.1', '/api/admin/system_log_sms/del', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Del-fm', '{\"id\":23}', '', 1, '2024-08-10 15:25:04', '2024-08-10 15:25:04', 13, '2024-08-10 15:25:04');
INSERT INTO `x_system_log_operate` VALUES (381, 1, 'POST', '系统短信日志删除', '127.0.0.1', '/api/admin/system_log_sms/del', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Del-fm', '{\"id\":22}', '', 1, '2024-08-10 15:25:07', '2024-08-10 15:25:07', 10, '2024-08-10 15:25:07');
INSERT INTO `x_system_log_operate` VALUES (382, 1, 'POST', '系统短信日志删除', '127.0.0.1', '/api/admin/system_log_sms/del', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Del-fm', '{\"id\":21}', '', 1, '2024-08-10 15:25:09', '2024-08-10 15:25:09', 3, '2024-08-10 15:25:09');
INSERT INTO `x_system_log_operate` VALUES (383, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e2\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"2\",\"results\":\"2\",\"scene\":2,\"send_time\":2,\"status\":1}', '', 1, '2024-08-10 15:25:26', '2024-08-10 15:26:17', 50735, '2024-08-10 15:26:17');
INSERT INTO `x_system_log_operate` VALUES (384, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"11\",\"results\":\"1\",\"scene\":\"1\",\"send_time\":1,\"status\":2}', '', 1, '2024-08-10 15:26:52', '2024-08-10 15:27:45', 53775, '2024-08-10 15:27:45');
INSERT INTO `x_system_log_operate` VALUES (385, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\",\"id\":\"\",\"mobile\":\"4\",\"results\":\"\",\"scene\":\"4\",\"send_time\":0,\"status\":\"\"}', '', 1, '2024-08-10 15:29:18', '2024-08-10 15:30:01', 43022, '2024-08-10 15:30:01');
INSERT INTO `x_system_log_operate` VALUES (386, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e5\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"5\",\"results\":\"5\",\"scene\":\"5\",\"send_time\":5,\"status\":2}', '', 1, '2024-08-10 17:10:58', '2024-08-10 17:11:30', 32400, '2024-08-10 17:11:30');
INSERT INTO `x_system_log_operate` VALUES (387, 1, 'POST', '系统短信日志删除', '127.0.0.1', '/api/admin/system_log_sms/del', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Del-fm', '{\"id\":28}', '', 1, '2024-08-10 17:11:41', '2024-08-10 17:11:41', 11, '2024-08-10 17:11:41');
INSERT INTO `x_system_log_operate` VALUES (388, 1, 'POST', '系统短信日志删除', '127.0.0.1', '/api/admin/system_log_sms/del', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Del-fm', '{\"id\":27}', '', 1, '2024-08-10 17:11:44', '2024-08-10 17:11:44', 5, '2024-08-10 17:11:44');
INSERT INTO `x_system_log_operate` VALUES (389, 1, 'POST', '系统短信日志删除', '127.0.0.1', '/api/admin/system_log_sms/del', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Del-fm', '{\"id\":26}', '', 1, '2024-08-10 17:11:48', '2024-08-10 17:11:48', 2, '2024-08-10 17:11:48');
INSERT INTO `x_system_log_operate` VALUES (390, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e4\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"4\",\"results\":\"4\",\"scene\":\"4\",\"send_time\":4,\"status\":2}', '', 1, '2024-08-10 17:43:47', '2024-08-10 17:50:38', 410879, '2024-08-10 17:50:38');
INSERT INTO `x_system_log_operate` VALUES (391, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e5\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"5\",\"results\":\"5\",\"scene\":\"5\",\"send_time\":5,\"status\":1}', '', 1, '2024-08-10 18:00:13', '2024-08-10 18:02:09', 116044, '2024-08-10 18:02:09');
INSERT INTO `x_system_log_operate` VALUES (392, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=beaeebf3739cf9a587e811b671fee427OWWsFO&content&create_timeEnd&create_timeStart&mobile=5&results&scene&send_time&status&update_timeEnd&update_timeStart', 'Error #01: 500:导出失败\n', 2, '2024-08-10 18:59:32', '2024-08-10 18:59:32', 5, '2024-08-10 18:59:32');
INSERT INTO `x_system_log_operate` VALUES (393, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=beaeebf3739cf9a587e811b671fee427OWWsFO', 'Error #01: 500:导出失败\n', 2, '2024-08-10 19:00:05', '2024-08-10 19:00:05', 6, '2024-08-10 19:00:05');
INSERT INTO `x_system_log_operate` VALUES (394, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=beaeebf3739cf9a587e811b671fee427OWWsFO', 'Error #01: 500:导出失败\n', 2, '2024-08-10 19:00:10', '2024-08-10 19:00:10', 30, '2024-08-10 19:00:10');
INSERT INTO `x_system_log_operate` VALUES (395, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=beaeebf3739cf9a587e811b671fee427OWWsFO', 'Error #01: 500:导出失败\n', 2, '2024-08-10 19:00:18', '2024-08-10 19:00:18', 3, '2024-08-10 19:00:18');
INSERT INTO `x_system_log_operate` VALUES (396, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=beaeebf3739cf9a587e811b671fee427OWWsFO', 'Error #01: 500:导出失败\n', 2, '2024-08-10 19:00:19', '2024-08-10 19:00:19', 3, '2024-08-10 19:00:19');
INSERT INTO `x_system_log_operate` VALUES (397, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=beaeebf3739cf9a587e811b671fee427OWWsFO', 'Error #01: 500:导出失败\n', 2, '2024-08-10 19:00:53', '2024-08-10 19:00:53', 3, '2024-08-10 19:00:53');
INSERT INTO `x_system_log_operate` VALUES (398, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=beaeebf3739cf9a587e811b671fee427OWWsFO', 'Error #01: 500:导出失败\n', 2, '2024-08-10 19:01:44', '2024-08-10 19:09:47', 483034, '2024-08-10 19:09:47');
INSERT INTO `x_system_log_operate` VALUES (399, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-08-11 02:54:23', '2024-08-11 02:54:23', 156, '2024-08-11 02:54:23');
INSERT INTO `x_system_log_operate` VALUES (400, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-08-11 02:55:39', '2024-08-11 02:55:39', 2, '2024-08-11 02:55:39');
INSERT INTO `x_system_log_operate` VALUES (401, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-08-11 02:57:14', '2024-08-11 02:57:14', 11, '2024-08-11 02:57:14');
INSERT INTO `x_system_log_operate` VALUES (402, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e6\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"6\",\"results\":\"6\",\"scene\":\"6\",\"send_time\":6,\"status\":4}', 'Error #01: 500:Struct field json tag don\'t match map key : send_time in obj\n', 2, '2024-08-11 02:58:26', '2024-08-11 03:01:17', 170605, '2024-08-11 03:01:17');
INSERT INTO `x_system_log_operate` VALUES (403, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\",\"id\":\"\",\"mobile\":\"\",\"results\":\"\",\"scene\":\"\",\"send_time\":\"\",\"status\":\"\"}', 'Error #01: 500:Map value type don\'t match struct field type\n', 2, '2024-08-11 03:01:17', '2024-08-11 03:01:20', 3354, '2024-08-11 03:01:20');
INSERT INTO `x_system_log_operate` VALUES (404, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\",\"id\":\"\",\"mobile\":\"\",\"results\":\"\",\"scene\":\"\",\"send_time\":\"\",\"status\":\"\"}', 'Error #01: 500:Map value type don\'t match struct field type\n', 2, '2024-08-11 03:04:22', '2024-08-11 03:05:11', 48865, '2024-08-11 03:05:11');
INSERT INTO `x_system_log_operate` VALUES (405, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":\"1\",\"results\":\"1\",\"scene\":\"1\",\"send_time\":1,\"status\":1}', '', 1, '2024-08-11 03:06:31', '2024-08-11 03:12:34', 363162, '2024-08-11 03:12:34');
INSERT INTO `x_system_log_operate` VALUES (406, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-08-11 03:13:17', '2024-08-11 03:13:17', 10, '2024-08-11 03:13:17');
INSERT INTO `x_system_log_operate` VALUES (407, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-08-11 03:31:18', '2024-08-11 03:31:18', 9, '2024-08-11 03:31:18');
INSERT INTO `x_system_log_operate` VALUES (408, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-08-11 03:38:14', '2024-08-11 03:38:14', 10, '2024-08-11 03:38:14');
INSERT INTO `x_system_log_operate` VALUES (409, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-08-11 03:42:39', '2024-08-11 03:42:39', 3, '2024-08-11 03:42:39');
INSERT INTO `x_system_log_operate` VALUES (410, 1, 'POST', '相册分类新增', '127.0.0.1', '/api/admin/common/album/cateAdd', 'x_admin/admin/common/album.albumHandler.cateAdd-fm', '{\"name\":\"1\",\"pid\":0,\"type\":10}', '', 1, '2024-08-11 03:43:51', '2024-08-11 03:43:51', 2, '2024-08-11 03:43:51');
INSERT INTO `x_system_log_operate` VALUES (411, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=57b4495eae5c178e3fd68bf8a0e50e36NC66Qt&content&create_timeEnd&create_timeStart&mobile&results&scene&send_time&status&update_timeEnd&update_timeStart', 'Error #01: 500:导出失败\n', 2, '2024-08-11 03:58:18', '2024-08-11 03:58:18', 5, '2024-08-11 03:58:18');
INSERT INTO `x_system_log_operate` VALUES (412, 1, 'GET', '错误收集error导出', '127.0.0.1', '/api/admin/monitor_web/ExportFile', 'x_admin/admin/monitor_web.MonitorWebHandler.ExportFile-fm', 'token=57b4495eae5c178e3fd68bf8a0e50e36NC66Qt&clientTimeEnd=&clientTimeStart=&createTimeEnd=&createTimeStart=&eventType=&message=&page=&stack=', '', 1, '2024-08-11 03:59:00', '2024-08-11 03:59:00', 13, '2024-08-11 03:59:00');
INSERT INTO `x_system_log_operate` VALUES (413, 1, 'GET', '错误收集error导出', '127.0.0.1', '/api/admin/monitor_web/ExportFile', 'x_admin/admin/monitor_web.MonitorWebHandler.ExportFile-fm', 'token=57b4495eae5c178e3fd68bf8a0e50e36NC66Qt&clientTimeEnd=&clientTimeStart=&createTimeEnd=&createTimeStart=&eventType=&message=&page=&stack=', '', 1, '2024-08-11 04:00:27', '2024-08-11 04:00:27', 19, '2024-08-11 04:00:27');
INSERT INTO `x_system_log_operate` VALUES (414, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=57b4495eae5c178e3fd68bf8a0e50e36NC66Qt&content&create_timeEnd&create_timeStart&mobile&results&scene&send_time&status&update_timeEnd&update_timeStart', 'Error #01: 500:导出失败\n', 2, '2024-08-11 04:01:17', '2024-08-11 04:01:17', 4, '2024-08-11 04:01:17');
INSERT INTO `x_system_log_operate` VALUES (415, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=57b4495eae5c178e3fd68bf8a0e50e36NC66Qt', 'Error #01: 500:导出失败\n', 2, '2024-08-11 04:01:31', '2024-08-11 04:01:31', 4, '2024-08-11 04:01:31');
INSERT INTO `x_system_log_operate` VALUES (416, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=57b4495eae5c178e3fd68bf8a0e50e36NC66Qt', 'Error #01: 500:导出失败\n', 2, '2024-08-11 04:02:03', '2024-08-11 04:02:57', 53918, '2024-08-11 04:02:57');
INSERT INTO `x_system_log_operate` VALUES (417, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=57b4495eae5c178e3fd68bf8a0e50e36NC66Qt', 'Error #01: 500:导出失败\n', 2, '2024-08-11 04:02:57', '2024-08-11 04:05:59', 182252, '2024-08-11 04:05:59');
INSERT INTO `x_system_log_operate` VALUES (418, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=57b4495eae5c178e3fd68bf8a0e50e36NC66Qt', 'Error #01: 500:导出失败\n', 2, '2024-08-11 04:06:14', '2024-08-11 04:06:54', 40243, '2024-08-11 04:06:54');
INSERT INTO `x_system_log_operate` VALUES (419, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=57b4495eae5c178e3fd68bf8a0e50e36NC66Qt', '', 1, '2024-08-11 04:39:03', '2024-08-11 04:39:03', 15, '2024-08-11 04:39:03');
INSERT INTO `x_system_log_operate` VALUES (420, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=57b4495eae5c178e3fd68bf8a0e50e36NC66Qt', '', 1, '2024-08-11 04:40:43', '2024-08-11 04:40:43', 17, '2024-08-11 04:40:43');
INSERT INTO `x_system_log_operate` VALUES (421, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=57b4495eae5c178e3fd68bf8a0e50e36NC66Qt', '', 1, '2024-08-11 04:42:14', '2024-08-11 04:42:14', 16, '2024-08-11 04:42:14');
INSERT INTO `x_system_log_operate` VALUES (422, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=57b4495eae5c178e3fd68bf8a0e50e36NC66Qt', '', 1, '2024-08-11 04:49:45', '2024-08-11 04:49:45', 19, '2024-08-11 04:49:45');
INSERT INTO `x_system_log_operate` VALUES (423, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=57b4495eae5c178e3fd68bf8a0e50e36NC66Qt', '', 1, '2024-08-11 04:52:12', '2024-08-11 04:58:08', 356164, '2024-08-11 04:58:08');
INSERT INTO `x_system_log_operate` VALUES (424, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=568e0cb5c8c35d770b59660d042353117meo88&content&create_timeEnd&create_timeStart&mobile&results&scene&send_time&status&update_timeEnd&update_timeStart', '', 1, '2024-08-11 19:02:43', '2024-08-11 19:02:47', 3557, '2024-08-11 19:02:47');
INSERT INTO `x_system_log_operate` VALUES (425, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=568e0cb5c8c35d770b59660d042353117meo88', '', 1, '2024-08-11 19:03:41', '2024-08-11 19:03:44', 3384, '2024-08-11 19:03:44');
INSERT INTO `x_system_log_operate` VALUES (426, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=568e0cb5c8c35d770b59660d042353117meo88', '', 1, '2024-08-11 19:04:29', '2024-08-11 19:05:26', 56526, '2024-08-11 19:05:26');
INSERT INTO `x_system_log_operate` VALUES (427, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=568e0cb5c8c35d770b59660d042353117meo88', '', 1, '2024-08-11 19:11:01', '2024-08-11 19:11:11', 10589, '2024-08-11 19:11:11');
INSERT INTO `x_system_log_operate` VALUES (428, 1, 'POST', '系统短信日志导入', '127.0.0.1', '/api/admin/system_log_sms/ImportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ImportFile-fm', '', '', 1, '2024-08-11 19:16:58', '2024-08-11 19:16:58', 1, '2024-08-11 19:16:58');
INSERT INTO `x_system_log_operate` VALUES (429, 1, 'POST', '系统短信日志导入', '127.0.0.1', '/api/admin/system_log_sms/ImportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ImportFile-fm', '', '', 1, '2024-08-11 19:17:25', '2024-08-11 19:17:25', 0, '2024-08-11 19:17:25');
INSERT INTO `x_system_log_operate` VALUES (430, 1, 'POST', '系统短信日志导入', '127.0.0.1', '/api/admin/system_log_sms/ImportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ImportFile-fm', '', '', 1, '2024-08-11 19:17:52', '2024-08-11 19:17:52', 0, '2024-08-11 19:17:52');
INSERT INTO `x_system_log_operate` VALUES (431, 1, 'POST', '系统短信日志导入', '127.0.0.1', '/api/admin/system_log_sms/ImportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ImportFile-fm', '', '', 1, '2024-08-11 19:18:33', '2024-08-11 19:20:36', 123770, '2024-08-11 19:20:36');
INSERT INTO `x_system_log_operate` VALUES (432, 1, 'GET', '错误收集error导出', '127.0.0.1', '/api/admin/monitor_web/ExportFile', 'x_admin/admin/monitor_web.MonitorWebHandler.ExportFile-fm', 'token=568e0cb5c8c35d770b59660d042353117meo88&clientTimeEnd=&clientTimeStart=&createTimeEnd=&createTimeStart=&eventType=&message=&page=&stack=', '', 1, '2024-08-11 19:21:02', '2024-08-11 19:21:11', 8976, '2024-08-11 19:21:11');
INSERT INTO `x_system_log_operate` VALUES (433, 1, 'POST', '系统短信日志导入', '127.0.0.1', '/api/admin/system_log_sms/ImportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ImportFile-fm', '', '', 1, '2024-08-11 19:22:21', '2024-08-11 19:22:26', 5049, '2024-08-11 19:22:26');
INSERT INTO `x_system_log_operate` VALUES (434, 1, 'POST', '系统短信日志导入', '127.0.0.1', '/api/admin/system_log_sms/ImportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ImportFile-fm', '', '', 1, '2024-08-12 01:00:16', '2024-08-12 01:00:19', 3002, '2024-08-12 01:00:19');
INSERT INTO `x_system_log_operate` VALUES (435, 1, 'POST', '系统短信日志导入', '127.0.0.1', '/api/admin/system_log_sms/ImportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ImportFile-fm', '', '', 1, '2024-08-12 01:03:25', '2024-08-12 01:20:26', 1021043, '2024-08-12 01:20:26');
INSERT INTO `x_system_log_operate` VALUES (436, 1, 'POST', '系统短信日志导入', '127.0.0.1', '/api/admin/system_log_sms/ImportFile', 'x_admin/admin/system_log_sms.SystemLogSmsHandler.ImportFile-fm', '', '', 1, '2024-08-12 01:37:06', '2024-08-12 01:37:22', 16470, '2024-08-12 01:37:22');
INSERT INTO `x_system_log_operate` VALUES (437, 1, 'POST', '系统短信日志导入', '127.0.0.1', '/api/admin/system_log_sms/ImportFile', 'x_admin/admin/system_log_sms.SystemLogSmsHandler.ImportFile-fm', '', '', 1, '2024-08-12 01:37:48', '2024-08-12 01:38:42', 53888, '2024-08-12 01:38:42');
INSERT INTO `x_system_log_operate` VALUES (438, 1, 'POST', '系统短信日志导入', '127.0.0.1', '/api/admin/system_log_sms/ImportFile', 'x_admin/admin/system_log_sms.SystemLogSmsHandler.ImportFile-fm', '', '', 1, '2024-08-12 01:38:51', '2024-08-12 01:38:54', 2370, '2024-08-12 01:38:54');
INSERT INTO `x_system_log_operate` VALUES (439, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=d281f50ba07e5531c5fc5e8fa96e9447v1pwKS&content&create_timeEnd&create_timeStart&mobile&results&scene&send_time&status&update_timeEnd&update_timeStart', '', 1, '2024-08-12 01:57:10', '2024-08-12 01:57:16', 5455, '2024-08-12 01:57:16');
INSERT INTO `x_system_log_operate` VALUES (440, 1, 'GET', '错误收集error导出', '127.0.0.1', '/api/admin/monitor_web/ExportFile', 'x_admin/admin/monitor_web.MonitorWebHandler.ExportFile-fm', 'token=d281f50ba07e5531c5fc5e8fa96e9447v1pwKS&clientTimeEnd=&clientTimeStart=&createTimeEnd=&createTimeStart=&eventType=&message=&page=&stack=', '', 1, '2024-08-12 02:00:55', '2024-08-12 02:00:55', 19, '2024-08-12 02:00:55');
INSERT INTO `x_system_log_operate` VALUES (441, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"content\":\"\\u003cp\\u003e2\\u003c/p\\u003e\",\"id\":38,\"mobile\":\"2\",\"results\":\"2\",\"scene\":2,\"send_time\":2,\"status\":2}', '', 1, '2024-08-12 14:43:06', '2024-08-12 14:43:06', 12, '2024-08-12 14:43:06');
INSERT INTO `x_system_log_operate` VALUES (442, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"content\":\"\\u003cp\\u003e2\\u003c/p\\u003e\",\"id\":38,\"mobile\":\"2\",\"results\":\"2\",\"scene\":2,\"send_time\":2,\"status\":2}', '', 1, '2024-08-12 14:43:39', '2024-08-12 14:43:39', 6, '2024-08-12 14:43:39');
INSERT INTO `x_system_log_operate` VALUES (443, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=f175e1edeb086b6cfc2c6ee3fa45192egFCKQv&content&create_timeEnd&create_timeStart&mobile&results&scene&send_time&status&update_timeEnd&update_timeStart', '', 1, '2024-08-12 14:43:53', '2024-08-12 14:43:56', 3235, '2024-08-12 14:43:56');
INSERT INTO `x_system_log_operate` VALUES (444, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=f175e1edeb086b6cfc2c6ee3fa45192egFCKQv&content&create_timeEnd&create_timeStart&mobile&results&scene&send_time&status&update_timeEnd&update_timeStart', '', 1, '2024-08-12 14:48:09', '2024-08-12 14:48:57', 48035, '2024-08-12 14:48:57');
INSERT INTO `x_system_log_operate` VALUES (445, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=f175e1edeb086b6cfc2c6ee3fa45192egFCKQv&', '', 1, '2024-08-12 14:53:34', '2024-08-12 14:53:38', 4094, '2024-08-12 14:53:38');
INSERT INTO `x_system_log_operate` VALUES (446, 1, 'GET', '错误收集error导出', '127.0.0.1', '/api/admin/monitor_web/ExportFile', 'x_admin/admin/monitor_web.MonitorWebHandler.ExportFile-fm', 'token=49f5bd20023d168fb0adb9fabf476dbcB7zeUv&clientTimeEnd=&clientTimeStart=&createTimeEnd=&createTimeStart=&eventType=&message=&page=&stack=', '', 1, '2024-08-12 15:04:08', '2024-08-12 15:04:08', 25, '2024-08-12 15:04:08');
INSERT INTO `x_system_log_operate` VALUES (447, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=49f5bd20023d168fb0adb9fabf476dbcB7zeUv&', '', 1, '2024-08-12 15:05:16', '2024-08-12 15:07:50', 154017, '2024-08-12 15:07:50');
INSERT INTO `x_system_log_operate` VALUES (448, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=49f5bd20023d168fb0adb9fabf476dbcB7zeUv&', '', 1, '2024-08-12 15:09:28', '2024-08-12 15:09:55', 27044, '2024-08-12 15:09:55');
INSERT INTO `x_system_log_operate` VALUES (449, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=49f5bd20023d168fb0adb9fabf476dbcB7zeUv&', '', 1, '2024-08-12 15:10:34', '2024-08-12 15:10:38', 4246, '2024-08-12 15:10:38');
INSERT INTO `x_system_log_operate` VALUES (450, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=49f5bd20023d168fb0adb9fabf476dbcB7zeUv&', '', 1, '2024-08-12 15:14:20', '2024-08-12 15:14:51', 31857, '2024-08-12 15:14:51');
INSERT INTO `x_system_log_operate` VALUES (451, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=49f5bd20023d168fb0adb9fabf476dbcB7zeUv&', '', 1, '2024-08-12 15:46:52', '2024-08-12 15:47:00', 8289, '2024-08-12 15:47:00');
INSERT INTO `x_system_log_operate` VALUES (452, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=49f5bd20023d168fb0adb9fabf476dbcB7zeUv&', '', 1, '2024-08-12 16:00:46', '2024-08-12 16:02:52', 125948, '2024-08-12 16:02:52');
INSERT INTO `x_system_log_operate` VALUES (453, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=49f5bd20023d168fb0adb9fabf476dbcB7zeUv&', '', 1, '2024-08-12 16:03:00', '2024-08-12 16:03:05', 4329, '2024-08-12 16:03:05');
INSERT INTO `x_system_log_operate` VALUES (454, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=49f5bd20023d168fb0adb9fabf476dbcB7zeUv&', '', 1, '2024-08-12 16:03:48', '2024-08-12 16:03:51', 3007, '2024-08-12 16:03:51');
INSERT INTO `x_system_log_operate` VALUES (455, 1, 'GET', '错误收集error导出', '127.0.0.1', '/api/admin/monitor_web/ExportFile', 'x_admin/admin/monitor_web.MonitorWebHandler.ExportFile-fm', 'token=4fd909f7b9234741a963df802da6e9819NiLXy&clientTimeEnd=&clientTimeStart=&createTimeEnd=&createTimeStart=&eventType=&message=&page=&stack=', '', 1, '2024-08-12 16:37:01', '2024-08-12 16:37:01', 21, '2024-08-12 16:37:01');
INSERT INTO `x_system_log_operate` VALUES (456, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=ce8b9b89dacc25748f7c8b33f3af55cdyEUfGd&', '', 1, '2024-08-12 17:12:37', '2024-08-12 17:12:42', 5372, '2024-08-12 17:12:42');
INSERT INTO `x_system_log_operate` VALUES (457, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=ce8b9b89dacc25748f7c8b33f3af55cdyEUfGd&', '', 1, '2024-08-12 17:16:02', '2024-08-12 17:16:17', 15744, '2024-08-12 17:16:17');
INSERT INTO `x_system_log_operate` VALUES (458, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=ce8b9b89dacc25748f7c8b33f3af55cdyEUfGd&', '', 1, '2024-08-12 17:18:42', '2024-08-12 17:19:22', 39711, '2024-08-12 17:19:22');
INSERT INTO `x_system_log_operate` VALUES (459, 1, 'POST', '系统短信日志删除', '127.0.0.1', '/api/admin/system_log_sms/del', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Del-fm', '{\"id\":42}', '', 1, '2024-08-12 18:08:30', '2024-08-12 18:08:30', 13, '2024-08-12 18:08:30');
INSERT INTO `x_system_log_operate` VALUES (460, 1, 'POST', '系统短信日志删除', '127.0.0.1', '/api/admin/system_log_sms/del', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Del-fm', '{\"id\":41}', '', 1, '2024-08-12 18:08:32', '2024-08-12 18:08:32', 3, '2024-08-12 18:08:32');
INSERT INTO `x_system_log_operate` VALUES (461, 1, 'POST', '系统短信日志删除', '127.0.0.1', '/api/admin/system_log_sms/del', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Del-fm', '{\"id\":40}', '', 1, '2024-08-12 18:08:34', '2024-08-12 18:08:34', 2, '2024-08-12 18:08:34');
INSERT INTO `x_system_log_operate` VALUES (462, 1, 'POST', '系统短信日志删除', '127.0.0.1', '/api/admin/system_log_sms/del', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Del-fm', '{\"id\":39}', '', 1, '2024-08-12 18:08:39', '2024-08-12 18:08:39', 3, '2024-08-12 18:08:39');
INSERT INTO `x_system_log_operate` VALUES (463, 1, 'POST', '系统短信日志删除', '127.0.0.1', '/api/admin/system_log_sms/del', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Del-fm', '{\"id\":38}', '', 1, '2024-08-12 18:08:41', '2024-08-12 18:08:41', 11, '2024-08-12 18:08:41');
INSERT INTO `x_system_log_operate` VALUES (464, 1, 'POST', '系统短信日志删除', '127.0.0.1', '/api/admin/system_log_sms/del', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Del-fm', '{\"id\":37}', '', 1, '2024-08-12 18:08:44', '2024-08-12 18:08:44', 12, '2024-08-12 18:08:44');
INSERT INTO `x_system_log_operate` VALUES (465, 1, 'POST', '错误收集error删除', '127.0.0.1', '/api/admin/monitor_web/del', 'x_admin/admin/monitor_web.MonitorWebHandler.Del-fm', '{\"id\":36}', '', 1, '2024-08-12 22:49:12', '2024-08-12 22:49:12', 14, '2024-08-12 22:49:12');
INSERT INTO `x_system_log_operate` VALUES (466, 1, 'POST', '错误收集error删除', '127.0.0.1', '/api/admin/monitor_web/del', 'x_admin/admin/monitor_web.MonitorWebHandler.Del-fm', '{\"id\":35}', '', 1, '2024-08-12 22:49:14', '2024-08-12 22:49:14', 10, '2024-08-12 22:49:14');
INSERT INTO `x_system_log_operate` VALUES (467, 1, 'POST', '错误收集error删除', '127.0.0.1', '/api/admin/monitor_web/del', 'x_admin/admin/monitor_web.MonitorWebHandler.Del-fm', '{\"id\":34}', '', 1, '2024-08-12 22:49:20', '2024-08-12 22:49:20', 10, '2024-08-12 22:49:20');
INSERT INTO `x_system_log_operate` VALUES (468, 1, 'POST', '错误收集error删除', '127.0.0.1', '/api/admin/monitor_web/del', 'x_admin/admin/monitor_web.MonitorWebHandler.Del-fm', '{\"id\":33}', '', 1, '2024-08-12 22:49:22', '2024-08-12 22:49:22', 10, '2024-08-12 22:49:22');
INSERT INTO `x_system_log_operate` VALUES (469, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=ed6e3df505acf64d775d7e2743a9485fLaQgux&', '', 1, '2024-08-13 01:18:49', '2024-08-13 01:18:49', 21, '2024-08-13 01:18:49');
INSERT INTO `x_system_log_operate` VALUES (470, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-08-13 12:16:14', '2024-08-13 12:16:14', 3, '2024-08-13 12:16:14');
INSERT INTO `x_system_log_operate` VALUES (471, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=14b29187d670995a458f54dc98227d89LcaHKq&', '', 1, '2024-08-13 12:19:14', '2024-08-13 12:19:14', 19, '2024-08-13 12:19:14');
INSERT INTO `x_system_log_operate` VALUES (472, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-08-13 12:31:08', '2024-08-13 12:31:08', 12, '2024-08-13 12:31:08');
INSERT INTO `x_system_log_operate` VALUES (473, 1, 'POST', '错误收集error删除', '127.0.0.1', '/api/admin/monitor_web/del', 'x_admin/admin/monitor_web.MonitorWebHandler.Del-fm', '{\"id\":19}', '', 1, '2024-08-13 12:41:47', '2024-08-13 12:41:47', 14, '2024-08-13 12:41:47');
INSERT INTO `x_system_log_operate` VALUES (474, 1, 'POST', '错误收集error删除', '127.0.0.1', '/api/admin/monitor_web/del', 'x_admin/admin/monitor_web.MonitorWebHandler.Del-fm', '{\"id\":21}', '', 1, '2024-08-13 12:41:49', '2024-08-13 12:41:49', 5, '2024-08-13 12:41:49');
INSERT INTO `x_system_log_operate` VALUES (475, 1, 'POST', '错误收集error删除', '127.0.0.1', '/api/admin/monitor_web/del', 'x_admin/admin/monitor_web.MonitorWebHandler.Del-fm', '{\"id\":20}', '', 1, '2024-08-13 12:41:53', '2024-08-13 12:41:53', 2, '2024-08-13 12:41:53');
INSERT INTO `x_system_log_operate` VALUES (476, 1, 'POST', '错误收集error删除', '127.0.0.1', '/api/admin/monitor_web/del', 'x_admin/admin/monitor_web.MonitorWebHandler.Del-fm', '{\"id\":24}', '', 1, '2024-08-13 12:41:55', '2024-08-13 12:41:55', 11, '2024-08-13 12:41:55');
INSERT INTO `x_system_log_operate` VALUES (477, 1, 'POST', '错误收集error删除', '127.0.0.1', '/api/admin/monitor_web/del', 'x_admin/admin/monitor_web.MonitorWebHandler.Del-fm', '{\"id\":27}', '', 1, '2024-08-13 12:41:57', '2024-08-13 12:41:57', 10, '2024-08-13 12:41:57');
INSERT INTO `x_system_log_operate` VALUES (478, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=14b29187d670995a458f54dc98227d89LcaHKq&', '', 1, '2024-08-13 13:58:37', '2024-08-13 13:58:37', 18, '2024-08-13 13:58:37');
INSERT INTO `x_system_log_operate` VALUES (479, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=14b29187d670995a458f54dc98227d89LcaHKq&', '', 1, '2024-08-13 14:28:17', '2024-08-13 14:28:17', 15, '2024-08-13 14:28:17');
INSERT INTO `x_system_log_operate` VALUES (480, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":1,\"results\":\"1\",\"scene\":\"1\",\"send_time\":\"2024-08-13 07:03:31\",\"status\":2}', '', 1, '2024-08-13 19:03:32', '2024-08-13 19:03:32', 3, '2024-08-13 19:03:32');
INSERT INTO `x_system_log_operate` VALUES (481, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e2\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":2,\"results\":\"2\",\"scene\":\"2\",\"send_time\":\"2024-08-14 12:00:00\",\"status\":2}', '', 1, '2024-08-13 19:04:49', '2024-08-13 19:04:49', 3, '2024-08-13 19:04:49');
INSERT INTO `x_system_log_operate` VALUES (482, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e3\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":3,\"results\":\"3\",\"scene\":\"3\",\"send_time\":\"2024-08-12 12:00:00\",\"status\":3}', '', 1, '2024-08-13 19:23:50', '2024-08-13 19:37:36', 826325, '2024-08-13 19:37:36');
INSERT INTO `x_system_log_operate` VALUES (483, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e2\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":2,\"results\":\"2\",\"scene\":\"2\",\"send_time\":\"2024-08-28 12:00:00\",\"status\":2}', '', 1, '2024-08-13 22:18:28', '2024-08-13 22:19:35', 67001, '2024-08-13 22:19:35');
INSERT INTO `x_system_log_operate` VALUES (484, 1, 'POST', '系统短信日志删除', '127.0.0.1', '/api/admin/system_log_sms/del', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Del-fm', '{\"id\":63}', '', 1, '2024-08-13 22:20:36', '2024-08-13 22:20:36', 15, '2024-08-13 22:20:36');
INSERT INTO `x_system_log_operate` VALUES (485, 1, 'POST', '系统短信日志删除', '127.0.0.1', '/api/admin/system_log_sms/del', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Del-fm', '{\"id\":65}', '', 1, '2024-08-13 22:20:39', '2024-08-13 22:20:39', 11, '2024-08-13 22:20:39');
INSERT INTO `x_system_log_operate` VALUES (486, 1, 'POST', '系统短信日志删除', '127.0.0.1', '/api/admin/system_log_sms/del', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Del-fm', '{\"id\":67}', '', 1, '2024-08-13 22:20:40', '2024-08-13 22:20:40', 3, '2024-08-13 22:20:40');
INSERT INTO `x_system_log_operate` VALUES (487, 1, 'POST', '系统短信日志删除', '127.0.0.1', '/api/admin/system_log_sms/del', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Del-fm', '{\"id\":66}', '', 1, '2024-08-13 22:20:43', '2024-08-13 22:20:43', 2, '2024-08-13 22:20:43');
INSERT INTO `x_system_log_operate` VALUES (488, 1, 'POST', '系统短信日志删除', '127.0.0.1', '/api/admin/system_log_sms/del', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Del-fm', '{\"id\":64}', '', 1, '2024-08-13 22:20:45', '2024-08-13 22:20:45', 11, '2024-08-13 22:20:45');
INSERT INTO `x_system_log_operate` VALUES (489, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":1,\"results\":\"1\",\"scene\":\"1\",\"send_time\":\"2024-08-12 12:00:00\",\"status\":2}', '', 1, '2024-08-13 22:20:58', '2024-08-13 22:21:06', 8101, '2024-08-13 22:21:06');
INSERT INTO `x_system_log_operate` VALUES (490, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":1,\"results\":\"1\",\"scene\":\"1\",\"send_time\":\"2024-08-14 12:00:00\",\"status\":2}', '', 1, '2024-08-13 22:25:26', '2024-08-13 22:25:43', 16845, '2024-08-13 22:25:43');
INSERT INTO `x_system_log_operate` VALUES (491, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":1,\"results\":\"1\",\"scene\":\"1\",\"send_time\":\"2024-08-12 12:00:00\",\"status\":1}', '', 1, '2024-08-13 22:25:58', '2024-08-13 22:26:47', 48720, '2024-08-13 22:26:47');
INSERT INTO `x_system_log_operate` VALUES (492, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":1,\"results\":\"1\",\"scene\":\"1\",\"send_time\":\"2024-08-12 12:00:00\",\"status\":2}', '', 1, '2024-08-13 22:34:40', '2024-08-13 22:35:08', 27344, '2024-08-13 22:35:08');
INSERT INTO `x_system_log_operate` VALUES (493, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e2\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":2,\"results\":\"2\",\"scene\":\"2\",\"send_time\":\"2024-08-20 12:00:00\",\"status\":2}', '', 1, '2024-08-13 22:36:58', '2024-08-13 22:37:58', 60710, '2024-08-13 22:37:58');
INSERT INTO `x_system_log_operate` VALUES (494, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e3\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":3,\"results\":\"3\",\"scene\":\"3\",\"send_time\":\"2024-08-23 12:00:00\",\"status\":2}', '', 1, '2024-08-13 22:52:42', '2024-08-13 22:52:46', 3942, '2024-08-13 22:52:46');
INSERT INTO `x_system_log_operate` VALUES (495, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"content\":\"\\u003cp\\u003e3\\u003c/p\\u003e\",\"id\":\"\",\"mobile\":3,\"results\":\"3\",\"scene\":\"3\",\"send_time\":\"2024-08-21 12:00:00\",\"status\":2}', 'reflect: call of reflect.Value.Type on zero Value', 2, '2024-08-13 23:25:28', '2024-08-13 23:25:42', 13454, '2024-08-13 23:25:42');
INSERT INTO `x_system_log_operate` VALUES (496, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"Content\":\"\\u003cp\\u003e2\\u003c/p\\u003e\",\"Id\":\"2\",\"Mobile\":2,\"Results\":\"2\",\"Scene\":\"2\",\"SendTime\":\"2024-08-14 12:00:00\",\"Status\":3}', '', 1, '2024-08-13 23:34:26', '2024-08-13 23:34:53', 27646, '2024-08-13 23:34:53');
INSERT INTO `x_system_log_operate` VALUES (497, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"Content\":\"\\u003cp\\u003e3\\u003c/p\\u003e\",\"Id\":\"3\",\"Mobile\":3,\"Results\":\"3\",\"Scene\":\"3\",\"SendTime\":\"2024-08-13 11:35:52\",\"Status\":1}', 'Error #01: 310:参数校验错误\n', 2, '2024-08-13 23:44:16', '2024-08-13 23:44:16', 1, '2024-08-13 23:44:16');
INSERT INTO `x_system_log_operate` VALUES (498, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"Content\":\"\\u003cp\\u003e3\\u003c/p\\u003e\",\"Id\":\"3\",\"Mobile\":3,\"Results\":\"3\",\"Scene\":\"4\",\"SendTime\":\"2024-08-13 11:35:52\",\"Status\":1}', 'Error #01: 310:参数校验错误\n', 2, '2024-08-13 23:45:39', '2024-08-13 23:45:39', 0, '2024-08-13 23:45:39');
INSERT INTO `x_system_log_operate` VALUES (499, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"Content\":\"\\u003cp\\u003e3\\u003c/p\\u003e\",\"Id\":\"3\",\"Mobile\":3,\"Results\":\"3\",\"Scene\":\"4\",\"SendTime\":\"2024-08-13 11:35:52\",\"Status\":1}', 'Error #01: 310:参数校验错误\n', 2, '2024-08-13 23:46:42', '2024-08-13 23:46:42', 0, '2024-08-13 23:46:42');
INSERT INTO `x_system_log_operate` VALUES (500, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"Content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"Id\":\"1\",\"Mobile\":1,\"Results\":\"1\",\"Scene\":1,\"SendTime\":\"\",\"Status\":\"\"}', 'Error #01: 310:参数校验错误\n', 2, '2024-08-13 23:46:57', '2024-08-13 23:46:57', 0, '2024-08-13 23:46:57');
INSERT INTO `x_system_log_operate` VALUES (501, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"Content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"Id\":\"1\",\"Mobile\":1,\"Results\":\"1\",\"Scene\":1,\"SendTime\":\"\",\"Status\":\"\"}', 'Error #01: 310:参数校验错误\n', 2, '2024-08-13 23:47:36', '2024-08-13 23:47:36', 0, '2024-08-13 23:47:36');
INSERT INTO `x_system_log_operate` VALUES (502, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"Content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"Id\":\"1\",\"Mobile\":\"1\",\"Results\":\"\",\"Scene\":1,\"SendTime\":\"\",\"Status\":\"\"}', 'Error #01: 310:参数校验错误\n', 2, '2024-08-13 23:47:48', '2024-08-13 23:47:48', 0, '2024-08-13 23:47:48');
INSERT INTO `x_system_log_operate` VALUES (503, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"Content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"Id\":\"1\",\"Mobile\":\"1\",\"Results\":null,\"Scene\":1,\"SendTime\":null,\"Status\":null}', '', 1, '2024-08-13 23:48:24', '2024-08-13 23:48:45', 21419, '2024-08-13 23:48:45');
INSERT INTO `x_system_log_operate` VALUES (504, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"Content\":null,\"Id\":null,\"Mobile\":null,\"Results\":null,\"Scene\":null,\"SendTime\":null,\"Status\":null}', '', 1, '2024-08-13 23:53:57', '2024-08-13 23:54:01', 3980, '2024-08-13 23:54:01');
INSERT INTO `x_system_log_operate` VALUES (505, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"Content\":null,\"Id\":null,\"Mobile\":null,\"Results\":null,\"Scene\":null,\"SendTime\":null,\"Status\":null}', '', 1, '2024-08-13 23:54:29', '2024-08-13 23:54:32', 3735, '2024-08-13 23:54:32');
INSERT INTO `x_system_log_operate` VALUES (506, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"Content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"Id\":\"1\",\"Mobile\":\"1\",\"Results\":\"1\",\"Scene\":1,\"SendTime\":\"2024-08-14 12:00:00\",\"Status\":2}', '', 1, '2024-08-13 23:56:44', '2024-08-13 23:57:06', 21779, '2024-08-13 23:57:06');
INSERT INTO `x_system_log_operate` VALUES (507, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"Content\":\"\\u003cp\\u003e2\\u003c/p\\u003e\",\"Id\":\"2\",\"Mobile\":\"2\",\"Results\":\"2\",\"Scene\":2,\"SendTime\":\"2024-08-12 12:00:00\",\"Status\":2}', '', 1, '2024-08-13 23:58:07', '2024-08-13 23:59:26', 79176, '2024-08-13 23:59:26');
INSERT INTO `x_system_log_operate` VALUES (508, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=0ef82602ce5eab30a3cb66eea6eadfa5exATip&', '', 1, '2024-08-14 00:01:24', '2024-08-14 00:01:24', 13, '2024-08-14 00:01:24');
INSERT INTO `x_system_log_operate` VALUES (509, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=0ef82602ce5eab30a3cb66eea6eadfa5exATip&', '', 1, '2024-08-14 00:03:48', '2024-08-14 00:03:48', 16, '2024-08-14 00:03:48');
INSERT INTO `x_system_log_operate` VALUES (510, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"Content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"Id\":\"1\",\"Mobile\":\"1\",\"Results\":\"1\",\"Scene\":1,\"SendTime\":\"2024-08-14 12:11:57\",\"Status\":2}', '', 1, '2024-08-14 00:12:01', '2024-08-14 00:12:05', 4097, '2024-08-14 00:12:05');
INSERT INTO `x_system_log_operate` VALUES (511, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"Content\":\"\",\"Id\":78,\"Mobile\":\"\",\"Results\":\"\",\"Scene\":0,\"SendTime\":null,\"Status\":0}', '', 1, '2024-08-14 00:14:14', '2024-08-14 00:14:14', 37, '2024-08-14 00:14:14');
INSERT INTO `x_system_log_operate` VALUES (512, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"Content\":\"\\u003cp\\u003e2\\u003c/p\\u003e\",\"Id\":\"2\",\"Mobile\":\"2\",\"Results\":null,\"Scene\":2,\"SendTime\":null,\"Status\":null}', '500:系统错误', 2, '2024-08-14 00:27:41', '2024-08-14 00:27:45', 4623, '2024-08-14 00:27:45');
INSERT INTO `x_system_log_operate` VALUES (513, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"Content\":\"\\u003cp\\u003e2\\u003c/p\\u003e\",\"Id\":\"2\",\"Mobile\":\"2\",\"Results\":null,\"Scene\":2,\"SendTime\":null,\"Status\":null}', '500:系统错误', 2, '2024-08-14 00:28:27', '2024-08-14 00:28:33', 5868, '2024-08-14 00:28:33');
INSERT INTO `x_system_log_operate` VALUES (514, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"Content\":\"\\u003cp\\u003e2\\u003c/p\\u003e\",\"Id\":\"2\",\"Mobile\":\"2\",\"Results\":null,\"Scene\":2,\"SendTime\":null,\"Status\":null}', '500:系统错误', 2, '2024-08-14 00:28:42', '2024-08-14 00:28:50', 7799, '2024-08-14 00:28:50');
INSERT INTO `x_system_log_operate` VALUES (515, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"Content\":\"\\u003cp\\u003e2\\u003c/p\\u003e\",\"Id\":\"2\",\"Mobile\":\"2\",\"Results\":null,\"Scene\":2,\"SendTime\":null,\"Status\":null}', '500:系统错误', 2, '2024-08-14 00:29:10', '2024-08-14 00:29:27', 17714, '2024-08-14 00:29:27');
INSERT INTO `x_system_log_operate` VALUES (516, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"Content\":\"\\u003cp\\u003e2\\u003c/p\\u003e\",\"Id\":\"2\",\"Mobile\":\"2\",\"Results\":null,\"Scene\":2,\"SendTime\":null,\"Status\":null}', '500:系统错误', 2, '2024-08-14 00:30:25', '2024-08-14 00:30:39', 14335, '2024-08-14 00:30:39');
INSERT INTO `x_system_log_operate` VALUES (517, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"Content\":\"\\u003cp\\u003e2\\u003c/p\\u003e\",\"Id\":\"2\",\"Mobile\":\"2\",\"Results\":null,\"Scene\":2,\"SendTime\":null,\"Status\":null}', '500:系统错误', 2, '2024-08-14 00:30:48', '2024-08-14 00:32:02', 73697, '2024-08-14 00:32:02');
INSERT INTO `x_system_log_operate` VALUES (518, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"Content\":\"\\u003cp\\u003e2\\u003c/p\\u003e\",\"Id\":\"2\",\"Mobile\":\"2\",\"Results\":null,\"Scene\":2,\"SendTime\":null,\"Status\":null}', '500:系统错误', 2, '2024-08-14 00:33:32', '2024-08-14 00:36:43', 190961, '2024-08-14 00:36:43');
INSERT INTO `x_system_log_operate` VALUES (519, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"Content\":\"\\u003cp\\u003e2\\u003c/p\\u003e\",\"Id\":\"2\",\"Mobile\":\"22\",\"Results\":null,\"Scene\":2,\"SendTime\":null,\"Status\":null}', '500:系统错误', 2, '2024-08-14 00:37:51', '2024-08-14 00:38:11', 20418, '2024-08-14 00:38:11');
INSERT INTO `x_system_log_operate` VALUES (520, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"Content\":\"\\u003cp\\u003e2\\u003c/p\\u003e\",\"Id\":\"2\",\"Mobile\":\"22\",\"Results\":null,\"Scene\":2,\"SendTime\":null,\"Status\":null}', '500:系统错误', 2, '2024-08-14 00:39:58', '2024-08-14 00:40:24', 25595, '2024-08-14 00:40:24');
INSERT INTO `x_system_log_operate` VALUES (521, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"Content\":\"\\u003cp\\u003e2\\u003c/p\\u003e\",\"Id\":\"2\",\"Mobile\":\"22\",\"Results\":null,\"Scene\":2,\"SendTime\":null,\"Status\":null}', '500:系统错误', 2, '2024-08-14 00:46:43', '2024-08-14 00:46:53', 9794, '2024-08-14 00:46:53');
INSERT INTO `x_system_log_operate` VALUES (522, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"Content\":\"\\u003cp\\u003e2\\u003c/p\\u003e\",\"Id\":\"2\",\"Mobile\":\"22\",\"Results\":null,\"Scene\":2,\"SendTime\":null,\"Status\":null}', '500:系统错误', 2, '2024-08-14 00:46:54', '2024-08-14 00:47:49', 54230, '2024-08-14 00:53:51');
INSERT INTO `x_system_log_operate` VALUES (523, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"Content\":\"\\u003cp\\u003e2\\u003c/p\\u003e\",\"Id\":\"2\",\"Mobile\":\"22\",\"Results\":null,\"Scene\":2,\"SendTime\":null,\"Status\":null}', '500:系统错误', 2, '2024-08-14 00:53:58', '2024-08-14 00:56:11', 132896, '2024-08-14 00:56:11');
INSERT INTO `x_system_log_operate` VALUES (524, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"Content\":\"\\u003cp\\u003e2\\u003c/p\\u003e\",\"Id\":\"2\",\"Mobile\":\"22\",\"Results\":null,\"Scene\":2,\"SendTime\":null,\"Status\":null}', '', 1, '2024-08-14 00:57:42', '2024-08-14 00:58:00', 17352, '2024-08-14 00:58:00');
INSERT INTO `x_system_log_operate` VALUES (525, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"Content\":\"\\u003cp\\u003e11\\u003c/p\\u003e\",\"Id\":\"1\",\"Mobile\":\"11\",\"Results\":\"11\",\"Scene\":11,\"SendTime\":\"2024-08-14 12:58:37\",\"Status\":2}', '', 1, '2024-08-14 00:58:43', '2024-08-14 00:58:45', 2650, '2024-08-14 00:58:45');
INSERT INTO `x_system_log_operate` VALUES (526, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"Content\":\"\\u003cp\\u003e11\\u003c/p\\u003e\",\"Id\":83,\"Mobile\":\"11\",\"Results\":\"11\",\"Scene\":11,\"SendTime\":null,\"Status\":2}', '', 1, '2024-08-14 00:59:11', '2024-08-14 00:59:11', 18, '2024-08-14 00:59:11');
INSERT INTO `x_system_log_operate` VALUES (527, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"Content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"Id\":\"1\",\"Mobile\":\"1\",\"Results\":null,\"Scene\":1,\"SendTime\":null,\"Status\":null}', '500:系统错误', 2, '2024-08-14 01:13:31', '2024-08-14 01:13:37', 5411, '2024-08-14 01:13:37');
INSERT INTO `x_system_log_operate` VALUES (528, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"Content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"Id\":\"1\",\"Mobile\":\"1\",\"Results\":null,\"Scene\":1,\"SendTime\":\"2024-08-12 12:00:00\",\"Status\":null}', '500:系统错误', 2, '2024-08-14 01:13:45', '2024-08-14 01:13:50', 4543, '2024-08-14 01:13:50');
INSERT INTO `x_system_log_operate` VALUES (529, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"Content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"Id\":\"1\",\"Mobile\":\"1\",\"Results\":null,\"Scene\":1,\"SendTime\":\"2024-08-12 12:00:00\",\"Status\":2}', '500:系统错误', 2, '2024-08-14 01:13:59', '2024-08-14 01:14:06', 6465, '2024-08-14 01:14:06');
INSERT INTO `x_system_log_operate` VALUES (530, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"Content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"Id\":\"1\",\"Mobile\":\"1\",\"Results\":\"1\",\"Scene\":1,\"SendTime\":\"2024-08-12 12:00:00\",\"Status\":2}', '', 1, '2024-08-14 01:15:53', '2024-08-14 01:17:03', 70112, '2024-08-14 01:17:03');
INSERT INTO `x_system_log_operate` VALUES (531, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"Content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"Id\":84,\"Mobile\":\"1\",\"Results\":\"1\",\"Scene\":1,\"SendTime\":\"2024-08-12 12:00:00\",\"Status\":2}', '', 1, '2024-08-14 01:17:19', '2024-08-14 01:17:19', 4, '2024-08-14 01:17:19');
INSERT INTO `x_system_log_operate` VALUES (532, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"Content\":null,\"Id\":null,\"Mobile\":null,\"Results\":null,\"Scene\":null,\"SendTime\":null,\"Status\":null}', '500:系统错误', 2, '2024-08-14 01:18:03', '2024-08-14 01:19:05', 62291, '2024-08-14 01:19:05');
INSERT INTO `x_system_log_operate` VALUES (533, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"Content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"Id\":\"1\",\"Mobile\":\"1\",\"Results\":\"1\",\"Scene\":1,\"SendTime\":\"2024-08-14 01:19:00\",\"Status\":1}', '', 1, '2024-08-14 01:19:05', '2024-08-14 01:19:26', 20324, '2024-08-14 01:19:26');
INSERT INTO `x_system_log_operate` VALUES (534, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"Content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"Id\":85,\"Mobile\":\"1\",\"Results\":\"1\",\"Scene\":1,\"SendTime\":\"2024-08-14 01:19:00\",\"Status\":2}', '', 1, '2024-08-14 01:22:20', '2024-08-14 01:22:20', 30, '2024-08-14 01:22:20');
INSERT INTO `x_system_log_operate` VALUES (535, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"Content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"Id\":85,\"Mobile\":\"1\",\"Results\":\"1\",\"Scene\":1,\"SendTime\":\"2024-08-18 01:19:00\",\"Status\":1}', '', 1, '2024-08-14 01:22:28', '2024-08-14 01:22:28', 10, '2024-08-14 01:22:28');
INSERT INTO `x_system_log_operate` VALUES (536, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"Content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"Id\":85,\"Mobile\":\"1\",\"Results\":\"1\",\"Scene\":1,\"SendTime\":\"2024-08-18 01:19:00\",\"Status\":1}', '', 1, '2024-08-14 01:22:36', '2024-08-14 01:22:36', 15, '2024-08-14 01:22:36');
INSERT INTO `x_system_log_operate` VALUES (537, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"Content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"Id\":85,\"Mobile\":\"1\",\"Results\":\"1\",\"Scene\":1,\"SendTime\":\"2024-08-04 01:19:00\",\"Status\":1}', '', 1, '2024-08-14 01:22:45', '2024-08-14 01:22:45', 11, '2024-08-14 01:22:45');
INSERT INTO `x_system_log_operate` VALUES (538, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"Content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"Id\":85,\"Mobile\":\"1\",\"Results\":\"1\",\"Scene\":1,\"SendTime\":\"2024-08-11 01:19:00\",\"Status\":1}', '', 1, '2024-08-14 01:22:54', '2024-08-14 01:22:54', 8, '2024-08-14 01:22:54');
INSERT INTO `x_system_log_operate` VALUES (539, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"Content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"Id\":85,\"Mobile\":\"1\",\"Results\":\"1\",\"Scene\":1,\"SendTime\":\"2024-08-12 01:19:00\",\"Status\":1}', '', 1, '2024-08-14 01:24:17', '2024-08-14 01:24:17', 13, '2024-08-14 01:24:17');
INSERT INTO `x_system_log_operate` VALUES (540, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"Content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"Id\":85,\"Mobile\":\"1\",\"Results\":\"1\",\"Scene\":1,\"SendTime\":\"2024-08-11 01:19:00\",\"Status\":1}', '', 1, '2024-08-14 01:25:04', '2024-08-14 01:27:20', 136055, '2024-08-14 01:27:20');
INSERT INTO `x_system_log_operate` VALUES (541, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"Content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"Id\":85,\"Mobile\":\"1\",\"Results\":\"1\",\"Scene\":1,\"SendTime\":\"2024-08-11 01:19:00\",\"Status\":1}', '', 1, '2024-08-14 01:30:07', '2024-08-14 01:31:18', 71632, '2024-08-14 01:31:18');
INSERT INTO `x_system_log_operate` VALUES (542, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"Content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"Id\":85,\"Mobile\":\"1\",\"Results\":\"1\",\"Scene\":1,\"SendTime\":\"2024-08-04 01:19:00\",\"Status\":1}', '', 1, '2024-08-14 01:32:09', '2024-08-14 01:32:12', 2483, '2024-08-14 01:32:12');
INSERT INTO `x_system_log_operate` VALUES (543, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"Content\":\"\\u003cp\\u003e2\\u003c/p\\u003e\",\"Id\":85,\"Mobile\":\"2\",\"Results\":\"2\",\"Scene\":2,\"SendTime\":\"2024-08-02 01:19:00\",\"Status\":1}', '', 1, '2024-08-14 01:32:48', '2024-08-14 01:32:59', 11323, '2024-08-14 01:32:59');
INSERT INTO `x_system_log_operate` VALUES (544, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"Content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"Id\":null,\"Mobile\":\"1\",\"Results\":null,\"Scene\":1,\"SendTime\":null,\"Status\":1}', '', 1, '2024-08-14 01:36:37', '2024-08-14 01:36:37', 3, '2024-08-14 01:36:37');
INSERT INTO `x_system_log_operate` VALUES (545, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"Content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"Id\":null,\"Mobile\":\"1\",\"Results\":\"1\",\"Scene\":1,\"SendTime\":\"2024-08-01 12:00:00\",\"Status\":1}', '', 1, '2024-08-14 01:36:51', '2024-08-14 01:36:51', 2, '2024-08-14 01:36:51');
INSERT INTO `x_system_log_operate` VALUES (546, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"Content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"Id\":87,\"Mobile\":\"1\",\"Results\":\"1\",\"Scene\":1,\"SendTime\":\"2024-08-04 12:00:00\",\"Status\":1}', '', 1, '2024-08-14 01:51:05', '2024-08-14 01:51:08', 3283, '2024-08-14 01:51:08');
INSERT INTO `x_system_log_operate` VALUES (547, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"Content\":\"\\u003cp\\u003e1222\\u003c/p\\u003e\",\"Id\":87,\"Mobile\":\"2\",\"Results\":\"122\",\"Scene\":2,\"SendTime\":\"2024-08-04 12:00:00\",\"Status\":1}', '', 1, '2024-08-14 01:52:21', '2024-08-14 01:52:24', 2762, '2024-08-14 01:52:24');
INSERT INTO `x_system_log_operate` VALUES (548, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"Content\":\"\\u003cp\\u003e1222\\u003c/p\\u003e\",\"Id\":87,\"Mobile\":\"2\",\"Results\":\"122\",\"Scene\":2,\"SendTime\":null,\"Status\":1}', '', 1, '2024-08-14 01:52:34', '2024-08-14 01:52:36', 2016, '2024-08-14 01:52:36');
INSERT INTO `x_system_log_operate` VALUES (549, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-08-14 02:11:55', '2024-08-14 02:11:55', 2, '2024-08-14 02:11:55');
INSERT INTO `x_system_log_operate` VALUES (550, 1, 'GET', '服务监控', '127.0.0.1', '/api/admin/monitor/server', 'x_admin/admin/monitor.monitorHandler.server-fm', '', '', 1, '2024-08-14 02:12:04', '2024-08-14 02:12:04', 9, '2024-08-14 02:12:04');
INSERT INTO `x_system_log_operate` VALUES (551, 1, 'GET', '缓存监控', '127.0.0.1', '/api/admin/monitor/cache', 'x_admin/admin/monitor.monitorHandler.cache-fm', '', '', 1, '2024-08-14 02:12:05', '2024-08-14 02:12:05', 1, '2024-08-14 02:12:05');
INSERT INTO `x_system_log_operate` VALUES (552, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"Content\":\"\\u003cp\\u003e11\\u003c/p\\u003e\",\"Id\":88,\"Mobile\":\"11\",\"Results\":\"1\",\"Scene\":111,\"SendTime\":\"2024-08-25 12:00:00\",\"Status\":2}', '', 1, '2024-08-14 02:24:02', '2024-08-14 02:24:05', 2458, '2024-08-14 02:24:05');
INSERT INTO `x_system_log_operate` VALUES (553, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-08-14 10:50:42', '2024-08-14 10:50:42', 7, '2024-08-14 10:50:42');
INSERT INTO `x_system_log_operate` VALUES (554, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"Content\":\"\\u003cp\\u003e2\\u003c/p\\u003e\",\"Id\":null,\"Mobile\":\"2\",\"Results\":\"2\",\"Scene\":2,\"SendTime\":\"2024-08-16 12:00:00\",\"Status\":2}', '', 1, '2024-08-14 11:11:34', '2024-08-14 11:11:34', 15, '2024-08-14 11:11:34');
INSERT INTO `x_system_log_operate` VALUES (555, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"Content\":\"\\u003cp\\u003e\\u003cbr\\u003e\\u003c/p\\u003e\",\"Id\":89,\"Mobile\":\"\",\"Results\":\"\",\"Scene\":\"\",\"SendTime\":null,\"Status\":2}', 'Error #01: 500:编辑失败\n', 2, '2024-08-14 11:12:01', '2024-08-14 11:12:05', 3793, '2024-08-14 11:12:05');
INSERT INTO `x_system_log_operate` VALUES (556, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"Content\":\"\\u003cp\\u003e\\u003cbr\\u003e\\u003c/p\\u003e\",\"Id\":89,\"Mobile\":\"\",\"Results\":\"\",\"Scene\":\"\",\"SendTime\":null,\"Status\":2}', 'Error #01: 310:参数校验错误\n', 2, '2024-08-14 11:18:37', '2024-08-14 11:18:37', 0, '2024-08-14 11:18:37');
INSERT INTO `x_system_log_operate` VALUES (557, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"Content\":\"\\u003cp\\u003e\\u003cbr\\u003e\\u003c/p\\u003e\",\"Id\":89,\"Mobile\":\"\",\"Results\":\"\",\"Scene\":\"\",\"SendTime\":null,\"Status\":2}', 'Error #01: 310:参数校验错误\n', 2, '2024-08-14 11:21:08', '2024-08-14 11:21:08', 1, '2024-08-14 11:21:08');
INSERT INTO `x_system_log_operate` VALUES (558, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"Content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"Id\":null,\"Mobile\":\"1\",\"Results\":\"1\",\"Scene\":1,\"SendTime\":\"2024-08-15 12:00:00\",\"Status\":2}', '', 1, '2024-08-14 11:46:13', '2024-08-14 11:47:08', 54788, '2024-08-14 11:47:08');
INSERT INTO `x_system_log_operate` VALUES (559, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"Content\":null,\"Id\":null,\"Mobile\":null,\"Results\":null,\"Scene\":null,\"SendTime\":null,\"Status\":null}', '', 1, '2024-08-14 11:48:20', '2024-08-14 11:49:19', 59355, '2024-08-14 11:49:19');
INSERT INTO `x_system_log_operate` VALUES (560, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"Content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"Id\":null,\"Mobile\":\"1\",\"Results\":\"1\",\"Scene\":1,\"SendTime\":\"2024-08-14 11:50:15\",\"Status\":2}', 'Error #01: 500:添加失败\n', 2, '2024-08-14 11:50:18', '2024-08-14 11:50:22', 4128, '2024-08-14 11:50:22');
INSERT INTO `x_system_log_operate` VALUES (561, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"Content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"Id\":null,\"Mobile\":\"1\",\"Results\":\"1\",\"Scene\":1,\"SendTime\":\"2024-08-14 11:50:15\",\"Status\":2}', 'Error #01: 500:添加失败\n', 2, '2024-08-14 11:56:30', '2024-08-14 11:57:46', 76733, '2024-08-14 11:57:46');
INSERT INTO `x_system_log_operate` VALUES (562, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"Content\":\"\\u003cp\\u003e2\\u003c/p\\u003e\",\"Id\":null,\"Mobile\":\"2\",\"Results\":\"2\",\"Scene\":2,\"SendTime\":\"2024-08-15 12:00:00\",\"Status\":2}', 'Error #01: 500:添加失败\n', 2, '2024-08-14 12:00:51', '2024-08-14 12:01:36', 45472, '2024-08-14 12:01:36');
INSERT INTO `x_system_log_operate` VALUES (563, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"Content\":\"\\u003cp\\u003e2\\u003c/p\\u003e\",\"Id\":null,\"Mobile\":\"2\",\"Results\":\"2\",\"Scene\":2,\"SendTime\":\"2024-08-14 12:02:14\",\"Status\":2}', 'Error #01: 500:添加失败\n', 2, '2024-08-14 12:02:16', '2024-08-14 12:03:18', 61644, '2024-08-14 12:03:18');
INSERT INTO `x_system_log_operate` VALUES (564, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"Content\":\"\\u003cp\\u003e2\\u003c/p\\u003e\",\"Id\":null,\"Mobile\":\"2\",\"Results\":\"2\",\"Scene\":2,\"SendTime\":\"2024-08-14 12:02:14\",\"Status\":2}', '', 1, '2024-08-14 12:07:35', '2024-08-14 12:08:13', 37916, '2024-08-14 12:08:13');
INSERT INTO `x_system_log_operate` VALUES (565, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"Content\":\"\\u003cp\\u003e2\\u003c/p\\u003e\",\"Id\":null,\"Mobile\":\"2\",\"Results\":\"2\",\"Scene\":2,\"SendTime\":\"2024-08-13 12:00:00\",\"Status\":2}', '', 1, '2024-08-14 12:09:17', '2024-08-14 12:09:57', 40183, '2024-08-14 12:09:57');
INSERT INTO `x_system_log_operate` VALUES (566, 1, 'POST', '系统短信日志删除', '127.0.0.1', '/api/admin/system_log_sms/del', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Del-fm', '{}', 'Error #01: 500:数据不存在!\n', 2, '2024-08-14 12:10:34', '2024-08-14 12:10:34', 3, '2024-08-14 12:10:34');
INSERT INTO `x_system_log_operate` VALUES (567, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"Content\":null,\"Id\":null,\"Mobile\":null,\"Results\":null,\"Scene\":null,\"SendTime\":null,\"Status\":null}', '', 1, '2024-08-14 12:10:56', '2024-08-14 12:11:10', 13801, '2024-08-14 12:11:10');
INSERT INTO `x_system_log_operate` VALUES (568, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"Content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"Id\":96,\"Mobile\":\"1\",\"Results\":\"1\",\"Scene\":1,\"SendTime\":\"2024-08-14 12:12:28\",\"Status\":2}', '', 1, '2024-08-14 12:12:30', '2024-08-14 12:12:32', 2459, '2024-08-14 12:12:32');
INSERT INTO `x_system_log_operate` VALUES (569, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"Content\":\"\\u003cp\\u003e\\u003cbr\\u003e\\u003c/p\\u003e\",\"Id\":100,\"Mobile\":\"11\",\"Results\":\"\",\"Scene\":\"\",\"SendTime\":\"2024-08-12 19:09:14\",\"Status\":2}', 'Error #01: 310:参数校验错误\n', 2, '2024-08-14 12:35:46', '2024-08-14 12:35:46', 0, '2024-08-14 12:35:46');
INSERT INTO `x_system_log_operate` VALUES (570, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"Content\":\"\\u003cp\\u003e\\u003cbr\\u003e\\u003c/p\\u003e\",\"Id\":100,\"Mobile\":\"11\",\"Results\":\"\",\"Scene\":\"\",\"SendTime\":\"2024-08-12 19:09:14\",\"Status\":2}', 'Error #01: 310:参数校验错误\n', 2, '2024-08-14 12:36:29', '2024-08-14 12:36:29', 1, '2024-08-14 12:36:29');
INSERT INTO `x_system_log_operate` VALUES (571, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"Content\":\"\\u003cp\\u003e\\u003cbr\\u003e\\u003c/p\\u003e\",\"Id\":100,\"Mobile\":\"11\",\"Results\":\"\",\"Scene\":\"\",\"SendTime\":\"2024-08-12 19:09:14\",\"Status\":2}', 'Error #01: 500:编辑失败\n', 2, '2024-08-14 12:39:15', '2024-08-14 12:39:50', 34416, '2024-08-14 12:39:50');
INSERT INTO `x_system_log_operate` VALUES (572, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"Content\":\"\\u003cp\\u003e\\u003cbr\\u003e\\u003c/p\\u003e\",\"Id\":100,\"Mobile\":\"11\",\"Results\":\"\",\"Scene\":\"\",\"SendTime\":\"2024-08-12 19:09:14\",\"Status\":2}', 'Error #01: 500:编辑失败\n', 2, '2024-08-14 13:05:29', '2024-08-14 13:07:28', 118277, '2024-08-14 13:07:28');
INSERT INTO `x_system_log_operate` VALUES (573, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"Content\":\"\\u003cp\\u003e11\\u003c/p\\u003e\",\"Id\":100,\"Results\":\"1\",\"SendTime\":\"2024-08-12 19:09:14\",\"Status\":2}', '', 1, '2024-08-14 13:59:56', '2024-08-14 13:59:56', 5, '2024-08-14 13:59:56');
INSERT INTO `x_system_log_operate` VALUES (574, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"Content\":\"\\u003cp\\u003e11\\u003c/p\\u003e\",\"Id\":100,\"Results\":\"1\",\"SendTime\":\"2024-08-12 19:09:14\",\"Status\":2}', '', 1, '2024-08-14 14:00:33', '2024-08-14 14:01:30', 57449, '2024-08-14 14:01:30');
INSERT INTO `x_system_log_operate` VALUES (575, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"Content\":\"\\u003cp\\u003e11\\u003c/p\\u003e\",\"Id\":100,\"Mobile\":\"11\",\"Results\":\"1\",\"SendTime\":\"2024-08-12 19:09:14\",\"Status\":2}', '', 1, '2024-08-14 14:02:20', '2024-08-14 14:13:35', 675382, '2024-08-14 14:13:35');
INSERT INTO `x_system_log_operate` VALUES (576, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"Content\":\"\\u003cp\\u003e11\\u003c/p\\u003e\",\"Id\":100,\"Mobile\":\"11\",\"Results\":\"1\",\"Scene\":111,\"SendTime\":\"2024-08-12 19:09:14\",\"Status\":2}', '', 1, '2024-08-14 16:54:54', '2024-08-14 16:55:47', 53553, '2024-08-14 16:55:47');
INSERT INTO `x_system_log_operate` VALUES (577, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=d3aa184384b3f23b3e879b1e48472919hB5HhK&', '', 1, '2024-08-14 17:03:43', '2024-08-14 17:03:43', 14, '2024-08-14 17:03:43');
INSERT INTO `x_system_log_operate` VALUES (578, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=d3aa184384b3f23b3e879b1e48472919hB5HhK&', '', 1, '2024-08-14 17:14:38', '2024-08-14 17:14:38', 11, '2024-08-14 17:14:38');
INSERT INTO `x_system_log_operate` VALUES (579, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=d3aa184384b3f23b3e879b1e48472919hB5HhK&', 'interface conversion: interface {} is nil, not string', 2, '2024-08-14 17:24:14', '2024-08-14 17:24:14', 5, '2024-08-14 17:24:14');
INSERT INTO `x_system_log_operate` VALUES (580, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=d3aa184384b3f23b3e879b1e48472919hB5HhK&', 'interface conversion: interface {} is nil, not string', 2, '2024-08-14 17:24:23', '2024-08-14 17:24:23', 5, '2024-08-14 17:24:23');
INSERT INTO `x_system_log_operate` VALUES (581, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=d3aa184384b3f23b3e879b1e48472919hB5HhK&', '', 1, '2024-08-14 17:29:54', '2024-08-14 17:30:47', 53595, '2024-08-14 17:30:47');
INSERT INTO `x_system_log_operate` VALUES (582, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=d3aa184384b3f23b3e879b1e48472919hB5HhK&', '', 1, '2024-08-14 17:34:52', '2024-08-14 17:34:52', 12, '2024-08-14 17:34:52');
INSERT INTO `x_system_log_operate` VALUES (583, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=d3aa184384b3f23b3e879b1e48472919hB5HhK&', '', 1, '2024-08-14 17:34:53', '2024-08-14 17:34:53', 14, '2024-08-14 17:34:53');
INSERT INTO `x_system_log_operate` VALUES (584, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=d3aa184384b3f23b3e879b1e48472919hB5HhK&', '', 1, '2024-08-14 17:41:33', '2024-08-14 17:41:33', 12, '2024-08-14 17:41:33');
INSERT INTO `x_system_log_operate` VALUES (585, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=d3aa184384b3f23b3e879b1e48472919hB5HhK&', '', 1, '2024-08-14 17:46:28', '2024-08-14 17:46:37', 8905, '2024-08-14 17:46:37');
INSERT INTO `x_system_log_operate` VALUES (586, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=d3aa184384b3f23b3e879b1e48472919hB5HhK&', '', 1, '2024-08-14 17:48:06', '2024-08-14 17:48:06', 13, '2024-08-14 17:48:06');
INSERT INTO `x_system_log_operate` VALUES (587, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"Content\":\"\\u003cp\\u003e11\\u003c/p\\u003e\",\"Id\":105,\"Mobile\":\"11\",\"Results\":\"1\",\"Status\":2}', '', 1, '2024-08-14 17:50:33', '2024-08-14 17:50:36', 2889, '2024-08-14 17:50:36');
INSERT INTO `x_system_log_operate` VALUES (588, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"Content\":\"\\u003cp\\u003e11\\u003c/p\\u003e\",\"Id\":105,\"Mobile\":\"11\",\"Results\":\"1\",\"Scene\":11,\"SendTime\":null,\"Status\":2}', '', 1, '2024-08-14 17:51:08', '2024-08-14 17:51:12', 4574, '2024-08-14 17:51:12');
INSERT INTO `x_system_log_operate` VALUES (589, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"Content\":\"\\u003cp\\u003e11\\u003c/p\\u003e\",\"Id\":105,\"Mobile\":\"11\",\"Results\":\"1\",\"Scene\":\"\",\"SendTime\":null,\"Status\":2}', '', 1, '2024-08-14 17:51:17', '2024-08-14 17:51:19', 2222, '2024-08-14 17:51:19');
INSERT INTO `x_system_log_operate` VALUES (590, 1, 'GET', '角色列表', '127.0.0.1', '/api/admin/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, '2024-08-15 11:10:36', '2024-08-15 11:10:36', 3, '2024-08-15 11:10:36');
INSERT INTO `x_system_log_operate` VALUES (591, 1, 'POST', '系统短信日志删除', '127.0.0.1', '/api/admin/system_log_sms/del', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Del-fm', '{}', 'Error #01: 500:数据不存在!\n', 2, '2024-08-15 16:15:51', '2024-08-15 16:15:51', 33, '2024-08-15 16:15:51');
INSERT INTO `x_system_log_operate` VALUES (592, 1, 'POST', '系统短信日志删除', '127.0.0.1', '/api/admin/system_log_sms/del', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Del-fm', '{}', 'Error #01: 500:数据不存在!\n', 2, '2024-08-15 16:16:36', '2024-08-15 16:16:36', 16, '2024-08-15 16:16:36');
INSERT INTO `x_system_log_operate` VALUES (593, 1, 'POST', '系统短信日志删除', '127.0.0.1', '/api/admin/system_log_sms/del', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Del-fm', '{\"id\":106}', '', 1, '2024-08-15 16:16:56', '2024-08-15 16:16:56', 132, '2024-08-15 16:16:56');
INSERT INTO `x_system_log_operate` VALUES (594, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"Content\":\"\\u003cp\\u003e11\\u003c/p\\u003e\",\"Id\":105,\"Mobile\":\"11\",\"Results\":\"1\",\"Scene\":\"\",\"SendTime\":null,\"Status\":2}', '', 1, '2024-08-15 16:43:44', '2024-08-15 16:44:32', 48566, '2024-08-15 16:44:32');
INSERT INTO `x_system_log_operate` VALUES (595, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '', 'Error #01: 310:参数校验错误\n', 2, '2024-08-15 17:06:02', '2024-08-15 17:06:02', 25, '2024-08-15 17:06:02');
INSERT INTO `x_system_log_operate` VALUES (596, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '', 'Error #01: 310:参数校验错误\n', 2, '2024-08-15 17:06:44', '2024-08-15 17:06:44', 0, '2024-08-15 17:06:44');
INSERT INTO `x_system_log_operate` VALUES (597, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '', 'Error #01: 310:参数校验错误\n', 2, '2024-08-15 17:07:12', '2024-08-15 17:07:12', 1, '2024-08-15 17:07:12');
INSERT INTO `x_system_log_operate` VALUES (598, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '', 'Error #01: 310:参数校验错误\n', 2, '2024-08-15 17:07:17', '2024-08-15 17:07:17', 0, '2024-08-15 17:07:17');
INSERT INTO `x_system_log_operate` VALUES (599, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '', 'Error #01: 310:参数校验错误\n', 2, '2024-08-15 17:07:22', '2024-08-15 17:07:22', 0, '2024-08-15 17:07:22');
INSERT INTO `x_system_log_operate` VALUES (600, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '', 'Error #01: 310:参数校验错误\n', 2, '2024-08-15 17:07:25', '2024-08-15 17:07:25', 0, '2024-08-15 17:07:25');
INSERT INTO `x_system_log_operate` VALUES (601, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '', 'Error #01: 310:参数校验错误\n', 2, '2024-08-15 17:07:29', '2024-08-15 17:07:29', 0, '2024-08-15 17:07:29');
INSERT INTO `x_system_log_operate` VALUES (602, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '', 'Error #01: 310:参数校验错误\n', 2, '2024-08-15 17:07:34', '2024-08-15 17:07:34', 0, '2024-08-15 17:07:34');
INSERT INTO `x_system_log_operate` VALUES (603, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '', 'Error #01: 310:参数校验错误\n', 2, '2024-08-15 17:08:02', '2024-08-15 17:08:02', 0, '2024-08-15 17:08:02');
INSERT INTO `x_system_log_operate` VALUES (604, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '{\"Content\":\"\\u003cp\\u003e1\\u003c/p\\u003e\",\"Id\":null,\"Mobile\":\"1\",\"Results\":\"1\",\"Scene\":1,\"SendTime\":null,\"Status\":1}', '', 1, '2024-08-15 17:09:21', '2024-08-15 17:09:38', 17517, '2024-08-15 17:09:38');
INSERT INTO `x_system_log_operate` VALUES (605, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '', 'Error #01: 310:参数校验错误\n', 2, '2024-08-15 17:09:56', '2024-08-15 17:09:56', 0, '2024-08-15 17:09:56');
INSERT INTO `x_system_log_operate` VALUES (606, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '', 'Error #01: 310:参数校验错误\n', 2, '2024-08-15 17:11:17', '2024-08-15 17:11:17', 1, '2024-08-15 17:11:17');
INSERT INTO `x_system_log_operate` VALUES (607, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '', 'Error #01: 310:参数校验错误\n', 2, '2024-08-15 17:11:23', '2024-08-15 17:11:23', 1, '2024-08-15 17:11:23');
INSERT INTO `x_system_log_operate` VALUES (608, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '', 'Error #01: 310:参数校验错误\n', 2, '2024-08-15 17:12:30', '2024-08-15 17:25:51', 800665, '2024-08-15 17:25:51');
INSERT INTO `x_system_log_operate` VALUES (609, 1, 'POST', '系统短信日志新增', '127.0.0.1', '/api/admin/system_log_sms/add', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Add-fm', '', 'Error #01: 310:参数校验错误\n', 2, '2024-08-15 17:25:51', '2024-08-15 17:25:53', 1555, '2024-08-15 17:25:53');
INSERT INTO `x_system_log_operate` VALUES (610, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"Content\":\"\\u003cp\\u003e123\\u003c/p\\u003e\",\"Id\":107,\"Mobile\":\"23\",\"Results\":\"1123\",\"Scene\":\"\",\"SendTime\":\"2024-07-11 12:00:00\",\"Status\":1}', 'runtime error: invalid memory address or nil pointer dereference', 2, '2024-08-15 22:17:14', '2024-08-15 22:17:54', 40287, '2024-08-15 22:17:54');
INSERT INTO `x_system_log_operate` VALUES (611, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"Content\":\"\\u003cp\\u003e123\\u003c/p\\u003e\",\"Id\":107,\"Mobile\":\"23\",\"Results\":\"1123\",\"Scene\":\"\",\"SendTime\":\"2024-07-11 12:00:00\",\"Status\":1}', 'runtime error: invalid memory address or nil pointer dereference', 2, '2024-08-15 22:18:51', '2024-08-15 22:20:10', 78588, '2024-08-15 22:20:10');
INSERT INTO `x_system_log_operate` VALUES (612, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"Content\":\"\\u003cp\\u003e123\\u003c/p\\u003e\",\"Id\":107,\"Mobile\":\"23\",\"Results\":\"1123\",\"Scene\":\"\",\"SendTime\":\"2024-07-11 12:00:00\",\"Status\":1}', 'runtime error: invalid memory address or nil pointer dereference', 2, '2024-08-15 22:21:25', '2024-08-15 22:22:44', 79232, '2024-08-15 22:22:44');
INSERT INTO `x_system_log_operate` VALUES (613, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"Content\":\"\\u003cp\\u003e123\\u003c/p\\u003e\",\"Id\":107,\"Mobile\":\"23\",\"Results\":\"1123\",\"Scene\":\"\",\"SendTime\":\"2024-07-11 12:00:00\",\"Status\":1}', 'runtime error: invalid memory address or nil pointer dereference', 2, '2024-08-15 22:29:26', '2024-08-15 22:32:28', 181990, '2024-08-15 22:32:28');
INSERT INTO `x_system_log_operate` VALUES (614, 1, 'POST', '系统短信日志编辑', '127.0.0.1', '/api/admin/system_log_sms/edit', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).Edit-fm', '{\"Content\":\"\\u003cp\\u003e123\\u003c/p\\u003e\",\"Id\":107,\"Mobile\":\"23\",\"Results\":\"1123\",\"Scene\":\"\",\"SendTime\":\"2024-07-11 12:00:00\",\"Status\":1}', 'runtime error: invalid memory address or nil pointer dereference', 2, '2024-08-15 22:33:08', '2024-08-15 22:33:28', 19408, '2024-08-15 22:33:28');
INSERT INTO `x_system_log_operate` VALUES (615, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=f5328d3fa0d1a653efe73bd36a3b58ffEvK8uD&', 'runtime error: invalid memory address or nil pointer dereference', 2, '2024-08-16 12:18:00', '2024-08-16 12:18:06', 5871, '2024-08-16 12:18:06');
INSERT INTO `x_system_log_operate` VALUES (616, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=f5328d3fa0d1a653efe73bd36a3b58ffEvK8uD&', 'runtime error: invalid memory address or nil pointer dereference', 2, '2024-08-16 12:20:02', '2024-08-16 12:20:02', 11, '2024-08-16 12:20:02');
INSERT INTO `x_system_log_operate` VALUES (617, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=f5328d3fa0d1a653efe73bd36a3b58ffEvK8uD&', 'runtime error: invalid memory address or nil pointer dereference', 2, '2024-08-16 12:20:59', '2024-08-16 12:20:59', 9, '2024-08-16 12:20:59');
INSERT INTO `x_system_log_operate` VALUES (618, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=f5328d3fa0d1a653efe73bd36a3b58ffEvK8uD&', 'runtime error: invalid memory address or nil pointer dereference', 2, '2024-08-16 12:21:35', '2024-08-16 12:24:26', 171123, '2024-08-16 12:24:26');
INSERT INTO `x_system_log_operate` VALUES (619, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=f5328d3fa0d1a653efe73bd36a3b58ffEvK8uD&', 'runtime error: invalid memory address or nil pointer dereference', 2, '2024-08-16 12:25:37', '2024-08-16 12:25:51', 13702, '2024-08-16 12:25:51');
INSERT INTO `x_system_log_operate` VALUES (620, 1, 'GET', '系统短信日志导出', '127.0.0.1', '/api/admin/system_log_sms/ExportFile', 'x_admin/admin/system_log_sms.(*SystemLogSmsHandler).ExportFile-fm', 'token=f5328d3fa0d1a653efe73bd36a3b58ffEvK8uD&', 'runtime error: invalid memory address or nil pointer dereference', 2, '2024-08-16 12:25:53', '2024-08-16 12:27:31', 98793, '2024-08-16 12:27:31');
INSERT INTO `x_system_log_operate` VALUES (621, 1, 'POST', '监控-错误列新增', '127.0.0.1', '/api/admin/monitor_error/add', 'x_admin/admin/monitor_error.(*MonitorErrorHandler).Add-fm', '{\"ClientId\":\"1\",\"ClientTime\":null,\"EventType\":\"1\",\"Id\":null,\"Md5\":null,\"Message\":null,\"Path\":null,\"ProjectKey\":\"2a01a60efed14d6eab279ba5c664ceac\",\"Stack\":null}', 'runtime error: invalid memory address or nil pointer dereference', 2, '2024-09-25 18:30:03', '2024-09-25 18:30:03', 1, '2024-09-25 18:30:03');
INSERT INTO `x_system_log_operate` VALUES (622, 1, 'POST', '监控-错误列新增', '127.0.0.1', '/api/admin/monitor_error/add', 'x_admin/admin/monitor_error.(*MonitorErrorHandler).Add-fm', '{\"ClientId\":\"1\",\"ClientTime\":null,\"EventType\":\"1\",\"Id\":null,\"Md5\":null,\"Message\":null,\"Path\":\"1\",\"ProjectKey\":\"2a01a60efed14d6eab279ba5c664ceac\",\"Stack\":null}', 'runtime error: invalid memory address or nil pointer dereference', 2, '2024-09-25 18:30:51', '2024-09-25 18:30:51', 0, '2024-09-25 18:30:51');

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
INSERT INTO `x_system_log_sms` VALUES (109, 2, '2', '<p>2</p>', 2, '2', '2024-08-01 12:00:00', '2024-08-16 12:07:57', '2024-08-16 12:31:55');
INSERT INTO `x_system_log_sms` VALUES (110, 2, '2', '2', 2, '2', NULL, '2024-08-16 12:07:57', '2024-08-16 12:07:57');
INSERT INTO `x_system_log_sms` VALUES (111, 2, '2', '<p>2</p>', 2, '2', '2024-08-01 12:00:00', '2024-08-16 12:07:57', '2024-08-16 12:31:55');
INSERT INTO `x_system_log_sms` VALUES (112, 23, '23', '<p>23</p>', 2, '3', '2024-07-04 12:00:00', '2024-08-16 12:07:57', '2024-08-16 17:49:28');
INSERT INTO `x_system_log_sms` VALUES (113, 2, '2', '<p>2</p>', 2, '2', '2024-08-01 12:00:00', '2024-08-16 12:07:57', '2024-08-16 12:31:55');
INSERT INTO `x_system_log_sms` VALUES (114, 2, '2', '2', 2, '2', NULL, '2024-08-16 12:07:57', '2024-08-16 12:07:57');
INSERT INTO `x_system_log_sms` VALUES (115, 2, '2', '<p>2</p>', 2, '2', '2024-08-01 12:00:00', '2024-08-16 12:07:57', '2024-08-16 12:31:55');
INSERT INTO `x_system_log_sms` VALUES (116, 23, '23', '<p>23</p>', 2, '3', '2024-07-04 12:00:00', '2024-08-16 12:07:57', '2024-08-16 17:49:28');

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
