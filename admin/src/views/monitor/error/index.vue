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
                <el-form-item label="项目" prop="ProjectKey" class="w-[280px]">
                    <el-select
                        v-model="queryParams.ProjectKey"
                        clearable
                        :empty-values="[null, undefined]"
                    >
                        <el-option label="全部" value="" />
                        <el-option
                            v-for="(item, index) in listAllData.monitor_project_listAll"
                            :key="index"
                            :label="item.ProjectName"
                            :value="item.ProjectKey"
                        />
                    </el-select>
                </el-form-item>
                <!-- <el-form-item label="md5" prop="Md5" class="w-[280px]">
                    <el-input v-model="queryParams.Md5" />
                </el-form-item> -->
                <el-form-item label="创建时间" prop="CreateTime" class="w-[280px]">
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

            <el-table
                class="mt-4"
                size="large"
                v-loading="pager.loading"
                :data="pager.lists"
                @selection-change="handleSelectionChange"
            >
                <el-table-column type="selection" width="55" />
                <el-table-column label="序号" type="index" :index="handleIndex" width="80" />
                <el-table-column label="项目" prop="ProjectKey">
                    <template #default="{ row }">
                        <dict-value
                            :options="listAllData.monitor_project_listAll"
                            :value="row.ProjectKey"
                            labelKey="ProjectName"
                            valueKey="ProjectKey"
                        />
                    </template>
                </el-table-column>
                <el-table-column label="事件类型" prop="EventType" width="170" />
                <!-- <el-table-column label="URL地址" prop="Path" min-width="130" /> -->
                <el-table-column label="错误消息" prop="Message" min-width="150" />

                <!-- <el-table-column label="md5" prop="Md5" min-width="130" /> -->
                <el-table-column label="创建时间" prop="CreateTime" width="170" />

                <el-table-column label="操作" width="120" fixed="right">
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
    ProjectKey: undefined,
    EventType: undefined,
    Path: undefined,
    Message: undefined,
    Stack: undefined,
    Md5: undefined,
    CreateTimeStart: undefined,
    CreateTimeEnd: undefined
})

const { pager, getLists, resetPage, resetParams, handleIndex } = usePaging<type_monitor_error>({
    fetchFun: monitor_error_list,
    params: queryParams
})
const { listAllData } = useListAllData<{
    monitor_project_listAll: any[]
}>({
    monitor_project_listAll: '/monitor_project/list_all'
})

const editRef = shallowRef<InstanceType<typeof EditPopup>>()

const multipleSelection = ref<type_monitor_error[]>([])
const handleSelectionChange = (val: type_monitor_error[]) => {
    console.log(val)
    multipleSelection.value = val
}

const detailsRef = shallowRef<InstanceType<typeof DetailsPopup>>()
const handleDetails = async (row: type_monitor_error) => {
    showDetails.value = true
    await nextTick()
    detailsRef.value?.open('details')
    detailsRef.value?.getDetail(row)
}
const handleDelete = async (Id: number) => {
    try {
        await feedback.confirm('确定要删除？')
        await monitor_error_delete(Id)
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
        await monitor_error_delete_batch({
            Ids: multipleSelection.value.map((item) => item.Id).join(',')
        })
        feedback.msgSuccess('删除成功')
        getLists()
    } catch (error) {}
}

getLists()
</script>
