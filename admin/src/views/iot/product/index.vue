<template>
    <div class="product-lists">
        <el-card class="border-none!" shadow="never">
            <el-form ref="formRef" class="mb-[-16px]" :model="queryParams" :inline="true">
                <el-form-item class="w-[240px]" label="产品编码">
                    <el-input v-model="queryParams.code" clearable @keyup.enter="resetPage" />
                </el-form-item>
                <el-form-item class="w-[240px]" label="产品名称">
                    <el-input v-model="queryParams.name" clearable @keyup.enter="resetPage" />
                </el-form-item>
                <el-form-item class="w-[240px]" label="所属分类">
                    <el-tree-select
                        v-model="queryParams.category_id"
                        :data="categoryTreeOptions"
                        clearable
                        node-key="id"
                        :props="{ label: 'name' }"
                        check-strictly
                        placeholder="请选择分类"
                    />
                </el-form-item>
                <el-form-item class="w-[180px]" label="状态">
                    <el-select v-model="queryParams.is_stop" clearable placeholder="全部">
                        <el-option label="正常" :value="0" />
                        <el-option label="停用" :value="1" />
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
                <el-button v-perms="['admin:iot:product:add']" type="primary" @click="handleAdd()">
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
                show-overflow="title"
            >
                <vxe-column title="产品编码" field="code" min-width="120" />
                <vxe-column title="产品名称" field="name" min-width="140" />
                <vxe-column title="所属分类" field="category_name" min-width="120" />
                <vxe-column title="节点类型" field="device_type" min-width="120">
                    <template #default="{ row }">
                        {{ NodeTypeMap[row.device_type as NodeType] || '未设置' }}
                    </template>
                </vxe-column>
                <vxe-column title="接入协议" field="protocol" min-width="90" />
                <vxe-column title="备注" field="remark" min-width="140" show-overflow="title" />
                <vxe-column title="状态" field="is_stop" min-width="90">
                    <template #default="{ row }">
                        <el-tag class="ml-2" :type="row.is_stop ? 'danger' : 'primary'">
                            {{ row.is_stop ? '停用' : '正常' }}
                        </el-tag>
                    </template>
                </vxe-column>
                <vxe-column title="创建时间" field="create_time" min-width="180" />
                <vxe-column title="操作" width="120" fixed="right">
                    <template #default="{ row }">
                        <el-button
                            v-perms="['admin:iot:product:edit']"
                            type="primary"
                            link
                            @click="handleEdit(row)"
                        >
                            编辑
                        </el-button>
                        <el-button
                            v-perms="['admin:iot:product:del']"
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
import { ref, shallowRef, reactive, nextTick, onMounted } from 'vue'
import {
    productLists,
    productDel,
    NodeTypeMap,
    type NodeType,
    type type_iot_product_list,
    type type_iot_product
} from '@/api/iot/product'
import { productCategoryAll } from '@/api/iot/product_category'
import { usePaging } from '@/hooks/usePaging'
import feedback from '@/utils/feedback'
import { arrayToTree } from '@/utils/util'
import EditPopup from './edit.vue'

defineOptions({
    name: 'iotProduct'
})

const editRef = shallowRef<InstanceType<typeof EditPopup>>()
const showEdit = ref(false)
const queryParams = reactive<type_iot_product_list>({
    code: '',
    name: '',
    category_id: '',
    is_stop: undefined
})

const { pager, getLists, resetPage, resetParams } = usePaging({
    fetchFun: productLists,
    params: queryParams
})

const categoryTreeOptions = ref<any[]>([])
onMounted(async () => {
    const list = await productCategoryAll()
    categoryTreeOptions.value = arrayToTree<any>(list as any[], '')
})

const handleAdd = async () => {
    showEdit.value = true
    await nextTick()
    editRef.value?.open('add')
}

const handleEdit = async (data: type_iot_product) => {
    showEdit.value = true
    await nextTick()
    editRef.value?.open('edit')
    editRef.value?.getDetail(data)
}

const handleDelete = async (id: string) => {
    try {
        await feedback.confirm('确定要删除？')
        await productDel({ id })
        feedback.msgSuccess('删除成功')
        getLists()
    } catch (error) {
        console.error('产品删除失败:', error)
    }
}

getLists()
</script>
