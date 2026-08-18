<template>
    <div class="dict-type">
        <popup
            ref="popupRef"
            :title="typeTitle"
            :confirmButtonText="false"
            :async="true"
            width="850px"
        >
            <div>
                <el-button
                    v-perms="['admin:setting:dict:data:add']"
                    type="primary"
                    @click="handleAdd"
                >
                    <template #icon>
                        <icon name="el-icon-Plus" />
                    </template>
                    添加数据
                </el-button>
                <el-button
                    v-perms="['admin:setting:dict:data:del']"
                    :disabled="!selectData.length"
                    type="danger"
                    @click="handleDelete(selectData)"
                >
                    <template #icon>
                        <icon name="el-icon-Delete" />
                    </template>
                    删除
                </el-button>
            </div>
            <div class="mt-4">
                <div>
                    <vxe-table
                        ref="tableRef"
                        :data="lists"
                        :row-config="{ keyField: 'id' }"
                        :checkbox-config="{ checkRowKeys: [] }"
                        @checkbox-change="selectData = getCheckedIds()"
                        @checkbox-all="selectData = getCheckedIds()"
                        :border="'inner'"
                    >
                        <vxe-column type="checkbox" width="55" />
                        <!-- <vxe-column title="ID" field="id" /> -->
                        <vxe-column title="数据名称" field="name" min-width="120">
                            <template v-slot="{ row }">
                                <span :style="{ color: row.color }">{{ row.name }}</span>
                            </template>
                        </vxe-column>
                        <vxe-column title="数据值" field="value" min-width="120" />
                        <!-- <vxe-column title="颜色" field="color" min-width="120" /> -->
                        <vxe-column title="状态">
                            <template v-slot="{ row }">
                                <el-tag v-if="row.status == 1" type="primary">正常</el-tag>
                                <el-tag v-else type="danger">停用</el-tag>
                            </template>
                        </vxe-column>
                        <vxe-column title="备注" field="remark" min-width="120" show-overflow />
                        <vxe-column title="排序" field="sort" />
                        <vxe-column title="操作" width="120" fixed="right">
                            <template #default="{ row }">
                                <el-button
                                    v-perms="['admin:setting:dict:data:edit']"
                                    link
                                    type="primary"
                                    @click="handleEdit(row)"
                                >
                                    编辑
                                </el-button>
                                <el-button
                                    v-perms="['admin:setting:dict:data:del']"
                                    link
                                    type="danger"
                                    @click="handleDelete([row.id])"
                                >
                                    删除
                                </el-button>
                            </template>
                        </vxe-column>
                    </vxe-table>
                </div>
            </div>
        </popup>
        <edit-popup v-if="showEdit" ref="editRef" @success="getLists" @close="showEdit = false" />
    </div>
</template>

<script lang="ts" setup>
import { computed, ref, shallowRef, nextTick } from 'vue'

import Popup from '@/components/popup/index.vue'
import { dictDataDelete, dictDataAll } from '@/api/setting/dict'
import { type_setting_dict_data_resp } from '@/api/setting/dict'
// import { useDictOptions } from '@/hooks/useDictOptions'

import feedback from '@/utils/feedback'
import EditPopup from './edit.vue'

defineOptions({
    name: 'dictData'
})

const popupRef = shallowRef<InstanceType<typeof Popup>>()
const showEdit = ref(false)
const editRef = shallowRef<InstanceType<typeof EditPopup>>()
const tableRef = ref<any>()

const selectRow = ref<any>()
const typeTitle = computed(() => {
    return `${selectRow.value?.dict_name} [ ${selectRow.value?.dict_type} ]`
})

const lists = ref<type_setting_dict_data_resp[]>([])
function getLists() {
    lists.value = []
    dictDataAll({
        dict_type: selectRow.value?.dict_type
    }).then((res) => {
        console.log(res)
        lists.value = res
    })
}

const open = (row: any) => {
    selectRow.value = row

    getLists()
    popupRef.value?.open()
}

const selectData = ref<string[]>([])

const getCheckedIds = () => {
    return (tableRef.value?.getCheckboxRecords() ?? []).map((item: any) => item.id)
}

const handleAdd = async () => {
    showEdit.value = true
    await nextTick()
    editRef.value?.setFormData({
        type_id: selectRow.value?.id
    })
    editRef.value?.open('add')
}

const handleEdit = async (data: type_setting_dict_data_resp) => {
    showEdit.value = true
    await nextTick()
    editRef.value?.open('edit')
    editRef.value?.setFormData({ ...data, type_id: selectRow.value?.id })
}

const handleDelete = async (ids: string[]) => {
    try {
        await feedback.confirm('确定要删除？')
        await dictDataDelete({ ids })
        feedback.msgSuccess('删除成功')
        getLists()
    } catch (error) {
        console.error('字典数据删除失败:', error)
    }
}

defineExpose({
    open
})
</script>
