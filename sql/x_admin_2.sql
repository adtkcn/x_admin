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

 Date: 08/11/2024 17:50:01
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
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '申请流程' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_flow_apply
-- ----------------------------
INSERT INTO `x_flow_apply` VALUES (1, 1, 1, 'admin', 'a', 1, 'a', '{\"widgetList\":[{\"key\":33700,\"type\":\"input\",\"icon\":\"text-field\",\"formItemFlag\":true,\"options\":{\"name\":\"input97761\",\"label\":\"input\",\"labelAlign\":\"\",\"type\":\"text\",\"defaultValue\":\"\",\"placeholder\":\"\",\"columnWidth\":\"200px\",\"size\":\"\",\"labelWidth\":null,\"labelHidden\":false,\"readonly\":false,\"disabled\":false,\"hidden\":false,\"clearable\":true,\"showPassword\":false,\"required\":false,\"requiredHint\":\"\",\"validation\":\"\",\"validationHint\":\"\",\"customClass\":\"\",\"labelIconClass\":null,\"labelIconPosition\":\"rear\",\"labelTooltip\":null,\"minLength\":null,\"maxLength\":null,\"showWordLimit\":false,\"prefixIcon\":\"\",\"suffixIcon\":\"\",\"appendButton\":false,\"appendButtonDisabled\":false,\"buttonIcon\":\"custom-search\",\"onCreated\":\"\",\"onMounted\":\"\",\"onInput\":\"\",\"onChange\":\"\",\"onFocus\":\"\",\"onBlur\":\"\",\"onValidate\":\"\",\"onAppendButtonClick\":\"\"},\"id\":\"input97761\"}],\"formConfig\":{\"modelName\":\"formData\",\"refName\":\"vForm\",\"rulesName\":\"rules\",\"labelWidth\":80,\"labelPosition\":\"left\",\"size\":\"\",\"labelAlign\":\"label-left-align\",\"cssCode\":\"\",\"customClass\":\"\",\"functions\":\"\",\"layoutType\":\"PC\",\"jsonVersion\":3,\"onFormCreated\":\"\",\"onFormMounted\":\"\",\"onFormDataChange\":\"\"}}', '{\"nodes\":[{\"id\":\"Event_a677124\",\"type\":\"bpmn:startEvent\",\"x\":300,\"y\":100,\"properties\":{\"width\":36,\"height\":36},\"zIndex\":1013,\"text\":{\"x\":300,\"y\":140,\"value\":\"开始\"}},{\"id\":\"Activity_2c996a7\",\"type\":\"bpmn:userTask\",\"x\":500,\"y\":120,\"properties\":{\"width\":100,\"height\":80},\"zIndex\":1015,\"text\":{\"x\":500,\"y\":120,\"value\":\"审批\"}},{\"id\":\"Event_129de4e\",\"type\":\"bpmn:endEvent\",\"x\":730,\"y\":130,\"properties\":{\"width\":36,\"height\":36},\"zIndex\":1012,\"text\":{\"x\":730,\"y\":170,\"value\":\"结束\"}}],\"edges\":[{\"id\":\"2b28919f-30fe-4feb-9db9-0863359f87ed\",\"type\":\"pro-polyline\",\"properties\":{},\"sourceNodeId\":\"Event_a677124\",\"targetNodeId\":\"Activity_2c996a7\",\"startPoint\":{\"x\":318,\"y\":100},\"endPoint\":{\"x\":450,\"y\":120},\"zIndex\":1014,\"pointsList\":[{\"x\":318,\"y\":100},{\"x\":384,\"y\":100},{\"x\":384,\"y\":120},{\"x\":450,\"y\":120}]},{\"id\":\"dd95a2cd-2ecc-4e03-9d86-932673dd677b\",\"type\":\"pro-polyline\",\"properties\":{},\"sourceNodeId\":\"Activity_2c996a7\",\"targetNodeId\":\"Event_129de4e\",\"startPoint\":{\"x\":550,\"y\":120},\"endPoint\":{\"x\":712,\"y\":130},\"zIndex\":1016,\"pointsList\":[{\"x\":550,\"y\":120},{\"x\":631,\"y\":120},{\"x\":631,\"y\":130},{\"x\":712,\"y\":130}]}]}', '[{\"id\":\"Event_a677124\",\"pid\":0,\"label\":\"开始\",\"type\":\"bpmn:startEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_2c996a7\",\"pid\":\"Event_a677124\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_129de4e\",\"pid\":\"Activity_2c996a7\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":null}]}]},{\"id\":\"Activity_2c996a7\",\"pid\":\"Event_a677124\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_129de4e\",\"pid\":\"Activity_2c996a7\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":null}]},{\"id\":\"Event_129de4e\",\"pid\":\"Activity_2c996a7\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":null},{\"id\":\"Activity_2c996a7\",\"pid\":0,\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_129de4e\",\"pid\":\"Activity_2c996a7\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":null}]},{\"id\":\"Event_129de4e\",\"pid\":\"Activity_2c996a7\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":null},{\"id\":\"Event_129de4e\",\"pid\":0,\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":null}]', '', 1, 0, '2024-11-07 15:01:53', '2024-11-07 15:01:53', NULL);

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
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '流程历史' ROW_FORMAT = DYNAMIC;

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
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '流程模板' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_flow_template
-- ----------------------------
INSERT INTO `x_flow_template` VALUES (1, 'a', 1, 'a', '{\"widgetList\":[{\"key\":33700,\"type\":\"input\",\"icon\":\"text-field\",\"formItemFlag\":true,\"options\":{\"name\":\"input97761\",\"label\":\"input\",\"labelAlign\":\"\",\"type\":\"text\",\"defaultValue\":\"\",\"placeholder\":\"\",\"columnWidth\":\"200px\",\"size\":\"\",\"labelWidth\":null,\"labelHidden\":false,\"readonly\":false,\"disabled\":false,\"hidden\":false,\"clearable\":true,\"showPassword\":false,\"required\":false,\"requiredHint\":\"\",\"validation\":\"\",\"validationHint\":\"\",\"customClass\":\"\",\"labelIconClass\":null,\"labelIconPosition\":\"rear\",\"labelTooltip\":null,\"minLength\":null,\"maxLength\":null,\"showWordLimit\":false,\"prefixIcon\":\"\",\"suffixIcon\":\"\",\"appendButton\":false,\"appendButtonDisabled\":false,\"buttonIcon\":\"custom-search\",\"onCreated\":\"\",\"onMounted\":\"\",\"onInput\":\"\",\"onChange\":\"\",\"onFocus\":\"\",\"onBlur\":\"\",\"onValidate\":\"\",\"onAppendButtonClick\":\"\"},\"id\":\"input97761\"}],\"formConfig\":{\"modelName\":\"formData\",\"refName\":\"vForm\",\"rulesName\":\"rules\",\"labelWidth\":80,\"labelPosition\":\"left\",\"size\":\"\",\"labelAlign\":\"label-left-align\",\"cssCode\":\"\",\"customClass\":\"\",\"functions\":\"\",\"layoutType\":\"PC\",\"jsonVersion\":3,\"onFormCreated\":\"\",\"onFormMounted\":\"\",\"onFormDataChange\":\"\"}}', '{\"nodes\":[{\"id\":\"Event_a677124\",\"type\":\"bpmn:startEvent\",\"x\":300,\"y\":100,\"properties\":{\"width\":36,\"height\":36},\"zIndex\":1013,\"text\":{\"x\":300,\"y\":140,\"value\":\"开始\"}},{\"id\":\"Activity_2c996a7\",\"type\":\"bpmn:userTask\",\"x\":500,\"y\":120,\"properties\":{\"width\":100,\"height\":80},\"zIndex\":1015,\"text\":{\"x\":500,\"y\":120,\"value\":\"审批\"}},{\"id\":\"Event_129de4e\",\"type\":\"bpmn:endEvent\",\"x\":730,\"y\":130,\"properties\":{\"width\":36,\"height\":36},\"zIndex\":1012,\"text\":{\"x\":730,\"y\":170,\"value\":\"结束\"}}],\"edges\":[{\"id\":\"2b28919f-30fe-4feb-9db9-0863359f87ed\",\"type\":\"pro-polyline\",\"properties\":{},\"sourceNodeId\":\"Event_a677124\",\"targetNodeId\":\"Activity_2c996a7\",\"startPoint\":{\"x\":318,\"y\":100},\"endPoint\":{\"x\":450,\"y\":120},\"zIndex\":1014,\"pointsList\":[{\"x\":318,\"y\":100},{\"x\":384,\"y\":100},{\"x\":384,\"y\":120},{\"x\":450,\"y\":120}]},{\"id\":\"dd95a2cd-2ecc-4e03-9d86-932673dd677b\",\"type\":\"pro-polyline\",\"properties\":{},\"sourceNodeId\":\"Activity_2c996a7\",\"targetNodeId\":\"Event_129de4e\",\"startPoint\":{\"x\":550,\"y\":120},\"endPoint\":{\"x\":712,\"y\":130},\"zIndex\":1016,\"pointsList\":[{\"x\":550,\"y\":120},{\"x\":631,\"y\":120},{\"x\":631,\"y\":130},{\"x\":712,\"y\":130}]}]}', '[{\"id\":\"Event_a677124\",\"pid\":0,\"label\":\"开始\",\"type\":\"bpmn:startEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_2c996a7\",\"pid\":\"Event_a677124\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_129de4e\",\"pid\":\"Activity_2c996a7\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":null}]}]},{\"id\":\"Activity_2c996a7\",\"pid\":\"Event_a677124\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_129de4e\",\"pid\":\"Activity_2c996a7\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":null}]},{\"id\":\"Event_129de4e\",\"pid\":\"Activity_2c996a7\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":null},{\"id\":\"Activity_2c996a7\",\"pid\":0,\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_129de4e\",\"pid\":\"Activity_2c996a7\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":null}]},{\"id\":\"Event_129de4e\",\"pid\":\"Activity_2c996a7\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":null},{\"id\":\"Event_129de4e\",\"pid\":0,\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":null}]', 0, '2024-11-07 15:01:46', '2024-11-07 15:01:46', NULL);

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
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '代码生成业务表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_gen_table
-- ----------------------------
INSERT INTO `x_gen_table` VALUES (1, 'x_monitor_client', '监控-客户端信息', '', '', '', 'monitorClient', 'monitor_client', '监控-客户端信息', '', '', '', 'crud', '', '2024-10-30 11:03:27', '2024-10-30 11:03:27');
INSERT INTO `x_gen_table` VALUES (2, 'x_album', '相册管理表', '', '', '', 'album', 'album', '相册管理', '', '', '', 'crud', '', '2024-11-05 17:53:14', '2024-11-05 17:53:14');
INSERT INTO `x_gen_table` VALUES (3, 'x_album_cate', '相册分类表', '', '', '', 'albumCate', 'album_cate', '相册分类', '', '', '', 'crud', '', '2024-11-05 17:53:14', '2024-11-05 17:53:14');
INSERT INTO `x_gen_table` VALUES (4, 'x_monitor_slow', '监控-错误列表', '', '', '', 'monitorSlow', 'monitor_slow', '监控-错误列', '', '', '', 'crud', '', '2024-11-06 12:13:50', '2024-11-06 12:13:50');
INSERT INTO `x_gen_table` VALUES (5, 'x_notice_setting', '消息通知设置表', '', '', '', 'noticeSetting', 'notice_setting', '消息通知设置', '', '', '', 'crud', '', '2024-11-08 14:12:22', '2024-11-08 14:12:22');

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
) ENGINE = InnoDB AUTO_INCREMENT = 58 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '代码生成字段表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_gen_table_column
-- ----------------------------
INSERT INTO `x_gen_table_column` VALUES (1, 1, 'id', 'uuid', '11', 'int', 'int', 'id', 1, 1, 1, 0, 1, 0, 0, '=', 'number', '', '', 1, '2024-10-30 11:03:27', '2024-10-30 11:03:27');
INSERT INTO `x_gen_table_column` VALUES (2, 1, 'project_key', '项目key', '128', 'varchar', 'string', 'project_key', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', '', 2, '2024-10-30 11:03:27', '2024-10-30 11:03:27');
INSERT INTO `x_gen_table_column` VALUES (3, 1, 'client_id', 'sdk生成的客户端id', '128', 'varchar', 'string', 'client_id', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', '', 3, '2024-10-30 11:03:27', '2024-10-30 11:03:27');
INSERT INTO `x_gen_table_column` VALUES (4, 1, 'user_id', '用户id', '128', 'varchar', 'string', 'user_id', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', '', 4, '2024-10-30 11:03:27', '2024-10-30 11:03:27');
INSERT INTO `x_gen_table_column` VALUES (5, 1, 'os', '系统', '30', 'varchar', 'string', 'os', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', '', 5, '2024-10-30 11:03:27', '2024-10-30 11:03:27');
INSERT INTO `x_gen_table_column` VALUES (6, 1, 'browser', '浏览器', '30', 'varchar', 'string', 'browser', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', '', 6, '2024-10-30 11:03:27', '2024-10-30 11:03:27');
INSERT INTO `x_gen_table_column` VALUES (7, 1, 'country', '国家', '50', 'varchar', 'string', 'country', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', '', 7, '2024-10-30 11:03:27', '2024-10-30 11:03:27');
INSERT INTO `x_gen_table_column` VALUES (8, 1, 'province', '省份', '50', 'varchar', 'string', 'province', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', '', 8, '2024-10-30 11:03:27', '2024-10-30 11:03:27');
INSERT INTO `x_gen_table_column` VALUES (9, 1, 'city', '城市', '50', 'varchar', 'string', 'city', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', '', 9, '2024-10-30 11:03:27', '2024-10-30 11:03:27');
INSERT INTO `x_gen_table_column` VALUES (10, 1, 'operator', '电信运营商', '50', 'varchar', 'string', 'operator', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', '', 10, '2024-10-30 11:03:27', '2024-10-30 11:03:27');
INSERT INTO `x_gen_table_column` VALUES (11, 1, 'ip', 'ip', '50', 'varchar', 'string', 'ip', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', '', 11, '2024-10-30 11:03:27', '2024-10-30 11:03:27');
INSERT INTO `x_gen_table_column` VALUES (12, 1, 'width', '屏幕', '10', 'smallint', 'int', 'width', 0, 0, 1, 1, 1, 1, 1, '=', 'number', '', '', 12, '2024-10-30 11:03:27', '2024-10-30 11:03:27');
INSERT INTO `x_gen_table_column` VALUES (13, 1, 'height', '屏幕高度', '10', 'smallint', 'int', 'height', 0, 0, 1, 1, 1, 1, 1, '=', 'number', '', '', 13, '2024-10-30 11:03:27', '2024-10-30 11:03:27');
INSERT INTO `x_gen_table_column` VALUES (14, 1, 'ua', 'ua记录', '128', 'varchar', 'string', 'ua', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', '', 14, '2024-10-30 11:03:27', '2024-10-30 11:03:27');
INSERT INTO `x_gen_table_column` VALUES (15, 1, 'create_time', '创建时间', '0', 'datetime', 'core.NullTime', 'create_time', 0, 0, 0, 0, 0, 1, 1, '=', 'datetime', '', '', 15, '2024-10-30 11:03:27', '2024-10-30 11:03:27');
INSERT INTO `x_gen_table_column` VALUES (16, 2, 'id', '主键ID', '10', 'int', 'int', 'id', 1, 1, 1, 0, 1, 0, 0, '=', 'number', '', '', 1, '2024-11-05 17:53:14', '2024-11-05 17:53:14');
INSERT INTO `x_gen_table_column` VALUES (17, 2, 'cid', '类目ID', '10', 'int', 'int', 'cid', 0, 0, 1, 1, 1, 1, 1, '=', 'number', '', '', 2, '2024-11-05 17:53:14', '2024-11-05 17:53:14');
INSERT INTO `x_gen_table_column` VALUES (18, 2, 'aid', '管理员ID', '10', 'int', 'int', 'aid', 0, 0, 1, 1, 1, 1, 1, '=', 'number', '', '', 3, '2024-11-05 17:53:14', '2024-11-05 17:53:14');
INSERT INTO `x_gen_table_column` VALUES (19, 2, 'uid', '用户ID', '10', 'int', 'int', 'uid', 0, 0, 1, 1, 1, 1, 1, '=', 'number', '', '', 4, '2024-11-05 17:53:14', '2024-11-05 17:53:14');
INSERT INTO `x_gen_table_column` VALUES (20, 2, 'type', '文件类型: [10=图片, 20=视频]', '2', 'tinyint', 'int', 'type', 0, 0, 1, 1, 1, 1, 1, '=', 'select', '', '', 5, '2024-11-05 17:53:14', '2024-11-05 17:53:14');
INSERT INTO `x_gen_table_column` VALUES (21, 2, 'name', '文件名称', '100', 'varchar', 'string', 'name', 0, 0, 1, 1, 1, 1, 1, 'LIKE', 'input', '', '', 6, '2024-11-05 17:53:14', '2024-11-05 17:53:14');
INSERT INTO `x_gen_table_column` VALUES (22, 2, 'uri', '文件路径', '200', 'varchar', 'string', 'uri', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', '', 7, '2024-11-05 17:53:14', '2024-11-05 17:53:14');
INSERT INTO `x_gen_table_column` VALUES (23, 2, 'ext', '文件扩展', '10', 'varchar', 'string', 'ext', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', '', 8, '2024-11-05 17:53:14', '2024-11-05 17:53:14');
INSERT INTO `x_gen_table_column` VALUES (24, 2, 'size', '文件大小', '10', 'int', 'int', 'size', 0, 0, 1, 1, 1, 1, 1, '=', 'number', '', '', 9, '2024-11-05 17:53:14', '2024-11-05 17:53:14');
INSERT INTO `x_gen_table_column` VALUES (25, 2, 'is_delete', '是否删除: 0=否, 1=是', '1', 'int', 'int', 'is_delete', 0, 0, 0, 0, 0, 0, 0, '=', 'number', '', '', 10, '2024-11-05 17:53:14', '2024-11-05 17:53:14');
INSERT INTO `x_gen_table_column` VALUES (26, 2, 'create_time', '创建时间', '0', 'datetime', 'core.NullTime', 'create_time', 0, 0, 0, 0, 0, 1, 1, '=', 'datetime', '', '', 11, '2024-11-05 17:53:14', '2024-11-05 17:53:14');
INSERT INTO `x_gen_table_column` VALUES (27, 2, 'update_time', '更新时间', '0', 'datetime', 'core.NullTime', 'update_time', 0, 0, 0, 0, 0, 1, 1, '=', 'datetime', '', '', 12, '2024-11-05 17:53:14', '2024-11-05 17:53:14');
INSERT INTO `x_gen_table_column` VALUES (28, 2, 'delete_time', '删除时间', '0', 'datetime', 'core.NullTime', 'delete_time', 0, 0, 0, 0, 0, 0, 0, '=', 'datetime', '', '', 13, '2024-11-05 17:53:14', '2024-11-05 17:53:14');
INSERT INTO `x_gen_table_column` VALUES (29, 3, 'id', '主键ID', '10', 'int', 'int', 'id', 1, 1, 1, 0, 1, 0, 0, '=', 'number', '', '', 1, '2024-11-05 17:53:14', '2024-11-05 17:53:14');
INSERT INTO `x_gen_table_column` VALUES (30, 3, 'pid', '父级ID', '10', 'int', 'int', 'pid', 0, 0, 1, 1, 1, 1, 1, '=', 'number', '', '', 2, '2024-11-05 17:53:14', '2024-11-05 17:53:14');
INSERT INTO `x_gen_table_column` VALUES (31, 3, 'type', '类型: [10=图片, 20=视频]', '2', 'tinyint', 'int', 'type', 0, 0, 1, 1, 1, 1, 1, '=', 'select', '', '', 3, '2024-11-05 17:53:14', '2024-11-05 17:53:14');
INSERT INTO `x_gen_table_column` VALUES (32, 3, 'name', '分类名称', '32', 'varchar', 'string', 'name', 0, 0, 1, 1, 1, 1, 1, 'LIKE', 'input', '', '', 4, '2024-11-05 17:53:14', '2024-11-05 17:53:14');
INSERT INTO `x_gen_table_column` VALUES (33, 3, 'is_delete', '是否删除: [0=否, 1=是]', '1', 'tinyint', 'int', 'is_delete', 0, 0, 0, 0, 0, 0, 0, '=', 'number', '', '', 5, '2024-11-05 17:53:14', '2024-11-05 17:53:14');
INSERT INTO `x_gen_table_column` VALUES (34, 3, 'create_time', '创建时间', '0', 'datetime', 'core.NullTime', 'create_time', 0, 0, 0, 0, 0, 1, 1, '=', 'datetime', '', '', 6, '2024-11-05 17:53:14', '2024-11-05 17:53:14');
INSERT INTO `x_gen_table_column` VALUES (35, 3, 'update_time', '更新时间', '0', 'datetime', 'core.NullTime', 'update_time', 0, 0, 0, 0, 0, 1, 1, '=', 'datetime', '', '', 7, '2024-11-05 17:53:14', '2024-11-05 17:53:14');
INSERT INTO `x_gen_table_column` VALUES (36, 3, 'delete_time', '删除时间', '0', 'datetime', 'core.NullTime', 'delete_time', 0, 0, 0, 0, 0, 0, 0, '=', 'datetime', '', '', 8, '2024-11-05 17:53:14', '2024-11-05 17:53:14');
INSERT INTO `x_gen_table_column` VALUES (37, 4, 'id', '错误id', '11', 'int', 'int', 'id', 1, 1, 1, 0, 1, 0, 0, '=', 'number', '', '', 1, '2024-11-06 12:13:50', '2024-11-06 12:13:50');
INSERT INTO `x_gen_table_column` VALUES (38, 4, 'project_key', '项目key', '128', 'varchar', 'string', 'project_key', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', '', 2, '2024-11-06 12:13:50', '2024-11-06 12:13:50');
INSERT INTO `x_gen_table_column` VALUES (39, 4, 'client_id', 'sdk生成的客户端id', '128', 'varchar', 'string', 'client_id', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', '', 3, '2024-11-06 12:13:50', '2024-11-06 12:13:50');
INSERT INTO `x_gen_table_column` VALUES (40, 4, 'user_id', '用户id', '128', 'varchar', 'string', 'user_id', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', '', 4, '2024-11-06 12:13:50', '2024-11-06 12:13:50');
INSERT INTO `x_gen_table_column` VALUES (41, 4, 'path', 'URL地址', '1000', 'varchar', 'string', 'path', 0, 0, 1, 1, 1, 1, 1, '=', 'textarea', '', '', 5, '2024-11-06 12:13:50', '2024-11-06 12:13:50');
INSERT INTO `x_gen_table_column` VALUES (42, 4, 'time', '时间', '0', 'float', 'float64', 'time', 0, 0, 1, 1, 1, 1, 1, '=', 'number', '', '', 6, '2024-11-06 12:13:50', '2024-11-06 12:13:50');
INSERT INTO `x_gen_table_column` VALUES (43, 4, 'create_time', '创建时间', '0', 'datetime', 'core.NullTime', 'create_time', 0, 0, 0, 0, 0, 1, 1, '=', 'datetime', '', '', 7, '2024-11-06 12:13:50', '2024-11-06 12:13:50');
INSERT INTO `x_gen_table_column` VALUES (44, 5, 'id', '', '11', 'int', 'int', 'id', 1, 1, 1, 0, 1, 0, 0, '=', 'number', '', '', 1, '2024-11-08 14:12:22', '2024-11-08 14:12:22');
INSERT INTO `x_gen_table_column` VALUES (45, 5, 'scene', '场景编号', '10', 'int', 'int', 'scene', 0, 0, 1, 1, 1, 1, 1, '=', 'number', '', '', 2, '2024-11-08 14:12:22', '2024-11-08 14:12:22');
INSERT INTO `x_gen_table_column` VALUES (46, 5, 'name', '场景名称', '100', 'varchar', 'string', 'name', 0, 0, 1, 1, 1, 1, 1, 'LIKE', 'input', '', '', 3, '2024-11-08 14:12:22', '2024-11-08 14:12:22');
INSERT INTO `x_gen_table_column` VALUES (47, 5, 'remarks', '场景描述', '200', 'varchar', 'string', 'remarks', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', '', 4, '2024-11-08 14:12:22', '2024-11-08 14:12:22');
INSERT INTO `x_gen_table_column` VALUES (48, 5, 'recipient', '接收人员: [1=用户, 2=平台]', '1', 'tinyint', 'int', 'recipient', 0, 0, 1, 1, 1, 1, 1, '=', 'number', '', '', 5, '2024-11-08 14:12:22', '2024-11-08 14:12:22');
INSERT INTO `x_gen_table_column` VALUES (49, 5, 'type', '通知类型: [1=业务, 2=验证]', '1', 'tinyint', 'int', 'type', 0, 0, 1, 1, 1, 1, 1, '=', 'select', '', '', 6, '2024-11-08 14:12:22', '2024-11-08 14:12:22');
INSERT INTO `x_gen_table_column` VALUES (50, 5, 'system_notice', '系统的通知设置', '0', 'text', 'string', 'system_notice', 0, 0, 1, 1, 1, 1, 1, '=', 'textarea', '', '', 7, '2024-11-08 14:12:22', '2024-11-08 14:12:22');
INSERT INTO `x_gen_table_column` VALUES (51, 5, 'sms_notice', '短信的通知设置', '0', 'text', 'string', 'sms_notice', 0, 0, 1, 1, 1, 1, 1, '=', 'textarea', '', '', 8, '2024-11-08 14:12:22', '2024-11-08 14:12:22');
INSERT INTO `x_gen_table_column` VALUES (52, 5, 'oa_notice', '公众号通知设置', '0', 'text', 'string', 'oa_notice', 0, 0, 1, 1, 1, 1, 1, '=', 'textarea', '', '', 9, '2024-11-08 14:12:22', '2024-11-08 14:12:22');
INSERT INTO `x_gen_table_column` VALUES (53, 5, 'mnp_notice', '小程序通知设置', '0', 'text', 'string', 'mnp_notice', 0, 0, 1, 1, 1, 1, 1, '=', 'textarea', '', '', 10, '2024-11-08 14:12:22', '2024-11-08 14:12:22');
INSERT INTO `x_gen_table_column` VALUES (54, 5, 'is_delete', '是否删除', '1', 'tinyint', 'int', 'is_delete', 0, 0, 0, 0, 0, 0, 0, '=', 'number', '', '', 11, '2024-11-08 14:12:22', '2024-11-08 14:12:22');
INSERT INTO `x_gen_table_column` VALUES (55, 5, 'create_time', '创建时间', '0', 'datetime', 'core.NullTime', 'create_time', 0, 0, 0, 0, 0, 1, 1, '=', 'datetime', '', '', 12, '2024-11-08 14:12:22', '2024-11-08 14:12:22');
INSERT INTO `x_gen_table_column` VALUES (56, 5, 'update_time', '更新时间', '0', 'datetime', 'core.NullTime', 'update_time', 0, 0, 0, 0, 0, 1, 1, '=', 'datetime', '', '', 13, '2024-11-08 14:12:22', '2024-11-08 14:12:22');
INSERT INTO `x_gen_table_column` VALUES (57, 5, 'delete_time', '删除时间', '0', 'datetime', 'core.NullTime', 'delete_time', 0, 0, 0, 0, 0, 0, 0, '=', 'datetime', '', '', 14, '2024-11-08 14:12:22', '2024-11-08 14:12:22');

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
  `country` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '国家',
  `province` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '省份',
  `city` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '城市',
  `operator` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '电信运营商',
  `ip` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'ip',
  `width` smallint(10) UNSIGNED NULL DEFAULT 0 COMMENT '屏幕',
  `height` smallint(10) UNSIGNED NULL DEFAULT 0 COMMENT '屏幕高度',
  `ua` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'ua记录',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `project_key`(`project_key`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  INDEX `client_id`(`client_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 84 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '监控-客户端信息' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_monitor_client
-- ----------------------------
INSERT INTO `x_monitor_client` VALUES (37, 'e19e3be20de94f49b68fafb4c30668bc', '81949910-96ae-11ef-9c27-e79173e89fc1', '10', 'Windows', 'Edge', '', '', '', '', '127.0.0.1', 1536, 727, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0', '2024-10-31 12:55:58');
INSERT INTO `x_monitor_client` VALUES (38, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '10', 'Windows', 'Edge', '', '', '', '', '127.0.0.1', 1536, 727, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0', '2024-11-04 19:05:47');
INSERT INTO `x_monitor_client` VALUES (39, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '10', 'Windows', 'Edge', '', '', '', '', '127.0.0.1', 994, 727, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0', '2024-11-05 12:12:19');
INSERT INTO `x_monitor_client` VALUES (40, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '10', 'Windows', 'Edge', '', '', '', '', '127.0.0.1', 715, 727, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0', '2024-11-05 14:36:42');
INSERT INTO `x_monitor_client` VALUES (41, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '10', 'Windows', 'Edge', '', '', '', '', '127.0.0.1', 715, 7270, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0', '2024-11-05 15:07:15');
INSERT INTO `x_monitor_client` VALUES (42, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '10', 'Windows', 'Edge', '', '', '', '', '127.0.0.1', 715, 72, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0', '2024-11-05 15:08:08');
INSERT INTO `x_monitor_client` VALUES (43, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '10', 'Windows', 'Edge', '', '', '', '', '127.0.0.1', 715, 782, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0', '2024-11-05 15:08:54');
INSERT INTO `x_monitor_client` VALUES (44, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '10', 'Windows', 'Edge', '', '', '', '', '127.0.0.1', 715, 7822, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0', '2024-11-05 15:10:25');
INSERT INTO `x_monitor_client` VALUES (45, 'e19e3be20de94f49b68fafb4c30668bc', '81949910-96ae-11ef-9c27-e79173e89fc1', '10', 'Windows', 'Edge', '', '', '北京', '', '', 1536, 727, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0', '2024-10-31 12:55:58');
INSERT INTO `x_monitor_client` VALUES (46, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '10', 'Windows', 'Edge', '', '', '北京', '', '', 1536, 727, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0', '2024-11-04 19:05:47');
INSERT INTO `x_monitor_client` VALUES (47, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '10', 'Windows', 'Edge', '', '', '天津', '', '', 994, 727, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0', '2024-11-05 12:12:19');
INSERT INTO `x_monitor_client` VALUES (48, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '10', 'Windows', 'Edge', '', '', '上海', '', '', 715, 727, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0', '2024-11-05 14:36:42');
INSERT INTO `x_monitor_client` VALUES (49, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '10', 'Windows', 'Edge', '', '', '', '', '', 715, 7270, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0', '2024-11-05 15:07:15');
INSERT INTO `x_monitor_client` VALUES (50, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '10', 'Windows', 'Edge', '', '', '', '', '', 715, 72, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0', '2024-11-05 15:08:08');
INSERT INTO `x_monitor_client` VALUES (51, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '10', 'Windows', 'Edge', '', '', '', '', '', 715, 782, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0', '2024-11-05 15:08:54');
INSERT INTO `x_monitor_client` VALUES (52, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '10', 'Windows', 'Edge', '', '', '', '', '', 715, 7822, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0', '2024-11-05 15:10:25');
INSERT INTO `x_monitor_client` VALUES (53, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '10', 'Windows', 'Edge', '', '', '', '', '127.0.0.1', 1536, 727, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0', '2024-11-05 17:00:35');
INSERT INTO `x_monitor_client` VALUES (54, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '10', 'Windows', 'Edge', '', '', '', '', '127.0.0.1', 934, 727, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0', '2024-11-05 17:29:02');
INSERT INTO `x_monitor_client` VALUES (55, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '10', 'Windows', 'Edge', '', '', '', '', '127.0.0.1', 1536, 727, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0', '2024-11-05 17:51:13');
INSERT INTO `x_monitor_client` VALUES (56, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '10', 'Windows', 'Edge', '', '', '', '', '127.0.0.1', 934, 727, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0', '2024-11-05 19:13:29');
INSERT INTO `x_monitor_client` VALUES (57, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '10', 'Windows', 'Edge', '', '', '', '', '127.0.0.1', 1536, 727, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0', '2024-11-05 19:20:40');
INSERT INTO `x_monitor_client` VALUES (58, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '10', 'Windows', 'Edge', '', '', '', '', '127.0.0.1', 934, 727, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0', '2024-11-05 19:26:27');
INSERT INTO `x_monitor_client` VALUES (59, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '10', 'Windows', 'Edge', '', '', '', '', '127.0.0.1', 1536, 727, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0', '2024-11-05 19:27:11');
INSERT INTO `x_monitor_client` VALUES (60, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '10', 'Windows', 'Edge', '', '', '', '', '127.0.0.1', 934, 727, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0', '2024-11-05 23:37:45');
INSERT INTO `x_monitor_client` VALUES (61, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '', 'Windows', 'Edge', '', '', '', '', '127.0.0.1', 934, 727, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0', '2024-11-05 23:40:30');
INSERT INTO `x_monitor_client` VALUES (62, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '', 'Windows', 'Edge', '', '', '', '', '127.0.0.1', 1536, 727, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0', '2024-11-06 11:13:42');
INSERT INTO `x_monitor_client` VALUES (63, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '', 'Windows', 'Edge', '', '', '', '', '127.0.0.1', 1076, 727, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0', '2024-11-06 11:45:21');
INSERT INTO `x_monitor_client` VALUES (64, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '', 'Windows', 'Edge', '', '', '', '', '127.0.0.1', 854, 727, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0', '2024-11-06 11:46:20');
INSERT INTO `x_monitor_client` VALUES (65, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '', 'Windows', 'Edge', '', '', '', '', '127.0.0.1', 606, 727, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0', '2024-11-06 11:48:05');
INSERT INTO `x_monitor_client` VALUES (66, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '', 'Windows', 'Edge', '', '', '', '', '127.0.0.1', 1536, 727, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0', '2024-11-06 11:57:55');
INSERT INTO `x_monitor_client` VALUES (67, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '', 'Windows', 'Edge', '', '', '', '', '127.0.0.1', 704, 727, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0', '2024-11-06 12:39:47');
INSERT INTO `x_monitor_client` VALUES (68, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '', 'Windows', 'Edge', '', '', '', '', '127.0.0.1', 1536, 727, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0', '2024-11-06 12:41:11');
INSERT INTO `x_monitor_client` VALUES (69, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '', 'Windows', 'Edge', '', '', '', '', '127.0.0.1', 1011, 727, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0', '2024-11-07 16:31:35');
INSERT INTO `x_monitor_client` VALUES (70, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '', 'Windows', 'Edge', '', '', '', '', '127.0.0.1', 1059, 727, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0', '2024-11-07 16:32:59');
INSERT INTO `x_monitor_client` VALUES (71, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '', 'Windows', 'Edge', '', '', '', '', '127.0.0.1', 667, 727, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0', '2024-11-07 16:34:46');
INSERT INTO `x_monitor_client` VALUES (72, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '', 'Windows', 'Edge', '', '', '', '', '127.0.0.1', 1536, 727, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0', '2024-11-07 16:49:47');
INSERT INTO `x_monitor_client` VALUES (73, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '1', 'Windows', 'Edge', '', '', '', '', '127.0.0.1', 1536, 727, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0', '2024-11-07 16:52:00');
INSERT INTO `x_monitor_client` VALUES (74, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '', 'Windows', 'Edge', '', '', '', '', '127.0.0.1', 1536, 727, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0', '2024-11-07 16:52:00');
INSERT INTO `x_monitor_client` VALUES (75, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '', 'Windows', 'Edge', '', '', '', '', '127.0.0.1', 1536, 727, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0', '2024-11-07 16:52:00');
INSERT INTO `x_monitor_client` VALUES (76, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '1', 'Windows', 'Edge', '', '', '', '', '127.0.0.1', 1536, 727, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0', '2024-11-07 16:52:10');
INSERT INTO `x_monitor_client` VALUES (77, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '1', 'Windows', 'Edge', '', '', '', '', '127.0.0.1', 903, 727, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0', '2024-11-07 17:15:07');
INSERT INTO `x_monitor_client` VALUES (78, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '1', 'Windows', 'Edge', '', '', '', '', '127.0.0.1', 1536, 332, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0', '2024-11-07 17:24:01');
INSERT INTO `x_monitor_client` VALUES (79, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '1', 'Windows', 'Edge', '', '', '', '', '127.0.0.1', 1536, 362, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0', '2024-11-07 18:26:57');
INSERT INTO `x_monitor_client` VALUES (80, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '2', 'Windows', 'Edge', '', '', '', '', '127.0.0.1', 1536, 362, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0', '2024-11-07 18:26:57');
INSERT INTO `x_monitor_client` VALUES (81, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '2', 'Windows', 'Edge', '', '', '', '', '127.0.0.1', 1536, 727, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0', '2024-11-07 18:33:15');
INSERT INTO `x_monitor_client` VALUES (82, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '2', 'Windows', 'Edge', '', '', '', '', '127.0.0.1', 1536, 362, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0', '2024-11-08 14:05:06');
INSERT INTO `x_monitor_client` VALUES (83, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '2', 'Windows', 'Edge', '', '', '', '', '127.0.0.1', 1536, 727, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0', '2024-11-08 14:18:55');

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
  `message` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '错误消息',
  `stack` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '错误堆栈',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `project_key`(`project_key`, `md5`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 155 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '监控-错误列表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_monitor_error
-- ----------------------------
INSERT INTO `x_monitor_error` VALUES (93, 'e19e3be20de94f49b68fafb4c30668bc', '02fd9c50b0dfcf319577a9e47e693b18', 'error', 'http://localhost:5174/dev_tools/error/test', 'Uncaught Error: 主动抛error', 'Error: 主动抛error\n    at error (http://localhost:5174/src/views/error/test.vue?t=1730367699656:24:13)\n    at callWithErrorHandling (http://localhost:5174/node_modules/.vite/deps/chunk-W2D5EV7B.js?v=2db2a83d:2477:19)\n    at callWithAsyncErrorHandling (http://localhost:5174/node_modules/.vite/deps/chunk-W2D5EV7B.js?v=2db2a83d:2484:17)', '2024-10-31 17:44:37');
INSERT INTO `x_monitor_error` VALUES (94, 'e19e3be20de94f49b68fafb4c30668bc', 'aefe6052f362cff581c583a0b1220210', 'error', 'http://localhost:5174/dev_tools/error/test', 'Uncaught Error: 主动抛error', 'Error: 主动抛error\n    at error (http://localhost:5174/src/views/error/test.vue?t=1730718341248:23:13)\n    at callWithErrorHandling (http://localhost:5174/node_modules/.vite/deps/chunk-W2D5EV7B.js?v=2db2a83d:2477:19)\n    at callWithAsyncErrorHandling (http://localhost:5174/node_modules/.vite/deps/chunk-W2D5EV7B.js?v=2db2a83d:2484:17)', '2024-11-04 19:06:07');
INSERT INTO `x_monitor_error` VALUES (95, 'e19e3be20de94f49b68fafb4c30668bc', 'd92f2036acfb736fb91b46419bb350ae', 'error', 'http://localhost:5174/dev_tools/error/test', 'Uncaught Error: 主动抛error', 'Error: 主动抛error\n    at error (http://localhost:5174/src/views/error/test.vue?t=1730778883476:10:13)\n    at callWithErrorHandling (http://localhost:5174/node_modules/.vite/deps/chunk-W2D5EV7B.js?v=2db2a83d:2477:19)\n    at callWithAsyncErrorHandling (http://localhost:5174/node_modules/.vite/deps/chunk-W2D5EV7B.js?v=2db2a83d:2484:17)', '2024-11-05 14:07:02');
INSERT INTO `x_monitor_error` VALUES (96, 'e19e3be20de94f49b68fafb4c30668bc', '30d04ccb00905006fcc3308541250c67', 'error', 'http://localhost:5174/dev_tools/error/test', 'Uncaught Error: 主动抛error', 'Error: 主动抛error\n    at error (http://localhost:5174/src/views/error/test.vue?t=1730786719985:10:13)\n    at callWithErrorHandling (http://localhost:5174/node_modules/.vite/deps/chunk-W2D5EV7B.js?v=2db2a83d:2477:19)\n    at callWithAsyncErrorHandling (http://localhost:5174/node_modules/.vite/deps/chunk-W2D5EV7B.js?v=2db2a83d:2484:17)', '2024-11-05 14:48:05');
INSERT INTO `x_monitor_error` VALUES (97, 'e19e3be20de94f49b68fafb4c30668bc', 'a5c695f6bcbee6ed4fc15f438dbe750b', 'error', 'http://localhost:5174/monitor/project/index', 'ResizeObserver loop completed with undelivered notifications.', '', '2024-11-05 17:27:34');
INSERT INTO `x_monitor_error` VALUES (113, 'e19e3be20de94f49b68fafb4c30668bc', '5dbe27c63610e9fc315a047041bab6a7', 'whiteScreen', 'http://localhost:5174/monitor/error/index', '白屏时间：0ms', '', '2024-11-05 19:25:09');
INSERT INTO `x_monitor_error` VALUES (114, 'e19e3be20de94f49b68fafb4c30668bc', '987445c92a18acfab2493eeb881631a0', 'whiteScreen', 'http://localhost:5174/monitor/error/index', '白屏时间：559.9000000357628ms', '', '2024-11-05 19:26:00');
INSERT INTO `x_monitor_error` VALUES (115, 'e19e3be20de94f49b68fafb4c30668bc', '16bd6d91c4a212d2cfac903514da2357', 'whiteScreen', 'http://localhost:5174/monitor/error/index', '白屏时间：368.19999998807907ms', '', '2024-11-05 19:26:11');
INSERT INTO `x_monitor_error` VALUES (116, 'e19e3be20de94f49b68fafb4c30668bc', '1b2553f51f9ab7791e7162219ee2568f', 'error', 'http://localhost:5174/monitor/error/index', 'ResizeObserver loop completed with undelivered notifications.', '', '2024-11-05 19:26:21');
INSERT INTO `x_monitor_error` VALUES (117, 'e19e3be20de94f49b68fafb4c30668bc', '6275faaaa0d45dda8fad3e60b6156cee', 'whiteScreen', 'http://localhost:5174/monitor/error/index', '白屏时间：2393.300000011921ms', '', '2024-11-05 19:26:37');
INSERT INTO `x_monitor_error` VALUES (118, 'e19e3be20de94f49b68fafb4c30668bc', '822ea1ded1482eb8ca8d798b44c56cc7', 'whiteScreen', 'http://localhost:5174/monitor/error/index', '白屏时间：547.1999999880791ms', '', '2024-11-05 19:27:25');
INSERT INTO `x_monitor_error` VALUES (119, 'e19e3be20de94f49b68fafb4c30668bc', 'b51eb156afdee7850a0376e510ce5ead', 'whiteScreen', 'http://localhost:5174/monitor/error/index', '白屏时间：635.5ms', '', '2024-11-05 19:27:25');
INSERT INTO `x_monitor_error` VALUES (120, 'e19e3be20de94f49b68fafb4c30668bc', 'c1ce04e9775a09f4298f556bf11179f2', 'whiteScreen', 'http://localhost:5174/monitor/error/index', '白屏时间：438.10000002384186ms', '', '2024-11-05 19:27:43');
INSERT INTO `x_monitor_error` VALUES (121, 'e19e3be20de94f49b68fafb4c30668bc', 'cb462e986c3872255ee4ef536fb93c32', 'whiteScreen', 'http://localhost:5174/monitor/error/index', '白屏时间：444.39999997615814ms', '', '2024-11-05 19:27:43');
INSERT INTO `x_monitor_error` VALUES (122, 'e19e3be20de94f49b68fafb4c30668bc', 'bdbb8d8362d7c910e13fe2b5ed365824', 'onloadTime', 'http://localhost:5174/monitor/error/index', '白屏时间：435.7000000476837ms', '', '2024-11-05 19:28:02');
INSERT INTO `x_monitor_error` VALUES (123, 'e19e3be20de94f49b68fafb4c30668bc', '97d0da90fade3a5b55063c5cc830a12f', 'onloadTime', 'http://localhost:5174/monitor/error/index', 'onload时间：540.8999999761581ms', '', '2024-11-05 19:28:28');
INSERT INTO `x_monitor_error` VALUES (124, 'e19e3be20de94f49b68fafb4c30668bc', '287a31dc5083d1f028b77c3510188780', 'onloadTime', 'http://localhost:5174/monitor/error/index', '时间：456.60000002384186ms', '', '2024-11-05 19:28:52');
INSERT INTO `x_monitor_error` VALUES (125, 'e19e3be20de94f49b68fafb4c30668bc', '98592ff4989ccaf00a1bc30dc49dd519', 'onloadTime', 'http://localhost:5174/monitor/error/index', '时间：500.0999999642372ms', '', '2024-11-05 19:29:55');
INSERT INTO `x_monitor_error` VALUES (126, 'e19e3be20de94f49b68fafb4c30668bc', 'b353c4513002e844462d28ba134c06dc', 'onloadTime', 'http://localhost:5174/monitor/error/index', '时间：460.7ms', '', '2024-11-05 19:30:21');
INSERT INTO `x_monitor_error` VALUES (127, 'e19e3be20de94f49b68fafb4c30668bc', '7aaeeb95968a3dcaed20f9a3bf2af9df', 'onloadTime', 'http://localhost:5174/monitor/error/index', '时间：651.8ms', '', '2024-11-05 22:29:58');
INSERT INTO `x_monitor_error` VALUES (128, 'e19e3be20de94f49b68fafb4c30668bc', '34693ff0ad3c53cb32ce1063a99e0aaa', 'onloadTime', 'http://localhost:5174/monitor/error/index', '时间：940.8ms', '', '2024-11-05 22:31:17');
INSERT INTO `x_monitor_error` VALUES (129, 'e19e3be20de94f49b68fafb4c30668bc', 'e22326fc31689e7419add91c837c2218', 'onloadTime', 'http://localhost:5174/monitor/error/index', '时间：1064.3ms', '', '2024-11-05 22:31:51');
INSERT INTO `x_monitor_error` VALUES (130, 'e19e3be20de94f49b68fafb4c30668bc', 'f32c8682bc033e759f6047ccf7d20c4b', 'onloadTime', 'http://localhost:5174/workbench', '时间：3101.6ms', '', '2024-11-05 23:40:46');
INSERT INTO `x_monitor_error` VALUES (131, 'e19e3be20de94f49b68fafb4c30668bc', 'afcf8d7d254a40d772a24a380c3775df', 'onloadTime', 'http://localhost:5174/', '时间：10721.7ms', '', '2024-11-06 11:13:52');
INSERT INTO `x_monitor_error` VALUES (132, 'e19e3be20de94f49b68fafb4c30668bc', 'b4f43b1db77fcb30def44f9a7893ebcc', 'onloadTime', 'http://localhost:5174/monitor/error/index', '时间：8199.9ms', '', '2024-11-06 11:46:49');
INSERT INTO `x_monitor_error` VALUES (133, 'e19e3be20de94f49b68fafb4c30668bc', '0c3f557b31e895fd3ded693cb0ddd847', 'onloadTime', 'http://localhost:5174/monitor/error/index', '时间：8441.3ms', '', '2024-11-06 11:47:34');
INSERT INTO `x_monitor_error` VALUES (134, 'e19e3be20de94f49b68fafb4c30668bc', '99e4b9a6c4f89b036ca5d9bfa7b8ce40', 'onloadTime', 'http://localhost:5174/monitor/error/index', '时间：8188ms', '', '2024-11-06 11:48:15');
INSERT INTO `x_monitor_error` VALUES (135, 'e19e3be20de94f49b68fafb4c30668bc', 'd63648211078130f9e9f558158e86192', 'unhandledrejection', 'http://localhost:5174/monitor/monitor/slow/index', 'nextTick is not defined', 'ReferenceError: nextTick is not defined\n    at Proxy.handleAdd (http://localhost:5174/src/views/monitor/slow/index.vue:35:7)\n    at _createBlock.onClick._cache.<computed>._cache.<computed> (http://localhost:5174/src/views/monitor/slow/index.vue:218:67)\n    at callWithErrorHandling (http://localhost:5174/node_modules/.vite/deps/chunk-W2D5EV7B.js?v=2db2a83d:2477:19)', '2024-11-06 12:37:43');
INSERT INTO `x_monitor_error` VALUES (136, 'e19e3be20de94f49b68fafb4c30668bc', 'ff93a0dcd262e61a52151ef662d75353', 'error', 'http://localhost:5174/monitor/monitor/slow/index', 'ResizeObserver loop completed with undelivered notifications.', '', '2024-11-06 12:37:43');
INSERT INTO `x_monitor_error` VALUES (137, 'e19e3be20de94f49b68fafb4c30668bc', 'ba922c6841afabfb0e1d5eb1a90f873a', 'unhandledrejection', 'http://localhost:5174/monitor/monitor/slow/index', '', '', '2024-11-06 12:48:11');
INSERT INTO `x_monitor_error` VALUES (138, 'e19e3be20de94f49b68fafb4c30668bc', '11b6b34938947c386f816deb4ad36c5f', 'unhandledrejection', 'http://localhost:5174/monitor/monitor/slow/index', 'Request failed with status code 500', 'AxiosError: Request failed with status code 500\n    at settle (http://localhost:5174/node_modules/.vite/deps/axios.js?v=2db2a83d:1230:12)\n    at XMLHttpRequest.onloadend (http://localhost:5174/node_modules/.vite/deps/axios.js?v=2db2a83d:1593:7)\n    at Axios.request (http://localhost:5174/node_modules/.vite/deps/axios.js?v=2db2a83d:2145:41)', '2024-11-06 12:52:01');
INSERT INTO `x_monitor_error` VALUES (139, 'e19e3be20de94f49b68fafb4c30668bc', 'ac41faabe5c45af29b696b37119b7389', 'unhandledrejection', 'http://localhost:5174/monitor/monitor/slow/index', 'timeout of 120000ms exceeded', 'AxiosError: timeout of 120000ms exceeded\n    at XMLHttpRequest.handleTimeout (http://localhost:5174/node_modules/.vite/deps/axios.js?v=2db2a83d:1632:14)\n    at Axios.request (http://localhost:5174/node_modules/.vite/deps/axios.js?v=2db2a83d:2145:41)', '2024-11-06 13:15:37');
INSERT INTO `x_monitor_error` VALUES (140, 'e19e3be20de94f49b68fafb4c30668bc', 'b57c3d565abfba586e8a617649b592a1', 'onloadTime', 'http://localhost:5174/', '时间：13914.8ms', '', '2024-11-07 12:32:39');
INSERT INTO `x_monitor_error` VALUES (141, 'e19e3be20de94f49b68fafb4c30668bc', 'ee461d81b0ef6728a89640be5f794675', 'onloadTime', 'http://localhost:5174/', '时间：5872.9ms', '', '2024-11-07 14:46:41');
INSERT INTO `x_monitor_error` VALUES (142, 'e19e3be20de94f49b68fafb4c30668bc', '3d201a726fbb323c5fc792acb9854884', 'img', 'http://localhost:5174/workbench', '', '', '2024-11-07 15:17:30');
INSERT INTO `x_monitor_error` VALUES (143, 'e19e3be20de94f49b68fafb4c30668bc', 'ed7d8ba6b8ca8dc9e3d5a51258aca428', 'img', 'http://localhost:5174/api/static/backend_backdrop.png', '', '', '2024-11-07 16:37:36');
INSERT INTO `x_monitor_error` VALUES (144, 'e19e3be20de94f49b68fafb4c30668bc', '0b520372ea4f67141414e50085c44951', 'unhandledrejection', 'http://localhost:5174/login?redirect=/monitor/monitor/slow/index', 'Request failed with status code 500', 'AxiosError: Request failed with status code 500\n    at settle (http://localhost:5174/node_modules/.vite/deps/axios.js?v=2db2a83d:1230:12)\n    at XMLHttpRequest.onloadend (http://localhost:5174/node_modules/.vite/deps/axios.js?v=2db2a83d:1593:7)\n    at Axios.request (http://localhost:5174/node_modules/.vite/deps/axios.js?v=2db2a83d:2145:41)', '2024-11-07 16:40:46');
INSERT INTO `x_monitor_error` VALUES (145, 'e19e3be20de94f49b68fafb4c30668bc', 'c66cf37bb5dc338b63678d00a7054ecf', 'img', 'http://localhost:5174/login?redirect=/monitor/monitor/slow/index', '', '', '2024-11-07 16:40:46');
INSERT INTO `x_monitor_error` VALUES (146, 'e19e3be20de94f49b68fafb4c30668bc', 'b4ab03a2fd2c320812854d55fe6abb4f', 'unhandledrejection', 'http://localhost:5174/monitor/monitor/slow/index', 'Cannot read properties of undefined (reading \'monitor_project_listAll\')', 'TypeError: Cannot read properties of undefined (reading \'monitor_project_listAll\')\n    at http://localhost:5174/src/views/monitor/slow/index.vue?t=1730969588901:148:52\n    at renderFnWithContext (http://localhost:5174/node_modules/.vite/deps/chunk-W2D5EV7B.js?v=2db2a83d:2926:13)\n    at renderSlot (http://localhost:5174/node_modules/.vite/deps/chunk-W2D5EV7B.js?v=2db2a83d:4537:53)', '2024-11-07 16:53:10');
INSERT INTO `x_monitor_error` VALUES (147, 'e19e3be20de94f49b68fafb4c30668bc', '6697e34d8798e9c5bfb8edce0b869870', 'unhandledrejection', 'http://localhost:5174/monitor/monitor/slow/index', 'Cannot read properties of null (reading \'parentNode\')', 'TypeError: Cannot read properties of null (reading \'parentNode\')\n    at parentNode (http://localhost:5174/node_modules/.vite/deps/chunk-W2D5EV7B.js?v=2db2a83d:11830:34)\n    at ReactiveEffect.componentUpdateFn [as fn] (http://localhost:5174/node_modules/.vite/deps/chunk-W2D5EV7B.js?v=2db2a83d:6633:11)\n    at ReactiveEffect.run (http://localhost:5174/node_modules/.vite/deps/chunk-W2D5EV7B.js?v=2db2a83d:1770:23)', '2024-11-07 16:53:50');
INSERT INTO `x_monitor_error` VALUES (148, 'e19e3be20de94f49b68fafb4c30668bc', 'a359051a7d0903201728fe227bf8d091', 'unhandledrejection', 'http://localhost:5174/monitor/monitor/slow/index', 'Request failed with status code 404', 'AxiosError: Request failed with status code 404\n    at settle (http://localhost:5174/node_modules/.vite/deps/axios.js?v=2db2a83d:1230:12)\n    at XMLHttpRequest.onloadend (http://localhost:5174/node_modules/.vite/deps/axios.js?v=2db2a83d:1593:7)\n    at Axios.request (http://localhost:5174/node_modules/.vite/deps/axios.js?v=2db2a83d:2145:41)', '2024-11-07 16:54:55');
INSERT INTO `x_monitor_error` VALUES (149, 'e19e3be20de94f49b68fafb4c30668bc', 'f47f07c13732e818c84dc087b7169615', 'unhandledrejection', 'http://localhost:5174/monitor/monitor/slow/index', 'Cannot read properties of undefined (reading \'Vue\')', 'TypeError: Cannot read properties of undefined (reading \'Vue\')\n    at initBackend (chrome-extension://olofadcdnkkjdfgjcmjaadnlehnnihnl/build/backend.js:2119:28)', '2024-11-07 17:28:58');
INSERT INTO `x_monitor_error` VALUES (150, 'e19e3be20de94f49b68fafb4c30668bc', 'fe7b096876be7f647d94ce84a053dab9', 'script', 'http://localhost:5174/src/main.ts?t=1730975599285', '', '', '2024-11-08 13:27:15');
INSERT INTO `x_monitor_error` VALUES (151, 'e19e3be20de94f49b68fafb4c30668bc', '58a3679c7266f040d3e8116412c22ed8', 'script', 'http://localhost:5174/@vite/client', '', '', '2024-11-08 14:03:52');
INSERT INTO `x_monitor_error` VALUES (152, 'e19e3be20de94f49b68fafb4c30668bc', 'f970e7ec49c1227292840651f9abb7b6', 'script', 'http://localhost:5174/src/main.ts?t=1731043835935', '', '', '2024-11-08 14:03:52');
INSERT INTO `x_monitor_error` VALUES (153, 'e19e3be20de94f49b68fafb4c30668bc', '1079deb09a1e4fd73b20649d73e9dd4d', 'unhandledrejection', 'http://localhost:5174/setting/system/environment', 'Network Error', 'AxiosError: Network Error\n    at XMLHttpRequest.handleError (http://localhost:5174/node_modules/.vite/deps/axios.js?v=2db2a83d:1623:14)\n    at Axios.request (http://localhost:5174/node_modules/.vite/deps/axios.js?v=2db2a83d:2145:41)', '2024-11-08 15:09:40');
INSERT INTO `x_monitor_error` VALUES (154, 'e19e3be20de94f49b68fafb4c30668bc', '3858135af1f6deb9079a34b7a5b42440', 'error', 'http://localhost:5174/dev_tools/error/test', 'Uncaught Error: 主动抛error', 'Error: 主动抛error\n    at error (http://localhost:5174/src/views/error/test.vue:10:13)\n    at callWithErrorHandling (http://localhost:5174/node_modules/.vite/deps/chunk-W2D5EV7B.js?v=2db2a83d:2477:19)\n    at callWithAsyncErrorHandling (http://localhost:5174/node_modules/.vite/deps/chunk-W2D5EV7B.js?v=2db2a83d:2484:17)', '2024-11-08 16:44:27');

-- ----------------------------
-- Table structure for x_monitor_error_list
-- ----------------------------
DROP TABLE IF EXISTS `x_monitor_error_list`;
CREATE TABLE `x_monitor_error_list`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `eid` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '错误表id',
  `cid` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '客户端表id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `uid`(`cid`) USING BTREE,
  INDEX `eid`(`eid`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 404 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '错误对应的用户记录' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of x_monitor_error_list
-- ----------------------------
INSERT INTO `x_monitor_error_list` VALUES (238, '84', '37', '2024-10-31 13:05:54');
INSERT INTO `x_monitor_error_list` VALUES (239, '85', '37', '2024-10-31 13:05:54');
INSERT INTO `x_monitor_error_list` VALUES (240, '84', '37', '2024-10-31 13:05:54');
INSERT INTO `x_monitor_error_list` VALUES (241, '84', '37', '2024-10-31 13:05:54');
INSERT INTO `x_monitor_error_list` VALUES (242, '84', '37', '2024-10-31 13:05:54');
INSERT INTO `x_monitor_error_list` VALUES (243, '84', '37', '2024-10-31 13:05:54');
INSERT INTO `x_monitor_error_list` VALUES (244, '84', '37', '2024-10-31 13:05:57');
INSERT INTO `x_monitor_error_list` VALUES (245, '84', '37', '2024-10-31 13:05:57');
INSERT INTO `x_monitor_error_list` VALUES (246, '84', '37', '2024-10-31 13:05:57');
INSERT INTO `x_monitor_error_list` VALUES (247, '84', '37', '2024-10-31 13:05:57');
INSERT INTO `x_monitor_error_list` VALUES (248, '84', '37', '2024-10-31 13:05:57');
INSERT INTO `x_monitor_error_list` VALUES (249, '84', '37', '2024-10-31 13:05:57');
INSERT INTO `x_monitor_error_list` VALUES (250, '84', '37', '2024-10-31 13:06:07');
INSERT INTO `x_monitor_error_list` VALUES (251, '84', '37', '2024-10-31 13:06:28');
INSERT INTO `x_monitor_error_list` VALUES (252, '84', '37', '2024-10-31 13:06:28');
INSERT INTO `x_monitor_error_list` VALUES (253, '85', '37', '2024-10-31 13:07:01');
INSERT INTO `x_monitor_error_list` VALUES (254, '86', '37', '2024-10-31 13:07:01');
INSERT INTO `x_monitor_error_list` VALUES (255, '87', '37', '2024-10-31 13:07:12');
INSERT INTO `x_monitor_error_list` VALUES (256, '84', '37', '2024-10-31 13:11:48');
INSERT INTO `x_monitor_error_list` VALUES (257, '84', '37', '2024-10-31 13:11:48');
INSERT INTO `x_monitor_error_list` VALUES (258, '88', '37', '2024-10-31 15:01:25');
INSERT INTO `x_monitor_error_list` VALUES (259, '89', '37', '2024-10-31 15:01:25');
INSERT INTO `x_monitor_error_list` VALUES (260, '88', '37', '2024-10-31 15:19:35');
INSERT INTO `x_monitor_error_list` VALUES (261, '90', '37', '2024-10-31 15:19:35');
INSERT INTO `x_monitor_error_list` VALUES (262, '88', '37', '2024-10-31 15:29:55');
INSERT INTO `x_monitor_error_list` VALUES (263, '90', '37', '2024-10-31 15:29:55');
INSERT INTO `x_monitor_error_list` VALUES (264, '88', '37', '2024-10-31 15:46:27');
INSERT INTO `x_monitor_error_list` VALUES (265, '90', '37', '2024-10-31 15:46:27');
INSERT INTO `x_monitor_error_list` VALUES (266, '88', '37', '2024-10-31 15:46:38');
INSERT INTO `x_monitor_error_list` VALUES (267, '90', '37', '2024-10-31 15:46:38');
INSERT INTO `x_monitor_error_list` VALUES (268, '91', '37', '2024-10-31 15:48:25');
INSERT INTO `x_monitor_error_list` VALUES (269, '92', '37', '2024-10-31 15:48:52');
INSERT INTO `x_monitor_error_list` VALUES (270, '93', '37', '2024-10-31 17:44:37');
INSERT INTO `x_monitor_error_list` VALUES (271, '93', '37', '2024-10-31 17:44:37');
INSERT INTO `x_monitor_error_list` VALUES (272, '94', '38', '2024-11-04 19:06:07');
INSERT INTO `x_monitor_error_list` VALUES (273, '95', '39', '2024-11-05 14:07:03');
INSERT INTO `x_monitor_error_list` VALUES (274, '95', '39', '2024-11-05 14:07:33');
INSERT INTO `x_monitor_error_list` VALUES (275, '95', '39', '2024-11-05 14:07:33');
INSERT INTO `x_monitor_error_list` VALUES (276, '95', '39', '2024-11-05 14:07:33');
INSERT INTO `x_monitor_error_list` VALUES (277, '95', '39', '2024-11-05 14:07:33');
INSERT INTO `x_monitor_error_list` VALUES (278, '95', '39', '2024-11-05 14:07:33');
INSERT INTO `x_monitor_error_list` VALUES (279, '96', '40', '2024-11-05 14:48:05');
INSERT INTO `x_monitor_error_list` VALUES (280, '96', '40', '2024-11-05 14:48:05');
INSERT INTO `x_monitor_error_list` VALUES (281, '96', '40', '2024-11-05 14:48:05');
INSERT INTO `x_monitor_error_list` VALUES (282, '96', '40', '2024-11-05 14:48:13');
INSERT INTO `x_monitor_error_list` VALUES (283, '96', '40', '2024-11-05 14:48:13');
INSERT INTO `x_monitor_error_list` VALUES (284, '96', '40', '2024-11-05 14:48:13');
INSERT INTO `x_monitor_error_list` VALUES (285, '96', '40', '2024-11-05 14:48:15');
INSERT INTO `x_monitor_error_list` VALUES (286, '96', '40', '2024-11-05 14:48:26');
INSERT INTO `x_monitor_error_list` VALUES (287, '96', '40', '2024-11-05 14:48:26');
INSERT INTO `x_monitor_error_list` VALUES (288, '96', '40', '2024-11-05 14:48:26');
INSERT INTO `x_monitor_error_list` VALUES (289, '96', '40', '2024-11-05 14:51:52');
INSERT INTO `x_monitor_error_list` VALUES (290, '96', '40', '2024-11-05 14:51:52');
INSERT INTO `x_monitor_error_list` VALUES (291, '96', '40', '2024-11-05 14:51:52');
INSERT INTO `x_monitor_error_list` VALUES (292, '97', '53', '2024-11-05 17:27:34');
INSERT INTO `x_monitor_error_list` VALUES (293, '97', '54', '2024-11-05 17:29:54');
INSERT INTO `x_monitor_error_list` VALUES (294, '97', '54', '2024-11-05 17:34:04');
INSERT INTO `x_monitor_error_list` VALUES (295, '97', '54', '2024-11-05 17:35:24');
INSERT INTO `x_monitor_error_list` VALUES (296, '98', '55', '2024-11-05 17:53:26');
INSERT INTO `x_monitor_error_list` VALUES (297, '98', '55', '2024-11-05 17:53:26');
INSERT INTO `x_monitor_error_list` VALUES (298, '98', '55', '2024-11-05 17:53:26');
INSERT INTO `x_monitor_error_list` VALUES (299, '98', '55', '2024-11-05 17:53:26');
INSERT INTO `x_monitor_error_list` VALUES (300, '98', '55', '2024-11-05 17:53:36');
INSERT INTO `x_monitor_error_list` VALUES (301, '98', '55', '2024-11-05 17:53:36');
INSERT INTO `x_monitor_error_list` VALUES (302, '99', '55', '2024-11-05 17:57:06');
INSERT INTO `x_monitor_error_list` VALUES (303, '100', '55', '2024-11-05 17:57:16');
INSERT INTO `x_monitor_error_list` VALUES (304, '99', '55', '2024-11-05 17:57:26');
INSERT INTO `x_monitor_error_list` VALUES (305, '101', '55', '2024-11-05 19:00:34');
INSERT INTO `x_monitor_error_list` VALUES (306, '102', '55', '2024-11-05 19:05:15');
INSERT INTO `x_monitor_error_list` VALUES (307, '103', '55', '2024-11-05 19:11:32');
INSERT INTO `x_monitor_error_list` VALUES (308, '104', '55', '2024-11-05 19:12:59');
INSERT INTO `x_monitor_error_list` VALUES (309, '105', '55', '2024-11-05 19:12:59');
INSERT INTO `x_monitor_error_list` VALUES (310, '106', '55', '2024-11-05 19:12:59');
INSERT INTO `x_monitor_error_list` VALUES (311, '107', '55', '2024-11-05 19:13:10');
INSERT INTO `x_monitor_error_list` VALUES (312, '108', '55', '2024-11-05 19:13:20');
INSERT INTO `x_monitor_error_list` VALUES (313, '109', '56', '2024-11-05 19:13:39');
INSERT INTO `x_monitor_error_list` VALUES (314, '108', '56', '2024-11-05 19:13:59');
INSERT INTO `x_monitor_error_list` VALUES (315, '110', '57', '2024-11-05 19:20:53');
INSERT INTO `x_monitor_error_list` VALUES (316, '111', '57', '2024-11-05 19:20:53');
INSERT INTO `x_monitor_error_list` VALUES (317, '112', '57', '2024-11-05 19:24:16');
INSERT INTO `x_monitor_error_list` VALUES (318, '112', '57', '2024-11-05 19:24:38');
INSERT INTO `x_monitor_error_list` VALUES (319, '112', '57', '2024-11-05 19:24:38');
INSERT INTO `x_monitor_error_list` VALUES (320, '112', '57', '2024-11-05 19:24:38');
INSERT INTO `x_monitor_error_list` VALUES (321, '113', '57', '2024-11-05 19:25:09');
INSERT INTO `x_monitor_error_list` VALUES (322, '114', '57', '2024-11-05 19:26:00');
INSERT INTO `x_monitor_error_list` VALUES (323, '115', '57', '2024-11-05 19:26:11');
INSERT INTO `x_monitor_error_list` VALUES (324, '116', '57', '2024-11-05 19:26:21');
INSERT INTO `x_monitor_error_list` VALUES (325, '117', '58', '2024-11-05 19:26:37');
INSERT INTO `x_monitor_error_list` VALUES (326, '116', '58', '2024-11-05 19:26:57');
INSERT INTO `x_monitor_error_list` VALUES (327, '118', '59', '2024-11-05 19:27:25');
INSERT INTO `x_monitor_error_list` VALUES (328, '119', '59', '2024-11-05 19:27:25');
INSERT INTO `x_monitor_error_list` VALUES (329, '120', '59', '2024-11-05 19:27:43');
INSERT INTO `x_monitor_error_list` VALUES (330, '121', '59', '2024-11-05 19:27:43');
INSERT INTO `x_monitor_error_list` VALUES (331, '122', '59', '2024-11-05 19:28:02');
INSERT INTO `x_monitor_error_list` VALUES (332, '123', '59', '2024-11-05 19:28:28');
INSERT INTO `x_monitor_error_list` VALUES (333, '124', '59', '2024-11-05 19:28:52');
INSERT INTO `x_monitor_error_list` VALUES (334, '125', '59', '2024-11-05 19:29:55');
INSERT INTO `x_monitor_error_list` VALUES (335, '126', '59', '2024-11-05 19:30:21');
INSERT INTO `x_monitor_error_list` VALUES (336, '127', '59', '2024-11-05 22:29:58');
INSERT INTO `x_monitor_error_list` VALUES (337, '128', '59', '2024-11-05 22:31:17');
INSERT INTO `x_monitor_error_list` VALUES (338, '129', '59', '2024-11-05 22:31:51');
INSERT INTO `x_monitor_error_list` VALUES (339, '130', '61', '2024-11-05 23:40:46');
INSERT INTO `x_monitor_error_list` VALUES (340, '97', '61', '2024-11-05 23:42:49');
INSERT INTO `x_monitor_error_list` VALUES (341, '131', '62', '2024-11-06 11:13:52');
INSERT INTO `x_monitor_error_list` VALUES (342, '116', '62', '2024-11-06 11:30:32');
INSERT INTO `x_monitor_error_list` VALUES (343, '116', '63', '2024-11-06 11:45:31');
INSERT INTO `x_monitor_error_list` VALUES (344, '132', '64', '2024-11-06 11:46:49');
INSERT INTO `x_monitor_error_list` VALUES (345, '133', '64', '2024-11-06 11:47:34');
INSERT INTO `x_monitor_error_list` VALUES (346, '134', '65', '2024-11-06 11:48:15');
INSERT INTO `x_monitor_error_list` VALUES (347, '116', '65', '2024-11-06 11:57:15');
INSERT INTO `x_monitor_error_list` VALUES (348, '135', '66', '2024-11-06 12:37:43');
INSERT INTO `x_monitor_error_list` VALUES (349, '135', '66', '2024-11-06 12:37:43');
INSERT INTO `x_monitor_error_list` VALUES (350, '135', '66', '2024-11-06 12:37:43');
INSERT INTO `x_monitor_error_list` VALUES (351, '136', '66', '2024-11-06 12:37:43');
INSERT INTO `x_monitor_error_list` VALUES (352, '116', '67', '2024-11-06 12:39:57');
INSERT INTO `x_monitor_error_list` VALUES (353, '116', '67', '2024-11-06 12:39:57');
INSERT INTO `x_monitor_error_list` VALUES (354, '116', '67', '2024-11-06 12:39:57');
INSERT INTO `x_monitor_error_list` VALUES (355, '116', '67', '2024-11-06 12:40:07');
INSERT INTO `x_monitor_error_list` VALUES (356, '137', '68', '2024-11-06 12:48:11');
INSERT INTO `x_monitor_error_list` VALUES (357, '136', '68', '2024-11-06 12:48:11');
INSERT INTO `x_monitor_error_list` VALUES (358, '137', '68', '2024-11-06 12:48:41');
INSERT INTO `x_monitor_error_list` VALUES (359, '137', '68', '2024-11-06 12:49:01');
INSERT INTO `x_monitor_error_list` VALUES (360, '138', '68', '2024-11-06 12:52:01');
INSERT INTO `x_monitor_error_list` VALUES (361, '138', '68', '2024-11-06 12:52:01');
INSERT INTO `x_monitor_error_list` VALUES (362, '137', '68', '2024-11-06 12:52:01');
INSERT INTO `x_monitor_error_list` VALUES (363, '137', '68', '2024-11-06 12:52:01');
INSERT INTO `x_monitor_error_list` VALUES (364, '137', '68', '2024-11-06 12:52:41');
INSERT INTO `x_monitor_error_list` VALUES (365, '116', '68', '2024-11-06 12:53:11');
INSERT INTO `x_monitor_error_list` VALUES (366, '137', '68', '2024-11-06 12:53:31');
INSERT INTO `x_monitor_error_list` VALUES (367, '136', '68', '2024-11-06 13:05:11');
INSERT INTO `x_monitor_error_list` VALUES (368, '139', '68', '2024-11-06 13:15:37');
INSERT INTO `x_monitor_error_list` VALUES (369, '139', '68', '2024-11-06 13:15:37');
INSERT INTO `x_monitor_error_list` VALUES (370, '139', '68', '2024-11-06 13:22:12');
INSERT INTO `x_monitor_error_list` VALUES (371, '139', '68', '2024-11-06 13:22:12');
INSERT INTO `x_monitor_error_list` VALUES (372, '140', '68', '2024-11-07 12:32:39');
INSERT INTO `x_monitor_error_list` VALUES (373, '136', '68', '2024-11-07 14:22:45');
INSERT INTO `x_monitor_error_list` VALUES (374, '136', '68', '2024-11-07 14:31:55');
INSERT INTO `x_monitor_error_list` VALUES (375, '141', '68', '2024-11-07 14:46:41');
INSERT INTO `x_monitor_error_list` VALUES (376, '142', '68', '2024-11-07 15:17:30');
INSERT INTO `x_monitor_error_list` VALUES (377, '136', '68', '2024-11-07 16:30:58');
INSERT INTO `x_monitor_error_list` VALUES (378, '143', '71', '2024-11-07 16:37:36');
INSERT INTO `x_monitor_error_list` VALUES (379, '144', '71', '2024-11-07 16:40:46');
INSERT INTO `x_monitor_error_list` VALUES (380, '144', '71', '2024-11-07 16:40:46');
INSERT INTO `x_monitor_error_list` VALUES (381, '145', '71', '2024-11-07 16:40:46');
INSERT INTO `x_monitor_error_list` VALUES (382, '136', '71', '2024-11-07 16:48:56');
INSERT INTO `x_monitor_error_list` VALUES (383, '146', '76', '2024-11-07 16:53:10');
INSERT INTO `x_monitor_error_list` VALUES (384, '146', '76', '2024-11-07 16:53:10');
INSERT INTO `x_monitor_error_list` VALUES (385, '146', '76', '2024-11-07 16:53:10');
INSERT INTO `x_monitor_error_list` VALUES (386, '146', '76', '2024-11-07 16:53:10');
INSERT INTO `x_monitor_error_list` VALUES (387, '146', '76', '2024-11-07 16:53:10');
INSERT INTO `x_monitor_error_list` VALUES (388, '146', '76', '2024-11-07 16:53:10');
INSERT INTO `x_monitor_error_list` VALUES (389, '147', '76', '2024-11-07 16:53:50');
INSERT INTO `x_monitor_error_list` VALUES (390, '147', '76', '2024-11-07 16:54:15');
INSERT INTO `x_monitor_error_list` VALUES (391, '148', '76', '2024-11-07 16:54:55');
INSERT INTO `x_monitor_error_list` VALUES (392, '136', '76', '2024-11-07 17:00:45');
INSERT INTO `x_monitor_error_list` VALUES (393, '136', '76', '2024-11-07 17:14:58');
INSERT INTO `x_monitor_error_list` VALUES (394, '136', '77', '2024-11-07 17:15:57');
INSERT INTO `x_monitor_error_list` VALUES (395, '136', '77', '2024-11-07 17:16:37');
INSERT INTO `x_monitor_error_list` VALUES (396, '136', '77', '2024-11-07 17:18:48');
INSERT INTO `x_monitor_error_list` VALUES (397, '149', '78', '2024-11-07 17:28:58');
INSERT INTO `x_monitor_error_list` VALUES (398, '149', '78', '2024-11-07 17:30:18');
INSERT INTO `x_monitor_error_list` VALUES (399, '150', '81', '2024-11-08 13:27:15');
INSERT INTO `x_monitor_error_list` VALUES (400, '151', '81', '2024-11-08 14:03:52');
INSERT INTO `x_monitor_error_list` VALUES (401, '152', '81', '2024-11-08 14:03:52');
INSERT INTO `x_monitor_error_list` VALUES (402, '153', '83', '2024-11-08 15:09:40');
INSERT INTO `x_monitor_error_list` VALUES (403, '154', '83', '2024-11-08 16:44:27');

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
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '监控项目' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_monitor_project
-- ----------------------------
INSERT INTO `x_monitor_project` VALUES (1, '2a01a60efed14d6eab279ba5c664ceac', 'x_admin', 'node', 0, 1, '2024-05-18 15:59:53', '2024-07-09 12:06:42', '2024-10-29 16:11:37');
INSERT INTO `x_monitor_project` VALUES (2, '9a1d5d657d884bbfb7df0ddf032a8504', '2', 'go', 0, 1, '2024-06-17 18:58:57', '2024-06-29 00:31:27', NULL);
INSERT INTO `x_monitor_project` VALUES (3, '603055b4b53e45a4b6fb725147d7fd42', 'ad', 'web', 0, 1, '2024-07-08 15:26:56', '2024-07-10 12:21:11', NULL);
INSERT INTO `x_monitor_project` VALUES (4, '95cf6a6a37e549019c90c2ed34704f1f', 'a', 'uniapp', 0, 1, '2024-07-09 00:45:18', '2024-07-09 00:49:27', NULL);
INSERT INTO `x_monitor_project` VALUES (5, '57377a40408f4bfd974047fbf0a18e24', 'go项目', 'go', 0, 1, '2024-07-10 12:20:37', '2024-09-25 16:50:45', '2024-10-29 16:11:37');
INSERT INTO `x_monitor_project` VALUES (6, '6217ea4ea0044014831bd25121a3113c', 'go', 'go', 0, 0, '2024-07-12 23:17:23', '2024-11-07 15:43:08', NULL);
INSERT INTO `x_monitor_project` VALUES (7, 'e19e3be20de94f49b68fafb4c30668bc', 'web项目', 'web', 1, 0, '2024-07-13 20:56:21', '2024-09-25 16:58:58', NULL);
INSERT INTO `x_monitor_project` VALUES (8, 'e0cb2d091ce1408babcf17b48eccebf6', '1', 'web', 1, 1, '2024-10-12 10:35:32', '2024-10-12 10:35:38', '2024-10-12 10:35:38');
INSERT INTO `x_monitor_project` VALUES (9, '72ed4cafe3c141e9ba662224651a22f0', 'a', 'web', 1, 1, '2024-10-29 16:16:08', '2024-10-29 16:16:22', '2024-10-29 16:16:22');

-- ----------------------------
-- Table structure for x_monitor_slow
-- ----------------------------
DROP TABLE IF EXISTS `x_monitor_slow`;
CREATE TABLE `x_monitor_slow`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '错误id',
  `project_key` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '项目key',
  `client_id` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'sdk生成的客户端id',
  `user_id` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户id',
  `path` varchar(1000) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'URL地址',
  `time` float(10, 2) NOT NULL COMMENT '时间',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `project_key`(`project_key`) USING BTREE,
  INDEX `client_id`(`client_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 167 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '监控-慢接口' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_monitor_slow
-- ----------------------------
INSERT INTO `x_monitor_slow` VALUES (136, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '', 'http://localhost:5174/login?redirect=/monitor/monitor/slow/index', 8515.50, '2024-11-07 16:47:42');
INSERT INTO `x_monitor_slow` VALUES (137, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '', 'http://localhost:5174/monitor/monitor/slow/index', 702.10, '2024-11-07 16:52:00');
INSERT INTO `x_monitor_slow` VALUES (138, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '1', 'http://localhost:5174/monitor/monitor/slow/index', 2648.20, '2024-11-07 16:52:02');
INSERT INTO `x_monitor_slow` VALUES (139, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '1', 'http://localhost:5174/monitor/monitor/slow/index', 393.70, '2024-11-07 16:52:13');
INSERT INTO `x_monitor_slow` VALUES (140, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '1', 'http://localhost:5174/monitor/monitor/slow/index', 635.30, '2024-11-07 16:54:10');
INSERT INTO `x_monitor_slow` VALUES (141, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '1', 'http://localhost:5174/monitor/monitor/slow/index', 620.50, '2024-11-07 16:56:25');
INSERT INTO `x_monitor_slow` VALUES (142, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '1', 'http://localhost:5174/monitor/monitor/slow/index', 1083.00, '2024-11-07 17:11:11');
INSERT INTO `x_monitor_slow` VALUES (143, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '1', 'http://localhost:5174/monitor/monitor/slow/index', 656.50, '2024-11-07 17:12:08');
INSERT INTO `x_monitor_slow` VALUES (144, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '1', 'http://localhost:5174/monitor/monitor/slow/index', 8103.20, '2024-11-07 17:15:07');
INSERT INTO `x_monitor_slow` VALUES (145, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '1', 'http://localhost:5174/monitor/monitor/slow/index', 8614.60, '2024-11-07 17:24:01');
INSERT INTO `x_monitor_slow` VALUES (146, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '1', 'http://localhost:5174/monitor/monitor/slow/index', 9773.30, '2024-11-07 17:25:51');
INSERT INTO `x_monitor_slow` VALUES (147, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '1', 'http://localhost:5174/monitor/monitor/slow/index', 8567.50, '2024-11-07 17:26:09');
INSERT INTO `x_monitor_slow` VALUES (148, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '1', 'http://localhost:5174/monitor/monitor/slow/index', 9582.50, '2024-11-07 17:26:25');
INSERT INTO `x_monitor_slow` VALUES (149, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '1', 'http://localhost:5174/monitor/monitor/slow/index', 8762.30, '2024-11-07 17:27:20');
INSERT INTO `x_monitor_slow` VALUES (150, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '1', 'http://localhost:5174/monitor/monitor/slow/index', 8923.60, '2024-11-07 17:27:39');
INSERT INTO `x_monitor_slow` VALUES (151, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '1', 'http://localhost:5174/monitor/monitor/slow/index', 8564.70, '2024-11-07 17:28:16');
INSERT INTO `x_monitor_slow` VALUES (152, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '1', 'http://localhost:5174/monitor/monitor/slow/index', 8892.90, '2024-11-07 17:28:49');
INSERT INTO `x_monitor_slow` VALUES (153, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '1', 'http://localhost:5174/monitor/monitor/slow/index', 8643.80, '2024-11-07 17:29:29');
INSERT INTO `x_monitor_slow` VALUES (154, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '1', 'http://localhost:5174/monitor/monitor/slow/index', 8916.20, '2024-11-07 17:30:08');
INSERT INTO `x_monitor_slow` VALUES (155, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '1', 'http://localhost:5174/monitor/monitor/slow/index', 8609.90, '2024-11-07 17:37:47');
INSERT INTO `x_monitor_slow` VALUES (156, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '1', 'http://localhost:5174/monitor/monitor/slow/index', 8719.30, '2024-11-07 18:13:16');
INSERT INTO `x_monitor_slow` VALUES (157, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '1', 'http://localhost:5174/monitor/monitor/slow/index', 8694.10, '2024-11-07 18:16:45');
INSERT INTO `x_monitor_slow` VALUES (158, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '1', 'http://localhost:5174/monitor/monitor/slow/index', 8556.70, '2024-11-07 18:17:57');
INSERT INTO `x_monitor_slow` VALUES (159, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '1', 'http://localhost:5174/monitor/monitor/slow/index', 8470.90, '2024-11-07 18:23:04');
INSERT INTO `x_monitor_slow` VALUES (160, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '2', 'http://localhost:5174/monitor/monitor/slow/index', 8797.10, '2024-11-07 18:27:05');
INSERT INTO `x_monitor_slow` VALUES (161, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '2', 'http://localhost:5174/monitor/monitor/slow/index', 2031.20, '2024-11-07 18:33:16');
INSERT INTO `x_monitor_slow` VALUES (162, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '2', 'http://localhost:5174/monitor/monitor/slow/index', 909.20, '2024-11-07 18:33:18');
INSERT INTO `x_monitor_slow` VALUES (163, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '2', 'http://localhost:5174/monitor/monitor/slow/index', 810.80, '2024-11-07 18:33:20');
INSERT INTO `x_monitor_slow` VALUES (164, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '2', 'http://localhost:5174/monitor/monitor/slow/index', 1180.90, '2024-11-08 13:27:05');
INSERT INTO `x_monitor_slow` VALUES (165, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '2', 'http://localhost:5174/monitor/monitor/slow/index', 927.50, '2024-11-08 13:29:29');
INSERT INTO `x_monitor_slow` VALUES (166, 'e19e3be20de94f49b68fafb4c30668bc', 'be924f50-9a9c-11ef-b456-837484cb12c1', '2', 'http://localhost:5174/setting/permission/menu', 12655.80, '2024-11-08 14:05:10');

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
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统管理成员表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_system_auth_admin
-- ----------------------------
INSERT INTO `x_system_auth_admin` VALUES (1, 1, 3, 'admin', 'admin', '81a13dd8e25644a8823082573ca973f7', '/image/20232911/27699a1e991e4ba1ba13ee6d45185db6.jpg', '0', 'WFdiD', 1, 1, 0, 0, '127.0.0.1', '2024-11-08 13:29:42', '2024-01-02 03:04:05', '2024-11-08 13:29:42', NULL);

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
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统部门管理表' ROW_FORMAT = DYNAMIC;

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
) ENGINE = InnoDB AUTO_INCREMENT = 873 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统菜单管理表' ROW_FORMAT = DYNAMIC;

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
INSERT INTO `x_system_auth_menu` VALUES (610, 600, 'C', '代码生成器', '', 0, 'admin:gen:list', 'code', 'dev_tools/code/index', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-10-31 14:59:14');
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
INSERT INTO `x_system_auth_menu` VALUES (783, 832, 'C', '项目', '', 0, 'admin:monitor_project:list', 'project/index', 'monitor/project/index', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-10-29 15:40:44');
INSERT INTO `x_system_auth_menu` VALUES (784, 832, 'C', '用户端', '', 0, '', 'client/index', 'monitor/client/index', '', '', 1, 1, 0, '2024-01-02 03:04:05', '2024-10-29 15:40:54');
INSERT INTO `x_system_auth_menu` VALUES (785, 794, 'A', '错误收集error添加', '', 0, 'admin:monitor_web:add', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (786, 794, 'A', '错误收集error编辑', '', 0, 'admin:monitor_web:edit', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (787, 794, 'A', '错误收集error删除', '', 0, 'admin:monitor_web:del', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (788, 794, 'A', '错误收集error列表', '', 0, 'admin:monitor_web:list', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (789, 794, 'A', '错误收集error全部列表', '', 0, 'admin:monitor_web:listAll', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (790, 794, 'A', '错误收集error详情', '', 0, 'admin:monitor_web:detail', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (791, 794, 'A', '错误收集error导出excel', '', 0, 'admin:monitor_web:ExportFile', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (792, 794, 'A', '错误收集error导入excel', '', 0, 'admin:monitor_web:ImportFile', '', '', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-01-02 03:04:05');
INSERT INTO `x_system_auth_menu` VALUES (794, 832, 'C', '错误列表', '', 0, '', 'error/index', 'monitor/error/index', '', '', 0, 1, 0, '2024-01-02 03:04:05', '2024-10-29 15:39:16');
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
INSERT INTO `x_system_auth_menu` VALUES (832, 0, 'M', '错误监控', 'el-icon-Memo', 0, '', 'monitor', '', '', '', 1, 1, 0, '2024-10-29 15:32:33', '2024-10-29 15:32:33');
INSERT INTO `x_system_auth_menu` VALUES (833, 600, 'C', '错误捕获', '', 0, '', 'error/test', 'error/test', '', '', 1, 1, 0, '2024-10-30 18:16:56', '2024-10-30 18:16:56');
INSERT INTO `x_system_auth_menu` VALUES (834, 832, 'C', '慢接口', '', 0, '', 'slow/index', 'monitor/slow/index', '', '', 0, 1, 0, '2024-11-06 12:30:19', '2024-11-08 13:45:22');
INSERT INTO `x_system_auth_menu` VALUES (844, 834, 'A', '慢接口添加', '', 0, 'admin:monitor_slow:add', '', '', '', '', 0, 1, 0, '2024-11-06 12:34:55', '2024-11-06 12:34:55');
INSERT INTO `x_system_auth_menu` VALUES (845, 834, 'A', '慢接口编辑', '', 0, 'admin:monitor_slow:edit', '', '', '', '', 0, 1, 0, '2024-11-06 12:34:55', '2024-11-06 12:34:55');
INSERT INTO `x_system_auth_menu` VALUES (846, 834, 'A', '慢接口删除', '', 0, 'admin:monitor_slow:del', '', '', '', '', 0, 1, 0, '2024-11-06 12:34:55', '2024-11-06 12:34:55');
INSERT INTO `x_system_auth_menu` VALUES (847, 834, 'A', '慢接口删除-批量', '', 0, 'admin:monitor_slow:delBatch', '', '', '', '', 0, 1, 0, '2024-11-06 12:34:55', '2024-11-06 12:34:55');
INSERT INTO `x_system_auth_menu` VALUES (848, 834, 'A', '慢接口列表', '', 0, 'admin:monitor_slow:list', '', '', '', '', 0, 1, 0, '2024-11-06 12:34:55', '2024-11-06 12:34:55');
INSERT INTO `x_system_auth_menu` VALUES (849, 834, 'A', '慢接口全部列表', '', 0, 'admin:monitor_slow:listAll', '', '', '', '', 0, 1, 0, '2024-11-06 12:34:55', '2024-11-06 12:34:55');
INSERT INTO `x_system_auth_menu` VALUES (850, 834, 'A', '慢接口详情', '', 0, 'admin:monitor_slow:detail', '', '', '', '', 0, 1, 0, '2024-11-06 12:34:55', '2024-11-06 12:34:55');
INSERT INTO `x_system_auth_menu` VALUES (851, 834, 'A', '慢接口导出excel', '', 0, 'admin:monitor_slow:ExportFile', '', '', '', '', 0, 1, 0, '2024-11-06 12:34:55', '2024-11-06 12:34:55');
INSERT INTO `x_system_auth_menu` VALUES (852, 834, 'A', '慢接口导入excel', '', 0, 'admin:monitor_slow:ImportFile', '', '', '', '', 0, 1, 0, '2024-11-06 12:34:55', '2024-11-06 12:34:55');
INSERT INTO `x_system_auth_menu` VALUES (864, 784, 'A', '监控-客户端信息添加', '', 0, 'admin:monitor_client:add', '', '', '', '', 0, 1, 0, '2024-11-08 14:15:28', '2024-11-08 14:15:28');
INSERT INTO `x_system_auth_menu` VALUES (865, 784, 'A', '监控-客户端信息编辑', '', 0, 'admin:monitor_client:edit', '', '', '', '', 0, 1, 0, '2024-11-08 14:15:28', '2024-11-08 14:15:28');
INSERT INTO `x_system_auth_menu` VALUES (866, 784, 'A', '监控-客户端信息删除', '', 0, 'admin:monitor_client:del', '', '', '', '', 0, 1, 0, '2024-11-08 14:15:28', '2024-11-08 14:15:28');
INSERT INTO `x_system_auth_menu` VALUES (867, 784, 'A', '监控-客户端信息删除-批量', '', 0, 'admin:monitor_client:delBatch', '', '', '', '', 0, 1, 0, '2024-11-08 14:15:28', '2024-11-08 14:15:28');
INSERT INTO `x_system_auth_menu` VALUES (868, 784, 'A', '监控-客户端信息列表', '', 0, 'admin:monitor_client:list', '', '', '', '', 0, 1, 0, '2024-11-08 14:15:28', '2024-11-08 14:15:28');
INSERT INTO `x_system_auth_menu` VALUES (869, 784, 'A', '监控-客户端信息全部列表', '', 0, 'admin:monitor_client:listAll', '', '', '', '', 0, 1, 0, '2024-11-08 14:15:28', '2024-11-08 14:15:28');
INSERT INTO `x_system_auth_menu` VALUES (870, 784, 'A', '监控-客户端信息详情', '', 0, 'admin:monitor_client:detail', '', '', '', '', 0, 1, 0, '2024-11-08 14:15:28', '2024-11-08 14:15:28');
INSERT INTO `x_system_auth_menu` VALUES (871, 784, 'A', '监控-客户端信息导出excel', '', 0, 'admin:monitor_client:ExportFile', '', '', '', '', 0, 1, 0, '2024-11-08 14:15:28', '2024-11-08 14:15:28');
INSERT INTO `x_system_auth_menu` VALUES (872, 784, 'A', '监控-客户端信息导入excel', '', 0, 'admin:monitor_client:ImportFile', '', '', '', '', 0, 1, 0, '2024-11-08 14:15:28', '2024-11-08 14:15:28');

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
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统岗位管理表' ROW_FORMAT = DYNAMIC;

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
) ENGINE = InnoDB AUTO_INCREMENT = 20 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统登录日志表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_system_log_login
-- ----------------------------
INSERT INTO `x_system_log_login` VALUES (1, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-10-29 14:27:41');
INSERT INTO `x_system_log_login` VALUES (2, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-10-29 14:46:35');
INSERT INTO `x_system_log_login` VALUES (3, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-10-29 23:05:34');
INSERT INTO `x_system_log_login` VALUES (4, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-10-30 09:49:22');
INSERT INTO `x_system_log_login` VALUES (5, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-10-30 10:05:00');
INSERT INTO `x_system_log_login` VALUES (6, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-10-30 18:15:47');
INSERT INTO `x_system_log_login` VALUES (7, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-10-30 19:02:50');
INSERT INTO `x_system_log_login` VALUES (8, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-10-31 10:57:33');
INSERT INTO `x_system_log_login` VALUES (9, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-10-31 14:56:49');
INSERT INTO `x_system_log_login` VALUES (10, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-11-04 14:59:54');
INSERT INTO `x_system_log_login` VALUES (11, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-11-04 19:02:36');
INSERT INTO `x_system_log_login` VALUES (12, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-11-05 14:05:30');
INSERT INTO `x_system_log_login` VALUES (13, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-11-05 22:22:23');
INSERT INTO `x_system_log_login` VALUES (14, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-11-06 11:13:57');
INSERT INTO `x_system_log_login` VALUES (15, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-11-07 12:32:52');
INSERT INTO `x_system_log_login` VALUES (16, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-11-07 13:01:06');
INSERT INTO `x_system_log_login` VALUES (17, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-11-07 15:17:26');
INSERT INTO `x_system_log_login` VALUES (18, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-11-07 16:47:51');
INSERT INTO `x_system_log_login` VALUES (19, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, '2024-11-08 13:29:42');

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
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统操作日志表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_system_log_operate
-- ----------------------------
INSERT INTO `x_system_log_operate` VALUES (1, 0, 'GET', '监控-客户端信息新增', '127.0.0.1', '/api/admin/monitor_client/add', 'x_admin/admin/monitor_client.(*MonitorClientHandler).Add-fm', 'data=[{\"type\":\"error\",\"eventType\":\"error\",\"message\":\"ResizeObserver loop completed with undelivered notifications.\",\"stack\":\"\"}]', 'runtime error: invalid memory address or nil pointer dereference', 2, '2024-10-30 23:30:23', '2024-10-30 23:30:29', 6622, '2024-10-30 23:30:30');

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
