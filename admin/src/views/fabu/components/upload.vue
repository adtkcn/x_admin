<template>
    <div class="upload-popup">
        <popup
            ref="popupRef"
            title="上传安装包"
            width="520px"
            :confirm-button-text="false"
            @close="handleClose"
        >
            <div class="mb-3 text-xs text-info">
                上传后自动解析应用信息（名称 / BundleId / 平台 / 版本 / 图标），按 BundleId
                自动创建应用并生成首个版本；若 BundleId 已存在则仅追加新版本。
            </div>
            <upload-chunk :ext="['ipa', 'apk']" @change="handleUploaded" @error="handleError" />
        </popup>
    </div>
</template>
<script lang="ts" setup>
import { shallowRef } from 'vue'
import { fabuVersionUpload } from '@/api/fabu'
import Popup from '@/components/popup/index.vue'
import UploadChunk from '@/components/upload-chunk'
import type { ChunkUploadResult } from '@/components/upload-chunk'
import feedback from '@/utils/feedback'

const emit = defineEmits(['success', 'close'])
const popupRef = shallowRef<InstanceType<typeof Popup>>()

// 分片上传完成后，通知后端按文件引用解析并创建应用/版本
const handleUploaded = async (result: ChunkUploadResult) => {
    if (!result.fileHashId) {
        feedback.msgError('上传异常，未获取到文件标识')
        return
    }
    try {
        await fabuVersionUpload({
            file_hash_id: result.fileHashId,
            file_name: result.fileName
        })
        feedback.msgSuccess('应用创建成功')
        popupRef.value?.close()
        emit('success')
    } catch (error) {
        console.error(error)
    }
}
const handleError = (error: Error) => {
    console.error(error)
}
const open = () => {
    popupRef.value?.open()
}
const handleClose = () => emit('close')
defineExpose({ open })
</script>
