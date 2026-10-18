<!-- 网站设置（网站信息 + 备案信息） -->
<template>
    <div class="website-setting">
        <el-card shadow="never" class="border-none!">
            <el-tabs v-model="activeTab" class="website-setting__tabs">
                <el-tab-pane label="网站信息" name="information">
                    <el-form ref="formRef" :rules="rules" :model="formData" label-width="120px">
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
                    </el-form>
                </el-tab-pane>

                <el-tab-pane label="备案信息" name="filing">
                    <el-form class="ls-form" label-width="100px">
                        <template v-for="(item, index) in filingData" :key="index">
                            <del-wrap
                                class="mb-4"
                                :show-close="filingData.length > 1"
                                @close="handleFilingDelete(index)"
                            >
                                <div class="filing-item bg-fill-lighter">
                                    <el-form-item label="显示内容" prop="name">
                                        <div class="w-80">
                                            <el-input
                                                v-model="item.name"
                                                placeholder="请输入内容 ，例如： ICP备案号"
                                            />
                                        </div>
                                    </el-form-item>
                                    <el-form-item label="跳转链接" prop="link">
                                        <div class="w-80">
                                            <el-input
                                                v-model="item.link"
                                                placeholder="请输入链接，例如：http://www.beian.gov.cn"
                                            />
                                            <div class="form-tips">跳转链接不设置，则不跳转</div>
                                        </div>
                                    </el-form-item>
                                </div>
                            </del-wrap>
                        </template>
                        <el-button type="primary" plain @click="handleFilingAdd">
                            <icon name="el-icon-Plus" />
                            添加
                        </el-button>
                    </el-form>
                </el-tab-pane>
            </el-tabs>
        </el-card>

        <footer-btns v-if="activeTab === 'information'" v-perms="['admin:setting:website:save']">
            <el-button type="primary" @click="handleSubmit">保存</el-button>
        </footer-btns>
        <footer-btns v-else v-perms="['admin:setting:copyright:save']">
            <el-button type="primary" @click="handleFilingSubmit">保存</el-button>
        </footer-btns>
    </div>
</template>

<script lang="ts" setup>
import { useTemplateRef, reactive, ref } from 'vue'
import {
    getWebsite,
    setWebsite,
    getCopyright,
    setCopyright,
    type type_setting_copyright_item
} from '@/api/setting/website'
import useAppStore from '@/stores/modules/app'
import feedback from '@/utils/feedback'
import type { FormInstance } from 'element-plus'
const formRef = useTemplateRef<FormInstance>('formRef')
const { getConfig } = useAppStore()

defineOptions({
    name: 'webInformation'
})

const activeTab = ref('information')

// 网站信息表单数据
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

// 备案信息表单数据
const filingData = ref<type_setting_copyright_item[]>([{ name: '', link: '' }])

// 获取网站信息
const getWebsiteData = async () => {
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

// 设置网站信息
const handleSubmit = async () => {
    try {
        await formRef.value?.validate()
        await setWebsite(formData)
        feedback.msgSuccess('操作成功')
        getConfig()
        getWebsiteData()
    } catch (error) {
        console.error('网站信息保存失败:', error)
    }
}

// 获取备案信息
const getFilingData = async () => {
    try {
        const data = await getCopyright()
        if (!data.length) return
        filingData.value = data
    } catch (error) {
        console.error('版权备案信息获取失败:', error)
    }
}

const handleFilingAdd = () => {
    filingData.value.push({ name: '', link: '' })
}

const handleFilingDelete = (index: number) => {
    if (filingData.value.length <= 1) {
        return feedback.msgError('至少保留一个')
    }
    filingData.value.splice(index, 1)
}

// 设置备案信息
const handleFilingSubmit = async () => {
    try {
        await setCopyright(filingData.value)
        feedback.msgSuccess('操作成功')
        getFilingData()
    } catch (error) {
        console.error('版权备案信息保存失败:', error)
    }
}

getWebsiteData()
getFilingData()
</script>

<style lang="scss" scoped>
.website-setting {
    // 底部固定保存栏高 60px，预留空间避免遮挡内容
    padding-bottom: 20px;

    &__tabs {
        :deep(.el-tabs__header) {
            margin-bottom: 24px;
        }
        :deep(.el-tabs__item) {
            font-size: 15px;
        }
    }
    .ls-form{
        padding:10px 10px;
    }
    // 备案条目卡片：统一内边距与圆角，去掉冗余的表单项底边距
    .filing-item {
        padding: 16px 16px 0;
        border-radius: 8px;
        background-color: #f2f2f2;

        :deep(.el-form-item) {
            margin-bottom: 16px;
        }
    }
}
</style>
