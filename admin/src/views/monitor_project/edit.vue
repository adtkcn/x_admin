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
                <!-- <el-form-item label="项目Key" prop="projectKey">
                    <el-input v-model="formData.projectKey" placeholder="请输入项目uuid" />
                </el-form-item> -->
                <el-form-item label="项目名称" prop="projectName">
                    <el-input v-model="formData.projectName" placeholder="请输入项目名称" />
                </el-form-item>
                <el-form-item label="项目类型" prop="projectType">
                    <el-select
                        class="flex-1"
                        v-model="formData.projectType"
                        placeholder="请选择项目类型"
                    >
                        <el-option
                            v-for="(item, index) in dictData.project_type"
                            :key="index"
                            :label="item.name"
                            :value="item.value"
                            clearable
                            :disabled="!item.status"
                        />
                    </el-select>
                </el-form-item>
            </el-form>
        </popup>
    </div>
</template>
<script lang="ts" setup>
import type { FormInstance } from 'element-plus'
import {
    monitor_project_edit,
    monitor_project_add,
    monitor_project_detail
} from '@/api/monitor_project'
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
    return mode.value == 'edit' ? '编辑错误项目' : '新增错误项目'
})

const formData = reactive({
    id: '',
    // projectKey: '',
    projectName: null,
    projectType: ''
})

const formRules = {
    id: [
        {
            required: true,
            message: '请输入项目id',
            trigger: ['blur']
        }
    ],
    // projectKey: [
    //     {
    //         required: true,
    //         message: '请输入项目Key',
    //         trigger: ['blur']
    //     }
    // ],
    projectName: [
        {
            required: true,
            message: '请输入项目名称',
            trigger: ['blur']
        }
    ],
    projectType: [
        {
            required: true,
            message: '请选择项目类型go java web node php 等',
            trigger: ['blur']
        }
    ]
}

const handleSubmit = async () => {
    try {
        await formRef.value?.validate()
        const data: any = { ...formData }
        mode.value == 'edit' ? await monitor_project_edit(data) : await monitor_project_add(data)
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
        const data = await monitor_project_detail({
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
