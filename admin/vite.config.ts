import { fileURLToPath, URL } from 'url'

import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueJsx from '@vitejs/plugin-vue-jsx'
// import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'

// import { ElementPlusResolver } from 'unplugin-vue-components/resolvers'
// import { createStyleImportPlugin, ElementPlusResolve } from 'vite-plugin-style-import'
import { createSvgIconsPlugin } from 'vite-plugin-svg-icons'

// import viteCompression from 'vite-plugin-compression'
import { visualizer } from 'rollup-plugin-visualizer'

// https://vitejs.dev/config/
export default ({ mode }) => {
    const env = loadEnv(mode, process.cwd())
    console.log(env)

    return defineConfig({
        optimizeDeps: {
            // 依赖预构建，避免开发刷新
            include: ['@wangeditor/editor-for-vue', 'vuedraggable', 'vue-echarts', 'crypto-js']
        },

        base: '/',
        build: {
            sourcemap: true,
            rollupOptions: {
                external: ['XErr'],
                output: {
                    manualChunks: {
                        vue: ['vue'],
                        'vue-router': ['vue-router'],
                        pinia: ['pinia'],
                        axios: ['axios'],
                        dayjs: ['dayjs'],
                        // echarts: ['echarts'],
                        // 'highlight.js': ['highlight.js'],
                        'element-plus': ['element-plus']

                        // 'lodash-es': ['lodash-es'],
                        // vuedraggable: ['vuedraggable'],
                        // 'vform3-builds': ['vform3-builds']
                    }
                }
            }
        },

        server: {
            open: true,
            host: '0.0.0.0',
            port: 5174,
            proxy: {
                '/api': {
                    target: env.VITE_APP_BASE_URL,
                    changeOrigin: true,
                    ws: true
                }
            }
        },
        plugins: [
            vue(),
            vueJsx(),
            // AutoImport({
            //     imports: ['vue', 'vue-router'],
            //     // resolvers: [ElementPlusResolver()],
            //     eslintrc: {
            //         enabled: true
            //     }
            // }),
            Components({
                directoryAsNamespace: true
                // resolvers: [ElementPlusResolver()]
            }),
            // createStyleImportPlugin({
            //     resolves: [ElementPlusResolve()]
            // }),
            createSvgIconsPlugin({
                // 配置路劲在你的src里的svg存放文件
                iconDirs: [fileURLToPath(new URL('./src/assets/icons', import.meta.url))],
                symbolId: 'local-icon-[dir]-[name]'
            }),
            // viteCompression({
            //     algorithm: 'gzip'
            // })
            // viteCompression({
            //     algorithm: 'brotliCompress'
            // })
            visualizer({
                gzipSize: false,
                brotliSize: false,
                emitFile: false,
                filename: 'test.html', //分析图生成的文件名
                open: true
            })
        ],
        resolve: {
            alias: {
                '@': fileURLToPath(new URL('./src', import.meta.url))
            }
        }
    })
}
