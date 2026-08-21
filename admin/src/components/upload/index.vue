<template>
    <div class="upload">
        <!-- 触发区：默认/自定义插槽作为点击触发 -->
        <span class="upload__trigger" @click="triggerSelect">
            <slot></slot>
        </span>
        <input
            ref="fileInputRef"
            type="file"
            class="upload__input"
            :multiple="multiple"
            :accept="getAccept"
            @change="handleFileChange"
        />
        <el-dialog
            v-if="showProgress && fileList.length"
            v-model="visible"
            title="上传进度"
            :close-on-click-modal="false"
            width="500px"
            @close="handleClose"
        >
            <div class="file-list p-4">
                <template v-for="(item, index) in fileList" :key="index">
                    <div class="mb-5">
                        <div class="flex items-center justify-between">
                            <span class="mr-2">{{ item.name }}</span>
                            <el-tag v-if="item.status === 'instant'" size="small" type="success"
                                >秒传</el-tag
                            >
                            <el-tag v-else-if="item.status === 'fail'" size="small" type="danger"
                                >失败</el-tag
                            >
                        </div>
                        <el-progress
                            :percentage="item.percentage"
                            :status="
                                item.status === 'fail'
                                    ? 'exception'
                                    : item.status === 'success' || item.status === 'instant'
                                      ? 'success'
                                      : ''
                            "
                        />
                    </div>
                </template>
            </div>
        </el-dialog>
    </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import type { PropType } from 'vue'
import SparkMD5 from 'spark-md5'
import axios from 'axios'
import config from '@/config'
import { getToken } from '@/utils/auth'
import feedback from '@/utils/feedback'
import { RequestCodeEnum } from '@/enums/requestEnums'

const props = defineProps({
    // 上传地址（相对路径会拼接 baseUrl + urlPrefix）
    url: {
        type: String,
        default: ''
    },
    // 允许的文件扩展名，如 ['png', 'jpg']
    ext: {
        type: Array as PropType<string[]>,
        default: () => []
    },
    // 是否支持多选
    multiple: {
        type: Boolean,
        default: true
    },
    // 多选时最多选择几条
    limit: {
        type: Number,
        default: 10
    },
    // 上传时的额外参数（随文件一起提交）
    data: {
        type: Object,
        default: () => ({})
    },
    // 是否显示上传进度弹窗
    showProgress: {
        type: Boolean,
        default: false
    }
})

const emit = defineEmits<{
    (e: 'change', fileLists: any[]): void
    (e: 'error'): void
}>()

const fileInputRef = ref<HTMLInputElement>()
const visible = ref(false)
const fileList = ref<any[]>([])
// 上传序号：用于取消进行中的任务
let uploadSeq = 0

const fullUrl = computed(() => {
    if (props.url && props.url.startsWith('http')) return props.url
    return `${config.baseUrl}${config.urlPrefix}${props.url || '/common/upload/file'}`
})

const checkInstantUrl = computed(() => {
    return `${config.baseUrl}${config.urlPrefix}/common/upload/checkInstant`
})

const getAccept = computed(() => {
    if (props.ext.length) {
        return props.ext.map((item) => `.${item.replace(/^\./, '')}`).join(',')
    }
    return ''
})

function triggerSelect() {
    fileInputRef.value?.click()
}

function getExt(name: string) {
    const i = name.lastIndexOf('.')
    return i >= 0 ? name.slice(i + 1) : ''
}

function handleFileChange(e: Event) {
    const input = e.target as HTMLInputElement
    const files = input.files ? Array.from(input.files) : []
    input.value = '' // 允许重复选择同一文件
    if (!files.length) return
    let selected = files
    if (!props.multiple && selected.length > 1) selected = selected.slice(0, 1)
    if (selected.length > props.limit) {
        feedback.msgError(`超出上传上限${props.limit}，请重新上传`)
        return
    }
    startUpload(selected)
}

// 计算文件 MD5
function calcMD5(file: File): Promise<string> {
    return new Promise((resolve, reject) => {
        const spark = new SparkMD5.ArrayBuffer()
        const reader = new FileReader()
        const chunk = 10 * 1024 * 1024
        const total = Math.ceil(file.size / chunk)
        let idx = 0
        reader.onload = () => {
            spark.append(reader.result as ArrayBuffer)
            idx++
            if (idx < total) loadNext()
            else resolve(spark.end())
        }
        reader.onerror = () => reject(new Error('文件读取错误'))
        const loadNext = () => {
            const s = idx * chunk
            reader.readAsArrayBuffer(file.slice(s, Math.min(s + chunk, file.size)))
        }
        loadNext()
    })
}

// 秒传检查
async function checkInstant(
    md5: string,
    file_name: string
): Promise<{
    instant: boolean
    file_hash_id?: string
    // path?: string
    size?: number
    url?: string
}> {
    try {
        const res = await axios.post(
            checkInstantUrl.value,
            {
                file_md5: md5,
                file_name
            },
            {
                headers: { token: getToken(), version: config.version }
            }
        )
        // 后端统一返回 { code, msg, show, data: CommonUploadFileResp }
        return res.data?.data || { instant: false }
    } catch (error) {
        console.error('秒传检查失败:', error)
        return { instant: false }
    }
}

// 实际上传（落盘）
async function uploadOne(file: File, item: any): Promise<any> {
    const formData = new FormData()
    formData.append('file', file)
    Object.keys(props.data).forEach((k) => formData.append(k, (props.data as any)[k]))
    const resp = await axios.post(fullUrl.value, formData, {
        headers: { token: getToken(), version: config.version },
        onUploadProgress: (e) => {
            if (e.total) item.percentage = Math.floor((e.loaded / e.total) * 100)
        }
    })
    return resp.data
}

async function startUpload(files: File[]) {
    const seq = ++uploadSeq
    visible.value = props.showProgress
    fileList.value = files.map((f, i) => ({
        uid: `${seq}-${i}`,
        name: f.name,
        size: f.size,
        status: 'computing',
        percentage: 0
    }))

    const results: any[] = []
    for (let i = 0; i < files.length; i++) {
        if (seq !== uploadSeq) return // 已被取消
        const file = files[i]
        const item = fileList.value[i]
        try {
            // 1. 计算 MD5
            const md5 = await calcMD5(file)
            if (seq !== uploadSeq) return

            // 2. 秒传检查：命中则跳过实际上传
            const instant = await checkInstant(md5, file.name)
            if (instant.instant && instant.file_hash_id) {
                item.status = 'instant'
                item.percentage = 100
                results.push({
                    name: file.name,
                    response: {
                        data: {
                            file_hash_id: instant.file_hash_id,
                            name: file.name,
                            ext: getExt(file.name),
                            size: instant.size,
                            url: instant.url
                            // path: instant.path
                        }
                    }
                })
                continue
            }

            // 3. 实际上传
            item.status = 'uploading'
            const data = await uploadOne(file, item)
            if (seq !== uploadSeq) return
            if (data?.code === RequestCodeEnum.FAILED) {
                item.status = 'fail'
                feedback.msgError(data.message || '上传失败')
                emit('error')
                return
            }
            item.status = 'success'
            item.percentage = 100
            // 保持与 el-upload 一致的响应结构：response.data 为后端返回的业务数据
            results.push({ name: file.name, response: data })
        } catch (err: any) {
            item.status = 'fail'
            feedback.msgError(err?.message || '上传失败')
            emit('error')
            return
        }
    }

    visible.value = false
    emit('change', results)
}

function handleClose() {
    uploadSeq++ // 取消进行中的上传
    fileList.value = []
    visible.value = false
}

defineExpose({ triggerSelect })
</script>

<style lang="scss" scoped>
.upload {
    display: inline-block;
    vertical-align: middle;
}
.upload__trigger {
    display: inline-block;
    cursor: pointer;
}
.upload__input {
    display: none;
}
</style>
