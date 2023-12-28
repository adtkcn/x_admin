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
                <el-form-item label="审批流" prop="templateId">
                    <el-select
                        v-model="formData.templateId"
                        placeholder="请选择审批流"
                        style="width: 100%"
                        @change="handleTemplateChange"
                    >
                        <el-option
                            v-for="(item, index) in flow_template"
                            :key="index"
                            :label="item.flowName"
                            :value="item.id"
                            clearable
                        />
                    </el-select>
                </el-form-item>

                <el-form-item label="流程名称" prop="flowName">
                    <el-input v-model="formData.flowName" placeholder="请输入流程名称" />
                </el-form-item>
            </el-form>
        </popup>
    </div>
</template>
<script lang="ts" setup>
import type { FormInstance } from 'element-plus'
import { flow_apply_edit, flow_apply_add, flow_apply_detail } from '@/api/flow/flow_apply'
import { flow_template_lists_all } from '@/api/flow/flow_template'

import Popup from '@/components/popup/index.vue'
import feedback from '@/utils/feedback'
import { shallowRef, ref, computed, reactive } from 'vue'
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
    return mode.value == 'edit' ? '编辑申请流程' : '新增申请流程'
})

const formData = reactive({
    id: '',
    templateId: '',
    // applyUserId: '',
    // applyUserNickname: '',
    flowName: '',
    // flowGroup: '',
    // flowRemark: '',
    // flowFormData: '',
    // flowProcessData: '',
    status: 0
    // formValue: ''
})

const formRules = {
    id: [
        {
            required: true,
            message: '请输入',
            trigger: ['blur']
        }
    ],
    templateId: [
        {
            required: true,
            message: '请输入模板',
            trigger: ['blur']
        }
    ],
    applyUserId: [
        {
            required: true,
            message: '请输入申请人id',
            trigger: ['blur']
        }
    ],
    applyUserNickname: [
        {
            required: true,
            message: '请输入申请人昵称',
            trigger: ['blur']
        }
    ],
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
    ],
    status: [
        {
            required: true,
            message: '请选择状态',
            trigger: ['blur']
        }
    ]
}
const flow_template = ref([])
const get_flow_template = async () => {
    flow_template.value = await flow_template_lists_all()
}
function handleTemplateChange(id: number) {
    console.log(id)
    flow_template.value.find((item: any) => {
        if (item.id == id) {
            formData.flowName = item.flowName
            // formData.flowGroup = item.flowGroup
            // formData.flowRemark = item.flowRemark
            // formData.flowFormData = item.flowFormData
            // formData.flowProcessData = item.flowProcessData
            return true
        }
    })
}
get_flow_template()
const handleSubmit = async () => {
    await formRef.value?.validate()
    const data: any = { ...formData }
    // if ( !data.id) {
    //     delete data.id
    // }
    mode.value == 'edit' ? await flow_apply_edit(data) : await flow_apply_add(data)
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
    const data = await flow_apply_detail({
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
