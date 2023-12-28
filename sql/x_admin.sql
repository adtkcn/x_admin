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

 Date: 29/12/2023 00:52:04
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
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  `delete_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_cid`(`cid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '相册管理表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_album
-- ----------------------------
INSERT INTO `x_album` VALUES (1, 0, 1, 0, 10, 'Alx_gp73pq.png', 'image/20230911/34a557325c004f498b1da01bb068f919.png', 'png', 7138309, 1, 1699499781, 1700821723, 1700821723);
INSERT INTO `x_album` VALUES (2, 0, 1, 0, 20, '素材中心 和另外 1 个页面 - 个人 - Microsoft​ Edge 2023-11-20 15-37-39.mp4', 'video/20232011/e5f53b824e314ab7992ef4b6e7595b86.mp4', 'mp4', 4053956, 0, 1700465870, 1700465870, 0);
INSERT INTO `x_album` VALUES (3, 1, 1, 0, 10, 'COMPUMAX_1p2x61.jpg', 'image/20232411/cd61498e40f545f9863e6251af587e67.jpg', 'jpg', 613802, 1, 1700821741, 1700821809, 1700821809);
INSERT INTO `x_album` VALUES (4, 1, 1, 0, 10, '截图 2023-11-24 12.05.30-fullpage.png', 'image/20232511/2c719c9ae2e14ed39d9710426df63c3f.png', 'png', 155658, 1, 1700924913, 1701224132, 1701224132);
INSERT INTO `x_album` VALUES (5, 1, 1, 0, 10, 'COMPUMAX_1p2x61.jpg', 'image/20232511/8510ef876e5b458d8e5c4291d9ac12e1.jpg', 'jpg', 613802, 1, 1700924945, 1701224129, 1701224129);
INSERT INTO `x_album` VALUES (6, 0, 1, 0, 10, 'COMPUMAX_1p2x61.jpg', 'image/20232911/27699a1e991e4ba1ba13ee6d45185db6.jpg', 'jpg', 613802, 0, 1701224137, 1701224137, 0);

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
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  `delete_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '相册分类表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_album_cate
-- ----------------------------
INSERT INTO `x_album_cate` VALUES (1, 0, 10, '1', 0, 1699519913, 1699519913, 0);

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
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  `delete_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `cid_idx`(`cid`) USING BTREE COMMENT '分类索引'
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文章资讯表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_article
-- ----------------------------
INSERT INTO `x_article` VALUES (1, 1, '让生活更精致！五款居家好物推荐，实用性超高', '##好物推荐🔥', '随着当代生活节奏的忙碌，很多人在闲暇之余都想好好的享受生活。随着科技的发展，也出现了越来越多可以帮助我们提升幸福感，让生活变得更精致的产品，下面周周就给大家盘点五款居家必备的好物，都是实用性很高的产品，周周可以保证大家买了肯定会喜欢。', '/api/static/article01.png', '<p><img src=\"https://likeadmin-java.yixiangonline.com/api/uploads/image/20220916/46d29489-4260-4917-8eca-d0f6cba6af23.png\" alt=\"\" data-href=\"\" style=\"\"/></p><p>拥有一台投影仪，闲暇时可以在家里直接看影院级别的大片，光是想想都觉得超级爽。市面上很多投影仪大几千，其实周周觉得没必要，选泰捷这款一千多的足够了，性价比非常高。</p><p>泰捷的专业度很高，在电视TV领域研发已经十年，有诸多专利和技术创新，荣获国内外多项技术奖项，拿下了腾讯创新工场投资，打造的泰捷视频TV端和泰捷电视盒子都获得了极高评价。</p><p>这款投影仪的分辨率在3000元内无敌，做到了真1080P高分辨率，也就是跟市场售价三千DLP投影仪一样的分辨率，真正做到了分毫毕现，像桌布的花纹、天空的云彩等，这些细节都清晰可见。</p><p>亮度方面，泰捷达到了850ANSI流明，同价位一般是200ANSI。这是因为泰捷为了提升亮度和LCD技术透射率低的问题，首创高功率LED灯源，让其亮度做到同价位最好。专业媒体也进行了多次对比，效果与3000元价位投影仪相当。</p><p>操作系统周周也很喜欢，完全不卡。泰捷作为资深音视频品牌，在系统优化方面有十年的研发经验，打造出的“零极”系统是业内公认效率最高、速度最快的系统，用户也评价它流畅度能一台顶三台，而且为了解决行业广告多这一痛点，系统内不植入任何广告。</p>', '红花', 9, 0, 1, 0, 1663317759, 1663322726, 0);
INSERT INTO `x_article` VALUES (2, 1, '埋葬UI设计师的坟墓不是内卷，而是免费模式', '', '本文从另外一个角度，聊聊作者对UI设计师职业发展前景的担忧，欢迎从事UI设计的同学来参与讨论，会有赠书哦', '/api/static/article02.jpeg', '<p><br></p><p style=\"text-align: justify;\">一个职业，卷，根本就没什么大不了的，尤其是成熟且收入高的职业，不卷才不符合事物发展的规律。何况 UI 设计师的人力市场到今天也和 5 年前一样，还是停留在大型菜鸡互啄的场面。远不能和医疗、证券、教师或者演艺练习生相提并论。</p><p style=\"text-align: justify;\">真正会让我对 <a href=\"https://www.uisdc.com/tag/ui\" target=\"_blank\">UI</a> 设计师发展前景觉得悲观的事情就只有一件 —— 国内的互联网产品免费机制。这也是一个我一直以来想讨论的话题，就在这次写一写。</p><p style=\"text-align: justify;\">国内互联网市场的发展，是一部浩瀚的 “免费经济” 发展史。虽然今天免费已经是深入国内民众骨髓的认知，但最早的中文互联网也是需要付费的，网游也都是要花钱的。</p><p style=\"text-align: justify;\">只是自有国情在此，付费确实阻碍了互联网行业的扩张和普及，一批创业家就开始通过免费的模式为用户提供服务，从而扩大了自己的产品覆盖面和普及程度。</p><p style=\"text-align: justify;\">印象最深的就是免费急先锋周鸿祎，和现在鲜少出现在公众视野不同，一零年前他是当之无愧的互联网教主，因为他开发出了符合中国国情的互联网产品 “打法”，让 360 的发展如日中天。</p><p style=\"text-align: justify;\">就是他在自传中提到：</p><p style=\"text-align: justify;\">只要是在互联网上每个人都需要的服务，我们就认为它是基础服务，基础服务一定是免费的，这样的话不会形成价值歧视。就是说，只要这种服务是每个人都一定要用的，我一定免费提供，而且是无条件免费。增值服务不是所有人都需要的，这个比例可能会相当低，它只是百分之几甚至更少比例的人需要，所以这种服务一定要收费……</p><p style=\"text-align: justify;\">这就是互联网的游戏规则，它决定了要想建立一个有效的商业模式，就一定要有海量的用户基数……</p>', '一一', 23, 0, 1, 0, 1663320938, 1663322854, 0);
INSERT INTO `x_article` VALUES (3, 2, '金山电池公布“沪广深市民绿色生活方式”调查结果', '', '60%以上受访者认为高质量的10分钟足以完成“自我充电”', '/api/static/article03.png', '<p style=\"text-align: left;\"><strong>深圳，2021年10月22日）</strong>生活在一线城市的沪广深市民一向以效率见称，工作繁忙和快节奏的生活容易缺乏充足的休息。近日，一项针对沪广深市民绿色生活方式而展开的网络问卷调查引起了大家的注意。问卷的问题设定集中于市民对休息时间的看法，以及从对循环充电电池的使用方面了解其对绿色生活方式的态度。该调查采用随机抽样的模式，并对最终收集的1,500份有效问卷进行专业分析后发现，超过60%的受访者表示，在每天的工作时段能拥有10分钟高质量的休息时间，就可以高效“自我充电”。该调查结果反映出，在快节奏时代下，人们需要高质量的休息时间，也要学会利用高效率的休息方式和工具来应对快节奏的生活，以时刻保持“满电”状态。</p><p style=\"text-align: left;\">　　<strong>60%以上受访者认为高质量的10分钟足以完成“自我充电”</strong></p><p style=\"text-align: left;\">　　这次调查超过1,500人，主要聚焦18至85岁的沪广深市民，了解他们对于休息时间的观念及使用充电电池的习惯，结果发现：</p><p style=\"text-align: left;\">　　· 90%以上有工作受访者每天工作时间在7小时以上，平均工作时间为8小时，其中43%以上的受访者工作时间超过9小时</p><p style=\"text-align: left;\">　　· 70%受访者认为在工作期间拥有10分钟“自我充电”时间不是一件困难的事情</p><p style=\"text-align: left;\">　　· 60%受访者认为在工作期间有10分钟休息时间足以为自己快速充电</p><p style=\"text-align: left;\">　　临床心理学家黄咏诗女士在发布会上分享为自己快速充电的实用技巧，她表示：“事实上，只要选择正确的休息方法，10分钟也足以为自己充电。以喝咖啡为例，我们可以使用心灵休息法 ── 静观呼吸，慢慢感受咖啡的温度和气味，如果能配合着聆听流水或海洋的声音，能够有效放松大脑及心灵。”</p><p style=\"text-align: left;\">　　这次调查结果反映出沪广深市民的希望在繁忙的工作中适时停下来，抽出10分钟喝杯咖啡、聆听音乐或小睡片刻，为自己充电。金山电池全新推出的“绿再十分充”超快速充电器仅需10分钟就能充好电，喝一杯咖啡的时间既能完成“自我充电”，也满足设备使用的用电需求，为提升工作效率和放松身心注入新能量。</p><p style=\"text-align: left;\">　　<strong>金山电池推出10分钟超快电池充电器*绿再十分充，以创新科技为市场带来革新体验</strong></p><p style=\"text-align: left;\">　　该问卷同时从沪广深市民对循环充电电池的使用方面进行了调查，以了解其对绿色生活方式的态度：</p><p style=\"text-align: left;\">　　· 87%受访者目前没有使用充电电池，其中61%表示会考虑使用充电电池</p><p style=\"text-align: left;\">　　· 58%受访者过往曾使用过充电电池，却只有20%左右市民仍在使用</p><p style=\"text-align: left;\">　　· 60%左右受访者认为充电电池尚未被广泛使用，主要障碍来自于充电时间过长、缺乏相关教育</p><p style=\"text-align: left;\">　　· 90%以上受访者认为充电电池充满电需要1小时或更长的时间</p><p style=\"text-align: left;\">　　金山电池一直致力于为大众提供安全可靠的充电电池，并与消费者的需求和生活方式一起演变及进步。今天，金山电池宣布推出10分钟超快电池充电器*绿再十分充，只需10分钟*即可将4粒绿再十分充充电电池充好电，充电速度比其他品牌提升3倍**。充电器的LED灯可以显示每粒电池的充电状态和模式，并提示用户是否错误插入已损坏电池或一次性电池。尽管其体型小巧，却具备多项创新科技 ，如拥有独特的充电算法以优化充电电流，并能根据各个电池类型、状况和温度用最短的时间为充电电池充好电;绿再十分充内置横流扇，有效防止电池温度过热和提供低噪音的充电环境等。<br></p>', '中网资讯科技', 3, 0, 1, 0, 1663322665, 1663322665, 0);

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
  `create_time` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '更新时间',
  `delete_time` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文章分类表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_article_category
-- ----------------------------
INSERT INTO `x_article_category` VALUES (1, '文章资讯', 0, 1, 0, 1663317280, 1663317282, 0);
INSERT INTO `x_article_category` VALUES (2, '社会热点', 0, 1, 0, 1663321464, 1663321494, 0);

-- ----------------------------
-- Table structure for x_article_collect
-- ----------------------------
DROP TABLE IF EXISTS `x_article_collect`;
CREATE TABLE `x_article_collect`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
  `article_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '文章ID',
  `is_delete` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除',
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  `delete_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文章收藏表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_article_collect
-- ----------------------------
INSERT INTO `x_article_collect` VALUES (1, 0, 0, 1, 0, 0, 0);
INSERT INTO `x_article_collect` VALUES (5, 1, 1, 1, 0, 0, 0);

-- ----------------------------
-- Table structure for x_decorate_page
-- ----------------------------
DROP TABLE IF EXISTS `x_decorate_page`;
CREATE TABLE `x_decorate_page`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `page_type` tinyint(2) UNSIGNED NOT NULL DEFAULT 10 COMMENT '页面类型',
  `page_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '页面名称',
  `page_data` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '页面数据',
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` int(10) UNSIGNED NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '页面装修表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_decorate_page
-- ----------------------------
INSERT INTO `x_decorate_page` VALUES (1, 1, '商城首页', '[{\"title\":\"搜索\",\"name\":\"search\",\"disabled\":1,\"content\":{},\"styles\":{}},{\"title\":\"首页轮播图\",\"name\":\"banner\",\"content\":{\"enabled\":1,\"data\":[{\"image\":\"/api/static/banner01.png\",\"name\":\"\",\"link\":{\"path\":\"/pages/index/index\",\"name\":\"商城首页\",\"type\":\"shop\"}},{\"image\":\"/api/static/banner02.png\",\"name\":\"\",\"link\":{}}]},\"styles\":{}},{\"title\":\"导航菜单\",\"name\":\"nav\",\"content\":{\"enabled\":1,\"data\":[{\"image\":\"/api/static/nav01.png\",\"name\":\"资讯中心\",\"link\":{\"path\":\"/pages/news/news\",\"name\":\"文章资讯\",\"type\":\"shop\"}},{\"image\":\"/api/static/nav02.png\",\"name\":\"我的收藏\",\"link\":{\"path\":\"/pages/collection/collection\",\"name\":\"我的收藏\",\"type\":\"shop\"}},{\"image\":\"/api/static/nav03.png\",\"name\":\"个人设置\",\"link\":{\"path\":\"/pages/user_set/user_set\",\"name\":\"个人设置\",\"type\":\"shop\"}},{\"image\":\"/api/static/nav04.png\",\"name\":\"联系客服\",\"link\":{\"path\":\"/pages/customer_service/customer_service\",\"name\":\"联系客服\",\"type\":\"shop\"}},{\"image\":\"/api/static/nav05.png\",\"name\":\"关于我们\",\"link\":{\"path\":\"/pages/as_us/as_us\",\"name\":\"关于我们\",\"type\":\"shop\"}}]},\"styles\":{}},{\"id\":\"l84almsk2uhyf\",\"title\":\"资讯\",\"name\":\"news\",\"disabled\":1,\"content\":{},\"styles\":{}}]', 1661757188, 1663321380);
INSERT INTO `x_decorate_page` VALUES (2, 2, '个人中心', '[{\"title\":\"用户信息\",\"name\":\"user-info\",\"disabled\":1,\"content\":{},\"styles\":{}},{\"title\":\"我的服务\",\"name\":\"my-service\",\"content\":{\"style\":2,\"title\":\"服务中心\",\"data\":[{\"image\":\"/api/static/user_collect.png\",\"name\":\"我的收藏\",\"link\":{\"path\":\"/pages/collection/collection\",\"name\":\"我的收藏\",\"type\":\"shop\"}},{\"image\":\"/api/static/user_setting.png\",\"name\":\"个人设置\",\"link\":{\"path\":\"/pages/user_set/user_set\",\"name\":\"个人设置\",\"type\":\"shop\"}},{\"image\":\"/api/static/user_kefu.png\",\"name\":\"联系客服\",\"link\":{\"path\":\"/pages/customer_service/customer_service\",\"name\":\"联系客服\",\"type\":\"shop\"}}]},\"styles\":{}},{\"title\":\"个人中心广告图\",\"name\":\"user-banner\",\"content\":{\"enabled\":1,\"data\":[{\"image\":\"/api/static/ad01.jpg\",\"name\":\"\",\"link\":{}}]},\"styles\":{}}]', 1661757188, 1663320728);
INSERT INTO `x_decorate_page` VALUES (3, 3, '客服设置', '[{\"title\":\"客服设置\",\"name\":\"customer-service\",\"content\":{\"title\":\"添加客服二维码\",\"time\":\"早上 9:00 - 22:00\",\"mobile\":\"13800138000\",\"qrcode\":\"\"},\"styles\":{}}]', 1661757188, 1662689155);

-- ----------------------------
-- Table structure for x_decorate_tabbar
-- ----------------------------
DROP TABLE IF EXISTS `x_decorate_tabbar`;
CREATE TABLE `x_decorate_tabbar`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '导航名称',
  `selected` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '未选图标',
  `unselected` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '已选图标',
  `link` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '链接地址',
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '底部装修表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_decorate_tabbar
-- ----------------------------
INSERT INTO `x_decorate_tabbar` VALUES (13, '首页', '/api/static/tabbar_home_sel.png', '/api/static/tabbar_home.png', '{\"path\":\"/pages/index/index\",\"name\":\"商城首页\",\"type\":\"shop\"}', 1662688157, 1662688157);
INSERT INTO `x_decorate_tabbar` VALUES (14, '资讯', '/api/static/tabbar_text_sel.png', '/api/static/tabbar_text.png', '{\"path\":\"/pages/news/news\",\"name\":\"文章资讯\",\"type\":\"shop\"}', 1662688157, 1662688157);
INSERT INTO `x_decorate_tabbar` VALUES (15, '我的', '/api/static/tabbar_me_sel.png', '/api/static/tabbar_me.png', '{\"path\":\"/pages/user/user\",\"name\":\"个人中心\",\"type\":\"shop\"}', 1662688157, 1662688157);

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
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  `delete_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 17 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '字典数据表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_dict_data
-- ----------------------------
INSERT INTO `x_dict_data` VALUES (1, 2, '待提交', '1', '#6D85FC', '', 0, 1, 0, 0, 1703780989, 0);
INSERT INTO `x_dict_data` VALUES (2, 2, '审批中', '2', NULL, '', 0, 1, 0, 0, 1702902752, 0);
INSERT INTO `x_dict_data` VALUES (3, 2, '审批成功', '3', 'green', '', 0, 1, 0, 0, 1703171982, 0);
INSERT INTO `x_dict_data` VALUES (4, 2, '失败', '4', 'red', '', 0, 1, 0, 0, 1703171974, 0);
INSERT INTO `x_dict_data` VALUES (5, 3, '待处理', '1', '#087BF6', '', 0, 1, 0, 0, 1703227511, 0);
INSERT INTO `x_dict_data` VALUES (6, 3, '通过', '2', 'green', '', 0, 1, 0, 0, 1703171473, 0);
INSERT INTO `x_dict_data` VALUES (7, 3, '拒绝', '3', 'red', '', 0, 1, 0, 0, 1703171461, 0);
INSERT INTO `x_dict_data` VALUES (8, 2, '666', '6', '123', '', 0, 1, 1, 1703172304, 1703172368, 1703172368);
INSERT INTO `x_dict_data` VALUES (9, 3, '0', '0', '', '', 0, 1, 1, 1703227539, 1703227544, 1703227544);
INSERT INTO `x_dict_data` VALUES (10, 4, '假勤管理', '1', '', '', 0, 1, 0, 1703767772, 1703767772, 0);
INSERT INTO `x_dict_data` VALUES (11, 4, '人事管理', '2', '', '', 0, 1, 0, 1703767780, 1703767780, 0);
INSERT INTO `x_dict_data` VALUES (12, 4, '财务管理', '3', '', '', 0, 1, 0, 1703767789, 1703767789, 0);
INSERT INTO `x_dict_data` VALUES (13, 4, '业务管理', '4', '', '', 0, 1, 0, 1703767797, 1703767797, 0);
INSERT INTO `x_dict_data` VALUES (14, 4, '行政管理', '5', '', '', 0, 1, 0, 1703767807, 1703767807, 0);
INSERT INTO `x_dict_data` VALUES (15, 4, '法务管理', '6', '', '', 0, 1, 0, 1703767821, 1703767821, 0);
INSERT INTO `x_dict_data` VALUES (16, 4, '其他', '7', '', '', 0, 1, 0, 1703767829, 1703767829, 0);

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
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  `delete_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '字典类型表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_dict_type
-- ----------------------------
INSERT INTO `x_dict_type` VALUES (2, '审批申请状态', 'flow_apply_status', '0待提交，1审批中，2审批完成，3审批失败', 1, 0, 1702207019, 1702207019, 0);
INSERT INTO `x_dict_type` VALUES (3, '审批历史状态', 'flow_history_status', '', 1, 0, 1702958280, 1702958280, 0);
INSERT INTO `x_dict_type` VALUES (4, '流程分类', 'flow_group', '1假勤管理,2人事管理3财务管理4业务管理5行政管理6法务管理7其他', 1, 0, 1703767753, 1703767753, 0);

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
  `update_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `delete_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 28 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '申请流程' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of x_flow_apply
-- ----------------------------
INSERT INTO `x_flow_apply` VALUES (25, 2, 1, 'admin', '1', 2, '2', '{\"widgetList\":[{\"key\":93225,\"type\":\"textarea\",\"icon\":\"textarea-field\",\"formItemFlag\":true,\"options\":{\"name\":\"textarea53158\",\"label\":\"textarea\",\"labelAlign\":\"\",\"rows\":3,\"defaultValue\":\"\",\"placeholder\":\"\",\"columnWidth\":\"200px\",\"size\":\"\",\"labelWidth\":null,\"labelHidden\":false,\"readonly\":false,\"disabled\":false,\"hidden\":false,\"required\":false,\"requiredHint\":\"\",\"validation\":\"\",\"validationHint\":\"\",\"customClass\":[],\"labelIconClass\":null,\"labelIconPosition\":\"rear\",\"labelTooltip\":null,\"minLength\":null,\"maxLength\":null,\"showWordLimit\":false,\"onCreated\":\"\",\"onMounted\":\"\",\"onInput\":\"\",\"onChange\":\"\",\"onFocus\":\"\",\"onBlur\":\"\",\"onValidate\":\"\"},\"id\":\"textarea53158\"},{\"key\":35486,\"type\":\"select\",\"icon\":\"select-field\",\"formItemFlag\":true,\"options\":{\"name\":\"select101798\",\"label\":\"select\",\"labelAlign\":\"\",\"defaultValue\":\"\",\"placeholder\":\"\",\"columnWidth\":\"200px\",\"size\":\"\",\"labelWidth\":null,\"labelHidden\":false,\"disabled\":false,\"hidden\":false,\"clearable\":true,\"filterable\":false,\"allowCreate\":false,\"remote\":false,\"automaticDropdown\":false,\"multiple\":false,\"multipleLimit\":0,\"optionItems\":[{\"label\":\"select 1\",\"value\":1},{\"label\":\"select 2\",\"value\":2},{\"label\":\"select 3\",\"value\":3}],\"required\":false,\"requiredHint\":\"\",\"validation\":\"\",\"validationHint\":\"\",\"customClass\":\"\",\"labelIconClass\":null,\"labelIconPosition\":\"rear\",\"labelTooltip\":null,\"onCreated\":\"\",\"onMounted\":\"\",\"onRemoteQuery\":\"\",\"onChange\":\"\",\"onFocus\":\"\",\"onBlur\":\"\",\"onValidate\":\"\"},\"id\":\"select101798\"},{\"key\":85807,\"type\":\"input\",\"icon\":\"text-field\",\"formItemFlag\":true,\"options\":{\"name\":\"input47807\",\"label\":\"input\",\"labelAlign\":\"\",\"type\":\"text\",\"defaultValue\":\"\",\"placeholder\":\"\",\"columnWidth\":\"200px\",\"size\":\"\",\"labelWidth\":null,\"labelHidden\":false,\"readonly\":false,\"disabled\":false,\"hidden\":false,\"clearable\":true,\"showPassword\":false,\"required\":false,\"requiredHint\":\"\",\"validation\":\"\",\"validationHint\":\"\",\"customClass\":\"\",\"labelIconClass\":null,\"labelIconPosition\":\"rear\",\"labelTooltip\":null,\"minLength\":null,\"maxLength\":null,\"showWordLimit\":false,\"prefixIcon\":\"\",\"suffixIcon\":\"\",\"appendButton\":false,\"appendButtonDisabled\":false,\"buttonIcon\":\"custom-search\",\"onCreated\":\"\",\"onMounted\":\"\",\"onInput\":\"\",\"onChange\":\"\",\"onFocus\":\"\",\"onBlur\":\"\",\"onValidate\":\"\",\"onAppendButtonClick\":\"\"},\"id\":\"input47807\"}],\"formConfig\":{\"modelName\":\"formData\",\"refName\":\"vForm\",\"rulesName\":\"rules\",\"labelWidth\":80,\"labelPosition\":\"left\",\"size\":\"\",\"labelAlign\":\"label-left-align\",\"cssCode\":\"\",\"customClass\":[],\"functions\":\"\",\"layoutType\":\"H5\",\"jsonVersion\":3,\"onFormCreated\":\"\",\"onFormMounted\":\"\",\"onFormDataChange\":\"\"}}', '{\"nodes\":[{\"id\":\"Event_5026d13\",\"type\":\"bpmn:startEvent\",\"x\":260,\"y\":250,\"properties\":{\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":\"\",\"userId\":\"\",\"deptId\":\"\",\"postId\":\"\"},\"zIndex\":1005,\"text\":{\"x\":260,\"y\":290,\"value\":\"开始\"}},{\"id\":\"Activity_88d7be8\",\"type\":\"bpmn:userTask\",\"x\":810,\"y\":180,\"properties\":{\"fieldAuth\":{\"textarea53158\":2,\"textarea9642\":2,\"select101798\":2,\"input47807\":2},\"userId\":1,\"deptId\":3,\"postId\":2,\"userType\":3,\"gateway\":[]},\"zIndex\":1011,\"text\":{\"x\":810,\"y\":180,\"value\":\"审批1\"}},{\"id\":\"Event_af85d64\",\"type\":\"bpmn:endEvent\",\"x\":1190,\"y\":360,\"properties\":{\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":\"\",\"userId\":\"\",\"deptId\":\"\",\"postId\":\"\"},\"zIndex\":1015,\"text\":{\"x\":1190,\"y\":400,\"value\":\"结束\"}},{\"id\":\"Gateway_5c0df6e\",\"type\":\"bpmn:exclusiveGateway\",\"x\":410,\"y\":180,\"properties\":{\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userId\":\"\",\"deptId\":\"\",\"postId\":\"\",\"userType\":\"\",\"gateway\":[{\"id\":\"textarea53158\",\"value\":\"1\",\"condition\":\"==\"}]},\"zIndex\":1005,\"text\":{\"x\":410,\"y\":220,\"value\":\"网关\"}},{\"id\":\"Activity_0c4ec56\",\"type\":\"bpmn:serviceTask\",\"x\":580,\"y\":180,\"properties\":{\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userId\":\"\",\"deptId\":\"\",\"postId\":\"\",\"userType\":\"\",\"gateway\":[]},\"zIndex\":1003,\"text\":{\"x\":580,\"y\":180,\"value\":\"系统任务\"}},{\"id\":\"Gateway_9e9bedf\",\"type\":\"bpmn:exclusiveGateway\",\"x\":410,\"y\":360,\"properties\":{\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":\"\",\"userId\":\"\",\"deptId\":\"\",\"postId\":\"\",\"gateway\":[{\"id\":\"textarea53158\",\"value\":\"2\",\"condition\":\"==\"}]},\"zIndex\":1001,\"text\":{\"x\":410,\"y\":400,\"value\":\"网关\"}},{\"id\":\"Activity_0c8eb76\",\"type\":\"bpmn:userTask\",\"x\":590,\"y\":360,\"properties\":{\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":1,\"userId\":\"\",\"deptId\":2,\"postId\":3,\"gateway\":[]},\"zIndex\":1012,\"text\":{\"x\":590,\"y\":360,\"value\":\"审批\"}},{\"id\":\"Activity_f09936b\",\"type\":\"bpmn:userTask\",\"x\":1030,\"y\":180,\"properties\":{\"userType\":3,\"userId\":1,\"deptId\":\"\",\"postId\":\"\",\"fieldAuth\":{},\"gateway\":[]},\"zIndex\":1001,\"text\":{\"x\":1030,\"y\":180,\"value\":\"审批\"}}],\"edges\":[{\"id\":\"c7d055f8-bbce-40e5-a2b4-5b0fb1ee9aea\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Event_5026d13\",\"targetNodeId\":\"Gateway_5c0df6e\",\"startPoint\":{\"x\":278,\"y\":250},\"endPoint\":{\"x\":385,\"y\":180},\"properties\":{},\"zIndex\":1009,\"pointsList\":[{\"x\":278,\"y\":250},{\"x\":355,\"y\":250},{\"x\":355,\"y\":180},{\"x\":385,\"y\":180}]},{\"id\":\"ad29adbd-060c-4e0e-9c54-a5512f2542b0\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Gateway_5c0df6e\",\"targetNodeId\":\"Activity_0c4ec56\",\"startPoint\":{\"x\":435,\"y\":180},\"endPoint\":{\"x\":530,\"y\":180},\"properties\":{},\"zIndex\":1016,\"pointsList\":[{\"x\":435,\"y\":180},{\"x\":530,\"y\":180}]},{\"id\":\"2146dbe0-9e6d-4c20-ab84-6f66d3f5f4ea\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Event_5026d13\",\"targetNodeId\":\"Gateway_9e9bedf\",\"startPoint\":{\"x\":278,\"y\":250},\"endPoint\":{\"x\":385,\"y\":360},\"properties\":{},\"zIndex\":1008,\"pointsList\":[{\"x\":278,\"y\":250},{\"x\":355,\"y\":250},{\"x\":355,\"y\":360},{\"x\":385,\"y\":360}]},{\"id\":\"7bfc16b1-a40b-4a53-8599-498a73dede74\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Gateway_9e9bedf\",\"targetNodeId\":\"Activity_0c8eb76\",\"startPoint\":{\"x\":435,\"y\":360},\"endPoint\":{\"x\":540,\"y\":360},\"properties\":{},\"zIndex\":1012,\"pointsList\":[{\"x\":435,\"y\":360},{\"x\":540,\"y\":360}]},{\"id\":\"3f872b96-1abf-4051-9c12-b6fe4e159b00\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Activity_0c8eb76\",\"targetNodeId\":\"Event_af85d64\",\"startPoint\":{\"x\":640,\"y\":360},\"endPoint\":{\"x\":1172,\"y\":360},\"properties\":{},\"zIndex\":1027,\"pointsList\":[{\"x\":640,\"y\":360},{\"x\":1172,\"y\":360}]},{\"id\":\"16833e07-5735-48dd-83f2-e14c04b8fe9b\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Activity_0c4ec56\",\"targetNodeId\":\"Activity_88d7be8\",\"startPoint\":{\"x\":630,\"y\":180},\"endPoint\":{\"x\":760,\"y\":180},\"properties\":{},\"zIndex\":1004,\"pointsList\":[{\"x\":630,\"y\":180},{\"x\":760,\"y\":180}]},{\"id\":\"15929f6f-cf6e-465c-848c-c31fdd4e89d2\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Activity_88d7be8\",\"targetNodeId\":\"Activity_f09936b\",\"startPoint\":{\"x\":860,\"y\":180},\"endPoint\":{\"x\":980,\"y\":180},\"properties\":{},\"zIndex\":1009,\"pointsList\":[{\"x\":860,\"y\":180},{\"x\":980,\"y\":180}]},{\"id\":\"0ba7f6ae-1101-4c7d-ac4f-b01f755f518f\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Activity_f09936b\",\"targetNodeId\":\"Event_af85d64\",\"startPoint\":{\"x\":1080,\"y\":180},\"endPoint\":{\"x\":1190,\"y\":342},\"properties\":{},\"zIndex\":1002,\"pointsList\":[{\"x\":1080,\"y\":180},{\"x\":1190,\"y\":180},{\"x\":1190,\"y\":342}]}]}', '[{\"id\":\"Event_5026d13\",\"pid\":0,\"label\":\"开始\",\"type\":\"bpmn:startEvent\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Gateway_5c0df6e\",\"pid\":\"Event_5026d13\",\"label\":\"网关\",\"type\":\"bpmn:exclusiveGateway\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"gateway\":[{\"id\":\"textarea53158\",\"value\":\"1\",\"condition\":\"==\"}],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_0c4ec56\",\"pid\":\"Gateway_5c0df6e\",\"label\":\"系统任务\",\"type\":\"bpmn:serviceTask\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"gateway\":[],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_88d7be8\",\"pid\":\"Activity_0c4ec56\",\"label\":\"审批1\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"textarea53158\":2,\"textarea9642\":2,\"select101798\":2,\"input47807\":2},\"gateway\":[],\"userType\":3,\"userId\":1,\"deptId\":3,\"postId\":2,\"children\":[{\"id\":\"Activity_f09936b\",\"pid\":\"Activity_88d7be8\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{},\"gateway\":[],\"userType\":3,\"userId\":1,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_af85d64\",\"pid\":\"Activity_f09936b\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0}]}]}]}]},{\"id\":\"Gateway_9e9bedf\",\"pid\":\"Event_5026d13\",\"label\":\"网关\",\"type\":\"bpmn:exclusiveGateway\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"gateway\":[{\"id\":\"textarea53158\",\"value\":\"2\",\"condition\":\"==\"}],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_0c8eb76\",\"pid\":\"Gateway_9e9bedf\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"gateway\":[],\"userType\":1,\"userId\":0,\"deptId\":2,\"postId\":3,\"children\":[{\"id\":\"Event_af85d64\",\"pid\":\"Activity_0c8eb76\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0}]}]}]},{\"id\":\"Gateway_5c0df6e\",\"pid\":\"Event_5026d13\",\"label\":\"网关\",\"type\":\"bpmn:exclusiveGateway\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"gateway\":[{\"id\":\"textarea53158\",\"value\":\"1\",\"condition\":\"==\"}],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_0c4ec56\",\"pid\":\"Gateway_5c0df6e\",\"label\":\"系统任务\",\"type\":\"bpmn:serviceTask\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"gateway\":[],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_88d7be8\",\"pid\":\"Activity_0c4ec56\",\"label\":\"审批1\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"textarea53158\":2,\"textarea9642\":2,\"select101798\":2,\"input47807\":2},\"gateway\":[],\"userType\":3,\"userId\":1,\"deptId\":3,\"postId\":2,\"children\":[{\"id\":\"Activity_f09936b\",\"pid\":\"Activity_88d7be8\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{},\"gateway\":[],\"userType\":3,\"userId\":1,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_af85d64\",\"pid\":\"Activity_f09936b\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0}]}]}]}]},{\"id\":\"Activity_0c4ec56\",\"pid\":\"Gateway_5c0df6e\",\"label\":\"系统任务\",\"type\":\"bpmn:serviceTask\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"gateway\":[],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_88d7be8\",\"pid\":\"Activity_0c4ec56\",\"label\":\"审批1\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"textarea53158\":2,\"textarea9642\":2,\"select101798\":2,\"input47807\":2},\"gateway\":[],\"userType\":3,\"userId\":1,\"deptId\":3,\"postId\":2,\"children\":[{\"id\":\"Activity_f09936b\",\"pid\":\"Activity_88d7be8\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{},\"gateway\":[],\"userType\":3,\"userId\":1,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_af85d64\",\"pid\":\"Activity_f09936b\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0}]}]}]},{\"id\":\"Activity_88d7be8\",\"pid\":\"Activity_0c4ec56\",\"label\":\"审批1\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"textarea53158\":2,\"textarea9642\":2,\"select101798\":2,\"input47807\":2},\"gateway\":[],\"userType\":3,\"userId\":1,\"deptId\":3,\"postId\":2,\"children\":[{\"id\":\"Activity_f09936b\",\"pid\":\"Activity_88d7be8\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{},\"gateway\":[],\"userType\":3,\"userId\":1,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_af85d64\",\"pid\":\"Activity_f09936b\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0}]}]},{\"id\":\"Activity_f09936b\",\"pid\":\"Activity_88d7be8\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{},\"gateway\":[],\"userType\":3,\"userId\":1,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_af85d64\",\"pid\":\"Activity_f09936b\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0}]},{\"id\":\"Event_af85d64\",\"pid\":\"Activity_f09936b\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0},{\"id\":\"Gateway_9e9bedf\",\"pid\":\"Event_5026d13\",\"label\":\"网关\",\"type\":\"bpmn:exclusiveGateway\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"gateway\":[{\"id\":\"textarea53158\",\"value\":\"2\",\"condition\":\"==\"}],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_0c8eb76\",\"pid\":\"Gateway_9e9bedf\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"gateway\":[],\"userType\":1,\"userId\":0,\"deptId\":2,\"postId\":3,\"children\":[{\"id\":\"Event_af85d64\",\"pid\":\"Activity_0c8eb76\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0}]}]},{\"id\":\"Activity_0c8eb76\",\"pid\":\"Gateway_9e9bedf\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"gateway\":[],\"userType\":1,\"userId\":0,\"deptId\":2,\"postId\":3,\"children\":[{\"id\":\"Event_af85d64\",\"pid\":\"Activity_0c8eb76\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0}]},{\"id\":\"Event_af85d64\",\"pid\":\"Activity_0c8eb76\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0}]', '{\"textarea53158\":\"1\",\"select101798\":\"\",\"input47807\":\"\"}', 3, 1703781731, 1703761264, 0);
INSERT INTO `x_flow_apply` VALUES (26, 2, 1, 'admin', '2', 2, '2', '{\"widgetList\":[{\"key\":93225,\"type\":\"textarea\",\"icon\":\"textarea-field\",\"formItemFlag\":true,\"options\":{\"name\":\"textarea53158\",\"label\":\"textarea\",\"labelAlign\":\"\",\"rows\":3,\"defaultValue\":\"\",\"placeholder\":\"\",\"columnWidth\":\"200px\",\"size\":\"\",\"labelWidth\":null,\"labelHidden\":false,\"readonly\":false,\"disabled\":false,\"hidden\":false,\"required\":false,\"requiredHint\":\"\",\"validation\":\"\",\"validationHint\":\"\",\"customClass\":[],\"labelIconClass\":null,\"labelIconPosition\":\"rear\",\"labelTooltip\":null,\"minLength\":null,\"maxLength\":null,\"showWordLimit\":false,\"onCreated\":\"\",\"onMounted\":\"\",\"onInput\":\"\",\"onChange\":\"\",\"onFocus\":\"\",\"onBlur\":\"\",\"onValidate\":\"\"},\"id\":\"textarea53158\"},{\"key\":35486,\"type\":\"select\",\"icon\":\"select-field\",\"formItemFlag\":true,\"options\":{\"name\":\"select101798\",\"label\":\"select\",\"labelAlign\":\"\",\"defaultValue\":\"\",\"placeholder\":\"\",\"columnWidth\":\"200px\",\"size\":\"\",\"labelWidth\":null,\"labelHidden\":false,\"disabled\":false,\"hidden\":false,\"clearable\":true,\"filterable\":false,\"allowCreate\":false,\"remote\":false,\"automaticDropdown\":false,\"multiple\":false,\"multipleLimit\":0,\"optionItems\":[{\"label\":\"select 1\",\"value\":1},{\"label\":\"select 2\",\"value\":2},{\"label\":\"select 3\",\"value\":3}],\"required\":false,\"requiredHint\":\"\",\"validation\":\"\",\"validationHint\":\"\",\"customClass\":\"\",\"labelIconClass\":null,\"labelIconPosition\":\"rear\",\"labelTooltip\":null,\"onCreated\":\"\",\"onMounted\":\"\",\"onRemoteQuery\":\"\",\"onChange\":\"\",\"onFocus\":\"\",\"onBlur\":\"\",\"onValidate\":\"\"},\"id\":\"select101798\"},{\"key\":85807,\"type\":\"input\",\"icon\":\"text-field\",\"formItemFlag\":true,\"options\":{\"name\":\"input47807\",\"label\":\"input\",\"labelAlign\":\"\",\"type\":\"text\",\"defaultValue\":\"\",\"placeholder\":\"\",\"columnWidth\":\"200px\",\"size\":\"\",\"labelWidth\":null,\"labelHidden\":false,\"readonly\":false,\"disabled\":false,\"hidden\":false,\"clearable\":true,\"showPassword\":false,\"required\":false,\"requiredHint\":\"\",\"validation\":\"\",\"validationHint\":\"\",\"customClass\":\"\",\"labelIconClass\":null,\"labelIconPosition\":\"rear\",\"labelTooltip\":null,\"minLength\":null,\"maxLength\":null,\"showWordLimit\":false,\"prefixIcon\":\"\",\"suffixIcon\":\"\",\"appendButton\":false,\"appendButtonDisabled\":false,\"buttonIcon\":\"custom-search\",\"onCreated\":\"\",\"onMounted\":\"\",\"onInput\":\"\",\"onChange\":\"\",\"onFocus\":\"\",\"onBlur\":\"\",\"onValidate\":\"\",\"onAppendButtonClick\":\"\"},\"id\":\"input47807\"}],\"formConfig\":{\"modelName\":\"formData\",\"refName\":\"vForm\",\"rulesName\":\"rules\",\"labelWidth\":80,\"labelPosition\":\"left\",\"size\":\"\",\"labelAlign\":\"label-left-align\",\"cssCode\":\"\",\"customClass\":[],\"functions\":\"\",\"layoutType\":\"H5\",\"jsonVersion\":3,\"onFormCreated\":\"\",\"onFormMounted\":\"\",\"onFormDataChange\":\"\"}}', '{\"nodes\":[{\"id\":\"Event_5026d13\",\"type\":\"bpmn:startEvent\",\"x\":260,\"y\":250,\"properties\":{\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":\"\",\"userId\":\"\",\"deptId\":\"\",\"postId\":\"\"},\"zIndex\":1005,\"text\":{\"x\":260,\"y\":290,\"value\":\"开始\"}},{\"id\":\"Activity_88d7be8\",\"type\":\"bpmn:userTask\",\"x\":810,\"y\":180,\"properties\":{\"fieldAuth\":{\"textarea53158\":2,\"textarea9642\":2,\"select101798\":2,\"input47807\":2},\"userId\":1,\"deptId\":3,\"postId\":2,\"userType\":3,\"gateway\":[]},\"zIndex\":1011,\"text\":{\"x\":810,\"y\":180,\"value\":\"审批1\"}},{\"id\":\"Event_af85d64\",\"type\":\"bpmn:endEvent\",\"x\":1190,\"y\":360,\"properties\":{\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":\"\",\"userId\":\"\",\"deptId\":\"\",\"postId\":\"\"},\"zIndex\":1015,\"text\":{\"x\":1190,\"y\":400,\"value\":\"结束\"}},{\"id\":\"Gateway_5c0df6e\",\"type\":\"bpmn:exclusiveGateway\",\"x\":410,\"y\":180,\"properties\":{\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userId\":\"\",\"deptId\":\"\",\"postId\":\"\",\"userType\":\"\",\"gateway\":[{\"id\":\"textarea53158\",\"value\":\"1\",\"condition\":\"==\"}]},\"zIndex\":1005,\"text\":{\"x\":410,\"y\":220,\"value\":\"网关\"}},{\"id\":\"Activity_0c4ec56\",\"type\":\"bpmn:serviceTask\",\"x\":580,\"y\":180,\"properties\":{\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userId\":\"\",\"deptId\":\"\",\"postId\":\"\",\"userType\":\"\",\"gateway\":[]},\"zIndex\":1003,\"text\":{\"x\":580,\"y\":180,\"value\":\"系统任务\"}},{\"id\":\"Gateway_9e9bedf\",\"type\":\"bpmn:exclusiveGateway\",\"x\":410,\"y\":360,\"properties\":{\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":\"\",\"userId\":\"\",\"deptId\":\"\",\"postId\":\"\",\"gateway\":[{\"id\":\"textarea53158\",\"value\":\"2\",\"condition\":\"==\"}]},\"zIndex\":1001,\"text\":{\"x\":410,\"y\":400,\"value\":\"网关\"}},{\"id\":\"Activity_0c8eb76\",\"type\":\"bpmn:userTask\",\"x\":590,\"y\":360,\"properties\":{\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":1,\"userId\":\"\",\"deptId\":2,\"postId\":3,\"gateway\":[]},\"zIndex\":1012,\"text\":{\"x\":590,\"y\":360,\"value\":\"审批\"}},{\"id\":\"Activity_f09936b\",\"type\":\"bpmn:userTask\",\"x\":1030,\"y\":180,\"properties\":{\"userType\":3,\"userId\":1,\"deptId\":\"\",\"postId\":\"\",\"fieldAuth\":{},\"gateway\":[]},\"zIndex\":1001,\"text\":{\"x\":1030,\"y\":180,\"value\":\"审批\"}}],\"edges\":[{\"id\":\"c7d055f8-bbce-40e5-a2b4-5b0fb1ee9aea\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Event_5026d13\",\"targetNodeId\":\"Gateway_5c0df6e\",\"startPoint\":{\"x\":278,\"y\":250},\"endPoint\":{\"x\":385,\"y\":180},\"properties\":{},\"zIndex\":1009,\"pointsList\":[{\"x\":278,\"y\":250},{\"x\":355,\"y\":250},{\"x\":355,\"y\":180},{\"x\":385,\"y\":180}]},{\"id\":\"ad29adbd-060c-4e0e-9c54-a5512f2542b0\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Gateway_5c0df6e\",\"targetNodeId\":\"Activity_0c4ec56\",\"startPoint\":{\"x\":435,\"y\":180},\"endPoint\":{\"x\":530,\"y\":180},\"properties\":{},\"zIndex\":1016,\"pointsList\":[{\"x\":435,\"y\":180},{\"x\":530,\"y\":180}]},{\"id\":\"2146dbe0-9e6d-4c20-ab84-6f66d3f5f4ea\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Event_5026d13\",\"targetNodeId\":\"Gateway_9e9bedf\",\"startPoint\":{\"x\":278,\"y\":250},\"endPoint\":{\"x\":385,\"y\":360},\"properties\":{},\"zIndex\":1008,\"pointsList\":[{\"x\":278,\"y\":250},{\"x\":355,\"y\":250},{\"x\":355,\"y\":360},{\"x\":385,\"y\":360}]},{\"id\":\"7bfc16b1-a40b-4a53-8599-498a73dede74\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Gateway_9e9bedf\",\"targetNodeId\":\"Activity_0c8eb76\",\"startPoint\":{\"x\":435,\"y\":360},\"endPoint\":{\"x\":540,\"y\":360},\"properties\":{},\"zIndex\":1012,\"pointsList\":[{\"x\":435,\"y\":360},{\"x\":540,\"y\":360}]},{\"id\":\"3f872b96-1abf-4051-9c12-b6fe4e159b00\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Activity_0c8eb76\",\"targetNodeId\":\"Event_af85d64\",\"startPoint\":{\"x\":640,\"y\":360},\"endPoint\":{\"x\":1172,\"y\":360},\"properties\":{},\"zIndex\":1027,\"pointsList\":[{\"x\":640,\"y\":360},{\"x\":1172,\"y\":360}]},{\"id\":\"16833e07-5735-48dd-83f2-e14c04b8fe9b\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Activity_0c4ec56\",\"targetNodeId\":\"Activity_88d7be8\",\"startPoint\":{\"x\":630,\"y\":180},\"endPoint\":{\"x\":760,\"y\":180},\"properties\":{},\"zIndex\":1004,\"pointsList\":[{\"x\":630,\"y\":180},{\"x\":760,\"y\":180}]},{\"id\":\"15929f6f-cf6e-465c-848c-c31fdd4e89d2\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Activity_88d7be8\",\"targetNodeId\":\"Activity_f09936b\",\"startPoint\":{\"x\":860,\"y\":180},\"endPoint\":{\"x\":980,\"y\":180},\"properties\":{},\"zIndex\":1009,\"pointsList\":[{\"x\":860,\"y\":180},{\"x\":980,\"y\":180}]},{\"id\":\"0ba7f6ae-1101-4c7d-ac4f-b01f755f518f\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Activity_f09936b\",\"targetNodeId\":\"Event_af85d64\",\"startPoint\":{\"x\":1080,\"y\":180},\"endPoint\":{\"x\":1190,\"y\":342},\"properties\":{},\"zIndex\":1002,\"pointsList\":[{\"x\":1080,\"y\":180},{\"x\":1190,\"y\":180},{\"x\":1190,\"y\":342}]}]}', '[{\"id\":\"Event_5026d13\",\"pid\":0,\"label\":\"开始\",\"type\":\"bpmn:startEvent\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Gateway_5c0df6e\",\"pid\":\"Event_5026d13\",\"label\":\"网关\",\"type\":\"bpmn:exclusiveGateway\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"gateway\":[{\"id\":\"textarea53158\",\"value\":\"1\",\"condition\":\"==\"}],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_0c4ec56\",\"pid\":\"Gateway_5c0df6e\",\"label\":\"系统任务\",\"type\":\"bpmn:serviceTask\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"gateway\":[],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_88d7be8\",\"pid\":\"Activity_0c4ec56\",\"label\":\"审批1\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"textarea53158\":2,\"textarea9642\":2,\"select101798\":2,\"input47807\":2},\"gateway\":[],\"userType\":3,\"userId\":1,\"deptId\":3,\"postId\":2,\"children\":[{\"id\":\"Activity_f09936b\",\"pid\":\"Activity_88d7be8\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{},\"gateway\":[],\"userType\":3,\"userId\":1,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_af85d64\",\"pid\":\"Activity_f09936b\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0}]}]}]}]},{\"id\":\"Gateway_9e9bedf\",\"pid\":\"Event_5026d13\",\"label\":\"网关\",\"type\":\"bpmn:exclusiveGateway\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"gateway\":[{\"id\":\"textarea53158\",\"value\":\"2\",\"condition\":\"==\"}],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_0c8eb76\",\"pid\":\"Gateway_9e9bedf\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"gateway\":[],\"userType\":1,\"userId\":0,\"deptId\":2,\"postId\":3,\"children\":[{\"id\":\"Event_af85d64\",\"pid\":\"Activity_0c8eb76\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0}]}]}]},{\"id\":\"Gateway_5c0df6e\",\"pid\":\"Event_5026d13\",\"label\":\"网关\",\"type\":\"bpmn:exclusiveGateway\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"gateway\":[{\"id\":\"textarea53158\",\"value\":\"1\",\"condition\":\"==\"}],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_0c4ec56\",\"pid\":\"Gateway_5c0df6e\",\"label\":\"系统任务\",\"type\":\"bpmn:serviceTask\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"gateway\":[],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_88d7be8\",\"pid\":\"Activity_0c4ec56\",\"label\":\"审批1\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"textarea53158\":2,\"textarea9642\":2,\"select101798\":2,\"input47807\":2},\"gateway\":[],\"userType\":3,\"userId\":1,\"deptId\":3,\"postId\":2,\"children\":[{\"id\":\"Activity_f09936b\",\"pid\":\"Activity_88d7be8\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{},\"gateway\":[],\"userType\":3,\"userId\":1,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_af85d64\",\"pid\":\"Activity_f09936b\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0}]}]}]}]},{\"id\":\"Activity_0c4ec56\",\"pid\":\"Gateway_5c0df6e\",\"label\":\"系统任务\",\"type\":\"bpmn:serviceTask\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"gateway\":[],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_88d7be8\",\"pid\":\"Activity_0c4ec56\",\"label\":\"审批1\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"textarea53158\":2,\"textarea9642\":2,\"select101798\":2,\"input47807\":2},\"gateway\":[],\"userType\":3,\"userId\":1,\"deptId\":3,\"postId\":2,\"children\":[{\"id\":\"Activity_f09936b\",\"pid\":\"Activity_88d7be8\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{},\"gateway\":[],\"userType\":3,\"userId\":1,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_af85d64\",\"pid\":\"Activity_f09936b\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0}]}]}]},{\"id\":\"Activity_88d7be8\",\"pid\":\"Activity_0c4ec56\",\"label\":\"审批1\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"textarea53158\":2,\"textarea9642\":2,\"select101798\":2,\"input47807\":2},\"gateway\":[],\"userType\":3,\"userId\":1,\"deptId\":3,\"postId\":2,\"children\":[{\"id\":\"Activity_f09936b\",\"pid\":\"Activity_88d7be8\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{},\"gateway\":[],\"userType\":3,\"userId\":1,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_af85d64\",\"pid\":\"Activity_f09936b\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0}]}]},{\"id\":\"Activity_f09936b\",\"pid\":\"Activity_88d7be8\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{},\"gateway\":[],\"userType\":3,\"userId\":1,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_af85d64\",\"pid\":\"Activity_f09936b\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0}]},{\"id\":\"Event_af85d64\",\"pid\":\"Activity_f09936b\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0},{\"id\":\"Gateway_9e9bedf\",\"pid\":\"Event_5026d13\",\"label\":\"网关\",\"type\":\"bpmn:exclusiveGateway\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"gateway\":[{\"id\":\"textarea53158\",\"value\":\"2\",\"condition\":\"==\"}],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_0c8eb76\",\"pid\":\"Gateway_9e9bedf\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"gateway\":[],\"userType\":1,\"userId\":0,\"deptId\":2,\"postId\":3,\"children\":[{\"id\":\"Event_af85d64\",\"pid\":\"Activity_0c8eb76\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0}]}]},{\"id\":\"Activity_0c8eb76\",\"pid\":\"Gateway_9e9bedf\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"gateway\":[],\"userType\":1,\"userId\":0,\"deptId\":2,\"postId\":3,\"children\":[{\"id\":\"Event_af85d64\",\"pid\":\"Activity_0c8eb76\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0}]},{\"id\":\"Event_af85d64\",\"pid\":\"Activity_0c8eb76\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0}]', '{\"textarea53158\":\"2\",\"select101798\":\"\",\"input47807\":\"\"}', 2, 1703767009, 1703766995, 0);
INSERT INTO `x_flow_apply` VALUES (27, 2, 1, 'admin', '2', 2, '2', '{\"widgetList\":[{\"key\":93225,\"type\":\"textarea\",\"icon\":\"textarea-field\",\"formItemFlag\":true,\"options\":{\"name\":\"textarea53158\",\"label\":\"textarea\",\"labelAlign\":\"\",\"rows\":3,\"defaultValue\":\"\",\"placeholder\":\"\",\"columnWidth\":\"200px\",\"size\":\"\",\"labelWidth\":null,\"labelHidden\":false,\"readonly\":false,\"disabled\":false,\"hidden\":false,\"required\":false,\"requiredHint\":\"\",\"validation\":\"\",\"validationHint\":\"\",\"customClass\":[],\"labelIconClass\":null,\"labelIconPosition\":\"rear\",\"labelTooltip\":null,\"minLength\":null,\"maxLength\":null,\"showWordLimit\":false,\"onCreated\":\"\",\"onMounted\":\"\",\"onInput\":\"\",\"onChange\":\"\",\"onFocus\":\"\",\"onBlur\":\"\",\"onValidate\":\"\"},\"id\":\"textarea53158\"},{\"key\":35486,\"type\":\"select\",\"icon\":\"select-field\",\"formItemFlag\":true,\"options\":{\"name\":\"select101798\",\"label\":\"select\",\"labelAlign\":\"\",\"defaultValue\":\"\",\"placeholder\":\"\",\"columnWidth\":\"200px\",\"size\":\"\",\"labelWidth\":null,\"labelHidden\":false,\"disabled\":false,\"hidden\":false,\"clearable\":true,\"filterable\":false,\"allowCreate\":false,\"remote\":false,\"automaticDropdown\":false,\"multiple\":false,\"multipleLimit\":0,\"optionItems\":[{\"label\":\"select 1\",\"value\":1},{\"label\":\"select 2\",\"value\":2},{\"label\":\"select 3\",\"value\":3}],\"required\":false,\"requiredHint\":\"\",\"validation\":\"\",\"validationHint\":\"\",\"customClass\":\"\",\"labelIconClass\":null,\"labelIconPosition\":\"rear\",\"labelTooltip\":null,\"onCreated\":\"\",\"onMounted\":\"\",\"onRemoteQuery\":\"\",\"onChange\":\"\",\"onFocus\":\"\",\"onBlur\":\"\",\"onValidate\":\"\"},\"id\":\"select101798\"},{\"key\":85807,\"type\":\"input\",\"icon\":\"text-field\",\"formItemFlag\":true,\"options\":{\"name\":\"input47807\",\"label\":\"input\",\"labelAlign\":\"\",\"type\":\"text\",\"defaultValue\":\"\",\"placeholder\":\"\",\"columnWidth\":\"200px\",\"size\":\"\",\"labelWidth\":null,\"labelHidden\":false,\"readonly\":false,\"disabled\":false,\"hidden\":false,\"clearable\":true,\"showPassword\":false,\"required\":false,\"requiredHint\":\"\",\"validation\":\"\",\"validationHint\":\"\",\"customClass\":\"\",\"labelIconClass\":null,\"labelIconPosition\":\"rear\",\"labelTooltip\":null,\"minLength\":null,\"maxLength\":null,\"showWordLimit\":false,\"prefixIcon\":\"\",\"suffixIcon\":\"\",\"appendButton\":false,\"appendButtonDisabled\":false,\"buttonIcon\":\"custom-search\",\"onCreated\":\"\",\"onMounted\":\"\",\"onInput\":\"\",\"onChange\":\"\",\"onFocus\":\"\",\"onBlur\":\"\",\"onValidate\":\"\",\"onAppendButtonClick\":\"\"},\"id\":\"input47807\"}],\"formConfig\":{\"modelName\":\"formData\",\"refName\":\"vForm\",\"rulesName\":\"rules\",\"labelWidth\":80,\"labelPosition\":\"left\",\"size\":\"\",\"labelAlign\":\"label-left-align\",\"cssCode\":\"\",\"customClass\":[],\"functions\":\"\",\"layoutType\":\"H5\",\"jsonVersion\":3,\"onFormCreated\":\"\",\"onFormMounted\":\"\",\"onFormDataChange\":\"\"}}', '{\"nodes\":[{\"id\":\"Event_5026d13\",\"type\":\"bpmn:startEvent\",\"x\":260,\"y\":250,\"properties\":{\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":\"\",\"userId\":\"\",\"deptId\":\"\",\"postId\":\"\"},\"zIndex\":1005,\"text\":{\"x\":260,\"y\":290,\"value\":\"开始\"}},{\"id\":\"Activity_88d7be8\",\"type\":\"bpmn:userTask\",\"x\":810,\"y\":180,\"properties\":{\"fieldAuth\":{\"textarea53158\":2,\"textarea9642\":2,\"select101798\":2,\"input47807\":2},\"userId\":1,\"deptId\":3,\"postId\":2,\"userType\":3,\"gateway\":[]},\"zIndex\":1011,\"text\":{\"x\":810,\"y\":180,\"value\":\"审批1\"}},{\"id\":\"Event_af85d64\",\"type\":\"bpmn:endEvent\",\"x\":1190,\"y\":360,\"properties\":{\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":\"\",\"userId\":\"\",\"deptId\":\"\",\"postId\":\"\"},\"zIndex\":1015,\"text\":{\"x\":1190,\"y\":400,\"value\":\"结束\"}},{\"id\":\"Gateway_5c0df6e\",\"type\":\"bpmn:exclusiveGateway\",\"x\":410,\"y\":180,\"properties\":{\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userId\":\"\",\"deptId\":\"\",\"postId\":\"\",\"userType\":\"\",\"gateway\":[{\"id\":\"textarea53158\",\"value\":\"1\",\"condition\":\"==\"}]},\"zIndex\":1005,\"text\":{\"x\":410,\"y\":220,\"value\":\"网关\"}},{\"id\":\"Activity_0c4ec56\",\"type\":\"bpmn:serviceTask\",\"x\":580,\"y\":180,\"properties\":{\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userId\":\"\",\"deptId\":\"\",\"postId\":\"\",\"userType\":\"\",\"gateway\":[]},\"zIndex\":1003,\"text\":{\"x\":580,\"y\":180,\"value\":\"系统任务\"}},{\"id\":\"Gateway_9e9bedf\",\"type\":\"bpmn:exclusiveGateway\",\"x\":410,\"y\":360,\"properties\":{\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":\"\",\"userId\":\"\",\"deptId\":\"\",\"postId\":\"\",\"gateway\":[{\"id\":\"textarea53158\",\"value\":\"2\",\"condition\":\"==\"}]},\"zIndex\":1001,\"text\":{\"x\":410,\"y\":400,\"value\":\"网关\"}},{\"id\":\"Activity_0c8eb76\",\"type\":\"bpmn:userTask\",\"x\":590,\"y\":360,\"properties\":{\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":1,\"userId\":\"\",\"deptId\":2,\"postId\":3,\"gateway\":[]},\"zIndex\":1012,\"text\":{\"x\":590,\"y\":360,\"value\":\"审批\"}},{\"id\":\"Activity_f09936b\",\"type\":\"bpmn:userTask\",\"x\":1030,\"y\":180,\"properties\":{\"userType\":3,\"userId\":1,\"deptId\":\"\",\"postId\":\"\",\"fieldAuth\":{},\"gateway\":[]},\"zIndex\":1001,\"text\":{\"x\":1030,\"y\":180,\"value\":\"审批\"}}],\"edges\":[{\"id\":\"c7d055f8-bbce-40e5-a2b4-5b0fb1ee9aea\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Event_5026d13\",\"targetNodeId\":\"Gateway_5c0df6e\",\"startPoint\":{\"x\":278,\"y\":250},\"endPoint\":{\"x\":385,\"y\":180},\"properties\":{},\"zIndex\":1009,\"pointsList\":[{\"x\":278,\"y\":250},{\"x\":355,\"y\":250},{\"x\":355,\"y\":180},{\"x\":385,\"y\":180}]},{\"id\":\"ad29adbd-060c-4e0e-9c54-a5512f2542b0\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Gateway_5c0df6e\",\"targetNodeId\":\"Activity_0c4ec56\",\"startPoint\":{\"x\":435,\"y\":180},\"endPoint\":{\"x\":530,\"y\":180},\"properties\":{},\"zIndex\":1016,\"pointsList\":[{\"x\":435,\"y\":180},{\"x\":530,\"y\":180}]},{\"id\":\"2146dbe0-9e6d-4c20-ab84-6f66d3f5f4ea\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Event_5026d13\",\"targetNodeId\":\"Gateway_9e9bedf\",\"startPoint\":{\"x\":278,\"y\":250},\"endPoint\":{\"x\":385,\"y\":360},\"properties\":{},\"zIndex\":1008,\"pointsList\":[{\"x\":278,\"y\":250},{\"x\":355,\"y\":250},{\"x\":355,\"y\":360},{\"x\":385,\"y\":360}]},{\"id\":\"7bfc16b1-a40b-4a53-8599-498a73dede74\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Gateway_9e9bedf\",\"targetNodeId\":\"Activity_0c8eb76\",\"startPoint\":{\"x\":435,\"y\":360},\"endPoint\":{\"x\":540,\"y\":360},\"properties\":{},\"zIndex\":1012,\"pointsList\":[{\"x\":435,\"y\":360},{\"x\":540,\"y\":360}]},{\"id\":\"3f872b96-1abf-4051-9c12-b6fe4e159b00\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Activity_0c8eb76\",\"targetNodeId\":\"Event_af85d64\",\"startPoint\":{\"x\":640,\"y\":360},\"endPoint\":{\"x\":1172,\"y\":360},\"properties\":{},\"zIndex\":1027,\"pointsList\":[{\"x\":640,\"y\":360},{\"x\":1172,\"y\":360}]},{\"id\":\"16833e07-5735-48dd-83f2-e14c04b8fe9b\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Activity_0c4ec56\",\"targetNodeId\":\"Activity_88d7be8\",\"startPoint\":{\"x\":630,\"y\":180},\"endPoint\":{\"x\":760,\"y\":180},\"properties\":{},\"zIndex\":1004,\"pointsList\":[{\"x\":630,\"y\":180},{\"x\":760,\"y\":180}]},{\"id\":\"15929f6f-cf6e-465c-848c-c31fdd4e89d2\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Activity_88d7be8\",\"targetNodeId\":\"Activity_f09936b\",\"startPoint\":{\"x\":860,\"y\":180},\"endPoint\":{\"x\":980,\"y\":180},\"properties\":{},\"zIndex\":1009,\"pointsList\":[{\"x\":860,\"y\":180},{\"x\":980,\"y\":180}]},{\"id\":\"0ba7f6ae-1101-4c7d-ac4f-b01f755f518f\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Activity_f09936b\",\"targetNodeId\":\"Event_af85d64\",\"startPoint\":{\"x\":1080,\"y\":180},\"endPoint\":{\"x\":1190,\"y\":342},\"properties\":{},\"zIndex\":1002,\"pointsList\":[{\"x\":1080,\"y\":180},{\"x\":1190,\"y\":180},{\"x\":1190,\"y\":342}]}]}', '[{\"id\":\"Event_5026d13\",\"pid\":0,\"label\":\"开始\",\"type\":\"bpmn:startEvent\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Gateway_5c0df6e\",\"pid\":\"Event_5026d13\",\"label\":\"网关\",\"type\":\"bpmn:exclusiveGateway\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"gateway\":[{\"id\":\"textarea53158\",\"value\":\"1\",\"condition\":\"==\"}],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_0c4ec56\",\"pid\":\"Gateway_5c0df6e\",\"label\":\"系统任务\",\"type\":\"bpmn:serviceTask\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"gateway\":[],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_88d7be8\",\"pid\":\"Activity_0c4ec56\",\"label\":\"审批1\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"textarea53158\":2,\"textarea9642\":2,\"select101798\":2,\"input47807\":2},\"gateway\":[],\"userType\":3,\"userId\":1,\"deptId\":3,\"postId\":2,\"children\":[{\"id\":\"Activity_f09936b\",\"pid\":\"Activity_88d7be8\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{},\"gateway\":[],\"userType\":3,\"userId\":1,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_af85d64\",\"pid\":\"Activity_f09936b\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0}]}]}]}]},{\"id\":\"Gateway_9e9bedf\",\"pid\":\"Event_5026d13\",\"label\":\"网关\",\"type\":\"bpmn:exclusiveGateway\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"gateway\":[{\"id\":\"textarea53158\",\"value\":\"2\",\"condition\":\"==\"}],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_0c8eb76\",\"pid\":\"Gateway_9e9bedf\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"gateway\":[],\"userType\":1,\"userId\":0,\"deptId\":2,\"postId\":3,\"children\":[{\"id\":\"Event_af85d64\",\"pid\":\"Activity_0c8eb76\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0}]}]}]},{\"id\":\"Gateway_5c0df6e\",\"pid\":\"Event_5026d13\",\"label\":\"网关\",\"type\":\"bpmn:exclusiveGateway\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"gateway\":[{\"id\":\"textarea53158\",\"value\":\"1\",\"condition\":\"==\"}],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_0c4ec56\",\"pid\":\"Gateway_5c0df6e\",\"label\":\"系统任务\",\"type\":\"bpmn:serviceTask\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"gateway\":[],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_88d7be8\",\"pid\":\"Activity_0c4ec56\",\"label\":\"审批1\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"textarea53158\":2,\"textarea9642\":2,\"select101798\":2,\"input47807\":2},\"gateway\":[],\"userType\":3,\"userId\":1,\"deptId\":3,\"postId\":2,\"children\":[{\"id\":\"Activity_f09936b\",\"pid\":\"Activity_88d7be8\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{},\"gateway\":[],\"userType\":3,\"userId\":1,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_af85d64\",\"pid\":\"Activity_f09936b\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0}]}]}]}]},{\"id\":\"Activity_0c4ec56\",\"pid\":\"Gateway_5c0df6e\",\"label\":\"系统任务\",\"type\":\"bpmn:serviceTask\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"gateway\":[],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_88d7be8\",\"pid\":\"Activity_0c4ec56\",\"label\":\"审批1\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"textarea53158\":2,\"textarea9642\":2,\"select101798\":2,\"input47807\":2},\"gateway\":[],\"userType\":3,\"userId\":1,\"deptId\":3,\"postId\":2,\"children\":[{\"id\":\"Activity_f09936b\",\"pid\":\"Activity_88d7be8\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{},\"gateway\":[],\"userType\":3,\"userId\":1,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_af85d64\",\"pid\":\"Activity_f09936b\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0}]}]}]},{\"id\":\"Activity_88d7be8\",\"pid\":\"Activity_0c4ec56\",\"label\":\"审批1\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"textarea53158\":2,\"textarea9642\":2,\"select101798\":2,\"input47807\":2},\"gateway\":[],\"userType\":3,\"userId\":1,\"deptId\":3,\"postId\":2,\"children\":[{\"id\":\"Activity_f09936b\",\"pid\":\"Activity_88d7be8\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{},\"gateway\":[],\"userType\":3,\"userId\":1,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_af85d64\",\"pid\":\"Activity_f09936b\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0}]}]},{\"id\":\"Activity_f09936b\",\"pid\":\"Activity_88d7be8\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{},\"gateway\":[],\"userType\":3,\"userId\":1,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_af85d64\",\"pid\":\"Activity_f09936b\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0}]},{\"id\":\"Event_af85d64\",\"pid\":\"Activity_f09936b\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0},{\"id\":\"Gateway_9e9bedf\",\"pid\":\"Event_5026d13\",\"label\":\"网关\",\"type\":\"bpmn:exclusiveGateway\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"gateway\":[{\"id\":\"textarea53158\",\"value\":\"2\",\"condition\":\"==\"}],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_0c8eb76\",\"pid\":\"Gateway_9e9bedf\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"gateway\":[],\"userType\":1,\"userId\":0,\"deptId\":2,\"postId\":3,\"children\":[{\"id\":\"Event_af85d64\",\"pid\":\"Activity_0c8eb76\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0}]}]},{\"id\":\"Activity_0c8eb76\",\"pid\":\"Gateway_9e9bedf\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"gateway\":[],\"userType\":1,\"userId\":0,\"deptId\":2,\"postId\":3,\"children\":[{\"id\":\"Event_af85d64\",\"pid\":\"Activity_0c8eb76\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0}]},{\"id\":\"Event_af85d64\",\"pid\":\"Activity_0c8eb76\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0}]', '', 1, 1703767519, 1703767519, 0);

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
  `update_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `delete_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `apply_id`(`apply_id`) USING BTREE,
  INDEX `approver_id`(`approver_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 82 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '流程历史' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_flow_history
-- ----------------------------
INSERT INTO `x_flow_history` VALUES (71, 25, 2, 1, 'admin', 0, '', 'Event_5026d13', 'bpmn:startEvent', '开始', '{\"textarea53158\":\"1\",\"select101798\":\"\",\"input47807\":\"\"}', 2, '', 1703761271, 1703761271, 0);
INSERT INTO `x_flow_history` VALUES (72, 25, 2, 1, 'admin', 0, '', 'Gateway_5c0df6e', 'bpmn:exclusiveGateway', '网关', '{\"textarea53158\":\"1\",\"select101798\":\"\",\"input47807\":\"\"}', 2, '', 1703761271, 1703761271, 0);
INSERT INTO `x_flow_history` VALUES (73, 25, 2, 1, 'admin', 0, '', 'Activity_0c4ec56', 'bpmn:serviceTask', '系统任务', '{\"textarea53158\":\"1\",\"select101798\":\"\",\"input47807\":\"\"}', 1, '', 1703761271, 1703761271, 0);
INSERT INTO `x_flow_history` VALUES (74, 25, 2, 1, 'admin', 1, '', 'Activity_88d7be8', 'bpmn:userTask', '审批1', '{\"textarea53158\":\"1\",\"select101798\":\"\",\"input47807\":\"\"}', 2, '1', 1703761304, 1703761271, 0);
INSERT INTO `x_flow_history` VALUES (75, 25, 2, 1, 'admin', 1, '', 'Activity_f09936b', 'bpmn:userTask', '审批', '{\"textarea53158\":\"1\",\"select101798\":\"\",\"input47807\":\"\"}', 3, '123', 1703761569, 1703761304, 0);
INSERT INTO `x_flow_history` VALUES (76, 25, 2, 1, 'admin', 1, '', 'Activity_88d7be8', 'bpmn:userTask', '审批1', '{\"textarea53158\":\"1\",\"select101798\":2,\"input47807\":\"6\"}', 2, '9', 1703781154, 1703761569, 0);
INSERT INTO `x_flow_history` VALUES (77, 26, 2, 1, 'admin', 0, '', 'Event_5026d13', 'bpmn:startEvent', '开始', '{\"textarea53158\":\"2\",\"select101798\":\"\",\"input47807\":\"\"}', 2, '', 1703767009, 1703767009, 0);
INSERT INTO `x_flow_history` VALUES (78, 26, 2, 1, 'admin', 0, '', 'Gateway_9e9bedf', 'bpmn:exclusiveGateway', '网关', '{\"textarea53158\":\"2\",\"select101798\":\"\",\"input47807\":\"\"}', 2, '', 1703767009, 1703767009, 0);
INSERT INTO `x_flow_history` VALUES (79, 26, 2, 1, 'admin', 2, '', 'Activity_0c8eb76', 'bpmn:userTask', '审批', '{\"textarea53158\":\"2\",\"select101798\":\"\",\"input47807\":\"\"}', 1, '', 1703767009, 1703767009, 0);
INSERT INTO `x_flow_history` VALUES (80, 25, 2, 1, 'admin', 1, '', 'Activity_f09936b', 'bpmn:userTask', '审批', '{\"textarea53158\":\"1\",\"select101798\":2,\"input47807\":\"6\"}', 2, '6', 1703781731, 1703781154, 0);
INSERT INTO `x_flow_history` VALUES (81, 25, 2, 1, 'admin', 0, '', 'Event_af85d64', 'bpmn:endEvent', '结束', '{\"textarea53158\":\"1\",\"select101798\":2,\"input47807\":\"6\"}', 2, '', 1703781731, 1703781731, 0);

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
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  `delete_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '流程模板' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of x_flow_template
-- ----------------------------
INSERT INTO `x_flow_template` VALUES (1, '1', 1, '1', '{\"widgetList\":[{\"key\":40082,\"type\":\"table\",\"category\":\"container\",\"icon\":\"table\",\"rows\":[{\"cols\":[{\"type\":\"table-cell\",\"category\":\"container\",\"icon\":\"table-cell\",\"internal\":true,\"widgetList\":[],\"merged\":false,\"options\":{\"name\":\"table-cell-42449\",\"cellWidth\":\"\",\"cellHeight\":\"\",\"colspan\":1,\"rowspan\":1,\"customClass\":\"\"},\"id\":\"table-cell-42449\"}],\"id\":\"table-row-100944\",\"merged\":false}],\"options\":{\"name\":\"table100586\",\"hidden\":false,\"customClass\":[]},\"id\":\"table100586\"},{\"key\":47677,\"type\":\"time-range\",\"icon\":\"time-range-field\",\"formItemFlag\":true,\"options\":{\"name\":\"timerange73058\",\"label\":\"time-range\",\"labelAlign\":\"\",\"defaultValue\":null,\"startPlaceholder\":\"\",\"endPlaceholder\":\"\",\"columnWidth\":\"200px\",\"size\":\"\",\"autoFullWidth\":true,\"labelWidth\":null,\"labelHidden\":false,\"readonly\":false,\"disabled\":false,\"hidden\":false,\"clearable\":true,\"editable\":false,\"format\":\"HH:mm:ss\",\"required\":false,\"requiredHint\":\"\",\"validation\":\"\",\"validationHint\":\"\",\"customClass\":[],\"labelIconClass\":null,\"labelIconPosition\":\"rear\",\"labelTooltip\":null,\"onCreated\":\"\",\"onMounted\":\"\",\"onChange\":\"\",\"onFocus\":\"\",\"onBlur\":\"\",\"onValidate\":\"\"},\"id\":\"timerange73058\"},{\"key\":93225,\"type\":\"textarea\",\"icon\":\"textarea-field\",\"formItemFlag\":true,\"options\":{\"name\":\"textarea53158\",\"label\":\"textarea\",\"labelAlign\":\"\",\"rows\":3,\"defaultValue\":\"\",\"placeholder\":\"\",\"columnWidth\":\"200px\",\"size\":\"\",\"labelWidth\":null,\"labelHidden\":false,\"readonly\":false,\"disabled\":false,\"hidden\":false,\"required\":false,\"requiredHint\":\"\",\"validation\":\"\",\"validationHint\":\"\",\"customClass\":[],\"labelIconClass\":null,\"labelIconPosition\":\"rear\",\"labelTooltip\":null,\"minLength\":null,\"maxLength\":null,\"showWordLimit\":false,\"onCreated\":\"\",\"onMounted\":\"\",\"onInput\":\"\",\"onChange\":\"\",\"onFocus\":\"\",\"onBlur\":\"\",\"onValidate\":\"\"},\"id\":\"textarea53158\"}],\"formConfig\":{\"modelName\":\"formData\",\"refName\":\"vForm\",\"rulesName\":\"rules\",\"labelWidth\":80,\"labelPosition\":\"left\",\"size\":\"\",\"labelAlign\":\"label-left-align\",\"cssCode\":\"\",\"customClass\":[],\"functions\":\"\",\"layoutType\":\"H5\",\"jsonVersion\":3,\"onFormCreated\":\"\",\"onFormMounted\":\"\",\"onFormDataChange\":\"\"}}', '{\"nodes\":[{\"id\":\"Event_be521c6\",\"type\":\"bpmn:startEvent\",\"x\":560,\"y\":250,\"properties\":{},\"zIndex\":1010,\"text\":{\"x\":560,\"y\":290,\"value\":\"开始\"}},{\"id\":\"Activity_af6a39d\",\"type\":\"bpmn:userTask\",\"x\":820,\"y\":290,\"properties\":{},\"zIndex\":1012,\"text\":{\"x\":820,\"y\":290,\"value\":\"审批\"}},{\"id\":\"Event_07c8d97\",\"type\":\"bpmn:endEvent\",\"x\":600,\"y\":480,\"properties\":{},\"zIndex\":1009,\"text\":{\"x\":600,\"y\":520,\"value\":\"结束\"}}],\"edges\":[{\"id\":\"2c5d59e5-9391-41eb-8965-ad7d96ec606e\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Event_be521c6\",\"targetNodeId\":\"Activity_af6a39d\",\"startPoint\":{\"x\":578,\"y\":250},\"endPoint\":{\"x\":770,\"y\":290},\"properties\":{},\"zIndex\":1011,\"pointsList\":[{\"x\":578,\"y\":250},{\"x\":674,\"y\":250},{\"x\":674,\"y\":290},{\"x\":770,\"y\":290}]},{\"id\":\"b43de65b-6489-4e2d-aa2b-d298a430f912\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Activity_af6a39d\",\"targetNodeId\":\"Event_07c8d97\",\"startPoint\":{\"x\":820,\"y\":330},\"endPoint\":{\"x\":618,\"y\":480},\"properties\":{},\"zIndex\":1013,\"pointsList\":[{\"x\":820,\"y\":330},{\"x\":820,\"y\":480},{\"x\":618,\"y\":480}]}]}', NULL, 0, 0, 0);
INSERT INTO `x_flow_template` VALUES (2, '2', 2, '2', '{\"widgetList\":[{\"key\":93225,\"type\":\"textarea\",\"icon\":\"textarea-field\",\"formItemFlag\":true,\"options\":{\"name\":\"textarea53158\",\"label\":\"textarea\",\"labelAlign\":\"\",\"rows\":3,\"defaultValue\":\"\",\"placeholder\":\"\",\"columnWidth\":\"200px\",\"size\":\"\",\"labelWidth\":null,\"labelHidden\":false,\"readonly\":false,\"disabled\":false,\"hidden\":false,\"required\":false,\"requiredHint\":\"\",\"validation\":\"\",\"validationHint\":\"\",\"customClass\":[],\"labelIconClass\":null,\"labelIconPosition\":\"rear\",\"labelTooltip\":null,\"minLength\":null,\"maxLength\":null,\"showWordLimit\":false,\"onCreated\":\"\",\"onMounted\":\"\",\"onInput\":\"\",\"onChange\":\"\",\"onFocus\":\"\",\"onBlur\":\"\",\"onValidate\":\"\"},\"id\":\"textarea53158\"},{\"key\":35486,\"type\":\"select\",\"icon\":\"select-field\",\"formItemFlag\":true,\"options\":{\"name\":\"select101798\",\"label\":\"select\",\"labelAlign\":\"\",\"defaultValue\":\"\",\"placeholder\":\"\",\"columnWidth\":\"200px\",\"size\":\"\",\"labelWidth\":null,\"labelHidden\":false,\"disabled\":false,\"hidden\":false,\"clearable\":true,\"filterable\":false,\"allowCreate\":false,\"remote\":false,\"automaticDropdown\":false,\"multiple\":false,\"multipleLimit\":0,\"optionItems\":[{\"label\":\"select 1\",\"value\":1},{\"label\":\"select 2\",\"value\":2},{\"label\":\"select 3\",\"value\":3}],\"required\":false,\"requiredHint\":\"\",\"validation\":\"\",\"validationHint\":\"\",\"customClass\":\"\",\"labelIconClass\":null,\"labelIconPosition\":\"rear\",\"labelTooltip\":null,\"onCreated\":\"\",\"onMounted\":\"\",\"onRemoteQuery\":\"\",\"onChange\":\"\",\"onFocus\":\"\",\"onBlur\":\"\",\"onValidate\":\"\"},\"id\":\"select101798\"},{\"key\":85807,\"type\":\"input\",\"icon\":\"text-field\",\"formItemFlag\":true,\"options\":{\"name\":\"input47807\",\"label\":\"input\",\"labelAlign\":\"\",\"type\":\"text\",\"defaultValue\":\"\",\"placeholder\":\"\",\"columnWidth\":\"200px\",\"size\":\"\",\"labelWidth\":null,\"labelHidden\":false,\"readonly\":false,\"disabled\":false,\"hidden\":false,\"clearable\":true,\"showPassword\":false,\"required\":false,\"requiredHint\":\"\",\"validation\":\"\",\"validationHint\":\"\",\"customClass\":\"\",\"labelIconClass\":null,\"labelIconPosition\":\"rear\",\"labelTooltip\":null,\"minLength\":null,\"maxLength\":null,\"showWordLimit\":false,\"prefixIcon\":\"\",\"suffixIcon\":\"\",\"appendButton\":false,\"appendButtonDisabled\":false,\"buttonIcon\":\"custom-search\",\"onCreated\":\"\",\"onMounted\":\"\",\"onInput\":\"\",\"onChange\":\"\",\"onFocus\":\"\",\"onBlur\":\"\",\"onValidate\":\"\",\"onAppendButtonClick\":\"\"},\"id\":\"input47807\"}],\"formConfig\":{\"modelName\":\"formData\",\"refName\":\"vForm\",\"rulesName\":\"rules\",\"labelWidth\":80,\"labelPosition\":\"left\",\"size\":\"\",\"labelAlign\":\"label-left-align\",\"cssCode\":\"\",\"customClass\":[],\"functions\":\"\",\"layoutType\":\"H5\",\"jsonVersion\":3,\"onFormCreated\":\"\",\"onFormMounted\":\"\",\"onFormDataChange\":\"\"}}', '{\"nodes\":[{\"id\":\"Event_5026d13\",\"type\":\"bpmn:startEvent\",\"x\":260,\"y\":250,\"properties\":{\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":\"\",\"userId\":\"\",\"deptId\":\"\",\"postId\":\"\"},\"zIndex\":1005,\"text\":{\"x\":260,\"y\":290,\"value\":\"开始\"}},{\"id\":\"Activity_88d7be8\",\"type\":\"bpmn:userTask\",\"x\":810,\"y\":180,\"properties\":{\"fieldAuth\":{\"textarea53158\":2,\"textarea9642\":2,\"select101798\":2,\"input47807\":2},\"userId\":1,\"deptId\":3,\"postId\":2,\"userType\":3,\"gateway\":[]},\"zIndex\":1011,\"text\":{\"x\":810,\"y\":180,\"value\":\"审批1\"}},{\"id\":\"Event_af85d64\",\"type\":\"bpmn:endEvent\",\"x\":1190,\"y\":360,\"properties\":{\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":\"\",\"userId\":\"\",\"deptId\":\"\",\"postId\":\"\"},\"zIndex\":1015,\"text\":{\"x\":1190,\"y\":400,\"value\":\"结束\"}},{\"id\":\"Gateway_5c0df6e\",\"type\":\"bpmn:exclusiveGateway\",\"x\":410,\"y\":180,\"properties\":{\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userId\":\"\",\"deptId\":\"\",\"postId\":\"\",\"userType\":\"\",\"gateway\":[{\"id\":\"textarea53158\",\"value\":\"1\",\"condition\":\"==\"}]},\"zIndex\":1005,\"text\":{\"x\":410,\"y\":220,\"value\":\"网关\"}},{\"id\":\"Activity_0c4ec56\",\"type\":\"bpmn:serviceTask\",\"x\":580,\"y\":180,\"properties\":{\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userId\":\"\",\"deptId\":\"\",\"postId\":\"\",\"userType\":\"\",\"gateway\":[]},\"zIndex\":1003,\"text\":{\"x\":580,\"y\":180,\"value\":\"系统任务\"}},{\"id\":\"Gateway_9e9bedf\",\"type\":\"bpmn:exclusiveGateway\",\"x\":410,\"y\":360,\"properties\":{\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":\"\",\"userId\":\"\",\"deptId\":\"\",\"postId\":\"\",\"gateway\":[{\"id\":\"textarea53158\",\"value\":\"2\",\"condition\":\"==\"}]},\"zIndex\":1001,\"text\":{\"x\":410,\"y\":400,\"value\":\"网关\"}},{\"id\":\"Activity_0c8eb76\",\"type\":\"bpmn:userTask\",\"x\":590,\"y\":360,\"properties\":{\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":1,\"userId\":\"\",\"deptId\":2,\"postId\":3,\"gateway\":[]},\"zIndex\":1012,\"text\":{\"x\":590,\"y\":360,\"value\":\"审批\"}},{\"id\":\"Activity_f09936b\",\"type\":\"bpmn:userTask\",\"x\":1030,\"y\":180,\"properties\":{\"userType\":3,\"userId\":1,\"deptId\":\"\",\"postId\":\"\",\"fieldAuth\":{},\"gateway\":[]},\"zIndex\":1001,\"text\":{\"x\":1030,\"y\":180,\"value\":\"审批\"}}],\"edges\":[{\"id\":\"c7d055f8-bbce-40e5-a2b4-5b0fb1ee9aea\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Event_5026d13\",\"targetNodeId\":\"Gateway_5c0df6e\",\"startPoint\":{\"x\":278,\"y\":250},\"endPoint\":{\"x\":385,\"y\":180},\"properties\":{},\"zIndex\":1009,\"pointsList\":[{\"x\":278,\"y\":250},{\"x\":355,\"y\":250},{\"x\":355,\"y\":180},{\"x\":385,\"y\":180}]},{\"id\":\"ad29adbd-060c-4e0e-9c54-a5512f2542b0\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Gateway_5c0df6e\",\"targetNodeId\":\"Activity_0c4ec56\",\"startPoint\":{\"x\":435,\"y\":180},\"endPoint\":{\"x\":530,\"y\":180},\"properties\":{},\"zIndex\":1016,\"pointsList\":[{\"x\":435,\"y\":180},{\"x\":530,\"y\":180}]},{\"id\":\"2146dbe0-9e6d-4c20-ab84-6f66d3f5f4ea\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Event_5026d13\",\"targetNodeId\":\"Gateway_9e9bedf\",\"startPoint\":{\"x\":278,\"y\":250},\"endPoint\":{\"x\":385,\"y\":360},\"properties\":{},\"zIndex\":1008,\"pointsList\":[{\"x\":278,\"y\":250},{\"x\":355,\"y\":250},{\"x\":355,\"y\":360},{\"x\":385,\"y\":360}]},{\"id\":\"7bfc16b1-a40b-4a53-8599-498a73dede74\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Gateway_9e9bedf\",\"targetNodeId\":\"Activity_0c8eb76\",\"startPoint\":{\"x\":435,\"y\":360},\"endPoint\":{\"x\":540,\"y\":360},\"properties\":{},\"zIndex\":1012,\"pointsList\":[{\"x\":435,\"y\":360},{\"x\":540,\"y\":360}]},{\"id\":\"3f872b96-1abf-4051-9c12-b6fe4e159b00\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Activity_0c8eb76\",\"targetNodeId\":\"Event_af85d64\",\"startPoint\":{\"x\":640,\"y\":360},\"endPoint\":{\"x\":1172,\"y\":360},\"properties\":{},\"zIndex\":1027,\"pointsList\":[{\"x\":640,\"y\":360},{\"x\":1172,\"y\":360}]},{\"id\":\"16833e07-5735-48dd-83f2-e14c04b8fe9b\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Activity_0c4ec56\",\"targetNodeId\":\"Activity_88d7be8\",\"startPoint\":{\"x\":630,\"y\":180},\"endPoint\":{\"x\":760,\"y\":180},\"properties\":{},\"zIndex\":1004,\"pointsList\":[{\"x\":630,\"y\":180},{\"x\":760,\"y\":180}]},{\"id\":\"15929f6f-cf6e-465c-848c-c31fdd4e89d2\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Activity_88d7be8\",\"targetNodeId\":\"Activity_f09936b\",\"startPoint\":{\"x\":860,\"y\":180},\"endPoint\":{\"x\":980,\"y\":180},\"properties\":{},\"zIndex\":1009,\"pointsList\":[{\"x\":860,\"y\":180},{\"x\":980,\"y\":180}]},{\"id\":\"0ba7f6ae-1101-4c7d-ac4f-b01f755f518f\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Activity_f09936b\",\"targetNodeId\":\"Event_af85d64\",\"startPoint\":{\"x\":1080,\"y\":180},\"endPoint\":{\"x\":1190,\"y\":342},\"properties\":{},\"zIndex\":1002,\"pointsList\":[{\"x\":1080,\"y\":180},{\"x\":1190,\"y\":180},{\"x\":1190,\"y\":342}]}]}', '[{\"id\":\"Event_5026d13\",\"pid\":0,\"label\":\"开始\",\"type\":\"bpmn:startEvent\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Gateway_5c0df6e\",\"pid\":\"Event_5026d13\",\"label\":\"网关\",\"type\":\"bpmn:exclusiveGateway\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"gateway\":[{\"id\":\"textarea53158\",\"value\":\"1\",\"condition\":\"==\"}],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_0c4ec56\",\"pid\":\"Gateway_5c0df6e\",\"label\":\"系统任务\",\"type\":\"bpmn:serviceTask\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"gateway\":[],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_88d7be8\",\"pid\":\"Activity_0c4ec56\",\"label\":\"审批1\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"textarea53158\":2,\"textarea9642\":2,\"select101798\":2,\"input47807\":2},\"gateway\":[],\"userType\":3,\"userId\":1,\"deptId\":3,\"postId\":2,\"children\":[{\"id\":\"Activity_f09936b\",\"pid\":\"Activity_88d7be8\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{},\"gateway\":[],\"userType\":3,\"userId\":1,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_af85d64\",\"pid\":\"Activity_f09936b\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0}]}]}]}]},{\"id\":\"Gateway_9e9bedf\",\"pid\":\"Event_5026d13\",\"label\":\"网关\",\"type\":\"bpmn:exclusiveGateway\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"gateway\":[{\"id\":\"textarea53158\",\"value\":\"2\",\"condition\":\"==\"}],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_0c8eb76\",\"pid\":\"Gateway_9e9bedf\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"gateway\":[],\"userType\":1,\"userId\":0,\"deptId\":2,\"postId\":3,\"children\":[{\"id\":\"Event_af85d64\",\"pid\":\"Activity_0c8eb76\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0}]}]}]},{\"id\":\"Gateway_5c0df6e\",\"pid\":\"Event_5026d13\",\"label\":\"网关\",\"type\":\"bpmn:exclusiveGateway\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"gateway\":[{\"id\":\"textarea53158\",\"value\":\"1\",\"condition\":\"==\"}],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_0c4ec56\",\"pid\":\"Gateway_5c0df6e\",\"label\":\"系统任务\",\"type\":\"bpmn:serviceTask\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"gateway\":[],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_88d7be8\",\"pid\":\"Activity_0c4ec56\",\"label\":\"审批1\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"textarea53158\":2,\"textarea9642\":2,\"select101798\":2,\"input47807\":2},\"gateway\":[],\"userType\":3,\"userId\":1,\"deptId\":3,\"postId\":2,\"children\":[{\"id\":\"Activity_f09936b\",\"pid\":\"Activity_88d7be8\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{},\"gateway\":[],\"userType\":3,\"userId\":1,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_af85d64\",\"pid\":\"Activity_f09936b\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0}]}]}]}]},{\"id\":\"Activity_0c4ec56\",\"pid\":\"Gateway_5c0df6e\",\"label\":\"系统任务\",\"type\":\"bpmn:serviceTask\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"gateway\":[],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_88d7be8\",\"pid\":\"Activity_0c4ec56\",\"label\":\"审批1\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"textarea53158\":2,\"textarea9642\":2,\"select101798\":2,\"input47807\":2},\"gateway\":[],\"userType\":3,\"userId\":1,\"deptId\":3,\"postId\":2,\"children\":[{\"id\":\"Activity_f09936b\",\"pid\":\"Activity_88d7be8\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{},\"gateway\":[],\"userType\":3,\"userId\":1,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_af85d64\",\"pid\":\"Activity_f09936b\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0}]}]}]},{\"id\":\"Activity_88d7be8\",\"pid\":\"Activity_0c4ec56\",\"label\":\"审批1\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"textarea53158\":2,\"textarea9642\":2,\"select101798\":2,\"input47807\":2},\"gateway\":[],\"userType\":3,\"userId\":1,\"deptId\":3,\"postId\":2,\"children\":[{\"id\":\"Activity_f09936b\",\"pid\":\"Activity_88d7be8\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{},\"gateway\":[],\"userType\":3,\"userId\":1,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_af85d64\",\"pid\":\"Activity_f09936b\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0}]}]},{\"id\":\"Activity_f09936b\",\"pid\":\"Activity_88d7be8\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{},\"gateway\":[],\"userType\":3,\"userId\":1,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Event_af85d64\",\"pid\":\"Activity_f09936b\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0}]},{\"id\":\"Event_af85d64\",\"pid\":\"Activity_f09936b\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0},{\"id\":\"Gateway_9e9bedf\",\"pid\":\"Event_5026d13\",\"label\":\"网关\",\"type\":\"bpmn:exclusiveGateway\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"gateway\":[{\"id\":\"textarea53158\",\"value\":\"2\",\"condition\":\"==\"}],\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_0c8eb76\",\"pid\":\"Gateway_9e9bedf\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"gateway\":[],\"userType\":1,\"userId\":0,\"deptId\":2,\"postId\":3,\"children\":[{\"id\":\"Event_af85d64\",\"pid\":\"Activity_0c8eb76\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0}]}]},{\"id\":\"Activity_0c8eb76\",\"pid\":\"Gateway_9e9bedf\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"gateway\":[],\"userType\":1,\"userId\":0,\"deptId\":2,\"postId\":3,\"children\":[{\"id\":\"Event_af85d64\",\"pid\":\"Activity_0c8eb76\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0}]},{\"id\":\"Event_af85d64\",\"pid\":\"Activity_0c8eb76\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"fieldAuth\":{\"textarea53158\":1,\"textarea9642\":1,\"select101798\":1,\"input47807\":1},\"userType\":0,\"userId\":0,\"deptId\":0,\"postId\":0}]', 0, 0, 0);
INSERT INTO `x_flow_template` VALUES (3, '3', 2, '3', '{\"widgetList\":[{\"key\":47677,\"type\":\"time-range\",\"icon\":\"time-range-field\",\"formItemFlag\":true,\"options\":{\"name\":\"timerange73058\",\"label\":\"时间\",\"labelAlign\":\"\",\"defaultValue\":null,\"startPlaceholder\":\"\",\"endPlaceholder\":\"\",\"columnWidth\":\"200px\",\"size\":\"\",\"autoFullWidth\":true,\"labelWidth\":null,\"labelHidden\":false,\"readonly\":false,\"disabled\":false,\"hidden\":false,\"clearable\":true,\"editable\":false,\"format\":\"HH:mm:ss\",\"required\":false,\"requiredHint\":\"\",\"validation\":\"\",\"validationHint\":\"\",\"customClass\":[],\"labelIconClass\":null,\"labelIconPosition\":\"rear\",\"labelTooltip\":null,\"onCreated\":\"\",\"onMounted\":\"\",\"onChange\":\"\",\"onFocus\":\"\",\"onBlur\":\"\",\"onValidate\":\"\"},\"id\":\"timerange73058\"},{\"key\":93225,\"type\":\"textarea\",\"icon\":\"textarea-field\",\"formItemFlag\":true,\"options\":{\"name\":\"textarea53158\",\"label\":\"文本\",\"labelAlign\":\"\",\"rows\":3,\"defaultValue\":\"\",\"placeholder\":\"\",\"columnWidth\":\"200px\",\"size\":\"\",\"labelWidth\":null,\"labelHidden\":false,\"readonly\":false,\"disabled\":false,\"hidden\":false,\"required\":false,\"requiredHint\":\"\",\"validation\":\"\",\"validationHint\":\"\",\"customClass\":[],\"labelIconClass\":null,\"labelIconPosition\":\"rear\",\"labelTooltip\":null,\"minLength\":null,\"maxLength\":null,\"showWordLimit\":false,\"onCreated\":\"\",\"onMounted\":\"\",\"onInput\":\"\",\"onChange\":\"\",\"onFocus\":\"\",\"onBlur\":\"\",\"onValidate\":\"\"},\"id\":\"textarea53158\"}],\"formConfig\":{\"modelName\":\"formData\",\"refName\":\"vForm\",\"rulesName\":\"rules\",\"labelWidth\":80,\"labelPosition\":\"left\",\"size\":\"\",\"labelAlign\":\"label-left-align\",\"cssCode\":\"\",\"customClass\":[],\"functions\":\"\",\"layoutType\":\"H5\",\"jsonVersion\":3,\"onFormCreated\":\"\",\"onFormMounted\":\"\",\"onFormDataChange\":\"\"}}', '{\"nodes\":[{\"id\":\"Event_4b6614e\",\"type\":\"bpmn:startEvent\",\"x\":240,\"y\":200,\"properties\":{\"fieldAuth\":{\"timerange73058\":1,\"textarea53158\":1},\"user\":\"\"},\"zIndex\":1041,\"text\":{\"x\":240,\"y\":240,\"value\":\"开始\"}},{\"id\":\"Gateway_9c32347\",\"type\":\"bpmn:exclusiveGateway\",\"x\":470,\"y\":120,\"properties\":{\"fieldAuth\":{\"timerange73058\":1,\"textarea53158\":1},\"user\":\"\"},\"zIndex\":1014,\"text\":{\"x\":470,\"y\":160,\"value\":\"网关1\"}},{\"id\":\"Event_d6eb84a\",\"type\":\"bpmn:endEvent\",\"x\":1050,\"y\":200,\"properties\":{},\"zIndex\":1021,\"text\":{\"x\":1050,\"y\":240,\"value\":\"结束\"}},{\"id\":\"Gateway_bca6386\",\"type\":\"bpmn:exclusiveGateway\",\"x\":470,\"y\":280,\"properties\":{\"fieldAuth\":{\"timerange73058\":1,\"textarea53158\":1},\"user\":\"\"},\"zIndex\":1012,\"text\":{\"x\":470,\"y\":320,\"value\":\"网关2\"}},{\"id\":\"Activity_9afc160\",\"type\":\"bpmn:userTask\",\"x\":890,\"y\":120,\"properties\":{\"fieldAuth\":{\"timerange73058\":1,\"textarea53158\":1},\"user\":\"\",\"userId\":1,\"deptId\":1,\"postId\":2},\"zIndex\":1003,\"text\":{\"x\":890,\"y\":120,\"value\":\"审批\"}},{\"id\":\"Activity_3a97a80\",\"type\":\"bpmn:userTask\",\"x\":890,\"y\":280,\"properties\":{\"fieldAuth\":{\"timerange73058\":1,\"textarea53158\":1},\"user\":\"\",\"userId\":1,\"deptId\":1,\"postId\":2},\"zIndex\":1020,\"text\":{\"x\":890,\"y\":280,\"value\":\"审批\"}},{\"id\":\"Activity_bcd84a7\",\"type\":\"bpmn:serviceTask\",\"x\":680,\"y\":280,\"properties\":{\"fieldAuth\":{\"timerange73058\":1,\"textarea53158\":1},\"user\":\"\"},\"zIndex\":1004,\"text\":{\"x\":680,\"y\":280,\"value\":\"系统任务\"}},{\"id\":\"Activity_1b66e34\",\"type\":\"bpmn:serviceTask\",\"x\":680,\"y\":120,\"properties\":{\"fieldAuth\":{\"timerange73058\":1,\"textarea53158\":1},\"user\":\"\"},\"zIndex\":1002,\"text\":{\"x\":680,\"y\":120,\"value\":\"系统任务\"}}],\"edges\":[{\"id\":\"ad328b0f-d643-4405-a572-a46231d85b20\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Event_4b6614e\",\"targetNodeId\":\"Gateway_9c32347\",\"startPoint\":{\"x\":258,\"y\":200},\"endPoint\":{\"x\":445,\"y\":120},\"properties\":{},\"zIndex\":1040,\"pointsList\":[{\"x\":258,\"y\":200},{\"x\":288,\"y\":200},{\"x\":288,\"y\":120},{\"x\":445,\"y\":120}]},{\"id\":\"0921f2e8-c802-4c2d-974d-3f598d6678fc\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Event_4b6614e\",\"targetNodeId\":\"Gateway_bca6386\",\"startPoint\":{\"x\":258,\"y\":200},\"endPoint\":{\"x\":445,\"y\":280},\"properties\":{},\"zIndex\":1042,\"pointsList\":[{\"x\":258,\"y\":200},{\"x\":288,\"y\":200},{\"x\":288,\"y\":280},{\"x\":445,\"y\":280}]},{\"id\":\"079a0882-5226-4606-b3a1-ca862276a298\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Activity_9afc160\",\"targetNodeId\":\"Event_d6eb84a\",\"startPoint\":{\"x\":940,\"y\":120},\"endPoint\":{\"x\":1032,\"y\":200},\"properties\":{},\"zIndex\":1052,\"pointsList\":[{\"x\":940,\"y\":120},{\"x\":1002,\"y\":120},{\"x\":1002,\"y\":200},{\"x\":1032,\"y\":200}]},{\"id\":\"1c2d49d2-2b61-41ce-9d22-444f53b0932f\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Activity_3a97a80\",\"targetNodeId\":\"Event_d6eb84a\",\"startPoint\":{\"x\":940,\"y\":280},\"endPoint\":{\"x\":1032,\"y\":200},\"properties\":{},\"zIndex\":1054,\"pointsList\":[{\"x\":940,\"y\":280},{\"x\":1002,\"y\":280},{\"x\":1002,\"y\":200},{\"x\":1032,\"y\":200}]},{\"id\":\"e224d2f9-5526-4e49-8859-e7f5cc8c4f35\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Gateway_bca6386\",\"targetNodeId\":\"Activity_bcd84a7\",\"startPoint\":{\"x\":495,\"y\":280},\"endPoint\":{\"x\":630,\"y\":280},\"properties\":{},\"zIndex\":1063,\"pointsList\":[{\"x\":495,\"y\":280},{\"x\":630,\"y\":280}]},{\"id\":\"b950c08c-a7d9-4eef-92b7-67c3d29ea08f\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Activity_bcd84a7\",\"targetNodeId\":\"Activity_3a97a80\",\"startPoint\":{\"x\":730,\"y\":280},\"endPoint\":{\"x\":840,\"y\":280},\"properties\":{},\"zIndex\":1065,\"pointsList\":[{\"x\":730,\"y\":280},{\"x\":840,\"y\":280}]},{\"id\":\"c7e981c9-1129-4e59-b741-d69436207741\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Gateway_9c32347\",\"targetNodeId\":\"Activity_1b66e34\",\"startPoint\":{\"x\":495,\"y\":120},\"endPoint\":{\"x\":630,\"y\":120},\"properties\":{},\"zIndex\":1071,\"pointsList\":[{\"x\":495,\"y\":120},{\"x\":630,\"y\":120}]},{\"id\":\"b09f0803-cf02-4ea3-9647-fd5edb0e7b55\",\"type\":\"pro-polyline\",\"sourceNodeId\":\"Activity_1b66e34\",\"targetNodeId\":\"Activity_9afc160\",\"startPoint\":{\"x\":730,\"y\":120},\"endPoint\":{\"x\":840,\"y\":120},\"properties\":{},\"zIndex\":1073,\"pointsList\":[{\"x\":730,\"y\":120},{\"x\":840,\"y\":120}]}]}', '[{\"id\":\"Event_4b6614e\",\"pid\":0,\"label\":\"开始\",\"type\":\"bpmn:startEvent\",\"fieldAuth\":{\"timerange73058\":1,\"textarea53158\":1},\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Gateway_9c32347\",\"pid\":\"Event_4b6614e\",\"label\":\"网关1\",\"type\":\"bpmn:exclusiveGateway\",\"fieldAuth\":{\"timerange73058\":1,\"textarea53158\":1},\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_1b66e34\",\"pid\":\"Gateway_9c32347\",\"label\":\"系统任务\",\"type\":\"bpmn:serviceTask\",\"fieldAuth\":{\"timerange73058\":1,\"textarea53158\":1},\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_9afc160\",\"pid\":\"Activity_1b66e34\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"timerange73058\":1,\"textarea53158\":1},\"userId\":1,\"deptId\":1,\"postId\":2,\"children\":[{\"id\":\"Event_d6eb84a\",\"pid\":\"Activity_9afc160\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userId\":0,\"deptId\":0,\"postId\":0}]}]}]},{\"id\":\"Gateway_bca6386\",\"pid\":\"Event_4b6614e\",\"label\":\"网关2\",\"type\":\"bpmn:exclusiveGateway\",\"fieldAuth\":{\"timerange73058\":1,\"textarea53158\":1},\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_bcd84a7\",\"pid\":\"Gateway_bca6386\",\"label\":\"系统任务\",\"type\":\"bpmn:serviceTask\",\"fieldAuth\":{\"timerange73058\":1,\"textarea53158\":1},\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_3a97a80\",\"pid\":\"Activity_bcd84a7\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"timerange73058\":1,\"textarea53158\":1},\"userId\":1,\"deptId\":1,\"postId\":2,\"children\":[{\"id\":\"Event_d6eb84a\",\"pid\":\"Activity_3a97a80\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userId\":0,\"deptId\":0,\"postId\":0}]}]}]}]},{\"id\":\"Gateway_9c32347\",\"pid\":\"Event_4b6614e\",\"label\":\"网关1\",\"type\":\"bpmn:exclusiveGateway\",\"fieldAuth\":{\"timerange73058\":1,\"textarea53158\":1},\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_1b66e34\",\"pid\":\"Gateway_9c32347\",\"label\":\"系统任务\",\"type\":\"bpmn:serviceTask\",\"fieldAuth\":{\"timerange73058\":1,\"textarea53158\":1},\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_9afc160\",\"pid\":\"Activity_1b66e34\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"timerange73058\":1,\"textarea53158\":1},\"userId\":1,\"deptId\":1,\"postId\":2,\"children\":[{\"id\":\"Event_d6eb84a\",\"pid\":\"Activity_9afc160\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userId\":0,\"deptId\":0,\"postId\":0}]}]}]},{\"id\":\"Activity_1b66e34\",\"pid\":\"Gateway_9c32347\",\"label\":\"系统任务\",\"type\":\"bpmn:serviceTask\",\"fieldAuth\":{\"timerange73058\":1,\"textarea53158\":1},\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_9afc160\",\"pid\":\"Activity_1b66e34\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"timerange73058\":1,\"textarea53158\":1},\"userId\":1,\"deptId\":1,\"postId\":2,\"children\":[{\"id\":\"Event_d6eb84a\",\"pid\":\"Activity_9afc160\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userId\":0,\"deptId\":0,\"postId\":0}]}]},{\"id\":\"Activity_9afc160\",\"pid\":\"Activity_1b66e34\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"timerange73058\":1,\"textarea53158\":1},\"userId\":1,\"deptId\":1,\"postId\":2,\"children\":[{\"id\":\"Event_d6eb84a\",\"pid\":\"Activity_9afc160\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userId\":0,\"deptId\":0,\"postId\":0}]},{\"id\":\"Event_d6eb84a\",\"pid\":\"Activity_9afc160\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userId\":0,\"deptId\":0,\"postId\":0},{\"id\":\"Gateway_bca6386\",\"pid\":\"Event_4b6614e\",\"label\":\"网关2\",\"type\":\"bpmn:exclusiveGateway\",\"fieldAuth\":{\"timerange73058\":1,\"textarea53158\":1},\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_bcd84a7\",\"pid\":\"Gateway_bca6386\",\"label\":\"系统任务\",\"type\":\"bpmn:serviceTask\",\"fieldAuth\":{\"timerange73058\":1,\"textarea53158\":1},\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_3a97a80\",\"pid\":\"Activity_bcd84a7\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"timerange73058\":1,\"textarea53158\":1},\"userId\":1,\"deptId\":1,\"postId\":2,\"children\":[{\"id\":\"Event_d6eb84a\",\"pid\":\"Activity_3a97a80\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userId\":0,\"deptId\":0,\"postId\":0}]}]}]},{\"id\":\"Activity_bcd84a7\",\"pid\":\"Gateway_bca6386\",\"label\":\"系统任务\",\"type\":\"bpmn:serviceTask\",\"fieldAuth\":{\"timerange73058\":1,\"textarea53158\":1},\"userId\":0,\"deptId\":0,\"postId\":0,\"children\":[{\"id\":\"Activity_3a97a80\",\"pid\":\"Activity_bcd84a7\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"timerange73058\":1,\"textarea53158\":1},\"userId\":1,\"deptId\":1,\"postId\":2,\"children\":[{\"id\":\"Event_d6eb84a\",\"pid\":\"Activity_3a97a80\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userId\":0,\"deptId\":0,\"postId\":0}]}]},{\"id\":\"Activity_3a97a80\",\"pid\":\"Activity_bcd84a7\",\"label\":\"审批\",\"type\":\"bpmn:userTask\",\"fieldAuth\":{\"timerange73058\":1,\"textarea53158\":1},\"userId\":1,\"deptId\":1,\"postId\":2,\"children\":[{\"id\":\"Event_d6eb84a\",\"pid\":\"Activity_3a97a80\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userId\":0,\"deptId\":0,\"postId\":0}]},{\"id\":\"Event_d6eb84a\",\"pid\":\"Activity_3a97a80\",\"label\":\"结束\",\"type\":\"bpmn:endEvent\",\"userId\":0,\"deptId\":0,\"postId\":0}]', 0, 0, 0);

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
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 17 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '代码生成业务表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_gen_table
-- ----------------------------
INSERT INTO `x_gen_table` VALUES (11, 'x_album', '相册管理表', '', '', '', 'album', 'album', '相册管理', '', '', '', 'crud', '', 1701241021, 1701241021);
INSERT INTO `x_gen_table` VALUES (14, 'x_flow_template', '流程模板', '', '', '', 'flowTemplate', 'flow_template', '流程模板', '', '', '', 'crud', '', 1702106174, 1702106174);
INSERT INTO `x_gen_table` VALUES (15, 'x_flow_apply', '申请流程', '', '', '', 'flowApply', 'flow_apply', '申请流程', '', '', '', 'crud', '', 1702206515, 1702207125);
INSERT INTO `x_gen_table` VALUES (16, 'x_flow_history', '流程历史', '', '', '', 'flowHistory', 'flow_history', '流程历史', '', '', '', 'crud', '', 1702207304, 1702207304);

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
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 149 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '代码生成字段表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_gen_table_column
-- ----------------------------
INSERT INTO `x_gen_table_column` VALUES (84, 10, 'id', '主键', '10', 'int', 'int', 'id', 1, 1, 0, 0, 0, 0, 0, '=', 'input', '', 0, 0, 1701240746);
INSERT INTO `x_gen_table_column` VALUES (85, 10, 'user_id', '用户ID', '10', 'int', 'int', 'user_id', 0, 0, 1, 1, 1, 1, 1, '=', 'select', 'a', 0, 0, 1701240746);
INSERT INTO `x_gen_table_column` VALUES (86, 10, 'article_id', '文章ID', '10', 'int', 'int', 'article_id', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', 0, 0, 1701240746);
INSERT INTO `x_gen_table_column` VALUES (87, 10, 'is_delete', '是否删除', '1', 'tinyint', 'int', 'is_delete', 0, 0, 0, 0, 0, 0, 0, '=', 'input', '', 0, 0, 1701240746);
INSERT INTO `x_gen_table_column` VALUES (88, 10, 'create_time', '创建时间', '10', 'int', 'core.TsTime', 'create_time', 0, 0, 0, 0, 0, 1, 0, '=', 'datetime', '', 0, 0, 1701240746);
INSERT INTO `x_gen_table_column` VALUES (89, 10, 'update_time', '更新时间', '10', 'int', 'core.TsTime', 'update_time', 0, 0, 0, 0, 0, 1, 0, '=', 'datetime', '', 0, 0, 1701240746);
INSERT INTO `x_gen_table_column` VALUES (90, 10, 'delete_time', '是否删除', '10', 'int', 'core.TsTime', 'delete_time', 0, 0, 0, 0, 0, 0, 0, '=', 'datetime', '', 0, 0, 1701240746);
INSERT INTO `x_gen_table_column` VALUES (91, 11, 'id', '主键ID', '10', 'int', 'int', 'id', 1, 1, 1, 0, 1, 0, 0, '=', 'input', '', 1, 1701241021, 1701241021);
INSERT INTO `x_gen_table_column` VALUES (92, 11, 'cid', '类目ID', '10', 'int', 'int', 'cid', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', 2, 1701241021, 1701241021);
INSERT INTO `x_gen_table_column` VALUES (93, 11, 'aid', '管理员ID', '10', 'int', 'int', 'aid', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', 3, 1701241021, 1701241021);
INSERT INTO `x_gen_table_column` VALUES (94, 11, 'uid', '用户ID', '10', 'int', 'int', 'uid', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', 4, 1701241021, 1701241021);
INSERT INTO `x_gen_table_column` VALUES (95, 11, 'type', '文件类型: [10=图片, 20=视频]', '2', 'tinyint', 'int', 'type', 0, 0, 1, 1, 1, 1, 1, '=', 'select', '', 5, 1701241021, 1701241021);
INSERT INTO `x_gen_table_column` VALUES (96, 11, 'name', '文件名称', '100', 'varchar', 'string', 'name', 0, 0, 1, 1, 1, 1, 1, 'LIKE', 'input', '', 6, 1701241021, 1701241021);
INSERT INTO `x_gen_table_column` VALUES (97, 11, 'uri', '文件路径', '200', 'varchar', 'string', 'uri', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', 7, 1701241021, 1701241021);
INSERT INTO `x_gen_table_column` VALUES (98, 11, 'ext', '文件扩展', '10', 'varchar', 'string', 'ext', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', 8, 1701241021, 1701241021);
INSERT INTO `x_gen_table_column` VALUES (99, 11, 'size', '文件大小', '10', 'int', 'int', 'size', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', 9, 1701241021, 1701241021);
INSERT INTO `x_gen_table_column` VALUES (100, 11, 'is_delete', '是否删除: 0=否, 1=是', '1', 'int', 'int', 'is_delete', 0, 0, 0, 0, 0, 0, 0, '=', 'input', '', 10, 1701241021, 1701241021);
INSERT INTO `x_gen_table_column` VALUES (101, 11, 'create_time', '创建时间', '10', 'int', 'core.TsTime', 'create_time', 0, 0, 0, 0, 0, 1, 0, '=', 'datetime', '', 11, 1701241021, 1701241021);
INSERT INTO `x_gen_table_column` VALUES (102, 11, 'update_time', '更新时间', '10', 'int', 'core.TsTime', 'update_time', 0, 0, 0, 0, 0, 1, 0, '=', 'datetime', '', 12, 1701241021, 1701241021);
INSERT INTO `x_gen_table_column` VALUES (103, 11, 'delete_time', '删除时间', '10', 'int', 'core.TsTime', 'delete_time', 0, 0, 0, 0, 0, 0, 0, '=', 'datetime', '', 13, 1701241021, 1701241021);
INSERT INTO `x_gen_table_column` VALUES (116, 14, 'id', '', '11', 'int', 'int', 'id', 1, 1, 1, 0, 1, 0, 0, '=', 'input', '', 1, 1702106174, 1702106174);
INSERT INTO `x_gen_table_column` VALUES (117, 14, 'flow_name', '流程名称', '255', 'varchar', 'string', 'flow_name', 0, 0, 1, 1, 1, 1, 1, 'LIKE', 'input', '', 2, 1702106174, 1702106174);
INSERT INTO `x_gen_table_column` VALUES (118, 14, 'flow_group', '流程分类', '2', 'tinyint', 'int', 'flow_group', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', 3, 1702106174, 1702106174);
INSERT INTO `x_gen_table_column` VALUES (119, 14, 'flow_remark', '流程描述', '255', 'varchar', 'string', 'flow_remark', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', 4, 1702106174, 1702106174);
INSERT INTO `x_gen_table_column` VALUES (120, 14, 'flow_form_data', '表单配置', '0', 'longtext', 'string', 'flow_form_data', 0, 0, 1, 1, 1, 1, 1, '=', 'textarea', '', 5, 1702106174, 1702106174);
INSERT INTO `x_gen_table_column` VALUES (121, 14, 'flow_process_data', '流程配置', '0', 'longtext', 'string', 'flow_process_data', 0, 0, 1, 1, 1, 1, 1, '=', 'textarea', '', 6, 1702106174, 1702106174);
INSERT INTO `x_gen_table_column` VALUES (122, 15, 'id', '', '10', 'int', 'int', 'id', 1, 0, 1, 0, 1, 0, 0, '=', 'input', '', 0, 0, 1702207125);
INSERT INTO `x_gen_table_column` VALUES (123, 15, 'template_id', '模板', '10', 'int', 'int', 'template_id', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', 0, 0, 1702207125);
INSERT INTO `x_gen_table_column` VALUES (124, 15, 'apply_user_id', '申请人id', '10', 'int', 'int', 'apply_user_id', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', 0, 0, 1702207125);
INSERT INTO `x_gen_table_column` VALUES (125, 15, 'apply_user_nickname', '申请人昵称', '32', 'varchar', 'string', 'apply_user_nickname', 0, 0, 1, 1, 1, 1, 1, 'LIKE', 'input', '', 0, 0, 1702207125);
INSERT INTO `x_gen_table_column` VALUES (126, 15, 'flow_name', '流程名称', '255', 'varchar', 'string', 'flow_name', 0, 0, 1, 1, 1, 1, 1, 'LIKE', 'input', '', 0, 0, 1702207125);
INSERT INTO `x_gen_table_column` VALUES (127, 15, 'flow_group', '流程分类', '2', 'tinyint', 'int', 'flow_group', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', 0, 0, 1702207125);
INSERT INTO `x_gen_table_column` VALUES (128, 15, 'flow_remark', '流程描述', '255', 'varchar', 'string', 'flow_remark', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', 0, 0, 1702207125);
INSERT INTO `x_gen_table_column` VALUES (129, 15, 'flow_form_data', '表单配置', '0', 'longtext', 'string', 'flow_form_data', 0, 0, 1, 1, 1, 1, 1, '=', 'textarea', '', 0, 0, 1702207125);
INSERT INTO `x_gen_table_column` VALUES (130, 15, 'flow_process_data', '流程配置', '0', 'longtext', 'string', 'flow_process_data', 0, 0, 1, 1, 1, 1, 1, '=', 'textarea', '', 0, 0, 1702207125);
INSERT INTO `x_gen_table_column` VALUES (131, 15, 'status', '状态：0待提交，1审批中，2审批完成，3审批失败', '2', 'tinyint', 'int', 'status', 0, 0, 1, 1, 1, 1, 1, '=', 'select', 'flow_apply_status', 0, 0, 1702207125);
INSERT INTO `x_gen_table_column` VALUES (132, 15, 'update_time', '更新时间', '10', 'int', 'core.TsTime', 'update_time', 0, 0, 0, 0, 0, 1, 0, '=', 'datetime', '', 0, 0, 1702207125);
INSERT INTO `x_gen_table_column` VALUES (133, 15, 'create_time', '创建时间', '10', 'int', 'core.TsTime', 'create_time', 0, 0, 0, 0, 0, 1, 0, '=', 'datetime', '', 0, 0, 1702207125);
INSERT INTO `x_gen_table_column` VALUES (134, 15, 'delete_time', '删除时间', '10', 'int', 'core.TsTime', 'delete_time', 0, 0, 0, 0, 0, 0, 0, '=', 'datetime', '', 0, 0, 1702207125);
INSERT INTO `x_gen_table_column` VALUES (135, 16, 'id', '历史id', '10', 'int', 'int', 'id', 1, 0, 1, 0, 1, 0, 0, '=', 'input', '', 1, 1702207304, 1702207304);
INSERT INTO `x_gen_table_column` VALUES (136, 16, 'apply_id', '申请id', '10', 'int', 'int', 'apply_id', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', 2, 1702207304, 1702207304);
INSERT INTO `x_gen_table_column` VALUES (137, 16, 'template_id', '模板id', '10', 'int', 'int', 'template_id', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', 3, 1702207304, 1702207304);
INSERT INTO `x_gen_table_column` VALUES (138, 16, 'apply_user_id', '申请人id', '10', 'int', 'int', 'apply_user_id', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', 4, 1702207304, 1702207304);
INSERT INTO `x_gen_table_column` VALUES (139, 16, 'apply_user_nickname', '申请人昵称', '32', 'varchar', 'string', 'apply_user_nickname', 0, 0, 1, 1, 1, 1, 1, 'LIKE', 'input', '', 5, 1702207304, 1702207304);
INSERT INTO `x_gen_table_column` VALUES (140, 16, 'approver_id', '审批人id', '10', 'int', 'int', 'approver_id', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', 6, 1702207304, 1702207304);
INSERT INTO `x_gen_table_column` VALUES (141, 16, 'approver_nickname', '审批用户昵称', '32', 'varchar', 'string', 'approver_nickname', 0, 0, 1, 1, 1, 1, 1, 'LIKE', 'input', '', 7, 1702207304, 1702207304);
INSERT INTO `x_gen_table_column` VALUES (142, 16, 'node_id', '节点', '50', 'varchar', 'string', 'node_id', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', 8, 1702207304, 1702207304);
INSERT INTO `x_gen_table_column` VALUES (143, 16, 'form_value', '表单值', '0', 'mediumtext', 'string', 'form_value', 0, 0, 1, 1, 1, 1, 1, '=', 'textarea', '', 9, 1702207304, 1702207304);
INSERT INTO `x_gen_table_column` VALUES (144, 16, 'pass_status', '通过状态：0待处理，1通过，2拒绝', '1', 'tinyint', 'int', 'pass_status', 0, 0, 1, 1, 1, 1, 1, '=', 'radio', '', 10, 1702207304, 1702207304);
INSERT INTO `x_gen_table_column` VALUES (145, 16, 'pass_remark', '通过备注', '200', 'varchar', 'string', 'pass_remark', 0, 0, 1, 1, 1, 1, 1, '=', 'input', '', 11, 1702207304, 1702207304);
INSERT INTO `x_gen_table_column` VALUES (146, 16, 'update_time', '更新时间', '10', 'int', 'core.TsTime', 'update_time', 0, 0, 0, 0, 0, 1, 0, '=', 'datetime', '', 12, 1702207304, 1702207304);
INSERT INTO `x_gen_table_column` VALUES (147, 16, 'create_time', '创建时间', '10', 'int', 'core.TsTime', 'create_time', 0, 0, 0, 0, 0, 1, 0, '=', 'datetime', '', 13, 1702207304, 1702207304);
INSERT INTO `x_gen_table_column` VALUES (148, 16, 'delete_time', '删除时间', '10', 'int', 'core.TsTime', 'delete_time', 0, 0, 0, 0, 0, 0, 0, '=', 'datetime', '', 14, 1702207304, 1702207304);

-- ----------------------------
-- Table structure for x_hot_search
-- ----------------------------
DROP TABLE IF EXISTS `x_hot_search`;
CREATE TABLE `x_hot_search`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '关键词',
  `sort` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序号',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '热门搜索配置表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_hot_search
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
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  `delete_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '消息通知设置表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_notice_setting
-- ----------------------------
INSERT INTO `x_notice_setting` VALUES (1, 101, '登录验证码', '用户手机号码登录时发送', 1, 2, '{}', '{\"type\":\"sms\",\"templateId\":\"SMS_222458159\",\"content\":\"您正在登录，验证码${code}，切勿将验证码泄露于他人，本条验证码有效期5分钟。\",\"tips\":[\"可选变量 验证码:code\",\"示例：您正在登录，验证码${code}，切勿将验证码泄露于他人，本条验证码有效期5分钟。\",\"生效条件：1、管理后台完成短信设置。2、第三方短信平台申请模板。\"],\"status\":\"1\"}', '{}', '{}', 0, 1648696695, 1648696695, 0);
INSERT INTO `x_notice_setting` VALUES (2, 102, '绑定手机验证码', '用户绑定手机号码时发送', 1, 2, '{}', '{\"type\":\"sms\",\"templateId\":\"SMS_175615069\",\"content\":\"您正在绑定手机号，验证码${code}，切勿将验证码泄露于他人，本条验证码有效期5分钟。\",\"tips\":[\"可选变量 验证码:code\",\"示例：您正在绑定手机号，验证码${code}，切勿将验证码泄露于他人，本条验证码有效期5分钟。\",\"生效条件：1、管理后台完成短信设置。2、第三方短信平台申请模板。\"],\"status\":\"1\"}', '{}', '{}', 0, 1648696695, 1648696695, 0);
INSERT INTO `x_notice_setting` VALUES (3, 103, '变更手机验证码', '用户变更手机号码时发送', 1, 2, '{}', '{\"type\":\"sms\",\"templateId\":\"SMS_207952628\",\"content\":\"您正在变更手机号，验证码${code}，切勿将验证码泄露于他人，本条验证码有效期5分钟。\",\"tips\":[\"可选变量 验证码:code\",\"示例：您正在变更手机号，验证码${code}，切勿将验证码泄露于他人，本条验证码有效期5分钟。\",\"生效条件：1、管理后台完成短信设置。2、第三方短信平台申请模板。\"],\"status\":\"1\"}', '{}', '{}', 0, 1648696695, 1648696695, 0);
INSERT INTO `x_notice_setting` VALUES (4, 104, '找回登录密码验证码', '用户找回登录密码号码时发送', 1, 2, '{}', '{\"type\":\"sms\",\"templateId\":\"SMS_175615069\",\"content\":\"您正在找回登录密码，验证码${code}，切勿将验证码泄露于他人，本条验证码有效期5分钟。\",\"tips\":[\"可选变量 验证码:code\",\"示例：您正在找回登录密码，验证码${code}，切勿将验证码泄露于他人，本条验证码有效期5分钟。\",\"条验证码有效期5分钟。\"],\"status\":\"1\"}', '{}', '{}', 0, 1648696695, 1648696695, 0);

-- ----------------------------
-- Table structure for x_oauth2_
-- ----------------------------
DROP TABLE IF EXISTS `x_oauth2_`;
CREATE TABLE `x_oauth2_`  (
  `id` int(11) NOT NULL COMMENT 'id',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `date` datetime NULL DEFAULT NULL,
  `del_flag` tinyint(1) NULL DEFAULT NULL,
  `json` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of x_oauth2_
-- ----------------------------

-- ----------------------------
-- Table structure for x_official_reply
-- ----------------------------
DROP TABLE IF EXISTS `x_official_reply`;
CREATE TABLE `x_official_reply`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '规则名',
  `keyword` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '关键词',
  `reply_type` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '回复类型: [1=关注回复 2=关键字回复, 3=默认回复]',
  `matching_type` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '匹配方式: [1=全匹配, 2=模糊匹配]',
  `content_type` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '内容类型: [1=文本]',
  `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '启动状态: [1=启动, 0=关闭]',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '回复内容',
  `sort` int(11) UNSIGNED NOT NULL DEFAULT 50 COMMENT '排序编号',
  `is_delete` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除',
  `create_time` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '更新时间',
  `delete_time` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '公众号的回复表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_official_reply
-- ----------------------------

-- ----------------------------
-- Table structure for x_product
-- ----------------------------
DROP TABLE IF EXISTS `x_product`;
CREATE TABLE `x_product`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '0' COMMENT '产品名称',
  `category` int(11) UNSIGNED NOT NULL COMMENT '产品分类',
  `pics` mediumtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  PRIMARY KEY (`id`) USING BTREE,
  FULLTEXT INDEX `name`(`name`)
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '产品表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_product
-- ----------------------------
INSERT INTO `x_product` VALUES (1, 'aaa', 2, '11');
INSERT INTO `x_product` VALUES (2, 'a', 0, '3');

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
  `last_login_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '最后登录',
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  `delete_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统管理成员表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_system_auth_admin
-- ----------------------------
INSERT INTO `x_system_auth_admin` VALUES (1, 1, 3, 'admin', 'admin', '7ca7e19452aa2366068785be5c7ded35', '/image/20232911/27699a1e991e4ba1ba13ee6d45185db6.jpg', '0', 'HOEp0', 1, 1, 0, 0, '127.0.0.1', 1703773877, 1642321599, 1703773877, 0);
INSERT INTO `x_system_auth_admin` VALUES (2, 2, 3, 'zhihuibu01', '指挥部01', 'ea7e7f97957b7cdd2b245abc31cebaa4', '/image/20232911/27699a1e991e4ba1ba13ee6d45185db6.jpg', '1', 'JUD5n', 1, 1, 0, 0, '', 0, 1702883992, 1703599817, 0);

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
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  `delete_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统部门管理表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_system_auth_dept
-- ----------------------------
INSERT INTO `x_system_auth_dept` VALUES (1, 0, '默认部门', 1, 'admin', '18327647788', 10, 0, 0, 1649841995, 1703523269, 0);
INSERT INTO `x_system_auth_dept` VALUES (2, 1, '指挥部', 2, '指挥部01', '17608390654', 0, 0, 0, 1702212515, 1703501462, 0);
INSERT INTO `x_system_auth_dept` VALUES (3, 2, '指挥部子级', 0, '', '', 0, 0, 0, 1703500730, 1703501512, 0);

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
  `perms` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '权限标识',
  `paths` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '路由地址',
  `component` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '前端组件',
  `selected` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '选中路径',
  `params` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '路由参数',
  `is_cache` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否缓存: 0=否, 1=是',
  `is_show` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '是否显示: 0=否, 1=是',
  `is_disable` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否禁用: 0=否, 1=是',
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 783 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统菜单管理表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_system_auth_menu
-- ----------------------------
INSERT INTO `x_system_auth_menu` VALUES (1, 0, 'C', '工作台', 'el-icon-Monitor', 50, 'common:index:console', 'workbench', 'workbench/index', '', '', 1, 1, 0, 1650341765, 1668672757);
INSERT INTO `x_system_auth_menu` VALUES (100, 0, 'M', '权限管理', 'el-icon-Lock', 44, '', 'permission', '', '', '', 0, 1, 0, 1650341765, 1662626201);
INSERT INTO `x_system_auth_menu` VALUES (101, 100, 'C', '管理员', 'local-icon-wode', 0, 'system:admin:list', 'admin', 'permission/admin/index', '', '', 1, 1, 0, 1650341765, 1663301404);
INSERT INTO `x_system_auth_menu` VALUES (102, 101, 'A', '管理员详情', '', 0, 'system:admin:detail', '', '', '', '', 0, 1, 0, 1650341765, 1660201785);
INSERT INTO `x_system_auth_menu` VALUES (103, 101, 'A', '管理员新增', '', 0, 'system:admin:add', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `x_system_auth_menu` VALUES (104, 101, 'A', '管理员编辑', '', 0, 'system:admin:edit', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `x_system_auth_menu` VALUES (105, 101, 'A', '管理员删除', '', 0, 'system:admin:del', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `x_system_auth_menu` VALUES (106, 101, 'A', '管理员状态', '', 0, 'system:admin:disable', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `x_system_auth_menu` VALUES (110, 100, 'C', '角色管理', 'el-icon-Female', 0, 'system:role:list', 'role', 'permission/role/index', '', '', 1, 1, 0, 1650341765, 1663301451);
INSERT INTO `x_system_auth_menu` VALUES (111, 110, 'A', '角色详情', '', 0, 'system:role:detail', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `x_system_auth_menu` VALUES (112, 110, 'A', '角色新增', '', 0, 'system:role:add', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `x_system_auth_menu` VALUES (113, 110, 'A', '角色编辑', '', 0, 'system:role:edit', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `x_system_auth_menu` VALUES (114, 110, 'A', '角色删除', '', 0, 'system:role:del', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `x_system_auth_menu` VALUES (120, 100, 'C', '菜单管理', 'el-icon-Operation', 0, 'system:menu:list', 'menu', 'permission/menu/index', '', '', 1, 1, 0, 1650341765, 1663301388);
INSERT INTO `x_system_auth_menu` VALUES (121, 120, 'A', '菜单详情', '', 0, 'system:menu:detail', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `x_system_auth_menu` VALUES (122, 120, 'A', '菜单新增', '', 0, 'system:menu:add', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `x_system_auth_menu` VALUES (123, 120, 'A', '菜单编辑', '', 0, 'system:menu:edit', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `x_system_auth_menu` VALUES (124, 120, 'A', '菜单删除', '', 0, 'system:menu:del', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `x_system_auth_menu` VALUES (130, 0, 'M', '组织管理', 'el-icon-OfficeBuilding', 45, '', 'organization', '', '', '', 0, 1, 0, 1650341765, 1664416715);
INSERT INTO `x_system_auth_menu` VALUES (131, 130, 'C', '部门管理', 'el-icon-Coordinate', 0, 'system:dept:list', 'department', 'organization/department/index', '', '', 1, 1, 0, 1650341765, 1660201994);
INSERT INTO `x_system_auth_menu` VALUES (132, 131, 'A', '部门详情', '', 0, 'system:dept:detail', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `x_system_auth_menu` VALUES (133, 131, 'A', '部门新增', '', 0, 'system:dept:add', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `x_system_auth_menu` VALUES (134, 131, 'A', '部门编辑', '', 0, 'system:dept:edit', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `x_system_auth_menu` VALUES (135, 131, 'A', '部门删除', '', 0, 'system:dept:del', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `x_system_auth_menu` VALUES (140, 130, 'C', '岗位管理', 'el-icon-PriceTag', 0, 'system:post:list', 'post', 'organization/post/index', '', '', 1, 1, 0, 1650341765, 1660202057);
INSERT INTO `x_system_auth_menu` VALUES (141, 140, 'A', '岗位详情', '', 0, 'system:post:detail', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `x_system_auth_menu` VALUES (142, 140, 'A', '岗位新增', '', 0, 'system:post:add', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `x_system_auth_menu` VALUES (143, 140, 'A', '岗位编辑', '', 0, 'system:post:edit', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `x_system_auth_menu` VALUES (144, 140, 'A', '岗位删除', '', 0, 'system:post:del', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `x_system_auth_menu` VALUES (200, 0, 'M', '其它管理', '', 0, '', '', '', '', '', 0, 0, 0, 1650341765, 1660636870);
INSERT INTO `x_system_auth_menu` VALUES (201, 200, 'M', '图库管理', '', 0, '', '', '', '', '', 0, 0, 0, 1650341765, 1650341765);
INSERT INTO `x_system_auth_menu` VALUES (202, 201, 'A', '文件列表', '', 0, 'albums:albumList', '', '', '', '', 0, 0, 0, 1650341765, 1650341765);
INSERT INTO `x_system_auth_menu` VALUES (203, 201, 'A', '文件命名', '', 0, 'albums:albumRename', '', '', '', '', 0, 0, 0, 1650341765, 1650341765);
INSERT INTO `x_system_auth_menu` VALUES (204, 201, 'A', '文件移动', '', 0, 'albums:albumMove', '', '', '', '', 0, 0, 0, 1650341765, 1650341765);
INSERT INTO `x_system_auth_menu` VALUES (205, 201, 'A', '文件删除', '', 0, 'albums:albumDel', '', '', '', '', 0, 0, 0, 1650341765, 1650341765);
INSERT INTO `x_system_auth_menu` VALUES (206, 201, 'A', '分类列表', '', 0, 'albums:cateList', '', '', '', '', 0, 0, 0, 1650341765, 1650341765);
INSERT INTO `x_system_auth_menu` VALUES (207, 201, 'A', '分类新增', '', 0, 'albums:cateAdd', '', '', '', '', 0, 0, 0, 1650341765, 1650341765);
INSERT INTO `x_system_auth_menu` VALUES (208, 201, 'A', '分类命名', '', 0, 'albums:cateRename', '', '', '', '', 0, 0, 0, 1650341765, 1650341765);
INSERT INTO `x_system_auth_menu` VALUES (209, 201, 'A', '分类删除', '', 0, 'albums:cateDel', '', '', '', '', 0, 0, 0, 1650341765, 1650341765);
INSERT INTO `x_system_auth_menu` VALUES (215, 200, 'M', '上传管理', '', 0, '', '', '', '', '', 0, 0, 0, 1650341765, 1650341765);
INSERT INTO `x_system_auth_menu` VALUES (216, 215, 'A', '上传图片', '', 0, 'upload:image', '', '', '', '', 0, 0, 0, 1650341765, 1650341765);
INSERT INTO `x_system_auth_menu` VALUES (217, 215, 'A', '上传视频', '', 0, 'upload:video', '', '', '', '', 0, 0, 0, 1650341765, 1650341765);
INSERT INTO `x_system_auth_menu` VALUES (500, 0, 'M', '系统设置', 'el-icon-Setting', 0, '', 'setting', '', '', '', 0, 1, 0, 1650341765, 1662626322);
INSERT INTO `x_system_auth_menu` VALUES (501, 500, 'M', '网站设置', 'el-icon-Basketball', 10, '', 'website', '', '', '', 0, 1, 0, 1650341765, 1663233572);
INSERT INTO `x_system_auth_menu` VALUES (502, 501, 'C', '网站信息', '', 0, 'setting:website:detail', 'information', 'setting/website/information', '', '', 0, 1, 0, 1650341765, 1660202218);
INSERT INTO `x_system_auth_menu` VALUES (503, 502, 'A', '保存配置', '', 0, 'setting:website:save', '', '', '', '', 0, 0, 0, 1650341765, 1650341765);
INSERT INTO `x_system_auth_menu` VALUES (505, 501, 'C', '网站备案', '', 0, 'setting:copyright:detail', 'filing', 'setting/website/filing', '', '', 0, 1, 0, 1650341765, 1660202294);
INSERT INTO `x_system_auth_menu` VALUES (506, 505, 'A', '备案保存', '', 0, 'setting:copyright:save', '', 'setting/website/protocol', '', '', 0, 0, 0, 1650341765, 1650341765);
INSERT INTO `x_system_auth_menu` VALUES (510, 501, 'C', '政策协议', '', 0, 'setting:protocol:detail', 'protocol', 'setting/website/protocol', '', '', 0, 1, 0, 1660027606, 1660202312);
INSERT INTO `x_system_auth_menu` VALUES (511, 510, 'A', '协议保存', '', 0, 'setting:protocol:save', '', '', '', '', 0, 0, 0, 1660027606, 1663670865);
INSERT INTO `x_system_auth_menu` VALUES (515, 600, 'C', '字典管理', 'el-icon-Box', 0, 'setting:dict:type:list', 'dict', 'setting/dict/type/index', '', '', 0, 1, 0, 1660035436, 1663226087);
INSERT INTO `x_system_auth_menu` VALUES (516, 515, 'A', '字典类型新增', '', 0, 'setting:dict:type:add', '', '', '', '', 0, 1, 0, 1660202761, 1660202761);
INSERT INTO `x_system_auth_menu` VALUES (517, 515, 'A', '字典类型编辑', '', 0, 'setting:dict:type:edit', '', '', '', '', 0, 1, 0, 1660202842, 1660202842);
INSERT INTO `x_system_auth_menu` VALUES (518, 515, 'A', '字典类型删除', '', 0, 'setting:dict:type:del', '', '', '', '', 0, 1, 0, 1660202903, 1660202903);
INSERT INTO `x_system_auth_menu` VALUES (519, 600, 'C', '字典数据管理', '', 0, 'setting:dict:data:list', 'dict/data', 'setting/dict/data/index', '/dev_tools/dict', '', 0, 0, 0, 1660202948, 1663309252);
INSERT INTO `x_system_auth_menu` VALUES (520, 515, 'A', '字典数据新增', '', 0, 'setting:dict:data:add', '', '', '', '', 0, 1, 0, 1660203117, 1660203117);
INSERT INTO `x_system_auth_menu` VALUES (521, 515, 'A', '字典数据编辑', '', 0, 'setting:dict:data:edit', '', '', '', '', 0, 1, 0, 1660203142, 1660203142);
INSERT INTO `x_system_auth_menu` VALUES (522, 515, 'A', '字典数据删除', '', 0, 'setting:dict:data:del', '', '', '', '', 0, 1, 0, 1660203159, 1660203159);
INSERT INTO `x_system_auth_menu` VALUES (550, 500, 'M', '系统维护', 'el-icon-SetUp', 0, '', 'system', '', '', '', 0, 1, 0, 1650341765, 1660202466);
INSERT INTO `x_system_auth_menu` VALUES (551, 550, 'C', '系统环境', '', 0, 'monitor:server', 'environment', 'setting/system/environment', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `x_system_auth_menu` VALUES (552, 550, 'C', '系统缓存', '', 0, 'monitor:cache', 'cache', 'setting/system/cache', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `x_system_auth_menu` VALUES (553, 550, 'C', '系统日志', '', 0, 'system:log:operate', 'journal', 'setting/system/journal', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `x_system_auth_menu` VALUES (555, 500, 'C', '存储设置', 'el-icon-FolderOpened', 6, 'setting:storage:list', 'storage', 'setting/storage/index', '', '', 0, 1, 0, 1650341765, 1663312996);
INSERT INTO `x_system_auth_menu` VALUES (556, 555, 'A', '保存配置', '', 0, 'setting:storage:edit', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `x_system_auth_menu` VALUES (600, 0, 'M', '开发工具', 'el-icon-EditPen', 0, '', 'dev_tools', '', '', '', 0, 1, 0, 1660027606, 1664335701);
INSERT INTO `x_system_auth_menu` VALUES (610, 600, 'C', '代码生成器', 'el-icon-DocumentAdd', 0, 'gen:list', 'code', 'dev_tools/code/index', '', '', 0, 1, 0, 1660028954, 1660532510);
INSERT INTO `x_system_auth_menu` VALUES (611, 610, 'A', '导入数据表', '', 0, 'gen:importTable', '', '', '', '', 0, 1, 0, 1660532389, 1660532389);
INSERT INTO `x_system_auth_menu` VALUES (612, 610, 'A', '生成代码', '', 0, 'gen:genCode', '', '', '', '', 0, 1, 0, 1660532421, 1660532421);
INSERT INTO `x_system_auth_menu` VALUES (613, 610, 'A', '下载代码', '', 0, 'gen:downloadCode', '', '', '', '', 0, 1, 0, 1660532437, 1660532437);
INSERT INTO `x_system_auth_menu` VALUES (614, 610, 'A', '预览代码', '', 0, 'gen:previewCode', '', '', '', '', 0, 1, 0, 1660532549, 1660532549);
INSERT INTO `x_system_auth_menu` VALUES (616, 610, 'A', '同步表结构', '', 0, 'gen:syncTable', '', '', '', '', 0, 1, 0, 1660532781, 1660532781);
INSERT INTO `x_system_auth_menu` VALUES (617, 610, 'A', '删除数据表', '', 0, 'gen:delTable', '', '', '', '', 0, 1, 0, 1660532800, 1660532800);
INSERT INTO `x_system_auth_menu` VALUES (618, 610, 'A', '数据表详情', '', 0, 'gen:detail', '', '', '', '', 0, 1, 0, 1660532964, 1660532977);
INSERT INTO `x_system_auth_menu` VALUES (700, 0, 'M', '素材管理', 'el-icon-Picture', 43, '', 'material', '', '', '', 0, 1, 0, 1660203293, 1663300847);
INSERT INTO `x_system_auth_menu` VALUES (701, 700, 'C', '素材中心', 'el-icon-PictureRounded', 0, '', 'index', 'material/index', '', '', 0, 1, 0, 1660203402, 1663301493);
INSERT INTO `x_system_auth_menu` VALUES (775, 600, 'C', '代码生成器编辑', 'el-icon-EditPen', 0, 'gen:editTable', 'dev_tools/code/edit', 'dev_tools/code/edit', '', '', 0, 0, 0, 1699344389, 1699344389);
INSERT INTO `x_system_auth_menu` VALUES (776, 778, 'C', '流程模板', '', 0, '', 'flow_template/index', 'flow/flow_template/index', '', '', 1, 1, 0, 1702105748, 1702105748);
INSERT INTO `x_system_auth_menu` VALUES (777, 778, 'C', '我的流程', '', 0, '', 'flow_apply/index', 'flow/flow_apply/index', '', '', 1, 1, 0, 1702280379, 1702280379);
INSERT INTO `x_system_auth_menu` VALUES (778, 0, 'M', '审批流', 'el-icon-Coordinate', 0, '', 'flow', '', '', '', 1, 1, 0, 1702309255, 1702309255);
INSERT INTO `x_system_auth_menu` VALUES (780, 778, 'C', '我的代办', '', 0, 'flow_history:todo', 'flow_history/todo', 'flow/flow_history/todo', '', '', 1, 1, 0, 1702900195, 1702900195);
INSERT INTO `x_system_auth_menu` VALUES (781, 778, 'C', '已处理', '', 0, '', 'flow_history/done', 'flow/flow_history/done', '', '', 0, 1, 0, 1703609354, 1703609354);
INSERT INTO `x_system_auth_menu` VALUES (782, 778, 'C', '已完成流程', '', 0, '', 'flow_apply/finish', 'flow/flow_apply/finish', '', '', 1, 1, 0, 1703781608, 1703781608);

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
INSERT INTO `x_system_auth_perm` VALUES ('012e4f07410f4efc85c534559219d8d5', 1, 142);
INSERT INTO `x_system_auth_perm` VALUES ('025098bd36de47bea08c3f4a9a4c8a62', 1, 140);
INSERT INTO `x_system_auth_perm` VALUES ('0ae9224d817e483a8d5f54a95f64472a', 1, 202);
INSERT INTO `x_system_auth_perm` VALUES ('0dbe4f35d6194c8aa9f7bdd70bcff056', 1, 113);
INSERT INTO `x_system_auth_perm` VALUES ('0ee45a508c0e4125bae3f9dcefc34157', 1, 520);
INSERT INTO `x_system_auth_perm` VALUES ('188c6fea6a6a447197a6b88938eac583', 1, 209);
INSERT INTO `x_system_auth_perm` VALUES ('1ecffa13b04a4f769b60342c9cfdecf3', 1, 550);
INSERT INTO `x_system_auth_perm` VALUES ('20967a0955394976a699d7dce28f85dd', 1, 203);
INSERT INTO `x_system_auth_perm` VALUES ('20c5ec0caa0e469db054929c5b317172', 1, 100);
INSERT INTO `x_system_auth_perm` VALUES ('21def579cf48498498958751af06538c', 1, 131);
INSERT INTO `x_system_auth_perm` VALUES ('234a5e62fa654d82864fb225e02f7b2e', 1, 141);
INSERT INTO `x_system_auth_perm` VALUES ('279ba9e84174420bb73a696eb659b8d2', 1, 215);
INSERT INTO `x_system_auth_perm` VALUES ('2c9120e6983b49e693dc52d32634de88', 1, 105);
INSERT INTO `x_system_auth_perm` VALUES ('31f300bf95844745a9ee5520755e0503', 1, 207);
INSERT INTO `x_system_auth_perm` VALUES ('32b216d4108a4462a74d084e1645b952', 1, 121);
INSERT INTO `x_system_auth_perm` VALUES ('32e9f41b980040459664f66e4e10f504', 1, 501);
INSERT INTO `x_system_auth_perm` VALUES ('337d7351a88e4ef0bbe81d3976c4f068', 1, 114);
INSERT INTO `x_system_auth_perm` VALUES ('33c15e5cbefa4adfb729fb335a7fd5b8', 1, 120);
INSERT INTO `x_system_auth_perm` VALUES ('46634386fa54440a8dddd626c35b1c98', 1, 617);
INSERT INTO `x_system_auth_perm` VALUES ('4a2d53ff7c8b4b6eba267f97f41f3aa6', 1, 103);
INSERT INTO `x_system_auth_perm` VALUES ('4b2d5ab21aa844b199ef302a87485938', 1, 511);
INSERT INTO `x_system_auth_perm` VALUES ('5b3568e801ba4114ab7b13fedf403027', 1, 517);
INSERT INTO `x_system_auth_perm` VALUES ('5cb2c5b13aaf4b409c1dddc1154ebc7c', 1, 551);
INSERT INTO `x_system_auth_perm` VALUES ('5dddaee2734746ce83e6b6d18eb2e877', 1, 614);
INSERT INTO `x_system_auth_perm` VALUES ('5ffd556c671d4a71b21fdb8e6bb7fef1', 1, 775);
INSERT INTO `x_system_auth_perm` VALUES ('605751a943304549bc8631398edebb60', 1, 217);
INSERT INTO `x_system_auth_perm` VALUES ('60d006c0a459456b925c95bd5500c089', 1, 610);
INSERT INTO `x_system_auth_perm` VALUES ('67dcd12537e447a38311e114cadd37fc', 1, 205);
INSERT INTO `x_system_auth_perm` VALUES ('687717786d804f52819f2f51aa377c9f', 1, 701);
INSERT INTO `x_system_auth_perm` VALUES ('6cdff8dcbbe04e64bc5ae3d749c58bcd', 1, 618);
INSERT INTO `x_system_auth_perm` VALUES ('700152c35a3848dd9c5faba4a62694bf', 1, 130);
INSERT INTO `x_system_auth_perm` VALUES ('716f5d2c17864d75849b8c387bf1ffb4', 1, 781);
INSERT INTO `x_system_auth_perm` VALUES ('728c984cde6340debc8b3713b831e224', 1, 613);
INSERT INTO `x_system_auth_perm` VALUES ('728f20230b49432db64365a4a3d4c593', 1, 700);
INSERT INTO `x_system_auth_perm` VALUES ('77f062375d2845d5993da143ea61f13a', 1, 780);
INSERT INTO `x_system_auth_perm` VALUES ('7f8c8e94e8b74746b9086f425d32d52c', 1, 506);
INSERT INTO `x_system_auth_perm` VALUES ('7ffc097d72704ac3afc01495c0dd8722', 1, 200);
INSERT INTO `x_system_auth_perm` VALUES ('82148c8343694cbd954984a5173b4307', 1, 134);
INSERT INTO `x_system_auth_perm` VALUES ('85a2b96e7ba74b5eb5821aa46b969a06', 1, 133);
INSERT INTO `x_system_auth_perm` VALUES ('86c24a38da39491e92494f2f3723ee05', 1, 110);
INSERT INTO `x_system_auth_perm` VALUES ('87240bae344946d1af5c54ab7f781308', 1, 500);
INSERT INTO `x_system_auth_perm` VALUES ('891a71afd2564ef3a6ae942085d03e13', 1, 505);
INSERT INTO `x_system_auth_perm` VALUES ('89fd1415ae80468fb1bdbc111156b6fd', 1, 522);
INSERT INTO `x_system_auth_perm` VALUES ('8b585e3cb25f4a7f9417761dac35b7dd', 1, 553);
INSERT INTO `x_system_auth_perm` VALUES ('8c390106e0474d858ad6c83d3be38856', 1, 132);
INSERT INTO `x_system_auth_perm` VALUES ('92983d912b2244a99e7efcb8ebe95bb2', 1, 1);
INSERT INTO `x_system_auth_perm` VALUES ('9b73f189d5364693868e814dae03ace8', 1, 112);
INSERT INTO `x_system_auth_perm` VALUES ('9d7b6cba9ab1425fb1a2a2452f1ebb1d', 1, 208);
INSERT INTO `x_system_auth_perm` VALUES ('a145cf6dc5d84011847f4db5bb667e44', 1, 521);
INSERT INTO `x_system_auth_perm` VALUES ('a1dfa995f2f846688959950ee34fdb03', 1, 778);
INSERT INTO `x_system_auth_perm` VALUES ('a2455fb8841043b282681a613421aa04', 1, 502);
INSERT INTO `x_system_auth_perm` VALUES ('a3ee8518aa964517a18342e5910127bc', 1, 143);
INSERT INTO `x_system_auth_perm` VALUES ('a4d612f834a2447ea964697a5cb25e27', 1, 516);
INSERT INTO `x_system_auth_perm` VALUES ('a55298adb7a64cce8f27a2a990ebb3bc', 1, 515);
INSERT INTO `x_system_auth_perm` VALUES ('a8faa21b39884c0e992a9972b3d73e6d', 1, 519);
INSERT INTO `x_system_auth_perm` VALUES ('b05b62185a4d49648df54eab445e05a4', 1, 122);
INSERT INTO `x_system_auth_perm` VALUES ('baa1543f551743619110f2db43e01410', 1, 102);
INSERT INTO `x_system_auth_perm` VALUES ('bd3382bd2c2c41dc8c7c30a63779ab8c', 1, 216);
INSERT INTO `x_system_auth_perm` VALUES ('bf9ae4b1c1d1487197dae35414c27764', 1, 518);
INSERT INTO `x_system_auth_perm` VALUES ('c50b243d1f0c472b8afd76edd37de1a9', 1, 556);
INSERT INTO `x_system_auth_perm` VALUES ('c9ba2240d8a248f59331f143a32aad7d', 1, 201);
INSERT INTO `x_system_auth_perm` VALUES ('cb8230cad0574c2a8247432e37efe24b', 1, 503);
INSERT INTO `x_system_auth_perm` VALUES ('cd38190fe74b4cc6b17b4a49cc6ba93e', 1, 206);
INSERT INTO `x_system_auth_perm` VALUES ('cf7c38297d444245a4b731253f0e5cb0', 1, 600);
INSERT INTO `x_system_auth_perm` VALUES ('d18d9e4531004bcea6d3306deed2bc4d', 1, 123);
INSERT INTO `x_system_auth_perm` VALUES ('d32e4b32fb9b418c95ae41af8c2a090d', 1, 101);
INSERT INTO `x_system_auth_perm` VALUES ('d53fcc6ef5d34d93ad7a6fc5f3092eeb', 1, 611);
INSERT INTO `x_system_auth_perm` VALUES ('dba25132d9f24e0491d011775aedb793', 1, 776);
INSERT INTO `x_system_auth_perm` VALUES ('dc1bd69205934489b9d2eaae2793a323', 1, 777);
INSERT INTO `x_system_auth_perm` VALUES ('def74fdf7022474aaeb603715fad332f', 1, 616);
INSERT INTO `x_system_auth_perm` VALUES ('df8d02c85e43483fad98bb8548d503a8', 1, 612);
INSERT INTO `x_system_auth_perm` VALUES ('e43c347605844f2b8e225c824baff10a', 1, 111);
INSERT INTO `x_system_auth_perm` VALUES ('ec25bbe0a2ea4d9a913b6e0fd6916400', 1, 124);
INSERT INTO `x_system_auth_perm` VALUES ('eed789ac91ad47d2af770e6d0038e7cb', 1, 144);
INSERT INTO `x_system_auth_perm` VALUES ('f01cab016d544b709983ad21d25b1b2d', 1, 204);
INSERT INTO `x_system_auth_perm` VALUES ('f19a655fd0bb47c19bcecfc4c8378f7f', 1, 135);
INSERT INTO `x_system_auth_perm` VALUES ('f1dcfa38b2e74f3299a0af75d58cd09a', 1, 106);
INSERT INTO `x_system_auth_perm` VALUES ('f64353ddcf754c35be0a35aeec3d7afc', 1, 510);
INSERT INTO `x_system_auth_perm` VALUES ('f81f2dfaa3d949458d4b346d2891ccef', 1, 104);
INSERT INTO `x_system_auth_perm` VALUES ('fab1d825eab54e1b97e9df67095972cb', 1, 555);
INSERT INTO `x_system_auth_perm` VALUES ('fb936c77e8644d2a9738802213a231ba', 1, 552);

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
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  `delete_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统岗位管理表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_system_auth_post
-- ----------------------------
INSERT INTO `x_system_auth_post` VALUES (2, 'gw0001', '默认岗位', '', 3, 0, 0, 1700821779, 1702212419, 0);
INSERT INTO `x_system_auth_post` VALUES (3, 'zhihuibu01', '指挥部岗位', '', 0, 0, 0, 1702884035, 1702884035, 0);

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
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统角色管理表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_system_auth_role
-- ----------------------------
INSERT INTO `x_system_auth_role` VALUES (1, '审核员', '审核数据', 0, 0, 1668679451, 1703780348);

-- ----------------------------
-- Table structure for x_system_config
-- ----------------------------
DROP TABLE IF EXISTS `x_system_config`;
CREATE TABLE `x_system_config`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `type` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '类型',
  `name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '键',
  `value` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '值',
  `create_time` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 81 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统全局配置表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_system_config
-- ----------------------------
INSERT INTO `x_system_config` VALUES (1, 'storage', 'default', '', 1660620367, 1702113968);
INSERT INTO `x_system_config` VALUES (2, 'storage', 'local', '{\"name\":\"本地存储\"}', 1660620367, 1702113968);
INSERT INTO `x_system_config` VALUES (3, 'storage', 'qiniu', '{\"name\":\"七牛云存储\",\"bucket\":\"\",\"secretKey\":\"\",\"accessKey\":\"\",\"domain\":\"\"}', 1660620367, 1660620367);
INSERT INTO `x_system_config` VALUES (4, 'storage', 'aliyun', '{\"name\":\"阿里云存储\",\"bucket\":\"\",\"secretKey\":\"\",\"accessKey\":\"\",\"domain\":\"\"}', 1660620367, 1662620071);
INSERT INTO `x_system_config` VALUES (5, 'storage', 'qcloud', '{\"name\":\"腾讯云存储\",\"bucket\":\"\",\"secretKey\":\"\",\"accessKey\":\"\",\"domain\":\"\",\"region\":\"\"}', 1660620367, 1660620367);
INSERT INTO `x_system_config` VALUES (6, 'sms', 'default', 'aliyun', 1660620367, 1660620367);
INSERT INTO `x_system_config` VALUES (7, 'sms', 'aliyun', '{\"name\":\"阿里云短信\",\"alias\":\"aliyun\",\"sign\":\"\",\"appKey\":\"\",\"secretKey\":\"\"}', 1660620367, 1660620367);
INSERT INTO `x_system_config` VALUES (8, 'sms', 'tencent', '{\"name\":\"腾讯云短信\",\"alias\":\"tencent\",\"sign\":\"\",\"appId\":\"\",\"secretId\":\"\",\"secretKey\":\"\"}', 1660620367, 1660620367);
INSERT INTO `x_system_config` VALUES (9, 'sms', 'huawei', '{\"name\":\"华为云短信\",\"alias\":\"huawei\"}', 1660620367, 1660620367);
INSERT INTO `x_system_config` VALUES (10, 'website', 'name', 'Admin开源后台', 1660620367, 1699524151);
INSERT INTO `x_system_config` VALUES (11, 'website', 'logo', '/api/static/backend_logo.png', 1660620367, 1699524151);
INSERT INTO `x_system_config` VALUES (12, 'website', 'favicon', '/api/static/backend_favicon.ico', 1660620367, 1699524151);
INSERT INTO `x_system_config` VALUES (13, 'website', 'backdrop', '/api/static/backend_backdrop.png', 1660620367, 1699524151);
INSERT INTO `x_system_config` VALUES (14, 'website', 'copyright', '[{\"name\":\"LikeAdmin开源系统\",\"link\":\"http://www.beian.gov.cn\"}]', 1660620367, 1660620367);
INSERT INTO `x_system_config` VALUES (15, 'website', 'shopName', 'Admin开源系统', 1631255140, 1699524151);
INSERT INTO `x_system_config` VALUES (16, 'website', 'shopLogo', '/api/static/shop_logo.png', 1631255140, 1699524151);
INSERT INTO `x_system_config` VALUES (17, 'protocol', 'service', '{\"name\":\"服务协议\",\"content\":\"\\u003cp\\u003e服务协议\\u003c/p\\u003e\"}', 1660620367, 1699496132);
INSERT INTO `x_system_config` VALUES (18, 'protocol', 'privacy', '{\"name\":\"隐私协议\",\"content\":\"\\u003cp\\u003e隐私协议\\u003c/p\\u003e\"}', 1660620367, 1699496132);
INSERT INTO `x_system_config` VALUES (19, 'tabbar', 'style', '{\"defaultColor\":\"#4A5DFF\",\"selectedColor\":\"#EA5455\"}', 1660620367, 1662544900);
INSERT INTO `x_system_config` VALUES (20, 'search', 'isHotSearch', '0', 1660620367, 1662546997);
INSERT INTO `x_system_config` VALUES (30, 'h5_channel', 'status', '1', 1660620367, 1660620367);
INSERT INTO `x_system_config` VALUES (31, 'h5_channel', 'close', '0', 1660620367, 1660620367);
INSERT INTO `x_system_config` VALUES (32, 'h5_channel', 'url', '', 1660620367, 1660620367);
INSERT INTO `x_system_config` VALUES (40, 'mp_channel', 'name', '', 1660620367, 1662551403);
INSERT INTO `x_system_config` VALUES (41, 'mp_channel', 'primaryId', '', 1660620367, 1662551403);
INSERT INTO `x_system_config` VALUES (42, 'mp_channel', 'appId', '', 1660620367, 1662551403);
INSERT INTO `x_system_config` VALUES (43, 'mp_channel', 'appSecret', '', 1660620367, 1662551403);
INSERT INTO `x_system_config` VALUES (44, 'mp_channel', 'qrCode', '', 1660620367, 1662551403);
INSERT INTO `x_system_config` VALUES (50, 'wx_channel', 'appId', '', 1660620367, 1660620367);
INSERT INTO `x_system_config` VALUES (51, 'wx_channel', 'appSecret', '', 1660620367, 1660620367);
INSERT INTO `x_system_config` VALUES (55, 'oa_channel', 'name', '', 1660620367, 1662551337);
INSERT INTO `x_system_config` VALUES (56, 'oa_channel', 'primaryId', ' ', 1660620367, 1662551337);
INSERT INTO `x_system_config` VALUES (57, 'oa_channel', 'qrCode', '', 1662551337, 1662551337);
INSERT INTO `x_system_config` VALUES (58, 'oa_channel', 'appId', '', 1660620367, 1662551337);
INSERT INTO `x_system_config` VALUES (59, 'oa_channel', 'appSecret', '', 1660620367, 1662551337);
INSERT INTO `x_system_config` VALUES (60, 'oa_channel', 'url', '', 1660620367, 1662551337);
INSERT INTO `x_system_config` VALUES (61, 'oa_channel', 'token', '', 1660620367, 1662551337);
INSERT INTO `x_system_config` VALUES (62, 'oa_channel', 'encodingAesKey', '', 1660620367, 1662551337);
INSERT INTO `x_system_config` VALUES (63, 'oa_channel', 'encryptionType', '1', 1660620367, 1662551337);
INSERT INTO `x_system_config` VALUES (64, 'oa_channel', 'menus', '[]', 1631255140, 1663118712);
INSERT INTO `x_system_config` VALUES (70, 'login', 'loginWay', '1,2', 1660620367, 1662538771);
INSERT INTO `x_system_config` VALUES (71, 'login', 'forceBindMobile', '0', 1660620367, 1662538771);
INSERT INTO `x_system_config` VALUES (72, 'login', 'openAgreement', '1', 1660620367, 1662538771);
INSERT INTO `x_system_config` VALUES (73, 'login', 'openOtherAuth', '1', 1660620367, 1662538771);
INSERT INTO `x_system_config` VALUES (74, 'login', 'autoLoginAuth', '1,2', 1660620367, 1662538771);
INSERT INTO `x_system_config` VALUES (80, 'user', 'defaultAvatar', '/api/static/default_avatar.png', 1660620367, 1662535156);

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
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 20 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统登录日志表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_system_log_login
-- ----------------------------
INSERT INTO `x_system_log_login` VALUES (1, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, 1699342613);
INSERT INTO `x_system_log_login` VALUES (2, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, 1703215805);
INSERT INTO `x_system_log_login` VALUES (3, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, 1703222067);
INSERT INTO `x_system_log_login` VALUES (4, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, 1703229521);
INSERT INTO `x_system_log_login` VALUES (5, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, 1703256368);
INSERT INTO `x_system_log_login` VALUES (6, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, 1703305890);
INSERT INTO `x_system_log_login` VALUES (7, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, 1703313437);
INSERT INTO `x_system_log_login` VALUES (8, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, 1703327546);
INSERT INTO `x_system_log_login` VALUES (9, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, 1703344025);
INSERT INTO `x_system_log_login` VALUES (10, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, 1703495127);
INSERT INTO `x_system_log_login` VALUES (11, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, 1703520093);
INSERT INTO `x_system_log_login` VALUES (12, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, 1703581331);
INSERT INTO `x_system_log_login` VALUES (13, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, 1703598707);
INSERT INTO `x_system_log_login` VALUES (14, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, 1703649048);
INSERT INTO `x_system_log_login` VALUES (15, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, 1703656752);
INSERT INTO `x_system_log_login` VALUES (16, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, 1703736103);
INSERT INTO `x_system_log_login` VALUES (17, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, 1703752713);
INSERT INTO `x_system_log_login` VALUES (18, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, 1703760991);
INSERT INTO `x_system_log_login` VALUES (19, 1, 'admin', '127.0.0.1', 'Windows', 'Edge', 1, 1703773878);

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
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 32 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统操作日志表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_system_log_operate
-- ----------------------------
INSERT INTO `x_system_log_operate` VALUES (1, 1, 'GET', '角色列表', '127.0.0.1', '/api/system/role/list', 'likeadmin/admin/routers/system.roleHandler.list-fm', 'pageNo=1&pageSize=15', '', 1, 1699343137, 1699343138, 4, 1699343138);
INSERT INTO `x_system_log_operate` VALUES (2, 1, 'GET', '角色列表', '127.0.0.1', '/api/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, 1703224047, 1703224047, 6, 1703224047);
INSERT INTO `x_system_log_operate` VALUES (3, 1, 'GET', '角色列表', '127.0.0.1', '/api/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, 1703231551, 1703231551, 5, 1703231551);
INSERT INTO `x_system_log_operate` VALUES (4, 1, 'GET', '角色列表', '127.0.0.1', '/api/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, 1703313449, 1703313449, 3, 1703313449);
INSERT INTO `x_system_log_operate` VALUES (5, 1, 'GET', '服务监控', '127.0.0.1', '/api/monitor/server', 'x_admin/admin/monitor.monitorHandler.server-fm', '', '', 1, 1703313470, 1703313470, 73, 1703313470);
INSERT INTO `x_system_log_operate` VALUES (6, 1, 'GET', '缓存监控', '127.0.0.1', '/api/monitor/cache', 'x_admin/admin/monitor.monitorHandler.cache-fm', '', '', 1, 1703313473, 1703313473, 3, 1703313473);
INSERT INTO `x_system_log_operate` VALUES (7, 1, 'GET', '服务监控', '127.0.0.1', '/api/monitor/server', 'x_admin/admin/monitor.monitorHandler.server-fm', '', '', 1, 1703313960, 1703313960, 12, 1703313960);
INSERT INTO `x_system_log_operate` VALUES (8, 1, 'GET', '角色列表', '127.0.0.1', '/api/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, 1703354126, 1703354126, 42, 1703354126);
INSERT INTO `x_system_log_operate` VALUES (9, 1, 'GET', '角色列表', '127.0.0.1', '/api/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, 1703354155, 1703354155, 3, 1703354155);
INSERT INTO `x_system_log_operate` VALUES (10, 1, 'GET', '角色列表', '127.0.0.1', '/api/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, 1703354159, 1703354159, 2, 1703354159);
INSERT INTO `x_system_log_operate` VALUES (11, 1, 'GET', '角色列表', '127.0.0.1', '/api/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, 1703354293, 1703354293, 3, 1703354293);
INSERT INTO `x_system_log_operate` VALUES (12, 1, 'GET', '角色列表', '127.0.0.1', '/api/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, 1703354311, 1703354311, 85, 1703354311);
INSERT INTO `x_system_log_operate` VALUES (13, 1, 'GET', '角色详情', '127.0.0.1', '/api/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=1', '', 1, 1703354313, 1703354313, 5, 1703354313);
INSERT INTO `x_system_log_operate` VALUES (14, 1, 'POST', '角色编辑', '127.0.0.1', '/api/system/role/edit', 'x_admin/admin/system/role.RoleHandler.Edit-fm', '{\"id\":1,\"isDisable\":0,\"menuIds\":\"1,130,131,132,133,134,135,140,141,142,143,144,100,101,102,103,104,105,106,110,111,112,113,114,120,121,122,123,124,700,701,200,201,202,203,204,205,206,207,208,209,215,216,217,500,501,502,503,505,506,510,511,555,556,550,551,552,553,600,515,516,517,518,520,521,522,519,610,611,612,613,614,616,617,618,775,778,776,777,779,780\",\"menus\":[1,130,131,132,133,134,135,140,141,142,143,144,100,101,102,103,104,105,106,110,111,112,113,114,120,121,122,123,124,700,701,200,201,202,203,204,205,206,207,208,209,215,216,217,500,501,502,503,505,506,510,511,555,556,550,551,552,553,600,515,516,517,518,520,521,522,519,610,611,612,613,614,616,617,618,775,778,776,777,779,780],\"name\":\"审核员\",\"remark\":\"审核数据\",\"sort\":0}', '', 1, 1703354326, 1703354326, 63, 1703354326);
INSERT INTO `x_system_log_operate` VALUES (15, 1, 'GET', '角色列表', '127.0.0.1', '/api/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, 1703354326, 1703354326, 2, 1703354326);
INSERT INTO `x_system_log_operate` VALUES (16, 1, 'GET', '角色详情', '127.0.0.1', '/api/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=1', '', 1, 1703354342, 1703354342, 4, 1703354342);
INSERT INTO `x_system_log_operate` VALUES (17, 1, 'GET', '角色详情', '127.0.0.1', '/api/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=1', '', 1, 1703354352, 1703354352, 3, 1703354352);
INSERT INTO `x_system_log_operate` VALUES (18, 1, 'GET', '角色列表', '127.0.0.1', '/api/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, 1703495181, 1703495181, 1, 1703495181);
INSERT INTO `x_system_log_operate` VALUES (19, 1, 'GET', '服务监控', '127.0.0.1', '/api/monitor/server', 'x_admin/admin/monitor.monitorHandler.server-fm', '', '', 1, 1703495196, 1703495196, 8, 1703495196);
INSERT INTO `x_system_log_operate` VALUES (20, 1, 'GET', '缓存监控', '127.0.0.1', '/api/monitor/cache', 'x_admin/admin/monitor.monitorHandler.cache-fm', '', '', 1, 1703495198, 1703495198, 1, 1703495198);
INSERT INTO `x_system_log_operate` VALUES (21, 1, 'GET', '角色列表', '127.0.0.1', '/api/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, 1703581345, 1703581345, 2, 1703581345);
INSERT INTO `x_system_log_operate` VALUES (22, 1, 'GET', '角色列表', '127.0.0.1', '/api/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, 1703581355, 1703581355, 2, 1703581355);
INSERT INTO `x_system_log_operate` VALUES (23, 1, 'POST', '管理员编辑', '127.0.0.1', '/api/system/admin/edit', 'x_admin/admin/system/admin.AdminHandler.Edit-fm', '{\"avatar\":\"/api/uploads/image/20232911/27699a1e991e4ba1ba13ee6d45185db6.jpg\",\"deptId\":2,\"id\":2,\"isDisable\":0,\"isMultipoint\":1,\"nickname\":\"指挥部01\",\"password\":\"\",\"passwordConfirm\":\"\",\"postId\":3,\"role\":1,\"sort\":1,\"username\":\"zhihuibu01\"}', '', 1, 1703599817, 1703599817, 51, 1703599817);
INSERT INTO `x_system_log_operate` VALUES (24, 1, 'GET', '服务监控', '127.0.0.1', '/api/monitor/server', 'x_admin/admin/monitor.monitorHandler.server-fm', '', '', 1, 1703774438, 1703774438, 8, 1703774438);
INSERT INTO `x_system_log_operate` VALUES (25, 1, 'GET', '角色列表', '127.0.0.1', '/api/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, 1703775460, 1703775460, 1, 1703775460);
INSERT INTO `x_system_log_operate` VALUES (26, 1, 'GET', '角色列表', '127.0.0.1', '/api/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, 1703775472, 1703775472, 1, 1703775472);
INSERT INTO `x_system_log_operate` VALUES (27, 1, 'GET', '角色列表', '127.0.0.1', '/api/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, 1703780333, 1703780333, 4, 1703780333);
INSERT INTO `x_system_log_operate` VALUES (28, 1, 'GET', '角色详情', '127.0.0.1', '/api/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=1', '', 1, 1703780337, 1703780337, 34, 1703780337);
INSERT INTO `x_system_log_operate` VALUES (29, 1, 'GET', '角色详情', '127.0.0.1', '/api/system/role/detail', 'x_admin/admin/system/role.RoleHandler.Detail-fm', 'id=1', '', 1, 1703780339, 1703780339, 4, 1703780339);
INSERT INTO `x_system_log_operate` VALUES (30, 1, 'POST', '角色编辑', '127.0.0.1', '/api/system/role/edit', 'x_admin/admin/system/role.RoleHandler.Edit-fm', '{\"id\":1,\"isDisable\":0,\"menuIds\":\"778,1,130,131,132,133,134,135,140,141,142,143,144,100,101,102,103,104,105,106,110,111,112,113,114,120,121,122,123,124,700,701,200,201,202,203,204,205,206,207,208,209,215,216,217,500,501,502,503,505,506,510,511,555,556,550,551,552,553,600,515,516,517,518,520,521,522,519,610,611,612,613,614,616,617,618,775,776,777,780,781\",\"menus\":[778,1,130,131,132,133,134,135,140,141,142,143,144,100,101,102,103,104,105,106,110,111,112,113,114,120,121,122,123,124,700,701,200,201,202,203,204,205,206,207,208,209,215,216,217,500,501,502,503,505,506,510,511,555,556,550,551,552,553,600,515,516,517,518,520,521,522,519,610,611,612,613,614,616,617,618,775,776,777,780,781],\"name\":\"审核员\",\"remark\":\"审核数据\",\"sort\":0}', '', 1, 1703780348, 1703780348, 108, 1703780348);
INSERT INTO `x_system_log_operate` VALUES (31, 1, 'GET', '角色列表', '127.0.0.1', '/api/system/role/list', 'x_admin/admin/system/role.RoleHandler.List-fm', 'pageNo=1&pageSize=15', '', 1, 1703780348, 1703780348, 2, 1703780348);

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
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
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
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  `delete_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
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
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `openid`(`openid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户授权表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of x_user_auth
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
