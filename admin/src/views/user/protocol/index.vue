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
                <el-form-item label="标题" prop="Title" class="w-[280px]">
                    <el-input v-model="queryParams.Title" />
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
        <el-card class="!border-none mt-4" shadow="never">
            <div class="text-right">
                <el-button
                    v-perms="['admin:user_protocol:add']"
                    type="primary"
                    @click="handleAdd()"
                >
                    <template #icon>
                        <icon name="el-icon-Plus" />
                    </template>
                    新增
                </el-button>
                <upload
                    v-perms="['admin:user_protocol:ImportFile']"
                    class="ml-3 mr-3"
                    :url="user_protocol_import_file"
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
                    v-perms="['admin:user_protocol:ExportFile']"
                    type="primary"
                    @click="exportFile"
                >
                    <template #icon>
                        <icon name="el-icon-Download" />
                    </template>
                    导出
                </el-button>
                <el-button
                    v-perms="['admin:user_protocol:delBatch']"
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
                <el-table-column label="标题" prop="Title" min-width="130" />
                <el-table-column label="协议内容" prop="Content" min-width="130" />
                <el-table-column label="排序" prop="Sort" min-width="130" />
                <el-table-column label="创建时间" prop="CreateTime" min-width="130" />
                <el-table-column label="更新时间" prop="UpdateTime" min-width="130" />
                <el-table-column label="操作" width="120" fixed="right">
                    <template #default="{ row }">
                        <el-button
                            v-perms="['admin:user_protocol:edit']"
                            type="primary"
                            link
                            @click="handleEdit(row)"
                        >
                            编辑
                        </el-button>
                        <el-button
                            v-perms="['admin:user_protocol:del']"
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
import { ref, shallowRef, reactive, nextTick } from 'vue'
import {
    user_protocol_delete,
    user_protocol_delete_batch,
    user_protocol_list,
    user_protocol_import_file,
    user_protocol_export_file
} from '@/api/user/protocol'
import type { type_user_protocol, type_user_protocol_query } from '@/api/user/protocol'

import { usePaging } from '@/hooks/usePaging'
import feedback from '@/utils/feedback'
import EditPopup from './edit.vue'
defineOptions({
    name: 'user_protocol'
})
const editRef = shallowRef<InstanceType<typeof EditPopup>>()
const showEdit = ref(false)
const queryParams = reactive<type_user_protocol_query>({
    Title: null,
    Content: null,
    Sort: null,
    CreateTimeStart: null,
    CreateTimeEnd: null,
    UpdateTimeStart: null,
    UpdateTimeEnd: null
})

const { pager, getLists, resetPage, resetParams } = usePaging<type_user_protocol>({
    fetchFun: user_protocol_list,
    params: queryParams
})

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
const multipleSelection = ref<type_user_protocol[]>([])
const handleSelectionChange = (val: type_user_protocol[]) => {
    console.log(val)
    multipleSelection.value = val
}

const handleDelete = async (Id: number) => {
    try {
        await feedback.confirm('确定要删除？')
        await user_protocol_delete(Id)
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
        await user_protocol_delete_batch({
            Ids: multipleSelection.value.map((item) => item.Id).join(',')
        })
        feedback.msgSuccess('删除成功')
        getLists()
    } catch (error) {}
}

const exportFile = async () => {
    try {
        await feedback.confirm('确定要导出？')
        await user_protocol_export_file(queryParams)
    } catch (error) {}
}
getLists()
</script>
