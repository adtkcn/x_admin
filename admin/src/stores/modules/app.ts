import { getConfig } from '@/api/app'
import { defineStore } from 'pinia'
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
            return url ? `${this.config.ossDomain}${url}` : ''
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
