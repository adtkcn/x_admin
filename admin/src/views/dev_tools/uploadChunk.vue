<template>
    <div class="upload-chunk-wrapper">
        <div class="upload-card">
            <div class="upload-card__header">
                <h3>S3 分片上传（@aws-sdk/lib-storage + 秒传）</h3>
            </div>
            <div class="upload-card__body">
                <!-- 文件选择 -->
                <div class="upload-area" @click="triggerFileInput" @dragover.prevent @drop.prevent="handleDrop">
                    <input ref="fileInputRef" type="file" class="upload-area__input" @change="handleFileChange" />
                    <div class="upload-area__content">
                        <el-icon :size="40" color="#909399"><UploadFilled /></el-icon>
                        <div class="upload-area__text">点击或拖拽文件到此处上传</div>
                        <div v-if="selectedFile" class="upload-area__file">
                            {{ selectedFile.name }}（{{ formatSize(selectedFile.size) }}）
                        </div>
                    </div>
                </div>

                <!-- MD5 计算进度 -->
                <div v-if="md5Percent > 0 && md5Percent < 100" class="progress-section">
                    <span class="progress-label">计算文件指纹：</span>
                    <el-progress :percentage="md5Percent" :stroke-width="12" />
                </div>

                <!-- 上传进度 -->
                <div v-if="uploading && !instantHit" class="progress-section">
                    <span class="progress-label">上传进度：</span>
                    <el-progress :percentage="uploadPercent" :stroke-width="12" />
                    <span class="progress-detail">{{ formatSize(uploadedSize) }} / {{ formatSize(totalSize) }}</span>
                </div>

                <!-- 操作按钮 -->
                <div class="upload-actions">
                    <el-button type="primary" :loading="uploading" :disabled="!selectedFile" @click="startUpload">
                        {{ uploading ? '上传中...' : '开始上传' }}
                    </el-button>
                    <el-button type="danger" :disabled="!uploading" @click="cancelUpload">
                        取消上传
                    </el-button>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { UploadFilled } from '@element-plus/icons-vue'
import { Upload } from '@aws-sdk/lib-storage'
import SparkMD5 from 'spark-md5'
import s3Client from '@/utils/s3Client'
import axios from 'axios'

const fileInputRef = ref<HTMLInputElement>()
const selectedFile = ref<File | null>(null)
const uploading = ref(false)
const instantHit = ref(false)
const md5Percent = ref(0)
const uploadPercent = ref(0)
const uploadedSize = ref(0)
const totalSize = ref(0)

function formatSize(bytes: number): string {
    if (bytes < 1024) return bytes + ' B'
    if (bytes < 1024 * 1024) return (bytes / 1024).toFixed(1) + ' KB'
    if (bytes < 1024 * 1024 * 1024) return (bytes / 1024 / 1024).toFixed(1) + ' MB'
    return (bytes / 1024 / 1024 / 1024).toFixed(1) + ' GB'
}

function triggerFileInput() { fileInputRef.value?.click() }
function handleFileChange(e: Event) { const files = (e.target as HTMLInputElement).files; if (files?.[0]) selectedFile.value = files[0] }
function handleDrop(e: DragEvent) { const files = e.dataTransfer?.files; if (files?.[0]) selectedFile.value = files[0] }

/** 计算文件 MD5 */
function calcMD5(file: File): Promise<string> {
    return new Promise((resolve, reject) => {
        const spark = new SparkMD5.ArrayBuffer()
        const reader = new FileReader(); const chunk = 10 * 1024 * 1024
        let idx = 0; const total = Math.ceil(file.size / chunk)
        reader.onload = () => {
            spark.append(reader.result as ArrayBuffer); idx++
            md5Percent.value = Math.floor((idx / total) * 100)
            if (idx < total) loadNext(); else resolve(spark.end())
        }
        reader.onerror = () => reject(new Error('文件读取错误'))
        const loadNext = () => { const s = idx * chunk; reader.readAsArrayBuffer(file.slice(s, Math.min(s + chunk, file.size))) }
        loadNext()
    })
}

/** 前置：秒传检查 */
async function checkInstant(md5: string, fileName: string): Promise<{ filePath: string; url: string } | null> {
    try {
        const res = await axios.post('/api/admin/s3/checkInstant', { fileMd5: md5, fileName })
        if (res.data?.instant) return { filePath: res.data.filePath, url: res.data.url }
    } catch { /* ignore */ }
    return null
}

/** 后置：注册文件哈希 */
async function registerHash(md5: string, key: string, fileName: string, fileSize: number) {
    try {
        await axios.post('/api/admin/s3/registerHash', { fileMd5: md5, fileKey: key, fileName, fileSize })
    } catch { /* ignore */ }
}

/** 生成 S3 Key */
async function generateKey(fileName: string): Promise<string> {
    const res = await axios.post('/api/admin/s3/generateKey', { fileName })
    return res.data.data?.key || res.data.key
}

async function startUpload() {
    if (!selectedFile.value) { ElMessage.warning('请先选择文件'); return }
    const file = selectedFile.value
    uploading.value = true; instantHit.value = false
    md5Percent.value = 0; uploadPercent.value = 0; uploadedSize.value = 0; totalSize.value = file.size

    try {
        // 1. 计算 MD5
        const md5 = await calcMD5(file)

        // 2. 前置：秒传检查
        const instant = await checkInstant(md5, file.name)
        if (instant) {
            instantHit.value = true; uploading.value = false; md5Percent.value = 100
            ElMessage.success(`秒传命中！${instant.filePath}`)
            return
        }

        // 3. 生成 S3 Key
        const key = await generateKey(file.name)

        // 4. 标准 S3 上传（不注入任何自定义头）
        const upload = new Upload({
            client: s3Client,
            params: { Bucket: 'files', Key: key, Body: file },
            queueSize: 3, partSize: 5 * 1024 * 1024, leavePartsOnError: false,
        })
        upload.on('httpUploadProgress', (p) => {
            if (p.loaded && p.total) { uploadedSize.value = p.loaded; totalSize.value = p.total; uploadPercent.value = Math.floor((p.loaded / p.total) * 100) }
        })
        const result = await upload.done()

        // 5. 后置：注册文件哈希（供下次秒传使用）
        await registerHash(md5, key, file.name, file.size)

        uploading.value = false; uploadPercent.value = 100
        ElMessage.success(`上传成功！${result.Location || key}`)
    } catch (error: any) {
        uploading.value = false
        if (error.name === 'AbortError') ElMessage.info('上传已取消')
        else ElMessage.error('上传失败：' + (error.message || '未知错误'))
    }
}

function cancelUpload() { uploading.value = false; ElMessage.info('正在取消...') }
</script>

<style scoped>
.upload-chunk-wrapper { display: flex; justify-content: center; padding: 40px 16px; background: #f5f7fa; min-height: 100vh; }
.upload-card { width: 520px; background: #fff; border-radius: 8px; box-shadow: 0 2px 12px rgba(0,0,0,0.08); }
.upload-card__header { padding: 16px 24px; border-bottom: 1px solid #ebeef5; }
.upload-card__header h3 { margin: 0; font-size: 16px; color: #303133; }
.upload-card__body { padding: 24px; }
.upload-area { border: 2px dashed #dcdfe6; border-radius: 8px; padding: 40px 16px; text-align: center; cursor: pointer; transition: border-color .2s; }
.upload-area:hover { border-color: #409eff; }
.upload-area__input { display: none; }
.upload-area__content { display: flex; flex-direction: column; align-items: center; gap: 12px; color: #909399; }
.upload-area__text { font-size: 14px; }
.upload-area__file { margin-top: 8px; font-size: 13px; color: #606266; word-break: break-all; }
.progress-section { margin-top: 20px; }
.progress-label { font-size: 13px; color: #606266; }
.progress-detail { margin-left: 12px; font-size: 12px; color: #909399; }
.upload-actions { margin-top: 24px; display: flex; gap: 12px; }
</style>
