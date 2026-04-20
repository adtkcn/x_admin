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
                <el-form-item label="字典名称" prop="dictName">
                    <el-input v-model="formData.dictName" placeholder="请输入字典名称" clearable />
                </el-form-item>
                <el-form-item label="字典类型" prop="dictType">
                    <el-input v-model="formData.dictType" placeholder="请输入字典类型" clearable />
                </el-form-item>
                <el-form-item label="字典状态" required prop="dictStatus">
                    <el-radio-group v-model="formData.dictStatus">
                        <el-radio :value="1">正常</el-radio>
                        <el-radio :value="0">停用</el-radio>
                    </el-radio-group>
                </el-form-item>
                <el-form-item label="备注" prop="dictRemark">
                    <el-input
                        v-model="formData.dictRemark"
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
import { ref, computed, reactive, shallowRef } from 'vue'
import type { FormInstance } from 'element-plus'
import Popup from '@/components/popup/index.vue'
import {
    dictTypeAdd,
    dictTypeEdit,
    type type_setting_dict_type_edit,
    type type_setting_dict_type_resp
} from '@/api/setting/dict'
import feedback from '@/utils/feedback'
const emit = defineEmits(['success', 'close'])
const formRef = shallowRef<FormInstance>()
const popupRef = shallowRef<InstanceType<typeof Popup>>()
const mode = ref('add')
const popupTitle = computed(() => {
    return mode.value == 'edit' ? '编辑字典类型' : '新增字典类型'
})

const formData = reactive<type_setting_dict_type_edit>({
    id: '',
    dictName: '',
    dictType: '',
    dictStatus: 1,
    dictRemark: ''
})

const rules = {
    dictName: [
        {
            required: true,
            message: '请输入字典名称',
            trigger: ['blur']
        }
    ],
    dictType: [
        {
            required: true,
            message: '请输入字典类型',
            trigger: ['blur']
        }
    ]
}

const handleSubmit = async () => {
    await formRef.value?.validate()
    mode.value == 'edit' ? await dictTypeEdit(formData) : await dictTypeAdd(formData)
    popupRef.value?.close()
    feedback.msgSuccess('操作成功')
    emit('success')
}

const handleClose = () => {
    emit('close')
}

const open = (type = 'add') => {
    mode.value = type
    popupRef.value?.open()
}

const setFormData = (data: type_setting_dict_type_resp) => {
    for (const key in formData) {
        // if (Object.hasOwnProperty.call(formData, key)) {
        if (
            data[key as keyof type_setting_dict_type_resp] != null &&
            data[key as keyof type_setting_dict_type_resp] != undefined
        ) {
            formData[key] = data[key as keyof type_setting_dict_type_resp]
        }
        // }
    }
}

defineExpose({
    open,
    setFormData
})
</script>
