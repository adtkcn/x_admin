<template>
    <el-dialog v-model="visible" :title="`设备详情：${deviceName}`" width="920px" destroy-on-close>
        <el-tabs v-model="tab">
            <el-tab-pane label="实时状态" name="state">
                <device-state :device-id="deviceId" :thing-props="thingProps" />
            </el-tab-pane>
            <el-tab-pane label="历史趋势" name="trend" lazy>
                <device-trend :device-id="deviceId" :thing-props="thingProps" />
            </el-tab-pane>
            <el-tab-pane label="告警记录" name="alert" lazy>
                <device-alert :device-id="deviceId" :thing-props="thingProps" />
            </el-tab-pane>
        </el-tabs>
    </el-dialog>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { productDetail, type type_iot_device } from '@/api/iot'
import { parseThingModel, type ThingProperty } from '@/views/iot/product/components/thing-model'
import DeviceState from './components/device-state.vue'
import DeviceTrend from './components/device-trend.vue'
import DeviceAlert from './components/device-alert.vue'

defineOptions({ name: 'DeviceDetailPopup' })

const visible = ref(false)
const tab = ref('state')
const deviceId = ref('')
const deviceName = ref('')
// 所属产品物模型属性，驱动各 tab 按属性显示（空表示未配物模型，各组件回退原始键值）
const thingProps = ref<ThingProperty[]>([])

const open = async (row: type_iot_device) => {
    deviceId.value = row.id || ''
    deviceName.value = row.name || row.sn || row.id || ''
    tab.value = 'state'
    thingProps.value = []
    visible.value = true
    // 各 tab 组件因 destroy-on-close 每次打开都重新挂载，内部数据自动重置，无需手工清空
    if (row.product_id) {
        const product = await productDetail({ id: row.product_id })
        thingProps.value = parseThingModel(product.thing_model || '').properties
    }
}

defineExpose({ open })
</script>
