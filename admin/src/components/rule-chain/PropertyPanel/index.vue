<template>
    <el-drawer
        v-model="drawerVisible"
        size="600px"
        :title="title"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
    >
        <!-- 节点显示名：双击只打开面板，改名统一在此进行 -->
        <el-form label-width="70px" class="rc-name">
            <el-form-item label="显示名">
                <el-input v-model="nodeName" maxlength="30" placeholder="节点在图上显示的文字" />
            </el-form-item>
        </el-form>

        <!-- 节点属性：按命名空间分发子面板 -->
        <div v-if="nodeKey === 'rule_trigger'" class="rc-tip">输入节点：属性上报由此进入规则链，不可删除、无需配置。</div>
        <template v-else>
            <Script
                v-if="nodeKey === 'rule_decode'"
                v-model="nodeProps"
                :hint="RC_SCRIPT_HINTS.rule_decode"
                :template="RC_SCRIPT_TEMPLATES.rule_decode"
                placeholder="function decode(values, sn) {&#10;  // values: 初始上报值;  return { temp: 25, hum: 60 }&#10;}"
            ></Script>
            <Script
                v-else-if="nodeKey === 'rule_filter'"
                v-model="nodeProps"
                :hint="RC_SCRIPT_HINTS.rule_filter"
                :template="RC_SCRIPT_TEMPLATES.rule_filter"
                placeholder="function filter(values, sn) {&#10;  return values.temp > 0&#10;}"
            ></Script>
            <div v-else-if="nodeKey === 'rule_parallel'" class="rc-tip">并行节点：报文同时流向所有后续分支，是唯一允许连接多个后续节点的节点，无需配置。</div>
            <Script
                v-else-if="nodeKey === 'rule_bool'"
                v-model="nodeProps"
                :hint="RC_SCRIPT_HINTS.rule_bool"
                :template="RC_SCRIPT_TEMPLATES.rule_bool"
                placeholder="function filter(values, sn) {&#10;  return values.temp > 30&#10;}"
            ></Script>
            <Alarm v-else-if="nodeKey === 'rule_alarm'" v-model="nodeProps"></Alarm>
            <div v-else-if="nodeKey === 'rule_state'" class="rc-tip">状态快照节点：把流经的最新属性写入 Redis 快照，供管理端与大屏即时读取，无需配置。</div>
            <div v-else-if="nodeKey === 'rule_history'" class="rc-tip">历史存储节点：把流经的属性追加写入时序历史表，无需配置。</div>
            <div v-else-if="nodeKey === 'rule_push'" class="rc-tip">大屏推送节点：把流经的属性广播到设备所属区域的地图大屏房间，无需配置。</div>
            <Notify v-else-if="nodeKey === 'rule_notify'" v-model="nodeProps"></Notify>
        </template>

        <template #footer>
            <div style="text-align: right">
                <el-button @click="cancel">取消</el-button>
                <el-button type="primary" @click="confirm">确定</el-button>
            </div>
        </template>
    </el-drawer>
</template>

<script setup lang="ts">
import { ref, toRaw, computed } from 'vue'
import Script from './Script.vue'
import Alarm from './Alarm.vue'
import Notify from './Notify.vue'
import type { RcNodeKey, RcNodeTypeData } from '../rule.type'
import { rcNodeKeyOf, rcDefaultProps, rcMetaOf, RC_SCRIPT_TEMPLATES, RC_SCRIPT_HINTS } from '../rule.type'

defineOptions({ name: 'RcPropertyPanel' })
const emit = defineEmits(['setProperties', 'renameNode'])

const drawerVisible = ref(false)

// 节点编辑态
const node = ref<RcNodeTypeData>({})
const nodeKey = ref<RcNodeKey>()
const nodeProps = ref<any>({})
const nodeName = ref('')

const title = computed(
    () => '节点：' + (node.value?.text?.value ?? rcMetaOf(node.value.type)?.label ?? '')
)

// 打开节点属性面板
const open = (newNode: RcNodeTypeData) => {
    const key = rcNodeKeyOf(newNode.type)
    if (!key) return
    node.value = newNode
    nodeKey.value = key
    nodeName.value = newNode.text?.value ?? rcMetaOf(newNode.type)?.label ?? ''
    nodeProps.value = {
        ...rcDefaultProps(key),
        ...structuredClone(toRaw((newNode.properties as any)?.[key]) ?? {})
    }
    drawerVisible.value = true
}

const confirm = () => {
    emit('renameNode', toRaw(node.value).id, nodeName.value)
    emit('setProperties', toRaw(node.value), {
        [nodeKey.value as string]: toRaw(nodeProps.value)
    })
    drawerVisible.value = false
}

const cancel = () => {
    drawerVisible.value = false
}

defineExpose({ open })
</script>

<style scoped>
.rc-tip {
    font-size: 13px;
    color: var(--el-text-color-secondary);
    line-height: 1.6;
}
</style>
