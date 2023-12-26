<template>
    <el-dialog
        v-model="dialogVisible"
        :show-close="false"
        :fullscreen="false"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
        :destroy-on-close="true"
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
            <el-button v-if="applyDetail.status == 1" type="primary" @click="onSubmit">
                确定
            </el-button>
        </template>
    </el-dialog>
</template>

<script setup>
import { ref } from 'vue'
import 'vform3-builds/dist/designer.style.css' //引入VForm3样式

const formJson = ref({})
const formData = ref({})
const optionData = reactive({})
const vFormRef = ref(null)

const dialogVisible = ref(false)
const applyDetail = ref({})

const props = defineProps({
    save: {
        type: Function,
        default: () => {}
    }
})
function open(row, form_json, form_data) {
    applyDetail.value = row

    formData.value = form_data
    formJson.value = form_json
    console.log('open')
    dialogVisible.value = true
}
function closeFn() {
    dialogVisible.value = false
    applyDetail.value = {}
    formData.value = {}
    formJson.value = {}
}
function onSubmit() {
    vFormRef.value.getFormData().then((formData) => {
        console.log('formData', formData)
        props
            .save(applyDetail.value?.id, formData)
            .then(() => {
                closeFn()
            })
            .catch(() => {})
    })
}
defineExpose({
    open
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
