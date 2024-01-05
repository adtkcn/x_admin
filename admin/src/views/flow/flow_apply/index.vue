<template>
    <div class="index-lists">
        <el-card class="!border-none" shadow="never">
            <el-form ref="formRef" class="mb-[-16px]" :model="queryParams" :inline="true">
                <!-- <el-form-item label="模板" prop="templateId">
                    <el-input v-model="queryParams.templateId" />
                </el-form-item> -->
                <!-- <el-form-item label="申请人id" prop="applyUserId">
                    <el-input   v-model="queryParams.applyUserId" />
                </el-form-item>
                <el-form-item label="申请人昵称" prop="applyUserNickname">
                    <el-input   v-model="queryParams.applyUserNickname" />
                </el-form-item> -->
                <el-form-item label="流程名称" prop="flowName">
                    <el-input v-model="queryParams.flowName" />
                </el-form-item>
                <el-form-item label="流程分类" prop="flowGroup">
                    <el-select v-model="queryParams.flowGroup" clearable>
                        <el-option label="全部" value="" />
                        <el-option
                            v-for="(item, index) in dictData.flow_group"
                            :key="index"
                            :label="item.name"
                            :value="item.value"
                        />
                    </el-select>
                </el-form-item>
                <!-- <el-form-item label="流程描述" prop="flowRemark">
                    <el-input v-model="queryParams.flowRemark" />
                </el-form-item> -->
                <el-form-item label="状态" prop="status">
                    <el-select v-model="queryParams.status" clearable>
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
        <el-card class="!border-none mt-4" shadow="never">
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
            <el-table class="mt-4" size="large" v-loading="pager.loading" :data="pager.lists">
                <el-table-column label="申请人昵称" prop="applyUserNickname" min-width="100" />
                <el-table-column label="流程名称" prop="flowName" min-width="100" />
                <el-table-column label="流程分类" prop="flowGroup" min-width="100">
                    <template #default="{ row }">
                        <dict-value :options="dictData.flow_group" :value="row.flowGroup" />
                    </template>
                </el-table-column>
                <el-table-column label="流程描述" prop="flowRemark" min-width="100" />
                <!-- <el-table-column label="formValue" prop="formValue" min-width="100" /> -->

                <el-table-column label="状态" prop="status" min-width="100">
                    <template #default="{ row }">
                        <dict-value :options="dictData.flow_apply_status" :value="row.status" />
                    </template>
                </el-table-column>
                <el-table-column label="更新时间" prop="updateTime" min-width="130" />
                <el-table-column label="创建时间" prop="createTime" min-width="130" />
                <el-table-column label="操作" width="180" fixed="right">
                    <template #default="{ row }">
                        <el-button
                            v-perms="['admin:flow:flow_apply:edit']"
                            type="primary"
                            link
                            @click="OpenViewForm(row)"
                        >
                            {{ row.status == 1 ? '编辑' : '预览' }}
                        </el-button>
                        <el-button
                            v-if="row.status == 1 && row.formValue"
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
                </el-table-column>
            </el-table>
            <div class="flex justify-end mt-4">
                <pagination v-model="pager" @change="getLists" />
            </div>
        </el-card>
        <edit-popup
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
    flow_apply_delete,
    flow_apply_lists,
    flow_apply_edit,
    flow_apply_detail
} from '@/api/flow/flow_apply'

import { useDictData } from '@/hooks/useDictOptions'
import { usePaging } from '@/hooks/usePaging'
import feedback from '@/utils/feedback'
import EditPopup from './edit.vue'

import ApplySubmit from './components/apply_submit.vue'
import ViewForm from './components/ViewForm.vue'

import useUserStore from '@/stores/modules/user'

const userStore = useUserStore()

defineOptions({
    name: 'flow_apply'
})
const viewFormRef = shallowRef<InstanceType<typeof ViewForm>>()
const ApplySubmitRef = shallowRef<InstanceType<typeof ApplySubmit>>()
const editRef = shallowRef<InstanceType<typeof EditPopup>>()
const showEdit = ref(false)
const queryParams = reactive({
    templateId: '',
    applyUserId: userStore.userInfo?.id,
    applyUserNickname: '',
    flowName: '',
    flowGroup: '',
    flowRemark: '',
    flowFormData: '',
    flowProcessData: '',
    status: ''
})

const { pager, getLists, resetPage, resetParams } = usePaging({
    fetchFun: flow_apply_lists,
    params: queryParams
})
const { dictData } = useDictData<{
    flow_apply_status: any[]
    flow_group: any[]
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

const handleDelete = async (id: number) => {
    await feedback.confirm('确定要删除？')
    await flow_apply_delete({ id })
    feedback.msgSuccess('删除成功')
    getLists()
}

const OpenViewForm = async (row: any) => {
    const detail = await flow_apply_detail({ id: row.id })
    let form_data = {}
    try {
        form_data = JSON.parse(row.formValue)
    } catch (error) {
        // 解析失败
    }
    let form_json = {}
    try {
        form_json = JSON.parse(detail.flowFormData)
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
const SaveViewForm = (id, form_data) => {
    return new Promise((resolve, reject) => {
        flow_apply_edit({
            id: id,
            formValue: JSON.stringify(form_data)
        })
            .then(async () => {
                feedback.msgSuccess('保存成功')
                await getLists()

                const row = pager.lists.find((item) => item.id === id)

                ApplySubmitRef.value?.open(row.id)

                resolve(true)
            })
            .catch((err) => {
                feedback.msgError(err.message)
                reject()
            })
    })
}
getLists()
</script>
