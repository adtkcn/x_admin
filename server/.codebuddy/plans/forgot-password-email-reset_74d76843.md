---
name: forgot-password-email-reset
overview: 为 likeadmin_go 项目增加"通过邮件验证码重置密码"功能，包含后端接口和 admin(Vue3) 前端页面。流程：用户输入邮箱 → 发送验证码到邮箱 → 输入验证码+新密码 → 重置成功。
design:
  architecture:
    framework: vue
  styleKeywords:
    - Minimalism
    - Clean
    - Form-focused
  fontSystem:
    fontFamily: PingFang SC, Helvetica Neue, Helvetica, Arial, sans-serif
    heading:
      size: 30px
      weight: 600
    subheading:
      size: 16px
      weight: 500
    body:
      size: 14px
      weight: 400
  colorSystem:
    primary:
      - "#409EFF"
      - "#66B1FF"
    background:
      - "#F8F8F8"
      - "#FFFFFF"
    text:
      - "#303133"
      - "#606266"
      - "#909399"
    functional:
      - "#67C23A"
      - "#F56C6C"
      - "#E6A23C"
      - "#409EFF"
todos:
  - id: step-1-schema
    content: 在 login_schema.go 中新增忘记密码请求结构体（SendCodeReq 和 ResetReq）
    status: completed
  - id: step-2-service
    content: 新建 forget_pwd_service.go，实现 SendResetCode 和 ResetPassword 方法
    status: completed
    dependencies:
      - step-1-schema
  - id: step-3-controller
    content: 在 login.go controller 中新增 ForgotPwdSendCode 和 ForgotPwdReset 两个处理方法
    status: completed
    dependencies:
      - step-2-service
  - id: step-4-routes
    content: 在 system_route.go 的 initLoginRoute 中注册两个忘记密码路由（无需认证）
    status: completed
    dependencies:
      - step-3-controller
  - id: step-5-api-frontend
    content: 在 admin/src/api/user.ts 中新增 forgotPwdSendCode 和 forgotPwdReset 两个 API 函数
    status: completed
  - id: step-6-page-enum
    content: 在 pageEnum.ts 中新增 FORGOT_PASSWORD 路由枚举
    status: completed
    dependencies:
      - step-5-api-frontend
  - id: step-7-forgot-page
    content: 新建 forgot-password.vue 页面，实现邮箱输入、验证码发送、密码重置完整交互
    status: completed
    dependencies:
      - step-6-page-enum
  - id: step-8-route-register
    content: 在 routes.ts 中注册忘记密码页面路由，并在登录页添加"忘记密码"跳转链接
    status: completed
    dependencies:
      - step-7-forgot-page
---

## 用户需求

实现通过邮件重置密码功能，具体需求如下：

1. **场景**：用户忘记密码，通过邮箱重置（无需登录）
2. **验证方式**：发送验证码到邮箱（输入验证码 + 新密码重置）
3. **前端**：需要配套前端页面，使用现有的 `admin`（Vue3 + Element Plus）前端项目

## 功能概述

- 用户在登录页点击"忘记密码"链接，跳转到忘记密码页面
- 用户输入注册邮箱，点击发送验证码（后端校验邮箱是否存在、限流、发邮件）
- 用户输入收到的验证码和新密码，提交重置
- 后端校验验证码正确性、更新密码（重新生成 salt 并加密）

## 核心功能

- 发送重置密码验证码邮件（后端 API + 邮件发送）
- 校验验证码并重置密码（后端 API）
- 忘记密码前端页面（邮箱输入 → 验证码+新密码输入）
- 登录页添加"忘记密码"跳转链接

## 技术栈

- **后端**：Go 1.26 + Gin + GORM + Redis + go-mail（已有，无需新增依赖）
- **前端**：Vue3 + TypeScript + Element Plus + Vite（已有，无需新增依赖）

## 实现方案

### 后端实现

#### 1. 新增忘记密码 Schema（`app/schema/system_schema/login_schema.go`）

在现有 `login_schema.go` 中追加两个请求结构体：

- `SystemForgotPwdSendCodeReq`：邮箱字段（binding: required,email）
- `SystemForgotPwdResetReq`：邮箱、验证码、新密码字段

#### 2. 新增忘记密码 Service（`app/service/system_service/forget_pwd_service.go`）

新建文件，实现 `ForgetPwdService` 单例，包含两种方法：

- **SendResetCode(email) error**：

1. 调用 `AdminService.FindByEmail(email)` 校验邮箱是否存在，不存在则返回"邮箱未注册"
2. 使用 `middleware.LimitEmail` 相同的限流逻辑（或复用 Redis 计数器），防止频繁发送
3. 生成 6 位随机数字验证码
4. 用 `util.RedisUtil.Set("forgot_pwd:code:"+email, code, 300)` 存入 Redis，有效期 5 分钟
5. 调用 `util.EmailUtil.SendEmail()` 发送邮件，邮件内容包含验证码和有效期说明

- **ResetPassword(email, code, newPassword) error**：

1. 从 Redis 获取验证码：`util.RedisUtil.Get("forgot_pwd:code:" + email)`
2. 验证码不存在或已过期：返回"验证码已过期"
3. 验证码不匹配：返回"验证码错误"
4. 验证通过：生成新 salt（`util.ToolsUtil.RandomString(5)`），计算新密码 `MD5(MD5(newPassword) + salt)`
5. 更新数据库：`db.Model(&SystemAuthAdmin{}).Where("email = ?", email).Updates(map[string]interface{}{"password": newPwd, "salt": salt})`
6. 删除 Redis 中的验证码（一次性使用）
7. 可选：清除该用户的所有登录 token（强制重新登录）

#### 3. 注册路由（`routes/admin_route/system_route.go`）

在 `initLoginRoute` 函数中新增两个无需认证的路由：

- `loginRg.POST("/forgot-pwd/send-code", handleLogin.ForgotPwdSendCode)`
- `loginRg.POST("/forgot-pwd/reset", handleLogin.ForgotPwdReset)`

#### 4. 新增 Controller 方法（`app/controller/admin_ctl/system_controller/login.go`）

在现有 `LoginHandler` 中追加两个方法：

- `ForgotPwdSendCode(c *gin.Context)`：绑定参数、调用 Service、返回响应
- `ForgotPwdReset(c *gin.Context)`：绑定参数（含新密码）、调用 Service、返回响应

#### 5. 邮件内容模板（`util/email.go` 或 Service 中）

重置密码邮件使用 HTML 格式，内容包含：

- 标题：密码重置验证码
- 验证码（大字体显示）
- 提示：5 分钟内有效，请勿泄露给他人

### 前端实现

#### 1. 新增 API 函数（`admin/src/api/user.ts`）

追加两个 API 函数：

- `forgotPwdSendCode(data: { email: string })` → POST `/system/forgot-pwd/send-code`
- `forgotPwdReset(data: { email: string, code: string, password: string })` → POST `/system/forgot-pwd/reset`

#### 2. 新增忘记密码页面（`admin/src/views/account/forgot-password.vue`）

分步骤表单：

- **第一步**：输入邮箱 → 点击"发送验证码"按钮（按钮显示倒计时）
- **第二步**：输入验证码 + 新密码 + 确认新密码 → 点击"重置密码"
- 使用 Element Plus 的 `el-form`、`el-input`、`el-button` 组件
- 密码字段使用 `encryptPassword` 加密后再提交（与登录保持一致）
- 重置成功后，提示"密码已重置，即将跳转到登录页"，3 秒后跳转

#### 3. 注册路由（`admin/src/router/routes.ts`）

在 `constantRoutes` 中新增：

```ts
{
    path: PageEnum.FORGOT_PASSWORD,
    component: () => import('@/views/account/forgot-password.vue')
}
```

在 `admin/src/enums/pageEnum.ts` 中新增 `FORGOT_PASSWORD = '/forgot-password'`

#### 4. 登录页添加跳转链接（`admin/src/views/account/login.vue`）

在登录按钮下方添加"忘记密码？"链接，点击跳转到 `/forgot-password`

## 关键设计说明

- **密码加密方式**：与现有系统保持一致，使用 `MD5(MD5(password) + salt)`，前端提交时也用同样的双重 MD5 加密（`encryptPassword` 函数）
- **验证码存储**：使用 Redis，key 为 `forgot_pwd:code:{email}`，有效期 300 秒
- **限流**：复用项目已有的 `middleware.LimitEmail(2, 60)` 中间件，限制同一邮箱 60 秒内最多发送 2 次验证码
- **安全性**：验证码一次性使用（验证通过后立即删除），密码重置后清除该用户所有 token

## 页面设计：忘记密码页面（forgot-password.vue）

### 整体布局

采用与登录页一致的左右分栏布局：左侧为装饰性背景图，右侧为表单卡片。整体风格与现有登录页保持一致。

### 页面区块设计（从上到下）

#### 区块1：页面标题

- 居中显示"忘记密码"大标题（text-3xl font-medium）
- 副标题："请输入注册邮箱，我们将发送验证码帮助您重置密码"

#### 区块2：邮箱输入表单

- 使用 `el-form` 包裹
- **邮箱输入框**（`el-input` + prefix icon `el-icon-Message`）
- placeholder: "请输入注册邮箱"
- 校验规则：必填、邮箱格式
- **发送验证码按钮**（`el-button`，与输入框在同一行或用单独一行）
- 点击后调用 API 发送验证码
- 发送成功后按钮变为倒计时（60秒），期间不可点击
- 按钮文字：发送前"发送验证码" / 倒计时中"重新发送(60s)"

#### 区块3：验证码 + 新密码输入

- **验证码输入框**（`el-input`）
- placeholder: "请输入6位验证码"
- 长度限制：6位
- **新密码输入框**（`el-input` type="password"，show-password）
- placeholder: "请输入新密码（6-32位）"
- **确认新密码输入框**（`el-input` type="password"，show-password）
- placeholder: "请再次输入新密码"
- 前端校验：两次密码是否一致

#### 区块4：操作按钮

- **重置密码按钮**（`el-button type="primary"`，大号，通栏）
- 点击后校验表单 → 调用 `forgotPwdReset` API
- 成功后：弹窗提示"密码重置成功"，3秒后跳转到登录页
- **返回登录链接**：按钮下方显示"返回登录"，点击跳转到 `/login`

### 交互设计

- 表单校验：使用 Element Plus 的 `rules` 进行前端校验
- 加载状态：发送验证码和重置密码时按钮显示 loading 状态
- 响应式：移动端适配（单行布局，无左侧背景图）

### 登录页改动

在登录按钮下方、注册链接（如有）旁新增"忘记密码？"链接文字，点击 `router.push('/forgot-password')`。

## Agent Extensions

### Skill

- **design-to-code-workflows**
- Purpose: 辅助生成忘记密码页面的 Vue3 + Element Plus 代码组件
- Expected outcome: 生成符合项目代码风格的 `forgot-password.vue` 单文件组件

### SubAgent

- **code-explorer**
- Purpose: 深度探索后端代码库，定位需要修改的确切位置和参考现有实现模式
- Expected outcome: 确认 `login_schema.go`、`admin_service.go`、`system_route.go` 的精确修改点，确保与现有代码风格一致