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
                <el-form-item label="标题" prop="Title">
                    <el-input v-model="formData.Title" placeholder="请输入标题" />
                </el-form-item>
                <el-form-item label="协议内容" prop="Content">
                    <editor v-model="formData.Content" :height="500" />
                </el-form-item>
                <el-form-item label="排序" prop="Sort">
                    <el-input v-model="formData.Sort" type="number" placeholder="请输入排序" />
                </el-form-item>
            </el-form>
        </popup>
    </div>
</template>
<script lang="ts" setup>
import type { FormInstance } from 'element-plus'
import { user_protocol_edit, user_protocol_add, user_protocol_detail } from '@/api/user/protocol'
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
    return mode.value == 'edit' ? '编辑用户协议' : '新增用户协议'
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
    ]
    // Sort: [
    //     {
    //         required: true,
    //         message: '请输入排序',
    //         trigger: ['blur']
    //     }
    // ],
}

const handleSubmit = async () => {
    try {
        await formRef.value?.validate()
        const data: any = { ...formData }
        mode.value == 'edit' ? await user_protocol_edit(data) : await user_protocol_add(data)
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
        const data = await user_protocol_detail(row.Id)
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
