<template>
    <div class="index-lists">
        <el-card class="!border-none" shadow="never">
            <el-form
                ref="formRef"
                class="mb-[-16px]"
                :model="queryParams"
                :inline="true"
                label-width="90px"
                label-position="left"
            >
                <el-form-item label="项目" prop="ProjectKey" class="w-[270px]">
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
                <!-- <el-form-item label="客户端id" prop="ClientId" class="w-[270px]">
                    <el-input v-model="queryParams.ClientId" />
                </el-form-item> -->
                <el-form-item label="用户id" prop="UserId" class="w-[270px]">
                    <el-input v-model="queryParams.UserId" />
                </el-form-item>
                <el-form-item label="系统" prop="Os" class="w-[270px]">
                    <el-input v-model="queryParams.Os" />
                </el-form-item>
                <el-form-item label="浏览器" prop="Browser" class="w-[270px]">
                    <el-input v-model="queryParams.Browser" />
                </el-form-item>
                <el-form-item label="国家" prop="Country" class="w-[270px]">
                    <el-input v-model="queryParams.Country" />
                </el-form-item>
                <el-form-item label="省份" prop="Province" class="w-[270px]">
                    <el-input v-model="queryParams.Province" />
                </el-form-item>
                <el-form-item label="城市" prop="City" class="w-[270px]">
                    <el-input v-model="queryParams.City" />
                </el-form-item>
                <el-form-item label="电信运营商" prop="Operator" class="w-[270px]">
                    <el-input v-model="queryParams.Operator" />
                </el-form-item>
                <el-form-item label="ip" prop="Ip" class="w-[270px]">
                    <el-input v-model="queryParams.Ip" />
                </el-form-item>
                <!-- <el-form-item label="ua记录" prop="Ua" class="w-[270px]">
                    <el-input v-model="queryParams.Ua" />
                </el-form-item> -->
                <el-form-item label="创建时间" prop="CreateTime" class="w-[425px]">
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
                <el-button
                    v-perms="['admin:monitor_client:delBatch']"
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
                <el-table-column label="项目" prop="ProjectKey" min-width="100">
                    <template #default="{ row }">
                        <dict-value
                            :options="listAllData.monitor_project_listAll"
                            :value="row.ProjectKey"
                            labelKey="ProjectName"
                            valueKey="ProjectKey"
                        />
                    </template>
                </el-table-column>
                <!-- <el-table-column label="客户端id" prop="ClientId" min-width="130" /> -->
                <el-table-column label="用户id" prop="UserId">
                    <template #default="{ row }">
                        <el-popover
                            placement="top-start"
                            :width="500"
                            trigger="hover"
                            :content="row.ClientId"
                        >
                            <template #reference>
                                <el-link type="primary">{{ row.UserId }}</el-link>
                            </template>
                            <div>用 户 ID ：{{ row.UserId }}</div>
                            <div>客户端ID：{{ row.ClientId }}</div>
                        </el-popover>
                    </template>
                </el-table-column>

                <el-table-column label="浏览器" prop="Browser" min-width="150">
                    <template #default="{ row }">
                        <el-popover
                            placement="top-start"
                            title="浏览器ua"
                            :width="500"
                            trigger="hover"
                            :content="row.Ua"
                        >
                            <template #reference>
                                <el-link type="primary">{{ row.Os }} / {{ row.Browser }}</el-link>
                            </template>
                        </el-popover>
                    </template>
                </el-table-column>
                <el-table-column label="IP" prop="Ip" width="145" />
                <el-table-column label="区域" min-width="150">
                    <template #default="{ row }">
                        {{ row.Country }}{{ row.Province }}{{ row.City }}
                    </template>
                </el-table-column>

                <!-- <el-table-column label="省份" prop="Province" />
                <el-table-column label="城市" prop="City" /> -->
                <el-table-column label="运营商" prop="Operator" />

                <!-- <el-table-column label="ua记录" prop="Ua" min-width="380" /> -->
                <el-table-column label="创建时间" prop="CreateTime" width="180" />

                <el-table-column label="操作" width="80" fixed="right">
                    <template #default="{ row }">
                        <!-- <el-button
                            v-perms="['admin:monitor_client:edit']"
                            type="primary"
                            link
                            @click="handleEdit(row)"
                        >
                            编辑
                        </el-button> -->
                        <el-button
                            v-perms="['admin:monitor_client:del']"
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
import { ref, shallowRef, reactive } from 'vue'
import {
    monitor_client_delete,
    monitor_client_delete_batch,
    monitor_client_list
} from '@/api/monitor/client'
import type { type_monitor_client, type_monitor_client_query } from '@/api/monitor/client'

import { useListAllData } from '@/hooks/useDictOptions'
// import type { type_dict } from '@/hooks/useDictOptions'

import { usePaging } from '@/hooks/usePaging'
import feedback from '@/utils/feedback'
import EditPopup from './edit.vue'
defineOptions({
    name: 'monitor_client'
})
const editRef = shallowRef<InstanceType<typeof EditPopup>>()
const showEdit = ref(false)
const queryParams = reactive<type_monitor_client_query>({
    ProjectKey: null,
    ClientId: null,
    UserId: null,
    Os: null,
    Browser: null,
    Country: null,
    Province: null,
    City: null,
    Operator: null,
    Ip: null,

    Ua: null,
    CreateTimeStart: null,
    CreateTimeEnd: null
})

const { pager, getLists, resetPage, resetParams } = usePaging<type_monitor_client>({
    fetchFun: monitor_client_list,
    params: queryParams
})
const { listAllData } = useListAllData<{
    monitor_project_listAll: any[]
}>({
    monitor_project_listAll: '/monitor_project/listAll'
})

const multipleSelection = ref<type_monitor_client[]>([])
const handleSelectionChange = (val: type_monitor_client[]) => {
    console.log(val)
    multipleSelection.value = val
}

const handleDelete = async (Id: number) => {
    try {
        await feedback.confirm('确定要删除？')
        await monitor_client_delete(Id)
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
        await monitor_client_delete_batch({
            Ids: multipleSelection.value.map((item) => item.Id).join(',')
        })
        feedback.msgSuccess('删除成功')
        getLists()
    } catch (error) {}
}
getLists()
</script>
