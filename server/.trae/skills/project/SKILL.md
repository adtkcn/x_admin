---
name: "项目目录文件结构"
description: "项目目录文件结构"
---

## server目录结构描述

```
├── server
│   ├── app // 应用程序的核心代码
│   │   ├───controller //控制器
│   │   │   ├── admin_ctl //后台控制器
│   │   │   ├── web_ctl //前端控制器
│   │   ├──────service // 服务
│   │   ├──────schema // 请求与返回的结构体定义
│   │   ├──────cron // 数据库模型
│   │   ├──────middleware // 中间件
│   │   model // 定时任务
│   ├── config // 配置
│   ├── core // core
│   ├── util // 工具包
│   ├── router // 路由目录
│   │   ├──────admin_route // /api/admin 后台路由
│   │   ├──────web_route // /api/web 路由
│   │   ├── api.go // /api入口文件
│   │   ├── route.go // 路由入口
│   ├── public 
│   │   ├──resources // 验证码等依赖文件
│   │   ├──static // 静态文件，访问路由`/api/static/*`
│   ├── main.go // 入口
│   ├── .env.yaml // 配置文件，注意不提交git
```
