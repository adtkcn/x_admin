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
                <el-form-item label="项目" prop="ProjectKey" class="w-[280px]">
                    <el-select v-model="queryParams.ProjectKey" clearable>
                        <el-option label="全部" value="" />
                        <el-option
                            v-for="(item, index) in listAllData.monitor_project_listAll"
                            :key="index"
                            :label="item.ProjectName"
                            :value="item.ProjectKey"
                        />
                    </el-select>
                </el-form-item>

                <el-form-item label="用户id" prop="UserId" class="w-[280px]">
                    <el-input v-model="queryParams.UserId" />
                </el-form-item>
                <el-form-item label="创建时间" prop="CreateTime" class="w-[420px]">
                    <daterange-picker
                        v-model:startTime="queryParams.CreateTimeStart"
                        v-model:endTime="queryParams.CreateTimeEnd"
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
                <!-- <el-button v-perms="['admin:monitor_slow:add']" type="primary" @click="handleAdd()">
                    <template #icon>
                        <icon name="el-icon-Plus" />
                    </template>
                    新增
                </el-button> -->
                <!-- <upload
                    v-perms="['admin:monitor_slow:ImportFile']"
                    class="ml-3 mr-3"
                    :url="monitor_slow_import_file"
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
                </upload> -->
                <!-- <el-button
                    v-perms="['admin:monitor_slow:ExportFile']"
                    type="primary"
                    @click="exportFile"
                >
                    <template #icon>
                        <icon name="el-icon-Download" />
                    </template>
                    导出
                </el-button> -->
                <el-button
                    v-perms="['admin:monitor_slow:delBatch']"
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
                <el-table-column label="项目" prop="ProjectKey" width="150">
                    <template #default="{ row }">
                        <dict-value
                            :options="listAllData.monitor_project_listAll"
                            :value="row.ProjectKey"
                            labelKey="ProjectName"
                            valueKey="ProjectKey"
                        />
                    </template>
                </el-table-column>

                <el-table-column label="用户id" prop="UserId" width="150" />
                <el-table-column label="URL地址" prop="Path" />
                <el-table-column label="时间(毫秒)" prop="Time" width="130" />
                <el-table-column label="创建时间" prop="CreateTime" width="180" />
                <el-table-column label="操作" width="80" fixed="right">
                    <template #default="{ row }">
                        <!-- <el-button
                            v-perms="['admin:monitor_slow:edit']"
                            type="primary"
                            link
                            @click="handleEdit(row)"
                        >
                            编辑
                        </el-button> -->
                        <el-button
                            v-perms="['admin:monitor_slow:del']"
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
        <edit-popup v-if="showEdit" ref="editRef" @success="getLists" @close="showEdit = false" />
    </div>
</template>
<script lang="ts" setup>
import { ref, reactive, shallowRef } from 'vue'
import {
    monitor_slow_delete,
    monitor_slow_delete_batch,
    monitor_slow_list
} from '@/api/monitor/slow'
import type { type_monitor_slow, type_monitor_slow_query } from '@/api/monitor/slow'

import { useListAllData } from '@/hooks/useDictOptions'
import { usePaging } from '@/hooks/usePaging'
import feedback from '@/utils/feedback'
import EditPopup from './edit.vue'
defineOptions({
    name: 'monitor_slow'
})
const editRef = shallowRef<InstanceType<typeof EditPopup>>()
const showEdit = ref(false)
const queryParams = reactive<type_monitor_slow_query>({
    ProjectKey: null,
    ClientId: null,
    UserId: null,
    Path: null,
    Time: null,
    CreateTimeStart: null,
    CreateTimeEnd: null
})

const { pager, getLists, resetPage, resetParams } = usePaging<type_monitor_slow>({
    fetchFun: monitor_slow_list,
    params: queryParams
})
const { listAllData } = useListAllData<{
    monitor_project_listAll: any[]
}>({
    monitor_project_listAll: '/monitor_project/listAll'
})

const multipleSelection = ref<type_monitor_slow[]>([])
const handleSelectionChange = (val: type_monitor_slow[]) => {
    console.log(val)
    multipleSelection.value = val
}

const handleDelete = async (Id: number) => {
    try {
        await feedback.confirm('确定要删除？')
        await monitor_slow_delete(Id)
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
        await monitor_slow_delete_batch({
            Ids: multipleSelection.value.map((item) => item.Id).join(',')
        })
        feedback.msgSuccess('删除成功')
        getLists()
    } catch (error) {}
}

getLists()
</script>
