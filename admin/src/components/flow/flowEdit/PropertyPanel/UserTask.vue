<template>
    <div class="user-task" style="padding-bottom: 10px">
        <el-card header="审批用户">
            <!-- 审批节点 -->
            <!-- <div>设置审批人（具体人员，部门（负责人），岗位？）</div> -->
            <!-- {{ adminUserList }} -->
            <el-form label-width="80px">
                <el-form-item label="">
                    <el-radio-group v-model="props.properties.userType">
                        <el-radio :label="1">指定部门、岗位</el-radio>
                        <el-radio :label="2">用户部门负责人</el-radio>
                        <el-radio :label="3">指定审批人</el-radio>
                    </el-radio-group>
                </el-form-item>

                <el-form-item label="指定部门" v-if="props.properties.userType == 1">
                    <!-- <el-select
                        v-model="props.properties.deptId"
                        placeholder="请选择审批部门"
                        style="width: 100%"
                    >
                        <el-option
                            v-for="item in deptList"
                            :key="item.value"
                            :label="item.label"
                            :value="item.value"
                        />
                    </el-select> -->
                    <el-tree-select
                        v-model="props.properties.deptId"
                        :data="deptList"
                        :check-strictly="true"
                        default-expand-all
                        :render-after-expand="false"
                        style="width: 100%"
                    />
                </el-form-item>
                <el-form-item label="岗位" v-if="[1].includes(props.properties.userType)">
                    <el-select v-model="props.properties.postId" placeholder="请选择岗位">
                        <el-option
                            v-for="item in postList"
                            :key="item.value"
                            :label="item.label"
                            :value="item.value"
                        />
                    </el-select>
                </el-form-item>

                <el-form-item label="审批人" v-if="props.properties.userType == 3">
                    <el-select
                        v-model="props.properties.userId"
                        placeholder="请选择审批人"
                        style="width: 100%"
                        clearable
                        filterable
                    >
                        <el-option
                            v-for="item in adminUserList"
                            :key="item.value"
                            :label="item.label"
                            :value="item.value"
                        />
                    </el-select>
                </el-form-item>
            </el-form>
        </el-card>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { adminListAll } from '@/api/perms/admin'
import { deptAll } from '@/api/org/department'
import { postAll } from '@/api/org/post'
import { arrayToTree } from '@/utils/util'
import type { NodeType, PropertiesType, FieldListType } from './property.type'

const props = defineProps<{
    node?: NodeType
    fieldList?: FieldListType[]
    properties?: PropertiesType
}>()

const adminUserList = ref([])
const deptList = ref([])
const postList = ref([])

function getAdminList() {
    adminListAll({}).then((res) => {
        adminUserList.value = res.map((item) => {
            return {
                value: item.id,
                label: item.nickname + ' (' + item.username + ')'
            }
        })
    })
}
function getDeptList() {
    deptAll().then((res) => {
        const list = res.map((item) => {
            return {
                value: item.id,
                label: item.name,
                ...item
            }
        })
        deptList.value = arrayToTree(list, '')
    })
}
function getPostList() {
    postAll().then((res) => {
        postList.value = res.map((item) => {
            return {
                value: item.id,
                label: item.name
            }
        })
    })
}
onMounted(() => {
    getAdminList()
    getDeptList()
    getPostList()
})
</script>

<style lang="scss">
.user-task {
    user-select: none;
}
</style>
