<template>
    <div class="fabu-wgt">
        <popup
            ref="popupRef"
            title="热更新包"
            width="860px"
            :confirm-button-text="false"
            :cancel-button-text="false"
            @close="handleClose"
        >
            <div class="mb-4">
                <div class="mb-3 text-xs text-info">
                    选择 .wgt 后分片上传，自动读取 manifest.json 中的版本信息，归属于当前版本。
                </div>
                <upload-chunk :ext="['wgt']" @change="handleUploaded" @error="handleUploadError" />
            </div>
            <vxe-table
                v-loading="loading"
                :data="lists"
                :row-config="{ keyField: 'id' }"
                :border="'inner'"
                show-overflow="title"
            >
                <vxe-column title="版本" field="version" min-width="120" />
                <vxe-column title="版本Code" field="version_code" min-width="110" />
                <vxe-column title="大小" min-width="110">
                    <template #default="{ row }">{{ formatSize(row.size, 2) }}</template>
                </vxe-column>
                <vxe-column title="MD5" field="md5" min-width="200" />
                <vxe-column title="已发布" min-width="90">
                    <template #default="{ row }">
                        <el-switch
                            :model-value="row.released"
                            @change="handleRelease(row, $event as boolean)"
                        />
                    </template>
                </vxe-column>
                <vxe-column title="创建时间" field="create_time" min-width="170" />
                <vxe-column title="操作" width="160" fixed="right">
                    <template #default="{ row }">
                        <el-button type="primary" link @click="copyLink(row)">复制</el-button>
                        <el-button type="danger" link @click="handleDel(row)">删除</el-button>
                    </template>
                </vxe-column>
            </vxe-table>
        </popup>
    </div>
</template>
<script lang="ts" setup>
import { ref, shallowRef } from 'vue'
import {
    fabuWgtLists,
    fabuWgtUpload,
    fabuWgtRelease,
    fabuWgtDel,
    type type_fabu_wgt_resp
} from '@/api/fabu'
import Popup from '@/components/popup/index.vue'
import UploadChunk from '@/components/upload-chunk'
import type { ChunkUploadResult } from '@/components/upload-chunk'
import feedback from '@/utils/feedback'
import { formatSize } from '@/utils/file'

const props = defineProps<{ versionId?: string }>()
const emit = defineEmits(['close'])
const popupRef = shallowRef<InstanceType<typeof Popup>>()
const lists = ref<type_fabu_wgt_resp[]>([])
const loading = ref(false)
const getLists = async () => {
    if (!props.versionId) return
    loading.value = true
    try {
        const res = await fabuWgtLists({ version_id: props.versionId })
        lists.value = res.lists || []
    } catch (error) {
        console.error(error)
    } finally {
        loading.value = false
    }
}
// 分片上传完成后，通知后端按文件引用解析 manifest 并登记 wgt 记录
const handleUploaded = async (result: ChunkUploadResult) => {
    if (!props.versionId || !result.fileHashId) return
    try {
        await fabuWgtUpload({
            version_id: props.versionId,
            file_hash_id: result.fileHashId,
            file_name: result.fileName
        })
        feedback.msgSuccess('上传成功')
        getLists()
    } catch (error) {
        console.error(error)
    }
}
const handleUploadError = (error: Error) => {
    console.error(error)
}
// 切换发布状态（一个版本仅一个已发布包，发布时会取消同版本其它包）
const handleRelease = async (row: type_fabu_wgt_resp, released: boolean) => {
    const tip = released
        ? `确定发布热更包「${row.version} (${row.version_code})」？同版本下其它已发布包将被自动取消。`
        : `确定取消发布热更包「${row.version} (${row.version_code})」？`
    try {
        await feedback.confirm(tip)
    } catch {
        getLists() // 取消切换，回读列表使开关回弹
        return
    }
    try {
        await fabuWgtRelease({ id: row.id, released })
        feedback.msgSuccess(released ? '发布成功' : '已取消发布')
        getLists()
    } catch (error) {
        console.error(error)
        getLists()
    }
}
const copyLink = async (row: type_fabu_wgt_resp) => {
    const url = window.location.origin + row.download_url
    try {
        await navigator.clipboard.writeText(url)
        feedback.msgSuccess('下载链接已复制')
    } catch {
        feedback.msgError(url)
    }
}
const handleDel = async (row: type_fabu_wgt_resp) => {
    try {
        await feedback.confirm('确定删除该热更新包？')
        await fabuWgtDel({ id: row.id })
        feedback.msgSuccess('删除成功')
        getLists()
    } catch (error) {
        console.error(error)
    }
}
const open = () => {
    getLists()
    popupRef.value?.open()
}
const handleClose = () => emit('close')
defineExpose({ open })
</script>
