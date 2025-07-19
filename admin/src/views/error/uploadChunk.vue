<template>
    <div class="card">
        <input type="file" ref="fileInput" @change="handleChange" />
        <el-button @click="btn">上传</el-button>
        <el-button @click="merge">合并</el-button>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import FileUploader from '@/utils/FileUploader'

const fileInput = ref<HTMLInputElement>()

let fileUploader: FileUploader | null = null
function handleChange(e) {
    const files = (e.target as HTMLInputElement).files
    // console.log('e.target', e.target)
    console.log('files', files)
    if (files) {
        fileUploader = new FileUploader(files[0], { chunkSize: 1024 * 1024 * 1 })
    }
}
function btn() {
    fileUploader.start()
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
