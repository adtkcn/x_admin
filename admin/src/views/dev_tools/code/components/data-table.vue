<template>
    <div class="data-table">
        <popup
            ref="popupRef"
            :clickModalClose="false"
            title="选择表"
            width="900px"
            :async="true"
            @confirm="handleConfirm"
        >
            <template #trigger>
                <slot></slot>
            </template>
            <el-form class="ls-form" :model="formData" inline>
                <el-form-item label="表名称">
                    <el-input v-model="formData.table_name" clearable @keyup.enter="resetPage" />
                </el-form-item>
                <el-form-item label="表描述">
                    <el-input v-model="formData.table_comment" clearable @keyup.enter="resetPage" />
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="resetPage">查询</el-button>
                    <el-button @click="resetParams">重置</el-button>
                </el-form-item>
            </el-form>
            <div class="m-4" v-loading="pager.loading">
                <vxe-table
                    ref="tableRef"
                    height="400"
                    :data="pager.lists"
                    :row-config="{ keyField: 'table_name' }"
                    :checkbox-config="{ checkRowKeys: [] }"
                    :border="'inner'"
                >
                    <vxe-column type="checkbox" width="55" />
                    <vxe-column title="表名称" field="table_name" min-width="150" />
                    <vxe-column title="表描述" field="table_comment" min-width="160" />
                    <vxe-column title="创建时间" field="create_time" min-width="180" />
                </vxe-table>
            </div>
            <div class="flex justify-end mt-4">
                <pagination v-model="pager" @change="getLists" />
            </div>
        </popup>
    </div>
</template>

<script lang="ts" setup>
import { shallowRef, reactive, watch, ref } from 'vue'
import Popup from '@/components/popup/index.vue'
import Pagination from '@/components/pagination/index.vue'
import { usePaging } from '@/hooks/usePaging'
import { dataTable, selectTable } from '@/api/tools/code'
import feedback from '@/utils/feedback'

const emit = defineEmits<{
    (event: 'success'): void
}>()

const popupRef = shallowRef<InstanceType<typeof Popup>>()
const tableRef = ref<any>()

const formData = reactive({
    table_name: '', // 表名称
    table_comment: '' // 表描述
})

const { pager, getLists, resetParams, resetPage } = usePaging({
    fetchFun: dataTable,
    params: formData,
    size: 10
})

const selectData = ref<any[]>([])

const handleConfirm = async () => {
    const checkedRows = tableRef.value?.getCheckboxRecords() ?? []
    selectData.value = checkedRows.map(({ table_name }: any) => table_name)
    if (!selectData.value.length) return feedback.msgError('请选择数据表')
    await selectTable({
        tables: selectData.value.join()
    })
    feedback.msgSuccess('导入成功')
    popupRef.value?.close()
    emit('success')
}

watch(
    () => popupRef.value?.visible,
    (value) => {
        if (value) getLists()
    }
)
</script>
