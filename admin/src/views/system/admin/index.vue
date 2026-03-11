<template>
    <div class="admin">
        <el-card class="border-none!" shadow="never">
            <el-form class="mb-[-16px]" :model="formData" inline>
                <el-form-item label="管理员账号" class="w-[280px]">
                    <el-input v-model="formData.username" clearable @keyup.enter="resetPage" />
                </el-form-item>
                <el-form-item label="管理员名称" class="w-[280px]">
                    <el-input v-model="formData.nickname" clearable @keyup.enter="resetPage" />
                </el-form-item>
                <el-form-item label="管理员角色" class="w-[280px]">
                    <el-select v-model="formData.roleId" :empty-values="[null, undefined]">
                        <el-option label="全部" value="" />
                        <el-option
                            v-for="(item, index) in optionsData.role"
                            :key="index"
                            :label="item.name"
                            :value="item.id"
                        />
                    </el-select>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="resetPage">查询</el-button>
                    <el-button @click="resetParams">重置</el-button>
                </el-form-item>
            </el-form>
        </el-card>
        <el-card v-loading="pager.loading" class="mt-2 border-none!" shadow="never">
            <div>
                <el-button v-perms="['admin:system:admin:add']" type="primary" @click="handleAdd">
                    <template #icon>
                        <icon name="el-icon-Plus" />
                    </template>
                    新增
                </el-button>

                <Upload
                    class="ml-3 mr-3"
                    :url="adminImportFile"
                    :ext="['xlsx']"
                    :show-progress="true"
                    @change="resetPage"
                >
                    <el-button type="primary">
                        <template #icon>
                            <icon name="el-icon-Upload" />
                        </template>
                        导入
                    </el-button>
                </Upload>

                <el-button type="primary" @click="exportFile">
                    <template #icon>
                        <icon name="el-icon-Download" />
                    </template>
                    导出
                </el-button>
            </div>

            <div class="mt-4" style="height: calc(100vh - 340px)">
                <vxe-table
                    :data="pager.lists"
                    :row-config="{
                        keyField: 'id'
                    }"
                    max-height="100%"
                    :border="'inner'"
                >
                    <vxe-column title="头像" width="80">
                        <template #default="{ row }">
                            <el-avatar :size="40" :src="row.avatar"></el-avatar>
                        </template>
                    </vxe-column>
                    <vxe-column title="账号" field="username" min-width="100" />
                    <vxe-column title="名称" field="nickname" min-width="100" />
                    <vxe-column title="角色" field="role" min-width="100" />
                    <vxe-column title="部门" field="dept" min-width="100" />
                    <vxe-column title="创建时间" field="createTime" width="150" />
                    <vxe-column title="最近登录时间" field="lastLoginTime" width="150" />
                    <vxe-column title="最近登录IP" field="lastLoginIp" width="120" />
                    <vxe-column title="状态" width="80">
                        <template #default="{ row }">
                            <el-switch
                                v-perms="['admin:system:admin:disable']"
                                v-if="row.id != 1"
                                :model-value="row.isDisable"
                                :active-value="0"
                                :inactive-value="1"
                                @change="(val) => changeStatus(val as number, row.id)"
                            />
                        </template>
                    </vxe-column>
                    <vxe-column title="操作" width="120" fixed="right">
                        <template #default="{ row }">
                            <el-button
                                v-perms="['admin:system:admin:edit']"
                                type="primary"
                                link
                                @click="handleEdit(row)"
                            >
                                编辑
                            </el-button>
                            <el-button
                                v-if="row.id != 1"
                                v-perms="['admin:system:admin:del']"
                                type="danger"
                                link
                                @click="handleDelete(row.id)"
                            >
                                删除
                            </el-button>
                        </template>
                    </vxe-column>
                </vxe-table>
            </div>
            <div class="flex mt-4 justify-end">
                <pagination v-model="pager" @change="getLists" />
            </div>
        </el-card>
        <edit-popup v-if="showEdit" ref="editRef" @success="getLists" @close="showEdit = false" />
    </div>
</template>

<script lang="ts" setup>
import { ref, shallowRef, reactive, nextTick, onMounted } from 'vue'
import {
    adminLists,
    adminDelete,
    adminStatus,
    adminExportFile,
    adminImportFile,
    type type_system_admin_list,
    type type_system_admin_resp
} from '@/api/perms/admin'
import { roleAll, type type_system_role_simple_resp } from '@/api/perms/role'
import { useDictOptions } from '@/hooks/useDictOptions'
import { usePaging } from '@/hooks/usePaging'
import feedback from '@/utils/feedback'
import EditPopup from './edit.vue'
defineOptions({
    name: 'SystemAdmin'
})
const editRef = shallowRef<InstanceType<typeof EditPopup>>()
// 表单数据
const formData = reactive<type_system_admin_list>({
    username: '',
    nickname: '',
    roleId: ''
})
const showEdit = ref(false)
const { pager, getLists, resetParams, resetPage } = usePaging({
    fetchFun: adminLists,
    params: formData
})

const changeStatus = async (active: number, id: string) => {
    try {
        await feedback.confirm(`确定${active ? '停用' : '开启'}当前管理员？`)
        await adminStatus({ id })
        feedback.msgSuccess('修改成功')
        getLists()
    } catch (error) {
        getLists()
    }
}
const handleAdd = async () => {
    showEdit.value = true
    await nextTick()
    editRef.value?.open('add')
}

const exportFile = async () => {
    await feedback.confirm('确定要导出？')
    await adminExportFile(formData)
}
const handleEdit = async (data: type_system_admin_resp) => {
    showEdit.value = true
    await nextTick()
    editRef.value?.open('edit')
    editRef.value?.setFormData(data)
}

const handleDelete = async (id: string) => {
    try {
        await feedback.confirm('确定要删除？')
        await adminDelete({ id })
        feedback.msgSuccess('删除成功')
        getLists()
    } catch (error) {}
}
const { optionsData } = useDictOptions<{
    role: type_system_role_simple_resp[]
}>({
    role: {
        api: roleAll
    }
})

onMounted(() => {
    getLists()
})
</script>
