<template>
    <div class="fabu-app">
        <el-card class="border-none!" shadow="never">
            <el-form :model="queryParams" :inline="true" class="mb-[-16px]">
                <el-form-item label="关键词">
                    <el-input
                        v-model="queryParams.keyword"
                        clearable
                        placeholder="应用名/BundleId"
                        @keyup.enter="resetPage"
                    />
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="resetPage">查询</el-button>
                    <el-button @click="resetParams">重置</el-button>
                    <el-button type="primary" @click="handleUpload">上传安装包</el-button>
                </el-form-item>
            </el-form>
        </el-card>
        <el-card class="border-none! mt-4" shadow="never">
            <vxe-table
                class="mt-4"
                v-loading="pager.loading"
                :data="pager.lists"
                :row-config="{ keyField: 'id' }"
                :border="'inner'"
                show-overflow="title"
            >
                <vxe-column title="图标" min-width="80">
                    <template #default="{ row }">
                        <el-image
                            v-if="row.icon"
                            :src="row.icon"
                            fit="cover"
                            style="width: 40px; height: 40px; border-radius: 6px"
                        />
                        <span v-else>-</span>
                    </template>
                </vxe-column>
                <vxe-column title="应用名称" field="name" min-width="140" />
                <vxe-column title="平台" field="platform" min-width="90">
                    <template #default="{ row }">
                        <el-tag :type="row.platform === 'ios' ? 'warning' : 'success'">
                            {{ row.platform }}
                        </el-tag>
                    </template>
                </vxe-column>
                <vxe-column title="BundleId" field="bundle_id" min-width="200" />
                <vxe-column title="短链" field="short_url" min-width="120">
                    <template #default="{ row }">
                        <el-link
                            v-if="row.short_url"
                            type="primary"
                            :underline="false"
                            @click="handleOpenDownload(row)"
                        >
                            {{ row.short_url }}
                        </el-link>
                        <span v-else>-</span>
                    </template>
                </vxe-column>
                <vxe-column title="下载次数" field="download_times" min-width="100" />
                <vxe-column title="创建时间" field="create_time" min-width="170" />
                <vxe-column title="操作" width="120" fixed="right">
                    <template #default="{ row }">
                        <el-button type="primary" link @click="handleVersions(row)">版本</el-button>
                        <el-button type="danger" link @click="handleDel(row)">删除</el-button>
                    </template>
                </vxe-column>
            </vxe-table>
            <div class="flex justify-end mt-4">
                <pagination v-model="pager" @change="getLists" />
            </div>
        </el-card>
        <upload-popup
            v-if="showUpload"
            ref="uploadRef"
            @success="getLists"
            @close="showUpload = false"
        />
    </div>
</template>
<script lang="ts" setup>
import { ref, shallowRef, reactive, nextTick } from 'vue'
import { useRouter } from 'vue-router'
import {
    fabuAppLists,
    fabuAppDel,
    type type_fabu_app_list,
    type type_fabu_app_resp
} from '@/api/fabu'
import { usePaging } from '@/hooks/usePaging'
import feedback from '@/utils/feedback'
import UploadPopup from '../components/upload.vue'

defineOptions({ name: 'fabuApp' })
const router = useRouter()
const uploadRef = shallowRef<InstanceType<typeof UploadPopup>>()
const showUpload = ref(false)
const queryParams = reactive<type_fabu_app_list>({ keyword: '' })
const { pager, getLists, resetPage, resetParams } = usePaging({
    fetchFun: fabuAppLists,
    params: queryParams
})

const handleUpload = async () => {
    showUpload.value = true
    await nextTick()
    uploadRef.value?.open()
}
const handleVersions = (row: type_fabu_app_resp) => {
    router.push(`/fabu/version?appId=${row.id}`)
}
const handleOpenDownload = (row: type_fabu_app_resp) => {
    window.open(`${location.origin}/fabu/url/${row.short_url}`, '_blank')
}
const handleDel = async (row: type_fabu_app_resp) => {
    try {
        await feedback.confirm(`确定删除应用「${row.name}」？该应用下的版本与热更新包将一并删除。`)
        await fabuAppDel({ id: row.id })
        feedback.msgSuccess('删除成功')
        getLists()
    } catch (error) {
        console.error(error)
    }
}
getLists()
</script>
