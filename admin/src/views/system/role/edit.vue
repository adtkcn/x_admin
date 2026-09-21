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
                label-width="60px"
            >
                <el-form-item label="名称" prop="name">
                    <el-input
                        class="ls-input"
                        v-model="formData.name"
                        placeholder="请输入名称"
                        clearable
                    />
                </el-form-item>
                <el-form-item label="备注" prop="remark">
                    <el-input
                        v-model="formData.remark"
                        type="textarea"
                        :autosize="{ minRows: 4, maxRows: 6 }"
                        placeholder="请输入备注"
                        maxlength="200"
                        show-word-limit
                    />
                </el-form-item>
                <el-form-item label="排序" prop="sort">
                    <el-input-number v-model="formData.sort" />
                </el-form-item>
                <el-form-item label="状态" prop="sort">
                    <el-radio-group v-model="formData.is_disable">
                        <el-radio :value="0">正常</el-radio>
                        <el-radio :value="1">停用</el-radio>
                    </el-radio-group>
                </el-form-item>
            </el-form>
        </popup>
    </div>
</template>
<script lang="ts" setup>
import { ref, computed, useTemplateRef } from 'vue'
import type { FormInstance } from 'element-plus'
import {
    roleAdd,
    roleDetail,
    roleEdit,
    type type_system_role_add,
    type type_system_role_edit,
    type type_system_role_resp
} from '@/api/perms/role'
import Popup from '@/components/popup/index.vue'
import feedback from '@/utils/feedback'
import { useReactiveWithReset } from '@/hooks/useReactiveWithReset'
const emit = defineEmits(['success', 'close'])
const formRef = useTemplateRef<FormInstance>('formRef')
const popupRef = useTemplateRef<InstanceType<typeof Popup>>('popupRef')
const mode = ref('add')
const popupTitle = computed(() => {
    return mode.value == 'edit' ? '编辑角色' : '新增角色'
})

type type_role_form = type_system_role_edit & {
    menus: string[]
}

const {
    state: formData,
    reset,
    setState
} = useReactiveWithReset<type_role_form>({
    id: '',
    name: '',
    remark: '',
    sort: 0,
    is_disable: 0,
    menus: []
})

const rules = {
    name: [
        {
            required: true,
            message: '请输入名称',
            trigger: ['blur']
        }
    ]
}

const handleSubmit = async () => {
    try {
        await formRef.value?.validate()
        const params = { ...formData, menuIds: formData.menus.join() }
        mode.value == 'edit' ? await roleEdit(params) : await roleAdd(params)
        popupRef.value?.close()
        feedback.msgSuccess('操作成功')
        emit('success')
    } catch (error) {
        console.error('角色保存失败:', error)
    }
}

const handleClose = () => {
    emit('close')
}

const open = (type = 'add') => {
    mode.value = type
    popupRef.value?.open()
}

const setFormData = async (row: type_system_role_resp) => {
    try {
        const data = await roleDetail({
            id: row.id
        })
        setState(data)
    } catch (error) {
        console.error('角色详情获取失败:', error)
    }
}

defineExpose({
    open,
    setFormData
})
</script>
