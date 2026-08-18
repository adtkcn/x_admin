<template>
    <div class="user-task" style="padding-bottom: 10px">
        <el-card header="审批用户">
            <!-- 审批节点 -->
            <!-- <div>设置审批人（具体人员，部门（负责人），岗位？）</div> -->
            <!-- {{ adminUserList }} -->
            <el-form label-width="80px">
                <el-form-item label="">
                    <el-radio-group v-model="model.user_type">
                        <el-radio :value="1">指定部门、岗位</el-radio>
                        <el-radio :value="2">用户部门负责人</el-radio>
                        <el-radio :value="3">指定审批人</el-radio>
                    </el-radio-group>
                </el-form-item>

                <el-form-item label="指定部门" v-if="model.user_type == 1">
                    <!-- <el-select
                        v-model="model.dept_id"
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
                        v-model="model.dept_id"
                        :data="deptList"
                        :check-strictly="true"
                        default-expand-all
                        :render-after-expand="false"
                        style="width: 100%"
                    />
                </el-form-item>
                <el-form-item label="岗位" v-if="model.user_type == 1">
                    <el-select v-model="model.post_id" placeholder="请选择岗位">
                        <el-option
                            v-for="item in postList"
                            :key="item.value"
                            :label="item.label"
                            :value="item.value"
                        />
                    </el-select>
                </el-form-item>

                <el-form-item label="审批人" v-if="model.user_type == 3">
                    <el-select
                        v-model="model.user_id"
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
import type { FieldListType, UserTaskProps } from './property.type'
import { type type_system_dept_resp } from '@/api/org/department'

// defineModel 直接暴露父层 v-model="nodeProps" 绑定的 user_task 私有属性，可读写
const model = defineModel<UserTaskProps>({ required: true })
const props = defineProps<{
    fieldList?: FieldListType[]
}>()
type LabelValue = {
    label: string
    value: any
}
const adminUserList = ref<LabelValue[]>([])
const deptList = ref<type_system_dept_resp[]>([])
const postList = ref<LabelValue[]>([])

function getAdminList() {
    adminListAll({}).then((res) => {
        adminUserList.value = res.map((item) => {
            return {
                value: item.id,
                label: item.nickname + ' (' + item.email + ')'
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
        deptList.value = arrayToTree<type_system_dept_resp>(list, '')
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
