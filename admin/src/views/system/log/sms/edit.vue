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
                <el-form-item label="场景编号" prop="scene">
                    <el-input v-model="formData.scene" type="number" placeholder="请输入场景编号" />
                </el-form-item>
                <el-form-item label="手机号码" prop="mobile">
                    <el-input v-model="formData.mobile" placeholder="请输入手机号码" />
                </el-form-item>
                <el-form-item label="发送内容" prop="content">
                    <editor v-model="formData.content" :height="500" />
                </el-form-item>
                <el-form-item label="发送状态：[0=发送中, 1=发送成功, 2=发送失败]" prop="status">
                    <el-select
                        class="flex-1"
                        v-model="formData.status"
                        placeholder="请选择发送状态：[0=发送中, 1=发送成功, 2=发送失败]"
                    >
                        <el-option
                            v-for="(item, index) in dictData.flow_apply_status"
                            :key="index"
                            :label="item.name"
                            :value="parseInt(item.value)"
                            clearable
                            :disabled="!item.status"
                        />
                    </el-select>
                </el-form-item>
                <el-form-item label="短信结果" prop="results">
                    <el-input
                        v-model="formData.results"
                        placeholder="请输入短信结果"
                        type="textarea"
                        :autosize="{ minRows: 4, maxRows: 6 }"
                    />
                </el-form-item>
                <el-form-item label="发送时间" prop="send_time">
                    <el-input
                        v-model.number="formData.send_time"
                        type="number"
                        placeholder="请输入发送时间"
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
    id: '',
    scene: 0,
    mobile: '',
    content: '',
    status: '',
    results: '',
    send_time: 0
})

const formRules = {
    id: [
        {
            required: true,
            message: '请输入id',
            trigger: ['blur']
        }
    ]
    // scene: [
    //     {
    //         required: true,
    //         message: '请输入场景编号',
    //         trigger: ['blur']
    //     }
    // ],
    // mobile: [
    //     {
    //         required: true,
    //         message: '请输入手机号码',
    //         trigger: ['blur']
    //     }
    // ],
    // content: [
    //     {
    //         required: true,
    //         message: '请输入发送内容',
    //         trigger: ['blur']
    //     }
    // ],
    // status: [
    //     {
    //         required: true,
    //         message: '请选择发送状态：[0=发送中, 1=发送成功, 2=发送失败]',
    //         trigger: ['blur']
    //     }
    // ],
    // results: [
    //     {
    //         required: true,
    //         message: '请输入短信结果',
    //         trigger: ['blur']
    //     }
    // ],
    // send_time: [
    //     {
    //         required: true,
    //         message: '请输入发送时间',
    //         trigger: ['blur']
    //     }
    // ]
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
        const data = await system_log_sms_detail(row.id)
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
