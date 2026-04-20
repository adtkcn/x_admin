---
name: "项目目录文件结构"
description: "x_admin 后台管理系统 - Go/Gin/GORM 项目完整技术文档"
---

# x_admin 项目技术文档

## 项目概述

x_admin 是一个基于 **Go 1.26** 的后台管理系统后端服务（LikeAdmin Go 版本），采用前后端分离架构。

**模块名**: `x_admin`  
**入口文件**: `main.go`  
**配置文件**: `.env.yaml`（不提交 git）  
**API 文档**: `/api/static/api/index.html`（Swagger）

## 技术栈

| 类别 | 技术 | 版本 |
|------|------|------|
| Web 框架 | [Gin](https://github.com/gin-gonic/gin) | v1.12.0 |
| ORM | [GORM](https://gorm.io) | v1.31.1 |
| 数据库 | MySQL (go-sql-driver) | v1.9.3 |
| 缓存 | Redis (go-redis) | v9.18.0 |
| 配置管理 | Viper (YAML) | v1.21.0 |
| 日志 | Zap + Lumberjack | v1.27.1 |
| 参数校验 | go-playground/validator | v10.30.1 |
| Excel 处理 | excelize | v2.10.1 |
| API 文档 | Swaggo | v1.16.6 |
| 定时任务 | robfig/cron | v3.0.1 |
| 邮件 | go-mail | v0.7.2 |
| WebSocket | gorilla/websocket | v1.5.3 |
| 验证码 | aj-captcha-go (自研) | - |

## 目录结构

```
server/
├── app/                          # 应用程序核心代码
│   ├── controller/               # 控制器层 (MVC - Controller)
│   │   ├── admin_ctl/            # 后台管理控制器
│   │   │   ├── commonController/ # 通用：上传、相册、首页、个推
│   │   │   ├── flowController/   # 审批流程
│   │   │   ├── generatorController/# 代码生成器
│   │   │   ├── monitorController/ # 系统监控
│   │   │   ├── settingController/ # 系统设置
│   │   │   ├── systemController/  # 系统管理（用户/角色/菜单/部门等）
│   │   │   ├── system_corn_ctl.go # 定时任务控制
│   │   │   └── user_protocol_ctl.go # 用户协议
│   │   └── ws.go                 # WebSocket 控制器
│   ├── service/                  # 服务层 (业务逻辑)
│   │   ├── commonService/        # 通用服务
│   │   ├── cornService/          # 定时任务服务
│   │   ├── flowService/          # 流程审批服务
│   │   ├── generatorService/     # 代码生成器服务（含 .tpl 模板）
│   │   ├── monitorService/       # 监控服务
│   │   ├── settingService/       # 设置服务
│   │   ├── systemService/        # 系统管理服务
│   │   └── user_protocol_service.go
│   ├── schema/                   # 请求与响应结构体定义 (DTO)
│   │   ├── common_schema/
│   │   ├── flow_schema/
│   │   ├── generator_schema/
│   │   ├── monitor_schema/
│   │   ├── setting_schema/
│   │   ├── system_schema/
│   │   ├── system_corn_schema.go
│   │   └── user_protocol_schema.go
│   ├── model/                    # 数据库模型层 (Entity)
│   │   ├── common_model/
│   │   ├── gen_model/
│   │   ├── setting_model/
│   │   ├── system_model/
│   │   ├── flow_*.go             # 流程相关模型
│   │   ├── monitor_*.go          # 监控相关模型
│   │   ├── system_corn.go        # 定时任务模型
│   │   └── user_protocol.go      # 用户协议模型
│   ├── middleware/               # 中间件
│   │   ├── auth.go               # 认证中间件 (Token + 权限校验)
│   │   ├── cors.go               # 跨域中间件
│   │   ├── error.go              # 错误处理中间件
│   │   ├── log.go                # 操作日志中间件
│   │   └── ratelimit.go         # 限流中间件
│   └── corn/                     # 定时任务注册目录
├── config/                       # 配置管理
│   ├── init.go                   # 配置初始化 (Viper)
│   ├── app.go                    # 应用配置 (端口、模式等)
│   ├── admin.go                  # 管理员配置 (Token、权限等)
│   ├── db.go                     # 数据库配置
│   ├── redis.go                  # Redis 配置
│   ├── file.go                   # 文件上传配置
│   ├── email.go                  # 邮件配置
│   ├── geTui.go                  # 个推配置
│   ├── log.go                    # 日志配置
│   └── constant.go               # 常量定义
├── core/                         # 核心组件
│   ├── db.go                     # GORM 数据库初始化
│   ├── redis.go                  # Redis 连接初始化
│   ├── logger.go                 # Zap 日志初始化
│   ├── ws.go                     # WebSocket 管理
│   ├── request/common.go         # 请求封装
│   ├── response/                 # 统一响应封装
│   │   ├── common.go            # 分页响应 / Ok / Fail 快捷函数
│   │   ├── response.go          # Response 结构体
│   │   └── error.go             # 错误码定义 (RespType 接口)
├── routes/                       # 路由注册
│   ├── route.go                  # Gin 引擎初始化 & 根路由
│   ├── api.go                    # /api 入口路由
│   ├── admin_route/              # /api/admin 后台路由组
│   │   ├── route.go             # 主路由注册 (含 Autoload 自动加载)
│   │   ├── flow_route.go
│   │   ├── monitor_client_route.go
│   │   ├── monitor_error_route.go
│   │   ├── monitor_project_route.go
│   │   ├── system_corn_route.go
│   │   └── user_protocol_route.go
│   └── web_route/                # /api/web 前端路由组 (预留)
├── util/                         # 工具包
│   ├── tools.go                  # 通用工具函数 (JSON转换、包含判断等)
│   ├── verify.go                 # 校验工具
│   ├── string.go                 # 字符串处理
│   ├── array.go                  # 数组处理
│   ├── cache.go                  # 缓存操作
│   ├── redis.go                  # Redis 封装
│   ├── redisLock.go              # Redis 分布式锁
│   ├── ip.go                     # IP 解析 (ip2region)
│   ├── url.go                    # URL 处理
│   ├── email.go                  # 邮件发送
│   ├── server.go                 # 服务器信息
│   ├── uaparser.go              # UA 解析
│   ├── aj-captcha-go/           # 行为验证码 (滑块验证码、文字点选验证码)
│   ├── convert_util/            # 类型转换
│   ├── excel2/                  # Excel 导入导出
│   ├── img_util/                # 图片处理
│   └── ws_util/                 # WebSocket 工具
├── plugin/                       # GORM 插件 (Null 类型 Validator)
├── public/                       # 公共资源
│   ├── resources/                # 验证码字体等依赖
│   └── static/                   # 静态文件 (/api/static/*)
│       └── api/index.html       # Swagger API 文档
├── main.go                       # 程序入口
├── go.mod                        # Go 模块依赖
├── pack.bat                      # 打包脚本
└── upx.exe                       # 可执行文件压缩工具
```

## 架构设计

### 分层架构 (MVC + Service)

```
请求 → Routes → Middleware → Controller → Service → Model(ORM) → MySQL
                              ↓
                        Schema (DTO校验)
```

各层职责：

| 层 | 目录 | 职责 |
|----|------|------|
| **Controller** | `app/controller` | 接收请求、参数绑定、调用 Service、返回响应 |
| **Service** | `app/service` | 业务逻辑处理、事务管理 |
| **Schema** | `app/schema` | 请求/响应结构体定义、参数校验标签 |
| **Model** | `app/model` | GORM 数据库模型定义、表关联 |
| **Middleware** | `app/middleware` | 认鉴权、CORS、日志、限流等横切关注点 |

### 认证机制

- **认证方式**: Token (存储于 Redis)
- **获取方式**: Header `token` 或 Query Param `token`
- **自动续签**: Token 剩余有效期 < 30 分钟时自动续签
- **两种中间件**:
  - `LoginAuth()`: 仅校验 Token 有效性
  - `PermAuth()`: 校验 Token + 接口权限（路由转权限标识，如 `/api/admin/user/list` → `admin:user:list`）
- **超级管理员**: 跳过权限检查 (`config.AdminConfig.SuperAdminId`)

### 统一响应格式

```json
{
  "code": 200,
  "message": "成功",
  "data": {}
}
```

**分页响应** (`PageResp`):
```json
{
  "code": 200,
  "message": "成功",
  "data": {
    "count": 100,
    "pageNo": 1,
    "pageSize": 15,
    "lists": []
  }
}
```

### 错误码约定

| code | 含义 |
|------|------|
| 200 | 成功 |
| 300 | 业务失败 (Fail) |
| 其他 | 具体错误 (response/error.go 中 RespType 定义) |

### 配置管理

使用 **Viper** 读取 YAML 配置文件，支持命令行参数 `-env` 指定配置文件路径：

```bash
# 默认读取 .env.yaml
go run main.go
# 指定配置文件
go run main.go -env=.env.prod.yaml
```

**配置模块**: `APP` | `DB` | `REDIS` | `FILE` | `GeTui` | `Email` | `Log`

### 数据库配置

- **表前缀**: 通过 `config.DBConfig.TablePrefix` 配置
- **命名策略**: SingularTable (单数表名, `User` → `user`)
- **连接池**: 可配置 MaxIdleConns / MaxOpenConns / ConnMaxLifetime
- **日志**: 支持配置慢 SQL 阈值和日志级别
- **软删除**: 使用 `gorm.io/plugin/soft_delete`

### 路由注册机制

支持 **自动加载** 模式 — 各子路由文件通过 `init()` 函数将注册函数追加到全局 `routeHandlers` 切片，主路由文件调用 `Autoload()` 统一执行。

```go
// 子路由文件中
func init() {
    routeHandlers = append(routeHandlers, myRouteHandler)
}
```

## 业务功能模块

### 系统管理 (systemController)
- 管理员登录 / 信息管理
- 角色 / 权限管理 (RBAC)
- 菜单管理
- 部门管理
- 岗位管理
- 操作日志

### 系统设置 (settingController)
- 网站设置
- 版权设置
- 字典管理 (字典类型 + 字典数据)

### 代码生成器 (generatorController)
- 从数据库表生成 CRUD 代码
- 支持模板自定义 (`generatorService/*.tpl`)
- 包含 Controller / Service / Schema / Model / Route 全套生成

### 系统监控 (monitorController)
- 项目列表 / 客户端管理
- 错误日志收集与分析

### 审批流程 (flowController)
- 流程模板配置
- 流程申请 / 审核
- 流程历史记录

### 通用功能 (commonController)
- 文件上传 (单文件 / 分片上传)
- 相册管理
- 后台首页统计数据
- 个推消息推送

### 其他
- **定时任务**: 基于 `robfig/cron`，注册入口在 `app/corn/`
- **用户协议**: 协议内容管理
- **WebSocket**: 支持 WS 连接 (`core/ws.go`, `controller/ws.go`)

## 常用命令

```bash
# 运行项目
go run .

# 格式化 Swagger 注释
swag fmt

# 生成 Swagger 文档
swag init
```

## 关键开发约定

1. **新增业务模块**需创建: Controller / Service / Schema / Model / Route 五件套
2. **路由注册**: 在对应 route 文件中使用 `init()` + `routeHandlers` 追加，或直接在 `admin_route/route.go` 的 `RegisterRoute` 中注册
3. **中间件选择**: 登录即可访问的接口用 `LoginAuth()`，需要权限校验的用 `PermAuth()`
4. **响应统一**: 使用 `response.Ok(c, data)` / `response.Fail(c, msg)` 快捷函数
5. **日志**: 使用 `core.Logger` (zap SugaredLogger)，支持控制台 + 文件双输出
6. **数据库操作**: 通过 `core.GetDB()` 获取 *gorm.DB 实例
7. **Redis 操作**: 封装在 `util.RedisUtil` 中
8. **参数校验**: 在 Schema 结构体中使用 validator 标签
