import { defineStore } from 'pinia'
import cache from '@/utils/cache'
import type { RouteRecordRaw } from 'vue-router'
import {
    getUserInfo,
    login,
    logout,
    getMenu,
    type type_system_login,
    type type_system_admin_self
} from '@/api/system/user'
import router, { filterAsyncRoutes } from '@/router'
import { TOKEN_KEY } from '@/enums/cacheEnums'
import { PageEnum } from '@/enums/pageEnum'
import { clearAuthInfo, getToken } from '@/utils/auth'
import type { type_system_menu_resp } from '@/api/perms/menu'
export interface UserState {
    token: string
    userInfo: type_system_admin_self['user']
    routes: RouteRecordRaw[]
    menu: type_system_menu_resp[]
    perms: string[]
}

const useUserStore = defineStore('user', {
    state: (): UserState => {
        return {
            token: getToken() || '',
            // 用户信息
            userInfo: {} as type_system_admin_self['user'],
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
            this.userInfo = {} as type_system_admin_self['user']
            this.perms = []
        },
        login(info: type_system_login): Promise<string> {
            return new Promise((resolve, reject) => {
                login(info)
                    .then((data) => {
                        this.token = data.token
                        cache.set(TOKEN_KEY, data.token)
                        resolve(data.token)
                    })
                    .catch((error: Error) => {
                        reject(error)
                    })
            })
        },
        logout(): Promise<void> {
            return new Promise((resolve, reject) => {
                logout()
                    .then(async () => {
                        this.token = ''
                        await router.push(PageEnum.LOGIN)
                        clearAuthInfo()
                        resolve()
                    })
                    .catch((error: Error) => {
                        reject(error)
                    })
            })
        },
        getUserInfo(): Promise<type_system_admin_self> {
            return new Promise((resolve, reject) => {
                getUserInfo()
                    .then((data: type_system_admin_self) => {
                        this.userInfo = data.user
                        const permissions: string[] = []
                        data.permissions &&
                            data.permissions.forEach((item: string) => {
                                if (item) {
                                    item.split(',').forEach((perm: string) => {
                                        permissions.push(perm)
                                    })
                                }
                            })
                        this.perms = permissions
                        resolve(data)
                    })
                    .catch((error: Error) => {
                        reject(error)
                    })
            })
        },
        getMenu(): Promise<type_system_menu_resp[]> {
            return new Promise((resolve, reject) => {
                getMenu()
                    .then((data) => {
                        this.menu = data
                        this.routes = filterAsyncRoutes(data)
                        resolve(data)
                    })
                    .catch((error: Error) => {
                        reject(error)
                    })
            })
        }
    }
})

export default useUserStore
