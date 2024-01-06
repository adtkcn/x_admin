<template>
    <div class="login flex flex-col">
        <div class="flex-1 flex items-center justify-center">
            <div class="login-card flex rounded-md">
                <div class="flex-1 h-full hidden md:inline-block">
                    <image-contain :src="config.webBackdrop" :width="400" height="100%" />
                </div>
                <div
                    class="login-form bg-body flex flex-col justify-center px-10 py-10 md:w-[400px] w-[375px] flex-none mx-auto"
                >
                    <div class="text-center text-3xl font-medium mb-8">{{ config.webName }}</div>
                    <el-form ref="formRef" :model="formData" size="large" :rules="rules">
                        <el-form-item prop="username">
                            <el-input v-model.trim="formData.username" placeholder="请输入账号">
                                <template #prepend>
                                    <icon name="el-icon-User" />
                                </template>
                            </el-input>
                        </el-form-item>
                        <el-form-item prop="password">
                            <el-input
                                ref="passwordRef"
                                v-model="formData.password"
                                show-password
                                placeholder="请输入密码"
                                @keyup.enter="handleLogin"
                            >
                                <template #prepend>
                                    <icon name="el-icon-Lock" />
                                </template>
                            </el-input>
                        </el-form-item>
                    </el-form>
                    <div class="mb-5">
                        <el-checkbox v-model="remAccount" label="记住账号"></el-checkbox>
                    </div>
                    <el-button
                        type="primary"
                        size="large"
                        :loading="isLock"
                        @click="onShowCaptcha('blockPuzzle')"
                    >
                        登录
                    </el-button>

                    <!-- <button @click="onShowCaptcha('blockPuzzle')">滑块</button>
                    <button @click="onShowCaptcha('clickWord')">点击文字</button> -->
                    <Verify
                        mode="pop"
                        :captchaType="captchaType"
                        :imgSize="{ width: '400px', height: '200px' }"
                        ref="verify"
                        @success="handleSuccess"
                        @error="
                            (e) => {
                                console.log(e)
                            }
                        "
                        @ready="
                            (e) => {
                                console.log(e)
                            }
                        "
                    ></Verify>
                </div>
            </div>
        </div>
        <layout-footer />
    </div>
</template>

<script lang="ts" setup>
import { computed, onMounted, reactive, ref, shallowRef } from 'vue'
import type { InputInstance, FormInstance } from 'element-plus'
import LayoutFooter from '@/layout/components/footer.vue'
import useAppStore from '@/stores/modules/app'
import useUserStore from '@/stores/modules/user'
import cache from '@/utils/cache'
import { ACCOUNT_KEY } from '@/enums/cacheEnums'
import { PageEnum } from '@/enums/pageEnum'
import { useLockFn } from '@/hooks/useLockFn'

import Verify from '@/components/verifition/Verify.vue'

const verify = ref(null)
const captchaType = ref('')
const onShowCaptcha = (type) => {
    captchaType.value = type
    verify.value.show()
}
let verifition = null
const handleSuccess = (res) => {
    console.log(res)
    console.log('sucess')
    verifition = res
    lockLogin(res)
}

const passwordRef = shallowRef<InputInstance>()
const formRef = shallowRef<FormInstance>()
const appStore = useAppStore()
const userStore = useUserStore()
const route = useRoute()
const router = useRouter()
const remAccount = ref(false)
const config = computed(() => appStore.config)
const formData = reactive({
    username: '',
    password: ''
})
const rules = {
    username: [
        {
            required: true,
            message: '请输入账号',
            trigger: ['blur']
        }
    ],
    password: [
        {
            required: true,
            message: '请输入密码',
            trigger: ['blur']
        }
    ]
}

// 登录处理
const handleLogin = async (captchaInfo) => {
    console.log('captchaInfo', captchaInfo, { ...formData, ...captchaInfo })

    await formRef.value?.validate()
    // 记住账号，缓存
    cache.set(ACCOUNT_KEY, {
        remember: remAccount.value,
        username: remAccount.value ? formData.username : ''
    })
    await userStore.login({ ...formData, ...verifition })
    const {
        query: { redirect }
    } = route
    const path = typeof redirect === 'string' ? redirect : PageEnum.INDEX
    router.push(path)
}
const { isLock, lockFn: lockLogin } = useLockFn(handleLogin)

onMounted(() => {
    const value = cache.get(ACCOUNT_KEY)
    if (value?.remember) {
        remAccount.value = value.remember
        formData.username = value.username
    }
})
</script>

<style lang="scss" scoped>
.login {
    background-image: url('./images/login_bg.png');
    background-repeat: no-repeat;
    background-size: cover;
    height: 100vh;

    .login-card {
        height: 400px;
    }
}
</style>
