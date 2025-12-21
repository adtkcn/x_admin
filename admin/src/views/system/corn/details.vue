<template>
    <div class="edit-popup">
        <popup
            ref="popupRef"
            :title="popupTitle"
            :async="true"
            width="550px"
            :clickModalClose="true"
            :confirmButtonText="false"
  
            @close="handleClose"
        >
            <el-form ref="formRef" :model="formData" label-width="84px" :rules="formRules">
                        <el-form-item label="任务名称" prop="TaskName">
                            <span v-text="formData.TaskName"></span>
                        </el-form-item>
                        <el-form-item label="任务编码" prop="TaskCode">
                            <span v-text="formData.TaskCode"></span>
                        </el-form-item>
                        <el-form-item label="corn表达式" prop="CornExpr">
                            <span v-text="formData.CornExpr"></span>
                        </el-form-item>
                        <el-form-item label="禁用" prop="Disabled">
                            <span v-text="formData.Disabled"></span>
                        </el-form-item>
            
            </el-form>
        </popup>
    </div>
</template>
<script lang="ts" setup>
import type { FormInstance } from 'element-plus'
import {  system_corn_detail } from '@/api/system/corn'
import Popup from '@/components/popup/index.vue'
 
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
const emit = defineEmits(['close'])
const formRef = shallowRef<FormInstance>()
const popupRef = shallowRef<InstanceType<typeof Popup>>()

const popupTitle = computed(() => {
    return '预览定时任务'
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

const open = () => {
    popupRef.value?.open()
}
const getDetail = async (row: Record<string, any>) => {
     try {
        const data = await system_corn_detail(row.Id)
        setFormData(data)
     } catch (error) {}
}
const setFormData = async (data: Record<string, any>) => {
    for (const key in formData) {
        if (data[key] != null && data[key] != undefined) {
            //@ts-ignore
            formData[key] = data[key]
        }
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
 