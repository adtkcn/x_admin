<template>
    <el-drawer v-model="drawerVisible" size="500px" :title="node?.text?.value" @close="close">
        <div class="setting-block">
            <!-- 开始节点 -->
            <!-- {{ node }} -->

            <div v-if="node.type == 'bpmn:startEvent'">开始节点</div>
            <div v-if="node.type == 'bpmn:userTask'">
                审批节点
                <!-- <div>设置审批人（具体人员，部门（负责人），岗位？）</div> -->
                <!-- {{ adminUserList }} -->
                <el-form label-width="80px">
                    <el-form-item label="审批人">
                        <el-select v-model="properties.userId" placeholder="请选择审批人">
                            <el-option
                                v-for="item in adminUserList"
                                :key="item.value"
                                :label="item.label"
                                :value="item.value"
                            />
                        </el-select>
                    </el-form-item>

                    <el-form-item label="审批部门">
                        <el-select v-model="properties.deptId" placeholder="请选择审批部门">
                            <el-option
                                v-for="item in deptList"
                                :key="item.value"
                                :label="item.label"
                                :value="item.value"
                            />
                        </el-select>
                    </el-form-item>
                    <el-form-item label="岗位">
                        <el-select v-model="properties.postId" placeholder="请选择岗位">
                            <el-option
                                v-for="item in postList"
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
                设置优先级
                <el-table fit size="small" :data="fieldList" style="width: 100%">
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

            <!-- 网关和结束节点不需要表单权限设置 -->
            <el-table
                v-if="['bpmn:startEvent', 'bpmn:userTask'].includes(node.type)"
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
import { deptLists } from '@/api/org/department'
import { postAll } from '@/api/org/post'

export default {
    name: 'PropertyPanel',
    props: {},
    data() {
        return {
            drawerVisible: false,
            adminUserList: [],
            deptList: [],
            postList: [],

            node: {},
            properties: {
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
            conditionsList: [],

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

            this.properties.userId = node?.properties?.userId || ''
            this.properties.deptId = node?.properties?.deptId || ''
            this.properties.postId = node?.properties?.postId || ''
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
            this.getDeptList()
            this.getPostList()
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
            this.setProperties('userId', this.properties.userId)
            this.setProperties('deptId', this.properties.deptId)
            this.setProperties('postId', this.properties.postId)
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
        getDeptList() {
            deptLists().then((res) => {
                console.log('res', res)
                this.deptList = res.map((item) => {
                    return {
                        value: item.id,
                        label: item.name
                    }
                })
            })
        },
        getPostList() {
            postAll().then((res) => {
                console.log('res', res)
                this.postList = res.map((item) => {
                    return {
                        value: item.id,
                        label: item.name
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
