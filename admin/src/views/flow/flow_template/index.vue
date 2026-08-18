<template>
    <div class="index-lists">
        <el-card class="border-none!" shadow="never">
            <el-form ref="formRef" class="mb-[-16px]" :model="queryParams" :inline="true">
                <el-form-item class="w-[280px]" label="流程名称" prop="flow_name">
                    <el-input v-model="queryParams.flow_name" />
                </el-form-item>
                <el-form-item class="w-[280px]" label="流程分类" prop="flow_group">
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
                <el-form-item class="w-[280px]" label="流程描述" prop="flow_remark">
                    <el-input v-model="queryParams.flow_remark" />
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
            <vxe-table
                class="mt-4"
                v-loading="pager.loading"
                :data="pager.lists"
                :row-config="{ keyField: 'id' }"
                :scroll-y="{ enabled: false }"
                :border="'inner'"
            >
                <vxe-column title="流程名称" field="flow_name" min-width="100" />
                <vxe-column title="流程分类" field="flow_group" min-width="100">
                    <template #default="{ row }">
                        <dict-value :options="dictData.flow_group" :value="row.flow_group" />
                    </template>
                </vxe-column>
                <vxe-column title="流程描述" field="flow_remark" min-width="100" />
                <!-- <vxe-column title="表单配置" field="flow_form_data" min-width="100" />
                <vxe-column title="流程配置" field="flow_process_data" min-width="100" /> -->
                <vxe-column title="操作" fixed="right" width="140">
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
                </vxe-column>
            </vxe-table>
            <div class="flex justify-end mt-4">
                <pagination v-model="pager" @change="getLists" />
            </div>
        </el-card>

        <!-- <edit-popup v-if="showEdit" ref="editRef" @success="getLists" @close="showEdit = false" /> -->
        <Approver ref="approverRef" :save="save"></Approver>
    </div>
</template>
<script lang="ts" setup>
import { shallowRef, reactive, defineAsyncComponent } from 'vue'
import {
    flow_template_delete,
    flow_template_lists,
    flow_template_edit,
    flow_template_add
} from '@/api/flow/flow_template'
import type { type_flow_template } from '@/api/flow/flow_template'
// 仅用于类型标注，运行时仍走下方的异步加载
import type ApproverComponent from '@/components/flow/Approver.vue'
import { useDictData } from '@/hooks/useDictOptions'
import { usePaging } from '@/hooks/usePaging'
import feedback from '@/utils/feedback'
const Approver = defineAsyncComponent(() => import('@/components/flow/Approver.vue'))

defineOptions({
    name: 'flow_template'
})

/** 字典项 */
interface DictItem {
    name: string
    value: number
}

/** 流程模板列表项（复用 API 类型） */
type FlowTemplateItem = type_flow_template

/** Approver 发布回调传出的数据 */
interface FlowSaveInfo {
    id?: string
    basicSetting: {
        flow_name?: string
        flow_group?: number
        flow_remark?: string
    }
    /** 表单设计器字段列表 */
    flow_form_data: any[]
    /** 流程设计器节点树 */
    flow_process_data: Record<string, any>
    /** 流程节点拍平列表 */
    flow_process_data_list?: any[]
}
// const editRef = shallowRef<InstanceType<typeof EditPopup>>()
// const showEdit = ref(false)
const queryParams = reactive({
    flow_name: '',
    flow_group: '',
    flow_remark: ''
})

const { pager, getLists, resetPage, resetParams } = usePaging<FlowTemplateItem>({
    fetchFun: flow_template_lists,
    params: queryParams
})
const { dictData } = useDictData<{
    flow_group: DictItem[]
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

const handleDelete = async (id: string) => {
    await feedback.confirm('确定要删除？')
    await flow_template_delete(id)
    feedback.msgSuccess('删除成功')
    getLists()
}
function save(info: FlowSaveInfo): Promise<boolean> {
    return new Promise<boolean>((resolve, reject) => {
        if (info.id) {
            flow_template_edit({
                id: info.id,
                flow_name: info.basicSetting.flow_name,
                flow_group: info.basicSetting.flow_group,
                flow_remark: info.basicSetting.flow_remark,
                flow_form_data: JSON.stringify(info.flow_form_data),
                flow_process_data: JSON.stringify(info.flow_process_data),
                flow_process_data_list: JSON.stringify(info.flow_process_data_list)
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
                flow_name: info.basicSetting.flow_name,
                flow_group: info.basicSetting.flow_group,
                flow_remark: info.basicSetting.flow_remark,
                flow_form_data: JSON.stringify(info.flow_form_data),
                flow_process_data: JSON.stringify(info.flow_process_data),
                flow_process_data_list: JSON.stringify(info.flow_process_data_list)
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
const approverRef = shallowRef<InstanceType<typeof ApproverComponent>>()
const handleConfig = async (data: FlowTemplateItem) => {
    approverRef.value?.open({
        id: data.id,
        basicSetting: {
            flow_name: data.flow_name,
            flow_group: data.flow_group,
            flow_remark: data.flow_remark
        },

        flow_form_data: JSON.parse(data.flow_form_data || '[]'),
        flow_process_data: JSON.parse(data.flow_process_data || '{}')
    })
}
getLists()
</script>
