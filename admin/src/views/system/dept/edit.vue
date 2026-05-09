<template>
    <div class="edit-popup">
        <popup
            ref="popupRef"
            :title="popupTitle"
            :async="true"
            width="550px"
            @confirm="handleSubmit"
            @close="handleClose"
        >
            <el-form ref="formRef" :model="formData" label-width="110px" :rules="formRules">
                <el-form-item label="上级部门" prop="pid" v-if="formData.pid !== ''">
                    <el-tree-select
                        class="flex-1"
                        v-model="formData.pid"
                        :data="optionsData.dept"
                        clearable
                        node-key="id"
                        :props="{
                            label: 'name'
                        }"
                        check-strictly
                        :default-expand-all="true"
                        placeholder="请选择上级部门"
                    />
                </el-form-item>
                <el-form-item label="部门名称" prop="name">
                    <el-input
                        v-model="formData.name"
                        placeholder="请输入部门名称"
                        clearable
                        :maxlength="100"
                    />
                </el-form-item>
                <el-form-item label="负责人" prop="duty_id" v-if="formData.id">
                    <el-select
                        class="flex-1"
                        v-model="formData.dutyId"
                        clearable
                        placeholder="请选择上级部门"
                        @change="dutyChange"
                    >
                        <el-option label="请先给管理员绑定部门" value="" />

                        <el-option
                            v-for="item in DeptUsers"
                            :key="item.id"
                            :label="item.nickname"
                            :value="item.id"
                        />
                    </el-select>
                </el-form-item>
                <el-form-item label="部门电话" prop="mobile">
                    <el-input v-model="formData.mobile" placeholder="请输入联系电话" clearable />
                </el-form-item>
                <el-form-item label="排序" prop="sort">
                    <div>
                        <el-input-number v-model="formData.sort" :min="0" :max="9999" />
                        <div class="form-tips">默认为0， 数值越大越排前</div>
                    </div>
                </el-form-item>
                <el-form-item label="部门状态" prop="isStop">
                    <el-switch v-model="formData.isStop" :active-value="0" :inactive-value="1" />
                </el-form-item>
            </el-form>
        </popup>
    </div>
</template>
<script lang="ts" setup>
import { ref, computed, shallowRef, reactive } from 'vue'
import type { FormInstance } from 'element-plus'
import {
    deptLists,
    deptEdit,
    deptAdd,
    deptDetail,
    type type_system_dept_add,
    type type_system_dept_edit,
    type type_system_dept_resp
} from '@/api/org/department'
import { adminListByDeptId, type type_system_admin_resp } from '@/api/perms/admin'

import Popup from '@/components/popup/index.vue'
import { useDictOptions } from '@/hooks/useDictOptions'
import feedback from '@/utils/feedback'
const emit = defineEmits(['success', 'close'])
const formRef = shallowRef<FormInstance>()
const popupRef = shallowRef<InstanceType<typeof Popup>>()
const mode = ref('add')
const popupTitle = computed(() => {
    return mode.value == 'edit' ? '编辑部门' : '新增部门'
})

const formData = reactive<type_system_dept_edit>({
    id: '',
    pid: '',
    name: '',
    dutyId: '',
    duty: '',
    mobile: '',
    sort: 0,
    isStop: 0
})
const DeptUsers = ref<type_system_admin_resp[]>([])
// 部门下的管理员
async function getDeptUsers(deptId: string) {
    const users = await adminListByDeptId({ deptId: deptId })
    DeptUsers.value = users
}
function dutyChange(id: string) {
    if (id) {
        const duty = DeptUsers.value.find((item) => item.id == id)
        formData.duty = duty?.nickname || ''
    } else {
        formData.duty = ''
    }
}
const checkMobile = (rule: any, value: any, callback: any) => {
    if (!value) {
        return callback()
    } else {
        const reg = /^[1][3,4,5,6,7,8,9][0-9]{9}$/
        console.log(reg.test(value))
        if (reg.test(value)) {
            callback()
        } else {
            return callback(new Error('请输入正确的手机号'))
        }
    }
}
const formRules = {
    pid: [
        {
            required: true,
            message: '请选择上级部门',
            trigger: ['change']
        }
    ],
    name: [
        {
            required: true,
            message: '请输入部门名称',
            trigger: ['blur']
        }
    ],
    duty: [
        {
            required: true,
            message: '请输入负责人姓名',
            trigger: ['blur']
        }
    ],
    mobile: [
        // {
        //     required: true,
        //     message: '请输入联系电话',
        //     trigger: ['blur']
        // },
        {
            validator: checkMobile,
            trigger: ['blur']
        }
    ]
}

const { optionsData } = useDictOptions<{
    dept: type_system_dept_resp[]
}>({
    dept: {
        api: deptLists
    }
})

const handleSubmit = async () => {
    await formRef.value?.validate()
    mode.value == 'edit' ? await deptEdit(formData) : await deptAdd(formData)
    popupRef.value?.close()
    feedback.msgSuccess('操作成功')
    emit('success')
}

const open = (type = 'add') => {
    mode.value = type
    popupRef.value?.open()
}

const setFormData = (data: Partial<type_system_dept_edit>) => {
    for (const key in formData) {
        if (data[key] != null && data[key] != undefined) {
            formData[key] = data[key]
        }
    }
}

const getDetail = async (row: type_system_dept_resp) => {
    const data = await deptDetail({
        id: row.id
    })
    getDeptUsers(data.id)
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
