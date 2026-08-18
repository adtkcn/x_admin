<template>
    <div class="index-lists">
        <el-card class="!border-none" shadow="never">
            <el-form ref="formRef" class="mb-[-16px]" :model="queryParams" :inline="true" label-width="90px"
                label-position="right">
            {{{- range .Columns }}}
            {{{- if eq .IsQuery 1 }}}
                {{{- if eq .HtmlType "datetime" }}}
                <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ .TsField }}}" class="w-[400px]">
                    <daterange-picker
                        v-model:startTime="queryParams.{{{ .TsField }}}_start"
                        v-model:endTime="queryParams.{{{ .TsField }}}_end"
                    />
                </el-form-item>
                {{{- else if or (eq .HtmlType "select") (eq .HtmlType "radio") }}}
                <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ .TsField }}}"  class="w-[280px]">
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
                            :label="item.ID"
                            :value="item.ID"
                        />
                        {{{- else }}}
                        <el-option label="请选择字典生成" value="" />
                        {{{- end }}}
                    </el-select>
                </el-form-item>
                {{{- else if eq .HtmlType "input" }}}
                <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ .TsField }}}" class="w-[280px]">
                    <el-input  v-model="queryParams.{{{ .TsField }}}" />
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
                    <Upload
                    v-perms="['admin:{{{ .ModuleName }}}:import_file']"
                    class="ml-3 mr-3"
                    :url="{{{.ModuleName}}}_import_file"
                    :ext="['xlsx']"
                    :show-progress="true"
                    @change="resetPage"
                >
                    <el-button type="primary">
                        <template #icon>
                            <icon name="el-icon-Upload" />
                        </template>
                        导入
                    </el-button>
                </Upload>
                <el-button v-perms="['admin:{{{ .ModuleName }}}:export_file']" type="primary" @click="export_file">
                    <template #icon>
                        <icon name="el-icon-Download" />
                    </template>
                    导出
                </el-button>
                <el-button
                    v-perms="['admin:{{{ .ModuleName }}}:del_batch']"
                    type="danger"
                    :disabled="!multipleSelection.length"
                    @click="deleteBatch"
                >
                    批量删除
                </el-button>
            </div>
            <vxe-table
                ref="tableRef"
                class="mt-4"
                 
                :loading="pager.loading"
                :data="pager.lists"
                :row-config="{ keyField: '{{{ .PrimaryTsField }}}' }"
                :border="'inner'"
                @checkbox-change="handleSelectionChange"
                @checkbox-all="handleSelectionChange"
            >
                <vxe-column type="checkbox" width="55" />
                <vxe-column type="seq" title="序号" min-width="60" />
            {{{- range .Columns }}}
            {{{- if and .IsList .IsListShow }}}
                {{{- if and (ne .DictType "") (or (eq .HtmlType "select") (eq .HtmlType "radio") (eq .HtmlType "checkbox")) }}}
                <vxe-column title="{{{ .ColumnComment }}}" field="{{{.TableColumnProp}}}" width="100">
                    <template #default="{ row }">
                       <dict-value :options="dictData.{{{ .DictType }}}" :value="row.{{{.TableColumnProp}}}" />
                    </template>
                </vxe-column>
                {{{- else if and (ne .ListAllApi "") (or (eq .HtmlType "select") (eq .HtmlType "radio") (eq .HtmlType "checkbox")) }}}
                <vxe-column title="{{{ .ColumnComment }}}" field="{{{.TableColumnProp}}}" min-width="100">
                    <template #default="{ row }">
                        <dict-value :options="listAllData.{{{pathToName .ListAllApi }}}" :value="row.{{{.TableColumnProp}}}" labelKey='ID' valueKey='ID' />
                    </template>
                </vxe-column>

                {{{- else if eq .HtmlType "imageUpload" }}}
                <vxe-column title="{{{ .ColumnComment }}}" field="{{{.TableColumnProp}}}" min-width="100">
                    <template #default="{ row }">
                        <image-contain
                            :height="100%"
                            :src="row.{{{.TableColumnProp}}}"
                            :preview-src-list="[row.{{{.TableColumnProp}}}]"
                            preview-teleported
                            hide-on-click-modal
                        />
                    </template>
                </vxe-column>
                {{{- else }}}
                <vxe-column title="{{{ .ColumnComment }}}" field="{{{.TableColumnProp}}}" min-width="130" show-overflow />
                {{{- end }}}
            {{{- end }}}
            {{{- end }}}
                <vxe-column title="操作" width="160" fixed="right">
                    <template #default="{ row }">
                        <el-button
                            v-perms="['admin:{{{ .ModuleName }}}:detail']"
                            type="primary"
                            link
                            @click="viewDetails(row)"
                        >详情</el-button>
                        <el-button
                            v-perms="['admin:{{{ .ModuleName }}}:edit','admin:{{{ .ModuleName }}}:detail']"
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
                            @click="handleDelete(row.{{{ .PrimaryTsField }}})"
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
import { ref,reactive,shallowRef,nextTick,useTemplateRef  } from 'vue'
import { {{{ .ModuleName }}}_delete,{{{ .ModuleName }}}_delete_batch, {{{ .ModuleName }}}_list,{{{.ModuleName}}}_import_file, {{{.ModuleName}}}_export_file } from '@/api/{{{.Domain}}}/{{{.ModuleName}}}'
import type { type_{{{ .ModuleName }}},type_{{{.ModuleName}}}_query	} from "@/api/{{{.Domain}}}/{{{.ModuleName}}}";

import { VxeTableInstance,VxeTablePropTypes } from 'vxe-table'

import { useDictData,useListAllData } from '@/hooks/useDictOptions'
import type { type_dict } from '@/hooks/useDictOptions'

import { usePaging } from '@/hooks/usePaging'
import feedback from '@/utils/feedback'
import EditPopup from './edit.vue'
import DetailsPopup from './details.vue'
defineOptions({
    name:"{{{ .ModuleName }}}"
})
const editRef = shallowRef<InstanceType<typeof EditPopup>>()
const showEdit = ref(false)
const detailsRef = shallowRef<InstanceType<typeof DetailsPopup>>()
const showDetails = ref(false)
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
const tableRef = useTemplateRef<VxeTableInstance<type_{{{ .ModuleName }}}>>('tableRef')
const multipleSelection = ref<type_{{{ .ModuleName }}}[]>([])
const handleSelectionChange = () => {
    multipleSelection.value = tableRef.value?.getCheckboxRecords() || []
}

const handleDelete = async ({{{ .PrimaryTsField }}}: {{{.PrimaryTsType}}}) => {
    try {
        await feedback.confirm('确定要删除？')
        await {{{ .ModuleName }}}_delete( {{{ .PrimaryTsField }}} )
        feedback.msgSuccess('删除成功')
        getLists()
    } catch (error) {
        console.error('删除失败:', error)
    }
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
            Ids: multipleSelection.value.map((item) => item.{{{ .PrimaryTsField }}}).join(',')
        })
        feedback.msgSuccess('删除成功')
        getLists()
    } catch (error) {
        console.error('批量删除失败:', error)
    }
}

const export_file = async () => {
    try {
        await feedback.confirm('确定要导出？')
        await {{{.ModuleName}}}_export_file(queryParams)
    } catch (error) {
        console.error('导出失败:', error)
    }
}
getLists()
</script>
