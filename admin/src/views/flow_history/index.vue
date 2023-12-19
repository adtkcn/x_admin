<template>
    <div class="index-lists">
        <el-card class="!border-none" shadow="never">
            <el-form ref="formRef" class="mb-[-16px]" :model="queryParams" :inline="true">
                <el-form-item label="申请id" prop="applyId">
                    <el-input v-model="queryParams.applyId" />
                </el-form-item>

                <el-form-item label="申请人昵称" prop="applyUserNickname">
                    <el-input v-model="queryParams.applyUserNickname" />
                </el-form-item>

                <el-form-item label="审批用户昵称" prop="approverNickname">
                    <el-input v-model="queryParams.approverNickname" />
                </el-form-item>
                <el-form-item label="节点" prop="nodeId">
                    <el-input v-model="queryParams.nodeId" />
                </el-form-item>
                <el-form-item label="通过状态" prop="passStatus">
                    <el-select v-model="queryParams.passStatus" clearable>
                        <el-option label="请选择字典生成" value="" />
                    </el-select>
                </el-form-item>
                <el-form-item label="通过备注" prop="passRemark">
                    <el-input v-model="queryParams.passRemark" />
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
                <el-table-column label="申请id" prop="applyId" min-width="100" />
                <el-table-column label="模板id" prop="templateId" min-width="100" />
                <el-table-column label="申请人id" prop="applyUserId" min-width="100" />
                <el-table-column label="申请人昵称" prop="applyUserNickname" min-width="100" />
                <el-table-column label="审批人id" prop="approverId" min-width="100" />
                <el-table-column label="审批用户昵称" prop="approverNickname" min-width="100" />
                <el-table-column label="节点" prop="nodeId" min-width="100" />
                <el-table-column label="表单值" prop="formValue" min-width="100" />
                <el-table-column
                    label="通过状态：1待处理，2通过，3拒绝"
                    prop="passStatus"
                    min-width="100"
                />
                <el-table-column label="通过备注" prop="passRemark" min-width="100" />
                <el-table-column label="更新时间" prop="updateTime" min-width="100" />
                <el-table-column label="创建时间" prop="createTime" min-width="100" />
                <el-table-column label="操作" width="120" fixed="right">
                    <template #default="{ row }">
                        <el-button
                            v-perms="['flow_history:edit']"
                            type="primary"
                            link
                            @click="handleEdit(row)"
                        >
                            编辑
                        </el-button>
                        <el-button
                            v-perms="['flow_history:del']"
                            type="danger"
                            link
                            @click="handleDelete(row.id)"
                        >
                            删除
                        </el-button>
                    </template>
                </el-table-column>
            </el-table>
            <div class="flex justify-end mt-4">
                <pagination v-model="pager" @change="getLists" />
            </div>
        </el-card>
        <edit-popup v-if="showEdit" ref="editRef" @success="getLists" @close="showEdit = false" />
    </div>
</template>
<script lang="ts" setup>
import { flow_history_delete, flow_history_list } from '@/api/flow_history'
import { usePaging } from '@/hooks/usePaging'
import feedback from '@/utils/feedback'
import EditPopup from './edit.vue'
defineOptions({
    name: 'flow_history'
})
const editRef = shallowRef<InstanceType<typeof EditPopup>>()
const showEdit = ref(false)
const queryParams = reactive({
    applyId: '',
    templateId: '',
    applyUserId: '',
    applyUserNickname: '',
    approverId: '',
    approverNickname: '',
    nodeId: '',
    formValue: '',
    passStatus: '',
    passRemark: ''
})

const { pager, getLists, resetPage, resetParams } = usePaging({
    fetchFun: flow_history_list,
    params: queryParams
})

const handleAdd = async () => {
    showEdit.value = true
    await nextTick()
    editRef.value?.open('add')
}

const handleEdit = async (data: any) => {
    showEdit.value = true
    await nextTick()
    editRef.value?.open('edit')
    editRef.value?.getDetail(data)
}

const handleDelete = async (id: number) => {
    await feedback.confirm('确定要删除？')
    await flow_history_delete({ id })
    feedback.msgSuccess('删除成功')
    getLists()
}

getLists()
</script>
