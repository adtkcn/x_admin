---
# https://vitepress.dev/reference/default-theme-home-page
layout: home

hero:
  name: "x_admin"
  text: "Go语言 通用管理后台"
  tagline: x_admin是一套快速开发通用管理后台，使用流行的技术栈Go、TypeScript、Vue3、vite、Element Plus。一键生成前端、后端、uniapp代码
  actions:
    - theme: brand
      text: 前端文档
      link: /admin/准备.html
    - theme: brand
      text: 后端文档
      link: /server/准备.html

features:
  - title: 方便部署
    details: 未使用cgo，Go可以跨平台打包，不需要运行环境Sdk
  - title: 轻量级
    details: 运行时不到50M内存占用，暴打Java
  - title: 快速生成CRUD代码
    details: 通过已有数据库表生成CRUD代码，减少开发成本
---

