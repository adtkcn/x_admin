<!-- src/components/BaseChart.vue -->
<template>
    <div ref="chartRef" class="base-chart" :style="{ width, height }"></div>
</template>

<script setup lang="ts">
import { onMounted, onBeforeUnmount, watch, nextTick, ref, toRaw, shallowRef } from 'vue'
import { echarts, ECOption } from '@/utils/echart'

// Props
const props = withDefaults(
    defineProps<{
        option?: any
        width?: string
        height?: string
        loading?: boolean
        autoresize?: boolean // 是否自动 resize
    }>(),
    {
        width: '100%',
        height: '100%',
        loading: false,
        autoresize: true
    }
)

// Emits
const emit = defineEmits<{
    (e: 'chartInstance', instance: echarts.ECharts): void
}>()

// Refs
const chartRef = ref<HTMLDivElement | null>(null)
const chartInstance = shallowRef<echarts.ECharts | null>(null)

// 初始化图表
const initChart = () => {
    if (!chartRef.value) return

    // 销毁已有实例
    if (chartInstance.value) {
        chartInstance.value.dispose()
    }

    // 创建新实例
    chartInstance.value = echarts.init(chartRef.value)
    emit('chartInstance', chartInstance.value)

    // 监听 option 变化
    setOption(props.option)
}

// 设置配置项
const setOption = (option: ECOption) => {
    if (!chartInstance.value || !option) return
    chartInstance.value.setOption(toRaw(option), {
        notMerge: false, // 合并配置（可设为 true 实现增量更新）
        lazyUpdate: false
    })
}

// 处理 loading
watch(
    () => props.loading,
    (loading) => {
        if (chartInstance.value) {
            loading ? chartInstance.value.showLoading() : chartInstance.value.hideLoading()
        }
    }
)

// 响应 option 变化
watch(
    () => props.option,
    () => {
        nextTick(() => {
            setOption(props.option)
        })
    },
    { deep: true }
)

// 自动 resize
let resizeObserver: ResizeObserver | null = null
const handleResize = () => {
    chartInstance.value?.resize()
}

onMounted(() => {
    initChart()

    // 方式1：监听 window resize（简单场景）
    if (props.autoresize) {
        window.addEventListener('resize', handleResize)
    }

    // 方式2（更精准）：使用 ResizeObserver（推荐）
    if (typeof ResizeObserver !== 'undefined' && chartRef.value) {
        resizeObserver = new ResizeObserver(handleResize)
        resizeObserver.observe(chartRef.value)
    }
})

onBeforeUnmount(() => {
    if (chartInstance.value) {
        chartInstance.value.dispose()
    }
    window.removeEventListener('resize', handleResize)
    if (resizeObserver) {
        resizeObserver.disconnect()
    }
})
defineExpose({
    setOption
})
</script>

<style scoped>
.base-chart {
    width: 100%;
    height: 100%;
}
</style>
