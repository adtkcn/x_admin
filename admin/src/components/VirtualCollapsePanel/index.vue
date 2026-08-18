<!-- 虚拟折叠面板 -->
<template>
    <DynamicScroller
        ref="scrollbarRef"
        class="scroller"
        :items="list"
        :min-item-size="props.minItemSize"
        :key-field="props.keyField"
        v-slot="{ item, active }"
        style="height: 100%"
        @scroll="onScroll"
    >
        <DynamicScrollerItem :item="item" :active="active" :size-dependencies="[item.__expanded]">
            <div class="collapse-header" @click="toggleExpand(item)">
                <div class="collapse-header-left">
                    <slot name="header" :item="item">{{ item[props.nameField] }}</slot>
                </div>

                <div class="collapse-header-right" v-if="props.showArrowRight">
                    <el-icon
                        class="arrow-right"
                        :class="{ 'rotate-icon': item.__expanded }"
                        :size="12"
                        :color="'#c1c1c1'"
                    >
                        <ArrowRight />
                    </el-icon>
                </div>
            </div>

            <Transition name="fade" mode="out-in">
                <div class="collapse-content" v-show="item.__expanded">
                    <slot name="content" :item="item"></slot>
                </div>
            </Transition>
        </DynamicScrollerItem>
    </DynamicScroller>
</template>
<script setup lang="ts">
import { ref, computed, onBeforeUnmount, onDeactivated, useTemplateRef } from 'vue'
import type { PropType } from 'vue'
import { ArrowRight } from '@element-plus/icons-vue'
import 'vue-virtual-scroller/dist/vue-virtual-scroller.css'
import { DynamicScroller, DynamicScrollerItem } from 'vue-virtual-scroller'
type ListItem = Record<string, any> & { __expanded?: boolean }

const emit = defineEmits<{
    scroll: [e: Event]
}>()
const props = defineProps({
    data: {
        type: Array as PropType<ListItem[]>,
        default: () => []
    },
    keyField: {
        type: String,
        default: 'id'
    },
    nameField: {
        type: String,
        default: 'name'
    },
    /** 每个项的最小高度 */
    minItemSize: {
        type: Number,
        default: 48
    },
    showArrowRight: {
        type: Boolean,
        default: true
    }
})

const itemHeight = computed(() => props.minItemSize + 'px')

const activeItem = ref<Record<string, any> | null>(null)
const list = computed<ListItem[]>(() => {
    console.time('list')
    const activeItemId = activeItem.value ? activeItem.value[props.keyField] : null
    const newList = props.data.map((listItem) => {
        return {
            ...listItem,
            __expanded: !!(activeItemId && listItem[props.keyField] === activeItemId)
        }
    })
    console.timeEnd('list')
    return newList
})

const scrollbarRef = useTemplateRef('scrollbarRef')
const scrollbarScrollTop = ref(0)
function onScroll(e: Event) {
    scrollbarScrollTop.value = (e.target as HTMLElement).scrollTop
    emit('scroll', e)
}
/**
 * 切换折叠面板展开状态
 * @param item 要切换展开状态的项
 */
function toggleExpand(item: ListItem) {
    activeItem.value = item.__expanded ? null : item
}
/**
 * 滚动到指定项
 * @param item 要滚动到的项
 */
function scrollToItem(item: ListItem) {
    if (!item) return
    activeItem.value = item
    const findIndex = list.value.findIndex((i) => i[props.keyField] === item[props.keyField])
    if (findIndex === -1) {
        return
    }
    const top = props.minItemSize * findIndex
    // 取消之前的动画
    cancelAnimationFrame(AnimationId.value as number)
    scrollTo(scrollbarScrollTop.value, top, scrollbarScrollTop.value)
}
/**
 * 滚动到指定索引项
 * @param index 要滚动到的索引项
 */
function scrollToIndex(index: number) {
    if (index < 0 || index >= list.value.length) return

    activeItem.value = list.value[index]

    const top = props.minItemSize * index
    // 取消之前的动画
    cancelAnimationFrame(AnimationId.value as number)
    scrollTo(scrollbarScrollTop.value, top, scrollbarScrollTop.value)
}
/**
 * 滚动到指定位置
 * @param position 要滚动到的位置
 */
function scrollToPosition(position: number) {
    if (position < 0) return
    // 取消之前的动画
    cancelAnimationFrame(AnimationId.value as number)
    scrollTo(scrollbarScrollTop.value, position, scrollbarScrollTop.value)
}

// 滚动到指定位置
const AnimationId = ref<number>()
function scrollTo(from: number, to: number, current: number) {
    if (scrollbarScrollTop.value === to) {
        return
    }
    const speed = (to - from) / 30

    if (speed < 0) {
        if (current <= to) {
            return
        }
    } else if (speed > 0) {
        if (current >= to) {
            return
        }
    }
    current = current + speed

    // DynamicScroller内部用的RecycleScroller组件，所以滚动到指定位置用RecycleScroller方法
    ;(scrollbarRef.value as any)?.$refs?.scroller?.scrollToPosition(current)
    AnimationId.value = requestAnimationFrame(() => {
        scrollTo(from, to, current)
    })
}
onDeactivated(() => {
    // 组件销毁时取消动画
    cancelAnimationFrame(AnimationId.value as number)
})
onBeforeUnmount(() => {
    // 组件销毁时取消动画
    cancelAnimationFrame(AnimationId.value as number)
})
defineExpose({
    toggleExpand,
    scrollToItem,
    scrollToIndex,
    scrollToPosition
})
</script>
<style lang="scss" scoped>
.collapse-header {
    display: flex;
    flex-direction: row;
    align-items: center;
    width: 100%;
    overflow: hidden;

    height: v-bind(itemHeight);
    padding: 0 4px;
    cursor: pointer;
    border-bottom: 1px solid rgba(0, 0, 0, 0.1);
    box-sizing: border-box;
    .collapse-header-left {
        font-size: 12px;
        color: #303133;
        flex: 1;
    }
}
.collapse-content {
    transform-origin: top;
    border-bottom: 1px solid rgba(0, 0, 0, 0.1);
}
// collapse-content显示隐藏动画
.fade-enter-active,
.fade-leave-active {
    transition: all 0.2s ease;
}

.fade-enter-from,
.fade-leave-to {
    opacity: 0;
    transform: scaleY(0);
}
.arrow-right {
    transition: all 0.2s ease;
}
.rotate-icon {
    transform: rotate(90deg);
}
::-webkit-scrollbar-thumb {
    background: rgba(0, 0, 0, 0.1) !important;
    border-radius: 4px;
}

::-webkit-scrollbar {
    width: 8px;
    background-color: #f5f5f5;
}
</style>
