<template>
    <div class="layout-default flex h-screen w-full">
        <div class="app-aside" :style="{ width: MainLayoutWidth }">
            <layout-sidebar />
        </div>

        <div class="flex-1 flex flex-col min-w-0">
            <div class="app-header">
                <layout-header />
            </div>
            <div class="app-main flex-1 min-h-0">
                <layout-main />
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { computed, watch, ref } from 'vue'
// 节流
// import { throttle } from 'lodash-es'
import LayoutMain from './components/main.vue'
import LayoutSidebar from './components/sidebar/index.vue'
import LayoutHeader from './components/header/index.vue'

import useSettingStore from '@/stores/modules/setting'
import useAppStore from '@/stores/modules/app'

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
</script>
