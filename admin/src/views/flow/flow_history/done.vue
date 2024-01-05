<!-- 已完成审批 -->
<template>
    <div class="index-lists">
        <el-card class="!border-none" shadow="never">
            <el-form ref="formRef" class="mb-[-16px]" :model="queryParams" :inline="true">
                <el-form-item label="申请人昵称" prop="applyUserNickname">
                    <el-input v-model="queryParams.applyUserNickname" />
                </el-form-item>

                <el-form-item>
                    <el-button type="primary" @click="resetPage">查询</el-button>
                    <el-button @click="resetParams">重置</el-button>
                </el-form-item>
            </el-form>
        </el-card>
        <el-card class="!border-none mt-4" shadow="never">
            <!-- <div></div> -->
            <el-table class="mt-4" size="large" v-loading="pager.loading" :data="pager.lists">
                <el-table-column label="申请人" prop="applyUserNickname" min-width="100" />

                <!-- <el-table-column label="表单值" prop="formValue" min-width="100" /> -->
                <el-table-column label="通过状态" prop="passStatus" min-width="100">
                    <template #default="{ row }">
                        <dict-value
                            :options="dictData.flow_history_status"
                            :value="row.passStatus"
                        />
                    </template>
                </el-table-column>
                <el-table-column label="审批备注" prop="passRemark" min-width="100" />
                <el-table-column label="更新时间" prop="updateTime" min-width="150" />
                <el-table-column label="创建时间" prop="createTime" min-width="150" />
                <el-table-column label="操作" width="120" fixed="right">
                    <template #default="{ row }">
                        <!-- <el-button
                            v-perms="['admin:flow_history:edit']"
                            type="primary"
                            link
                            @click="handleOpen(row)"
                        >
                            审批
                        </el-button> -->
                        <el-button
                            v-perms="['admin:flow:flow_apply:edit']"
                            type="primary"
                            link
                            @click="OpenViewForm(row)"
                        >
                            {{ row.passStatus == 1 ? '审批' : '预览' }}
                        </el-button>
                        <!-- <el-button
                            v-perms="['admin:flow:flow_apply:edit']"
                            type="primary"
                            link
                            @click="OpenApplySubmit(row)"
                        >
                            审批
                        </el-button> -->
                    </template>
                </el-table-column>
            </el-table>
            <div class="flex justify-end mt-4">
                <pagination v-model="pager" @change="getLists" />
            </div>
        </el-card>

        <!-- <Approve ref="ApproveRef"></Approve> -->

        <ViewForm ref="viewFormRef"></ViewForm>
        <!-- <ApplySubmit ref="ApplySubmitRef" title="审批" @close="getLists"></ApplySubmit> -->
    </div>
</template>
<script lang="ts" setup>
import { flow_apply_detail } from '@/api/flow/flow_apply'
import { flow_history_list } from '@/api/flow/flow_history'
import { useDictData } from '@/hooks/useDictOptions'
import { usePaging } from '@/hooks/usePaging'
import feedback from '@/utils/feedback'
import useUserStore from '@/stores/modules/user'
// import ApplySubmit from '@/views/flow/flow_apply/components/apply_submit.vue'
import ViewForm from './components/ViewForm.vue'

const userStore = useUserStore()

defineOptions({
    name: 'todo'
})
const ApproveRef = shallowRef<InstanceType<typeof ApproveRef>>()
const viewFormRef = shallowRef<InstanceType<typeof ViewForm>>()

// const ApplySubmitRef = shallowRef<InstanceType<typeof ApplySubmit>>()

const queryParams = reactive({
    approverId: userStore?.userInfo?.id,
    applyUserNickname: '',
    passStatus: 2
})

const { pager, getLists, resetPage, resetParams } = usePaging({
    fetchFun: flow_history_list,
    params: queryParams
})
const { dictData } = useDictData<{
    flow_history_status: any[]
}>(['flow_history_status'])
// const handleOpen = async (row) => {
//     ApproveRef.value?.open(toRaw(row))
// }
const OpenViewForm = async (row: any) => {
    const applyDetail = await flow_apply_detail({ id: row.applyId })

    let form_data = {}
    try {
        form_data = JSON.parse(row.formValue)
    } catch (error) {
        // 解析失败
    }
    let form_json = {}
    try {
        form_json = JSON.parse(applyDetail.flowFormData)
    } catch (error) {
        // 解析失败
    }

    console.log(applyDetail, row, form_data, form_json)

    viewFormRef.value?.open(applyDetail, row, form_json, form_data)
}

getLists()
</script>
