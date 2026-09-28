<template>
    <div class="device-lists">
        <el-card class="border-none!" shadow="never">
            <el-form ref="formRef" class="mb-[-16px]" :model="queryParams" :inline="true">
                <el-form-item class="w-[220px]" label="设备名称">
                    <el-input v-model="queryParams.name" clearable @keyup.enter="resetPage" />
                </el-form-item>
                <el-form-item class="w-[220px]" label="设备编号">
                    <el-input v-model="queryParams.sn" clearable @keyup.enter="resetPage" />
                </el-form-item>
                <el-form-item class="w-[220px]" label="所属产品">
                    <el-select v-model="queryParams.product_id" clearable placeholder="请选择产品">
                        <el-option v-for="p in products" :key="p.id" :label="p.name" :value="(p.id as string)" />
                    </el-select>
                </el-form-item>
                <el-form-item class="w-[220px]" label="所属区域">
                    <el-tree-select
                        v-model="queryParams.space_id"
                        :data="spaceTreeOptions"
                        clearable
                        node-key="id"
                        :props="{ label: 'name' }"
                        check-strictly
                        placeholder="请选择区域"
                    />
                </el-form-item>
                <el-form-item class="w-[180px]" label="在线状态">
                    <el-select v-model="queryParams.status" clearable placeholder="全部">
                        <el-option label="在线" :value="1" />
                        <el-option label="离线" :value="0" />
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
                <el-button v-perms="['admin:iot:device:add']" type="primary" @click="handleAdd()">
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
                <vxe-column title="设备名称" field="name" min-width="140" />
                <vxe-column title="设备编号" field="sn" min-width="140" />
                <vxe-column title="所属产品" field="product_name" min-width="120" />
                <vxe-column title="所属区域" field="space_name" min-width="120" />
                <vxe-column title="网关设备" field="gateway_name" min-width="120">
                    <template #default="{ row }">
                        {{ row.gateway_name || '-' }}
                    </template>
                </vxe-column>
                <vxe-column title="在线状态" field="status" min-width="90">
                    <template #default="{ row }">
                        <el-tag class="ml-2" :type="row.status ? 'success' : 'info'">
                            {{ row.status ? '在线' : '离线' }}
                        </el-tag>
                    </template>
                </vxe-column>
                <vxe-column title="最后上线时间" field="last_online_time" min-width="180">
                    <template #default="{ row }">
                        {{ row.last_online_time || '-' }}
                    </template>
                </vxe-column>
                <vxe-column title="创建时间" field="create_time" min-width="180" />
                <vxe-column title="操作" width="120" fixed="right">
                    <template #default="{ row }">
                        <el-button
                            v-perms="['admin:iot:device:edit']"
                            type="primary"
                            link
                            @click="handleEdit(row)"
                        >
                            编辑
                        </el-button>
                        <el-button
                            v-perms="['admin:iot:device:del']"
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
    deviceLists,
    deviceDel,
    productAll,
    deviceSpaceAll,
    type type_iot_device_list,
    type type_iot_device,
    type type_iot_product
} from '@/api/iot'
import { usePaging } from '@/hooks/usePaging'
import feedback from '@/utils/feedback'
import { arrayToTree } from '@/utils/util'
import EditPopup from './edit.vue'

defineOptions({
    name: 'iotDevice'
})

const editRef = shallowRef<InstanceType<typeof EditPopup>>()
const showEdit = ref(false)
const queryParams = reactive<type_iot_device_list>({
    name: '',
    sn: '',
    product_id: '',
    space_id: '',
    status: -1
})

const { pager, getLists, resetPage, resetParams } = usePaging({
    fetchFun: deviceLists,
    params: queryParams
})

const products = ref<type_iot_product[]>([])
const spaceTreeOptions = ref<any[]>([])
onMounted(async () => {
    products.value = await productAll()
    const spaces = await deviceSpaceAll()
    spaceTreeOptions.value = arrayToTree<any>(spaces as any[], '')
})

const handleAdd = async () => {
    showEdit.value = true
    await nextTick()
    editRef.value?.open('add')
}

const handleEdit = async (data: type_iot_device) => {
    showEdit.value = true
    await nextTick()
    editRef.value?.open('edit')
    editRef.value?.getDetail(data)
}

const handleDelete = async (id: string) => {
    try {
        await feedback.confirm('确定要删除？')
        await deviceDel({ id })
        feedback.msgSuccess('删除成功')
        getLists()
    } catch (error) {
        console.error('设备删除失败:', error)
    }
}

getLists()
</script>
