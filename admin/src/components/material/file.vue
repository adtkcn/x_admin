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
                .{{ extName }}
                <!-- {{ ext }} -->
                <!-- {{ fileType }} -->
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
        // 访问 URL（推荐：ossDomain + /api/uploads/<id>，由路由返回文件流）
        uri: {
            type: String,
            default: ''
        },
        // 图片尺寸
        fileSize: {
            type: String,
            default: '100px'
        },
        // 文件扩展名（不含点）。id 形式的 uri 无法从后缀推断类型，需显式传入。
        ext: {
            type: String,
            default: ''
        }
    },
    computed: {
        fileType() {
            // 优先用 ext prop 推断；空时回退到 uri 后缀（兼容旧 url 形式）
            const key = this.ext || (this.uri ? this.uri.split('.').pop() : '')
            const type = GetFileType(key || '')
            // 兜底：id 形式 uri（无后缀，GetFileType 走 default 返回 'file'）
            // 按 image 渲染，因为绝大多数上传场景是图片，且后端 ServeContent 会下发
            // 正确 Content-Type，浏览器能正确解码图片；非图片（视频/音频）场景仍按 ext 优先
            // 走对应分支，不会被这里覆盖。
            if (
                type === 'file' &&
                this.uri &&
                !this.uri.includes('.') &&
                !this.uri.includes(':')
            ) {
                return 'image'
            }
            return type
        },
        extName() {
            return this.ext || (this.uri ? this.uri.split('.').pop() : '')
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
