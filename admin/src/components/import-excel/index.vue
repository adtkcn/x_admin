<template>
    <div class="import-excel">
        <!-- 触发区：默认插槽作为点击触发，作用域暴露导入中状态 -->
        <span class="import-excel__trigger" @click="triggerSelect">
            <slot :loading="loading"></slot>
        </span>
        <input
            ref="fileInputRef"
            type="file"
            class="import-excel__input"
            :multiple="multiple"
            :accept="getAccept"
            @change="handleFileChange"
        />
        <el-dialog
            v-if="showProgress && fileList.length"
            v-model="visible"
            title="导入进度"
            :close-on-click-modal="false"
            width="500px"
            @close="handleClose"
        >
            <div class="p-4">
                <div v-for="(item, index) in fileList" :key="index" class="mb-5">
                    <div class="flex items-center justify-between">
                        <span class="mr-2">{{ item.name }}</span>
                        <el-tag v-if="item.status === 'fail'" size="small" type="danger">
                            失败
                        </el-tag>
                        <el-tag v-else-if="item.status === 'success'" size="small" type="success">
                            成功
                        </el-tag>
                    </div>
                    <el-progress
                        :percentage="item.percentage"
                        :status="
                            item.status === 'fail'
                                ? 'exception'
                                : item.status === 'success'
                                  ? 'success'
                                  : ''
                        "
                    />
                </div>
            </div>
        </el-dialog>
    </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import type { PropType } from 'vue'
import axios from 'axios'
import config from '@/config'
import { getToken } from '@/utils/auth'
import feedback from '@/utils/feedback'
import { RequestCodeEnum } from '@/enums/requestEnums'

const props = defineProps({
    // 导入接口地址（相对路径会拼接 baseUrl + urlPrefix）
    url: {
        type: String,
        required: true
    },
    // 允许的文件扩展名，默认 Excel
    ext: {
        type: Array as PropType<string[]>,
        default: () => ['xlsx', 'xls']
    },
    // 是否支持多选
    multiple: {
        type: Boolean,
        default: false
    },
    // 最多可导入的文件数
    limit: {
        type: Number,
        default: 1
    },
    // 导入时随文件一起提交的额外参数
    data: {
        type: Object,
        default: () => ({})
    },
    // 是否显示导入进度弹窗
    showProgress: {
        type: Boolean,
        default: true
    }
})

const emit = defineEmits<{
    (e: 'change', fileLists: any[]): void
    (e: 'error'): void
}>()

const fileInputRef = ref<HTMLInputElement>()
const visible = ref(false)
const loading = ref(false)
const fileList = ref<any[]>([])
// 导入序号：用于取消进行中的任务
let importSeq = 0

const fullUrl = computed(() => {
    if (props.url.startsWith('http')) return props.url
    return `${config.baseUrl}${config.urlPrefix}${props.url}`
})

const getAccept = computed(() => {
    if (props.ext.length) {
        return props.ext.map((item) => `.${item.replace(/^\./, '')}`).join(',')
    }
    return ''
})

function triggerSelect() {
    if (loading.value) return
    fileInputRef.value?.click()
}

// 校验文件扩展名
function checkExt(name: string) {
    if (!props.ext.length) return true
    const i = name.lastIndexOf('.')
    const ext = i >= 0 ? name.slice(i + 1).toLowerCase() : ''
    return props.ext.some((item) => item.replace(/^\./, '').toLowerCase() === ext)
}

function handleFileChange(e: Event) {
    const input = e.target as HTMLInputElement
    const files = input.files ? Array.from(input.files) : []
    input.value = '' // 允许重复选择同一文件
    if (!files.length) return

    let selected = files
    if (!props.multiple) selected = selected.slice(0, 1)
    if (selected.length > props.limit) {
        feedback.msgError(`最多只能导入${props.limit}个文件`)
        return
    }
    if (!selected.every((file) => checkExt(file.name))) {
        feedback.msgError(`仅支持${props.ext.join('/')}格式的文件`)
        return
    }
    startImport(selected)
}

// 提交单个文件
async function importOne(file: File, item: any): Promise<any> {
    const formData = new FormData()
    formData.append('file', file)
    Object.keys(props.data).forEach((k) => formData.append(k, (props.data as any)[k]))
    const resp = await axios.post(fullUrl.value, formData, {
        headers: { token: getToken(), version: config.version },
        onUploadProgress: (e) => {
            // 留 1% 给服务端解析 Excel 的时间
            if (e.total) item.percentage = Math.min(99, Math.floor((e.loaded / e.total) * 100))
        }
    })
    return resp.data
}

async function startImport(files: File[]) {
    const seq = ++importSeq
    visible.value = props.showProgress
    loading.value = true
    fileList.value = files.map((f, i) => ({
        uid: `${seq}-${i}`,
        name: f.name,
        size: f.size,
        status: 'uploading',
        percentage: 0
    }))

    const results: any[] = []
    for (let i = 0; i < files.length; i++) {
        if (seq !== importSeq) return // 已被取消
        const file = files[i]
        const item = fileList.value[i]
        try {
            const data = await importOne(file, item)
            if (seq !== importSeq) return
            // 后端统一返回 { code, message, data }
            if (data?.code !== RequestCodeEnum.SUCCESS) {
                item.status = 'fail'
                loading.value = false
                feedback.msgError(data?.message || '导入失败')
                emit('error')
                return
            }
            item.status = 'success'
            item.percentage = 100
            results.push({ name: file.name, response: data })
        } catch (err: any) {
            item.status = 'fail'
            loading.value = false
            feedback.msgError(err?.message || '导入失败')
            emit('error')
            return
        }
    }

    loading.value = false
    visible.value = false
    emit('change', results)
}

function handleClose() {
    importSeq++ // 取消进行中的导入
    fileList.value = []
    visible.value = false
}

defineExpose({ triggerSelect })
</script>

<style lang="scss" scoped>
.import-excel {
    display: inline-block;
    vertical-align: middle;
}
.import-excel__trigger {
    display: inline-block;
    cursor: pointer;
}
.import-excel__input {
    display: none;
}
</style>
