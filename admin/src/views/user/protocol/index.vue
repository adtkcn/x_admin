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
                <el-form-item label="标题" prop="title" class="w-[280px]">
                    <el-input v-model="queryParams.title" />
                </el-form-item>
                <el-form-item label="版本" prop="version" class="w-[280px]">
                    <el-input v-model="queryParams.version" />
                </el-form-item>

                <el-form-item label="创建时间" prop="create_time" class="w-[280px]">
                    <daterange-picker
                        v-model:startTime="queryParams.create_time_start"
                        v-model:endTime="queryParams.create_time_end"
                        type="daterange"
                    />
                </el-form-item>
                <el-form-item label="更新时间" prop="update_time" class="w-[280px]">
                    <daterange-picker
                        v-model:startTime="queryParams.update_time_start"
                        v-model:endTime="queryParams.update_time_end"
                        type="daterange"
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
                    v-perms="['admin:user_protocol:add']"
                    type="primary"
                    @click="handleAdd()"
                >
                    <template #icon>
                        <icon name="el-icon-Plus" />
                    </template>
                    新增
                </el-button>
                <Upload
                    v-perms="['admin:user_protocol:import_file']"
                    class="ml-3 mr-3"
                    :url="user_protocol_import_file"
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
                    v-perms="['admin:user_protocol:export_file']"
                    type="primary"
                    @click="export_file"
                >
                    <template #icon>
                        <icon name="el-icon-Download" />
                    </template>
                    导出
                </el-button>
                <el-button
                    v-perms="['admin:user_protocol:del_batch']"
                    type="danger"
                    :disabled="!multipleSelection.length"
                    @click="deleteBatch"
                >
                    批量删除
                </el-button>
            </div>
            <vxe-table
                ref="tableRef"
                :border="'inner'"
                size="medium"
                class="mt-4"
                :data="pager.lists"
                auto-resize
                @checkbox-change="handleSelectionChange"
                @checkbox-all="handleSelectionChange"
            >
                <vxe-column type="checkbox" width="55"></vxe-column>
                <vxe-column field="title" title="标题" min-width="130"></vxe-column>
                <vxe-column field="tag" title="标识" min-width="130"></vxe-column>
                <vxe-column field="version" title="版本" width="100"></vxe-column>
                <vxe-column
                    field="created_by_user.nickname"
                    title="创建人"
                    width="120"
                ></vxe-column>

                <vxe-column field="create_time" title="创建时间" width="180"></vxe-column>
                <vxe-column field="update_time" title="更新时间" width="180"></vxe-column>
                <vxe-column title="操作" width="160" fixed="right">
                    <template #default="{ row }">
                        <el-button
                            v-perms="['admin:user_protocol:detail']"
                            type="primary"
                            link
                            @click="viewDetails(row)"
                            >详情</el-button
                        >
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
                </vxe-column>
            </vxe-table>
            <div class="flex justify-end mt-4">
                <pagination v-model="pager" @change="getLists" />
            </div>
        </el-card>
        <edit-popup v-if="showEdit" ref="editRef" @success="getLists" @close="showEdit = false" />
        <DetailsPopup v-if="showDetails" ref="detailsRef" @close="showDetails = false" />
    </div>
</template>
<script lang="ts" setup>
import { ref, reactive, nextTick, useTemplateRef } from 'vue'
import {
    user_protocol_delete,
    user_protocol_delete_batch,
    user_protocol_list,
    user_protocol_import_file,
    user_protocol_export_file
} from '@/api/user/protocol'
import type { type_user_protocol, type_user_protocol_query } from '@/api/user/protocol'
// import type { VxeUI } from 'vxe-table'
import { VxeUI, VxeTableInstance, VxeTableEvents, VxeTablePropTypes } from 'vxe-table'

import { usePaging } from '@/hooks/usePaging'
import feedback from '@/utils/feedback'
import EditPopup from './edit.vue'
import DetailsPopup from './details.vue'
defineOptions({
    name: 'user_protocol'
})
const editRef = useTemplateRef<InstanceType<typeof EditPopup>>('editRef')
const showEdit = ref(false)
const detailsRef = useTemplateRef<InstanceType<typeof DetailsPopup>>('detailsRef')
const showDetails = ref(false)
const queryParams = reactive<type_user_protocol_query>({
    title: undefined,
    content: undefined,
    version: undefined,
    create_time_start: undefined,
    create_time_end: undefined,
    update_time_start: undefined,
    update_time_end: undefined
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
const viewDetails = async (data: any) => {
    showDetails.value = true
    await nextTick()
    detailsRef.value?.open()
    detailsRef.value?.getDetail(data)
}

const tableRef = useTemplateRef<VxeTableInstance<type_user_protocol>>('tableRef')
const multipleSelection = ref<type_user_protocol[]>([])
const handleSelectionChange = () => {
    if (tableRef.value) {
        multipleSelection.value = tableRef.value.getCheckboxRecords()
    }
}

const handleDelete = async (id: string) => {
    try {
        await feedback.confirm('确定要删除？')
        await user_protocol_delete(id)
        feedback.msgSuccess('删除成功')
        getLists()
    } catch (error) {
        console.error('协议删除失败:', error)
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
        await user_protocol_delete_batch({
            ids: multipleSelection.value.map((item) => item.id).join(',')
        })
        feedback.msgSuccess('删除成功')
        getLists()
    } catch (error) {
        console.error('协议批量删除失败:', error)
    }
}

const export_file = async () => {
    try {
        await feedback.confirm('确定要导出？')
        await user_protocol_export_file(queryParams)
    } catch (error) {
        console.error('协议导出失败:', error)
    }
}
getLists()
</script>
