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
                <el-form-item label="手机号码" prop="Mobile" class="w-[280px]">
                    <el-input v-model="queryParams.Mobile" />
                </el-form-item>
                <el-form-item
                    label="发送状态：[0=发送中, 1=发送成功, 2=发送失败]"
                    prop="Status"
                    class="w-[280px]"
                >
                    <el-select v-model="queryParams.Status" clearable>
                        <el-option label="全部" value="" />
                        <el-option
                            v-for="(item, index) in dictData.flow_apply_status"
                            :key="index"
                            :label="item.name"
                            :value="item.value"
                        />
                    </el-select>
                </el-form-item>
                <el-form-item label="发送时间" prop="SendTime" class="w-[280px]">
                    <daterange-picker
                        v-model:startTime="queryParams.SendTime_start"
                        v-model:endTime="queryParams.SendTime_end"
                    />
                </el-form-item>
                <el-form-item label="创建时间" prop="CreateTime" class="w-[280px]">
                    <daterange-picker
                        v-model:startTime="queryParams.CreateTime_start"
                        v-model:endTime="queryParams.CreateTime_end"
                    />
                </el-form-item>
                <el-form-item label="更新时间" prop="UpdateTime" class="w-[280px]">
                    <daterange-picker
                        v-model:startTime="queryParams.UpdateTime_start"
                        v-model:endTime="queryParams.UpdateTime_end"
                    />
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="resetPage">查询</el-button>
                    <el-button @click="resetParams">重置</el-button>
                </el-form-item>
            </el-form>
        </el-card>
        <el-card class="!border-none mt-4" shadow="never">
            <div>
                <el-button
                    v-perms="['admin:system_log_sms:add']"
                    type="primary"
                    @click="handleAdd()"
                >
                    <template #icon>
                        <icon name="el-icon-Plus" />
                    </template>
                    新增
                </el-button>
                <upload
                    class="ml-3 mr-3"
                    :url="system_log_sms_import_file"
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
                <el-button type="primary" @click="exportFile">
                    <template #icon>
                        <icon name="el-icon-Download" />
                    </template>
                    导出
                </el-button>
            </div>
            <el-table class="mt-4" size="large" v-loading="pager.loading" :data="pager.lists">
                <el-table-column label="场景编号" prop="Scene" min-width="130" />
                <el-table-column label="手机号码" prop="Mobile" min-width="130" />
                <el-table-column label="发送内容" prop="Content" min-width="130" />
                <el-table-column
                    label="发送状态：[0=发送中, 1=发送成功, 2=发送失败]"
                    prop="Status"
                    min-width="100"
                >
                    <template #default="{ row }">
                        <dict-value :options="dictData.flow_apply_status" :value="row.Status" />
                    </template>
                </el-table-column>
                <el-table-column label="短信结果" prop="Results" min-width="130" />
                <el-table-column label="发送时间" prop="SendTime" min-width="130" />
                <el-table-column label="创建时间" prop="CreateTime" min-width="130" />
                <el-table-column label="更新时间" prop="UpdateTime" min-width="130" />
                <el-table-column label="操作" width="120" fixed="right">
                    <template #default="{ row }">
                        <el-button
                            v-perms="['admin:system_log_sms:edit']"
                            type="primary"
                            link
                            @click="handleEdit(row)"
                        >
                            编辑
                        </el-button>
                        <el-button
                            v-perms="['admin:system_log_sms:del']"
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
import {
    system_log_sms_delete,
    system_log_sms_list,
    system_log_sms_import_file,
    system_log_sms_export_file
} from '@/api/system_log_sms'
import type { type_system_log_sms, type_system_log_sms_query } from '@/api/system_log_sms'

import { useDictData } from '@/hooks/useDictOptions'
import type { type_dict } from '@/hooks/useDictOptions'

import { usePaging } from '@/hooks/usePaging'
import feedback from '@/utils/feedback'
import EditPopup from './edit.vue'
defineOptions({
    name: 'system_log_sms'
})
const editRef = shallowRef<InstanceType<typeof EditPopup>>()
const showEdit = ref(false)
const queryParams = reactive<type_system_log_sms_query>({
    Scene: null,
    Mobile: null,
    Content: null,
    Status: null,
    Results: null,
    SendTime_start: null,
    SendTime_end: null,
    CreateTime_start: null,
    CreateTime_end: null,
    UpdateTime_start: null,
    UpdateTime_end: null
})

const { pager, getLists, resetPage, resetParams } = usePaging<type_system_log_sms>({
    fetchFun: system_log_sms_list,
    params: queryParams
})
const { dictData } = useDictData<{
    flow_apply_status: type_dict[]
}>(['flow_apply_status'])

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
    try {
        await feedback.confirm('确定要删除？')
        await system_log_sms_delete(id)
        feedback.msgSuccess('删除成功')
        getLists()
    } catch (error) {}
}
const exportFile = async () => {
    try {
        await feedback.confirm('确定要导出？')
        await system_log_sms_export_file(queryParams)
    } catch (error) {}
}
getLists()
</script>
