---
name: "项目总导航"
description: "x_admin 仓库根目录导航 — 介绍 monorepo 各子项目构成，并指引到子目录（server 等）的 skills"
---

# x_admin 项目总览（导航入口）

x_admin 是一个免费开源的前后端分离通用管理后台（Go + Gin + Vue3 + Element Plus + uni-app），本仓库为 monorepo，包含以下子项目。

## 仓库目录构成

| 目录 | 技术栈 | 说明 |
|------|--------|------|
| `server/` | Go 1.27 / Gin / GORM / Redis | **后端服务**：后台管理 API（`/api/admin`）、C端用户 API、IoT 物联网域、发布域、WebSocket、消息队列、定时任务 |
| `admin/` | Vue3 / TypeScript / Vite / Element Plus | **后台管理前端**：系统管理、IoT 大屏、审批流、代码生成器、监控等页面 |
| `x_admin_app/` | uni-app x (UTS/uvue) | **移动端 App**：登录、工作台、监控、视频等，配套 uvue skill 与 WGT 热更升级 SDK |
| `docs/` | VitePress | **开发部署文档站**（https://x.adtk.cn/x_admin/ 的源码），含 `docs/sql/` 初始化脚本 |
| `x_err_sdk/` | TypeScript / Vite | **错误上报 SDK**：web / uni / react-native 三套产物，配合 `plugin/` 监控体系 |
| `iot-mock/` | Go | **IoT 设备模拟器**：向后端投递模拟报文，用于联调规则链 / 告警 |

## Skills 导航

各子项目的 agents skills 统一放在**所在目录的 `.agents/skills/` 下**，在对应目录中工作时优先使用：

### 后端（`server/.agents/skills/`）

| Skill | 什么时候用 |
|-------|-----------|
| [server-project](../../../server/.agents/skills/server-project/SKILL.md) | **后端导航入口 + 完整技术文档**：后端有哪些东西（路由域/业务模块/基础设施），及技术栈、目录结构、架构、子 skills 导航 |
| [add-api](../../../server/.agents/skills/add-api/SKILL.md) | 在已有模块中新增接口：Schema / Service / Controller / Route 四件套模板 |
| [user-system](../../../server/.agents/skills/user-system/SKILL.md) | C端用户系统：JWT 鉴权、邮箱/手机/微信登录、验证码、踢人下线 |
| [wechatpay-payment-integration](../../../server/.agents/skills/wechatpay-payment-integration/SKILL.md) | 微信支付接入、产品选型、排障与质量检查 |

### 管理端前端（`admin/.agents/skills/`）

| Skill | 什么时候用 |
|-------|-----------|
| [admin-project](../../admin/.agents/skills/admin-project/SKILL.md) | **管理端前端导航 + 技术文档**：Vue3/Vite/Element Plus 目录结构、分层约定、业务域、组件与 hooks |

### App 端（`x_admin_app/.agents/skills/`）

| Skill | 什么时候用 |
|-------|-----------|
| [app-project](../../x_admin_app/.agents/skills/app-project/SKILL.md) | **App 端导航 + 技术文档**：uni-app x 目录结构、页面域、请求/升级 SDK、hb-cli 打包 |
| [uvue](../../x_admin_app/.agents/skills/uvue/SKILL.md) | 写 uvue 页面/组件必读：UTS 严格类型、条件编译、滚动容器等编码规范 |
