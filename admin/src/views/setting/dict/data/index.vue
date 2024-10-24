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
                    <el-table :data="lists" size="large" @selection-change="handleSelectionChange">
                        <el-table-column type="selection" width="55" />
                        <!-- <el-table-column label="ID" prop="id" /> -->
                        <el-table-column label="数据名称" prop="name" min-width="120">
                            <template v-slot="{ row }">
                                <span :style="{ color: row.color }">{{ row.name }}</span>
                            </template>
                        </el-table-column>
                        <el-table-column label="数据值" prop="value" min-width="120" />
                        <!-- <el-table-column label="颜色" prop="color" min-width="120" /> -->
                        <el-table-column label="状态">
                            <template v-slot="{ row }">
                                <el-tag v-if="row.status == 1" type="primary">正常</el-tag>
                                <el-tag v-else type="danger">停用</el-tag>
                            </template>
                        </el-table-column>
                        <el-table-column
                            label="备注"
                            prop="remark"
                            min-width="120"
                            show-tooltip-when-overflow
                        />
                        <el-table-column label="排序" prop="sort" />
                        <el-table-column label="操作" width="120" fixed="right">
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
                        </el-table-column>
                    </el-table>
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
// import { useDictOptions } from '@/hooks/useDictOptions'

import feedback from '@/utils/feedback'
import EditPopup from './edit.vue'

defineOptions({
    name: 'dictData'
})

const popupRef = shallowRef<InstanceType<typeof Popup>>()
const showEdit = ref(false)
const editRef = shallowRef<InstanceType<typeof EditPopup>>()

const selectRow = ref<any>()
const typeTitle = computed(() => {
    return `${selectRow.value?.dictName} [ ${selectRow.value?.dictType} ]`
})

const lists = ref([])
function getLists() {
    lists.value = []
    dictDataAll({
        dictType: selectRow.value?.dictType
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

const selectData = ref<any[]>([])

const handleSelectionChange = (val: any[]) => {
    selectData.value = val.map(({ id }) => id)
}

const handleAdd = async () => {
    showEdit.value = true
    await nextTick()
    editRef.value?.setFormData({
        typeValue: selectRow.value?.dictType,
        typeId: selectRow.value?.id
    })
    editRef.value?.open('add')
}

const handleEdit = async (data: any) => {
    showEdit.value = true
    await nextTick()
    editRef.value?.open('edit')
    editRef.value?.setFormData({ ...data, typeValue: selectRow.value?.dictType })
}

const handleDelete = async (ids: any[] | number) => {
    try {
        await feedback.confirm('确定要删除？')
        await dictDataDelete({ ids })
        feedback.msgSuccess('删除成功')
        getLists()
    } catch (error) {
        console.log(error)
    }
}

defineExpose({
    open
})
</script>
