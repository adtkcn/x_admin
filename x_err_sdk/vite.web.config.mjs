import { defineConfig } from 'vite'

// Web：UMD/ES 产物，可通过 <script src="XErr.umd.js"> 直接使用全局 XErr
export default defineConfig({
  build: {
    // 三套产物共用 dist，禁止互相清空（默认会清空 outDir）
    emptyOutDir: false,
    lib: {
      entry: 'web/index.ts',
      name: 'XErr',
      formats: ['es', 'umd'],
      fileName: (format) => `web/XErr.${format}.js`,
    },
    rolldownOptions: {
      output: {
        globals: {},
      },
    },
  },
})
