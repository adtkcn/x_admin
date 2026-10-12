<template>
    <div class="map-pop">
        <div class="pop__title">{{ title }}</div>
        <div>SN：{{ sn || '-' }}</div>
        <div>状态：{{ status === 1 ? '在线' : '离线' }}</div>
        <div v-if="severity !== 0" class="pop__alert" :style="{ color: sevColor }">
            告警中·{{ SeverityMap[severity] || severity }}
        </div>
        <template v-if="hasValues">
            <div v-for="(v, k) in values" :key="k">{{ k }} = {{ String(v) }}</div>
        </template>
        <div v-else>暂无实时值</div>
    </div>
</template>

<script lang="ts" setup>
import { computed } from 'vue'
import { SeverityMap, SeverityColor } from '@/api/iot'

defineOptions({ name: 'DevicePopupCard' })

const props = defineProps<{
    id?: string
    name?: string
    sn?: string
    status: number
    severity: number
    values?: Record<string, any>
}>()

const title = computed(() => props.name || props.sn || props.id || '')
// 危害序含负数档，未收录的更高级别一律取顶格红
const sevColor = computed(() => SeverityColor[props.severity] || SeverityColor[2])
const hasValues = computed(() => !!props.values && Object.keys(props.values).length > 0)
</script>

<style scoped>
.pop__title,
.pop__alert {
    font-weight: 600;
}
</style>
