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
                    <el-input v-model="queryParams.projectKey" />
                </el-form-item>
                <el-form-item label="客户端id" prop="clientId" class="w-[280px]">
                    <el-input v-model="queryParams.clientId" />
                </el-form-item>
                <el-form-item label="用户id" prop="userId" class="w-[280px]">
                    <el-input v-model="queryParams.userId" />
                </el-form-item>

                <el-form-item label="城市" prop="city" class="w-[280px]">
                    <el-input v-model="queryParams.city" />
                </el-form-item>

                <el-form-item label="时间" prop="clientTime" class="w-[280px]">
                    <el-input v-model="queryParams.clientTime" />
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
            </div>
            <el-table class="mt-4" size="large" v-loading="pager.loading" :data="pager.lists">
                <el-table-column label="项目key" prop="projectKey" min-width="130" />
                <el-table-column label="客户端id" prop="clientId" min-width="130" />
                <el-table-column label="用户id" prop="userId" min-width="130" />
                <el-table-column label="系统" prop="os" min-width="130" />
                <el-table-column label="浏览器" prop="browser" min-width="130" />
                <el-table-column label="城市" prop="city" min-width="130" />
                <el-table-column label="屏幕" prop="width" min-width="130" />
                <el-table-column label="屏幕高度" prop="height" min-width="130" />
                <el-table-column label="ua记录" prop="ua" min-width="130" />
                <el-table-column label="客户端时间" prop="clientTime" min-width="130" />
                <!-- <el-table-column label="创建时间" prop="createTime" min-width="130" /> -->
                <el-table-column label="操作" width="160" fixed="right">
                    <template #default="{ row }">
                        <el-button type="primary" link @click="to_monitor_web(row)">日志</el-button>

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
import { monitor_client_delete, monitor_client_list } from '@/api/monitor_client'
import { usePaging } from '@/hooks/usePaging'
import feedback from '@/utils/feedback'
import EditPopup from './edit.vue'

const router = useRouter()

defineOptions({
    name: 'monitor_client'
})
const editRef = shallowRef<InstanceType<typeof EditPopup>>()
const showEdit = ref(false)
const queryParams = reactive({
    projectKey: '',
    clientId: '',
    userId: '',
    os: '',
    browser: '',
    city: '',
    width: '',
    height: '',
    ua: '',
    clientTime: ''
})

const { pager, getLists, resetPage, resetParams } = usePaging({
    fetchFun: monitor_client_list,
    params: queryParams
})

const handleAdd = async () => {
    showEdit.value = true
    await nextTick()
    editRef.value?.open('add')
}

const handleDelete = async (id: number) => {
    await feedback.confirm('确定要删除？')
    await monitor_client_delete({ id })
    feedback.msgSuccess('删除成功')
    getLists()
}
const to_monitor_web = async (data: any) => {
    router.push(
        `/setting/monitor_web/index?clientId=${data.clientId}&projectKey=${data.projectKey}`
    )
}
getLists()
</script>
