<template>
    <div class="index-lists">
        <el-card class="!border-none" shadow="never">
            <el-form
                ref="formRef"
                class="mb-[-16px]"
                :model="queryParams"
                :inline="true"
                label-width="70px"
                label-position="left"
            >
                <el-form-item label="项目key" prop="ProjectKey" class="w-[280px]">
                    <el-input v-model="queryParams.ProjectKey" />
                </el-form-item>
                <el-form-item label="项目名称" prop="ProjectName" class="w-[280px]">
                    <el-input v-model="queryParams.ProjectName" />
                </el-form-item>
                <el-form-item label="项目类型" prop="ProjectType" class="w-[280px]">
                    <el-select v-model="queryParams.ProjectType" clearable>
                        <el-option label="全部" value="" />
                        <el-option
                            v-for="(item, index) in dictData.project_type"
                            :key="index"
                            :label="item.name"
                            :value="item.value"
                        />
                    </el-select>
                </el-form-item>
                <el-form-item label="是否启用" prop="Status" class="w-[280px]">
                    <el-select v-model="queryParams.Status" clearable>
                        <el-option label="全部" value="" />
                        <el-option
                            v-for="(item, index) in dictData.status"
                            :key="index"
                            :label="item.name"
                            :value="item.value"
                        />
                    </el-select>
                </el-form-item>
                <!-- <el-form-item label="创建时间" prop="CreateTime" class="w-[430px]">
                    <daterange-picker
                        v-model:startTime="queryParams.CreateTimeStart"
                        v-model:endTime="queryParams.CreateTimeEnd"
                    />
                </el-form-item> -->
                <el-form-item label="更新时间" prop="UpdateTime" class="w-[420px]">
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
        <el-card class="!border-none mt-4" shadow="never">
            <div class="text-right">
                <el-button
                    v-perms="['admin:monitor_project:add']"
                    type="primary"
                    @click="handleAdd()"
                >
                    <template #icon>
                        <icon name="el-icon-Plus" />
                    </template>
                    新增
                </el-button>
                <upload
                    v-perms="['admin:monitor_project:ImportFile']"
                    class="ml-3 mr-3"
                    :url="monitor_project_import_file"
                    :data="{ cid: 0 }"
                    type="file"
                    :show-progress="true"
                    @change="resetPage"
                >
                    <el-button type="primary">
                        <template #icon>
                            <icon name="el-icon-Upload" />
                        </template>
                        导入
                    </el-button>
                </upload>
                <el-button
                    v-perms="['admin:monitor_project:ExportFile']"
                    type="primary"
                    @click="exportFile"
                >
                    <template #icon>
                        <icon name="el-icon-Download" />
                    </template>
                    导出
                </el-button>
                <el-button
                    v-perms="['admin:monitor_project:delBatch']"
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
                <el-table-column label="项目uuid" prop="ProjectKey" min-width="130" />
                <el-table-column label="项目名称" prop="ProjectName" min-width="130" />
                <el-table-column label="项目类型" prop="ProjectType" width="100">
                    <template #default="{ row }">
                        <dict-value :options="dictData.project_type" :value="row.ProjectType" />
                    </template>
                </el-table-column>
                <el-table-column label="是否启用" prop="Status" width="100">
                    <template #default="{ row }">
                        <dict-value :options="dictData.status" :value="row.Status" />
                    </template>
                </el-table-column>
                <!-- <el-table-column label="创建时间" prop="CreateTime" min-width="130" /> -->
                <el-table-column label="更新时间" prop="UpdateTime" width="180" />
                <el-table-column label="操作" width="120" fixed="right">
                    <template #default="{ row }">
                        <el-button
                            v-perms="['admin:monitor_project:edit']"
                            type="primary"
                            link
                            @click="handleEdit(row)"
                        >
                            编辑
                        </el-button>
                        <el-button
                            v-perms="['admin:monitor_project:del']"
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
        <edit-popup
            v-if="showEdit"
            ref="editRef"
            :dict-data="dictData"
            @success="getLists"
            @close="showEdit = false"
        />
    </div>
</template>
<script lang="ts" setup>
import { ref, shallowRef, reactive, nextTick } from 'vue'
import {
    monitor_project_delete,
    monitor_project_delete_batch,
    monitor_project_list,
    monitor_project_import_file,
    monitor_project_export_file
} from '@/api/monitor/project'
import type { type_monitor_project, type_monitor_project_query } from '@/api/monitor/project'

import { useDictData } from '@/hooks/useDictOptions'
import type { type_dict } from '@/hooks/useDictOptions'

import { usePaging } from '@/hooks/usePaging'
import feedback from '@/utils/feedback'
import EditPopup from './edit.vue'
defineOptions({
    name: 'monitor_project'
})
const editRef = shallowRef<InstanceType<typeof EditPopup>>()
const showEdit = ref(false)
const queryParams = reactive<type_monitor_project_query>({
    ProjectKey: null,
    ProjectName: null,
    ProjectType: null,
    Status: null,
    CreateTimeStart: null,
    CreateTimeEnd: null,
    UpdateTimeStart: null,
    UpdateTimeEnd: null
})

const { pager, getLists, resetPage, resetParams } = usePaging<type_monitor_project>({
    fetchFun: monitor_project_list,
    params: queryParams
})
const { dictData } = useDictData<{
    project_type: type_dict[]
    status: type_dict[]
}>(['project_type', 'status'])

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
const multipleSelection = ref<type_monitor_project[]>([])
const handleSelectionChange = (val: type_monitor_project[]) => {
    console.log(val)
    multipleSelection.value = val
}

const handleDelete = async (Id: number) => {
    try {
        await feedback.confirm('确定要删除？')
        await monitor_project_delete(Id)
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
        await monitor_project_delete_batch({
            Ids: multipleSelection.value.map((item) => item.Id).join(',')
        })
        feedback.msgSuccess('删除成功')
        getLists()
    } catch (error) {}
}

const exportFile = async () => {
    try {
        await feedback.confirm('确定要导出？')
        await monitor_project_export_file(queryParams)
    } catch (error) {}
}
getLists()
</script>
