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
                    <div class="w-1/2 md:w-1/4">
                        <div class="leading-10">访问量(人)</div>
                        <div class="text-6xl">{{ workbenchData.today.todayVisits }}</div>
                        <div class="text-tx-secondary text-xs">
                            总访问量：{{ workbenchData.today.totalVisits }}
                        </div>
                    </div>
                    <div class="w-1/2 md:w-1/4">
                        <div class="leading-10">待办审批</div>
                        <div class="text-6xl">{{ workbenchData.today.flow_todo }}</div>
                    </div>
                    <div class="w-1/2 md:w-1/4">
                        <div class="leading-10">订单量(笔)</div>
                        <div class="text-6xl">{{ workbenchData.today.todayOrder }}</div>
                        <div class="text-tx-secondary text-xs">
                            总订单量：{{ workbenchData.today.totalOrder }}
                        </div>
                    </div>
                    <div class="w-1/2 md:w-1/4">
                        <div class="leading-10">新增用户</div>
                        <div class="text-6xl">{{ workbenchData.today.todayUsers }}</div>
                        <div class="text-tx-secondary text-xs">
                            总访用户：{{ workbenchData.today.totalUsers }}
                        </div>
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
import { reactive, onActivated, onMounted, useTemplateRef } from 'vue'
import { getWorkbench } from '@/api/app'
import type { Console } from '@/api/app'
import { onWsMessage } from '@/hooks/useGlobalWs'

import type { ECOption } from '@/utils/echart'

defineOptions({
    name: 'workbench'
})

// 订阅 WS 在线人数消息（组件卸载自动取消订阅）
onWsMessage('onlineCount', (msg) => {
    updateChart(msg.data.count)
})

// 表单数据
const workbenchData: Console = reactive({
    version: {
        version: '',
        name: ''
    },

    today: {
        time: '',
        todayVisits: 0,
        totalVisits: 0,
        flow_todo: 0,
        todayOrder: 0,
        totalOrder: 0,
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
    workbenchData.version = res.version
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
    console.log('onActivated')
})
onMounted(() => {
    getData()
})
</script>

<style lang="scss" scoped></style>
