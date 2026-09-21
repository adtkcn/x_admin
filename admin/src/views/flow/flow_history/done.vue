<!-- 已完成审批 -->
<template>
    <div class="index-lists">
        <el-card class="border-none!" shadow="never">
            <el-form ref="formRef" class="mb-[-16px]" :model="queryParams" :inline="true">
                <el-form-item label="申请人昵称" prop="apply_user_nickname">
                    <el-input v-model="queryParams.apply_user_nickname" />
                </el-form-item>

                <el-form-item>
                    <el-button type="primary" @click="resetPage">查询</el-button>
                    <el-button @click="resetParams">重置</el-button>
                </el-form-item>
            </el-form>
        </el-card>
        <el-card class="border-none! mt-4" shadow="never">
            <!-- <div></div> -->
            <vxe-table
                class="mt-4"
                v-loading="pager.loading"
                :data="pager.lists"
                :row-config="{ keyField: 'id' }"
                :scroll-y="{ enabled: false }"
                :border="'inner'"
            >
                <vxe-column title="申请人" field="apply_user_nickname" min-width="100" />

                <!-- <vxe-column title="表单值" field="formValue" min-width="100" /> -->
                <vxe-column title="通过状态" field="pass_status" min-width="100">
                    <template #default="{ row }">
                        <dict-value
                            :options="dictData.flow_history_status"
                            :value="row.pass_status"
                        />
                    </template>
                </vxe-column>
                <vxe-column title="审批备注" field="pass_remark" min-width="100" />
                <vxe-column title="更新时间" field="update_time" min-width="150" />
                <vxe-column title="创建时间" field="create_time" min-width="150" />
                <vxe-column title="操作" width="140" fixed="right">
                    <template #default="{ row }">
                        <el-button
                            v-perms="['admin:flow:flow_apply:edit']"
                            type="primary"
                            link
                            @click="OpenViewForm(row)"
                        >
                            {{ row.pass_status == 1 ? '审批' : '预览' }}
                        </el-button>
                        <el-button
                            v-perms="['admin:flow:flow_history:del']"
                            type="danger"
                            link
                            @click="handleDelete(row)"
                        >
                            删除
                        </el-button>
                    </template>
                </vxe-column>
            </vxe-table>
            <div class="flex justify-end mt-4">
                <pagination v-model="pager" @change="getLists" />
            </div>
        </el-card>

        <ViewForm ref="viewFormRef"></ViewForm>
    </div>
</template>
<script lang="ts" setup>
import { shallowRef, reactive, defineAsyncComponent, onMounted, onActivated } from 'vue'
import { flow_apply_detail } from '@/api/flow/flow_apply'
import { flow_history_list, flow_history_done_hidden } from '@/api/flow/flow_history'
import type { type_flow_apply } from '@/api/flow/flow_apply'

import { useDictData } from '@/hooks/useDictOptions'
import { usePaging } from '@/hooks/usePaging'
import useUserStore from '@/stores/modules/user'
import feedback from '@/utils/feedback'
const ViewForm = defineAsyncComponent(() => import('./components/ViewForm.vue'))
const userStore = useUserStore()

defineOptions({
    name: 'done'
})
const viewFormRef = shallowRef<InstanceType<typeof ViewForm>>()

// const ApplySubmitRef = shallowRef<InstanceType<typeof ApplySubmit>>()

const queryParams = reactive({
    approver_id: String(userStore?.userInfo?.id),
    apply_user_nickname: '',
    pass_status: 2,
    is_show: 1
})

const { pager, getLists, resetPage, resetParams } = usePaging<type_flow_apply>({
    fetchFun: flow_history_list,
    params: queryParams
})
const { dictData } = useDictData<{
    flow_history_status: any[]
}>(['flow_history_status'])

const OpenViewForm = async (row: any) => {
    const applyDetail = await flow_apply_detail({ id: row.apply_id })

    let form_data = {}
    try {
        form_data = JSON.parse(row.form_value)
    } catch (error) {
        // 解析失败
    }
    let form_json = []
    try {
        if (applyDetail.flow_form_data) {
            form_json = JSON.parse(applyDetail.flow_form_data)
        }
    } catch (error) {
        // 解析失败
    }

    console.log(applyDetail, row, form_data, form_json)

    viewFormRef.value?.open(applyDetail, row, form_json, form_data)
}

const handleDelete = async (row: any) => {
    await feedback.confirm('确定要隐藏这条记录？')
    await flow_history_done_hidden(row.id)
    feedback.msgSuccess('操作成功')
    getLists()
}

onMounted(() => {
    getLists()
})
onActivated(() => {
    if (!pager.loading) {
        getLists()
    }
})
</script>
