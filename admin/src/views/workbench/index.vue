<template>
    <div class="workbench">
        <div class="md:flex">
            <el-card class="border-none! mb-4 md:mr-4" shadow="never">
                <template #header>
                    <span class="card-title">版本信息</span>
                </template>
                <div>
                    <div class="flex leading-9">
                        <div class="w-20 flex-none">当前版本</div>
                        <span> {{ workbenchData.version.version }}</span>
                    </div>
                    <div class="flex leading-9">
                        <div class="w-20 flex-none">基于框架</div>
                        <span> {{ workbenchData.version.based }}</span>
                    </div>
                    <div class="flex leading-9">
                        <div class="w-20 felx-none">获取渠道</div>
                        <div>
                            <a :href="workbenchData.version.channel.website" target="_blank">
                                <el-button type="success" size="small">官网</el-button>
                            </a>
                            <a
                                class="ml-3"
                                href="https://github.com/adtkcn/x_admin.git"
                                target="_blank"
                            >
                                <el-button type="danger" size="small">Github</el-button>
                            </a>
                        </div>
                    </div>
                </div>
            </el-card>
            <el-card class="border-none! mb-4 flex-1" shadow="never">
                <template #header>
                    <div>
                        <span class="card-title">今日数据</span>
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
                        <div class="leading-10">销售额(元)</div>
                        <div class="text-6xl">{{ workbenchData.today.todaySales }}</div>
                        <div class="text-tx-secondary text-xs">
                            总销售额：{{ workbenchData.today.totalSales }}
                        </div>
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
                    <span>访问量趋势图</span>
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
import { reactive, onDeactivated, onActivated, onMounted, useTemplateRef, onUnmounted } from 'vue'
import { getWorkbench } from '@/api/app'

// import feedback from '@/utils/feedback'
import { useWebSocket } from '@vueuse/core'
import useUserStore from '@/stores/modules/user'

import type { ECOption } from '@/utils/echart'

const userStore = useUserStore()
defineOptions({
    name: 'workbench'
})

// 表单数据
const workbenchData: any = reactive({
    version: {
        version: '', // 版本号
        website: '', // 官网
        based: '',
        channel: {
            gitee: '',
            website: ''
        }
    },

    today: {}, // 今日数据

    visitor: [], // 访问量
    article: [] // 文章阅读量
})
const visitorOption = {
    xAxis: {
        type: 'category',
        data: []
    },
    yAxis: {
        type: 'value'
    },
    legend: {
        data: ['访问量']
    },
    itemStyle: {
        // 点的颜色。
        color: 'red'
    },
    tooltip: {
        trigger: 'axis'
    },
    series: [
        {
            name: '访问量',
            data: [],
            type: 'line',
            smooth: true
        }
    ]
}
const visitorChartRef = useTemplateRef('visitorChartRef')
// 获取工作台主页数据
const getData = async () => {
    const res = await getWorkbench()
    workbenchData.version = res.version
    workbenchData.today = res.today
    workbenchData.visitor = res.visitor

    // 写入从后台拿来的数据
    visitorOption.xAxis.data = res.visitor.date
    visitorOption.series[0].data = res.visitor.list
    visitorChartRef.value?.setOption(visitorOption as ECOption)
}

function updateChart(val) {
    visitorOption.xAxis.data.push(new Date().toLocaleTimeString())
    visitorOption.series[0].data.push(val)

    // 保持数据长度在10个
    if (visitorOption.xAxis.data.length > 20) {
        visitorOption.xAxis.data.shift()
        visitorOption.series[0].data.shift()
    }
    visitorChartRef.value?.setOption(visitorOption as ECOption)
}

// 定义你的消息类型
interface ChatMessage {
    onlineCount: number
}
const ws = useWebSocket(`ws://localhost:8080/api/ws?token=${userStore.token}&room=room1`, {
    heartbeat: {
        message: 'ping',
        interval: 10000,
        pongTimeout: 1000
    },
    autoReconnect: true,

    onMessage(ws, e) {
        if (e.data === 'pong') {
            console.log('Received pong message')
            return
        }
        try {
            const data = JSON.parse(e.data) as ChatMessage
            updateChart(data.onlineCount)
        } catch (error) {
            console.error('JSON parse error:', error)
            return
        }
    },
    onError: (ws, event) => {
        console.error('WebSocket error:', event)
    },

    onDisconnected: (ws, event) => {
        console.log('WebSocket closed:', event)
    }
})
// setInterval(() => {
//     ws.send('ping')
// }, 1000)
onActivated(() => {
    console.log('onActivated')
})
onDeactivated(() => {
    // ws.close()
})
onMounted(() => {
    console.log('onMounted')
    // ws.connect()
    getData()
    // updateChart()
})
onUnmounted(() => {
    console.log('onUnmounted')
})
</script>

<style lang="scss" scoped></style>
