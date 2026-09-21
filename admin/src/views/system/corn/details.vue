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
            <el-form ref="formRef" :model="formData" label-width="110px" :rules="formRules">
                <el-form-item label="任务名称" prop="task_name">
                    <span v-text="formData.task_name"></span>
                </el-form-item>
                <el-form-item label="任务编码" prop="task_code">
                    <span v-text="formData.task_code"></span>
                </el-form-item>
                <el-form-item label="corn表达式" prop="corn_expr">
                    <span v-text="formData.corn_expr"></span>
                </el-form-item>
                <el-form-item label="禁用" prop="status">
                    <span v-text="formData.status"></span>
                </el-form-item>
            </el-form>
        </popup>
    </div>
</template>
<script lang="ts" setup>
import type { FormInstance } from 'element-plus'
import { system_corn_detail } from '@/api/system/corn'
import Popup from '@/components/popup/index.vue'

import { ref, shallowRef, computed, reactive } from 'vue'
import type { PropType } from 'vue'
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
const emit = defineEmits(['close'])
const formRef = shallowRef<FormInstance>()
const popupRef = shallowRef<InstanceType<typeof Popup>>()

const popupTitle = computed(() => {
    return '预览定时任务'
})

const formData = reactive({
    id: null,
    task_name: null,
    task_code: null,
    corn_expr: null,
    status: null
})

const formRules = {
    id: [
        {
            required: true,
            message: '请输入',
            trigger: ['blur']
        }
    ],
    task_name: [
        {
            required: true,
            message: '请输入任务名称',
            trigger: ['blur']
        }
    ],
    task_code: [
        {
            required: true,
            message: '请输入任务编码',
            trigger: ['blur']
        }
    ],
    corn_expr: [
        {
            required: true,
            message: '请输入corn表达式',
            trigger: ['blur']
        }
    ],
    status: [
        {
            required: true,
            message: '请输入禁用',
            trigger: ['blur']
        }
    ]
}

const open = () => {
    popupRef.value?.open()
}
const getDetail = async (row: Record<string, any>) => {
    try {
        const data = await system_corn_detail(row.id)
        setFormData(data)
    } catch (error) {
        console.error('定时任务详情获取失败:', error)
    }
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
