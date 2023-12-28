<template>
    <el-dialog
        v-model="dialogVisible"
        :fullscreen="false"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
        :destroy-on-close="true"
        :title="title"
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
                    style="width: 100%"
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

            <el-form-item label="审批意见" prop="passRemark" v-if="props.showRemark">
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
            <el-button type="primary" @click="submit"> 通过 </el-button>
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
    })
    flow_history_get_approver({ applyId: applyId }).then((user) => {
        console.log('user', user)
        approverUserList.value = user
        if (user && user.length == 1) {
            formData.applyUserId = user[0].id
        }
    })
}
function BeforeClose() {
    dialogVisible.value = false

    // formData = {}
}
function submit() {
    console.log('submit', next_nodes)

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
