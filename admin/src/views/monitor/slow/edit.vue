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
                        <el-form-item label="项目key" prop="ProjectKey">
                            <el-input v-model="formData.ProjectKey" placeholder="请输入项目key" />
                        </el-form-item>
                        <el-form-item label="sdk生成的客户端id" prop="ClientId">
                            <el-input v-model="formData.ClientId" placeholder="请输入sdk生成的客户端id" />
                        </el-form-item>
                        <el-form-item label="用户id" prop="UserId">
                            <el-input v-model="formData.UserId" placeholder="请输入用户id" />
                        </el-form-item>
                        <el-form-item label="URL地址" prop="Path">
                            <el-input
                                v-model="formData.Path"
                                placeholder="请输入URL地址"
                                type="textarea"
                                :autosize="{ minRows: 4, maxRows: 6 }"
                            />
                        </el-form-item>
                        <el-form-item label="时间" prop="Time">
                            <el-input v-model="formData.Time" type="number" placeholder="请输入时间" />
                        </el-form-item>
            
            </el-form>
        </popup>
    </div>
</template>
<script lang="ts" setup>
import type { FormInstance } from 'element-plus'
import {  monitor_slow_edit, monitor_slow_add, monitor_slow_detail } from '@/api/monitor/slow'
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
    return mode.value == 'edit' ? '编辑监控-错误列' : '新增监控-错误列'
})

const formData = reactive({
    Id: null,
    ProjectKey: null,
    ClientId: null,
    UserId: null,
    Path: null,
    Time: null,
})

const formRules = {
    Id: [
        {
            required: true,
            message: '请输入错误id',
            trigger: ['blur']
        }
    ],
    ProjectKey: [
        {
            required: true,
            message: '请输入项目key',
            trigger: ['blur']
        }
    ],
    ClientId: [
        {
            required: true,
            message: '请输入sdk生成的客户端id',
            trigger: ['blur']
        }
    ],
    UserId: [
        {
            required: true,
            message: '请输入用户id',
            trigger: ['blur']
        }
    ],
    Path: [
        {
            required: true,
            message: '请输入URL地址',
            trigger: ['blur']
        }
    ],
    Time: [
        {
            required: true,
            message: '请输入时间',
            trigger: ['blur']
        }
    ],
}

const handleSubmit = async () => {
     try {
        await formRef.value?.validate()
        const data: any = { ...formData }
        mode.value == 'edit' ? await monitor_slow_edit(data) : await monitor_slow_add(data)
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
        const data = await monitor_slow_detail(row.Id)
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
