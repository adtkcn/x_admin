<template>
    <el-dialog v-model="visible" :title="`告警趋势：${deviceName}`" width="920px" destroy-on-close>
        <device-trend :device-id="deviceId" :thing-props="thingProps" :initial-range="range" :highlight-range="highlightRange" />
    </el-dialog>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { productDetail, type type_iot_alert } from '@/api/iot'
import { parseThingModel, type ThingProperty } from '@/views/iot/product/components/thing-model'
import DeviceTrend from '@/views/iot/device/components/device-trend.vue'

defineOptions({ name: 'AlertTrendPopup' })

// 告警时间段前后各延伸的秒数
const PAD_MS = 30 * 1000

const visible = ref(false)
const deviceId = ref('')
const deviceName = ref('')
const thingProps = ref<ThingProperty[]>([])
const range = ref<[Date, Date]>()
const highlightRange = ref<[number, number]>()

// "YYYY-MM-DD HH:mm:ss" -> Date（替换分隔符以兼容本地时区解析）
const toDate = (s?: string | null) => (s ? new Date(s.replace(/-/g, '/')) : null)

const open = async (row: type_iot_alert) => {
    deviceId.value = row.device_id || ''
    deviceName.value = row.sn || row.device_id || ''
    thingProps.value = []
    // 触发时间前 30s ~（已恢复取恢复时间，否则取当前时间）后 30s
    const start = toDate(row.triggered_at)
    const end = toDate(row.recovered_at) ?? new Date()
    range.value = start ? [new Date(start.getTime() - PAD_MS), new Date(end.getTime() + PAD_MS)] : undefined
    // 实际告警时间段（不延伸），用于图表 markArea 高亮
    highlightRange.value = start ? [start.getTime(), end.getTime()] : undefined
    visible.value = true
    // destroy-on-close：每次打开 trend 组件重新挂载，range 作为初始区间生效
    if (row.product_id) {
        const product = await productDetail({ id: row.product_id })
        thingProps.value = parseThingModel(product.thing_model || '').properties
    }
}

defineExpose({ open })
</script>
