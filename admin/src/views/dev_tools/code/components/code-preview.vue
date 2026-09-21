<template>
    <div class="code-preview">
        <el-dialog
            v-model="show"
            width="98%"
            :title="'代码预览【' + showItem.label + '】'"
            top="0px"
            draggable
        >
            <el-container style="height: calc(100vh - 200px)">
                <el-aside
                    width="400px"
                    style="padding: 10px 0; margin-right: 20px; border: 1px solid #dcdfe6"
                >
                    <el-tree
                        :data="treeData"
                        :props="{ label: 'label' }"
                        highlight-current
                        @node-click="handleNodeClick"
                    />
                </el-aside>
                <el-main style="padding: 0; overflow: hidden; height: 100%">
                    <highlight-code
                        :code="showItem.value"
                        :lang="lang"
                        class="code-body"
                    ></highlight-code>
                </el-main>
            </el-container>
            <template v-slot:footer>
                <div>
                    <el-button
                        icon="el-icon-CopyDocument"
                        type="primary"
                        @click="handleCopy(showItem.value)"
                        >复制</el-button
                    >
                    <el-button @click="show = false">关闭</el-button>
                </div>
            </template>
        </el-dialog>
    </div>
</template>

<script lang="ts" setup>
import { onMounted, ref, computed } from 'vue'
import feedback from '@/utils/feedback'
import useClipboard from 'vue-clipboard3'

const props = defineProps<{
    modelValue: boolean
    code: Record<string, string>
}>()

const emit = defineEmits<{
    (event: 'update:modelValue', value: boolean): void
}>()
const { toClipboard } = useClipboard()

const handleCopy = async (text: string) => {
    try {
        await toClipboard(text)
        feedback.msgSuccess('复制成功')
    } catch (e) {
        feedback.msgError('复制失败')
    }
}
const showItem = ref({
    label: '',
    value: ''
})
const lang = computed(() => {
    const ext = showItem.value.label?.split('.')?.[1]
    switch (ext) {
        case 'js':
            return 'javascript'
        case 'ts':
            return 'typescript'
        case 'vue':
            return 'xml'
        case 'go':
            return 'go'
        default:
            return 'javascript'
    }
    // return ext == 'js' ? 'javascript' : ext
})
const treeData = computed(() => {
    return Object.keys(props.code).map((key) => {
        return {
            label: key,
            value: props.code[key]
        }
    })
})

onMounted(() => {
    if (treeData.value.length === 0) {
        return
    }
    showItem.value = {
        label: treeData.value[0]?.label,
        value: treeData.value[0]?.value
    }
})

function handleNodeClick(params: any) {
    console.log('params', params)
    showItem.value = {
        label: params.label,
        value: params.value
    }
}
const show = computed<boolean>({
    get() {
        return props.modelValue
    },
    set(value) {
        emit('update:modelValue', value)
    }
})
</script>
<style lang="scss" scoped>
.code-preview {
    :deep(.el-dialog__body) {
        padding-top: 10px;
    }
    .code-body {
        height: 100%;
        width: 100%;
    }
}
</style>
