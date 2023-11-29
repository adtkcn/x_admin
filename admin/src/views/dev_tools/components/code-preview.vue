<template>
    <div class="code-preview">
        <el-dialog v-model="show" width="95%" title="代码预览">
            <el-container style="height: 75vh">
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
                <el-main style="padding: 0; border: 1px solid #dcdfe6">
                    <div class="flex flex-col h-[100%]">
                        <div class="flex">
                            <div class="flex-1 p-4">{{ showItem.label }}</div>
                            <div>
                                <el-button @click="handleCopy(showItem.value)" type="primary" link>
                                    <template #icon>
                                        <icon name="el-icon-CopyDocument" />
                                    </template>
                                    复制
                                </el-button>
                            </div>
                        </div>

                        <div class="flex-1 overflow-auto">
                            <highlightjs autodetect :code="showItem.value" language="javascript" />
                        </div>
                    </div>
                </el-main>
            </el-container>
        </el-dialog>
    </div>
</template>

<script lang="ts" setup>
import feedback from '@/utils/feedback'
import useClipboard from 'vue-clipboard3'

import 'highlight.js/styles/github.css'
import hljs from 'highlight.js/lib/common'
import javascript from 'highlight.js/lib/languages/javascript'

// Then register the languages you need
hljs.registerLanguage('javascript', javascript)

import hljsVuePlugin from '@highlightjs/vue-plugin'
const highlightjs = hljsVuePlugin.component

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
