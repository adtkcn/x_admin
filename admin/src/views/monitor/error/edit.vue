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
                    <el-select
                        class="flex-1"
                        v-model="formData.ProjectKey"
                        placeholder="请选择项目key"
                    >
                        <el-option
                            v-for="(item, index) in listAllData.monitor_project_listAll"
                            :key="index"
                            :label="item.ProjectName"
                            :value="String(item.ProjectKey)"
                            clearable
                        />
                    </el-select>
                </el-form-item>
                <el-form-item label="客户端id" prop="ClientId">
                    <el-input v-model="formData.ClientId" placeholder="请输入客户端id" />
                </el-form-item>

                <el-form-item label="事件类型" prop="EventType">
                    <el-input
                        v-model="formData.EventType"
                        placeholder="请输入事件类型"
                        type="textarea"
                        :autosize="{ minRows: 4, maxRows: 6 }"
                    />
                </el-form-item>
                <el-form-item label="URL地址" prop="Path">
                    <el-input
                        v-model="formData.Path"
                        placeholder="请输入URL地址"
                        type="textarea"
                        :autosize="{ minRows: 4, maxRows: 6 }"
                    />
                </el-form-item>
                <el-form-item label="错误消息" prop="Message">
                    <el-input
                        v-model="formData.Message"
                        placeholder="请输入错误消息"
                        type="textarea"
                        :autosize="{ minRows: 4, maxRows: 6 }"
                    />
                </el-form-item>
                <el-form-item label="错误堆栈" prop="Stack">
                    <el-input
                        v-model="formData.Stack"
                        placeholder="请输入错误堆栈"
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
import { monitor_error_add, monitor_error_detail } from '@/api/monitor/error'
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
    return mode.value == 'edit' ? '编辑监控-错误列' : '新增监控-错误列'
})

const formData = reactive({
    Id: null,
    ProjectKey: null,
    ClientId: null,
    EventType: null,
    Path: null,
    Message: null,
    Stack: null,
    Md5: null,
    ClientTime: null
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
            message: '请选择项目key',
            trigger: ['blur']
        }
    ]
    // EventType: [
    //     {
    //         required: true,
    //         message: '请输入事件类型',
    //         trigger: ['blur']
    //     }
    // ],
    // Path: [
    //     {
    //         required: true,
    //         message: '请输入URL地址',
    //         trigger: ['blur']
    //     }
    // ],
    // Message: [
    //     {
    //         required: true,
    //         message: '请输入错误消息',
    //         trigger: ['blur']
    //     }
    // ],
    // Stack: [
    //     {
    //         required: true,
    //         message: '请输入错误堆栈',
    //         trigger: ['blur']
    //     }
    // ]
}

const handleSubmit = async () => {
    try {
        await formRef.value?.validate()
        const data: any = { ...formData }
        await monitor_error_add(data)
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
        const data = await monitor_error_detail(row.Id)
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
