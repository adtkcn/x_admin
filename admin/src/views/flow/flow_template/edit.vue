<template>
    <div class="edit-popup">
        <popup
            ref="popupRef"
            :title="popupTitle"
            :async="true"
            width="550px"
            :clickModalClose="true"
            @confirm="handleSubmit"
            @close="handleClose"
        >
            <el-form ref="formRef" :model="formData" label-width="110px" :rules="formRules">
                <el-form-item label="流程名称" prop="flow_name">
                    <el-input v-model="formData.flow_name" placeholder="请输入流程名称" />
                </el-form-item>
                <el-form-item label="流程分类" prop="flow_group">
                    <el-input v-model="formData.flow_group" placeholder="请输入流程分类" />
                </el-form-item>
                <el-form-item label="流程描述" prop="flow_remark">
                    <el-input v-model="formData.flow_remark" placeholder="请输入流程描述" />
                </el-form-item>
                <el-form-item label="表单配置" prop="flow_form_data">
                    <el-input
                        v-model="formData.flow_form_data"
                        placeholder="请输入表单配置"
                        type="textarea"
                        :autosize="{ minRows: 4, maxRows: 6 }"
                    />
                </el-form-item>
                <el-form-item label="流程配置" prop="flow_process_data">
                    <el-input
                        v-model="formData.flow_process_data"
                        placeholder="请输入流程配置"
                        type="textarea"
                        :autosize="{ minRows: 4, maxRows: 6 }"
                    />
                </el-form-item>
            </el-form>
        </popup>
    </div>
</template>
<script lang="ts" setup>
import type { FormInstance } from 'element-plus'
import {
    flow_template_edit,
    flow_template_add,
    flow_template_detail,
    type type_flow_template_edit
} from '@/api/flow/flow_template'
import Popup from '@/components/popup/index.vue'
import feedback from '@/utils/feedback'
import { computed, ref, shallowRef } from 'vue'
import { useReactiveWithReset } from '@/hooks/useReactiveWithReset'
import type { PropType } from 'vue'
defineProps({
    dictData: {
        type: Object as PropType<Record<string, any[]>>,
        default: () => ({})
    }
})
const emit = defineEmits(['success', 'close'])
const formRef = shallowRef<FormInstance>()
const popupRef = shallowRef<InstanceType<typeof Popup>>()
const mode = ref('add')
const popupTitle = computed(() => {
    return mode.value == 'edit' ? '编辑流程模板' : '新增流程模板'
})

const { state: formData, setState } = useReactiveWithReset<type_flow_template_edit>({
    id: '',
    flow_name: undefined,
    flow_group: undefined,
    flow_remark: undefined,
    flow_form_data: undefined,
    flow_process_data: undefined
})

const formRules = {
    flow_name: [
        {
            required: true,
            message: '请输入流程名称',
            trigger: ['blur']
        }
    ],
    flow_group: [
        {
            required: true,
            message: '请输入流程分类',
            trigger: ['blur']
        }
    ],
    flow_remark: [
        {
            required: true,
            message: '请输入流程描述',
            trigger: ['blur']
        }
    ],
    flow_form_data: [
        {
            required: true,
            message: '请输入表单配置',
            trigger: ['blur']
        }
    ],
    flow_process_data: [
        {
            required: true,
            message: '请输入流程配置',
            trigger: ['blur']
        }
    ]
}

const handleSubmit = async () => {
    await formRef.value?.validate()
    const data: any = { ...formData }
    mode.value == 'edit' ? await flow_template_edit(data) : await flow_template_add(data)
    popupRef.value?.close()
    feedback.msgSuccess('操作成功')
    emit('success')
}

const open = (type = 'add') => {
    mode.value = type
    popupRef.value?.open()
}

const setFormData = async (data: type_flow_template_edit) => {
    setState(data)
}

const getDetail = async (row: Record<string, any>) => {
    const data = await flow_template_detail(row.id)
    setFormData(data)
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
