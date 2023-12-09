<template>
    <div class="index-lists">
        <el-card class="!border-none" shadow="never">
            <el-form ref="formRef" class="mb-[-16px]" :model="queryParams" :inline="true">
                <el-form-item label="流程名称" prop="flowName">
                    <el-input class="w-[280px]" v-model="queryParams.flowName" />
                </el-form-item>
                <el-form-item label="流程分类" prop="flowGroup">
                    <el-input class="w-[280px]" v-model="queryParams.flowGroup" />
                </el-form-item>
                <el-form-item label="流程描述" prop="flowRemark">
                    <el-input class="w-[280px]" v-model="queryParams.flowRemark" />
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="resetPage">查询</el-button>
                    <el-button @click="resetParams">重置</el-button>
                </el-form-item>
            </el-form>
        </el-card>
        <el-card class="!border-none mt-4" shadow="never">
            <div>
                <el-button v-perms="['flow_template:add']" type="primary" @click="handleAdd()">
                    <template #icon>
                        <icon name="el-icon-Plus" />
                    </template>
                    新增
                </el-button>
            </div>
            <el-table
                class="mt-4"
                size="large"
                v-loading="pager.loading"
                :data="pager.lists"
            >
                <el-table-column label="流程名称" prop="flowName" min-width="100" />
                <el-table-column label="流程分类" prop="flowGroup" min-width="100" />
                <el-table-column label="流程描述" prop="flowRemark" min-width="100" />
                <el-table-column label="表单配置" prop="flowFormData" min-width="100" />
                <el-table-column label="流程配置" prop="flowProcessData" min-width="100" />
                <el-table-column label="操作" width="120" fixed="right">
                    <template #default="{ row }">
                        <el-button
                            v-perms="['flow_template:edit']"
                            type="primary"
                            link
                            @click="handleEdit(row)"
                        >
                            编辑
                        </el-button>
                        <el-button
                            v-perms="['flow_template:del']"
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
        <edit-popup
            v-if="showEdit"
            ref="editRef"
            @success="getLists"
            @close="showEdit = false"
        />
    </div>
</template>
<script lang="ts" setup name="flow_template">
import { flow_template_delete, flow_template_lists } from '@/api/flow_template'
import { usePaging } from '@/hooks/usePaging'
import feedback from '@/utils/feedback'
import EditPopup from './edit.vue'
const editRef = shallowRef<InstanceType<typeof EditPopup>>()
const showEdit = ref(false)
const queryParams = reactive({
    flowName: '',
    flowGroup: '',
    flowRemark: '',
    flowFormData: '',
    flowProcessData: '',
})

const { pager, getLists, resetPage, resetParams } = usePaging({
    fetchFun: flow_template_lists,
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
    await flow_template_delete({ id })
    feedback.msgSuccess('删除成功')
    getLists()
}

getLists()
</script>
