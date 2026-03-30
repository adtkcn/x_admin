---
name: "laravel-admin-migration"
description: "Generates Go+Vue code from Laravel-Admin PHP code. Invoke when user wants to migrate code from PHP Laravel-Admin project to likeadmin Go."
---

# Laravel-Admin 代码迁移指南

本技能用于从 PHP Laravel-Admin 项目迁移代码到 likeadmin Go 项目。

## 迁移流程概述

### 步骤一：分析 Laravel-Admin 代码

分析 Laravel-Admin 项目中的以下文件：
- **Controller**：获取接口列表、业务逻辑
- **Model**：获取数据库字段定义
- **Form/Grid**：获取表单字段、验证规则

### 步骤二：使用现有代码生成系统

项目已内置完整的代码生成系统，位于：
```
app/service/generatorService/tpl_utils/templates/
```

#### Go 后端模板（直接使用）

| 模板文件 | 说明 |
|----------|------|
| [gocode/schema.go.tpl](file:///f:/xiangheng/go/likeadmin_go/server/app/service/generatorService/tpl_utils/templates/gocode/schema.go.tpl) | 数据结构定义 |
| [gocode/model.go.tpl](file:///f:/xiangheng/go/likeadmin_go/server/app/service/generatorService/tpl_utils/templates/gocode/model.go.tpl) | 数据模型 |
| [gocode/service.go.tpl](file:///f:/xiangheng/go/likeadmin_go/server/app/service/generatorService/tpl_utils/templates/gocode/service.go.tpl) | 业务逻辑服务 |
| [gocode/controller.go.tpl](file:///f:/xiangheng/go/likeadmin_go/server/app/service/generatorService/tpl_utils/templates/gocode/controller.go.tpl) | 控制器 |
| [gocode/route.go.tpl](file:///f:/xiangheng/go/likeadmin_go/server/app/service/generatorService/tpl_utils/templates/gocode/route.go.tpl) | 路由注册 |

#### Vue 前端模板（直接使用）

| 模板文件 | 说明 |
|----------|------|
| [vue/index.vue.tpl](file:///f:/xiangheng/go/likeadmin_go/server/app/service/generatorService/tpl_utils/templates/vue/index.vue.tpl) | 列表页 |
| [vue/index-tree.vue.tpl](file:///f:/xiangheng/go/likeadmin_go/server/app/service/generatorService/tpl_utils/templates/vue/index-tree.vue.tpl) | 树状列表页 |
| [vue/edit.vue.tpl](file:///f:/xiangheng/go/likeadmin_go/server/app/service/generatorService/tpl_utils/templates/vue/edit.vue.tpl) | 编辑页 |
| [vue/details.vue.tpl](file:///f:/xiangheng/go/likeadmin_go/server/app/service/generatorService/tpl_utils/templates/vue/details.vue.tpl) | 详情页 |
| [vue/api.ts.tpl](file:///f:/xiangheng/go/likeadmin_go/server/app/service/generatorService/tpl_utils/templates/vue/api.ts.tpl) | API 接口 |


### 步骤三：参考现有实现

参考项目现有实现创建业务代码：

#### 参考文件

| 文件路径 | 说明 |
|----------|------|
| [app/schema/user_protocol_schema.go](file:///f:/xiangheng/go/likeadmin_go/server/app/schema/user_protocol_schema.go) | Schema 定义示例 |
| [app/service/user_protocol_service.go](file:///f:/xiangheng/go/likeadmin_go/server/app/service/user_protocol_service.go) | Service 实现示例 |
| [app/controller/admin_ctl/user_protocol_ctl.go](file:///f:/xiangheng/go/likeadmin_go/server/app/controller/admin_ctl/user_protocol_ctl.go) | Controller 实现示例 |
| [routes/adminRoute/user_protocol_route.go](file:///f:/xiangheng/go/likeadmin_go/server/routes/adminRoute/user_protocol_route.go) | 路由注册示例 |

### 步骤四：数据库表结构

根据 Laravel-Admin 的迁移文件生成 MySQL 表结构：

```sql
CREATE TABLE `x_{table_name}` (
  `id` varchar(36) NOT NULL AUTO_INCREMENT COMMENT '主键',

  `is_delete` tinyint(1) DEFAULT 0 COMMENT '是否删除',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
  `created_by` varchar(36) DEFAULT NULL COMMENT '创建人',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='{comment}';
```

### 步骤五：菜单 SQL

```sql
-- 菜单
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, paths, component, is_cache, is_show, is_disable, create_time, update_time) 
VALUES (0, 'C', '{MenuName}', '{module}/index', '{module}/index', 0, 1, 0, now(), now());

-- 按钮（替换pid为菜单id）
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms, is_cache, is_show, is_disable, create_time, update_time) 
VALUES (0, 'A', '{MenuName}添加', 'admin:{module}:add', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms, is_cache, is_show, is_disable, create_time, update_time) 
VALUES (0, 'A', '{MenuName}编辑', 'admin:{module}:edit', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms, is_cache, is_show, is_disable, create_time, update_time) 
VALUES (0, 'A', '{MenuName}删除', 'admin:{module}:del', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms, is_cache, is_show, is_disable, create_time, update_time) 
VALUES (0, 'A', '{MenuName}批量删除', 'admin:{module}:delBatch', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms, is_cache, is_show, is_disable, create_time, update_time) 
VALUES (0, 'A', '{MenuName}列表', 'admin:{module}:list', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms, is_cache, is_show, is_disable, create_time, update_time) 
VALUES (0, 'A', '{MenuName}全部列表', 'admin:{module}:listAll', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms, is_cache, is_show, is_disable, create_time, update_time) 
VALUES (0, 'A', '{MenuName}详情', 'admin:{module}:detail', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms, is_cache, is_show, is_disable, create_time, update_time) 
VALUES (0, 'A', '{MenuName}导出', 'admin:{module}:ExportFile', 0, 1, 0, now(), now());
INSERT INTO x_system_auth_menu (pid, menu_type, menu_name, perms, is_cache, is_show, is_disable, create_time, update_time) 
VALUES (0, 'A', '{MenuName}导入', 'admin:{module}:ImportFile', 0, 1, 0, now(), now());
```

## 代码实现要点

### 1. Schema 定义

```go
// app/schema/{module}_schema.go
package schema

import (
    "x_admin/app/schema/systemSchema"
    "x_admin/core"
)

type {Module}Primarykey struct {
    Id string
}

type {Module}ListReq struct {
    Title   x_null.String
    Content x_null.String
    CreateTimeStart x_null.String
    CreateTimeEnd   x_null.String
}

type {Module}AddReq struct {
    Title   x_null.String
    Content x_null.String
    Sort    x_null.Int64
}

type {Module}EditReq struct {
    {Module}Primarykey
    {Module}AddReq
}

type {Module}DelBatchReq struct {
    Ids string
}

type {Module}Resp struct {
    {Module}Primarykey
    Title      x_null.String
    Content    x_null.String
    Sort       x_null.Int64
    CreateTime x_null.Time
    UpdateTime x_null.Time
    CreatedBy  string
    CreatedByUser systemSchema.SystemAuthAdminSimpleInfo
}
```

### 2. Controller 实现

关键点：
- 使用 `response.CheckAndRespWithData` 进行响应
- 使用 `util.VerifyUtil.VerifyJSON/VerifyQuery` 进行参数验证
- 使用单飞缓存防止缓存击穿：`singleflight.Group`

### 3. Service 实现

关键点：
- 使用 `response.CheckErr`、`response.CheckMysqlErr`、`response.CheckDBNotRecord` 处理错误
- 使用 `convert_util.Copy` 进行对象拷贝
- 使用软删除（is_delete）机制
- 实现缓存机制提高性能
- 使用 map 更新字段以区分空值

### 4. Route 注册

关键点：
- 使用 `init()` 函数自动注册路由
- 使用 `middleware.PermAuth()` 进行认证
- 使用 `middleware.RecordLog` 记录操作日志

## 迁移占位符说明

| 占位符 | 说明 | 示例 |
|--------|------|------|
| `{Module}` | 模块名称（首字母大写） | UserProtocol |
| `{module}` | 模块名称（全小写） | user_protocol |
| `{ModuleName}` | 模块中文名称 | 用户协议 |
| `{MenuName}` | 菜单名称 | 用户协议 |
| `{ModelName}` | 模型名称 | UserProtocol |
| `{table_name}` | 数据库表名 | user_protocol |

## 使用示例

当用户需要从 Laravel-Admin 迁移代码时：

1. **分析 Laravel-Admin 代码**：获取字段列表、接口定义
2. **创建数据库表**：根据字段生成 MySQL 表结构
3. **生成后端代码**：参考现有模板生成 Go 代码
4. **生成前端代码**：使用现有 Vue/Uniapp 模板
5. **注册路由**：在 route.go 中注册新路由
6. **创建菜单**：执行菜单 SQL

## 注意事项

1. 保持与现有代码风格一致
2. 使用 `response.CheckAndRespWithData` 进行响应
3. 使用 `util.VerifyUtil.VerifyJSON/VerifyQuery` 进行参数验证
4. 使用 `convert_util.Copy` 进行对象拷贝
5. 使用软删除（is_delete）机制
6. 实现缓存机制提高性能
7. 路由注册通过 `init()` 函数自动收集
