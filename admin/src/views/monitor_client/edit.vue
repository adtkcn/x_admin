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
                <el-form-item label="客户端id" prop="clientId">
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
                    <el-input v-model.number="formData.width" placeholder="请输入屏幕" />
                </el-form-item>
                <el-form-item label="屏幕高度" prop="height">
                    <el-input v-model.number="formData.height" placeholder="请输入屏幕高度" />
                </el-form-item>
                <el-form-item label="ua记录" prop="ua">
                    <el-input v-model="formData.ua" placeholder="请输入ua记录" />
                </el-form-item>
                <el-form-item label="时间" prop="clientTime">
                    <el-input v-model="formData.clientTime" placeholder="请输入客户端时间" />
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
} from '@/api/monitor_client'
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
    return mode.value == 'edit' ? '编辑客户端信息' : '新增客户端信息'
})

const formData = reactive({
    id: '',
    projectKey: '',
    clientId: '',
    userId: '',
    os: '',
    browser: '',
    city: '',
    width: '',
    height: '',
    ua: '',
    clientTime: ''
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
            // required: true,
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
            message: '请输入客户端时间',
            trigger: ['blur']
        }
    ]
}

const handleSubmit = async () => {
    formRef.value
        ?.validate()
        .then(() => {
            try {
                const data: any = { ...formData }
                let req = null
                if (mode.value == 'edit') {
                    req = monitor_client_edit(data)
                } else {
                    req = monitor_client_add(data)
                }
                req.then(() => {
                    popupRef.value?.close()
                    feedback.msgSuccess('操作成功')
                    emit('success')
                }).catch((err) => {
                    feedback.msgError(err.msg)
                })
            } catch (error) {}
        })
        .catch(() => {})
}

const open = (type = 'add') => {
    mode.value = type
    popupRef.value?.open()
}

const setFormData = async (data: Record<string, any>) => {
    for (const key in formData) {
        if (data[key] != null && data[key] != undefined) {
            formData[key] = data[key]
        }
    }
}

const getDetail = async (row: Record<string, any>) => {
    try {
        const data = await monitor_client_detail({
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
