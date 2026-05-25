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
                <el-form-item label="账号" prop="username">
                    <el-input
                        v-model="formData.username"
                        :disabled="isRoot"
                        placeholder="请输入账号"
                        clearable
                    />
                </el-form-item>
                <el-form-item label="头像" prop="avatar">
                    <div>
                        <div>
                            <material-picker v-model="formData.avatar" :limit="1" />
                        </div>
                        <div class="form-tips">建议尺寸：100*100px，支持jpg，jpeg，png格式</div>
                    </div>
                </el-form-item>
                <el-form-item label="名称" prop="nickname">
                    <el-input v-model="formData.nickname" placeholder="请输入名称" clearable />
                </el-form-item>

                <el-form-item label="角色" prop="roleIds">
                    <el-select
                        v-model="formData.roleIds"
                        :disabled="isRoot"
                        class="flex-1"
                        multiple
                        clearable
                        placeholder="请选择角色"
                    >
                        <el-option v-if="isRoot" label="系统管理员" :value="0" />
                        <el-option
                            v-for="(item, index) in optionsData.role"
                            :key="index"
                            :label="item.name"
                            :value="item.id"
                        />
                    </el-select>
                </el-form-item>

                <el-form-item label="归属部门" prop="deptId">
                    <el-tree-select
                        class="flex-1"
                        v-model="formData.deptId"
                        :data="deptTreeList"
                        clearable
                        node-key="id"
                        :props="{
                            // value: 'id',
                            label: 'name',
                            disabled(data: any) {
                                return !!data.isStop
                            }
                        }"
                        :disabled="isRoot"
                        check-strictly
                        :default-expand-all="true"
                        placeholder="请选择上级部门"
                    />
                </el-form-item>
                <el-form-item label="岗位" prop="postId">
                    <el-select
                        class="flex-1"
                        clearable
                        v-model="formData.postId"
                        placeholder="请选择岗位"
                        :disabled="isRoot"
                    >
                        <!-- multiple -->
                        <el-option
                            v-for="(item, index) in optionsData.post"
                            :key="index"
                            :label="item.name"
                            :value="item.id"
                        />
                    </el-select>
                </el-form-item>

                <el-form-item label="密码" prop="password">
                    <el-input
                        v-model.trim="formData.password"
                        show-password
                        clearable
                        placeholder="请输入密码"
                    />
                </el-form-item>

                <el-form-item label="确认密码" prop="passwordConfirm">
                    <el-input
                        v-model.trim="formData.passwordConfirm"
                        show-password
                        clearable
                        placeholder="请输入确认密码"
                    />
                </el-form-item>

                <el-form-item label="状态" v-if="!isRoot">
                    <el-switch v-model="formData.isDisable" :active-value="0" :inactive-value="1" />
                </el-form-item>
            </el-form>
        </popup>
    </div>
</template>
<script lang="ts" setup>
import { ref, computed, reactive, shallowRef } from 'vue'
import type { FormInstance } from 'element-plus'
import Popup from '@/components/popup/index.vue'
import {
    adminAdd,
    adminEdit,
    adminDetail,
    type type_system_admin_add,
    type type_system_admin_resp
} from '@/api/perms/admin'
import { useDictOptions } from '@/hooks/useDictOptions'
import { roleAll, type type_system_role_simple_resp } from '@/api/perms/role'
import { postAll, type type_system_post_resp } from '@/api/org/post'
import { deptLists, type type_system_dept_resp } from '@/api/org/department'
import feedback from '@/utils/feedback'
import { encryptPassword, arrayToTree } from '@/utils/util'

const emit = defineEmits(['success', 'close'])
const formRef = shallowRef<FormInstance>()
const popupRef = shallowRef<InstanceType<typeof Popup>>()
const mode = ref('add')
const popupTitle = computed(() => {
    return mode.value == 'edit' ? '编辑管理员' : '新增管理员'
})

type type_admin_form = type_system_admin_add & {
    id: string
    passwordConfirm: string
}

const formData = reactive<type_admin_form>({
    id: '',
    username: '',
    nickname: '',
    deptId: '',
    postId: '',
    roleIds: [],
    avatar: '',
    password: '',
    passwordConfirm: '',
    isDisable: 0,
    sort: 1
})

const isRoot = computed(() => {
    return formData.id == '1'
})

const passwordConfirmValidator = (rule: object, value: string, callback: any) => {
    if (formData.password) {
        if (!value) callback(new Error('请再次输入密码'))
        if (value !== formData.password) callback(new Error('两次输入密码不一致!'))
    }
    callback()
}
const formRules = reactive({
    username: [
        {
            required: true,
            message: '请输入账号',
            trigger: ['blur']
        }
    ],
    nickname: [
        {
            required: true,
            message: '请输入名称',
            trigger: ['blur']
        }
    ],
    password: [
        {
            required: true,
            message: '请输入密码',
            trigger: 'blur'
        }
    ] as any[],
    passwordConfirm: [
        {
            required: true,
            message: '请再次输入密码',
            trigger: 'blur'
        },
        {
            validator: passwordConfirmValidator,
            trigger: 'blur'
        }
    ] as any[]
})

const { optionsData } = useDictOptions<{
    role: type_system_role_simple_resp[]
    post: type_system_post_resp[]
    dept: type_system_dept_resp[]
}>({
    role: {
        api: roleAll
    },
    post: {
        api: postAll
    },
    dept: {
        api: deptLists
    }
})
const deptTreeList = computed(() => {
    const treeList = arrayToTree(optionsData.dept, '')
    return treeList
})
const handleSubmit = async () => {
    await formRef.value?.validate()
    const data: any = {
        ...formData
    }
    delete data.passwordConfirm
    if (formData.password) {
        data.password = encryptPassword(formData.password)
    }
    if (mode.value == 'edit') {
        await adminEdit(data)
    } else {
        await adminAdd(data)
    }
    popupRef.value?.close()
    feedback.msgSuccess('操作成功')
    emit('success')
}

const open = (type = 'add') => {
    mode.value = type
    popupRef.value?.open()
}

const setFormData = async (row: type_system_admin_resp) => {
    const data = await adminDetail({
        id: row.id
    })
    for (const key in formData) {
        if (
            data[key as keyof type_system_admin_resp] != null &&
            data[key as keyof type_system_admin_resp] != undefined
        ) {
            //@ts-ignore
            formData[key as keyof type_admin_form] = data[key as keyof type_system_admin_resp]
        }
    }
    formRules.password = []
    formRules.passwordConfirm = [
        {
            validator: passwordConfirmValidator,
            trigger: 'blur'
        }
    ]
}

const handleClose = () => {
    emit('close')
}

defineExpose({
    open,
    setFormData
})
</script>
