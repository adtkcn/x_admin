<template>
    <div class="login flex flex-col">
        <div class="flex-1 flex items-center justify-center">
            <div class="login-card flex rounded-md">
                <div class="flex-1 h-full hidden md:inline-block">
                    <ImageContain :src="config.webBackdrop" :width="400" height="100%" />
                </div>
                <div
                    class="login-form bg-body flex flex-col justify-center px-10 py-10 md:w-[400px] w-[375px] flex-none mx-auto"
                >
                    <div class="text-center text-3xl font-medium mb-8">{{ config.webName }}</div>

                    <el-form ref="formRef" :model="formData" :rules="rules" autoComplete="off">
                        <el-form-item prop="email">
                            <el-input
                                v-model.trim="formData.email"
                                type="text"
                                autocomplete="off"
                                placeholder="请输入邮箱"
                                tabindex="1"
                            >
                                <template #prepend>
                                    <icon name="el-icon-User" />
                                </template>
                            </el-input>
                        </el-form-item>
                        <input
                            v-model.trim="formData.email"
                            type="text"
                            name="email-hide"
                            class="hide-input"
                        />
                        <input v-model="formData.password" type="password" class="hide-input" />
                        <el-form-item prop="password">
                            <el-input
                                v-model="formData.password"
                                tabindex="1"
                                show-password
                                placeholder="请输入密码"
                                @keyup.enter="handleLoginClick"
                            >
                                <template #prepend>
                                    <icon name="el-icon-Lock" />
                                </template>
                            </el-input>
                        </el-form-item>
                    </el-form>

                    <el-button
                        type="primary"
                        tabindex="1"
                        :loading="isLock"
                        @click="handleLoginClick"
                    >
                        登录
                    </el-button>

                    <div class="mt-3 text-right">
                        <el-button link type="primary" @click="goForgotPwd">忘记密码？</el-button>
                    </div>

                    <Verify
                        mode="pop"
                        captchaType="clickWord"
                        :imgSize="{ width: 400, height: 200 }"
                        ref="verifyRef"
                        @success="handleCaptchaSuccess"
                    ></Verify>
                </div>
            </div>
        </div>
        <layout-footer />
    </div>
</template>

<script lang="ts" setup>
import { computed, onMounted, reactive, useTemplateRef, defineAsyncComponent } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import type { FormInstance } from 'element-plus'
import LayoutFooter from '@/layout/components/footer.vue'
import useAppStore from '@/stores/modules/app'
import useUserStore from '@/stores/modules/user'
import cache from '@/utils/cache'
import { ACCOUNT_KEY } from '@/enums/cacheEnums'
import { PageEnum } from '@/enums/pageEnum'
import { useLockFn } from '@/hooks/useLockFn'
import { encryptPassword } from '@/utils/util'

defineOptions({
    name: 'AccountLogin'
})

const Verify = defineAsyncComponent(() => import('@/components/verify/Verify.vue'))
const ImageContain = defineAsyncComponent(() => import('@/components/image-contain/index.vue'))

const formRef = useTemplateRef<FormInstance>('formRef')
const verifyRef = useTemplateRef<InstanceType<typeof Verify>>('verifyRef')
const appStore = useAppStore()
const userStore = useUserStore()
const route = useRoute()
const router = useRouter()
const config = computed(() => appStore.config)

const formData = reactive({
    email: 'x@qq.com',
    password: '123456'
})

const rules = {
    email: [{ required: true, message: '请输入邮箱', trigger: ['blur'] }],
    password: [{ required: true, message: '请输入密码', trigger: ['blur'] }]
}

// 点击登录：先校验表单，通过后再弹出验证码
const handleLoginClick = async () => {
    await formRef.value?.validate()
    verifyRef.value?.show()
}

// 验证码通过后执行登录
const handleCaptchaSuccess = (captcha: Record<string, any>) => {
    lockLogin(captcha)
}

// 登录处理
const handleLogin = async (captcha: Record<string, any>) => {
    await userStore.login({
        email: formData.email,
        password: encryptPassword(formData.password),
        ...captcha
    })

    cache.set(ACCOUNT_KEY, { email: formData.email })

    const path = typeof route.query.redirect === 'string' ? route.query.redirect : PageEnum.INDEX
    router.push(path)
}
const { isLock, lockFn: lockLogin } = useLockFn(handleLogin)

const goForgotPwd = () => {
    router.push(PageEnum.FORGOT_PASSWORD)
}

onMounted(() => {
    const value: { email?: string } | null = cache.get(ACCOUNT_KEY)
    formData.email = value?.email || ''
})
</script>

<style lang="scss" scoped>
.login {
    background-color: #f8f8f8;
    background-repeat: no-repeat;
    background-size: cover;
    height: 100vh;
    background-image: url(./images/login_bg.png);

    .login-card {
        height: 400px;
    }
}

.hide-input {
    position: absolute;
    height: 0;
    width: 0;
    line-height: 0;
    border: 0;
    opacity: 0;
    overflow: hidden;
}
</style>
