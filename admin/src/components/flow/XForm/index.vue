<template>
    <v-form-designer ref="designerRef" :designer-config="designerConfig"></v-form-designer>
</template>

<script setup>
import { ref } from 'vue'
import 'vform3-builds/dist/designer.style.css' //引入VForm3样式
const designerRef = ref(null)
function setData(json) {
    console.log('setFormJson', json)
    designerRef.value.setFormJson(json)
}
function getFieldWidgets() {
    const fieldList = designerRef.value.getFieldWidgets()
    console.log('getFieldWidgets', fieldList)
    return fieldList
}
function getData() {
    return new Promise((resolve, reject) => {
        try {
            const jsonData = designerRef.value.getFormJson()
            console.log('jsonData', jsonData)
            resolve({ formData: jsonData })
        } catch (error) {
            reject(error)
        }
    })
}
const props = defineProps({
    conf: {
        type: Object,
        default: () => {}
    }
})
const designerConfig = {
    languageMenu: false,
    externalLink: false,
    formTemplates: false
}
onMounted(() => {
    setData(props.conf)
})
defineExpose({
    setData,
    getData,
    getFieldWidgets
})
</script>

<style lang="scss">
.el-range-editor.el-input__wrapper {
    box-sizing: border-box;
}
</style>
