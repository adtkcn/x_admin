<template>
    <div class="workbench">
        <div class="md:flex">
            <el-card class="!border-none mb-4 md:mr-4" shadow="never">
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
            <el-card class="!border-none mb-4 flex-1" shadow="never">
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
            <el-card class="flex-1 !border-none md:mr-4 mb-4" shadow="never">
                <template #header>
                    <span>访问量趋势图</span>
                </template>
                <div>
                    <v-charts
                        style="height: 350px"
                        :option="workbenchData.visitorOption"
                        :autoresize="true"
                    />
                </div>
            </el-card>
        </div>
    </div>
</template>

<script lang="ts" setup name="workbench">
import { getWorkbench } from '@/api/app'
import '@/utils/echart'
import vCharts from 'vue-echarts'

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
    article: [], // 文章阅读量

    visitorOption: {
        xAxis: {
            type: 'category',
            data: [0]
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
                data: [0],
                type: 'line',
                smooth: true
            }
        ]
    }
})

// 获取工作台主页数据
const getData = async () => {
    const res = await getWorkbench()
    workbenchData.version = res.version
    workbenchData.today = res.today
    workbenchData.visitor = res.visitor

    // 清空echarts 数据
    workbenchData.visitorOption.xAxis.data = []
    workbenchData.visitorOption.series[0].data = []

    // 写入从后台拿来的数据
    workbenchData.visitorOption.xAxis.data = res.visitor.date
    workbenchData.visitorOption.series[0].data = res.visitor.list
}

getData()
</script>

<style lang="scss" scoped></style>
