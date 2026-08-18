<template>
    <el-dialog
        v-model="dialogVisible"
        :fullscreen="false"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
        :destroy-on-close="true"
        draggable
        :title="applyDetail.flow_name"
    >
        <FormCreate
            v-if="dialogVisible"
            :rule="formJson"
            v-model="formData"
            v-model:api="api"
            :option="options"
        ></FormCreate>

        <template #footer>
            <el-button @click="dialogVisible = false">关闭</el-button>
            <el-button v-if="historyDetail.pass_status == 1" type="warning" @click="onBack">
                驳回
            </el-button>
            <el-button v-if="historyDetail.pass_status == 1" type="primary" @click="onSubmit">
                确定
            </el-button>
        </template>
    </el-dialog>
</template>

<script setup lang="ts">
import { ref, defineAsyncComponent } from 'vue'
const FormCreate = defineAsyncComponent(() => import('@form-create/element-ui'))
import type { Api } from '@form-create/element-ui'

// import { flow_apply_detail } from '@/api/flow/flow_apply'

const api = ref<Api>()
// 表单组件配置
const formJson = ref([])
// 表单数据
const formData = ref<Record<string, any>>({})
const options = ref({
    submitBtn: {
        show: false
    }
})

const dialogVisible = ref(false)
const applyDetail = ref({
    flow_name: ''
})

const historyDetail = ref<{ id: string | null; pass_status: number | null }>({
    id: null,
    pass_status: null
})

const props = defineProps({
    save: {
        type: Function,
        default: () => {}
    }
})
const emit = defineEmits(['back'])

function open(
    row: any,
    history: { id: string | null; pass_status: number | null },
    form_json: [],
    form_data: Record<string, any>
) {
    applyDetail.value = row
    historyDetail.value = history
    formData.value = form_data
    formJson.value = form_json
    console.log('open')

    dialogVisible.value = true
}

function closeFn() {
    dialogVisible.value = false
    applyDetail.value = { flow_name: '' }
    formData.value = {}
    formJson.value = []
    historyDetail.value = {
        id: null,
        pass_status: null
    }
}
function onBack() {
    emit('back', historyDetail.value)
}
function onSubmit() {
    console.log('formData', formData.value)
    api.value?.validate().then(() => {
        //todo 验证通过
        props
            .save(historyDetail.value.id, formData.value)
            .then(() => {
                closeFn()
            })
            .catch(() => {})
    })
}
defineExpose({
    open,

    closeFn
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
