import { defineConfig } from 'vitepress'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  // base:'/docs/',
  title: "x_admin",
  description: "x_admin 开发文档",
  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    nav: [
      { text: 'Home', link: '/' },
      { text: '后端', link: '/server/准备.md' },
      { text: '前端', link: '/admin/准备.md' }
    ],

    sidebar: [
      {
        text: '后端',
        items: [
          // { text: '后端', link: '/markdown-examples' },
          // { text: '前端', link: '/api-examples' },
          
          { text: '准备', link: '/server/准备.md' },
          { text: 'Go打包', link: '/server/部署Go.md' },
          { text: 'nginx配置', link: '/server/nginx配置.md' },
        ]
      },
      {
        text: '前端',
        items: [          
          { text: '准备', link: '/admin/准备.md' },
  
          { text: 'nginx配置', link: '/admin/nginx配置.md' },
        ]
      }
    ],

    socialLinks: [
      { icon: 'github', link: 'https://github.com/adtkcn/x_admin' }
    ]
  }
})
