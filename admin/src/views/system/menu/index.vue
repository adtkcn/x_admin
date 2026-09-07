<template>
    <div class="menu-lists p-4 h-full box-border flex flex-col">
        <div>
            <el-input
                v-model="menuName"
                style="width: 240px"
                clearable
                placeholder="请输入菜单名称"
            />

            <el-button
                v-perms="['admin:system:menu:add']"
                type="primary"
                @click="handleAdd()"
                class="ml-4"
            >
                <template #icon>
                    <icon name="el-icon-Plus" />
                </template>
                新增
            </el-button>
            <el-button @click="handleExpand"> 展开/收起 </el-button>
        </div>

        <div class="mt-4" style="height: 100%">
            <vxe-table
                ref="tableRef"
                :row-config="rowConfig"
                :row-drag-config="rowDragConfig"
                :tree-config="treeConfig"
                :data="filterList"
                :border="'inner'"
                height="100%"
                :virtual-y-config="{ enabled: true, gt: 0 }"
                @row-dragend="rowDragend"
            >
                <vxe-column type="seq" width="60"></vxe-column>
                <vxe-column
                    field="menu_name"
                    title="菜单名称"
                    min-width="200"
                    tree-node
                ></vxe-column>
                <vxe-column field="menu_type" title="类型" width="60">
                    <template #default="{ row }">
                        <div v-if="row.menu_type == MenuEnum.CATALOGUE">目录</div>
                        <div v-else-if="row.menu_type == MenuEnum.MENU">菜单</div>
                        <div v-else-if="row.menu_type == MenuEnum.BUTTON">按钮</div>
                    </template>
                </vxe-column>
                <vxe-column field="menu_icon" title="图标" width="60">
                    <template #default="{ row }">
                        <div class="flex">
                            <icon :name="row.menu_icon" :size="20" />
                        </div>
                    </template>
                </vxe-column>
                <vxe-column field="paths" title="路径" min-width="100"></vxe-column>
                <vxe-column field="permsArr" title="权限标识" min-width="120">
                    <template #default="{ row }">
                        <span v-if="row.perms" type="info">{{ row.perms }}</span>
                    </template>
                </vxe-column>
                <vxe-column field="is_disable" title="状态" width="80">
                    <template #default="{ row }">
                        <el-tag v-if="row.is_disable == 0" type="primary">正常</el-tag>
                        <el-tag v-else type="danger">停用</el-tag>
                    </template>
                </vxe-column>
                <vxe-column title="排序" width="60" drag-sort></vxe-column>
                <vxe-column title="操作" width="160" align="right">
                    <template #default="{ row }">
                        <el-button
                            v-perms="['admin:system:menu:add']"
                            v-show="row.menu_type != MenuEnum.BUTTON"
                            type="primary"
                            link
                            @click="handleAdd(row.id)"
                        >
                            新增
                        </el-button>
                        <el-button
                            v-perms="['admin:system:menu:edit']"
                            type="primary"
                            link
                            @click="handleEdit(row)"
                        >
                            编辑
                        </el-button>
                        <el-button
                            v-perms="['admin:system:menu:del']"
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
    </div>
    <EditPopup v-if="showEdit" ref="editRef" @success="getLists" @close="showEdit = false" />
</template>
<script lang="ts" setup>
import { ref, reactive, useTemplateRef, nextTick, computed, toRaw } from 'vue'
import {
    menuDelete,
    menuLists,
    menuSort,
    type type_system_menu_resp,
    type type_system_menu_edit
} from '@/api/perms/menu'
import { MenuEnum } from '@/enums/appEnums'
import EditPopup from './edit.vue'
import feedback from '@/utils/feedback'
import { queryHierarchy } from '@/utils/flatTreeUtils'

import { VxeUI, VxeTableInstance, VxeTableEvents, VxeTablePropTypes } from 'vxe-table'

defineOptions({
    name: 'MenuView'
})
const rowConfig = {
    keyField: 'id',
    drag: true
}
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
    // 顶级菜单无父行，从全量树形根级取同级；子级取父行的 children
    const parentRow = $table.getTreeParentRow(oldRow)
    const siblings = (parentRow ? parentRow.children : $table.getFullData()) as any[]
    const ids = siblings.map((item) => item.id)
    try {
        await menuSort({ ids })
        feedback.msgSuccess('排序已保存')
    } finally {
        getLists() // 重新拉取以应用最新排序
    }
}
const treeConfig = reactive<VxeTablePropTypes.TreeConfig>({
    rowField: 'id',
    childrenField: 'children',
    indent: 10,
    reserve: true,
    lazy: true,
    transform: true,

    parentField: 'pid'
})
const tableRef = useTemplateRef<VxeTableInstance<type_system_menu_resp>>('tableRef')
const editRef = useTemplateRef<InstanceType<typeof EditPopup>>('editRef')

const loading = ref(false)
const showEdit = ref(false)
const lists = ref<type_system_menu_resp[]>([])
const menuName = ref('')
const filterList = computed(() => {
    if (!menuName.value) {
        return lists.value
    }
    const raw = toRaw(lists.value)

    const { all } = queryHierarchy(raw, menuName.value, {
        fields: ['menu_name'],
        exact: false
    })
    return all
})
const getLists = async () => {
    loading.value = true
    try {
        const data = await menuLists()
        lists.value = data
        loading.value = false
    } catch (error) {
        console.error('菜单列表获取失败:', error)
        loading.value = false
    }
}

const handleAdd = async (id?: string) => {
    showEdit.value = true
    await nextTick()
    if (id) {
        editRef.value?.setFormData({
            pid: id
        } as type_system_menu_edit)
    }
    editRef.value?.open('add')
}

const handleEdit = async (data: type_system_menu_resp) => {
    showEdit.value = true
    await nextTick()
    editRef.value?.open('edit')
    editRef.value?.getDetail(data)
}

const handleDelete = async (id: string) => {
    try {
        await feedback.confirm('确定要删除？')
        await menuDelete({ id })
        feedback.msgSuccess('删除成功')
        getLists()
    } catch (error) {
        console.error('菜单删除失败:', error)
    }
}

let isExpand = false
const handleExpand = () => {
    const $table = tableRef.value
    if ($table) {
        isExpand = !isExpand
        if (isExpand) {
            $table.setAllTreeExpand(true)
        } else {
            $table.clearTreeExpand()
        }
    }
}
getLists()
</script>
<style scoped lang="scss">
.menu-lists {
    background-color: white;
}
</style>
