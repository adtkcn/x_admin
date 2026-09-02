<template>
    <div class="upload-chunk">
        <!-- 文件选择 -->
        <div
            class="upload-chunk__area"
            @click="triggerFileInput"
            @dragover.prevent
            @drop.prevent="handleDrop"
        >
            <input
                ref="fileInputRef"
                type="file"
                class="upload-chunk__input"
                :accept="acceptValue"
                @change="handleFileChange"
            />
            <div class="upload-chunk__content">
                <el-icon :size="40" color="#909399"><UploadFilled /></el-icon>
                <div class="upload-chunk__text">点击或拖拽文件到此处上传</div>
                <div v-if="selectedFile" class="upload-chunk__file">
                    {{ selectedFile.name }}（{{ formatSize(selectedFile.size) }}）
                </div>
            </div>
        </div>

        <!-- MD5 计算进度 -->
        <div v-if="md5Percent > 0 && md5Percent < 100" class="upload-chunk__progress">
            <span class="upload-chunk__label">计算文件指纹：</span>
            <el-progress :percentage="md5Percent" :stroke-width="12" />
        </div>

        <!-- 上传进度 -->
        <div v-if="uploading && !instantHit" class="upload-chunk__progress">
            <span class="upload-chunk__label">上传进度：</span>
            <el-progress :percentage="uploadPercent" :stroke-width="12" />
            <span class="upload-chunk__detail"
                >{{ formatSize(uploadedSize) }} / {{ formatSize(totalSize) }}</span
            >
        </div>

        <!-- 操作按钮 -->
        <div class="upload-chunk__actions">
            <slot name="actions" :file="selectedFile" :uploading="uploading"></slot>
            <el-button
                type="primary"
                :loading="uploading"
                :disabled="!selectedFile"
                @click="startUpload"
            >
                {{ uploading ? '上传中...' : '开始上传' }}
            </el-button>
            <el-button type="danger" :disabled="!uploading" @click="cancelUpload">
                取消上传
            </el-button>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import type { PropType } from 'vue'
import { ElMessage } from 'element-plus'
import { UploadFilled } from '@element-plus/icons-vue'
import { Upload } from '@aws-sdk/lib-storage'
import SparkMD5 from 'spark-md5'
import s3Client from '@/utils/s3Client'
import axios from 'axios'
import type { ChunkUploadResult } from './type'

const props = defineProps({
    // S3 Bucket 名称
    bucket: {
        type: String,
        default: 'files'
    },
    // 接收的文件类型（input accept）
    accept: {
        type: String,
        default: ''
    },
    // 允许的文件扩展名，如 ['png', 'jpg']（与 upload 组件保持一致）
    ext: {
        type: Array as PropType<string[]>,
        default: () => []
    },
    // S3 管理接口前缀（秒传检查 / 生成 Key / 注册哈希）
    endpoint: {
        type: String,
        default: '/api/admin/upload_chunk'
    },
    // 分片大小
    partSize: {
        type: Number,
        default: 5 * 1024 * 1024
    },
    // 并发队列大小
    queueSize: {
        type: Number,
        default: 3
    }
})

const emit = defineEmits<{
    (e: 'change', result: ChunkUploadResult): void
    (e: 'error', error: Error): void
}>()

const fileInputRef = ref<HTMLInputElement>()
const selectedFile = ref<File | null>(null)
const uploading = ref(false)
const instantHit = ref(false)
const md5Percent = ref(0)
const uploadPercent = ref(0)
const uploadedSize = ref(0)
const totalSize = ref(0)
let currentUpload: Upload | null = null

const acceptValue = computed(() => {
    if (props.ext.length) {
        return props.ext.map((item) => `.${item.replace(/^\./, '')}`).join(',')
    }
    return props.accept
})

function formatSize(bytes: number): string {
    if (bytes < 1024) return bytes + ' B'
    if (bytes < 1024 * 1024) return (bytes / 1024).toFixed(1) + ' KB'
    if (bytes < 1024 * 1024 * 1024) return (bytes / 1024 / 1024).toFixed(1) + ' MB'
    return (bytes / 1024 / 1024 / 1024).toFixed(1) + ' GB'
}

function triggerFileInput() {
    fileInputRef.value?.click()
}
function handleFileChange(e: Event) {
    const files = (e.target as HTMLInputElement).files
    if (files?.[0]) selectedFile.value = files[0]
}
function handleDrop(e: DragEvent) {
    const files = e.dataTransfer?.files
    if (files?.[0]) selectedFile.value = files[0]
}

/** 计算文件 MD5 */
function calcMD5(file: File): Promise<string> {
    return new Promise((resolve, reject) => {
        const spark = new SparkMD5.ArrayBuffer()
        const reader = new FileReader()
        const chunk = 10 * 1024 * 1024
        let idx = 0
        const total = Math.ceil(file.size / chunk)
        reader.onload = () => {
            spark.append(reader.result as ArrayBuffer)
            idx++
            md5Percent.value = Math.floor((idx / total) * 100)
            if (idx < total) loadNext()
            else resolve(spark.end())
        }
        reader.onerror = () => reject(new Error('文件读取错误'))
        const loadNext = () => {
            const s = idx * chunk
            reader.readAsArrayBuffer(file.slice(s, Math.min(s + chunk, file.size)))
        }
        loadNext()
    })
}

/** 前置：秒传检查（返回统一 CommonUploadFileResp 的解包结构） */
async function checkInstant(
    md5: string,
    fileName: string
): Promise<{ fileHashId: string | null; url: string } | null> {
    try {
        const res = await axios.post(`${props.endpoint}/checkInstant`, {
            file_md5: md5,
            file_name: fileName
        })
        const data = res.data?.data
        console.log('checkInstant res:', res)
        if (data?.instant) {
            return {
                url: data.url, // 完整访问地址
                fileHashId: data.file_hash_id ?? null
            }
        }
    } catch (error) {
        console.error('秒传检查失败:', error)
    }
    return null
}

/** 后置：注册文件哈希，返回统一 CommonUploadFileResp 的解包结构（file_hash_id + url） */
async function registerHash(
    md5: string,
    key: string,
    fileName: string,
    fileSize: number
): Promise<{ fileHashId: string | null; url: string }> {
    try {
        const res = await axios.post(`${props.endpoint}/registerHash`, {
            file_md5: md5,
            file_key: key,
            file_name: fileName,
            file_size: fileSize
        })
        console.log('registerHash res:', res)
        const data = res.data?.data
        return { fileHashId: data?.file_hash_id ?? null, url: data?.url ?? '' }
    } catch (error) {
        console.error('文件哈希注册失败:', error)
        return { fileHashId: null, url: '' }
    }
}

/** 生成 S3 Key */
async function generateKey(fileName: string): Promise<string> {
    try {
        const res = await axios.post(`${props.endpoint}/generateKey`, { file_name: fileName })
        return res.data.data?.key || res.data.key
    } catch (error) {
        console.error('生成S3 Key失败:', error)
        return ''
    }
}

async function startUpload() {
    if (!selectedFile.value) {
        ElMessage.warning('请先选择文件')
        return
    }
    const file = selectedFile.value
    uploading.value = true
    instantHit.value = false
    md5Percent.value = 0
    uploadPercent.value = 0
    uploadedSize.value = 0
    totalSize.value = file.size

    try {
        // 1. 计算 MD5
        const md5 = await calcMD5(file)

        // 2. 前置：秒传检查
        const instant = await checkInstant(md5, file.name)
        if (instant) {
            instantHit.value = true
            uploading.value = false
            md5Percent.value = 100
            ElMessage.success(`秒传命中！${instant.url}`)
            emit('change', {
                md5,
                key: '',
                fileName: file.name,
                fileSize: file.size,
                fileHashId: instant.fileHashId,
                instant: true,
                location: instant.url
            })
            return
        }

        // 3. 生成 S3 Key
        const key = await generateKey(file.name)

        // 4. 标准 S3 上传（不注入任何自定义头）
        const upload = new Upload({
            client: s3Client,
            params: { Bucket: props.bucket, Key: key, Body: file },
            queueSize: props.queueSize,
            partSize: props.partSize,
            leavePartsOnError: false
        })
        currentUpload = upload
        upload.on('httpUploadProgress', (p) => {
            if (p.loaded && p.total) {
                uploadedSize.value = p.loaded
                totalSize.value = p.total
                uploadPercent.value = Math.floor((p.loaded / p.total) * 100)
            }
        })
        const result = await upload.done()
        currentUpload = null

        // 5. 后置：注册文件哈希（供下次秒传使用），返回统一结构（file_hash_id + url）
        const hashRes = await registerHash(md5, key, file.name, file.size)

        uploading.value = false
        uploadPercent.value = 100
        ElMessage.success(`上传成功！${hashRes.url}`)
        emit('change', {
            md5,
            key,
            fileName: file.name,
            fileSize: file.size,
            fileHashId: hashRes.fileHashId,
            instant: false,
            location: hashRes.url
        })
    } catch (error: any) {
        uploading.value = false
        currentUpload = null
        if (error.name === 'AbortError') ElMessage.info('上传已取消')
        else {
            ElMessage.error('上传失败：' + (error.message || '未知错误'))
            emit('error', error)
        }
    }
}

function cancelUpload() {
    currentUpload?.abort()
    currentUpload = null
    uploading.value = false
    ElMessage.info('正在取消...')
}

defineExpose({ startUpload, cancelUpload, triggerFileInput })
</script>

<style scoped>
.upload-chunk__area {
    border: 2px dashed #dcdfe6;
    border-radius: 8px;
    padding: 40px 16px;
    text-align: center;
    cursor: pointer;
    transition: border-color 0.2s;
}
.upload-chunk__area:hover {
    border-color: #409eff;
}
.upload-chunk__input {
    display: none;
}
.upload-chunk__content {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 12px;
    color: #909399;
}
.upload-chunk__text {
    font-size: 14px;
}
.upload-chunk__file {
    margin-top: 8px;
    font-size: 13px;
    color: #606266;
    word-break: break-all;
}
.upload-chunk__progress {
    margin-top: 20px;
}
.upload-chunk__label {
    font-size: 13px;
    color: #606266;
}
.upload-chunk__detail {
    margin-left: 12px;
    font-size: 12px;
    color: #909399;
}
.upload-chunk__actions {
    margin-top: 24px;
    display: flex;
    gap: 12px;
}
</style>
