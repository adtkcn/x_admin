import { fileURLToPath, URL } from 'url'

import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueJsx from '@vitejs/plugin-vue-jsx'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers'
import { createStyleImportPlugin, ElementPlusResolve } from 'vite-plugin-style-import'
import { createSvgIconsPlugin } from 'vite-plugin-svg-icons'

// https://vitejs.dev/config/
export default ({ mode }) => {
    const env = loadEnv(mode, process.cwd())
    console.log(env)

    return defineConfig({
        // base: '/admin/',
        server: {
            open: true,
            host: '0.0.0.0',
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
            AutoImport({
                imports: ['vue', 'vue-router'],
                resolvers: [ElementPlusResolver()],
                eslintrc: {
                    enabled: true
                }
            }),
            Components({
                directoryAsNamespace: true,
                resolvers: [ElementPlusResolver()]
            }),
            createStyleImportPlugin({
                resolves: [ElementPlusResolve()]
            }),
            createSvgIconsPlugin({
                // 配置路劲在你的src里的svg存放文件
                iconDirs: [fileURLToPath(new URL('./src/assets/icons', import.meta.url))],
                symbolId: 'local-icon-[dir]-[name]'
            })
        ],
        resolve: {
            alias: {
                '@': fileURLToPath(new URL('./src', import.meta.url))
            }
        }
    })
}
