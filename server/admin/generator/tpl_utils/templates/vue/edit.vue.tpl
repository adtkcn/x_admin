<template>
    <div class="edit-popup">
        <popup
            ref="popupRef"
            :title="popupTitle"
            :async="true"
            width="550px"
            :clickModalClose="true"
            @confirm="handleSubmit"
            @close="handleClose"
        >
            <el-form ref="formRef" :model="formData" label-width="84px" :rules="formRules">
                {{{- if and .Table.TreePrimary .Table.TreeParent }}}
                    <el-form-item label="父级" prop="{{{ (toUpperCamelCase .Table.TreeParent) }}}">
                        <el-tree-select
                            class="flex-1"
                            v-model="formData.{{{ (toUpperCamelCase .Table.TreeParent) }}}"
                            :data="treeList"
                            clearable
                            node-key="{{{ .Table.TreePrimary }}}"
                            :props="{ label: '{{{ (toUpperCamelCase .Table.TreeName) }}}', value: '{{{ (toUpperCamelCase .Table.TreePrimary) }}}', children: 'children' }"
                            :default-expand-all="true"
                            placeholder="请选择父级"
                            check-strictly
                        />
                    </el-form-item>
                {{{- end }}}
            {{{- range .Columns }}}
                {{{- if .IsEdit }}}
                {{{- if ne (toUpperCamelCase .GoField) "Id" }}}
                    {{{- if eq .HtmlType "input" }}}
                        <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ (toUpperCamelCase .GoField) }}}">
                            <el-input v-model="formData.{{{ (toUpperCamelCase .GoField) }}}" placeholder="请输入{{{ .ColumnComment }}}" />
                        </el-form-item>
                    {{{- else if eq .HtmlType "number" }}}
                        <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ (toUpperCamelCase .GoField) }}}">
                            <el-input v-model="formData.{{{ (toUpperCamelCase .GoField) }}}" type="number" placeholder="请输入{{{ .ColumnComment }}}" />
                        </el-form-item>
                    {{{- else if eq .HtmlType "textarea" }}}
                        <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ (toUpperCamelCase .GoField) }}}">
                            <el-input
                                v-model="formData.{{{ (toUpperCamelCase .GoField) }}}"
                                placeholder="请输入{{{ .ColumnComment }}}"
                                type="textarea"
                                :autosize="{ minRows: 4, maxRows: 6 }"
                            />
                        </el-form-item>
                    {{{- else if eq .HtmlType "checkbox" }}}
                        <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ (toUpperCamelCase .GoField) }}}">
                            <el-checkbox-group v-model="formData.{{{ (toUpperCamelCase .GoField) }}}" placeholder="请选择{{{ .ColumnComment }}}">
                                {{{- if ne .DictType "" }}}
                                <el-checkbox
                                    v-for="(item, index) in dictData.{{{ .DictType }}}"
                                    :key="index"
                                    :label="item.name"
                                    :value="item.value"
                                    :disabled="!item.status"
                                ></el-checkbox>
                                {{{- else if ne .ListAllApi "" }}}
                                <el-checkbox
                                    v-for="(item, index) in listAllData.{{{pathToName .ListAllApi }}}"
                                    :key="index"
                                    :label="item.Id"
                                    :value="item.Id"
                                ></el-checkbox>
                                {{{- else }}}
                                <el-checkbox>请选择字典生成</el-checkbox>
                                {{{- end }}}
                            </el-checkbox-group>
                        </el-form-item>
                    {{{- else if eq .HtmlType "select" }}}
                        <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ (toUpperCamelCase .GoField) }}}">
                            <el-select class="flex-1" v-model="formData.{{{ (toUpperCamelCase .GoField) }}}" placeholder="请选择{{{ .ColumnComment }}}">
                                {{{- if ne .DictType "" }}}
                                <el-option
                                    v-for="(item, index) in dictData.{{{ .DictType }}}"
                                    :key="index"
                                    :label="item.name"
                                    {{{- if eq .GoType "int" }}}
                                    :value="parseInt(item.value)"
                                    {{{- else }}}
                                    :value="item.value"
                                    {{{- end }}}
                                    clearable
                                    :disabled="!item.status"
                                />
                                 {{{- else if ne .ListAllApi "" }}}
                                 <el-option
                                    v-for="(item, index) in listAllData.{{{pathToName .ListAllApi }}}"
                                    :key="index"
                                    :label="item.Id"
                                    {{{- if eq .GoType "int" }}}
                                    :value="parseInt(item.Id)"
                                    {{{- else }}}
                                    :value="String(item.Id)"
                                    {{{- end }}}
                                    clearable
                                />
                                {{{- else }}}
                                <el-option label="请选择字典生成" value="" />
                                {{{- end }}}
                            </el-select>
                        </el-form-item>
                    {{{- else if eq .HtmlType "radio" }}}
                        <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ (toUpperCamelCase .GoField) }}}">
                            <el-radio-group v-model="formData.{{{ (toUpperCamelCase .GoField) }}}" placeholder="请选择{{{ .ColumnComment }}}">
                                {{{- if ne .DictType "" }}}
                                <el-radio
                                    v-for="(item, index) in dictData.{{{ .DictType }}}"
                                    :key="index"
                                    :label="item.name"
                                    {{{- if eq .GoType "int" }}}
                                    :value="parseInt(item.value)"
                                    {{{- else }}}
                                    :value="item.value"
                                    {{{- end }}}
                                    :disabled="!item.status"
                                ></el-radio>
                                {{{- else if ne .ListAllApi "" }}}
                                <el-radio
                                    v-for="(item, index) in listAllData.{{{ pathToName .ListAllApi }}}"
                                    :key="index"
                                    :label="item.name"
                                    {{{- if eq .GoType "int" }}}
                                    :value="parseInt(item.Id)"
                                    {{{- else }}}
                                    :value="item.Id"
                                    {{{- end }}}
                                >
                                    {{ item.Id }}
                                </el-radio>
                                {{{- else }}}
                                <el-radio label="0">请选择字典生成</el-radio>
                                {{{- end }}}
                            </el-radio-group>
                        </el-form-item>
                    {{{- else if eq .HtmlType "datetime" }}}
                        <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ (toUpperCamelCase .GoField) }}}">
                            <el-date-picker
                                class="flex-1 !flex"
                                v-model="formData.{{{ (toUpperCamelCase .GoField) }}}"
                                type="datetime"
                                clearable
                                value-format="YYYY-MM-DD hh:mm:ss"
                                placeholder="请选择{{{ .ColumnComment }}}"
                            />
                        </el-form-item>
                    {{{- else if eq .HtmlType "editor" }}}
                        <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ (toUpperCamelCase .GoField) }}}">
                            <editor v-model="formData.{{{ (toUpperCamelCase .GoField) }}}" :height="500" />
                        </el-form-item>
                    {{{- else if eq .HtmlType "imageUpload" }}}
                        <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ (toUpperCamelCase .GoField) }}}">
                            <material-picker v-model="formData.{{{ (toUpperCamelCase .GoField) }}}" />
                        </el-form-item>
                    {{{- end }}}
                {{{- end }}}
                 {{{- end }}}
            {{{- end }}}
            
            </el-form>
        </popup>
    </div>
</template>
<script lang="ts" setup>
import type { FormInstance } from 'element-plus'
import { {{{ if and .Table.TreePrimary .Table.TreeParent }}}{{{ .ModuleName }}}_list_all,{{{ end }}} {{{ .ModuleName }}}_edit, {{{ .ModuleName }}}_add, {{{ .ModuleName }}}_detail } from '@/api/{{{nameToPath .ModuleName }}}'
import Popup from '@/components/popup/index.vue'
import feedback from '@/utils/feedback'
import { ref, shallowRef, computed, reactive } from 'vue'
import type { PropType } from 'vue'
defineProps({
    dictData: {
        type: Object as PropType<Record<string, any[]>>,
        default: () => ({})
    },
    listAllData:{
        type: Object as PropType<Record<string, any[]>>,
        default: () => ({})
    }
})
const emit = defineEmits(['success', 'close'])
const formRef = shallowRef<FormInstance>()
const popupRef = shallowRef<InstanceType<typeof Popup>>()
{{{- if and .Table.TreePrimary .Table.TreeParent }}}
const treeList = ref<any[]>([])
{{{- end }}}
const mode = ref('add')
const popupTitle = computed(() => {
    return mode.value == 'edit' ? '编辑{{{ .FunctionName }}}' : '新增{{{ .FunctionName }}}'
})

const formData = reactive({
    {{{- range .Columns }}}
    {{{- if eq (toUpperCamelCase .GoField) $.PrimaryKey }}}
    {{{ $.PrimaryKey }}}: '',
    {{{- else if .IsEdit }}}
    {{{- if eq .HtmlType "checkbox" }}}
    {{{ (toUpperCamelCase .GoField) }}}: [],
    {{{- else if eq .HtmlType "number" }}}
    {{{ (toUpperCamelCase .GoField) }}}: null,
    {{{- else }}}
    {{{ (toUpperCamelCase .GoField) }}}: null,
    {{{- end }}}
    {{{- end }}}
    {{{- end }}}
})

const formRules = {
    {{{- range .Columns }}}
    {{{- if and .IsEdit }}}
    {{{ (toUpperCamelCase .GoField) }}}: [
        {
            required: {{{- if eq .IsRequired 1}}} true {{{- else}}} false {{{- end }}},
            {{{- if or (eq .HtmlType "checkbox") (eq .HtmlType "datetime") (eq .HtmlType "radio") (eq .HtmlType "select") (eq .HtmlType "imageUpload") }}}
            message: '请选择{{{ .ColumnComment }}}',
            {{{- else }}}
            message: '请输入{{{ .ColumnComment }}}',
            {{{- end }}}
            trigger: ['blur']
        }
    ],
    {{{- end }}}
    {{{- end }}}
}

const handleSubmit = async () => {
     try {
        await formRef.value?.validate()
        const data: any = { ...formData }
        {{{- range .Columns }}}
        {{{- if eq .HtmlType "checkbox" }}}
        data.{{{ (toUpperCamelCase .GoField) }}} = data.{{{ (toUpperCamelCase .GoField) }}}.join(',')
        {{{- end }}}
        {{{- end }}}
        mode.value == 'edit' ? await {{{ .ModuleName }}}_edit(data) : await {{{ .ModuleName }}}_add(data)
        popupRef.value?.close()
        feedback.msgSuccess('操作成功')
        emit('success')
     } catch (error) {}
}

const open = (type = 'add') => {
    mode.value = type
    popupRef.value?.open()
}

const setFormData = async (data: Record<string, any>) => {
    for (const key in formData) {
        if (data[key] != null && data[key] != undefined) {
            //@ts-ignore
            formData[key] = data[key]
            {{{- range .Columns }}}
            {{{- if eq .HtmlType "checkbox" }}}
            //@ts-ignore
            formData.{{{ (toUpperCamelCase .GoField) }}} = String(data.{{{ (toUpperCamelCase .GoField) }}}).split(',')
            {{{- end }}}
            {{{- end }}}
        }
    }
}

const getDetail = async (row: Record<string, any>) => {
     try {
        const data = await {{{ .ModuleName }}}_detail(row.{{{toUpperCamelCase .PrimaryKey }}})
        setFormData(data)
     } catch (error) {}
}

const handleClose = () => {
    emit('close')
}
{{{- if and .Table.TreePrimary .Table.TreeParent }}}

const getLists = async () => {
    const data: any = await {{{ .ModuleName }}}_list_all()
    const item = { {{{ .Table.TreePrimary }}}: 0, {{{ .Table.TreeName }}}: '顶级', children: [] }
    item.children = data
    treeList.value.push(item)
}

getLists()
{{{- end }}}

defineExpose({
    open,
    setFormData,
    getDetail
})
</script>
