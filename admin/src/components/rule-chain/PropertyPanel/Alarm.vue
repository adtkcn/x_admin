<template>
    <el-form label-width="110px">
        <el-form-item label="告警脚本">
            <div class="rc-alarm__script">
                <div class="rc-tip">
                    实现 <code>alert(values, sn)</code>，返回数字即告警级别（0=不告警/恢复，1/2/3=级别）；兼容返回布尔（true=默认级别）。留空则本节点不触发告警。
                </div>
                <ScriptEditor v-model="model.script" height="260px" placeholder="function alert(values, sn) {&#10;  return values.temp > 80 ? 2 : 0&#10;}" />
            </div>
        </el-form-item>
        <el-form-item label="持续时长(秒)">
            <el-input-number v-model="model.duration" :min="0" :max="86400" :controls="false" />
            <div class="rc-tip">异常需连续保持该时长才升级为告警，0=命中即告</div>
        </el-form-item>
        <el-form-item label="静默时长(秒)">
            <el-input-number v-model="model.silence" :min="0" :max="86400" :controls="false" />
            <div class="rc-tip">同一告警触发后的静默窗口，避免刷屏，0=不静默</div>
        </el-form-item>
    </el-form>
</template>

<script setup lang="ts">
import type { RcAlarmProps } from '../rule.type'
import ScriptEditor from './ScriptEditor.vue'
const model = defineModel<RcAlarmProps>({ required: true })
</script>

<style scoped>
.rc-alarm__script {
    width: 100%;
}
.rc-tip {
    font-size: 12px;
    color: var(--el-text-color-secondary);
    line-height: 1.6;
    margin-bottom: 6px;
}
code {
    background: var(--el-fill-color-light);
    padding: 0 4px;
    border-radius: 3px;
}
</style>
