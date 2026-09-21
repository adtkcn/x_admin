<template>
    <div class="edit-popup">
        <popup
            ref="popupRef"
            :title="popuptitle"
            :async="true"
            width="550px"
            :clickModalClose="true"
            :confirmButtonText="false"
            @close="handleClose"
        >
            <el-form ref="formRef" :model="formData" label-width="110px" :rules="formRules">
                <el-form-item label="标题" prop="title">
                    <span v-text="formData.title"></span>
                </el-form-item>
                <el-form-item label="协议内容" prop="content">
                    <div class="rich-text-style" v-html="formData.content"></div>
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
import '@/components/editor/rich-text-style.css'

import { useTemplateRef, computed, reactive } from 'vue'
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
const formRef = useTemplateRef<FormInstance>('formRef')
const popupRef = useTemplateRef<InstanceType<typeof Popup>>('popupRef')

const popuptitle = computed(() => {
    return '预览用户协议'
})

const formData = reactive({
    id: null,
    title: null,
    content: null,
    Sort: null
})

const formRules = {
    id: [
        {
            required: true,
            message: '请输入',
            trigger: ['blur']
        }
    ],
    title: [
        {
            required: true,
            message: '请输入标题',
            trigger: ['blur']
        }
    ],
    content: [
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
        const data = await user_protocol_detail(row.id)
        setFormData(data)
    } catch (error) {
        console.error('协议详情获取失败:', error)
    }
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
