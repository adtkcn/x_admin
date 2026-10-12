<template>
    <div class="rc-lists">
        <el-card class="border-none!" shadow="never">
            <el-form class="mb-[-16px]" :model="queryParams" :inline="true">
                <el-form-item class="w-[240px]" label="规则链名称">
                    <el-input v-model="queryParams.name" clearable @keyup.enter="resetPage" />
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
                <el-button v-perms="['admin:iot:ruleChain:add']" type="primary" @click="handleAdd()">
                    <template #icon>
                        <icon name="el-icon-Plus" />
                    </template>
                    新增规则链
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
                <vxe-column title="规则链名称" field="name" min-width="160" />
                <vxe-column title="备注" field="remark" min-width="200" show-overflow="title" />
                <vxe-column title="状态" field="is_stop" min-width="90">
                    <template #default="{ row }">
                        <el-tag :type="row.is_stop ? 'danger' : 'primary'">
                            {{ row.is_stop ? '停用' : '正常' }}
                        </el-tag>
                    </template>
                </vxe-column>
                <vxe-column title="更新时间" field="update_time" min-width="180" />
                <vxe-column title="操作" width="160" fixed="right">
                    <template #default="{ row }">
                        <el-button
                            v-perms="['admin:iot:ruleChain:edit']"
                            type="primary"
                            link
                            @click="handleConfig(row)"
                        >
                            配置
                        </el-button>
                        <el-button
                            v-perms="['admin:iot:ruleChain:del']"
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

        <EditPopup v-if="showEdit" ref="editRef" @success="getLists" @close="showEdit = false" />
    </div>
</template>
<script lang="ts" setup>
import { ref, shallowRef, reactive, nextTick, defineComponent } from 'vue'
import { ruleChainLists, ruleChainDel, type type_rule_chain, type type_rule_chain_list } from '@/api/iot/rule_chain'
import { usePaging } from '@/hooks/usePaging'
import feedback from '@/utils/feedback'

import EditPopup from './edit.vue'

defineOptions({
    name: 'iotRuleChain'
})

const editRef = shallowRef<InstanceType<typeof EditPopup>>()
const showEdit = ref(false)
const queryParams = reactive<type_rule_chain_list>({
    name: '',
    is_stop: undefined
})

const { pager, getLists, resetPage, resetParams } = usePaging<type_rule_chain>({
    fetchFun: ruleChainLists,
    params: queryParams
})

const handleAdd = async () => {
    showEdit.value = true
    await nextTick()
    editRef.value?.open()
}

const handleConfig = async (row: type_rule_chain) => {
    showEdit.value = true
    await nextTick()
    editRef.value?.open(row)
}

const handleDelete = async (id: string) => {
    try {
        await feedback.confirm('确定要删除该规则链？被产品引用的规则链无法删除')
        await ruleChainDel({ id })
        feedback.msgSuccess('删除成功')
        getLists()
    } catch (error) {
        console.error('规则链删除失败:', error)
    }
}

getLists()
</script>
