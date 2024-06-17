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
                        <el-form-item label="sdk生成的客户端id" prop="clientId">
                            <el-input v-model="formData.clientId" placeholder="请输入sdk生成的客户端id" />
                        </el-form-item>
                        <el-form-item label="用户id" prop="userId">
                            <el-input v-model="formData.userId" placeholder="请输入用户id" />
                        </el-form-item>
                        <el-form-item label="系统" prop="os">
                            <el-input v-model="formData.os" placeholder="请输入系统" />
                        </el-form-item>
                        <el-form-item label="浏览器" prop="browser">
                            <el-input v-model="formData.browser" placeholder="请输入浏览器" />
                        </el-form-item>
                        <el-form-item label="城市" prop="city">
                            <el-input v-model="formData.city" placeholder="请输入城市" />
                        </el-form-item>
                        <el-form-item label="屏幕" prop="width">
                            <el-input v-model="formData.width" placeholder="请输入屏幕" />
                        </el-form-item>
                        <el-form-item label="屏幕高度" prop="height">
                            <el-input v-model="formData.height" placeholder="请输入屏幕高度" />
                        </el-form-item>
                        <el-form-item label="ua记录" prop="ua">
                            <el-input v-model="formData.ua" placeholder="请输入ua记录" />
                        </el-form-item>
                        <el-form-item label="更新时间" prop="clientTime">
                            <el-date-picker
                                class="flex-1 !flex"
                                v-model="formData.clientTime"
                                type="datetime"
                                clearable
                                value-format="YYYY-MM-DD hh:mm:ss"
                                placeholder="请选择更新时间"
                            />
                        </el-form-item>
            
            </el-form>
        </popup>
    </div>
</template>
<script lang="ts" setup>
import type { FormInstance } from 'element-plus'
import {  monitor_client_edit, monitor_client_add, monitor_client_detail } from '@/api/monitor_client'
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
    return mode.value == 'edit' ? '编辑监控-客户端信息' : '新增监控-客户端信息'
})

const formData = reactive({
    id: '',
    clientId: '',
    userId: '',
    os: '',
    browser: '',
    city: '',
    width: '',
    height: '',
    ua: '',
    clientTime: '',
})

const formRules = {
    clientId: [
        {
            required: true,
            message: '请输入sdk生成的客户端id',
            trigger: ['blur']
        }
    ],
    userId: [
        {
            required: true,
            message: '请输入用户id',
            trigger: ['blur']
        }
    ],
    os: [
        {
            required: true,
            message: '请输入系统',
            trigger: ['blur']
        }
    ],
    browser: [
        {
            required: true,
            message: '请输入浏览器',
            trigger: ['blur']
        }
    ],
    city: [
        {
            required: true,
            message: '请输入城市',
            trigger: ['blur']
        }
    ],
    width: [
        {
            required: true,
            message: '请输入屏幕',
            trigger: ['blur']
        }
    ],
    height: [
        {
            required: true,
            message: '请输入屏幕高度',
            trigger: ['blur']
        }
    ],
    ua: [
        {
            required: true,
            message: '请输入ua记录',
            trigger: ['blur']
        }
    ],
    clientTime: [
        {
            required: true,
            message: '请选择更新时间',
            trigger: ['blur']
        }
    ],
}

const handleSubmit = async () => {
     try {
        await formRef.value?.validate()
        const data: any = { ...formData }
        mode.value == 'edit' ? await monitor_client_edit(data) : await monitor_client_add(data)
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
        const data = await monitor_client_detail(row.id)
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
