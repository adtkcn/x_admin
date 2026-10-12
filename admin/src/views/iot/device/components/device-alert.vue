<template>
    <div>
        <el-table :data="alerts" size="small" v-loading="loading" max-height="380">
            <el-table-column label="级别" width="80">
                <template #default="{ row }">
                    <el-tag :type="SeverityTagType[row.severity] || 'danger'" size="small">
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
                <template #default="{ row }">{{ snapText(row.values, thingMap) }}</template>
            </el-table-column>
        </el-table>
        <div class="flex justify-end mt-2">
            <el-pagination
                v-model:current-page="page.pageNo"
                :page-size="page.pageSize"
                :total="page.count"
                layout="prev, pager, next"
                size="small"
                @current-change="load"
            />
        </div>
    </div>
</template>

<script lang="ts" setup>
import { ref, computed, onMounted } from 'vue'
import { alertLists, SeverityMap, SeverityTagType, type type_iot_alert } from '@/api/iot/alert'
import type { ThingProperty } from '@/views/iot/product/components/thing-model'
import { buildThingMap, snapText } from './thing-display'

const props = defineProps<{ deviceId: string; thingProps: ThingProperty[] }>()

const alerts = ref<type_iot_alert[]>([])
const loading = ref(false)
const page = ref({ pageNo: 1, pageSize: 10, count: 0 })
// 属性快照里的 identifier 依物模型映射成中文名展示
const thingMap = computed(() => buildThingMap(props.thingProps))

const load = async () => {
    if (!props.deviceId) return
    loading.value = true
    try {
        const res = await alertLists({
            device_id: props.deviceId,
            pageNo: page.value.pageNo,
            pageSize: page.value.pageSize
        })
        alerts.value = res.lists
        page.value.count = Number(res.count)
    } finally {
        loading.value = false
    }
}

onMounted(load)
</script>
