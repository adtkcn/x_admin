import { createRouter, createWebHistory } from 'vue-router'
import { constantRoutes, INDEX_ROUTE_NAME } from './routes'
import useUserStore from '@/stores/modules/user'

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: constantRoutes
})

// 重置路由
export function resetRouter() {
    router.removeRoute(INDEX_ROUTE_NAME)
    const { routes } = useUserStore()
    routes.forEach((route) => {
        const name = route.name
        if (name && router.hasRoute(name)) {
            router.removeRoute(name)
        }
    })
}

export default router
