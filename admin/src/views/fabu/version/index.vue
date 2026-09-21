<template>
    <div class="fabu-version">
        <el-card class="border-none!" shadow="never">
            <el-form :inline="true" class="mb-[-16px]">
                <el-form-item label="应用">
                    <el-select
                        v-model="appId"
                        filterable
                        placeholder="选择应用"
                        style="width: 320px"
                        @change="onAppChange"
                    >
                        <el-option
                            v-for="a in appOptions"
                            :key="a.id"
                            :label="`${a.name} (${a.platform})`"
                            :value="a.id"
                        />
                    </el-select>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" :disabled="!appId" @click="handleUpload"
                        >上传版本</el-button
                    >
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
                <vxe-column title="版本" field="version" min-width="120" />
                <vxe-column title="版本Code" field="version_code" min-width="110" />
                <vxe-column title="大小" min-width="110">
                    <template #default="{ row }">{{ formatSize(row.size, 2) }}</template>
                </vxe-column>
                <!-- <vxe-column title="MD5" field="md5" min-width="220" /> -->
                <vxe-column title="已发布" min-width="100">
                    <template #default="{ row }">
                        <el-tag :type="row.released ? 'success' : 'info'">
                            {{ row.released ? '已发布' : '未发布' }}
                        </el-tag>
                    </template>
                </vxe-column>
                <vxe-column title="灰度" min-width="90">
                    <template #default="{ row }">
                        <el-switch
                            :model-value="row.gray"
                            @change="(v: boolean | string | number) => handleGray(row, Boolean(v))"
                        />
                    </template>
                </vxe-column>
                <vxe-column title="更新模式" min-width="130">
                    <template #default="{ row }">
                        <el-select
                            :model-value="row.update_mode"
                            style="width: 110px"
                            @change="(v: number) => handleUpdateMode(row, v)"
                        >
                            <el-option label="普通" :value="0" />
                            <el-option label="强制" :value="1" />
                        </el-select>
                    </template>
                </vxe-column>
                <vxe-column title="下载次数" field="download_times" min-width="100" />
                <vxe-column title="创建时间" field="create_time" min-width="170" />
                <vxe-column title="操作" width="240" fixed="right">
                    <template #default="{ row }">
                        <el-button
                            v-if="!row.released"
                            type="primary"
                            link
                            @click="handleRelease(row)"
                        >
                            发布
                        </el-button>
                        <el-button v-else type="warning" link @click="handleCancel(row)"
                            >取消发布</el-button
                        >
                        <el-button type="primary" link @click="handleWgt(row)">热更新</el-button>
                        <el-button type="primary" link @click="copyLink(row)">下载</el-button>
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
        <wgt-popup
            v-if="showWgt"
            ref="wgtRef"
            :version-id="wgtVersionId"
            @close="showWgt = false"
        />
    </div>
</template>
<script lang="ts" setup>
import { ref, shallowRef, reactive, onMounted, nextTick } from 'vue'
import { useRoute } from 'vue-router'
import {
    fabuAppLists,
    fabuVersionLists,
    fabuVersionRelease,
    fabuVersionCancel,
    fabuVersionGray,
    fabuVersionUpdateMode,
    fabuVersionDel,
    type type_fabu_app_resp,
    type type_fabu_version_resp,
    type type_fabu_version_list
} from '@/api/fabu'
import { usePaging } from '@/hooks/usePaging'
import feedback from '@/utils/feedback'
import { formatSize } from '@/utils/file'
import UploadPopup from '../components/upload.vue'
import WgtPopup from './wgt.vue'

defineOptions({ name: 'fabuVersion' })
const route = useRoute()
const uploadRef = shallowRef<InstanceType<typeof UploadPopup>>()
const showUpload = ref(false)
const wgtRef = shallowRef<InstanceType<typeof WgtPopup>>()
const showWgt = ref(false)
const wgtVersionId = ref('')
const appId = ref<string>((route.query.appId as string) || '')
const queryParams = reactive<type_fabu_version_list>({ app_id: appId.value })
const appOptions = ref<type_fabu_app_resp[]>([])

const { pager, getLists, resetPage } = usePaging({
    fetchFun: fabuVersionLists,
    params: queryParams
})

const loadApps = async () => {
    try {
        const res = await fabuAppLists({ pageSize: 100 })
        appOptions.value = res.lists || []
    } catch (error) {
        console.error(error)
    }
}
const onAppChange = (val: string) => {
    queryParams.app_id = val
    resetPage()
}
const handleUpload = async () => {
    showUpload.value = true
    await nextTick()
    uploadRef.value?.open()
}
const handleWgt = async (row: type_fabu_version_resp) => {
    wgtVersionId.value = row.id
    showWgt.value = true
    await nextTick()
    wgtRef.value?.open()
}
const handleRelease = async (row: type_fabu_version_resp) => {
    try {
        await feedback.confirm('确定发布该版本为当前线上版本？')
        await fabuVersionRelease({ app_id: appId.value, id: row.id })
        feedback.msgSuccess('发布成功')
        getLists()
    } catch (error) {
        console.error(error)
    }
}
const handleCancel = async (row: type_fabu_version_resp) => {
    try {
        await feedback.confirm('确定取消发布该版本？')
        await fabuVersionCancel({ app_id: appId.value, id: row.id })
        feedback.msgSuccess('已取消发布')
        getLists()
    } catch (error) {
        console.error(error)
    }
}
const handleGray = async (row: type_fabu_version_resp, val: boolean) => {
    try {
        await fabuVersionGray({ app_id: appId.value, id: row.id, gray: val })
        feedback.msgSuccess('操作成功')
        getLists()
    } catch (error) {
        console.error(error)
    }
}
const handleUpdateMode = async (row: type_fabu_version_resp, val: number) => {
    try {
        await fabuVersionUpdateMode({ app_id: appId.value, id: row.id, update_mode: val })
        feedback.msgSuccess('操作成功')
        getLists()
    } catch (error) {
        console.error(error)
    }
}
const copyLink = async (row: type_fabu_version_resp) => {
    const url = window.location.origin + row.download_url
    try {
        await navigator.clipboard.writeText(url)
        feedback.msgSuccess('下载链接已复制')
    } catch {
        feedback.msgError(url)
    }
}
const handleDel = async (row: type_fabu_version_resp) => {
    try {
        await feedback.confirm('确定删除该版本？')
        await fabuVersionDel({ app_id: appId.value, id: row.id })
        feedback.msgSuccess('删除成功')
        getLists()
    } catch (error) {
        console.error(error)
    }
}
onMounted(() => {
    loadApps()
    if (appId.value) getLists()
})
</script>
