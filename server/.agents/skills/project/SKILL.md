---
name: "项目目录文件结构"
description: "x_admin 后台管理系统 - Go/Gin/GORM 项目完整技术文档"
---

# x_admin 项目技术文档

## 项目概述

x_admin 是一个基于 **Go 1.26** 的后台管理系统后端服务，采用前后端分离架构，同时提供 C 端用户系统（JWT 鉴权、邮箱/手机/微信登录）。

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
| 微信 SDK | PowerWeChat/v3 | v3.4.41 |
| 验证码 | aj-captcha-go (自研) | - |

## 目录结构

```
server/
├── app/                              # 应用程序核心代码
│   ├── controller/                   # 控制器层
│   │   ├── admin_ctl/                # 后台管理控制器
│   │   │   ├── common_controller/    # 通用功能
│   │   │   │   ├── album.go          #   相册管理
│   │   │   │   ├── captcha.go        #   验证码
│   │   │   │   ├── ge_tui.go         #   个推推送
│   │   │   │   ├── index.go          #   后台首页统计
│   │   │   │   ├── s3_upload.go      #   S3 风格直传
│   │   │   │   └── upload.go         #   文件上传
│   │   │   ├── flow_controller/      # 审批流程
│   │   │   │   ├── flow_apply_ctl.go
│   │   │   │   ├── flow_history_ctl.go
│   │   │   │   └── flow_template_ctl.go
│   │   │   ├── generator_controller/ # 代码生成器
│   │   │   │   └── generator.go
│   │   │   ├── monitor_controller/   # 系统监控
│   │   │   │   ├── monitor.go
│   │   │   │   ├── monitor_client_ctl.go
│   │   │   │   ├── monitor_error_ctl.go
│   │   │   │   └── monitor_project_ctl.go
│   │   │   ├── setting_controller/   # 系统设置
│   │   │   │   ├── copyright.go
│   │   │   │   ├── dict_data.go
│   │   │   │   ├── dict_type.go
│   │   │   │   └── website.go
│   │   │   ├── system_controller/    # 系统管理
│   │   │   │   ├── admin.go          #   管理员
│   │   │   │   ├── dept.go           #   部门
│   │   │   │   ├── log.go            #   操作日志
│   │   │   │   ├── login.go          #   登录
│   │   │   │   ├── menu.go           #   菜单
│   │   │   │   ├── notice.go         #   通知
│   │   │   │   ├── post.go           #   岗位
│   │   │   │   └── role.go           #   角色
│   │   │   ├── system_corn_ctl.go    # 定时任务管理
│   │   │   └── user_protocol_ctl.go  # 用户协议
│   │   ├── user_ctl/                 # C端用户控制器
│   │   │   ├── auth_controller.go    #   绑定/解绑（手机、邮箱、微信）
│   │   │   └── user_controller.go    #   注册/登录/手机验证码/密码重置
│   │   └── ws.go                     # WebSocket 升级处理器
│   │
│   ├── service/                      # 服务层（业务逻辑）
│   │   ├── common_service/           # 通用服务
│   │   │   ├── album_service.go      #   相册管理
│   │   │   ├── captcha_service.go    #   验证码
│   │   │   ├── file_hash_service.go  #   文件哈希（秒传）
│   │   │   ├── file_ref_service.go   #   文件关联引用管理
│   │   │   ├── file_cleanup_service.go#  文件定时清理（7天无引用）
│   │   │   ├── ge_tui_service.go     #   个推推送（两级缓存）
│   │   │   ├── index_service.go      #   首页统计
│   │   │   └── upload_service.go     #   文件上传
│   │   ├── corn_service/             # 定时任务服务
│   │   │   ├── system_corn_service.go
│   │   │   └── task.go               #   Task 结构定义
│   │   ├── flow_service/             # 审批流程服务
│   │   │   ├── flow_apply_service.go
│   │   │   ├── flow_history_service.go
│   │   │   └── flow_template_service.go
│   │   ├── generator_service/        # 代码生成器服务
│   │   │   ├── generator_service.go
│   │   │   └── tpl_utils/            #   模板工具
│   │   │       ├── constants.go
│   │   │       ├── tpl.go
│   │   │       ├── utils.go
│   │   │       └── templates/        #   代码模板
│   │   │           ├── gocode/       #     Go 模板 (controller/model/route/schema/service/sql.tpl)
│   │   │           ├── vue/          #     Vue 模板 (api/details/edit/index.tpl)
│   │   │           └── uniapp/       #     UniApp 模板 (api/details/edit/index.tpl)
│   │   ├── monitor_service/          # 监控服务
│   │   │   ├── monitor_client_service.go
│   │   │   ├── monitor_error_list_service.go
│   │   │   ├── monitor_error_service.go
│   │   │   ├── monitor_project_service.go
│   │   │   └── monitor_server_service.go
│   │   ├── notice_service/           # 通知服务
│   │   │   └── notice_service.go     #   站内通知 + 邮件延迟补推
│   │   ├── setting_service/          # 设置服务
│   │   │   ├── copyright_service.go
│   │   │   ├── dict_data_service.go
│   │   │   ├── dict_type_service.go
│   │   │   ├── system_config_service.go
│   │   │   └── website_service.go
│   │   ├── system_service/           # 系统管理服务
│   │   │   ├── admin_service.go      #   管理员 CRUD
│   │   │   ├── admin_role_service.go #   管理员-角色关联
│   │   │   ├── dept_service.go       #   部门
│   │   │   ├── forget_pwd_service.go #   忘记密码
│   │   │   ├── log_service.go        #   操作日志
│   │   │   ├── login_service.go      #   登录认证
│   │   │   ├── menu_service.go       #   菜单
│   │   │   ├── perm_service.go       #   权限
│   │   │   ├── post_service.go       #   岗位
│   │   │   └── role_service.go       #   角色
│   │   ├── user_service/             # C端用户服务
│   │   │   ├── auth_service.go       #   绑定/解绑手机、邮箱
│   │   │   ├── user_service.go       #   注册/登录/手机验证码/密码重置
│   │   │   └── wechat_service.go     #   微信小程序/公众号登录绑定
│   │   └── user_protocol_service.go  # 用户协议
│   │
│   ├── schema/                       # 请求与响应结构体 (DTO)
│   │   ├── common_schema/
│   │   │   ├── album_schema.go
│   │   │   ├── captcha_schema.go
│   │   │   └── file_hash_schema.go
│   │   ├── flow_schema/
│   │   │   ├── flow_apply_schema.go
│   │   │   ├── flow_history_schema.go
│   │   │   └── flow_template_schema.go
│   │   ├── generator_schema/
│   │   │   └── schema.go
│   │   ├── monitor_schema/
│   │   │   ├── monitor_client_schema.go
│   │   │   ├── monitor_error_list_schema.go
│   │   │   ├── monitor_error_schema.go
│   │   │   └── monitor_project_schema.go
│   │   ├── setting_schema/
│   │   │   ├── copyright_schema.go
│   │   │   ├── dict_data_schema.go
│   │   │   ├── dict_type_schema.go
│   │   │   └── website_schema.go
│   │   ├── system_schema/
│   │   │   ├── admin_schema.go
│   │   │   ├── dept_schema.go
│   │   │   ├── log_schema.go
│   │   │   ├── login_schema.go
│   │   │   ├── menu_schema.go
│   │   │   ├── notice_schema.go
│   │   │   ├── post_schema.go
│   │   │   └── role_schema.go
│   │   ├── user_schema/              # C端用户 DTO
│   │   │   ├── auth.go               #   绑定/解绑请求
│   │   │   ├── email_code.go         #   邮箱/短信验证码
│   │   │   ├── login.go              #   登录请求/响应 (含 LoginResp.IsNew)
│   │   │   ├── register.go           #   注册请求
│   │   │   ├── user.go               #   用户信息
│   │   │   └── wechat.go             #   微信登录/绑定请求
│   │   ├── system_corn_schema.go
│   │   └── user_protocol_schema.go
│   │
│   ├── model/                        # 数据库模型层
│   │   ├── common_model/
│   │   │   ├── album.go              #   相册
│   │   │   ├── file_hash.go          #   文件哈希（秒传）
│   │   │   └── file_ref.go           #   文件关联引用
│   │   ├── gen_model/
│   │   │   └── gen.go                #   代码生成器元数据
│   │   ├── setting_model/
│   │   │   ├── dict_data.go
│   │   │   └── dict_type.go
│   │   ├── system_model/
│   │   │   ├── admin_role.go         #   管理员-角色关联
│   │   │   ├── auth_admin.go         #   管理员
│   │   │   └── system.go             #   系统配置/角色/菜单/部门/岗位/日志
│   │   ├── user_model/               # C端用户模型
│   │   │   ├── user.go               #   用户主表
│   │   │   └── user_auth.go          #   第三方认证绑定
│   │   ├── flow_apply.go
│   │   ├── flow_history.go
│   │   ├── flow_template.go
│   │   ├── monitor_client.go
│   │   ├── monitor_error.go
│   │   ├── monitor_error_list.go
│   │   ├── monitor_project.go
│   │   ├── system_corn.go
│   │   ├── system_notice.go
│   │   └── user_protocol.go
│   │
│   ├── middleware/                    # 中间件
│   │   ├── auth.go                   #   后台认证 (Token + 权限校验)
│   │   ├── user_auth.go              #   C端用户 JWT 认证
│   │   ├── cors.go                   #   跨域
│   │   ├── error.go                  #   Panic 恢复
│   │   ├── log.go                    #   操作日志
│   │   ├── limit_email.go            #   邮箱限频
│   │   ├── limit_ip.go               #   IP 限频
│   │   └── limit_rate.go             #   通用限流
│   │
│   └── corn/                         # 定时任务注册
│       ├── corn_manager.go           #   CronManager (基于 robfig/cron)
│       ├── fixed_task.go             #   固定任务（清理分片/在线人数/监控/文件清理等）
│       └── dynamic_task.go           #   动态任务（从数据库加载）
│
├── config/                           # 配置管理
│   ├── init.go                       #   Viper 初始化入口
│   ├── app.go                        #   应用配置 (端口/模式/名称)
│   ├── admin.go                      #   管理员配置 (Token/权限/超级管理员)
│   ├── db.go                         #   数据库配置
│   ├── redis.go                      #   Redis 配置
│   ├── jwt.go                        #   JWT 配置 (密钥/过期时间)
│   ├── file.go                       #   文件上传配置 (大小限制/分片/预签名)
│   ├── email.go                      #   邮件配置
│   ├── wechat.go                     #   微信配置 (小程序/公众号 AppID+Secret)
│   ├── ge_tui.go                     #   个推配置
│   ├── notice.go                     #   通知配置
│   ├── log.go                        #   日志配置
│   └── constant.go                   #   常量定义
│
├── core/                             # 核心组件
│   ├── db.go                         #   GORM 数据库初始化
│   ├── redis.go                      #   Redis 连接初始化
│   ├── logger.go                     #   Zap 日志初始化
│   ├── ws.go                         #   WebSocket 全局管理器入口
│   ├── ws/                           #   WebSocket 实现
│   │   ├── client.go                 #     客户端（sync.Once 安全关闭）
│   │   └── manager.go               #     连接管理器（单推/群推/全推）
│   ├── request/
│   │   └── common.go                 #   分页请求 PageReq
│   └── response/                     #   统一响应封装
│       ├── common.go                 #     Ok/Fail/NoRoute 快捷函数
│       ├── response.go               #     Response/PageResp 结构体
│       └── error.go                  #     错误码定义 (RespType 接口)
│
├── routes/                           # 路由注册
│   ├── init.go                       #   Gin 引擎初始化
│   ├── api.go                        #   /api 入口路由注册
│   ├── admin_route/                  #   /api/admin 后台路由组
│   │   ├── init.go                   #     Autoload 自动加载机制
│   │   ├── common_route.go           #     通用：上传/相册/个推
│   │   ├── flow_route.go             #     审批流程
│   │   ├── generator_route.go        #     代码生成器
│   │   ├── monitor_route.go          #     系统监控
│   │   ├── notice_route.go           #     通知
│   │   ├── s3_route.go               #     S3 风格直传
│   │   ├── setting_route.go          #     系统设置
│   │   ├── system_route.go           #     系统管理
│   │   ├── system_corn_route.go      #     定时任务
│   │   └── user_protocol_route.go    #     用户协议
│   └── user_route/                   #   /api/user C端用户路由
│       └── user_route.go             #     注册/登录/绑定/微信
│
├── util/                             # 工具包
│   ├── tools.go                      #   通用工具 (UUID/JSON/MD5/随机等)
│   ├── verify.go                     #   校验工具
│   ├── string.go                     #   字符串处理
│   ├── array.go                      #   数组处理
│   ├── cache.go                      #   缓存操作
│   ├── redis.go                      #   Redis 封装 (Set/Get/Hash/List 等)
│   ├── redis_lock.go                 #   Redis 分布式锁 (Lua 原子解锁)
│   ├── jwt.go                        #   JWT 生成/解析/刷新
│   ├── email.go                      #   邮件发送
│   ├── email_code.go                 #   邮箱验证码 (Redis 存储+限频)
│   ├── sms_code.go                   #   短信验证码 (Redis 存储+限频, TODO 对接服务商)
│   ├── wechat.go                     #   PowerWeChat 客户端初始化 (sync.Once)
│   ├── ip.go                         #   IP 解析 (ip2region)
│   ├── url.go                        #   URL 处理 (绝对路径转换)
│   ├── server.go                     #   服务器信息采集
│   ├── uaparser.go                   #   UA 解析
│   ├── null_time.go                  #   Null 时间类型工具
│   ├── convert_util/                 #   类型转换 (struct copy)
│   ├── excel2/                       #   Excel 导入导出
│   ├── img_util/                     #   图片处理
│   └── aj-captcha-go/               #   行为验证码
│       ├── captcha_config/           #     配置与常量
│       ├── captcha_service/          #     滑块/点选验证码实现
│       ├── model/vo/                 #     值对象
│       └── util/                     #     AES/图片/字体/缓存工具
│
├── plugin/                           # 插件
│   ├── storage.go                    #   文件上传驱动
│   ├── storage_interface.go          #   存储引擎接口 (含分片上传)
│   ├── storage_local.go              #   本地文件存储实现
│   ├── storage_factory.go            #   存储引擎工厂
│   ├── presign.go                    #   S3 预签名 URL 生成
│   ├── cron_cleanup.go              #   分片临时目录定时清理
│   └── validator_null.go            #   GORM Null 类型 Validator 插件
│
├── public/                           # 公共资源
│   ├── resources/                    #   验证码字体等
│   └── static/                       #   静态文件 (/api/static/*)
│       └── api/index.html            #     Swagger API 文档
│
├── docs/                             # 文档
│   ├── docs.go                       #   Swagger 生成代码
│   ├── swagger.json
│   ├── swagger.yaml
│   └── migration/                    #   数据库迁移 SQL
│       └── add_common_file_ref_table.sql
│
├── main.go                           # 程序入口
├── go.mod                            # Go 模块依赖
├── .env.yaml                         # 本地配置 (不提交 git)
├── .env.yaml.example                 # 配置模板
├── pack.bat                          # 打包脚本
└── upx.exe                           # 可执行文件压缩
```

## 架构设计

### 分层架构 (MVC + Service)

```
请求 → Routes → Middleware → Controller → Service → Model(ORM) → MySQL
                              ↓
                        Schema (DTO校验)
```

| 层 | 目录 | 职责 |
|----|------|------|
| **Controller** | `app/controller` | 接收请求、参数绑定、调用 Service、返回响应 |
| **Service** | `app/service` | 业务逻辑处理、事务管理 |
| **Schema** | `app/schema` | 请求/响应结构体定义、参数校验标签 |
| **Model** | `app/model` | GORM 数据库模型定义、表关联 |
| **Middleware** | `app/middleware` | 认证鉴权、CORS、日志、限流等横切关注点 |

### 双认证体系

| 体系 | 中间件 | 适用场景 |
|------|--------|----------|
| **后台管理** | `auth.go` → `LoginAuth()` / `PermAuth()` | Token 存 Redis，自动续签，权限校验 |
| **C端用户** | `user_auth.go` → `UserJWTAuth()` | JWT 双 Token (access+refresh)，token_version 踢人下线 |

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

### 配置管理

使用 **Viper** 读取 YAML 配置文件：

```bash
go run main.go                  # 默认 .env.yaml
go run main.go -env=.env.prod.yaml  # 指定环境
```

**配置模块**: `APP` | `DB` | `REDIS` | `JWT` | `FILE` | `WECHAT` | `GeTui` | `Email` | `Notice` | `Log`

### 路由注册机制

支持 **自动加载** 模式 — 子路由文件通过 `init()` 将注册函数追加到全局 `routeHandlers`，由 `admin_route/init.go` 的 `Autoload()` 统一执行。

## 业务功能模块

### 后台管理系统

| 模块 | Controller | 功能 |
|------|------------|------|
| 系统管理 | `system_controller/` | 管理员/角色/菜单/部门/岗位/操作日志 |
| 系统设置 | `setting_controller/` | 网站设置/版权/字典管理 |
| 代码生成器 | `generator_controller/` | 从数据库表生成 CRUD（Go/Vue/UniApp） |
| 系统监控 | `monitor_controller/` | 项目/客户端/错误日志收集与分析 |
| 审批流程 | `flow_controller/` | 流程模板/申请/审核/历史 |
| 通用功能 | `common_controller/` | 文件上传(含S3直传)/相册/个推/首页统计 |
| 定时任务 | `system_corn_ctl.go` | 动态定时任务管理 |
| 用户协议 | `user_protocol_ctl.go` | 协议内容管理 |

### C端用户系统

| 功能 | 路由 | 说明 |
|------|------|------|
| 邮箱注册/登录 | `/api/user/register`, `/api/user/login` | bcrypt + salt |
| 手机号+密码登录 | `/api/user/phoneLogin` | 手机号非必填 |
| 手机号+验证码登录 | `/api/user/phoneCodeLogin` | 短信验证码 (Redis 限频) |
| 微信小程序登录 | `/api/user/wechatMiniLogin` | PowerWeChat SDK，自动注册 |
| 微信公众号登录 | `/api/user/wechatMpLogin` | OAuth 网页授权 |
| 绑定/解绑 | `/api/user/bind*`, `/api/user/unbind*` | 手机(需短信码)/邮箱(需验证码)/微信 |
| 密码重置 | `/api/user/resetPassword`, `/api/user/resetPhonePassword` | 邮箱验证码/短信验证码 |
| 踢人下线 | token_version 自增 | 使旧 JWT 失效 |

## 关键开发约定

1. **新增业务模块**需创建: Controller / Service / Schema / Model / Route 五件套
2. **路由注册**: 在对应 route 文件中 `init()` + `routeHandlers` 追加
3. **中间件选择**: 后台接口用 `LoginAuth()`/`PermAuth()`，C端接口用 `UserJWTAuth()`
4. **响应统一**: 使用 `response.Ok(c, data)` / `response.Fail(c, msg)`
5. **日志**: 使用 `core.Logger` (zap SugaredLogger)
6. **数据库**: 通过 `core.GetDB()` 获取 `*gorm.DB`
7. **Redis**: 封装在 `util.RedisUtil`，分布式锁用 `util.NewRedisLock()`
8. **参数校验**: Schema 结构体中使用 `binding` + `validator` 标签
9. **文件关联**: 业务使用文件后调用 `FileRefService.AddRef()` 建立引用，定时清理无引用文件
10. **主键**: UUID v7 (`char(36)`)，模型中通过 `BeforeCreate` 钩子自动生成
