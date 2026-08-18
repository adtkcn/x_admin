<template>
    <div class="p-4 upload-chunk-demo">
        <el-card>
            <template #header>
                <div class="flex items-center justify-between">
                    <span class="font-medium">分片上传 / 秒传（调试）</span>
                </div>
            </template>

            <UploadChunk
                ref="uploadChunkRef"
                bucket="files"
                @change="handleUploadResult"
                @error="handleError"
            >
                <template #actions="{ file }">
                    <span v-if="file" class="text-info mr-2">已选：{{ file.name }}</span>
                </template>
            </UploadChunk>
        </el-card>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import UploadChunk, { ChunkUploadResult } from '@/components/upload-chunk'
import feedback from '@/utils/feedback'

const uploadChunkRef = ref<InstanceType<typeof UploadChunk>>()

const handleUploadResult = async (result: ChunkUploadResult) => {
    if (!result.fileHashId) {
        feedback.msgSuccess('未生成文件哈希（秒传命中或注册失败），不挂载相册')
        return
    }
    feedback.msgSuccess('已上传')
}

const handleError = (e: Error) => {
    feedback.msgError('上传失败：' + (e?.message || '未知错误'))
}
</script>

<style scoped>
.upload-chunk-demo {
    max-width: 720px;
    margin: 0 auto;
}
</style>
