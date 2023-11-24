<template>
    <router-view v-if="isRouteShow" v-slot="{ Component, route }">
        <keep-alive :include="includeList" :max="20">
            <component :is="Component" :key="route.fullPath" />
        </keep-alive>
    </router-view>
</template>

<script setup lang="ts">
import useAppStore from '@/stores/modules/app'
import useTabsStore from '@/stores/modules/multipleTabs'
import useSettingStore from '@/stores/modules/setting'
const appStore = useAppStore()
const tabsStore = useTabsStore()
const settingStore = useSettingStore()
const isRouteShow = computed(() => appStore.isRouteShow)
const includeList = computed(() => (settingStore.openMultipleTabs ? tabsStore.getCacheTabList : []))
</script>

<style></style>
