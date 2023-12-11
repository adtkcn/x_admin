<template>
    <el-drawer v-model="drawerVisible" size="500px" :title="node?.text?.value" @close="close">
        <div class="setting-block">
            <!-- 开始节点 -->
            {{ node }}

            <div v-if="node.type == 'bpmn:startEvent'">开始节点</div>
            <div v-if="node.type == 'bpmn:userTask'">
                审批节点
                <div>设置审批人（具体人员，部门（负责人），岗位？）</div>
                <!-- {{ adminUserList }} -->
                <el-form>
                    <el-form-item label="审批人">
                        <el-select v-model="properties.user" placeholder="请选择审批人">
                            <el-option
                                v-for="item in adminUserList"
                                :key="item.value"
                                :label="item.label"
                                :value="item.value"
                            />
                        </el-select>
                    </el-form-item>
                </el-form>
            </div>

            <div v-if="node.type == 'bpmn:serviceTask'">
                <div>系统任务</div>
                <div>抄送</div>
                <div>发送邮件</div>
                <div>发送短信</div>
                <div>发送站内消息</div>
            </div>
            <div v-if="node.type == 'bpmn:exclusiveGateway'">
                <div>网关，只能有一个网关通过</div>
                <div>从form取值判断</div>
            </div>

            <div v-if="node.type == 'bpmn:endEvent'">结束</div>

            <!-- 网关和结束节点不需要表单权限设置 -->
            <el-table
                v-if="['bpmn:startEvent', 'bpmn:userTask', 'bpmn:serviceTask'].includes(node.type)"
                fit
                size="small"
                :data="fieldList"
                style="width: 100%"
            >
                <el-table-column prop="label" label="表单"></el-table-column>
                <el-table-column label="权限">
                    <template #default="{ row }">
                        <el-radio-group v-model="row.auth">
                            <el-radio :label="1">读写</el-radio>
                            <el-radio :label="2">可读</el-radio>
                            <el-radio :label="3">隐藏</el-radio>
                        </el-radio-group>
                    </template>
                </el-table-column>
            </el-table>
            <!-- {{ fieldList }} -->
            <!-- 用户审批节点 -->

            <!-- 系统任务节点 -->
        </div>
    </el-drawer>
</template>

<script>
import { adminLists } from '@/api/perms/admin'

export default {
    name: 'PropertyPanel',
    props: {},
    data() {
        return {
            drawerVisible: false,
            adminUserList: [],

            node: {},
            properties: {
                user: '', //审批人id
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
            fieldList: []
        }
    },

    methods: {
        open(node, fieldList) {
            this.node = node

            this.properties.user = node?.properties?.user || ''
            this.properties.fieldAuth = node?.properties?.fieldAuth || {}
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
            this.getAdminList()
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
            this.setProperties('user', this.properties.user)
        },
        getAdminList() {
            adminLists().then((res) => {
                console.log('res', res)
                this.adminUserList = res.lists.map((item) => {
                    return {
                        value: item.id,
                        label: item.nickname
                    }
                })
            })
        },
        setProperties(key, val) {
            this.$emit('setProperties', this.node, {
                [key]: val
            })
        }
    },
    components: {}
}
</script>

<style scoped></style>
