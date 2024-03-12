import { defineStore } from 'pinia'
import cache from '@/utils/cache'
import type { RouteRecordRaw } from 'vue-router'
import { getUserInfo, login, logout, getMenu } from '@/api/user'
import router, { filterAsyncRoutes } from '@/router'
import { TOKEN_KEY } from '@/enums/cacheEnums'
import { PageEnum } from '@/enums/pageEnum'
import { clearAuthInfo, getToken } from '@/utils/auth'
export interface UserState {
    token: string
    userInfo: Record<string, any>
    routes: RouteRecordRaw[]
    menu: any[]
    perms: string[]
}

const useUserStore = defineStore('user', {
    state: (): UserState => {
        return {
            token: getToken() || '',
            // 用户信息
            userInfo: {},
            // 路由
            routes: [],
            menu: [],
            // 权限
            perms: []
        }
    },
    getters: {},
    actions: {
        resetState() {
            this.token = ''
            this.userInfo = {}
            this.perms = []
        },
        login(playload: any) {
            return new Promise((resolve, reject) => {
                login({
                    ...playload
                })
                    .then((data) => {
                        this.token = data.token
                        cache.set(TOKEN_KEY, data.token)
                        resolve(data)
                        // reject(data)
                    })
                    .catch((error) => {
                        reject(error)
                    })
            })
        },
        logout() {
            return new Promise((resolve, reject) => {
                logout()
                    .then(async (data) => {
                        this.token = ''
                        await router.push(PageEnum.LOGIN)
                        clearAuthInfo()
                        resolve(data)
                    })
                    .catch((error) => {
                        reject(error)
                    })
            })
        },
        getUserInfo() {
            return new Promise((resolve, reject) => {
                getUserInfo()
                    .then((data) => {
                        this.userInfo = data.user
                        const permissions = []
                        data.permissions.forEach((item: any) => {
                            if (item) {
                                item.split(',').forEach((item: any) => {
                                    permissions.push(item)
                                })
                            }
                        })
                        this.perms = permissions
                        resolve(data)
                    })
                    .catch((error) => {
                        reject(error)
                    })
            })
        },
        getMenu() {
            return new Promise((resolve, reject) => {
                getMenu()
                    .then((data) => {
                        this.menu = data
                        this.routes = filterAsyncRoutes(data)
                        resolve(data)
                    })
                    .catch((error) => {
                        reject(error)
                    })
            })
        }
    }
})

export default useUserStore
