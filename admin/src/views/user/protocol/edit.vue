<template>
    <div class="edit-popup">
        <popup
            ref="popupRef"
            :title="popuptitle"
            :async="true"
            width="900px"
            :clickModalClose="true"
            @confirm="handleSubmit"
            @close="handleClose"
        >
            <el-form ref="formRef" :model="formData" label-width="110px" :rules="formRules">
                <el-form-item label="标识" prop="tag" borderBottom>
                    <el-input v-model="formData.tag" placeholder="请输入标识" />
                </el-form-item>
                <el-form-item label="版本" prop="Version" borderBottom>
                    <el-input v-model="formData.version" type="number" placeholder="请输入版本" />
                </el-form-item>
                <el-form-item label="标题" prop="title">
                    <el-input v-model="formData.title" placeholder="请输入标题" />
                </el-form-item>
                <el-form-item label="协议内容" prop="content">
                    <editor v-model="formData.content" :height="500" />
                </el-form-item>
            </el-form>
        </popup>
    </div>
</template>
<script lang="ts" setup>
import { ref, computed, useTemplateRef } from 'vue'
import type { FormInstance, FormRules } from 'element-plus'
import {
    user_protocol_edit,
    user_protocol_add,
    user_protocol_detail,
    type type_user_protocol_edit
} from '@/api/user/protocol'
import Popup from '@/components/popup/index.vue'
import feedback from '@/utils/feedback'
import type { PropType } from 'vue'
import { useReactiveWithReset } from '@/hooks/useReactiveWithReset'
defineProps({
    dictData: {
        type: Object as PropType<Record<string, any[]>>,
        default: () => ({})
    },
    listAllData: {
        type: Object as PropType<Record<string, any[]>>,
        default: () => ({})
    }
})
const emit = defineEmits(['success', 'close'])
const formRef = useTemplateRef<FormInstance>('formRef')
const popupRef = useTemplateRef<InstanceType<typeof Popup>>('popupRef')
const mode = ref('add')
const popuptitle = computed(() => {
    return mode.value == 'edit' ? '编辑用户协议' : '新增用户协议'
})

const { state: formData, setState } = useReactiveWithReset<type_user_protocol_edit>({
    tag: undefined,
    version: undefined,
    id: undefined,
    title: undefined,
    content: undefined
})

const formRules: FormRules = {
    // id: [
    //     {
    //         required: true,
    //         message: '请输入',
    //         trigger: ['blur']
    //     }
    // ]
    // title: [
    //     {
    //         required: true,
    //         message: '请输入标题',
    //         trigger: ['blur']
    //     }
    // ],
    // content: [
    //     {
    //         required: true,
    //         message: '请输入协议内容',
    //         trigger: ['blur']
    //     }
    // ]
    // Sort: [
    //     {
    //         required: true,
    //         message: '请输入排序',
    //         trigger: ['blur']
    //     }
    // ]
}

const handleSubmit = async () => {
    try {
        await formRef.value?.validate()
        const data: any = { ...formData }
        // delete data.title
        data.Version = data.Version ? Number(data.Version) : null
        // data.Version = null
        mode.value == 'edit' ? await user_protocol_edit(data) : await user_protocol_add(data)
        popupRef.value?.close()
        feedback.msgSuccess('操作成功')
        emit('success')
    } catch (error) {
        console.error('协议保存失败:', error)
    }
}

const open = (type = 'add') => {
    mode.value = type
    popupRef.value?.open()
}

const setFormData = async (data: type_user_protocol_edit) => {
    setState(data)
}

const getDetail = async (row: Record<string, any>) => {
    try {
        const data = await user_protocol_detail(row.id)
        setFormData(data)
    } catch (error) {
        console.error('协议详情获取失败:', error)
    }
}

const handleClose = () => {
    emit('close')
}

defineExpose({
    open,
    setFormData,
    getDetail
})
</script>
