<!-- 系统环境 -->
<template>
    <div class="system-environment" v-loading="loading">
        <el-tabs v-model="activeIp" type="card">
            <el-tab-pane v-for="ip in ips" :key="ip" :label="ip" :name="ip"></el-tab-pane>
        </el-tabs>

        <div class="lg:flex">
            <el-card class="border-none! flex-1 mb-4 lg:mr-4" shadow="never">
                <div>CPU</div>
                <div class="mt-4">
                    <div class="flex flex-wrap">
                        <div class="sm:flex-1 w-1/2 mb-4">
                            <div class="text-4xl mb-3">{{ info.cpu.cpu_num }}</div>
                            <div class="text-tx-regular">核心数</div>
                        </div>

                        <div class="sm:flex-1 w-1/2 mb-4">
                            <div class="text-4xl mb-3">
                                {{ info.cpu.used ? `${info.cpu.used}%` : '-' }}
                            </div>
                            <div class="text-tx-regular">用户使用率</div>
                        </div>

                        <div class="sm:flex-1 w-1/2 mb-4">
                            <div class="text-4xl mb-3">
                                {{ info.cpu.sys ? `${info.cpu.sys}%` : '-' }}
                            </div>
                            <div class="text-tx-regular">系统使用率</div>
                        </div>

                        <div class="sm:flex-1 w-1/2 mb-4">
                            <div class="text-4xl mb-3">
                                {{ info.cpu.free ? `${info.cpu.free}%` : '-' }}
                            </div>
                            <div class="text-tx-regular">当前空闲率</div>
                        </div>
                    </div>
                </div>
            </el-card>
            <el-card class="border-none! flex-1 mb-4" shadow="never">
                <div>内存</div>
                <div class="mt-4">
                    <div class="flex flex-wrap">
                        <div class="sm:flex-1 w-1/2 mb-4">
                            <div class="text-4xl mb-3">{{ info.mem.total }}</div>
                            <div class="text-tx-regular">总内存</div>
                        </div>

                        <div class="sm:flex-1 w-1/2 mb-4">
                            <div class="text-4xl mb-3">
                                {{ info.mem.used ? `${info.mem.used}` : '-' }}
                            </div>
                            <div class="text-tx-regular">已用内存</div>
                        </div>

                        <div class="sm:flex-1 w-1/2 mb-4">
                            <div class="text-4xl mb-3">
                                {{ info.mem.free ? `${info.mem.free}` : '-' }}
                            </div>
                            <div class="text-tx-regular">剩余内存</div>
                        </div>

                        <div class="sm:flex-1 w-1/2 mb-4">
                            <div class="text-4xl mb-3">
                                {{ info.mem.usage ? `${info.mem.usage}%` : '-' }}
                            </div>
                            <div class="text-tx-regular">使用率</div>
                        </div>
                    </div>
                </div>
            </el-card>
        </div>
        <el-card class="border-none!" shadow="never">
            <div>服务器信息</div>
            <div class="mt-4">
                <vxe-table :data="[info.sys]" :row-config="{ keyField: 'computerName' }" border>
                    <vxe-column field="computerName" title="服务器名称" min-width="150" />
                    <vxe-column field="computerIp" title="服务器IP" min-width="120" />
                    <vxe-column field="osName" title="操作系统" min-width="100" />
                    <vxe-column field="osArch" title="系统架构" min-width="100" />
                    <vxe-column field="userDir" title="项目路径" min-width="250" />
                </vxe-table>
            </div>
        </el-card>

        <el-card shadow="never" class="border-none! mt-4">
            <div>go环境信息</div>
            <div class="mt-4">
                <vxe-table :data="[info.go]" :row-config="{ keyField: 'name' }" border>
                    <vxe-column field="name" title="go名称" min-width="120" />
                    <vxe-column field="startTime" title="启动时间" min-width="120" />
                    <vxe-column field="home" title="安装路径" min-width="120" />
                    <vxe-column field="inputArgs" title="运行参数" min-width="120" />
                    <vxe-column field="version" title="go版本" min-width="120" />
                    <vxe-column field="runTime" title="运行时长" min-width="120" />
                </vxe-table>
            </div>
        </el-card>

        <el-card shadow="never" class="border-none! mt-4">
            <div>硬盘状态</div>
            <div class="mt-4">
                <vxe-table :data="info.disk" :row-config="{ keyField: 'dirName' }" border>
                    <vxe-column field="dirName" title="盘符路径" min-width="100" />
                    <vxe-column field="sysTypeName" title="文件系统" min-width="100" />
                    <vxe-column field="typeName" title="盘符类型" min-width="100" />
                    <vxe-column field="total" title="总大小" min-width="100" />
                    <vxe-column field="free" title="可用大小" min-width="100" />
                    <vxe-column field="used" title="已用大小" min-width="100" />
                    <vxe-column field="usage" title="已用百分比" min-width="100">
                        <template #default="{ row }"> {{ row.usage }}% </template>
                    </vxe-column>
                </vxe-table>
            </div>
        </el-card>
    </div>
</template>

<script lang="ts" setup>
import { ref, onBeforeUnmount, computed } from 'vue'
import { systemInfo } from '@/api/setting/system'
defineOptions({
    name: 'environment'
})
const loading = ref(true)
const result = ref<Record<string, any>>({})
const ips = computed(() => {
    return Object.keys(result.value)
})
const activeIp = ref<string>('')
const info = computed(() => {
    if (result.value && activeIp.value && result.value[activeIp.value]) {
        return result.value[activeIp.value]
    }
    return {
        cpu: {} as any,
        disk: [],
        go: {},
        mem: {} as any,
        sys: {}
    }
})

const getSystemInfo = async () => {
    try {
        // loading.value = true
        const data = await systemInfo()
        result.value = data || {}
        if (activeIp.value == '' && ips.value.length > 0) {
            activeIp.value = ips.value[0]
        }
        loading.value = false
    } catch (error) {
        console.error('系统环境信息获取失败:', error)
        loading.value = false
    }
}

getSystemInfo()
const timer = setInterval(() => {
    getSystemInfo()
}, 2000)
onBeforeUnmount(() => {
    clearInterval(timer)
})
</script>

<style lang="scss" scoped></style>
