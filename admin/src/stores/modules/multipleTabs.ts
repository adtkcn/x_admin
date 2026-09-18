import { unref } from 'vue'
import { defineStore } from 'pinia'
import { isExternal } from '@/utils/validate'
import type {
    LocationQuery,
    RouteLocationNormalized,
    RouteParamsRaw,
    Router,
    RouteRecordName
} from 'vue-router'
import { PageEnum } from '@/enums/pageEnum'

// 多标签页（tabs）状态管理：维护已打开的标签列表与 keep-alive 组件缓存

// 单个标签项，字段直接对应可恢复的路由信息
interface TabItem {
    name: RouteRecordName // 路由名称
    fullPath: string // 含 query 的完整路径，作为标签唯一标识
    path: string // 路径（不含 query）
    title?: string // 标签标题
    query?: LocationQuery // 查询参数
    params?: RouteParamsRaw // 路由参数
}

interface TabsSate {
    cacheTabList: Set<string> // 需 keep-alive 缓存的组件名集合
    tabList: TabItem[] // 当前打开的标签列表（有序）
    tasMap: Record<string, TabItem> // fullPath -> 标签，便于快速查找
    indexRouteName: RouteRecordName // 首页路由名称（用于"关闭全部"时判断）
}

// 按 fullPath 查找标签下标，未找到返回 -1
const getHasTabIndex = (fullPath: string, tabList: TabItem[]) => {
    return tabList.findIndex((item) => item.fullPath == fullPath)
}

// 判断路由是否不应作为标签打开：外链、隐藏标签、动态未注册、登录/403 页
const isCannotAddRoute = (route: RouteLocationNormalized, router: Router) => {
    const { path, meta, name } = route
    if (!path || isExternal(path)) return true
    if (meta?.hideTab) return true
    if (!router.hasRoute(name!)) return true
    if (([PageEnum.LOGIN, PageEnum.ERROR_403] as string[]).includes(path)) {
        return true
    }
    return false
}

// 按 fullPath 查找标签下标（与 getHasTabIndex 等价，用于移除场景）
const findTabsIndex = (fullPath: string, tabList: TabItem[]) => {
    return tabList.findIndex((item) => item.fullPath === fullPath)
}

// 取路由匹配到的最深一层组件名（keep-alive 缓存以此为准）
const getComponentName = (route: RouteLocationNormalized) => {
    return route.matched[route.matched.length - 1]?.components?.default?.name
}

// 将标签项转换为 router.push 可用的路由参数
export const getRouteParams = (tabItem: TabItem) => {
    const { params, path, query } = tabItem
    return {
        params: params || {},
        path,
        query: query || {}
    }
}

const useTabsStore = defineStore('tabs', {
    state: (): TabsSate => ({
        cacheTabList: new Set(),
        tabList: [],
        tasMap: {},
        indexRouteName: ''
    }),
    getters: {
        getTabList(): TabItem[] {
            return this.tabList
        },
        getCacheTabList(): string[] {
            return Array.from(this.cacheTabList)
        }
    },
    actions: {
        // 记录首页路由名称
        setRouteName(name: RouteRecordName) {
            this.indexRouteName = name
        },
        // 加入 keep-alive 缓存
        addCache(componentName?: string) {
            if (componentName) this.cacheTabList.add(componentName)
        },
        // 移出 keep-alive 缓存
        removeCache(componentName?: string) {
            if (componentName && this.cacheTabList.has(componentName)) {
                this.cacheTabList.delete(componentName)
            }
            console.log(this.cacheTabList)
        },
        // 清空全部缓存
        clearCache() {
            this.cacheTabList.clear()
        },
        // 重置整个 store（如退出登录）
        resetState() {
            this.cacheTabList = new Set()
            this.tabList = []
            this.tasMap = {}
            this.indexRouteName = ''
        },
        // 将当前路由加入标签：已存在则仅更新映射/缓存，不重复追加
        addTab(router: Router) {
            const route = unref(router.currentRoute)
            const { name, query, meta, params, fullPath, path } = route
            if (isCannotAddRoute(route, router)) return
            const hasTabIndex = getHasTabIndex(fullPath!, this.tabList)
            const componentName = getComponentName(route)
            const tabItem = {
                name: name!,
                path,
                fullPath,
                title: meta?.title,
                query,
                params
            }
            this.tasMap[fullPath] = tabItem
            if (meta?.keepAlive) {
                this.addCache(componentName)
            }
            if (hasTabIndex != -1) {
                return
            }

            this.tabList.push(tabItem)
        },
        // 移除指定标签；若移除的是当前激活标签，则跳转到相邻标签
        removeTab(fullPath: string, router: Router) {
            const { currentRoute, push } = router
            const index = findTabsIndex(fullPath, this.tabList)
            // 移除tab
            if (this.tabList.length > 1) {
                index !== -1 && this.tabList.splice(index, 1)
            }
            const componentName = getComponentName(currentRoute.value)
            this.removeCache(componentName)
            if (fullPath !== currentRoute.value.fullPath) {
                return
            }
            // 删除选中的tab
            let toTab: TabItem | null = null

            if (index === 0) {
                toTab = this.tabList[index]
            } else {
                toTab = this.tabList[index - 1]
            }

            const toRoute = getRouteParams(toTab)
            push(toRoute)
        },
        // 只保留指定路由对应的标签，关闭其余标签及其缓存
        removeOtherTab(route: RouteLocationNormalized) {
            this.tabList = this.tabList.filter((item) => item.fullPath == route.fullPath)
            const componentName = getComponentName(route)
            this.cacheTabList.forEach((name) => {
                if (componentName !== name) {
                    this.removeCache(name)
                }
            })
        },
        // 关闭全部标签：当前已在首页则等价于只保留首页，否则清空并跳回首页
        removeAllTab(router: Router) {
            const { push, currentRoute } = router
            const { name } = unref(currentRoute)
            if (name == this.indexRouteName) {
                this.removeOtherTab(currentRoute.value)
                return
            }
            this.tabList = []
            this.clearCache()
            push(PageEnum.INDEX)
        }
    }
})

export default useTabsStore
