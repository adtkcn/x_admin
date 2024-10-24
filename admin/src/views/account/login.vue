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

                    <el-form
                        ref="formRef"
                        :model="formData"
                        size="large"
                        :rules="rules"
                        autoComplete="off"
                    >
                        <el-form-item prop="username">
                            <el-input
                                v-model.trim="formData.username"
                                type="text"
                                autocomplete="off"
                                placeholder="请输入账号"
                                tabindex="1"
                            >
                                <template #prepend>
                                    <icon name="el-icon-User" />
                                </template>
                            </el-input>
                        </el-form-item>
                        <input
                            v-model.trim="formData.username"
                            type="text"
                            name="username-hide"
                            tabindex="-1"
                            style="
                                position: absolute;
                                height: 0;
                                width: 0;
                                line-height: 0;
                                border: 0;
                                opacity: 0;
                                overflow: hidden;
                            "
                        />
                        <input
                            v-model="formData.password"
                            type="password"
                            tabindex="-1"
                            style="
                                position: absolute;
                                height: 0;
                                width: 0;
                                line-height: 0;
                                border: 0;
                                opacity: 0;
                                overflow: hidden;
                            "
                        />
                        <el-form-item prop="password">
                            <el-input
                                ref="passwordRef"
                                v-model="formData.password"
                                tabindex="1"
                                show-password
                                placeholder="请输入密码"
                                @keyup.enter="onShowCaptcha"
                            >
                                <template #prepend>
                                    <icon name="el-icon-Lock" />
                                </template>
                            </el-input>
                        </el-form-item>
                    </el-form>
                    <!-- <div class="mb-5">
                        <el-checkbox v-model="remAccount" label="记住账号"></el-checkbox>
                    </div> -->
                    <el-button
                        type="primary"
                        size="large"
                        tabindex="1"
                        :loading="isLock"
                        @click="onShowCaptcha"
                    >
                        登录
                    </el-button>

                    <Verify
                        mode="pop"
                        captchaType="clickWord"
                        :imgSize="{ width: 400, height: 200 }"
                        ref="verifyRef"
                        @success="handleSuccess"
                    ></Verify>
                </div>
            </div>
        </div>
        <layout-footer />
    </div>
</template>

<script lang="ts" setup>
import { computed, onMounted, reactive, shallowRef } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import type { InputInstance, FormInstance } from 'element-plus'
import LayoutFooter from '@/layout/components/footer.vue'
import useAppStore from '@/stores/modules/app'
import useUserStore from '@/stores/modules/user'
import cache from '@/utils/cache'
import { ACCOUNT_KEY } from '@/enums/cacheEnums'
import { PageEnum } from '@/enums/pageEnum'
import { useLockFn } from '@/hooks/useLockFn'
import { encryptPassword } from '@/utils/util'
import Verify from '@/components/verify/Verify.vue'
import ImageContain from '@/components/image-contain/index.vue'
// const verifyRef = ref(null)
const verifyRef = shallowRef<InstanceType<typeof Verify>>()
const onShowCaptcha = () => {
    verifyRef.value.show()
}
let verifyInfo = null
const handleSuccess = (res) => {
    console.log(res)
    verifyInfo = res
    lockLogin(res)
}

const passwordRef = shallowRef<InputInstance>()
const formRef = shallowRef<FormInstance>()
const appStore = useAppStore()
const userStore = useUserStore()
const route = useRoute()
const router = useRouter()
// const remAccount = ref(false)
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
    console.log('captchaInfo', {
        username: formData.username,
        password: encryptPassword(formData.password),
        ...captchaInfo
    })

    await formRef.value?.validate()
    // 记住账号，缓存
    cache.set(ACCOUNT_KEY, {
        username: formData.username
    })
    await userStore.login({
        username: formData.username,
        password: encryptPassword(formData.password),
        ...verifyInfo
    })
    const {
        query: { redirect }
    } = route
    const path = typeof redirect === 'string' ? redirect : PageEnum.INDEX
    router.push(path)
}
const { isLock, lockFn: lockLogin } = useLockFn(handleLogin)

onMounted(() => {
    const value = cache.get(ACCOUNT_KEY)

    formData.username = value?.username
})
</script>

<style lang="scss" scoped>
.login {
    // background-image: url('./images/login_bg.png');
    background-color: #f8f8f8;
    background-repeat: no-repeat;
    background-size: cover;
    height: 100vh;

    .login-card {
        height: 400px;
    }
}
</style>
