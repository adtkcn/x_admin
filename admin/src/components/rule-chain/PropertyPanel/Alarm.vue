<template>
    <el-form label-width="110px">
        <el-form-item label="告警脚本">
            <div class="rc-alarm__script">
                <div class="rc-tip">
                    实现 <code>alert(values, sn)</code>，只返回数字：0=不告警/恢复，非 0 即告警级别（危害升序：-1=故障、1=轻微、2=严重，3+ 预留；故障危害最轻但优先于其他告警）；返回其它类型按不告警处理。留空则本节点不触发告警。
                </div>
                <ScriptEditor
                    v-model="model.script"
                    height="260px"
                    :template="RC_SCRIPT_TEMPLATES.rule_alarm"
                    placeholder="function alert(values, sn) {&#10;  return values.temp > 80 ? 2 : 0&#10;}"
                />
            </div>
        </el-form-item>
        <el-form-item label="持续时长(秒)">
            <el-input-number v-model="model.duration" :min="0" :max="86400" :controls="false" />
            <div class="rc-tip">异常需连续保持该时长才升级为告警，0=命中即告</div>
        </el-form-item>
        <el-form-item label="恢复确认(秒)">
            <el-input-number v-model="model.recover" :min="0" :max="86400" :controls="false" />
            <div class="rc-tip">报警后需恢复正常连续满该时长才收尾，期间再命中则重新计时，0=恢复正常即收</div>
        </el-form-item>
    </el-form>
</template>

<script setup lang="ts">
import type { RcAlarmProps } from '../rule.type'
import { RC_SCRIPT_TEMPLATES } from '../rule.type'
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
