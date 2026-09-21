<template>
    <div class="edit-popup">
        <popup
            ref="popupRef"
            :title="popupTitle"
            :async="true"
            width="550px"
            @confirm="handleSubmit"
            @close="handleClose"
        >
            <el-form
                class="ls-form"
                ref="formRef"
                :rules="rules"
                :model="formData"
                label-width="110px"
            >
                <el-form-item label="数据名称" prop="name">
                    <el-input v-model="formData.name" placeholder="请输入数据名称" clearable />
                </el-form-item>
                <el-form-item label="数据值" prop="value">
                    <el-input v-model="formData.value" placeholder="请输入数据值" clearable />
                </el-form-item>
                <el-form-item label="颜色" prop="color">
                    <el-color-picker v-model="formData.color" />
                </el-form-item>
                <el-form-item label="排序" prop="sort">
                    <div>
                        <el-input-number v-model="formData.sort" :min="0" :max="9999" />
                        <div class="form-tips">数值越大越排前</div>
                    </div>
                </el-form-item>
                <el-form-item label="状态" required prop="status">
                    <el-radio-group v-model="formData.status">
                        <el-radio :value="1">正常</el-radio>
                        <el-radio :value="0">停用</el-radio>
                    </el-radio-group>
                </el-form-item>
                <el-form-item label="备注" prop="remark">
                    <el-input
                        v-model="formData.remark"
                        type="textarea"
                        :autosize="{ minRows: 4, maxRows: 6 }"
                        clearable
                        maxlength="200"
                        show-word-limit
                    />
                </el-form-item>
            </el-form>
        </popup>
    </div>
</template>
<script lang="ts" setup>
import type { FormInstance } from 'element-plus'
import Popup from '@/components/popup/index.vue'
import { dictDataAdd, dictDataEdit, type type_setting_dict_data_edit } from '@/api/setting/dict'
import feedback from '@/utils/feedback'
import { computed, ref, shallowRef } from 'vue'
import { useReactiveWithReset } from '@/hooks/useReactiveWithReset'

const emit = defineEmits(['success', 'close'])
const formRef = shallowRef<FormInstance>()
const popupRef = shallowRef<InstanceType<typeof Popup>>()
const mode = ref('add')
const popupTitle = computed(() => {
    return mode.value == 'edit' ? '编辑字典数据' : '新增字典数据'
})
const { state: formData, setState } = useReactiveWithReset<type_setting_dict_data_edit>({
    id: '',
    type_id: '',
    name: '',
    value: '',
    color: '',
    sort: 0,
    status: 1,
    remark: ''
})

const rules = {
    name: [
        {
            required: true,
            message: '请输入数据名称',
            trigger: ['blur']
        }
    ],
    value: [
        {
            required: true,
            message: '请输入数据值',
            trigger: ['blur']
        }
    ]
}

const handleSubmit = async () => {
    try {
        await formRef.value?.validate()
        mode.value == 'edit' ? await dictDataEdit(formData) : await dictDataAdd(formData)
        popupRef.value?.close()
        feedback.msgSuccess('操作成功')
        emit('success')
    } catch (error) {
        console.error('字典数据保存失败:', error)
    }
}

const handleClose = () => {
    emit('close')
}

const open = (type = 'add') => {
    mode.value = type
    popupRef.value?.open()
}

const setFormData = (data: Partial<type_setting_dict_data_edit>) => {
    setState(data as type_setting_dict_data_edit)
}

defineExpose({
    open,
    setFormData
})
</script>
