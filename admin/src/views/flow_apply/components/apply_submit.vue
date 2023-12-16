<template>
    <el-dialog
        v-model="dialogVisible"
        :show-close="false"
        :fullscreen="false"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
        :destroy-on-close="true"
        top="1px"
    >
        <el-form ref="formRef" :model="formData" label-width="84px" :rules="formRules">
            <el-form-item label="审批节点" prop="flowName">
                <el-input v-model="formData.flowName" placeholder="请输入流程名称" />
            </el-form-item>
            <el-form-item label="审批人" prop="applyUserId">
                <el-select
                    class="flex-1"
                    v-model="formData.applyUserId"
                    placeholder="请选择审批人"
                    @change="handleTemplateChange"
                >
                    <el-option
                        v-for="(item, index) in flow_template"
                        :key="index"
                        :label="item.flowName"
                        :value="item.id"
                        clearable
                    />
                </el-select>
            </el-form-item>
        </el-form>

        <template #footer>
            <el-button @click="dialogVisible = false">取消</el-button>
            <el-button type="primary" @click="getData"> 确定 </el-button>
        </template>
    </el-dialog>
</template>

<script setup>
import { ref } from 'vue'
import { flow_apply_next_node } from '@/api/flow_history'

const formRef = ref(null)

const dialogVisible = ref(false)

const props = defineProps({
    save: {
        type: Function,
        default: () => {}
    }
})
const formData = reactive({
    id: '',
    templateId: '',
    // applyUserId: '',
    // applyUserNickname: '',
    flowName: '',
    flowGroup: '',
    flowRemark: '',
    flowFormData: '',
    flowProcessData: '',
    status: 0,
    formValue: ''
})
const formRules = {
    id: [
        {
            required: true,
            message: '请输入',
            trigger: ['blur']
        }
    ]
}
function open(row) {
    console.log('open')
    formData.value = row
    dialogVisible.value = true
    flow_apply_next_node({
        id: row.id,
        historyId: ''
    }).then((res) => {
        console.log('res', res)
    })
}
function closeFn() {
    dialogVisible.value = false

    formData.value = {}
}
function getData() {
    formRef.value.getFormData().then((formData) => {
        console.log('formData', formData)
        props
            .save(formData)
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
