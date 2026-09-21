import { defineConfig } from 'vitepress'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  base:'/x_admin/',
  title: "x_admin",
  description: "x_admin 开发文档",
  ignoreDeadLinks: true,
  markdown: {
    lineNumbers: true
  },
  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    nav: [
      { text: 'Home', link: '/' },
      { text: '后端', link: '/server/功能说明/模块概述.md' },
      { text: '前端', link: '/admin/准备.md' },
      { text: '在线体验', link: 'https://x.adtk.cn/' },
    ],

    sidebar: [
      {
        text: '后端-server',
        collapsed: false,
        items: [
          { text: '功能总览', link: '/server/功能说明/模块概述.md' },
          {
            text: '功能说明',
            collapsed: false,
            items: [
              { text: '模块概述', link: '/server/功能说明/模块概述.md' },
              { text: '定时任务', link: '/server/功能说明/定时任务.md' },
              { text: '异步队列', link: '/server/功能说明/异步队列.md' },
              { text: '邮件发送', link: '/server/功能说明/邮件发送.md' },
              { text: 'WebSocket', link: '/server/功能说明/WebSocket.md' },
              { text: '事件总线 pubsub', link: '/server/功能说明/事件总线pubsub.md' },
              { text: '存储引擎', link: '/server/功能说明/存储引擎.md' },
              { text: '工具函数与核心组件', link: '/server/功能说明/工具函数与核心组件.md' },
            ]
          },
          {
            text: '开发指南',
            collapsed: true,
            items: [
              { text: '准备', link: '/server/准备.md' },
              { text: '结构说明', link: '/server/结构说明.md' },
              { text: '权限验证', link: '/server/权限验证.md' },
              { text: '注意事项', link: '/server/注意事项.md' },
              { text: '数据库null值', link: '/server/数据库null值.md' },
              { text: 'swaggo', link: '/server/swaggo.md' },
              { text: '导入导出excel', link: '/server/导入导出excel.md' },
            ]
          },
          {
            text: '部署运维',
            collapsed: true,
            items: [
              { text: 'Go打包', link: '/server/部署Go.md' },
              { text: 'nginx配置', link: '/server/nginx配置.md' },
              { text: '.env配置选项', link: '/server/环境变量.md' },
            ]
          },
        ]
      },
      {
        text: '前端-admin',
        items: [          
          { text: '准备', link: '/admin/准备.md' },
  
          { text: 'nginx配置', link: '/admin/nginx配置.md' },

          { text: '路由', link: '/admin/路由.md' },
          { text: '嵌入iframe', link: '/admin/嵌入iframe.md' },

          { text: '自定义hooks', link: '/admin/自定义hooks.md' },
        ]
      }
    ],

    socialLinks: [
      { icon: 'github', link: 'https://github.com/adtkcn/x_admin' }
    ]
  }
})
