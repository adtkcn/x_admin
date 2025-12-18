<template>
    <div class="menu-lists p-4 h-full box-border flex flex-col">
        <div>
            <el-input v-model="menuName" style="width: 240px" placeholder="Please input" />

            <el-button v-perms="['admin:system:menu:add']" type="primary" @click="handleAdd()">
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
                :tree-config="treeConfig"
                :data="lists"
                :border="'inner'"
                height="100%"
                :virtual-y-config="{ enabled: true, gt: 0 }"
            >
                <vxe-column type="seq" width="60"></vxe-column>
                <vxe-column
                    field="menuName"
                    title="菜单名称"
                    min-width="200"
                    tree-node
                ></vxe-column>
                <vxe-column field="menuType" title="类型" width="60">
                    <template #default="{ row }">
                        <div v-if="row.menuType == MenuEnum.CATALOGUE">目录</div>
                        <div v-else-if="row.menuType == MenuEnum.MENU">菜单</div>
                        <div v-else-if="row.menuType == MenuEnum.BUTTON">按钮</div>
                    </template>
                </vxe-column>
                <vxe-column field="menuIcon" title="图标" width="60">
                    <template #default="{ row }">
                        <div class="flex">
                            <icon :name="row.menuIcon" :size="20" />
                        </div>
                    </template>
                </vxe-column>
                <vxe-column field="paths" title="路径" min-width="100"></vxe-column>
                <vxe-column field="permsArr" title="权限标识" min-width="120">
                    <template #default="{ row }">
                        <span v-if="row.perms" type="info">{{ row.perms }}</span>
                    </template>
                </vxe-column>
                <vxe-column field="isDisable" title="状态" width="80">
                    <template #default="{ row }">
                        <el-tag v-if="row.isDisable == 0" type="primary">正常</el-tag>
                        <el-tag v-else type="danger">停用</el-tag>
                    </template>
                </vxe-column>
                <vxe-column field="menuSort" title="排序" width="60"></vxe-column>
                <vxe-column title="操作" width="160">
                    <template #default="{ row }">
                        <el-button
                            v-perms="['admin:system:menu:add']"
                            v-show="row.menuType != MenuEnum.BUTTON"
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
        <!-- <el-table
                v-loading="loading"
                ref="tableRef"
                class="mt-4"
                size="large"
                :data="lists"
                :lazy="true"
                row-key="id"
                :tree-props="{ children: 'children', hasChildren: 'hasChildren' }"
                height="calc(100vh - 220px)"
            >
                <el-table-column
                    label="菜单名称"
                    prop="menuName"
                    min-width="200"
                    show-overflow-tooltip
                ></el-table-column>
                <el-table-column label="类型" prop="menuType" width="60">
                    <template #default="{ row }">
                        <div v-if="row.menuType == MenuEnum.CATALOGUE">目录</div>
                        <div v-else-if="row.menuType == MenuEnum.MENU">菜单</div>
                        <div v-else-if="row.menuType == MenuEnum.BUTTON">按钮</div>
                    </template>
                </el-table-column>
                <el-table-column label="图标" prop="menuIcon" width="60">
                    <template #default="{ row }">
                        <div class="flex">
                            <icon :name="row.menuIcon" :size="20" />
                        </div>
                    </template>
                </el-table-column>

                <el-table-column label="路径" prop="paths" width="200" />
                <el-table-column label="权限标识" prop="permsArr" width="220">
                    <template #default="{ row }">
                        <el-tag v-if="row.perms" type="info">{{ row.perms }}</el-tag>
                    </template>
                </el-table-column>
                <el-table-column label="状态" prop="isDisable" width="100">
                    <template #default="{ row }">
                        <el-tag v-if="row.isDisable == 0" type="primary">正常</el-tag>
                        <el-tag v-else type="danger">停用</el-tag>
                    </template>
                </el-table-column>
                <el-table-column label="排序" prop="menuSort" width="80" />

                <el-table-column label="操作" width="160">
                    <template #default="{ row }">
                        <el-button
                            v-perms="['admin:system:menu:add']"
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
                </el-table-column>
            </el-table> -->
    </div>
    <EditPopup v-if="showEdit" ref="editRef" @success="getLists" @close="showEdit = false" />
</template>
<script lang="ts" setup>
import { ref, useTemplateRef, nextTick, computed } from 'vue'
import { menuDelete, menuLists, SystemAuthMenuResp } from '@/api/perms/menu'
// import { arrayToTree } from '@/utils/util'
import { MenuEnum } from '@/enums/appEnums'
import EditPopup from './edit.vue'
import feedback from '@/utils/feedback'

import { VxeTableInstance } from 'vxe-table'

defineOptions({
    name: 'MenuView'
})
const rowConfig = {
    keyField: 'id'
}
const treeConfig = {
    rowField: 'id',
    childrenField: 'children',
    indent: 10,
    reserve: true,
    lazy: true,
    transform: true,

    parentField: 'pid'
}
const tableRef = useTemplateRef<VxeTableInstance<any>>('tableRef')
const editRef = useTemplateRef<InstanceType<typeof EditPopup>>('editRef')

const loading = ref(false)
const showEdit = ref(false)
const lists = ref<SystemAuthMenuResp[]>([])
const menuName = ref('')
// const filterList = computed(() => {
//     if (!menuName.value) {
//         return lists.value
//     }
//     return lists.value.filter((item) => item.menuName.includes(menuName.value))
// })
const getLists = async () => {
    loading.value = true
    try {
        const data = await menuLists()
        lists.value = data
        // lists.value = arrayToTree(data)
        //  .map((item: any) => {
        // return item
        // })
        loading.value = false
    } catch (error) {
        loading.value = false
    }
}

const handleAdd = async (id?: number) => {
    showEdit.value = true
    await nextTick()
    if (id) {
        editRef.value?.setFormData({
            pid: id
        })
    }
    editRef.value?.open('add')
}

const handleEdit = async (data: any) => {
    showEdit.value = true
    await nextTick()
    editRef.value?.open('edit')
    editRef.value?.getDetail(data)
}

const handleDelete = async (id: number) => {
    await feedback.confirm('确定要删除？')
    await menuDelete({ id })
    feedback.msgSuccess('删除成功')
    getLists()
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
