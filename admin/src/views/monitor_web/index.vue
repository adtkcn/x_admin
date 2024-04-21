<template>
    <div class="index-lists">
        <el-card class="!border-none" shadow="never">
            <el-form ref="formRef" class="mb-[-16px]" :model="queryParams" :inline="true">
                <el-form-item label="项目Key" prop="projectKey" class="w-[280px]">
                    <el-input v-model="queryParams.projectKey" />
                </el-form-item>
                <el-form-item label="客户端id" prop="clientId" class="w-[280px]">
                    <el-input v-model="queryParams.clientId" />
                </el-form-item>
                <!-- <el-form-item label="事件类型" prop="eventType" class="w-[280px]">
                    <el-select v-model="queryParams.eventType" clearable>
                        <el-option label="请选择字典生成" value="" />
                    </el-select>
                </el-form-item> -->
                <el-form-item label="时间" prop="clientTime" class="w-[400px]">
                    <daterange-picker
                        v-model:startTime="queryParams.clientTimeStart"
                        v-model:endTime="queryParams.clientTimeEnd"
                    />
                </el-form-item>
                <!-- <el-form-item label="创建时间" prop="createTime" class="w-[400px]">
                    <daterange-picker
                        v-model:startTime="queryParams.createTimeStart"
                        v-model:endTime="queryParams.createTimeEnd"
                    />
                </el-form-item> -->
                <el-form-item>
                    <el-button type="primary" @click="resetPage">查询</el-button>
                    <el-button @click="resetParams">重置</el-button>
                </el-form-item>
            </el-form>
        </el-card>
        <el-card class="!border-none mt-4" shadow="never">
            <div>
                <el-button v-perms="['admin:monitor_web:add']" type="primary" @click="handleAdd()">
                    <template #icon>
                        <icon name="el-icon-Plus" />
                    </template>
                    新增
                </el-button>
                <upload
                    class="ml-3 mr-3"
                    :url="monitor_web_import_file"
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
                <el-table-column type="expand">
                    <template #default="props">
                        <el-alert :title="props.row.stack" type="error" :closable="false" />
                    </template>
                </el-table-column>
                <el-table-column label="项目key" prop="projectKey" min-width="130" />
                <el-table-column label="客户端id" prop="clientId" min-width="130" />
                <el-table-column label="事件类型" prop="eventType" min-width="130" />
                <el-table-column label="URL地址" prop="page" min-width="130" />
                <el-table-column label="错误消息" prop="message" min-width="130" />

                <el-table-column label="时间" prop="clientTime" min-width="130" />
                <!-- <el-table-column label="创建时间" prop="createTime" min-width="130" /> -->
                <el-table-column label="操作" width="100" fixed="right">
                    <template #default="{ row }">
                        <el-button
                            v-perms="['admin:monitor_web:del']"
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
    monitor_web_delete,
    monitor_web_list,
    monitor_web_import_file,
    monitor_web_export_file
} from '@/api/monitor_web'
import { usePaging } from '@/hooks/usePaging'
import feedback from '@/utils/feedback'
import EditPopup from './edit.vue'

const route = useRoute()

defineOptions({
    name: 'monitor_web'
})
const editRef = shallowRef<InstanceType<typeof EditPopup>>()
const showEdit = ref(false)
const queryParams = reactive({
    projectKey: route.query?.projectKey as string,
    clientId: route.query?.clientId as string,
    eventType: '',
    page: '',
    message: '',
    stack: '',
    clientTimeStart: '',
    clientTimeEnd: '',
    createTimeStart: '',
    createTimeEnd: ''
})

const { pager, getLists, resetPage, resetParams } = usePaging({
    fetchFun: monitor_web_list,
    params: queryParams
})

const handleAdd = async () => {
    showEdit.value = true
    await nextTick()
    editRef.value?.open('add')
}

const handleDelete = async (id: number) => {
    try {
        await feedback.confirm('确定要删除？')
        await monitor_web_delete({ id })
        feedback.msgSuccess('删除成功')
        getLists()
    } catch (error) {}
}
const exportFile = async () => {
    try {
        await feedback.confirm('确定要导出？')
        await monitor_web_export_file(queryParams)
    } catch (error) {
        console.error(error)
    }
}
getLists()
</script>
