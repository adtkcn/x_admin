<template>
    <div class="workbench">
        <div class="md:flex">
            <el-card class="border-none! mb-4 flex-1" shadow="never">
                <template #header>
                    <div>
                        <span class="card-title">数据统计</span>
                        <span class="text-tx-secondary text-xs ml-4">
                            更新时间：{{ workbenchData.today.time }}
                        </span>
                    </div>
                </template>

                <div class="flex flex-wrap">
                    <div class="w-1/3 md:w-1/6">
                        <div class="leading-10">轻微告警</div>
                        <div class="text-4xl" :style="{ color: SeverityColor[1] }">
                            {{ workbenchData.today.alertMinor }}
                        </div>
                    </div>
                    <div class="w-1/3 md:w-1/6">
                        <div class="leading-10">严重告警</div>
                        <div class="text-4xl" :style="{ color: SeverityColor[2] }">
                            {{ workbenchData.today.alertSerious }}
                        </div>
                    </div>
                    <div class="w-1/3 md:w-1/6">
                        <div class="leading-10">故障告警</div>
                        <div class="text-4xl" :style="{ color: SeverityColor[-1] }">
                            {{ workbenchData.today.alertFault }}
                        </div>
                    </div>
                    <div class="w-1/3 md:w-1/6">
                        <div class="leading-10">待办审批</div>
                        <div class="text-4xl">{{ workbenchData.today.flow_todo }}</div>
                    </div>
                    <div class="w-1/3 md:w-1/6">
                        <div class="leading-10">在线设备(台)</div>
                        <div class="text-4xl">{{ workbenchData.today.deviceOnline }}/{{ workbenchData.today.deviceTotal }}</div>
               
                    </div>
                    <div class="w-1/3 md:w-1/6">
                        <div class="leading-10">今日新增用户</div>
                        <div class="text-4xl">{{ workbenchData.today.todayUsers }}/{{ workbenchData.today.totalUsers }}</div>
 
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
import { getWorkbench } from '@/api/app'
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
    today: {
        time: '',
        flow_todo: 0,
        deviceTotal: 0,
        deviceOnline: 0,
        alertMinor: 0,
        alertSerious: 0,
        alertFault: 0,
        todayUsers: 0,
        totalUsers: 0
    },

    visitor: {
        date: [],
        list: []
    }
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

// 获取工作台主页数据
const getData = async () => {
    try {
        const res = await getWorkbench()
    workbenchData.today = res.today
    workbenchData.visitor = res.visitor

    // 写入从后台拿来的数据
    visitorOption.xAxis.data = res.visitor.date
    visitorOption.series[0].data = res.visitor.list
    visitorChartRef.value?.setOption(visitorOption as ECOption)
    } catch (error) {
        console.error('工作台数据获取失败:', error)
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
})
onDeactivated(() => {
    leaveOnlineRoom()
})
onMounted(() => {
    getData()
    joinOnlineRoom()
})
onUnmounted(() => {
    leaveOnlineRoom()
})
</script>

<style lang="scss" scoped></style>
