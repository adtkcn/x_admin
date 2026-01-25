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
            <el-form ref="formRef" :model="formData" label-width="110px" :rules="formRules">
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
                        check-strictly
                        disabled
                    />
                </el-form-item>
            {{{- end }}}
            {{{- range .Columns }}}
                {{{- if .IsEdit }}}
                {{{- if not .IsPk }}}               
                    {{{- if eq .HtmlType "input" }}}
                        <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ .TsField }}}">
                            <span v-text="formData.{{{ .TsField }}}"></span>
                        </el-form-item>
                    {{{- else if eq .HtmlType "number" }}}
                        <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ .TsField }}}">
                            <span v-text="formData.{{{ .TsField }}}"></span>
                        </el-form-item>
                    {{{- else if eq .HtmlType "textarea" }}}
                        <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ .TsField }}}">
                            <span v-text="formData.{{{ .TsField }}}"></span>
                        </el-form-item>
                    {{{- else if eq .HtmlType "checkbox" }}}
                        <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ .TsField }}}">
                                {{{- if ne .DictType "" }}}
                                <dict-value :options="dictData.{{{ .DictType }}}" :value="formData.{{{ .TsField }}}" />

                                {{{- else if ne .ListAllApi "" }}}
                                <dict-value :options="listAllData.{{{pathToName .ListAllApi }}}" :value="formData.{{{ .TsField }}}" />
                                {{{- end }}}
                    
                        </el-form-item>
                    {{{- else if eq .HtmlType "select" }}}
                        <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ .TsField }}}"> 
                                {{{- if ne .DictType "" }}}
                                <dict-value :options="dictData.{{{ .DictType }}}" :value="formData.{{{ .TsField }}}" />
         
                                 {{{- else if ne .ListAllApi "" }}}
                                 <dict-value :options="listAllData.{{{pathToName .ListAllApi }}}" :value="formData.{{{ .TsField }}}" />
                                {{{- end }}} 
                        </el-form-item>
                    {{{- else if eq .HtmlType "radio" }}}
                        <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ .TsField }}}">
                                {{{- if ne .DictType "" }}}
                                <dict-value :options="dictData.{{{ .DictType }}}" :value="formData.{{{ .TsField }}}" />
                                {{{- else if ne .ListAllApi "" }}}
                                <dict-value :options="listAllData.{{{pathToName .ListAllApi }}}" :value="formData.{{{ .TsField }}}" />
                                {{{- end }}}
                        </el-form-item>
                    {{{- else if eq .HtmlType "datetime" }}}
                        <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ .TsField }}}">
                            <span v-text="formData.{{{ .TsField }}}"></span>
                        </el-form-item>
                    {{{- else if eq .HtmlType "editor" }}}
                        <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ .TsField }}}">
                            <div v-html="formData.{{{ .TsField }}}"></div>
                        </el-form-item>
                    {{{- else if eq .HtmlType "imageUpload" }}}
                        <el-form-item label="{{{ .ColumnComment }}}" prop="{{{ .TsField }}}">
                            <image-contain
                                :width="40"
                                :height="40"
                                :src="formData.{{{ .TsField }}}"
                                :preview-src-list="[formData.{{{ .TsField }}}]"
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
import { {{{ if and .Table.TreePrimary .Table.TreeParent }}}{{{ .ModuleName }}}_list_all,{{{ end }}} {{{ .ModuleName }}}_detail } from '@/api/{{{nameToPath .ModuleName }}}'
import Popup from '@/components/popup/index.vue'
 
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
const emit = defineEmits(['close'])
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
    {{{- if eq .TsField $.PrimaryKey }}}
    {{{ $.PrimaryKey }}}: '',
    {{{- else if .IsEdit }}}
    {{{- if eq .HtmlType "checkbox" }}}
    {{{ .TsField }}}: [],
    {{{- else if eq .HtmlType "number" }}}
    {{{ .TsField }}}: null,
    {{{- else }}}
    {{{ .TsField }}}: null,
    {{{- end }}}
    {{{- end }}}
    {{{- end }}}
})

const formRules = {
    {{{- range .Columns }}}
    {{{- if and .IsEdit }}}
    {{{ .TsField }}}: [
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
const getDetail = async (row: Record<string, any>) => {
     try {
        const data = await {{{ .ModuleName }}}_detail(row.{{{toUpperCamelCase .PrimaryKey }}})
        setFormData(data)
     } catch (error) {}
}
const setFormData = async (data: Record<string, any>) => {
    for (const key in formData) {
        if (data[key] != null && data[key] != undefined) {
            //@ts-ignore
            formData[key] = data[key]
            {{{- range .Columns }}}
            {{{- if eq .HtmlType "checkbox" }}}
            //@ts-ignore
            formData.{{{ .TsField }}} = String(data.{{{ .TsField }}}).split(',')
            {{{- end }}}
            {{{- end }}}
        }
    }
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
 