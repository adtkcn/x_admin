<template>
    <div class="dept-lists">
        <el-card class="border-none!" shadow="never">
            <div>
                <el-button v-perms="['admin:system:dept:add']" type="primary" @click="handleAdd()">
                    <template #icon>
                        <icon name="el-icon-Plus" />
                    </template>
                    新增
                </el-button>
                <el-button @click="handleExpand"> 展开/折叠 </el-button>
            </div>
            <vxe-table
                ref="tableRef"
                class="mt-4"
                v-loading="loading"
                :data="lists"
                :row-config="{ keyField: 'id', drag: true }"
                :tree-config="{
                    transform: true,
                    rowField: 'id',
                    parentField: 'pid',
                    childrenField: 'children',
                    expandAll: true
                }"
                :row-drag-config="rowDragConfig"
                :border="'inner'"
                @row-dragend="rowDragend"
            >
                <vxe-column
                    title="部门名称"
                    field="name"
                    min-width="150"
                    show-overflow="title"
                    tree-node
                />
                <vxe-column title="负责人" field="duty" min-width="150" show-overflow="title" />

                <vxe-column title="部门状态" field="is_stop" width="100">
                    <template #default="{ row }">
                        <el-tag class="ml-2" :type="row.is_stop ? 'danger' : 'primary'">
                            {{ row.is_stop ? '停用' : '正常' }}
                        </el-tag>
                    </template>
                </vxe-column>
                <vxe-column
                    title="排序"
                    width="100"
                    drag-sort
                    v-perms="['admin:system:dept:sort']"
                />
                <vxe-column title="更新时间" field="update_time" width="180" />
                <vxe-column title="操作" width="160" fixed="right">
                    <template #default="{ row }">
                        <el-button
                            v-perms="['admin:system:dept:add']"
                            type="primary"
                            link
                            @click="handleAdd(row.id)"
                        >
                            新增
                        </el-button>
                        <el-button
                            v-perms="['admin:system:dept:edit']"
                            type="primary"
                            link
                            @click="handleEdit(row)"
                        >
                            编辑
                        </el-button>
                        <el-button
                            v-if="row.pid !== 0"
                            v-perms="['admin:system:dept:del']"
                            type="danger"
                            link
                            @click="handleDelete(row.id)"
                        >
                            删除
                        </el-button>
                    </template>
                </vxe-column>
            </vxe-table>
        </el-card>
        <edit-popup v-if="showEdit" ref="editRef" @success="getLists" @close="showEdit = false" />
    </div>
</template>
<script lang="ts" setup>
import { ref, shallowRef, reactive, nextTick, onMounted } from 'vue'
import EditPopup from './edit.vue'
import { deptDelete, deptAll, deptSort, type type_system_dept_resp } from '@/api/org/department'
import feedback from '@/utils/feedback'
defineOptions({
    name: 'department'
})
const tableRef = ref<any>()
const editRef = shallowRef<InstanceType<typeof EditPopup>>()
const loading = ref(false)
const lists = ref<type_system_dept_resp[]>([])

// 行拖拽配置：仅允许同级之间拖拽排序
const rowDragConfig = reactive({
    isPeerDrag: true,
    showGuidesStatus: true,
    showIcon: true
})

// 拖拽结束：提取被拖拽行的同级新顺序并持久化
const rowDragend = async ({ oldRow }: any) => {
    const $table = tableRef.value
    if (!$table) return
    const parentRow = $table.getTreeParentRow(oldRow)
    // 顶级部门无父行，从全量树形根级取同级；子级取父行的 children
    const siblings = (parentRow ? parentRow.children : $table.getFullData()) as any[]
    const ids = siblings.map((item) => item.id)
    try {
        await deptSort({ ids })
        feedback.msgSuccess('排序已保存')
    } finally {
        getLists() // 重新拉取以应用最新排序
    }
}

const showEdit = ref(false)
const getLists = async () => {
    loading.value = true
    try {
        lists.value = await deptAll()
    } catch (error) {
        console.error('部门列表获取失败:', error)
    }
    loading.value = false
}

const handleAdd = async (id?: string) => {
    showEdit.value = true
    await nextTick()
    if (id) {
        editRef.value?.setFormData({
            pid: id
        })
    }
    editRef.value?.open('add')
}

const handleEdit = async (data: type_system_dept_resp) => {
    showEdit.value = true
    await nextTick()
    editRef.value?.open('edit')
    editRef.value?.getDetail(data)
}

const handleDelete = async (id: string) => {
    try {
        await feedback.confirm('确定要删除？')
        await deptDelete({ id })
        feedback.msgSuccess('删除成功')
        getLists()
    } catch (error) {
        console.error('部门删除失败:', error)
    }
}

let isExpand = true
const handleExpand = () => {
    const $table = tableRef.value
    if (!$table) return
    isExpand = !isExpand
    if (isExpand) {
        $table.setAllTreeExpand(true)
    } else {
        $table.clearTreeExpand()
    }
}

onMounted(async () => {
    await getLists()
    // 等待 vxe-table 内部完成 tree 数据初始化后再展开，避免 setAllTreeExpand 被静默丢弃
    // await nextTick()
    // setTimeout(() => {
    //     handleExpand()
    // }, 0)
})
</script>
