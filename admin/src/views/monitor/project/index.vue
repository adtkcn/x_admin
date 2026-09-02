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
                <el-form-item label="项目key" prop="project_key" class="w-[280px]">
                    <el-input v-model="queryParams.project_key" />
                </el-form-item>
                <el-form-item label="项目名称" prop="project_name" class="w-[280px]">
                    <el-input v-model="queryParams.project_name" />
                </el-form-item>
                <el-form-item label="项目类型" prop="project_type" class="w-[280px]">
                    <el-select
                        v-model="queryParams.project_type"
                        clearable
                        :empty-values="[null, undefined]"
                    >
                        <el-option label="全部" value="" />
                        <el-option
                            v-for="(item, index) in dictData.project_type"
                            :key="index"
                            :label="item.name"
                            :value="item.value"
                        />
                    </el-select>
                </el-form-item>
                <el-form-item label="是否启用" prop="status" class="w-[280px]">
                    <el-select
                        v-model="queryParams.status"
                        clearable
                        :empty-values="[null, undefined]"
                    >
                        <el-option label="全部" value="" />
                        <el-option
                            v-for="(item, index) in dictData.status"
                            :key="index"
                            :label="item.name"
                            :value="item.value"
                        />
                    </el-select>
                </el-form-item>
                <!-- <el-form-item label="创建时间" prop="create_time" class="w-[430px]">
                    <daterange-picker
                        v-model:startTime="queryParams.create_time_start"
                        v-model:endTime="queryParams.create_time_end"
                    />
                </el-form-item> -->
                <el-form-item label="更新时间" prop="update_time" class="w-[420px]">
                    <daterange-picker
                        v-model:startTime="queryParams.update_time_start"
                        v-model:endTime="queryParams.update_time_end"
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
                <ImportExcel
                    v-perms="['admin:monitor_project:import_file']"
                    class="ml-3 mr-3"
                    :url="monitor_project_import_file"
                    @change="resetPage"
                >
                    <el-button type="primary">
                        <template #icon>
                            <icon name="el-icon-Upload" />
                        </template>
                        导入
                    </el-button>
                </ImportExcel>
                <el-button
                    v-perms="['admin:monitor_project:export_file']"
                    type="primary"
                    @click="export_file"
                >
                    <template #icon>
                        <icon name="el-icon-Download" />
                    </template>
                    导出
                </el-button>
                <el-button
                    v-perms="['admin:monitor_project:del_batch']"
                    type="danger"
                    :disabled="!multipleSelection.length"
                    @click="deleteBatch"
                >
                    批量删除
                </el-button>
            </div>
            <vxe-table
                ref="tableRef"
                class="mt-4"
                v-loading="pager.loading"
                :data="pager.lists"
                :row-config="{ keyField: 'id' }"
                :checkbox-config="{ checkRowKeys: [] }"
                @checkbox-change="multipleSelection = $event.$table.getCheckboxRecords()"
                @checkbox-all="multipleSelection = $event.$table.getCheckboxRecords()"
                :border="'inner'"
            >
                <vxe-column type="checkbox" width="55" />
                <vxe-column title="项目uuid" field="project_key" min-width="130" />
                <vxe-column title="项目名称" field="project_name" min-width="130" />
                <vxe-column title="项目类型" field="project_type" width="100">
                    <template #default="{ row }">
                        <dict-value :options="dictData.project_type" :value="row.project_type" />
                    </template>
                </vxe-column>
                <vxe-column title="是否启用" field="status" width="100">
                    <template #default="{ row }">
                        <dict-value :options="dictData.status" :value="row.status" />
                    </template>
                </vxe-column>
                <!-- <vxe-column title="创建时间" field="create_time" min-width="130" /> -->
                <vxe-column title="更新时间" field="update_time" width="180" />
                <vxe-column title="操作" width="120" fixed="right">
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
                            @click="handleDelete(row.id)"
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
const tableRef = ref<any>()
const queryParams = reactive<type_monitor_project_query>({
    project_key: undefined,
    project_name: undefined,
    project_type: undefined,
    status: undefined,
    create_time_start: undefined,
    create_time_end: undefined,
    update_time_start: undefined,
    update_time_end: undefined
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

const handleDelete = async (id: string) => {
    try {
        await feedback.confirm('确定要删除？')
        await monitor_project_delete(id)
        feedback.msgSuccess('删除成功')
        getLists()
    } catch (error) {
        console.error('监控项目删除失败:', error)
    }
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
            ids: multipleSelection.value.map((item) => item.id).join(',')
        })
        feedback.msgSuccess('删除成功')
        getLists()
    } catch (error) {
        console.error('监控项目批量删除失败:', error)
    }
}

const export_file = async () => {
    try {
        await feedback.confirm('确定要导出？')
        await monitor_project_export_file(queryParams)
    } catch (error) {
        console.error('监控项目导出失败:', error)
    }
}
getLists()
</script>
