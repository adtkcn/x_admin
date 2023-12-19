<!-- 审批 -->

<!-- 预览和编辑
 -->
<!-- 审批记录备注 -->
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
        <v-form-render
            v-if="render"
            :form-json="formJson"
            :form-data="formRenderData"
            :option-data="optionData"
            ref="vFormRef"
        >
        </v-form-render>

        <el-divider> 审批 </el-divider>
        <el-form
            ref="formRef"
            :model="formData"
            label-position="top"
            label-width="110px"
            :rules="formRules"
        >
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
        <template #footer>
            <el-button @click="dialogVisible = false">关闭</el-button>
            <el-button type="primary" @click="save"> 通过 </el-button>
        </template>
    </el-dialog>
</template>

<script setup lang="ts">
import 'vform3-builds/dist/designer.style.css' //引入VForm3样式
import {
    flow_history_next_node,
    flow_history_get_approver,
    flow_history_pass
} from '@/api/flow_history'
import { flow_apply_detail } from '@/api/flow_apply'
import { defineExpose, defineProps, defineOptions, reactive, ref, computed } from 'vue'
defineOptions({
    name: 'todo-approve'
})
// 表单
const formJson = ref({})
const formRenderData = ref({})
const optionData = reactive({})
const vFormRef = ref(null)

const dialogVisible = ref(false)

// 审批
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
const render = ref(false)
function reset() {
    render.value = false
}
async function open(row) {
    reset()
    if (row) {
        dialogVisible.value = true

        const flowFormData = await flow_apply_detail({
            id: row.applyId
        })

        formRenderData.value = JSON.parse(row.formValue)
        formJson.value = JSON.parse(flowFormData.flowFormData)

        render.value = true
        flow_history_next_node({
            applyId: row.applyId,
            currentNodeId: row.nodeId
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
}
function save() {
    vFormRef.value.getFormData().then((formData) => {
        console.log('formData', formData)

        flow_history_pass({
            applyId: formData.id,
            passRemark: formData.passRemark,
            applyUserId: formData.applyUserId
        }).then((res) => {
            console.log('res', res)
            dialogVisible.value = false
        })
    })
}
defineExpose({
    open
})
</script>

<style lang="scss">
.el-range-editor.el-input__wrapper {
    box-sizing: border-box;
}
</style>
