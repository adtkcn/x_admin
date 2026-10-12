<template>
    <div v-loading="loading">
        <el-descriptions :column="3" border size="small">
            <el-descriptions-item label="在线">
                <el-tag :type="state.online ? 'success' : 'info'" size="small">
                    {{ state.online ? '在线' : '离线' }}
                </el-tag>
            </el-descriptions-item>
            <el-descriptions-item label="最后报文">{{ state.updated_at || '-' }}</el-descriptions-item>
            <el-descriptions-item label="最后上线">{{ fmt(state.last_online_time) }}</el-descriptions-item>
            <el-descriptions-item label="最后离线">{{ fmt(state.last_offline_time) }}</el-descriptions-item>
            <el-descriptions-item label="判定窗口">{{ state.offline_timeout ?? '-' }}s</el-descriptions-item>
        </el-descriptions>
        <div class="mt-4 text-sm font-medium">最新属性快照</div>
        <el-table :data="rows" size="small" class="mt-2" max-height="320">
            <el-table-column prop="label" label="属性" width="220" />
            <el-table-column prop="value" label="值" />
        </el-table>
    </div>
</template>

<script lang="ts" setup>
import { ref, computed, onMounted, onBeforeUnmount } from 'vue'
import { deviceState, type type_device_state } from '@/api/iot/device'
import type { ThingProperty } from '@/views/iot/product/components/thing-model'
import { displayValue } from './thing-display'

const props = defineProps<{ deviceId: string; thingProps: ThingProperty[] }>()

const state = ref<type_device_state>({})
const loading = ref(false)
let timer: ReturnType<typeof setInterval> | null = null

// 有物模型则按其属性逐条展示（属性名 / 单位来自物模型），未配物模型时回退展示快照全部键值
const rows = computed(() => {
    const vals = state.value.values || {}
    if (props.thingProps.length) {
        return props.thingProps.map((p) => ({
            label: p.name === p.identifier ? p.identifier : `${p.name}（${p.identifier}）`,
            value: displayValue(vals[p.identifier], p)
        }))
    }
    return Object.entries(vals).map(([k, v]) => ({ label: k, value: displayValue(v) }))
})

const fmt = (v?: any) => (typeof v === 'string' && v ? v : '-')

const load = async () => {
    if (!props.deviceId) return
    loading.value = !state.value.device_id
    try {
        state.value = await deviceState({ id: props.deviceId })
    } finally {
        loading.value = false
    }
}

onMounted(() => {
    load()
    // 5s 轮询承接实时值：规则链「状态快照」节点写 Redis，设备房间不再逐条广播
    timer = setInterval(load, 5000)
})
onBeforeUnmount(() => {
    if (timer) clearInterval(timer)
    timer = null
})
</script>
