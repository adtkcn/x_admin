<template>
    <div class="device-space-lists">
        <el-card class="border-none!" shadow="never">
            <div>
                <el-button v-perms="['admin:iot:deviceSpace:add']" type="primary" @click="handleAdd()">
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
                    title="区域名称"
                    field="name"
                    min-width="200"
                    show-overflow="title"
                    tree-node
                />
                <vxe-column title="区域范围" field="geojson" width="120">
                    <template #default="{ row }">
                        <el-tag :type="row.geojson ? 'success' : 'info'">
                            {{ row.geojson ? '已绘制' : '未绘制' }}
                        </el-tag>
                    </template>
                </vxe-column>
                <vxe-column title="备注" field="remark" min-width="150" show-overflow="title" />
                <vxe-column
                    title="排序"
                    width="100"
                    drag-sort
                    v-perms="['admin:iot:deviceSpace:sort']"
                />
                <vxe-column title="更新时间" field="update_time" width="180" />
                <vxe-column title="操作" width="160" fixed="right">
                    <template #default="{ row }">
                        <el-button
                            v-perms="['admin:iot:deviceSpace:add']"
                            type="primary"
                            link
                            @click="handleAdd(row.id)"
                        >
                            新增
                        </el-button>
                        <el-button
                            v-perms="['admin:iot:deviceSpace:edit']"
                            type="primary"
                            link
                            @click="handleEdit(row)"
                        >
                            编辑
                        </el-button>
                        <el-button
                            v-perms="['admin:iot:deviceSpace:del']"
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
import {
    deviceSpaceAll,
    deviceSpaceDel,
    deviceSpaceSort,
    type type_iot_device_space
} from '@/api/iot'
import feedback from '@/utils/feedback'

defineOptions({
    name: 'iotDeviceSpace'
})

const tableRef = ref<any>()
const editRef = shallowRef<InstanceType<typeof EditPopup>>()
const loading = ref(false)
const lists = ref<type_iot_device_space[]>([])

const rowDragConfig = reactive({
    isPeerDrag: true,
    showGuidesStatus: true,
    showIcon: true
})

const rowDragend = async ({ oldRow }: any) => {
    const $table = tableRef.value
    if (!$table) return
    const parentRow = $table.getTreeParentRow(oldRow)
    const siblings = (parentRow ? parentRow.children : $table.getFullData()) as any[]
    const ids = siblings.map((item) => item.id)
    try {
        await deviceSpaceSort({ ids })
        feedback.msgSuccess('排序已保存')
    } finally {
        getLists()
    }
}

const showEdit = ref(false)
const getLists = async () => {
    loading.value = true
    try {
        lists.value = await deviceSpaceAll()
    } catch (error) {
        console.error('设备区域列表获取失败:', error)
    }
    loading.value = false
}

const handleAdd = async (id?: string) => {
    showEdit.value = true
    await nextTick()
    if (id) {
        editRef.value?.setFormData({ pid: id })
    }
    editRef.value?.open('add')
}

const handleEdit = async (data: type_iot_device_space) => {
    showEdit.value = true
    await nextTick()
    editRef.value?.open('edit')
    editRef.value?.getDetail(data)
}

const handleDelete = async (id: string) => {
    try {
        await feedback.confirm('确定要删除？')
        await deviceSpaceDel({ id })
        feedback.msgSuccess('删除成功')
        getLists()
    } catch (error) {
        console.error('区域删除失败:', error)
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

onMounted(() => {
    getLists()
})
</script>
