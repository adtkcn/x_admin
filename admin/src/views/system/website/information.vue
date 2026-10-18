<!-- 网站信息 -->
<template>
    <div class="website-information">
        <el-form ref="formRef" :rules="rules" :model="formData" label-width="120px">
            <el-card shadow="never" class="border-none!">
                <div class="text-xl font-medium mb-[20px]">后台设置</div>
                <el-form-item label="网站名称" prop="name">
                    <div class="w-80">
                        <el-input
                            v-model="formData.name"
                            placeholder="请输入网站名称"
                            maxlength="30"
                            show-word-limit
                        ></el-input>
                    </div>
                </el-form-item>
                <el-form-item label="网站图标" prop="favicon">
                    <div>
                        <material-picker v-model="formData.favicon" :limit="1" />
                        <div class="form-tips">建议尺寸：100*100像素，支持jpg，jpeg，png格式</div>
                    </div>
                </el-form-item>
                <el-form-item label="网站logo" prop="logo">
                    <div>
                        <material-picker v-model="formData.logo" :limit="1" />
                        <div class="form-tips">建议尺寸：200*200像素，支持jpg，jpeg，png格式</div>
                    </div>
                </el-form-item>
                <el-form-item label="登录页广告图" prop="backdrop">
                    <div>
                        <material-picker v-model="formData.backdrop" :limit="1" />
                        <div class="form-tips">建议尺寸：400*400像素，支持jpg，jpeg，png格式</div>
                    </div>
                </el-form-item>
            </el-card>
        </el-form>
        <footer-btns v-perms="['admin:setting:website:save']">
            <el-button type="primary" @click="handleSubmit">保存</el-button>
        </footer-btns>
    </div>
</template>

<script lang="ts" setup>
import { useTemplateRef, reactive } from 'vue'
import { getWebsite, setWebsite } from '@/api/setting/website'
import useAppStore from '@/stores/modules/app'
import feedback from '@/utils/feedback'
import type { FormInstance } from 'element-plus'
const formRef = useTemplateRef<FormInstance>('formRef')
const { getConfig } = useAppStore()

defineOptions({
    name: 'webInformation'
})
// 表单数据
const formData = reactive({
    name: '', // 网站名称
    favicon: '', // 网站图标
    logo: '', // 网站logo
    backdrop: '' // 登录页广告图
})

// 表单验证
const rules = {
    name: [
        {
            required: true,
            message: '请输入网站名称',
            trigger: ['blur']
        }
    ],
    favicon: [
        {
            required: true,
            message: '请选择网站图标',
            trigger: ['change']
        }
    ],
    logo: [
        {
            required: true,
            message: '请选择网站logo',
            trigger: ['change']
        }
    ],
    backdrop: [
        {
            required: true,
            message: '请选择登录页广告图',
            trigger: ['change']
        }
    ]
}

// 获取备案信息
const getData = async () => {
    try {
        const data = await getWebsite()
        for (const key in formData) {
            //@ts-ignore
            formData[key] = data[key]
        }
    } catch (error) {
        console.error('网站信息获取失败:', error)
    }
}

// 设置备案信息
const handleSubmit = async () => {
    try {
        await formRef.value?.validate()
        await setWebsite(formData)
        feedback.msgSuccess('操作成功')
        getConfig()
        getData()
    } catch (error) {
        console.error('网站信息保存失败:', error)
    }
}

getData()
</script>

<style lang="scss" scoped></style>
