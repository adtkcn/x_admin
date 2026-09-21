<template>
    <div class="dict-type">
        <el-card class="border-none!" shadow="never">
            <el-form ref="formRef" class="mb-[-16px]" :model="queryParams" inline>
                <el-form-item class="w-[280px]" label="字典名称">
                    <el-input v-model="queryParams.dict_name" clearable @keyup.enter="resetPage" />
                </el-form-item>
                <el-form-item class="w-[280px]" label="字典类型">
                    <el-input v-model="queryParams.dict_type" clearable @keyup.enter="resetPage" />
                </el-form-item>
                <el-form-item class="w-[280px]" label="状态">
                    <el-select v-model="queryParams.dict_status">
                        <el-option label="正常" :value="1" />
                        <el-option label="停用" :value="0" />
                    </el-select>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="resetPage">查询</el-button>
                    <el-button @click="resetParams">重置</el-button>
                </el-form-item>
            </el-form>
        </el-card>

        <el-card class="border-none! mt-4" shadow="never">
            <div>
                <el-button
                    v-perms="['admin:setting:dict:type:add']"
                    type="primary"
                    @click="handleAdd"
                >
                    <template #icon>
                        <icon name="el-icon-Plus" />
                    </template>
                    新增
                </el-button>
                <el-button
                    v-perms="['admin:setting:dict:type:list']"
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
            <div class="mt-4" v-loading="pager.loading">
                <div>
                    <vxe-table
                        ref="tableRef"
                        :data="pager.lists"
                        :row-config="{ keyField: 'id' }"
                        :checkbox-config="{ checkRowKeys: [] }"
                        @checkbox-change="selectData = getCheckedIds()"
                        @checkbox-all="selectData = getCheckedIds()"
                        :border="'inner'"
                    >
                        <vxe-column type="checkbox" width="55" />
                        <!-- <vxe-column title="ID" field="id" width="100" /> -->
                        <vxe-column title="字典名称" field="dict_name" />
                        <vxe-column title="字典类型" field="dict_type" />
                        <vxe-column title="状态">
                            <template v-slot="{ row }">
                                <el-tag v-if="row.dict_status == 1" type="primary">正常</el-tag>
                                <el-tag v-else type="danger">停用</el-tag>
                            </template>
                        </vxe-column>
                        <vxe-column title="备注" field="dictRemark" show-overflow="title" />
                        <vxe-column title="创建时间" field="create_time" />
                        <vxe-column title="操作" width="190" fixed="right">
                            <template #default="{ row }">
                                <el-button
                                    v-perms="['admin:setting:dict:type:edit']"
                                    link
                                    type="primary"
                                    @click="handleEdit(row)"
                                >
                                    编辑
                                </el-button>
                                <el-button
                                    v-perms="['admin:setting:dict:data:list']"
                                    type="primary"
                                    link
                                    @click="openDataEdit(row)"
                                >
                                    数据管理
                                </el-button>
                                <el-button
                                    v-perms="['admin:setting:dict:type:del']"
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
                <div class="flex justify-end mt-4">
                    <pagination v-model="pager" @change="getLists" />
                </div>
            </div>
        </el-card>

        <edit-popup v-if="showEdit" ref="editRef" @success="getLists" @close="showEdit = false" />
        <Data ref="dataRef" @success="getLists" @close="showDataEdit = false" />
    </div>
</template>

<script lang="ts" setup>
import { ref, shallowRef, reactive, nextTick } from 'vue'
import {
    dictTypeDelete,
    dictTypeLists,
    type type_setting_dict_type_list,
    type type_setting_dict_type_resp,
    type type_setting_dict_type_del
} from '@/api/setting/dict'
import { usePaging } from '@/hooks/usePaging'
import feedback from '@/utils/feedback'
import EditPopup from './edit.vue'

import Data from '../data/index.vue'

defineOptions({
    name: 'dictType'
})

const editRef = shallowRef<InstanceType<typeof EditPopup>>()
const showEdit = ref(false)

const dataRef = shallowRef<InstanceType<typeof Data>>()
const showDataEdit = ref(false)
const tableRef = ref<any>()
const queryParams = reactive<type_setting_dict_type_list>({
    dict_name: '',
    dict_type: '',
    dict_status: 1
})

const { pager, getLists, resetPage, resetParams } = usePaging({
    fetchFun: dictTypeLists,
    params: queryParams
})

const selectData = ref<string[]>([])
function openDataEdit(row: type_setting_dict_type_resp) {
    dataRef.value?.open(row)
}
const getCheckedIds = () => {
    return (tableRef.value?.getCheckboxRecords() ?? []).map((item: any) => item.id)
}

const handleAdd = async () => {
    showEdit.value = true
    await nextTick()
    editRef.value?.open('add')
}

const handleEdit = async (data: type_setting_dict_type_resp) => {
    showEdit.value = true
    await nextTick()
    editRef.value?.open('edit')
    editRef.value?.setFormData(data)
}

// 删除角色
const handleDelete = async (ids: string[]) => {
    try {
        await feedback.confirm('确定要删除？')
        await dictTypeDelete({ ids })
        feedback.msgSuccess('删除成功')
        getLists()
    } catch (error) {
        console.error('字典类型删除失败:', error)
    }
}

getLists()
</script>
