<template>
    <aside class="sidebar h-full">
        <el-drawer
            v-if="isMobile"
            v-model="showMenuDrawer"
            direction="ltr"
            :size="drawerSize"
            title="主题设置"
            :with-header="false"
        >
            <Side />
        </el-drawer>
        <Side v-if="!isMobile" />
    </aside>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import Side from './side.vue'
import useAppStore from '@/stores/modules/app'
import useSettingStore from '@/stores/modules/setting'
defineOptions({
    name: 'LayoutDefaultSidebar'
})
const appStore = useAppStore()
const settingStore = useSettingStore()
const isMobile = computed(() => appStore.isMobile)
const showMenuDrawer = computed({
    get() {
        return !appStore.isCollapsed && isMobile.value
    },
    set(value) {
        appStore.toggleCollapsed(!value)
    }
})

const drawerSize = computed(() => {
    return `${settingStore.sideWidth + 1}px`
})
</script>

<style lang="scss" scoped>
.sidebar {
    :deep(.el-drawer__body) {
        padding: 0;
    }
}
</style>
