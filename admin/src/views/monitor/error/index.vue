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
                <el-form-item label="项目" prop="project_key" class="w-[280px]">
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
                <!-- <el-form-item label="md5" prop="md5" class="w-[280px]">
                    <el-input v-model="queryParams.md5" />
                </el-form-item> -->
                <el-form-item label="创建时间" prop="create_time" class="w-[280px]">
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
                    v-perms="['admin:monitor_error:del_batch']"
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
                <vxe-column type="seq" title="序号" width="80" />
                <vxe-column title="项目" field="project_key">
                    <template #default="{ row }">
                        <dict-value
                            :options="listAllData.monitor_project_listAll"
                            :value="row.project_key"
                            labelKey="project_name"
                            valueKey="project_key"
                        />
                    </template>
                </vxe-column>
                <vxe-column title="事件类型" field="event_type" width="170" />
                <!-- <vxe-column title="URL地址" field="path" min-width="130" /> -->
                <vxe-column title="错误消息" field="message" min-width="150" show-overflow />

                <!-- <vxe-column title="md5" field="md5" min-width="130" /> -->
                <vxe-column title="创建时间" field="create_time" width="170" />

                <vxe-column title="操作" width="120" fixed="right">
                    <template #default="{ row }">
                        <el-button
                            v-perms="['admin:monitor_error:detail']"
                            type="primary"
                            link
                            @click="handleDetails(row)"
                        >
                            详情
                        </el-button>
                        <el-button
                            v-perms="['admin:monitor_error:del']"
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
            :list-all-data="listAllData"
            @success="getLists"
            @close="showEdit = false"
        />
        <DetailsPopup
            v-if="showDetails"
            ref="detailsRef"
            :list-all-data="listAllData"
            @close="showDetails = false"
        />
    </div>
</template>
<script lang="ts" setup>
import { ref, reactive, shallowRef, nextTick } from 'vue'
import {
    monitor_error_delete,
    monitor_error_delete_batch,
    monitor_error_list
} from '@/api/monitor/error'
import type { type_monitor_error, type_monitor_error_query } from '@/api/monitor/error'

import { useListAllData } from '@/hooks/useDictOptions'
// import type { type_dict } from '@/hooks/useDictOptions'

import { usePaging } from '@/hooks/usePaging'
import feedback from '@/utils/feedback'
import EditPopup from './edit.vue'
import DetailsPopup from './details.vue'

defineOptions({
    name: 'monitor_error'
})

const showEdit = ref(false)
const showDetails = ref(false)

const queryParams = reactive<type_monitor_error_query>({
    project_key: undefined,
    event_type: undefined,
    path: undefined,
    message: undefined,
    stack: undefined,
    md5: undefined,
    create_time_start: undefined,
    create_time_end: undefined
})

const { pager, getLists, resetPage, resetParams } = usePaging<type_monitor_error>({
    fetchFun: monitor_error_list,
    params: queryParams
})
const { listAllData } = useListAllData<{
    monitor_project_listAll: any[]
}>({
    monitor_project_listAll: '/monitor_project/list_all'
})

const editRef = shallowRef<InstanceType<typeof EditPopup>>()
const tableRef = ref<any>()

const multipleSelection = ref<type_monitor_error[]>([])

const detailsRef = shallowRef<InstanceType<typeof DetailsPopup>>()
const handleDetails = async (row: type_monitor_error) => {
    showDetails.value = true
    await nextTick()
    detailsRef.value?.open('details')
    detailsRef.value?.getDetail(row)
}
const handleDelete = async (id: string) => {
    try {
        await feedback.confirm('确定要删除？')
        await monitor_error_delete(id)
        feedback.msgSuccess('删除成功')
        getLists()
    } catch (error) {
        console.error('监控错误删除失败:', error)
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
        await monitor_error_delete_batch({
            ids: multipleSelection.value.map((item) => item.id).join(',')
        })
        feedback.msgSuccess('删除成功')
        getLists()
    } catch (error) {
        console.error('监控错误批量删除失败:', error)
    }
}

getLists()
</script>
