<template>
    <div class="edit-popup">
        <popup ref="popupRef" :title="popupTitle" :async="true" width="550px" @confirm="handleSubmit" @close="handleClose">
            <el-form ref="formRef" :model="formData" label-width="110px" :rules="formRules">
                <el-form-item label="昵称" prop="nickname">
                    <el-input v-model="formData.nickname" placeholder="请输入昵称" clearable />
                </el-form-item>
                <el-form-item label="头像" prop="avatar">
                    <el-input v-model="formData.avatar" placeholder="头像URL" clearable />
                </el-form-item>
                <el-form-item label="手机号" prop="phone">
                    <el-input v-model="formData.phone" placeholder="手机号" clearable />
                </el-form-item>
                <el-form-item label="手机区号" prop="phone_code">
                    <el-input v-model="formData.phone_code" placeholder="如 +86" clearable />
                </el-form-item>
                <el-form-item label="状态" prop="status">
                    <el-switch v-model="formData.status" :active-value="1" :inactive-value="0" />
                </el-form-item>
            </el-form>
        </popup>
    </div>
</template>
<script lang="ts" setup>
import { ref, computed, shallowRef } from 'vue'
import type { FormInstance } from 'element-plus'
import {
    userEdit,
    userDetail,
    type type_user_edit,
    type type_user_resp
} from '@/api/user'
import Popup from '@/components/popup/index.vue'
import feedback from '@/utils/feedback'
import { useReactiveWithReset } from '@/hooks/useReactiveWithReset'
const emit = defineEmits(['success', 'close'])
const formRef = shallowRef<FormInstance>()
const popupRef = shallowRef<InstanceType<typeof Popup>>()
const mode = ref('edit')
const popupTitle = computed(() => '编辑用户')

const {
    state: formData,
    setState
} = useReactiveWithReset<type_user_edit>({
    id: '',
    nickname: '',
    avatar: '',
    phone: '',
    phone_code: '',
    status: 1
})

const formRules = {
    nickname: [{ required: true, message: '请输入昵称', trigger: ['blur'] }]
}

const handleSubmit = async () => {
    try {
        await formRef.value?.validate()
        await userEdit(formData)
        feedback.msgSuccess('操作成功')
        popupRef.value?.close()
        emit('success')
    } catch (error) {
        console.error('用户保存失败:', error)
    }
}
const open = (type = 'edit') => {
    mode.value = type
    popupRef.value?.open()
}
const setFormData = (data: Partial<type_user_edit>) => {
    setState(data)
}
const getDetail = async (row: type_user_resp) => {
    try {
        const data = await userDetail({ id: row.id })
        setFormData(data)
    } catch (error) {
        console.error('用户详情获取失败:', error)
    }
}
const handleClose = () => emit('close')
defineExpose({ open, setFormData, getDetail })
</script>
