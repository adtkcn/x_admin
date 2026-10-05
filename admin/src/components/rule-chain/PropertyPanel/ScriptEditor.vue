<template>
    <div class="se" :style="{ height }">
        <div class="se__bar">
            <span class="se__lang">javascript</span>
            <span class="se__meta">{{ lineCount }} 行 · {{ (modelValue || '').length }} 字符</span>
        </div>
        <div class="se__body">
            <!-- 行号槽：与代码区共用滚动偏移 -->
            <div ref="gutterRef" class="se__gutter">
                <div v-for="n in lineCount" :key="n" class="se__lineno">{{ n }}</div>
            </div>
            <div class="se__code">
                <pre class="se__hl" aria-hidden="true"><code v-html="highlighted"></code></pre>
                <textarea
                    ref="taRef"
                    class="se__ta"
                    spellcheck="false"
                    wrap="off"
                    :value="modelValue"
                    :placeholder="placeholder"
                    @input="onInput"
                    @keydown.tab.prevent="onTab"
                    @keydown.enter.prevent="onEnter"
                    @scroll="syncScroll"
                ></textarea>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { computed, ref, nextTick, onMounted } from 'vue'
import hljs from 'highlight.js/lib/core'
import javascript from 'highlight.js/lib/languages/javascript'
import 'highlight.js/styles/monokai.min.css'

hljs.registerLanguage('javascript', javascript)

const props = defineProps<{
    modelValue: string
    placeholder?: string
    height?: string
}>()
const emit = defineEmits(['update:modelValue'])

const taRef = ref<HTMLTextAreaElement>()
const gutterRef = ref<HTMLDivElement>()

// 末尾补一个换行，保证最后一行高亮与行号对齐渲染
const highlighted = computed(() => {
    const code = (props.modelValue ?? '') + '\n'
    return hljs.highlight(code, { language: 'javascript', ignoreIllegals: true }).value
})
const lineCount = computed(() => (props.modelValue ?? '').split('\n').length || 1)

function onInput(e: Event) {
    emit('update:modelValue', (e.target as HTMLTextAreaElement).value)
    nextTick(syncScroll)
}

// Tab 缩进 / Shift+Tab 反缩进：按整行处理，选区跨行时批量缩进
function onTab(e: KeyboardEvent) {
    const ta = taRef.value!
    const { selectionStart: s, selectionEnd: en, value } = ta
    const shift = e.shiftKey
    // 扩展选区到整行，保证多行缩进符合编辑器习惯
    const lineStart = value.lastIndexOf('\n', s - 1) + 1
    let lineEnd = value.indexOf('\n', en)
    if (lineEnd === -1) lineEnd = value.length
    const block = value.slice(lineStart, lineEnd)
    const lines = block.split('\n')
    let delta = 0
    const out = lines.map((ln) => {
        if (shift) {
            const m = ln.match(/^ {1,2}/)
            if (m) {
                delta -= m[0].length
                return ln.slice(m[0].length)
            }
            return ln
        }
        delta += 2
        return '  ' + ln
    })
    const newBlock = out.join('\n')
    const nv = value.slice(0, lineStart) + newBlock + value.slice(lineEnd)
    emit('update:modelValue', nv)
    nextTick(() => {
        ta.selectionStart = Math.max(lineStart, s + (shift ? Math.min(0, delta) : 2))
        ta.selectionEnd = Math.max(lineStart, en + delta)
    })
}

// 回车：自动继承上一行缩进；上一行以 { 结尾则再进一级
function onEnter() {
    const ta = taRef.value!
    const { selectionStart: s, selectionEnd: en, value } = ta
    const lineStart = value.lastIndexOf('\n', s - 1) + 1
    const curLine = value.slice(lineStart, s)
    const indent = (curLine.match(/^\s*/) || [''])[0]
    const extra = /{\s*$/.test(curLine) ? '  ' : ''
    const insert = '\n' + indent + extra
    const nv = value.slice(0, s) + insert + value.slice(en)
    emit('update:modelValue', nv)
    nextTick(() => {
        const pos = s + insert.length
        ta.selectionStart = ta.selectionEnd = pos
    })
}

// 高亮层与行号槽跟随 textarea 滚动，保持像素对齐
function syncScroll() {
    const ta = taRef.value
    if (!ta) return
    const codeEl = ta.previousElementSibling as HTMLElement | null
    if (codeEl) {
        codeEl.scrollTop = ta.scrollTop
        codeEl.scrollLeft = ta.scrollLeft
    }
    if (gutterRef.value) gutterRef.value.scrollTop = ta.scrollTop
}

onMounted(syncScroll)
</script>

<style scoped>
.se {
    display: flex;
    flex-direction: column;
    width: 100%;
    border: 1px solid var(--el-border-color);
    border-radius: 4px;
    overflow: hidden;
    background: #272822;
}
.se__bar {
    display: flex;
    align-items: center;
    gap: 12px;
    height: 30px;
    padding: 0 10px;
    font-size: 12px;
    color: #cfcfc2;
    background: #1e1f1a;
    flex-shrink: 0;
}
.se__meta {
    margin-left: auto;
}
.se__body {
    position: relative;
    flex: 1;
    display: flex;
    min-height: 0;
}
.se__gutter {
    width: 42px;
    flex-shrink: 0;
    overflow: hidden;
    text-align: right;
    padding: 8px 6px 8px 0;
    background: #20211b;
    color: #6b6d60;
    user-select: none;
    font-family: 'JetBrains Mono', Consolas, Menlo, monospace;
    font-size: 13px;
    line-height: 1.6;
}
.se__lineno {
    height: 20.8px;
}
.se__code {
    position: relative;
    flex: 1;
    min-width: 0;
}
.se__hl,
.se__ta {
    position: absolute;
    inset: 0;
    margin: 0;
    padding: 8px 12px;
    font-family: 'JetBrains Mono', Consolas, Menlo, monospace;
    font-size: 13px;
    line-height: 1.6;
    tab-size: 2;
    white-space: pre;
    word-wrap: normal;
    border: 0;
    box-sizing: border-box;
}
.se__hl {
    overflow: hidden;
    pointer-events: none;
    background: transparent;
    color: #f8f8f2;
}
.se__hl :deep(code) {
    display: block;
    background: transparent;
    padding: 0;
    font: inherit;
    white-space: pre;
}
.se__ta {
    overflow: auto;
    resize: none;
    background: transparent;
    color: transparent;
    caret-color: #fff;
    outline: none;
}
.se__ta::placeholder {
    color: #6b6d60;
}
</style>
