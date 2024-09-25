<template>
    <div class="index-lists">
        <el-card class="!border-none" shadow="never">
            <el-form ref="formRef" class="mb-[-16px]" :model="queryParams" :inline="true" label-width="70px"
                label-position="left">
            {{{- range .Columns }}}
            {{{- if eq .IsQuery 1 }}}
                {{{- if eq .HtmlType "datetime" }}}
                <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ (toUpperCamelCase .GoField) }}}" class="w-[280px]">
                    <daterange-picker
                        v-model:startTime="queryParams.{{{ (toUpperCamelCase .GoField) }}}Start"
                        v-model:endTime="queryParams.{{{ (toUpperCamelCase .GoField) }}}End"
                    />
                </el-form-item>
                {{{- else if or (eq .HtmlType "select") (eq .HtmlType "radio") }}}
                <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ (toUpperCamelCase .GoField) }}}"  class="w-[280px]">
                    <el-select
                        v-model="queryParams.{{{ (toUpperCamelCase .GoField) }}}"
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
                <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ (toUpperCamelCase .GoField) }}}" class="w-[280px]">
                    <el-input  v-model="queryParams.{{{ (toUpperCamelCase .GoField) }}}" />
                </el-form-item>
                {{{- end }}}
            {{{- end }}}
            {{{- end }}}
                <el-form-item>
                    <el-button type="primary" @click="resetPage">查询</el-button>
                    <el-button @click="resetParams">重置</el-button>
                </el-form-item>
            </el-form>
        </el-card>
        <el-card class="!border-none mt-4" shadow="never">
            <div class="text-right">
                <el-button v-perms="['admin:{{{ .ModuleName }}}:add']" type="primary" @click="handleAdd()">
                    <template #icon>
                        <icon name="el-icon-Plus" />
                    </template>
                    新增
                </el-button>
                    <upload
                    v-perms="['admin:{{{ .ModuleName }}}:ImportFile']"
                    class="ml-3 mr-3"
                    :url="{{{.ModuleName}}}_import_file"
                    :data="{ cid: 0 }"
                    type="file"
                    :show-progress="true"
                    @change="resetPage"
                >
                    <el-button type="primary">
                        <template #icon>
                            <icon name="el-icon-Upload" />
                        </template>
                        导入
                    </el-button>
                </upload>
                <el-button v-perms="['admin:{{{ .ModuleName }}}:ExportFile']" type="primary" @click="exportFile">
                    <template #icon>
                        <icon name="el-icon-Download" />
                    </template>
                    导出
                </el-button>
                <el-button
                    v-perms="['admin:{{{ .ModuleName }}}:delBatch']"
                    type="danger"
                    :disabled="!multipleSelection.length"
                    @click="deleteBatch"
                >
                    批量删除
                </el-button>
            </div>
            <el-table
                class="mt-4"
                size="large"
                v-loading="pager.loading"
                :data="pager.lists"
                @selection-change="handleSelectionChange"
            >
                <el-table-column type="selection" width="55" />
            {{{- range .Columns }}}
            {{{- if .IsList }}}
                {{{- if and (ne .DictType "") (or (eq .HtmlType "select") (eq .HtmlType "radio") (eq .HtmlType "checkbox")) }}}
                <el-table-column label="{{{ .ColumnComment }}}" prop="{{{ (toUpperCamelCase .GoField) }}}" min-width="100">
                    <template #default="{ row }">
                       <dict-value :options="dictData.{{{ .DictType }}}" :value="row.{{{ (toUpperCamelCase .GoField) }}}" />
                    </template>
                </el-table-column>
                {{{- else if and (ne .ListAllApi "") (or (eq .HtmlType "select") (eq .HtmlType "radio") (eq .HtmlType "checkbox")) }}}
                <el-table-column label="{{{ .ColumnComment }}}" prop="{{{ (toUpperCamelCase .GoField) }}}" min-width="100">
                    <template #default="{ row }">
                        <dict-value :options="listAllData.{{{pathToName .ListAllApi }}}" :value="row.{{{ (toUpperCamelCase .GoField) }}}" labelKey='{{{toUpperCamelCase .PrimaryKey }}}' valueKey='{{{toUpperCamelCase .PrimaryKey }}}' />
                    </template>
                </el-table-column>

                {{{- else if eq .HtmlType "imageUpload" }}}
                <el-table-column label="{{{ .ColumnComment }}}" prop="{{{ (toUpperCamelCase .GoField) }}}" min-width="100">
                    <template #default="{ row }">
                        <image-contain
                            :width="40"
                            :height="40"
                            :src="row.{{{ (toUpperCamelCase .GoField) }}}"
                            :preview-src-list="[row.{{{ (toUpperCamelCase .GoField) }}}]"
                            preview-teleported
                            hide-on-click-modal
                        />
                    </template>
                </el-table-column>
                {{{- else }}}
                <el-table-column label="{{{ .ColumnComment }}}" prop="{{{ (toUpperCamelCase .GoField) }}}" min-width="130" />
                {{{- end }}}
            {{{- end }}}
            {{{- end }}}
                <el-table-column label="操作" width="120" fixed="right">
                    <template #default="{ row }">
                        <el-button
                            v-perms="['admin:{{{ .ModuleName }}}:edit']"
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
                            @click="handleDelete(row.{{{toUpperCamelCase .PrimaryKey }}})"
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
            {{{- if ge (len .DictFields) 1 }}}
            :dict-data="dictData"
            {{{- end }}}
            {{{- if ge (len .ListAllFields) 1 }}}
            :list-all-data="listAllData"
            {{{- end }}}
            @success="getLists"
            @close="showEdit = false"
        />
    </div>
</template>
<script lang="ts" setup>
import { {{{ .ModuleName }}}_delete,{{{ .ModuleName }}}_delete_batch, {{{ .ModuleName }}}_list,{{{.ModuleName}}}_import_file, {{{.ModuleName}}}_export_file } from '@/api/{{{nameToPath .ModuleName }}}'
import type { type_{{{ .ModuleName }}},type_{{{.ModuleName}}}_query	} from "@/api/{{{nameToPath .ModuleName }}}";


import { useDictData,useListAllData } from '@/hooks/useDictOptions'
import type { type_dict } from '@/hooks/useDictOptions'

import { usePaging } from '@/hooks/usePaging'
import feedback from '@/utils/feedback'
import EditPopup from './edit.vue'
defineOptions({
    name:"{{{ .ModuleName }}}"
})
const editRef = shallowRef<InstanceType<typeof EditPopup>>()
const showEdit = ref(false)
const queryParams = reactive<type_{{{.ModuleName}}}_query>({
{{{- range .Columns }}}
{{{- if .IsQuery }}}
    {{{- if eq .HtmlType "datetime" }}}
    {{{ (toUpperCamelCase .GoField) }}}Start: null,
    {{{ (toUpperCamelCase .GoField) }}}End: null,
    {{{- else }}}
    {{{ (toUpperCamelCase .GoField) }}}: null,
    {{{- end }}}
{{{- end }}}
{{{- end }}}
})

const { pager, getLists, resetPage, resetParams } = usePaging<type_{{{ .ModuleName }}}>({
    fetchFun: {{{ .ModuleName }}}_list,
    params: queryParams
})

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


const handleAdd = async () => {
    showEdit.value = true
    await nextTick()
    editRef.value?.open('add')
}

const handleEdit = async (data: any) => {
    showEdit.value = true
    await nextTick()
    editRef.value?.open('edit')
    editRef.value?.getDetail(data)
}
const multipleSelection = ref<type_{{{ .ModuleName }}}[]>([])
const handleSelectionChange = (val: type_{{{ .ModuleName }}}[]) => {
    console.log(val)
    multipleSelection.value = val
}

const handleDelete = async ({{{toUpperCamelCase .PrimaryKey }}}: number) => {
    try {
        await feedback.confirm('确定要删除？')
        await {{{ .ModuleName }}}_delete( {{{toUpperCamelCase .PrimaryKey }}} )
        feedback.msgSuccess('删除成功')
        getLists()
    } catch (error) {}
}
// 批量删除
const deleteBatch = async () => {
    if (multipleSelection.value.length === 0) {
        feedback.msgError('请选择要删除的数据')
        return
    }
    try {
        await feedback.confirm('确定要删除？')
        await {{{ .ModuleName }}}_delete_batch({
            Ids: multipleSelection.value.map((item) => item.{{{toUpperCamelCase .PrimaryKey }}}).join(',')
        })
        feedback.msgSuccess('删除成功')
        getLists()
    } catch (error) {}
}

const exportFile = async () => {
    try {
        await feedback.confirm('确定要导出？')
        await {{{.ModuleName}}}_export_file(queryParams)
    } catch (error) {}
}
getLists()
</script>
