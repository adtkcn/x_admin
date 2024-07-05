<template>
    <el-drawer
        v-model="drawerVisible"
        size="600px"
        :title="'节点：' + node?.text?.value"
        @close="close"
    >
        <!-- fieldList:{{ fieldList }}
            <div>properties:{{ properties }}</div> -->

        <!-- 开始节点 -->
        <!-- {{ node }} -->

        <div v-if="node.type == 'bpmn:startEvent'">
            开始节点
            <FieldAuth :node="node" :properties="properties" :fieldList="fieldList"></FieldAuth>
        </div>
        <div v-if="node.type == 'bpmn:userTask'">
            <UserTask :node="node" :properties="properties" :fieldList="fieldList"></UserTask>
            <FieldAuth :node="node" :properties="properties" :fieldList="fieldList"></FieldAuth>
        </div>

        <div v-if="node.type == 'bpmn:serviceTask'">
            <div>系统任务</div>
            <div>抄送</div>
            <div>发送邮件</div>
            <div>发送短信</div>
            <div>发送站内消息</div>
            <div>数据入库</div>
        </div>
        <div v-if="node.type == 'bpmn:exclusiveGateway'">
            <Gateway :node="node" :properties="properties" :fieldList="fieldList"></Gateway>
        </div>

        <div v-if="node.type == 'bpmn:endEvent'">结束</div>
    </el-drawer>
</template>
<script setup lang="ts">
import { ref } from 'vue'
import UserTask from './UserTask.vue'
import FieldAuth from './FieldAuth.vue'
import Gateway from './Gateway.vue'

const emit = defineEmits(['setProperties'])

const drawerVisible = ref(false)
const node = ref<{
    type?: string
    text?: {
        value?: string
    }
}>({})
const properties = reactive({
    userType: '',
    userId: '',
    deptId: '',
    postId: '',
    fieldAuth: {},
    gateway: []
})
const fieldList = ref([])

const open = (newNode, newFieldList) => {
    node.value = newNode

    properties.userType = newNode?.properties?.userType || ''
    properties.userId = newNode?.properties?.userId || ''
    properties.deptId = newNode?.properties?.deptId || ''
    properties.postId = newNode?.properties?.postId || ''
    properties.gateway = newNode?.properties?.gateway || []

    properties.fieldAuth = newNode?.properties?.fieldAuth
        ? { ...newNode?.properties?.fieldAuth }
        : {}

    fieldList.value = newFieldList.map((item) => ({
        id: item?.field?.id,
        label: item?.field?.options?.label,
        auth: newNode?.properties?.fieldAuth?.[item?.field?.id] || 1
    }))
    drawerVisible.value = true
}

const close = () => {
    const fieldAuth = {}
    fieldList.value.forEach((item) => {
        fieldAuth[item.id] = item.auth
    })

    emit('setProperties', node.value, { ...properties })
    drawerVisible.value = false
}

defineExpose({
    open
})
</script>

<style scoped></style>
