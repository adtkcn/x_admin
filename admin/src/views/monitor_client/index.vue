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
                <el-form-item label="项目key" prop="projectKey" class="w-[280px]">
                    <el-select v-model="queryParams.projectKey" clearable>
                        <el-option label="请选择字典生成" value="" />
                    </el-select>
                </el-form-item>
                <el-form-item label="sdk生成的客户端id" prop="clientId" class="w-[280px]">
                    <el-input v-model="queryParams.clientId" />
                </el-form-item>
                <el-form-item label="用户id" prop="userId" class="w-[280px]">
                    <el-input v-model="queryParams.userId" />
                </el-form-item>
                <el-form-item label="系统" prop="os" class="w-[280px]">
                    <el-input v-model="queryParams.os" />
                </el-form-item>
                <el-form-item label="浏览器" prop="browser" class="w-[280px]">
                    <el-input v-model="queryParams.browser" />
                </el-form-item>
                <el-form-item label="城市" prop="city" class="w-[280px]">
                    <el-input v-model="queryParams.city" />
                </el-form-item>
                <el-form-item label="屏幕" prop="width" class="w-[280px]">
                    <el-input v-model="queryParams.width" />
                </el-form-item>
                <el-form-item label="屏幕高度" prop="height" class="w-[280px]">
                    <el-input v-model="queryParams.height" />
                </el-form-item>
                <el-form-item label="ua记录" prop="ua" class="w-[280px]">
                    <el-input v-model="queryParams.ua" />
                </el-form-item>
                <el-form-item label="创建时间" prop="createTime" class="w-[280px]">
                    <daterange-picker
                        v-model:startTime="queryParams.createTimeStart"
                        v-model:endTime="queryParams.createTimeEnd"
                    />
                </el-form-item>
                <el-form-item label="更新时间" prop="clientTime" class="w-[280px]">
                    <daterange-picker
                        v-model:startTime="queryParams.clientTimeStart"
                        v-model:endTime="queryParams.clientTimeEnd"
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
                    v-perms="['admin:monitor_client:add']"
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
                    :url="monitor_client_import_file"
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
                <el-table-column label="项目key" prop="projectKey" min-width="130" />
                <el-table-column label="sdk生成的客户端id" prop="clientId" min-width="130" />
                <el-table-column label="用户id" prop="userId" min-width="130" />
                <el-table-column label="系统" prop="os" min-width="130" />
                <el-table-column label="浏览器" prop="browser" min-width="130" />
                <el-table-column label="城市" prop="city" min-width="130" />
                <el-table-column label="屏幕" prop="width" min-width="130" />
                <el-table-column label="屏幕高度" prop="height" min-width="130" />
                <el-table-column label="ua记录" prop="ua" min-width="130" />
                <el-table-column label="创建时间" prop="createTime" min-width="130" />
                <el-table-column label="更新时间" prop="clientTime" min-width="130" />
                <el-table-column label="操作" width="120" fixed="right">
                    <template #default="{ row }">
                        <el-button
                            v-perms="['admin:monitor_client:edit']"
                            type="primary"
                            link
                            @click="handleEdit(row)"
                        >
                            编辑
                        </el-button>
                        <el-button
                            v-perms="['admin:monitor_client:del']"
                            type="danger"
                            link
                            @click="handleDelete(row.id)"
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
import {
    monitor_client_delete,
    monitor_client_list,
    monitor_client_import_file,
    monitor_client_export_file
} from '@/api/monitor_client'
import type { type_monitor_client, type_monitor_client_query } from '@/api/monitor_client'
import { usePaging } from '@/hooks/usePaging'
import feedback from '@/utils/feedback'
import EditPopup from './edit.vue'
defineOptions({
    name: 'monitor_client'
})
const editRef = shallowRef<InstanceType<typeof EditPopup>>()
const showEdit = ref(false)
const queryParams = reactive<type_monitor_client_query>({
    projectKey: '',
    clientId: '',
    userId: '',
    os: '',
    browser: '',
    city: '',
    width: null,
    height: null,
    ua: '',
    createTimeStart: '',
    createTimeEnd: '',
    clientTimeStart: '',
    clientTimeEnd: ''
})

const { pager, getLists, resetPage, resetParams } = usePaging<type_monitor_client>({
    fetchFun: monitor_client_list,
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

const handleDelete = async (id: number) => {
    try {
        await feedback.confirm('确定要删除？')
        await monitor_client_delete(id)
        feedback.msgSuccess('删除成功')
        getLists()
    } catch (error) {}
}
const exportFile = async () => {
    try {
        await feedback.confirm('确定要导出？')
        await monitor_client_export_file(queryParams)
    } catch (error) {}
}
getLists()
</script>
