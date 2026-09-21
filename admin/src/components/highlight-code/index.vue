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
import hljs from 'highlight.js/lib/core'
import javascript from 'highlight.js/lib/languages/javascript'
import go from 'highlight.js/lib/languages/go'
import xml from 'highlight.js/lib/languages/xml'
defineOptions({
    name: 'HighlightCode'
})

hljs.registerLanguage('javascript', javascript)
hljs.registerLanguage('go', go)
hljs.registerLanguage('xml', xml)

import hljsVuePlugin from '@highlightjs/vue-plugin'
const highlightjs = hljsVuePlugin.component

interface Props {
    code: string
    lang: string
}
withDefaults(defineProps<Props>(), {
    code: '',
    lang: 'javascript'
})
</script>
