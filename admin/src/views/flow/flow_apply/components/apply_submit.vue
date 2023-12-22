<template>
    <el-dialog
        v-model="dialogVisible"
        :show-close="true"
        :fullscreen="false"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
        :destroy-on-close="true"
        title="审批"
        top="1px"
    >
        <el-form
            ref="formRef"
            :model="formData"
            label-position="top"
            label-width="110px"
            :rules="formRules"
        >
            <!-- {{ userTask }} -->
            <el-form-item v-if="userTask" :label="`${userTask?.label}节点审批人`" prop="templateId">
                <el-select
                    class="flex-1"
                    v-if="approverUserList.length"
                    v-model="formData.applyUserId"
                    placeholder="请选择审批人"
                >
                    <el-option
                        v-for="(item, index) in approverUserList"
                        :key="index"
                        :label="item.nickname"
                        :value="item.id"
                        clearable
                    />
                </el-select>
            </el-form-item>
            <el-form-item label="审批意见" prop="passRemark">
                <el-input
                    v-model="formData.passRemark"
                    :rows="2"
                    type="textarea"
                    placeholder="请输入审批意见"
                />
            </el-form-item>
        </el-form>
        <!-- <el-steps direction="horizontal" :active="next_nodes.length">
            <el-step :title="node.label" v-for="node of next_nodes" :key="node.id" />
        </el-steps> -->

        <template #footer>
            <el-button @click="dialogVisible = false">关闭</el-button>
            <el-button type="primary" @click="getData"> 通过 </el-button>
        </template>
    </el-dialog>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import feedback from '@/utils/feedback'
import {
    flow_history_next_node,
    flow_history_get_approver,
    flow_history_pass
} from '@/api/flow/flow_history'

const dialogVisible = ref(false)

// const props = defineProps({
//     save: {
//         type: Function,
//         default: () => {}
//     }
// })

class formDataState {
    id = ''
    passRemark = ''
    applyUserId = ''
}
const formData = reactive(new formDataState())
const next_nodes = ref([])
const userTask = computed(() => {
    return next_nodes.value.find((item) => item.type == 'bpmn:userTask')
})
const approverUserList = ref([])
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
    formData.id = applyId
    dialogVisible.value = true

    flow_history_next_node({
        applyId: applyId
    }).then((res) => {
        console.log('res', res)
        next_nodes.value = res

        res.map((item) => {
            if (item.type == 'bpmn:userTask') {
                flow_history_get_approver(item).then((user) => {
                    console.log('user', user)
                    approverUserList.value = user
                })
            }
        })
    })
}
function BeforeClose() {
    dialogVisible.value = false

    // formData = {}
}
function getData() {
    console.log('getData', next_nodes)

    if (userTask.value && !formData.applyUserId) {
        feedback.msgWarning('请选择审批人')
        return
    }
    flow_history_pass({
        applyId: formData.id,
        nextNodeAdminId: formData.applyUserId || 0,
        passRemark: formData.passRemark
    }).then(() => {
        BeforeClose()
    })
    // formRef.value.getFormData().then((formData) => {
    //     console.log('formData', formData)
    //     props
    //         .save(formData)
    //         .then(() => {
    //             BeforeClose()
    //         })
    //         .catch(() => {})
    // })
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
@/api/flow/flow_history
