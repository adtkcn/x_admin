<!-- 系统缓存 -->
<template>
    <div class="cache">
        <el-card class="border-none!" shadow="never">
            <div>
                <div class="mb-4 lg">基本信息</div>
                <el-row :gutter="20">
                    <el-col :xl="4" :lg="6" :md="8" :sm="12" class="info-item">
                        <div class="info-label">Redis版本</div>
                        <div class="info-value">{{ baseInfo.redis_version }}</div>
                    </el-col>
                    <el-col :xl="4" :lg="6" :md="8" :sm="12" class="info-item">
                        <div class="info-label">运行模式</div>
                        <div class="info-value">
                            {{ baseInfo.redis_mode == 'standalone' ? '单机' : '集群' }}
                        </div>
                    </el-col>
                    <el-col :xl="4" :lg="6" :md="8" :sm="12" class="info-item">
                        <div class="info-label">端口</div>
                        <div class="info-value">{{ baseInfo.tcp_port }}</div>
                    </el-col>
                    <el-col :xl="4" :lg="6" :md="8" :sm="12" class="info-item">
                        <div class="info-label">客户端数</div>
                        <div class="info-value">{{ baseInfo.connected_clients }}</div>
                    </el-col>
                    <el-col :xl="4" :lg="6" :md="8" :sm="12" class="info-item">
                        <div class="info-label">运行时间(天)</div>
                        <div class="info-value">{{ baseInfo.uptime_in_days }}</div>
                    </el-col>
                    <el-col :xl="4" :lg="6" :md="8" :sm="12" class="info-item">
                        <div class="info-label">使用内存</div>
                        <div class="info-value">{{ baseInfo.used_memory_human }}</div>
                    </el-col>
                    <el-col :xl="4" :lg="6" :md="8" :sm="12" class="info-item">
                        <div class="info-label">使用CPU</div>
                        <div class="info-value">{{ baseInfo.used_cpu_user_children }}</div>
                    </el-col>
                    <el-col :xl="4" :lg="6" :md="8" :sm="12" class="info-item">
                        <div class="info-label">内存配置</div>
                        <div class="info-value">{{ baseInfo.maxmemory_human }}</div>
                    </el-col>
                    <el-col :xl="4" :lg="6" :md="8" :sm="12" class="info-item">
                        <div class="info-label">AOF是否开启</div>
                        <div class="info-value">
                            {{ baseInfo.aof_enabled == 0 ? '开启' : '关闭' }}
                        </div>
                    </el-col>
                    <el-col :xl="4" :lg="6" :md="8" :sm="12" class="info-item">
                        <div class="info-label">RDB是否成功</div>
                        <div class="info-value">
                            {{ baseInfo.aof_enabled == 'ok' ? '成功' : '失败' }}
                        </div>
                    </el-col>
                    <el-col :xl="4" :lg="6" :md="8" :sm="12" class="info-item">
                        <div class="info-label">Key数量</div>
                        <div class="info-value">{{ dbSize }}</div>
                    </el-col>
                    <el-col :xl="4" :lg="6" :md="8" :sm="12" class="info-item">
                        <div class="info-label">网络入/出口</div>
                        <div class="info-value">
                            {{ baseInfo.instantaneous_input_kbps }} /
                            {{ baseInfo.instantaneous_output_kbps }}
                        </div>
                    </el-col>
                </el-row>
            </div>
        </el-card>

        <div class="sm:flex">
            <!-- 命令统计 -->
            <el-card class="sm:mr-4 flex-1 border-none! mt-4" shadow="never">
                <div>
                    <div class="mb-10">命令统计</div>
                    <div class="flex h-[300px] items-center">
                        <echart-component
                            :option="chartOptions.commandChartOption"
                            :autoresize="true"
                        />
                    </div>
                </div>
            </el-card>

            <!-- 内存信息 -->
            <el-card class="flex-1 border-none! mt-4" shadow="never">
                <div>
                    <div class="mb-10">内存信息</div>
                    <div class="flex h-[300px] items-center">
                        <!-- {{ chartOptions.memoryChartOption }} -->
                        <echart-component
                            :option="chartOptions.memoryChartOption"
                            :autoresize="true"
                        />
                    </div>
                </div>
            </el-card>
        </div>
    </div>
</template>

<script setup lang="ts">
import { systemCache } from '@/api/setting/system'
import type { type_info, type_commandStats, type_dbSize } from '@/api/setting/system'

import { reactive, ref } from 'vue'
// import { ElTable } from 'element-plus'
defineOptions({
    name: 'cache'
})
const baseInfo = ref<type_info>({})
const dbSize = ref<type_dbSize>(0)

const chartOptions = reactive({
    commandChartOption: {
        tooltip: {
            trigger: 'item'
            // formatter: '{b} : {d}%'
        },

        series: [
            {
                label: {
                    show: true
                },
                labelLine: {
                    show: true
                },
                type: 'pie',
                radius: '85%',
                color: [
                    '#0D47A1',
                    '#1565C0',
                    '#1976D2',
                    '#1E88E5',
                    '#2196F3',
                    '#42A5F5',
                    '#64B5F6',
                    '#90CAF9',
                    '#BBDEFB',
                    '#E3F2FD',
                    '#CAF0F8',
                    '#ADE8F4',
                    '#90E0EF',
                    '#48CAE4',
                    '#00B4D8',
                    '#0096C7',
                    '#0077B6',
                    '#023E8A',
                    '#03045E',
                    '#8ecae6',
                    '#98c1d9',
                    '#D9ED92',
                    '#B5E48C',
                    '#99D98C',
                    '#76C893',
                    '#52B69A',
                    '#34A0A4',
                    '#168AAD',
                    '#1A759F',
                    '#1E6091',
                    '#184E77',
                    '#457b9d'
                ],
                data: [
                    {
                        value: '',
                        name: ''
                    }
                ],
                emphasis: {
                    itemStyle: {
                        shadowBlur: 10,
                        shadowOffsetX: 0,
                        shadowColor: 'rgba(0, 0, 0, 0.5)'
                    }
                }
            }
        ]
    },

    memoryChartOption: {
        tooltip: {
            formatter: '{a} <br/>{b} : {c}%'
        },
        series: [
            {
                name: 'Pressure',
                type: 'gauge',
                radius: '100%',
                detail: {
                    formatter: '{value}'
                },
                data: [
                    {
                        value: '',
                        name: '内存消耗'
                    }
                ]
            }
        ]
    }
})

const getSystemCache = async () => {
    try {
        const data = await systemCache()
        baseInfo.value = data.info
        dbSize.value = data.dbSize

        chartOptions.commandChartOption.series[0].data = data.commandStats

        chartOptions.memoryChartOption.series[0].data[0].value = (
            Number(data.info.used_memory) /
            1024 /
            1024
        ).toFixed(2)
        chartOptions.memoryChartOption.series[0].detail.formatter = '{value}' + 'M'
    } catch (error) {
        console.error('系统缓存信息获取失败:', error)
    }
}

getSystemCache()
</script>

<style scoped>
.el-table .el-table__cell {
    min-width: 120px;
}
.info-item {
    display: flex;
    padding: 12px 0;
    border-bottom: 1px solid #ebeef5;
}

.info-label {
    font-weight: 500;
    margin-bottom: 4px;
    color: #303133;
    width: 110px;
}

.info-value {
    color: #949597;
}
</style>
