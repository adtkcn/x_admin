<template>
    <el-image :style="styles" v-bind="props"> </el-image>
</template>

<script lang="ts" setup>
import { computed } from 'vue'
import type { CSSProperties } from 'vue'
import { addUnit } from '@/utils/util'
import { imageProps } from 'element-plus'
const props = defineProps({
    width: {
        type: [String, Number],
        default: 'auto'
    },
    height: {
        type: [String, Number],
        default: 'auto'
    },
    radius: {
        type: [String, Number],
        default: 0
    },
    ...imageProps,
    hideOnClickModal: {
        type: Boolean,
        default: true
    }
})

const styles = computed<CSSProperties>(() => {
    return {
        width: addUnit(props.width),
        height: addUnit(props.height),
        borderRadius: addUnit(props.radius)
    }
})
</script>

<style lang="scss" scoped>
.el-image {
    display: block;
    .el-image__error {
        font-size: var(--el-font-size-extra-small);
    }
}
</style>
