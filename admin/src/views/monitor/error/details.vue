<template>
    <div class="edit-popup">
        <popup
            ref="popupRef"
            :title="popupTitle"
            :async="true"
            width="98%"
            :clickModalClose="true"
            @confirm="handleSubmit"
            @close="handleClose"
        >
            <el-row :gutter="10">
                <el-col :xs="24" :md="14">
                    <el-card>
                        <template #header>
                            <div class="card-header">
                                <span
                                    >【
                                    <dict-value
                                        :options="listAllData.monitor_project_listAll"
                                        :value="formData.project_key"
                                        labelKey="project_name"
                                        valueKey="project_key"
                                    />】错误详情</span
                                >
                            </div>
                        </template>

                        <el-form ref="formRef">
                            <!-- <el-form-item label="项目key" prop="project_key"> </el-form-item> -->

                            <el-form-item label="事件类型：" prop="event_type">
                                {{ formData.event_type }}
                            </el-form-item>
                            <el-form-item label="URL地址：" prop="path">
                                {{ formData.path }}
                            </el-form-item>
                            <el-form-item label="错误消息：" prop="message">
                                {{ formData.message }}
                            </el-form-item>
                            <el-form-item label="" prop="stack">
                                {{ formData.stack }}
                            </el-form-item>
                        </el-form>
                    </el-card>
                </el-col>
                <el-col :xs="24" :md="10">
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
                                    :key="user.id"
                                    :title="user.create_time"
                                    :name="index"
                                >
                                    <template #title>
                                        <div class="collapse-title">
                                            <span>
                                                {{ user.city }} {{ user.browser }}：{{ user.ip }}
                                            </span>

                                            <span>
                                                {{ user.create_time }}
                                            </span>
                                        </div>
                                    </template>
                                    <el-descriptions border :column="2">
                                        <el-descriptions-item label="省市区">
                                            {{ user.country }}{{ user.province }}{{ user.city }}
                                        </el-descriptions-item>
                                        <el-descriptions-item label="浏览器">
                                            {{ user.os }}/{{ user.browser }}
                                        </el-descriptions-item>
                                        <el-descriptions-item label="网络">{{
                                            user.operator
                                        }}</el-descriptions-item>
                                        <el-descriptions-item label="IP">{{
                                            user.ip
                                        }}</el-descriptions-item>

                                        <el-descriptions-item label="业务ID">{{
                                            user.user_id
                                        }}</el-descriptions-item>
                                        <el-descriptions-item label="屏幕">
                                            {{ user.width }}*{{ user.height }}
                                        </el-descriptions-item>
                                        <el-descriptions-item label="userAgent">
                                            {{ user.ua }}
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
import { computed, ref, reactive, shallowRef } from 'vue'
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
    id: null,
    project_key: null,
    client_id: null,
    event_type: null,
    path: null,
    message: null,
    stack: null,
    md5: null,
    client_time: null
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
    } catch (error) {
        console.error('监控错误保存失败:', error)
    }
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
        const data = await monitor_error_detail(row.id)

        users.value = await monitor_client_errorUsers(row.id)
        console.error('监控错误用户', users.value)

        setFormData(data)
    } catch (error) {
        console.error('监控错误详情获取失败:', error)
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
<style lang="scss">
.collapse-title {
    display: flex;
    justify-content: space-between;
}
</style>
