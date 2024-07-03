import { defineConfig } from 'vitepress'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  base:'/x_admin/',
  title: "x_admin",
  description: "x_admin 开发文档",

  markdown: {
    lineNumbers: true
  },
  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    nav: [
      { text: 'Home', link: '/' },
      { text: '后端', link: '/server/准备.md' },
      { text: '前端', link: '/admin/准备.md' },
      { text: '在线体验', link: 'https://x.adtk.cn/' },
    ],

    sidebar: [
      {
        text: '后端-server',
        items: [
          // { text: '后端', link: '/markdown-examples' },
          // { text: '前端', link: '/api-examples' },
          
          { text: '准备', link: '/server/准备.md' },
          { text: 'Go打包', link: '/server/部署Go.md' },
          { text: 'nginx配置', link: '/server/nginx配置.md' },
          { text: '.env配置选项', link: '/server/环境变量.md' },
          { text: '权限验证', link: '/server/权限验证.md' },
          { text: '结构说明', link: '/server/结构说明.md' },
          { text: '注意事项', link: '/server/注意事项.md' },
          { text: '数据库null值', link: '/server/数据库null值.md' },
        ]
      },
      {
        text: '前端-admin',
        items: [          
          { text: '准备', link: '/admin/准备.md' },
  
          { text: 'nginx配置', link: '/admin/nginx配置.md' },

          { text: '路由', link: '/admin/路由.md' },
        ]
      }
    ],

    socialLinks: [
      { icon: 'github', link: 'https://github.com/adtkcn/x_admin' }
    ]
  }
})
