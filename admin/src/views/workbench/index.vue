<template>
    <div class="workbench">
        <div class="md:flex">
            <el-card class="border-none! mb-4 flex-1" shadow="never">
                <template #header>
                    <div>
                        <span class="card-title">数据统计</span>
                        <span class="text-tx-secondary text-xs ml-4">
                            更新时间：{{ workbenchData.time }}
                        </span>
                    </div>
                </template>

                <div class="flex flex-wrap">
                    <div class="w-1/3 md:w-1/6">
                        <div class="leading-10">轻微告警</div>
                        <div class="text-4xl" :style="{ color: SeverityColor[1] }">
                            {{ workbenchData.alertMinor }}
                        </div>
                    </div>
                    <div class="w-1/3 md:w-1/6">
                        <div class="leading-10">严重告警</div>
                        <div class="text-4xl" :style="{ color: SeverityColor[2] }">
                            {{ workbenchData.alertSerious }}
                        </div>
                    </div>
                    <div class="w-1/3 md:w-1/6">
                        <div class="leading-10">故障告警</div>
                        <div class="text-4xl" :style="{ color: SeverityColor[-1] }">
                            {{ workbenchData.alertFault }}
                        </div>
                    </div>
                    <div class="w-1/3 md:w-1/6">
                        <div class="leading-10">待办审批</div>
                        <div class="text-4xl">{{ workbenchData.flow_todo }}</div>
                    </div>
                    <div class="w-1/3 md:w-1/6">
                        <div class="leading-10">在线设备(台)</div>
                        <div class="text-4xl">{{ workbenchData.deviceOnline }}/{{ workbenchData.deviceTotal }}</div>
               
                    </div>
                    <div class="w-1/3 md:w-1/6">
                        <div class="leading-10">今日新增用户</div>
                        <div class="text-4xl">{{ workbenchData.todayUsers }}/{{ workbenchData.totalUsers }}</div>
 
                    </div>
                </div>
            </el-card>
        </div>

        <div class="md:flex">
            <el-card class="flex-1 border-none! mb-4" shadow="never">
                <template #header>
                    <span>在线人数趋势</span>
                </template>

                <div>
                    <echart-component
                        ref="visitorChartRef"
                        height="350px"
                        :option="visitorOption"
                        :autoresize="true"
                    />
                </div>
            </el-card>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { reactive, watch, onActivated, onDeactivated, onMounted, onUnmounted, useTemplateRef } from 'vue'
import { getWorkbench, getOnlineTrend } from '@/api/app'
import type { Console } from '@/api/app'
import { SeverityColor } from '@/api/iot/alert'
import { onWsMessage, useGlobalWs } from '@/hooks/useGlobalWs'

import type { ECOption } from '@/utils/echart'

defineOptions({
    name: 'workbench'
})

// 在线人数按房间推送：仅首页订阅，进页加入、离页退出，无关连接不再收此消息
const ONLINE_ROOM = 'onlineCount'
const { send, status: wsStatus } = useGlobalWs()
const joinOnlineRoom = () => send(JSON.stringify({ type: 'join_room', data: ONLINE_ROOM }))
const leaveOnlineRoom = () => send(JSON.stringify({ type: 'leave_room', data: ONLINE_ROOM }))

// 房间成员关系存于服务端连接：断线自动重连后需重新加入
watch(wsStatus, (s) => {
    if (s === 'open') joinOnlineRoom()
})

// 订阅 WS 在线人数消息（组件卸载自动取消订阅）
onWsMessage('onlineCount', (msg) => {
    updateChart(msg.data.count)
})

// 表单数据
const workbenchData: Console = reactive({
    time: '',
    flow_todo: 0,
    deviceTotal: 0,
    deviceOnline: 0,
    alertMinor: 0,
    alertSerious: 0,
    alertFault: 0,
    todayUsers: 0,
    totalUsers: 0
})
const visitorOption = {
    xAxis: {
        type: 'category',
        data: [] as string[]
    },
    yAxis: {
        type: 'value'
    },
    legend: {
        data: ['在线人数']
    },
    itemStyle: {
        color: 'red'
    },
    tooltip: {
        trigger: 'axis'
    },
    grid: {
        left: '3%',
        right: '4%',
        containLabel: true
    },
    series: [
        {
            name: '在线人数',
            data: [] as number[],
            type: 'line',
            smooth: true,
            symbol: 'none'
        }
    ]
}
const visitorChartRef = useTemplateRef('visitorChartRef')

// 拉取统计数据（定时刷新）
const getData = async () => {
    try {
        const res = await getWorkbench()
        Object.assign(workbenchData, res)
    } catch (error) {
        console.error('工作台统计数据获取失败:', error)
    }
}

// 拉取在线人数曲线初始数据（仅进页取一次，后续增量由 WS 推送）
const getTrend = async () => {
    try {
        const res = await getOnlineTrend()
        visitorOption.xAxis.data = res.date
        visitorOption.series[0].data = res.list
        visitorChartRef.value?.setOption(visitorOption as ECOption)
    } catch (error) {
        console.error('在线人数趋势获取失败:', error)
    }
}

// 统计数据定时刷新（页面激活期间开启，离页/卸载停止）
const STAT_REFRESH_MS = 5 * 1000
let statTimer: ReturnType<typeof setInterval> | null = null
const startStatTimer = () => {
    stopStatTimer()
    statTimer = setInterval(getData, STAT_REFRESH_MS)
}
const stopStatTimer = () => {
    if (statTimer) {
        clearInterval(statTimer)
        statTimer = null
    }
}

function updateChart(val: number) {
    visitorOption.xAxis.data.push(new Date().toLocaleTimeString())
    visitorOption.series[0].data.push(val)

    // 保持数据长度在20个
    if (visitorOption.xAxis.data.length > 20) {
        visitorOption.xAxis.data.shift()
        visitorOption.series[0].data.shift()
    }
    visitorChartRef.value?.setOption(visitorOption as ECOption)
}

onActivated(() => {
    joinOnlineRoom()
    startStatTimer()
})
onDeactivated(() => {
    leaveOnlineRoom()
    stopStatTimer()
})
onMounted(() => {
    getData()
    getTrend()
    joinOnlineRoom()
    startStatTimer()
})
onUnmounted(() => {
    leaveOnlineRoom()
    stopStatTimer()
})
</script>

<style lang="scss" scoped></style>
