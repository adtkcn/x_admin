<template>
    <div class="map-pop">
        <div class="pop__title">{{ title }}</div>
        <div>SN：{{ sn || '-' }}</div>
        <div>状态：{{ status === 1 ? '在线' : '离线' }}</div>
        <div v-if="severity !== 0" class="pop__alert" :style="{ color: sevColor }">
            告警中·{{ SeverityMap[severity] || severity }}
        </div>
        <template v-if="rows.length">
            <div v-for="r in rows" :key="r.label">{{ r.label }} = {{ r.value }}</div>
        </template>
        <div v-else>暂无实时值</div>
    </div>
</template>

<script lang="ts" setup>
import { computed } from 'vue'
import { SeverityMap, SeverityColor } from '@/api/iot/alert'
import type { ThingProperty } from '@/views/iot/product/components/thing-model'
import { displayValue } from '@/views/iot/device/components/thing-display'

defineOptions({ name: 'DevicePopupCard' })

const props = defineProps<{
    id?: string
    name?: string
    sn?: string
    status: number
    severity: number
    values?: Record<string, any>
    thingProps?: ThingProperty[]
}>()

const title = computed(() => props.name || props.sn || props.id || '')
// 危害序含负数档，未收录的更高级别一律取顶格红
const sevColor = computed(() => SeverityColor[props.severity] || SeverityColor[2])

// 有物模型按属性逐条展示（名称/单位来自物模型，缺失值显示 '-'），未配物模型回退原始键值
const rows = computed(() => {
    const vals = props.values || {}
    if (props.thingProps?.length) {
        return props.thingProps.map((p) => ({
            label: p.name === p.identifier ? p.identifier : `${p.name}（${p.identifier}）`,
            value: displayValue(vals[p.identifier], p)
        }))
    }
    return Object.entries(vals).map(([k, v]) => ({ label: k, value: displayValue(v) }))
})
</script>

<style scoped>
.pop__title,
.pop__alert {
    font-weight: 600;
}
</style>
