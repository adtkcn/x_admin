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
            <el-button type="primary" size="small" @click="load">查询</el-button>
        </div>
        <div ref="chartRef" v-loading="loading" class="trend-chart"></div>
        <div v-if="trendRows.length" class="text-xs mt-1 text-gray-400">
            共取 {{ trendRows.length }} 条，曲线按时间正序绘制
        </div>
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

const load = async () => {
    if (!props.deviceId || !range.value) return
    loading.value = true
    try {
        // 不传 limit：区间内全量回读（后端 limit<=0 即不限制条数）
        trendRows.value = await deviceData({
            id: props.deviceId,
            start_ms: range.value[0].getTime(),
            end_ms: range.value[1].getTime()
        })
        if (!picked.value.length) picked.value = attrOptions.value.slice(0, 4).map((o) => o.value)
    } finally {
        loading.value = false
    }
}

const render = () => {
    if (!chartRef.value) return
    if (!chart.value || chart.value.getDom() !== chartRef.value) {
        chart.value?.dispose()
        chart.value = echarts.init(chartRef.value)
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
        return
    }
    const rows = [...trendRows.value].sort((a, b) => a.ts - b.ts)
    chart.value.setOption(
        {
            tooltip: { trigger: 'axis' },
            legend: { data: picked.value.map(nameOf) },
            grid: { left: 50, right: 20, top: 36, bottom: 60 },
            dataZoom: [{ type: 'inside' }],
            xAxis: { type: 'time' },
            yAxis: { type: 'value', scale: true },
            series: picked.value.map((k) => ({
                name: nameOf(k),
                type: 'line',
                showSymbol: false,
                data: rows
                    .filter((r) => r.values && k in r.values)
                    .map((r) => [r.ts, Number(r.values![k])])
            }))
        },
        true
    )
}

// 勾选属性或区间数据变化即重绘（同一 tick 内多处改动只触发一次），切换属性无需重新请求
watch([picked, trendRows], render)

const resize = () => chart.value?.resize()
onMounted(() => {
    load()
    window.addEventListener('resize', resize)
})
onBeforeUnmount(() => {
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
