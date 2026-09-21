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
                <el-form-item label="邮箱(账号)" prop="email">
                    <el-input
                        v-model="formData.email"
                        :disabled="isRoot"
                        placeholder="请输入邮箱"
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

                <el-form-item label="角色" prop="role_ids">
                    <el-select
                        v-model="formData.role_ids"
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

                <el-form-item label="归属部门" prop="dept_id">
                    <el-tree-select
                        class="flex-1"
                        v-model="formData.dept_id"
                        :data="deptTreeList"
                        clearable
                        node-key="id"
                        :props="{
                            // value: 'id',
                            label: 'name',
                            disabled(data: any) {
                                return !!data.is_stop
                            }
                        }"
                        :disabled="isRoot"
                        check-strictly
                        :default-expand-all="true"
                        placeholder="请选择上级部门"
                    />
                </el-form-item>
                <el-form-item label="岗位" prop="post_id">
                    <el-select
                        class="flex-1"
                        clearable
                        v-model="formData.post_id"
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
                    <el-switch
                        v-model="formData.is_disable"
                        :active-value="0"
                        :inactive-value="1"
                    />
                </el-form-item>
            </el-form>
        </popup>
    </div>
</template>
<script lang="ts" setup>
import { ref, computed, shallowRef, reactive } from 'vue'
import type { FormInstance } from 'element-plus'
import Popup from '@/components/popup/index.vue'
import {
    adminAdd,
    adminEdit,
    adminDetail,
    type type_system_admin_edit,
    type type_system_admin_resp
} from '@/api/perms/admin'
import { useDictOptions } from '@/hooks/useDictOptions'
import { roleAll, type type_system_role_simple_resp } from '@/api/perms/role'
import { postAll, type type_system_post_resp } from '@/api/org/post'
import { deptLists, type type_system_dept_resp } from '@/api/org/department'
import feedback from '@/utils/feedback'
import { encryptPassword, arrayToTree } from '@/utils/util'
import { useReactiveWithReset } from '@/hooks/useReactiveWithReset'

const emit = defineEmits(['success', 'close'])
const formRef = shallowRef<FormInstance>()
const popupRef = shallowRef<InstanceType<typeof Popup>>()
const mode = ref('add')
const popupTitle = computed(() => {
    return mode.value == 'edit' ? '编辑管理员' : '新增管理员'
})

type type_admin_form = type_system_admin_edit & {
    passwordConfirm?: string
}

const { state: formData, setState } = useReactiveWithReset<type_admin_form>({
    id: '',
    email: '',
    nickname: '',
    dept_id: '',
    post_id: '',
    role_ids: [],
    avatar: '',
    password: '',
    passwordConfirm: '',
    is_disable: 0,
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
    email: [
        {
            required: true,
            message: '请输入邮箱',
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
    try {
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
    } catch (error) {
        console.error('管理员保存失败:', error)
    }
}

const open = (type = 'add') => {
    mode.value = type
    popupRef.value?.open()
}

const setFormData = async (row: type_system_admin_resp) => {
    try {
        const data = await adminDetail({
            id: row.id
        })
        setState(data)
        formRules.password = []
        formRules.passwordConfirm = [
            {
                validator: passwordConfirmValidator,
                trigger: 'blur'
            }
        ]
    } catch (error) {
        console.error('管理员详情获取失败:', error)
    }
}

const handleClose = () => {
    emit('close')
}

defineExpose({
    open,
    setFormData
})
</script>
