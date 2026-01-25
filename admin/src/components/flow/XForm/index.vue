<template>
    <!-- <v-form-designer ref="designerRef" :designer-config="designerConfig"></v-form-designer> -->
    <fc-designer ref="designerRef" :config="config" />
</template>

<script setup lang="ts">
import { onMounted, useTemplateRef } from 'vue'
import FcDesigner from '@form-create/designer'

const designerRef = useTemplateRef<InstanceType<typeof FcDesigner>>('designerRef')
const config = {}
function setData(json: any[]) {
    console.log('setFormJson', json)
    json && designerRef.value.setRule(json)

    // 使用 getJson和 getOptionsJson导出数据。
    // 使用 setRule和 setOptions方法回显数据。
}
function getFieldWidgets() {
    const description = designerRef.value.getDescription()
    console.log('description', description)
    const fieldList: { id: string; name: string }[] = []
    function deepChild(item: any) {
        if (item.children) {
            item.children.forEach((child: any) => {
                deepChild(child)
            })
        } else {
            fieldList.push({
                id: item.field,
                name: item.title
            })
        }
    }
    description &&
        description.forEach((item) => {
            deepChild(item)
        })

    console.log('getFieldWidgets', fieldList)
    return fieldList
}
function getData() {
    return new Promise<{
        formData: any
    }>((resolve, reject) => {
        try {
            const jsonData = designerRef.value.getRule()
            const getOption = designerRef.value.getOption()
            const getDescription = designerRef.value.getDescription()
            // 表单组件的层级结构数据
            const getFormDescription = designerRef.value.getFormDescription()
            console.log('jsonData', jsonData)
            console.log('getOption', getOption)
            console.log('getDescription', getDescription)
            console.log('getFormDescription', getFormDescription)
            resolve({ formData: jsonData })
        } catch (error) {
            reject(error)
        }
    })
}
const props = defineProps({
    conf: {
        type: Array,
        default: () => []
    }
})

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
