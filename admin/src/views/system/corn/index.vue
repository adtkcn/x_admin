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
                <el-form-item label="任务名称" prop="task_name" class="w-[280px]">
                    <el-input v-model="queryParams.task_name" />
                </el-form-item>
                <el-form-item label="任务编码" prop="task_code" class="w-[280px]">
                    <el-input v-model="queryParams.task_code" />
                </el-form-item>
                <el-form-item label="corn表达式" prop="corn_expr" class="w-[280px]">
                    <el-input v-model="queryParams.corn_expr" />
                </el-form-item>
                <el-form-item label="创建人" prop="created_by" class="w-[280px]">
                    <el-input v-model="queryParams.created_by" />
                </el-form-item>
                <el-form-item label="创建人名称" prop="nickname" class="w-[280px]">
                    <el-input v-model="queryParams.nickname" />
                </el-form-item>

                <el-form-item label="创建时间" prop="create_time" class="w-[280px]">
                    <daterange-picker
                        v-model:startTime="queryParams.create_time_start"
                        v-model:endTime="queryParams.create_time_end"
                    />
                </el-form-item>
                <el-form-item label="更新时间" prop="update_time" class="w-[280px]">
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
                <el-button v-perms="['admin:system_corn:add']" type="primary" @click="handleAdd()">
                    <template #icon>
                        <icon name="el-icon-Plus" />
                    </template>
                    新增
                </el-button>
                <Upload
                    v-perms="['admin:system_corn:import_file']"
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
                    v-perms="['admin:system_corn:export_file']"
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
            <vxe-table
                ref="tableRef"
                class="mt-4"
                v-loading="pager.loading"
                :data="pager.lists"
                :row-config="{ keyField: 'Id' }"
                :checkbox-config="{ checkRowKeys: [] }"
                @checkbox-change="multipleSelection = $event.$table.getCheckboxRecords()"
                @checkbox-all="multipleSelection = $event.$table.getCheckboxRecords()"
                :border="'inner'"
            >
                <vxe-column type="checkbox" width="55" />
                <vxe-column title="任务名称" field="task_name" min-width="130" />
                <vxe-column title="任务编码" field="task_code" min-width="130" />
                <vxe-column title="corn表达式" field="corn_expr" min-width="130" />
                <vxe-column title="状态" field="status" min-width="130">
                    <template #default="{ row }">
                        <dict-value :options="dictData.status" :value="row.status" />
                    </template>
                </vxe-column>
                <vxe-column title="创建人" field="created_by_user.nickname" min-width="130" />
                <vxe-column title="创建时间" field="create_time" min-width="130" />
                <vxe-column title="更新时间" field="update_time" min-width="130" />
                <vxe-column title="操作" width="160" fixed="right">
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
                </vxe-column>
            </vxe-table>
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
const tableRef = ref<any>()
const queryParams = reactive<type_system_corn_query>({
    task_name: undefined,
    task_code: undefined,
    corn_expr: undefined,
    status: undefined,
    created_by: undefined,
    nickname: undefined,
    create_time_start: undefined,
    create_time_end: undefined,
    update_time_start: undefined,
    update_time_end: undefined
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

const handleDelete = async (Id: string) => {
    try {
        await feedback.confirm('确定要删除？')
        await system_corn_delete(Id)
        feedback.msgSuccess('删除成功')
        getLists()
    } catch (error) {
        console.error('定时任务删除失败:', error)
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
        await system_corn_delete_batch({
            ids: multipleSelection.value.map((item) => item.id).join(',')
        })
        feedback.msgSuccess('删除成功')
        getLists()
    } catch (error) {
        console.error('定时任务批量删除失败:', error)
    }
}

const export_file = async () => {
    try {
        await feedback.confirm('确定要导出？')
        await system_corn_export_file(queryParams)
    } catch (error) {
        console.error('定时任务导出失败:', error)
    }
}
getLists()
</script>
