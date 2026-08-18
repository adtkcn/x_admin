<template>
    <div class="code-edit">
        <el-card class="border-none!" shadow="never">
            <el-page-header content="编辑数据表" @back="$router.back()" />
        </el-card>
        <el-card class="mt-4 border-none!" shadow="never">
            <el-form
                ref="formRef"
                class="ls-form"
                :model="formData"
                label-width="100px"
                :rules="rules"
            >
                <el-tabs v-model="activeName">
                    <el-tab-pane label="基础信息" name="base">
                        <el-form-item label="表名称" prop="base.table_name">
                            <div class="w-80">
                                <el-input
                                    v-model="formData.base.table_name"
                                    placeholder="请输入表名称"
                                    clearable
                                />
                            </div>
                        </el-form-item>
                        <el-form-item label="表描述" prop="base.table_comment">
                            <div class="w-80">
                                <el-input
                                    v-model="formData.base.table_comment"
                                    placeholder="请输入表描述"
                                    clearable
                                />
                            </div>
                        </el-form-item>
                        <el-form-item label="实体类名称" prop="base.entity_name">
                            <div class="w-80">
                                <el-input
                                    v-model="formData.base.entity_name"
                                    placeholder="请输入实体类名称"
                                    clearable
                                />
                            </div>
                        </el-form-item>

                        <el-form-item label="作者" prop="base.author_name">
                            <div class="w-80">
                                <el-input
                                    v-model="formData.base.author_name"
                                    placeholder="请输入作者"
                                    clearable
                                />
                            </div>
                        </el-form-item>
                        <el-form-item label="备注">
                            <div class="w-80">
                                <el-input
                                    v-model="formData.base.remarks"
                                    class="w-full"
                                    type="textarea"
                                    :autosize="{ minRows: 4, maxRows: 4 }"
                                    maxlength="200"
                                    show-word-limit
                                    clearable
                                />
                            </div>
                        </el-form-item>
                    </el-tab-pane>
                    <el-tab-pane label="字段管理" name="column">
                        <vxe-table
                            :data="formData.column"
                            :row-config="{ keyField: 'id' }"
                            :border="'inner'"
                        >
                            <vxe-column title="字段列名" field="column_name" min-width="120" />
                            <vxe-column title="字段描述" field="column_comment" min-width="120">
                                <template #default="{ row }">
                                    <el-input v-model="row.column_comment"></el-input>
                                </template>
                            </vxe-column>
                            <vxe-column title="sql类型" field="column_type" min-width="100" />
                            <vxe-column title="go类型" min-width="140">
                                <template #default="{ row }">
                                    <el-select v-model="row.go_type">
                                        <el-option label="int" value="int" />
                                        <el-option label="int64" value="int64" />
                                        <el-option label="string" value="string" />
                                        <el-option label="float64" value="float64" />
                                        <!-- <el-option label="rune" value="rune" /> -->
                                        <el-option label="bool" value="bool" />
                                        <el-option label="time.Time" value="time.Time" />
                                    </el-select>
                                </template>
                            </vxe-column>
                            <!-- <vxe-column title="go属性" min-width="100">
                                <template #default="{ row }">
                                    <el-input v-model="row.go_field" />
                                </template>
                            </vxe-column> -->
                            <vxe-column title="必填" width="80">
                                <template #default="{ row }">
                                    <el-checkbox
                                        v-model="row.is_required"
                                        :true-label="1"
                                        :false-label="0"
                                    />
                                </template>
                            </vxe-column>
                            <vxe-column title="插入" width="80">
                                <template #default="{ row }">
                                    <el-checkbox
                                        v-model="row.is_insert"
                                        :true-label="1"
                                        :false-label="0"
                                    />
                                </template>
                            </vxe-column>
                            <vxe-column title="编辑" width="80">
                                <template #default="{ row }">
                                    <el-checkbox
                                        v-model="row.is_edit"
                                        :true-label="1"
                                        :false-label="0"
                                    />
                                </template>
                            </vxe-column>
                            <vxe-column title="列表" width="80">
                                <template #default="{ row }">
                                    <el-checkbox
                                        v-model="row.is_list"
                                        :true-label="1"
                                        :false-label="0"
                                    />
                                </template>
                            </vxe-column>
                            <vxe-column title="查询" width="80">
                                <template #default="{ row }">
                                    <el-checkbox
                                        v-model="row.is_query"
                                        :true-label="1"
                                        :false-label="0"
                                    />
                                </template>
                            </vxe-column>
                            <vxe-column title="查询方式" min-width="140">
                                <template #default="{ row }">
                                    <el-select v-model="row.query_type">
                                        <el-option label="=" value="EQ" />
                                        <el-option label="!=" value="NE" />
                                        <el-option label=">" value="GT" />
                                        <el-option label=">=" value="GTE" />
                                        <el-option label="<" value="LT" />
                                        <el-option label="<=" value="LTE" />
                                        <el-option label="LIKE" value="LIKE" />
                                        <el-option label="BETWEEN" value="BETWEEN" />
                                    </el-select>
                                </template>
                            </vxe-column>
                            <vxe-column title="显示类型" min-width="140">
                                <template #default="{ row }">
                                    <el-select v-model="row.html_type">
                                        <el-option label="文本框" value="input" />
                                        <el-option label="数字框" value="number" />
                                        <el-option label="文本域" value="textarea" />
                                        <el-option label="下拉框" value="select" />
                                        <el-option label="单选框" value="radio" />
                                        <el-option label="复选框" value="checkbox" />
                                        <el-option label="日期控件" value="datetime" />
                                        <el-option label="图片选择控件" value="imageUpload" />
                                        <el-option label="富文本控件" value="editor" />
                                    </el-select>
                                </template>
                            </vxe-column>
                            <vxe-column title="字典类型" min-width="140">
                                <template #default="{ row }">
                                    <el-select
                                        v-model="row.dict_type"
                                        clearable
                                        v-if="
                                            (row.html_type == 'select' ||
                                                row.html_type == 'radio' ||
                                                row.html_type == 'checkbox') &&
                                            !row.list_all_api
                                        "
                                        placeholder="字典类型"
                                        @change="row.list_all_api = null"
                                    >
                                        <el-option
                                            v-for="(item, index) in optionsData.dictType"
                                            :key="index"
                                            :label="item.dict_name"
                                            :value="item.dict_type"
                                            :disabled="!item.dict_status"
                                        />
                                    </el-select>
                                </template>
                            </vxe-column>

                            <vxe-column title="数据来源（字典类型优先）" min-width="280">
                                <template #default="{ row }">
                                    <el-select
                                        v-model="row.list_all_api"
                                        clearable
                                        filterable
                                        v-if="
                                            (row.html_type == 'select' ||
                                                row.html_type == 'radio' ||
                                                row.html_type == 'checkbox') &&
                                            !row.dict_type
                                        "
                                        placeholder="字典类型"
                                        @change="row.dict_type = null"
                                    >
                                        <el-option
                                            v-for="(item, index) in optionsData.ApiList"
                                            :key="index"
                                            :label="item"
                                            :value="item"
                                        />
                                    </el-select>
                                </template>
                            </vxe-column>
                        </vxe-table>
                    </el-tab-pane>
                    <el-tab-pane label="生成配置" name="config">
                        <el-form-item label="模板类型" prop="gen.gen_tpl" required>
                            <el-radio-group v-model="formData.gen.gen_tpl">
                                <el-radio :value="GenTpl.CRUD">单表（增删改查）</el-radio>
                                <el-radio :value="GenTpl.TREE">树表（增删改查）</el-radio>
                            </el-radio-group>
                        </el-form-item>
                        <el-form-item label="模块名" prop="gen.module_name">
                            <div class="w-80">
                                <el-input
                                    v-model="formData.gen.module_name"
                                    placeholder="请输入模块名"
                                    clearable
                                />
                                <div class="form-tips">生成文件所在模块名</div>
                            </div>
                        </el-form-item>
                        <el-form-item label="功能名称" prop="gen.function_name">
                            <div class="w-80">
                                <el-input
                                    v-model="formData.gen.function_name"
                                    placeholder="请输入功能名称"
                                    clearable
                                />
                            </div>
                        </el-form-item>

                        <template v-if="formData.gen.gen_tpl == GenTpl.TREE">
                            <el-form-item label="树主键字段" prop="gen.tree_primary">
                                <el-select
                                    class="w-80"
                                    v-model="formData.gen.tree_primary"
                                    clearable
                                >
                                    <el-option
                                        v-for="item in formData.column"
                                        :key="item.id"
                                        :value="item.column_name"
                                        :label="`${item.column_name}：${item.column_comment}`"
                                    />
                                </el-select>
                            </el-form-item>
                            <el-form-item label="树父级字段" prop="gen.tree_parent">
                                <el-select
                                    class="w-80"
                                    v-model="formData.gen.tree_parent"
                                    clearable
                                >
                                    <el-option
                                        v-for="item in formData.column"
                                        :key="item.id"
                                        :value="item.column_name"
                                        :label="`${item.column_name}：${item.column_comment}`"
                                    />
                                </el-select>
                            </el-form-item>
                            <el-form-item label="树名称字段" prop="gen.tree_name">
                                <el-select class="w-80" v-model="formData.gen.tree_name" clearable>
                                    <el-option
                                        v-for="item in formData.column"
                                        :key="item.id"
                                        :value="item.column_name"
                                        :label="`${item.column_name}：${item.column_comment}`"
                                    />
                                </el-select>
                            </el-form-item>
                        </template>
                    </el-tab-pane>
                </el-tabs>
            </el-form>
        </el-card>
        <footer-btns>
            <el-button type="primary" @click="handleSave">保存</el-button>
        </footer-btns>
    </div>
</template>

<script lang="ts" setup>
import { useRoute, useRouter } from 'vue-router'
import { ref, shallowRef, reactive } from 'vue'

import { generateEdit, tableDetail, type type_gen_edit_column } from '@/api/tools/code'
import { dictTypeAll } from '@/api/setting/dict'
import type { FormInstance } from 'element-plus'
import feedback from '@/utils/feedback'
import { menuLists } from '@/api/perms/menu'
import { getApiList } from '@/api/setting/website'
import { useDictOptions } from '@/hooks/useDictOptions'
import useMultipleTabs from '@/hooks/useMultipleTabs'
import { useReactiveWithReset } from '@/hooks/useReactiveWithReset'
const GenTpl = {
    CRUD: 'crud',
    TREE: 'tree'
}
defineOptions({
    name: 'tableEdit'
})
const route = useRoute()
const router = useRouter()
const { removeTab } = useMultipleTabs()
const activeName = ref('column')
const { state: formData } = useReactiveWithReset({
    base: {
        id: '',
        table_name: '',
        table_comment: '',
        entity_name: '',
        author_name: '',
        remarks: ''
    },
    column: [] as type_gen_edit_column[],
    gen: {
        function_name: '',

        gen_tpl: '',
        module_name: '',
        sub_table_fk: '',
        sub_table_name: '',
        tree_parent: '',
        tree_primary: '',
        tree_name: ''
    }
})

const formRef = shallowRef<FormInstance>()
const rules = reactive({
    ['base.table_name']: [{ required: true, message: '请输入表名称', trigger: 'blur' }],
    ['base.table_comment']: [{ required: true, message: '请输入表描述', trigger: 'blur' }],
    ['base.entity_name']: [{ required: true, message: '请输入实体类名称', trigger: 'blur' }],
    // ['base.author_name']: [{ required: true, message: '请输入作者', trigger: 'blur' }],
    ['gen.module_name']: [{ required: true, message: '请输入模块名', trigger: 'blur' }],
    ['gen.function_name']: [{ required: true, message: '请输入功能名称', trigger: 'blur' }],
    ['gen.tree_primary']: [{ required: true, message: '请选择树主键字段', trigger: 'blur' }],
    ['gen.tree_parent']: [{ required: true, message: '请选择树父级字段', trigger: 'blur' }],
    ['gen.tree_name']: [{ required: true, message: '请选择树名称字段', trigger: 'blur' }]
})

const getDetails = async () => {
    const data = await tableDetail({
        id: route.query.id as string
    })
    Object.keys(formData).forEach((key) => {
        ;(formData as any)[key] = (data as any)[key]
    })
}

const { optionsData } = useDictOptions<{
    dictType: any[]
    menu: any[]
    ApiList: string[]
}>({
    dictType: {
        api: dictTypeAll
    },
    menu: {
        api: menuLists,
        transformData(data: any) {
            const menu = { id: '', name: '顶级', children: [] }
            menu.children = data
            return menu
        }
    },
    ApiList: {
        api: getApiList,
        transformData(data: any) {
            return data.filter((item: any) => {
                return item.endsWith('list_all')
            })
        }
    }
})

const handleSave = async () => {
    try {
        await formRef.value?.validate()
        const { base, column, gen } = formData
        await generateEdit({ ...base, ...gen, columns: column })
        feedback.msgSuccess('操作成功')
        removeTab()
        router.back()
    } catch (error: any) {
        for (const err in error) {
            const isInRules = Object.keys(rules).includes(err)
            if (isInRules) {
                feedback.msgError(error[err][0]?.message)
            }
        }
    }
}

getDetails()
</script>
