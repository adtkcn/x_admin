<template>
    <div class="edit-popup">
        <popup
            ref="popupRef"
            :title="popupTitle"
            :async="true"
            width="680px"
            @confirm="handleSubmit"
            @close="handleClose"
        >
            <el-form ref="formRef" :model="formData" label-width="110px" :rules="formRules">
                <el-form-item label="产品名称" prop="name">
                    <el-input v-model="formData.name" placeholder="请输入产品名称" clearable :maxlength="100" />
                </el-form-item>
                <el-form-item label="所属分类" prop="category_id">
                    <el-tree-select
                        class="flex-1"
                        v-model="formData.category_id"
                        :data="categoryTreeOptions"
                        clearable
                        node-key="id"
                        :props="{ label: 'name' }"
                        check-strictly
                        :default-expand-all="true"
                        placeholder="请选择所属分类"
                    />
                </el-form-item>
                <el-form-item label="节点类型" prop="device_type">
                    <div class="w-full">
                        <el-select v-model="formData.device_type" placeholder="请选择节点类型" clearable class="w-full">
                            <el-option label="直连设备" value="DEVICE" />
                            <el-option label="网关设备" value="GATEWAY" />
                            <el-option label="边缘主机" value="EDGE" />
                        </el-select>
                        <div class="form-tips">网关设备/边缘主机可作为其他设备的网关，直连设备不能挂载子设备</div>
                    </div>
                </el-form-item>
                <el-form-item label="接入协议" prop="protocol">
                    <el-select v-model="formData.protocol" placeholder="请选择接入协议" clearable class="flex-1">
                        <el-option label="MQTT" value="mqtt" />
                        <el-option label="TCP" value="tcp" />
                        <el-option label="HTTP" value="http" />
                        <el-option label="WebSocket" value="ws" />
                    </el-select>
                </el-form-item>
                <el-form-item label="离线判定时长" prop="offline_timeout">
                    <div class="w-full">
                        <el-input-number
                            v-model="formData.offline_timeout"
                            :min="1"
                            :max="86400"
                            :step="10"
                            :controls="false"
                            placeholder="秒"
                            class="w-full"
                        />
                        <div class="form-tips">
                            必填：连续多少秒未收到报文则判为离线（1~86400 秒）；设备可单独配置，设备留 0 即用本产品的值
                        </div>
                    </div>
                </el-form-item>
                <el-form-item label="物模型" prop="thing_model">
                    <div class="w-full">
                        <div class="flex items-center gap-2">
                            <el-button type="primary" plain @click="handleOpenThingModel">
                                <template #icon><icon name="el-icon-Setting" /></template>
                                可视化设计
                            </el-button>
                            <el-tag v-if="thingModelCount" type="success">{{ thingModelCount }} 个属性</el-tag>
                            <el-tag v-else-if="formData.thing_model" type="warning">自定义 JSON</el-tag>
                            <el-tag v-else type="info">未配置</el-tag>
                            <el-button v-if="formData.thing_model" link type="danger" @click="formData.thing_model = ''">
                                清空
                            </el-button>
                        </div>
                        <div class="form-tips">
                            声明产品的属性白名单，未声明的属性上报时会被丢弃；不配置则不限制上报字段
                        </div>
                    </div>
                </el-form-item>
                <el-form-item label="规则链" prop="rule_chain_id">
                    <div class="w-full">
                        <el-select
                            v-model="formData.rule_chain_id"
                            placeholder="请选择可视化规则链（处理解码/过滤/告警/存储/通知等）"
                            clearable
                            class="w-full"
                        >
                            <el-option
                                v-for="item in ruleChainOptions"
                                :key="item.id"
                                :label="item.name"
                                :value="item.id as string"
                            />
                        </el-select>
                        <div class="form-tips">
                            解码/过滤/告警/存储/推送/通知等处理逻辑均由规则链图描述，请到「IoT · 规则链」配置；未选择则产品上报会被丢弃
                        </div>
                    </div>
                </el-form-item>
                <el-form-item label="备注" prop="remark">
                    <el-input v-model="formData.remark" type="textarea" :rows="2" placeholder="请输入备注" :maxlength="255" />
                </el-form-item>
                <el-form-item label="状态" prop="is_stop">
                    <el-switch v-model="formData.is_stop" :active-value="0" :inactive-value="1" />
                </el-form-item>
            </el-form>
        </popup>
        <thing-model-dialog ref="thingModelRef" v-model="formData.thing_model" />
    </div>
</template>
<script lang="ts" setup>
import { ref, computed, shallowRef } from 'vue'
import type { FormInstance } from 'element-plus'
import { productAdd, productEdit, productDetail, type type_iot_product } from '@/api/iot/product'
import { productCategoryAll } from '@/api/iot/product_category'
import { ruleChainAll, type type_rule_chain } from '@/api/iot/rule_chain'
import Popup from '@/components/popup/index.vue'
import ThingModelDialog from './components/thing-model.vue'
import { countProperties } from './components/thing-model'
import feedback from '@/utils/feedback'
import { arrayToTree } from '@/utils/util'
import { useReactiveWithReset } from '@/hooks/useReactiveWithReset'

const emit = defineEmits(['success', 'close'])
const formRef = shallowRef<FormInstance>()
const popupRef = shallowRef<InstanceType<typeof Popup>>()
const thingModelRef = shallowRef<InstanceType<typeof ThingModelDialog>>()
const handleOpenThingModel = () => thingModelRef.value?.open()
const mode = ref('add')
const popupTitle = computed(() => (mode.value == 'edit' ? '编辑产品' : '新增产品'))

const { state: formData, setState } = useReactiveWithReset<type_iot_product>({
    id: '',
    name: '',
    category_id: '',
    device_type: '',
    protocol: '',
    // 新增时先给与后端兜底值一致的默认时长，避免表单直接报必填错
    offline_timeout: 300,
    thing_model: '',
    rule_chain_id: '',
    remark: '',
    is_stop: 0
})

// 规则链下拉：产品引用一条可视化规则链完成解码/过滤/告警等处理
const ruleChainOptions = ref<type_rule_chain[]>([])

// 属性条数仅供展示：自定义结构的 JSON 解析不出属性时返回 0，表单仍按原文保存
const thingModelCount = computed(() => countProperties(formData.thing_model))

// 物模型必填且必须是合法 JSON：处理层按物模型白名单裁剪属性，未建模产品会被后端拒绝
const validateThingModel = (_rule: any, value: string, callback: (e?: Error) => void) => {
    if (!value || !value.trim()) return callback(new Error('物模型必须配置'))
    try {
        JSON.parse(value)
        callback()
    } catch {
        callback(new Error('物模型必须是合法 JSON'))
    }
}

const formRules = {
    name: [{ required: true, message: '请输入产品名称', trigger: ['blur'] }],
    category_id: [{ required: true, message: '请选择所属分类', trigger: ['change'] }],
    device_type: [{ required: true, message: '请选择节点类型', trigger: ['change'] }],
    protocol: [{ required: true, message: '请选择接入协议', trigger: ['change'] }],
    offline_timeout: [
        {
            required: true,
            validator: (_rule: any, value: number, callback: (e?: Error) => void) =>
                value >= 1 && value <= 86400
                    ? callback()
                    : callback(new Error('请输入 1~86400 之间的离线判定时长')),
            trigger: ['blur', 'change']
        }
    ],
    thing_model: [{ required: true, validator: validateThingModel, trigger: ['blur', 'change'] }]
}

const categoryFlat = ref<any[]>([])
const categoryTreeOptions = computed(() =>
    categoryFlat.value.length ? arrayToTree<any>(categoryFlat.value.map((i) => ({ ...i })) as any[], '') : []
)

const handleSubmit = async () => {
    await formRef.value?.validate()
    mode.value == 'edit' ? await productEdit(formData) : await productAdd(formData)
    popupRef.value?.close()
    feedback.msgSuccess('操作成功')
    emit('success')
}

const open = async (type = 'add') => {
    mode.value = type
    categoryFlat.value = await productCategoryAll()
    ruleChainOptions.value = await ruleChainAll()
    popupRef.value?.open()
}

const setFormData = (data: Partial<type_iot_product>) => {
    setState(data)
}

const getDetail = async (row: type_iot_product) => {
    const data = await productDetail({ id: row.id as string })
    setState(data)
}

const handleClose = () => {
    emit('close')
}

defineExpose({ open, setFormData, getDetail })
</script>
