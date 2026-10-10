<template>
    <div class="edit-popup">
        <popup
            ref="popupRef"
            :title="popupTitle"
            :async="true"
            width="600px"
            @confirm="handleSubmit"
            @close="handleClose"
        >
            <el-form ref="formRef" :model="formData" label-width="110px" :rules="formRules">
                <el-form-item label="设备名称" prop="name">
                    <el-input v-model="formData.name" placeholder="请输入设备名称" clearable :maxlength="100" />
                </el-form-item>
                <el-form-item label="设备编号" prop="sn">
                    <el-input v-model="formData.sn" placeholder="请输入设备编号(SN)" clearable :maxlength="64" />
                </el-form-item>
                <el-form-item label="所属产品" prop="product_id">
                    <el-select v-model="formData.product_id" placeholder="请选择产品" clearable class="flex-1">
                        <el-option
                            v-for="p in products"
                            :key="p.id"
                            :label="p.name"
                            :value="(p.id as string)"
                        />
                    </el-select>
                </el-form-item>
                <el-form-item label="所属区域" prop="space_id">
                    <el-tree-select
                        class="flex-1"
                        v-model="formData.space_id"
                        :data="spaceTreeOptions"
                        clearable
                        node-key="id"
                        :props="{ label: 'name' }"
                        check-strictly
                        :default-expand-all="true"
                        placeholder="请选择区域"
                    />
                </el-form-item>
                <el-form-item label="网关设备" prop="gateway_id">
                    <div class="w-full">
                        <el-select v-model="formData.gateway_id" placeholder="请选择网关设备" clearable class="w-full">
                            <el-option
                                v-for="d in gatewayOptions"
                                :key="d.id"
                                :label="d.name"
                                :value="(d.id as string)"
                            />
                        </el-select>
                        <div class="form-tips">仅可选择所属产品为网关设备/边缘主机的设备</div>
                    </div>
                </el-form-item>
                <el-form-item label="离线判定时长" prop="offline_timeout">
                    <div class="w-full">
                        <el-input-number
                            v-model="formData.offline_timeout"
                            :min="0"
                            :max="86400"
                            :step="10"
                            :controls="false"
                            placeholder="秒"
                            class="w-full"
                        />
                        <div class="form-tips">
                            连续多少秒未收到本机报文则判为离线；留 0 表示跟随所属产品的配置，填具体值则为本机自定义（不随产品变更）
                        </div>
                    </div>
                </el-form-item>
                <el-form-item label="经纬度">
                    <div class="flex items-center gap-2 w-full">
                        <el-input-number v-model="formData.longitude" :precision="7" :step="0.0001" :controls="false" placeholder="经度" class="flex-1" />
                        <el-input-number v-model="formData.latitude" :precision="7" :step="0.0001" :controls="false" placeholder="纬度" class="flex-1" />
                        <el-button type="primary" plain @click="handleOpenMap">
                            <template #icon><icon name="el-icon-Location" /></template>
                            选点
                        </el-button>
                    </div>
                </el-form-item>
                <el-form-item v-if="formData.access_token" label="接入凭证">
                    <div class="flex items-center gap-2 w-full">
                        <el-input v-model="formData.access_token" readonly class="flex-1" />
                        <el-popover placement="right-start" :width="440" trigger="hover">
                            <template #reference>
                                <el-button link type="primary">
                                    <template #icon><icon name="el-icon-Document" /></template>
                                    接入示例
                                </el-button>
                            </template>
                            <div class="access-example">
                                <div class="access-example__desc">
                                    设备通过请求头 <code>X-Device-Sn</code>、<code>X-Access-Token</code> 鉴权上报，与后台登录态无关；所属产品的接入协议须为
                                    <code>http</code>（或留空）。示例已填入本设备凭证，可直接复制到设备端联调。
                                </div>
                                <div v-for="ex in accessExamples" :key="ex.title" class="access-example__block">
                                    <div class="access-example__head">
                                        <span class="access-example__title">{{ ex.title }}</span>
                                        <el-button v-copy="ex.text" link type="primary" size="small">
                                            <template #icon><icon name="el-icon-CopyDocument" /></template>
                                            复制
                                        </el-button>
                                    </div>
                                    <pre class="access-example__code">{{ ex.text }}</pre>
                                </div>
                            </div>
                        </el-popover>
                    </div>
                </el-form-item>
                <el-form-item label="备注" prop="remark">
                    <el-input v-model="formData.remark" type="textarea" :rows="2" placeholder="请输入备注" :maxlength="255" />
                </el-form-item>
            </el-form>
        </popup>
        <map-picker ref="mapPickerRef" :model-value="currentCoord" @confirm="handleMapConfirm" />
    </div>
</template>
<script lang="ts" setup>
import { ref, computed, shallowRef } from 'vue'
import type { FormInstance } from 'element-plus'
import {
    deviceAdd,
    deviceEdit,
    deviceDetail,
    deviceLists,
    productAll,
    deviceSpaceAll,
    nodeTypeSupportsGateway,
    type type_iot_device,
    type type_iot_product
} from '@/api/iot'
import Popup from '@/components/popup/index.vue'
import MapPicker from '@/components/map-picker/index.vue'
import type { MapCoordinate } from '@/components/map-picker/type'
import feedback from '@/utils/feedback'
import { arrayToTree } from '@/utils/util'
import { useReactiveWithReset } from '@/hooks/useReactiveWithReset'

const emit = defineEmits(['success', 'close'])
const formRef = shallowRef<FormInstance>()
const popupRef = shallowRef<InstanceType<typeof Popup>>()
const mode = ref('add')
const popupTitle = computed(() => (mode.value == 'edit' ? '编辑设备' : '新增设备'))

const { state: formData, setState } = useReactiveWithReset<type_iot_device>({
    id: '',
    name: '',
    sn: '',
    product_id: '',
    space_id: '',
    gateway_id: '',
    status: 0,
    offline_timeout: 0,
    longitude: 0,
    latitude: 0,
    remark: ''
})

const formRules = {
    name: [{ required: true, message: '请输入设备名称', trigger: ['blur'] }],
    sn: [{ required: true, message: '请输入设备编号', trigger: ['blur'] }],
    product_id: [{ required: true, message: '请选择所属产品', trigger: ['change'] }]
}

const products = ref<type_iot_product[]>([])
const spaceTreeOptions = ref<any[]>([])
const gatewayOptions = ref<type_iot_device[]>([])

const mapPickerRef = shallowRef<InstanceType<typeof MapPicker>>()
// 当前经纬度（供地图选点弹窗回显）
const currentCoord = computed<MapCoordinate>(() => ({
    longitude: formData.longitude || 0,
    latitude: formData.latitude || 0
}))
const handleOpenMap = () => mapPickerRef.value?.open()
const handleMapConfirm = (val: MapCoordinate) => {
    formData.longitude = val.longitude
    formData.latitude = val.latitude
}

// 接入本设备示例：将真实 sn / access_token 填入 curl，供设备端直接复制联调
const accessExamples = computed(() => {
    const sn = formData.sn || '<设备编号>'
    const token = formData.access_token || '<接入凭证>'
    return [
        {
            title: '属性上报（telemetry）',
            text: [
                'curl -X POST "http://<平台接入地址>/api/iot/telemetry" \\',
                '  -H "Content-Type: application/json" \\',
                `  -H "X-Device-Sn: ${sn}" \\`,
                `  -H "X-Access-Token: ${token}" \\`,
                '  -d \'{"values":{"temperature":26.3},"ts":1710000000000}\''
            ].join('\n')
        },
        {
            title: '心跳（heartbeat）',
            text: [
                'curl -X POST "http://<平台接入地址>/api/iot/heartbeat" \\',
                '  -H "Content-Type: application/json" \\',
                `  -H "X-Device-Sn: ${sn}" \\`,
                `  -H "X-Access-Token: ${token}" \\`,
                '  -d \'{"ts":1710000000000}\''
            ].join('\n')
        }
    ]
})

const handleSubmit = async () => {
    await formRef.value?.validate()
    // el-input-number 被清空时是 undefined/null：归 0 再提交（0 = 跟随产品值），避免后端收到 null
    const payload = { ...formData, offline_timeout: formData.offline_timeout || 0 }
    mode.value == 'edit' ? await deviceEdit(payload) : await deviceAdd(payload)
    popupRef.value?.close()
    feedback.msgSuccess('操作成功')
    emit('success')
}

const open = async (type = 'add') => {
    mode.value = type
    products.value = await productAll()
    const spaces = await deviceSpaceAll()
    spaceTreeOptions.value = arrayToTree<any>(spaces as any[], '')
    const res = await deviceLists({ pageNo: 1, pageSize: 1000 })
    // 网关设备候选：仅保留所属产品具备网关能力的设备，并排除当前编辑设备自身
    gatewayOptions.value = (res.lists || []).filter(
        (d) => d.id !== formData.id && nodeTypeSupportsGateway(d.product_node_type)
    )
    popupRef.value?.open()
}

const setFormData = (data: Partial<type_iot_device>) => {
    setState(data)
}

const getDetail = async (row: type_iot_device) => {
    const data = await deviceDetail({ id: row.id as string })
    setState(data)
}

const handleClose = () => {
    emit('close')
}

defineExpose({ open, setFormData, getDetail })
</script>
<style lang="scss" scoped>
.access-example {
    line-height: 1.6;
    &__desc {
        margin-bottom: 10px;
        font-size: 12px;
        color: var(--el-text-color-secondary);
        code {
            padding: 1px 4px;
            font-size: 12px;
            color: var(--el-color-primary);
            background: var(--el-fill-color-light);
            border-radius: 3px;
        }
    }
    &__block + &__block {
        margin-top: 10px;
    }
    &__head {
        display: flex;
        align-items: center;
        justify-content: space-between;
        margin-bottom: 4px;
    }
    &__title {
        font-size: 13px;
        font-weight: 500;
    }
    &__code {
        margin: 0;
        padding: 10px 12px;
        font-family: Consolas, Monaco, monospace;
        font-size: 12px;
        line-height: 1.5;
        white-space: pre;
        overflow-x: auto;
        background: var(--el-fill-color-light);
        border-radius: 6px;
    }
}
</style>
