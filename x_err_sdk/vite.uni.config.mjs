import { defineConfig } from 'vite'

// uniapp：ES/CJS 产物，供 uni-app 项目 import 使用
export default defineConfig({
  build: {
    // 三套产物共用 dist，禁止互相清空（默认会清空 outDir）
    emptyOutDir: false,
    lib: {
      entry: 'uni/index.ts',
      formats: ['es'],
      fileName: (format) => `uni/XErr.${format}.js`,
    },
    rolldownOptions: {
      output: {},
    },
  },
})
