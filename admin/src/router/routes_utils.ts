/**
 * 路由纯函数工具集
 *
 * 这些函数只依赖「视图扫描结果」与「菜单路由配置」，不引用 router 实例、userStore 等运行时状态，
 * 因此被单独拆出，以避免 router/index.ts 与 stores/modules/user 之间形成循环依赖。
 */
import { type RouteRecordRaw } from 'vue-router'
import { MenuEnum } from '@/enums/appEnums'
import { isExternal } from '@/utils/validate'
import { LAYOUT, Empty } from './routes'
import qs from 'query-string'

// 编译期通过 Vite glob 收集 /src/views 下所有 .vue 视图模块：key 为文件路径，value 为动态导入函数
const modules = import.meta.glob('/src/views/**/*.vue')

// 获取所有可作为路由的视图模块路径（剔除编辑页与公共 component 目录）
export function getModulesKey() {
    return Object.keys(modules)
        .filter((item) => !item.endsWith('edit.vue') && item.indexOf('/component/') == -1)
        .map((item) => item.replace('/src/views/', '').replace('.vue', ''))
}

// 将后端菜单树递归转换为 vue-router 可用的路由记录数组
export function filterAsyncRoutes(routes: any[], firstRoute = true) {
    return routes.map((route) => {
        const routeRecord = createRouteRecord(route, firstRoute)
        if (route.children != null && route.children && route.children.length) {
            routeRecord.children = filterAsyncRoutes(route.children, false)
        }
        return routeRecord
    })
}

// 根据单个菜单节点创建一条路由记录（RouteRecordRaw）
export function createRouteRecord(route: any, firstRoute: boolean): RouteRecordRaw {
    // 菜单配置中携带的 query 参数，先取原始值
    let query = route.params

    try {
        if (route.params) {
            query = decodeURIComponent(qs.stringify(JSON.parse(route.params)))
            console.log(query)
        }
    } catch (error) {
        // params 解析失败则忽略，query 保留原始值
    }
    // 菜单为后端动态结构，注入 meta 时类型无法静态校验
    // @ts-ignore
    const routeRecord: RouteRecordRaw = {
        path: isExternal(route.paths) ? route.paths : firstRoute ? `/${route.paths}` : route.paths,
        name: Symbol(route.paths),
        meta: {
            hidden: !route.is_show,
            keepAlive: !!route.is_cache,
            title: route.menu_name,
            perms: route.perms,
            query: query,
            icon: route.menu_icon,
            type: route.menu_type,
            activeMenu: route.selected
        }
    }
    // 根据菜单类型决定渲染组件：目录用布局/空组件包裹，普通菜单加载具体视图
    switch (route.menu_type) {
        case MenuEnum.CATALOGUE:
            routeRecord.component = firstRoute ? LAYOUT : Empty
            if (!route.children) {
                routeRecord.component = Empty
            }
            break
        case MenuEnum.MENU:
            routeRecord.component = loadRouteView(route.component)
            break
    }
    return routeRecord
}

// 根据组件路径从已扫描的视图模块中匹配并返回对应的动态导入函数（找不到则回退 Empty）
export function loadRouteView(component: string) {
    try {
        const key = Object.keys(modules).find((key) => {
            return key.includes(`${component}.vue`)
        })
        if (key) {
            return modules[key]
        }
        throw Error(`找不到组件${component}，请确保组件路径正确`)
    } catch (error) {
        console.error(error)
        return Empty
    }
}

// 递归查找第一个「可访问的普通菜单」路由名称，用于登录成功后的默认跳转
export function findFirstValidRoute(routes: RouteRecordRaw[]): string | undefined {
    for (const route of routes) {
        if (route.meta?.type == MenuEnum.MENU && !route.meta?.hidden && !isExternal(route.path)) {
            return route.name as string
        }
        if (route.children) {
            const name = findFirstValidRoute(route.children)
            if (name) {
                return name
            }
        }
    }
}
