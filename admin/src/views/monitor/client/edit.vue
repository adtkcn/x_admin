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
                <el-form-item label="客户端id" prop="ClientId">
                    <el-input v-model="formData.ClientId" placeholder="请输入客户端id" />
                </el-form-item>
                <el-form-item label="用户id" prop="UserId">
                    <el-input v-model="formData.UserId" placeholder="请输入用户id" />
                </el-form-item>
                <el-form-item label="系统" prop="Os">
                    <el-input v-model="formData.Os" placeholder="请输入系统" />
                </el-form-item>
                <el-form-item label="浏览器" prop="Browser">
                    <el-input v-model="formData.Browser" placeholder="请输入浏览器" />
                </el-form-item>
                <el-form-item label="城市" prop="City">
                    <el-input v-model="formData.City" placeholder="请输入城市" />
                </el-form-item>
                <el-form-item label="屏幕" prop="Width">
                    <el-input v-model="formData.Width" type="number" placeholder="请输入屏幕" />
                </el-form-item>
                <el-form-item label="屏幕高度" prop="Height">
                    <el-input
                        v-model="formData.Height"
                        type="number"
                        placeholder="请输入屏幕高度"
                    />
                </el-form-item>
                <el-form-item label="ua记录" prop="Ua">
                    <el-input v-model="formData.Ua" placeholder="请输入ua记录" />
                </el-form-item>
            </el-form>
        </popup>
    </div>
</template>
<script lang="ts" setup>
import type { FormInstance } from 'element-plus'
import {
    monitor_client_edit,
    monitor_client_add,
    monitor_client_detail
} from '@/api/monitor/client'
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
    return mode.value == 'edit' ? '编辑监控-客户端信息' : '新增监控-客户端信息'
})

const formData = reactive({
    Id: null,
    ProjectKey: null,
    ClientId: null,
    UserId: null,
    Os: null,
    Browser: null,
    City: null,
    Width: null,
    Height: null,
    Ua: null,
    ClientTime: null
})

const formRules = {
    Id: [
        {
            required: true,
            message: '请输入uuid',
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
            message: '请输入客户端id',
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
    Os: [
        {
            required: true,
            message: '请输入系统',
            trigger: ['blur']
        }
    ],
    Browser: [
        {
            required: true,
            message: '请输入浏览器',
            trigger: ['blur']
        }
    ],
    City: [
        {
            required: true,
            message: '请输入城市',
            trigger: ['blur']
        }
    ],
    Width: [
        {
            required: true,
            message: '请输入屏幕',
            trigger: ['blur']
        }
    ],
    Height: [
        {
            required: true,
            message: '请输入屏幕高度',
            trigger: ['blur']
        }
    ],
    Ua: [
        {
            required: true,
            message: '请输入ua记录',
            trigger: ['blur']
        }
    ],
    ClientTime: [
        {
            required: true,
            message: '请选择更新时间',
            trigger: ['blur']
        }
    ]
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
        const data = await monitor_client_detail(row.Id)
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
