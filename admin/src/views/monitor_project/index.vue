<template>
    <div class="index-lists">
        <el-card class="!border-none" shadow="never">
            <el-form ref="formRef" class="mb-[-16px]" :model="queryParams" :inline="true">
                <el-form-item label="项目Key" prop="projectKey">
                    <el-input v-model="queryParams.projectKey" />
                </el-form-item>
                <el-form-item label="项目名称" prop="projectName">
                    <el-input v-model="queryParams.projectName" />
                </el-form-item>
                <el-form-item label="项目类型" prop="projectType" class="w-[280px]">
                    <el-select v-model="queryParams.projectType" clearable>
                        <el-option label="全部" value="" />
                        <el-option
                            v-for="(item, index) in dictData.project_type"
                            :key="index"
                            :label="item.name"
                            :value="item.value"
                        />
                    </el-select>
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
                    v-perms="['admin:monitor_project:add']"
                    type="primary"
                    @click="handleAdd()"
                >
                    <template #icon>
                        <icon name="el-icon-Plus" />
                    </template>
                    新增
                </el-button>
            </div>
            <el-table class="mt-4" size="large" v-loading="pager.loading" :data="pager.lists">
                <el-table-column label="项目Key" prop="projectKey" min-width="130" />
                <el-table-column label="项目名称" prop="projectName" min-width="130" />
                <el-table-column label="项目类型" prop="projectType" min-width="100">
                    <template #default="{ row }">
                        <dict-value :options="dictData.project_type" :value="row.projectType" />
                    </template>
                </el-table-column>
                <el-table-column label="更新时间" prop="updateTime" min-width="130" />
                <!-- <el-table-column label="创建时间" prop="createTime" min-width="130" /> -->
                <el-table-column label="操作" width="240" fixed="right">
                    <template #default="{ row }">
                        <el-button type="primary" link @click="to_monitor_web(row)">日志</el-button>

                        <el-button
                            v-perms="['admin:monitor_project:edit']"
                            type="primary"
                            link
                            @click="handleEdit(row)"
                        >
                            编辑
                        </el-button>
                        <el-button
                            v-perms="['admin:monitor_project:del']"
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
import { monitor_project_delete, monitor_project_list } from '@/api/monitor_project'
import { useDictData } from '@/hooks/useDictOptions'
import { usePaging } from '@/hooks/usePaging'
import feedback from '@/utils/feedback'
import EditPopup from './edit.vue'
defineOptions({
    name: 'monitor_project'
})

const router = useRouter()

const editRef = shallowRef<InstanceType<typeof EditPopup>>()
const showEdit = ref(false)
const queryParams = reactive({
    projectKey: '',
    projectName: '',
    projectType: ''
})

const { pager, getLists, resetPage, resetParams } = usePaging({
    fetchFun: monitor_project_list,
    params: queryParams
})
const { dictData } = useDictData<{
    project_type: any[]
}>(['project_type'])

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
    await feedback.confirm('确定要删除？')
    await monitor_project_delete({ id })
    feedback.msgSuccess('删除成功')
    getLists()
}

const to_monitor_web = async (data: any) => {
    router.push('/setting/monitor_web/index?projectKey=' + data.projectKey)
}
getLists()
</script>
