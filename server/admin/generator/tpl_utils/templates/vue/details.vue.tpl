<template>
    <div class="edit-popup">
        <popup
            ref="popupRef"
            :title="popupTitle"
            :async="true"
            width="550px"
            :clickModalClose="true"
            :confirmButtonText="false"
  
            @close="handleClose"
        >
            <el-form ref="formRef" :model="formData" label-width="84px" :rules="formRules">
            {{{- range .Columns }}}
                {{{- if .IsEdit }}}
                {{{- if ne (toUpperCamelCase .GoField) "Id" }}}
                    {{{- if and (ne $.Table.TreeParent "") (eq (toUpperCamelCase .GoField) $.Table.TreeParent) }}}
                        <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ (toUpperCamelCase .GoField) }}}">
                            <el-tree-select
                                class="flex-1"
                                v-model="formData.{{{ (toUpperCamelCase .GoField) }}}"
                                :data="treeList"
                                clearable
                                node-key="{{{ .Table.TreePrimary }}}"
                                :props="{ label: '{{{ .Table.TreeName }}}', value: '{{{ .Table.TreePrimary }}}', children: 'children' }"
                                :default-expand-all="true"
                                placeholder="请选择{{{ .ColumnComment }}}"
                                check-strictly
                                disabled
                            />
                        </el-form-item>
                    {{{- else if eq .HtmlType "input" }}}
                        <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ (toUpperCamelCase .GoField) }}}">
                            <span v-text="formData.{{{ (toUpperCamelCase .GoField) }}}"></span>
                        </el-form-item>
                    {{{- else if eq .HtmlType "number" }}}
                        <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ (toUpperCamelCase .GoField) }}}">
                            <span v-text="formData.{{{ (toUpperCamelCase .GoField) }}}"></span>
                        </el-form-item>
                    {{{- else if eq .HtmlType "textarea" }}}
                        <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ (toUpperCamelCase .GoField) }}}">
                            <span v-text="formData.{{{ (toUpperCamelCase .GoField) }}}"></span>
                        </el-form-item>
                    {{{- else if eq .HtmlType "checkbox" }}}
                        <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ (toUpperCamelCase .GoField) }}}">
                      
                      
                                {{{- if ne .DictType "" }}}
                                <dict-value :options="dictData.{{{ .DictType }}}" :value="formData.{{{ (toUpperCamelCase .GoField) }}}" />
                                {{!-- <el-checkbox
                                    v-for="(item, index) in dictData.{{{ .DictType }}}"
                                    :key="index"
                                    :label="item.name"
                                    :value="item.value"
                                    :disabled="!item.status"
                                ></el-checkbox> --}}
                                {{{- else if ne .ListAllApi "" }}}
                                <dict-value :options="listAllData.{{{pathToName .ListAllApi }}}" :value="formData.{{{ (toUpperCamelCase .GoField) }}}" />
                                {{!-- <el-checkbox
                                    v-for="(item, index) in listAllData.{{{pathToName .ListAllApi }}}"
                                    :key="index"
                                    :label="item.Id"
                                    :value="item.Id"
                                ></el-checkbox> --}}
                       
                                {{!-- <el-checkbox>请选择字典生成</el-checkbox> --}}
                                {{{- end }}}
                    
                        </el-form-item>
                    {{{- else if eq .HtmlType "select" }}}
                        <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ (toUpperCamelCase .GoField) }}}">
                            {{!-- <el-select class="flex-1" v-model="formData.{{{ (toUpperCamelCase .GoField) }}}" placeholder="请选择{{{ .ColumnComment }}}"> --}}
                                {{{- if ne .DictType "" }}}
                                <dict-value :options="dictData.{{{ .DictType }}}" :value="formData.{{{ (toUpperCamelCase .GoField) }}}" />
                                {{!-- <el-option
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
                                /> --}}
                                 {{{- else if ne .ListAllApi "" }}}
                                 <dict-value :options="listAllData.{{{pathToName .ListAllApi }}}" :value="formData.{{{ (toUpperCamelCase .GoField) }}}" />
                                 {{!-- <el-option
                                    v-for="(item, index) in listAllData.{{{pathToName .ListAllApi }}}"
                                    :key="index"
                                    :label="item.Id"
                                    {{{- if eq .GoType "int" }}}
                                    :value="parseInt(item.Id)"
                                    {{{- else }}}
                                    :value="String(item.Id)"
                                    {{{- end }}}
                                    clearable
                                /> --}}
                             
                                {{!-- <el-option label="请选择字典生成" value="" /> --}}
                                {{{- end }}}
                            {{!-- </el-select> --}}
                        </el-form-item>
                    {{{- else if eq .HtmlType "radio" }}}
                        <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ (toUpperCamelCase .GoField) }}}">
                            {{!-- <el-radio-group v-model="formData.{{{ (toUpperCamelCase .GoField) }}}" placeholder="请选择{{{ .ColumnComment }}}"> --}}
                                {{{- if ne .DictType "" }}}
                                <dict-value :options="dictData.{{{ .DictType }}}" :value="formData.{{{ (toUpperCamelCase .GoField) }}}" />
                                {{!-- <el-radio
                                    v-for="(item, index) in dictData.{{{ .DictType }}}"
                                    :key="index"
                                    :label="item.name"
                                    {{{- if eq .GoType "int" }}}
                                    :value="parseInt(item.value)"
                                    {{{- else }}}
                                    :value="item.value"
                                    {{{- end }}}
                                    :disabled="!item.status"
                                ></el-radio> --}}
                                {{{- else if ne .ListAllApi "" }}}
                                <dict-value :options="listAllData.{{{pathToName .ListAllApi }}}" :value="formData.{{{ (toUpperCamelCase .GoField) }}}" />
                                {{!-- <el-radio
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
                                </el-radio> --}}
                      
                                {{!-- <el-radio label="0">请选择字典生成</el-radio> --}}
                                {{{- end }}}
                            {{!-- </el-radio-group> --}}
                        </el-form-item>
                    {{{- else if eq .HtmlType "datetime" }}}
                        <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ (toUpperCamelCase .GoField) }}}">
                            <span v-text="formData.{{{ (toUpperCamelCase .GoField) }}}"></span>
                            {{!-- <el-date-picker
                                class="flex-1 !flex"
                                v-model="formData.{{{ (toUpperCamelCase .GoField) }}}"
                                type="datetime"
                                clearable
                                value-format="YYYY-MM-DD hh:mm:ss"
                                placeholder="请选择{{{ .ColumnComment }}}"
                            /> --}}
                        </el-form-item>
                    {{{- else if eq .HtmlType "editor" }}}
                        <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ (toUpperCamelCase .GoField) }}}">
                            <div v-html="formData.{{{ (toUpperCamelCase .GoField) }}}"></div>
                        </el-form-item>
                    {{{- else if eq .HtmlType "imageUpload" }}}
                        <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ (toUpperCamelCase .GoField) }}}">
                            <image-contain
                                :width="40"
                                :height="40"
                                :src="formData.{{{ (toUpperCamelCase .GoField) }}}"
                                :preview-src-list="[formData.{{{ (toUpperCamelCase .GoField) }}}]"
                                preview-teleported
                                hide-on-click-modal
                            />
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
import { {{{ if and .Table.TreePrimary .Table.TreeParent }}}{{{ .ModuleName }}}_list_all,{{{ end }}} ,{{{ .ModuleName }}}_detail } from '@/api/{{{nameToPath .ModuleName }}}'
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

const popupTitle = computed(() => {
    return '预览{{{ .FunctionName }}}'
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

const open = () => {
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
