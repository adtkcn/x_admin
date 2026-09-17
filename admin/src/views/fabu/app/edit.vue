<template>
    <div class="edit-popup">
        <popup
            ref="popupRef"
            :title="popupTitle"
            :async="true"
            width="650px"
            @confirm="handleSubmit"
            @close="handleClose"
        >
            <el-form ref="formRef" :model="formData" label-width="110px" :rules="formRules">
                <el-form-item label="应用名称" prop="name">
                    <el-input v-model="formData.name" placeholder="应用名称" clearable />
                </el-form-item>
                <el-form-item label="平台" prop="platform">
                    <el-select v-model="formData.platform" placeholder="选择平台" :disabled="mode === 'edit'">
                        <el-option label="iOS" value="ios" />
                        <el-option label="Android" value="android" />
                    </el-select>
                </el-form-item>
                <el-form-item label="BundleId" prop="bundle_id">
                    <el-input v-model="formData.bundle_id" placeholder="包名 / BundleId" :disabled="mode === 'edit'" clearable />
                </el-form-item>
                <el-form-item label="BundleName">
                    <el-input v-model="formData.bundle_name" placeholder="BundleName" clearable />
                </el-form-item>
                <el-form-item label="版本号">
                    <el-input v-model="formData.version" placeholder="如 1.0.0" clearable />
                </el-form-item>
                <el-form-item label="版本Code">
                    <el-input v-model="formData.version_code" type="number" placeholder="如 1" />
                </el-form-item>
                <el-form-item label="短链">
                    <el-input v-model="formData.short_url" placeholder="自定义短链，留空自动生成" clearable />
                </el-form-item>
                <el-form-item label="图标">
                    <material-picker v-model="formData.icon" :limit="1" />
                </el-form-item>
            </el-form>
        </popup>
    </div>
</template>
<script lang="ts" setup>
import { ref, computed, shallowRef } from 'vue'
import type { FormInstance } from 'element-plus'
import {
    fabuAppAdd,
    fabuAppEdit,
    fabuAppDetail,
    type type_fabu_app_edit,
    type type_fabu_app_resp
} from '@/api/fabu'
import Popup from '@/components/popup/index.vue'
import feedback from '@/utils/feedback'
import { useReactiveWithReset } from '@/hooks/useReactiveWithReset'

const emit = defineEmits(['success', 'close'])
const formRef = shallowRef<FormInstance>()
const popupRef = shallowRef<InstanceType<typeof Popup>>()
const mode = ref('add')
const popupTitle = computed(() => (mode.value === 'edit' ? '编辑应用' : '新增应用'))

const { state: formData, setState } = useReactiveWithReset<type_fabu_app_edit>({
    id: '',
    name: '',
    platform: 'android',
    bundle_id: '',
    bundle_name: '',
    version: '',
    version_code: 1,
    short_url: '',
    icon: ''
})

const formRules = {
    name: [{ required: true, message: '请输入应用名称', trigger: ['blur'] }],
    platform: [{ required: true, message: '请选择平台', trigger: ['blur'] }],
    bundle_id: [{ required: true, message: '请输入BundleId', trigger: ['blur'] }]
}

const handleSubmit = async () => {
    try {
        await formRef.value?.validate()
        const data = { ...formData }
        if (mode.value === 'edit') {
            await fabuAppEdit(data)
        } else {
            await fabuAppAdd(data)
        }
        feedback.msgSuccess('操作成功')
        popupRef.value?.close()
        emit('success')
    } catch (error) {
        console.error(error)
    }
}
const open = (type = 'add') => {
    mode.value = type
    popupRef.value?.open()
}
const setFormData = (data: Partial<type_fabu_app_edit>) => setState(data)
const getDetail = async (row: type_fabu_app_resp) => {
    try {
        const data = await fabuAppDetail({ id: row.id })
        setFormData(data)
    } catch (error) {
        console.error(error)
    }
}
const handleClose = () => emit('close')
defineExpose({ open, setFormData, getDetail })
</script>
