<template>
    <el-dialog
        v-model="dialogVisible"
        :fullscreen="false"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
        :destroy-on-close="true"
        draggable
        :title="applyDetail.flowName"
    >
        <v-form-render
            :form-json="formJson"
            :form-data="formData"
            :option-data="optionData"
            ref="vFormRef"
        >
        </v-form-render>

        <template #footer>
            <el-button @click="dialogVisible = false">关闭</el-button>
            <el-button v-if="historyDetail.passStatus == 1" type="warning" @click="onBack">
                驳回
            </el-button>
            <el-button v-if="historyDetail.passStatus == 1" type="primary" @click="onSubmit">
                确定
            </el-button>
        </template>
    </el-dialog>
</template>

<script setup>
import { ref, reactive } from 'vue'
import 'vform3-builds/dist/designer.style.css' //引入VForm3样式
// import { flow_apply_detail } from '@/api/flow/flow_apply'

const formJson = ref({})
const formData = ref({})
const optionData = reactive({})
const vFormRef = ref(null)

const dialogVisible = ref(false)
const applyDetail = ref({})
const historyDetail = ref({})

const props = defineProps({
    save: {
        type: Function,
        default: () => {}
    }
})
const emit = defineEmits(['back'])

function open(row, history, form_json, form_data) {
    applyDetail.value = row
    historyDetail.value = history
    formData.value = form_data
    formJson.value = form_json
    console.log('open')

    dialogVisible.value = true
}

function disableWidgets(widgetNames) {
    vFormRef.value.disableWidgets(widgetNames)
}
function hideWidgets(widgetNames) {
    vFormRef.value.hideWidgets(widgetNames)
}
function closeFn() {
    dialogVisible.value = false
    applyDetail.value = {}
    formData.value = {}
    formJson.value = {}
    historyDetail.value = {}
}
function onBack() {
    emit('back', historyDetail.value)
}
function onSubmit() {
    vFormRef.value.getFormData().then((formData) => {
        console.log('formData', formData)
        props
            .save(historyDetail.value?.id, formData)
            .then(() => {
                closeFn()
            })
            .catch(() => {})
    })
}
defineExpose({
    open,
    disableWidgets,
    hideWidgets
})
</script>

<style lang="scss">
// body {
//   margin: 0; /* 如果页面出现垂直滚动条，则加入此行CSS以消除之 */
// }
.el-range-editor.el-input__wrapper {
    box-sizing: border-box;
}
</style>
