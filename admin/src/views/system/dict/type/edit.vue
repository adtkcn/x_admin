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
                <el-form-item label="字典名称" prop="dict_name">
                    <el-input v-model="formData.dict_name" placeholder="请输入字典名称" clearable />
                </el-form-item>
                <el-form-item label="字典类型" prop="dict_type">
                    <el-input v-model="formData.dict_type" placeholder="请输入字典类型" clearable />
                </el-form-item>
                <el-form-item label="字典状态" required prop="dict_status">
                    <el-radio-group v-model="formData.dict_status">
                        <el-radio :value="1">正常</el-radio>
                        <el-radio :value="0">停用</el-radio>
                    </el-radio-group>
                </el-form-item>
                <el-form-item label="备注" prop="dict_remark">
                    <el-input
                        v-model="formData.dict_remark"
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
import { ref, computed, shallowRef } from 'vue'
import type { FormInstance } from 'element-plus'
import Popup from '@/components/popup/index.vue'
import {
    dictTypeAdd,
    dictTypeEdit,
    type type_setting_dict_type_edit,
    type type_setting_dict_type_resp
} from '@/api/setting/dict'
import feedback from '@/utils/feedback'
import { useReactiveWithReset } from '@/hooks/useReactiveWithReset'
const emit = defineEmits(['success', 'close'])
const formRef = shallowRef<FormInstance>()
const popupRef = shallowRef<InstanceType<typeof Popup>>()
const mode = ref('add')
const popupTitle = computed(() => {
    return mode.value == 'edit' ? '编辑字典类型' : '新增字典类型'
})

const {
    state: formData,
    setState
} = useReactiveWithReset<type_setting_dict_type_edit>({
    id: '',
    dict_name: '',
    dict_type: '',
    dict_status: 1,
    dict_remark: ''
})

const rules = {
    dict_name: [
        {
            required: true,
            message: '请输入字典名称',
            trigger: ['blur']
        }
    ],
    dict_type: [
        {
            required: true,
            message: '请输入字典类型',
            trigger: ['blur']
        }
    ]
}

const handleSubmit = async () => {
    try {
        await formRef.value?.validate()
        mode.value == 'edit' ? await dictTypeEdit(formData) : await dictTypeAdd(formData)
        popupRef.value?.close()
        feedback.msgSuccess('操作成功')
        emit('success')
    } catch (error) {
        console.error('字典类型保存失败:', error)
    }
}

const handleClose = () => {
    emit('close')
}

const open = (type = 'add') => {
    mode.value = type
    popupRef.value?.open()
}

const setFormData = (data: type_setting_dict_type_resp) => {
    setState(data)
}

defineExpose({
    open,
    setFormData
})
</script>
