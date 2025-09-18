<template>
    <highlightjs
        style="width: 100%; height: 100%; overflow: auto"
        autodetect
        :code="code"
        :language="lang"
    />
</template>
<script lang="ts" setup>
import 'highlight.js/styles/monokai.min.css'
import hljs from 'highlight.js/lib/common'
import javascript from 'highlight.js/lib/languages/javascript'
import go from 'highlight.js/lib/languages/go'

hljs.registerLanguage('javascript', javascript)
hljs.registerLanguage('go', go)

import hljsVuePlugin from '@highlightjs/vue-plugin'
const highlightjs = hljsVuePlugin.component

interface Props {
    code: string
    lang: 'javascript' | 'go'
}
withDefaults(defineProps<Props>(), {
    code: '',
    lang: 'javascript'
})
</script>
