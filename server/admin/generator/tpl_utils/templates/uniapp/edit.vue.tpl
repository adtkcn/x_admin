<template>
	<view class="page-content">
		<uv-form labelPosition="left" :model="form" :rules="formRules" ref="formRef">
			{{{- range .Columns }}}
            {{{- if .IsEdit }}}			
			<uv-form-item label="{{{ .ColumnComment }}}" prop="{{{ (toUpperCamelCase .GoField) }}}" borderBottom>
                {{{- if eq .HtmlType "input" }}}
					<uv-input v-model="form.{{{ (toUpperCamelCase .GoField) }}}" border="surround"></uv-input>
				{{{- else if eq .HtmlType "number" }}}
					<uv-number-box v-model="form.{{{ (toUpperCamelCase .GoField) }}}" :min="-99999999" :max="99999999" :integer="true"></uv-number-box>
				{{{- else if eq .HtmlType "textarea" }}}
					<uv-textarea v-model="form.{{{ (toUpperCamelCase .GoField) }}}" border="surround"></uv-textarea>
				{{{- else if eq .HtmlType "datetime" }}}
					<x-date v-model:time="form.{{{ (toUpperCamelCase .GoField) }}}"></x-date>
				{{{- else if or (eq .HtmlType "checkbox") (eq .HtmlType "radio") (eq .HtmlType "select")}}}
					{{{- if ne .DictType "" }}}
						<x-picker v-model="form.{{{ (toUpperCamelCase .GoField) }}}" valueKey="value" labelKey="name" :columns="dictData.{{{ .DictType }}}"></x-picker>
					{{{- else if ne .ListAllApi "" }}}
						<x-picker v-model="form.{{{ (toUpperCamelCase .GoField) }}}" valueKey="{{{toUpperCamelCase .PrimaryKey}}}" labelKey="{{{toUpperCamelCase .PrimaryKey}}}" :columns="listAllData.{{{pathToName .ListAllApi}}}"></x-picker>
					{{{- else }}}
						请选择字典生成代码
					{{{- end }}}
				{{{- end }}}
			</uv-form-item>
			{{{- end }}}
			{{{- end }}}

			<uv-button   type="primary" text="提交" customStyle="margin: 40rpx 0"
				@click="submit"></uv-button>
		</uv-form>
	</view>
</template>

<script setup lang="ts">
	import { ref } from "vue";
	import {
		onLoad
	} from "@dcloudio/uni-app";
 	import {
		{{{ .ModuleName }}}_detail,
		{{{ .ModuleName }}}_edit,
		{{{ .ModuleName }}}_add
	} from "@/api/{{{ .ModuleName }}}";
	import type { type_{{{ .ModuleName }}}_edit	} from "@/api/{{{nameToPath .ModuleName }}}";

	import {
		toast,
		alert
	} from "@/utils/utils";
	import {
		useDictData,useListAllData
	} from "@/hooks/useDictOptions";
	import type { type_dict } from '@/hooks/useDictOptions'

	let formRef = ref();
	let form = ref<type_{{{ .ModuleName }}}_edit>({
	{{{- range .Columns }}}
    {{{- if eq (toUpperCamelCase .GoField) $.PrimaryKey }}}
    {{{ $.PrimaryKey }}}: '',
    {{{- else if .IsEdit }}}
    {{{- if eq .HtmlType "checkbox" }}}
    {{{ (toUpperCamelCase .GoField) }}}: [],
    {{{- else if eq .HtmlType "number" }}}
    {{{ (toUpperCamelCase .GoField) }}}: 0,
    {{{- else }}}
    {{{ (toUpperCamelCase .GoField) }}}: '',
    {{{- end }}}
    {{{- end }}}
    {{{- end }}}
	});
	const formRules = {
		{{{- range .Columns }}}
		{{{- if and .IsEdit .IsRequired }}}
		{{{ (toUpperCamelCase .GoField) }}}: [
			{
				required: true,
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
	onLoad((e) => {
		console.log("onLoad", e);
		if (e.{{{toUpperCamelCase .PrimaryKey}}}) {
			getDetails(e.{{{toUpperCamelCase .PrimaryKey}}});
		}
	});


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

	function getDetails({{{toUpperCamelCase .PrimaryKey}}}) {
		{{{ .ModuleName }}}_detail({{{toUpperCamelCase .PrimaryKey}}}).then((res) => {
            if (res.code == 200) {
                if (res?.data) {
                    form.value = res?.data
                }
            } else {
                toast(res.message);
            }
        })
        .catch((err) => {
            toast(err.message||"网络错误");
        });
	}

	function submit() {
		console.log("submit", form.value);
		formRef.value.validate().then(() => {
			if (form.value.{{{toUpperCamelCase .PrimaryKey}}}) {
				{{{ .ModuleName }}}_edit(form.value).then((res) => {
					if (res.code == 200) {
						toast("编辑成功");				
						getDetails(form.value?.{{{toUpperCamelCase .PrimaryKey}}});
					} else {
						toast(res.message);
					}
				});
			}else{
				{{{ .ModuleName }}}_add(form.value).then((res) => {
					if (res.code == 200) {
						toast("添加成功");
						uni.navigateBack();
					} else {
						toast(res.message);
					}
				});
			}

		})
	}
</script>

<style lang="scss" scoped>
	.page-content {
		padding: 10rpx 20rpx 300rpx;
	}
</style>