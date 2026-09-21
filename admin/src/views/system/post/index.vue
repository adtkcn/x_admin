<template>
    <div class="post-lists">
        <el-card class="border-none!" shadow="never">
            <el-form ref="formRef" class="mb-[-16px]" :model="queryParams" :inline="true">
                <el-form-item class="w-[280px]" label="岗位编码">
                    <el-input v-model="queryParams.code" clearable @keyup.enter="resetPage" />
                </el-form-item>
                <el-form-item class="w-[280px]" label="岗位名称">
                    <el-input v-model="queryParams.name" clearable @keyup.enter="resetPage" />
                </el-form-item>

                <el-form-item>
                    <el-button type="primary" @click="resetPage">查询</el-button>
                    <el-button @click="resetParams">重置</el-button>
                </el-form-item>
            </el-form>
        </el-card>
        <el-card class="border-none! mt-4" shadow="never">
            <div>
                <el-button v-perms="['admin:system:post:add']" type="primary" @click="handleAdd()">
                    <template #icon>
                        <icon name="el-icon-Plus" />
                    </template>
                    新增
                </el-button>
            </div>
            <vxe-table
                class="mt-4"
                v-loading="pager.loading"
                :data="pager.lists"
                :row-config="{ keyField: 'id' }"
                :scroll-y="{ enabled: false }"
                :border="'inner'"
            >
                <vxe-column title="岗位编码" field="code" min-width="100" />
                <vxe-column title="岗位名称" field="name" min-width="100" />
                <vxe-column title="排序" field="sort" min-width="100" />
                <vxe-column title="备注" field="remarks" min-width="100" show-overflow="title" />
                <vxe-column title="添加时间" field="create_time" min-width="180" />
                <vxe-column title="岗位状态" field="is_stop" min-width="100">
                    <template #default="{ row }">
                        <el-tag class="ml-2" :type="row.is_stop ? 'danger' : 'primary'">
                            {{ row.is_stop ? '停用' : '正常' }}
                        </el-tag>
                    </template>
                </vxe-column>
                <vxe-column title="操作" width="120" fixed="right">
                    <template #default="{ row }">
                        <el-button
                            v-perms="['admin:system:post:edit']"
                            type="primary"
                            link
                            @click="handleEdit(row)"
                        >
                            编辑
                        </el-button>
                        <el-button
                            v-perms="['admin:system:post:del']"
                            type="danger"
                            link
                            @click="handleDelete(row.id)"
                        >
                            删除
                        </el-button>
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
    postDelete,
    postLists,
    type type_system_post_list,
    type type_system_post_resp
} from '@/api/org/post'
import { usePaging } from '@/hooks/usePaging'
import feedback from '@/utils/feedback'
import EditPopup from './edit.vue'

defineOptions({
    name: 'post'
})

const editRef = shallowRef<InstanceType<typeof EditPopup>>()
const showEdit = ref(false)
const queryParams = reactive<type_system_post_list>({
    code: '',
    name: '',
    is_stop: -1
})

const { pager, getLists, resetPage, resetParams } = usePaging({
    fetchFun: postLists,
    params: queryParams
})

const handleAdd = async () => {
    showEdit.value = true
    await nextTick()
    editRef.value?.open('add')
}

const handleEdit = async (data: type_system_post_resp) => {
    showEdit.value = true
    await nextTick()
    editRef.value?.open('edit')
    editRef.value?.getDetail(data)
}

const handleDelete = async (id: string) => {
    try {
        await feedback.confirm('确定要删除？')
        await postDelete({ id })
        feedback.msgSuccess('删除成功')
        getLists()
    } catch (error) {
        console.error('岗位删除失败:', error)
    }
}

getLists()
</script>
