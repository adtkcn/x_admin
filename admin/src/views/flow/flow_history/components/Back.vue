<template>
    <el-dialog
        v-model="dialogVisible"
        :show-close="true"
        :fullscreen="false"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
        :destroy-on-close="true"
        :title="title"
        draggable
    >
        <el-form
            ref="formRef"
            :model="formData"
            label-position="top"
            label-width="110px"
            :rules="formRules"
        >
            <el-form-item :label="`驳回到`" prop="templateId">
                <el-select
                    style="width: 100%"
                    v-model="formData.historyId"
                    placeholder="请选择驳回节点"
                >
                    <el-option label="发起人" :value="0" clearable />
                    <el-option
                        v-for="item in back_nodes"
                        :key="item.id"
                        :label="item.nodeLabel"
                        :value="item.id"
                        clearable
                    />
                </el-select>
            </el-form-item>

            <el-form-item label="审批意见" prop="remark" v-if="props.showRemark">
                <el-input
                    v-model="formData.remark"
                    :rows="2"
                    type="textarea"
                    placeholder="请输入审批意见"
                />
            </el-form-item>
        </el-form>

        <template #footer>
            <el-button @click="dialogVisible = false">关闭</el-button>
            <el-button type="warning" @click="submit"> 驳回 </el-button>
        </template>
    </el-dialog>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { flow_history_back, flow_history_list_all } from '@/api/flow/flow_history'
const props = defineProps({
    title: {
        type: String,
        default: ''
    },
    showRemark: {
        type: Boolean,
        default: true
    }
})
const dialogVisible = ref(false)
// const props = defineProps({
//     save: {
//         type: Function,
//         default: () => {}
//     }
// })

class formDataState {
    applyId = '' // 申请id
    historyId: 0 //审批节点,0为发起人
    remark = '' // 备注
}
const formData = reactive(new formDataState())
const back_nodes = ref([])

const formRules = {
    applyUserId: [
        {
            required: true,
            message: '请选择',
            trigger: ['blur']
        }
    ]
}
function open(applyId) {
    console.log('open')
    Object.assign(formData, new formDataState())
    formData.applyId = applyId
    dialogVisible.value = true

    flow_history_list_all({
        applyId: applyId,
        nodeType: 'bpmn:userTask',
        passStatus: 2
    }).then((userNode) => {
        console.log('userNode', userNode)
        back_nodes.value = userNode
    })
}
function BeforeClose() {
    dialogVisible.value = false
}
function submit() {
    flow_history_back(formData).then(() => {
        BeforeClose()
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
</style>
