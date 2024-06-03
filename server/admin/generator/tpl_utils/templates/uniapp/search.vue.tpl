<template>
	<view class="page-content">
		<uv-form labelPosition="left" labelWidth="80" :model="form"  ref="formRef">
			{{{- range .Columns }}}
			{{{- if eq .IsQuery 1 }}}
			<uv-form-item label="{{{ .ColumnComment }}}" prop="{{{ (toCamelCase .GoField) }}}" borderBottom>
				{{{- if eq .HtmlType "datetime" }}}
					<x-date-range v-model:startTime="form.{{{ (toCamelCase .GoField) }}}Start"
							v-model:endTime="form.{{{ (toCamelCase .GoField) }}}End"></x-date-range>
				{{{- else if or (eq .HtmlType "select") (eq .HtmlType "radio") }}}
				{{{- if ne .DictType "" }}}
					<x-picker v-model="form.{{{ (toCamelCase .GoField) }}}" valueKey="value" labelKey="name" :columns="dictData.{{{ .DictType }}}"></x-picker>
				{{{- end }}}
				{{{- else if eq .HtmlType "input" }}}
					<uv-input v-model="form.{{{ (toCamelCase .GoField) }}}"> </uv-input>
				{{{- end }}}
			</uv-form-item>
			{{{- end }}}
			{{{- end }}}

			<uv-button type="primary" text="搜索" customStyle="margin-top: 20rpx" @click="submit"></uv-button>
		</uv-form>
	</view>
</template>

<script setup>
 
	import {
		onLoad
	} from "@dcloudio/uni-app";
	import {
		reactive,
		ref,
		computed
	} from "vue";
	import {
		toPath,
		toast,
		clearObjEmpty
	} from "@/utils/utils";
	import {
		useDictData
	} from "@/hooks/useDictOptions";
	
{{{- if ge (len .DictFields) 1 }}}
{{{- $dictSize := sub (len .DictFields) 1 }}}
const { dictData } = useDictData([{{{- range .DictFields }}}'{{{ . }}}'{{{- if ne (index $.DictFields $dictSize) . }}},{{{- end }}}{{{- end }}}])
{{{- end }}}

	let formRef = ref();
	let form = ref({
{{{- range .Columns }}}
{{{- if .IsQuery }}}
    {{{- if eq .HtmlType "datetime" }}}
    {{{ (toCamelCase .GoField) }}}Start: '',
    {{{ (toCamelCase .GoField) }}}End: '',
    {{{- else }}}
    {{{ (toCamelCase .GoField) }}}: '',
    {{{- end }}}
{{{- end }}}
{{{- end }}}
	});

	function submit() {
		console.log("submit", form.value);

		const search = clearObjEmpty(form.value);

		if (Object.keys(search).length === 0) {
			return toast("请输入查询条件");
		}

		toPath("/pages/{{{ .ModuleName }}}/index", search);
	}
</script>

<style lang="scss" scoped>
	.page-content {
		padding: 10rpx 20rpx 300rpx;
	}
</style>