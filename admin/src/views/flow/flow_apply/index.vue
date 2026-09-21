<template>
    <div class="index-lists">
        <el-card class="border-none!" shadow="never">
            <el-form
                ref="formRef"
                class="mb-[-16px]"
                label-width="80px"
                :model="queryParams"
                :inline="true"
            >
                <!-- <el-form-item label="模板" prop="templateId">
                    <el-input v-model="queryParams.template_id" />
                </el-form-item> -->
                <!-- <el-form-item label="申请人id" prop="applyUserId">
                    <el-input   v-model="queryParams.apply_user_id" />
                </el-form-item>
                <el-form-item label="申请人昵称" prop="apply_user_nickname">
                    <el-input   v-model="queryParams.apply_user_nickname" />
                </el-form-item> -->
                <el-form-item label="流程名称" prop="flow_name" class="w-[280px]">
                    <el-input v-model="queryParams.flow_name" />
                </el-form-item>
                <el-form-item label="流程分类" prop="flow_group" class="w-[280px]">
                    <el-select
                        v-model="queryParams.flow_group"
                        clearable
                        :empty-values="[null, undefined]"
                    >
                        <el-option label="全部" value="" />
                        <el-option
                            v-for="(item, index) in dictData.flow_group"
                            :key="index"
                            :label="item.name"
                            :value="item.value"
                        />
                    </el-select>
                </el-form-item>
                <!-- <el-form-item label="流程描述" prop="flow_remark">
                    <el-input v-model="queryParams.flow_remark" />
                </el-form-item> -->
                <el-form-item label="状态" prop="status" class="w-[280px]">
                    <el-select
                        v-model="queryParams.status"
                        clearable
                        :empty-values="[null, undefined]"
                    >
                        <el-option label="全部" value="" />
                        <el-option
                            v-for="(item, index) in dictData.flow_apply_status"
                            :key="index"
                            :label="item.name"
                            :value="item.value"
                        />
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
                    v-perms="['admin:flow:flow_apply:add']"
                    type="primary"
                    @click="handleAdd()"
                >
                    <template #icon>
                        <icon name="el-icon-Plus" />
                    </template>
                    申请
                </el-button>
            </div>
            <vxe-table
                class="mt-4"
                v-loading="pager.loading"
                :data="pager.lists"
                :row-config="{ keyField: 'id' }"
                :scroll-y="{ enabled: false }"
                :border="'inner'"
            >
                <vxe-column title="申请人昵称" field="apply_user_nickname" min-width="100" />
                <vxe-column title="流程名称" field="flow_name" min-width="160" />
                <vxe-column title="流程分类" field="flow_group" min-width="100">
                    <template #default="{ row }">
                        <dict-value :options="dictData.flow_group" :value="row.flow_group" />
                    </template>
                </vxe-column>
                <vxe-column title="流程描述" field="flow_remark" min-width="100" />
                <!-- <vxe-column title="formValue" field="formValue" min-width="100" /> -->

                <vxe-column title="状态" field="status" width="100">
                    <template #default="{ row }">
                        <dict-value :options="dictData.flow_apply_status" :value="row.status" />
                    </template>
                </vxe-column>
                <vxe-column title="更新时间" field="update_time" width="180" />
                <vxe-column title="创建时间" field="create_time" width="180" />
                <vxe-column title="操作" width="150" fixed="right">
                    <template #default="{ row }">
                        <el-button
                            v-perms="['admin:flow:flow_apply:edit']"
                            type="primary"
                            link
                            @click="OpenViewForm(row)"
                        >
                            {{ row.status == 1 || row.status == 4 ? '编辑' : '预览' }}
                        </el-button>
                        <el-button
                            v-if="(row.status == 1 || row.status == 4) && row.form_value"
                            v-perms="['admin:flow:flow_apply:edit']"
                            type="primary"
                            link
                            @click="OpenApplySubmit(row)"
                        >
                            提交
                        </el-button>

                        <!-- <el-button
                            v-perms="['admin:flow:flow_apply:edit']"
                            type="primary"
                            link
                            @click="handleEdit(row)"
                        >
                            编辑
                        </el-button> -->
                        <el-button
                            v-perms="['admin:flow:flow_apply:del']"
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
        <EditPopup
            v-if="showEdit"
            ref="editRef"
            :dict-data="dictData"
            @success="getLists"
            @close="showEdit = false"
        />
        <ViewForm ref="viewFormRef" :save="SaveViewForm"></ViewForm>
        <ApplySubmit
            ref="ApplySubmitRef"
            title="提交申请"
            :showRemark="false"
            @close="getLists"
        ></ApplySubmit>
    </div>
</template>
<script lang="ts" setup>
import {
    ref,
    reactive,
    shallowRef,
    nextTick,
    defineAsyncComponent,
    onMounted,
    onActivated
} from 'vue'
import {
    flow_apply_delete,
    flow_apply_lists,
    flow_apply_edit,
    flow_apply_detail
} from '@/api/flow/flow_apply'

import type { type_flow_apply, type_flow_apply_query } from '@/api/flow/flow_apply'

import { useDictData } from '@/hooks/useDictOptions'
import type { type_dict } from '@/hooks/useDictOptions'
import { usePaging } from '@/hooks/usePaging'
import feedback from '@/utils/feedback'
const EditPopup = defineAsyncComponent(() => import('./edit.vue'))

const ApplySubmit = defineAsyncComponent(() => import('./components/apply_submit.vue'))
const ViewForm = defineAsyncComponent(() => import('./components/ViewForm.vue'))

import useUserStore from '@/stores/modules/user'

const userStore = useUserStore()

defineOptions({
    name: 'flow_apply'
})
const viewFormRef = shallowRef<InstanceType<typeof ViewForm>>()
const ApplySubmitRef = shallowRef<InstanceType<typeof ApplySubmit>>()
const editRef = shallowRef<InstanceType<typeof EditPopup>>()
const showEdit = ref(false)
const queryParams = reactive<type_flow_apply_query>({
    template_id: undefined,
    apply_user_id: String(userStore.userInfo?.id),
    apply_user_nickname: undefined,
    flow_name: undefined,
    flow_group: undefined,
    flow_remark: undefined,
    flow_form_data: undefined,
    flow_process_data: undefined,
    status: undefined
})

const { pager, getLists, resetPage, resetParams } = usePaging<type_flow_apply>({
    fetchFun: flow_apply_lists,
    params: queryParams
})
const { dictData } = useDictData<{
    flow_apply_status: type_dict[]
    flow_group: type_dict[]
}>(['flow_apply_status', 'flow_group'])

const handleAdd = async () => {
    showEdit.value = true
    await nextTick()
    editRef.value?.open('add')
}

// const handleEdit = async (data: any) => {
//     showEdit.value = true
//     await nextTick()
//     editRef.value?.open('edit')
//     editRef.value?.getDetail(data)
// }

const handleDelete = async (id: string) => {
    await feedback.confirm('确定要删除？')
    await flow_apply_delete(id)
    feedback.msgSuccess('删除成功')
    getLists()
}

const OpenViewForm = async (row: any) => {
    const detail = await flow_apply_detail({ id: row.id })
    let form_data = {}
    try {
        form_data = JSON.parse(row.form_value)
    } catch (error) {
        // 解析失败
    }
    let form_json = {}
    try {
        detail.flow_form_data && (form_json = JSON.parse(detail.flow_form_data))
    } catch (error) {
        // 解析失败
    }

    console.log(detail, form_data, form_json)
    viewFormRef.value?.open(detail, form_json, form_data)
}
const OpenApplySubmit = async (data: any) => {
    console.log('OpenApplySubmit')

    ApplySubmitRef.value?.open(data.id)
}
const SaveViewForm = (id: string, form_data: Record<string, any>) => {
    return new Promise((resolve, reject) => {
        flow_apply_edit({
            id: id,
            form_value: JSON.stringify(form_data)
        })
            .then(async () => {
                feedback.msgSuccess('保存成功')
                await getLists()

                const row = pager.lists.find((item) => item.id === id)
                if (!row) {
                    feedback.msgError('申请不存在')
                    reject()
                    return
                }
                ApplySubmitRef.value?.open(row.id)

                resolve(true)
            })
            .catch((err) => {
                feedback.msgError(err.message)
                reject()
            })
    })
}
onMounted(() => {
    getLists()
})
onActivated(() => {
    if (!pager.loading) {
        getLists()
    }
})
</script>
