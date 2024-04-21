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
                        <el-form-item label="项目key" prop="projectKey">
                            <el-input v-model="formData.projectKey" placeholder="请输入项目key" />
                        </el-form-item>
                        <el-form-item label="sdk生成的客户端id" prop="clientId">
                            <el-input v-model="formData.clientId" placeholder="请输入sdk生成的客户端id" />
                        </el-form-item>
                        <el-form-item label="事件类型" prop="eventType">
                            <el-select class="flex-1" v-model="formData.eventType" placeholder="请选择事件类型">
                                <el-option label="请选择字典生成" value="" />
                            </el-select>
                        </el-form-item>
                        <el-form-item label="URL地址" prop="page">
                            <el-input
                                v-model="formData.page"
                                placeholder="请输入URL地址"
                                type="textarea"
                                :autosize="{ minRows: 4, maxRows: 6 }"
                            />
                        </el-form-item>
                        <el-form-item label="错误消息" prop="message">
                            <el-input
                                v-model="formData.message"
                                placeholder="请输入错误消息"
                                type="textarea"
                                :autosize="{ minRows: 4, maxRows: 6 }"
                            />
                        </el-form-item>
                        <el-form-item label="错误堆栈" prop="stack">
                            <el-input
                                v-model="formData.stack"
                                placeholder="请输入错误堆栈"
                                type="textarea"
                                :autosize="{ minRows: 4, maxRows: 6 }"
                            />
                        </el-form-item>
                        <el-form-item label="客户端时间" prop="clientTime">
                            <el-input v-model="formData.clientTime" placeholder="请输入客户端时间" />
                        </el-form-item>
            
            </el-form>
        </popup>
    </div>
</template>
<script lang="ts" setup>
import type { FormInstance } from 'element-plus'
import {  monitor_web_edit, monitor_web_add, monitor_web_detail } from '@/api/monitor_web'
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
    return mode.value == 'edit' ? '编辑错误收集error' : '新增错误收集error'
})

const formData = reactive({
    id: '',
    projectKey: '',
    clientId: '',
    eventType: '',
    page: '',
    message: '',
    stack: '',
    clientTime: '',
})

const formRules = {
    id: [
        {
            required: true,
            message: '请输入uuid',
            trigger: ['blur']
        }
    ],
    projectKey: [
        {
            required: true,
            message: '请输入项目key',
            trigger: ['blur']
        }
    ],
    clientId: [
        {
            required: true,
            message: '请输入sdk生成的客户端id',
            trigger: ['blur']
        }
    ],
    eventType: [
        {
            required: true,
            message: '请选择事件类型',
            trigger: ['blur']
        }
    ],
    page: [
        {
            required: true,
            message: '请输入URL地址',
            trigger: ['blur']
        }
    ],
    message: [
        {
            required: true,
            message: '请输入错误消息',
            trigger: ['blur']
        }
    ],
    stack: [
        {
            required: true,
            message: '请输入错误堆栈',
            trigger: ['blur']
        }
    ],
    clientTime: [
        {
            required: true,
            message: '请输入客户端时间',
            trigger: ['blur']
        }
    ],
}

const handleSubmit = async () => {
     try {
        await formRef.value?.validate()
        const data: any = { ...formData }
        mode.value == 'edit' ? await monitor_web_edit(data) : await monitor_web_add(data)
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
        const data = await monitor_web_detail({
            id: row.id
        })
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
