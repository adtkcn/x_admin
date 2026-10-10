<template>
    <div class="app-tabs flex">
        <div class="flex-1 min-w-0 select-none">
            <el-tabs
                :model-value="currentTab"
                :closable="tabsLists.length > 1"
                @tab-change="handleChange"
                @tab-remove="removeTab($event)"
            >
                <el-tab-pane v-for="item in tabsLists" :key="item.fullPath" :name="item.fullPath">
                    <template #label>
                        <span @contextmenu.prevent="openContextMenu(item.fullPath, $event)">
                            {{ item.title }}
                        </span>
                    </template>
                </el-tab-pane>
            </el-tabs>
        </div>

        <!-- 标签右键上下文菜单：关闭当前 / 关闭其他 / 关闭全部 -->
        <Teleport to="body">
            <ul
                v-show="menuVisible"
                class="tab-context-menu"
                :style="{ left: `${menuLeft}px`, top: `${menuTop}px` }"
                @click.stop
            >
                <li @click="handleMenuCommand('closeCurrent')">关闭当前</li>
                <li @click="handleMenuCommand('closeOther')">关闭其他</li>
                <li @click="handleMenuCommand('closeAll')">关闭全部</li>
            </ul>
        </Teleport>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted, nextTick } from 'vue'
import { useRouter } from 'vue-router'
import useMultipleTabs from '@/hooks/useMultipleTabs'
import { useWatchRoute } from '@/hooks/useWatchRoute'
import useTabsStore, { getRouteParams } from '@/stores/modules/multipleTabs'
const router = useRouter()
const tabsStore = useTabsStore()
const { removeOtherTab, addTab, removeAllTab, removeTab, tabsLists, currentTab } = useMultipleTabs()
useWatchRoute(() => {
    addTab()
})

const handleChange = (fullPath: any) => {
    const tabItem = tabsStore.tasMap[fullPath]
    router.push(getRouteParams(tabItem))
}

// 标签右键菜单状态：记录被右键的 fullPath 与鼠标定位
const menuVisible = ref(false)
const menuLeft = ref(0)
const menuTop = ref(0)
const menuTarget = ref('')

const closeContextMenu = () => {
    menuVisible.value = false
}

const openContextMenu = (fullPath: string, e: MouseEvent) => {
    menuTarget.value = fullPath
    menuLeft.value = e.clientX
    menuTop.value = e.clientY
    menuVisible.value = true
}

const handleMenuCommand = async (command: string) => {
    const fullPath = menuTarget.value
    switch (command) {
        case 'closeCurrent':
            // removeTab 支持传入具体 fullPath，关闭非激活标签时不会误跳转
            removeTab(fullPath)
            break
        case 'closeOther':
            // 关闭其他：若右键的不是当前激活标签，先切过去，保证保留的是被右键的标签
            if (fullPath !== currentTab.value) {
                await router.push(getRouteParams(tabsStore.tasMap[fullPath]))
                await nextTick()
            }
            removeOtherTab()
            break
        case 'closeAll':
            removeAllTab()
            break
    }
    closeContextMenu()
}

// 点击页面任意处或窗口尺寸变化时收起右键菜单
onMounted(() => {
    document.addEventListener('click', closeContextMenu)
    window.addEventListener('resize', closeContextMenu)
})
onUnmounted(() => {
    document.removeEventListener('click', closeContextMenu)
    window.removeEventListener('resize', closeContextMenu)
})
</script>
<style lang="scss" scoped>
.app-tabs {
    border-top: 1px solid rgba($color: #ccc, $alpha: 0.4);
    padding: 0 12px;
    :deep(.el-tabs) {
        height: 40px;
        .el-tabs {
            &__header {
                margin-bottom: 0;
            }
            &__content {
                display: none;
            }
            &__nav-next,
            &__nav-prev {
                font-size: var(--el-font-size-large);
            }

            &__nav-wrap::after {
                height: 0;
            }
            &__item {
                font-weight: normal;
                padding: 0 15px !important;
                box-sizing: border-box;
                &.is-active {
                    background-color: var(--el-color-primary-light-9);
                    &::before {
                        content: '';
                        display: inline-block;
                        width: 6px;
                        height: 6px;
                        background-color: var(--el-color-primary);
                        margin-right: 6px;
                        border-radius: 50%;
                        vertical-align: 2px;
                    }
                    &::after {
                        position: absolute;
                        content: '';
                        display: block;
                        top: 0;
                        height: 2px;
                        left: 0;
                        width: 100%;
                        background-color: var(--el-color-primary);
                    }
                }
                .is-icon-close {
                    color: var(--el-text-color-regular);
                    vertical-align: -2px;
                    &:hover {
                        color: var(--color-white);
                        background-color: var(--el-color-danger);
                    }
                }
            }
            &__active-bar {
                display: none;
            }
        }
    }
}
.tab-context-menu {
    position: fixed;
    z-index: 3000;
    min-width: 120px;
    margin: 0;
    padding: 5px 0;
    list-style: none;
    font-size: 14px;
    color: var(--el-text-color-regular);
    background-color: var(--el-bg-color-overlay, #fff);
    border: 1px solid var(--el-border-color-light);
    border-radius: 4px;
    box-shadow: var(--el-box-shadow-light);
    li {
        padding: 7px 16px;
        cursor: pointer;
        &:hover {
            color: var(--el-color-primary);
            background-color: var(--el-fill-color-light);
        }
    }
}
</style>
