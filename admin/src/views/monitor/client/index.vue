<template>
    <div class="index-lists">
        <el-card class="border-none!" shadow="never">
            <el-form
                ref="formRef"
                class="mb-[-16px]"
                :model="queryParams"
                :inline="true"
                label-width="90px"
                label-position="right"
            >
                <el-form-item label="项目" prop="project_key" class="w-[270px]">
                    <el-select
                        v-model="queryParams.project_key"
                        clearable
                        :empty-values="[null, undefined]"
                    >
                        <el-option label="全部" value="" />
                        <el-option
                            v-for="(item, index) in listAllData.monitor_project_listAll"
                            :key="index"
                            :label="item.project_name"
                            :value="item.project_key"
                        />
                    </el-select>
                </el-form-item>

                <el-form-item label="系统" prop="os" class="w-[270px]">
                    <el-input v-model="queryParams.os" />
                </el-form-item>
                <el-form-item label="浏览器" prop="browser" class="w-[270px]">
                    <el-input v-model="queryParams.browser" />
                </el-form-item>

                <el-form-item label="创建时间" prop="create_time" class="w-[425px]">
                    <daterange-picker
                        v-model:startTime="queryParams.create_time_start"
                        v-model:endTime="queryParams.create_time_end"
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
                    v-perms="['admin:monitor_client:del_batch']"
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
                <vxe-column title="项目" field="project_key" min-width="100">
                    <template #default="{ row }">
                        <dict-value
                            :options="listAllData.monitor_project_listAll"
                            :value="row.project_key"
                            labelKey="project_name"
                            valueKey="project_key"
                        />
                    </template>
                </vxe-column>
                <vxe-column title="客户端id" field="client_id" min-width="130" />

                <vxe-column title="浏览器" field="browser" min-width="150">
                    <template #default="{ row }">
                        <el-popover
                            placement="top-start"
                            title="浏览器ua"
                            :width="500"
                            trigger="hover"
                            :content="row.ua"
                        >
                            <template #reference>
                                <el-link type="primary">{{ row.os }} / {{ row.browser }}</el-link>
                            </template>
                        </el-popover>
                    </template>
                </vxe-column>

                <vxe-column title="创建时间" field="create_time" width="180" />

                <vxe-column title="操作" width="80" fixed="right">
                    <template #default="{ row }">
                        <el-button
                            v-perms="['admin:monitor_client:del']"
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
    </div>
</template>
<script lang="ts" setup>
import { ref, reactive } from 'vue'
import {
    monitor_client_delete,
    monitor_client_delete_batch,
    monitor_client_list
} from '@/api/monitor/client'
import type { type_monitor_client, type_monitor_client_query } from '@/api/monitor/client'
import type { type_monitor_project } from '@/api/monitor/project'

import { useListAllData } from '@/hooks/useDictOptions'
// import type { type_dict } from '@/hooks/useDictOptions'

import { usePaging } from '@/hooks/usePaging'
import feedback from '@/utils/feedback'

defineOptions({
    name: 'monitor_client'
})

const queryParams = reactive<type_monitor_client_query>({
    project_key: undefined,
    os: undefined,
    browser: undefined,
    ua: undefined,
    create_time_start: undefined,
    create_time_end: undefined
})

const { pager, getLists, resetPage, resetParams } = usePaging<type_monitor_client>({
    fetchFun: monitor_client_list,
    params: queryParams
})
const { listAllData } = useListAllData<{
    monitor_project_listAll: type_monitor_project[]
}>({
    monitor_project_listAll: '/monitor_project/list_all'
})

const tableRef = ref<any>()
const multipleSelection = ref<type_monitor_client[]>([])

const handleDelete = async (id: string) => {
    try {
        await feedback.confirm('确定要删除？')
        await monitor_client_delete(id)
        feedback.msgSuccess('删除成功')
        getLists()
    } catch (error) {
        console.error('监控客户端删除失败:', error)
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
        await monitor_client_delete_batch({
            ids: multipleSelection.value.map((item) => item.id).join(',')
        })
        feedback.msgSuccess('删除成功')
        getLists()
    } catch (error) {
        console.error('监控客户端批量删除失败:', error)
    }
}
getLists()
</script>
