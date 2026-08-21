import { getConfig } from '@/api/app'
import { defineStore } from 'pinia'
import { nextTick } from 'vue'
interface AppSate {
    config: Record<string, any>
    isMobile: boolean
    isCollapsed: boolean
    isRouteShow: boolean
}

const useAppStore = defineStore('app', {
    state: (): AppSate => {
        return {
            config: {},
            isMobile: true,
            isCollapsed: false,
            isRouteShow: true
        }
    },
    actions: {
        getImageUrl(url: string) {
            if (!url) return ''
            // 兼容：id（不含 / 与 :）→ 拼 /api/uploads/<id>，由文件流路由返回物理文件
            if (!url.includes('/') && !url.includes(':')) {
                return `${this.config.ossDomain}/api/uploads/${url}`
            }
            // 兼容：绝对 URL（http://、https://）原样返回
            if (url.startsWith('http://') || url.startsWith('https://')) {
                return url
            }
            // 兼容：相对路径（如 /api/static/backend_avatar.png、/api/uploads/abc）原样拼接 ossDomain
            return `${this.config.ossDomain}${url}`
        },
        getConfig() {
            return new Promise((resolve, reject) => {
                getConfig()
                    .then((data) => {
                        this.config = data
                        resolve(data)
                    })
                    .catch((err) => {
                        reject(err)
                    })
            })
        },
        setMobile(value: boolean) {
            this.isMobile = value
        },
        toggleCollapsed(toggle?: boolean) {
            this.isCollapsed = toggle ?? !this.isCollapsed
        },
        refreshView() {
            this.isRouteShow = false
            nextTick(() => {
                this.isRouteShow = true
            })
        }
    }
})

export default useAppStore
