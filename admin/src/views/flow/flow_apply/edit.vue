<template>
    <div class="edit-popup">
        <popup
            ref="popupRef"
            :title="popupTitle"
            :async="true"
            width="550px"
            :clickModalClose="true"
            confirmButtonText="发起申请"
            @confirm="handleSubmit"
            @close="handleClose"
        >
            <el-form ref="formRef" :model="formData" label-width="110px" :rules="formRules">
                <el-form-item label="审批流" prop="template_id">
                    <el-select
                        v-model="formData.template_id"
                        placeholder="请选择审批流"
                        style="width: 100%"
                        @change="handleTemplateChange"
                    >
                        <el-option
                            v-for="(item, index) in flow_template"
                            :key="index"
                            :label="item.flow_name"
                            :value="item.id"
                            clearable
                        />
                    </el-select>
                </el-form-item>

                <el-form-item label="流程名称" prop="flow_name">
                    <el-input v-model="formData.flow_name" placeholder="请输入流程名称" />
                </el-form-item>
            </el-form>
        </popup>
    </div>
</template>
<script lang="ts" setup>
import type { FormInstance } from 'element-plus'
import {
    flow_apply_edit,
    flow_apply_add,
    flow_apply_detail,
    type type_flow_apply_edit
} from '@/api/flow/flow_apply'
import { flow_template_lists_all } from '@/api/flow/flow_template'
import type { type_flow_template } from '@/api/flow/flow_template'

import useUserStore from '@/stores/modules/user'
import Popup from '@/components/popup/index.vue'
import feedback from '@/utils/feedback'
import { shallowRef, ref, computed } from 'vue'
import { useReactiveWithReset } from '@/hooks/useReactiveWithReset'
import type { PropType } from 'vue'

defineProps({
    dictData: {
        type: Object as PropType<Record<string, any[]>>,
        default: () => ({})
    }
})
const userStore = useUserStore()
const emit = defineEmits(['success', 'close'])
const formRef = shallowRef<FormInstance>()
const popupRef = shallowRef<InstanceType<typeof Popup>>()
const mode = ref('add')
const popupTitle = computed(() => {
    return mode.value == 'edit' ? '编辑申请流程' : '新增申请流程'
})

const { state: formData, setState } = useReactiveWithReset<type_flow_apply_edit>({
    id: '',
    template_id: '',
    flow_name: '',
    status: 0
})

const formRules = {
    id: [
        {
            required: true,
            message: '请输入',
            trigger: ['blur']
        }
    ],
    template_id: [
        {
            required: true,
            message: '请输入模板',
            trigger: ['blur']
        }
    ],
    apply_user_id: [
        {
            required: true,
            message: '请输入申请人id',
            trigger: ['blur']
        }
    ],
    apply_user_nickname: [
        {
            required: true,
            message: '请输入申请人昵称',
            trigger: ['blur']
        }
    ],
    flow_name: [
        {
            required: true,
            message: '请输入流程名称',
            trigger: ['blur']
        }
    ],
    flow_group: [
        {
            required: true,
            message: '请输入流程分类',
            trigger: ['blur']
        }
    ],
    flow_remark: [
        {
            required: true,
            message: '请输入流程描述',
            trigger: ['blur']
        }
    ],
    flow_form_data: [
        {
            required: true,
            message: '请输入表单配置',
            trigger: ['blur']
        }
    ],
    flow_process_data: [
        {
            required: true,
            message: '请输入流程配置',
            trigger: ['blur']
        }
    ],
    status: [
        {
            required: true,
            message: '请选择状态',
            trigger: ['blur']
        }
    ]
}
const flow_template = ref<type_flow_template[]>([])
const get_flow_template = async () => {
    flow_template.value = await flow_template_lists_all()
}
function handleTemplateChange(id: string) {
    flow_template.value.find((item: any) => {
        if (item.id == id) {
            formData.flow_name = item.flow_name
            // '【' +
            // item.flow_name +
            // '】' +
            // userStore.userInfo.nickname +
            // ' - ' +
            // dayjs().format('YYYY-MM-DD')
            // formData.flow_group = item.flow_group
            // formData.flow_remark = item.flow_remark
            // formData.flow_form_data = item.flow_form_data
            // formData.flow_process_data = item.flow_process_data
            return true
        }
    })
}
get_flow_template()
const handleSubmit = async () => {
    await formRef.value?.validate()
    const data: any = { ...formData }
    // if ( !data.id) {
    //     delete data.id
    // }
    mode.value == 'edit' ? await flow_apply_edit(data) : await flow_apply_add(data)
    popupRef.value?.close()
    feedback.msgSuccess('操作成功')
    emit('success')
}

const open = (type = 'add') => {
    mode.value = type

    popupRef.value?.open()
}

const setFormData = async (data: type_flow_apply_edit) => {
    setState(data)
}

const getDetail = async (row: Record<string, any>) => {
    const data = await flow_apply_detail({
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
