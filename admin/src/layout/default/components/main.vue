<template>
    <main class="h-full bg-page pl-3 pr-3 pb-3 pt-1 box-border overflow-auto">
        <router-view v-if="isRouteShow" v-slot="{ Component, route }">
            <keep-alive :include="includeList" :max="20">
                <component :is="Component" :key="route.fullPath" />
            </keep-alive>
        </router-view>
    </main>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { RouterView } from 'vue-router'
import useAppStore from '@/stores/modules/app'
import useTabsStore from '@/stores/modules/multipleTabs'
import useSettingStore from '@/stores/modules/setting'
defineOptions({
    name: 'LayoutDefaultMain'
})

const appStore = useAppStore()
const tabsStore = useTabsStore()
const settingStore = useSettingStore()
const isRouteShow = computed(() => appStore.isRouteShow)
const includeList = computed(() => (settingStore.openMultipleTabs ? tabsStore.getCacheTabList : []))
</script>
