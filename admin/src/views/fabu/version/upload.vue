<template>
    <div class="upload-popup">
        <popup
            ref="popupRef"
            title="上传安装包"
            :async="true"
            width="520px"
            @confirm="handleSubmit"
            @close="handleClose"
        >
            <el-upload
                drag
                :auto-upload="false"
                :limit="1"
                v-model:file-list="fileList"
                accept=".ipa,.apk"
                :on-change="onChange"
                :on-exceed="onExceed"
            >
                <div class="el-upload__text">将 .ipa / .apk 拖到此处，或<em>点击上传</em></div>
                <template #tip>
                    <div class="text-xs text-info">
                        上传后自动解析包信息（名称 / BundleId / 版本 / 图标），按 BundleId 自动归属已有应用或新建应用。
                    </div>
                </template>
            </el-upload>
        </popup>
    </div>
</template>
<script lang="ts" setup>
import { ref, shallowRef } from 'vue'
import type { UploadUserFile, UploadFile } from 'element-plus'
import { fabuVersionUpload } from '@/api/fabu'
import Popup from '@/components/popup/index.vue'
import feedback from '@/utils/feedback'

const props = defineProps<{ appId?: string }>()
const emit = defineEmits(['success', 'close'])
const popupRef = shallowRef<InstanceType<typeof Popup>>()
const fileList = ref<UploadUserFile[]>([])
const fileRaw = ref<File | null>(null)

const onChange = (file: UploadFile) => {
    fileRaw.value = (file.raw as File) || null
}
const onExceed = () => {
    feedback.msgError('只能上传一个文件')
}
const handleSubmit = async () => {
    if (!fileRaw.value) {
        feedback.msgError('请先选择安装包')
        return
    }
    try {
        const fd = new FormData()
        fd.append('file', fileRaw.value)
        await fabuVersionUpload(fd)
        feedback.msgSuccess('上传成功')
        popupRef.value?.close()
        emit('success')
    } catch (error) {
        console.error(error)
    }
}
const open = () => {
    fileList.value = []
    fileRaw.value = null
    popupRef.value?.open()
}
const handleClose = () => emit('close')
defineExpose({ open })
</script>
