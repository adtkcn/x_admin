<template>
    <div class="user-lists">
        <el-card class="border-none!" shadow="never">
            <el-form class="mb-[-16px]" :model="queryParams" :inline="true">
                <el-form-item class="w-[220px]" label="关键词">
                    <el-input
                        v-model="queryParams.keyword"
                        clearable
                        placeholder="昵称/邮箱/手机"
                        @keyup.enter="resetPage"
                    />
                </el-form-item>
                <el-form-item class="w-[180px]" label="状态">
                    <el-select v-model="queryParams.status" clearable placeholder="全部">
                        <el-option label="启用" :value="1" />
                        <el-option label="禁用" :value="0" />
                    </el-select>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="resetPage">查询</el-button>
                    <el-button @click="resetParams">重置</el-button>
                </el-form-item>
            </el-form>
        </el-card>
        <el-card class="border-none! mt-4" shadow="never">
            <vxe-table
                class="mt-4"
                v-loading="pager.loading"
                :data="pager.lists"
                :row-config="{ keyField: 'id' }"
                :scroll-y="{ enabled: false }"
                :border="'inner'"
                :tooltip-config="tooltipConfig"
                show-overflow
            >
                <vxe-column title="用户ID" field="id" min-width="200" show-overflow />
                <vxe-column title="昵称" field="nickname" min-width="140" />
                <vxe-column title="邮箱" field="email" min-width="180" show-overflow />
                <vxe-column title="手机" min-width="140">
                    <template #default="{ row }">{{ row.phone_code }} {{ row.phone }}</template>
                </vxe-column>
                <vxe-column title="状态" field="status" min-width="100">
                    <template #default="{ row }">
                        <el-tag :type="row.status ? 'success' : 'info'">
                            {{ row.status ? '启用' : '禁用' }}
                        </el-tag>
                    </template>
                </vxe-column>
                <vxe-column title="最后登录" field="last_login_time" min-width="180" />
                <vxe-column title="创建时间" field="create_time" min-width="180" />
                <vxe-column title="操作" width="220" fixed="right">
                    <template #default="{ row }">
                        <el-button
                            v-perms="['admin:user:edit']"
                            type="primary"
                            link
                            @click="handleEdit(row)"
                            >编辑</el-button
                        >
                        <el-button
                            v-perms="['admin:user:disable']"
                            :type="row.status ? 'warning' : 'success'"
                            link
                            @click="handleDisable(row)"
                        >
                            {{ row.status ? '禁用' : '启用' }}
                        </el-button>
                        <el-button
                            v-perms="['admin:user:kick']"
                            type="danger"
                            link
                            @click="handleKick(row)"
                            >踢下线</el-button
                        >
                    </template>
                </vxe-column>
            </vxe-table>
            <div class="flex justify-end mt-4">
                <pagination v-model="pager" @change="getLists" />
            </div>
        </el-card>
        <edit-popup v-if="showEdit" ref="editRef" @success="getLists" @close="showEdit = false" />
    </div>
</template>
<script lang="ts" setup>
import { ref, shallowRef, reactive, nextTick } from 'vue'
import {
    userLists,
    userDisable,
    userKick,
    type type_user_list,
    type type_user_resp
} from '@/api/user'
import { usePaging } from '@/hooks/usePaging'
import feedback from '@/utils/feedback'
import EditPopup from './edit.vue'
import type { VxeTablePropTypes } from 'vxe-table'

defineOptions({ name: 'user' })
const tooltipConfig = reactive<VxeTablePropTypes.TooltipConfig<type_user_resp>>({
    enterable: true
})
const editRef = shallowRef<InstanceType<typeof EditPopup>>()
const showEdit = ref(false)
const queryParams = reactive<type_user_list>({ keyword: '', status: '' })

const { pager, getLists, resetPage, resetParams } = usePaging({
    fetchFun: userLists,
    params: queryParams
})

const handleEdit = async (data: type_user_resp) => {
    showEdit.value = true
    await nextTick()
    editRef.value?.open('edit')
    editRef.value?.getDetail(data)
}
const handleDisable = async (row: type_user_resp) => {
    try {
        await feedback.confirm(row.status ? '确定要禁用该用户？' : '确定要启用该用户？')
        await userDisable({ id: row.id, status: row.status ? 0 : 1 })
        feedback.msgSuccess('操作成功')
        getLists()
    } catch (error) {
        console.error('用户禁用/启用失败:', error)
    }
}
const handleKick = async (row: type_user_resp) => {
    try {
        await feedback.confirm('确定要踢该用户下线？')
        await userKick({ id: row.id })
        feedback.msgSuccess('操作成功')
        getLists()
    } catch (error) {
        console.error('用户踢下线失败:', error)
    }
}
getLists()
</script>
