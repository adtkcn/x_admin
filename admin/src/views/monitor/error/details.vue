<template>
    <div class="edit-popup">
        <popup
            ref="popupRef"
            :title="popupTitle"
            :async="true"
            width="90%"
            :clickModalClose="true"
            @confirm="handleSubmit"
            @close="handleClose"
        >
            <el-row :gutter="20">
                <el-col :span="12">
                    <el-card>
                        <template #header>
                            <div class="card-header">
                                <span>错误详情</span>
                            </div>
                        </template>

                        <el-form ref="formRef">
                            <el-form-item label="项目key" prop="ProjectKey">
                                <dict-value
                                    :options="listAllData.monitor_project_listAll"
                                    :value="formData.ProjectKey"
                                    labelKey="ProjectName"
                                    valueKey="ProjectKey"
                                />
                            </el-form-item>

                            <el-form-item label="事件类型" prop="EventType">
                                {{ formData.EventType }}
                            </el-form-item>
                            <el-form-item label="URL地址" prop="Path">
                                {{ formData.Path }}
                            </el-form-item>
                            <el-form-item label="错误消息" prop="Message">
                                {{ formData.Message }}
                            </el-form-item>
                            <el-form-item label="错误堆栈" prop="Stack">
                                {{ formData.Stack }}
                            </el-form-item>
                        </el-form>
                        <!-- <template #footer>Footer content</template> -->
                    </el-card>
                </el-col>
                <el-col :span="12">
                    <el-card>
                        <template #header>
                            <div class="card-header">
                                <span>用户列表</span>
                            </div>
                        </template>
                        <el-scrollbar height="400px">
                            <el-collapse v-model="activeNames">
                                <el-collapse-item
                                    v-for="(user, index) in users"
                                    :key="user.Id"
                                    :title="user.CreateTime"
                                    :name="index"
                                >
                                    <template #title>
                                        <div class="flex-1 text-left">
                                            {{ user.UserId }}
                                        </div>
                                        <span class="">
                                            {{ user.CreateTime }}
                                        </span>
                                    </template>
                                    <el-descriptions border>
                                        <el-descriptions-item label="城市：">{{
                                            user.City
                                        }}</el-descriptions-item>

                                        <el-descriptions-item label="OS：">{{
                                            user.Os
                                        }}</el-descriptions-item>
                                        <el-descriptions-item label="屏幕尺寸：">
                                            <el-tag size="small"
                                                >{{ user.Width }}*{{ user.Height }}</el-tag
                                            >
                                        </el-descriptions-item>
                                        <el-descriptions-item label="userAgent">
                                            {{ user.Ua }}
                                        </el-descriptions-item>
                                    </el-descriptions>
                                </el-collapse-item>
                            </el-collapse>
                        </el-scrollbar>
                    </el-card>
                </el-col>
            </el-row>
        </popup>
    </div>
</template>
<script lang="ts" setup>
import type { FormInstance } from 'element-plus'
import { monitor_error_add, monitor_error_detail } from '@/api/monitor/error'
import { monitor_client_errorUsers } from '@/api/monitor/client'
import type { type_monitor_client } from '@/api/monitor/client'

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
    return '监控详情'
})

const activeNames = ref<string[]>(['1'])

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
const users = ref<type_monitor_client[]>([])

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

        users.value = await monitor_client_errorUsers(row.Id)
        console.log('user', users.value)

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
