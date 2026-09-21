<template>
    <div class="edit-popup">
        <popup
            ref="popupRef"
            title="权限设置"
            :async="true"
            width="550px"
            @confirm="handleSubmit"
            @close="handleClose"
        >
            <el-form class="ls-form" ref="formRef" :model="formData">
                <el-form-item label="">
                    <el-checkbox label="展开/折叠" @change="handleExpand" />
                    <el-checkbox label="全选/不全选" @change="handleSelectAll" />
                    <el-checkbox v-model="checkStrictly" label="父子联动" />
                </el-form-item>
                <el-form-item label="">
                    <el-scrollbar class="h-[calc(100vh-270px)]! w-full">
                        <el-tree
                            ref="treeRef"
                            :data="menuTree"
                            :props="{
                                label: 'menu_name',
                                children: 'children'
                            }"
                            :check-strictly="!checkStrictly"
                            node-key="id"
                            :default-expand-all="isExpand"
                            show-checkbox
                        />
                    </el-scrollbar>
                </el-form-item>
            </el-form>
        </popup>
    </div>
</template>
<script lang="ts" setup>
import { ref, useTemplateRef, nextTick, reactive } from 'vue'
import type { CheckboxValueType, ElTree, FormInstance } from 'element-plus'
import { roleDetail, roleEdit } from '@/api/perms/role'
import { menuLists } from '@/api/perms/menu'
import Popup from '@/components/popup/index.vue'
import { treeToArray, arrayToTree } from '@/utils/util'
import feedback from '@/utils/feedback'
const emit = defineEmits(['success', 'close'])
const treeRef = useTemplateRef<InstanceType<typeof ElTree>>('treeRef')
const formRef = useTemplateRef<FormInstance>('formRef')
const popupRef = useTemplateRef<InstanceType<typeof Popup>>('popupRef')
const isExpand = ref(false)
const checkStrictly = ref(true)
const menuArray = ref<any[]>([])
const menuTree = ref<any[]>([])
const formData = reactive({
    id: '',
    name: '',
    remark: '',
    sort: 0,
    is_disable: 0,
    menus: [] as any[]
})

const getOptions = async () => {
    try {
        const data = await menuLists()
        menuTree.value = arrayToTree(data, '')
        menuArray.value = treeToArray(data)
    } catch (error) {
        console.error('菜单选项获取失败:', error)
    }
}

// 获取所有选择的节点包括半选中节点
const getDeptAllCheckedKeys = () => {
    const checkedKeys = treeRef.value?.getCheckedKeys()
    const halfCheckedKeys = treeRef.value?.getHalfCheckedKeys()!
    checkedKeys?.unshift.apply(checkedKeys, halfCheckedKeys)
    return checkedKeys
}

const setDeptAllCheckedKeys = () => {
    formData.menus.forEach((v) => {
        nextTick(() => {
            treeRef.value?.setChecked(v, true, false)
        })
    })
}

const handleExpand = (check: CheckboxValueType) => {
    //@ts-ignore
    const nodes = treeRef.value?.store._getAllNodes() as any[]
    nodes.forEach((node) => {
        node.expanded = check
    })
}

const handleSelectAll = (check: CheckboxValueType) => {
    if (check) {
        treeRef.value?.setCheckedKeys(menuArray.value.map((item) => item.id))
    } else {
        treeRef.value?.setCheckedKeys([])
    }
}

const handleSubmit = async () => {
    try {
        await formRef.value?.validate()
        formData.menus = getDeptAllCheckedKeys()!
        await roleEdit({ ...formData, menuIds: formData.menus.join() })
        popupRef.value?.close()
        feedback.msgSuccess('操作成功')
        emit('success')
    } catch (error) {
        console.error('角色权限保存失败:', error)
    }
}

const handleClose = () => {
    emit('close')
}

const open = () => {
    popupRef.value?.open()
}

const setFormData = async (row: Record<any, any>) => {
    try {
        await getOptions()
        const data = await roleDetail({
            id: row.id
        })
    for (const key in formData) {
        //@ts-ignore
        if (data[key] != null && data[key] != undefined) {
            //@ts-ignore
            formData[key] = data[key]
        }
    }
    nextTick(() => {
        setDeptAllCheckedKeys()
    })
    } catch (error) {
        console.error('角色权限详情获取失败:', error)
    }
}

defineExpose({
    open,
    setFormData
})
</script>
