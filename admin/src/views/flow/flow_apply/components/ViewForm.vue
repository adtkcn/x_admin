<template>
    <el-dialog
        v-model="dialogVisible"
        :fullscreen="false"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
        :destroy-on-close="true"
        :title="applyDetail.flowName"
    >
        <formCreate
            v-if="dialogVisible"
            :rule="formJson"
            v-model="formData"
            v-model:api="api"
            :option="options"
        ></formCreate>

        <!-- <v-form-render
            :form-json="formJson"
            :form-data="formData"
            :option-data="optionData"
            ref="vFormRef"
        >
        </v-form-render> -->

        <el-table size="large" :data="showHistoryList" v-if="showHistoryList.length">
            <el-table-column label="审批人" prop="approverNickname" />
            <el-table-column label="节点" prop="nodeLabel" />
            <el-table-column label="状态" prop="passStatus">
                <template #default="{ row }">
                    <dict-value :options="dictData.flow_history_status" :value="row.passStatus" />
                </template>
            </el-table-column>
            <el-table-column label="备注" prop="passRemark" />
        </el-table>

        <template #footer>
            <el-button @click="dialogVisible = false">关闭</el-button>
            <el-button
                v-if="applyDetail.status == 1 || applyDetail.status == 4"
                type="primary"
                @click="onSubmit"
            >
                确定
            </el-button>
        </template>
    </el-dialog>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'
import formCreate from '@form-create/element-ui'
import type { Api } from '@form-create/element-ui'
import { useDictData } from '@/hooks/useDictOptions'
import type { type_dict } from '@/hooks/useDictOptions'

import { flow_history_list_all } from '@/api/flow/flow_history'
const props = defineProps({
    save: {
        type: Function,
        default: () => {}
    }
})

const api = ref<Api>(null)
// 表单组件配置
const formJson = ref([])
// 表单数据
const formData = ref<Record<string, any>>({})
const options = ref({
    submitBtn: {
        show: false
    },
    onSubmit: (formData) => {
        console.log(JSON.stringify(formData))
        onSubmit()
    }
    // resetBtn: true
})

const dialogVisible = ref(false)
const applyDetail = ref({
    id: null,
    status: null,
    flowName: ''
})
const { dictData } = useDictData<{
    flow_history_status: type_dict[]
}>(['flow_history_status'])

const historyList = ref([])
const showHistoryList = computed(() => {
    return historyList.value.filter((item) => {
        return item.nodeType !== 'bpmn:startEvent' && item.nodeType !== 'bpmn:endEvent'
    })
})

/**
 *
 * @param row 申请详情
 * @param form_json 表单配置
 * @param form_data 表单数据
 */
function open(row, form_json, form_data: Record<string, any>) {
    applyDetail.value = row
    getHistoryList(row.id)
    formData.value = form_data
    formJson.value = form_json
    console.log('open')
    dialogVisible.value = true
}
async function getHistoryList(applyId) {
    try {
        historyList.value = await flow_history_list_all({
            applyId: applyId
        })
    } catch (error) {
        historyList.value = []
    }
}
function closeFn() {
    dialogVisible.value = false
    applyDetail.value = {
        id: null,
        status: null,
        flowName: ''
    }
    formData.value = {}
    formJson.value = []
    historyList.value = []
}
function onSubmit() {
    console.log('formData', formData.value)
    api.value.validate().then(() => {
        //todo 验证通过
        props
            .save(applyDetail.value?.id, formData.value)
            .then(() => {
                closeFn()
            })
            .catch(() => {})
    })
    // vFormRef.value.getFormData().then((formData) => {
    //     console.log('formData', formData)
    //     props
    //         .save(applyDetail.value?.id, formData)
    //         .then(() => {
    //             closeFn()
    //         })
    //         .catch(() => {})
    // })
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
