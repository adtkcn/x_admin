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
                <el-form-item label="申请id" prop="applyId">
                    <el-input v-model="formData.applyId" placeholder="请输入申请id" />
                </el-form-item>
                <el-form-item label="模板id" prop="templateId">
                    <el-input v-model="formData.templateId" placeholder="请输入模板id" />
                </el-form-item>
                <el-form-item label="申请人id" prop="applyUserId">
                    <el-input v-model="formData.applyUserId" placeholder="请输入申请人id" />
                </el-form-item>
                <el-form-item label="申请人昵称" prop="applyUserNickname">
                    <el-input v-model="formData.applyUserNickname" placeholder="请输入申请人昵称" />
                </el-form-item>
                <el-form-item label="审批人id" prop="approverId">
                    <el-input v-model="formData.approverId" placeholder="请输入审批人id" />
                </el-form-item>
                <el-form-item label="审批用户昵称" prop="approverNickname">
                    <el-input
                        v-model="formData.approverNickname"
                        placeholder="请输入审批用户昵称"
                    />
                </el-form-item>
                <el-form-item label="节点" prop="nodeId">
                    <el-input v-model="formData.nodeId" placeholder="请输入节点" />
                </el-form-item>
                <el-form-item label="表单值" prop="formValue">
                    <el-input
                        v-model="formData.formValue"
                        placeholder="请输入表单值"
                        type="textarea"
                        :autosize="{ minRows: 4, maxRows: 6 }"
                    />
                </el-form-item>
                <el-form-item label="通过状态：0待处理，1通过，2拒绝" prop="passStatus">
                    <el-radio-group
                        v-model="formData.passStatus"
                        placeholder="请选择通过状态：0待处理，1通过，2拒绝"
                    >
                        <el-radio label="0">请选择字典生成</el-radio>
                    </el-radio-group>
                </el-form-item>
                <el-form-item label="通过备注" prop="passRemark">
                    <el-input v-model="formData.passRemark" placeholder="请输入通过备注" />
                </el-form-item>
            </el-form>
        </popup>
    </div>
</template>
<script lang="ts" setup>
import type { FormInstance } from 'element-plus'
import { flow_history_edit, flow_history_add, flow_history_detail } from '@/api/flow/flow_history'
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
    return mode.value == 'edit' ? '编辑流程历史' : '新增流程历史'
})

const formData = reactive({
    id: '',
    applyId: '',
    templateId: '',
    applyUserId: '',
    applyUserNickname: '',
    approverId: '',
    approverNickname: '',
    nodeId: '',
    formValue: '',
    passStatus: '',
    passRemark: ''
})

const formRules = {
    id: [
        {
            required: true,
            message: '请输入历史id',
            trigger: ['blur']
        }
    ],
    applyId: [
        {
            required: true,
            message: '请输入申请id',
            trigger: ['blur']
        }
    ],
    templateId: [
        {
            required: true,
            message: '请输入模板id',
            trigger: ['blur']
        }
    ],
    applyUserId: [
        {
            required: true,
            message: '请输入申请人id',
            trigger: ['blur']
        }
    ],
    applyUserNickname: [
        {
            required: true,
            message: '请输入申请人昵称',
            trigger: ['blur']
        }
    ],
    approverId: [
        {
            required: true,
            message: '请输入审批人id',
            trigger: ['blur']
        }
    ],
    approverNickname: [
        {
            required: true,
            message: '请输入审批用户昵称',
            trigger: ['blur']
        }
    ],
    nodeId: [
        {
            required: true,
            message: '请输入节点',
            trigger: ['blur']
        }
    ],
    formValue: [
        {
            required: true,
            message: '请输入表单值',
            trigger: ['blur']
        }
    ],
    passStatus: [
        {
            required: true,
            message: '请选择通过状态：0待处理，1通过，2拒绝',
            trigger: ['blur']
        }
    ],
    passRemark: [
        {
            required: true,
            message: '请输入通过备注',
            trigger: ['blur']
        }
    ]
}

const handleSubmit = async () => {
    await formRef.value?.validate()
    const data: any = { ...formData }
    mode.value == 'edit' ? await flow_history_edit(data) : await flow_history_add(data)
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
    const data = await flow_history_detail({
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
@/api/flow/flow_history
