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
            <el-form ref="formRef" :model="formData" label-width="84px" :rules="formRules">
                <el-form-item label="流程名称" prop="flowName">
                    <el-input v-model="formData.flowName" placeholder="请输入流程名称" />
                </el-form-item>
                <el-form-item label="流程分类" prop="flowGroup">
                    <el-input v-model.number="formData.flowGroup" placeholder="请输入流程分类" />
                </el-form-item>
                <el-form-item label="流程描述" prop="flowRemark">
                    <el-input v-model="formData.flowRemark" placeholder="请输入流程描述" />
                </el-form-item>
                <el-form-item label="表单配置" prop="flowFormData">
                    <el-input
                        v-model="formData.flowFormData"
                        placeholder="请输入表单配置"
                        type="textarea"
                        :autosize="{ minRows: 4, maxRows: 6 }"
                    />
                </el-form-item>
                <el-form-item label="流程配置" prop="flowProcessData">
                    <el-input
                        v-model="formData.flowProcessData"
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
    flow_template_detail
} from '@/api/flow/flow_template'
import Popup from '@/components/popup/index.vue'
import feedback from '@/utils/feedback'
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

const formData = reactive({
    id: '',
    flowName: '',
    flowGroup: '',
    flowRemark: '',
    flowFormData: '',
    flowProcessData: ''
})

const formRules = {
    flowName: [
        {
            required: true,
            message: '请输入流程名称',
            trigger: ['blur']
        }
    ],
    flowGroup: [
        {
            required: true,
            message: '请输入流程分类',
            trigger: ['blur']
        }
    ],
    flowRemark: [
        {
            required: true,
            message: '请输入流程描述',
            trigger: ['blur']
        }
    ],
    flowFormData: [
        {
            required: true,
            message: '请输入表单配置',
            trigger: ['blur']
        }
    ],
    flowProcessData: [
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

const setFormData = async (data: Record<string, any>) => {
    for (const key in formData) {
        if (data[key] != null && data[key] != undefined) {
            //@ts-ignore
            formData[key] = data[key]
        }
    }
}

const getDetail = async (row: Record<string, any>) => {
    const data = await flow_template_detail({
        id: row.id
    })
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
