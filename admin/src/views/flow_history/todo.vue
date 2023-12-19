<!-- 代办 -->
<template>
    <div class="index-lists">
        <el-card class="!border-none" shadow="never">
            <el-form ref="formRef" class="mb-[-16px]" :model="queryParams" :inline="true">
                <el-form-item label="申请人昵称" prop="applyUserNickname">
                    <el-input v-model="queryParams.applyUserNickname" />
                </el-form-item>

                <el-form-item label="审批用户昵称" prop="approverNickname">
                    <el-input v-model="queryParams.approverNickname" />
                </el-form-item>
                <el-form-item label="节点" prop="nodeId">
                    <el-input v-model="queryParams.nodeId" />
                </el-form-item>

                <el-form-item>
                    <el-button type="primary" @click="resetPage">查询</el-button>
                    <el-button @click="resetParams">重置</el-button>
                </el-form-item>
            </el-form>
        </el-card>
        <el-card class="!border-none mt-4" shadow="never">
            <div>
                <el-button v-perms="['flow_history:add']" type="primary" @click="handleAdd()">
                    <template #icon>
                        <icon name="el-icon-Plus" />
                    </template>
                    新增
                </el-button>
            </div>
            <el-table class="mt-4" size="large" v-loading="pager.loading" :data="pager.lists">
                <el-table-column label="申请人昵称" prop="applyUserNickname" min-width="100" />
                <!-- <el-table-column label="审批人id" prop="approverId" min-width="100" />
                <el-table-column label="审批用户昵称" prop="approverNickname" min-width="100" /> -->
                <el-table-column label="节点" prop="nodeId" min-width="100" />
                <el-table-column label="表单值" prop="formValue" min-width="100" />
                <el-table-column
                    label="通过状态：1待处理，2通过，3拒绝"
                    prop="passStatus"
                    min-width="100"
                >
                    <template #default="{ row }">
                        <dict-value
                            :options="dictData.flow_history_status"
                            :value="row.passStatus"
                        />
                    </template>
                </el-table-column>
                <el-table-column label="通过备注" prop="passRemark" min-width="100" />
                <el-table-column label="更新时间" prop="updateTime" min-width="150" />
                <el-table-column label="创建时间" prop="createTime" min-width="150" />
                <el-table-column label="操作" width="120" fixed="right">
                    <template #default="{ row }">
                        <el-button
                            v-perms="['flow_history:edit']"
                            type="primary"
                            link
                            @click="handleOpen(row)"
                        >
                            审批
                        </el-button>
                    </template>
                </el-table-column>
            </el-table>
            <div class="flex justify-end mt-4">
                <pagination v-model="pager" @change="getLists" />
            </div>
        </el-card>

        <Approve ref="ApproveRef"></Approve>
    </div>
</template>
<script lang="ts" setup>
import { flow_history_list } from '@/api/flow_history'
import { useDictData } from '@/hooks/useDictOptions'
import { usePaging } from '@/hooks/usePaging'
import feedback from '@/utils/feedback'
import useUserStore from '@/stores/modules/user'
import Approve from './components/todo/approve.vue'

const userStore = useUserStore()

defineOptions({
    name: 'todo'
})
const ApproveRef = shallowRef<InstanceType<typeof ApproveRef>>()

const queryParams = reactive({
    applyId: '',
    templateId: '',
    applyUserId: '',
    applyUserNickname: '',
    approverId: userStore?.userInfo?.id,
    approverNickname: '',
    nodeId: '',
    formValue: '',
    passStatus: 1,
    passRemark: ''
})

const { pager, getLists, resetPage, resetParams } = usePaging({
    fetchFun: flow_history_list,
    params: queryParams
})
const { dictData } = useDictData<{
    flow_history_status: any[]
}>(['flow_history_status'])
const handleOpen = async (row) => {
    ApproveRef.value?.open(toRaw(row))
}

getLists()
</script>
