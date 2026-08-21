<template>
    <div v-show="modelValue">
        <div v-if="fileType == 'image'">
            <el-image-viewer
                v-if="previewLists.length"
                :url-list="previewLists"
                hide-on-click-modal
                @close="handleClose"
            />
        </div>
        <div v-else>
            <el-dialog v-model="visible" width="900px" title="文件预览" :before-close="handleClose">
                <video-player
                    v-if="fileType == 'video' || fileType == 'audio'"
                    ref="playerRef"
                    :src="url"
                    width="100%"
                    height="450px"
                />
                <div v-else style="padding: 20px; text-align: center">
                    <p>
                        无法预览该文件类型，请下载后查看。<a :href="url" target="_blank" download
                            >点击下载</a
                        >
                    </p>
                </div>
            </el-dialog>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { ref, useTemplateRef, watch, computed, nextTick } from 'vue'
import { GetFileType } from '@/enums/fileEnums'

const props = defineProps({
    modelValue: {
        type: Boolean,
        default: false
    },
    url: {
        type: String,
        default: ''
    },
    // 文件扩展名（不含点）。id 形式的 url 无法从后缀推断类型，需显式传入。
    ext: {
        type: String,
        default: ''
    }
})

const fileType = computed(() => {
    // 优先用 ext prop 推断；空时回退到 url 后缀（兼容旧 url 形式）
    const key = props.ext || (props.url ? props.url.split('.').pop() : '')
    return GetFileType(key || '')
})

const playerRef = useTemplateRef('playerRef')
const visible = defineModel({
    type: Boolean
})

const handleClose = () => {
    visible.value = false
}

const previewLists = ref<any[]>([])

watch(
    () => props.modelValue,
    (value) => {
        if (value) {
            nextTick(() => {
                previewLists.value = [props.url]
                playerRef.value?.play()
            })
        } else {
            nextTick(() => {
                previewLists.value = []
                playerRef.value?.pause()
            })
        }
    }
)
</script>
