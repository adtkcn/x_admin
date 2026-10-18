<template>
    <div class="edit-popup">
        <popup
            ref="popupRef"
            title="预览定时任务"
            :async="true"
            width="550px"
            :clickModalClose="true"
            :confirmButtonText="false"
            @close="handleClose"
        >
            <el-form label-width="110px">
                <el-form-item label="任务名称">
                    <span>{{ formData.task_name }}</span>
                </el-form-item>
                <el-form-item label="任务编码">
                    <span>{{ formData.task_code }}</span>
                </el-form-item>
                <el-form-item label="corn表达式">
                    <span>{{ formData.corn_expr }}</span>
                </el-form-item>
                <el-form-item label="状态">
                    <dict-value :options="dictData.status" :value="formData.status" />
                </el-form-item>
            </el-form>
        </popup>
    </div>
</template>

<script lang="ts" setup>
import { system_corn_detail } from '@/api/system/corn'
import Popup from '@/components/popup/index.vue'
import DictValue from '@/components/dict-value/index.vue'
import { useDictData } from '@/hooks/useDictOptions'
import type { type_dict } from '@/hooks/useDictOptions'
import { shallowRef, reactive } from 'vue'

const { dictData } = useDictData<{ status: type_dict[] }>(['status'])

const emit = defineEmits(['close'])
const popupRef = shallowRef<InstanceType<typeof Popup>>()

const formData = reactive<Record<string, any>>({})

const open = () => {
    popupRef.value?.open()
}

const setFormData = (data: Record<string, any>) => {
    Object.assign(formData, data)
}

const getDetail = async (row: Record<string, any>) => {
    const data = await system_corn_detail(row.id)
    setFormData(data)
}

const handleClose = () => {
    emit('close')
}

defineExpose({ open, setFormData, getDetail })
</script>
