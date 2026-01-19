import { fileURLToPath, URL } from 'url'

import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'
// import vueJsx from '@vitejs/plugin-vue-jsx'
// import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'

// import { ElementPlusResolver } from 'unplugin-vue-components/resolvers'
// import { createStyleImportPlugin, ElementPlusResolve } from 'vite-plugin-style-import'
import { createSvgIconsPlugin } from 'vite-plugin-svg-icons'

// import viteCompression from 'vite-plugin-compression'
import { visualizer } from 'rollup-plugin-visualizer'

// https://vitejs.dev/config/

export default defineConfig(({ mode }) => {
    const env = loadEnv(mode, process.cwd())
    console.log(env)

    return {
        experimental: {
            enableNativePlugin: true // 启用 Rust 原生插件（如 alias/resolve）
        },
        optimizeDeps: {
            // 依赖预构建，避免开发刷新
            include: ['@wangeditor/editor-for-vue', 'vuedraggable', 'crypto-js']
        },

        base: '/',
        build: {
            sourcemap: true,
            rolldownOptions: {
                external: ['XErr'],
                output: {
                    codeSplitting: {
                        groups: [
                            {
                                name: 'vue_vue-router_pinia',
                                test: /node_modules\/(vue|vue-router|pinia)\//
                            },
                            // {
                            //     name: 'vue-router',
                            //     test: /node_modules\/vue-router\//
                            // },
                            // pinia
                            // {
                            //     name: 'pinia',
                            //     test: /node_modules\/pinia\//
                            // },
                            // @vueuse/core
                            {
                                name: '@vueuse',
                                test: /node_modules\/@vueuse\//
                            },

                            {
                                name: 'element-plus-icons-vue',
                                test: /node_modules\/@element-plus\/icons-vue\//
                            },
                            {
                                name: 'element-plus',
                                test: /node_modules\/element-plus\//
                            },

                            {
                                name: 'axios',
                                test: /node_modules\/axios\//
                            },
                            {
                                name: 'dayjs',
                                test: /node_modules\/dayjs\//
                            },
                            // vuedraggable
                            {
                                name: 'vuedraggable',
                                test: /node_modules\/vuedraggable\//
                            },
                            // vue3-video-play
                            {
                                name: 'vue3-video-play',
                                test: /node_modules\/vue3-video-play\//
                            },

                            // zrender
                            {
                                name: 'zrender',
                                test: /node_modules\/zrender\//
                            },
                            // echarts
                            {
                                name: 'echarts',
                                test: /node_modules\/echarts\//
                            },
                            // highlight.js
                            {
                                name: 'highlight.js',
                                test: /node_modules\/highlight\.js\//
                            },
                            // lodash-es
                            {
                                name: 'lodash-es',
                                test: /node_modules\/lodash-es\//
                            },
                            // @logicflow/core
                            {
                                name: '@logicflow/core',
                                test: /node_modules\/@logicflow\/core\//
                            },
                            // @logicflow/extension
                            {
                                name: '@logicflow/extension',
                                test: /node_modules\/@logicflow\/extension\//
                            },
                            // @wangeditor/editor
                            {
                                name: '@wangeditor/editor',
                                test: /node_modules\/@wangeditor\//
                            },
                            // xe-utils
                            {
                                name: 'xe-utils',
                                test: /node_modules\/xe-utils\//
                            },
                            // vxe-table
                            {
                                name: 'vxe-table',
                                test: /node_modules\/vxe-table\//
                            },
                            // spark-md5
                            {
                                name: 'spark-md5',
                                test: /node_modules\/spark-md5\//
                            },
                            // crypto-js
                            {
                                name: 'crypto-js',
                                test: /node_modules\/crypto-js\//
                            }
                        ]
                    }
                }
            }
        },

        server: {
            open: true,
            host: '0.0.0.0',
            port: 5180,
            proxy: {
                '/api/': {
                    target: env.VITE_APP_BASE_URL + '/',
                    changeOrigin: true,
                    ws: true
                }
            }
        },
        plugins: [
            vue(),

            // vueJsx(),
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
    }
})
