<template>
    <div class="role-lists">
        <el-card class="!border-none" shadow="never">
            <div>
                <el-button v-perms="['admin:system:role:add']" type="primary" @click="handleAdd">
                    <template #icon>
                        <icon name="el-icon-Plus" />
                    </template>
                    新增
                </el-button>
            </div>
            <div class="mt-4">
                <div style="height: calc(100vh - 270px)">
                    <vxe-table
                        :data="pager.lists"
                        :row-config="{
                            keyField: 'id'
                        }"
                        max-height="100%"
                        :border="'inner'"
                        v-loading="pager.loading"
                    >
                        <vxe-column field="id" title="ID" min-width="60" />
                        <vxe-column field="name" title="名称" min-width="150" />
                        <vxe-column field="remark" title="备注" min-width="150" />
                        <vxe-column field="sort" title="排序" min-width="100" />
                        <vxe-column field="member" title="角色人数" min-width="100" />
                        <vxe-column title="岗位状态" field="isDisable" min-width="80">
                            <template #default="{ row }">
                                <el-tag class="ml-2" :type="row.isDisable ? 'danger' : 'primary'">
                                    {{ row.isDisable ? '停用' : '正常' }}
                                </el-tag>
                            </template>
                        </vxe-column>

                        <vxe-column field="createTime" title="创建时间" width="150" />
                        <vxe-column title="操作" width="190" fixed="right">
                            <template #default="{ row }">
                                <el-button
                                    v-perms="['admin:system:role:edit']"
                                    link
                                    type="primary"
                                    @click="handleEdit(row)"
                                >
                                    编辑
                                </el-button>
                                <el-button
                                    v-perms="['admin:system:role:edit']"
                                    link
                                    type="primary"
                                    @click="handleAuth(row)"
                                >
                                    权限设置
                                </el-button>
                                <el-button
                                    v-perms="['admin:system:role:del']"
                                    link
                                    type="danger"
                                    @click="handleDelete(row.id)"
                                >
                                    删除
                                </el-button>
                            </template>
                        </vxe-column>
                    </vxe-table>
                </div>
                <div class="flex justify-end mt-4">
                    <pagination v-model="pager" @change="getLists" />
                </div>
            </div>
        </el-card>
        <edit-popup v-if="showEdit" ref="editRef" @success="getLists" @close="showEdit = false" />
        <auth-popup v-if="showAuth" ref="authRef" @success="getLists" @close="showAuth = false" />
    </div>
</template>

<script lang="ts" setup>
import { ref, useTemplateRef, nextTick } from 'vue'
import { roleLists, roleDelete } from '@/api/perms/role'
import { usePaging } from '@/hooks/usePaging'
import feedback from '@/utils/feedback'
import EditPopup from './edit.vue'
import AuthPopup from './auth.vue'
defineOptions({
    name: 'role'
})

const editRef = useTemplateRef<InstanceType<typeof EditPopup>>('editRef')
const authRef = useTemplateRef<InstanceType<typeof AuthPopup>>('authRef')
const showEdit = ref(false)
const showAuth = ref(false)
const { pager, getLists } = usePaging({
    fetchFun: roleLists
})
const handleAdd = async () => {
    showEdit.value = true
    await nextTick()
    editRef.value?.open('add')
}

const handleEdit = async (data: any) => {
    showEdit.value = true
    await nextTick()
    editRef.value?.open('edit')
    editRef.value?.setFormData(data)
}

const handleAuth = async (data: any) => {
    showAuth.value = true
    await nextTick()
    authRef.value?.open()
    authRef.value?.setFormData(data)
}

// 删除角色
const handleDelete = async (id: number) => {
    await feedback.confirm('确定要删除？')
    await roleDelete({ id })
    feedback.msgSuccess('删除成功')
    getLists()
}

getLists()
</script>
