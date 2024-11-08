import { defineConfig, loadEnv } from 'vite'

export default defineConfig({
    build: {
      lib: {
        entry: 'web/index.ts', // 库的入口文件
        name: 'XErr', // 库的名称
        fileName: (format) => `web/XErr.${format}.js`, // 输出文件名
      },
      rollupOptions: {
        // 外部化处理不想打包进库的依赖
        // external: ['vue'],
        output: {
          globals: {
            // vue: 'Vue'
          }
        }
      }
    }
  })
  