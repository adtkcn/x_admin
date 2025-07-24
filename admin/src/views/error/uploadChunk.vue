<template>
    <div class="card">
        <input type="file" ref="fileInput" @change="handleChange" />
        <el-button type="primary" @click="btn">上传</el-button>
        <el-button type="primary" @click="cancel">取消</el-button>
        <el-button type="primary" @click="merge">合并</el-button>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import FileUploader from '@/utils/FileUploader'

const fileInput = ref<HTMLInputElement>()

const fileUploader = new FileUploader({
    chunkSize: 1024 * 1024 * 1,
    onSuccess(filePath) {
        ElMessage.success('上传成功:' + filePath)
    },
    onError(error) {
        // console.error('error', error)
        ElMessage.error(error.message)
    }
})
function handleChange(e) {
    const files = (e.target as HTMLInputElement).files
    // console.log('e.target', e.target)
    console.log('files', files)
    if (files) {
        fileUploader.loadFile(files[0])
    }
}
function btn() {
    fileUploader.start()
}
function cancel() {
    fileUploader.cancel()
}
function merge() {
    fileUploader.mergeChunk()
}
</script>

<style scoped>
button {
    color: #fff;
    background-color: rgb(150, 149, 149);
    border: 0;
    padding: 6px 10px;
    margin: 6px;
    font-size: 14px;
    line-height: 18px;
}
</style>
