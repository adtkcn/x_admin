<template>
    <div class="rc-sidebar">
        <!-- <div class="rc-sidebar__title">节点</div> -->
        <div
            class="rc-node"
            v-for="item of dragNodes"
            :key="item.type"
            :title="item.desc"
            @mousedown.stop.prevent="dragInNode(item.type, item.label)"
        >
            <span class="rc-node__dot" :style="{ background: item.color }"></span>
            <span class="rc-node__label">{{ item.label }}</span>
        </div>
    </div>
</template>

<script setup lang="ts">
import { RC_NODE_META } from './rule.type'

defineOptions({ name: 'RcDiagramSidebar' })
const emit = defineEmits(['dragInNode'])
// 「输入」是固定起点，由设计器自动放置，不进拖拽列表（保证入口唯一）
const dragNodes = RC_NODE_META.filter((m) => m.type !== 'rule:trigger')
function dragInNode(type: string, text: string) {
    emit('dragInNode', type, text)
}
</script>

<style scoped>
.rc-sidebar {
    user-select: none;
    overflow: auto;
    padding: 8px 6px;
}
/* .rc-sidebar__title {
    font-size: 12px;
    color: var(--el-text-color-secondary);
    text-align: center;
    margin-bottom: 8px;
} */
.rc-node {
    display: flex;
    align-items: center;
    gap: 6px;
    padding: 6px 8px;
    margin-bottom: 4px;
    border-radius: 4px;
    cursor: grab;
    transition: background 0.15s;
}
.rc-node:hover {
    background: var(--el-fill-color-light);
}
.rc-node__dot {
    width: 10px;
    height: 10px;
    border-radius: 50%;
    flex-shrink: 0;
}
.rc-node__label {
    font-size: 12px;
}
</style>
