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
            <el-form ref="formRef" :model="formData" label-width="80px" :rules="formRules">
                <el-form-item label="菜单类型" prop="menu_type" required>
                    <el-radio-group v-model="formData.menu_type">
                        <el-radio :value="MenuEnum.CATALOGUE">目录</el-radio>
                        <el-radio :value="MenuEnum.MENU">菜单</el-radio>
                        <el-radio :value="MenuEnum.BUTTON">按钮</el-radio>
                    </el-radio-group>
                </el-form-item>
                <el-form-item label="父级菜单" prop="pid">
                    <el-tree-select
                        class="flex-1"
                        v-model="formData.pid"
                        :data="menuTreeOptions"
                        clearable
                        node-key="id"
                        :props="{
                            label: 'menu_name',
                            disabled: 'disabled'
                        }"
                        :default-expand-all="true"
                        placeholder="请选择父级菜单"
                        check-strictly
                        :empty-values="[undefined, null]"
                    />
                </el-form-item>
                <el-form-item label="菜单名称" prop="menu_name">
                    <el-input v-model="formData.menu_name" placeholder="请输入菜单名称" clearable />
                </el-form-item>
                <el-form-item
                    v-if="formData.menu_type != MenuEnum.BUTTON"
                    label="菜单图标"
                    prop="menu_icon"
                >
                    <icon-picker class="flex-1" v-model="formData.menu_icon" />
                </el-form-item>

                <el-form-item
                    v-if="formData.menu_type == MenuEnum.MENU"
                    label="组件路径"
                    prop="component"
                >
                    <div class="flex-1">
                        <el-autocomplete
                            class="w-full"
                            v-model="formData.component"
                            :fetch-suggestions="querySearch"
                            clearable
                            placeholder="请输入组件路径"
                        />
                        <!-- <el-input v-model="formData.component" placeholder="请输入组件路径" /> -->
                        <div class="form-tips">
                            访问的组件路径，如：`permission/admin/index`，默认在`views`目录下
                        </div>
                    </div>
                </el-form-item>

                <el-form-item
                    v-if="formData.menu_type != MenuEnum.BUTTON"
                    label="路由路径"
                    prop="paths"
                >
                    <div class="flex-1">
                        <el-input v-model="formData.paths" placeholder="请输入路由路径" clearable />
                        <div class="form-tips">
                            访问的路由地址，如：`admin`，如外网地址需内链访问则以`http(s)://`开头
                        </div>
                    </div>
                </el-form-item>
                <el-form-item
                    v-if="formData.menu_type == MenuEnum.MENU"
                    label="路由参数"
                    prop="params"
                >
                    <div>
                        <div class="flex-1">
                            <el-input
                                v-model="formData.params"
                                placeholder="请输入路由参数"
                                clearable
                            />
                        </div>
                        <div class="form-tips">
                            访问路由的默认传递参数，如：`{"id": 1, "name":
                            "admin"}`或`id=1&name=admin`
                        </div>
                    </div>
                </el-form-item>
                <el-form-item
                    label="选中菜单"
                    prop="selected"
                    v-if="formData.menu_type == MenuEnum.MENU"
                >
                    <div class="flex-1">
                        <el-input
                            v-model="formData.selected"
                            placeholder="请输入路由路径"
                            clearable
                        />
                        <div class="form-tips">
                            访问详情页面，编辑页面时，菜单高亮显示，如`/consumer/lists`
                        </div>
                    </div>
                </el-form-item>
                <el-form-item
                    v-if="formData.menu_type != MenuEnum.CATALOGUE"
                    label="接口权限"
                    prop="perms"
                >
                    <div class="flex-1">
                        <!-- {{ formData.permsArr }} -->
                        <!-- <el-input v-model="formData.perms" placeholder="请输入接口权限" clearable /> -->
                        <el-select
                            v-model="formData.perms"
                            clearable
                            filterable
                            placeholder="请选择接口权限"
                            :style="{ width: '100%' }"
                        >
                            <el-option
                                v-for="(item, index) in permissionOptions"
                                :key="index"
                                :label="item.label"
                                :value="item.value"
                            ></el-option>
                        </el-select>

                        <!-- <div class="form-tips">
                            请求路径`api/admin/system/admin/list`权限为`admin:system:admin:list`
                        </div> -->
                    </div>
                </el-form-item>

                <el-form-item
                    v-if="formData.menu_type == MenuEnum.MENU"
                    label="是否缓存"
                    prop="is_cache"
                    required
                >
                    <div>
                        <el-radio-group v-model="formData.is_cache">
                            <el-radio :value="1">缓存</el-radio>
                            <el-radio :value="0">不缓存</el-radio>
                        </el-radio-group>
                        <div class="form-tips">选择缓存则会被`keep-alive`缓存</div>
                    </div>
                </el-form-item>
                <el-form-item
                    v-if="formData.menu_type != MenuEnum.BUTTON"
                    label="是否显示"
                    prop="is_show"
                    required
                >
                    <div>
                        <el-radio-group v-model="formData.is_show">
                            <el-radio :value="1">显示</el-radio>
                            <el-radio :value="0">隐藏</el-radio>
                        </el-radio-group>
                        <div class="form-tips">
                            选择隐藏则路由将不会出现在侧边栏，但仍然可以访问
                        </div>
                    </div>
                </el-form-item>
                <el-form-item
                    v-if="formData.menu_type != MenuEnum.BUTTON"
                    label="菜单状态"
                    prop="is_disable"
                    required
                >
                    <div>
                        <el-radio-group v-model="formData.is_disable">
                            <el-radio :value="0">正常</el-radio>
                            <el-radio :value="1">停用</el-radio>
                        </el-radio-group>
                        <div class="form-tips">选择停用则路由将不会出现在侧边栏，也不能被访问</div>
                    </div>
                </el-form-item>
                <el-form-item label="菜单排序" prop="menu_sort">
                    <div>
                        <el-input-number v-model="formData.menu_sort" :max="9999" />
                        <div class="form-tips">数值越大越排前</div>
                    </div>
                </el-form-item>
            </el-form>
        </popup>
    </div>
</template>
<script lang="ts" setup>
import { ref, computed, shallowRef } from 'vue'
import type { FormInstance } from 'element-plus'
import {
    menuLists,
    menuEdit,
    menuAdd,
    menuDetail,
    type type_system_menu_edit,
    type type_system_menu_resp
} from '@/api/perms/menu'
import { getApiList } from '@/api/setting/website'
import { getModulesKey } from '@/router'
import { MenuEnum } from '@/enums/appEnums'
import Popup from '@/components/popup/index.vue'
import feedback from '@/utils/feedback'
import { arrayToTree } from '@/utils/util'
import { useReactiveWithReset } from '@/hooks/useReactiveWithReset'

const emit = defineEmits(['success', 'close'])
const formRef = shallowRef<FormInstance>()
const popupRef = shallowRef<InstanceType<typeof Popup>>()
const mode = ref('add')
const popupTitle = computed(() => {
    return mode.value == 'edit' ? '编辑菜单' : '新增菜单'
})
const permissionOptions = ref<{ value: string; label: string }[]>([])

const componentsOptions = ref(getModulesKey())
const querySearch = (queryString: string, cb: any) => {
    const results = queryString
        ? componentsOptions.value.filter((item) =>
              item.toLowerCase().includes(queryString.toLowerCase())
          )
        : componentsOptions.value
    cb(results.map((item) => ({ value: item })))
}

const {
    state: formData,
    setState
} = useReactiveWithReset<type_system_menu_edit>({
    id: '',
    pid: '',
    menu_type: MenuEnum.CATALOGUE,
    menu_icon: '',
    menu_name: '',
    menu_sort: 0,
    paths: '',
    perms: '',
    component: '',
    selected: '',
    params: '',
    is_cache: 1,
    is_show: 1,
    is_disable: 0
})

const formRules = {
    // pid: [
    //     {
    //         required: true,
    //         message: '请选择父级菜单',
    //         trigger: ['blur', 'change']
    //     }
    // ],
    menu_name: [
        {
            required: true,
            message: '请输入菜单名称',
            trigger: 'blur'
        }
    ],
    paths: [
        {
            required: true,
            message: '请输入路由地址',
            trigger: 'blur'
        }
    ]
    // component: [
    //     {
    //         required: true,
    //         message: '请输入组件地址',
    //         trigger: 'blur'
    //     }
    // ]
}
const menuOptions = ref<type_system_menu_resp[]>([])
// 所有菜单的平铺数据，用于构建“当前菜单及其所有后代”的禁用集合
const allMenuData = ref<type_system_menu_resp[]>([])

const getMenu = async () => {
    const data = await menuLists()
    allMenuData.value = data.filter((item) => item.menu_type != MenuEnum.BUTTON)
    const menu: type_system_menu_resp = { id: '', menu_name: '顶级', children: [] } as any
    menu.children = arrayToTree(allMenuData.value, '')
    menuOptions.value = [menu]
}

// 收集自身及所有后代菜单 ID（基于 pid 关系，不依赖 children 字段名）
const collectSelfAndDescendants = (list: type_system_menu_resp[], currentId: string) => {
    const childrenMap = new Map<string, string[]>()
    list.forEach((item) => {
        if (!childrenMap.has(item.pid)) childrenMap.set(item.pid, [])
        childrenMap.get(item.pid)!.push(item.id)
    })
    const ids = new Set<string>()
    const queue = [currentId]
    while (queue.length) {
        const cur = queue.shift()!
        ids.add(cur)
        ;(childrenMap.get(cur) || []).forEach((childId) => {
            if (!ids.has(childId)) {
                ids.add(childId)
                queue.push(childId)
            }
        })
    }
    return ids
}

// 编辑模式下禁用当前菜单及其所有子级，避免选自己或自己的子孙作为上级形成环路
const markDisabled = (
    list: type_system_menu_resp[],
    disabledIds: Set<string>
): type_system_menu_resp[] => {
    return list.map((item) => {
        const children = item.children?.length
            ? markDisabled(item.children, disabledIds)
            : item.children
        return {
            ...item,
            disabled: disabledIds.has(item.id),
            children
        }
    })
}
const menuTreeOptions = computed(() => {
    const list = menuOptions.value
    if (!list?.length || !formData.id) return list
    const disabledIds = collectSelfAndDescendants(allMenuData.value, formData.id)
    return markDisabled(list, disabledIds)
})
function getApiListFn() {
    getApiList().then((res: string[]) => {
        const arr: { value: string; label: string }[] = []
        res.forEach((item: string) => {
            if (item.indexOf('/api/admin/') == 0) {
                const per = item.replace(/\/api\//, '').replace(/\//g, ':')
                arr.push({
                    value: per,
                    label: item
                })
            }
        })
        permissionOptions.value = arr.sort()
    })
}
const handleSubmit = async () => {
    try {
        await formRef.value?.validate()
    } catch (error) {
        // 表单校验未通过，Element 已自动提示，直接中断提交
        return
    }
    const data = { ...formData }

    try {
        if (mode.value == 'edit') {
            await menuEdit(data)
        } else {
            await menuAdd(data)
        }
        popupRef.value?.close()
        feedback.msgSuccess('操作成功')
        emit('success')
    } catch (error) {
        console.error('菜单提交失败:', error)
        feedback.msgError((error as Error)?.message || '操作失败,请稍后重试')
    }
}

const open = (type = 'add') => {
    mode.value = type
    popupRef.value?.open()
}

const setFormData = (data: type_system_menu_edit) => {
    setState(data)
}

const getDetail = async (row: type_system_menu_resp) => {
    const data = await menuDetail({
        id: row.id
    })
    setFormData(data)
}

const handleClose = () => {
    emit('close')
}

getMenu()
getApiListFn()

defineExpose({
    open,
    setFormData,
    getDetail
})
</script>
