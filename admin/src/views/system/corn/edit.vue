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
            <el-form ref="formRef" :model="formData" label-width="110px" :rules="formRules">
                <el-form-item label="任务名称" prop="task_name">
                    <el-input v-model="formData.task_name" placeholder="请输入任务名称" />
                </el-form-item>
                <el-form-item label="任务编码" prop="task_code">
                    <!-- <el-input v-model="formData.task_code" placeholder="请输入任务编码" /> -->

                    <el-select v-model="formData.task_code" placeholder="请选择任务编码">
                        <el-option
                            v-for="item in taskList"
                            :key="item.task_code"
                            :label="item.task_code + ' - ' + item.task_desc"
                            :value="item.task_code"
                        />
                    </el-select>
                </el-form-item>
                <el-form-item label="corn表达式" prop="corn_expr">
                    <el-autocomplete
                        v-model="formData.corn_expr"
                        :fetch-suggestions="queryCornSuggest"
                        placeholder="请输入corn表达式,秒级示例*/5 * * * * *"
                        clearable
                        style="width: 100%"
                    >
                        <template #default="{ item }">
                            <div class="corn-option">
                                <span class="corn-value">{{ item.value }}</span>
                                <span class="corn-label">{{ item.desc }}</span>
                            </div>
                        </template>
                    </el-autocomplete>
                </el-form-item>
                <el-form-item label="状态" prop="status">
                    <!-- <el-input v-model="formData.status" type="number" placeholder="请输入禁用" /> -->

                    <el-radio-group v-model="formData.status">
                        <el-radio :value="1">正常</el-radio>
                        <el-radio :value="0">停用</el-radio>
                    </el-radio-group>
                </el-form-item>
            </el-form>
        </popup>
    </div>
</template>
<script lang="ts" setup>
import type { FormInstance } from 'element-plus'
import {
    system_corn_edit,
    system_corn_add,
    system_corn_detail,
    system_corn_getTaskList,
    type type_system_corn_edit
} from '@/api/system/corn'
import type { type_task } from '@/api/system/corn'
import Popup from '@/components/popup/index.vue'
import feedback from '@/utils/feedback'
import { ref, shallowRef, computed } from 'vue'
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
    return mode.value == 'edit' ? '编辑定时任务' : '新增定时任务'
})

const { state: formData, setState } = useReactiveWithReset<type_system_corn_edit>({
    id: '',
    task_name: '',
    task_code: '',
    corn_expr: '',
    status: 1
})

const formRules = {
    id: [
        {
            required: true,
            message: '请输入',
            trigger: ['blur']
        }
    ],
    task_name: [
        {
            required: true,
            message: '请输入任务名称',
            trigger: ['blur']
        }
    ],
    task_code: [
        {
            required: true,
            message: '请输入任务编码',
            trigger: ['blur']
        }
    ],
    corn_expr: [
        {
            required: true,
            message: '请输入corn表达式',
            trigger: ['blur']
        }
    ],
    status: [
        {
            required: true,
            message: '请输入禁用',
            trigger: ['blur']
        }
    ]
}

const handleSubmit = async () => {
    try {
        await formRef.value?.validate()
        const data = { ...formData }
        if (mode.value == 'edit') {
            await system_corn_edit(data)
        } else {
            await system_corn_add(data)
        }
        popupRef.value?.close()
        feedback.msgSuccess('操作成功')
        emit('success')
    } catch (error) {
        console.error('定时任务保存失败:', error)
    }
}

const open = (type = 'add') => {
    mode.value = type
    popupRef.value?.open()
    getTaskList()
}

const setFormData = async (data: Record<string, any>) => {
    setState(data)
}

const getDetail = async (row: Record<string, any>) => {
    try {
        const data = await system_corn_detail(row.id)
        setFormData(data)
    } catch (error) {
        console.error('定时任务详情获取失败:', error)
    }
}

const taskList = ref<type_task[]>([])
const getTaskList = async () => {
    try {
        const data = await system_corn_getTaskList()
        taskList.value = data
    } catch (error) {
        console.error('定时任务可选列表获取失败:', error)
    }
}

const cornOptions: { value: string; desc: string }[] = [
    { value: '*/5 * * * * *', desc: '每5秒执行一次' },
    { value: '0 * * * * *', desc: '每分钟的第0秒执行' },
    { value: '0 0 * * * *', desc: '每小时整点执行' },
    { value: '0 0 0 * * *', desc: '每天0点执行' },
    { value: '0 0 1 * * *', desc: '每天1点执行' },
    { value: '0 0 0 * * 1', desc: '每周一零点执行' },
    { value: '0 0 0 1 * *', desc: '每月1号零点执行' },
    { value: '0 0/30 * * * *', desc: '每30分钟执行一次' },
    { value: '0 0 0/1 * * *', desc: '每1小时执行一次' },
    { value: '0 0 1 * * 0', desc: '每周日1点执行' },
    { value: '0 0 0 L * *', desc: '每月最后一天零点执行' }
]

const queryCornSuggest = (queryString: string, cb: (results: any[]) => void) => {
    if (queryString) {
        cb(
            cornOptions.filter(
                (item) => item.value.includes(queryString) || item.desc.includes(queryString)
            )
        )
    } else {
        cb(cornOptions)
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
<style lang="scss" scoped>
.corn-option {
    display: flex;
    align-items: center;
    justify-content: space-between;
    gap: 12px;
    .corn-value {
        font-family: monospace;
        color: var(--el-color-primary);
    }
    .corn-label {
        color: var(--el-text-color-secondary);
        font-size: 12px;
    }
}
</style>
