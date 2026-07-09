<template>
    <div class="layout-default flex h-screen">
        <div class="app-aside" :style="{ width: MainLayoutWidth }">
            <LayoutSidebar />
        </div>

        <div class="flex-1 flex flex-col min-w-0">
            <div class="app-header">
                <LayoutHeader />
            </div>
            <div class="app-main flex-1 min-h-0">
                <LayoutMain />
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { computed, watch, ref, onMounted, onUnmounted } from 'vue'
import LayoutMain from './components/main.vue'
import LayoutSidebar from './components/sidebar/index.vue'
import LayoutHeader from './components/header/index.vue'

import useSettingStore from '@/stores/modules/setting'
import useAppStore from '@/stores/modules/app'
import { initGlobalWs, destroyGlobalWs } from '@/hooks/useGlobalWs'

defineOptions({
    name: 'LayoutDefault'
})

const appStore = useAppStore()
const showMenuDrawer = computed(() => {
    if (appStore.isMobile) {
        return false
    } else {
        return appStore.isCollapsed
    }
})
const settingStore = useSettingStore()

const MainLayoutWidth = ref('auto')
watch(
    () => showMenuDrawer.value,
    () => {
        if (!appStore.isMobile) {
            if (appStore.isCollapsed) {
                MainLayoutWidth.value = `50px`
            } else {
                MainLayoutWidth.value = `${settingStore.sideWidth}px`
            }
        }

        setTimeout(() => {
            MainLayoutWidth.value = 'auto'
        }, 600)
    }
)

// 全局 WebSocket：登录后初始化，离开布局时销毁
onMounted(() => {
    initGlobalWs()
})
onUnmounted(() => {
    destroyGlobalWs()
})
</script>

<style scoped lang="scss">
.app-header {
    border-bottom: 8px solid var(--el-bg-color-page);
}
</style>
