<template>
    <div>
        <div class="flex items-center gap-2 mb-3">
            <el-date-picker
                v-model="range"
                type="datetimerange"
                size="small"
                :clearable="false"
                :default-time="defaultTime"
                range-separator="至"
                start-placeholder="开始时间"
                end-placeholder="结束时间"
            />
            <el-select
                v-model="picked"
                multiple
                collapse-tags
                size="small"
                class="flex-1"
                placeholder="选择属性"
            >
                <el-option v-for="opt in attrOptions" :key="opt.value" :label="opt.label" :value="opt.value" />
            </el-select>
            <el-button type="primary" size="small" @click="load()">查询</el-button>
        </div>
        <div ref="chartRef" v-loading="loading" class="trend-chart"></div>
    </div>
</template>

<script lang="ts" setup>
import { ref, computed, shallowRef, onMounted, onBeforeUnmount, watch } from 'vue'
import { deviceData, type type_device_data } from '@/api/iot'
import { echarts } from '@/utils/echart'
import { isNumericType, type ThingProperty } from '@/views/iot/product/components/thing-model'

const props = defineProps<{
    deviceId: string
    thingProps: ThingProperty[]
    // 可选初始区间：传入则以其为默认查询区间（如告警时间段前后延伸），不传沿用「当天」默认
    initialRange?: [Date, Date]
    // 高亮时间段 [startMs, endMs]：在图表中以半透明色带标识告警持续时间
    highlightRange?: [number, number]
}>()

// 区间必填、不可清空；默认取 initialRange，否则当天 00:00:00 ~ 23:59:59
const today = new Date()
const range = ref<[Date, Date]>(
    props.initialRange ?? [
        new Date(today.getFullYear(), today.getMonth(), today.getDate(), 0, 0, 0),
        new Date(today.getFullYear(), today.getMonth(), today.getDate(), 23, 59, 59)
    ]
)
// el-date-picker 的 default-time：面板里选完日期后时间部分的默认值（起 00:00:00，止 23:59:59）
const defaultTime: [Date, Date] = [
    new Date(2000, 0, 1, 0, 0, 0),
    new Date(2000, 0, 1, 23, 59, 59)
]
const trendRows = ref<type_device_data[]>([])
const picked = ref<string[]>([])
const loading = ref(false)
const chartRef = ref<HTMLElement>()
const chart = shallowRef<ReturnType<typeof echarts.init>>()

// 可勾选属性完全以产品物模型为准，只取其数值型属性（趋势曲线需可绘制），展示中文名
const thingNumeric = computed(() => props.thingProps.filter((p) => isNumericType(p.dataType)))
const attrOptions = computed(() =>
    thingNumeric.value.map((p) => ({
        value: p.identifier,
        label: p.name === p.identifier ? p.identifier : `${p.name}（${p.identifier}）`
    }))
)
const nameOf = (id: string) => props.thingProps.find((p) => p.identifier === id)?.name || id

// lastTs 已持有数据的最大时间戳（滚动续查的增量起点）；rolling 标记本次区间是否处于持续中（结束时间晚于当前）
let lastTs = 0
let rolling = false
// 上次渲染的勾选属性集合（结构指纹）：与其一致说明仅数据增长，走增量合并而非全量重建
let renderedKey = ''
const maxTs = (rows: type_device_data[]) => rows.reduce((m, r) => (r.ts > m ? r.ts : m), 0)

// silent=滚动续查（静默，不闪 loading）；手动/首次查询整段回读，并按「结束时间是否晚于当前」决定是否开启滚动
const load = async (silent = false) => {
    if (!props.deviceId || !range.value) return
    if (!silent) loading.value = true
    try {
        const [s, e] = range.value
        // 滚动续查：起点取已持有数据的最后时间之后、终点取当前时刻，仅追加增量，不整段重查
        const startMs = silent ? (lastTs > 0 ? lastTs + 1 : s.getTime()) : s.getTime()
        const endMs = silent ? Date.now() : e.getTime()
        // 不传 limit：区间内回读（后端 limit<=0 即不限制条数）
        const rows = await deviceData({ id: props.deviceId, start_ms: startMs, end_ms: endMs })
        if (silent) {
            if (rows.length) {
                trendRows.value = [...trendRows.value, ...rows]
                lastTs = Math.max(lastTs, maxTs(rows))
            }
        } else {
            trendRows.value = rows
            lastTs = maxTs(rows)
            rolling = e.getTime() > Date.now()
        }
        if (!picked.value.length) picked.value = attrOptions.value.slice(0, 4).map((o) => o.value)
    } finally {
        if (!silent) loading.value = false
    }
    if (!silent) startLive()
}

const render = () => {
    if (!chartRef.value) return
    if (!chart.value || chart.value.getDom() !== chartRef.value) {
        chart.value?.dispose()
        chart.value = echarts.init(chartRef.value)
        renderedKey = '' // 新实例：结构需整体重建
    }
    // 区间内无数据：留一张空图并居中提示，避免图表区域一片空白
    if (!trendRows.value.length) {
        chart.value.setOption(
            {
                title: {
                    text: '暂无数据',
                    left: 'center',
                    top: 'middle',
                    textStyle: { color: '#909399', fontSize: 14 }
                }
            },
            true
        )
        renderedKey = '' // 空态后若转为有数据，强制整体重建
        return
    }
    const rows = [...trendRows.value].sort((a, b) => a.ts - b.ts)
    const series: any[] = picked.value.map((k) => ({
        name: nameOf(k),
        type: 'line',
        showSymbol: false,
        data: rows
            .filter((r) => r.values && k in r.values)
            .map((r) => [r.ts, Number(r.values![k])])
    }))
    // 告警时间段高亮：把 markArea 挂到第一个有数据的 series 上
    if (props.highlightRange) {
        const target = series.find((s) => s.data.length > 0)
        if (target) {
            // 滚动持续中（结束时间晚于当前），高亮带右端跟随当前时刻延展
            const bandEnd = rolling ? Date.now() : props.highlightRange[1]
            target.markArea = {
                silent: true,
                itemStyle: { color: 'rgba(245, 108, 108, 0.15)' },
                label: {
                    show: true,
                    position: 'insideTop',
                    color: '#f56c6c',
                    fontSize: 11,
                    formatter: '告警期间'
                },
                data: [[{ xAxis: props.highlightRange[0] }, { xAxis: bandEnd }]]
            }
        }
    }
    // 结构未变（勾选属性集合一致）：多为滚动追加数据，仅按 series 合并更新——曲线平滑延伸、
    // 保留 dataZoom 缩放态，避免整图重建导致的闪烁；结构变化或首次才整体重建
    const key = picked.value.join(',')
    if (key === renderedKey) {
        chart.value.setOption({ series })
    } else {
        chart.value.setOption(
            {
                tooltip: { trigger: 'axis' },
                legend: { data: picked.value.map(nameOf), top: 5, left: 'center' },
                grid: { left: 50, right: 20, top: 36, bottom: 68 },
                // inside：图上滚轮/拖拽缩放；slider：底部可视化缩放轴（可拖拽区间、看缩略曲线）
                dataZoom: [{ type: 'inside' }, { type: 'slider', height: 22, bottom: 12, brushSelect: false }],
                xAxis: { type: 'time' },
                yAxis: { type: 'value', scale: true },
                series
            },
            true
        )
        renderedKey = key
    }
}

// 勾选属性或区间数据变化即重绘（同一 tick 内多处改动只触发一次），切换属性无需重新请求
watch([picked, trendRows], render)

// 滚动更新：区间结束时间晚于当前（视为持续中）时定时静默续查增量；重新查询/卸载即停止
const LIVE_POLL_MS = 5000
let liveTimer: ReturnType<typeof setInterval> | null = null
const startLive = () => {
    stopLive()
    if (!rolling) return
    liveTimer = setInterval(() => load(true), LIVE_POLL_MS)
}
const stopLive = () => {
    if (liveTimer) {
        clearInterval(liveTimer)
        liveTimer = null
    }
}

const resize = () => chart.value?.resize()
onMounted(() => {
    load()
    window.addEventListener('resize', resize)
})
onBeforeUnmount(() => {
    stopLive()
    window.removeEventListener('resize', resize)
    chart.value?.dispose()
    chart.value = undefined
})
</script>

<style scoped>
.trend-chart {
    height: 360px;
    width: 100%;
}
</style>
