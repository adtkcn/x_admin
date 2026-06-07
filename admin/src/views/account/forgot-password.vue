<template>
    <div class="forgot-wd flex flex-col">
        <div class="flex-1 flex items-center justify-center">
            <div class="forgot-card flex rounded-md">
                <div class="flex-1 h-full hidden md:inline-block">
                    <ImageContain :src="config.webBackdrop" :width="400" height="100%" />
                </div>
                <div
                    class="forgot-form bg-body flex flex-col justify-center px-10 py-10 md:w-[400px] w-[375px] flex-none mx-auto"
                >
                    <div class="text-center text-3xl font-medium mb-4">{{ config.webName }}</div>
                    <div class="text-center text-sm text-gray-500 mb-8">
                        忘记密码？请输入注册邮箱重置
                    </div>

                    <el-form
                        ref="formRef"
                        :model="formData"
                        :rules="rules"
                        size="large"
                        autoComplete="off"
                    >
                        <!-- 步骤1：邮箱 -->
                        <div v-show="step === 1">
                            <el-form-item prop="email">
                                <el-input
                                    v-model.trim="formData.email"
                                    type="email"
                                    autocomplete="off"
                                    placeholder="请输入注册邮箱"
                                    tabindex="1"
                                >
                                    <template #prepend>
                                        <icon name="el-icon-Message" />
                                    </template>
                                </el-input>
                            </el-form-item>
                            <el-button
                                type="primary"
                                size="large"
                                tabindex="2"
                                :loading="isLockSend"
                                :disabled="countdown > 0"
                                class="w-full"
                                @click="lockSendCode"
                            >
                                {{ countdown > 0 ? `重新发送(${countdown}s)` : '发送验证码' }}
                            </el-button>
                        </div>

                        <!-- 步骤2：验证码 + 新密码 -->
                        <div v-show="step === 2">
                            <el-form-item prop="code">
                                <el-input
                                    v-model.trim="formData.code"
                                    placeholder="请输入6位验证码"
                                    maxlength="6"
                                    tabindex="1"
                                >
                                    <template #prepend>
                                        <icon name="el-icon-ChatDotRound" />
                                    </template>
                                </el-input>
                            </el-form-item>
                            <el-form-item prop="password">
                                <el-input
                                    ref="passwordRef"
                                    v-model="formData.password"
                                    show-password
                                    placeholder="请输入新密码（6-32位）"
                                    tabindex="2"
                                >
                                    <template #prepend>
                                        <icon name="el-icon-Lock" />
                                    </template>
                                </el-input>
                            </el-form-item>
                            <el-form-item prop="confirmPassword">
                                <el-input
                                    v-model="formData.confirmPassword"
                                    show-password
                                    placeholder="请再次输入新密码"
                                    tabindex="3"
                                    @keyup.enter="lockResetFn"
                                >
                                    <template #prepend>
                                        <icon name="el-icon-Lock" />
                                    </template>
                                </el-input>
                            </el-form-item>
                            <el-button
                                type="primary"
                                size="large"
                                tabindex="4"
                                :loading="isLockReset"
                                class="w-full"
                                @click="lockResetFn"
                            >
                                重置密码
                            </el-button>
                        </div>
                    </el-form>

                    <div class="mt-5 text-center">
                        <el-button link type="primary" @click="goLogin">返回登录</el-button>
                    </div>
                </div>
            </div>
        </div>
        <LayoutFooter />
    </div>
</template>

<script lang="ts" setup>
import { ref, reactive, computed, onUnmounted } from 'vue'
import type { FormInstance } from 'element-plus'
import { ElMessage } from 'element-plus'
import LayoutFooter from '@/layout/components/footer.vue'
import ImageContain from '@/components/image-contain/index.vue'
import useAppStore from '@/stores/modules/app'
import { forgotPwdSendCode, forgotPwdReset } from '@/api/user'
import { encryptPassword } from '@/utils/util'
import { PageEnum } from '@/enums/pageEnum'
import { useRouter } from 'vue-router'
import { useLockFn } from '@/hooks/useLockFn'

const formRef = ref<FormInstance>()
const passwordRef = ref()
const appStore = useAppStore()
const router = useRouter()
const config = computed(() => appStore.config)

// 步骤：1=输入邮箱发验证码，2=输入验证码+新密码
const step = ref(1)

// 倒计时
const countdown = ref(0)
let timer: number | null = null
const startCountdown = () => {
    countdown.value = 60
    timer = window.setInterval(() => {
        countdown.value--
        if (countdown.value <= 0 && timer !== null) {
            clearInterval(timer)
            timer = null
        }
    }, 1000)
}
onUnmounted(() => {
    if (timer !== null) clearInterval(timer)
})

const formData = reactive({
    email: '',
    code: '',
    password: '',
    confirmPassword: ''
})

const validateConfirmPwd = (rule: any, value: string, callback: Function) => {
    if (value === '') {
        callback(new Error('请再次输入新密码'))
    } else if (value !== formData.password) {
        callback(new Error('两次输入的密码不一致'))
    } else {
        callback()
    }
}

const rules = computed(() => {
    if (step.value === 1) {
        return {
            email: [
                { required: true, message: '请输入邮箱', trigger: ['blur'] },
                { type: 'email', message: '请输入正确的邮箱格式', trigger: ['blur'] }
            ]
        }
    }
    return {
        code: [
            { required: true, message: '请输入验证码', trigger: ['blur'] },
            { len: 6, message: '验证码为6位数字', trigger: ['blur'] }
        ],
        password: [
            { required: true, message: '请输入新密码', trigger: ['blur'] },
            { min: 6, max: 32, message: '密码长度为6-32位', trigger: ['blur'] }
        ],
        confirmPassword: [{ required: true, validator: validateConfirmPwd, trigger: ['blur'] }]
    }
})

// 发送验证码
const handleSendCode = async () => {
    await formRef.value?.validate()
    await forgotPwdSendCode({ email: formData.email })
    step.value = 2
    startCountdown()
}
const { isLock: isLockSend, lockFn: lockSendCode } = useLockFn(handleSendCode)

// 重置密码
const handleReset = async () => {
    await formRef.value?.validate()
    await forgotPwdReset({
        email: formData.email,
        code: formData.code,
        password: encryptPassword(formData.password)
    })
    ElMessage.success('密码重置成功，即将跳转到登录页')
    setTimeout(() => {
        router.push(PageEnum.LOGIN)
    }, 3000)
}
const { isLock: isLockReset, lockFn: lockResetFn } = useLockFn(handleReset)

const goLogin = () => {
    router.push(PageEnum.LOGIN)
}
</script>

<style lang="scss" scoped>
.forgot-wd {
    background-color: #f8f8f8;
    background-repeat: no-repeat;
    background-size: cover;
    height: 100vh;

    .forgot-card {
        height: 400px;
    }
}
</style>
