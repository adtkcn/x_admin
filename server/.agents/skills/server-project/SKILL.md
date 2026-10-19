---
name: "后端导航与项目技术文档"
description: "x_admin 后端导航入口 + 完整技术文档 — 后端有哪些东西（路由域/业务模块/基础设施），含技术栈、目录结构、架构与子 skills 导航"
---

# x_admin 项目技术文档

> 导航：本页既是后端 skills 的导航索引，也是后端详细文档。子 skills 见下方「后端 Skills 导航」，仓库整体情况见 [项目总导航](../../../../.agents/skills/project-navigation/SKILL.md)。

## 项目概述

x_admin 是一个基于 **Go 1.27** 的后台管理系统后端服务，采用前后端分离架构，同时提供 C 端(Web)用户系统（JWT 鉴权、邮箱/手机/微信登录），并包含基于消息队列（`core/queue`）的异步任务处理。

**模块名**: `x_admin`  
**入口文件**: `main.go`  
**配置文件**: `.env.yaml`（不提交 git）  
**API 文档**: `/api/static/api/index.html`（Swagger）

## 后端导航

### 路由域（对外 API）

| 路由前缀 | 路由目录 | 控制器目录 | 说明 |
|----------|----------|------------|------|
| `/api/admin` | `routes/admin_route/` | `app/controller/admin_ctl/` | 后台管理接口（`LoginAuth`/`PermAuth` 鉴权） |
| `/api/web` | `routes/web_route/` | `app/controller/web_ctl/` | C端用户接口（JWT `UserLoginAuth` 鉴权） |
| `/api/iot` | `routes/iot_route/` | `app/controller/iot_ctl/` | IoT 物联网接口（设备/产品/区域/规则链/告警等） |
| `/api/static` | — | — | Swagger 文档与静态资源 |

### 基础设施（`core/`、`config/`、`util/`、`plugin/`）

- **core/**：GORM / Redis / Zap 初始化，`queue`（Redis 消息队列）、`ordered_queue`（分片有序队列）、`pubsub`（跨实例广播）、`ws`（WebSocket）、`iot_bus`（IoT 数据总线）
- **config/**：Viper 读取 `.env.yaml`，按 `APP/DB/REDIS/JWT/FILE/WECHAT/...` 分模块
- **util/**：JWT、Redis 封装与分布式锁、邮箱/短信验证码、Excel 导入导出、类型转换 `convert_util` 等
- **plugin/**：ASR 语音识别、微信、GORM Null 类型 Validator
- **app/queue/**：异步队列消费者统一注册启动；**app/corn/**：固定 + 动态定时任务

（逐文件说明见下方「目录结构」，业务模块清单见「业务功能模块」）

### 后端 Skills 导航

| Skill | 什么时候用 |
|-------|-----------|
| [add-api](../add-api/SKILL.md) | 在已有模块新增接口：Schema / Service / Controller / Route 四件套模板与权限标识规则 |
| [user-system](../user-system/SKILL.md) | C端用户系统：JWT 双 token、邮箱/手机/微信登录、验证码、踢人下线、数据库设计 |
| [wechatpay-payment-integration](../wechatpay-payment-integration/SKILL.md) | 微信支付接入 / 产品选型 / 排障 / 上线前质量检查 |

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
│   │   ├── ws.go                     # WebSocket 升级处理器
│   │   ├── admin_ctl/                # 后台管理控制器 (/api/admin)
│   │   │   ├── common_controller/    #   通用功能
│   │   │   │   ├── album.go          #     相册管理
│   │   │   │   ├── captcha.go        #     验证码
│   │   │   │   ├── ge_tui.go         #     个推推送
│   │   │   │   ├── index.go          #     后台首页统计
│   │   │   │   ├── upload.go         #     文件上传
│   │   │   │   └── upload_chunk.go   #     分片上传（直传存储引擎 plugin/storage）
│   │   │   ├── fabu_controller/      #   App 发布管理
│   │   │   │   ├── fabu_app_ctl.go
│   │   │   │   ├── fabu_version_ctl.go
│   │   │   │   └── fabu_wgt_ctl.go   #     WGT 热更包
│   │   │   ├── flow_controller/      #   审批流程
│   │   │   │   ├── flow_apply_ctl.go
│   │   │   │   ├── flow_history_ctl.go
│   │   │   │   └── flow_template_ctl.go
│   │   │   ├── generator_controller/ #   代码生成器
│   │   │   │   └── generator.go
│   │   │   ├── iot_controller/       #   IoT 后台管理
│   │   │   │   ├── product.go        #     产品（物模型）
│   │   │   │   ├── product_category.go #   产品分类
│   │   │   │   ├── device.go         #     设备
│   │   │   │   ├── device_space.go   #     设备空间/区域
│   │   │   │   ├── rule_chain.go     #     规则链（可视化编排）
│   │   │   │   └── alert.go          #     告警管理
│   │   │   ├── monitor_controller/   #   系统监控
│   │   │   │   ├── monitor.go
│   │   │   │   ├── monitor_client_ctl.go
│   │   │   │   ├── monitor_error_ctl.go
│   │   │   │   └── monitor_project_ctl.go
│   │   │   ├── setting_controller/   #   系统设置
│   │   │   │   ├── copyright.go
│   │   │   │   ├── dict_data.go
│   │   │   │   ├── dict_type.go
│   │   │   │   └── website.go
│   │   │   ├── system_controller/    #   系统管理
│   │   │   │   ├── admin.go          #     管理员
│   │   │   │   ├── dept.go           #     部门
│   │   │   │   ├── log.go            #     操作日志
│   │   │   │   ├── login.go          #     登录
│   │   │   │   ├── menu.go           #     菜单
│   │   │   │   ├── notice.go         #     通知
│   │   │   │   ├── post.go           #     岗位
│   │   │   │   ├── role.go           #     角色
│   │   │   │   └── system_corn_ctl.go #    定时任务管理
│   │   │   └── user_controller/      #   C端用户管理
│   │   │       ├── user_ctl.go       #     用户列表
│   │   │       └── user_protocol_ctl.go #  用户协议
│   │   ├── iot_ctl/                  # IoT 设备 HTTP 接入网关 (/api/iot，DeviceAuth)
│   │   │   └── http_gateway.go       #   遥测上报/心跳
│   │   └── web_ctl/                  # C端(Web)用户控制器 (/api/web)
│   │       ├── auth_controller.go    #   绑定/解绑（手机、邮箱、微信）
│   │       ├── fabu_controller.go    #   App 发布下载/更新检测
│   │       └── user_controller.go    #   注册/登录/手机验证码/密码重置
│   │
│   ├── service/                      # 服务层（业务逻辑）
│   │   ├── common_service/           #   通用服务（相册/验证码/文件哈希秒传/个推/首页/上传）
│   │   ├── corn_service/             #   定时任务服务
│   │   │   ├── system_corn_service.go
│   │   │   └── task.go               #     Task 结构定义
│   │   ├── fabu_service/             #   App 发布服务
│   │   │   ├── fabu_app_service.go
│   │   │   ├── fabu_version_service.go
│   │   │   ├── fabu_wgt_service.go   #     WGT 热更包
│   │   │   └── fabu_plist_service.go #     iOS plist 安装清单
│   │   ├── flow_service/             #   审批流程服务（apply/history/template）
│   │   ├── generator_service/        #   代码生成器服务
│   │   │   ├── generator_service.go
│   │   │   └── tpl_utils/            #     模板工具
│   │   ├── iot_service/              #   IoT 核心服务
│   │   │   ├── product_service.go    #     产品/物模型
│   │   │   ├── product_category_service.go
│   │   │   ├── device_service.go     #     设备 CRUD/在线状态
│   │   │   ├── device_auth_service.go #    设备接入鉴权
│   │   │   ├── device_space_service.go #   设备空间/区域
│   │   │   ├── rule_chain_service.go #     规则链管理
│   │   │   ├── alert_service.go      #     告警状态机
│   │   │   ├── local_cache.go        #     IoT 本地缓存
│   │   │   ├── tree_filter.go        #     区域/分类树过滤（列表 IN 查询）
│   │   │   └── pipeline/             #     IoT 报文处理管道（消费 core/ordered_queue）
│   │   │       ├── pipeline.go       #       管道编排入口
│   │   │       ├── rulechain/        #       规则链引擎（goja 脚本解码/节点执行）
│   │   │       ├── store/            #       时序存储（write 写入/query 查询/clean 清理）
│   │   │       ├── alert/            #       告警处理（状态升级/恢复确认）
│   │   │       └── online/           #       在线维护（离线判定 ZSET 快/慢路径）
│   │   ├── monitor_service/          #   监控服务（client/error/project/server）
│   │   ├── notice_service/           #   站内通知 + 邮件延迟补推
│   │   ├── setting_service/          #   设置服务（网站/版权/字典/系统配置）
│   │   ├── system_service/           #   系统管理服务（管理员/角色/菜单/部门/岗位/登录/权限/日志）
│   │   └── user_service/             #   C端用户服务
│   │       ├── user_service.go       #     注册/登录/验证码/重置密码
│   │       ├── user_admin_service.go #     后台端用户管理
│   │       ├── user_protocol_service.go #  用户协议
│   │       ├── auth_service.go       #     绑定/解绑手机、邮箱
│   │       └── wechat_service.go     #     微信小程序/公众号登录绑定
│   │
│   ├── queue/                        # 异步队列消费者（基于 core/queue，package queue）
│   │   ├── start.go                  #   统一注册并启动所有消费者（由启动流程调用）
│   │   ├── iot.go                    #   IoT 报文队列消费入口（接入队列 → pipeline）
│   │   ├── email_code.go             #   邮箱验证码发送
│   │   ├── flow_notify_email.go      #   流程邮件通知
│   │   ├── flow_notify_webhook.go    #   流程 Webhook 通知
│   │   ├── image_webp.go             #   图片异步转 webp（上传后投递，转码回写记录）
│   │   ├── notice_email.go           #   通知邮件补推
│   │   ├── notice_email_delay.go     #   通知邮件延迟补推
│   │   └── operate_log.go            #   操作日志落库（由中间件投递，解耦写库）
│   │
│   ├── schema/                       # 请求与响应结构体 (DTO)，按域分子目录
│   │   ├── common_schema/            #   相册/验证码/文件哈希
│   │   ├── fabu_schema/              #   App 发布 DTO
│   │   ├── flow_schema/              #   审批流程 DTO
│   │   ├── generator_schema/         #   代码生成器 DTO
│   │   ├── iot_schema/               #   IoT DTO（product/device/device_space/rule_chain/alert/ingest 报文）
│   │   ├── monitor_schema/           #   监控 DTO
│   │   ├── queue_schema/             #   异步队列 payload + 队列名常量 (queue.go)
│   │   ├── setting_schema/           #   设置 DTO
│   │   ├── system_schema/            #   系统管理 DTO（含 system_corn_schema.go）
│   │   └── user_schema/              #   C端用户 DTO（登录/注册/绑定/微信/协议 user_protocol.go）
│   │
│   ├── model/                        # 数据库模型层（均按域分子目录）
│   │   ├── common_model/             #   相册、文件哈希（秒传）
│   │   ├── fabu_model/               #   App 发布（fabu_app/fabu_app_version/fabu_wgt）
│   │   ├── flow_model/               #   审批流程（apply/history/template）
│   │   ├── gen_model/                #   代码生成器元数据
│   │   ├── iot_model/                #   IoT（product/product_category/device/device_space/device_data/rule_chain/alert_history）
│   │   ├── monitor_model/            #   监控（client/error/error_list/project）
│   │   ├── setting_model/            #   字典（dict_data/dict_type）
│   │   ├── system_model/             #   系统（auth_admin/admin_role/system/system_corn/system_notice）
│   │   └── user_model/               #   C端用户（user/user_auth/user_protocol）
│   │
│   ├── middleware/                   # 中间件
│   │   ├── admin_auth.go             #   后台认证 (Token + 权限校验 LoginAuth/PermAuth)
│   │   ├── user_auth.go              #   C端用户 JWT 认证
│   │   ├── device_auth.go            #   IoT 设备接入鉴权 (/api/iot)
│   │   ├── cors.go                   #   跨域
│   │   ├── error.go                  #   Panic 恢复
│   │   ├── request_cost.go           #   请求耗时统计
│   │   ├── log.go                    #   操作日志
│   │   ├── limit_email.go            #   邮箱限频
│   │   ├── limit_ip.go               #   IP 限频
│   │   └── limit_rate.go             #   通用限流
│   │
│   └── corn/                         # 定时任务注册 (package corn)
│       ├── corn_manager.go           #   CronManager (基于 robfig/cron 秒级)
│       ├── fixed_task.go             #   固定任务（在线人数/监控/清理/补推等）
│       └── dynamic_task.go           #   动态任务（从数据库加载）
│
├── config/                           # 配置管理 (Viper)
│   ├── init.go                       #   Viper 初始化入口
│   ├── app.go / admin.go             #   应用配置(端口/模式) / 管理员配置(Token/超级管理员)
│   ├── db.go / redis.go              #   MySQL / Redis 配置
│   ├── jwt.go                        #   JWT 配置 (密钥/过期时间)
│   ├── file.go / oss.go              #   文件上传配置 / OSS 存储配置
│   ├── email.go / notice.go          #   邮件配置 / 通知配置
│   ├── iot.go                        #   IoT 配置（离线超时/TTL/日志开关等）
│   ├── wechat.go / ge_tui.go         #   微信配置 / 个推配置
│   ├── log.go / constant.go          #   日志配置 / 常量定义
│
├── core/                             # 核心组件
│   ├── db.go / redis.go / logger.go  #   GORM / Redis / Zap 初始化
│   ├── iot_bus.go                    #   IoT 数据总线（设备报文接入队列）
│   ├── local_cache.go                #   全局本地缓存（ttlcache）
│   ├── queue.go / queue/             #   全局消息队列实例及实现 (Redis)
│   ├── ordered_queue.go / ordered_queue/ # 分片有序队列（IoT 报文按设备顺序消费）
│   ├── queue_delay/                  #   延迟队列（Redis 实现）
│   ├── pubsub/                       #   发布订阅 (基于 Redis，跨实例广播)
│   ├── ws.go / ws/                   #   WebSocket 全局管理器及实现（client/manager/room）
│   ├── request/                      #   分页请求 PageReq
│   └── response/                     #   统一响应封装 (Ok/Fail/PageResp/错误码)
│
├── routes/                           # 路由注册
│   ├── init.go                       #   Gin 引擎初始化
│   ├── api.go                        #   /api 入口（captcha/ws + admin/web/iot 三个路由组）
│   ├── admin_route/                  #   /api/admin 后台路由（Autoload 自动追加）
│   │   ├── init.go                   #     Autoload 自动加载机制
│   │   ├── common_route.go           #     通用：上传/分片上传/相册/个推
│   │   ├── fabu_route.go             #     App 发布管理
│   │   ├── flow_route.go             #     审批流程
│   │   ├── generator_route.go        #     代码生成器
│   │   ├── iot_route.go              #     IoT 后台管理（产品/设备/规则链/告警…）
│   │   ├── monitor_route.go          #     系统监控
│   │   ├── notice_route.go           #     通知
│   │   ├── setting_route.go          #     系统设置
│   │   ├── system_route.go           #     系统管理
│   │   ├── system_corn_route.go      #     定时任务
│   │   ├── user_route.go             #     C端用户管理
│   │   └── user_protocol_route.go    #     用户协议
│   ├── web_route/                    #   /api/web C端(Web)用户路由
│   │   ├── init.go                   #     Autoload 自动加载机制
│   │   ├── user_route.go             #     注册/登录/绑定/微信
│   │   └── fabu_route.go             #     App 发布下载/更新检测
│   └── iot_route/                    #   /api/iot 设备 HTTP 接入网关
│       ├── init.go                   #     Autoload 自动加载机制
│       └── device_route.go           #     遥测/心跳（DeviceAuth 鉴权）
│
├── util/                             # 工具包
│   ├── tools.go / verify.go          #   通用工具(UUID/JSON/MD5) / 校验工具 VerifyUtil
│   ├── string.go / array.go          #   字符串 / 数组处理
│   ├── cache.go                      #   缓存操作
│   ├── redis.go / redis_lock.go      #   Redis 封装 / 分布式锁 (Lua 原子解锁)
│   ├── jwt.go                        #   JWT 生成/解析/刷新
│   ├── ip.go / url.go                #   IP 解析(ip2region) / URL 处理
│   ├── server.go / uaparser.go       #   服务器信息 / UA 解析
│   ├── null_time.go                  #   Null 时间类型工具
│   ├── convert_util/                 #   类型转换 (struct copy)
│   ├── excel_util/                   #   Excel 导入导出
│   ├── file_util/                    #   文件操作
│   ├── img_util/                     #   图片处理（转 webp 等）
│   └── validation/                   #   GORM Null 类型 Validator（原 plugin/validator_null）
│
├── plugin/                           # 插件
│   ├── aj-captcha-go/                #   行为验证码（原 util 下迁至插件：AES/滑块/点选/缓存）
│   ├── email/                        #   邮件发送 + 邮箱验证码 (code.go，Redis 存储+限频)
│   ├── sms/                          #   短信验证码 (code.go，Redis 限频，TODO 对接服务商)
│   ├── ipapk/                        #   APK/IPA 包元数据解析（App 发布用）
│   ├── storage/                      #   存储引擎（local/oss 分桶、预签名直传、定时清理）
│   └── wechat/                       #   微信相关插件 (PowerWeChat 客户端/消息/支付)
│
├── public/                           # 公共资源
│   ├── resources/                    #   验证码字体/底图、ip2region 库
│   └── static/api/                   #   Swagger API 文档 (/api/static/*)
│
├── docs/                             # Swagger 生成代码 (docs.go/swagger.json/yaml)
│
├── main.go                           # 程序入口
├── bootstrap.go                      # 启动初始化（配置/DB/Redis/队列等）
├── shutdown.go                       # 优雅关闭
├── go.mod                            # Go 模块依赖
├── .env.yaml                         # 本地配置 (不提交 git)
├── .env.yaml.example                 # 配置模板
├── .env.online.yaml                  # 线上配置
└── pack.bat                          # 打包脚本
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
| **后台管理** | `admin_auth.go` → `LoginAuth()` / `PermAuth()` | Token 存 Redis，自动续签，权限校验 |
| **C端用户** | `user_auth.go` → `UserLoginAuth()` | JWT 双 Token (access+refresh)，token_version 踢人下线 |
| **IoT 设备** | `device_auth.go` → `DeviceAuth()` | 设备接入网关 `/api/iot` 鉴权 |

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
| 定时任务 | `system_corn_ctl.go` | 动态定时任务管理 (corn 包注册) |
| 用户协议 | `user_protocol_ctl.go` | 协议内容管理 |
| IoT 物联网 | `iot_ctl/` + `routes/iot_route/` | 设备/产品/区域/物模型/规则链引擎/告警状态机/大屏（`iot_service/`） |
| 发布域 | `fabu_service/` | 发布相关业务（控制器位于 `admin_ctl/` 对应子目录） |

### C端(Web)用户系统

C端路由前缀为 **`/api/web`**（控制器在 `web_ctl/`，路由在 `web_route/`）。

| 功能 | 路由 | 说明 |
|------|------|------|
| 邮箱注册/登录 | `/api/web/register`, `/api/web/login` | bcrypt + salt |
| 手机号+密码登录 | `/api/web/phoneLogin` | 手机号非必填 |
| 手机号+验证码登录 | `/api/web/phoneCodeLogin` | 短信验证码 (Redis 限频) |
| 微信小程序登录 | `/api/web/wechatMiniLogin` | PowerWeChat SDK，自动注册 |
| 微信公众号登录 | `/api/web/wechatMpLogin` | OAuth 网页授权 |
| 绑定/解绑 | `/api/web/bind*`, `/api/web/unbind*` | 手机(需短信码)/邮箱(需验证码)/微信 |
| 密码重置 | `/api/web/resetPassword`, `/api/web/resetPhonePassword` | 邮箱验证码/短信验证码 |
| 踢人下线 | token_version 自增 | 使旧 JWT 失效 |

## 关键开发约定

1. **新增业务模块**需创建: Controller / Service / Schema / Model / Route 五件套
2. **路由注册**: 在对应 route 文件中 `init()` + `routeHandlers` 追加
3. **中间件选择**: 后台接口用 `LoginAuth()`/`PermAuth()`，C端接口用 `UserLoginAuth()`，IoT 设备接入用 `DeviceAuth()`
4. **响应统一**: 使用 `response.Ok(c, data)` / `response.Fail(c, msg)`
5. **日志**: 使用 `core.Logger` (zap SugaredLogger)
6. **数据库**: 通过 `core.GetDB()` 获取 `*gorm.DB`
7. **Redis**: 封装在 `util.RedisUtil`，分布式锁用 `util.NewRedisLock()`
8. **参数校验**: Schema 结构体中使用 `binding` + `validator` 标签
9. **异步任务**: 耗时操作（AI处理/流程通知/图片转webp/操作日志等）通过 `core/queue` 消息队列异步处理，消费者在 `app/queue/start.go` 统一注册与启动
10. **主键**: UUID v7 (`char(36)`)，模型中通过 `BeforeCreate` 钩子自动生成

## 相关文档导航

> 以下文档位于仓库 `docs/` 目录（VitePress 文档站），本 SKILL 已为其补充 `name` + `description` 格式的 frontmatter，可作为子 skill 直接检索。相对路径从本文件 `server/.agents/skills/project/` 出发。

### 后端基础

- [前期准备](../../../../docs/server/准备.md) — Go/MySQL/Redis 安装与项目初始化
- [项目结构说明](../../../../docs/server/结构说明.md) — 目录结构与模块划分
- [环境变量](../../../../docs/server/环境变量.md) — Viper 配置与 `.env.yaml`
- [权限验证](../../../../docs/server/权限验证.md) — `LoginAuth`/`PermAuth` 中间件与菜单权限标识
- [开发注意事项](../../../../docs/server/注意事项.md) — 循环依赖/零值/软删除等坑
- [数据库 Null 值](../../../../docs/server/数据库null值.md) — GORM Null 类型自定义
- [Swaggo 接口文档](../../../../docs/server/swaggo.md) — 接口文档生成
- [打包部署](../../../../docs/server/部署Go.md) — Go 跨平台打包
- [Nginx 配置](../../../../docs/server/nginx配置.md) — 反向代理与静态资源
- [Excel 导入导出](../../../../docs/server/导入导出excel.md) — 基于 excelize

### 后端功能说明

- [系统功能模块概述](../../../../docs/server/功能说明/模块概述.md) — 模块总览
- [工具函数与核心组件](../../../../docs/server/功能说明/工具函数与核心组件.md) — `util` 与 `core` 基础设施
- [异步队列](../../../../docs/server/功能说明/异步队列.md) — `core/queue` 消费者注册
- [定时任务](../../../../docs/server/功能说明/定时任务.md) — `robfig/cron` 固定/动态任务
- [事件总线 PubSub](../../../../docs/server/功能说明/事件总线pubsub.md) — `core/pubsub` 跨实例广播
- [存储引擎](../../../../docs/server/功能说明/存储引擎.md) — 本地/OSS 与预签名直传
- [邮件发送](../../../../docs/server/功能说明/邮件发送.md) — 多账号与异步补推
- [WebSocket 实时通信](../../../../docs/server/功能说明/WebSocket.md) — `core/ws` 连接管理
