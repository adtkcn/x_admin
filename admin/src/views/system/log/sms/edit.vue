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
                <el-form-item label="场景编号" prop="Scene">
                    <el-input v-model="formData.Scene" type="number" placeholder="请输入场景编号" />
                </el-form-item>
                <el-form-item label="手机号码" prop="Mobile">
                    <el-input v-model="formData.Mobile" placeholder="请输入手机号码" />
                </el-form-item>
                <el-form-item label="发送内容" prop="Content">
                    <editor v-model="formData.Content" :height="500" />
                </el-form-item>
                <el-form-item label="发送状态：[0=发送中, 1=发送成功, 2=发送失败]" prop="Status">
                    <el-radio-group
                        v-model="formData.Status"
                        placeholder="请选择发送状态：[0=发送中, 1=发送成功, 2=发送失败]"
                    >
                        <el-radio label="0">请选择字典生成</el-radio>
                    </el-radio-group>
                </el-form-item>
                <el-form-item label="短信结果" prop="Results">
                    <el-input
                        v-model="formData.Results"
                        placeholder="请输入短信结果"
                        type="textarea"
                        :autosize="{ minRows: 4, maxRows: 6 }"
                    />
                </el-form-item>
                <el-form-item label="发送时间" prop="SendTime">
                    <el-date-picker
                        class="flex-1 !flex"
                        v-model="formData.SendTime"
                        type="datetime"
                        clearable
                        value-format="YYYY-MM-DD hh:mm:ss"
                        placeholder="请选择发送时间"
                    />
                </el-form-item>
            </el-form>
        </popup>
    </div>
</template>
<script lang="ts" setup>
import type { FormInstance } from 'element-plus'
import {
    system_log_sms_edit,
    system_log_sms_add,
    system_log_sms_detail
} from '@/api/system_log_sms'
import Popup from '@/components/popup/index.vue'
import feedback from '@/utils/feedback'
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
const emit = defineEmits(['success', 'close'])
const formRef = shallowRef<FormInstance>()
const popupRef = shallowRef<InstanceType<typeof Popup>>()
const mode = ref('add')
const popupTitle = computed(() => {
    return mode.value == 'edit' ? '编辑系统短信日志' : '新增系统短信日志'
})

const formData = reactive({
    Id: null,
    Scene: null,
    Mobile: null,
    Content: null,
    Status: null,
    Results: null,
    SendTime: null
})

const formRules = {
    Id: [
        {
            required: true,
            message: '请输入id',
            trigger: ['blur']
        }
    ],
    Scene: [
        {
            required: true,
            message: '请输入场景编号',
            trigger: ['blur']
        }
    ],
    Mobile: [
        {
            required: true,
            message: '请输入手机号码',
            trigger: ['blur']
        }
    ],
    Content: [
        {
            required: true,
            message: '请输入发送内容',
            trigger: ['blur']
        }
    ],
    Status: [
        {
            required: true,
            message: '请选择发送状态：[0=发送中, 1=发送成功, 2=发送失败]',
            trigger: ['blur']
        }
    ],
    Results: [
        {
            required: true,
            message: '请输入短信结果',
            trigger: ['blur']
        }
    ],
    SendTime: [
        {
            required: true,
            message: '请选择发送时间',
            trigger: ['blur']
        }
    ]
}

const handleSubmit = async () => {
    try {
        await formRef.value?.validate()
        const data: any = { ...formData }
        mode.value == 'edit' ? await system_log_sms_edit(data) : await system_log_sms_add(data)
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
        const data = await system_log_sms_detail(row.Id)
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
