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
                <el-form-item label="物模型" prop="thing_model">
                    <el-input v-model="formData.thing_model" type="textarea" :rows="4" placeholder="物模型 JSON" />
                </el-form-item>
                <el-form-item label="解码脚本" prop="decode_script">
                    <el-input v-model="formData.decode_script" type="textarea" :rows="3" placeholder="解码脚本" />
                </el-form-item>
                <el-form-item label="预处理脚本" prop="preprocess_script">
                    <el-input v-model="formData.preprocess_script" type="textarea" :rows="3" placeholder="预处理脚本(goja)" />
                </el-form-item>
                <el-form-item label="备注" prop="remark">
                    <el-input v-model="formData.remark" type="textarea" :rows="2" placeholder="请输入备注" :maxlength="255" />
                </el-form-item>
                <el-form-item label="状态" prop="is_stop">
                    <el-switch v-model="formData.is_stop" :active-value="0" :inactive-value="1" />
                </el-form-item>
            </el-form>
        </popup>
    </div>
</template>
<script lang="ts" setup>
import { ref, computed, shallowRef } from 'vue'
import type { FormInstance } from 'element-plus'
import {
    productAdd,
    productEdit,
    productDetail,
    productCategoryAll,
    type type_iot_product
} from '@/api/iot'
import Popup from '@/components/popup/index.vue'
import feedback from '@/utils/feedback'
import { arrayToTree } from '@/utils/util'
import { useReactiveWithReset } from '@/hooks/useReactiveWithReset'

const emit = defineEmits(['success', 'close'])
const formRef = shallowRef<FormInstance>()
const popupRef = shallowRef<InstanceType<typeof Popup>>()
const mode = ref('add')
const popupTitle = computed(() => (mode.value == 'edit' ? '编辑产品' : '新增产品'))

const { state: formData, setState } = useReactiveWithReset<type_iot_product>({
    id: '',
    name: '',
    category_id: '',
    device_type: '',
    protocol: '',
    thing_model: '',
    decode_script: '',
    preprocess_script: '',
    remark: '',
    is_stop: 0
})

const formRules = {
    name: [{ required: true, message: '请输入产品名称', trigger: ['blur'] }]
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
