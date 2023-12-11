<template>
    <el-dialog
        v-model="dialogVisible"
        :show-close="false"
        :fullscreen="false"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
        :destroy-on-close="true"
        top="1px"
        class="flow-config-dialog"
    >
        <v-form-render
            :form-json="formJson"
            :form-data="formData"
            :option-data="optionData"
            ref="vFormRef"
        >
        </v-form-render>

        <template #footer>
            <span class="dialog-footer">
                <el-button @click="dialogVisible = false">Cancel</el-button>
                <el-button type="primary" @click="getData"> Confirm </el-button>
            </span>
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
const rowId = ref(0)

const props = defineProps({
    save: {
        type: Function,
        default: () => {}
    }
})
function open(id, form_json, form_data) {
    rowId.value = id
    formData.value = form_data
    formJson.value = form_json
    console.log('open')
    dialogVisible.value = true
}
function closeFn() {
    dialogVisible.value = false
    rowId.value = 0
    formData.value = {}
    formJson.value = {}
}
function getData() {
    vFormRef.value.getFormData().then((formData) => {
        console.log('formData', formData)
        props
            .save(rowId.value, formData)
            .then(() => {
                closeFn()
            })
            .catch(() => {})
    })
}
defineExpose({
    getData,
    open
})
</script>

<style lang="scss">
// body {
//   margin: 0; /* 如果页面出现垂直滚动条，则加入此行CSS以消除之 */
// }
</style>
