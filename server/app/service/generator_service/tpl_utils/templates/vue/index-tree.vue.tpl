<template>
    <div class="index-tree">
        <el-card class="!border-none mb-4" shadow="never">
            <el-form ref="formRef" class="mb-[-16px]" :model="queryParams" :inline="true">
            {{{- range .Columns }}}
            {{{- if eq .IsQuery 1 }}}
                {{{- if eq .HtmlType "datetime" }}}
                <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ .TsField }}}">
                    <daterange-picker
                        v-model:startTime="queryParams.{{{ .TsField }}}_start"
                        v-model:endTime="queryParams.{{{ .TsField }}}_end"
                    />
                </el-form-item>
                {{{- else if or (eq .HtmlType "select") (eq .HtmlType "radio") }}}
                <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ .TsField }}}" class="w-[280px]">
                    <el-select
                        v-model="queryParams.{{{ .TsField }}}"
                        :empty-values="[null, undefined]"
                        clearable
                    >
                        {{{- if ne .DictType "" }}}
                        <el-option label="全部" value="" />
                        <el-option
                            v-for="(item, index) in dictData.{{{ .DictType }}}"
                            :key="index"
                            :label="item.name"
                            :value="item.value"
                        />
                         {{{- else if ne .ListAllApi ""}}}
                         <el-option label="全部" value="" />
                        <el-option
                            v-for="(item, index) in listAllData.{{{pathToName .ListAllApi}}}"
                            :key="index"
                            :label="item.{{{ .PrimaryTsField }}}"
                            :value="item.{{{ .PrimaryTsField }}}"
                        />
                        {{{- else }}}
                        <el-option label="请选择字典生成" value="" />
                        {{{- end }}}
                    </el-select>
                </el-form-item>
                {{{- else if eq .HtmlType "input" }}}
                <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ .TsField }}}" class="w-[280px]">
                    <el-input v-model="queryParams.{{{ .TsField }}}" />
                </el-form-item>
                {{{- end }}}
            {{{- end }}}
            {{{- end }}}
                <el-form-item>
                    <el-button type="primary" @click="getLists">查询</el-button>
                    <el-button @click="getLists">重置</el-button>
                </el-form-item>
            </el-form>
        </el-card>
        <el-card class="!border-none" shadow="never">
            <div>
                <el-button v-perms="['admin:{{{ .ModuleName }}}:add']" type="primary" @click="handleAdd()">
                    <template #icon>
                        <icon name="el-icon-Plus" />
                    </template>
                    新增
                </el-button>
                <el-button @click="handleExpand"> 展开/折叠 </el-button>
            </div>
            <vxe-table
                v-loading="loading"
                ref="tableRef"
                class="mt-4"
                 :border="'inner'"
                :data="lists"
                row-key="{{{ .Table.TreePrimary }}}"
                :tree-props="{ children: 'children', hasChildren: 'hasChildren' }"
            >
            {{{- range .Columns }}}
            {{{- if .IsList }}}
                {{{- if and (ne .DictType "") (or (eq .HtmlType "select") (eq .HtmlType "radio") (eq .HtmlType "checkbox")) }}}
                <vxe-column title="{{{ .ColumnComment }}}" field="{{{ .TableColumnProp }}}" min-width="100">
                    <template #default="{ row }">
                        <dict-value :options="dictData.{{{ .DictType }}}" :value="row.{{{ .TableColumnProp }}}" />
                    </template>
                </vxe-column>
                {{{- else if and (ne .ListAllApi "") (or (eq .HtmlType "select") (eq .HtmlType "radio") (eq .HtmlType "checkbox")) }}}
                <vxe-column title="{{{ .ColumnComment }}}" field="{{{ .TableColumnProp }}}" min-width="100">
                    <template #default="{ row }">
                        <dict-value :options="listAllData.{{{pathToName .ListAllApi }}}" :value="row.{{{ .TableColumnProp }}}" labelKey='{{{ .PrimaryTsField }}}' valueKey='{{{ .PrimaryTsField }}}' />
                    </template>
                </vxe-column>
                {{{- else if eq .HtmlType "imageUpload" }}}
                <vxe-column title="{{{ .ColumnComment }}}" field="{{{ .TableColumnProp }}}" min-width="100">
                    <template #default="{ row }">
                        <image-contain
                            :width="40"
                            :height="40"
                            :src="row.{{{ .TableColumnProp }}}"
                            :preview-src-list="[row.{{{ .TableColumnProp }}}]"
                            preview-teleported
                            hide-on-click-modal
                        />
                    </template>
                </vxe-column>
                {{{- else }}}
                <vxe-column title="{{{ .ColumnComment }}}" field="{{{ .TableColumnProp }}}" min-width="100" />
                {{{- end }}}
            {{{- end }}}
            {{{- end }}}
                <vxe-column title="操作" width="160" fixed="right">
                    <template #default="{ row }">
                        <el-button
                            v-perms="[{{{ if and .Table.TreePrimary .Table.TreeParent }}}'admin:{{{ .ModuleName }}}:list_all',{{{ end }}}'admin:{{{ .ModuleName }}}:detail']"
                            type="primary"
                            link
                            @click="viewDetails(row)"
                        >详情</el-button>
                        <el-button
                            v-perms="[{{{ if and .Table.TreePrimary .Table.TreeParent }}}'admin:{{{ .ModuleName }}}:list_all',{{{ end }}}'admin:{{{ .ModuleName }}}:add']"
                            type="primary"
                            link
                            @click="handleAdd(row.{{{ .Table.TreePrimary }}})"
                        >
                            新增
                        </el-button>
                        <el-button
                            v-perms="[{{{ if and .Table.TreePrimary .Table.TreeParent }}}'admin:{{{ .ModuleName }}}:list_all',{{{ end }}}'admin:{{{ .ModuleName }}}:edit','admin:{{{ .ModuleName }}}:detail']"
                            type="primary"
                            link
                            @click="handleEdit(row)"
                        >
                            编辑
                        </el-button>
                        <el-button
                            v-perms="['admin:{{{ .ModuleName }}}:del']"
                            type="danger"
                            link
                            @click="handleDelete(row.{{{ .PrimaryKey }}})"
                        >
                            删除
                        </el-button>
                    </template>
                </vxe-column>
            </vxe-table>
        </el-card>
        <edit-popup
            v-if="showEdit"
            ref="editRef"
            {{{- if ge (len .DictFields) 1 }}}
            :dict-data="dictData"
            {{{- end }}}
            {{{- if ge (len .ListAllFields) 1 }}}
            :list-all-data="listAllData"
            {{{- end }}}
            @success="getLists"
            @close="showEdit = false"
        />
        <DetailsPopup
            v-if="showDetails"
            ref="detailsRef"
            {{{- if ge (len .DictFields) 1 }}}
            :dict-data="dictData"
            {{{- end }}}
            {{{- if ge (len .ListAllFields) 1 }}}
            :list-all-data="listAllData"
            {{{- end }}}
            @close="showDetails = false"
        />
        
    </div>
</template>
<script lang="ts" setup>
import { ref, reactive, useTemplateRef, nextTick } from 'vue'
import { {{{ .ModuleName }}}_delete, {{{ .ModuleName }}}_list } from '@/api/{{{.Domain}}}/{{{.ModuleName}}}'
import type { type_{{{ .ModuleName }}},type_{{{.ModuleName}}}_query	} from "@/api/{{{.Domain}}}/{{{.ModuleName}}}";

import EditPopup from './edit.vue'
import DetailsPopup from './details.vue'
import feedback from '@/utils/feedback'

import { VxeTableInstance, VxeTablePropTypes } from 'vxe-table'

import { useDictData, useListAllData } from '@/hooks/useDictOptions'
import type { type_dict } from '@/hooks/useDictOptions'

defineOptions({
    name: "{{{ .ModuleName }}}"
})

const rowConfig = {
    keyField: '{{{ .Table.TreePrimary }}}'
}
const treeConfig = reactive<VxeTablePropTypes.TreeConfig>({
    rowField: '{{{ .Table.TreePrimary }}}',
    childrenField: 'children',
    indent: 10,
    reserve: true,
    lazy: true,
    transform: true,

    parentField: '{{{ .Table.TreeParent }}}'
})

const tableRef = useTemplateRef<VxeTableInstance<type_{{{ .ModuleName }}}}('tableRef')
const editRef = useTemplateRef<InstanceType<typeof EditPopup>>('editRef')

const showEdit = ref(false)

const detailsRef = useTemplateRef<InstanceType<typeof DetailsPopup>>('detailsRef')
const showDetails = ref(false)

const loading = ref(false)
const lists = ref<type_{{{ .ModuleName }}}[]>([])

const queryParams = reactive<type_{{{.ModuleName}}}_query>({
{{{- range .Columns }}}
{{{- if .IsQuery }}}
    {{{- if eq .HtmlType "datetime" }}}
    {{{ .TsField }}}_start: undefined,
    {{{ .TsField }}}_end: undefined,
    {{{- else }}}
    {{{ .TsField }}}: undefined,
    {{{- end }}}
{{{- end }}}
{{{- end }}}
})

const getLists = async () => {
    loading.value = true
    try {
        const data = await {{{ .ModuleName }}}_list(queryParams)
        lists.value = data
        loading.value = false
    } catch (error) {
        console.error('获取列表失败:', error)
        loading.value = false
    }
}

{{{- if ge (len .DictFields) 1 }}}
{{{- $dictSize := sub (len .DictFields) 1 }}}
const { dictData } = useDictData<{
{{{- range .DictFields }}}
    {{{ . }}}: type_dict[]
{{{- end }}}
}>([{{{- range .DictFields }}}'{{{ . }}}'{{{- if ne (index $.DictFields $dictSize) . }}},{{{- end }}}{{{- end }}}])
{{{- end }}}

{{{- if ge (len .ListAllFields) 1 }}}
{{{- $list_all_size := sub (len .ListAllFields) 1 }}}
const { listAllData } = useListAllData<{
    {{{- range .ListAllFields }}}
    {{{pathToName . }}}: any[]
    {{{- end }}}
}>({ 
	{{{- range .ListAllFields }}}
		{{{pathToName . }}}:'{{{deletePathPrefix . }}}',
	{{{- end }}}
})
{{{- end }}}

const handleAdd = async ({{{ .Table.TreePrimary }}}?: number) => {
    showEdit.value = true
    await nextTick()
    if ({{{ .Table.TreePrimary }}}) {
        editRef.value?.setFormData({
            {{{ .Table.TreeParent }}}: {{{ .Table.TreePrimary }}}
        } as type_{{{ .ModuleName }}}_edit)
    }
    editRef.value?.open('add')
}

const handleEdit = async (data: type_{{{ .ModuleName }}}) => {
    showEdit.value = true
    await nextTick()
    editRef.value?.open('edit')
    editRef.value?.getDetail(data)
}
const viewDetails = async (data: type_{{{ .ModuleName }}}) => {
    showDetails.value = true
    await nextTick()
    detailsRef.value?.open()
    detailsRef.value?.getDetail(data)
}
const handleDelete = async ({{{ .PrimaryTsField }}}: {{{.PrimaryTsType}}}) => {
    try {
        await feedback.confirm('确定要删除？')
        await {{{ .ModuleName }}}_delete({ {{{ .PrimaryTsField }}} })
        feedback.msgSuccess('删除成功')
        getLists()
    } catch (error) {
        console.error('删除失败:', error)
    }
}

let isExpand = false
const handleExpand = () => {
    const $table = tableRef.value
    if ($table) {
        isExpand = !isExpand
        if (isExpand) {
            $table.setAllTreeExpand(true)
        } else {
            $table.clearTreeExpand()
        }
    }
}

getLists()
</script>
