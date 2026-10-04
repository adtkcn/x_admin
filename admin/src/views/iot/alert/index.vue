<template>
    <div class="alert-lists">
        <el-card class="border-none!" shadow="never">
            <el-form ref="formRef" class="mb-[-16px]" :model="queryParams" :inline="true">
                <el-form-item class="w-[220px]" label="设备编号">
                    <el-input v-model="queryParams.sn" clearable @keyup.enter="resetPage" />
                </el-form-item>
                <el-form-item class="w-[220px]" label="所属产品">
                    <el-select v-model="queryParams.product_id" clearable placeholder="请选择产品">
                        <el-option v-for="p in products" :key="p.id" :label="p.name" :value="(p.id as string)" />
                    </el-select>
                </el-form-item>
                <el-form-item class="w-[220px]" label="所属区域">
                    <el-tree-select
                        v-model="queryParams.space_id"
                        :data="spaceTreeOptions"
                        clearable
                        node-key="id"
                        :props="{ label: 'name' }"
                        check-strictly
                        placeholder="请选择区域"
                    />
                </el-form-item>
                <el-form-item class="w-[160px]" label="告警状态">
                    <el-select v-model="queryParams.status" clearable placeholder="全部">
                        <el-option label="告警中" :value="1" />
                        <el-option label="已恢复" :value="0" />
                    </el-select>
                </el-form-item>
                <el-form-item class="w-[160px]" label="告警级别">
                    <el-select v-model="queryParams.severity" clearable placeholder="全部">
                        <el-option v-for="(label, val) in SeverityMap" :key="val" :label="label" :value="Number(val)" />
                    </el-select>
                </el-form-item>
                <el-form-item label="触发时间">
                    <el-date-picker
                        v-model="timeRange"
                        type="datetimerange"
                        value-format="YYYY-MM-DD HH:mm:ss"
                        range-separator="至"
                        start-placeholder="开始时间"
                        end-placeholder="结束时间"
                        clearable
                        @change="handleTimeChange"
                    />
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="resetPage">查询</el-button>
                    <el-button @click="handleReset">重置</el-button>
                </el-form-item>
            </el-form>
        </el-card>
        <el-card class="border-none! mt-4" shadow="never">
            <vxe-table
                v-loading="pager.loading"
                :data="pager.lists"
                :row-config="{ keyField: 'id' }"
                :scroll-y="{ enabled: false }"
                :border="'inner'"
                show-overflow="title"
            >
                <vxe-column title="设备编号" field="sn" min-width="140" />
                <vxe-column title="所属产品" field="product_name" min-width="120" />
                <vxe-column title="所属区域" field="space_name" min-width="120">
                    <template #default="{ row }">
                        {{ row.space_name || '-' }}
                    </template>
                </vxe-column>
                <vxe-column title="状态" field="status" min-width="90">
                    <template #default="{ row }">
                        <el-tag :type="row.status ? 'danger' : 'info'">
                            {{ row.status ? '告警中' : '已恢复' }}
                        </el-tag>
                    </template>
                </vxe-column>
                <vxe-column title="级别" field="severity" min-width="90">
                    <template #default="{ row }">
                        <el-tag :type="SeverityTagType[row.severity] || 'info'">
                            {{ SeverityMap[row.severity] || '-' }}
                        </el-tag>
                    </template>
                </vxe-column>
                <vxe-column title="触发时间" field="triggered_at" min-width="180" />
                <vxe-column title="恢复时间" field="recovered_at" min-width="180">
                    <template #default="{ row }">
                        {{ row.recovered_at || '-' }}
                    </template>
                </vxe-column>
                <vxe-column title="属性快照" field="values" min-width="200">
                    <template #default="{ row }">
                        {{ formatValues(row.values) }}
                    </template>
                </vxe-column>
                <vxe-column title="操作" width="100" fixed="right">
                    <template #default="{ row }">
                        <el-button
                            v-perms="['admin:iot:alert:del']"
                            type="danger"
                            link
                            @click="handleDelete(row.id)"
                        >
                            删除
                        </el-button>
                    </template>
                </vxe-column>
            </vxe-table>
            <div class="flex justify-end mt-4">
                <pagination v-model="pager" @change="getLists" />
            </div>
        </el-card>
    </div>
</template>
<script lang="ts" setup>
import { ref, reactive, onMounted } from 'vue'
import {
    alertLists,
    alertDel,
    productAll,
    deviceSpaceAll,
    SeverityMap,
    SeverityTagType,
    type type_iot_alert_list,
    type type_iot_product
} from '@/api/iot'
import { usePaging } from '@/hooks/usePaging'
import feedback from '@/utils/feedback'
import { arrayToTree } from '@/utils/util'

defineOptions({
    name: 'iotAlert'
})

const queryParams = reactive<type_iot_alert_list>({
    sn: '',
    product_id: '',
    space_id: '',
    status: -1,
    severity: -1,
    start_time: '',
    end_time: ''
})

const { pager, getLists, resetPage, resetParams } = usePaging({
    fetchFun: alertLists,
    params: queryParams
})

// 触发时间区间：日期选择器绑定数组，变化时拆进查询参数
const timeRange = ref<string[]>([])
const handleTimeChange = (val: string[] | null) => {
    queryParams.start_time = val?.[0] ?? ''
    queryParams.end_time = val?.[1] ?? ''
    resetPage()
}

const handleReset = () => {
    timeRange.value = []
    resetParams()
}

const formatValues = (values?: Record<string, any>) => {
    if (!values) return '-'
    return Object.entries(values)
        .map(([k, v]) => `${k}=${v}`)
        .join(', ')
}

const products = ref<type_iot_product[]>([])
const spaceTreeOptions = ref<any[]>([])
onMounted(async () => {
    products.value = await productAll()
    const spaces = await deviceSpaceAll()
    spaceTreeOptions.value = arrayToTree<any>(spaces as any[], '')
})

const handleDelete = async (id: string) => {
    try {
        await feedback.confirm('确定要删除该告警记录？')
        await alertDel({ ids: [id] })
        feedback.msgSuccess('删除成功')
        getLists()
    } catch (error) {
        console.error('告警记录删除失败:', error)
    }
}

getLists()
</script>
