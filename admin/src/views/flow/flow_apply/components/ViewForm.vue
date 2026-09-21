<template>
    <el-dialog
        v-model="dialogVisible"
        :fullscreen="false"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
        :destroy-on-close="true"
        :title="applyDetail.flow_name || '流程详情'"
        class="view-form-dialog"
        width="90%"
    >
        <div class="view-form-body">
            <!-- 左侧：表单 -->
            <div class="view-form-main">
                <FormCreate
                    v-if="dialogVisible"
                    :rule="formJson"
                    v-model="formData"
                    v-model:api="api"
                    :option="options"
                ></FormCreate>
            </div>

            <!-- 右侧：审批流程时间线 -->
            <div class="view-form-steps" v-if="showHistoryList.length">
                <el-timeline>
                    <el-timeline-item
                        v-for="(item, index) in showHistoryList"
                        :key="index"
                        :color="stepColor(item.pass_status)"
                        :hollow="item.pass_status === 1"
                    >
                        <div class="view-form-step__node">
                            <div class="view-form-step__row">
                                <span class="view-form-step__label">
                                    {{ item.node_label || '未命名节点' }}
                                </span>
                                <span class="view-form-step__status">
                                    <dict-value
                                        :options="dictData.flow_history_status"
                                        :value="item.pass_status"
                                    />
                                </span>
                            </div>
                            <div class="view-form-step__row">
                                <span class="view-form-step__approver">
                                    {{ item.approver_nickname || '' }}
                                </span>
                                <span class="view-form-step__time" v-if="item.create_time">
                                    {{ item.create_time }}
                                </span>
                            </div>
                            <div class="view-form-step__remark" v-if="item.pass_remark">
                                备注：{{ item.pass_remark }}
                            </div>
                        </div>
                    </el-timeline-item>
                </el-timeline>
            </div>
        </div>

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
import { computed, ref, defineAsyncComponent } from 'vue'
// import formCreate from '@form-create/element-ui'
const FormCreate = defineAsyncComponent(() => import('@form-create/element-ui'))
import type { Api } from '@form-create/element-ui'
import { useDictData } from '@/hooks/useDictOptions'
import type { type_dict } from '@/hooks/useDictOptions'

import { flow_history_list_all, type type_flow_history } from '@/api/flow/flow_history'
import type { type_flow_apply } from '@/api/flow/flow_apply'
const props = defineProps({
    save: {
        type: Function,
        default: () => {}
    }
})

const api = ref<Api | null>(null)
// 表单组件配置
const formJson = ref([])
// 表单数据
const formData = ref<Record<string, any>>({})
const options = ref({
    submitBtn: {
        show: false
    },
    onSubmit: (formData: Record<string, any>) => {
        console.log(JSON.stringify(formData))
        onSubmit()
    }
    // resetBtn: true
})

const dialogVisible = ref(false)
const applyDetail = ref<type_flow_apply>({
    id: '',
    template_id: null,
    apply_user_id: null,
    apply_user_nickname: null,
    flow_name: '',
    flow_group: null,
    flow_remark: null,
    flow_form_data: null,
    flow_process_data: null,
    flow_process_data_list: null,
    form_value: null,
    status: null,
    create_time: null,
    update_time: null
})
const { dictData } = useDictData<{
    flow_history_status: type_dict[]
}>(['flow_history_status'])

interface FlowHistoryItem {
    node_type: string | null
    node_label: string | null
    approver_nickname: string | null
    pass_status: number | null
    pass_remark: string | null
    [key: string]: any
}
const historyList = ref<type_flow_history[]>([])
const showHistoryList = computed<FlowHistoryItem[]>(() => {
    return historyList.value.map((item) => ({ ...item }))
})

/**
 *
 * @param row 申请详情
 * @param form_json 表单配置
 * @param form_data 表单数据
 */
function open(row: type_flow_apply, form_json: any, form_data: Record<string, any>) {
    applyDetail.value = row
    getHistoryList(row.id)
    formData.value = form_data
    formJson.value = form_json
    console.log('open')
    dialogVisible.value = true
}
async function getHistoryList(applyId: string) {
    try {
        historyList.value = await flow_history_list_all({
            apply_id: applyId
        })
    } catch (error) {
        historyList.value = []
    }
}
function closeFn() {
    dialogVisible.value = false
    applyDetail.value = {
        id: '',
        template_id: null,
        apply_user_id: null,
        apply_user_nickname: null,
        flow_name: '',
        flow_group: null,
        flow_remark: null,
        flow_form_data: null,
        flow_process_data: null,
        flow_process_data_list: null,
        form_value: null,
        status: null,
        create_time: null,
        update_time: null
    }
    formData.value = {}
    formJson.value = []
    historyList.value = []
}
function onSubmit() {
    console.log('formData', formData.value)
    api.value?.validate().then(() => {
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
// 根据审批状态返回时间线节点颜色
function stepColor(status: number | null): string {
    switch (status) {
        case 2: // 通过
            return '#67C23A'
        case 3: // 驳回
            return '#F56C6C'
        case 4: // 撤销
            return '#909399'
        default: // 待审批/未开始
            return '#409EFF'
    }
}
defineExpose({
    open
})
</script>

<style lang="scss">
.el-range-editor.el-input__wrapper {
    box-sizing: border-box;
}

// 对话框内容区左右布局
.view-form-dialog {
    .el-dialog__body {
        padding: 12px 16px;
    }
}
.view-form-body {
    display: flex;
    gap: 16px;
    min-height: 420px;
    max-height: 70vh;
}
// 右侧审批流程时间线
.view-form-steps {
    width: 280px;
    flex-shrink: 0;
    overflow-y: auto;
    border-left: 1px solid var(--el-border-color-lighter);
    padding-left: 12px;
    // 让时间线贴左，避免右侧留白过多
    :deep(.el-timeline) {
        padding-left: 0;
    }
}
.view-form-step {
    &__node {
        display: flex;
        flex-direction: column;
        gap: 2px;
        padding: 4px 0;
    }
    &__row {
        display: flex;
        align-items: baseline;
        justify-content: space-between;
        gap: 8px;
    }
    &__label {
        font-weight: 600;
        color: var(--el-text-color-primary);
    }
    &__time {
        flex-shrink: 0;
        font-size: 12px;
        color: var(--el-text-color-secondary);
    }
    &__approver {
        font-size: 13px;
        color: var(--el-text-color-regular);
    }
    &__remark {
        margin-top: 2px;
        font-size: 12px;
        color: var(--el-text-color-secondary);
    }
    &__status {
        flex-shrink: 0;
    }
}
// 右侧表单
.view-form-main {
    flex: 1;
    min-width: 0;
    overflow-y: auto;
}
</style>
