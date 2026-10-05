<template>
    <el-dialog
        v-model="visible"
        :title="`设备详情：${deviceName}`"
        width="920px"
        destroy-on-close
        @closed="handleClosed"
    >
        <el-tabs v-model="tab" @tab-change="handleTabChange">
            <!-- 实时状态：Redis 运行态快照，打开期间 5s 轮询 -->
            <el-tab-pane label="实时状态" name="state">
                <div v-loading="stateLoading">
                    <el-descriptions :column="3" border size="small">
                        <el-descriptions-item label="在线">
                            <el-tag :type="state.online ? 'success' : 'info'" size="small">
                                {{ state.online ? '在线' : '离线' }}
                            </el-tag>
                        </el-descriptions-item>
                        <el-descriptions-item label="距判离线">{{ state.offline_in ?? '-' }}s</el-descriptions-item>
                        <el-descriptions-item label="最后报文">{{ state.updated_at || '-' }}</el-descriptions-item>
                        <el-descriptions-item label="最后上线">{{ fmt(state.last_online_time) }}</el-descriptions-item>
                        <el-descriptions-item label="最后离线">{{ fmt(state.last_offline_time) }}</el-descriptions-item>
                        <el-descriptions-item label="判定窗口">{{ state.offline_timeout ?? '-' }}s</el-descriptions-item>
                    </el-descriptions>
                    <div class="mt-4 text-sm font-medium">最新属性快照</div>
                    <el-table :data="stateRows" size="small" class="mt-2" max-height="320">
                        <el-table-column prop="key" label="属性" width="220" />
                        <el-table-column prop="value" label="值" />
                    </el-table>
                </div>
            </el-tab-pane>

            <!-- 历史趋势：时序区间回读 + echarts 多属性曲线 -->
            <el-tab-pane label="历史趋势" name="trend">
                <div class="flex items-center gap-2 mb-3">
                    <el-date-picker
                        v-model="range"
                        type="datetimerange"
                        size="small"
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
                        placeholder="属性(先查询后选择)"
                    >
                        <el-option v-for="k in identifiers" :key="k" :label="k" :value="k" />
                    </el-select>
                    <el-button type="primary" size="small" @click="loadTrend">查询</el-button>
                </div>
                <div ref="chartRef" v-loading="trendLoading" class="trend-chart"></div>
                <div v-if="trendRows.length" class="text-xs mt-1 text-gray-400">
                    共取 {{ trendRows.length }} 条（按时间倒序上限 5000），曲线按时间正序绘制
                </div>
            </el-tab-pane>

            <!-- 告警记录：该设备的告警事件历史 -->
            <el-tab-pane label="告警记录" name="alert">
                <el-table :data="alerts" size="small" v-loading="alertLoading" max-height="380">
                    <el-table-column label="级别" width="80">
                        <template #default="{ row }">
                            <el-tag :type="SeverityTagType[row.severity] || 'info'" size="small">
                                {{ SeverityMap[row.severity] || row.severity }}
                            </el-tag>
                        </template>
                    </el-table-column>
                    <el-table-column label="状态" width="80">
                        <template #default="{ row }">
                            <el-tag :type="row.status ? 'danger' : 'success'" size="small">
                                {{ row.status ? '告警中' : '已恢复' }}
                            </el-tag>
                        </template>
                    </el-table-column>
                    <el-table-column label="触发时间" prop="triggered_at" width="170" />
                    <el-table-column label="恢复时间" width="170">
                        <template #default="{ row }">{{ row.recovered_at || '-' }}</template>
                    </el-table-column>
                    <el-table-column label="属性快照">
                        <template #default="{ row }">{{ snapText(row.values) }}</template>
                    </el-table-column>
                </el-table>
                <div class="flex justify-end mt-2">
                    <el-pagination
                        v-model:current-page="alertPage.pageNo"
                        :page-size="alertPage.pageSize"
                        :total="alertPage.count"
                        layout="prev, pager, next"
                        small
                        @current-change="loadAlerts"
                    />
                </div>
            </el-tab-pane>
        </el-tabs>
    </el-dialog>
</template>

<script lang="ts" setup>
import { ref, computed, shallowRef, nextTick, onBeforeUnmount } from 'vue'
import {
    deviceState,
    deviceData,
    alertLists,
    SeverityMap,
    SeverityTagType,
    type type_device_state,
    type type_device_data,
    type type_iot_alert
} from '@/api/iot'
import { echarts } from '@/utils/echart'

defineOptions({ name: 'DeviceDetailPopup' })

const visible = ref(false)
const tab = ref('state')
const deviceId = ref('')
const deviceName = ref('')

// ---------- 实时状态 ----------
const state = ref<type_device_state>({})
const stateLoading = ref(false)
let pollTimer: ReturnType<typeof setInterval> | null = null

const stateRows = computed(() =>
    Object.entries(state.value.values || {}).map(([key, value]) => ({ key, value }))
)

const fmt = (v?: any) => (typeof v === 'string' && v ? v : '-')
const snapText = (vals?: Record<string, any>) =>
    vals ? Object.entries(vals).map(([k, v]) => `${k}=${v}`).join('; ') : '-'

const loadState = async () => {
    if (!deviceId.value) return
    stateLoading.value = !state.value.device_id
    try {
        state.value = await deviceState({ id: deviceId.value })
    } finally {
        stateLoading.value = false
    }
}

const startPoll = () => {
    stopPoll()
    // 5s 轮询承接实时值：规则链「状态快照」节点写 Redis，设备房间不再逐条广播
    pollTimer = setInterval(loadState, 5000)
}
const stopPoll = () => {
    if (pollTimer) clearInterval(pollTimer)
    pollTimer = null
}

// ---------- 历史趋势 ----------
const range = ref<[Date, Date]>([new Date(Date.now() - 3600e3), new Date()])
const trendRows = ref<type_device_data[]>([])
const identifiers = ref<string[]>([])
const picked = ref<string[]>([])
const trendLoading = ref(false)
const chartRef = ref<HTMLElement>()
const chart = shallowRef<ReturnType<typeof echarts.init>>()

const loadTrend = async () => {
    if (!deviceId.value || !range.value) return
    trendLoading.value = true
    try {
        trendRows.value = await deviceData({
            id: deviceId.value,
            start_ms: new Date(range.value[0]).getTime(),
            end_ms: new Date(range.value[1]).getTime(),
            limit: 2000
        })
        // 汇总区间内出现过的数值型属性供勾选
        const seen: Record<string, true> = {}
        for (const row of trendRows.value) {
            for (const [k, v] of Object.entries(row.values || {})) {
                if (typeof v === 'number' || (v !== null && Number.isFinite(Number(v)))) seen[k] = true
            }
        }
        identifiers.value = Object.keys(seen)
        if (!picked.value.length) picked.value = identifiers.value.slice(0, 4)
        renderChart()
    } finally {
        trendLoading.value = false
    }
}

const renderChart = () => {
    if (!chartRef.value || !trendRows.value.length) return
    if (!chart.value || chart.value.getDom() !== chartRef.value) {
        chart.value?.dispose()
        chart.value = echarts.init(chartRef.value)
    }
    const rows = [...trendRows.value].sort((a, b) => a.ts - b.ts)
    chart.value.setOption(
        {
            tooltip: { trigger: 'axis' },
            legend: { data: picked.value },
            grid: { left: 50, right: 20, top: 36, bottom: 30 },
            dataZoom: [{ type: 'inside' }],
            xAxis: { type: 'time' },
            yAxis: { type: 'value', scale: true },
            series: picked.value.map((k) => ({
                name: k,
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

// ---------- 告警记录 ----------
const alerts = ref<type_iot_alert[]>([])
const alertLoading = ref(false)
const alertPage = ref({ pageNo: 1, pageSize: 10, count: 0 })

const loadAlerts = async () => {
    if (!deviceId.value) return
    alertLoading.value = true
    try {
        const res = await alertLists({
            device_id: deviceId.value,
            pageNo: alertPage.value.pageNo,
            pageSize: alertPage.value.pageSize
        })
        alerts.value = res.lists
        alertPage.value.count = Number(res.count)
    } finally {
        alertLoading.value = false
    }
}

// ---------- 生命周期 ----------
const handleTabChange = async (name: string | number) => {
    if (name === 'trend') {
        await nextTick()
        chart.value?.resize()
        if (!trendRows.value.length) loadTrend()
    } else if (name === 'alert') {
        if (!alerts.value.length) loadAlerts()
    }
}

const open = async (row: { id?: string; name?: string; sn?: string }) => {
    deviceId.value = row.id || ''
    deviceName.value = row.name || row.sn || row.id || ''
    tab.value = 'state'
    state.value = {}
    trendRows.value = []
    identifiers.value = []
    picked.value = []
    alerts.value = []
    alertPage.value.pageNo = 1
    visible.value = true
    await loadState()
    startPoll()
}

const handleClosed = () => {
    stopPoll()
    chart.value?.dispose()
    chart.value = undefined
}

onBeforeUnmount(() => {
    stopPoll()
    chart.value?.dispose()
})

defineExpose({ open })
</script>

<style scoped>
.trend-chart {
    height: 360px;
    width: 100%;
}
</style>
