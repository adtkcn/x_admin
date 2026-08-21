<template>
    <div class="upload-dropzone">
        <div
            class="upload-dropzone__area"
            :class="{ 'is-dragover': dragover }"
            @click="triggerFileInput"
            @dragover.prevent="dragover = true"
            @dragleave.prevent="dragover = false"
            @drop.prevent="onDrop"
        >
            <input
                ref="fileInputRef"
                type="file"
                class="upload-dropzone__input"
                :accept="acceptValue"
                :multiple="multiple"
                @change="onFileChange"
            />
            <el-icon :size="40" color="#909399"><UploadFilled /></el-icon>
            <div class="upload-dropzone__text">点击或拖拽文件到此处上传</div>
            <div v-if="selectedFiles.length" class="upload-dropzone__hint">
                已选择 {{ selectedFiles.length }} 个文件，开始上传…
            </div>
        </div>

        <div v-if="md5Percent > 0 && md5Percent < 100" class="upload-dropzone__progress">
            <span class="upload-dropzone__label">计算文件指纹：</span>
            <el-progress :percentage="md5Percent" :stroke-width="10" />
        </div>

        <div v-if="uploading && !instantHit" class="upload-dropzone__progress">
            <span class="upload-dropzone__label">上传进度：</span>
            <el-progress :percentage="uploadPercent" :stroke-width="10" />
            <span class="upload-dropzone__detail"
                >{{ formatSize(uploadedSize) }} / {{ formatSize(totalSize) }}</span
            >
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { UploadFilled } from '@element-plus/icons-vue'
import SparkMD5 from 'spark-md5'
import axios from 'axios'
import type { PropType } from 'vue'
import config from '@/config'
import { getToken } from '@/utils/auth'
import feedback from '@/utils/feedback'
import { RequestCodeEnum } from '@/enums/requestEnums'

const props = defineProps({
    // 允许的后缀，如 ['png','jpg']
    ext: {
        type: Array as PropType<string[]>,
        default: () => []
    },
    multiple: {
        type: Boolean,
        default: true
    },
    limit: {
        type: Number,
        default: 10
    }
})

const emit = defineEmits<{
    (e: 'change', fileLists: any[]): void
    (e: 'error'): void
}>()

const fileInputRef = ref<HTMLInputElement>()
const selectedFiles = ref<File[]>([])
const uploading = ref(false)
const instantHit = ref(false)
const dragover = ref(false)
const md5Percent = ref(0)
const uploadPercent = ref(0)
const uploadedSize = ref(0)
const totalSize = ref(0)

const acceptValue = computed(() => {
    if (props.ext.length) {
        return props.ext.map((item) => `.${item.replace(/^\./, '')}`).join(',')
    }
    return ''
})

const effectiveLimit = computed(() => (props.limit && props.limit > 0 ? props.limit : 999))

const fullUrl = `${config.baseUrl}${config.urlPrefix}/common/upload/file`
const checkInstantUrl = `${config.baseUrl}${config.urlPrefix}/common/upload/checkInstant`

function formatSize(bytes: number): string {
    if (bytes < 1024) return bytes + ' B'
    if (bytes < 1024 * 1024) return (bytes / 1024).toFixed(1) + ' KB'
    if (bytes < 1024 * 1024 * 1024) return (bytes / 1024 / 1024).toFixed(1) + ' MB'
    return (bytes / 1024 / 1024 / 1024).toFixed(1) + ' GB'
}

function getExt(name: string) {
    const i = name.lastIndexOf('.')
    return i >= 0 ? name.slice(i + 1) : ''
}

function triggerFileInput() {
    fileInputRef.value?.click()
}

function onFileChange(e: Event) {
    const input = e.target as HTMLInputElement
    const files = input.files ? Array.from(input.files) : []
    input.value = ''
    if (files.length) addFiles(files)
}

function onDrop(e: DragEvent) {
    dragover.value = false
    const files = e.dataTransfer?.files ? Array.from(e.dataTransfer.files) : []
    if (files.length) addFiles(files)
}

function addFiles(files: File[]) {
    if (uploading.value) {
        feedback.msgError('正在上传中，请稍候')
        return
    }
    let next = [...selectedFiles.value, ...files]
    if (next.length > effectiveLimit.value) {
        feedback.msgError(`最多上传 ${effectiveLimit.value} 个文件`)
        next = next.slice(0, effectiveLimit.value)
    }
    selectedFiles.value = next
    if (next.length) startUpload()
}

function calcMD5(file: File): Promise<string> {
    return new Promise((resolve, reject) => {
        const spark = new SparkMD5.ArrayBuffer()
        const reader = new FileReader()
        const chunkSize = 10 * 1024 * 1024
        const total = Math.ceil(file.size / chunkSize) || 1
        let idx = 0
        const loadNext = () => {
            const start = idx * chunkSize
            reader.readAsArrayBuffer(file.slice(start, Math.min(start + chunkSize, file.size)))
        }
        reader.onload = () => {
            spark.append(reader.result as ArrayBuffer)
            idx++
            md5Percent.value = Math.floor((idx / total) * 100)
            if (idx < total) loadNext()
            else resolve(spark.end())
        }
        reader.onerror = () => reject(new Error('文件读取错误'))
        loadNext()
    })
}

async function checkInstant(md5: string, file_name: string) {
    try {
        const res = await axios.post(
            checkInstantUrl,
            { file_md5: md5, file_name },
            { headers: { token: getToken(), version: config.version } }
        )
        // 后端统一返回 { code, msg, show, data: CommonUploadFileResp }
        return res.data?.data || { instant: false }
    } catch (error) {
        console.error('秒传检查失败:', error)
        return { instant: false }
    }
}

async function uploadOne(file: File, completedBytes: number): Promise<any> {
    const formData = new FormData()
    formData.append('file', file)
    const resp = await axios.post(fullUrl, formData, {
        headers: { token: getToken(), version: config.version },
        onUploadProgress: (e) => {
            if (e.total) {
                uploadedSize.value = completedBytes + e.loaded
                uploadPercent.value = Math.min(
                    99,
                    Math.floor((uploadedSize.value / totalSize.value) * 100)
                )
            }
        }
    })
    return resp.data
}

async function startUpload() {
    const files = selectedFiles.value
    if (!files.length) return
    uploading.value = true
    instantHit.value = false
    md5Percent.value = 0
    uploadPercent.value = 0
    uploadedSize.value = 0
    totalSize.value = files.reduce((s, f) => s + f.size, 0)

    const results: any[] = []
    let completedBytes = 0
    for (let i = 0; i < files.length; i++) {
        const file = files[i]
        try {
            const md5 = await calcMD5(file)
            const instant = await checkInstant(md5, file.name)
            if (instant.instant && instant.file_hash_id) {
                instantHit.value = true
                results.push({
                    name: file.name,
                    response: {
                        data: {
                            file_hash_id: instant.file_hash_id,
                            name: file.name,
                            ext: getExt(file.name),
                            size: instant.size,
                            url: instant.url,
                            path: instant.path
                        }
                    }
                })
                completedBytes += file.size
                uploadedSize.value = completedBytes
                continue
            }
            const data = await uploadOne(file, completedBytes)
            if (data?.code === RequestCodeEnum.FAILED) {
                feedback.msgError(data.message || '上传失败')
                emit('error')
                uploading.value = false
                return
            }
            results.push({ name: file.name, response: data })
            completedBytes += file.size
            uploadedSize.value = completedBytes
            uploadPercent.value = Math.floor((completedBytes / totalSize.value) * 100)
        } catch (err: any) {
            feedback.msgError(err?.message || '上传失败')
            emit('error')
            uploading.value = false
            return
        }
    }
    uploading.value = false
    uploadPercent.value = 100
    selectedFiles.value = []
    emit('change', results)
}

defineExpose({ triggerFileInput })
</script>

<style scoped lang="scss">
.upload-dropzone {
    &__area {
        border: 2px dashed #dcdfe6;
        border-radius: 8px;
        padding: 40px 16px;
        text-align: center;
        cursor: pointer;
        transition: border-color 0.2s;
        &:hover,
        &.is-dragover {
            border-color: #409eff;
        }
    }
    &__input {
        display: none;
    }
    &__text {
        font-size: 14px;
        color: #606266;
        margin-top: 12px;
    }
    &__hint {
        margin-top: 8px;
        font-size: 13px;
        color: #909399;
    }
    &__progress {
        margin-top: 20px;
    }
    &__label {
        font-size: 13px;
        color: #606266;
    }
    &__detail {
        margin-left: 12px;
        font-size: 12px;
        color: #909399;
    }
}
</style>
