<template>
    <el-drawer
        v-model="drawerVisible"
        size="600px"
        :title="'节点：' + node?.text?.value"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
    >
        <!-- 按节点类型分发子面板，子面板通过 v-model 直接操作 nodeProps 私有数据 -->
        <div v-if="nodeKey === 'start_event'">
            <FieldAuth :fieldList="fieldList"></FieldAuth>
        </div>
        <div v-else-if="nodeKey === 'user_task'">
            <UserTask v-model="nodeProps" :fieldList="fieldList"></UserTask>
            <FieldAuth :fieldList="fieldList"></FieldAuth>
        </div>
        <div v-else-if="nodeKey === 'notify_task'">
            <NotifyTask v-model="nodeProps"></NotifyTask>
        </div>
        <div v-else-if="nodeKey === 'exclusive_gateway'">
            <Gateway v-model="nodeProps" :fieldList="fieldList"></Gateway>
        </div>

        <div v-if="node.type == 'bpmn:endEvent'">结束</div>

        <template #footer>
            <div style="text-align: right">
                <el-button @click="cancel">取消</el-button>
                <el-button type="primary" @click="confirm">确定</el-button>
            </div>
        </template>
    </el-drawer>
</template>
<script setup lang="ts">
import { ref, toRaw } from 'vue'
import UserTask from './UserTask.vue'
import FieldAuth from './FieldAuth.vue'
import Gateway from './Gateway.vue'
import NotifyTask from './NotifyTask.vue'
import type {
    NodeType,
    FormFieldListType,
    FieldListType,
    NodeKey,
    FieldAuthMap,
} from './property.type'
import { nodeKeyOf, defaultProps } from './property.type'

defineOptions({
    name: 'PropertyPanel'
})
const emit = defineEmits(['setProperties'])

const drawerVisible = ref(false)

const node = ref<NodeType>({})
// 当前节点类型对应的属性命名空间 key
const nodeKey = ref<NodeKey>()
// 当前节点类型的私有属性（仅装载自身那一份，其余 key 不关心）
const nodeProps = ref<any>({})

const fieldList = ref<FieldListType[]>([])

const open = (newNode: NodeType, newFieldList: FormFieldListType[]) => {
    const key = nodeKeyOf(newNode.type)
    // 结束节点无配置项，直接跳过
    if (!key || key === 'end_event') {
        return
    }
    node.value = newNode
    nodeKey.value = key
    // 深拷贝：避免子面板对 gateway / field_auth 的修改直接污染 LogicFlow 原始数据，
    // 保证「取消」时本次编辑不生效
    nodeProps.value = {
        ...defaultProps(key),
        ...structuredClone(toRaw(newNode.properties?.[key]) ?? {}),
    }

    fieldList.value = newFieldList.map((item) => {
        return {
            id: item?.id,
            name: item?.name,
            // 权限字段从 start_event / user_task 的 field_auth 读取
            auth:
                (key
                    ? (newNode?.properties?.[key as NodeKey] as any)?.field_auth?.[item?.id]
                    : undefined) || 1,
        }
    })
    drawerVisible.value = true
}

// 确定：回写 field_auth 并保存当前节点属性，仅写自身命名空间
const confirm = () => {
    if ('field_auth' in nodeProps.value) {
        const field_auth: FieldAuthMap = {}
        fieldList.value.forEach((item) => {
            if (item.id) field_auth[item.id] = item.auth
        })
        nodeProps.value.field_auth = field_auth
    }
    emit('setProperties', toRaw(node.value), {
        [nodeKey.value as string]: toRaw(nodeProps.value),
    })
    drawerVisible.value = false
}

// 取消：不保存，直接关闭
const cancel = () => {
    drawerVisible.value = false
}

defineExpose({
    open
})
</script>

<style scoped></style>
