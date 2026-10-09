<template>
    <div class="rc-script">
        <div class="rc-script__hint" v-html="hint"></div>
        <ScriptEditor
            v-model="model.script"
            :placeholder="placeholder"
            :template="template"
            height="300px"
        />
    </div>
</template>

<script setup lang="ts">
import type { RcScriptProps } from '../rule.type'
import ScriptEditor from './ScriptEditor.vue'
const model = defineModel<RcScriptProps>({ required: true })
defineProps<{ hint?: string; placeholder?: string; template?: string }>()
</script>

<style scoped>
.rc-script {
    width: 100%;
}
.rc-script__hint {
    font-size: 12px;
    color: var(--el-text-color-secondary);
    line-height: 1.6;
    margin-bottom: 8px;
}
</style>
