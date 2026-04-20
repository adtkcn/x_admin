<template>
    <div class="index-lists">
        <el-card class="border-none!" shadow="never">
            <el-form
                ref="formRef"
                class="mb-[-16px]"
                :model="queryParams"
                :inline="true"
                label-width="90px"
                label-position="left"
            >
                <el-form-item label="任务名称" prop="TaskName" class="w-[280px]">
                    <el-input v-model="queryParams.TaskName" />
                </el-form-item>
                <el-form-item label="任务编码" prop="TaskCode" class="w-[280px]">
                    <el-input v-model="queryParams.TaskCode" />
                </el-form-item>
                <el-form-item label="corn表达式" prop="CornExpr" class="w-[280px]">
                    <el-input v-model="queryParams.CornExpr" />
                </el-form-item>
                <el-form-item label="创建人" prop="CreatedBy" class="w-[280px]">
                    <el-input v-model="queryParams.CreatedBy" />
                </el-form-item>
                <el-form-item label="创建人名称" prop="CreatedByNickname" class="w-[280px]">
                    <el-input v-model="queryParams.CreatedByNickname" />
                </el-form-item>

                <el-form-item label="创建时间" prop="CreateTime" class="w-[280px]">
                    <daterange-picker
                        v-model:startTime="queryParams.CreateTimeStart"
                        v-model:endTime="queryParams.CreateTimeEnd"
                    />
                </el-form-item>
                <el-form-item label="更新时间" prop="UpdateTime" class="w-[280px]">
                    <daterange-picker
                        v-model:startTime="queryParams.UpdateTimeStart"
                        v-model:endTime="queryParams.UpdateTimeEnd"
                    />
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="resetPage">查询</el-button>
                    <el-button @click="resetParams">重置</el-button>
                </el-form-item>
            </el-form>
        </el-card>
        <el-card class="border-none! mt-4" shadow="never">
            <div class="text-right">
                <el-button v-perms="['admin:system_corn:add']" type="primary" @click="handleAdd()">
                    <template #icon>
                        <icon name="el-icon-Plus" />
                    </template>
                    新增
                </el-button>
                <Upload
                    v-perms="['admin:system_corn:ImportFile']"
                    class="ml-3 mr-3"
                    :url="system_corn_import_file"
                    :ext="['xlsx']"
                    :show-progress="true"
                    @change="resetPage"
                >
                    <el-button type="primary">
                        <template #icon>
                            <icon name="el-icon-Upload" />
                        </template>
                        导入
                    </el-button>
                </Upload>
                <el-button
                    v-perms="['admin:system_corn:ExportFile']"
                    type="primary"
                    @click="export_file"
                >
                    <template #icon>
                        <icon name="el-icon-Download" />
                    </template>
                    导出
                </el-button>
                <el-button
                    v-perms="['admin:system_corn:del_batch']"
                    type="danger"
                    :disabled="!multipleSelection.length"
                    @click="deleteBatch"
                >
                    批量删除
                </el-button>
            </div>
            <el-table
                class="mt-4"
                size="large"
                v-loading="pager.loading"
                :data="pager.lists"
                @selection-change="handleSelectionChange"
            >
                <el-table-column type="selection" width="55" />
                <el-table-column label="任务名称" prop="TaskName" min-width="130" />
                <el-table-column label="任务编码" prop="TaskCode" min-width="130" />
                <el-table-column label="corn表达式" prop="CornExpr" min-width="130" />
                <el-table-column label="状态" prop="Status" min-width="130">
                    <template #default="{ row }">
                        <dict-value :options="dictData.status" :value="row.Status" />
                    </template>
                </el-table-column>
                <el-table-column label="创建人" prop="CreatedByUser.nickname" min-width="130" />
                <el-table-column label="创建时间" prop="CreateTime" min-width="130" />
                <el-table-column label="更新时间" prop="UpdateTime" min-width="130" />
                <el-table-column label="操作" width="160" fixed="right">
                    <template #default="{ row }">
                        <el-button
                            v-perms="['admin:system_corn:detail']"
                            type="primary"
                            link
                            @click="viewDetails(row)"
                            >详情</el-button
                        >
                        <el-button
                            v-perms="['admin:system_corn:edit', 'admin:system_corn:detail']"
                            type="primary"
                            link
                            @click="handleEdit(row)"
                        >
                            编辑
                        </el-button>
                        <el-button
                            v-perms="['admin:system_corn:del']"
                            type="danger"
                            link
                            @click="handleDelete(row.Id)"
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
        <EditPopup v-if="showEdit" ref="editRef" @success="getLists" @close="showEdit = false" />
        <DetailsPopup v-if="showDetails" ref="detailsRef" @close="showDetails = false" />
    </div>
</template>
<script lang="ts" setup>
import { ref, reactive, shallowRef, nextTick } from 'vue'
import {
    system_corn_delete,
    system_corn_delete_batch,
    system_corn_list,
    system_corn_import_file,
    system_corn_export_file
} from '@/api/system/corn'
import type { type_system_corn, type_system_corn_query } from '@/api/system/corn'

import { useDictData } from '@/hooks/useDictOptions'
import type { type_dict } from '@/hooks/useDictOptions'
const { dictData } = useDictData<{
    status: type_dict[]
}>(['status'])
import { usePaging } from '@/hooks/usePaging'
import feedback from '@/utils/feedback'
import EditPopup from './edit.vue'
import DetailsPopup from './details.vue'
defineOptions({
    name: 'system_corn'
})
const editRef = shallowRef<InstanceType<typeof EditPopup>>()
const showEdit = ref(false)
const detailsRef = shallowRef<InstanceType<typeof DetailsPopup>>()
const showDetails = ref(false)
const queryParams = reactive<type_system_corn_query>({
    TaskName: undefined,
    TaskCode: undefined,
    CornExpr: undefined,
    Status: undefined,
    CreatedBy: undefined,
    CreatedByNickname: undefined,
    CreateTimeStart: undefined,
    CreateTimeEnd: undefined,
    UpdateTimeStart: undefined,
    UpdateTimeEnd: undefined
})

const { pager, getLists, resetPage, resetParams } = usePaging<type_system_corn>({
    fetchFun: system_corn_list,
    params: queryParams
})

const handleAdd = async () => {
    showEdit.value = true
    await nextTick()
    editRef.value?.open('add')
}

const handleEdit = async (data: type_system_corn) => {
    showEdit.value = true
    await nextTick()
    editRef.value?.open('edit')
    editRef.value?.getDetail(data)
}
const viewDetails = async (data: type_system_corn) => {
    showDetails.value = true
    await nextTick()
    detailsRef.value?.open()
    detailsRef.value?.getDetail(data)
}
const multipleSelection = ref<type_system_corn[]>([])
const handleSelectionChange = (val: type_system_corn[]) => {
    multipleSelection.value = val
}

const handleDelete = async (Id: string) => {
    try {
        await feedback.confirm('确定要删除？')
        await system_corn_delete(Id)
        feedback.msgSuccess('删除成功')
        getLists()
    } catch (error) {}
}
// 批量删除
const deleteBatch = async () => {
    if (multipleSelection.value.length === 0) {
        feedback.msgError('请选择要删除的数据')
        return
    }
    try {
        await feedback.confirm('确定要删除？')
        await system_corn_delete_batch({
            Ids: multipleSelection.value.map((item) => item.Id).join(',')
        })
        feedback.msgSuccess('删除成功')
        getLists()
    } catch (error) {}
}

const export_file = async () => {
    try {
        await feedback.confirm('确定要导出？')
        await system_corn_export_file(queryParams)
    } catch (error) {}
}
getLists()
</script>
