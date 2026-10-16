import { defineConfig } from 'vite'

// react-native：ES/CJS 产物，react-native 作为宿主依赖外部化，不打包进 SDK
export default defineConfig({
  build: {
    // 三套产物共用 dist，禁止互相清空（默认会清空 outDir）
    emptyOutDir: false,
    lib: {
      entry: 'react-native/index.ts',
      formats: ['es'],
      fileName: (format) => `react-native/XErr.${format}.js`,
    },
    rolldownOptions: {
      external: ['react-native'],
      output: {},
    },
  },
})
