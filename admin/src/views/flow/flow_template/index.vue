<template>
    <div class="index-lists">
        <el-card class="!border-none" shadow="never">
            <el-form ref="formRef" class="mb-[-16px]" :model="queryParams" :inline="true">
                <el-form-item class="w-[280px]" label="流程名称" prop="flowName">
                    <el-input v-model="queryParams.flowName" />
                </el-form-item>
                <el-form-item class="w-[280px]" label="流程分类" prop="flowGroup">
                    <el-input v-model="queryParams.flowGroup" />
                </el-form-item>
                <el-form-item class="w-[280px]" label="流程描述" prop="flowRemark">
                    <el-input v-model="queryParams.flowRemark" />
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="resetPage">查询</el-button>
                    <el-button @click="resetParams">重置</el-button>
                </el-form-item>
            </el-form>
        </el-card>
        <el-card class="!border-none mt-4" shadow="never">
            <div>
                <el-button
                    v-perms="['admin:flow_template:add']"
                    type="primary"
                    @click="handleAdd()"
                >
                    <template #icon>
                        <icon name="el-icon-Plus" />
                    </template>
                    新增
                </el-button>
            </div>
            <el-table class="mt-4" size="large" v-loading="pager.loading" :data="pager.lists">
                <el-table-column label="流程名称" prop="flowName" min-width="100" />
                <el-table-column label="流程分类" prop="flowGroup" min-width="100">
                    <template #default="{ row }">
                        <dict-value :options="dictData.flow_group" :value="row.flowGroup" />
                    </template>
                </el-table-column>
                <el-table-column label="流程描述" prop="flowRemark" min-width="100" />
                <!-- <el-table-column label="表单配置" prop="flowFormData" min-width="100" />
                <el-table-column label="流程配置" prop="flowProcessData" min-width="100" /> -->
                <el-table-column label="操作" fixed="right">
                    <template #default="{ row }">
                        <el-button
                            v-perms="['admin:flow_template:edit']"
                            type="primary"
                            link
                            @click="handleConfig(row)"
                        >
                            配置
                        </el-button>
                        <!-- <el-button
                            v-perms="['admin:flow_template:edit']"
                            type="primary"
                            link
                            @click="handleEdit(row)"
                        >
                            编辑
                        </el-button> -->
                        <el-button
                            v-perms="['admin:flow_template:del']"
                            type="danger"
                            link
                            @click="handleDelete(row.id)"
                        >
                            删除
                        </el-button>
                    </template>
                </el-table-column>
            </el-table>
            <div class="flex justify-end mt-4">
                <pagination v-model="pager" @change="getLists" />
            </div>
        </el-card>

        <edit-popup v-if="showEdit" ref="editRef" @success="getLists" @close="showEdit = false" />
        <Approver ref="approverRef" :save="save"></Approver>
    </div>
</template>
<script lang="ts" setup>
import {
    flow_template_delete,
    flow_template_lists,
    flow_template_edit,
    flow_template_add
} from '@/api/flow/flow_template'
import { useDictData } from '@/hooks/useDictOptions'
import { usePaging } from '@/hooks/usePaging'
import feedback from '@/utils/feedback'
import EditPopup from './edit.vue'
import Approver from '@/components/flow/Approver.vue'

defineOptions({
    name: 'flow_template'
})
const editRef = shallowRef<InstanceType<typeof EditPopup>>()
const showEdit = ref(false)
const queryParams = reactive({
    flowName: '',
    flowGroup: '',
    flowRemark: '',
    flowFormData: '',
    flowProcessData: ''
})

const { pager, getLists, resetPage, resetParams } = usePaging({
    fetchFun: flow_template_lists,
    params: queryParams
})
const { dictData } = useDictData<{
    flow_group: any[]
}>(['flow_group'])

const handleAdd = async () => {
    // showEdit.value = true
    // await nextTick()
    // editRef.value?.open('add')

    approverRef.value?.open()
}

// const handleEdit = async (data: any) => {
//     showEdit.value = true
//     await nextTick()
//     editRef.value?.open('edit')
//     editRef.value?.getDetail(data)
// }

const handleDelete = async (id: number) => {
    await feedback.confirm('确定要删除？')
    await flow_template_delete({ id })
    feedback.msgSuccess('删除成功')
    getLists()
}
function save(info) {
    return new Promise((resolve, reject) => {
        if (info.id) {
            flow_template_edit({
                id: info.id,
                flowName: info.basicSetting.flowName,
                flowGroup: info.basicSetting.flowGroup,
                flowRemark: info.basicSetting.flowRemark,
                flowFormData: JSON.stringify(info.flowFormData),
                flowProcessData: JSON.stringify(info.flowProcessData),
                flowProcessDataList: JSON.stringify(info.flowProcessDataList)
            })
                .then(() => {
                    feedback.msgSuccess('修改成功')
                    getLists()
                    resolve(true)
                })
                .catch((err) => {
                    feedback.msgError(err.message)
                    reject()
                })
        } else {
            flow_template_add({
                flowName: info.flowName,
                flowGroup: info.flowGroup,
                flowRemark: info.flowRemark,
                flowFormData: JSON.stringify(info.flowFormData),
                flowProcessData: JSON.stringify(info.flowProcessData)
            })
                .then(() => {
                    feedback.msgSuccess('新增成功')
                    getLists()
                    resolve(true)
                })
                .catch((err) => {
                    feedback.msgError(err.message)
                    reject()
                })
        }
    })
}
const approverRef = shallowRef<InstanceType<typeof EditPopup>>()
const handleConfig = async (data: any) => {
    console.log('toRaw(data)', toRaw(data))

    approverRef.value?.open({
        id: data.id,
        basicSetting: {
            flowName: data.flowName,
            flowGroup: data.flowGroup,
            flowRemark: data.flowRemark
        },

        flowFormData: JSON.parse(data.flowFormData),
        flowProcessData: JSON.parse(data.flowProcessData)
    } as any)
}
getLists()
</script>
