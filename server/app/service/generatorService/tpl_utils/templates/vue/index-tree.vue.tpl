<template>
    <div class="index-tree">
        <el-card class="!border-none mb-4" shadow="never">
            <el-form ref="formRef" class="mb-[-16px]" :model="queryParams" :inline="true">
            {{{- range .Columns }}}
            {{{- if eq .IsQuery 1 }}}
                {{{- if eq .HtmlType "datetime" }}}
                <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ .TsType }}}">
                    <daterange-picker
                        v-model:startTime="queryParams.{{{ .TsType }}}Start"
                        v-model:endTime="queryParams.{{{ .TsType }}}End"
                    />
                </el-form-item>
                {{{- else if or (eq .HtmlType "select") (eq .HtmlType "radio") }}}
                <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ .TsType }}}" class="w-[280px]">
                    <el-select
                        v-model="queryParams.{{{ .TsType }}}"
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
                            :label="item.{{{toUpperCamelCase .PrimaryKey }}}"
                            :value="item.{{{toUpperCamelCase .PrimaryKey }}}"
                        />
                        {{{- else }}}
                        <el-option label="请选择字典生成" value="" />
                        {{{- end }}}
                    </el-select>
                </el-form-item>
                {{{- else if eq .HtmlType "input" }}}
                <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ .TsType }}}" class="w-[280px]">
                    <el-input v-model="queryParams.{{{ .TsType }}}" />
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
                size="large"
                :data="lists"
                row-key="{{{ .Table.TreePrimary }}}"
                :tree-props="{ children: 'children', hasChildren: 'hasChildren' }"
            >
            {{{- range .Columns }}}
            {{{- if .IsList }}}
                {{{- if and (ne .DictType "") (or (eq .HtmlType "select") (eq .HtmlType "radio") (eq .HtmlType "checkbox")) }}}
                <vxe-column title="{{{ .ColumnComment }}}" field="{{{ .TsType }}}" min-width="100">
                    <template #default="{ row }">
                        <dict-value :options="dictData.{{{ .DictType }}}" :value="row.{{{ .TsType }}}" />
                    </template>
                </vxe-column>
                {{{- else if and (ne .ListAllApi "") (or (eq .HtmlType "select") (eq .HtmlType "radio") (eq .HtmlType "checkbox")) }}}
                <vxe-column title="{{{ .ColumnComment }}}" field="{{{ .TsType }}}" min-width="100">
                    <template #default="{ row }">
                        <dict-value :options="listAllData.{{{pathToName .ListAllApi }}}" :value="row.{{{ .TsType }}}" labelKey='{{{toUpperCamelCase .PrimaryKey }}}' valueKey='{{{toUpperCamelCase .PrimaryKey }}}' />
                    </template>
                </vxe-column>
                {{{- else if eq .HtmlType "imageUpload" }}}
                <vxe-column title="{{{ .ColumnComment }}}" field="{{{ .TsType }}}" min-width="100">
                    <template #default="{ row }">
                        <image-contain
                            :width="40"
                            :height="40"
                            :src="row.{{{ .TsType }}}"
                            :preview-src-list="[row.{{{ .TsType }}}]"
                            preview-teleported
                            hide-on-click-modal
                        />
                    </template>
                </vxe-column>
                {{{- else }}}
                <vxe-column title="{{{ .ColumnComment }}}" field="{{{ .TsType }}}" min-width="100" />
                {{{- end }}}
            {{{- end }}}
            {{{- end }}}
                <vxe-column title="操作" width="160" fixed="right">
                    <template #default="{ row }">
                        <el-button
                            v-perms="[{{{ if and .Table.TreePrimary .Table.TreeParent }}}'admin:{{{ .ModuleName }}}:listAll',{{{ end }}}'admin:{{{ .ModuleName }}}:detail']"
                            type="primary"
                            link
                            @click="viewDetails(row)"
                        >详情</el-button>
                        <el-button
                            v-perms="[{{{ if and .Table.TreePrimary .Table.TreeParent }}}'admin:{{{ .ModuleName }}}:listAll',{{{ end }}}'admin:{{{ .ModuleName }}}:add']"
                            type="primary"
                            link
                            @click="handleAdd(row.{{{ .Table.TreePrimary }}})"
                        >
                            新增
                        </el-button>
                        <el-button
                            v-perms="[{{{ if and .Table.TreePrimary .Table.TreeParent }}}'admin:{{{ .ModuleName }}}:listAll',{{{ end }}}'admin:{{{ .ModuleName }}}:edit','admin:{{{ .ModuleName }}}:detail']"
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
import { ref,shallowRef,reactive,onMounted,watch } from 'vue'
import { {{{ .ModuleName }}}_delete, {{{ .ModuleName }}}_list } from '@/api/{{{nameToPath .ModuleName }}}'
import type { type_{{{ .ModuleName }}},type_{{{.ModuleName}}}_query	} from "@/api/{{{nameToPath .ModuleName }}}";

import EditPopup from './edit.vue'
import DetailsPopup from './details.vue'
import feedback from '@/utils/feedback'

import { VxeTableInstance } from 'vxe-table'

import { useDictData,useListAllData } from '@/hooks/useDictOptions'
import type { type_dict } from '@/hooks/useDictOptions'

//import type { ElTable } from 'element-plus'
defineOptions({
    name:"{{{ .ModuleName }}}"
})
//const tableRef = shallowRef<InstanceType<typeof ElTable>>()
//const editRef = shallowRef<InstanceType<typeof EditPopup>>()

const rowConfig = {
    keyField: '{{{ .Table.TreePrimary }}}'
}
const treeConfig = {
    rowField: '{{{ .Table.TreePrimary }}}',
    childrenField: 'children',
    indent: 10,
    reserve: true,
    lazy: true,
    transform: true,

    parentField: '{{{ .Table.TreeParent }}}'
}

const tableRef = useTemplateRef<VxeTableInstance<any>>('tableRef')
const editRef = useTemplateRef<InstanceType<typeof EditPopup>>('editRef')

const showEdit = ref(false)

const detailsRef = shallowRef<InstanceType<typeof DetailsPopup>>()
const showDetails = ref(false)

const loading = ref(false)
const lists = ref<type_{{{ .ModuleName }}}[]>([])

const queryParams = reactive<type_{{{.ModuleName}}}_query>({
{{{- range .Columns }}}
{{{- if .IsQuery }}}
    {{{- if eq .HtmlType "datetime" }}}
    {{{ .TsType }}}Start: null,
    {{{ .TsType }}}End: null,
    {{{- else }}}
    {{{ .TsType }}}: null,
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
        })
    }
    editRef.value?.open('add')
}

const handleEdit = async (data: any) => {
    showEdit.value = true
    await nextTick()
    editRef.value?.open('edit')
    editRef.value?.getDetail(data)
}
const viewDetails = async (data: any) => {
    showDetails.value = true
    await nextTick()
    detailsRef.value?.open()
    detailsRef.value?.getDetail(data)
}
const handleDelete = async ({{{toUpperCamelCase .PrimaryKey }}}: {{{goToTsType .PrimaryKeyGoType}}}) => {
    await feedback.confirm('确定要删除？')
    await {{{ .ModuleName }}}_delete({ {{{toUpperCamelCase .PrimaryKey }}} })
    feedback.msgSuccess('删除成功')
    getLists()
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
