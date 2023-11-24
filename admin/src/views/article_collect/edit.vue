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
                <el-form-item label="用户ID" prop="userId">
                    <el-input v-model.number="formData.userId" placeholder="请输入用户ID" />
                </el-form-item>
                <el-form-item label="文章ID" prop="articleId">
                    <el-input v-model.number="formData.articleId" placeholder="请输入文章ID" />
                </el-form-item>
            </el-form>
        </popup>
    </div>
</template>
<script lang="ts" setup>
import type { FormInstance } from 'element-plus'
import {
    article_collect_edit,
    article_collect_add,
    article_collect_detail
} from '@/api/article_collect'
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
    return mode.value == 'edit' ? '编辑文章收藏' : '新增文章收藏'
})

const formData = reactive({
    id: '',
    userId: '',
    articleId: ''
})

const formRules = {
    id: [
        {
            required: true,
            message: '请输入主键',
            trigger: ['blur']
        }
    ],
    userId: [
        {
            required: true,
            message: '请输入用户ID',
            trigger: ['blur']
        }
    ],
    articleId: [
        {
            required: true,
            message: '请输入文章ID',
            trigger: ['blur']
        }
    ]
}

const handleSubmit = async () => {
    await formRef.value?.validate()
    const data: any = { ...formData }
    mode.value == 'edit' ? await article_collect_edit(data) : await article_collect_add(data)
    popupRef.value?.close()
    feedback.msgSuccess('操作成功')
    emit('success')
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
    const data = await article_collect_detail({
        id: row.id
    })
    setFormData(data)
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
