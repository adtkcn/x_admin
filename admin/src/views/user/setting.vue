<!-- 个人资料 -->
<template>
    <div class="user-setting">
        <!-- 基本信息 -->
        <el-card class="border-none! mb-4" shadow="never">
            <template #header>
                <span class="font-medium text-base">基本信息</span>
            </template>
            <el-form
                ref="baseFormRef"
                :model="formData"
                :rules="baseRules"
                label-width="90px"
                class="max-w-[520px]"
            >
                <el-form-item label="头像" prop="avatar">
                    <material-picker v-model="formData.avatar" :limit="1" />
                </el-form-item>

                <el-form-item label="昵称" prop="nickname">
                    <el-input v-model="formData.nickname" placeholder="请输入昵称" />
                </el-form-item>

                <el-form-item label="邮箱(账号)" prop="email">
                    <div class="flex gap-2 w-full">
                        <el-input
                            v-model="formData.email"
                            placeholder="请输入新邮箱"
                            @input="onEmailChange"
                        />
                        <el-button
                            type="primary"
                            :loading="sendingCode"
                            :disabled="!canSendCode || countdown > 0"
                            @click="handleSendCode"
                        >
                            {{ countdown > 0 ? `${countdown}s 后重发` : '发送验证码' }}
                        </el-button>
                    </div>
                </el-form-item>

                <el-form-item
                    v-if="needVerify"
                    label="验证码"
                    prop="emailCode"
                >
                    <el-input
                        v-model="formData.emailCode"
                        placeholder="请输入邮箱验证码"
                        maxlength="6"
                        class="w-40"
                    />
                </el-form-item>
            </el-form>
        </el-card>

        <!-- 修改密码 -->
        <el-card class="border-none! mb-4" shadow="never">
            <template #header>
                <span class="font-medium text-base">修改密码</span>
            </template>
            <el-form
                ref="pwdFormRef"
                :model="formData"
                :rules="pwdRules"
                label-width="90px"
                class="max-w-[520px]"
            >
                <input
                    type="password"
                    style="position:absolute;height:0;width:0;line-height:0;border:0;opacity:0;overflow:hidden"
                />
                <el-form-item label="当前密码" prop="currPassword">
                    <el-input
                        v-model.trim="formData.currPassword"
                        type="password"
                        show-password
                        placeholder="修改密码时必填"
                        autocomplete="new-password"
                    />
                </el-form-item>

                <el-form-item label="新密码" prop="password">
                    <el-input
                        v-model.trim="formData.password"
                        type="password"
                        show-password
                        placeholder="修改密码时必填"
                    />
                </el-form-item>

                <el-form-item label="确认密码" prop="passwordConfirm">
                    <el-input
                        v-model.trim="formData.passwordConfirm"
                        type="password"
                        show-password
                        placeholder="请再次输入新密码"
                    />
                </el-form-item>
            </el-form>
        </el-card>

        <div class="flex justify-center">
            <el-button type="primary" size="large" @click="handleSubmit">保 存</el-button>
        </div>
    </div>
</template>

<script setup lang="ts">
import { reactive, ref, computed, useTemplateRef } from 'vue'
import { setUserInfo, sendEmailCode } from '@/api/user'
import useUserStore from '@/stores/modules/user'
import feedback from '@/utils/feedback'
import { encryptPassword } from '@/utils/util'
import type { FormInstance } from 'element-plus'

defineOptions({ name: 'userSetting' })

const baseFormRef = useTemplateRef<FormInstance>('baseFormRef')
const pwdFormRef = useTemplateRef<FormInstance>('pwdFormRef')
const userStore = useUserStore()

const originalEmail = ref('')
const sendingCode = ref(false)
const countdown = ref(0)
let timer: ReturnType<typeof setInterval> | null = null

const formData = reactive({
    avatar: '',
    email: '',
    emailCode: '',
    nickname: '',
    currPassword: '',
    password: '',
    passwordConfirm: ''
})

// 邮箱是否有变更
const emailChanged = computed(() => formData.email !== originalEmail.value)
// 是否可以发送验证码
const canSendCode = computed(() => emailChanged.value && formData.email.includes('@'))
// 是否需要验证码（邮箱变更且已填入新值）
const needVerify = computed(() => emailChanged.value && formData.email !== '')

const baseRules = {
    avatar: [{ required: true, message: '头像不能为空', trigger: 'blur' }],
    nickname: [{ required: true, message: '请输入昵称', trigger: 'blur' }],
    email: [
        { required: true, message: '邮箱不能为空', trigger: 'blur' },
        { type: 'email', message: '请输入正确的邮箱格式', trigger: 'blur' }
    ],
    emailCode: [{
        validator: (_r: any, v: string, cb: any) => {
            if (needVerify.value && !v) return cb(new Error('请输入验证码'))
            cb()
        },
        trigger: 'blur'
    }]
}

const pwdRules = {
    currPassword: [{
        validator: (_r: any, v: string, cb: any) => {
            if (formData.password && !v) return cb(new Error('请输入当前密码'))
            cb()
        },
        trigger: 'blur'
    }],
    password: [{
        validator: (_r: any, v: string, cb: any) => {
            if (formData.currPassword && !v) return cb(new Error('请输入新密码'))
            cb()
        },
        trigger: 'blur'
    }],
    passwordConfirm: [{
        validator: (_r: any, v: string, cb: any) => {
            if (formData.password) {
                if (!v) return cb(new Error('请再次输入密码'))
                if (v !== formData.password) return cb(new Error('两次输入密码不一致'))
            }
            cb()
        },
        trigger: 'blur'
    }]
}

function onEmailChange() {
    formData.emailCode = ''
}

async function handleSendCode() {
    if (!canSendCode.value || countdown.value > 0) return
    sendingCode.value = true
    try {
        await sendEmailCode({ email: formData.email })
        feedback.msgSuccess('验证码已发送')
        countdown.value = 60
        timer = setInterval(() => {
            countdown.value--
            if (countdown.value <= 0) {
                clearInterval(timer!)
                timer = null
            }
        }, 1000)
    } catch {
    } finally {
        sendingCode.value = false
    }
}

const getUser = async () => {
    const userInfo = userStore.userInfo
    formData.avatar = userInfo.avatar || ''
    formData.email = userInfo.email || ''
    formData.nickname = userInfo.nickname || ''
    formData.currPassword = ''
    formData.password = ''
    formData.passwordConfirm = ''
    formData.emailCode = ''
    originalEmail.value = userInfo.email || ''
}

async function handleSubmit() {
    await Promise.all([baseFormRef.value?.validate(), pwdFormRef.value?.validate()])

    const info: Record<string, any> = {
        avatar: formData.avatar,
        nickname: formData.nickname
    }

    // 邮箱变更时附带新邮箱和验证码
    if (emailChanged.value) {
        info.email = formData.email
        info.emailCode = formData.emailCode
    }

    if (formData.password) {
        info.password = encryptPassword(formData.password)
        info.currPassword = encryptPassword(formData.currPassword)
    }

    await setUserInfo(info)
    feedback.msgSuccess('保存成功')
    await userStore.getUserInfo()
    originalEmail.value = formData.email
    formData.emailCode = ''
    formData.currPassword = ''
    formData.password = ''
    formData.passwordConfirm = ''
}

getUser()
</script>

<style lang="scss" scoped>
.user-setting {
    max-width: 680px;
    margin: 0 auto;
}
</style>
