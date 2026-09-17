<template>
    <div class="fabu-wgt">
        <popup
            ref="popupRef"
            title="热更新包"
            width="800px"
            :confirm-button-text="false"
            :cancel-button-text="false"
            @close="handleClose"
        >
            <div class="mb-4">
                <el-upload
                    drag
                    :auto-upload="false"
                    :limit="1"
                    v-model:file-list="fileList"
                    accept=".wgt"
                    :on-change="onChange"
                    :on-exceed="onExceed"
                >
                    <div class="el-upload__text">将 .wgt 拖到此处，或<em>点击上传</em></div>
                    <template #tip>
                        <div class="text-xs text-info">
                            uni-app 热更新包，自动读取 manifest.json 中的版本信息，归属于当前版本。
                        </div>
                    </template>
                </el-upload>
                <el-button
                    class="mt-2"
                    type="primary"
                    :disabled="!fileRaw"
                    :loading="uploading"
                    @click="handleUpload"
                >
                    上传
                </el-button>
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
                    <template #default="{ row }">{{ formatSize(row.size) }}</template>
                </vxe-column>
                <vxe-column title="MD5" field="md5" min-width="200" />
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
import type { UploadUserFile, UploadFile } from 'element-plus'
import { fabuWgtLists, fabuWgtUpload, fabuWgtDel, type type_fabu_wgt_resp } from '@/api/fabu'
import Popup from '@/components/popup/index.vue'
import feedback from '@/utils/feedback'

const props = defineProps<{ versionId?: string }>()
const emit = defineEmits(['close'])
const popupRef = shallowRef<InstanceType<typeof Popup>>()
const fileList = ref<UploadUserFile[]>([])
const fileRaw = ref<File | null>(null)
const lists = ref<type_fabu_wgt_resp[]>([])
const loading = ref(false)
const uploading = ref(false)

const formatSize = (bytes: number) => {
    if (!bytes) return '0 B'
    const units = ['B', 'KB', 'MB', 'GB']
    const i = Math.floor(Math.log(bytes) / Math.log(1024))
    return `${(bytes / Math.pow(1024, i)).toFixed(2)} ${units[i]}`
}
const onChange = (file: UploadFile) => {
    fileRaw.value = (file.raw as File) || null
}
const onExceed = () => {
    feedback.msgError('只能上传一个文件')
}
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
const handleUpload = async () => {
    if (!fileRaw.value || !props.versionId) return
    uploading.value = true
    try {
        const fd = new FormData()
        fd.append('file', fileRaw.value)
        fd.append('version_id', props.versionId)
        await fabuWgtUpload(fd)
        feedback.msgSuccess('上传成功')
        fileList.value = []
        fileRaw.value = null
        getLists()
    } catch (error) {
        console.error(error)
    } finally {
        uploading.value = false
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
