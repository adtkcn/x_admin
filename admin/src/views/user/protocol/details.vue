<template>
    <div class="edit-popup">
        <popup
            ref="popupRef"
            :title="popupTitle"
            :async="true"
            width="550px"
            :clickModalClose="true"
            :confirmButtonText="false"
            @close="handleClose"
        >
            <el-form ref="formRef" :model="formData" label-width="84px" :rules="formRules">
                <el-form-item label="标题" prop="Title">
                    <span v-text="formData.Title"></span>
                </el-form-item>
                <el-form-item label="协议内容" prop="Content">
                    <div v-html="formData.Content"></div>
                </el-form-item>
                <el-form-item label="排序" prop="Sort">
                    <span v-text="formData.Sort"></span>
                </el-form-item>
            </el-form>
        </popup>
    </div>
</template>
<script lang="ts" setup>
import type { FormInstance } from 'element-plus'
import { user_protocol_detail } from '@/api/user/protocol'
import Popup from '@/components/popup/index.vue'

import { shallowRef, computed, reactive } from 'vue'
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
const emit = defineEmits(['close'])
const formRef = shallowRef<FormInstance>()
const popupRef = shallowRef<InstanceType<typeof Popup>>()

const popupTitle = computed(() => {
    return '预览用户协议'
})

const formData = reactive({
    Id: null,
    Title: null,
    Content: null,
    Sort: null
})

const formRules = {
    Id: [
        {
            required: true,
            message: '请输入',
            trigger: ['blur']
        }
    ],
    Title: [
        {
            required: true,
            message: '请输入标题',
            trigger: ['blur']
        }
    ],
    Content: [
        {
            required: true,
            message: '请输入协议内容',
            trigger: ['blur']
        }
    ],
    Sort: [
        {
            required: true,
            message: '请输入排序',
            trigger: ['blur']
        }
    ]
}

const open = () => {
    popupRef.value?.open()
}
const getDetail = async (row: Record<string, any>) => {
    try {
        const data = await user_protocol_detail(row.Id)
        setFormData(data)
    } catch (error) {}
}
const setFormData = async (data: Record<string, any>) => {
    for (const key in formData) {
        if (data[key] != null && data[key] != undefined) {
            //@ts-ignore
            formData[key] = data[key]
        }
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
