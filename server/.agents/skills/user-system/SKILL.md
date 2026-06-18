---
name: "用户系统"
description: "C端用户系统完整技术文档 — JWT鉴权、邮箱/手机登录、验证码、踢人下线机制"
---

# 用户系统技术文档

## 概述

独立的 **C端用户体系**，与后台管理员系统 (`SystemAuthAdmin`) 完全隔离。  
邮箱为主账号，支持绑定手机号（可选）、微信（小程序/公众号/APP）、QQ 等第三方认证。  
采用 **JWT 无状态鉴权**，路由前缀 `/api/user/*`。

**核心特性：**
- JWT access_token + refresh_token 双 token 方案
- 邮箱注册/登录 + 手机号登录（密码/短信验证码两种方式）
- 邮箱验证码（注册/重置密码/解绑手机） + 短信验证码（绑定手机/手机登录/手机重置密码）
- token_version 机制实现踢人下线
- Redis 缓存 token_version，避免每次鉴权查 MySQL

---

## 文件清单

```
app/model/user_model/
├── user.go                          # User 主表（邮箱、密码、手机号、token_version）
└── user_auth.go                     # UserAuth 第三方绑定表（不含手机号）

app/schema/user_schema/
├── register.go                      # RegisterReq
├── login.go                         # LoginReq, PhoneLoginReq, PhoneCodeLoginReq, LoginResp
├── user.go                          # UserInfoResp, UpdateUserReq
├── auth.go                          # RefreshTokenReq, BindPhoneReq, UnbindPhoneReq, UserAuthItem
└── email_code.go                    # SendCodeReq, SendSmsCodeReq, ResetPasswordReq, ResetPhonePasswordReq

app/service/user_service/
├── user_service.go                  # 注册/邮箱登录/手机登录/短信登录/验证码/重置密码/踢人下线/token_version缓存
└── auth_service.go                  # 绑定手机号/解绑手机号/获取绑定列表

app/controller/user_ctl/
├── user_controller.go               # 用户相关 Handler（含 Swagger 注释）
└── auth_controller.go               # 绑定相关 Handler

app/middleware/
└── user_auth.go                     # UserLoginAuth() JWT鉴权中间件

routes/user_route/
└── user_route.go                    # /api/user 路由注册

config/
└── jwt.go                           # JWT 配置 + GetUserID() 辅助方法

util/
├── jwt.go                           # JWT 生成/解析/续签（UserClaims 仅含 userID+tokenVersion+tokenType）
├── email_code.go                    # 邮箱验证码发送/校验（Redis 存储+频率限制）
└── sms_code.go                      # 短信验证码发送/校验（Redis 存储+频率限制）

docs/migration/
└── add_user_tables.sql              # 手动建表 SQL（备用）
```

---

## 数据库设计

### user 表（用户主表，手机号唯一存储位置）

| 字段 | 类型 | 说明 |
|------|------|------|
| id | char(36) PK | UUID v7（BeforeCreate 自动生成） |
| email | varchar(128) UNIQUE | 邮箱（主账号，必填） |
| nickname | varchar(64) | 昵称 |
| avatar | varchar(255) | 头像 |
| password | varchar(255) | 密码（bcrypt） |
| salt | varchar(32) | 加密盐 |
| phone | varchar(20) INDEX | 手机号（可选，绑定后存入） |
| phone_code | varchar(10) | 区号（默认86） |
| status | tinyint | 0正常 1禁用 |
| token_version | bigint | 踢人下线版本号（自增使旧token失效） |
| last_login_ip | varchar(50) | 最后登录IP |
| last_login_time | datetime | 最后登录时间 |
| is_delete / create_time / update_time / delete_time | — | 软删除 + 时间戳 |

> **手机号存储规范**：手机号只存 `x_user` 表，不在 `x_user_auth` 表中冗余存储。

### user_auth 表（第三方认证绑定表，不含手机号）

| 字段 | 类型 | 说明 |
|------|------|------|
| id | char(36) PK | UUID v7 |
| user_id | char(36) INDEX | 用户ID |
| identity_type | varchar(20) | 认证类型常量 |
| identifier | varchar(128) | 标识（openid/unionid） |
| credential | varchar(255) | 凭证（部分类型为空） |
| extra | varchar(512) | 扩展信息 JSON |

**UNIQUE INDEX**: `(identity_type, identifier)` 防止重复绑定

### IdentityType 常量

```go
const (
    IdentityWechatMini = "wechat_mini" // 微信小程序
    IdentityWechatMp   = "wechat_mp"   // 微信公众号
    IdentityWechatApp  = "wechat_app"  // 微信APP
    IdentityQQ         = "qq"          // QQ
)
```

---

## JWT 鉴权机制

### Token 设计

| Token | 有效期 | 用途 |
|-------|--------|------|
| access_token | 2小时（可配） | 接口鉴权，放 Header `Authorization: Bearer xxx` |
| refresh_token | 7天（可配） | 换取新 token 对 |

### UserClaims 结构（JWT Payload）

```go
type UserClaims struct {
    UserID       string `json:"userId"`
    TokenVersion int64  `json:"tokenVersion"`
    TokenType    string `json:"tokenType"` // "access" | "refresh"
    jwt.RegisteredClaims
}
```

> **注意**：JWT 中**不存储 email**，仅含 userID + tokenVersion + tokenType，减小 token 体积。

### 鉴权流程（UserLoginAuth 中间件）

```
请求 → 解析 Bearer token
     → 验证 JWT 签名和有效期
     → 读 Redis 缓存 token_version（key=user:tv:{userID}，TTL 5min）
     → 缓存 miss 时回源 MySQL
     → 对比 JWT 中 tokenVersion 与缓存值
     → 不一致 → 401（踢人下线生效）
     → 一致 → 写入 gin.Context（仅 userID）
     → access_token 剩余<30分钟 → 响应头返回 X-New-Access-Token（自动续签）
```

### 性能优化规范

- **JWT Claims 不存 email**，仅保留 userID、tokenVersion、tokenType
- **中间件不直接查 MySQL**，通过 Redis 缓存 `user:tv:{userID}`（5分钟 TTL）
- **踢人下线 / 重置密码后立即清除缓存**，确保旧 token 即时失效

### 获取当前用户ID

```go
// Controller 中获取当前登录用户ID
userID := config.JWTConfig.GetUserID(c)
```

---

## 踢人下线机制

基于 `token_version` 字段实现：

1. 用户注册时 `token_version = 0`
2. 登录时 JWT Claims 写入当前 `token_version`
3. 每次鉴权中间件对比 JWT 中的 `token_version` 与数据库/缓存值
4. 踢人下线 / 重置密码 → `token_version++` + 清除 Redis 缓存
5. 所有旧 token 下次请求时 version 不匹配 → 401

---

## 验证码

### 邮箱验证码

```go
// 发送（场景: register/reset/bind/unbind）
util.EmailCodeUtil.SendCode(email, scene)

// 校验（成功后自动删除验证码）
util.EmailCodeUtil.VerifyCode(email, scene, code)
```

| 场景常量 | 用途 |
|----------|------|
| `CodeSceneRegister` | 邮箱注册 |
| `CodeSceneReset` | 邮箱重置密码 |
| `CodeSceneBind` | 绑定邮箱 |
| `CodeSceneUnbind` | 解绑手机（邮箱验证码确认身份） |

### 短信验证码

```go
// 发送（场景: sms_bind/sms_login/sms_reset）
util.SmsCodeUtil.SendCode(phone, scene)

// 校验（成功后自动删除验证码）
util.SmsCodeUtil.VerifyCode(phone, scene, code)
```

| 场景常量 | 用途 |
|----------|------|
| `SmsSceneBind` | 绑定手机号 |
| `SmsSceneLogin` | 手机号+短信验证码登录 |
| `SmsSceneReset` | 手机号重置密码 |

> **注意**：短信验证码当前仅记录日志（`core.Logger.Infof`），生产环境需对接阿里云/腾讯云等短信服务商。

### Redis Key 设计

| Key | 用途 | TTL |
|-----|------|-----|
| `user:code:{scene}:{email}` | 邮箱验证码值 | 5分钟 |
| `user:code:limit:{email}` | 邮箱60秒频率限制 | 60秒 |
| `user:code:daily:{email}:{date}` | 邮箱每日发送次数上限(10次) | 24小时 |
| `user:sms:{scene}:{phone}` | 短信验证码值 | 5分钟 |
| `user:sms:limit:{phone}` | 短信60秒频率限制 | 60秒 |
| `user:sms:daily:{phone}:{date}` | 短信每日发送次数上限(10次) | 24小时 |

---

## 接口一览

### 免登录接口

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /api/user/sendCode | 发送邮箱验证码（场景: register/reset/unbind） |
| POST | /api/user/sendSmsCode | 发送短信验证码（场景: sms_bind/sms_login/sms_reset） |
| POST | /api/user/register | 邮箱注册（需邮箱验证码） |
| POST | /api/user/login | 邮箱+密码登录 |
| POST | /api/user/phoneLogin | 手机号+密码登录 |
| POST | /api/user/phoneCodeLogin | 手机号+短信验证码登录 |
| POST | /api/user/refresh | 刷新token |
| POST | /api/user/resetPassword | 邮箱重置密码（邮箱验证码+新密码） |
| POST | /api/user/resetPhonePassword | 手机号重置密码（短信验证码+新密码） |

### 需要登录的接口（JWT）

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /api/user/info | 获取用户信息 |
| POST | /api/user/info | 更新昵称/头像 |
| POST | /api/user/kickOffline | 踢人下线 |
| POST | /api/user/bindPhone | 绑定手机号（需短信验证码） |
| POST | /api/user/unbindPhone | 解绑手机号（需邮箱验证码） |
| GET | /api/user/authList | 获取第三方绑定列表（微信/QQ，不含手机号） |

---

## 路由注册方式

用户系统路由**不走** admin_route 的 Autoload 机制，而是在 `routes/api.go` 中直接挂载：

```go
// routes/api.go
userRg := api.Group("/user")
user_route.UserRoute(userRg)
```

免登录接口在 `rg.POST(...)` 直接注册；需登录接口通过 `rg.Group("/", middleware.UserLoginAuth())` 分组。

---

## 配置

### .env.yaml

```yaml
JWT:
  AccessSecret: 'your_access_secret'
  RefreshSecret: 'your_refresh_secret'
  AccessExpireSec: 7200     # access_token 有效期(秒)
  RefreshExpireSec: 604800  # refresh_token 有效期(秒)
```

---

## 开发约定

### 新增用户系统接口

遵循与 admin 模块相同的四件套模式，区别在于：

| 对比项 | admin 模块 | user 模块 |
|--------|-----------|-----------|
| Controller 目录 | `app/controller/admin_ctl/` | `app/controller/user_ctl/` |
| Schema 目录 | `app/schema/<module>_schema/` | `app/schema/user_schema/` |
| Service 目录 | `app/service/<module>_service/` | `app/service/user_service/` |
| Route 目录 | `routes/admin_route/` | `routes/user_route/` |
| 鉴权中间件 | `middleware.LoginAuth()` / `PermAuth()` | `middleware.UserLoginAuth()` |
| 获取用户ID | `config.AdminConfig.GetAdminId(c)` | `config.JWTConfig.GetUserID(c)` |

### 密码处理

- 使用 `golang.org/x/crypto/bcrypt` 加密（非 admin 系统的 MD5+salt 方式）
- `user_service.hashPassword(password)` / `checkPassword(hashed, password)`

### 第三方绑定扩展

后续新增微信/QQ绑定时：
1. 在 `user_auth.go` 中使用已有的 `IdentityType` 常量（wechat_mini/wechat_mp/wechat_app/qq）
2. 在 `auth_service.go` 中新增对应绑定/解绑方法
3. 在 `auth_controller.go` 中新增 Handler + Swagger 注释
4. 在 `user_route.go` 的 auth 分组中追加路由

### 手机号绑定/解绑规则

- **绑定**：需要短信验证码（`sms_bind` 场景），直接写入 `x_user.phone`
- **解绑**：需要邮箱验证码（`unbind` 场景），因为手机号可能已注销无法接收短信
- **手机号非必填**：用户可以不绑定手机号，仅用邮箱登录

### 自动迁移

`main.go` 启动时自动建表：
```go
core.AutoMigrate(&user_model.User{}, &user_model.UserAuth{})
```
