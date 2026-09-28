<template>
    <div class="edit-popup">
        <popup
            ref="popupRef"
            :title="popupTitle"
            :async="true"
            width="550px"
            @confirm="handleSubmit"
            @close="handleClose"
        >
            <el-form ref="formRef" :model="formData" label-width="100px" :rules="formRules">
                <el-form-item label="上级分类" prop="pid">
                    <el-tree-select
                        class="flex-1"
                        v-model="formData.pid"
                        :data="categoryTreeOptions"
                        clearable
                        node-key="id"
                        :props="{ label: 'name', disabled: 'disabled' }"
                        check-strictly
                        :default-expand-all="true"
                        placeholder="请选择上级分类"
                    />
                </el-form-item>
                <el-form-item label="分类名称" prop="name">
                    <el-input
                        v-model="formData.name"
                        placeholder="请输入分类名称"
                        clearable
                        :maxlength="100"
                    />
                </el-form-item>
                <el-form-item label="排序" prop="sort">
                    <div>
                        <el-input-number v-model="formData.sort" :min="0" :max="9999" />
                        <div class="form-tips">默认为0，数值越大越排前</div>
                    </div>
                </el-form-item>
                <el-form-item label="备注" prop="remark">
                    <el-input v-model="formData.remark" type="textarea" :rows="2" placeholder="请输入备注" clearable :maxlength="255" />
                </el-form-item>
            </el-form>
        </popup>
    </div>
</template>
<script lang="ts" setup>
import { ref, computed, shallowRef } from 'vue'
import type { FormInstance } from 'element-plus'
import {
    productCategoryAll,
    productCategoryAdd,
    productCategoryEdit,
    productCategoryDetail,
    type type_iot_product_category
} from '@/api/iot'
import Popup from '@/components/popup/index.vue'
import feedback from '@/utils/feedback'
import { arrayToTree } from '@/utils/util'
import { useReactiveWithReset } from '@/hooks/useReactiveWithReset'

const emit = defineEmits(['success', 'close'])
const formRef = shallowRef<FormInstance>()
const popupRef = shallowRef<InstanceType<typeof Popup>>()
const mode = ref('add')
const popupTitle = computed(() => (mode.value == 'edit' ? '编辑分类' : '新增分类'))

const { state: formData, setState } = useReactiveWithReset<type_iot_product_category>({
    id: '',
    pid: '',
    name: '',
    sort: 0,
    remark: ''
})

const formRules = {
    name: [{ required: true, message: '请输入分类名称', trigger: ['blur'] }]
}

const flatList = ref<type_iot_product_category[]>([])
async function loadCategoryOptions() {
    flatList.value = await productCategoryAll()
}

// 收集某节点自身及所有后代 id（编辑时禁用，避免成环）
const collectSelfAndDescendants = (list: any[], currentId: string): Set<string> => {
    const ids = new Set<string>()
    const findAndCollect = (nodes: any[]) => {
        for (const n of nodes) {
            if (n.id === currentId || ids.has(n.id)) {
                const collectSub = (sub: any) => {
                    ids.add(sub.id)
                    ;(sub.children || []).forEach(collectSub)
                }
                collectSub(n)
            } else if (n.children?.length) {
                findAndCollect(n.children)
            }
        }
    }
    findAndCollect(list)
    return ids
}
const markDisabled = (list: any[], disabledIds: Set<string>): any[] =>
    list.map((item) => ({
        ...item,
        disabled: disabledIds.has(item.id),
        children: item.children?.length ? markDisabled(item.children, disabledIds) : item.children
    }))

const categoryTreeOptions = computed(() => {
    if (!flatList.value.length) return []
    const tree = arrayToTree<any>(
        flatList.value.map((item) => ({ ...item })) as any[],
        ''
    )
    if (!formData.id) return tree
    const disabledIds = collectSelfAndDescendants(tree, formData.id as string)
    return markDisabled(tree, disabledIds)
})

const handleSubmit = async () => {
    await formRef.value?.validate()
    mode.value == 'edit' ? await productCategoryEdit(formData) : await productCategoryAdd(formData)
    popupRef.value?.close()
    feedback.msgSuccess('操作成功')
    emit('success')
}

const open = (type = 'add') => {
    mode.value = type
    loadCategoryOptions()
    popupRef.value?.open()
}

const setFormData = (data: Partial<type_iot_product_category>) => {
    setState(data)
}

const getDetail = async (row: type_iot_product_category) => {
    const data = await productCategoryDetail({ id: row.id as string })
    setState(data)
}

const handleClose = () => {
    emit('close')
}

defineExpose({ open, setFormData, getDetail })
</script>
