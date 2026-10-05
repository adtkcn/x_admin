<template>
    <div class="h-full">
        <iframe :src="src" frameborder="0" width="100%" height="100%"></iframe>
    </div>
</template>

<script setup lang="ts">
import { ref, watch, onMounted, onUnmounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'

defineOptions({
    name: 'IframeComponent'
})

const route = useRoute()
const router = useRouter()

// iframe 的实际 src：独立于 query.url，避免子应用内部跳转回写 query 时误触发 iframe 重新加载
const src = ref('')
// 最近一次由子应用上报的 url，用于区分「父应用主动改 url」与「子应用内部跳转回写」
let innerUrl = ''

// 监听 query.url：主应用跳转 / 刷新浏览器 / 复制链接还原 都走这里
watch(
    () => route.query.url,
    (val) => {
        const url = (val as string) || ''
        // 子应用内部跳转回写的 url，iframe 已自行导航，无需重新加载
        if (url && url === innerUrl) return
        src.value = url
    },
    { immediate: true }
)

// 接收子应用内部跳转上报：子应用 postMessage({ type: 'iframe-url-change', url })
const handleMessage = (e: MessageEvent) => {
    const data = e.data
    if (!data || data.type !== 'iframe-url-change' || typeof data.url !== 'string') return
    innerUrl = data.url
    // 用 replace 不新增历史记录，刷新 / 复制链接都能还原到当前子应用页面
    router.replace({ query: { ...route.query, url: data.url } })
}

onMounted(() => window.addEventListener('message', handleMessage))
onUnmounted(() => window.removeEventListener('message', handleMessage))
</script>
