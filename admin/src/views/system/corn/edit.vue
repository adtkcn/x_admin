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
                        <el-form-item label="任务名称" prop="TaskName">
                            <el-input v-model="formData.TaskName" placeholder="请输入任务名称" />
                        </el-form-item>
                        <el-form-item label="任务编码" prop="TaskCode">
                            <el-input v-model="formData.TaskCode" placeholder="请输入任务编码" />
                        </el-form-item>
                        <el-form-item label="corn表达式" prop="CornExpr">
                            <el-input v-model="formData.CornExpr" placeholder="请输入corn表达式" />
                        </el-form-item>
                        <el-form-item label="禁用" prop="Disabled">
                            <el-input v-model="formData.Disabled" type="number" placeholder="请输入禁用" />
                        </el-form-item>
            
            </el-form>
        </popup>
    </div>
</template>
<script lang="ts" setup>
import type { FormInstance } from 'element-plus'
import {  system_corn_edit, system_corn_add, system_corn_detail } from '@/api/system/corn'
import Popup from '@/components/popup/index.vue'
import feedback from '@/utils/feedback'
import { ref, shallowRef, computed, reactive } from 'vue'
import type { PropType } from 'vue'
defineProps({
    dictData: {
        type: Object as PropType<Record<string, any[]>>,
        default: () => ({})
    },
    listAllData:{
        type: Object as PropType<Record<string, any[]>>,
        default: () => ({})
    }
})
const emit = defineEmits(['success', 'close'])
const formRef = shallowRef<FormInstance>()
const popupRef = shallowRef<InstanceType<typeof Popup>>()
const mode = ref('add')
const popupTitle = computed(() => {
    return mode.value == 'edit' ? '编辑定时任务' : '新增定时任务'
})

const formData = reactive({
    Id: null,
    TaskName: null,
    TaskCode: null,
    CornExpr: null,
    Disabled: null,
})

const formRules = {
    Id: [
        {
            required: true,
            message: '请输入taskid',
            trigger: ['blur']
        }
    ],
    TaskName: [
        {
            required: true,
            message: '请输入任务名称',
            trigger: ['blur']
        }
    ],
    TaskCode: [
        {
            required: true,
            message: '请输入任务编码',
            trigger: ['blur']
        }
    ],
    CornExpr: [
        {
            required: true,
            message: '请输入corn表达式',
            trigger: ['blur']
        }
    ],
    Disabled: [
        {
            required: true,
            message: '请输入禁用',
            trigger: ['blur']
        }
    ],
}

const handleSubmit = async () => {
     try {
        await formRef.value?.validate()
        const data: any = { ...formData }
        mode.value == 'edit' ? await system_corn_edit(data) : await system_corn_add(data)
        popupRef.value?.close()
        feedback.msgSuccess('操作成功')
        emit('success')
     } catch (error) {}
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
     try {
        const data = await system_corn_detail(row.Id)
        setFormData(data)
     } catch (error) {}
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
