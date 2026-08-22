<template>
    <div>
        <div class="file-item relative" :style="{ height: fileSize, width: fileSize }">
            <el-image
                class="image"
                v-if="fileType == 'image'"
                fit="contain"
                lazy
                :src="uri"
            ></el-image>
            <video class="video" v-else-if="fileType == 'video'" :src="uri"></video>
            <div
                v-if="fileType == 'video'"
                class="absolute left-1/2 top-1/2 rounded-full w-[30px] h-[30px] flex justify-center items-center bg-[rgba(0,0,0,0.3)]"
                style="transform: translate(-50%, -50%)"
            >
                <icon name="el-icon-CaretRight" :size="18" color="#fff" />
            </div>
            <div
                v-if="fileType == 'audio'"
                class="absolute left-1/2 top-1/2 rounded-full w-[30px] h-[30px] flex justify-center items-center bg-[rgba(0,0,0,0.3)]"
                style="transform: translate(-50%, -50%)"
            >
                <icon name="el-icon-CaretRight" :size="16" color="#fff" />
            </div>
            <div v-if="fileType == 'office' || fileType == 'file'" class="file-type-icon">
                <!-- <icon
                    :name="fileType == 'office' ? 'el-icon-Document' : 'el-icon-Files'"
                    :size="28"
                    color="#909399"
                /> -->
                .{{ ext }}
            </div>

            <slot></slot>
        </div>
    </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { GetFileType } from '@/enums/fileEnums'

export default defineComponent({
    props: {
        // 图片地址
        uri: {
            type: String,
            default: ''
        },
        // 图片尺寸
        fileSize: {
            type: String,
            default: '100px'
        }
        // // 文件类型
        // ext: {
        //     type: String,
        //     default: ''
        // }
    },
    computed: {
        fileType() {
            const fileType = GetFileType(this.uri)
            return fileType
        },
        ext() {
            return this.uri.split('.').pop()
        }
    }
})
</script>

<style scoped lang="scss">
.file-item {
    box-sizing: border-box;
    position: relative;
    border-radius: 4px;
    overflow: hidden;

    background-color: var(--el-border-color-extra-light);
    border-width: 1px;
    border-color: var(--el-border-color-extra-light);

    .image,
    .video {
        display: block;
        box-sizing: border-box;
        width: 100%;
        height: 100%;
    }

    .file-type-icon {
        display: flex;
        align-items: center;
        justify-content: center;
        width: 100%;
        height: 100%;
    }
}
</style>
