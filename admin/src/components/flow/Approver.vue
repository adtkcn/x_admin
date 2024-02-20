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
        <template #header="{ close }">
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

                <el-button class="publish-btn" @click="close">关闭</el-button>
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
            <Diagram
                ref="processDesign"
                v-show="activeStep === 'processDesign'"
                tabName="processDesign"
                :fieldList="fieldList"
                :conf="mockData.flowProcessData"
            ></Diagram>

            <AdvancedSetting
                ref="advancedSetting"
                :conf="mockData.advancedSetting"
                v-show="activeStep === 'advancedSetting'"
            />
        </div>
    </el-dialog>
</template>

<script>
import XForm from './XForm/index.vue'
import Diagram from './flowEdit/Diagram.vue'

import BasicSetting from './BasicSetting/index.vue'
import AdvancedSetting from './AdvancedSetting/index.vue'

const beforeUnload = function (e) {
    const confirmationMessage = '离开网站可能会丢失您编辑得内容'
    ;(e || window.event).returnValue = confirmationMessage // Gecko and Trident
    return confirmationMessage // Gecko and WebKit
}
export default {
    name: 'Approver',
    components: {
        // Process,
        Diagram,
        // DynamicForm,
        XForm,
        BasicSetting,
        AdvancedSetting
    },
    props: {
        title: {
            type: String,
            default: '流程配置'
        },
        save: {
            type: Function,
            default: () => {}
        }
    },
    data() {
        return {
            dialogVisible: false,

            activeStep: 'basicSetting', // 激活的步骤面板
            mockData: {
                basicSetting: {},
                flowFormData: {},
                flowProcessData: {},
                advancedSetting: {}
            },
            fieldList: [],
            steps: [
                { label: '基础设置', key: 'basicSetting' },
                { label: '表单设计', key: 'formDesign' },
                { label: '流程设计', key: 'processDesign' }
                // { label: "高级设置", key: "advancedSetting" },
            ]
        }
    },
    beforeRouteEnter(to, from, next) {
        window.addEventListener('beforeunload', beforeUnload)
        next()
    },
    beforeRouteLeave(to, from, next) {
        window.removeEventListener('beforeunload', beforeUnload)
        next()
    },
    computed: {},
    mounted() {},
    methods: {
        reset() {
            this.mockData = {}
            this.activeStep = 'basicSetting'
            this.fieldList = []
        },
        open(data) {
            this.reset()
            console.log('data', data)
            if (data) {
                this.mockData = { ...data }
            } else {
                this.mockData = {
                    basicSetting: {},
                    flowFormData: {},
                    flowProcessData: {},
                    advancedSetting: {}
                }
            }

            this.dialogVisible = true
        },
        async changeSteps(item) {
            const fieldList = this.$refs.formDesign.getFieldWidgets()

            this.fieldList = fieldList

            this.activeStep = item.key
        },
        publish() {
            const getCmpData = (name) => {
                return this.$refs[name].getData()
            }
            // basicSetting  formDesign processDesign 返回的是Promise 因为要做校验
            // advancedSetting返回的就是值
            const p1 = getCmpData('basicSetting')
            const p2 = getCmpData('formDesign')
            const p3 = getCmpData('processDesign')
            Promise.all([p1, p2, p3])
                .then((res) => {
                    console.log('res', res)
                    const param = {
                        id: this?.mockData?.id,
                        basicSetting: res[0].formData,
                        flowFormData: res[1].formData,
                        flowProcessData: res[2].formData,
                        flowProcessDataList: res[2].treeToList

                        // advancedSetting: getCmpData("advancedSetting"),
                    }
                    console.log(param)
                    this.sendToServer(param)
                })
                .catch((err) => {
                    console.error(err)
                    err.target && (this.activeStep = err.target)
                    err.msg && this.$message.error(err.msg)
                })
        },

        sendToServer(param) {
            this.$notify({
                title: '数据已整合完成',
                message: '请在控制台中查看数据输出',
                position: 'bottom-right'
            })
            this.save(param)
                .then(() => {
                    this.dialogVisible = false
                })
                .catch((err) => {
                    err.message && this.$message.error(err.message)
                })
            console.log('配置数据', param)
        },
        exit() {
            this.$confirm('离开此页面您得修改将会丢失, 是否继续?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            })
                .then(() => {
                    window.history.back()
                })
                .catch(() => {})
        }
    }
}
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
    background: #f5f5f7;
}

.publish-btn {
    margin-right: 20px;
    // color: #3296fa;
    padding: 7px 20px;
    border-radius: 4px;
    font-size: 14px;
}
</style>
