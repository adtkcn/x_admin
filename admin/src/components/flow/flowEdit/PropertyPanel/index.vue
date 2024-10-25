<template>
    <el-drawer
        v-model="drawerVisible"
        size="600px"
        :title="'节点：' + node?.text?.value"
        @close="close"
    >
        <div v-if="node.type == 'bpmn:startEvent'">
            <FieldAuth :node="node" :properties="properties" :fieldList="fieldList"></FieldAuth>
        </div>
        <div v-if="node.type == 'bpmn:userTask'">
            <UserTask :node="node" :properties="properties" :fieldList="fieldList"></UserTask>
            <FieldAuth :node="node" :properties="properties" :fieldList="fieldList"></FieldAuth>
        </div>

        <div v-if="node.type == 'bpmn:serviceTask'">
            <div>（都还不支持）系统任务</div>
            <div>（都还不支持）抄送</div>
            <div>（都还不支持）发送邮件</div>
            <div>（都还不支持）发送短信</div>
            <div>（都还不支持）发送站内消息</div>
            <div>（都还不支持）数据入库</div>
        </div>
        <div v-if="node.type == 'bpmn:exclusiveGateway'">
            <Gateway :node="node" :properties="properties" :fieldList="fieldList"></Gateway>
        </div>

        <div v-if="node.type == 'bpmn:endEvent'">结束</div>
    </el-drawer>
</template>
<script setup lang="ts">
import { ref, toRaw, reactive } from 'vue'
import UserTask from './UserTask.vue'
import FieldAuth from './FieldAuth.vue'
import Gateway from './Gateway.vue'
import type { NodeType, PropertiesType, FormFieldListType, FieldListType } from './property.type'
defineOptions({
    name: 'PropertyPanel'
})
const emit = defineEmits(['setProperties'])

const drawerVisible = ref(false)

const node = ref<NodeType>({})
const properties = reactive<PropertiesType>({
    userType: null,
    userId: null,
    deptId: null,
    postId: null,
    fieldAuth: {},
    gateway: []
})

const fieldList = ref<FieldListType[]>([])

const open = (newNode: NodeType, newFieldList: FormFieldListType[]) => {
    if (newNode.type == 'bpmn:endEvent') {
        return
    }
    node.value = newNode

    properties.userType = newNode?.properties?.userType || null
    properties.userId = newNode?.properties?.userId || null
    properties.deptId = newNode?.properties?.deptId || null
    properties.postId = newNode?.properties?.postId || null
    properties.gateway = newNode?.properties?.gateway || []

    properties.fieldAuth = newNode?.properties?.fieldAuth
        ? { ...newNode?.properties?.fieldAuth }
        : {}

    fieldList.value = newFieldList.map((item) => {
        return {
            id: item?.field?.id,
            label: item?.field?.options?.label,
            auth: newNode?.properties?.fieldAuth?.[item?.field?.id] || 1
        }
    })
    drawerVisible.value = true
}

const close = () => {
    const fieldAuth = {}
    fieldList.value.forEach((item) => {
        fieldAuth[item.id] = item.auth
    })
    properties.fieldAuth = fieldAuth
    emit('setProperties', toRaw(node.value), { ...toRaw(properties) })
    drawerVisible.value = false
}

defineExpose({
    open
})
</script>

<style scoped></style>
