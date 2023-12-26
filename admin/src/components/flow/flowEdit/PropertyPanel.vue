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

<script>
import UserTask from './PropertyPanel/UserTask.vue'
import FieldAuth from './PropertyPanel/FieldAuth.vue'
import Gateway from './PropertyPanel/Gateway.vue'

export default {
    name: 'PropertyPanel',
    props: {},
    components: {
        UserTask,
        FieldAuth,
        Gateway
    },
    data() {
        return {
            drawerVisible: false,

            node: {},
            properties: {
                userType: '',
                userId: '', //审批人id
                deptId: '', //审批部门id
                postId: '', //岗位id

                fieldAuth: {}, // 字段权限
                gateway: [] //网关条件列表
            },
            /**
             * 表单列表
             * [{
             *  id: 1,
             *  label: '表单1',
             *  auth: 1,
             * }]
             */
            fieldList: []
        }
    },

    methods: {
        open(node, fieldList) {
            this.node = node

            this.properties.userType = node?.properties?.userType || ''
            this.properties.userId = node?.properties?.userId || ''
            this.properties.deptId = node?.properties?.deptId || ''
            this.properties.postId = node?.properties?.postId || ''

            this.properties.gateway = node?.properties?.gateway || []

            this.properties.fieldAuth = node?.properties?.fieldAuth
                ? { ...node?.properties?.fieldAuth }
                : {}

            this.fieldList = fieldList.map((item) => {
                let auth = 1
                const formId = item?.field?.id
                if (node?.properties?.fieldAuth?.[formId]) {
                    auth = node.properties.fieldAuth[formId]
                }
                return {
                    id: formId,
                    label: item?.field?.options?.label,
                    auth: auth
                }
            })
            this.drawerVisible = true
        },
        close() {
            const fieldAuth = {}
            this.fieldList.forEach((item) => {
                fieldAuth[item.id] = item.auth
            })
            this.$emit('setProperties', this.node, { ...this.properties })

            // this.setProperties('fieldAuth', {
            //     ...fieldAuth
            // })
            // this.setProperties('userType', this.properties.userType)
            // this.setProperties('userId', this.properties.userId)
            // this.setProperties('deptId', this.properties.deptId)
            // this.setProperties('postId', this.properties.postId)
            // this.setProperties('gateway', this.properties.gateway)
        },

        setProperties(key, val) {
            this.$emit('setProperties', this.node, {
                [key]: val
            })
        }
    }
}
</script>

<style scoped></style>
