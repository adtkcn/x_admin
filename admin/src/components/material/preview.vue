<template>
    <div v-show="modelValue">
        <div v-if="type == 'image'">
            <el-image-viewer
                v-if="previewLists.length"
                :url-list="previewLists"
                hide-on-click-modal
                @close="handleClose"
            />
        </div>
        <div v-if="type == 'video'">
            <el-dialog v-model="visible" width="740px" title="视频预览" :before-close="handleClose">
                <video-player ref="playerRef" :src="url" width="100%" height="450px" />
            </el-dialog>
        </div>
    </div>
</template>

<script lang="ts" setup>
const props = defineProps({
    modelValue: {
        type: Boolean,
        default: false
    },
    url: {
        type: String,
        default: ''
    },
    type: {
        type: String,
        default: 'image'
    }
})

const playerRef = shallowRef()
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
