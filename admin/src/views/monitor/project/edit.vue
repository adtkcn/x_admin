<template>
    <div class="edit-popup">
        <popup
            ref="popupRef"
            :title="popupTitle"
            :async="true"
            width="700px"
            :clickModalClose="true"
            @confirm="handleSubmit"
            @close="handleClose"
        >
            <el-form ref="formRef" :model="formData" label-width="110px" :rules="formRules">
                <el-form-item label="项目uuid" prop="project_key" v-if="mode === 'edit'">
                    {{ formData.project_key }}
                </el-form-item>
                <el-form-item label="项目名称" prop="project_name">
                    <el-input v-model="formData.project_name" placeholder="请输入项目名称" />
                </el-form-item>
                <el-form-item label="项目类型" prop="project_type">
                    <el-select
                        class="flex-1"
                        v-model="formData.project_type"
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
                <el-form-item label="是否启用" prop="status">
                    <el-select
                        class="flex-1"
                        v-model="formData.status"
                        placeholder="请选择是否启用"
                    >
                        <el-option
                            v-for="(item, index) in dictData.status"
                            :key="index"
                            :label="item.name"
                            :value="parseInt(item.value)"
                            clearable
                            :disabled="!item.status"
                        />
                    </el-select>
                </el-form-item>
                <el-form-item
                    label="使用SDK"
                    v-if="mode == 'edit' && formData.project_type == 'web'"
                >
                    <highlight-code :code="code" lang="javascript"></highlight-code>
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
} from '@/api/monitor/project'
import Popup from '@/components/popup/index.vue'
import feedback from '@/utils/feedback'
import { computed, ref, shallowRef } from 'vue'
import type { PropType } from 'vue'
import { useReactiveWithReset } from '@/hooks/useReactiveWithReset'
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
    return mode.value == 'edit' ? '编辑监控项目' : '新增监控项目'
})

const {
    state: formData,
    setState
} = useReactiveWithReset({
    id: null,
    project_key: null,
    project_name: null,
    project_type: null,
    status: null
})
const code = computed(() => {
    return `import { Base, Web } from '../../x_err_sdk/web/index'
new Base(
    {
        Dns: '${location.origin}/api',
        Pid: '${formData.project_key}',
        Uid: ''
    },
    new Web()
)`
})
const formRules = {
    id: [
        {
            required: true,
            message: '请输入项目id',
            trigger: ['blur']
        }
    ],
    project_key: [
        {
            required: true,
            message: '请输入项目uuid',
            trigger: ['blur']
        }
    ],
    project_name: [
        {
            required: true,
            message: '请输入项目名称',
            trigger: ['blur']
        }
    ],
    project_type: [
        {
            required: true,
            message: '请选择项目类型',
            trigger: ['blur']
        }
    ],
    status: [
        {
            required: true,
            message: '请选择是否启用',
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
    } catch (error) {
        console.error('监控项目保存失败:', error)
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
        const data = await monitor_project_detail(row.id)
        setFormData(data)
    } catch (error) {
        console.error('监控项目详情获取失败:', error)
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
