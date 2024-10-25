<template>
    <el-dialog
        v-model="dialogVisible"
        append-to-body
        :mask="false"
        :show-close="false"
        :fullscreen="true"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
        :destroy-on-close="true"
        top="1px"
        style="height: 100vh"
        class="flow-config-dialog"
    >
        <template #header>
            <header class="page__header">
                <div class="page-actions">
                    {{ title }}
                </div>
                <div class="step-tab">
                    <div
                        v-for="(item, index) in steps"
                        :key="index"
                        class="step"
                        :class="[activeStep == item.key ? 'active' : '']"
                        @click="changeSteps(item)"
                    >
                        <span class="step-index">{{ index + 1 }}</span>
                        {{ item.label }}
                    </div>
                </div>

                <el-button class="publish-btn" @click="exit">关闭</el-button>
                <el-button class="publish-btn" type="primary" @click="publish">发布</el-button>
            </header>
        </template>

        <div class="page__content" v-if="mockData">
            <BasicSetting
                ref="basicSetting"
                :conf="mockData.basicSetting"
                v-show="activeStep === 'basicSetting'"
                tabName="basicSetting"
            />
            <XForm
                ref="formDesign"
                :conf="mockData.flowFormData"
                v-show="activeStep === 'formDesign'"
                tabName="formDesign"
            />
            <FlowEdit
                ref="flowEdit"
                v-show="activeStep === 'flowEdit'"
                tabName="flowEdit"
                :fieldList="fieldList"
                :conf="mockData.flowProcessData"
            ></FlowEdit>
        </div>
    </el-dialog>
</template>

<script setup lang="ts">
import { ref, shallowRef, watch } from 'vue'
import XForm from './XForm/index.vue'
import FlowEdit from './flowEdit/index.vue'
import BasicSetting from './BasicSetting/index.vue'

import feedback from '@/utils/feedback'

const beforeUnload = function (e) {
    e.preventDefault()

    e.returnValue = '离开网站可能会丢失您编辑得内容' // Gecko and Trident
    return false // Gecko and WebKit
}
defineOptions({
    name: 'Approver'
})
const props = defineProps({
    title: {
        type: String,
        default: '流程配置'
    },
    save: {
        type: Function,
        default: () => {}
    }
})

const basicSetting = shallowRef<InstanceType<typeof BasicSetting>>()
const formDesign = shallowRef<InstanceType<typeof XForm>>()
const flowEdit = shallowRef<InstanceType<typeof FlowEdit>>()
const dialogVisible = ref(false)
const activeStep = ref('basicSetting')
const mockData = ref({
    id: '',
    basicSetting: {},
    flowFormData: {},
    flowProcessData: {}
})
const fieldList = ref([])
const steps = [
    { label: '基础设置', key: 'basicSetting' },
    { label: '表单设计', key: 'formDesign' },
    { label: '流程设计', key: 'flowEdit' }
]
watch(
    () => dialogVisible.value,
    (val) => {
        if (!val) {
            window.removeEventListener('beforeunload', beforeUnload)
        } else {
            window.addEventListener('beforeunload', beforeUnload)
        }
    }
)
function reset() {
    mockData.value = {
        id: '',
        basicSetting: {},
        flowFormData: {},
        flowProcessData: {}
    }
    activeStep.value = 'basicSetting'
    fieldList.value = []
}
function open(data) {
    reset()
    console.log('data', data)
    if (data) {
        mockData.value = { ...data }
    } else {
        mockData.value = {
            id: '',
            basicSetting: {},
            flowFormData: {},
            flowProcessData: {}
        }
    }

    dialogVisible.value = true
}
function changeSteps(item) {
    fieldList.value = formDesign.value.getFieldWidgets()
    console.log('fieldList', fieldList.value)

    activeStep.value = item.key
}
// 发布
function publish() {
    const p1 = basicSetting.value.getData()
    const p2 = formDesign.value.getData()
    const p3 = flowEdit.value.getData()
    Promise.all([p1, p2, p3])
        .then((res) => {
            console.log('res', res)
            const param = {
                id: mockData.value?.id,
                basicSetting: res[0].formData,
                flowFormData: res[1].formData,
                flowProcessData: res[2].formData,
                flowProcessDataList: res[2].treeToList
            }
            console.log(param)
            sendToServer(param)
        })
        .catch((err) => {
            console.error(err)
            err.target && (activeStep.value = err.target)
            err.message && feedback.msgError(err.message)
        })
}

function sendToServer(param) {
    feedback.notify('数据已整合完成')
    props
        .save(param)
        .then(() => {
            dialogVisible.value = false
        })
        .catch((err) => {
            err.message && feedback.msgError(err.message)
        })
    console.log('配置数据', param)
}
function exit() {
    feedback
        .confirm('离开此页面您得修改将会丢失, 是否继续?')
        .then(() => {
            dialogVisible.value = false
        })
        .catch(() => {})
}
defineExpose({
    open
})
</script>
<style>
.flow-config-dialog {
    padding: 0;
}
.flow-config-dialog .el-dialog__header {
    font-size: var(--el-font-size-large);
    padding: 0;
    margin: 0;
}
</style>
<style lang="scss" scoped>
$header-height: 54px;

.page__header {
    position: fixed;
    z-index: 100;
    width: 100%;
    height: $header-height;
    display: flex;
    align-items: center;
    justify-content: center;
    // justify-content: space-between;
    box-sizing: border-box;
    color: white;
    background: #3296fa;
    font-size: 14px;
    user-select: none;

    .page-actions {
        line-height: $header-height;

        height: 100%;
        padding-left: 20px;
        padding-right: 20px;
    }

    .step-tab {
        flex: 1;
        display: flex;
        justify-content: center;
        height: 100%;
        position: relative;

        > .step {
            width: 140px;
            line-height: $header-height;
            padding-left: 30px;
            padding-right: 30px;
            cursor: pointer;
            position: relative;

            &.active > .step-index {
                background: white;
                color: #4483f2;
            }

            > .step-index {
                display: inline-block;
                width: 18px;
                height: 18px;
                border: 1px solid #fff;
                border-radius: 8px;
                line-height: 18px;
                text-align: center;
                box-sizing: border-box;
            }
        }
    }
}

// width: 100vw;
// height: 100vh;
// padding-top: $header-height;

.page__content {
    padding-top: $header-height;
    height: 100%;

    box-sizing: border-box;
    // background: #f5f5f7;
}

.publish-btn {
    margin-right: 20px;
    // color: #3296fa;
    padding: 7px 20px;
    border-radius: 4px;
    font-size: 14px;
}
</style>
