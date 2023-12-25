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
            <el-alert title="同一父级的网关只能有一个通过" type="warning" />

            <!-- 设置优先级 -->
            <div style="padding: 40px 0 20px">
                <el-select v-model="selectGateway" placeholder="请选择">
                    <el-option
                        v-for="item in fieldList"
                        :key="item.id"
                        :label="item.label"
                        :value="item.id"
                    />
                </el-select>
                <el-button type="primary" style="margin-left: 10px" @click="addCondition"
                    >添加条件</el-button
                >
            </div>

            <el-table fit size="small" :data="GatewayList" style="width: 100%">
                <el-table-column prop="label" label="表单"></el-table-column>
                <el-table-column label="权限">
                    <template #default="{ row }">
                        <el-select v-model="row.condition" placeholder="请选择审批人">
                            <el-option
                                v-for="item in conditionList"
                                :key="item.value"
                                :label="item.label"
                                :value="item.value"
                            />
                        </el-select>
                    </template>
                </el-table-column>
                <el-table-column label="值">
                    <template #default="{ row }">
                        <el-input v-model="row.conditionValue" placeholder="请输入"></el-input>
                    </template>
                </el-table-column>
            </el-table>
        </div>

        <div v-if="node.type == 'bpmn:endEvent'">结束</div>
    </el-drawer>
</template>

<script>
import UserTask from './PropertyPanel/UserTask.vue'
import FieldAuth from './PropertyPanel/FieldAuth.vue'
export default {
    name: 'PropertyPanel',
    props: {},
    components: {
        UserTask,
        FieldAuth
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

                fieldAuth: {} // 字段权限
            },
            /**
             * 表单列表
             * [{
             *  id: 1,
             *  label: '表单1',
             *  auth: 1,
             * }]
             */
            fieldList: [],
            // 网关条件列表
            GatewayList: [],
            selectGateway: '',

            conditionList: [
                {
                    value: '=',
                    label: '等于'
                },
                {
                    value: '>=',
                    label: '大于等于'
                },
                {
                    value: '<=',
                    label: '小于等于'
                },
                {
                    value: 'include',
                    label: '包含'
                }
            ]
        }
    },

    methods: {
        open(node, fieldList) {
            this.node = node

            this.properties.userType = node?.properties?.userType || ''
            this.properties.userId = node?.properties?.userId || ''
            this.properties.deptId = node?.properties?.deptId || ''
            this.properties.postId = node?.properties?.postId || ''

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

            this.setProperties('fieldAuth', {
                ...fieldAuth
            })
            this.setProperties('userType', this.properties.userType)
            this.setProperties('userId', this.properties.userId)
            this.setProperties('deptId', this.properties.deptId)
            this.setProperties('postId', this.properties.postId)
        },
        // getAdminList() {
        //     adminLists().then((res) => {
        //         console.log('res', res)
        //         this.adminUserList = res.lists.map((item) => {
        //             return {
        //                 value: item.id,
        //                 label: item.nickname
        //             }
        //         })
        //     })
        // },
        // getDeptList() {
        //     deptLists().then((res) => {
        //         console.log('res', res)
        //         this.deptList = res.map((item) => {
        //             return {
        //                 value: item.id,
        //                 label: item.name
        //             }
        //         })
        //     })
        // },
        // getPostList() {
        //     postAll().then((res) => {
        //         console.log('res', res)
        //         this.postList = res.map((item) => {
        //             return {
        //                 value: item.id,
        //                 label: item.name
        //             }
        //         })
        //     })
        // },
        setProperties(key, val) {
            this.$emit('setProperties', this.node, {
                [key]: val
            })
        },

        addCondition() {
            // this.selectGateway

            this.fieldList.find((item) => {
                if (item.id === this.selectGateway) {
                    // item.condition = this.GatewayList[this.GatewayList.length - 1].condition
                    // item.value = this.GatewayList[this.GatewayList.length - 1].value
                    this.GatewayList.push({
                        id: item.id,
                        label: item.label,
                        value: '',
                        condition: ''
                    })
                }
            })
        }
    }
}
</script>

<style scoped></style>
